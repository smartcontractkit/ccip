package v1_2_0

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/config"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store_1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
)

func TestCommitReportEncoding(t *testing.T) {
	t.Parallel()
	ctx := testutils.Context(t)
	report := cciptypes.CommitStoreReport{
		TokenPrices: []cciptypes.TokenPrice{
			{
				Token: cciptypes.Address(utils.RandomAddress().String()),
				Value: big.NewInt(9e18),
			},
			{
				Token: cciptypes.Address(utils.RandomAddress().String()),
				Value: big.NewInt(1e18),
			},
		},
		GasPrices: []cciptypes.GasPrice{
			{
				DestChainSelector: rand.Uint64(),
				Value:             big.NewInt(2000e9),
			},
			{
				DestChainSelector: rand.Uint64(),
				Value:             big.NewInt(3000e9),
			},
		},
		MerkleRoot: [32]byte{123},
		Interval:   cciptypes.CommitStoreInterval{Min: 1, Max: 10},
	}

	c, err := NewCommitStore(logger.TestLogger(t), utils.RandomAddress(), nil, mocks.NewLogPoller(t))
	assert.NoError(t, err)

	encodedReport, err := c.EncodeCommitReport(ctx, report)
	require.NoError(t, err)
	assert.Greater(t, len(encodedReport), 0)

	decodedReport, err := c.DecodeCommitReport(ctx, encodedReport)
	require.NoError(t, err)
	require.Equal(t, report, decodedReport)
}

func TestCommitStoreV120ffchainConfigEncoding(t *testing.T) {
	t.Parallel()
	validConfig := JSONCommitOffchainConfig{
		SourceFinalityDepth:      3,
		DestFinalityDepth:        4,
		GasPriceHeartBeat:        *config.MustNewDuration(1 * time.Minute),
		DAGasPriceDeviationPPB:   10,
		ExecGasPriceDeviationPPB: 11,
		TokenPriceHeartBeat:      *config.MustNewDuration(2 * time.Minute),
		TokenPriceDeviationPPB:   12,
		InflightCacheExpiry:      *config.MustNewDuration(3 * time.Minute),
	}

	require.NoError(t, validConfig.Validate())

	tests := []struct {
		name       string
		want       JSONCommitOffchainConfig
		errPattern string
	}{
		{
			name: "legacy offchain config format parses",
			want: validConfig,
		},
		{
			name: "can omit finality depth",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.SourceFinalityDepth = 0
				c.DestFinalityDepth = 0
			}),
		},
		{
			name: "can set PriceReportingDisabled",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.PriceReportingDisabled = true
			}),
		},
		{
			name: "must set GasPriceHeartBeat",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.GasPriceHeartBeat = *config.MustNewDuration(0)
			}),
			errPattern: "GasPriceHeartBeat",
		},
		{
			name: "must set ExecGasPriceDeviationPPB",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.ExecGasPriceDeviationPPB = 0
			}),
			errPattern: "ExecGasPriceDeviationPPB",
		},
		{
			name: "must set TokenPriceHeartBeat",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.TokenPriceHeartBeat = *config.MustNewDuration(0)
			}),
			errPattern: "TokenPriceHeartBeat",
		},
		{
			name: "must set TokenPriceDeviationPPB",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.TokenPriceDeviationPPB = 0
			}),
			errPattern: "TokenPriceDeviationPPB",
		},
		{
			name: "must set InflightCacheExpiry",
			want: modifyCopy(validConfig, func(c *JSONCommitOffchainConfig) {
				c.InflightCacheExpiry = *config.MustNewDuration(0)
			}),
			errPattern: "InflightCacheExpiry",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			exp := tc.want
			encode, err := ccipconfig.EncodeOffchainConfig(&exp)
			require.NoError(t, err)
			got, err := ccipconfig.DecodeOffchainConfig[JSONCommitOffchainConfig](encode)

			if tc.errPattern != "" {
				require.ErrorContains(t, err, tc.errPattern)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestCommitStoreV120ffchainConfigDecodingCompatibility(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                   string
		config                 []byte
		priceReportingDisabled bool
	}{
		{
			name: "with MaxGasPrice",
			config: []byte(`{
				"SourceFinalityDepth": 3,
				"DestFinalityDepth": 4,
				"GasPriceHeartBeat": "60s",
				"DAGasPriceDeviationPPB": 10,
				"ExecGasPriceDeviationPPB": 11,
				"TokenPriceHeartBeat": "120s",
				"TokenPriceDeviationPPB": 12,
				"MaxGasPrice": 100000000,
				"SourceMaxGasPrice": 100000000,
				"InflightCacheExpiry": "180s"
			}`),
			priceReportingDisabled: false,
		},
		{
			name: "without MaxGasPrice",
			config: []byte(`{
				"SourceFinalityDepth": 3,
				"DestFinalityDepth": 4,
				"GasPriceHeartBeat": "60s",
				"DAGasPriceDeviationPPB": 10,
				"ExecGasPriceDeviationPPB": 11,
				"TokenPriceHeartBeat": "120s",
				"TokenPriceDeviationPPB": 12,
				"InflightCacheExpiry": "180s"
			}`),
			priceReportingDisabled: false,
		},
		{
			name: "with PriceReportingDisabled",
			config: []byte(`{
				"SourceFinalityDepth": 3,
				"DestFinalityDepth": 4,
				"GasPriceHeartBeat": "60s",
				"DAGasPriceDeviationPPB": 10,
				"ExecGasPriceDeviationPPB": 11,
				"TokenPriceHeartBeat": "120s",
				"TokenPriceDeviationPPB": 12,
				"InflightCacheExpiry": "180s",
				"PriceReportingDisabled": true
			}`),
			priceReportingDisabled: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			decoded, err := ccipconfig.DecodeOffchainConfig[JSONCommitOffchainConfig](tc.config)
			require.NoError(t, err)
			require.Equal(t, JSONCommitOffchainConfig{
				SourceFinalityDepth:      3,
				DestFinalityDepth:        4,
				GasPriceHeartBeat:        *config.MustNewDuration(1 * time.Minute),
				DAGasPriceDeviationPPB:   10,
				ExecGasPriceDeviationPPB: 11,
				TokenPriceHeartBeat:      *config.MustNewDuration(2 * time.Minute),
				TokenPriceDeviationPPB:   12,
				InflightCacheExpiry:      *config.MustNewDuration(3 * time.Minute),
				PriceReportingDisabled:   tc.priceReportingDisabled,
			}, decoded)
		})
	}
}

func Test_CommitReportAccepted(t *testing.T) {
	ctx := testutils.Context(t)
	chainID := testutils.NewRandomEVMChainID()
	orm := logpoller.NewORM(chainID, pgtest.NewSqlxDB(t), logger.TestLogger(t))
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Hour,
		FinalityDepth:            2,
		BackfillBatchSize:        20,
		RpcBatchSize:             10,
		KeepFinalizedBlocksDepth: 1000,
	}
	lp := logpoller.NewLogPoller(orm, nil, logger.TestLogger(t), nil, lpOpts)

	commitStoreAddr := utils.RandomAddress()
	merkleRoot := [32]byte{123}
	inputLogs := []logpoller.Log{
		createReportAcceptedLog(t, chainID, commitStoreAddr, 2, 1, merkleRoot),
		createReportAcceptedLog(t, chainID, commitStoreAddr, 3, 1, merkleRoot),
		createReportAcceptedLog(t, chainID, commitStoreAddr, 5, 1, merkleRoot),
		createReportAcceptedLog(t, chainID, commitStoreAddr, 5, 2, utils.RandomBytes32()),
		createReportAcceptedLog(t, chainID, commitStoreAddr, 5, 3, utils.RandomBytes32()),
		createReportAcceptedLog(t, chainID, commitStoreAddr, 8, 1, utils.RandomBytes32()),
		createReportAcceptedLog(t, chainID, commitStoreAddr, 9, 1, utils.RandomBytes32()),
		createReportAcceptedLog(t, chainID, utils.RandomAddress(), 9, 1, utils.RandomBytes32()),
	}
	require.NoError(t, orm.InsertLogsWithBlock(ctx, inputLogs, logpoller.NewLogPollerBlock(utils.RandomBytes32(), 20, time.Now(), 5)))

	commitStoreABI := abihelpers.MustParseABI(commit_store_1_2_0.CommitStoreABI)
	eventSig := abihelpers.MustGetEventID("ReportAccepted", commitStoreABI)

	logs, err := lp.Logs(ctx, 0, 100, eventSig, commitStoreAddr)
	require.NoError(t, err)
	require.Len(t, logs, 7)

	logs, err = lp.LogsDataWordRange(ctx, eventSig, commitStoreAddr, 4, merkleRoot, merkleRoot, 0)
	require.NoError(t, err)
	require.Len(t, logs, 3)
}

func createReportAcceptedLog(t testing.TB, chainID *big.Int, address common.Address, blockNumber int64, logIndex int64, merkleRoot common.Hash) logpoller.Log {
	tAbi, err := commit_store_1_2_0.CommitStoreMetaData.GetAbi()
	require.NoError(t, err)
	eseEvent, ok := tAbi.Events["ReportAccepted"]
	require.True(t, ok)

	gasPriceUpdates := make([]commit_store_1_2_0.InternalGasPriceUpdate, 100)
	tokenPriceUpdates := make([]commit_store_1_2_0.InternalTokenPriceUpdate, 100)

	for i := 0; i < 100; i++ {
		gasPriceUpdates[i] = commit_store_1_2_0.InternalGasPriceUpdate{
			DestChainSelector: uint64(i),
			UsdPerUnitGas:     big.NewInt(int64(i)),
		}
		tokenPriceUpdates[i] = commit_store_1_2_0.InternalTokenPriceUpdate{
			SourceToken: utils.RandomAddress(),
			UsdPerToken: big.NewInt(int64(i)),
		}
	}

	message := commit_store_1_2_0.CommitStoreCommitReport{
		PriceUpdates: commit_store_1_2_0.InternalPriceUpdates{
			TokenPriceUpdates: tokenPriceUpdates,
			GasPriceUpdates:   gasPriceUpdates,
		},
		Interval:   commit_store_1_2_0.CommitStoreInterval{Min: 1, Max: 10},
		MerkleRoot: merkleRoot,
	}

	logData, err := eseEvent.Inputs.Pack(message)
	require.NoError(t, err)

	topic0 := commit_store_1_2_0.CommitStoreReportAccepted{}.Topic()

	return logpoller.Log{
		Topics: [][]byte{
			topic0[:],
		},
		Data:        logData,
		LogIndex:    logIndex,
		BlockHash:   utils.RandomBytes32(),
		BlockNumber: blockNumber,
		EventSig:    topic0,
		Address:     address,
		TxHash:      utils.RandomBytes32(),
		EvmChainId:  ubig.New(chainID),
	}
}
