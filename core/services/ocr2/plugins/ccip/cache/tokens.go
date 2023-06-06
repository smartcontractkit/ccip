package cache

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

// NewCachedFeeTokens cache fee tokens returned from PriceRegistry
func NewCachedFeeTokens(
	lp logpoller.LogPoller,
	priceRegistry price_registry.PriceRegistryInterface,
	optimisticConfirmations int64,
) *CachedChain[[]common.Address] {
	return &CachedChain[[]common.Address]{
		observedEvents: []common.Hash{
			abihelpers.EventSignatures.FeeTokenAdded,
			abihelpers.EventSignatures.FeeTokenRemoved,
		},
		logPoller:               lp,
		address:                 []common.Address{priceRegistry.Address()},
		optimisticConfirmations: optimisticConfirmations,
		lock:                    &sync.RWMutex{},
		value:                   []common.Address{},
		lastChangeBlock:         0,
		origin:                  &feeTokensOrigin{priceRegistry: priceRegistry},
	}
}

type CachedTokens struct {
	SupportedTokens map[common.Address]common.Address
	FeeTokens       []common.Address
}

// NewCachedTokens cache both fee tokens and supported tokens. Therefore, it uses 4 different events
// when checking for changes in logpoller.LogPoller
func NewCachedTokens(
	lp logpoller.LogPoller,
	offRamp evm_2_evm_offramp.EVM2EVMOffRampInterface,
	priceRegistry price_registry.PriceRegistryInterface,
	optimisticConfirmations int64,
) *CachedChain[CachedTokens] {
	return &CachedChain[CachedTokens]{
		observedEvents: []common.Hash{
			abihelpers.EventSignatures.FeeTokenAdded,
			abihelpers.EventSignatures.FeeTokenRemoved,
			abihelpers.EventSignatures.PoolAdded,
			abihelpers.EventSignatures.PoolRemoved,
		},
		logPoller:               lp,
		address:                 []common.Address{priceRegistry.Address(), offRamp.Address()},
		optimisticConfirmations: optimisticConfirmations,
		lock:                    &sync.RWMutex{},
		value:                   CachedTokens{},
		lastChangeBlock:         0,
		origin: &feeAndSupportedTokensOrigin{
			feeTokensOrigin:       feeTokensOrigin{priceRegistry: priceRegistry},
			supportedTokensOrigin: supportedTokensOrigin{offRamp: offRamp}},
	}
}

type supportedTokensOrigin struct {
	offRamp evm_2_evm_offramp.EVM2EVMOffRampInterface
}

func (t *supportedTokensOrigin) Copy(value map[common.Address]common.Address) map[common.Address]common.Address {
	return copyMap(value)
}

// CallOrigin Generates the source to dest token mapping based on the offRamp.
// NOTE: this queries the offRamp n+1 times, where n is the number of enabled tokens.
func (t *supportedTokensOrigin) CallOrigin(ctx context.Context) (map[common.Address]common.Address, error) {
	srcToDstTokenMapping := make(map[common.Address]common.Address)
	sourceTokens, err := t.offRamp.GetSupportedTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	for _, sourceToken := range sourceTokens {
		dst, err1 := t.offRamp.GetDestinationToken(&bind.CallOpts{Context: ctx}, sourceToken)
		if err1 != nil {
			return nil, err1
		}
		srcToDstTokenMapping[sourceToken] = dst
	}
	return srcToDstTokenMapping, nil
}

type feeTokensOrigin struct {
	priceRegistry price_registry.PriceRegistryInterface
}

func (t *feeTokensOrigin) Copy(value []common.Address) []common.Address {
	return copyArray(value)
}

func (t *feeTokensOrigin) CallOrigin(ctx context.Context) ([]common.Address, error) {
	return t.priceRegistry.GetFeeTokens(&bind.CallOpts{Context: ctx})
}

func copyArray(source []common.Address) []common.Address {
	dst := make([]common.Address, len(source))
	copy(dst, source)
	return dst
}

type feeAndSupportedTokensOrigin struct {
	feeTokensOrigin       feeTokensOrigin
	supportedTokensOrigin supportedTokensOrigin
}

func (t *feeAndSupportedTokensOrigin) Copy(value CachedTokens) CachedTokens {
	return CachedTokens{
		SupportedTokens: t.supportedTokensOrigin.Copy(value.SupportedTokens),
		FeeTokens:       t.feeTokensOrigin.Copy(value.FeeTokens),
	}
}

func (t *feeAndSupportedTokensOrigin) CallOrigin(ctx context.Context) (CachedTokens, error) {
	supportedTokens, err := t.supportedTokensOrigin.CallOrigin(ctx)
	if err != nil {
		return CachedTokens{}, err
	}
	feeToken, err := t.feeTokensOrigin.CallOrigin(ctx)
	if err != nil {
		return CachedTokens{}, err
	}
	return CachedTokens{
		SupportedTokens: supportedTokens,
		FeeTokens:       feeToken,
	}, nil
}

func copyMap[M ~map[K]V, K comparable, V any](m M) M {
	cpy := make(M)
	for k, v := range m {
		cpy[k] = v
	}
	return cpy
}
