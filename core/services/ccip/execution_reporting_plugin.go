package ccip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

const (
	ExecutionMaxInflightTimeSeconds = 180
)

var _ types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
var _ types.ReportingPlugin = &ExecutionReportingPlugin{}

type Message struct {
	SequenceNumber     *big.Int       `json:"sequenceNumber"`
	SourceChainId      *big.Int       `json:"sourceChainId"`
	DestinationChainId *big.Int       `json:"destinationChainId"`
	Sender             common.Address `json:"sender"`
	Payload            struct {
		Receiver common.Address   `json:"receiver"`
		Data     []uint8          `json:"data"`
		Tokens   []common.Address `json:"tokens"`
		Amounts  []*big.Int       `json:"amounts"`
		Executor common.Address   `json:"executor"`
		Options  []uint8          `json:"options"`
	} `json:"payload"`
}

type ExecutableMessage struct {
	Proof   [][32]byte `json:"proof"`
	Message Message    `json:"message"`
	Index   *big.Int   `json:"index"`
}

type ExecutableMessages []ExecutableMessage

func (ems ExecutableMessages) SeqNums() (nums []*big.Int) {
	for i := range ems {
		nums = append(nums, ems[i].Message.SequenceNumber)
	}
	return
}

// ExecutionObservation Note there can be gaps in this range of sequence numbers,
// indicative of some messages being non-DON executed.
type ExecutionObservation struct {
	MinSeqNum utils.Big `json:"minSeqNum"`
	MaxSeqNum utils.Big `json:"maxSeqNum"`
}

func makeExecutionReportArgs() abi.Arguments {
	mustType := func(ts string, components []abi.ArgumentMarshaling) abi.Type {
		ty, _ := abi.NewType(ts, "", components)
		return ty
	}
	return []abi.Argument{
		{
			Name: "executableMessages",
			Type: mustType("tuple[]", []abi.ArgumentMarshaling{
				{
					Name: "Proof",
					Type: "bytes32[]",
				},
				{
					Name: "Message",
					Type: "tuple",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "sequenceNumber",
							Type: "uint256",
						},
						{
							Name: "sourceChainId",
							Type: "uint256",
						},
						{
							Name: "destinationChainId",
							Type: "uint256",
						},
						{
							Name: "sender",
							Type: "address",
						},
						{
							Name: "payload",
							Type: "tuple",
							Components: []abi.ArgumentMarshaling{
								{
									Name: "receiver",
									Type: "address",
								},
								{
									Name: "data",
									Type: "bytes",
								},
								{
									Name: "tokens",
									Type: "address[]",
								},
								{
									Name: "amounts",
									Type: "uint256[]",
								},
								{
									Name: "executor",
									Type: "address",
								},
								{
									Name: "options",
									Type: "bytes",
								},
							},
						},
					},
				},
				{
					Name: "Index",
					Type: "uint256",
				},
			}),
		},
	}
}

func EncodeExecutionReport(ems []ExecutableMessage) (types.Report, error) {
	report, err := makeExecutionReportArgs().PackValues([]interface{}{ems})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func DecodeExecutionReport(report types.Report) ([]ExecutableMessage, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, nil
	}

	// Must be anonymous struct here
	msgs, ok := unpacked[0].([]struct {
		Proof   [][32]uint8 `json:"Proof"`
		Message struct {
			SequenceNumber     *big.Int       `json:"sequenceNumber"`
			SourceChainId      *big.Int       `json:"sourceChainId"`
			DestinationChainId *big.Int       `json:"destinationChainId"`
			Sender             common.Address `json:"sender"`
			Payload            struct {
				Receiver common.Address   `json:"receiver"`
				Data     []uint8          `json:"data"`
				Tokens   []common.Address `json:"tokens"`
				Amounts  []*big.Int       `json:"amounts"`
				Executor common.Address   `json:"executor"`
				Options  []uint8          `json:"options"`
			} `json:"payload"`
		} `json:"Message"`
		Index *big.Int `json:"Index"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var ems []ExecutableMessage
	for _, emi := range msgs {
		ems = append(ems, ExecutableMessage{
			Proof:   emi.Proof,
			Message: emi.Message,
			Index:   emi.Index,
		})
	}
	return ems, nil
}

//go:generate mockery --name OffRampLastReporter --output ./mocks/lastreporter --case=underscore
type OffRampLastReporter interface {
	GetLastReport(opts *bind.CallOpts) (single_token_offramp.CCIPRelayReport, error)
}

type ExecutionReportingPluginFactory struct {
	l            logger.Logger
	orm          ORM
	source, dest *big.Int
	lastReporter OffRampLastReporter
	executor     common.Address
}

func NewExecutionReportingPluginFactory(l logger.Logger, orm ORM, source, dest *big.Int, executor common.Address, lastReporter OffRampLastReporter) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{l: l, orm: orm, source: source, dest: dest, executor: executor, lastReporter: lastReporter}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	return ExecutionReportingPlugin{rf.l, config.F, rf.orm, rf.source, rf.dest, rf.executor, rf.lastReporter}, types.ReportingPluginInfo{
		Name:              "CCIPExecution",
		UniqueReports:     true,
		MaxQueryLen:       0,      // We do not use the query phase.
		MaxObservationLen: 100000, // TODO
		MaxReportLen:      100000, // TODO
	}, nil
}

type ExecutionReportingPlugin struct {
	l             logger.Logger
	F             int
	orm           ORM
	sourceChainId *big.Int
	destChainId   *big.Int
	executor      common.Address
	// We also use the offramp for defensive checks
	lastReporter OffRampLastReporter
}

func (r ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	// We want to execute any messages which satisfy the following:
	// 1. Have the executor field set to the DONs message executor contract
	// 2. There exists a confirmed relay report containing its sequence number, i.e. it's status is RequestStatusRelayConfirmed
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, nil, nil, RequestStatusRelayConfirmed, &r.executor, nil)
	if err != nil {
		return nil, err
	}
	// No request to process
	// Return an empty observation
	// which should not result in a report generated.
	if len(reqs) == 0 {
		return nil, fmt.Errorf("no requests for oracle execution")
	}
	// Double check the latest sequence number onchain is >= our max relayed seq num
	lr, err := r.lastReporter.GetLastReport(nil)
	if err != nil {
		return nil, err
	}
	if reqs[len(reqs)-1].SeqNum.ToInt().Cmp(lr.MaxSequenceNumber) > 0 {
		return nil, fmt.Errorf("invariant violated, mismatch between relay_confirmed requests and last report")
	}
	b, err := json.Marshal(&ExecutionObservation{
		MinSeqNum: reqs[0].SeqNum,
		MaxSeqNum: reqs[len(reqs)-1].SeqNum,
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	var nonEmptyObservations []ExecutionObservation
	for _, ao := range observations {
		var ob ExecutionObservation
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			r.l.Errorw("unmarshallable observation", "ao", ao.Observation, "err", err)
			continue
		}
		nonEmptyObservations = append(nonEmptyObservations, ob)
	}
	// Need at least F+1 observations
	if len(nonEmptyObservations) <= r.F {
		return false, nil, nil
	}
	// We have at least F+1 valid observations
	// Extract the min and max
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MinSeqNum.ToInt().Cmp(nonEmptyObservations[j].MinSeqNum.ToInt()) < 0
	})
	min := nonEmptyObservations[r.F].MinSeqNum.ToInt()
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MaxSeqNum.ToInt().Cmp(nonEmptyObservations[j].MaxSeqNum.ToInt()) < 0
	})
	max := nonEmptyObservations[r.F].MaxSeqNum.ToInt()
	if max.Cmp(min) < 0 {
		return false, nil, errors.New("max seq num smaller than min")
	}
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, max, RequestStatusRelayConfirmed, &r.executor, nil)
	if err != nil {
		return false, nil, err
	}
	// Cannot construct a report for which we haven't seen all the messages.
	if len(reqs) == 0 {
		return false, nil, fmt.Errorf("do not have all the messages in report, have zero messages, report has min %v max %v", min, max)
	}
	lr, err := r.lastReporter.GetLastReport(nil)
	if err != nil {
		return false, nil, err
	}
	if reqs[len(reqs)-1].SeqNum.ToInt().Cmp(lr.MaxSequenceNumber) > 0 {
		return false, nil, fmt.Errorf("invariant violated, mismatch between relay_confirmed requests (max %v) and last report (max %v)", reqs[len(reqs)-1].SeqNum, lr.MaxSequenceNumber)
	}
	report, err := r.buildReport(reqs)
	if err != nil {
		return false, nil, err
	}
	return true, report, nil
}

// For each message in the given range of sequence numbers (with potential holes):
// 1. Lookup the report associated with that sequence number
// 2. Generate a merkle proof that the message was in that report
// 3. Encode those proofs and messages into a report for the executor contract
// TODO: We may want to combine these queries for performance, hold off
// until we decide whether we move forward with batch proving.
func (r ExecutionReportingPlugin) buildReport(reqs []*Request) ([]byte, error) {
	var executable []ExecutableMessage
	for _, req := range reqs {
		// Look up all the messages that are in the same report
		// as this one (even externally executed ones), generate a Proof and double-check the root checks out.
		rep, err2 := r.orm.RelayReport(req.SeqNum.ToInt())
		if err2 != nil {
			r.l.Errorw("could not find relay report for request", "err", err2, "seq num", req.SeqNum.String())
			continue
		}
		allReqsInReport, err3 := r.orm.Requests(r.sourceChainId, r.destChainId, rep.MinSeqNum.ToInt(), rep.MaxSeqNum.ToInt(), "", nil, nil)
		if err3 != nil {
			continue
		}
		var leaves [][]byte
		for _, reqInReport := range allReqsInReport {
			leaves = append(leaves, reqInReport.Raw)
		}
		index := big.NewInt(0).Sub(req.SeqNum.ToInt(), rep.MinSeqNum.ToInt())
		root, proof := GenerateMerkleProof(32, leaves, int(index.Int64()))
		if !bytes.Equal(root[:], rep.Root[:]) {
			continue
		}
		executable = append(executable, ExecutableMessage{
			Proof:   proof.PathForExecute(),
			Message: req.ToMessage(),
			Index:   proof.Index(),
		})
	}

	report, err := EncodeExecutionReport(executable)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	ems, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If the report is stale, we do not accept it.
	stale, err := r.isStale(ems[0].Message.SequenceNumber)
	if err != nil {
		return !stale, err
	}
	if stale {
		return false, err
	}
	// Any timed out requests should be set back to RequestStatusExecutionPending so their execution can be retried in a subsequent report.
	if err = r.orm.ResetExpiredRequests(r.sourceChainId, r.destChainId, ExecutionMaxInflightTimeSeconds, RequestStatusExecutionPending, RequestStatusRelayConfirmed); err != nil {
		// Ok to continue here, we'll try to reset them again on the next round.
		r.l.Errorw("unable to reset expired requests", "err", err)
	}
	if err := r.orm.UpdateRequestSetStatus(r.sourceChainId, r.destChainId, ExecutableMessages(ems).SeqNums(), RequestStatusExecutionPending); err != nil {
		return false, err
	}
	return true, nil
}

func (r ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for bptxm,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStale(parsedReport[0].Message.SequenceNumber)
	return !stale, err
}

func (r ExecutionReportingPlugin) isStale(min *big.Int) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	req, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, min, "", nil, nil)
	if err != nil {
		// if we can't find the request, assume transient db issue
		// and wait until the next OCR2 round (don't submit)
		return true, err
	}
	if len(req) != 1 {
		// If we don't have the request at all, this likely means we never had the request to begin with
		// (say our eth subscription is down) and we want to let other oracles continue the protocol.
		return false, errors.New("could not find first message in execution report")
	}
	return req[0].Status == RequestStatusExecutionConfirmed, nil
}

func (r ExecutionReportingPlugin) Start() error {
	return nil
}

func (r ExecutionReportingPlugin) Close() error {
	return nil
}
