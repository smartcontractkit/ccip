package ccip

import (
	"sync"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	telemPb "github.com/smartcontractkit/chainlink/v2/core/services/synchronization/telem"
)

// TelemetryCollector is an interface for collecting telemetry data.
type TelemetryCollector interface {
	ReportCommit(observations map[commontypes.OracleID]CommitObservation, report ccipdata.CommitStoreReport, epochAndRound types.ReportTimestamp)
	ReportExec(observations map[commontypes.OracleID]ExecutionObservation, messages []ObservedMessage, epochAndRound types.ReportTimestamp)
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

// CollectCommit collects commit report data and sends it to the OTI monitoring endpoint.
func (tc *telemetryCollector) ReportCommit(
	observations map[commontypes.OracleID]CommitObservation,
	report ccipdata.CommitStoreReport,
	epochAndRound types.ReportTimestamp) {

	// collect telemetry data from valid observations
	obs := make([]*telemPb.CommitObservation, 0, len(observations))
	for oracleId, o := range observations {
		tps := make([]*telemPb.TokenPrice, 0, len(o.TokenPricesUSD))
		for addr, price := range o.TokenPricesUSD {
			tps = append(tps, &telemPb.TokenPrice{
				Address:  addr.Bytes(),
				PriceUsd: price.Bytes(),
			})
		}
		obs = append(obs, &telemPb.CommitObservation{
			IntervalMin:       o.Interval.Min,
			IntervalMax:       o.Interval.Max,
			TokenPrices:       tps,
			SourceGasPriceUsd: o.SourceGasPriceUSD.Bytes(),
			OracleID:          uint32(oracleId),
		})
	}

	// collect telemetry data from report
	telem := &telemPb.CCIPTelemWrapper{
		Msg: &telemPb.CCIPTelemWrapper_CommitReport{
			CommitReport: &telemPb.CCIPCommitReportSummary{
				LenTokenPrices: uint32(len(report.TokenPrices)),
				LenGasPrices:   uint32(len(report.GasPrices)), // XXX: if the len is short, would it be better to send the actual gas prices?
				IntervalMin:    report.Interval.Min,
				IntervalMax:    report.Interval.Max,
				Epoch:          epochAndRound.Epoch,
				Round:          uint32(epochAndRound.Round),
				Observations:   obs,
			},
		},
	}

	tc.maybeSend(telem)
}

// CollectExec collects execution report data and sends it to the OTI monitoring endpoint.
func (tc *telemetryCollector) ReportExec(observations map[commontypes.OracleID]ExecutionObservation, messages []ObservedMessage, epochAndRound types.ReportTimestamp) {
	obs := make([]*telemPb.ExecutionObservation, len(observations))
	for oracleId, observation := range observations {
		seqNrs := maps.Keys(observation.Messages)
		slices.Sort(seqNrs)
		obs = append(obs, &telemPb.ExecutionObservation{OracleID: uint32(oracleId), SeqNrs: seqNrs})
	}

	telem := &telemPb.CCIPTelemWrapper{
		Msg: &telemPb.CCIPTelemWrapper_ExecutionReport{
			ExecutionReport: &telemPb.CCIPExecutionReportSummary{
				LenObservedMessages: uint32(len(messages)),
				Epoch:               epochAndRound.Epoch,
				Round:               uint32(epochAndRound.Round),
				Observations:        obs,
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
