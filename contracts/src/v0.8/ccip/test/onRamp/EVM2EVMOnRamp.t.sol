// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMOnRampSetup.t.sol";
import {IEVM2EVMOnRamp} from "../../interfaces/onRamp/IEVM2EVMOnRamp.sol";

/// @notice #constructor
contract EVM2EVMOnRamp_constructor is EVM2EVMOnRampSetup {
  event ConfigSet(IEVM2EVMOnRamp.StaticConfig staticConfig, IEVM2EVMOnRamp.DynamicConfig dynamicConfig);

  function testConstructorSuccess() public {
    IEVM2EVMOnRamp.StaticConfig memory staticConfig = IEVM2EVMOnRamp.StaticConfig({
      linkToken: s_sourceTokens[0],
      chainId: SOURCE_CHAIN_ID,
      destChainId: DEST_CHAIN_ID,
      defaultTxGasLimit: GAS_LIMIT
    });
    IEVM2EVMOnRamp.DynamicConfig memory dynamicConfig = generateDynamicOnRampConfig(
      address(s_sourceRouter),
      address(s_priceRegistry),
      address(0),
      address(s_afn)
    );

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    s_onRamp = new EVM2EVMOnRamp(
      staticConfig,
      dynamicConfig,
      getTokensAndPools(s_sourceTokens, getCastedSourcePools()),
      new address[](0),
      rateLimiterConfig(),
      s_feeTokenConfigArgs,
      getNopsAndWeights()
    );

    IEVM2EVMOnRamp.StaticConfig memory gotStaticConfig = s_onRamp.getStaticConfig();

    assertEq(staticConfig.linkToken, gotStaticConfig.linkToken);
    assertEq(staticConfig.chainId, gotStaticConfig.chainId);
    assertEq(staticConfig.destChainId, gotStaticConfig.destChainId);
    assertEq(staticConfig.defaultTxGasLimit, gotStaticConfig.defaultTxGasLimit);

    IEVM2EVMOnRamp.DynamicConfig memory gotDynamicConfig = s_onRamp.getDynamicConfig();

    assertEq(dynamicConfig.router, gotDynamicConfig.router);
    assertEq(dynamicConfig.priceRegistry, gotDynamicConfig.priceRegistry);
    assertEq(dynamicConfig.maxDataSize, gotDynamicConfig.maxDataSize);
    assertEq(dynamicConfig.maxTokensLength, gotDynamicConfig.maxTokensLength);
    assertEq(dynamicConfig.maxGasLimit, gotDynamicConfig.maxGasLimit);
    assertEq(dynamicConfig.feeAdmin, gotDynamicConfig.feeAdmin);
    assertEq(dynamicConfig.afn, gotDynamicConfig.afn);

    // Tokens
    assertEq(s_sourceTokens, s_onRamp.getSupportedTokens());

    // Initial values
    assertEq("EVM2EVMOnRamp 1.0.0", s_onRamp.typeAndVersion());
    assertEq(OWNER, s_onRamp.owner());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());
  }
}

contract EVM2EVMOnRamp_payNops_fuzz is EVM2EVMOnRampSetup {
  function test_fuzz_NopPayNopsSuccess(uint96 nopFeesJuels) public {
    (IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal) = s_onRamp.getNops();
    // To avoid NoFeesToPay
    vm.assume(nopFeesJuels > weightsTotal);

    // Set Nop fee juels
    deal(s_sourceFeeToken, address(s_onRamp), nopFeesJuels);
    changePrank(address(s_sourceRouter));
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), nopFeesJuels, OWNER);

    changePrank(OWNER);

    uint256 totalJuels = s_onRamp.getNopFeesJuels();
    s_onRamp.payNops();
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      uint256 expectedPayout = (totalJuels * nopsAndWeights[i].weight) / weightsTotal;
      assertEq(IERC20(s_sourceFeeToken).balanceOf(nopsAndWeights[i].nop), expectedPayout);
    }
  }
}

contract EVM2EVMOnRamp_payNops is EVM2EVMOnRampSetup {
  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_sourceRouter));

    uint256 feeAmount = 1234567890;
    uint256 numberOfMessages = 5;

    // Send a bunch of messages, increasing the juels in the contract
    for (uint256 i = 0; i < numberOfMessages; ++i) {
      IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);
      s_onRamp.forwardFromRouter(_generateEmptyMessage(), feeAmount, OWNER);
    }

    assertEq(s_onRamp.getNopFeesJuels(), feeAmount * numberOfMessages);
    assertEq(IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp)), feeAmount * numberOfMessages);
  }

  function testOwnerPayNopsSuccess() public {
    changePrank(OWNER);

    uint256 totalJuels = s_onRamp.getNopFeesJuels();
    s_onRamp.payNops();
    (IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal) = s_onRamp.getNops();
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      uint256 expectedPayout = (nopsAndWeights[i].weight * totalJuels) / weightsTotal;
      assertEq(IERC20(s_sourceFeeToken).balanceOf(nopsAndWeights[i].nop), expectedPayout);
    }
  }

  function testFeeAdminPayNopsSuccess() public {
    changePrank(s_feeAdmin);

    uint256 totalJuels = s_onRamp.getNopFeesJuels();
    s_onRamp.payNops();
    (IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal) = s_onRamp.getNops();
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      uint256 expectedPayout = (nopsAndWeights[i].weight * totalJuels) / weightsTotal;
      assertEq(IERC20(s_sourceFeeToken).balanceOf(nopsAndWeights[i].nop), expectedPayout);
    }
  }

  function testNopPayNopsSuccess() public {
    changePrank(getNopsAndWeights()[0].nop);

    uint256 totalJuels = s_onRamp.getNopFeesJuels();
    s_onRamp.payNops();
    (IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal) = s_onRamp.getNops();
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      uint256 expectedPayout = (nopsAndWeights[i].weight * totalJuels) / weightsTotal;
      assertEq(IERC20(s_sourceFeeToken).balanceOf(nopsAndWeights[i].nop), expectedPayout);
    }
  }

  // Reverts

  function testInsufficientBalanceReverts() public {
    changePrank(address(s_onRamp));
    IERC20(s_sourceFeeToken).transfer(OWNER, IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp)));
    changePrank(OWNER);
    vm.expectRevert(IEVM2EVMOnRamp.InsufficientBalance.selector);
    s_onRamp.payNops();
  }

  function testWrongPermissionsReverts() public {
    changePrank(STRANGER);

    vm.expectRevert(IEVM2EVMOnRamp.OnlyCallableByOwnerOrFeeAdminOrNop.selector);
    s_onRamp.payNops();
  }

  function testNoFeesToPayReverts() public {
    changePrank(OWNER);
    s_onRamp.payNops();
    vm.expectRevert(IEVM2EVMOnRamp.NoFeesToPay.selector);
    s_onRamp.payNops();
  }

  function testNoNopsToPayReverts() public {
    changePrank(OWNER);
    IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = new IEVM2EVMOnRamp.NopAndWeight[](0);
    s_onRamp.setNops(nopsAndWeights);
    vm.expectRevert(IEVM2EVMOnRamp.NoNopsToPay.selector);
    s_onRamp.payNops();
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMOnRamp_forwardFromRouter is EVM2EVMOnRampSetup {
  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceTokens[1];
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_sourceRouter));
  }

  function testForwardFromRouterSuccessCustomExtraArgs() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT * 2, strict: true}));
    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

    vm.expectEmit();
    emit CCIPSendRequested(_messageToEvent(message, 1, 1, feeAmount, OWNER));

    s_onRamp.forwardFromRouter(message, feeAmount, OWNER);
  }

  function testForwardFromRouterSuccess() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

    vm.expectEmit();
    emit CCIPSendRequested(_messageToEvent(message, 1, 1, feeAmount, OWNER));

    s_onRamp.forwardFromRouter(message, feeAmount, OWNER);
  }

  function testShouldIncrementSeqNumAndNonceSuccess() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    for (uint64 i = 1; i < 4; ++i) {
      uint64 nonceBefore = s_onRamp.getSenderNonce(OWNER);

      vm.expectEmit();
      emit CCIPSendRequested(_messageToEvent(message, i, i, 0, OWNER));

      s_onRamp.forwardFromRouter(message, 0, OWNER);

      uint64 nonceAfter = s_onRamp.getSenderNonce(OWNER);
      assertEq(nonceAfter, nonceBefore + 1);
    }
  }

  event Transfer(address indexed from, address indexed to, uint256 value);

  function testShouldStoreLinkFees() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

    s_onRamp.forwardFromRouter(message, feeAmount, OWNER);

    assertEq(IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp)), feeAmount);
    assertEq(s_onRamp.getNopFeesJuels(), feeAmount);
  }

  function testShouldStoreNonLinkFees() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceTokens[1];

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceTokens[1]).transferFrom(OWNER, address(s_onRamp), feeAmount);

    s_onRamp.forwardFromRouter(message, feeAmount, OWNER);

    assertEq(IERC20(s_sourceTokens[1]).balanceOf(address(s_onRamp)), feeAmount);

    // Calculate conversion done by prices contract
    uint256 feeTokenPrice = s_priceRegistry.getTokenPrice(s_sourceTokens[1]).value;
    uint256 linkTokenPrice = s_priceRegistry.getTokenPrice(s_sourceFeeToken).value;
    uint256 conversionRate = (feeTokenPrice * 1e18) / linkTokenPrice;
    uint256 expectedJuels = (feeAmount * conversionRate) / 1e18;

    assertEq(s_onRamp.getNopFeesJuels(), expectedJuels);
  }

  // Make sure any valid sender, receiver and feeAmount can be handled.
  function test_fuzz_ForwardFromRouterSuccess(
    address originalSender,
    address receiver,
    uint96 feeTokenAmount
  ) public {
    // To avoid RouterMustSetOriginalSender
    vm.assume(originalSender != address(0));

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encode(receiver);

    // Make sure the tokens are in the contract
    deal(s_sourceFeeToken, address(s_onRamp), feeTokenAmount);

    Internal.EVM2EVMMessage memory expectedEvent = _messageToEvent(message, 1, 1, feeTokenAmount, originalSender);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(expectedEvent);

    // Assert the message Id is correct
    assertEq(expectedEvent.messageId, s_onRamp.forwardFromRouter(message, feeTokenAmount, originalSender));
    // Assert the fee token amount is correctly assigned to the nop fee pool
    assertEq(feeTokenAmount, s_onRamp.getNopFeesJuels());
  }

  // Reverts

  function testPausedReverts() public {
    changePrank(OWNER);
    s_onRamp.pause();
    vm.expectRevert("Pausable: paused");
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, OWNER);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(IEVM2EVMOnRamp.BadAFNSignal.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, OWNER);
  }

  function testPermissionsReverts() public {
    changePrank(OWNER);
    vm.expectRevert(IEVM2EVMOnRamp.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, OWNER);
  }

  function testOriginalSenderReverts() public {
    vm.expectRevert(IEVM2EVMOnRamp.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, address(0));
  }

  function testMessageTooLargeReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.data = new bytes(MAX_DATA_SIZE + 1);
    vm.expectRevert(
      abi.encodeWithSelector(IEVM2EVMOnRamp.MessageTooLarge.selector, MAX_DATA_SIZE, message.data.length)
    );

    s_onRamp.forwardFromRouter(message, 0, STRANGER);
  }

  function testTooManyTokensReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 tooMany = MAX_TOKENS_LENGTH + 1;
    message.tokenAmounts = new Client.EVMTokenAmount[](tooMany);
    vm.expectRevert(IEVM2EVMOnRamp.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, 0, STRANGER);
  }

  function testSenderNotAllowedReverts() public {
    changePrank(OWNER);
    s_onRamp.setAllowListEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.SenderNotAllowed.selector, STRANGER));
    changePrank(address(s_sourceRouter));
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    address wrongToken = address(1);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].token = wrongToken;
    message.tokenAmounts[0].amount = 1;

    // We need to set the price of this new token to be able to reach
    // the proper revert point. This must be called by the owner.
    changePrank(OWNER);
    uint256[] memory prices = new uint256[](1);
    prices[0] = 1;
    s_onRamp.setPrices(abi.decode(abi.encode(message.tokenAmounts), (IERC20[])), prices);

    // Change back to the router
    changePrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.UnsupportedToken.selector, wrongToken));

    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  function testValueExceedsCapacityReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].amount = 2**128;
    message.tokenAmounts[0].token = s_sourceTokens[0];

    IERC20(s_sourceTokens[0]).approve(address(s_onRamp), 2**128);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAggregateRateLimiter.ValueExceedsCapacity.selector,
        rateLimiterConfig().capacity,
        message.tokenAmounts[0].amount * getTokenPrices()[0]
      )
    );

    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  function testPriceNotFoundForTokenReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    address fakeToken = address(1);
    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].token = fakeToken;

    vm.expectRevert(abi.encodeWithSelector(IAggregateRateLimiter.PriceNotFoundForToken.selector, fakeToken));

    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  // Asserts gasLimit must be <=maxGasLimit
  function testMessageGasLimitTooHighReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: MAX_GAS_LIMIT + 1, strict: false}));
    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.MessageGasLimitTooHigh.selector));
    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  function testInvalidAddressEncodePackedReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encodePacked(address(234));

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.InvalidAddress.selector, message.receiver));

    s_onRamp.forwardFromRouter(message, 1, OWNER);
  }

  function testInvalidAddressReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encode(type(uint208).max);

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.InvalidAddress.selector, message.receiver));

    s_onRamp.forwardFromRouter(message, 1, OWNER);
  }
}

contract EVM2EVMOnRamp_setNops is EVM2EVMOnRampSetup {
  // Used because EnumerableMap doesn't guarantee order
  mapping(address => uint256) internal s_nopsToWeights;

  function testSetNopsSuccess() public {
    IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = getNopsAndWeights();
    nopsAndWeights[1].nop = USER_4;
    nopsAndWeights[1].weight = 20;
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      s_nopsToWeights[nopsAndWeights[i].nop] = nopsAndWeights[i].weight;
    }

    s_onRamp.setNops(nopsAndWeights);

    (IEVM2EVMOnRamp.NopAndWeight[] memory actual, ) = s_onRamp.getNops();
    for (uint256 i = 0; i < actual.length; ++i) {
      assertEq(actual[i].weight, s_nopsToWeights[actual[i].nop]);
    }
  }

  function testSetNopsRemovesOldNopsCompletelySuccess() public {
    IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = new IEVM2EVMOnRamp.NopAndWeight[](0);
    s_onRamp.setNops(nopsAndWeights);
    (IEVM2EVMOnRamp.NopAndWeight[] memory actual, uint256 totalWeight) = s_onRamp.getNops();
    assertEq(actual.length, 0);
    assertEq(totalWeight, 0);
  }

  function testNonOwnerReverts() public {
    IEVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = getNopsAndWeights();
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");

    s_onRamp.setNops(nopsAndWeights);
  }
}

/// @notice #withdrawNonLinkFees
contract EVM2EVMOnRamp_withdrawNonLinkFees is EVM2EVMOnRampSetup {
  IERC20 internal s_token;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();
    s_token = IERC20(s_sourceTokens[1]);
    changePrank(OWNER);
    s_token.transfer(address(s_onRamp), 100);
  }

  function testwithdrawNonLinkFeesSuccess() public {
    IEVM2EVMOnRamp(s_onRamp).withdrawNonLinkFees(address(s_token), address(this));

    assertEq(0, s_token.balanceOf(address(s_onRamp)));
    assertEq(100, s_token.balanceOf(address(this)));
  }

  function testNonOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");
    IEVM2EVMOnRamp(s_onRamp).withdrawNonLinkFees(address(s_token), address(this));
  }

  function testInvalidWithdrawalAddressReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.InvalidWithdrawalAddress.selector, address(0)));
    IEVM2EVMOnRamp(s_onRamp).withdrawNonLinkFees(address(s_token), address(0));
  }

  function testInvalidTokenReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.InvalidFeeToken.selector, s_sourceTokens[0]));
    IEVM2EVMOnRamp(s_onRamp).withdrawNonLinkFees(s_sourceTokens[0], address(this));
  }
}

/// @notice #setFeeConfig
contract EVM2EVMOnRamp_setFeeConfig is EVM2EVMOnRampSetup {
  event FeeConfigSet(IEVM2EVMOnRamp.FeeTokenConfigArgs[] feeConfig);

  function testSetFeeConfigSuccess() public {
    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeConfig;

    vm.expectEmit();
    emit FeeConfigSet(feeConfig);

    s_onRamp.setFeeConfig(feeConfig);
  }

  function testSetFeeConfigByFeeAdminSuccess() public {
    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeConfig;

    changePrank(s_feeAdmin);

    vm.expectEmit();
    emit FeeConfigSet(feeConfig);

    s_onRamp.setFeeConfig(feeConfig);
  }

  // Reverts

  function testOnlyCallableByOwnerOrFeeAdminReverts() public {
    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeConfig;
    changePrank(STRANGER);

    vm.expectRevert(IEVM2EVMOnRamp.OnlyCallableByOwnerOrFeeAdmin.selector);

    s_onRamp.setFeeConfig(feeConfig);
  }
}

// #getTokenPool
contract EVM2EVMOnRamp_getTokenPool is EVM2EVMOnRampSetup {
  function testSuccess() public {
    assertEq(s_sourcePools[0], address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[0]))));
    assertEq(s_sourcePools[1], address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[1]))));

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.UnsupportedToken.selector, IERC20(s_destTokens[0])));
    s_onRamp.getPoolBySourceToken(IERC20(s_destTokens[0]));
  }
}

contract EVM2EVMOnRamp_applyPoolUpdates is EVM2EVMOnRampSetup {
  event PoolAdded(address token, address pool);
  event PoolRemoved(address token, address pool);

  function testApplyPoolUpdatesSuccess() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0] = Internal.PoolUpdate({token: address(1), pool: address(2)});

    vm.expectEmit();
    emit PoolAdded(adds[0].token, adds[0].pool);

    s_onRamp.applyPoolUpdates(adds, new Internal.PoolUpdate[](0));

    assertEq(adds[0].pool, address(s_onRamp.getPoolBySourceToken(IERC20(adds[0].token))));

    vm.expectEmit();
    emit PoolRemoved(adds[0].token, adds[0].pool);

    s_onRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), adds);

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.UnsupportedToken.selector, adds[0].token));
    s_onRamp.getPoolBySourceToken(IERC20(adds[0].token));
  }

  // Reverts
  function testOnlyCallableByOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");

    s_onRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), new Internal.PoolUpdate[](0));
  }

  function testPoolAlreadyExistsReverts() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](2);
    adds[0] = Internal.PoolUpdate({token: address(1), pool: address(2)});
    adds[1] = Internal.PoolUpdate({token: address(1), pool: address(2)});

    vm.expectRevert(IEVM2EVMOnRamp.PoolAlreadyAdded.selector);

    s_onRamp.applyPoolUpdates(adds, new Internal.PoolUpdate[](0));
  }

  function testInvalidTokenPoolConfigReverts() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0] = Internal.PoolUpdate({token: address(0), pool: address(2)});

    vm.expectRevert(IEVM2EVMOnRamp.InvalidTokenPoolConfig.selector);

    s_onRamp.applyPoolUpdates(adds, new Internal.PoolUpdate[](0));

    adds[0] = Internal.PoolUpdate({token: address(1), pool: address(0)});

    vm.expectRevert(IEVM2EVMOnRamp.InvalidTokenPoolConfig.selector);

    s_onRamp.applyPoolUpdates(adds, new Internal.PoolUpdate[](0));
  }

  function testPoolDoesNotExistReverts() public {
    Internal.PoolUpdate[] memory removes = new Internal.PoolUpdate[](1);
    removes[0] = Internal.PoolUpdate({token: address(1), pool: address(2)});

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOnRamp.PoolDoesNotExist.selector, removes[0].token));

    s_onRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), removes);
  }

  function testTokenPoolMismatchReverts() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0] = Internal.PoolUpdate({token: address(1), pool: address(2)});
    Internal.PoolUpdate[] memory removes = new Internal.PoolUpdate[](1);
    removes[0] = Internal.PoolUpdate({token: address(1), pool: address(20)});

    vm.expectRevert(IEVM2EVMOnRamp.TokenPoolMismatch.selector);

    s_onRamp.applyPoolUpdates(adds, removes);
  }
}

// #getSupportedTokens
contract EVM2EVMOnRamp_getSupportedTokens is EVM2EVMOnRampSetup {
  function testGetSupportedTokensSuccess() public {
    address[] memory supportedTokens = s_onRamp.getSupportedTokens();

    assertEq(s_sourceTokens, supportedTokens);

    Internal.PoolUpdate[] memory removes = new Internal.PoolUpdate[](1);
    removes[0] = Internal.PoolUpdate({token: s_sourceTokens[0], pool: s_sourcePools[0]});

    s_onRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), removes);

    supportedTokens = s_onRamp.getSupportedTokens();

    assertEq(address(s_sourceTokens[1]), supportedTokens[0]);
    assertEq(s_sourceTokens.length - 1, supportedTokens.length);
  }
}

// #getExpectedNextSequenceNumber
contract EVM2EVMOnRamp_getExpectedNextSequenceNumber is EVM2EVMOnRampSetup {
  function testSuccess() public {
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());
  }
}

// #setDynamicConfig
contract EVM2EVMOnRamp_setDynamicConfig is EVM2EVMOnRampSetup {
  event ConfigSet(IEVM2EVMOnRamp.StaticConfig staticConfig, IEVM2EVMOnRamp.DynamicConfig dynamicConfig);

  function testSuccess() public {
    IEVM2EVMOnRamp.StaticConfig memory staticConfig = s_onRamp.getStaticConfig();
    IEVM2EVMOnRamp.DynamicConfig memory newConfig = IEVM2EVMOnRamp.DynamicConfig({
      router: address(2134),
      priceRegistry: address(23423),
      maxDataSize: 400,
      maxTokensLength: 14,
      maxGasLimit: MAX_GAS_LIMIT / 2,
      feeAdmin: address(98235),
      afn: address(11)
    });

    vm.expectEmit();
    emit ConfigSet(staticConfig, newConfig);

    s_onRamp.setDynamicConfig(newConfig);

    IEVM2EVMOnRamp.DynamicConfig memory gotDynamicConfig = s_onRamp.getDynamicConfig();
    assertEq(newConfig.router, gotDynamicConfig.router);
    assertEq(newConfig.priceRegistry, gotDynamicConfig.priceRegistry);
    assertEq(newConfig.maxDataSize, gotDynamicConfig.maxDataSize);
    assertEq(newConfig.maxTokensLength, gotDynamicConfig.maxTokensLength);
    assertEq(newConfig.maxGasLimit, gotDynamicConfig.maxGasLimit);
    assertEq(newConfig.feeAdmin, gotDynamicConfig.feeAdmin);
  }

  // Reverts
  function testSetConfigOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRamp.setDynamicConfig(generateDynamicOnRampConfig(address(1), address(2), address(3), address(4)));
  }
}

contract EVM2EVMOnRampWithAllowListSetup is EVM2EVMOnRampSetup {
  function setUp() public virtual override(EVM2EVMOnRampSetup) {
    EVM2EVMOnRampSetup.setUp();
    address[] memory allowedAddresses = new address[](1);
    allowedAddresses[0] = OWNER;
    s_onRamp.applyAllowListUpdates(allowedAddresses, new address[](0));
    s_onRamp.setAllowListEnabled(true);
  }
}

contract EVM2EVMOnRamp_setAllowListEnabled is EVM2EVMOnRampWithAllowListSetup {
  function testSuccess() public {
    assertTrue(s_onRamp.getAllowListEnabled());
    s_onRamp.setAllowListEnabled(false);
    assertFalse(s_onRamp.getAllowListEnabled());
    s_onRamp.setAllowListEnabled(true);
    assertTrue(s_onRamp.getAllowListEnabled());
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRamp.setAllowListEnabled(true);
  }
}

/// @notice #getAllowListEnabled
contract EVM2EVMOnRamp_getAllowListEnabled is EVM2EVMOnRampWithAllowListSetup {
  function testSuccess() public {
    assertTrue(s_onRamp.getAllowListEnabled());
    s_onRamp.setAllowListEnabled(false);
    assertFalse(s_onRamp.getAllowListEnabled());
    s_onRamp.setAllowListEnabled(true);
    assertTrue(s_onRamp.getAllowListEnabled());
  }
}

/// @notice #setAllowList
contract EVM2EVMOnRamp_applyAllowListUpdates is EVM2EVMOnRampWithAllowListSetup {
  event AllowListAdd(address sender);
  event AllowListRemove(address sender);

  function testSuccess() public {
    address[] memory newAddresses = new address[](2);
    newAddresses[0] = address(1);
    newAddresses[1] = address(2);

    for (uint256 i = 0; i < 2; ++i) {
      vm.expectEmit();
      emit AllowListAdd(newAddresses[i]);
    }

    s_onRamp.applyAllowListUpdates(newAddresses, new address[](0));
    address[] memory setAddresses = s_onRamp.getAllowList();

    // First address is owner.
    assertEq(newAddresses[0], setAddresses[1]);
    assertEq(newAddresses[1], setAddresses[2]);
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    address[] memory newAddresses = new address[](2);
    s_onRamp.applyAllowListUpdates(newAddresses, new address[](0));
  }
}

/// @notice #getAllowList
contract EVM2EVMOnRamp_getAllowList is EVM2EVMOnRampWithAllowListSetup {
  function testSuccess() public {
    address[] memory setAddresses = s_onRamp.getAllowList();
    assertEq(OWNER, setAddresses[0]);
  }
}

contract EVM2EVMOnRamp_afn is EVM2EVMOnRampSetup {
  function testAFN() public {
    // Test pausing
    assertEq(s_onRamp.paused(), false);
    s_onRamp.pause();
    assertEq(s_onRamp.paused(), true);
    s_onRamp.unpause();
    assertEq(s_onRamp.paused(), false);

    // Test afn
    assertEq(s_onRamp.isAFNHealthy(), true);
    s_afn.voteBad();
    assertEq(s_onRamp.isAFNHealthy(), false);
    s_afn.recoverFromBadSignal();
    assertEq(s_onRamp.isAFNHealthy(), true);
  }
}
