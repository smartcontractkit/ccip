// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {MockV3Aggregator} from "../../../tests/MockV3Aggregator.sol";
import {PriceRegistry} from "../../PriceRegistry.sol";
import {IPriceRegistry} from "../../interfaces/IPriceRegistry.sol";
import {Internal} from "../../libraries/Internal.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {Vm} from "forge-std/Vm.sol";

contract PriceRegistrySetup is TokenSetup {
  uint112 internal constant USD_PER_GAS = 1e6; // 0.001 gwei
  uint112 internal constant USD_PER_DATA_AVAILABILITY_GAS = 1e9; // 1 gwei

  // Encode L1 gas price and L2 gas price into a packed price.
  // L1 gas price is left-shifted to the higher-order bits.
  uint224 internal constant PACKED_USD_PER_GAS =
    (uint224(USD_PER_DATA_AVAILABILITY_GAS) << Internal.GAS_PRICE_BITS) + USD_PER_GAS;

  PriceRegistry internal s_priceRegistry;
  // Cheat to store the price updates in storage since struct arrays aren't supported.
  bytes internal s_encodedInitialPriceUpdates;
  address internal s_weth;

  address[] internal s_sourceFeeTokens;
  uint224[] internal s_sourceTokenPrices;
  address[] internal s_destFeeTokens;
  uint224[] internal s_destTokenPrices;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_weth = s_sourceRouter.getWrappedNative();

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
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](0);

    s_priceRegistry = new PriceRegistry(priceUpdaters, feeTokens, uint32(TWELVE_HOURS), tokenPriceFeedUpdates);
    s_priceRegistry.updatePrices(priceUpdates);
  }

  function _initialiseSingleTokenPriceFeed() internal returns (address) {
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]]);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    return s_sourceTokens[0];
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
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](2);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]]);
    tokenPriceFeedUpdates[1] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[1], s_dataFeedByToken[s_sourceTokens[1]]);

    s_priceRegistry = new PriceRegistry(priceUpdaters, feeTokens, uint32(TWELVE_HOURS), tokenPriceFeedUpdates);

    assertEq(feeTokens, s_priceRegistry.getFeeTokens());
    assertEq(uint32(TWELVE_HOURS), s_priceRegistry.getStalenessThreshold());
    assertEq(priceUpdaters, s_priceRegistry.getPriceUpdaters());
    assertEq(s_priceRegistry.typeAndVersion(), "PriceRegistry 1.6.0-dev");
    assertEq(tokenPriceFeedUpdates[0].dataFeedAddress, s_priceRegistry.getTokenPriceFeed(s_sourceTokens[0]));
    assertEq(tokenPriceFeedUpdates[1].dataFeedAddress, s_priceRegistry.getTokenPriceFeed(s_sourceTokens[1]));
  }

  function test_InvalidStalenessThreshold_Revert() public {
    vm.expectRevert(PriceRegistry.InvalidStalenessThreshold.selector);
    s_priceRegistry = new PriceRegistry(new address[](0), new address[](0), 0, new IPriceRegistry.TokenPriceFeedUpdate[](0));
  }
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
    assertEq(tokenPriceAnswer.timestamp, uint32(originalTimestampValue));
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

  function test_GetValidatedTokenPriceFromFeedMaxInt224Value_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, int256(uint256(type(uint224).max)));

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is: uint224.MAX_VALUE * (10 ** (36 - 18 - 18))
    assertEq(tokenPriceAnswer, uint224(type(uint224).max));
  }

  function test_GetValidatedTokenPriceFromFeedErc20Below18Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 6);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 8, 1e8);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e6) -> expected 1e30
    assertEq(tokenPriceAnswer, uint224(1e30));
  }

  function test_GetValidatedTokenPriceFromFeedErc20Above18Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 24);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 8, 1e8);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e24) -> expected 1e12
    assertEq(tokenPriceAnswer, uint224(1e12));
  }

  function test_GetValidatedTokenPriceFromFeedFeedAt18Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, 1e18);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e18) -> expected 1e18
    assertEq(tokenPriceAnswer, uint224(1e18));
  }

  function test_GetValidatedTokenPriceFromFeedFeedAt0Decimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 0);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 0, 1e31);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e31 (0 decimal token) - unit is (1e18 * 1e18 / 1e0) -> expected 1e36
    assertEq(tokenPriceAnswer, uint224(1e67));
  }

  function test_GetValidatedTokenPriceFromFeedFlippedDecimals_Success() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 20);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 20, 1e18);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    uint224 tokenPriceAnswer = s_priceRegistry.getValidatedTokenPrice(tokenAddress);

    // Price answer is 1e8 (6 decimal token) - unit is (1e18 * 1e18 / 1e20) -> expected 1e14
    assertEq(tokenPriceAnswer, uint224(1e14));
  }

  function test_OverflowFeedPrice_Revert() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, int256(uint256(type(uint224).max) + 1));

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    vm.expectRevert("Rebased data feed value does not fit in 224 bits");
    s_priceRegistry.getValidatedTokenPrice(tokenAddress);
  }

  function test_UnderflowFeedPrice_Revert() public {
    address tokenAddress = _deploySourceToken("testToken", 0, 18);
    address feedAddress = _deployTokenPriceDataFeed(tokenAddress, 18, -1);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, feedAddress);
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    vm.expectRevert("Data feed value must be positive");
    s_priceRegistry.getValidatedTokenPrice(tokenAddress);
  }

  function test_StaleFeedPrice_Revert() public {
    // Right above staleness threshold
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    address sourceToken = _initialiseSingleTokenPriceFeed();

    vm.expectRevert(
      abi.encodeWithSelector(PriceRegistry.StaleTokenPrice.selector, sourceToken, TWELVE_HOURS, TWELVE_HOURS + 1)
    );
    s_priceRegistry.getValidatedTokenPrice(sourceToken);
  }

  function test_StaleFeeToken_Revert() public {
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    vm.expectRevert(
      abi.encodeWithSelector(PriceRegistry.StaleTokenPrice.selector, s_sourceTokens[0], TWELVE_HOURS, TWELVE_HOURS + 1)
    );
    s_priceRegistry.getValidatedTokenPrice(s_sourceTokens[0]);
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

contract PriceRegistry_applyPriceUpdatersUpdates is PriceRegistrySetup {
  event PriceUpdaterSet(address indexed priceUpdater);
  event PriceUpdaterRemoved(address indexed priceUpdater);

  function test_ApplyPriceUpdaterUpdates_Success() public {
    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;

    vm.expectEmit();
    emit PriceUpdaterSet(STRANGER);

    s_priceRegistry.applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
    assertEq(s_priceRegistry.getPriceUpdaters().length, 1);
    assertEq(s_priceRegistry.getPriceUpdaters()[0], STRANGER);

    // add same priceUpdater is no-op
    s_priceRegistry.applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
    assertEq(s_priceRegistry.getPriceUpdaters().length, 1);
    assertEq(s_priceRegistry.getPriceUpdaters()[0], STRANGER);

    vm.expectEmit();
    emit PriceUpdaterRemoved(STRANGER);

    s_priceRegistry.applyPriceUpdatersUpdates(new address[](0), priceUpdaters);
    assertEq(s_priceRegistry.getPriceUpdaters().length, 0);

    // removing already removed priceUpdater is no-op
    s_priceRegistry.applyPriceUpdatersUpdates(new address[](0), priceUpdaters);
    assertEq(s_priceRegistry.getPriceUpdaters().length, 0);
  }

  function test_OnlyCallableByOwner_Revert() public {
    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;
    vm.startPrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_priceRegistry.applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
  }
}

contract PriceRegistry_applyFeeTokensUpdates is PriceRegistrySetup {
  event FeeTokenAdded(address indexed feeToken);
  event FeeTokenRemoved(address indexed feeToken);

  function test_ApplyFeeTokensUpdates_Success() public {
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceTokens[1];

    vm.expectEmit();
    emit FeeTokenAdded(feeTokens[0]);

    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));
    assertEq(s_priceRegistry.getFeeTokens().length, 3);
    assertEq(s_priceRegistry.getFeeTokens()[2], feeTokens[0]);

    // add same feeToken is no-op
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));
    assertEq(s_priceRegistry.getFeeTokens().length, 3);
    assertEq(s_priceRegistry.getFeeTokens()[2], feeTokens[0]);

    vm.expectEmit();
    emit FeeTokenRemoved(feeTokens[0]);

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
  event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp);
  event UsdPerUnitGasUpdated(uint64 indexed destChain, uint256 value, uint256 timestamp);

  function test_OnlyTokenPrice_Success() public {
    Internal.PriceUpdates memory update = Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](1),
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    });
    update.tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceTokens[0], usdPerToken: 4e18});

    vm.expectEmit();
    emit UsdPerTokenUpdated(
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
    emit UsdPerUnitGasUpdated(
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
      emit UsdPerTokenUpdated(
        update.tokenPriceUpdates[i].sourceToken, update.tokenPriceUpdates[i].usdPerToken, block.timestamp
      );
    }
    for (uint256 i = 0; i < gasPriceUpdates.length; ++i) {
      vm.expectEmit();
      emit UsdPerUnitGasUpdated(
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

  // Reverts

  function test_OnlyCallableByUpdaterOrOwner_Revert() public {
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: new Internal.TokenPriceUpdate[](0),
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    });

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.OnlyCallableByUpdaterOrOwner.selector));
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

  function test_StaleFeeToken_Revert() public {
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](1);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceTokens[0], usdPerToken: 4e18});
    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: tokenPriceUpdates, gasPriceUpdates: new Internal.GasPriceUpdate[](0)});
    s_priceRegistry.updatePrices(priceUpdates);

    vm.expectRevert(
      abi.encodeWithSelector(
        PriceRegistry.StaleTokenPrice.selector, s_weth, uint128(TWELVE_HOURS), uint128(TWELVE_HOURS + 1)
      )
    );
    s_priceRegistry.convertTokenAmount(s_weth, 3e16, s_sourceTokens[0]);
  }

  function test_LinkTokenNotSupported_Revert() public {
    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.convertTokenAmount(DUMMY_CONTRACT_ADDRESS, 3e16, s_sourceTokens[0]);

    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.convertTokenAmount(s_sourceTokens[0], 3e16, DUMMY_CONTRACT_ADDRESS);
  }

  function test_StaleLinkToken_Revert() public {
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](1);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_weth, usdPerToken: 18e17});
    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: tokenPriceUpdates, gasPriceUpdates: new Internal.GasPriceUpdate[](0)});
    s_priceRegistry.updatePrices(priceUpdates);

    vm.expectRevert(
      abi.encodeWithSelector(
        PriceRegistry.StaleTokenPrice.selector, s_sourceTokens[0], uint128(TWELVE_HOURS), uint128(TWELVE_HOURS + 1)
      )
    );
    s_priceRegistry.convertTokenAmount(s_weth, 3e16, s_sourceTokens[0]);
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

  function test_StaleTokenPrice_Revert() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);

    Internal.GasPriceUpdate[] memory gasPriceUpdates = new Internal.GasPriceUpdate[](1);
    gasPriceUpdates[0] =
      Internal.GasPriceUpdate({destChainSelector: DEST_CHAIN_SELECTOR, usdPerUnitGas: PACKED_USD_PER_GAS});

    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: new Internal.TokenPriceUpdate[](0), gasPriceUpdates: gasPriceUpdates});
    s_priceRegistry.updatePrices(priceUpdates);

    vm.expectRevert(
      abi.encodeWithSelector(PriceRegistry.StaleTokenPrice.selector, s_sourceTokens[0], TWELVE_HOURS, diff)
    );
    s_priceRegistry.getTokenAndGasPrices(s_sourceTokens[0], DEST_CHAIN_SELECTOR);
  }
}

contract PriceRegistry_updateTokenPriceFeeds is PriceRegistrySetup {
  event DataFeedPerTokenUpdated(address indexed token, address dataFeedAddress);

  function test_ZeroFeeds_Success() public {
    Vm.Log[] memory logEntries = vm.getRecordedLogs();

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](0);
    vm.recordLogs();
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    // Verify no log emissions
    assertEq(logEntries.length, 0);
  }

  function test_SingleFeedUpdate_Success() public {
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]]);

    assertEq(s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[0].sourceToken), address(0));

    vm.expectEmit();
    emit DataFeedPerTokenUpdated(tokenPriceFeedUpdates[0].sourceToken, tokenPriceFeedUpdates[0].dataFeedAddress);

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    assertEq(
      s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].dataFeedAddress
    );
  }

  function test_MultipleFeedUpdate_Success() public {
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](2);

    for (uint256 i = 0; i < 2; ++i) {
      tokenPriceFeedUpdates[i] =
        getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[i], s_dataFeedByToken[s_sourceTokens[i]]);

      assertEq(s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[i].sourceToken), address(0));

      vm.expectEmit();
      emit DataFeedPerTokenUpdated(tokenPriceFeedUpdates[i].sourceToken, tokenPriceFeedUpdates[i].dataFeedAddress);
    }

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    assertEq(
      s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].dataFeedAddress
    );
    assertEq(
      s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[1].sourceToken), tokenPriceFeedUpdates[1].dataFeedAddress
    );
  }

  function test_FeedUnset_Success() public {
    Internal.TimestampedPackedUint224 memory priceQueryInitial = s_priceRegistry.getTokenPrice(s_sourceTokens[0]);
    assertFalse(priceQueryInitial.value == 0);
    assertFalse(priceQueryInitial.timestamp == 0);

    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]]);

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    assertEq(
      s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[0].sourceToken), tokenPriceFeedUpdates[0].dataFeedAddress
    );

    tokenPriceFeedUpdates[0].dataFeedAddress = address(0);
    vm.expectEmit();
    emit DataFeedPerTokenUpdated(tokenPriceFeedUpdates[0].sourceToken, address(0));

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
    assertEq(s_priceRegistry.getTokenPriceFeed(tokenPriceFeedUpdates[0].sourceToken), address(0));

    // Price data should remain after a feed has been set->unset
    Internal.TimestampedPackedUint224 memory priceQueryPostUnsetFeed = s_priceRegistry.getTokenPrice(s_sourceTokens[0]);
    assertEq(priceQueryPostUnsetFeed.value, priceQueryInitial.value);
    assertEq(priceQueryPostUnsetFeed.timestamp, priceQueryInitial.timestamp);
  }

  function test_SingleFeedUpdateSkipTokenPriceUnset_Success() public {
    address tokenAddress = address(42);
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] = getSingleTokenPriceFeedUpdateStruct(tokenAddress, s_dataFeedByToken[s_sourceTokens[0]]);

    vm.expectEmit();
    emit DataFeedPerTokenUpdated(tokenAddress, tokenPriceFeedUpdates[0].dataFeedAddress);

    vm.recordLogs();
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    // Only 1 event for feed update
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 1);
  }

  // Reverts

  function test_FeedNotUpdated_Revert() public {
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]]);

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);

    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.DataFeedPerTokenNotUpdated.selector, s_sourceTokens[0]));
    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
  }

  function test_FeedUpdatedByNonOwner_Revert() public {
    IPriceRegistry.TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates = new IPriceRegistry.TokenPriceFeedUpdate[](1);
    tokenPriceFeedUpdates[0] =
      getSingleTokenPriceFeedUpdateStruct(s_sourceTokens[0], s_dataFeedByToken[s_sourceTokens[0]]);

    vm.startPrank(STRANGER);
    vm.expectRevert("Only callable by owner");

    s_priceRegistry.updateTokenPriceFeeds(tokenPriceFeedUpdates);
  }
}
