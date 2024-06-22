package ocrimpls

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type configTracker struct {
	cfg cctypes.OCRConfig
}

func NewConfigTracker(cfg cctypes.OCRConfig) *configTracker {
	return &configTracker{cfg: cfg}
}

// LatestBlockHeight implements types.ContractConfigTracker.
func (c *configTracker) LatestBlockHeight(ctx context.Context) (blockHeight uint64, err error) {
	return 0, nil
}

// LatestConfig implements types.ContractConfigTracker.
func (c *configTracker) LatestConfig(ctx context.Context, changedInBlock uint64) (types.ContractConfig, error) {
	return types.ContractConfig{
		ConfigDigest:          c.cfg.ConfigDigest(),
		ConfigCount:           c.cfg.ConfigCount(),
		Signers:               toOnchainPublicKeys(c.cfg.Signers()),
		Transmitters:          toOCRAccounts(c.cfg.Transmitters()),
		F:                     c.cfg.F(),
		OnchainConfig:         []byte{},
		OffchainConfigVersion: c.cfg.OffchainConfigVersion(),
		OffchainConfig:        c.cfg.OffchainConfig(),
	}, nil
}

// LatestConfigDetails implements types.ContractConfigTracker.
func (c *configTracker) LatestConfigDetails(ctx context.Context) (changedInBlock uint64, configDigest types.ConfigDigest, err error) {
	return 0, types.ConfigDigest{}, nil
}

// Notify implements types.ContractConfigTracker.
func (c *configTracker) Notify() <-chan struct{} {
	return nil
}

func toOnchainPublicKeys(signers [][]byte) []types.OnchainPublicKey {
	keys := make([]types.OnchainPublicKey, len(signers))
	for i, signer := range signers {
		keys[i] = types.OnchainPublicKey(signer)
	}
	return keys
}

func toOCRAccounts(transmitters [][]byte) []types.Account {
	accounts := make([]types.Account, len(transmitters))
	for _, transmitter := range transmitters {
		accounts = append(accounts, types.Account(hexutil.Encode(transmitter)))
	}
	return accounts
}

var _ types.ContractConfigTracker = (*configTracker)(nil)
