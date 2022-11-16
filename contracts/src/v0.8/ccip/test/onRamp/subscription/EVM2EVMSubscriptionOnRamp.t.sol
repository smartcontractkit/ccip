// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMSubscriptionOnRampSetup.t.sol";

/// @notice #constructor
contract EVM2EVMSubscriptionOnRamp_constructor is EVM2EVMSubscriptionOnRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMSubscriptionOnRamp 1.0.0", s_onRamp.typeAndVersion());

    // owner
    assertEq(OWNER, s_onRamp.owner());

    // baseOnRamp
    assertEq(COMMIT_FEE_JUELS, s_onRamp.getConfig().commitFeeJuels);
    assertEq(MAX_DATA_SIZE, s_onRamp.getConfig().maxDataSize);
    assertEq(MAX_TOKENS_LENGTH, s_onRamp.getConfig().maxTokensLength);
    assertEq(MAX_GAS_LIMIT, s_onRamp.getConfig().maxGasLimit);

    assertEq(SOURCE_CHAIN_ID, s_onRamp.i_chainId());
    assertEq(DEST_CHAIN_ID, s_onRamp.i_destinationChainId());

    assertEq(address(s_onRampRouter), s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());

    // HealthChecker
    assertEq(address(s_afn), address(s_onRamp.getAFN()));
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMSubscriptionOnRamp_forwardFromRouter is EVM2EVMSubscriptionOnRampSetup {
  using CCIP for CCIP.EVMExtraArgsV1;

  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_onRampRouter));
  }

  // Success

  // Asserts that forwardFromRouter succeeds when called from the
  // router.
  function testSuccess() public {
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  // Asserts that multiple forwardFromRouter calls should result in
  // incrementing sequence number values.
  function testShouldIncrementSeqNumSuccess() public {
    uint64 seqNum = s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
    assertEq(seqNum, 1);
    seqNum = s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
    assertEq(seqNum, 2);
    seqNum = s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
    assertEq(seqNum, 3);
  }

  // Asserts that forwardFromRouter emits the correct event when sending
  // properly approved tokens.
  function testExactApproveSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].amount = 2**64;
    message.tokensAndAmounts[0].token = s_sourceTokens[0];

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(this._messageToEvent(message, 1, 1));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Assert that sending messages increments the nonce and the sequence numbers
  // on the onramp. Sending to a different receiver should start at 1 again.
  function testShouldIncrementReceiverNonceSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    CCIP.EVM2EVMSubscriptionMessage memory tollEvent = this._messageToEvent(message, 1, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = _generateEmptyMessage();
    tollEvent = this._messageToEvent(message, 2, 2);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = _generateEmptyMessage();
    message.receiver = abi.encode(address(s_onRampRouter));
    tollEvent = this._messageToEvent(message, 3, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Reverts

  function testPausedReverts() public {
    changePrank(OWNER);
    s_onRamp.pause();
    vm.expectRevert("Pausable: paused");
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  function testSenderNotAllowedReverts() public {
    changePrank(OWNER);
    s_onRamp.setAllowlistEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(AllowListInterface.SenderNotAllowed.selector, STRANGER));
    changePrank(address(s_onRampRouter));
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    address wrongToken = address(1);

    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].token = wrongToken;
    message.tokensAndAmounts[0].amount = 1;

    uint256[] memory prices = new uint256[](1);
    prices[0] = 1;
    // We need to set the price of this new token to be able to reach
    // the proper revert point. This must be called by the owner.
    changePrank(OWNER);
    s_onRamp.setPrices(abi.decode(abi.encode(message.tokensAndAmounts), (IERC20[])), prices);

    // Change back to the router
    changePrank(address(s_onRampRouter));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, wrongToken));
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts that forwardFromRouter reverts when it's not called by
  // the router
  function testMustBeCalledByRouterReverts() public {
    vm.stopPrank();
    vm.expectRevert(BaseOnRampInterface.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  // Asserts that forwardFromRouter reverts when the original sender
  // is not set by the router.
  function testRouterMustSetOriginalSenderReverts() public {
    vm.expectRevert(BaseOnRampInterface.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), address(0));
  }

  // Asserts that forwardFromRouter reverts when the number of supplied tokensAndAmounts
  // is larger than the maxTokenLength.
  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](MAX_TOKENS_LENGTH + 1);
    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts that forwardFromRouter reverts when the data length is too long.
  function testMessageTooLargeReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.data = new bytes(MAX_DATA_SIZE + 1);
    vm.expectRevert(
      abi.encodeWithSelector(
        BaseOnRampInterface.MessageTooLarge.selector,
        onRampConfig().maxDataSize,
        message.data.length
      )
    );

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testValueExceedsCapacityReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].amount = 2**128;
    message.tokensAndAmounts[0].token = s_sourceTokens[0];

    IERC20(s_sourceTokens[0]).approve(address(s_onRamp), 2**128);

    vm.expectRevert(
      abi.encodeWithSelector(
        AggregateRateLimiterInterface.ValueExceedsCapacity.selector,
        rateLimiterConfig().capacity,
        message.tokensAndAmounts[0].amount * getTokenPrices()[0]
      )
    );

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testPriceNotFoundForTokenReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    address fakeToken = address(1);
    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].token = fakeToken;

    vm.expectRevert(abi.encodeWithSelector(AggregateRateLimiterInterface.PriceNotFoundForToken.selector, fakeToken));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts gasLimit must be <=20M
  function testMessageGasLimitTooHighReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.extraArgs = CCIP.EVMExtraArgsV1({gasLimit: MAX_GAS_LIMIT + 1})._toBytes();
    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.MessageGasLimitTooHigh.selector));
    s_onRamp.forwardFromRouter(message, OWNER);
  }
}
