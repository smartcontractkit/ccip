package testhelpers

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	types3 "github.com/ethereum/go-ethereum/core/types"
	"github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	types4 "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	ctfClient "github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/client"
	v2 "github.com/smartcontractkit/chainlink/core/chains/evm/config/v2"
	"github.com/smartcontractkit/chainlink/core/chains/evm/headtracker"
	types2 "github.com/smartcontractkit/chainlink/core/chains/evm/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/chains/evm/log"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"
	configv2 "github.com/smartcontractkit/chainlink/core/config/v2"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/logger/audit"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/validate"
	"github.com/smartcontractkit/chainlink/core/services/ocrbootstrap"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
)

type Node struct {
	App         chainlink.Application
	Transmitter common.Address
	KeyBundle   ocr2key.KeyBundle
}

func (node *Node) EventuallyHasReqSeqNum(t *testing.T, ccipContracts CCIPContracts, eventSignatures ccip.EventSignatures, onRamp common.Address, seqNum int) logpoller.Log {
	c, err := node.App.GetChains().EVM.Get(big.NewInt(0).SetUint64(ccipContracts.Source.ChainID))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		lgs, err := c.LogPoller().LogsDataWordRange(eventSignatures.SendRequested, onRamp, eventSignatures.SendRequestedSequenceNumberIndex, ccip.EvmWord(uint64(seqNum)), ccip.EvmWord(uint64(seqNum)), 1)
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

func (node *Node) EventuallyHasExecutedSeqNums(t *testing.T, ccipContracts CCIPContracts, eventSignatures ccip.EventSignatures, offRamp common.Address, minSeqNum int, maxSeqNum int) []logpoller.Log {
	c, err := node.App.GetChains().EVM.Get(big.NewInt(0).SetUint64(ccipContracts.Dest.ChainID))
	require.NoError(t, err)
	var logs []logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		lgs, err := c.LogPoller().IndexedLogsTopicRange(
			eventSignatures.ExecutionStateChanged,
			offRamp,
			eventSignatures.ExecutionStateChangedSequenceNumberIndex,
			ccip.EvmWord(uint64(minSeqNum)),
			ccip.EvmWord(uint64(maxSeqNum)),
			1)
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

func (node *Node) ConsistentlySeqNumHasNotBeenExecuted(t *testing.T, ccipContracts CCIPContracts, eventSignatures ccip.EventSignatures, offRamp common.Address, seqNum int) logpoller.Log {
	c, err := node.App.GetChains().EVM.Get(big.NewInt(0).SetUint64(ccipContracts.Dest.ChainID))
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Consistently(func() bool {
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		lgs, err := c.LogPoller().IndexedLogsTopicRange(
			eventSignatures.ExecutionStateChanged,
			offRamp,
			eventSignatures.ExecutionStateChangedSequenceNumberIndex,
			ccip.EvmWord(uint64(seqNum)),
			ccip.EvmWord(uint64(seqNum)),
			1)
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
	ccipJob, err := validate.ValidatedOracleSpecToml(node.App.GetConfig(), specString)
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

func AddAllJobs(t *testing.T, jobParams CCIPJobSpecParams, ccipContracts CCIPContracts, nodes []Node) {
	jobParams.OffRamp = ccipContracts.Dest.OffRamp.Address()
	jobParams.OnRamp = ccipContracts.Source.OnRamp.Address()

	commitSpec, err := jobParams.CommitJobSpec()
	require.NoError(t, err)
	geExecutionSpec, err := jobParams.ExecutionJobSpec()
	require.NoError(t, err)

	for i, node := range nodes {
		commitSpec.Name = fmt.Sprintf("ccip-commit-%d", i)
		node.AddJobsWithSpec(t, commitSpec)

		geExecutionSpec.Name = fmt.Sprintf("ccip-exec-ge-%d", i)
		node.AddJobsWithSpec(t, geExecutionSpec)
	}
}

func (node *Node) AddJobsWithSpec(t *testing.T, jobSpec *ctfClient.OCR2TaskJobSpec) {
	// set node specific values
	jobSpec.OCR2OracleSpec.OCRKeyBundleID.SetValid(node.KeyBundle.ID())
	jobSpec.OCR2OracleSpec.TransmitterID.SetValid(node.Transmitter.Hex())
	node.AddJob(t, jobSpec)
}

func SetupNodeCCIP(
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
	config, db := heavyweight.FullTestDBNoFixturesV2(t, fmt.Sprintf("%s%d", dbName, port), func(c *chainlink.Config, s *chainlink.Secrets) {
		p2pAddresses := []string{
			fmt.Sprintf("127.0.0.1:%d", port),
		}
		// Disables ocr spec validation, so we can have fast polling for the test.
		c.DevMode = trueRef
		c.Log.Level = &loglevel
		c.Feature.CCIP = &trueRef
		c.OCR.Enabled = &falseRef
		c.OCR2.Enabled = &trueRef
		c.Feature.LogPoller = &trueRef
		c.P2P.V2.Enabled = &trueRef
		c.P2P.V2.DeltaDial = models.MustNewDuration(500 * time.Millisecond)
		c.P2P.V2.DeltaReconcile = models.MustNewDuration(5 * time.Second)
		c.P2P.V2.ListenAddresses = &p2pAddresses
		c.P2P.V2.AnnounceAddresses = &p2pAddresses

		c.EVM = []*v2.EVMConfig{createConfigV2Chain(t, sourceChainID), createConfigV2Chain(t, destChainID)}

		if bootstrapPeerID != "" {
			// Supply the bootstrap IP and port as a V2 peer address
			c.P2P.V2.DefaultBootstrappers = &[]commontypes.BootstrapperLocator{{
				PeerID: bootstrapPeerID, Addrs: []string{
					fmt.Sprintf("127.0.0.1:%d", bootstrapPort),
				}},
			}
		}
	})

	var lggr = logger.TestLogger(t)

	eventBroadcaster := pg.NewEventBroadcaster(config.DatabaseURL(), 0, 0, lggr, uuid.NewV1())

	// We fake different chainIDs using the wrapped sim cltest.SimulatedBackend
	chainORM := evm.NewORM(db, lggr, config)
	err := chainORM.EnsureChains([]utils.Big{*utils.NewBig(sourceChainID), *utils.NewBig(destChainID)})
	require.NoError(t, err)
	sourceClient := client.NewSimulatedBackendClient(t, sourceChain, sourceChainID)
	destClient := client.NewSimulatedBackendClient(t, destChain, destChainID)
	keyStore := keystore.New(db, utils.FastScryptParams, lggr, config)
	simEthKeyStore := EthKeyStoreSim{Eth: keyStore.Eth()}

	evmCfg := evmtest.NewChainScopedConfig(t, config)

	// Create our chainset manually, so we can have custom eth clients
	// (the wrapped sims faking different chainIDs)
	var (
		sourceLp logpoller.LogPoller = logpoller.NewLogPoller(logpoller.NewORM(sourceChainID, db, lggr, config), sourceClient,
			lggr, 500*time.Millisecond, 10, 2, 2, int64(evmCfg.EvmLogKeepBlocksDepth()))
		destLp logpoller.LogPoller = logpoller.NewLogPoller(logpoller.NewORM(destChainID, db, lggr, config), destClient,
			lggr, 500*time.Millisecond, 10, 2, 2, int64(evmCfg.EvmLogKeepBlocksDepth()))
	)

	mailMon := utils.NewMailboxMonitor("CCIP")

	evmChain, err := evm.NewTOMLChainSet(context.Background(), evm.ChainSetOpts{
		Config:           config,
		Logger:           lggr,
		DB:               db,
		KeyStore:         simEthKeyStore,
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
		GenHeadTracker: func(chainID *big.Int, hb types2.HeadBroadcaster) types2.HeadTracker {
			if chainID.String() == sourceChainID.String() {
				return headtracker.NewHeadTracker(
					lggr,
					sourceClient,
					evmtest.NewChainScopedConfig(t, config),
					hb,
					headtracker.NewHeadSaver(lggr, headtracker.NewORM(db, lggr, pgtest.NewQConfig(falseRef), *sourceClient.ChainID()), evmCfg),
					mailMon,
				)
			} else if chainID.String() == destChainID.String() {
				return headtracker.NewHeadTracker(
					lggr,
					destClient,
					evmtest.NewChainScopedConfig(t, config),
					hb,
					headtracker.NewHeadSaver(lggr, headtracker.NewORM(db, lggr, pgtest.NewQConfig(falseRef), *destClient.ChainID()), evmCfg),
					mailMon,
				)
			}
			t.Fatalf("invalid chain ID %v", chainID.String())
			return nil
		},
		GenLogPoller: func(chainID *big.Int) logpoller.LogPoller {
			if chainID.String() == sourceChainID.String() {
				t.Log("Generating log broadcaster source")
				return sourceLp
			} else if chainID.String() == destChainID.String() {
				return destLp
			}
			t.Fatalf("invalid chain ID %v", chainID.String())
			return nil
		},
		GenLogBroadcaster: func(chainID *big.Int) log.Broadcaster {
			if chainID.String() == sourceChainID.String() {
				t.Log("Generating log broadcaster source")
				return log.NewBroadcaster(
					log.NewORM(db, lggr, config, *sourceChainID),
					sourceClient,
					evmtest.NewChainScopedConfig(t, config),
					lggr,
					nil, mailMon)
			} else if chainID.String() == destChainID.String() {
				return log.NewBroadcaster(
					log.NewORM(db, lggr, config, *destChainID),
					destClient,
					evmtest.NewChainScopedConfig(t, config),
					lggr,
					nil,
					mailMon)
			}
			t.Fatalf("invalid chain ID %v", chainID.String())
			return nil
		},
		GenTxManager: func(chainID *big.Int) txmgr.TxManager {
			if chainID.String() == sourceChainID.String() {
				return txmgr.NewTxm(db, sourceClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr, &txmgr.CheckerFactory{Client: sourceClient}, sourceLp)
			} else if chainID.String() == destChainID.String() {
				return txmgr.NewTxm(db, destClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr, &txmgr.CheckerFactory{Client: destClient}, destLp)
			}
			t.Fatalf("invalid chain ID %v", chainID.String())
			return nil
		},
		MailMon: mailMon,
	})
	if err != nil {
		lggr.Fatal(err)
	}

	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:           config,
		EventBroadcaster: eventBroadcaster,
		SqlxDB:           db,
		KeyStore:         keyStore,
		Chains: chainlink.Chains{
			EVM: evmChain,
		},
		Logger:                   lggr,
		ExternalInitiatorManager: nil,
		CloseLogger: func() error {
			return nil
		},
		UnrestrictedHTTPClient: &http.Client{},
		RestrictedHTTPClient:   &http.Client{},
		AuditLogger:            audit.NoopLogger,
		MailMon:                mailMon,
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

func createConfigV2Chain(t *testing.T, chainId *big.Int) *v2.EVMConfig {
	// NOTE: For the executor jobs, the default of 500k is insufficient for a 3 message batch
	defaultGasLimit := uint32(5000000)
	tr := true

	sourceC := v2.Defaults((*utils.Big)(chainId))
	sourceC.GasEstimator.LimitDefault = &defaultGasLimit
	fixedPrice := "FixedPrice"
	sourceC.GasEstimator.Mode = &fixedPrice
	return &v2.EVMConfig{
		ChainID: (*utils.Big)(chainId),
		Enabled: &tr,
		Chain:   sourceC,
		Nodes:   v2.EVMNodes{&v2.Node{}},
	}
}

func AllNodesHaveReqSeqNum(t *testing.T, ccipContracts CCIPContracts, eventSignatures ccip.EventSignatures, onRamp common.Address, nodes []Node, seqNum int) logpoller.Log {
	var log logpoller.Log
	for _, node := range nodes {
		log = node.EventuallyHasReqSeqNum(t, ccipContracts, eventSignatures, onRamp, seqNum)
	}
	return log
}

func AllNodesHaveExecutedSeqNums(t *testing.T, ccipContracts CCIPContracts, eventSignatures ccip.EventSignatures, offRamp common.Address, nodes []Node, minSeqNum int, maxSeqNum int) []logpoller.Log {
	var logs []logpoller.Log
	for _, node := range nodes {
		logs = node.EventuallyHasExecutedSeqNums(t, ccipContracts, eventSignatures, offRamp, minSeqNum, maxSeqNum)
	}
	return logs
}

func NoNodesHaveExecutedSeqNum(t *testing.T, ccipContracts CCIPContracts, eventSignatures ccip.EventSignatures, offRamp common.Address, nodes []Node, seqNum int) logpoller.Log {
	var log logpoller.Log
	for _, node := range nodes {
		log = node.ConsistentlySeqNumHasNotBeenExecuted(t, ccipContracts, eventSignatures, offRamp, seqNum)
	}
	return log
}

func EventuallyCommitReportAccepted(t *testing.T, ccipContracts CCIPContracts, currentBlock uint64) commit_store.ICommitStoreCommitReport {
	g := gomega.NewGomegaWithT(t)
	var report commit_store.ICommitStoreCommitReport
	g.Eventually(func() bool {
		it, err := ccipContracts.Dest.CommitStore.FilterReportAccepted(&bind.FilterOpts{Start: currentBlock})
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

func SetupAndStartNodes(ctx context.Context, t *testing.T, ccipContracts *CCIPContracts, bootstrapNodePort int64) (Node, []Node, int64) {
	appBootstrap, bootstrapPeerID, bootstrapTransmitter, bootstrapKb := SetupNodeCCIP(t, ccipContracts.Dest.User, bootstrapNodePort,
		"bootstrap_ccip", ccipContracts.Source.Chain, ccipContracts.Dest.Chain, big.NewInt(0).SetUint64(ccipContracts.Source.ChainID),
		big.NewInt(0).SetUint64(ccipContracts.Dest.ChainID), "", 0)
	var (
		oracles []confighelper.OracleIdentityExtra
		nodes   []Node
	)
	err := appBootstrap.Start(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		appBootstrap.Stop()
	})
	bootstrapNode := Node{
		appBootstrap, bootstrapTransmitter, bootstrapKb,
	}
	// Set up the minimum 4 oracles all funded with destination ETH
	for i := int64(0); i < 4; i++ {
		app, peerID, transmitter, kb := SetupNodeCCIP(
			t,
			ccipContracts.Dest.User,
			bootstrapNodePort+1+i,
			fmt.Sprintf("oracle_ccip%d", i),
			ccipContracts.Source.Chain,
			ccipContracts.Dest.Chain,
			big.NewInt(0).SetUint64(ccipContracts.Source.ChainID),
			big.NewInt(0).SetUint64(ccipContracts.Dest.ChainID),
			bootstrapPeerID,
			bootstrapNodePort,
		)
		nodes = append(nodes, Node{app, transmitter, kb})
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
			app.Stop()
		})
	}

	reportingPluginConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   1,
	}.Encode()
	require.NoError(t, err)
	configBlock := ccipContracts.SetupOnchainConfig(oracles, reportingPluginConfig)
	return bootstrapNode, nodes, configBlock
}
