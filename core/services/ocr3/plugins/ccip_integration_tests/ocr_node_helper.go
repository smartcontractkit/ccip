package ccip_integration_tests

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jmoiron/sqlx"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	v2toml "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"
	encodeutils "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	evmutils "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/logger/audit"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/plugins"
	"github.com/smartcontractkit/libocr/commontypes"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3confighelper"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

type ocr3Node struct {
	app          chainlink.Application
	peerID       string
	transmitters map[uint64]common.Address
	keybundle    ocr2key.KeyBundle
	db           *sqlx.DB
}

// setupNodeOCR3 creates a chainlink node and any associated keys in order to run
// ccip.
func setupNodeOCR3(
	t *testing.T,
	port int,
	p2pV2Bootstrappers []commontypes.BootstrapperLocator,
	universe map[uint64]onchainUniverse,
) *ocr3Node {
	// Do not want to load fixtures as they contain a dummy chainID.
	config, db := heavyweight.FullTestDBNoFixturesV2(t, func(c *chainlink.Config, s *chainlink.Secrets) {
		c.Insecure.OCRDevelopmentMode = ptr(true) // Disables ocr spec validation so we can have fast polling for the test.

		c.Feature.LogPoller = ptr(true)

		// P2P V2 configs.
		c.P2P.V2.Enabled = ptr(true)
		c.P2P.V2.DeltaDial = config.MustNewDuration(500 * time.Millisecond)
		c.P2P.V2.DeltaReconcile = config.MustNewDuration(5 * time.Second)
		c.P2P.V2.ListenAddresses = &[]string{fmt.Sprintf("127.0.0.1:%d", port)}
		if len(p2pV2Bootstrappers) > 0 {
			c.P2P.V2.DefaultBootstrappers = &p2pV2Bootstrappers
		}

		// OCR configs
		c.OCR.Enabled = ptr(false)
		c.OCR.DefaultTransactionQueueDepth = ptr(uint32(200))
		c.OCR2.Enabled = ptr(true)
		c.OCR2.ContractPollInterval = config.MustNewDuration(5 * time.Second)

		var chains v2toml.EVMConfigs
		for chainID := range universe {
			chains = append(chains, createConfigV2Chain(uBigInt(chainID)))
		}
		c.EVM = chains
	})

	lggr := logger.TestLogger(t)
	lggr.SetLogLevel(zapcore.DebugLevel)
	ctx := testutils.Context(t)
	clients := make(map[uint64]client.Client)

	for chainID, uni := range universe {
		clients[chainID] = client.NewSimulatedBackendClient(t, uni.backend, uBigInt(chainID))
	}

	master := keystore.New(db, utils.FastScryptParams, lggr)

	keystore := KeystoreSim{
		eks: &EthKeystoreSim{
			Eth: master.Eth(),
			t:   t,
		},
		csa: master.CSA(),
	}
	mailMon := mailbox.NewMonitor("ccip", lggr.Named("mailbox"))
	evmOpts := chainlink.EVMFactoryConfig{
		ChainOpts: legacyevm.ChainOpts{
			AppConfig: config,
			GenEthClient: func(i *big.Int) client.Client {
				t.Log("genning eth client for chain id:", i.String())
				client, ok := clients[i.Uint64()]
				if !ok {
					t.Fatal("no backend for chainID", i)
				}
				return client
			},
			MailMon: mailMon,
			DS:      db,
		},
		CSAETHKeystore: keystore,
	}
	relayerFactory := chainlink.RelayerFactory{
		Logger:       lggr,
		LoopRegistry: plugins.NewLoopRegistry(lggr.Named("LoopRegistry"), config.Tracing()),
		GRPCOpts:     loop.GRPCOpts{},
	}
	initOps := []chainlink.CoreRelayerChainInitFunc{chainlink.InitEVM(testutils.Context(t), relayerFactory, evmOpts)}
	rci, err := chainlink.NewCoreRelayerChainInteroperators(initOps...)
	require.NoError(t, err)

	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:                     config,
		DS:                         db,
		KeyStore:                   master,
		RelayerChainInteroperators: rci,
		Logger:                     lggr,
		ExternalInitiatorManager:   nil,
		CloseLogger:                lggr.Sync,
		UnrestrictedHTTPClient:     &http.Client{},
		RestrictedHTTPClient:       &http.Client{},
		AuditLogger:                audit.NoopLogger,
		MailMon:                    mailMon,
		LoopRegistry:               plugins.NewLoopRegistry(lggr, config.Tracing()),
	})
	require.NoError(t, err)
	require.NoError(t, app.GetKeyStore().Unlock(ctx, "password"))
	_, err = app.GetKeyStore().P2P().Create(ctx)
	require.NoError(t, err)

	p2pIDs, err := app.GetKeyStore().P2P().GetAll()
	require.NoError(t, err)
	require.Len(t, p2pIDs, 1)
	peerID := p2pIDs[0].PeerID()

	// create a transmitter for each chain
	transmitters := make(map[uint64]common.Address)
	for chainID, uni := range universe {
		backend := uni.backend
		owner := uni.owner
		cID := uBigInt(chainID)
		addrs, err2 := app.GetKeyStore().Eth().EnabledAddressesForChain(testutils.Context(t), cID)
		require.NoError(t, err2)
		if len(addrs) == 1 {
			// just fund the address
			fundAddress(t, owner, addrs[0], assets.Ether(10).ToInt(), backend)
			transmitters[chainID] = addrs[0]
		} else {
			// create key and fund it
			_, err3 := app.GetKeyStore().Eth().Create(testutils.Context(t), cID)
			require.NoError(t, err3, "failed to create key for chain", chainID)
			sendingKeys, err3 := app.GetKeyStore().Eth().EnabledAddressesForChain(testutils.Context(t), cID)
			require.NoError(t, err3)
			require.Len(t, sendingKeys, 1)
			fundAddress(t, owner, sendingKeys[0], assets.Ether(10).ToInt(), backend)
			transmitters[chainID] = sendingKeys[0]
		}
	}
	require.Len(t, transmitters, len(universe))

	keybundle, err := app.GetKeyStore().OCR2().Create(ctx, chaintype.EVM)
	require.NoError(t, err)

	return &ocr3Node{
		// can't use this app because it doesn't have the right toml config
		// missing bootstrapp
		app:          app,
		peerID:       peerID.Raw(),
		transmitters: transmitters,
		keybundle:    keybundle,
		db:           db,
	}
}

func ptr[T any](v T) *T { return &v }

var _ keystore.Eth = &EthKeystoreSim{}

type EthKeystoreSim struct {
	keystore.Eth
	t *testing.T
}

// override
func (e *EthKeystoreSim) SignTx(ctx context.Context, address common.Address, tx *gethtypes.Transaction, chainID *big.Int) (*gethtypes.Transaction, error) {
	// always sign with chain id 1337 for the simulated backend
	e.t.Log("always signing tx for chain id:", chainID.String(), "with chain id 1337, tx hash:", tx.Hash())
	return e.Eth.SignTx(ctx, address, tx, big.NewInt(1337))
}

type KeystoreSim struct {
	eks keystore.Eth
	csa keystore.CSA
}

func (e KeystoreSim) Eth() keystore.Eth {
	return e.eks
}

func (e KeystoreSim) CSA() keystore.CSA {
	return e.csa
}

func fundAddress(t *testing.T, from *bind.TransactOpts, to common.Address, amount *big.Int, backend *backends.SimulatedBackend) {
	nonce, err := backend.PendingNonceAt(testutils.Context(t), from.From)
	require.NoError(t, err)
	gp, err := backend.SuggestGasPrice(testutils.Context(t))
	require.NoError(t, err)
	rawTx := gethtypes.NewTx(&gethtypes.LegacyTx{
		Nonce:    nonce,
		GasPrice: gp,
		Gas:      21000,
		To:       &to,
		Value:    amount,
	})
	signedTx, err := from.Signer(from.From, rawTx)
	require.NoError(t, err)
	err = backend.SendTransaction(testutils.Context(t), signedTx)
	require.NoError(t, err)
	backend.Commit()
}

func createConfigV2Chain(chainID *big.Int) *v2toml.EVMConfig {
	chain := v2toml.Defaults((*evmutils.Big)(chainID))
	chain.GasEstimator.LimitDefault = ptr(uint64(5e6))
	chain.LogPollInterval = config.MustNewDuration(100 * time.Millisecond)
	chain.Transactions.ForwardersEnabled = ptr(false)
	chain.FinalityDepth = ptr(uint32(2))
	return &v2toml.EVMConfig{
		ChainID: (*evmutils.Big)(chainID),
		Enabled: ptr(true),
		Chain:   chain,
		Nodes:   v2toml.EVMNodes{&v2toml.Node{}},
	}
}

func donOCRConfig(t *testing.T, uni onchainUniverse, oracles []confighelper2.OracleIdentityExtra) []byte {
	var schedule []int
	for range oracles {
		schedule = append(schedule, 1)
	}
	offchainConfig, onchainConfig := []byte{}, []byte{}
	f := uint8(1)
	_, _, f, _, offchainConfigVersion, offchainConfig, err := ocr3confighelper.ContractSetConfigArgsForTests(
		30*time.Second, // deltaProgress
		10*time.Second, // deltaResend
		20*time.Second, // deltaInitial
		2*time.Second,  // deltaRound
		20*time.Second, // deltaGrace
		10*time.Second, // deltaCertifiedCommitRequest
		10*time.Second, // deltaStage
		3,              // rmax
		schedule,
		oracles,
		offchainConfig,
		50*time.Millisecond, // maxDurationQuery
		5*time.Second,       // maxDurationObservation
		10*time.Second,      // maxDurationShouldAcceptAttestedReport
		10*time.Second,      // maxDurationShouldTransmitAcceptedReport
		int(f),
		onchainConfig)
	require.NoError(t, err, "failed to create contract config")
	/*
		struct OCR3Config {
			PluginType pluginType; // ────────╮ The plugin that the configuration is for.
			uint64 chainSelector; //          | The (remote) chain that the configuration is for.
			uint8 F; //                       | The "big F" parameter for the role DON.
			uint64 offchainConfigVersion; // ─╯ The version of the offchain configuration.
			bytes32 offrampAddress; // The remote chain combined (offramp|commit store) address.
			bytes32[2][] signers; // An associative array that contains (p2p id, onchain signer public key) pairs.
			bytes32[2][] transmitters; // An associative array that contains (p2p id, transmitter) pairs.
			bytes offchainConfig; // The offchain configuration for the OCR3 protocol. Protobuf encoded.
		}
	*/
	ocrConfigABI := `
{
	[
		{
			"type": "uint8"
		},
		{
			"type": "uint64"
		},
		{
			"type": "uint8"
		},
		{
			"type": "uint64"
		},
		{
			"type": "bytes32"
		},
		{
			"type": "bytes32[2][]"
		},
		{
			"type": "bytes32[2][]"
		},
		{
			"type": "bytes"
		}
	]
}`
	chainSelector, ok := chainsel.EvmChainIdToChainSelector()[uni.chainID]
	require.True(t, ok, "chain selector not found for chain id", uni.chainID)
	var offrampAddressBytes32 [32]byte
	copy(offrampAddressBytes32[:], uni.offramp.Address().Bytes())
	commitConfig, err := encodeutils.ABIEncode(ocrConfigABI,
		uint8(0),              // pluginType
		chainSelector,         // chainSelector
		f,                     // F
		offchainConfigVersion, // offchainConfigVersion
		offrampAddressBytes32, // offrampAddress
		nil,                   // TODO signers
		nil,                   // TODO transmitters
		offchainConfig,        // offchainConfig
	)
	require.NoError(t, err, "failed to encode commit OCR3 config")
	// TODO: implement
	return commitConfig
}

func createCCIPSpecToml(nodeP2PID, bootstrapP2PID string, bootstrapPort int, ocrKeyBundleID string) string {
	return fmt.Sprintf(`
type = "ccip"
capabilityVersion = "v1.0.0"
capabilityLabelledName = "ccip"
p2pKeyID = "%s"
p2pV2Bootstrappers = ["%s"]
[ocrKeyBundleIDs]
evm = "%s"
[relayConfigs.evm.chainReaderConfig.contracts.Offramp]
contractABI = "the abi"

[relayConfigs.evm.chainReaderConfig.contracts.Offramp.configs.getStuff]
chainSpecificName = "getStuffEVM"

[pluginConfig]
tokenPricesPipeline = "the pipeline"`,
		nodeP2PID,
		fmt.Sprintf("%s@127.0.0.1:%d", bootstrapP2PID,
			bootstrapPort,
		),
		ocrKeyBundleID,
	)
}
