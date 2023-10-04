package integrationtesthelpers

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	types3 "github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	types4 "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"k8s.io/utils/pointer"

	"github.com/smartcontractkit/chainlink-relay/pkg/loop"
	ctfClient "github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	v2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	configv2 "github.com/smartcontractkit/chainlink/v2/core/config/toml"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/logger/audit"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/validate"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrbootstrap"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	evmrelay "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type Node struct {
	App             chainlink.Application
	Transmitter     common.Address
	PaymentReceiver common.Address
	KeyBundle       ocr2key.KeyBundle
}

func (node *Node) FindJobIDForContract(t *testing.T, addr common.Address) int32 {
	jobs := node.App.JobSpawner().ActiveJobs()
	for _, j := range jobs {
		if j.Type == job.OffchainReporting2 && j.OCR2OracleSpec.ContractID == addr.Hex() {
			return j.ID
		}
	}
	t.Fatalf("Could not find job for contract %s", addr.Hex())
	return 0
}

func (node *Node) EventuallyNodeUsesUpdatedPriceRegistry(t *testing.T, ccipContracts CCIPIntegrationTestHarness) logpoller.Log {
	c, err := node.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(ccipContracts.Dest.ChainID, 10))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		log, err := c.LogPoller().LatestLogByEventSigWithConfs(
			abihelpers.EventSignatures.UsdPerUnitGasUpdated,
			ccipContracts.Dest.PriceRegistry.Address(),
			0,
			pg.WithParentCtx(testutils.Context(t)),
		)
		// err can be transient errors such as sql row set empty
		if err != nil {
			return false
		}
		return log != nil
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "node is not using updated price registry %s", ccipContracts.Dest.PriceRegistry.Address().Hex())
	return log
}

func (node *Node) EventuallyNodeUsesNewCommitConfig(t *testing.T, ccipContracts CCIPIntegrationTestHarness, commitCfg ccipconfig.CommitOnchainConfig) logpoller.Log {
	c, err := node.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(ccipContracts.Dest.ChainID, 10))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		log, err := c.LogPoller().LatestLogByEventSigWithConfs(
			evmrelay.ConfigSet,
			ccipContracts.Dest.CommitStore.Address(),
			0,
			pg.WithParentCtx(testutils.Context(t)),
		)
		require.NoError(t, err)
		var latestCfg ccipconfig.CommitOnchainConfig
		if log != nil {
			latestCfg, err = DecodeCommitOnChainConfig(log.Data)
			require.NoError(t, err)
			return latestCfg == commitCfg
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "node is using old cfg")
	return log
}

func (node *Node) EventuallyNodeUsesNewExecConfig(t *testing.T, ccipContracts CCIPIntegrationTestHarness, execCfg ccipconfig.ExecOnchainConfig) logpoller.Log {
	c, err := node.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(ccipContracts.Dest.ChainID, 10))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		log, err := c.LogPoller().LatestLogByEventSigWithConfs(
			evmrelay.ConfigSet,
			ccipContracts.Dest.OffRamp.Address(),
			0,
			pg.WithParentCtx(testutils.Context(t)),
		)
		require.NoError(t, err)
		var latestCfg ccipconfig.ExecOnchainConfig
		if log != nil {
			latestCfg, err = DecodeExecOnChainConfig(log.Data)
			require.NoError(t, err)
			return latestCfg == execCfg
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "node is using old cfg")
	return log
}

func (node *Node) EventuallyHasReqSeqNum(t *testing.T, ccipContracts *CCIPIntegrationTestHarness, onRamp common.Address, seqNum int) logpoller.Log {
	c, err := node.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(ccipContracts.Source.ChainID, 10))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		lgs, err := c.LogPoller().LogsDataWordRange(
			ccipdata.CCIPSendRequestEventSigV1_2_0,
			onRamp,
			ccipdata.CCIPSendRequestSeqNumIndexV1_2_0,
			abihelpers.EvmWord(uint64(seqNum)),
			abihelpers.EvmWord(uint64(seqNum)),
			1,
			pg.WithParentCtx(testutils.Context(t)),
		)
		require.NoError(t, err)
		t.Log("Send requested", len(lgs))
		if len(lgs) == 1 {
			log = lgs[0]
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "eventually has seq num")
	return log
}

func (node *Node) EventuallyHasExecutedSeqNums(t *testing.T, ccipContracts *CCIPIntegrationTestHarness, offRamp common.Address, minSeqNum int, maxSeqNum int) []logpoller.Log {
	c, err := node.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(ccipContracts.Dest.ChainID, 10))
	require.NoError(t, err)
	var logs []logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		lgs, err := c.LogPoller().IndexedLogsTopicRange(
			abihelpers.EventSignatures.ExecutionStateChanged,
			offRamp,
			abihelpers.EventSignatures.ExecutionStateChangedSequenceNumberIndex,
			abihelpers.EvmWord(uint64(minSeqNum)),
			abihelpers.EvmWord(uint64(maxSeqNum)),
			1,
			pg.WithParentCtx(testutils.Context(t)),
		)
		require.NoError(t, err)
		t.Logf("Have executed logs %d want %d", len(lgs), maxSeqNum-minSeqNum+1)
		if len(lgs) == maxSeqNum-minSeqNum+1 {
			logs = lgs
			t.Logf("Seq Num %d-%d executed", minSeqNum, maxSeqNum)
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "eventually has not executed seq num")
	return logs
}

func (node *Node) ConsistentlySeqNumHasNotBeenExecuted(t *testing.T, ccipContracts *CCIPIntegrationTestHarness, offRamp common.Address, seqNum int) logpoller.Log {
	c, err := node.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(ccipContracts.Dest.ChainID, 10))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Consistently(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		lgs, err := c.LogPoller().IndexedLogsTopicRange(
			abihelpers.EventSignatures.ExecutionStateChanged,
			offRamp,
			abihelpers.EventSignatures.ExecutionStateChangedSequenceNumberIndex,
			abihelpers.EvmWord(uint64(seqNum)),
			abihelpers.EvmWord(uint64(seqNum)),
			1,
			pg.WithParentCtx(testutils.Context(t)),
		)
		require.NoError(t, err)
		t.Log("Executed logs", lgs)
		if len(lgs) == 1 {
			log = lgs[0]
			return true
		}
		return false
	}, 10*time.Second, 1*time.Second).Should(gomega.BeFalse(), "seq number got executed")
	return log
}

func (node *Node) AddJob(t *testing.T, spec *ctfClient.OCR2TaskJobSpec) {
	specString, err := spec.String()
	require.NoError(t, err)
	ccipJob, err := validate.ValidatedOracleSpecToml(node.App.GetConfig().OCR2(), node.App.GetConfig().Insecure(), specString)
	require.NoError(t, err)
	err = node.App.AddJobV2(context.Background(), &ccipJob)
	require.NoError(t, err)
}

func (node *Node) AddBootstrapJob(t *testing.T, spec *ctfClient.OCR2TaskJobSpec) {
	specString, err := spec.String()
	require.NoError(t, err)
	ccipJob, err := ocrbootstrap.ValidatedBootstrapSpecToml(specString)
	require.NoError(t, err)
	err = node.App.AddJobV2(context.Background(), &ccipJob)
	require.NoError(t, err)
}

func (node *Node) AddJobsWithSpec(t *testing.T, jobSpec *ctfClient.OCR2TaskJobSpec) {
	// set node specific values
	jobSpec.OCR2OracleSpec.OCRKeyBundleID.SetValid(node.KeyBundle.ID())
	jobSpec.OCR2OracleSpec.TransmitterID.SetValid(node.Transmitter.Hex())
	node.AddJob(t, jobSpec)
}

func setupNodeCCIP(
	t *testing.T,
	owner *bind.TransactOpts,
	port int64,
	dbName string,
	sourceChain *backends.SimulatedBackend, destChain *backends.SimulatedBackend,
	sourceChainID *big.Int, destChainID *big.Int,
	bootstrapPeerID string,
	bootstrapPort int64,
) (chainlink.Application, string, common.Address, ocr2key.KeyBundle) {
	trueRef, falseRef := true, false

	// Do not want to load fixtures as they contain a dummy chainID.
	loglevel := configv2.LogLevel(zap.DebugLevel)
	config, db := heavyweight.FullTestDBNoFixturesV2(t, fmt.Sprintf("%s%d", dbName, port), func(c *chainlink.Config, _ *chainlink.Secrets) {
		p2pAddresses := []string{
			fmt.Sprintf("127.0.0.1:%d", port),
		}
		c.Log.Level = &loglevel
		c.Feature.CCIP = &trueRef
		c.Feature.UICSAKeys = &trueRef
		c.OCR.Enabled = &falseRef
		c.OCR.DefaultTransactionQueueDepth = pointer.Uint32(200)
		c.OCR2.Enabled = &trueRef
		c.Feature.LogPoller = &trueRef
		c.P2P.V1.Enabled = &falseRef
		c.P2P.V2.Enabled = &trueRef
		c.P2P.V2.DeltaDial = models.MustNewDuration(500 * time.Millisecond)
		c.P2P.V2.DeltaReconcile = models.MustNewDuration(5 * time.Second)
		c.P2P.V2.ListenAddresses = &p2pAddresses
		c.P2P.V2.AnnounceAddresses = &p2pAddresses

		c.EVM = []*v2.EVMConfig{createConfigV2Chain(sourceChainID), createConfigV2Chain(destChainID)}

		if bootstrapPeerID != "" {
			// Supply the bootstrap IP and port as a V2 peer address
			c.P2P.V2.DefaultBootstrappers = &[]commontypes.BootstrapperLocator{
				{
					PeerID: bootstrapPeerID, Addrs: []string{
						fmt.Sprintf("127.0.0.1:%d", bootstrapPort),
					},
				},
			}
		}
	})

	lggr := logger.TestLogger(t)

	eventBroadcaster := pg.NewEventBroadcaster(config.Database().URL(), 0, 0, lggr, uuid.New())

	// The in-memory geth sim does not let you create a custom ChainID, it will always be 1337.
	// In particular this means that if you sign an eip155 tx, the chainID used MUST be 1337
	// and the CHAINID op code will always emit 1337. To work around this to simulate a "multichain"
	// test, we fake different chainIDs using the wrapped sim cltest.SimulatedBackend so the RPC
	// appears to operate on different chainIDs and we use an EthKeyStoreSim wrapper which always
	// signs 1337 see https://github.com/smartcontractkit/chainlink-ccip/blob/a24dd436810250a458d27d8bb3fb78096afeb79c/core/services/ocr2/plugins/ccip/testhelpers/simulated_backend.go#L35
	sourceClient := client.NewSimulatedBackendClient(t, sourceChain, sourceChainID)
	destClient := client.NewSimulatedBackendClient(t, destChain, destChainID)
	keyStore := keystore.New(db, utils.FastScryptParams, lggr, config.Database())
	simEthKeyStore := testhelpers.EthKeyStoreSim{
		ETHKS: keyStore.Eth(),
		CSAKS: keyStore.CSA(),
	}
	mailMon := utils.NewMailboxMonitor("CCIP")
	evmOpts := chainlink.EVMFactoryConfig{
		ChainOpts: evm.ChainOpts{
			AppConfig:        config,
			EventBroadcaster: eventBroadcaster,
			GenEthClient: func(chainID *big.Int) client.Client {
				if chainID.String() == sourceChainID.String() {
					return sourceClient
				} else if chainID.String() == destChainID.String() {
					return destClient
				}
				t.Fatalf("invalid chain ID %v", chainID.String())
				return nil
			},
			MailMon: mailMon,
			DB:      db,
		},
		CSAETHKeystore: simEthKeyStore,
	}
	loopRegistry := plugins.NewLoopRegistry(lggr.Named("LoopRegistry"))
	relayerFactory := chainlink.RelayerFactory{
		Logger:       lggr,
		LoopRegistry: loopRegistry,
		GRPCOpts:     loop.GRPCOpts{},
	}
	testCtx := testutils.Context(t)
	// evm alway enabled for backward compatibility
	initOps := []chainlink.CoreRelayerChainInitFunc{chainlink.InitEVM(testCtx, relayerFactory, evmOpts)}

	relayChainInterops, err := chainlink.NewCoreRelayerChainInteroperators(initOps...)
	if err != nil {
		t.Fatal(err)
	}

	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:                     config,
		EventBroadcaster:           eventBroadcaster,
		SqlxDB:                     db,
		KeyStore:                   keyStore,
		RelayerChainInteroperators: relayChainInterops,
		Logger:                     lggr,
		ExternalInitiatorManager:   nil,
		CloseLogger:                lggr.Sync,
		UnrestrictedHTTPClient:     &http.Client{},
		RestrictedHTTPClient:       &http.Client{},
		AuditLogger:                audit.NoopLogger,
		MailMon:                    mailMon,
		LoopRegistry:               plugins.NewLoopRegistry(lggr),
	})
	require.NoError(t, err)
	require.NoError(t, app.GetKeyStore().Unlock("password"))
	_, err = app.GetKeyStore().P2P().Create()
	require.NoError(t, err)

	p2pIDs, err := app.GetKeyStore().P2P().GetAll()
	require.NoError(t, err)
	require.Len(t, p2pIDs, 1)
	peerID := p2pIDs[0].PeerID()

	_, err = app.GetKeyStore().Eth().Create(destChainID)
	require.NoError(t, err)
	sendingKeys, err := app.GetKeyStore().Eth().EnabledKeysForChain(destChainID)
	require.NoError(t, err)
	require.Len(t, sendingKeys, 1)
	transmitter := sendingKeys[0].Address
	s, err := app.GetKeyStore().Eth().GetState(sendingKeys[0].ID(), destChainID)
	require.NoError(t, err)
	lggr.Debug(fmt.Sprintf("Transmitter address %s chainID %s", transmitter, s.EVMChainID.String()))

	// Fund the commitTransmitter address with some ETH
	n, err := destChain.NonceAt(context.Background(), owner.From, nil)
	require.NoError(t, err)

	tx := types3.NewTransaction(n, transmitter, big.NewInt(1000000000000000000), 21000, big.NewInt(1000000000), nil)
	signedTx, err := owner.Signer(owner.From, tx)
	require.NoError(t, err)
	err = destChain.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
	destChain.Commit()

	kb, err := app.GetKeyStore().OCR2().Create(chaintype.EVM)
	require.NoError(t, err)
	return app, peerID.Raw(), transmitter, kb
}

func createConfigV2Chain(chainId *big.Int) *v2.EVMConfig {
	// NOTE: For the executor jobs, the default of 500k is insufficient for a 3 message batch
	defaultGasLimit := uint32(5000000)
	tr := true

	sourceC := v2.Defaults((*utils.Big)(chainId))
	sourceC.GasEstimator.LimitDefault = &defaultGasLimit
	fixedPrice := "FixedPrice"
	sourceC.GasEstimator.Mode = &fixedPrice
	d, _ := models.MakeDuration(100 * time.Millisecond)
	sourceC.LogPollInterval = &d
	fd := uint32(2)
	sourceC.FinalityDepth = &fd
	return &v2.EVMConfig{
		ChainID: (*utils.Big)(chainId),
		Enabled: &tr,
		Chain:   sourceC,
		Nodes:   v2.EVMNodes{&v2.Node{}},
	}
}

type CCIPIntegrationTestHarness struct {
	testhelpers.CCIPContracts
	Nodes     []Node
	Bootstrap Node
}

func SetupCCIPIntegrationTH(t *testing.T, sourceChainID, sourceChainSelector, destChainId, destChainSelector uint64) CCIPIntegrationTestHarness {
	c := testhelpers.SetupCCIPContracts(t, sourceChainID, sourceChainSelector, destChainId, destChainSelector)
	return CCIPIntegrationTestHarness{
		CCIPContracts: c,
	}
}

func (c *CCIPIntegrationTestHarness) AddAllJobs(t *testing.T, jobParams CCIPJobSpecParams) {
	jobParams.OffRamp = c.Dest.OffRamp.Address()

	commitSpec, err := jobParams.CommitJobSpec()
	require.NoError(t, err)
	geExecutionSpec, err := jobParams.ExecutionJobSpec()
	require.NoError(t, err)
	nodes := c.Nodes
	for _, node := range nodes {
		node.AddJobsWithSpec(t, commitSpec)
		node.AddJobsWithSpec(t, geExecutionSpec)
	}
}

func (c *CCIPIntegrationTestHarness) AllNodesHaveReqSeqNum(t *testing.T, seqNum int, onRampOpts ...common.Address) logpoller.Log {
	var log logpoller.Log
	nodes := c.Nodes
	var onRamp common.Address
	if len(onRampOpts) > 0 {
		onRamp = onRampOpts[0]
	} else {
		require.NotNil(t, c.Source.OnRamp, "no onramp configured")
		onRamp = c.Source.OnRamp.Address()
	}
	for _, node := range nodes {
		log = node.EventuallyHasReqSeqNum(t, c, onRamp, seqNum)
	}
	return log
}

func (c *CCIPIntegrationTestHarness) AllNodesHaveExecutedSeqNums(t *testing.T, minSeqNum int, maxSeqNum int, offRampOpts ...common.Address) []logpoller.Log {
	var logs []logpoller.Log
	nodes := c.Nodes
	var offRamp common.Address

	if len(offRampOpts) > 0 {
		offRamp = offRampOpts[0]
	} else {
		require.NotNil(t, c.Dest.OffRamp, "no offramp configured")
		offRamp = c.Dest.OffRamp.Address()
	}
	for _, node := range nodes {
		logs = node.EventuallyHasExecutedSeqNums(t, c, offRamp, minSeqNum, maxSeqNum)
	}
	return logs
}

func (c *CCIPIntegrationTestHarness) NoNodesHaveExecutedSeqNum(t *testing.T, seqNum int, offRampOpts ...common.Address) logpoller.Log {
	var log logpoller.Log
	nodes := c.Nodes
	var offRamp common.Address
	if len(offRampOpts) > 0 {
		offRamp = offRampOpts[0]
	} else {
		require.NotNil(t, c.Dest.OffRamp, "no offramp configured")
		offRamp = c.Dest.OffRamp.Address()
	}
	for _, node := range nodes {
		log = node.ConsistentlySeqNumHasNotBeenExecuted(t, c, offRamp, seqNum)
	}
	return log
}

func (c *CCIPIntegrationTestHarness) EventuallyCommitReportAccepted(t *testing.T, currentBlock uint64, commitStoreOpts ...common.Address) commit_store.CommitStoreCommitReport {
	var commitStore *commit_store.CommitStore
	var err error
	if len(commitStoreOpts) > 0 {
		commitStore, err = commit_store.NewCommitStore(commitStoreOpts[0], c.Dest.Chain)
		require.NoError(t, err)
	} else {
		require.NotNil(t, c.Dest.CommitStore, "no commitStore configured")
		commitStore = c.Dest.CommitStore
	}
	g := gomega.NewGomegaWithT(t)
	var report commit_store.CommitStoreCommitReport
	g.Eventually(func() bool {
		it, err := commitStore.FilterReportAccepted(&bind.FilterOpts{Start: currentBlock})
		g.Expect(err).NotTo(gomega.HaveOccurred(), "Error filtering ReportAccepted event")
		g.Expect(it.Next()).To(gomega.BeTrue(), "No ReportAccepted event found")
		report = it.Event.Report
		if report.MerkleRoot != [32]byte{} {
			t.Log("Report Accepted by commitStore")
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "report has not been committed")
	return report
}

func (c *CCIPIntegrationTestHarness) EventuallyExecutionStateChangedToSuccess(t *testing.T, seqNum []uint64, blockNum uint64, offRampOpts ...common.Address) {
	var offRamp *evm_2_evm_offramp.EVM2EVMOffRamp
	var err error
	if len(offRampOpts) > 0 {
		offRamp, err = evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampOpts[0], c.Dest.Chain)
		require.NoError(t, err)
	} else {
		require.NotNil(t, c.Dest.OffRamp, "no offRamp configured")
		offRamp = c.Dest.OffRamp
	}
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		it, err := offRamp.FilterExecutionStateChanged(&bind.FilterOpts{Start: blockNum}, seqNum, [][32]byte{})
		require.NoError(t, err)
		for it.Next() {
			if abihelpers.MessageExecutionState(it.Event.State) == abihelpers.ExecutionStateSuccess {
				t.Logf("ExecutionStateChanged event found for seqNum %d", it.Event.SequenceNumber)
				return true
			}
		}
		c.Source.Chain.Commit()
		c.Dest.Chain.Commit()
		return false
	}, testutils.WaitTimeout(t), time.Second).
		Should(gomega.BeTrue(), "ExecutionStateChanged Event")
}

func (c *CCIPIntegrationTestHarness) EventuallyReportCommitted(t *testing.T, max int, commitStoreOpts ...common.Address) uint64 {
	var commitStore *commit_store.CommitStore
	var err error
	var committedSeqNum uint64
	if len(commitStoreOpts) > 0 {
		commitStore, err = commit_store.NewCommitStore(commitStoreOpts[0], c.Dest.Chain)
		require.NoError(t, err)
	} else {
		require.NotNil(t, c.Dest.CommitStore, "no commitStore configured")
		commitStore = c.Dest.CommitStore
	}
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		minSeqNum, err := commitStore.GetExpectedNextSequenceNumber(nil)
		require.NoError(t, err)
		c.Source.Chain.Commit()
		c.Dest.Chain.Commit()
		t.Log("next expected seq num reported", minSeqNum)
		committedSeqNum = minSeqNum
		return minSeqNum > uint64(max)
	}, testutils.WaitTimeout(t), time.Second).Should(gomega.BeTrue(), "report has not been committed")
	return committedSeqNum
}

func (c *CCIPIntegrationTestHarness) EventuallySendRequested(t *testing.T, seqNum uint64, onRampOpts ...common.Address) {
	var onRamp *evm_2_evm_onramp.EVM2EVMOnRamp
	var err error
	if len(onRampOpts) > 0 {
		onRamp, err = evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampOpts[0], c.Source.Chain)
		require.NoError(t, err)
	} else {
		require.NotNil(t, c.Source.OnRamp, "no onRamp configured")
		onRamp = c.Source.OnRamp
	}
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		it, err := onRamp.FilterCCIPSendRequested(nil)
		require.NoError(t, err)
		for it.Next() {
			if it.Event.Message.SequenceNumber == seqNum {
				t.Log("sendRequested generated for", seqNum)
				return true
			}
		}
		c.Source.Chain.Commit()
		c.Dest.Chain.Commit()
		return false
	}, testutils.WaitTimeout(t), time.Second).Should(gomega.BeTrue(), "sendRequested has not been generated")
}

func (c *CCIPIntegrationTestHarness) ConsistentlyReportNotCommitted(t *testing.T, max int, commitStoreOpts ...common.Address) {
	var commitStore *commit_store.CommitStore
	var err error
	if len(commitStoreOpts) > 0 {
		commitStore, err = commit_store.NewCommitStore(commitStoreOpts[0], c.Dest.Chain)
		require.NoError(t, err)
	} else {
		require.NotNil(t, c.Dest.CommitStore, "no commitStore configured")
		commitStore = c.Dest.CommitStore
	}
	gomega.NewGomegaWithT(t).Consistently(func() bool {
		minSeqNum, err := commitStore.GetExpectedNextSequenceNumber(nil)
		require.NoError(t, err)
		c.Source.Chain.Commit()
		c.Dest.Chain.Commit()
		t.Log("min seq num reported", minSeqNum)
		return minSeqNum > uint64(max)
	}, testutils.WaitTimeout(t), time.Second).Should(gomega.BeFalse(), "report has been committed")
}

func (c *CCIPIntegrationTestHarness) SetupAndStartNodes(ctx context.Context, t *testing.T, bootstrapNodePort int64) (Node, []Node, int64) {
	appBootstrap, bootstrapPeerID, bootstrapTransmitter, bootstrapKb := setupNodeCCIP(t, c.Dest.User, bootstrapNodePort,
		"bootstrap_ccip", c.Source.Chain, c.Dest.Chain, big.NewInt(0).SetUint64(c.Source.ChainID),
		big.NewInt(0).SetUint64(c.Dest.ChainID), "", 0)
	var (
		oracles []confighelper.OracleIdentityExtra
		nodes   []Node
	)
	err := appBootstrap.Start(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, appBootstrap.Stop())
	})
	bootstrapNode := Node{
		App:         appBootstrap,
		Transmitter: bootstrapTransmitter,
		KeyBundle:   bootstrapKb,
	}
	// Set up the minimum 4 oracles all funded with destination ETH
	for i := int64(0); i < 4; i++ {
		app, peerID, transmitter, kb := setupNodeCCIP(
			t,
			c.Dest.User,
			bootstrapNodePort+1+i,
			fmt.Sprintf("oracle_ccip%d", i),
			c.Source.Chain,
			c.Dest.Chain,
			big.NewInt(0).SetUint64(c.Source.ChainID),
			big.NewInt(0).SetUint64(c.Dest.ChainID),
			bootstrapPeerID,
			bootstrapNodePort,
		)
		nodes = append(nodes, Node{
			App:         app,
			Transmitter: transmitter,
			KeyBundle:   kb,
		})
		offchainPublicKey, _ := hex.DecodeString(strings.TrimPrefix(kb.OnChainPublicKey(), "0x"))
		oracles = append(oracles, confighelper.OracleIdentityExtra{
			OracleIdentity: confighelper.OracleIdentity{
				OnchainPublicKey:  offchainPublicKey,
				TransmitAccount:   types4.Account(transmitter.String()),
				OffchainPublicKey: kb.OffchainPublicKey(),
				PeerID:            peerID,
			},
			ConfigEncryptionPublicKey: kb.ConfigEncryptionPublicKey(),
		})
		err = app.Start(ctx)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, app.Stop())
		})
	}

	c.Oracles = oracles
	commitOnchainConfig := c.CreateDefaultCommitOnchainConfig(t)
	commitOffchainConfig := c.CreateDefaultCommitOffchainConfig(t)
	execOnchainConfig := c.CreateDefaultExecOnchainConfig(t)
	execOffchainConfig := c.CreateDefaultExecOffchainConfig(t)

	configBlock := c.SetupOnchainConfig(t, commitOnchainConfig, commitOffchainConfig, execOnchainConfig, execOffchainConfig)
	c.Nodes = nodes
	c.Bootstrap = bootstrapNode
	return bootstrapNode, nodes, configBlock
}

func (c *CCIPIntegrationTestHarness) SetUpNodesAndJobs(t *testing.T, pricePipeline string, bootstrapNodePort int64) CCIPJobSpecParams {
	// setup Jobs
	ctx := context.Background()
	// Starts nodes and configures them in the OCR contracts.
	bootstrapNode, _, configBlock := c.SetupAndStartNodes(ctx, t, bootstrapNodePort)

	jobParams := c.NewCCIPJobSpecParams(pricePipeline, configBlock)

	// Add the bootstrap job
	c.Bootstrap.AddBootstrapJob(t, jobParams.BootstrapJob(c.Dest.CommitStore.Address().Hex()))
	c.AddAllJobs(t, jobParams)

	// Replay for bootstrap.
	bc, err := bootstrapNode.App.GetRelayers().LegacyEVMChains().Get(strconv.FormatUint(c.Dest.ChainID, 10))
	require.NoError(t, err)
	require.NoError(t, bc.LogPoller().Replay(context.Background(), configBlock))
	c.Dest.Chain.Commit()

	return jobParams
}
func DecodeCommitOnChainConfig(encoded []byte) (ccipconfig.CommitOnchainConfig, error) {
	var onchainConfig ccipconfig.CommitOnchainConfig
	unpacked, err := abihelpers.DecodeOCR2Config(encoded)
	if err != nil {
		return onchainConfig, err
	}
	onChainCfg := unpacked.OnchainConfig
	onchainConfig, err = abihelpers.DecodeAbiStruct[ccipconfig.CommitOnchainConfig](onChainCfg)
	if err != nil {
		return onchainConfig, err
	}
	return onchainConfig, nil
}

func DecodeExecOnChainConfig(encoded []byte) (ccipconfig.ExecOnchainConfig, error) {
	var onchainConfig ccipconfig.ExecOnchainConfig
	unpacked, err := abihelpers.DecodeOCR2Config(encoded)
	if err != nil {
		return onchainConfig, errors.Wrap(err, "failed to unpack log data")
	}
	onChainCfg := unpacked.OnchainConfig
	onchainConfig, err = abihelpers.DecodeAbiStruct[ccipconfig.ExecOnchainConfig](onChainCfg)
	if err != nil {
		return onchainConfig, err
	}
	return onchainConfig, nil
}
