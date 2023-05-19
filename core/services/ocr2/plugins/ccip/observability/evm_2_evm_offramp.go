package observability

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
)

type ObservedEVM2EVMOfframp struct {
	evm_2_evm_offramp.EVM2EVMOffRampInterface
	histogram *prometheus.HistogramVec
}

func NewObservedEVM2EVMOffRamp(address common.Address, backend bind.ContractBackend) (evm_2_evm_offramp.EVM2EVMOffRampInterface, error) {
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(address, backend)
	if err != nil {
		return nil, err
	}
	return &ObservedEVM2EVMOfframp{
		EVM2EVMOffRampInterface: offRamp,
		histogram:               evm2evmOffRampHistogram,
	}, nil
}

func (o *ObservedEVM2EVMOfframp) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	return withObservedContract(o.histogram, "GetSupportedTokens", func() ([]common.Address, error) {
		return o.EVM2EVMOffRampInterface.GetSupportedTokens(opts)
	})
}

func (o *ObservedEVM2EVMOfframp) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	return withObservedContract(o.histogram, "GetDestinationTokens", func() ([]common.Address, error) {
		return o.EVM2EVMOffRampInterface.GetDestinationTokens(opts)
	})
}

func (o *ObservedEVM2EVMOfframp) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	return withObservedContract(o.histogram, "GetDestinationToken", func() (common.Address, error) {
		return o.EVM2EVMOffRampInterface.GetDestinationToken(opts, sourceToken)
	})
}

func (o *ObservedEVM2EVMOfframp) CurrentRateLimiterState(opts *bind.CallOpts) (evm_2_evm_offramp.RateLimiterTokenBucket, error) {
	return withObservedContract(o.histogram, "CurrentRateLimiterState", func() (evm_2_evm_offramp.RateLimiterTokenBucket, error) {
		return o.EVM2EVMOffRampInterface.CurrentRateLimiterState(opts)
	})
}
