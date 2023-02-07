// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMOnRampSetup.t.sol";

/// @notice #constructor
contract EVM2EVMOnRamp_constructor is EVM2EVMOnRampSetup {
  function testConstructorSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMOnRamp 1.0.0", s_onRamp.typeAndVersion());

    // owner
    assertEq(OWNER, s_onRamp.owner());

    // baseOnRamp
    IBaseOnRamp.OnRampConfig memory onRampConfig = onRampConfig();
    assertEq(onRampConfig.commitFeeJuels, s_onRamp.getOnRampConfig().commitFeeJuels);
    assertEq(onRampConfig.maxDataSize, s_onRamp.getOnRampConfig().maxDataSize);
    assertEq(onRampConfig.maxTokensLength, s_onRamp.getOnRampConfig().maxTokensLength);
    assertEq(onRampConfig.maxGasLimit, s_onRamp.getOnRampConfig().maxGasLimit);

    assertEq(SOURCE_CHAIN_ID, s_onRamp.getChainId());
    assertEq(DEST_CHAIN_ID, s_onRamp.getDestinationChainId());

    assertEq(address(s_sourceRouter), s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());

    assertEq(s_sourceTokens, s_onRamp.getSupportedTokens());

    // HealthChecker
    assertEq(address(s_afn), address(s_onRamp.getAFN()));
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMOnRamp_forwardFromRouter is EVM2EVMOnRampSetup {
  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_sourceRouter));
  }

  // Success

  function testForwardFromRouterSuccess() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1, feeAmount));

    s_onRamp.forwardFromRouter(message, feeAmount, OWNER);
  }

  function testShouldIncrementSeqNumAndNonceSuccess() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();

    for (uint64 i = 1; i < 4; i++) {
      uint64 nonceBefore = s_onRamp.getSenderNonce(OWNER);

      vm.expectEmit(false, false, false, true);
      emit CCIPSendRequested(_messageToEvent(message, i, i, 0));

      s_onRamp.forwardFromRouter(message, 0, OWNER);

      uint64 nonceAfter = s_onRamp.getSenderNonce(OWNER);
      assertEq(nonceAfter, nonceBefore + 1);
    }
  }

  event Transfer(address indexed from, address indexed to, uint256 value);

  function testShouldSendFeesToTheFeeManager() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRamp), feeAmount);

    vm.expectEmit(true, true, true, true);
    emit Transfer(address(s_onRamp), address(s_IFeeManager), feeAmount);

    s_onRamp.forwardFromRouter(message, feeAmount, OWNER);
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
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, OWNER);
  }

  function testPermissionsReverts() public {
    changePrank(OWNER);
    vm.expectRevert(IBaseOnRamp.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, OWNER);
  }

  function testOriginalSenderReverts() public {
    vm.expectRevert(IBaseOnRamp.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, address(0));
  }

  function testMessageTooLargeReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.data = new bytes(onRampConfig().maxDataSize + 1);
    vm.expectRevert(
      abi.encodeWithSelector(IBaseOnRamp.MessageTooLarge.selector, onRampConfig().maxDataSize, message.data.length)
    );

    s_onRamp.forwardFromRouter(message, 0, STRANGER);
  }

  function testTooManyTokensReverts() public {
    assertEq(MAX_TOKENS_LENGTH, s_onRamp.getOnRampConfig().maxTokensLength);
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 tooMany = MAX_TOKENS_LENGTH + 1;
    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](tooMany);
    vm.expectRevert(IBaseOnRamp.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, 0, STRANGER);
  }

  function testSenderNotAllowedReverts() public {
    changePrank(OWNER);
    s_onRamp.setAllowlistEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(IAllowList.SenderNotAllowed.selector, STRANGER));
    changePrank(address(s_sourceRouter));
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), 0, STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    address wrongToken = address(1);

    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].token = wrongToken;
    message.tokensAndAmounts[0].amount = 1;

    // We need to set the price of this new token to be able to reach
    // the proper revert point. This must be called by the owner.
    changePrank(OWNER);
    uint256[] memory prices = new uint256[](1);
    prices[0] = 1;
    s_onRamp.setPrices(abi.decode(abi.encode(message.tokensAndAmounts), (IERC20[])), prices);

    // Change back to the router
    changePrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(IBaseOnRamp.UnsupportedToken.selector, wrongToken));

    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  function testValueExceedsCapacityReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].amount = 2**128;
    message.tokensAndAmounts[0].token = s_sourceTokens[0];

    IERC20(s_sourceTokens[0]).approve(address(s_onRamp), 2**128);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAggregateRateLimiter.ValueExceedsCapacity.selector,
        rateLimiterConfig().capacity,
        message.tokensAndAmounts[0].amount * getTokenPrices()[0]
      )
    );

    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  function testPriceNotFoundForTokenReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();

    address fakeToken = address(1);
    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].token = fakeToken;

    vm.expectRevert(abi.encodeWithSelector(IAggregateRateLimiter.PriceNotFoundForToken.selector, fakeToken));

    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }

  // Asserts gasLimit must be <=maxGasLimit
  function testMessageGasLimitTooHighReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = Consumer._argsToBytes(Consumer.EVMExtraArgsV1({gasLimit: MAX_GAS_LIMIT + 1, strict: false}));
    vm.expectRevert(abi.encodeWithSelector(IBaseOnRamp.MessageGasLimitTooHigh.selector));
    s_onRamp.forwardFromRouter(message, 0, OWNER);
  }
}

/// @notice #setFeeAdmin
contract EVM2EVMOnRamp_setFeeAdmin is EVM2EVMOnRampSetup {
  event FeeAdminSet(address feeAdmin);

  function testOwnerSetFeeAdminSuccess() public {
    address newAdmin = address(13371337);

    vm.expectEmit(false, false, false, true);
    emit FeeAdminSet(newAdmin);

    s_onRamp.setFeeAdmin(newAdmin);
  }

  // Reverts

  function testOnlyCallableByOwnerReverts() public {
    address newAdmin = address(13371337);
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");

    s_onRamp.setFeeAdmin(newAdmin);
  }
}

/// @notice #setFeeConfig
contract EVM2EVMOnRamp_setFeeConfig is EVM2EVMOnRampSetup {
  event FeeConfigSet(IEVM2EVMOnRamp.FeeTokenConfigArgs[] feeConfig);

  function testSetFeeConfigSuccess() public {
    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeConfig;

    vm.expectEmit(false, false, false, true);
    emit FeeConfigSet(feeConfig);

    s_onRamp.setFeeConfig(feeConfig);
  }

  function testSetFeeConfigByFeeAdminSuccess() public {
    address newAdmin = address(13371337);
    s_onRamp.setFeeAdmin(newAdmin);

    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeConfig;

    changePrank(newAdmin);

    vm.expectEmit(false, false, false, true);
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
