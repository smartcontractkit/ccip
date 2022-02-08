package ccip

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

const (
	RelayMaxInflightTimeSeconds = 180
)

var _ types.ReportingPluginFactory = &RelayReportingPluginFactory{}
var _ types.ReportingPlugin = &RelayReportingPlugin{}

type Observation struct {
	MinSeqNum utils.Big `json:"minSeqNum"`
	MaxSeqNum utils.Big `json:"maxSeqNum"`
}

func makeRelayReportArgs() abi.Arguments {
	mustType := func(ts string) abi.Type {
		ty, _ := abi.NewType(ts, "", nil)
		return ty
	}
	return []abi.Argument{
		{
			Name: "merkleRoot",
			Type: mustType("bytes32"),
		},
		{
			Name: "minSequenceNumber",
			Type: mustType("uint256"),
		},
		{
			Name: "maxSequenceNumber",
			Type: mustType("uint256"),
		},
	}
}

func EncodeRelayReport(relayReport *single_token_offramp.CCIPRelayReport) (types.Report, error) {
	report, err := makeRelayReportArgs().PackValues([]interface{}{relayReport.MerkleRoot, relayReport.MinSequenceNumber, relayReport.MaxSequenceNumber})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func DecodeRelayReport(report types.Report) (*single_token_offramp.CCIPRelayReport, error) {
	unpacked, err := makeRelayReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 3 {
		return nil, errors.New("invalid num fields in report")
	}
	root, ok := unpacked[0].([32]byte)
	if !ok {
		return nil, errors.New("invalid root")
	}
	min, ok := unpacked[1].(*big.Int)
	if !ok {
		return nil, errors.New("invalid min")
	}
	max, ok := unpacked[2].(*big.Int)
	if !ok {
		return nil, errors.New("invalid max")
	}
	return &single_token_offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: min,
		MaxSequenceNumber: max,
	}, nil
}

type RelayReportingPluginFactory struct {
	l       logger.Logger
	orm     ORM
	jobID   int32
	offRamp *single_token_offramp.SingleTokenOffRamp
}

func NewRelayReportingPluginFactory(l logger.Logger, orm ORM, offRamp *single_token_offramp.SingleTokenOffRamp) types.ReportingPluginFactory {
	return &RelayReportingPluginFactory{l: l, orm: orm, offRamp: offRamp}
}

func (rf *RelayReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	destChainId, err := rf.offRamp.CHAINID(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	sourceChainId, err := rf.offRamp.SOURCECHAINID(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return RelayReportingPlugin{rf.l, config.F, rf.orm, sourceChainId, destChainId, rf.offRamp}, types.ReportingPluginInfo{
		Name:              "CCIPRelay",
		UniqueReports:     true,
		MaxQueryLen:       0,      // We do not use the query phase.
		MaxObservationLen: 100000, // TODO
		MaxReportLen:      100000, // TODO
	}, nil
}

type RelayReportingPlugin struct {
	l             logger.Logger
	F             int
	orm           ORM
	sourceChainId *big.Int
	destChainId   *big.Int
	offRamp       *single_token_offramp.SingleTokenOffRamp
}

func (r RelayReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r RelayReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lastReport, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		return nil, err
	}
	unstartedReqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, big.NewInt(0).Add(lastReport.MaxSequenceNumber, big.NewInt(1)), nil, RequestStatusUnstarted, nil, nil)
	if err != nil {
		return nil, err
	}
	// No request to process
	// Return an empty observation
	// which should not result in a report generated.
	if len(unstartedReqs) == 0 {
		return nil, fmt.Errorf("no requests with seq num greater than %v", lastReport.MaxSequenceNumber)
	}

	b, err := json.Marshal(&Observation{
		MinSeqNum: unstartedReqs[0].SeqNum,
		MaxSeqNum: unstartedReqs[len(unstartedReqs)-1].SeqNum,
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r RelayReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	// Need at least F+1 valid observations
	var nonEmptyObservations []Observation
	for _, ao := range observations {
		var ob Observation
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			r.l.Errorw("received unmarshallable observation", "err", err, "observation", string(ao.Observation))
			continue
		}
		nonEmptyObservations = append(nonEmptyObservations, ob)
	}
	if len(nonEmptyObservations) <= r.F {
		return false, nil, nil
	}
	// We have at least F+1 valid observations
	// Extract the min and max
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MinSeqNum.ToInt().Cmp(nonEmptyObservations[j].MinSeqNum.ToInt()) < 0
	})
	// r.F < len(nonEmptyObservations) because of the check above and therefore this is safe
	min := nonEmptyObservations[r.F].MinSeqNum.ToInt()
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MaxSeqNum.ToInt().Cmp(nonEmptyObservations[j].MaxSeqNum.ToInt()) < 0
	})
	max := nonEmptyObservations[r.F].MaxSeqNum.ToInt()
	if max.Cmp(min) < 0 {
		return false, nil, errors.New("max seq num smaller than min")
	}
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, nil, RequestStatusUnstarted, nil, nil)
	if err != nil {
		return false, nil, err
	}
	// Cannot construct a report for which we haven't seen all the messages.
	if len(reqs) == 0 {
		return false, nil, fmt.Errorf("do not have all the messages in report, have zero messages, report has min %v max %v", min, max)
	}
	if reqs[len(reqs)-1].SeqNum.ToInt().Cmp(max) < 0 {
		return false, nil, fmt.Errorf("do not have all the messages in report, our max %v reports max %v", reqs[len(reqs)-1].SeqNum, max)
	}
	if r.isStale(min) {
		return false, nil, nil
	}

	report, err := r.buildReport(min, max)
	if err != nil {
		return false, nil, err
	}
	return true, report, nil
}

func (r RelayReportingPlugin) isStale(minSeqNum *big.Int) bool {
	lastReport, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		// Assume its a transient issue getting the last report
		// Will try again on the next round
		return true
	}
	// If the last report onchain has a lower bound
	// strictly greater than this minSeqNum, then this minSeqNum
	// is stale.
	return lastReport.MinSequenceNumber.Cmp(minSeqNum) > 0
}

func (r RelayReportingPlugin) buildReport(min *big.Int, max *big.Int) ([]byte, error) {
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, max, "", nil, nil)
	if err != nil {
		return nil, err
	}
	// Take all these request and produce a merkle root of them
	var leaves [][]byte
	for _, req := range reqs {
		leaves = append(leaves, req.Raw)
	}

	// Note Index doesn't matter, we just want the root
	root, _ := GenerateMerkleProof(32, leaves, 0)
	report, err := EncodeRelayReport(&single_token_offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: min,
		MaxSequenceNumber: max,
	})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r RelayReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	// Note its ok to leave the unstarted requests behind, since the
	// Observe is always based on the last reports onchain min seq num.
	if r.isStale(parsedReport.MinSequenceNumber) {
		return false, nil
	}
	// Any timed out requests should be set back to RequestStatusExecutionPending so their execution can be retried in a subsequent report.
	if err = r.orm.ResetExpiredRequests(r.sourceChainId, r.destChainId, RelayMaxInflightTimeSeconds, RequestStatusRelayPending, RequestStatusUnstarted); err != nil {
		// Ok to continue here, we'll try to reset them again on the next round.
		r.l.Errorw("unable to reset expired requests", "err", err)
	}
	// Marking new requests as pending/in-flight
	err = r.orm.UpdateRequestStatus(r.sourceChainId, r.destChainId, parsedReport.MinSequenceNumber, parsedReport.MaxSequenceNumber, RequestStatusRelayPending)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r RelayReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the relayTransmitter enqueues the tx for bptxm,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStale(parsedReport.MinSequenceNumber), nil
}

func (r RelayReportingPlugin) Start() error {
	return nil
}

func (r RelayReportingPlugin) Close() error {
	return nil
}
