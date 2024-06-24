package oraclecreator_test

import (
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	evmrelayer "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

func genTestEVMRelayers(t *testing.T, opts legacyevm.ChainRelayExtenderConfig, ks evmrelayer.CSAETHKeystore) *chainlink.CoreRelayerChainInteroperators {
	f := chainlink.RelayerFactory{
		Logger:       opts.Logger,
		LoopRegistry: plugins.NewLoopRegistry(opts.Logger, opts.AppConfig.Tracing()),
	}

	relayers, err := chainlink.NewCoreRelayerChainInteroperators(chainlink.InitEVM(testutils.Context(t), f, chainlink.EVMFactoryConfig{
		ChainOpts:      opts.ChainOpts,
		CSAETHKeystore: ks,
	}))
	if err != nil {
		t.Fatal(err)
	}
	return relayers
}

func ptr[T any](v T) *T { return &v }

func TestOracleCreator_CreateCommitOracle(t *testing.T) {
	cfg := configtest.NewGeneralConfig(t, func(c *chainlink.Config, s *chainlink.Secrets) {
		s.Password.Keystore = models.NewSecret("dummy")
		c.EVM[0].Nodes[0].Name = ptr("fake")
		c.EVM[0].Nodes[0].HTTPURL = commonconfig.MustParseURL("http://fake.com")
		c.EVM[0].Nodes[0].WSURL = commonconfig.MustParseURL("WSS://fake.com/ws")
		// seems to be needed for config validate
		c.Insecure.OCRDevelopmentMode = nil
	})
	db := pgtest.NewSqlxDB(t)
	keyStore := cltest.NewKeyStore(t, db)

	lggr := logger.TestLogger(t)

	opts := legacyevm.ChainRelayExtenderConfig{
		Logger:   lggr,
		KeyStore: keyStore.Eth(),
		ChainOpts: legacyevm.ChainOpts{
			AppConfig: cfg,
			MailMon:   &mailbox.Monitor{},
			DS:        db,
		},
	}
	interops := genTestEVMRelayers(t, opts, keyStore)

	relayers, err := interops.GetIDToRelayerMap()
	require.NoError(t, err)

	relayID := types.NewRelayID("evm", "0")
	relayer, ok := relayers[relayID]
	require.True(t, ok)

	provider, err := relayer.NewPluginProvider(testutils.Context(t), types.RelayArgs{}, types.PluginArgs{})
	require.NoError(t, err)

	contractReader := provider.ChainReader()
	require.NotNil(t, contractReader)
}
