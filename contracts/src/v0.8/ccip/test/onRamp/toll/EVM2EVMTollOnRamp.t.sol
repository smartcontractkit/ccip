// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMTollOnRampSetup.t.sol";

/// @notice #constructor
contract EVM2EVMTollOnRamp_constructor is EVM2EVMTollOnRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMTollOnRamp 1.0.0", s_onRamp.typeAndVersion());

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

    // HealthChecker
    assertEq(HEARTBEAT, s_onRamp.getMaxSecondsWithoutAFNHeartbeat());
    assertEq(address(s_afn), address(s_onRamp.getAFN()));
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMTollOnRamp_forwardFromRouter is EVM2EVMTollOnRampSetup {
  function setUp() public virtual override {
    EVM2EVMTollOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_onRampRouter));
  }

  // Success

  function testSuccess() public {
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  // Reverts

  function testPausedReverts() public {
    changePrank(OWNER);
    s_onRamp.pause();
    vm.expectRevert("Pausable: paused");
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  function testPermissionsReverts() public {
    changePrank(OWNER);
    vm.expectRevert(BaseOnRampInterface.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), OWNER);
  }

  function testOriginalSenderReverts() public {
    vm.expectRevert(BaseOnRampInterface.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), address(0));
  }

  function testMessageTooLargeReverts() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.data = "000000000000000000000000000000000000000000000000000";
    vm.expectRevert(
      abi.encodeWithSelector(
        BaseOnRampInterface.MessageTooLarge.selector,
        s_onRampConfig.maxDataSize,
        message.data.length
      )
    );

    s_onRamp.forwardFromRouter(message, STRANGER);
  }

  function testTooManyTokensReverts() public {
    assertEq(3, s_onRamp.getConfig().maxTokensLength);
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    uint256 tooMany = 4;
    message.tokens = new IERC20[](tooMany);
    message.amounts = new uint256[](tooMany);
    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, STRANGER);
  }

  function testTokenNumberMismatchReverts() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.tokens = new IERC20[](1);
    message.amounts = new uint256[](2);
    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, STRANGER);
  }

  function testSenderNotAllowedReverts() public {
    changePrank(OWNER);
    s_onRamp.setAllowlistEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(AllowListInterface.SenderNotAllowed.selector, STRANGER));
    changePrank(address(s_onRampRouter));
    s_onRamp.forwardFromRouter(getEmptyMessage(), STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    IERC20 wrongToken = IERC20(address(1));

    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.tokens = new IERC20[](1);
    message.tokens[0] = wrongToken;
    message.amounts = new uint256[](1);

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, wrongToken));

    s_onRamp.forwardFromRouter(message, OWNER);
  }
}

/// @notice #getRequiredFee
contract EVM2EVMTollOnRamp_getRequiredFee is EVM2EVMTollOnRampSetup {
  // Success

  // Asserts that the fee is calculated correctly.
  function testGetRequiredFeeSuccess() public {
    uint256 fee = s_onRamp.getRequiredFee(s_sourceTokens[0]);
    uint256 expectedFee = s_onRampConfig.relayingFeeJuels * uint256(s_sourceFeeds[0].latestAnswer());
    assertEq(expectedFee, fee);
  }

  // Reverts

  // Asserts that the fee is calculated correctly.
  function testGetRequiredFeeUnsupportedFeeToken() public {
    IERC20 wrongToken = IERC20(address(1));

    vm.expectRevert(abi.encodeWithSelector(Any2EVMTollOnRampInterface.UnsupportedFeeToken.selector, wrongToken));

    s_onRamp.getRequiredFee(wrongToken);
  }
}
