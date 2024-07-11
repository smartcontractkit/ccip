// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {PriceRegistry} from "../../PriceRegistry.sol";
import {Client} from "../../libraries/Client.sol";

contract PriceRegistryHelper is PriceRegistry {
  constructor(
    StaticConfig memory staticConfig,
    address[] memory priceUpdaters,
    address[] memory feeTokens,
    TokenPriceFeedUpdate[] memory tokenPriceFeeds,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs,
    DestChainDynamicConfigArgs[] memory destChainConfigArgs
  )
    PriceRegistry(
      staticConfig,
      priceUpdaters,
      feeTokens,
      tokenPriceFeeds,
      tokenTransferFeeConfigArgs,
      premiumMultiplierWeiPerEthArgs,
      destChainConfigArgs
    )
  {}

  function getDataAvailabilityCost(
    uint64 destChainSelector,
    uint112 dataAvailabilityGasPrice,
    uint256 messageDataLength,
    uint256 numberOfTokens,
    uint32 tokenTransferBytesOverhead
  ) external view returns (uint256) {
    return _getDataAvailabilityCost(
      s_destChainDynamicConfigs[destChainSelector],
      dataAvailabilityGasPrice,
      messageDataLength,
      numberOfTokens,
      tokenTransferBytesOverhead
    );
  }

  function getTokenTransferCost(
    uint64 destChainSelector,
    address feeToken,
    uint224 feeTokenPrice,
    Client.EVMTokenAmount[] calldata tokenAmounts
  ) external view returns (uint256, uint32, uint32) {
    return _getTokenTransferCost(
      s_destChainDynamicConfigs[destChainSelector], destChainSelector, feeToken, feeTokenPrice, tokenAmounts
    );
  }

  function parseEVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint64 destChainSelector
  ) external view returns (Client.EVMExtraArgsV2 memory) {
    return _parseEVMExtraArgsFromBytes(extraArgs, s_destChainDynamicConfigs[destChainSelector]);
  }

  function parseEVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    DestChainDynamicConfig memory destChainDynamicConfig
  ) external pure returns (Client.EVMExtraArgsV2 memory) {
    return _parseEVMExtraArgsFromBytes(extraArgs, destChainDynamicConfig);
  }

  function validateDestFamilyAddress(bytes4 chainFamilySelector, bytes memory destAddress) external pure {
    _validateDestFamilyAddress(chainFamilySelector, destAddress);
  }
}
