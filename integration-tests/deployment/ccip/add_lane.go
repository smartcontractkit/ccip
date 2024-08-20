package ccipdeployment

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
)

func AddLane(e deployment.Environment, state CCIPOnChainState, from, to uint64) error {
	// TODO: Batch
	tx, err := state.Routers[from].ApplyRampUpdates(e.Chains[from].DeployerKey, []router.RouterOnRamp{
		{
			DestChainSelector: to,
			OnRamp:            state.EvmOnRampsV160[from].Address(),
		},
	}, []router.RouterOffRamp{}, []router.RouterOffRamp{})
	if err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}
	tx, err = state.EvmOnRampsV160[from].ApplyDestChainConfigUpdates(e.Chains[from].DeployerKey,
		[]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs{
			{
				DestChainSelector: to,
				Router:            state.Routers[from].Address(),
			},
		})
	if err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}

	_, err = state.PriceRegistries[from].UpdatePrices(
		e.Chains[from].DeployerKey, price_registry.InternalPriceUpdates{
			TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
				{
					SourceToken: state.LinkTokens[from].Address(),
					UsdPerToken: deployment.E18Mult(20),
				},
				{
					SourceToken: state.Weth9s[from].Address(),
					UsdPerToken: deployment.E18Mult(4000),
				},
			},
			GasPriceUpdates: []price_registry.InternalGasPriceUpdate{
				{
					DestChainSelector: to,
					UsdPerUnitGas:     big.NewInt(2e12),
				},
			}})
	if err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}

	// Enable dest in price registry
	tx, err = state.PriceRegistries[from].ApplyDestChainConfigUpdates(e.Chains[from].DeployerKey,
		[]price_registry.PriceRegistryDestChainConfigArgs{
			{
				DestChainSelector: to,
				DestChainConfig:   defaultPriceRegistryDestChainConfig(),
			},
		})
	if err := deployment.ConfirmIfNoError(e.Chains[from], tx, err); err != nil {
		return err
	}

	tx, err = state.EvmOffRampsV160[to].ApplySourceChainConfigUpdates(e.Chains[to].DeployerKey,
		[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{
			{
				Router:              state.Routers[to].Address(),
				SourceChainSelector: from,
				IsEnabled:           true,
				OnRamp:              common.LeftPadBytes(state.EvmOnRampsV160[from].Address().Bytes(), 32),
			},
		})
	if err := deployment.ConfirmIfNoError(e.Chains[to], tx, err); err != nil {
		return err
	}
	tx, err = state.Routers[to].ApplyRampUpdates(e.Chains[to].DeployerKey, []router.RouterOnRamp{}, []router.RouterOffRamp{}, []router.RouterOffRamp{
		{
			SourceChainSelector: from,
			OffRamp:             state.EvmOffRampsV160[to].Address(),
		},
	})
	return deployment.ConfirmIfNoError(e.Chains[to], tx, err)
}

func defaultPriceRegistryDestChainConfig() price_registry.PriceRegistryDestChainConfig {
	// https://github.com/smartcontractkit/ccip/blob/c4856b64bd766f1ddbaf5d13b42d3c4b12efde3a/contracts/src/v0.8/ccip/libraries/Internal.sol#L337-L337
	/*
		```Solidity
			// bytes4(keccak256("CCIP ChainFamilySelector EVM"))
			bytes4 public constant CHAIN_FAMILY_SELECTOR_EVM = 0x2812d52c;
		```
	*/
	evmFamilySelector, _ := hex.DecodeString("2812d52c")
	return price_registry.PriceRegistryDestChainConfig{
		IsEnabled:                         true,
		MaxNumberOfTokensPerMsg:           10,
		MaxDataBytes:                      256,
		MaxPerMsgGasLimit:                 3_000_000,
		DestGasOverhead:                   50_000,
		DefaultTokenFeeUSDCents:           1,
		DestGasPerPayloadByte:             10,
		DestDataAvailabilityOverheadGas:   0,
		DestGasPerDataAvailabilityByte:    100,
		DestDataAvailabilityMultiplierBps: 1,
		DefaultTokenDestGasOverhead:       125_000,
		DefaultTokenDestBytesOverhead:     32,
		DefaultTxGasLimit:                 200_000,
		GasMultiplierWeiPerEth:            1,
		NetworkFeeUSDCents:                1,
		ChainFamilySelector:               [4]byte(evmFamilySelector),
	}
}
