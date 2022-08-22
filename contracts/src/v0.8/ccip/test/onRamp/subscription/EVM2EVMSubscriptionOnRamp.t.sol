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
    assertEq(RELAYING_FEE_JUELS, s_onRamp.getConfig().relayingFeeJuels);
    assertEq(MAX_DATA_SIZE, s_onRamp.getConfig().maxDataSize);
    assertEq(MAX_TOKENS_LENGTH, s_onRamp.getConfig().maxTokensLength);

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
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**128;
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = s_sourceTokens[0];
    message.amounts = amounts;
    message.tokens = tokens;

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Assert that sending messages increments the nonce and the sequence numbers
  // on the onramp. Sending to a different receiver should start at 1 again.
  function testShouldIncrementReceiverNonceSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    CCIP.EVM2EVMSubscriptionEvent memory tollEvent = _messageToEvent(message, 1, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = _generateEmptyMessage();
    tollEvent = _messageToEvent(message, 2, 2);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = _generateEmptyMessage();
    message.receiver = address(s_onRampRouter);
    tollEvent = _messageToEvent(message, 3, 1);

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
    IERC20 wrongToken = IERC20(address(1));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, wrongToken));
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
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
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  // Asserts that forwardFromRouter reverts when the original sender
  // is not set by the router.
  function testRouterMustSetOriginalSenderReverts() public {
    vm.expectRevert(BaseOnRampInterface.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), address(0));
  }

  // Asserts that forwardFromRouter reverts when the number of supplied tokens
  // is larger than the maxTokenLength.
  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokens = new IERC20[](MAX_TOKENS_LENGTH + 1);
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
}
