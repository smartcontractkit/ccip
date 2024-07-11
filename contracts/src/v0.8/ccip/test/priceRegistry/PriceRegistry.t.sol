// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {AuthorizedCallers} from "../../../shared/access/AuthorizedCallers.sol";
import {MockV3Aggregator} from "../../../tests/MockV3Aggregator.sol";
import {PriceRegistry} from "../../PriceRegistry.sol";
import {IPriceRegistry} from "../../interfaces/IPriceRegistry.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {Pool} from "../../libraries/Pool.sol";
import {USDPriceWith18Decimals} from "../../libraries/USDPriceWith18Decimals.sol";

import {TokenSetup} from "../TokenSetup.t.sol";
import {PriceRegistryHelper} from "../helpers/PriceRegistryHelper.sol";
import {Vm} from "forge-std/Vm.sol";
import {console} from "forge-std/console.sol";

contract PriceRegistrySetup is TokenSetup {
  uint112 internal constant USD_PER_GAS = 1e6; // 0.001 gwei
  uint112 internal constant USD_PER_DATA_AVAILABILITY_GAS = 1e9; // 1 gwei

  address internal constant CUSTOM_TOKEN = address(12345);
  uint224 internal constant CUSTOM_TOKEN_PRICE = 1e17; // $0.1 CUSTOM

  // Encode L1 gas price and L2 gas price into a packed price.
  // L1 gas price is left-shifted to the higher-order bits.
  uint224 internal constant PACKED_USD_PER_GAS =
    (uint224(USD_PER_DATA_AVAILABILITY_GAS) << Internal.GAS_PRICE_BITS) + USD_PER_GAS;

  PriceRegistryHelper internal s_priceRegistry;
  // Cheat to store the price updates in storage since struct arrays aren't supported.
  bytes internal s_encodedInitialPriceUpdates;
  address internal s_weth;

  address[] internal s_sourceFeeTokens;
  uint224[] internal s_sourceTokenPrices;
  address[] internal s_destFeeTokens;
  uint224[] internal s_destTokenPrices;

  PriceRegistry.PremiumMultiplierWeiPerEthArgs[] internal s_priceRegistryPremiumMultiplierWeiPerEthArgs;
  PriceRegistry.TokenTransferFeeConfigArgs[] internal s_priceRegistryTokenTransferFeeConfigArgs;

  mapping(address token => address dataFeedAddress) internal s_dataFeedByToken;

  function setUp() public virtual override {
    TokenSetup.setUp();

    _deployTokenPriceDataFeed(s_sourceFeeToken, 8, 1e8);

    s_weth = s_sourceRouter.getWrappedNative();
    _deployTokenPriceDataFeed(s_weth, 8, 1e11);

    address[] memory sourceFeeTokens = new address[](3);
    sourceFeeTokens[0] = s_sourceTokens[0];
    sourceFeeTokens[1] = s_sourceTokens[1];
    sourceFeeTokens[2] = s_sourceRouter.getWrappedNative();
    s_sourceFeeTokens = sourceFeeTokens;

    uint224[] memory sourceTokenPrices = new uint224[](3);
    sourceTokenPrices[0] = 5e18;
    sourceTokenPrices[1] = 2000e18;
    sourceTokenPrices[2] = 2000e18;
    s_sourceTokenPrices = sourceTokenPrices;

    address[] memory destFeeTokens = new address[](3);
    destFeeTokens[0] = s_destTokens[0];
    destFeeTokens[1] = s_destTokens[1];
    destFeeTokens[2] = s_destRouter.getWrappedNative();
    s_destFeeTokens = destFeeTokens;

    uint224[] memory destTokenPrices = new uint224[](3);
    destTokenPrices[0] = 5e18;
    destTokenPrices[1] = 2000e18;
    destTokenPrices[2] = 2000e18;
    s_destTokenPrices = destTokenPrices;

    uint256 sourceTokenCount = sourceFeeTokens.length;
    uint256 destTokenCount = destFeeTokens.length;
    address[] memory pricedTokens = new address[](sourceTokenCount + destTokenCount);
    uint224[] memory tokenPrices = new uint224[](sourceTokenCount + destTokenCount);
    for (uint256 i = 0; i < sourceTokenCount; ++i) {
      pricedTokens[i] = sourceFeeTokens[i];
      tokenPrices[i] = sourceTokenPrices[i];
    }
    for (uint256 i = 0; i < destTokenCount; ++i) {
      pricedTokens[i + sourceTokenCount] = destFeeTokens[i];
      tokenPrices[i + sourceTokenCount] = destTokenPrices[i];
    }

    Internal.PriceUpdates memory priceUpdates = getPriceUpdatesStruct(pricedTokens, tokenPrices);
    priceUpdates.gasPriceUpdates =
      getSingleGasPriceUpdateStruct(DEST_CHAIN_SELECTOR, PACKED_USD_PER_GAS).gasPriceUpdates;

    s_encodedInitialPriceUpdates = abi.encode(priceUpdates);

    address[] memory priceUpdaters = new address[](0);
    address[] memory feeTokens = new address[](2);
    feeTokens[0] = s_sourceTokens[0];
    feeTokens[1] = s_weth;
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](0);

    s_priceRegistryPremiumMultiplierWeiPerEthArgs.push(
      PriceRegistry.PremiumMultiplierWeiPerEthArgs({
        token: s_sourceFeeToken,
        premiumMultiplierWeiPerEth: 5e17 // 0.5x
      })
    );
    s_priceRegistryPremiumMultiplierWeiPerEthArgs.push(
      PriceRegistry.PremiumMultiplierWeiPerEthArgs({
        token: s_sourceRouter.getWrappedNative(),
        premiumMultiplierWeiPerEth: 2e18 // 2x
      })
    );

    s_priceRegistryTokenTransferFeeConfigArgs.push();
    s_priceRegistryTokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    s_priceRegistryTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs.push(
      PriceRegistry.TokenTransferFeeConfigSingleTokenArgs({
        token: s_sourceFeeToken,
        tokenTransferFeeConfig: PriceRegistry.TokenTransferFeeConfig({
          minFeeUSDCents: 1_00, // 1 USD
          maxFeeUSDCents: 1000_00, // 1,000 USD
          deciBps: 2_5, // 2.5 bps, or 0.025%
          destGasOverhead: 40_000,
          destBytesOverhead: 32,
          isEnabled: true
        })
      })
    );
    s_priceRegistryTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs.push(
      PriceRegistry.TokenTransferFeeConfigSingleTokenArgs({
        token: s_sourceRouter.getWrappedNative(),
        tokenTransferFeeConfig: PriceRegistry.TokenTransferFeeConfig({
          minFeeUSDCents: 50, // 0.5 USD
          maxFeeUSDCents: 500_00, // 500 USD
          deciBps: 5_0, // 5 bps, or 0.05%
          destGasOverhead: 10_000,
          destBytesOverhead: 100,
          isEnabled: true
        })
      })
    );
    s_priceRegistryTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs.push(
      PriceRegistry.TokenTransferFeeConfigSingleTokenArgs({
        token: CUSTOM_TOKEN,
        tokenTransferFeeConfig: PriceRegistry.TokenTransferFeeConfig({
          minFeeUSDCents: 2_00, // 1 USD
          maxFeeUSDCents: 2000_00, // 1,000 USD
          deciBps: 10_0, // 10 bps, or 0.1%
          destGasOverhead: 1,
          destBytesOverhead: 200,
          isEnabled: true
        })
      })
    );

    s_priceRegistry = new PriceRegistryHelper(
      PriceRegistry.StaticConfig({linkToken: s_sourceTokens[0], maxFeeJuelsPerMsg: MAX_MSG_FEES_JUELS}),
      priceUpdaters,
      feeTokens,
      uint32(TWELVE_HOURS),
      tokenPriceFeedUpdates,
      s_priceRegistryTokenTransferFeeConfigArgs,
      s_priceRegistryPremiumMultiplierWeiPerEthArgs,
      _generatePriceRegistryDestChainDynamicConfigArgs()
    );
    s_priceRegistry.updatePrices(priceUpdates);
  }

  function _deployTokenPriceDataFeed(address token, uint8 decimals, int256 initialAnswer) internal returns (address) {
    MockV3Aggregator dataFeed = new MockV3Aggregator(decimals, initialAnswer);
    s_dataFeedByToken[token] = address(dataFeed);
    return address(dataFeed);
  }

  function getPriceUpdatesStruct(
    address[] memory tokens,
    uint224[] memory prices
  ) internal pure returns (Internal.PriceUpdates memory) {
    uint256 length = tokens.length;

    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](length);
    for (uint256 i = 0; i < length; ++i) {
      tokenPriceUpdates[i] = Internal.TokenPriceUpdate({sourceToken: tokens[i], usdPerToken: prices[i]});
    }
    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: tokenPriceUpdates, gasPriceUpdates: new Internal.GasPriceUpdate[](0)});

    return priceUpdates;
  }

  function getEmptyPriceUpdates() internal pure returns (Internal.PriceUpdates memory priceUpdates) {
    return Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](0),
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    });
  }

  function getSingleTokenPriceFeedUpdateStruct(
    address sourceToken,
    address dataFeedAddress,
    uint8 tokenDecimals
  ) internal pure returns (PriceRegistry.TokenPriceFeedUpdate memory) {
    return PriceRegistry.TokenPriceFeedUpdate({
      sourceToken: sourceToken,
      feedConfig: IPriceRegistry.TokenPriceFeedConfig({dataFeedAddress: dataFeedAddress, tokenDecimals: tokenDecimals})
    });
  }

  function _initialiseSingleTokenPriceFeed() internal returns (address) {
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]], 18);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    return s_sourceTokens[0];
  }

  function _generateTokenTransferFeeConfigArgs(
    uint256 destChainSelectorLength,
    uint256 tokenLength
  ) internal pure returns (PriceRegistry.TokenTransferFeeConfigArgs[] memory) {
    PriceRegistry.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
      new PriceRegistry.TokenTransferFeeConfigArgs[](destChainSelectorLength);
    for (uint256 i = 0; i < destChainSelectorLength; ++i) {
      tokenTransferFeeConfigArgs[i].tokenTransferFeeConfigs =
        new PriceRegistry.TokenTransferFeeConfigSingleTokenArgs[](tokenLength);
    }
    return tokenTransferFeeConfigArgs;
  }

  function _generatePriceRegistryDestChainDynamicConfigArgs()
    internal
    pure
    returns (PriceRegistry.DestChainDynamicConfigArgs[] memory)
  {
    PriceRegistry.DestChainDynamicConfigArgs[] memory destChainConfigs =
      new PriceRegistry.DestChainDynamicConfigArgs[](1);
    destChainConfigs[0] = PriceRegistry.DestChainDynamicConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      dynamicConfig: PriceRegistry.DestChainDynamicConfig({
        isEnabled: true,
        maxNumberOfTokensPerMsg: MAX_TOKENS_LENGTH,
        destGasOverhead: DEST_GAS_OVERHEAD,
        destGasPerPayloadByte: DEST_GAS_PER_PAYLOAD_BYTE,
        destDataAvailabilityOverheadGas: DEST_DATA_AVAILABILITY_OVERHEAD_GAS,
        destGasPerDataAvailabilityByte: DEST_GAS_PER_DATA_AVAILABILITY_BYTE,
        destDataAvailabilityMultiplierBps: DEST_GAS_DATA_AVAILABILITY_MULTIPLIER_BPS,
        maxDataBytes: MAX_DATA_SIZE,
        maxPerMsgGasLimit: MAX_GAS_LIMIT,
        defaultTokenFeeUSDCents: DEFAULT_TOKEN_FEE_USD_CENTS,
        defaultTokenDestGasOverhead: DEFAULT_TOKEN_DEST_GAS_OVERHEAD,
        defaultTokenDestBytesOverhead: DEFAULT_TOKEN_BYTES_OVERHEAD,
        defaultTxGasLimit: GAS_LIMIT,
        gasMultiplierWeiPerEth: 5e17,
        networkFeeUSDCents: 1_00,
        enforceOutOfOrder: false,
        chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM
      })
    });
    return destChainConfigs;
  }

  function _assertTokenPriceFeedConfigEquality(
    IPriceRegistry.TokenPriceFeedConfig memory config1,
    IPriceRegistry.TokenPriceFeedConfig memory config2
  ) internal pure virtual {
    assertEq(config1.dataFeedAddress, config2.dataFeedAddress);
    assertEq(config1.tokenDecimals, config2.tokenDecimals);
  }

  function _assertTokenPriceFeedConfigUnconfigured(IPriceRegistry.TokenPriceFeedConfig memory config)
    internal
    pure
    virtual
  {
    _assertTokenPriceFeedConfigEquality(
      config, IPriceRegistry.TokenPriceFeedConfig({dataFeedAddress: address(0), tokenDecimals: 0})
    );
  }

  function _assertTokenTransferFeeConfigEqual(
    PriceRegistry.TokenTransferFeeConfig memory a,
    PriceRegistry.TokenTransferFeeConfig memory b
  ) internal pure {
    assertEq(a.minFeeUSDCents, b.minFeeUSDCents);
    assertEq(a.maxFeeUSDCents, b.maxFeeUSDCents);
    assertEq(a.deciBps, b.deciBps);
    assertEq(a.destGasOverhead, b.destGasOverhead);
    assertEq(a.destBytesOverhead, b.destBytesOverhead);
    assertEq(a.isEnabled, b.isEnabled);
  }

  function _assertPriceRegistryStaticConfigsEqual(
    PriceRegistry.StaticConfig memory a,
    PriceRegistry.StaticConfig memory b
  ) internal pure {
    assertEq(a.linkToken, b.linkToken);
    assertEq(a.maxFeeJuelsPerMsg, b.maxFeeJuelsPerMsg);
  }

  function _assertPriceRegistryDestChainDynamicConfigsEqual(
    PriceRegistry.DestChainDynamicConfig memory a,
    PriceRegistry.DestChainDynamicConfig memory b
  ) internal pure {
    assertEq(a.isEnabled, b.isEnabled);
    assertEq(a.maxNumberOfTokensPerMsg, b.maxNumberOfTokensPerMsg);
    assertEq(a.maxDataBytes, b.maxDataBytes);
    assertEq(a.maxPerMsgGasLimit, b.maxPerMsgGasLimit);
    assertEq(a.destGasOverhead, b.destGasOverhead);
    assertEq(a.destGasPerPayloadByte, b.destGasPerPayloadByte);
    assertEq(a.destDataAvailabilityOverheadGas, b.destDataAvailabilityOverheadGas);
    assertEq(a.destGasPerDataAvailabilityByte, b.destGasPerDataAvailabilityByte);
    assertEq(a.destDataAvailabilityMultiplierBps, b.destDataAvailabilityMultiplierBps);
    assertEq(a.defaultTokenFeeUSDCents, b.defaultTokenFeeUSDCents);
    assertEq(a.defaultTokenDestGasOverhead, b.defaultTokenDestGasOverhead);
    assertEq(a.defaultTokenDestBytesOverhead, b.defaultTokenDestBytesOverhead);
    assertEq(a.defaultTxGasLimit, b.defaultTxGasLimit);
  }
}

contract PriceRegistryFeeSetup is PriceRegistrySetup {
  uint224 internal s_feeTokenPrice;
  uint224 internal s_wrappedTokenPrice;
  uint224 internal s_customTokenPrice;

  address internal s_selfServeTokenDefaultPricing = makeAddr("self-serve-token-default-pricing");

  function setUp() public virtual override {
    super.setUp();

    s_feeTokenPrice = s_sourceTokenPrices[0];
    s_wrappedTokenPrice = s_sourceTokenPrices[2];
    s_customTokenPrice = CUSTOM_TOKEN_PRICE;

    s_priceRegistry.updatePrices(getSingleTokenPriceUpdateStruct(CUSTOM_TOKEN, CUSTOM_TOKEN_PRICE));
  }

  function _setupFeeTokenPools() internal {
    // Add additional pool addresses for test tokens to mark them as supported
    s_tokenAdminRegistry.proposeAdministrator(s_sourceRouter.getWrappedNative(), OWNER);
    s_tokenAdminRegistry.acceptAdminRole(s_sourceRouter.getWrappedNative());
    s_tokenAdminRegistry.proposeAdministrator(CUSTOM_TOKEN, OWNER);
    s_tokenAdminRegistry.acceptAdminRole(CUSTOM_TOKEN);

    // LockReleaseTokenPool wrappedNativePool = new LockReleaseTokenPool(
    //   IERC20(s_sourceRouter.getWrappedNative()), new address[](0), address(s_mockRMN), true, address(s_sourceRouter)
    // );

    // TokenPool.ChainUpdate[] memory wrappedNativeChainUpdate = new TokenPool.ChainUpdate[](1);
    // wrappedNativeChainUpdate[0] = TokenPool.ChainUpdate({
    //   remoteChainSelector: DEST_CHAIN_SELECTOR,
    //   remotePoolAddress: abi.encode(s_destTokenPool),
    //   remoteTokenAddress: abi.encode(s_destToken),
    //   allowed: true,
    //   outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
    //   inboundRateLimiterConfig: getInboundRateLimiterConfig()
    // });
    // wrappedNativePool.applyChainUpdates(wrappedNativeChainUpdate);
    // s_tokenAdminRegistry.setPool(s_sourceRouter.getWrappedNative(), address(wrappedNativePool));

    // LockReleaseTokenPool customPool = new LockReleaseTokenPool(
    //   IERC20(CUSTOM_TOKEN), new address[](0), address(s_mockRMN), true, address(s_sourceRouter)
    // );
    // TokenPool.ChainUpdate[] memory customChainUpdate = new TokenPool.ChainUpdate[](1);
    // customChainUpdate[0] = TokenPool.ChainUpdate({
    //   remoteChainSelector: DEST_CHAIN_SELECTOR,
    //   remotePoolAddress: abi.encode(s_destTokenPool),
    //   remoteTokenAddress: abi.encode(s_destToken),
    //   allowed: true,
    //   outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
    //   inboundRateLimiterConfig: getInboundRateLimiterConfig()
    // });
    // customPool.applyChainUpdates(customChainUpdate);
    // s_tokenAdminRegistry.setPool(CUSTOM_TOKEN, address(customPool));

    // Ensure the self-serve token is set up on the admin registry
    // vm.mockCall(
    //   address(s_tokenAdminRegistry),
    //   abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_selfServeTokenDefaultPricing),
    //   abi.encode(makeAddr("self-serve-pool"))
    // );
  }

  function _generateEmptyMessage() public view returns (Client.EVM2AnyMessage memory) {
    return Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
  }

  function _generateSingleTokenMessage(
    address token,
    uint256 amount
  ) public view returns (Client.EVM2AnyMessage memory) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    return Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: tokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
  }

  function calcUSDValueFromTokenAmount(uint224 tokenPrice, uint256 tokenAmount) internal pure returns (uint256) {
    return (tokenPrice * tokenAmount) / 1e18;
  }

  function applyBpsRatio(uint256 tokenAmount, uint16 ratio) internal pure returns (uint256) {
    return (tokenAmount * ratio) / 1e5;
  }

  function configUSDCentToWei(uint256 usdCent) internal pure returns (uint256) {
    return usdCent * 1e16;
  }
}

contract PriceRegistry_constructor is PriceRegistrySetup {
  function test_Setup_Success() public virtual {
    address[] memory priceUpdaters = new address[](2);
    priceUpdaters[0] = STRANGER;
    priceUpdaters[1] = OWNER;
    address[] memory feeTokens = new address[](2);
    feeTokens[0] = s_sourceTokens[0];
    feeTokens[1] = s_sourceTokens[1];
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](2);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]], 18);
    tokenPriceFeedUpdates[1] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[1], s_dataFeedByToken[s_sourceTokens[1]], 6);

    vm.expectEmit();
    emit PriceRegistry.PremiumMultiplierWeiPerEthUpdated(
      s_priceRegistryPremiumMultiplierWeiPerEthArgs[0].token,
      s_priceRegistryPremiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth
    );
    vm.expectEmit();
    emit PriceRegistry.PremiumMultiplierWeiPerEthUpdated(
      s_priceRegistryPremiumMultiplierWeiPerEthArgs[1].token,
      s_priceRegistryPremiumMultiplierWeiPerEthArgs[1].premiumMultiplierWeiPerEth
    );

    s_priceRegistry = new PriceRegistryHelper(
      PriceRegistry.StaticConfig({linkToken: s_sourceTokens[0], maxFeeJuelsPerMsg: MAX_MSG_FEES_JUELS}),
      priceUpdaters,
      feeTokens,
      uint32(TWELVE_HOURS),
      tokenPriceFeedUpdates,
      s_priceRegistryTokenTransferFeeConfigArgs,
      s_priceRegistryPremiumMultiplierWeiPerEthArgs,
      _generatePriceRegistryDestChainDynamicConfigArgs()
    );

    // TODO: assert staticConfig equality
    // TODO: assert dynamicConfig equality
    // TODO: assert tokenTransferFeeConfig equality

    assertEq(feeTokens, s_priceRegistry.getFeeTokens());
    assertEq(uint32(TWELVE_HOURS), s_priceRegistry.getStalenessThreshold());
    assertEq(priceUpdaters, s_priceRegistry.getAllAuthorizedCallers());
    assertEq(s_priceRegistry.typeAndVersion(), "PriceRegistry 1.6.0-dev");

    _assertTokenPriceFeedConfigEquality(
      tokenPriceFeedUpdates[0].feedConfig, s_priceRegistry.getTokenPriceFeedConfig(s_sourceTokens[0])
    );

    _assertTokenPriceFeedConfigEquality(
      tokenPriceFeedUpdates[1].feedConfig, s_priceRegistry.getTokenPriceFeedConfig(s_sourceTokens[1])
    );

    uint64 gotFeeTokenConfig0 =
      s_priceRegistry.getPremiumMultiplierWeiPerEth(s_priceRegistryPremiumMultiplierWeiPerEthArgs[0].token);
    assertEq(s_priceRegistryPremiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth, gotFeeTokenConfig0);

    uint64 gotFeeTokenConfig1 =
      s_priceRegistry.getPremiumMultiplierWeiPerEth(s_priceRegistryPremiumMultiplierWeiPerEthArgs[1].token);
    assertEq(s_priceRegistryPremiumMultiplierWeiPerEthArgs[1].premiumMultiplierWeiPerEth, gotFeeTokenConfig1);
  }

  // TODO: re-add test
  // function test_InvalidStalenessThreshold_Revert() public {
  //   vm.expectRevert(PriceRegistry.InvalidStalenessThreshold.selector);
  //   s_priceRegistry =
  //     new PriceRegistryHelper(new address[](0), new address[](0), 0, new PriceRegistry.TokenPriceFeedUpdate[](0));
  // }

  // TODO: re-add test
  // function test_Constructor_InvalidConfigLinkTokenEqAddressZero_Revert() public {
  //   vm.expectRevert(EVM2EVMMultiOnRamp.InvalidConfig.selector);
  //   new EVM2EVMMultiOnRampHelper(
  //     EVM2EVMMultiOnRamp.StaticConfig({
  //       linkToken: address(0),
  //       chainSelector: SOURCE_CHAIN_SELECTOR,
  //       maxFeeJuelsPerMsg: MAX_NOP_FEES_JUELS,
  //       rmnProxy: address(s_mockRMN),
  //       nonceManager: address(s_outboundNonceManager),
  //       tokenAdminRegistry: address(s_tokenAdminRegistry)
  //     }),
  //     _generateDynamicMultiOnRampConfig(address(s_sourceRouter), address(s_priceRegistry)),
  //     _generateDestChainConfigArgs(),
  //     s_premiumMultiplierWeiPerEthArgs,
  //     s_tokenTransferFeeConfigArgs
  //   );
  // }
}

contract PriceRegistry_getTokenPrices is PriceRegistrySetup {
  function test_GetTokenPrices_Success() public view {
    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));

    address[] memory tokens = new address[](3);
    tokens[0] = s_sourceTokens[0];
    tokens[1] = s_sourceTokens[1];
    tokens[2] = s_weth;

    Internal.TimestampedPackedUint224[] memory tokenPrices = s_priceRegistry.getTokenPrices(tokens);

    assertEq(tokenPrices.length, 3);
    assertEq(tokenPrices[0].value, priceUpdates.tokenPriceUpdates[0].usdPerToken);
    assertEq(tokenPrices[1].value, priceUpdates.tokenPriceUpdates[1].usdPerToken);
    assertEq(tokenPrices[2].value, priceUpdates.tokenPriceUpdates[2].usdPerToken);
  }
}

contract PriceRegistry_getTokenPrice is PriceRegistrySetup {
  function test_GetTokenPriceFromFeed_Success() public {
    uint256 originalTimestampValue = block.timestamp;

    // Below staleness threshold
    vm.warp(originalTimestampValue + 1 hours);

    address sourceToken = _initialiseSingleTokenPriceFeed();
    Internal.TimestampedPackedUint224 memory tokenPriceAnswer = s_priceRegistry.getTokenPrice(sourceToken);

    // Price answer is 1e8 (18 decimal token) - unit is (1e18 * 1e18 / 1e18) -> expected 1e18
    assertEq(tokenPriceAnswer.value, uint224(1e18));
    assertEq(tokenPriceAnswer.timestamp, uint32(block.timestamp));
  }
}

contract PriceRegistry_getValidatedTokenPrice is PriceRegistrySetup {
  function test_GetValidatedTokenPrice_Success() public view {
    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));
    address token = priceUpdates.tokenPriceUpdates[0].sourceToken;

    uint224 tokenPrice = s_priceRegistry.getValidatedTokenPrice(token);

    assertEq(priceUpdates.tokenPriceUpdates[0].usdPerToken, tokenPrice);
  }

  function test_GetValidatedTokenPriceFromFeed_Success() public {
    uint256 originalTimestampValue = block.timestamp;

    // Right below staleness threshold
    vm.warp(originalTimestampValue + TWELVE_HOURS);

    address sourceToken = _initialiseSingleTokenPriceFeed();
    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(sourceToken);

    // Price answer is 1e8 (18 decimal token) - unit is (1e18 * 1e18 / 1e18) -> expected 1e18
    assertEq(tokenPriceAnswer, uint224(1e18));
  }

  function test_GetValidatedTokenPriceFromFeedOverStalenessPeriod_Success() public {
    uint256 originalTimestampValue = block.timestamp;

    // Right above staleness threshold
    vm.warp(originalTimestampValue + TWELVE_HOURS + 1);

    address sourceToken = _initialiseSingleTokenPriceFeed();
    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(sourceToken);

    // Price answer is 1e8 (18 decimal token) - unit is (1e18 * 1e18 / 1e18) -> expected 1e18
    assertEq(tokenPriceAnswer, uint224(1e18));
  }

  function test_GetValidatedTokenPriceFromFeedMaxInt224Value_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, int256(uint256(type(uint224).max)));

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 18);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is: uint224.MAX_VALUE * (10 ** (36 - 18 - 18))
    assertEq(tokenPriceAnswer, uint224(type(uint224).max));
  }

  function test_GetValidatedTokenPriceFromFeedErc20Below18Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 6);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 8, 1e8);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 6);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e6) -> expected 1e30
    assertEq(tokenPriceAnswer, uint224(1e30));
  }

  function test_GetValidatedTokenPriceFromFeedErc20Above18Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 24);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 8, 1e8);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 24);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e24) -> expected 1e12
    assertEq(tokenPriceAnswer, uint224(1e12));
  }

  function test_GetValidatedTokenPriceFromFeedFeedAt18Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, 1e18);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 18);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e18) -> expected 1e18
    assertEq(tokenPriceAnswer, uint224(1e18));
  }

  function test_GetValidatedTokenPriceFromFeedFeedAt0Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 0);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 0, 1e31);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 0);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e31 (0 decimal token) - unit is (1e18 * 1e18 / 1e0) -> expected 1e36
    assertEq(tokenPriceAnswer, uint224(1e67));
  }

  function test_GetValidatedTokenPriceFromFeedFlippedDecimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 20);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 20, 1e18);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 20);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e20) -> expected 1e14
    assertEq(tokenPriceAnswer, uint224(1e14));
  }

  function test_StaleFeeToken_Success() public {
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));
    address token = priceUpdates.tokenPriceUpdates[0].sourceToken;

    uint224 tokenPrice = s_priceRegistry.getValidatedTokenPrice(token);

    assertEq(priceUpdates.tokenPriceUpdates[0].usdPerToken, tokenPrice);
  }

  // Reverts

  function test_OverflowFeedPrice_Revert() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, int256(uint256(type(uint224).max) + 1));

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 18);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    vm.expectRevert(PriceRegistry.DataFeedValueOutOfUint224Range.selector);
    s_priceRegistry.getValidatedTokenPrice(tokenAddress);
  }

  function test_UnderflowFeedPrice_Revert() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, -1);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress, 18);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    vm.expectRevert(PriceRegistry.DataFeedValueOutOfUint224Range.selector);
    s_priceRegistry.getValidatedTokenPrice(tokenAddress);
  }

  function test_TokenNotSupported_Revert() public {
    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.getValidatedTokenPrice(DUMMY_CONTRACT_ADDRESS);
  }

  function test_TokenNotSupportedFeed_Revert() public {
    address sourceToken = _initialiseSingleTokenPriceFeed();
    MockV3Aggregator(s_dataFeedByToken[sourceToken]).updateAnswer(0);

    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.TokenNotSupported.selector, sourceToken));
    s_priceRegistry.getValidatedTokenPrice(sourceToken);
  }
}

contract PriceRegistry_applyFeeTokensUpdates is PriceRegistrySetup {
  function test_ApplyFeeTokensUpdates_Success() public {
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceTokens[1];

    vm.expectEmit();
    emit PriceRegistry.FeeTokenAdded(feeTokens[0]);

    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));
    assertEq(s_priceRegistry.getFeeTokens().length, 3);
    assertEq(s_priceRegistry.getFeeTokens()[2], feeTokens[0]);

    // add same feeToken is no-op
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));
    assertEq(s_priceRegistry.getFeeTokens().length, 3);
    assertEq(s_priceRegistry.getFeeTokens()[2], feeTokens[0]);

    vm.expectEmit();
    emit PriceRegistry.FeeTokenRemoved(feeTokens[0]);

    s_priceRegistry.applyFeeTokensUpdates(new address[](0), feeTokens);
    assertEq(s_priceRegistry.getFeeTokens().length, 2);

    // removing already removed feeToken is no-op
    s_priceRegistry.applyFeeTokensUpdates(new address[](0), feeTokens);
    assertEq(s_priceRegistry.getFeeTokens().length, 2);
  }

  function test_OnlyCallableByOwner_Revert() public {
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = STRANGER;
    vm.startPrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));
  }
}

contract PriceRegistry_updatePrices is PriceRegistrySetup {
  function test_OnlyTokenPrice_Success() public {
    Internal.PriceUpdates memory update = Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](1),
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    });
    update.tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceTokens[0], usdPerToken: 4e18});

    vm.expectEmit();
    emit PriceRegistry.UsdPerTokenUpdated(
      update.tokenPriceUpdates[0].sourceToken, update.tokenPriceUpdates[0].usdPerToken, block.timestamp
    );

    s_priceRegistry.updatePrices(update);

    assertEq(s_priceRegistry.getTokenPrice(s_sourceTokens[0]).value, update.tokenPriceUpdates[0].usdPerToken);
  }

  function test_OnlyGasPrice_Success() public {
    Internal.PriceUpdates memory update = Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](0),
      gasPriceUpdates: new Internal.GasPriceUpdate[](1)
    });
    update.gasPriceUpdates[0] =
      Internal.GasPriceUpdate({destChainSelector: DEST_CHAIN_SELECTOR, usdPerUnitGas: 2000e18});

    vm.expectEmit();
    emit PriceRegistry.UsdPerUnitGasUpdated(
      update.gasPriceUpdates[0].destChainSelector, update.gasPriceUpdates[0].usdPerUnitGas, block.timestamp
    );

    s_priceRegistry.updatePrices(update);

    assertEq(
      s_priceRegistry.getDestinationChainGasPrice(DEST_CHAIN_SELECTOR).value, update.gasPriceUpdates[0].usdPerUnitGas
    );
  }

  function test_UpdateMultiplePrices_Success() public {
    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](3);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceTokens[0], usdPerToken: 4e18});
    tokenPriceUpdates[1] = Internal.TokenPriceUpdate({sourceToken: s_sourceTokens[1], usdPerToken: 1800e18});
    tokenPriceUpdates[2] = Internal.TokenPriceUpdate({sourceToken: address(12345), usdPerToken: 1e18});

    Internal.GasPriceUpdate[] memory gasPriceUpdates = new Internal.GasPriceUpdate[](3);
    gasPriceUpdates[0] = Internal.GasPriceUpdate({destChainSelector: DEST_CHAIN_SELECTOR, usdPerUnitGas: 2e6});
    gasPriceUpdates[1] = Internal.GasPriceUpdate({destChainSelector: SOURCE_CHAIN_SELECTOR, usdPerUnitGas: 2000e18});
    gasPriceUpdates[2] = Internal.GasPriceUpdate({destChainSelector: 12345, usdPerUnitGas: 1e18});

    Internal.PriceUpdates memory update =
      Internal.PriceUpdates({tokenPriceUpdates: tokenPriceUpdates, gasPriceUpdates: gasPriceUpdates});

    for (uint256 i = 0; i < tokenPriceUpdates.length; ++i) {
      vm.expectEmit();
      emit PriceRegistry.UsdPerTokenUpdated(
        update.tokenPriceUpdates[i].sourceToken, update.tokenPriceUpdates[i].usdPerToken, block.timestamp
      );
    }
    for (uint256 i = 0; i < gasPriceUpdates.length; ++i) {
      vm.expectEmit();
      emit PriceRegistry.UsdPerUnitGasUpdated(
        update.gasPriceUpdates[i].destChainSelector, update.gasPriceUpdates[i].usdPerUnitGas, block.timestamp
      );
    }

    s_priceRegistry.updatePrices(update);

    for (uint256 i = 0; i < tokenPriceUpdates.length; ++i) {
      assertEq(
        s_priceRegistry.getTokenPrice(update.tokenPriceUpdates[i].sourceToken).value, tokenPriceUpdates[i].usdPerToken
      );
    }
    for (uint256 i = 0; i < gasPriceUpdates.length; ++i) {
      assertEq(
        s_priceRegistry.getDestinationChainGasPrice(update.gasPriceUpdates[i].destChainSelector).value,
        gasPriceUpdates[i].usdPerUnitGas
      );
    }
  }

  function test_UpdatableByAuthorizedCaller_Success() public {
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](1),
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    });
    priceUpdates.tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceTokens[0], usdPerToken: 4e18});

    // Revert when caller is not authorized
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_priceRegistry.updatePrices(priceUpdates);

    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;
    vm.startPrank(OWNER);
    s_priceRegistry.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: priceUpdaters, removedCallers: new address[](0)})
    );

    // Stranger is now an authorized caller to update prices
    vm.expectEmit();
    emit PriceRegistry.UsdPerTokenUpdated(
      priceUpdates.tokenPriceUpdates[0].sourceToken, priceUpdates.tokenPriceUpdates[0].usdPerToken, block.timestamp
    );
    s_priceRegistry.updatePrices(priceUpdates);

    assertEq(s_priceRegistry.getTokenPrice(s_sourceTokens[0]).value, priceUpdates.tokenPriceUpdates[0].usdPerToken);

    vm.startPrank(OWNER);
    s_priceRegistry.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: priceUpdaters})
    );

    // Revert when authorized caller is removed
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_priceRegistry.updatePrices(priceUpdates);
  }

  // Reverts

  function test_OnlyCallableByUpdaterOrOwner_Revert() public {
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](0),
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    });

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_priceRegistry.updatePrices(priceUpdates);
  }
}

contract PriceRegistry_convertTokenAmount is PriceRegistrySetup {
  function test_ConvertTokenAmount_Success() public view {
    Internal.PriceUpdates memory initialPriceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));
    uint256 amount = 3e16;
    uint256 conversionRate = (uint256(initialPriceUpdates.tokenPriceUpdates[2].usdPerToken) * 1e18)
      / uint256(initialPriceUpdates.tokenPriceUpdates[0].usdPerToken);
    uint256 expected = (amount * conversionRate) / 1e18;
    assertEq(s_priceRegistry.convertTokenAmount(s_weth, amount, s_sourceTokens[0]), expected);
  }

  function test_Fuzz_ConvertTokenAmount_Success(
    uint256 feeTokenAmount,
    uint224 usdPerFeeToken,
    uint160 usdPerLinkToken,
    uint224 usdPerUnitGas
  ) public {
    vm.assume(usdPerFeeToken > 0);
    vm.assume(usdPerLinkToken > 0);
    // We bound the max fees to be at most uint96.max link.
    feeTokenAmount = bound(feeTokenAmount, 0, (uint256(type(uint96).max) * usdPerLinkToken) / usdPerFeeToken);

    address feeToken = address(1);
    address linkToken = address(2);
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = feeToken;
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));

    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](2);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: feeToken, usdPerToken: usdPerFeeToken});
    tokenPriceUpdates[1] = Internal.TokenPriceUpdate({sourceToken: linkToken, usdPerToken: usdPerLinkToken});

    Internal.GasPriceUpdate[] memory gasPriceUpdates = new Internal.GasPriceUpdate[](1);
    gasPriceUpdates[0] = Internal.GasPriceUpdate({destChainSelector: DEST_CHAIN_SELECTOR, usdPerUnitGas: usdPerUnitGas});

    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: tokenPriceUpdates, gasPriceUpdates: gasPriceUpdates});

    s_priceRegistry.updatePrices(priceUpdates);

    uint256 linkFee = s_priceRegistry.convertTokenAmount(feeToken, feeTokenAmount, linkToken);
    assertEq(linkFee, (feeTokenAmount * usdPerFeeToken) / usdPerLinkToken);
  }

  // Reverts

  function test_LinkTokenNotSupported_Revert() public {
    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.convertTokenAmount(DUMMY_CONTRACT_ADDRESS, 3e16, s_sourceTokens[0]);

    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.convertTokenAmount(s_sourceTokens[0], 3e16, DUMMY_CONTRACT_ADDRESS);
  }
}

contract PriceRegistry_getTokenAndGasPrices is PriceRegistrySetup {
  function test_GetFeeTokenAndGasPrices_Success() public view {
    (uint224 feeTokenPrice, uint224 gasPrice) =
      s_priceRegistry.getTokenAndGasPrices(s_sourceFeeToken, DEST_CHAIN_SELECTOR);

    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));

    assertEq(feeTokenPrice, s_sourceTokenPrices[0]);
    assertEq(gasPrice, priceUpdates.gasPriceUpdates[0].usdPerUnitGas);
  }

  function test_ZeroGasPrice_Success() public {
    uint64 zeroGasDestChainSelector = 345678;
    Internal.GasPriceUpdate[] memory gasPriceUpdates = new Internal.GasPriceUpdate[](1);
    gasPriceUpdates[0] = Internal.GasPriceUpdate({destChainSelector: zeroGasDestChainSelector, usdPerUnitGas: 0});

    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: new Internal.TokenPriceUpdate[](0), gasPriceUpdates: gasPriceUpdates});
    s_priceRegistry.updatePrices(priceUpdates);

    (, uint224 gasPrice) = s_priceRegistry.getTokenAndGasPrices(s_sourceFeeToken, zeroGasDestChainSelector);

    assertEq(gasPrice, priceUpdates.gasPriceUpdates[0].usdPerUnitGas);
  }

  function test_UnsupportedChain_Revert() public {
    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.ChainNotSupported.selector, DEST_CHAIN_SELECTOR + 1));
    s_priceRegistry.getTokenAndGasPrices(s_sourceTokens[0], DEST_CHAIN_SELECTOR + 1);
  }

  function test_StaleGasPrice_Revert() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);
    vm.expectRevert(
      abi.encodeWithSelector(PriceRegistry.StaleGasPrice.selector, DEST_CHAIN_SELECTOR, TWELVE_HOURS, diff)
    );
    s_priceRegistry.getTokenAndGasPrices(s_sourceTokens[0], DEST_CHAIN_SELECTOR);
  }
}

contract PriceRegistry_updateTokenPriceFeeds is PriceRegistrySetup {
  function test_ZeroFeeds_Success() public {
    Vm.Log[] memory logEntries = vm.getRecordedLogs();

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](0);
    vm.recordLogs();
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    // Verify no log emissions
    assertEq(logEntries.length, 0);
  }

  function test_SingleFeedUpdate_Success() public {
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]], 18);

    _assertTokenPriceFeedConfigUnconfigured(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[0].sourceToken)
    );

    vm.expectEmit();
    emit PriceRegistry.PriceFeedPerTokenUpdated(
      tokenPriceFeedUpdates[0].sourceToken, tokenPriceFeedUpdates[0].feedConfig
    );

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    _assertTokenPriceFeedConfigEquality(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].feedConfig
    );
  }

  function test_MultipleFeedUpdate_Success() public {
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](2);

    for (uint256 i = 0; i < 2; ++i) {
      tokenPriceFeedUpdates[i] =
        getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[i], s_dataFeedByToken[s_sourceTokens[i]], 18);

      _assertTokenPriceFeedConfigUnconfigured(
        s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[i].sourceToken)
      );

      vm.expectEmit();
      emit PriceRegistry.PriceFeedPerTokenUpdated(
        tokenPriceFeedUpdates[i].sourceToken, tokenPriceFeedUpdates[i].feedConfig
      );
    }

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    _assertTokenPriceFeedConfigEquality(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].feedConfig
    );
    _assertTokenPriceFeedConfigEquality(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[1].sourceToken), tokenPriceFeedUpdates[1].feedConfig
    );
  }

  function test_FeedUnset_Success() public {
    Internal.TimestampedPackedUint224 memory priceQueryInitial = s_priceRegistry.getTokenPrice(s_sourceTokens[0]);
    assertFalse(priceQueryInitial.value == 0);
    assertFalse(priceQueryInitial.timestamp == 0);

    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]], 18);

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    _assertTokenPriceFeedConfigEquality(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].feedConfig
    );

    tokenPriceFeedUpdates[0].feedConfig.dataFeedAddress = address(0);
    vm.expectEmit();
    emit PriceRegistry.PriceFeedPerTokenUpdated(
      tokenPriceFeedUpdates[0].sourceToken, tokenPriceFeedUpdates[0].feedConfig
    );

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    _assertTokenPriceFeedConfigEquality(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].feedConfig
    );

    // Price data should remain after a feed has been set->unset
    Internal.TimestampedPackedUint224 memory priceQueryPostUnsetFeed = s_priceRegistry.getTokenPrice(s_sourceTokens[0]);
    assertEq(priceQueryPostUnsetFeed.value, priceQueryInitial.value);
    assertEq(priceQueryPostUnsetFeed.timestamp, priceQueryInitial.timestamp);
  }

  function test_FeedNotUpdated() public {
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]], 18);

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    _assertTokenPriceFeedConfigEquality(
      s_priceRegistry.getTokenPriceFeedConfig(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].feedConfig
    );
  }

  // Reverts

  function test_FeedUpdatedByNonOwner_Revert() public {
    PriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new PriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]], 18);

    vm.startPrank(STRANGER);
    vm.expectRevert("Only callable by owner");

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
  }
}

// TODO: add dest chain config tests
// contract EVM2EVMMultiOnRamp_applyDestChainConfigUpdates is EVM2EVMMultiOnRampSetup {
//   function test_Fuzz_applyDestChainConfigUpdates_Success(
//     EVM2EVMMultiOnRamp.DestChainConfigArgs memory destChainConfigArgs
//   ) public {
//     vm.assume(destChainConfigArgs.destChainSelector != 0);
//     vm.assume(destChainConfigArgs.dynamicConfig.defaultTxGasLimit != 0);
//     destChainConfigArgs.dynamicConfig.defaultTokenDestBytesOverhead = uint32(
//       bound(
//         destChainConfigArgs.dynamicConfig.defaultTokenDestBytesOverhead,
//         Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
//         type(uint32).max
//       )
//     );
//     destChainConfigArgs.dynamicConfig.chainFamilySelector = Internal.CHAIN_FAMILY_SELECTOR_EVM;

//     bool isNewChain = true;

//     if (destChainConfigArgs.destChainSelector == DEST_CHAIN_SELECTOR) {
//       destChainConfigArgs.prevOnRamp = address(0);
//       isNewChain = false;
//     }
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory newDestChainConfigArgs =
//       new EVM2EVMMultiOnRamp.DestChainConfigArgs[](1);
//     newDestChainConfigArgs[0] = destChainConfigArgs;
//     EVM2EVMMultiOnRamp.DestChainConfig memory expectedDestChainConfig = EVM2EVMMultiOnRamp.DestChainConfig({
//       dynamicConfig: destChainConfigArgs.dynamicConfig,
//       prevOnRamp: destChainConfigArgs.prevOnRamp,
//       sequenceNumber: 0,
//       metadataHash: keccak256(
//         abi.encode(
//           Internal.EVM_2_ANY_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, destChainConfigArgs.destChainSelector, address(s_onRamp)
//         )
//         )
//     });

//     if (isNewChain) {
//       vm.expectEmit();
//       emit EVM2EVMMultiOnRamp.DestChainAdded(destChainConfigArgs.destChainSelector, expectedDestChainConfig);
//     } else {
//       vm.expectEmit();
//       emit EVM2EVMMultiOnRamp.DestChainDynamicConfigUpdated(
//         destChainConfigArgs.destChainSelector, expectedDestChainConfig.dynamicConfig
//       );
//     }

//     s_onRamp.applyDestChainConfigUpdates(newDestChainConfigArgs);

//     _assertDestChainConfigsEqual(
//       expectedDestChainConfig, s_onRamp.getDestChainConfig(destChainConfigArgs.destChainSelector)
//     );
//   }

//   function test_applyDestChainConfigUpdates_Success() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs =
//       new EVM2EVMMultiOnRamp.DestChainConfigArgs[](2);
//     destChainConfigArgs[0] = _generateDestChainConfigArgs()[0];
//     destChainConfigArgs[0].dynamicConfig.isEnabled = false;
//     destChainConfigArgs[1] = _generateDestChainConfigArgs()[0];
//     destChainConfigArgs[1].destChainSelector = DEST_CHAIN_SELECTOR + 1;

//     EVM2EVMMultiOnRamp.DestChainConfig memory expectedDestChainConfig0 = EVM2EVMMultiOnRamp.DestChainConfig({
//       dynamicConfig: destChainConfigArgs[0].dynamicConfig,
//       prevOnRamp: address(0),
//       sequenceNumber: 0,
//       metadataHash: keccak256(
//         abi.encode(
//           Internal.EVM_2_ANY_MESSAGE_HASH,
//           SOURCE_CHAIN_SELECTOR,
//           destChainConfigArgs[0].destChainSelector,
//           address(s_onRamp)
//         )
//         )
//     });

//     EVM2EVMMultiOnRamp.DestChainConfig memory expectedDestChainConfig1 = EVM2EVMMultiOnRamp.DestChainConfig({
//       dynamicConfig: destChainConfigArgs[1].dynamicConfig,
//       prevOnRamp: address(0),
//       sequenceNumber: 0,
//       metadataHash: keccak256(
//         abi.encode(
//           Internal.EVM_2_ANY_MESSAGE_HASH,
//           SOURCE_CHAIN_SELECTOR,
//           destChainConfigArgs[1].destChainSelector,
//           address(s_onRamp)
//         )
//         )
//     });

//     vm.expectEmit();
//     emit EVM2EVMMultiOnRamp.DestChainDynamicConfigUpdated(DEST_CHAIN_SELECTOR, expectedDestChainConfig0.dynamicConfig);
//     vm.expectEmit();
//     emit EVM2EVMMultiOnRamp.DestChainAdded(DEST_CHAIN_SELECTOR + 1, expectedDestChainConfig1);

//     vm.recordLogs();
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

//     EVM2EVMMultiOnRamp.DestChainConfig memory gotDestChainConfig0 = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
//     EVM2EVMMultiOnRamp.DestChainConfig memory gotDestChainConfig1 = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR + 1);

//     assertEq(vm.getRecordedLogs().length, 2);
//     _assertDestChainConfigsEqual(expectedDestChainConfig0, gotDestChainConfig0);
//     _assertDestChainConfigsEqual(expectedDestChainConfig1, gotDestChainConfig1);
//   }

//   function test_applyDestChainConfigUpdatesZeroIntput() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs =
//       new EVM2EVMMultiOnRamp.DestChainConfigArgs[](0);

//     vm.recordLogs();
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

//     assertEq(vm.getRecordedLogs().length, 0);
//   }

//   // Reverts

//   function test_InvalidDestChainConfigDestChainSelectorEqZero_Revert() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//     EVM2EVMMultiOnRamp.DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[0];

//     destChainConfigArg.destChainSelector = 0;
//     vm.expectRevert(
//       abi.encodeWithSelector(EVM2EVMMultiOnRamp.InvalidDestChainConfig.selector, destChainConfigArg.destChainSelector)
//     );
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
//   }

//   function test_applyDestChainConfigUpdatesDefaultTxGasLimitEqZero_Revert() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//     EVM2EVMMultiOnRamp.DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[0];

//     destChainConfigArg.dynamicConfig.defaultTxGasLimit = 0;
//     vm.expectRevert(
//       abi.encodeWithSelector(EVM2EVMMultiOnRamp.InvalidDestChainConfig.selector, destChainConfigArg.destChainSelector)
//     );
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
//   }

//   function test_InvalidDestChainConfigNewPrevOnRampOnExistingChain_Revert() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//     EVM2EVMMultiOnRamp.DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[0];

//     destChainConfigArg.prevOnRamp = address(1);
//     vm.expectRevert(
//       abi.encodeWithSelector(EVM2EVMMultiOnRamp.InvalidDestChainConfig.selector, destChainConfigArg.destChainSelector)
//     );
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
//   }

//   function test_InvalidDestBytesOverhead_Revert() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//     EVM2EVMMultiOnRamp.DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[0];

//     destChainConfigArg.dynamicConfig.defaultTokenDestBytesOverhead = uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES - 1);

//     vm.expectRevert(
//       abi.encodeWithSelector(
//         EVM2EVMMultiOnRamp.InvalidDestBytesOverhead.selector,
//         address(0),
//         destChainConfigArg.dynamicConfig.defaultTokenDestBytesOverhead
//       )
//     );

//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
//   }

//   function test_InvalidChainFamilySelector_Revert() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//     EVM2EVMMultiOnRamp.DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[0];

//     destChainConfigArg.dynamicConfig.chainFamilySelector = bytes4(uint32(1));

//     vm.expectRevert(
//       abi.encodeWithSelector(EVM2EVMMultiOnRamp.InvalidDestChainConfig.selector, destChainConfigArg.destChainSelector)
//     );
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
//   }
// }

contract PriceRegistry_getDataAvailabilityCost is PriceRegistrySetup {
  function test_EmptyMessageCalculatesDataAvailabilityCost_Success() public {
    uint256 dataAvailabilityCostUSD =
      s_priceRegistry.getDataAvailabilityCost(DEST_CHAIN_SELECTOR, USD_PER_DATA_AVAILABILITY_GAS, 0, 0, 0);

    PriceRegistry.DestChainDynamicConfig memory destChainDynamicConfig =
      s_priceRegistry.getDestChainDynamicConfig(DEST_CHAIN_SELECTOR);

    uint256 dataAvailabilityGas = destChainDynamicConfig.destDataAvailabilityOverheadGas
      + destChainDynamicConfig.destGasPerDataAvailabilityByte * Internal.ANY_2_EVM_MESSAGE_FIXED_BYTES;
    uint256 expectedDataAvailabilityCostUSD = USD_PER_DATA_AVAILABILITY_GAS * dataAvailabilityGas
      * destChainDynamicConfig.destDataAvailabilityMultiplierBps * 1e14;

    assertEq(expectedDataAvailabilityCostUSD, dataAvailabilityCostUSD);

    // Test that the cost is destnation chain specific
    PriceRegistry.DestChainDynamicConfigArgs[] memory destChainConfigArgs =
      _generatePriceRegistryDestChainDynamicConfigArgs();
    destChainConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR + 1;
    destChainConfigArgs[0].dynamicConfig.destDataAvailabilityOverheadGas =
      destChainDynamicConfig.destDataAvailabilityOverheadGas * 2;
    destChainConfigArgs[0].dynamicConfig.destGasPerDataAvailabilityByte =
      destChainDynamicConfig.destGasPerDataAvailabilityByte * 2;
    destChainConfigArgs[0].dynamicConfig.destDataAvailabilityMultiplierBps =
      destChainDynamicConfig.destDataAvailabilityMultiplierBps * 2;
    s_priceRegistry.applyDestChainConfigUpdates(destChainConfigArgs);

    destChainDynamicConfig = s_priceRegistry.getDestChainDynamicConfig(DEST_CHAIN_SELECTOR + 1);
    uint256 dataAvailabilityCostUSD2 =
      s_priceRegistry.getDataAvailabilityCost(DEST_CHAIN_SELECTOR + 1, USD_PER_DATA_AVAILABILITY_GAS, 0, 0, 0);
    dataAvailabilityGas = destChainDynamicConfig.destDataAvailabilityOverheadGas
      + destChainDynamicConfig.destGasPerDataAvailabilityByte * Internal.ANY_2_EVM_MESSAGE_FIXED_BYTES;
    expectedDataAvailabilityCostUSD = USD_PER_DATA_AVAILABILITY_GAS * dataAvailabilityGas
      * destChainDynamicConfig.destDataAvailabilityMultiplierBps * 1e14;

    assertEq(expectedDataAvailabilityCostUSD, dataAvailabilityCostUSD2);
    assertFalse(dataAvailabilityCostUSD == dataAvailabilityCostUSD2);
  }

  function test_SimpleMessageCalculatesDataAvailabilityCost_Success() public view {
    uint256 dataAvailabilityCostUSD =
      s_priceRegistry.getDataAvailabilityCost(DEST_CHAIN_SELECTOR, USD_PER_DATA_AVAILABILITY_GAS, 100, 5, 50);

    PriceRegistry.DestChainDynamicConfig memory destChainDynamicConfig =
      s_priceRegistry.getDestChainDynamicConfig(DEST_CHAIN_SELECTOR);

    uint256 dataAvailabilityLengthBytes =
      Internal.ANY_2_EVM_MESSAGE_FIXED_BYTES + 100 + (5 * Internal.ANY_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN) + 50;
    uint256 dataAvailabilityGas = destChainDynamicConfig.destDataAvailabilityOverheadGas
      + destChainDynamicConfig.destGasPerDataAvailabilityByte * dataAvailabilityLengthBytes;
    uint256 expectedDataAvailabilityCostUSD = USD_PER_DATA_AVAILABILITY_GAS * dataAvailabilityGas
      * destChainDynamicConfig.destDataAvailabilityMultiplierBps * 1e14;

    assertEq(expectedDataAvailabilityCostUSD, dataAvailabilityCostUSD);
  }

  function test_SimpleMessageCalculatesDataAvailabilityCostUnsupportedDestChainSelector_Success() public view {
    uint256 dataAvailabilityCostUSD =
      s_priceRegistry.getDataAvailabilityCost(0, USD_PER_DATA_AVAILABILITY_GAS, 100, 5, 50);

    assertEq(dataAvailabilityCostUSD, 0);
  }

  function test_Fuzz_ZeroDataAvailabilityGasPriceAlwaysCalculatesZeroDataAvailabilityCost_Success(
    uint64 messageDataLength,
    uint32 numberOfTokens,
    uint32 tokenTransferBytesOverhead
  ) public view {
    uint256 dataAvailabilityCostUSD = s_priceRegistry.getDataAvailabilityCost(
      DEST_CHAIN_SELECTOR, 0, messageDataLength, numberOfTokens, tokenTransferBytesOverhead
    );

    assertEq(0, dataAvailabilityCostUSD);
  }

  function test_Fuzz_CalculateDataAvailabilityCost_Success(
    uint64 destChainSelector,
    uint32 destDataAvailabilityOverheadGas,
    uint16 destGasPerDataAvailabilityByte,
    uint16 destDataAvailabilityMultiplierBps,
    uint112 dataAvailabilityGasPrice,
    uint64 messageDataLength,
    uint32 numberOfTokens,
    uint32 tokenTransferBytesOverhead
  ) public {
    vm.assume(destChainSelector != 0);
    PriceRegistry.DestChainDynamicConfigArgs[] memory destChainConfigArgs =
      new PriceRegistry.DestChainDynamicConfigArgs[](1);
    PriceRegistry.DestChainDynamicConfig memory destChainConfig =
      s_priceRegistry.getDestChainDynamicConfig(destChainSelector);
    destChainConfigArgs[0] =
      PriceRegistry.DestChainDynamicConfigArgs({destChainSelector: destChainSelector, dynamicConfig: destChainConfig});
    destChainConfigArgs[0].dynamicConfig.destDataAvailabilityOverheadGas = destDataAvailabilityOverheadGas;
    destChainConfigArgs[0].dynamicConfig.destGasPerDataAvailabilityByte = destGasPerDataAvailabilityByte;
    destChainConfigArgs[0].dynamicConfig.destDataAvailabilityMultiplierBps = destDataAvailabilityMultiplierBps;
    destChainConfigArgs[0].dynamicConfig.defaultTxGasLimit = GAS_LIMIT;
    destChainConfigArgs[0].dynamicConfig.chainFamilySelector = Internal.CHAIN_FAMILY_SELECTOR_EVM;
    destChainConfigArgs[0].dynamicConfig.defaultTokenDestBytesOverhead = DEFAULT_TOKEN_BYTES_OVERHEAD;

    s_priceRegistry.applyDestChainConfigUpdates(destChainConfigArgs);

    uint256 dataAvailabilityCostUSD = s_priceRegistry.getDataAvailabilityCost(
      destChainConfigArgs[0].destChainSelector,
      dataAvailabilityGasPrice,
      messageDataLength,
      numberOfTokens,
      tokenTransferBytesOverhead
    );

    uint256 dataAvailabilityLengthBytes = Internal.ANY_2_EVM_MESSAGE_FIXED_BYTES + messageDataLength
      + (numberOfTokens * Internal.ANY_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN) + tokenTransferBytesOverhead;

    uint256 dataAvailabilityGas =
      destDataAvailabilityOverheadGas + destGasPerDataAvailabilityByte * dataAvailabilityLengthBytes;
    uint256 expectedDataAvailabilityCostUSD =
      dataAvailabilityGasPrice * dataAvailabilityGas * destDataAvailabilityMultiplierBps * 1e14;

    assertEq(expectedDataAvailabilityCostUSD, dataAvailabilityCostUSD);
  }
}

contract PriceRegistry_applyPremiumMultiplierWeiPerEthUpdates is PriceRegistrySetup {
  function test_Fuzz_applyPremiumMultiplierWeiPerEthUpdates_Success(
    PriceRegistry.PremiumMultiplierWeiPerEthArgs memory premiumMultiplierWeiPerEthArg
  ) public {
    PriceRegistry.PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs =
      new PriceRegistry.PremiumMultiplierWeiPerEthArgs[](1);
    premiumMultiplierWeiPerEthArgs[0] = premiumMultiplierWeiPerEthArg;

    vm.expectEmit();
    emit PriceRegistry.PremiumMultiplierWeiPerEthUpdated(
      premiumMultiplierWeiPerEthArg.token, premiumMultiplierWeiPerEthArg.premiumMultiplierWeiPerEth
    );

    s_priceRegistry.applyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs);

    assertEq(
      premiumMultiplierWeiPerEthArg.premiumMultiplierWeiPerEth,
      s_priceRegistry.getPremiumMultiplierWeiPerEth(premiumMultiplierWeiPerEthArg.token)
    );
  }

  function test_applyPremiumMultiplierWeiPerEthUpdatesSingleToken_Success() public {
    PriceRegistry.PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs =
      new PriceRegistry.PremiumMultiplierWeiPerEthArgs[](1);
    premiumMultiplierWeiPerEthArgs[0] = s_priceRegistryPremiumMultiplierWeiPerEthArgs[0];
    premiumMultiplierWeiPerEthArgs[0].token = vm.addr(1);

    vm.expectEmit();
    emit PriceRegistry.PremiumMultiplierWeiPerEthUpdated(
      vm.addr(1), premiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth
    );

    s_priceRegistry.applyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs);

    assertEq(
      s_priceRegistryPremiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth,
      s_priceRegistry.getPremiumMultiplierWeiPerEth(vm.addr(1))
    );
  }

  function test_applyPremiumMultiplierWeiPerEthUpdatesMultipleTokens_Success() public {
    PriceRegistry.PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs =
      new PriceRegistry.PremiumMultiplierWeiPerEthArgs[](2);
    premiumMultiplierWeiPerEthArgs[0] = s_priceRegistryPremiumMultiplierWeiPerEthArgs[0];
    premiumMultiplierWeiPerEthArgs[0].token = vm.addr(1);
    premiumMultiplierWeiPerEthArgs[1].token = vm.addr(2);

    vm.expectEmit();
    emit PriceRegistry.PremiumMultiplierWeiPerEthUpdated(
      vm.addr(1), premiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth
    );
    vm.expectEmit();
    emit PriceRegistry.PremiumMultiplierWeiPerEthUpdated(
      vm.addr(2), premiumMultiplierWeiPerEthArgs[1].premiumMultiplierWeiPerEth
    );

    s_priceRegistry.applyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs);

    assertEq(
      premiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth,
      s_priceRegistry.getPremiumMultiplierWeiPerEth(vm.addr(1))
    );
    assertEq(
      premiumMultiplierWeiPerEthArgs[1].premiumMultiplierWeiPerEth,
      s_priceRegistry.getPremiumMultiplierWeiPerEth(vm.addr(2))
    );
  }

  function test_applyPremiumMultiplierWeiPerEthUpdatesZeroInput() public {
    vm.recordLogs();
    s_priceRegistry.applyPremiumMultiplierWeiPerEthUpdates(new PriceRegistry.PremiumMultiplierWeiPerEthArgs[](0));

    assertEq(vm.getRecordedLogs().length, 0);
  }

  // Reverts

  function test_OnlyCallableByOwnerOrAdmin_Revert() public {
    PriceRegistry.PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs;
    vm.startPrank(STRANGER);

    vm.expectRevert("Only callable by owner");

    s_priceRegistry.applyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs);
  }
}

contract PriceRegistry_applyTokenTransferFeeConfigUpdates is PriceRegistrySetup {
  function test_Fuzz_ApplyTokenTransferFeeConfig_Success(
    PriceRegistry.TokenTransferFeeConfig[2] memory tokenTransferFeeConfigs
  ) public {
    PriceRegistry.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
      _generateTokenTransferFeeConfigArgs(2, 2);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[1].destChainSelector = DEST_CHAIN_SELECTOR + 1;

    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      for (uint256 j = 0; j < tokenTransferFeeConfigs.length; ++j) {
        tokenTransferFeeConfigs[j].destBytesOverhead = uint32(
          bound(tokenTransferFeeConfigs[j].destBytesOverhead, Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES, type(uint32).max)
        );
        address feeToken = s_sourceTokens[j];
        tokenTransferFeeConfigArgs[i].tokenTransferFeeConfigs[j].token = feeToken;
        tokenTransferFeeConfigArgs[i].tokenTransferFeeConfigs[j].tokenTransferFeeConfig = tokenTransferFeeConfigs[j];

        vm.expectEmit();
        emit PriceRegistry.TokenTransferFeeConfigUpdated(
          tokenTransferFeeConfigArgs[i].destChainSelector, feeToken, tokenTransferFeeConfigs[j]
        );
      }
    }

    s_priceRegistry.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](0)
    );

    for (uint256 i = 0; i < tokenTransferFeeConfigs.length; ++i) {
      _assertTokenTransferFeeConfigEqual(
        tokenTransferFeeConfigs[i],
        s_priceRegistry.getTokenTransferFeeConfig(
          tokenTransferFeeConfigArgs[0].destChainSelector,
          tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[i].token
        )
      );
    }
  }

  function test_ApplyTokenTransferFeeConfig_Success() public {
    PriceRegistry.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
      _generateTokenTransferFeeConfigArgs(1, 2);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = address(5);
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = PriceRegistry
      .TokenTransferFeeConfig({
      minFeeUSDCents: 6,
      maxFeeUSDCents: 7,
      deciBps: 8,
      destGasOverhead: 9,
      destBytesOverhead: 312,
      isEnabled: true
    });
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token = address(11);
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig = PriceRegistry
      .TokenTransferFeeConfig({
      minFeeUSDCents: 12,
      maxFeeUSDCents: 13,
      deciBps: 14,
      destGasOverhead: 15,
      destBytesOverhead: 394,
      isEnabled: true
    });

    vm.expectEmit();
    emit PriceRegistry.TokenTransferFeeConfigUpdated(
      tokenTransferFeeConfigArgs[0].destChainSelector,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig
    );
    vm.expectEmit();
    emit PriceRegistry.TokenTransferFeeConfigUpdated(
      tokenTransferFeeConfigArgs[0].destChainSelector,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig
    );

    PriceRegistry.TokenTransferFeeConfigRemoveArgs[] memory tokensToRemove =
      new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](0);
    s_priceRegistry.applyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs, tokensToRemove);

    PriceRegistry.TokenTransferFeeConfig memory config0 = s_priceRegistry.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    );
    PriceRegistry.TokenTransferFeeConfig memory config1 = s_priceRegistry.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token
    );

    _assertTokenTransferFeeConfigEqual(
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig, config0
    );
    _assertTokenTransferFeeConfigEqual(
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig, config1
    );

    // Remove only the first token and validate only the first token is removed
    tokensToRemove = new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](1);
    tokensToRemove[0] = PriceRegistry.TokenTransferFeeConfigRemoveArgs({
      destChainSelector: tokenTransferFeeConfigArgs[0].destChainSelector,
      token: tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    });

    vm.expectEmit();
    emit PriceRegistry.TokenTransferFeeConfigDeleted(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    );

    s_priceRegistry.applyTokenTransferFeeConfigUpdates(
      new PriceRegistry.TokenTransferFeeConfigArgs[](0), tokensToRemove
    );

    config0 = s_priceRegistry.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    );
    config1 = s_priceRegistry.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token
    );

    PriceRegistry.TokenTransferFeeConfig memory emptyConfig;

    _assertTokenTransferFeeConfigEqual(emptyConfig, config0);
    _assertTokenTransferFeeConfigEqual(
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig, config1
    );
  }

  function test_ApplyTokenTransferFeeZeroInput() public {
    vm.recordLogs();
    s_priceRegistry.applyTokenTransferFeeConfigUpdates(
      new PriceRegistry.TokenTransferFeeConfigArgs[](0), new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](0)
    );

    assertEq(vm.getRecordedLogs().length, 0);
  }

  // Reverts

  function test_OnlyCallableByOwnerOrAdmin_Revert() public {
    vm.startPrank(STRANGER);
    PriceRegistry.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs;

    vm.expectRevert("Only callable by owner");

    s_priceRegistry.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](0)
    );
  }

  function test_InvalidDestBytesOverhead_Revert() public {
    PriceRegistry.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
      _generateTokenTransferFeeConfigArgs(1, 1);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = address(5);
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = PriceRegistry
      .TokenTransferFeeConfig({
      minFeeUSDCents: 6,
      maxFeeUSDCents: 7,
      deciBps: 8,
      destGasOverhead: 9,
      destBytesOverhead: uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES - 1),
      isEnabled: true
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        PriceRegistry.InvalidDestBytesOverhead.selector,
        tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token,
        tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig.destBytesOverhead
      )
    );

    s_priceRegistry.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](0)
    );
  }
}

contract PriceRegistry_getTokenTransferCost is PriceRegistryFeeSetup {
  using USDPriceWith18Decimals for uint224;

  function test_NoTokenTransferChargesZeroFee_Success() public view {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    assertEq(0, feeUSDWei);
    assertEq(0, destGasOverhead);
    assertEq(0, destBytesOverhead);
  }

  function test_getTokenTransferCost_selfServeUsesDefaults_Success() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_selfServeTokenDefaultPricing, 1000);

    // Get config to assert it isn't set
    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    assertFalse(transferFeeConfig.isEnabled);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    // Assert that the default values are used
    assertEq(uint256(DEFAULT_TOKEN_FEE_USD_CENTS) * 1e16, feeUSDWei);
    assertEq(DEFAULT_TOKEN_DEST_GAS_OVERHEAD, destGasOverhead);
    assertEq(DEFAULT_TOKEN_BYTES_OVERHEAD, destBytesOverhead);
  }

  function test_SmallTokenTransferChargesMinFeeAndGas_Success() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1000);
    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    assertEq(configUSDCentToWei(transferFeeConfig.minFeeUSDCents), feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_ZeroAmountTokenTransferChargesMinFeeAndGas_Success() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 0);
    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    assertEq(configUSDCentToWei(transferFeeConfig.minFeeUSDCents), feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_LargeTokenTransferChargesMaxFeeAndGas_Success() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1e36);
    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    assertEq(configUSDCentToWei(transferFeeConfig.maxFeeUSDCents), feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_FeeTokenBpsFee_Success() public view {
    uint256 tokenAmount = 10000e18;

    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, tokenAmount);
    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    uint256 usdWei = calcUSDValueFromTokenAmount(s_feeTokenPrice, tokenAmount);
    uint256 bpsUSDWei = applyBpsRatio(
      usdWei, s_priceRegistryTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig.deciBps
    );

    assertEq(bpsUSDWei, feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_WETHTokenBpsFee_Success() public view {
    uint256 tokenAmount = 100e18;

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceRouter.getWrappedNative(),
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceRouter.getWrappedNative(), amount: tokenAmount});

    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) = s_priceRegistry.getTokenTransferCost(
      DEST_CHAIN_SELECTOR, message.feeToken, s_wrappedTokenPrice, message.tokenAmounts
    );

    uint256 usdWei = calcUSDValueFromTokenAmount(s_wrappedTokenPrice, tokenAmount);
    uint256 bpsUSDWei = applyBpsRatio(
      usdWei, s_priceRegistryTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig.deciBps
    );

    assertEq(bpsUSDWei, feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_CustomTokenBpsFee_Success() public view {
    uint256 tokenAmount = 200000e18;

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: CUSTOM_TOKEN, amount: tokenAmount});

    PriceRegistry.TokenTransferFeeConfig memory transferFeeConfig =
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    uint256 usdWei = calcUSDValueFromTokenAmount(s_customTokenPrice, tokenAmount);
    uint256 bpsUSDWei = applyBpsRatio(
      usdWei, s_priceRegistryTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[2].tokenTransferFeeConfig.deciBps
    );

    assertEq(bpsUSDWei, feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_ZeroFeeConfigChargesMinFee_Success() public {
    PriceRegistry.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
      _generateTokenTransferFeeConfigArgs(1, 1);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = s_sourceFeeToken;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = PriceRegistry
      .TokenTransferFeeConfig({
      minFeeUSDCents: 1,
      maxFeeUSDCents: 0,
      deciBps: 0,
      destGasOverhead: 0,
      destBytesOverhead: uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES),
      isEnabled: true
    });
    s_priceRegistry.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new PriceRegistry.TokenTransferFeeConfigRemoveArgs[](0)
    );

    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1e36);
    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, s_feeTokenPrice, message.tokenAmounts);

    // if token charges 0 bps, it should cost minFee to transfer
    assertEq(
      configUSDCentToWei(tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig.minFeeUSDCents),
      feeUSDWei
    );
    assertEq(0, destGasOverhead);
    assertEq(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES, destBytesOverhead);
  }

  function test_Fuzz_TokenTransferFeeDuplicateTokens_Success(uint256 transfers, uint256 amount) public view {
    // It shouldn't be possible to pay materially lower fees by splitting up the transfers.
    // Note it is possible to pay higher fees since the minimum fees are added.
    PriceRegistry.DestChainDynamicConfig memory dynamicConfig =
      s_priceRegistry.getDestChainDynamicConfig(DEST_CHAIN_SELECTOR);
    transfers = bound(transfers, 1, dynamicConfig.maxNumberOfTokensPerMsg);
    // Cap amount to avoid overflow
    amount = bound(amount, 0, 1e36);
    Client.EVMTokenAmount[] memory multiple = new Client.EVMTokenAmount[](transfers);
    for (uint256 i = 0; i < transfers; ++i) {
      multiple[i] = Client.EVMTokenAmount({token: s_sourceTokens[0], amount: amount});
    }
    Client.EVMTokenAmount[] memory single = new Client.EVMTokenAmount[](1);
    single[0] = Client.EVMTokenAmount({token: s_sourceTokens[0], amount: amount * transfers});

    address feeToken = s_sourceRouter.getWrappedNative();

    (uint256 feeSingleUSDWei, uint32 gasOverheadSingle, uint32 bytesOverheadSingle) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, feeToken, s_wrappedTokenPrice, single);
    (uint256 feeMultipleUSDWei, uint32 gasOverheadMultiple, uint32 bytesOverheadMultiple) =
      s_priceRegistry.getTokenTransferCost(DEST_CHAIN_SELECTOR, feeToken, s_wrappedTokenPrice, multiple);

    // Note that there can be a rounding error once per split.
    assertGe(feeMultipleUSDWei, (feeSingleUSDWei - dynamicConfig.maxNumberOfTokensPerMsg));
    assertEq(gasOverheadMultiple, gasOverheadSingle * transfers);
    assertEq(bytesOverheadMultiple, bytesOverheadSingle * transfers);
  }

  function test_MixedTokenTransferFee_Success() public view {
    address[3] memory testTokens = [s_sourceFeeToken, s_sourceRouter.getWrappedNative(), CUSTOM_TOKEN];
    uint224[3] memory tokenPrices = [s_feeTokenPrice, s_wrappedTokenPrice, s_customTokenPrice];
    PriceRegistry.TokenTransferFeeConfig[3] memory tokenTransferFeeConfigs = [
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[0]),
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[1]),
      s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[2])
    ];

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](3),
      feeToken: s_sourceRouter.getWrappedNative(),
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
    uint256 expectedTotalGas = 0;
    uint256 expectedTotalBytes = 0;

    // Start with small token transfers, total bps fee is lower than min token transfer fee
    for (uint256 i = 0; i < testTokens.length; ++i) {
      message.tokenAmounts[i] = Client.EVMTokenAmount({token: testTokens[i], amount: 1e14});
      expectedTotalGas += s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[i]).destGasOverhead;
      expectedTotalBytes +=
        s_priceRegistry.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[i]).destBytesOverhead;
    }
    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) = s_priceRegistry.getTokenTransferCost(
      DEST_CHAIN_SELECTOR, message.feeToken, s_wrappedTokenPrice, message.tokenAmounts
    );

    uint256 expectedFeeUSDWei = 0;
    for (uint256 i = 0; i < testTokens.length; ++i) {
      expectedFeeUSDWei += configUSDCentToWei(tokenTransferFeeConfigs[i].minFeeUSDCents);
    }

    assertEq(expectedFeeUSDWei, feeUSDWei);
    assertEq(expectedTotalGas, destGasOverhead);
    assertEq(expectedTotalBytes, destBytesOverhead);

    // Set 1st token transfer to a meaningful amount so its bps fee is now between min and max fee
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: testTokens[0], amount: 10000e18});

    (feeUSDWei, destGasOverhead, destBytesOverhead) = s_priceRegistry.getTokenTransferCost(
      DEST_CHAIN_SELECTOR, message.feeToken, s_wrappedTokenPrice, message.tokenAmounts
    );
    expectedFeeUSDWei = applyBpsRatio(
      calcUSDValueFromTokenAmount(tokenPrices[0], message.tokenAmounts[0].amount), tokenTransferFeeConfigs[0].deciBps
    );
    expectedFeeUSDWei += configUSDCentToWei(tokenTransferFeeConfigs[1].minFeeUSDCents);
    expectedFeeUSDWei += configUSDCentToWei(tokenTransferFeeConfigs[2].minFeeUSDCents);

    assertEq(expectedFeeUSDWei, feeUSDWei);
    assertEq(expectedTotalGas, destGasOverhead);
    assertEq(expectedTotalBytes, destBytesOverhead);

    // Set 2nd token transfer to a large amount that is higher than maxFeeUSD
    message.tokenAmounts[1] = Client.EVMTokenAmount({token: testTokens[1], amount: 1e36});

    (feeUSDWei, destGasOverhead, destBytesOverhead) = s_priceRegistry.getTokenTransferCost(
      DEST_CHAIN_SELECTOR, message.feeToken, s_wrappedTokenPrice, message.tokenAmounts
    );
    expectedFeeUSDWei = applyBpsRatio(
      calcUSDValueFromTokenAmount(tokenPrices[0], message.tokenAmounts[0].amount), tokenTransferFeeConfigs[0].deciBps
    );
    expectedFeeUSDWei += configUSDCentToWei(tokenTransferFeeConfigs[1].maxFeeUSDCents);
    expectedFeeUSDWei += configUSDCentToWei(tokenTransferFeeConfigs[2].minFeeUSDCents);

    assertEq(expectedFeeUSDWei, feeUSDWei);
    assertEq(expectedTotalGas, destGasOverhead);
    assertEq(expectedTotalBytes, destBytesOverhead);
  }
}

contract PriceRegistry_validateDestFamilyAddress is PriceRegistrySetup {
  function test_ValidEVMAddress_Success() public view {
    bytes memory encodedAddress = abi.encode(address(10000));
    s_priceRegistry.validateDestFamilyAddress(Internal.CHAIN_FAMILY_SELECTOR_EVM, encodedAddress);
  }

  function test_ValidNonEVMAddress_Success() public view {
    s_priceRegistry.validateDestFamilyAddress(bytes4(uint32(1)), abi.encode(type(uint208).max));
  }

  // Reverts

  function test_InvalidEVMAddress_Revert() public {
    bytes memory invalidAddress = abi.encode(type(uint208).max);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    s_priceRegistry.validateDestFamilyAddress(Internal.CHAIN_FAMILY_SELECTOR_EVM, invalidAddress);
  }
}

contract PriceRegistry_parseEVMExtraArgsFromBytes is PriceRegistrySetup {
  PriceRegistry.DestChainDynamicConfig private s_destChainDynamicConfig;

  function setUp() public virtual override {
    super.setUp();
    s_destChainDynamicConfig = _generatePriceRegistryDestChainDynamicConfigArgs()[0].dynamicConfig;
  }

  function test_EVMExtraArgsV1_Success() public view {
    Client.EVMExtraArgsV1 memory inputArgs = Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT});
    bytes memory inputExtraArgs = Client._argsToBytes(inputArgs);
    Client.EVMExtraArgsV2 memory expectedOutputArgs =
      Client.EVMExtraArgsV2({gasLimit: GAS_LIMIT, allowOutOfOrderExecution: false});

    vm.assertEq(
      abi.encode(s_priceRegistry.parseEVMExtraArgsFromBytes(inputExtraArgs, s_destChainDynamicConfig)),
      abi.encode(expectedOutputArgs)
    );
  }

  function test_EVMExtraArgsV2_Success() public view {
    Client.EVMExtraArgsV2 memory inputArgs =
      Client.EVMExtraArgsV2({gasLimit: GAS_LIMIT, allowOutOfOrderExecution: true});
    bytes memory inputExtraArgs = Client._argsToBytes(inputArgs);

    vm.assertEq(
      abi.encode(s_priceRegistry.parseEVMExtraArgsFromBytes(inputExtraArgs, s_destChainDynamicConfig)),
      abi.encode(inputArgs)
    );
  }

  function test_EVMExtraArgsDefault_Success() public view {
    Client.EVMExtraArgsV2 memory expectedOutputArgs =
      Client.EVMExtraArgsV2({gasLimit: s_destChainDynamicConfig.defaultTxGasLimit, allowOutOfOrderExecution: false});

    vm.assertEq(
      abi.encode(s_priceRegistry.parseEVMExtraArgsFromBytes("", s_destChainDynamicConfig)),
      abi.encode(expectedOutputArgs)
    );
  }

  // Reverts

  function test_EVMExtraArgsInvalidExtraArgsTag_Revert() public {
    Client.EVMExtraArgsV2 memory inputArgs =
      Client.EVMExtraArgsV2({gasLimit: GAS_LIMIT, allowOutOfOrderExecution: true});
    bytes memory inputExtraArgs = Client._argsToBytes(inputArgs);
    // Invalidate selector
    inputExtraArgs[0] = bytes1(uint8(0));

    vm.expectRevert(PriceRegistry.InvalidExtraArgsTag.selector);
    s_priceRegistry.parseEVMExtraArgsFromBytes(inputExtraArgs, s_destChainDynamicConfig);
  }

  function test_EVMExtraArgsEnforceOutOfOrder_Revert() public {
    Client.EVMExtraArgsV2 memory inputArgs =
      Client.EVMExtraArgsV2({gasLimit: GAS_LIMIT, allowOutOfOrderExecution: false});
    bytes memory inputExtraArgs = Client._argsToBytes(inputArgs);
    s_destChainDynamicConfig.enforceOutOfOrder = true;

    vm.expectRevert(PriceRegistry.ExtraArgOutOfOrderExecutionMustBeTrue.selector);
    s_priceRegistry.parseEVMExtraArgsFromBytes(inputExtraArgs, s_destChainDynamicConfig);
  }

  function test_EVMExtraArgsGasLimitTooHigh_Revert() public {
    Client.EVMExtraArgsV2 memory inputArgs =
      Client.EVMExtraArgsV2({gasLimit: s_destChainDynamicConfig.maxPerMsgGasLimit + 1, allowOutOfOrderExecution: true});
    bytes memory inputExtraArgs = Client._argsToBytes(inputArgs);

    vm.expectRevert(PriceRegistry.MessageGasLimitTooHigh.selector);
    s_priceRegistry.parseEVMExtraArgsFromBytes(inputExtraArgs, s_destChainDynamicConfig);
  }
}

// TODO: add validation tests (revisit placing them back to MultiOnRamp tests)
// {
// function test_Fuzz_EnforceOutOfOrder(bool enforce, bool allowOutOfOrderExecution) public {
//   // Update dynamic config to enforce allowOutOfOrderExecution = defaultVal.
//   vm.stopPrank();
//   vm.startPrank(OWNER);

//   EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//   destChainConfigArgs[0].dynamicConfig.enforceOutOfOrder = enforce;
//   s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

//   vm.stopPrank();

//   vm.startPrank(address(s_sourceRouter));
//   Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//   message.extraArgs = abi.encodeWithSelector(
//     Client.EVM_EXTRA_ARGS_V2_TAG,
//     Client.EVMExtraArgsV2({gasLimit: GAS_LIMIT * 2, allowOutOfOrderExecution: allowOutOfOrderExecution})
//   );
//   uint256 feeAmount = 1234567890;
//   IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

//   if (enforce) {
//     // If enforcement is on, only true should be allowed.
//     if (allowOutOfOrderExecution) {
//       vm.expectEmit();
//       emit EVM2EVMMultiOnRamp.CCIPSendRequested(DEST_CHAIN_SELECTOR, _messageToEvent(message, 1, 1, feeAmount, OWNER));
//       s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
//     } else {
//       vm.expectRevert(EVM2EVMMultiOnRamp.ExtraArgOutOfOrderExecutionMustBeTrue.selector);
//       s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
//     }
//   } else {
//     // no enforcement should allow any value.
//     vm.expectEmit();
//     emit EVM2EVMMultiOnRamp.CCIPSendRequested(DEST_CHAIN_SELECTOR, _messageToEvent(message, 1, 1, feeAmount, OWNER));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
//   }
// }

// function test_MessageTooLarge_Revert() public {
//   Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//   message.data = new bytes(MAX_DATA_SIZE + 1);
//   vm.expectRevert(
//     abi.encodeWithSelector(EVM2EVMMultiOnRamp.MessageTooLarge.selector, MAX_DATA_SIZE, message.data.length)
//   );

//   s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
// }

// function test_TooManyTokens_Revert() public {
//   Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//   uint256 tooMany = MAX_TOKENS_LENGTH + 1;
//   message.tokenAmounts = new Client.EVMTokenAmount[](tooMany);
//   vm.expectRevert(EVM2EVMMultiOnRamp.UnsupportedNumberOfTokens.selector);
//   s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
// }

// Asserts gasLimit must be <=maxGasLimit
//   function test_MessageGasLimitTooHigh_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     message.extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: MAX_GAS_LIMIT + 1}));
//     vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOnRamp.MessageGasLimitTooHigh.selector));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);
//   }

//   function test_InvalidAddressEncodePacked_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     message.receiver = abi.encodePacked(address(234));

//     vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, message.receiver));

//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1, OWNER);
//   }

//   function test_InvalidEVMAddress_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     message.receiver = abi.encode(type(uint208).max);

//     vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, message.receiver));

//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1, OWNER);
//   }

//   // We disallow sending to addresses 0-9.
//   function test_ZeroAddressReceiver_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();

//     for (uint160 i = 0; i < 10; ++i) {
//       message.receiver = abi.encode(address(i));

//       vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, message.receiver));

//       s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1, OWNER);
//     }
//   }

//   function test_MesssageFeeTooHigh_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();

//     vm.expectRevert(
//       abi.encodeWithSelector(EVM2EVMMultiOnRamp.MessageFeeTooHigh.selector, MAX_MSG_FEES_JUELS + 1, MAX_MSG_FEES_JUELS)
//     );

//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, MAX_MSG_FEES_JUELS + 1, OWNER);
//   }

//   function test_SourceTokenDataTooLarge_Revert() public {
//     address sourceETH = s_sourceTokens[1];
//     vm.stopPrank();
//     vm.startPrank(OWNER);

//     MaybeRevertingBurnMintTokenPool newPool = new MaybeRevertingBurnMintTokenPool(
//       BurnMintERC677(sourceETH), new address[](0), address(s_mockRMN), address(s_sourceRouter)
//     );
//     BurnMintERC677(sourceETH).grantMintAndBurnRoles(address(newPool));
//     deal(address(sourceETH), address(newPool), type(uint256).max);

//     // Add TokenPool to OnRamp
//     s_tokenAdminRegistry.setPool(sourceETH, address(newPool));

//     // Allow chain in TokenPool
//     TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
//     chainUpdates[0] = TokenPool.ChainUpdate({
//       remoteChainSelector: DEST_CHAIN_SELECTOR,
//       remotePoolAddress: abi.encode(s_destTokenPool),
//       remoteTokenAddress: abi.encode(s_destToken),
//       allowed: true,
//       outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
//       inboundRateLimiterConfig: getInboundRateLimiterConfig()
//     });
//     newPool.applyChainUpdates(chainUpdates);

//     Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(address(sourceETH), 1000);

//     // No data set, should succeed
//     vm.startPrank(address(s_sourceRouter));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);

//     // Set max data length, should succeed
//     vm.startPrank(OWNER);
//     newPool.setSourceTokenData(new bytes(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES));

//     vm.startPrank(address(s_sourceRouter));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);

//     // Set data to max length +1, should revert
//     vm.startPrank(OWNER);
//     newPool.setSourceTokenData(new bytes(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES + 1));

//     vm.startPrank(address(s_sourceRouter));
//     vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOnRamp.SourceTokenDataTooLarge.selector, sourceETH));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);

//     // Set token config to allow larger data
//     vm.startPrank(OWNER);
//     EVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
//       _generateTokenTransferFeeConfigArgs(1, 1);
//     tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
//     tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = sourceETH;
//     tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = EVM2EVMMultiOnRamp
//       .TokenTransferFeeConfig({
//       minFeeUSDCents: 1,
//       maxFeeUSDCents: 0,
//       deciBps: 0,
//       destGasOverhead: 0,
//       destBytesOverhead: uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) + 32,
//       isEnabled: true
//     });
//     s_onRamp.applyTokenTransferFeeConfigUpdates(
//       tokenTransferFeeConfigArgs, new EVM2EVMMultiOnRamp.TokenTransferFeeConfigRemoveArgs[](0)
//     );

//     vm.startPrank(address(s_sourceRouter));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);

//     // Set the token data larger than the configured token data, should revert
//     vm.startPrank(OWNER);
//     newPool.setSourceTokenData(new bytes(uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) + 32 + 1));

//     vm.startPrank(address(s_sourceRouter));
//     vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOnRamp.SourceTokenDataTooLarge.selector, sourceETH));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);
//   }

//   function test_InvalidEVMAddressDestToken_Revert() public {
//     address sourceETH = s_sourceTokens[1];
//     vm.stopPrank();
//     vm.startPrank(OWNER);

//     MaybeRevertingBurnMintTokenPool newPool = new MaybeRevertingBurnMintTokenPool(
//       BurnMintERC677(sourceETH), new address[](0), address(s_mockRMN), address(s_sourceRouter)
//     );
//     BurnMintERC677(sourceETH).grantMintAndBurnRoles(address(newPool));
//     deal(address(sourceETH), address(newPool), type(uint256).max);

//     // Add TokenPool to OnRamp
//     s_tokenAdminRegistry.setPool(sourceETH, address(newPool));

//     bytes memory nonEvmAddress = abi.encode(type(uint208).max);

//     // Allow chain in TokenPool
//     TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
//     chainUpdates[0] = TokenPool.ChainUpdate({
//       remoteChainSelector: DEST_CHAIN_SELECTOR,
//       remotePoolAddress: abi.encode(s_destTokenPool),
//       remoteTokenAddress: nonEvmAddress,
//       allowed: true,
//       outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
//       inboundRateLimiterConfig: getInboundRateLimiterConfig()
//     });
//     newPool.applyChainUpdates(chainUpdates);

//     Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(address(sourceETH), 1000);

//     vm.startPrank(address(s_sourceRouter));
//     vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, nonEvmAddress));
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, OWNER);
//   }

//   function test_EnforceOutOfOrder_Revert() public {
//     // Update dynamic config to enforce allowOutOfOrderExecution = true.
//     vm.stopPrank();
//     vm.startPrank(OWNER);

//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs = _generateDestChainConfigArgs();
//     destChainConfigArgs[0].dynamicConfig.enforceOutOfOrder = true;
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
//     vm.stopPrank();

//     vm.startPrank(address(s_sourceRouter));
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     // Empty extraArgs to should revert since it enforceOutOfOrder is true.
//     message.extraArgs = "";
//     uint256 feeAmount = 1234567890;
//     IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

//     vm.expectRevert(EVM2EVMMultiOnRamp.ExtraArgOutOfOrderExecutionMustBeTrue.selector);
//     s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
//   }
// }

// TODO: getValidatedFee
// {
// function test_EmptyMessage_Success() public view {
//   address[2] memory testTokens = [s_sourceFeeToken, s_sourceRouter.getWrappedNative()];
//   uint224[2] memory feeTokenPrices = [s_feeTokenPrice, s_wrappedTokenPrice];

//   for (uint256 i = 0; i < feeTokenPrices.length; ++i) {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     message.feeToken = testTokens[i];
//     uint64 premiumMultiplierWeiPerEth = s_priceRegistry.getPremiumMultiplierWeiPerEth(message.feeToken);
//     PriceRegistry.DestChainDynamicConfig memory destChainDynamicConfig =
//       s_priceRegistry.getDestChainConfig(DEST_CHAIN_SELECTOR);

//     uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

//     uint256 gasUsed = GAS_LIMIT + DEST_GAS_OVERHEAD;
//     uint256 gasFeeUSD = (gasUsed * destChainDynamicConfig.gasMultiplierWeiPerEth * USD_PER_GAS);
//     uint256 messageFeeUSD =
//       (configUSDCentToWei(destChainDynamicConfig.networkFeeUSDCents) * premiumMultiplierWeiPerEth);
//     uint256 dataAvailabilityFeeUSD = s_priceRegistry.getDataAvailabilityCost(
//       DEST_CHAIN_SELECTOR, USD_PER_DATA_AVAILABILITY_GAS, message.data.length, message.tokenAmounts.length, 0
//     );

//     uint256 totalPriceInFeeToken = (gasFeeUSD + messageFeeUSD + dataAvailabilityFeeUSD) / feeTokenPrices[i];
//     assertEq(totalPriceInFeeToken, feeAmount);
//   }
// }

//   function test_ZeroDataAvailabilityMultiplier_Success() public {
//     EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigArgs =
//       new EVM2EVMMultiOnRamp.DestChainConfigArgs[](1);
//     EVM2EVMMultiOnRamp.DestChainConfig memory destChainConfig = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
//     destChainConfigArgs[0] = EVM2EVMMultiOnRamp.DestChainConfigArgs({
//       destChainSelector: DEST_CHAIN_SELECTOR,
//       dynamicConfig: destChainConfig.dynamicConfig,
//       prevOnRamp: destChainConfig.prevOnRamp
//     });
//     destChainConfigArgs[0].dynamicConfig.destDataAvailabilityMultiplierBps = 0;
//     s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     uint64 premiumMultiplierWeiPerEth = s_onRamp.getPremiumMultiplierWeiPerEth(message.feeToken);
//     EVM2EVMMultiOnRamp.DestChainDynamicConfig memory destChainDynamicConfig =
//       s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).dynamicConfig;

//     uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

//     uint256 gasUsed = GAS_LIMIT + DEST_GAS_OVERHEAD;
//     uint256 gasFeeUSD = (gasUsed * destChainDynamicConfig.gasMultiplierWeiPerEth * USD_PER_GAS);
//     uint256 messageFeeUSD = (configUSDCentToWei(destChainDynamicConfig.networkFeeUSDCents) * premiumMultiplierWeiPerEth);

//     uint256 totalPriceInFeeToken = (gasFeeUSD + messageFeeUSD) / s_feeTokenPrice;
//     assertEq(totalPriceInFeeToken, feeAmount);
//   }

//   function test_HighGasMessage_Success() public view {
//     address[2] memory testTokens = [s_sourceFeeToken, s_sourceRouter.getWrappedNative()];
//     uint224[2] memory feeTokenPrices = [s_feeTokenPrice, s_wrappedTokenPrice];

//     uint256 customGasLimit = MAX_GAS_LIMIT;
//     uint256 customDataSize = MAX_DATA_SIZE;
//     for (uint256 i = 0; i < feeTokenPrices.length; ++i) {
//       Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
//         receiver: abi.encode(OWNER),
//         data: new bytes(customDataSize),
//         tokenAmounts: new Client.EVMTokenAmount[](0),
//         feeToken: testTokens[i],
//         extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: customGasLimit}))
//       });

//       uint64 premiumMultiplierWeiPerEth = s_onRamp.getPremiumMultiplierWeiPerEth(message.feeToken);
//       EVM2EVMMultiOnRamp.DestChainDynamicConfig memory destChainDynamicConfig =
//         s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).dynamicConfig;

//       uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
//       uint256 gasUsed = customGasLimit + DEST_GAS_OVERHEAD + customDataSize * DEST_GAS_PER_PAYLOAD_BYTE;
//       uint256 gasFeeUSD = (gasUsed * destChainDynamicConfig.gasMultiplierWeiPerEth * USD_PER_GAS);
//       uint256 messageFeeUSD =
//         (configUSDCentToWei(destChainDynamicConfig.networkFeeUSDCents) * premiumMultiplierWeiPerEth);
//       uint256 dataAvailabilityFeeUSD = s_onRamp.getDataAvailabilityCost(
//         DEST_CHAIN_SELECTOR, USD_PER_DATA_AVAILABILITY_GAS, message.data.length, message.tokenAmounts.length, 0
//       );

//       uint256 totalPriceInFeeToken = (gasFeeUSD + messageFeeUSD + dataAvailabilityFeeUSD) / feeTokenPrices[i];
//       assertEq(totalPriceInFeeToken, feeAmount);
//     }
//   }

//   function test_SingleTokenMessage_Success() public view {
//     address[2] memory testTokens = [s_sourceFeeToken, s_sourceRouter.getWrappedNative()];
//     uint224[2] memory feeTokenPrices = [s_feeTokenPrice, s_wrappedTokenPrice];

//     uint256 tokenAmount = 10000e18;
//     for (uint256 i = 0; i < feeTokenPrices.length; ++i) {
//       Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, tokenAmount);
//       message.feeToken = testTokens[i];
//       EVM2EVMMultiOnRamp.DestChainDynamicConfig memory destChainDynamicConfig =
//         s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).dynamicConfig;
//       uint32 destBytesOverhead =
//         s_onRamp.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token).destBytesOverhead;
//       uint32 tokenBytesOverhead =
//         destBytesOverhead == 0 ? uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) : destBytesOverhead;

//       uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

//       uint256 gasUsed = GAS_LIMIT + DEST_GAS_OVERHEAD
//         + s_onRamp.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token).destGasOverhead;
//       uint256 gasFeeUSD = (gasUsed * destChainDynamicConfig.gasMultiplierWeiPerEth * USD_PER_GAS);
//       (uint256 transferFeeUSD,,) =
//         s_onRamp.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, feeTokenPrices[i], message.tokenAmounts);
//       uint256 messageFeeUSD = (transferFeeUSD * s_onRamp.getPremiumMultiplierWeiPerEth(message.feeToken));
//       uint256 dataAvailabilityFeeUSD = s_onRamp.getDataAvailabilityCost(
//         DEST_CHAIN_SELECTOR,
//         USD_PER_DATA_AVAILABILITY_GAS,
//         message.data.length,
//         message.tokenAmounts.length,
//         tokenBytesOverhead
//       );

//       uint256 totalPriceInFeeToken = (gasFeeUSD + messageFeeUSD + dataAvailabilityFeeUSD) / feeTokenPrices[i];
//       assertEq(totalPriceInFeeToken, feeAmount);
//     }
//   }

//   function test_MessageWithDataAndTokenTransfer_Success() public view {
//     address[2] memory testTokens = [s_sourceFeeToken, s_sourceRouter.getWrappedNative()];
//     uint224[2] memory feeTokenPrices = [s_feeTokenPrice, s_wrappedTokenPrice];

//     uint256 customGasLimit = 1_000_000;
//     for (uint256 i = 0; i < feeTokenPrices.length; ++i) {
//       Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
//         receiver: abi.encode(OWNER),
//         data: "",
//         tokenAmounts: new Client.EVMTokenAmount[](2),
//         feeToken: testTokens[i],
//         extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: customGasLimit}))
//       });
//       uint64 premiumMultiplierWeiPerEth = s_onRamp.getPremiumMultiplierWeiPerEth(message.feeToken);
//       EVM2EVMMultiOnRamp.DestChainDynamicConfig memory destChainDynamicConfig =
//         s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).dynamicConfig;

//       message.tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceFeeToken, amount: 10000e18}); // feeTokenAmount
//       message.tokenAmounts[1] = Client.EVMTokenAmount({token: CUSTOM_TOKEN, amount: 200000e18}); // customTokenAmount
//       message.data = "random bits and bytes that should be factored into the cost of the message";

//       uint32 tokenGasOverhead = 0;
//       uint32 tokenBytesOverhead = 0;
//       for (uint256 j = 0; j < message.tokenAmounts.length; ++j) {
//         tokenGasOverhead +=
//           s_onRamp.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[j].token).destGasOverhead;
//         uint32 destBytesOverhead =
//           s_onRamp.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[j].token).destBytesOverhead;
//         tokenBytesOverhead += destBytesOverhead == 0 ? uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) : destBytesOverhead;
//       }

//       uint256 gasUsed =
//         customGasLimit + DEST_GAS_OVERHEAD + message.data.length * DEST_GAS_PER_PAYLOAD_BYTE + tokenGasOverhead;
//       uint256 gasFeeUSD = (gasUsed * destChainDynamicConfig.gasMultiplierWeiPerEth * USD_PER_GAS);
//       (uint256 transferFeeUSD,,) =
//         s_onRamp.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.feeToken, feeTokenPrices[i], message.tokenAmounts);
//       uint256 messageFeeUSD = (transferFeeUSD * premiumMultiplierWeiPerEth);
//       uint256 dataAvailabilityFeeUSD = s_onRamp.getDataAvailabilityCost(
//         DEST_CHAIN_SELECTOR,
//         USD_PER_DATA_AVAILABILITY_GAS,
//         message.data.length,
//         message.tokenAmounts.length,
//         tokenBytesOverhead
//       );

//       uint256 totalPriceInFeeToken = (gasFeeUSD + messageFeeUSD + dataAvailabilityFeeUSD) / feeTokenPrices[i];
//       assertEq(totalPriceInFeeToken, s_onRamp.getFee(DEST_CHAIN_SELECTOR, message));
//     }
//   }

//   function test_MessageTooLarge_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     message.data = new bytes(MAX_DATA_SIZE + 1);
//     vm.expectRevert(
//       abi.encodeWithSelector(EVM2EVMMultiOnRamp.MessageTooLarge.selector, MAX_DATA_SIZE, message.data.length)
//     );

//     s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
//   }

//   function test_TooManyTokens_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     uint256 tooMany = MAX_TOKENS_LENGTH + 1;
//     message.tokenAmounts = new Client.EVMTokenAmount[](tooMany);
//     vm.expectRevert(EVM2EVMMultiOnRamp.UnsupportedNumberOfTokens.selector);
//     s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
//   }

//   // Asserts gasLimit must be <=maxGasLimit
//   function test_MessageGasLimitTooHigh_Revert() public {
//     Client.EVM2AnyMessage memory message = _generateEmptyMessage();
//     message.extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: MAX_GAS_LIMIT + 1}));
//     vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOnRamp.MessageGasLimitTooHigh.selector));
//     s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
//   }
//
//   function test_NotAFeeToken_Revert() public {
//     address notAFeeToken = address(0x111111);
//     Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(notAFeeToken, 1);
//     message.feeToken = notAFeeToken;

//     vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOnRamp.NotAFeeToken.selector, notAFeeToken));

//     s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
//   }
// }
