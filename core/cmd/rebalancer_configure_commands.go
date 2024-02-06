package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/static"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type SetupRebalancerNodePayload struct {
	OnChainPublicKey  string
	OffChainPublicKey string
	ConfigPublicKey   string
	PeerID            string
	Transmitters      map[string]string // chain id -> transmitter address
	SendingKeys       []string
}

type rebalancerTemplateArgs struct {
	Name                    string
	ContractID              string
	OCRKeyBundleID          string
	TransmitterID           string
	ChainID                 int64
	LiquidityManagerAddress string
	LiquidityManagerNetwork uint64
	MaxNumTransfers         int64
	P2PV2BootstrapperPeerID string
	P2PV2BootstrapperPort   string
}

func (s *Shell) ConfigureRebalancerNode(
	c *cli.Context,
) (*SetupRebalancerNodePayload, error) {
	const (
		passwordArg                = "password"
		vrfPasswordArg             = "vrfpassword"
		l1ChainIDArg               = "l1ChainID"
		isBootstrapperArg          = "isBootstrapper"
		bootstrapperPeerIDArg      = "bootstrapperPeerID"
		jobTypeArg                 = "job-type"
		jobNameArg                 = "job-name"
		contractIDArg              = "contractID"
		liquidityManagerAddressArg = "liquidityManagerAddress"
		liquidityManagerNetworkArg = "liquidityManagerNetwork"
		maxNumTransfersArg         = "maxNumTransfers"
		bootstrapPortArg           = "bootstrapPort"
	)
	ctx := s.ctx()
	lggr := logger.Sugared(s.Logger.Named("ConfigureRebalancerNode"))
	lggr.Infow(
		fmt.Sprintf("Configuring Chainlink node for job type %s %s at commit %s", c.String("job-type"), static.Version, static.Sha),
		"Version", static.Version, "SHA", static.Sha)
	var pwd, vrfpwd *string
	if passwordFile := c.String(passwordArg); passwordFile != "" {
		p, err := utils.PasswordFromFile(passwordFile)
		if err != nil {
			return nil, errors.Wrap(err, "error reading password from file")
		}
		pwd = &p
	}
	if vrfPasswordFile := c.String(vrfPasswordArg); len(vrfPasswordFile) != 0 {
		p, err := utils.PasswordFromFile(vrfPasswordFile)
		if err != nil {
			return nil, errors.Wrapf(err, "error reading VRF password from vrfpassword file \"%s\"", vrfPasswordFile)
		}
		vrfpwd = &p
	}
	s.Config.SetPasswords(pwd, vrfpwd)
	err := s.Config.Validate()
	if err != nil {
		return nil, s.errorOut(errors.Wrap(err, "config validation failed"))
	}
	cfg := s.Config
	ldb := pg.NewLockedDB(cfg.AppID(), cfg.Database(), cfg.Database().Lock(), lggr)
	if err = ldb.Open(ctx); err != nil {
		return nil, s.errorOut(errors.Wrap(err, "opening db"))
	}
	defer lggr.ErrorIfFn(ldb.Close, "Error closing db")

	app, err := s.AppFactory.NewApplication(ctx, s.Config, lggr, ldb.DB())
	if err != nil {
		return nil, s.errorOut(errors.Wrap(err, "fatal error instantiating application"))
	}

	chainID := c.Int64(l1ChainIDArg)
	// Initialize keystore and generate keys.
	keyStore := app.GetKeyStore()
	err = setupKeystore(s, app, keyStore)
	if err != nil {
		return nil, s.errorOut(err)
	}

	// Start application.
	err = app.Start(ctx)
	if err != nil {
		return nil, s.errorOut(err)
	}

	// Defer close application.
	defer lggr.ErrorIfFn(app.Stop, "Failed to Stop application")

	// create sending keys for each chain enabled
	transmitters := make(map[string]string)
	chains, err := app.GetRelayers().LegacyEVMChains().List()
	if err != nil {
		return nil, s.errorOut(err)
	}
	for _, chain := range chains {
		ethKeys, err2 := app.GetKeyStore().Eth().EnabledKeysForChain(chain.ID())
		if err2 != nil {
			return nil, s.errorOut(err2)
		}
		if len(ethKeys) == 0 {
			return nil, s.errorOut(errors.New("no eth keys found"))
		}
		transmitters[chain.ID().String()] = ethKeys[0].Address.Hex()
	}

	// transmitterID on the job spec will be that of the main chain
	mainChainKeys, err := app.GetKeyStore().Eth().EnabledKeysForChain(big.NewInt(chainID))
	if err != nil {
		return nil, s.errorOut(err)
	}
	if len(mainChainKeys) == 0 {
		return nil, s.errorOut(errors.New("no eth keys found"))
	}
	transmitterID := mainChainKeys[0].Address.Hex()

	// Get all configuration parameters.
	p2p, _ := app.GetKeyStore().P2P().GetAll()
	ocr2List, _ := app.GetKeyStore().OCR2().GetAll()
	peerID := p2p[0].PeerID().Raw()
	if !c.Bool(isBootstrapperArg) {
		peerID = c.String(bootstrapperPeerIDArg)
	}

	// Find the EVM OCR2 bundle.
	var ocr2 ocr2key.KeyBundle
	for _, ocr2Item := range ocr2List {
		if ocr2Item.ChainType() == chaintype.EVM {
			ocr2 = ocr2Item
		}
	}
	if ocr2 == nil {
		return nil, s.errorOut(errors.Wrap(job.ErrNoSuchKeyBundle, "evm OCR2 key bundle not found"))
	}
	offChainPublicKey := ocr2.OffchainPublicKey()
	configPublicKey := ocr2.ConfigEncryptionPublicKey()

	if c.Bool(isBootstrapperArg) {
		// Set up bootstrapper job if bootstrapper.
		err = createRebalancerBootstrapperJob(ctx, lggr, chainID, c.String(contractIDArg), app)
	} else if c.String(jobTypeArg) == "rebalancer" {
		// Set up rebalancer job.
		err = createRebalancerJob(ctx, lggr, app, rebalancerTemplateArgs{
			Name:                    c.String(jobNameArg),
			ContractID:              c.String(contractIDArg),
			OCRKeyBundleID:          ocr2.ID(),
			TransmitterID:           transmitterID,
			ChainID:                 chainID,
			LiquidityManagerAddress: c.String(liquidityManagerAddressArg),
			LiquidityManagerNetwork: c.Uint64(liquidityManagerNetworkArg),
			MaxNumTransfers:         c.Int64(maxNumTransfersArg),
			P2PV2BootstrapperPeerID: peerID,
			P2PV2BootstrapperPort:   c.String(bootstrapPortArg),
		})
	} else {
		err = fmt.Errorf("unknown job type: %s", c.String(jobTypeArg))
	}

	if err != nil {
		return nil, err
	}

	return &SetupRebalancerNodePayload{
		OnChainPublicKey:  ocr2.OnChainPublicKey(),
		OffChainPublicKey: hex.EncodeToString(offChainPublicKey[:]),
		ConfigPublicKey:   hex.EncodeToString(configPublicKey[:]),
		PeerID:            p2p[0].PeerID().Raw(),
		Transmitters:      transmitters,
	}, nil
}

func createRebalancerJob(
	ctx context.Context,
	lggr logger.Logger,
	app chainlink.Application,
	args rebalancerTemplateArgs) error {
	const RebalancerTemplate = `
# Rebalancer Spec
type                 	= "offchainreporting2"
schemaVersion        	= 1
name                 	= "%s"
maxTaskDuration      	= "30s"
contractID           	= "%s"
ocrKeyBundleID       	= "%s"
relay                	= "evm"
pluginType           	= "rebalancer"
transmitterID        	= "%s"
forwardingAllowed       = false
contractConfigTrackerPollInterval = "15s"
# p2pv2Bootstrappers below
%s

[relayConfig]
chainID              	= %d
# This is the fromBlock for the main chain
# We set config after we launch the nodes, so this is not needed
# fromBlock               = blah
[relayConfig.fromBlocks]
# these are the fromBlock values for the follower chains
# We set config after we launch the nodes, so this is not needed

[pluginConfig]
liquidityManagerAddress = "%s"
liquidityManagerNetwork = "%d"
closePluginTimeoutSec = 10
[pluginConfig.rebalancerConfig]
type = "random"
[pluginConfig.rebalancerConfig.randomRebalancerConfig]
maxNumTransfers = %d
checkSourceDestEqual = false
`
	lggr.Info("Liquidity manager network:", args.LiquidityManagerNetwork)
	sp := fmt.Sprintf(RebalancerTemplate,
		args.Name,
		args.ContractID,
		args.OCRKeyBundleID,
		args.TransmitterID,
		fmt.Sprintf(`p2pv2Bootstrappers   = ["%s@127.0.0.1:%s"]`, args.P2PV2BootstrapperPeerID, args.P2PV2BootstrapperPort),
		args.ChainID,
		args.LiquidityManagerAddress,
		args.LiquidityManagerNetwork,
		args.MaxNumTransfers,
	)

	var jb job.Job
	err := toml.Unmarshal([]byte(sp), &jb)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal job spec")
	}
	var os job.OCR2OracleSpec
	err = toml.Unmarshal([]byte(sp), &os)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal job spec")
	}
	jb.OCR2OracleSpec = &os

	err = app.AddJobV2(ctx, &jb)
	if err != nil {
		return errors.Wrap(err, "failed to add job")
	}
	lggr.Info("rebalancer spec:", sp)

	return nil
}

func createRebalancerBootstrapperJob(
	ctx context.Context,
	lggr logger.Logger,
	l1ChainID int64,
	contractID string,
	app chainlink.Application,
) error {
	sp := fmt.Sprintf(BootstrapTemplate,
		l1ChainID,
		contractID,
		l1ChainID,
	)
	var jb job.Job
	err := toml.Unmarshal([]byte(sp), &jb)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal job spec")
	}
	var os job.BootstrapSpec
	err = toml.Unmarshal([]byte(sp), &os)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal job spec")
	}
	jb.BootstrapSpec = &os

	err = app.AddJobV2(ctx, &jb)
	if err != nil {
		return errors.Wrap(err, "failed to add job")
	}
	lggr.Info("rebalancer bootstrap spec:", sp)

	// Give a cooldown
	time.Sleep(time.Second)

	return nil
}
