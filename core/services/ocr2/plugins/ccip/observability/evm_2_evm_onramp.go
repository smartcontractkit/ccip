package observability

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
)

type ObservedEVM2EVMOnramp struct {
	evm_2_evm_onramp.EVM2EVMOnRampInterface
	metric metricDetails
}

func NewObservedEVM2EVMnRamp(address common.Address, pluginName string, backend bind.ContractBackend) (evm_2_evm_onramp.EVM2EVMOnRampInterface, error) {
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(address, backend)
	if err != nil {
		return nil, err
	}
	return &ObservedEVM2EVMOnramp{
		EVM2EVMOnRampInterface: onRamp,
		metric: metricDetails{
			histogram:  evm2evmOnRampHistogram,
			pluginName: pluginName,
		},
	}, nil
}

func (o *ObservedEVM2EVMOnramp) GetDynamicConfig(opts *bind.CallOpts) (evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig, error) {
	return withObservedContract(o.metric, "GetDynamicConfig", func() (evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig, error) {
		return o.EVM2EVMOnRampInterface.GetDynamicConfig(opts)
	})
}
