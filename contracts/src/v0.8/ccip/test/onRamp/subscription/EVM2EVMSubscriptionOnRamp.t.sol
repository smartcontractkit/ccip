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
    assertEq(s_onRampConfig.relayingFeeJuels, s_onRamp.getConfig().relayingFeeJuels);
    assertEq(s_onRampConfig.maxDataSize, s_onRamp.getConfig().maxDataSize);
    assertEq(s_onRampConfig.maxTokensLength, s_onRamp.getConfig().maxTokensLength);

    assertEq(SOURCE_CHAIN_ID, s_onRamp.CHAIN_ID());
    assertEq(DEST_CHAIN_ID, s_onRamp.DESTINATION_CHAIN_ID());

    assertEq(address(s_onRampRouter), s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());

    // TODO: AFN and Heartbeat timestamp getters
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMSubscriptionOnRamp_forwardFromRouter is EVM2EVMSubscriptionOnRampSetup {
  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    vm.stopPrank();
    vm.startPrank(address(s_onRampRouter));
  }

  // Success

  // Asserts that forwardFromRouter succeeds when called from the
  // router.
  function testSuccess() public {
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  // Asserts that multiple forwardFromRouter calls should result in
  // incrementing sequence number values.
  function testShouldIncrementSeqNumSuccess() public {
    uint64 seqNum = s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
    assertEq(seqNum, 1);
    seqNum = s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
    assertEq(seqNum, 2);
    seqNum = s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
    assertEq(seqNum, 3);
  }

  // Asserts that forwardFromRouter emits the correct event when sending
  // properly approved tokens.
  function testExactApproveSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = getEmptyMessage();
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**128;
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = s_sourceTokens[0];
    message.amounts = amounts;
    message.tokens = tokens;

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(messageToEvent(message, 1, 1));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Assert that sending messages increments the nonce and the sequence numbers
  // on the onramp. Sending to a different receiver should start at 1 again.
  function testShouldIncrementReceiverNonceSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = getEmptyMessage();
    CCIP.EVM2EVMSubscriptionEvent memory tollEvent = messageToEvent(message, 1, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = getEmptyMessage();
    tollEvent = messageToEvent(message, 2, 2);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = getEmptyMessage();
    message.receiver = address(s_onRampRouter);
    tollEvent = messageToEvent(message, 3, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Reverts

  function testPausedReverts() public {
    vm.stopPrank();
    vm.startPrank(OWNER);
    s_onRamp.pause();
    vm.expectRevert("Pausable: paused");
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  function testSenderNotAllowedReverts() public {
    vm.stopPrank();
    vm.prank(OWNER);
    s_onRamp.setAllowlistEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(AllowListInterface.SenderNotAllowed.selector, STRANGER));
    vm.prank(address(s_onRampRouter));
    s_onRamp.forwardFromRouter(getEmptyMessage(), STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    IERC20 wrongToken = IERC20(address(1));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, wrongToken));
    CCIP.EVM2AnySubscriptionMessage memory message = getEmptyMessage();
    message.tokens = new IERC20[](1);
    message.tokens[0] = wrongToken;
    message.amounts = new uint256[](1);

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts that forwardFromRouter reverts when it's not called by
  // the router
  function testMustBeCalledByRouterReverts() public {
    vm.stopPrank();
    vm.expectRevert(BaseOnRampInterface.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  // Asserts that forwardFromRouter reverts when the original sender
  // is not set by the router.
  function testRouterMustSetOriginalSenderReverts() public {
    vm.expectRevert(BaseOnRampInterface.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), address(0));
  }

  // Asserts that forwardFromRouter reverts when the number of supplied tokens
  // is larger than the maxTokenLength.
  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = getEmptyMessage();
    message.tokens = new IERC20[](s_onRampConfig.maxTokensLength + 1);
    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts that forwardFromRouter reverts when the data length is too long.
  function testMessageTooLargeReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = getEmptyMessage();
    message.data = "000000000000000000000000000000000000000000000000000";
    vm.expectRevert(
      abi.encodeWithSelector(
        BaseOnRampInterface.MessageTooLarge.selector,
        s_onRampConfig.maxDataSize,
        message.data.length
      )
    );

    s_onRamp.forwardFromRouter(message, OWNER);
  }
}
