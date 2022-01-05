package ccip

import (
	"math/big"
	"time"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	evmconfig "github.com/smartcontractkit/chainlink/core/chains/evm/config"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/store/models"

	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

// Fallback to config if explicit spec parameters are not set
func computeLocalConfig(config evmconfig.OCR2Config, dev bool, bt time.Duration, confs uint16, poll time.Duration) ocrtypes.LocalConfig {
	var blockchainTimeout time.Duration
	if bt != 0 {
		blockchainTimeout = bt
	} else {
		blockchainTimeout = config.OCR2BlockchainTimeout()
	}

	var contractConfirmations uint16
	if confs != 0 {
		contractConfirmations = confs
	} else {
		contractConfirmations = config.OCR2ContractConfirmations()
	}

	var contractConfigTrackerPollInterval time.Duration
	if poll != 0 {
		contractConfigTrackerPollInterval = poll
	} else {
		contractConfigTrackerPollInterval = config.OCR2ContractPollInterval()
	}

	lc := ocrtypes.LocalConfig{
		BlockchainTimeout:                  blockchainTimeout,
		ContractConfigConfirmations:        contractConfirmations,
		ContractConfigTrackerPollInterval:  contractConfigTrackerPollInterval,
		ContractTransmitterTransmitTimeout: config.OCR2ContractTransmitterTransmitTimeout(),
		DatabaseTimeout:                    config.OCR2DatabaseTimeout(),
	}
	if dev {
		// Skips config validation so we can use any config parameters we want.
		// For example to lower contractConfigTrackerPollInterval to speed up tests.
		lc.DevelopmentMode = ocrtypes.EnableDangerousDevelopmentMode
	}
	return lc
}

func parseBootstrapPeers(peers []string) (bootstrapPeers []ocrcommontypes.BootstrapperLocator, err error) {
	for _, bs := range peers {
		var bsl ocrcommontypes.BootstrapperLocator
		err = bsl.UnmarshalText([]byte(bs))
		if err != nil {
			return nil, err
		}
		bootstrapPeers = append(bootstrapPeers, bsl)
	}
	return
}

func getValidatedBootstrapPeers(specPeers []string, chain evm.Chain) ([]ocrcommontypes.BootstrapperLocator, error) {
	bootstrapPeers, err := parseBootstrapPeers(specPeers)
	if err != nil {
		return nil, err
	}
	if len(bootstrapPeers) == 0 {
		bootstrapPeers = chain.Config().P2PV2Bootstrappers()
		if err != nil {
			return nil, err
		}
	}
	return bootstrapPeers, nil
}

func validatePeerWrapper(specID *p2pkey.PeerID, chain evm.Chain, pw *ocrcommon.SingletonPeerWrapper) error {
	var peerID p2pkey.PeerID
	if specID != nil {
		peerID = *specID
	} else {
		peerID = chain.Config().P2PPeerID()
	}
	if !pw.IsStarted() {
		return errors.New("peerWrapper is not started. OCR2 jobs require a started and running peer. Did you forget to specify P2P_LISTEN_PORT?")
	} else if pw.PeerID != peerID {
		return errors.Errorf("given peer with ID '%s' does not match OCR2 configured peer with ID: %s", pw.PeerID.String(), peerID.String())
	}
	return nil
}

func getValidatedKeyBundle(specBundleID *models.Sha256Hash, chain evm.Chain, ks keystore.OCR2) (kb ocr2key.KeyBundle, err error) {
	var kbs string
	if specBundleID != nil {
		kbs = specBundleID.String()
	} else if kbs, err = chain.Config().OCR2KeyBundleID(); err != nil {
		return kb, err
	}
	key, err := ks.Get(kbs)
	if err != nil {
		return kb, err
	}
	return key, nil
}

func getTransmitterAddress(specAddress *ethkey.EIP55Address, chain evm.Chain) (ta ethkey.EIP55Address, err error) {
	if specAddress != nil {
		ta = *specAddress
	} else if ta, err = chain.Config().OCR2TransmitterAddress(); err != nil {
		return ta, err
	}
	return ta, nil
}

// Multi-chain tests using the sim have to be remapped to the default
// sim chainID because its a hardcoded constant in the geth code base and so
// and CHAINID op codes will ALWAYS be 1337.
func maybeRemapChainID(chainID *big.Int) *big.Int {
	testChainIDs := []*big.Int{big.NewInt(1000), big.NewInt(2000)}
	for _, testChainID := range testChainIDs {
		if chainID.Cmp(testChainID) == 0 {
			return big.NewInt(1337)
		}
	}
	return chainID
}
