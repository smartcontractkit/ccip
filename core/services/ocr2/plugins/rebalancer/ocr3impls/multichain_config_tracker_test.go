package ocr3impls_test

import (
	"encoding/hex"
	"strconv"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/no_op_ocr3"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager"
	lm_mocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/ocr3impls"
	mocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/rebalancermocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
)

func setupLogPoller[RI ocr3impls.MultichainMeta](t *testing.T, db *sqlx.DB, bs *keyringsAndSigners[RI]) (logpoller.LogPoller, testUniverse[RI]) {
	lggr := logger.TestLogger(t)

	o := logpoller.NewORM(testutils.SimulatedChainID, db, lggr, pgtest.NewQConfig(false))

	// create the universe which will deploy the OCR contract and set config
	// we will replay on the log poller to get the appropriate ConfigSet log
	uni := newTestUniverse[RI](t, bs)

	lp := logpoller.NewLogPoller(o, uni.simClient, lggr, 1*time.Second, false, 100, 100, 100, 200)
	return lp, uni
}

func TestConfigSet(t *testing.T) {
	require.Equal(t, no_op_ocr3.NoOpOCR3ConfigSet{}.Topic().Hex(), ocr3impls.ConfigSet.Hex())
}

func TestMultichainConfigTracker_New(t *testing.T) {
	t.Run("master chain not in log pollers", func(t *testing.T) {
		db := pgtest.NewSqlxDB(t)
		_, uni := setupLogPoller[multichainMeta](t, db, nil)

		masterChain := relay.ID{
			Network: relay.EVM,
			ChainID: testutils.SimulatedChainID.String(),
		}
		mockLMFactory := mocks.NewFactory(t)
		_, err := ocr3impls.NewMultichainConfigTracker(
			masterChain,
			logger.TestLogger(t),
			map[relay.ID]logpoller.LogPoller{},
			uni.simClient,
			uni.wrapper.Address(),
			mockLMFactory,
			ocr3impls.TransmitterCombiner,
			nil,
		)
		require.Error(t, err, "expected error creating multichain config tracker")
	})

	t.Run("combiner is nil", func(t *testing.T) {
		db := pgtest.NewSqlxDB(t)
		lp, uni := setupLogPoller[multichainMeta](t, db, nil)

		masterChain := relay.ID{
			Network: relay.EVM,
			ChainID: testutils.SimulatedChainID.String(),
		}
		mockLMFactory := mocks.NewFactory(t)
		_, err := ocr3impls.NewMultichainConfigTracker(
			masterChain,
			logger.TestLogger(t),
			map[relay.ID]logpoller.LogPoller{masterChain: lp},
			uni.simClient,
			uni.wrapper.Address(),
			mockLMFactory,
			nil,
			nil,
		)
		require.Error(t, err, "expected error creating multichain config tracker")
	})

	t.Run("factory is nil", func(t *testing.T) {
		db := pgtest.NewSqlxDB(t)
		lp, uni := setupLogPoller[multichainMeta](t, db, nil)

		masterChain := relay.ID{
			Network: relay.EVM,
			ChainID: testutils.SimulatedChainID.String(),
		}
		_, err := ocr3impls.NewMultichainConfigTracker(
			masterChain,
			logger.TestLogger(t),
			map[relay.ID]logpoller.LogPoller{masterChain: lp},
			uni.simClient,
			uni.wrapper.Address(),
			nil,
			ocr3impls.TransmitterCombiner,
			nil,
		)
		require.Error(t, err, "expected error creating multichain config tracker")
	})
}

func TestMultichainConfigTracker_SingleChain(t *testing.T) {
	db := pgtest.NewSqlxDB(t)
	lp, uni := setupLogPoller[multichainMeta](t, db, nil)
	require.NoError(t, lp.Start(testutils.Context(t)))
	t.Cleanup(func() { require.NoError(t, lp.Close()) })

	masterChain := relay.ID{
		Network: relay.EVM,
		ChainID: testutils.SimulatedChainID.String(),
	}
	// for this test only one LM is "deployed"
	// so the discovery will return a single LM which is the master LM
	reg := liquiditymanager.NewRegistry()
	reg.Add(models.NetworkSelector(mustStrToI64(t, masterChain.ChainID)), models.Address(uni.wrapper.Address()))
	mockMasterLM := lm_mocks.NewRebalancer(t)
	mockMasterLM.On("Discover", mock.Anything, mock.Anything).Return(reg, liquiditygraph.NewGraph(), nil)
	defer mockMasterLM.AssertExpectations(t)
	mockLMFactory := mocks.NewFactory(t)
	mockLMFactory.On("NewRebalancer", models.NetworkSelector(mustStrToI64(t, masterChain.ChainID)), models.Address(uni.wrapper.Address())).
		Return(mockMasterLM, nil)
	defer mockLMFactory.AssertExpectations(t)
	tracker, err := ocr3impls.NewMultichainConfigTracker(
		masterChain,
		logger.TestLogger(t),
		map[relay.ID]logpoller.LogPoller{masterChain: lp},
		uni.simClient,
		uni.wrapper.Address(),
		mockLMFactory,
		ocr3impls.TransmitterCombiner,
		nil,
	)
	require.NoError(t, err, "failed to create multichain config tracker")

	// Replay the log poller to get the ConfigSet log
	err = tracker.ReplayChain(testutils.Context(t), masterChain, 1)
	require.NoError(t, err, "failed to replay log poller")

	// fetch config digest from the tracker
	changedInBlock, configDigest, err := tracker.LatestConfigDetails(testutils.Context(t))
	require.NoError(t, err, "failed to get latest config details")
	c, err := uni.wrapper.LatestConfigDigestAndEpoch(nil)
	require.NoError(t, err, "failed to get latest config digest and epoch")
	require.Equal(t, hex.EncodeToString(c.ConfigDigest[:]), configDigest.Hex(), "expected latest config digest to match")

	// fetch config details from the tracker
	config, err := tracker.LatestConfig(testutils.Context(t), changedInBlock)
	require.NoError(t, err, "failed to get latest config")
	require.Equal(t, uint64(1), config.ConfigCount, "expected config count to match")
	require.Equal(t, configDigest, config.ConfigDigest, "expected config digest to match")
	require.Equal(t, uint8(1), config.F, "expected f to match")
	require.Equal(t, []byte{}, config.OnchainConfig, "expected onchain config to match")
	require.Equal(t, []byte{}, config.OffchainConfig, "expected offchain config to match")
	require.Equal(t, uint64(3), config.OffchainConfigVersion, "expected offchain config version to match")
	expectedSigners := func() []ocrtypes.OnchainPublicKey {
		var signers []ocrtypes.OnchainPublicKey
		for _, b := range uni.keyrings {
			signers = append(signers, b.PublicKey())
		}
		return signers
	}()
	expectedTransmitters := func() []ocrtypes.Account {
		var accounts []ocrtypes.Account
		for _, tm := range uni.transmitters {
			accounts = append(accounts, ocrtypes.Account(ocr3impls.EncodeTransmitter(masterChain, ocrtypes.Account(tm.From.Hex()))))
		}
		return accounts
	}()
	require.Equal(t, expectedSigners, config.Signers, "expected signers to match")
	require.Equal(t, expectedTransmitters, config.Transmitters, "expected transmitters to match")
}

func TestMultichainConfigTracker_Multichain(t *testing.T) {
	// create heavyweight db's because the log pollers need to have separate
	// databases to avoid conflicts.
	_, db1 := heavyweight.FullTestDBV2(t, nil)
	_, db2 := heavyweight.FullTestDBV2(t, nil)

	lp1, uni1 := setupLogPoller[multichainMeta](t, db1, nil)
	lp2, uni2 := setupLogPoller[multichainMeta](t, db2, &keyringsAndSigners[multichainMeta]{
		keyrings: uni1.keyrings,
		signers:  uni1.signers,
	})
	t.Cleanup(func() {
		require.NoError(t, lp1.Close())
		require.NoError(t, lp2.Close())
	})

	// finality depth
	uni2.backend.Commit()
	uni2.backend.Commit()

	// start the log pollers
	require.NoError(t, lp1.Start(testutils.Context(t)))
	require.NoError(t, lp2.Start(testutils.Context(t)))

	// create the multichain config tracker
	// the chain id's we're using in the mappings are different from the
	// simulated chain id but that should be fine for this test.
	masterChain := relay.ID{
		Network: relay.EVM,
		ChainID: testutils.NewRandomEVMChainID().String(),
	}
	secondChain := relay.ID{
		Network: relay.EVM,
		ChainID: testutils.NewRandomEVMChainID().String(),
	}
	reg := liquiditymanager.NewRegistry()
	reg.Add(models.NetworkSelector(mustStrToI64(t, masterChain.ChainID)), models.Address(uni1.wrapper.Address()))
	reg.Add(models.NetworkSelector(mustStrToI64(t, secondChain.ChainID)), models.Address(uni2.wrapper.Address()))
	mockMasterLM := lm_mocks.NewRebalancer(t)
	mockMasterLM.On("Discover", mock.Anything, mock.Anything).Return(reg, liquiditygraph.NewGraph(), nil)
	defer mockMasterLM.AssertExpectations(t)
	mockLMFactory := mocks.NewFactory(t)
	mockLMFactory.On("NewRebalancer", models.NetworkSelector(mustStrToI64(t, masterChain.ChainID)), models.Address(uni1.wrapper.Address())).
		Return(mockMasterLM, nil)
	defer mockLMFactory.AssertExpectations(t)
	tracker, err := ocr3impls.NewMultichainConfigTracker(
		masterChain,
		logger.TestLogger(t),
		map[relay.ID]logpoller.LogPoller{
			masterChain: lp1,
			secondChain: lp2,
		},
		uni1.simClient,
		uni1.wrapper.Address(),
		mockLMFactory,
		ocr3impls.TransmitterCombiner,
		nil, // we call replay explicitly below
	)
	require.NoError(t, err, "failed to create multichain config tracker")

	// Replay the log pollers to get the ConfigSet log
	// on each respective chain
	require.NoError(t, tracker.ReplayChain(testutils.Context(t), masterChain, 1), "failed to replay log poller on master chain")
	require.NoError(t, tracker.ReplayChain(testutils.Context(t), secondChain, 1), "failed to replay log poller on second chain")

	// fetch config digest from the tracker
	changedInBlock, configDigest, err := tracker.LatestConfigDetails(testutils.Context(t))
	require.NoError(t, err, "failed to get latest config details")
	c, err := uni1.wrapper.LatestConfigDigestAndEpoch(nil)
	require.NoError(t, err, "failed to get latest config digest and epoch")
	require.Equal(t, hex.EncodeToString(c.ConfigDigest[:]), configDigest.Hex(), "expected latest config digest to match")

	// fetch config details from the tracker
	config, err := tracker.LatestConfig(testutils.Context(t), changedInBlock)
	require.NoError(t, err, "failed to get latest config")
	require.Equal(t, uint64(1), config.ConfigCount, "expected config count to match")
	require.Equal(t, configDigest, config.ConfigDigest, "expected config digest to match")
	require.Equal(t, uint8(1), config.F, "expected f to match")
	require.Equal(t, []byte{}, config.OnchainConfig, "expected onchain config to match")
	require.Equal(t, []byte{}, config.OffchainConfig, "expected offchain config to match")
	require.Equal(t, uint64(3), config.OffchainConfigVersion, "expected offchain config version to match")
	expectedSigners := func() []ocrtypes.OnchainPublicKey {
		var signers []ocrtypes.OnchainPublicKey
		for _, b := range uni1.keyrings {
			signers = append(signers, b.PublicKey())
		}
		return signers
	}()
	require.Equal(t, expectedSigners, config.Signers, "expected signers to match")
	expectedTransmitters := func() []ocrtypes.Account {
		var accounts []ocrtypes.Account
		for i := range uni1.transmitters {
			t1 := ocr3impls.EncodeTransmitter(masterChain, ocrtypes.Account(uni1.transmitters[i].From.Hex()))
			t2 := ocr3impls.EncodeTransmitter(secondChain, ocrtypes.Account(uni2.transmitters[i].From.Hex()))
			accounts = append(accounts, ocrtypes.Account(ocr3impls.JoinTransmitters([]string{t1, t2})))
		}
		return accounts
	}()
	require.Equal(t, expectedTransmitters, config.Transmitters, "expected transmitters to match")
}

func mustStrToI64(t *testing.T, s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	require.NoError(t, err)
	return i
}
