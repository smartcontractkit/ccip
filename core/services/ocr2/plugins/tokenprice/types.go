package tokenprice

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	clcommontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var _ types.ReportingPluginFactory = (*TokenPriceFactory)(nil)

type TokenPriceFactory struct {
	Logger              commontypes.Logger
	ContractTransmitter types.ContractTransmitter
}

// NewReportingPlugin implements types.ReportingPluginFactory.
func (*TokenPriceFactory) NewReportingPlugin(types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	panic("unimplemented")
}

//go:generate mockery --name TokenInfoer --output ./mocks/ --case=underscore
type TokenInfoer interface {
	// Symbol returns the symbol of the ERC-20 token as a string.
	Symbol(ctx context.Context, address common.Address) (string, error)

	// EnabledTokens returns the list of tokens that are enabled.
	// These are tokens we need to fetch USD prices for.
	EnabledTokens(ctx context.Context) ([]common.Address, error)

	// OnchainPrices returns the on-chain prices of the provided tokens in USD to 18 decimals.
	// These are fetched from the price registry contract.
	OnchainPrices(ctx context.Context, tokens []common.Address) (map[string]*utilsbig.Big, error)
}

//go:generate mockery --name PriceServiceClient --output ./mocks/ --case=underscore
type PriceServiceClient interface {
	// LatestPrices returns the latest prices of the provided tokens in USD to 18 decimals.
	LatestPrices(ctx context.Context, tokens []string) (map[string]*utilsbig.Big, error)
}

type Observation struct {
	EnabledTokens []string                 `json:"enabledTokens"`
	Prices        map[string]*utilsbig.Big `json:"prices"`
	OnchainPrices map[string]*utilsbig.Big `json:"onchainPrices"`
}

type Outcome struct {
	MedianPrices        map[string]*utilsbig.Big `json:"prices"`
	MedianOnchainPrices map[string]*utilsbig.Big `json:"onchainPrices"`
}

// OffchainConfig is the offchain config for the token price job.
// This is posted onchain as part of setConfig transaction.
// It is posted as an abi-encoded tuple, i.e
// abi.encode([reportPPB, acceptPPB, deltaC])
// where all of these fields are uint64's.
type OffchainConfig struct {
	// ReportPPB determines the relative deviation between the median (i.e.
	// answer) in the contract and the current median of observations (offchain)
	// at which a report should be issued. That is, a report is issued if
	// abs((offchainMedian - contractMedian)/contractMedian) >= reportPPB.
	// PPB is parts-per-billion
	ReportPPB uint64
	// AcceptPPB determines the relative deviation between the median in a
	// newly generated report considered for transmission and the median of the
	// currently pending report. That is, a report is accepted for transmission
	// if abs((newMedian - pendingMedian)/pendingMedian) >= acceptPPB. If no
	// report is pending, this variable has no effect.
	// PPB is parts-per-billion
	AcceptPPB uint64
	// DeltaC is the maximum age of the latest report in the contract. If the
	// maximum age is exceeded, a new report will be created by the report
	// generation protocol.
	DeltaC time.Duration
}

// Encode encodes the offchain config as an abi-encoded tuple suitable
// for passing to the contract.
func (o OffchainConfig) Encode() ([]byte, error) {
	return utils.ABIEncode(
		`[{ "type": "uint64" }, { "type": "uint64" }, { "type": "uint64" }]`,
		o.ReportPPB,
		o.AcceptPPB,
		o.DeltaC,
	)
}

func DecodeOffchainConfig(encodedConfig []byte) (OffchainConfig, error) {
	decoded, err := utils.ABIDecode(
		`[{ "type": "uint64" }, { "type": "uint64" }, { "type": "uint64" }]`,
		encodedConfig)
	if err != nil {
		return OffchainConfig{}, err
	}
	reportPPB, ok := decoded[0].(uint64)
	if !ok {
		return OffchainConfig{}, fmt.Errorf("failed to decode reportPPB from %+v", decoded)
	}
	acceptPPB, ok := decoded[1].(uint64)
	if !ok {
		return OffchainConfig{}, fmt.Errorf("failed to decode acceptPPB from %+v", decoded)
	}
	deltaC, ok := decoded[2].(uint64)
	if !ok {
		return OffchainConfig{}, fmt.Errorf("failed to decode deltaC from %+v", decoded)
	}
	return OffchainConfig{
		ReportPPB: reportPPB,
		AcceptPPB: acceptPPB,
		DeltaC:    time.Duration(deltaC),
	}, nil
}

type TokenPriceContract interface {
	// LatestTransmissionDetails returns the latest transmission details from the contract.
	LatestTransmissionDetails(ctx context.Context) (
		configDigest types.ConfigDigest,
		seqNr uint64,
		latestTimestamp time.Time,
		err error,
	)
}

func newTokenPriceContract(chainReader clcommontypes.ChainReader, address common.Address) TokenPriceContract {
	contract := clcommontypes.BoundContract{Address: address.String(), Name: "tokenprice", Pending: true}
	return &tokenPriceContract{chainReader, contract}
}

type tokenPriceContract struct {
	chainReader clcommontypes.ChainReader
	contract    clcommontypes.BoundContract
}

var _ TokenPriceContract = &tokenPriceContract{}

func (t *tokenPriceContract) LatestTransmissionDetails(ctx context.Context) (
	configDigest types.ConfigDigest,
	seqNr uint64,
	latestTimestamp time.Time,
	err error,
) {
	var resp struct {
		ConfigDigest types.ConfigDigest
		SeqNr        uint64
		Timestamp    time.Time
	}
	err = t.chainReader.GetLatestValue(ctx, t.contract, "LatestTransmissionDetails", nil, &resp)
	if err != nil {
		return
	}
	return resp.ConfigDigest, resp.SeqNr, resp.Timestamp, nil
}
