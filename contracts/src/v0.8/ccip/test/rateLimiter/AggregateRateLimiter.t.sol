// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../mocks/MockERC20.sol";
import "../BaseTest.t.sol";
import "../helpers/AggregateRateLimiterHelper.sol";
import "../../interfaces/rateLimiter/AggregateRateLimiterInterface.sol";

contract AggregateTokenLimiterSetup is BaseTest {
  AggregateRateLimiterHelper s_rateLimiter;
  AggregateRateLimiterInterface.RateLimiterConfig s_config;

  IERC20 constant TOKEN = ERC20(0x21118E64E1fB0c487F25Dd6d3601FF6af8D32E4e);
  uint256 constant TOKEN_PRICE = 4;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_config = AggregateRateLimiterInterface.RateLimiterConfig({rate: 5, capacity: 100});
    s_rateLimiter = new AggregateRateLimiterHelper(s_config, TOKEN_LIMIT_ADMIN);
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = TOKEN;
    uint256[] memory prices = new uint256[](1);
    prices[0] = TOKEN_PRICE;
    s_rateLimiter.setPrices(tokens, prices);
  }
}

/// @notice #constructor
contract AggregateTokenLimiter_constructor is AggregateTokenLimiterSetup {
  // Success
  function testSuccess() public {
    assertEq(TOKEN_LIMIT_ADMIN, s_rateLimiter.getTokenLimitAdmin());
    assertEq(OWNER, s_rateLimiter.owner());

    AggregateRateLimiterInterface.TokenBucket memory tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();
    assertEq(s_config.rate, tokenBucket.rate);
    assertEq(s_config.capacity, tokenBucket.capacity);
    assertEq(s_config.capacity, tokenBucket.tokens);
    assertEq(BLOCK_TIME, tokenBucket.lastUpdated);
  }
}

/// @notice #getTokenLimitAdmin
contract AggregateTokenLimiter_getTokenLimitAdmin is AggregateTokenLimiterSetup {
  // Success
  function testSuccess() public {
    assertEq(TOKEN_LIMIT_ADMIN, s_rateLimiter.getTokenLimitAdmin());
  }
}

/// @notice #setTokenLimitAdmin
contract AggregateTokenLimiter_setTokenLimitAdmin is AggregateTokenLimiterSetup {
  // Success
  function testOwnerSuccess() public {
    s_rateLimiter.setTokenLimitAdmin(STRANGER);
    assertEq(STRANGER, s_rateLimiter.getTokenLimitAdmin());
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");

    s_rateLimiter.setTokenLimitAdmin(STRANGER);
  }
}

/// @notice #getTokenBucket
contract AggregateTokenLimiter_getTokenBucket is AggregateTokenLimiterSetup {
  // Success
  function testSuccess() public {
    AggregateRateLimiterInterface.TokenBucket memory tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();
    assertEq(s_config.rate, tokenBucket.rate);
    assertEq(s_config.capacity, tokenBucket.capacity);
    assertEq(s_config.capacity, tokenBucket.tokens);
    assertEq(BLOCK_TIME, tokenBucket.lastUpdated);
  }

  function testRefillSuccess() public {
    s_config.capacity = s_config.capacity * 2;
    s_rateLimiter.setRateLimiterConfig(s_config);

    AggregateRateLimiterInterface.TokenBucket memory tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();

    assertEq(s_config.rate, tokenBucket.rate);
    assertEq(s_config.capacity, tokenBucket.capacity);
    assertEq(s_config.capacity / 2, tokenBucket.tokens);
    assertEq(BLOCK_TIME, tokenBucket.lastUpdated);

    uint256 warpTime = 4;
    vm.warp(BLOCK_TIME + warpTime);

    tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();

    assertEq(s_config.rate, tokenBucket.rate);
    assertEq(s_config.capacity, tokenBucket.capacity);
    assertEq(s_config.capacity / 2 + warpTime * s_config.rate, tokenBucket.tokens);
    assertEq(BLOCK_TIME + warpTime, tokenBucket.lastUpdated);

    vm.warp(BLOCK_TIME + warpTime * 100);

    // Bucket overflow
    tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();
    assertEq(s_config.capacity, tokenBucket.tokens);
  }

  // Reverts

  function testTimeUnderflowReverts() public {
    vm.warp(BLOCK_TIME - 1);

    vm.expectRevert(stdError.arithmeticError);
    s_rateLimiter.calculateCurrentTokenBucketState();
  }
}

/// @notice #setRateLimiterConfig
contract AggregateTokenLimiter_setRateLimiterConfig is AggregateTokenLimiterSetup {
  event ConfigChanged(uint256 capacity, uint256 rate);

  // Success
  function testOwnerSuccess() public {
    setConfig();
  }

  function testTokenLimitAdminSuccess() public {
    changePrank(TOKEN_LIMIT_ADMIN);
    setConfig();
  }

  function setConfig() private {
    AggregateRateLimiterInterface.TokenBucket memory tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();
    assertEq(s_config.rate, tokenBucket.rate);
    assertEq(s_config.capacity, tokenBucket.capacity);

    s_config = AggregateRateLimiterInterface.RateLimiterConfig({
      rate: tokenBucket.rate * 2,
      capacity: tokenBucket.capacity * 8
    });

    vm.expectEmit(false, false, false, true);
    emit ConfigChanged(s_config.capacity, s_config.rate);

    s_rateLimiter.setRateLimiterConfig(s_config);

    tokenBucket = s_rateLimiter.calculateCurrentTokenBucketState();
    assertEq(s_config.rate, tokenBucket.rate);
    assertEq(s_config.capacity, tokenBucket.capacity);
  }

  // Reverts

  function testOnlyOnlyCallableByAdminOrOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert(AggregateRateLimiterInterface.OnlyCallableByAdminOrOwner.selector);

    s_rateLimiter.setRateLimiterConfig(s_config);
  }

  function testRefillRateTooHighReverts() public {
    AggregateRateLimiterInterface.RateLimiterConfig memory config = AggregateRateLimiterInterface.RateLimiterConfig({
      rate: 2**208,
      capacity: 100
    });

    vm.expectRevert(AggregateRateLimiterInterface.RefillRateTooHigh.selector);

    s_rateLimiter.setRateLimiterConfig(config);
  }
}

/// @notice #getPricesForTokens
contract AggregateTokenLimiter_getPricesForTokens is AggregateTokenLimiterSetup {
  // Success
  function testSuccess() public {
    IERC20[] memory tokens = new IERC20[](2);
    // Unknown tokens
    tokens[0] = ERC20(0x31118E64E1fb0c487f25DD6D3601FF6Af8D32e4E);
    // Zero token
    tokens[0] = ERC20(address(0));
    // Known token
    tokens[1] = TOKEN;
    uint256[] memory prices = new uint256[](2);
    prices[0] = 0;
    prices[0] = 0;
    prices[1] = TOKEN_PRICE;

    uint256[] memory actualPrices = s_rateLimiter.getPricesForTokens(tokens);

    assertEq(actualPrices, prices);
  }
}

/// @notice #setPrices
contract AggregateTokenLimiter_setPrices is AggregateTokenLimiterSetup {
  event TokenPriceChanged(address token, uint256 newPrice);
  IERC20[] s_tokens;
  uint256[] s_prices;

  function setUp() public virtual override {
    AggregateTokenLimiterSetup.setUp();

    uint256 numberOfTokens = 15;
    IERC20[] memory tokens = new IERC20[](numberOfTokens);
    uint256[] memory prices = new uint256[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      tokens[i] = IERC20(address(uint160(i + 1)));
      prices[i] = TOKEN_PRICE * (i + 1);
    }

    s_rateLimiter.setPrices(tokens, prices);

    s_tokens = tokens;
    s_prices = prices;
  }

  // Success
  function testOwnerSuccess() public {
    setPrice();
  }

  function testTokenLimitAdminSuccess() public {
    changePrank(TOKEN_LIMIT_ADMIN);
    setPrice();
  }

  function setPrice() private {
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = TOKEN;
    uint256[] memory prices = new uint256[](1);
    prices[0] = TOKEN_PRICE * 2;

    vm.expectEmit(false, false, false, true);
    emit TokenPriceChanged(address(TOKEN), TOKEN_PRICE * 2);

    s_rateLimiter.setPrices(tokens, prices);

    assertEq(TOKEN_PRICE * 2, s_rateLimiter.getPricesForTokens(tokens)[0]);
  }

  function testClearExistingTokens() public {
    IERC20[] memory tokens = s_tokens;
    IERC20[] memory unsetTokens = new IERC20[](1);
    unsetTokens[0] = tokens[0];

    tokens[0] = IERC20(address(10000));

    // Assert the token has a price before being unset
    assertEq(TOKEN_PRICE, s_rateLimiter.getPricesForTokens(unsetTokens)[0]);

    s_rateLimiter.setPrices(tokens, s_prices);

    // Assert the token not being sent in the new setPrices request has no
    // corresponding price after the request.
    assertEq(0, s_rateLimiter.getPricesForTokens(unsetTokens)[0]);
  }

  // Reverts

  function testAddressCannotBeZeroReverts() public {
    vm.expectRevert(AggregateRateLimiterInterface.AddressCannotBeZero.selector);

    s_rateLimiter.setPrices(new IERC20[](1), new uint256[](1));
  }

  function testOnlyOnlyCallableByAdminOrOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert(AggregateRateLimiterInterface.OnlyCallableByAdminOrOwner.selector);

    s_rateLimiter.setPrices(new IERC20[](1), new uint256[](1));
  }

  function testTokensAndPriceLengthMismatchReverts() public {
    vm.expectRevert(AggregateRateLimiterInterface.TokensAndPriceLengthMismatch.selector);

    s_rateLimiter.setPrices(new IERC20[](2), new uint256[](1));
  }
}

/// @notice #_removeTokens
contract AggregateTokenLimiter__removeTokens is AggregateTokenLimiterSetup {
  event TokensRemovedFromBucket(uint256 tokens);

  // Success
  function testSuccess() public {
    // 15 (tokens) * 4 (price) * 2 (number of times) > 100 (capacity)
    uint256 numberOfTokens = 15;
    uint256 value = numberOfTokens * TOKEN_PRICE;

    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    tokensAndAmounts[0].token = address(TOKEN);
    tokensAndAmounts[0].amount = numberOfTokens;

    vm.expectEmit(false, false, false, true);
    emit TokensRemovedFromBucket(value);

    // Remove the value from the pool
    s_rateLimiter.removeTokens(tokensAndAmounts);

    // Get the updated bucket status
    AggregateRateLimiterInterface.TokenBucket memory bucket = s_rateLimiter.calculateCurrentTokenBucketState();
    // Assert the proper value has been taken out of the bucket
    assertEq(bucket.capacity - value, bucket.tokens);

    // Since value * 2 > bucket.capacity we cannot take it out twice.
    // Expect a revert when we try, with a wait time.
    uint256 waitTime = 4;
    vm.expectRevert(
      abi.encodeWithSelector(AggregateRateLimiterInterface.ValueExceedsAllowedThreshold.selector, waitTime)
    );
    s_rateLimiter.removeTokens(tokensAndAmounts);

    // Move the block time forward by 10 so the bucket refills by 10 * rate
    vm.warp(BLOCK_TIME + waitTime);

    // The bucket has filled up enough so we can take out more tokens
    s_rateLimiter.removeTokens(tokensAndAmounts);
    bucket = s_rateLimiter.calculateCurrentTokenBucketState();
    assertEq(bucket.capacity - value + waitTime * s_config.rate - value, bucket.tokens);
  }

  // Reverts

  function testUnknownTokenReverts() public {
    vm.expectRevert(abi.encodeWithSelector(AggregateRateLimiterInterface.PriceNotFoundForToken.selector, address(0)));
    s_rateLimiter.removeTokens(new CCIP.EVMTokenAndAmount[](1));
  }

  function testValueExceedsCapacityReverts() public {
    AggregateRateLimiterInterface.TokenBucket memory bucket = s_rateLimiter.calculateCurrentTokenBucketState();

    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    tokensAndAmounts[0].token = address(TOKEN);
    tokensAndAmounts[0].amount = 100;

    vm.expectRevert(
      abi.encodeWithSelector(
        AggregateRateLimiterInterface.ValueExceedsCapacity.selector,
        bucket.capacity,
        tokensAndAmounts[0].amount * TOKEN_PRICE
      )
    );
    s_rateLimiter.removeTokens(tokensAndAmounts);
  }
}
