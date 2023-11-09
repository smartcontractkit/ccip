package ccip

import (
	"slices"
	"sync"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	telemPb "github.com/smartcontractkit/chainlink/v2/core/services/synchronization/telem"
)

// TelemetryCollector is an interface for collecting telemetry data. All underlying sends are non-blocking and asynchronous.
type TelemetryCollector interface {
	TrackCommitObservation(observation CommitObservation, epochAndRound types.ReportTimestamp)
	TrackExecObservation(observation ExecutionObservation, epochAndRound types.ReportTimestamp)
}

type telemetryCollector struct {
	monitoringEndpoint commontypes.MonitoringEndpoint
	lggr               logger.Logger
}

var (
	telemCollector telemetryCollector
	telemOnce      sync.Once
)

// NewTelemetryCollector creates a single telemetry collector. It's thread-safe.
func NewTelemetryCollector(monitoringEndpoint commontypes.MonitoringEndpoint, lggr logger.Logger) *telemetryCollector {
	telemOnce.Do(func() { // For Java/GOF fans -- it's a singleton.
		telemCollector = telemetryCollector{
			monitoringEndpoint: monitoringEndpoint,
			lggr:               lggr,
		}
	})
	return &telemCollector
}

func (tc *telemetryCollector) TrackCommitObservation(observation CommitObservation, epochAndRound types.ReportTimestamp) {
	telem := &telemPb.CCIPTelemWrapper{
		Msg: &telemPb.CCIPTelemWrapper_CommitObservation{
			CommitObservation: &telemPb.CCIPCommitObservation{
				// Mising fields
				LenTokenPrices: uint32(len(observation.TokenPricesUSD)),
				IntervalMin:    observation.Interval.Min,
				IntervalMax:    observation.Interval.Max,
				Epoch:          epochAndRound.Epoch,
				Round:          uint32(epochAndRound.Round),
			},
		},
	}

	tc.maybeSend(telem)
}

func (tc *telemetryCollector) TrackExecObservation(observation ExecutionObservation, epochAndRound types.ReportTimestamp) {
	observedSeqNrs := maps.Keys(observation.Messages)
	slices.Sort(observedSeqNrs)

	var lenTokenData uint32
	tokenData := make([][]byte, 0, len(observation.Messages))
	for _, msg := range observation.Messages {
		lenTokenData += uint32(len(msg.TokenData))
		tokenData = append(tokenData, msg.TokenData...)
	}

	telem := &telemPb.CCIPTelemWrapper{
		Msg: &telemPb.CCIPTelemWrapper_ExecutionObservation{
			ExecutionObservation: &telemPb.CCIPExecutionObservation{
				LenObservedMessages: uint32(len(observation.Messages)),
				LenTokenData:        lenTokenData,
				TokenData:           tokenData,
				SeqNrs:              observedSeqNrs,
				Epoch:               epochAndRound.Epoch,
				Round:               uint32(epochAndRound.Round),
			},
		},
	}

	tc.maybeSend(telem)
}

// maybeSend sends the telemetry data to the OTI monitoring endpoint.
func (tc *telemetryCollector) maybeSend(telemetry *telemPb.CCIPTelemWrapper) {
	bytes, err := proto.Marshal(telemetry)
	if err != nil || tc.monitoringEndpoint == nil {
		// Telemetry related errors are not critical and must not affect
		// execution, so we log them and continue.
		tc.lggr.Errorw("cannot marshal or send telemetry", "err", err)
	} else {
		tc.monitoringEndpoint.SendLog(bytes)
	}
}
