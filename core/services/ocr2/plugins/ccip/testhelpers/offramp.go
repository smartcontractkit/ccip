package testhelpers

import (
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type FakeOffRamp struct {
	*mock_contracts.EVM2EVMOffRampInterface

	rateLimiterState   evm_2_evm_offramp.RateLimiterTokenBucket
	senderNonces       map[common.Address]uint64
	tokenToPool        map[common.Address]common.Address
	dynamicConfig      evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig
	sourceToDestTokens map[common.Address]common.Address

	mu sync.RWMutex
}

func NewFakeOffRamp(t *testing.T) (*FakeOffRamp, common.Address) {
	addr := utils.RandomAddress()
	mockOffRamp := mock_contracts.NewEVM2EVMOffRampInterface(t)
	mockOffRamp.On("Address").Return(addr).Maybe()

	offRamp := &FakeOffRamp{EVM2EVMOffRampInterface: mockOffRamp}
	return offRamp, addr
}

func (o *FakeOffRamp) CurrentRateLimiterState(opts *bind.CallOpts) (evm_2_evm_offramp.RateLimiterTokenBucket, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) (evm_2_evm_offramp.RateLimiterTokenBucket, error) { return o.rateLimiterState, nil })
}

func (o *FakeOffRamp) SetRateLimiterState(state evm_2_evm_offramp.RateLimiterTokenBucket) {
	setOffRampVal(o, func(o *FakeOffRamp) { o.rateLimiterState = state })
}

func (o *FakeOffRamp) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) (uint64, error) { return o.senderNonces[sender], nil })
}

func (o *FakeOffRamp) SetSenderNonces(senderNonces map[common.Address]uint64) {
	setOffRampVal(o, func(o *FakeOffRamp) { o.senderNonces = senderNonces })
}

func (o *FakeOffRamp) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) (common.Address, error) {
		addr, exists := o.tokenToPool[destToken]
		if !exists {
			return common.Address{}, errors.New("not found")
		}
		return addr, nil
	})
}

func (o *FakeOffRamp) SetTokenPools(tokenToPool map[common.Address]common.Address) {
	setOffRampVal(o, func(o *FakeOffRamp) { o.tokenToPool = tokenToPool })
}

func (o *FakeOffRamp) GetDynamicConfig(opts *bind.CallOpts) (evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) (evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig, error) {
		return o.dynamicConfig, nil
	})
}

func (o *FakeOffRamp) SetDynamicConfig(cfg evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig) {
	setOffRampVal(o, func(o *FakeOffRamp) { o.dynamicConfig = cfg })
}

func (o *FakeOffRamp) SetSourceToDestTokens(m map[common.Address]common.Address) {
	setOffRampVal(o, func(o *FakeOffRamp) { o.sourceToDestTokens = m })
}

func (o *FakeOffRamp) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) ([]common.Address, error) {
		tks := make([]common.Address, 0, len(o.sourceToDestTokens))
		for tk := range o.sourceToDestTokens {
			tks = append(tks, tk)
		}
		return tks, nil
	})
}

func (o *FakeOffRamp) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) (common.Address, error) {
		addr, exists := o.sourceToDestTokens[sourceToken]
		if !exists {
			return common.Address{}, errors.New("token does not exist")
		}
		return addr, nil
	})
}

func (o *FakeOffRamp) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	return getOffRampVal(o, func(o *FakeOffRamp) ([]common.Address, error) {
		tokens := make([]common.Address, 0, len(o.sourceToDestTokens))
		for _, dst := range o.sourceToDestTokens {
			tokens = append(tokens, dst)
		}
		return tokens, nil
	})
}

func getOffRampVal[T any](o *FakeOffRamp, getter func(o *FakeOffRamp) (T, error)) (T, error) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return getter(o)
}

func setOffRampVal(o *FakeOffRamp, setter func(o *FakeOffRamp)) {
	o.mu.Lock()
	defer o.mu.Unlock()
	setter(o)
}
