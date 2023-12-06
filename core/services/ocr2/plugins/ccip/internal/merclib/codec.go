package merclib

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"

	mercuryutils "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/utils"
	v1report "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v1/types"
	v2report "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v2/types"
	v3report "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v3/types"
)

type ReportWithContext struct {
	FeedId      mercuryutils.FeedID
	FeedVersion mercuryutils.FeedVersion
	V1Report    *v1report.Report
	V2Report    *v2report.Report
	V3Report    *v3report.Report
	Round       uint8
	Epoch       uint32
	Digest      []byte
	FullReport  *FullReport
}

type FullReport struct {
	ReportContext [3][32]byte
	ReportBlob    []byte
	RawRs         [][32]byte
	RawSs         [][32]byte
	RawVs         [32]byte
}

func mustNewType(t string) abi.Type {
	result, err := abi.NewType(t, "", []abi.ArgumentMarshaling{})
	if err != nil {
		panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
	}
	return result
}

var schema = abi.Arguments{
	{Name: "reportContext", Type: mustNewType("bytes32[3]")},
	{Name: "reportBlob", Type: mustNewType("bytes")},
	{Name: "rawRs", Type: mustNewType("bytes32[]")},
	{Name: "rawSs", Type: mustNewType("bytes32[]")},
	{Name: "rawVs", Type: mustNewType("bytes32")},
}

// EncodeFullReport encodes a full Mercury report from the provided data.
func EncodeFullReport(
	reportContext [3][32]byte,
	reportBlob []byte,
	rawRs [][32]byte,
	rawSs [][32]byte,
	rawVs [32]byte,
) ([]byte, error) {
	return schema.Pack(reportContext, reportBlob, rawRs, rawSs, rawVs)
}

// DecodeFullReport reads the "fullReport" from the API response into a struct containing the report context, report data,
// and raw signatures. This functions requires no prep to use, because the schema for the "fullReport" blob is
// common among all report versions (basic, premium, etc),
func DecodeFullReport(fullReport []byte) (*FullReport, error) {
	values, err := schema.Unpack(fullReport)
	if err != nil {
		return nil, fmt.Errorf("failed to decode FullReport: %w", err)
	}
	decoded := new(FullReport)
	if err = schema.Copy(decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy FullReport values to struct: %w", err)
	}

	return decoded, nil
}

// EncodeReportDataV3 takes the report data and encodes it into a report blob that can be sent onchain.
func EncodeReportDataV3(
	feedID [32]byte,
	validFromTimestamp, observationsTimestamp uint32,
	nativeFee, linkFee *big.Int,
	expiresAt uint32,
	benchmarkPrice, bid, ask *big.Int) ([]byte, error) {
	return v3report.GetSchema().Pack(
		feedID,
		validFromTimestamp,
		observationsTimestamp,
		nativeFee,
		linkFee,
		expiresAt,
		benchmarkPrice,
		bid,
		ask,
	)
}

// DecodeReportData takes the report blob (FullReport.ReportBlob), extracts the feeds id, calculates the version from the feed id,
// and finally decodes the report blob using the lib that correlates with the version. The resulting interface can be cast into
// the correct report type as needed.
func DecodeReportData(reportBlob []byte) (mercuryutils.FeedID, interface{}, error) {
	feedIdAbi := abi.Arguments{
		{Name: "feedId", Type: mustNewType("bytes32")},
	}
	reportElements := map[string]interface{}{}
	if err := feedIdAbi.UnpackIntoMap(reportElements, reportBlob); err != nil {
		return mercuryutils.FeedID{}, nil, err
	}
	feedIdInterface, ok := reportElements["feedId"]
	if !ok {
		return mercuryutils.FeedID{}, nil, errors.Errorf("unpacked ReportBlob has no 'feedId'")
	}
	feedIdBytes, ok := feedIdInterface.([32]byte)
	if !ok {
		return mercuryutils.FeedID{}, nil, errors.Errorf("cannot cast ReportBlob feedId to [32]byte, type is %T", feedIdBytes)
	}
	feedID := mercuryutils.FeedID(feedIdBytes)

	switch feedID.Version() {
	case mercuryutils.REPORT_V1: // Legacy/Backward compatible report. Most customers won't use this.
		res, err := v1report.Decode(reportBlob)
		return feedID, res, err // Cast to v1report.Report
	case mercuryutils.REPORT_V2: // Basic report
		res, err := v2report.Decode(reportBlob)
		return feedID, res, err // Cast to v2report.Report
	case mercuryutils.REPORT_V3: // Premium report
		res, err := v3report.Decode(reportBlob)
		return feedID, res, err // Cast to v3report.Report
	default:
		return mercuryutils.FeedID{}, nil, errors.Errorf("unknown report version %d", feedID.Version())
	}
}

// DecodeFullReportAndReportData takes the full report payload, decodes the fullReport blob, and then decodes the report data.
func DecodeFullReportAndReportData(payload []byte) (*ReportWithContext, error) {
	fullReport, err := DecodeFullReport(payload)
	if err != nil {
		return nil, err
	}

	feedID, report, err := DecodeReportData(fullReport.ReportBlob)
	if err != nil {
		return nil, err
	}

	result := &ReportWithContext{
		FeedId:      feedID,
		FeedVersion: feedID.Version(),
		Digest:      fullReport.ReportContext[0][:],
		Round:       fullReport.ReportContext[1][31],
		Epoch:       binary.BigEndian.Uint32(fullReport.ReportContext[1][32-5 : 32-1]),
		FullReport:  fullReport,
	}

	switch feedID.Version() {
	case mercuryutils.REPORT_V1:
		result.V1Report = report.(*v1report.Report)
	case mercuryutils.REPORT_V2:
		result.V2Report = report.(*v2report.Report)
	case mercuryutils.REPORT_V3:
		result.V3Report = report.(*v3report.Report)
	default:
		return nil, errors.Errorf("unknown report version %d", feedID.Version())
	}

	return result, nil
}
