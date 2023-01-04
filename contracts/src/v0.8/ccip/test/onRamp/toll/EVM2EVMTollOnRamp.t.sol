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
    IBaseOnRamp.OnRampConfig memory onRampConfig = onRampConfig();
    assertEq(onRampConfig.commitFeeJuels, s_onRamp.getConfig().commitFeeJuels);
    assertEq(onRampConfig.maxDataSize, s_onRamp.getConfig().maxDataSize);
    assertEq(onRampConfig.maxTokensLength, s_onRamp.getConfig().maxTokensLength);
    assertEq(onRampConfig.maxGasLimit, s_onRamp.getConfig().maxGasLimit);

    assertEq(SOURCE_CHAIN_ID, s_onRamp.i_chainId());
    assertEq(DEST_CHAIN_ID, s_onRamp.i_destinationChainId());

    assertEq(address(s_onRampRouter), s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());

    // HealthChecker
    assertEq(address(s_afn), address(s_onRamp.getAFN()));
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMTollOnRamp_forwardFromRouter is EVM2EVMTollOnRampSetup {
  using TollConsumer for TollConsumer.EVMExtraArgsV1;

  function setUp() public virtual override {
    EVM2EVMTollOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_onRampRouter));
  }

  // Success

  function testSuccess() public {
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEventNoFee(message, 1));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testShouldIncrementSeqNumSuccess() public {
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEventNoFee(message, 1));

    s_onRamp.forwardFromRouter(message, OWNER);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEventNoFee(message, 2));

    s_onRamp.forwardFromRouter(message, OWNER);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEventNoFee(message, 3));

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

  function testPermissionsReverts() public {
    changePrank(OWNER);
    vm.expectRevert(IBaseOnRamp.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  function testOriginalSenderReverts() public {
    vm.expectRevert(IBaseOnRamp.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), address(0));
  }

  function testMessageTooLargeReverts() public {
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    message.data = new bytes(onRampConfig().maxDataSize + 1);
    vm.expectRevert(
      abi.encodeWithSelector(IBaseOnRamp.MessageTooLarge.selector, onRampConfig().maxDataSize, message.data.length)
    );

    s_onRamp.forwardFromRouter(message, STRANGER);
  }

  function testTooManyTokensReverts() public {
    assertEq(MAX_TOKENS_LENGTH, s_onRamp.getConfig().maxTokensLength);
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    uint256 tooMany = MAX_TOKENS_LENGTH + 1;
    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](tooMany);
    vm.expectRevert(IBaseOnRamp.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, STRANGER);
  }

  function testSenderNotAllowedReverts() public {
    changePrank(OWNER);
    s_onRamp.setAllowlistEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(IAllowList.SenderNotAllowed.selector, STRANGER));
    changePrank(address(s_onRampRouter));
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    address wrongToken = address(1);

    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();
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
    changePrank(address(s_onRampRouter));
    vm.expectRevert(abi.encodeWithSelector(IBaseOnRamp.UnsupportedToken.selector, wrongToken));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testValueExceedsCapacityReverts() public {
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();
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

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testPriceNotFoundForTokenReverts() public {
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();

    address fakeToken = address(1);
    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].token = fakeToken;

    vm.expectRevert(abi.encodeWithSelector(IAggregateRateLimiter.PriceNotFoundForToken.selector, fakeToken));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts gasLimit must be <=maxGasLimit
  function testMessageGasLimitTooHighReverts() public {
    TollConsumer.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    message.extraArgs = TollConsumer._argsToBytes(
      TollConsumer.EVMExtraArgsV1({gasLimit: MAX_GAS_LIMIT + 1, strict: false})
    );
    vm.expectRevert(abi.encodeWithSelector(IBaseOnRamp.MessageGasLimitTooHigh.selector));
    s_onRamp.forwardFromRouter(message, OWNER);
  }
}

/// @notice #getRequiredFee
contract EVM2EVMTollOnRamp_getRequiredFee is EVM2EVMTollOnRampSetup {
  // Success

  // Asserts that the fee is calculated correctly.
  function testGetRequiredFeeSuccess() public {
    uint256 fee = s_onRamp.getRequiredFee(IERC20(s_sourceTokens[0]));
    uint256 expectedFee = COMMIT_FEE_JUELS;
    assertEq(expectedFee, fee);
  }
}

/// @notice #setFeeConfig
contract EVM2EVMTollOnRamp_setFeeConfig is EVM2EVMTollOnRampSetup {
  IEVM2EVMTollOnRamp.FeeConfig s_feeConfig1;
  IEVM2EVMTollOnRamp.FeeConfig s_feeConfig2;
  uint256 constant FEE = 1;

  function setUp() public virtual override {
    EVM2EVMTollOnRampSetup.setUp();

    // Set up arguments for fee config.
    IERC20[] memory feeTokens1 = new IERC20[](1);
    feeTokens1[0] = IERC20(s_sourceTokens[0]);
    IERC20[] memory feeTokens2 = new IERC20[](1);
    feeTokens2[0] = IERC20(s_sourceTokens[1]);
    uint256[] memory fees = new uint256[](1);
    fees[0] = FEE;
    s_feeConfig1 = IEVM2EVMTollOnRamp.FeeConfig({feeTokens: feeTokens1, fees: fees});
    s_feeConfig2 = IEVM2EVMTollOnRamp.FeeConfig({feeTokens: feeTokens2, fees: fees});
  }

  // Success

  function testSetFeeConfigSuccess() public {
    s_onRamp.setFeeConfig(s_feeConfig1);
    // Only configured fee should be set.
    uint256 fee = s_onRamp.getRequiredFee(IERC20(s_sourceTokens[0]));
    assertEq(FEE, fee);
    uint256 fee2 = s_onRamp.getRequiredFee(IERC20(s_sourceTokens[1]));
    assertEq(0, fee2);

    // Should clear old fees upon setting.
    s_onRamp.setFeeConfig(s_feeConfig2);
    fee = s_onRamp.getRequiredFee(IERC20(s_sourceTokens[0]));
    assertEq(0, fee);
    fee2 = s_onRamp.getRequiredFee(IERC20(s_sourceTokens[1]));
    assertEq(FEE, fee2);
  }

  // Reverts

  function testSetFeeConfigNotOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_onRamp.setFeeConfig(s_feeConfig1);
  }
}
