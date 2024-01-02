package tokenprice

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_price_ocr"
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
	// Symbols returns the symbol of the provided ERC-20 tokens as strings.
	Symbols(ctx context.Context, address []common.Address) ([]string, error)

	// EnabledTokens returns the list of tokens that are enabled.
	// These are tokens we need to fetch USD prices for.
	EnabledTokens(ctx context.Context) ([]common.Address, error)

	// OnchainTokenPrices returns the on-chain prices of the provided tokens in USD to 18 decimals.
	// These are fetched from the price registry contract.
	OnchainTokenPrices(ctx context.Context, tokens []common.Address) (map[string]*utilsbig.Big, error)
}

//go:generate mockery --name PriceServiceClient --output ./mocks/ --case=underscore
type PriceServiceClient interface {
	// LatestPrices returns the latest prices of the provided tokens in USD to 18 decimals
	// from the price service.
	LatestPrices(ctx context.Context, tokens []string) (map[string]*utilsbig.Big, error)
}

type Observation struct {
	// EnabledTokens is a mapping from token symbol to token address.
	EnabledTokens map[string]common.Address `json:"enabledTokens"`
	// TokenPrices is a mapping from token symbol to token price in USD to 18 decimals.
	TokenPrices map[string]*utilsbig.Big `json:"tokenPrices"`
	// OnchainTokenPrices is a mapping from token symbol to token price in USD to 18 decimals.
	OnchainTokenPrices map[string]*utilsbig.Big `json:"onchainTokenPrices"`
	// Heartbeats is a mapping from token symbol to the time of the last price update.
	// Time is represented as an int64, that is, a unix timestamp in the UTC timezone.
	Heartbeats map[string]int64 `json:"heartbeats"`
}

type Outcome struct {
	// MedianTokenPrices is a mapping from token symbol to median token price in USD to 18 decimals.
	MedianTokenPrices map[string]*utilsbig.Big `json:"medianTokenPrices"`
	// MedianOnchainTokenPrices is a mapping from token symbol to median token price in USD to 18 decimals.
	MedianOnchainTokenPrices map[string]*utilsbig.Big `json:"medianOnchainTokenPrices"`
	// MedianTimestamps is a mapping from token symbol to median last updated timestamp.
	MedianTimestamps map[string]int64 `json:"medianTimestamps"`
	// QuorumTokens is a mapping from token symbol to token address.
	QuorumTokens map[string]common.Address `json:"quorumTokens"`
}

// OffchainConfig is the offchain config for the token price job.
// This is posted onchain as part of setConfig transaction.
// It is posted as an abi-encoded tuple, i.e
// abi.encode([reportPPB, deltaC, maxPriceUpdates])
// where all of these fields are uint64's.
type OffchainConfig struct {
	// ReportPPB determines the relative deviation between the median (i.e.
	// answer) in the contract and the current median of observations (offchain)
	// at which a report should be issued. That is, a report is issued if
	// abs((offchainMedian - contractMedian)/contractMedian) >= reportPPB.
	// PPB is parts-per-billion
	// TODO: different deviation per token?
	ReportPPB uint64
	// DeltaC is the maximum age of the latest report in the contract. If the
	// maximum age is exceeded, a new report will be created by the report
	// generation protocol.
	// TODO: different hearbeat per token?
	DeltaC time.Duration
	// MaxPriceUpdates is the maximum number of price updates in a single report.
	MaxPriceUpdates uint64
}

// Encode encodes the offchain config as an abi-encoded tuple suitable
// for passing to the contract.
func (o OffchainConfig) Encode() ([]byte, error) {
	return utils.ABIEncode(
		`[{ "type": "uint64" }, { "type": "uint64" }, { "type": "uint64" }]`,
		o.ReportPPB,
		uint64(o.DeltaC),
		o.MaxPriceUpdates,
	)
}

func DecodeOffchainConfig(encodedConfig []byte) (OffchainConfig, error) {
	decoded, err := utils.ABIDecode(
		`[{ "type": "uint64" }, { "type": "uint64" }, { "type": "uint64" }]`,
		encodedConfig)
	if err != nil {
		return OffchainConfig{}, err
	}
	if len(decoded) != 4 {
		return OffchainConfig{}, fmt.Errorf("expected 4 fields in decoded offchain config, got %d", len(decoded))
	}
	reportPPB, ok := decoded[0].(uint64)
	if !ok {
		return OffchainConfig{}, fmt.Errorf("failed to decode reportPPB from %+v", decoded)
	}
	deltaC, ok := decoded[1].(uint64)
	if !ok {
		return OffchainConfig{}, fmt.Errorf("failed to decode deltaC from %+v", decoded)
	}
	maxPriceUpdates, ok := decoded[2].(uint64)
	if !ok {
		return OffchainConfig{}, fmt.Errorf("failed to decode maxPriceUpdates from %+v", decoded)
	}
	return OffchainConfig{
		ReportPPB:       reportPPB,
		DeltaC:          time.Duration(deltaC),
		MaxPriceUpdates: maxPriceUpdates,
	}, nil
}

var tokenPriceABI = evmtypes.MustGetABI(token_price_ocr.TokenPriceOCRABI)

func ABIEncodeReportData(reportData token_price_ocr.TokenPriceOCRReport) ([]byte, error) {
	encoded, err := tokenPriceABI.Pack("exposeForEncoding", reportData)
	if err != nil {
		return nil, err
	}
	// The first 4 bytes are the function selector, so we drop them
	return encoded[4:], nil
}
