package ccip

import (
	"context"
	"encoding/json"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const RelayMaxInflightTimeSeconds = 180

var _ types.ReportingPluginFactory = &RelayReportingPluginFactory{}
var _ types.ReportingPlugin = &RelayReportingPlugin{}

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
			Type: mustType("uint64"),
		},
		{
			Name: "maxSequenceNumber",
			Type: mustType("uint64"),
		},
	}
}

// EncodeRelayReport abi encodes an offramp.CCIPRelayReport.
func EncodeRelayReport(relayReport *offramp.CCIPRelayReport) (types.Report, error) {
	report, err := makeRelayReportArgs().PackValues([]interface{}{relayReport.MerkleRoot, relayReport.MinSequenceNumber, relayReport.MaxSequenceNumber})
	if err != nil {
		return nil, err
	}
	return report, nil
}

// DecodeRelayReport abi decodes a types.Report to an offramp.CCIPRelayReport
func DecodeRelayReport(report types.Report) (*offramp.CCIPRelayReport, error) {
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
	min, ok := unpacked[1].(uint64)
	if !ok {
		return nil, errors.New("invalid min")
	}
	max, ok := unpacked[2].(uint64)
	if !ok {
		return nil, errors.New("invalid max")
	}
	return &offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: min,
		MaxSequenceNumber: max,
	}, nil
}

type RelayReportingPluginFactory struct {
	l       logger.Logger
	orm     ORM
	onRamp  common.Address
	offRamp *offramp.OffRamp
}

// NewRelayReportingPluginFactory return a new RelayReportingPluginFactory.
func NewRelayReportingPluginFactory(l logger.Logger, orm ORM, offRamp *offramp.OffRamp, onRamp common.Address) types.ReportingPluginFactory {
	return &RelayReportingPluginFactory{l: l, orm: orm, offRamp: offRamp, onRamp: onRamp}
}

// NewReportingPlugin returns the ccip RelayReportingPlugin and satisfies the ReportingPluginFactory interface.
// This function can error if the onRamp or offRamp chainIDs are not properly set.
func (rf *RelayReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	destChainId, err := rf.offRamp.CHAINID(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, errors.WithStack(err)
	}
	sourceChainId, err := rf.offRamp.SOURCECHAINID(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, errors.WithStack(err)
	}
	return RelayReportingPlugin{rf.l, config.F, rf.orm, sourceChainId, destChainId, rf.onRamp, rf.offRamp}, types.ReportingPluginInfo{
		Name:          "CCIPRelay",
		UniqueReports: true,
		Limits: types.ReportingPluginLimits{
			MaxQueryLength: 0, // We do not use the query phase.
			// TODO: https://app.shortcut.com/chainlinklabs/story/30171/define-report-plugin-limits
			MaxObservationLength: 100000, // TODO
			MaxReportLength:      100000, // TODO
		},
	}, nil
}

type RelayReportingPlugin struct {
	l             logger.Logger
	F             int
	orm           ORM
	sourceChainId *big.Int
	destChainId   *big.Int
	onRamp        common.Address
	offRamp       *offramp.OffRamp
}

func (r RelayReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// We don't use a query for this reporting plugin, so we can just leave it empty here
	return types.Query{}, nil
}

func (r RelayReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	nextMin, err := r.nextMinSeqNumForOffRamp()
	if err != nil {
		return nil, err
	}

	// Because we explicitly look for requests with status RequestStatusUnstarted, inflight requests
	// are ignored.
	unstartedReqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, r.onRamp, r.offRamp.Address(), nextMin, math.MaxInt64, RequestStatusUnstarted, nil, nil)
	if err != nil {
		return nil, err
	}

	// If there are no request to process, return an observation with MinSeqNum and MaxSeqNum equal to NoRequestsToProcess
	// which should not result in a new report being generated during the Report step.
	var (
		minSeqNum = NoRequestsToProcess
		maxSeqNum = NoRequestsToProcess
	)
	if len(unstartedReqs) != 0 {
		minSeqNum = unstartedReqs[0].SeqNum
		maxSeqNum = unstartedReqs[len(unstartedReqs)-1].SeqNum
	}
	return json.Marshal(&Observation{
		MinSeqNum: minSeqNum,
		MaxSeqNum: maxSeqNum,
	})
}

func (r RelayReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	nonEmptyObservations := getNonEmptyObservations(r.l, observations)
	// Need at least F+1 valid observations
	if len(nonEmptyObservations) <= r.F {
		r.l.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	minSeqNum, maxSeqNum, err := getMinMaxSequenceNumbers(nonEmptyObservations, r.F)
	if err != nil {
		return false, nil, errors.Wrap(err, "failed getting valid sequence numbers from observations")
	}
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, r.onRamp, r.offRamp.Address(), minSeqNum, maxSeqNum, RequestStatusUnstarted, nil, nil)
	if err != nil {
		return false, nil, err
	}
	// Cannot construct a report for which we haven't seen all the messages.
	if len(reqs) == 0 {
		return false, nil, errors.Errorf("do not have all the messages in report, have zero messages, report has min %v max %v", minSeqNum, maxSeqNum)
	}
	if reqs[len(reqs)-1].SeqNum < maxSeqNum {
		return false, nil, errors.Errorf("do not have all the messages in report, our max %v reports max %v", reqs[len(reqs)-1].SeqNum, maxSeqNum)
	}

	nextMin, err := r.nextMinSeqNumForOffRamp()
	if err != nil {
		return false, nil, err
	}
	if nextMin > minSeqNum {
		return false, nil, errors.Errorf("invalid min seq number got %v want %v", minSeqNum, nextMin)
	}
	encodedReport, err := EncodeRelayReport(r.buildReport(reqs))
	if err != nil {
		return false, nil, err
	}
	return true, encodedReport, nil
}

func (r RelayReportingPlugin) nextMinSeqNumForOffRamp() (int64, error) {
	lastReport, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		return 0, err
	}
	if lastReport.MerkleRoot == [32]byte{} {
		return 0, nil
	}
	if lastReport.MaxSequenceNumber+1 > math.MaxInt64 {
		return 0, errors.New("sequenceNumber is larger than max int64")
	}
	return int64(lastReport.MaxSequenceNumber + 1), nil
}

func (r RelayReportingPlugin) isStaleReport(report *offramp.CCIPRelayReport) bool {
	nextMin, err := r.nextMinSeqNumForOffRamp()
	if err != nil {
		// Assume it's a transient issue getting the last report
		// Will try again on the next round
		return true
	}
	// TODO(36248): Add is offramp healthy check
	// If the next min is already greater than this reports min,
	// this report is stale.
	return uint64(nextMin) > report.MinSequenceNumber
}

// buildReport assumes there is at least one message in reqs.
func (r RelayReportingPlugin) buildReport(reqs []*Request) *offramp.CCIPRelayReport {
	// Take all these request and produce a merkle root of them
	var leaves [][]byte
	for _, req := range reqs {
		leaves = append(leaves, req.Raw)
	}

	// Note Index doesn't matter, we just want the root
	root, _ := GenerateMerkleProof(32, leaves, 0)
	return &offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: uint64(reqs[0].SeqNum),
		MaxSequenceNumber: uint64(reqs[len(reqs)-1].SeqNum),
	}
}

func (r RelayReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	if parsedReport.MinSequenceNumber > math.MaxInt64 || parsedReport.MaxSequenceNumber > math.MaxInt64 {
		r.l.Errorw("SequenceNumber is larger than max int64")
		return false, errors.New("sequenceNumber is larger than max int64")
	}
	// Note it's ok to leave the unstarted requests behind, since the
	// 'Observe' is always based on the last reports onchain min seq num.
	if r.isStaleReport(parsedReport) {
		return false, nil
	}
	// Any timed out requests should be set back to RequestStatusExecutionPending so their execution can be retried in a subsequent report.
	if err = r.orm.ResetExpiredRequests(r.sourceChainId, r.destChainId, r.onRamp, r.offRamp.Address(), RelayMaxInflightTimeSeconds, RequestStatusRelayPending, RequestStatusUnstarted); err != nil {
		// Ok to continue here, we'll try to reset them again on the next round.
		r.l.Errorw("Unable to reset expired requests", "err", err)
	}
	// Marking new requests as pending/in-flight
	err = r.orm.UpdateRequestStatus(r.sourceChainId, r.destChainId, r.onRamp, r.offRamp.Address(), int64(parsedReport.MinSequenceNumber), int64(parsedReport.MaxSequenceNumber), RequestStatusRelayPending)
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
	return !r.isStaleReport(parsedReport), nil
}

func (r RelayReportingPlugin) Start() error {
	return nil
}

func (r RelayReportingPlugin) Close() error {
	return nil
}
