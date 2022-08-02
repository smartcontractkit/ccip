// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMSubscriptionOnRampSetup.t.sol";

/// @notice #constructor
contract EVM2AnySubscriptionOnRampRouter_constructor is EVM2EVMSubscriptionOnRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2AnySubscriptionOnRampRouter 1.0.0", s_onRampRouter.typeAndVersion());

    // owner
    assertEq(OWNER, s_onRampRouter.owner());
  }
}

/// @notice #ccipSend
contract EVM2AnySubscriptionOnRampRouter_ccipSend is EVM2EVMSubscriptionOnRampSetup {
  // Success

  // Asserts that ccipSend with proper arguments succeeds
  function testSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testTokensSuccess() public {
    uint256 poolBalance0Before = s_sourceTokens[0].balanceOf(address(s_sourcePools[0]));
    uint256 poolBalance1Before = s_sourceTokens[1].balanceOf(address(s_sourcePools[1]));

    uint256 userBalance0Before = s_sourceTokens[0].balanceOf(OWNER);
    uint256 userBalance1Before = s_sourceTokens[1].balanceOf(OWNER);

    s_sourceTokens[0].approve(address(s_onRampRouter), TOKEN_AMOUNT_0);
    s_sourceTokens[1].approve(address(s_onRampRouter), TOKEN_AMOUNT_1);
    CCIP.EVM2AnySubscriptionMessage memory message = _generateTokenMessage();

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
    // Assert the user balance is lowered by the tokens sent
    assertEq(userBalance0Before - message.amounts[0], s_sourceTokens[0].balanceOf(OWNER));
    assertEq(userBalance1Before - message.amounts[1], s_sourceTokens[1].balanceOf(OWNER));
    // Asserts the tokens are all sent to the proper pools
    assertEq(poolBalance0Before + TOKEN_AMOUNT_0, s_sourceTokens[0].balanceOf(address(s_sourcePools[0])));
    assertEq(poolBalance1Before + TOKEN_AMOUNT_1, s_sourceTokens[1].balanceOf(address(s_sourcePools[1])));
  }

  function testChargeSubscriptionFundingSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    uint96 newFee = 100;
    s_onRampRouter.setFee(newFee);
    s_onRampRouter.fundSubscription(newFee);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  // Reverts

  function testTokensNoApproveReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateTokenMessage();

    vm.expectRevert("ERC20: transfer amount exceeds allowance");

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  // Asserts that ccipSend can only be called for a supported destination chain.
  function testUnsupportedDestinationChainReverts() public {
    uint256 wrongChain = DEST_CHAIN_ID + 1;
    vm.expectRevert(abi.encodeWithSelector(BaseOnRampRouterInterface.UnsupportedDestinationChain.selector, wrongChain));
    s_onRampRouter.ccipSend(wrongChain, _generateEmptyMessage());
  }

  function testChargeSubscriptionFundingTooLowReverts() public {
    uint96 newFee = 100;
    s_onRampRouter.setFee(newFee);

    vm.expectRevert(stdError.arithmeticError);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage());
  }
}

/// @notice #setOnRamp
contract EVM2AnySubscriptionOnRampRouter_setOnRamp is EVM2EVMSubscriptionOnRampSetup {
  // Success

  // Asserts that setOnRamp changes the configured onramp. Also tests getOnRamp
  // and isChainSupported.
  function testSuccess() public {
    Any2EVMSubscriptionOnRampInterface onramp = Any2EVMSubscriptionOnRampInterface(address(1));
    uint256 chainId = 1337;
    Any2EVMSubscriptionOnRampInterface before = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_onRampRouter.isChainSupported(chainId));

    vm.expectEmit(true, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_onRampRouter.setOnRamp(chainId, onramp);
    Any2EVMSubscriptionOnRampInterface afterSet = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_onRampRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        Any2EVMSubscriptionOnRampRouterInterface.OnRampAlreadySet.selector,
        DEST_CHAIN_ID,
        s_onRamp
      )
    );

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);
  }

  // Asserts that setOnRamp can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRampRouter.setOnRamp(1337, Any2EVMSubscriptionOnRampInterface(address(1)));
  }
}

/// @notice #setFee
contract EVM2AnySubscriptionOnRampRouter_setFee is EVM2EVMSubscriptionOnRampSetup {
  // Success

  // Asserts that setFee sets the relay fee.
  function testSuccess() public {
    uint96 newFee = 100000;

    vm.expectEmit(false, false, false, true);
    emit FeeSet(newFee);

    s_onRampRouter.setFee(newFee);

    assertEq(newFee, s_onRampRouter.getFee());
  }

  // Reverts

  // Asserts that setFee can only be called by the feeAdmin.
  function testOnlyFeeAdminReverts() public {
    vm.stopPrank();
    vm.expectRevert(Any2EVMSubscriptionOnRampRouterInterface.OnlyCallableByFeeAdmin.selector);
    s_onRampRouter.setFee(1);
  }
}

/// @notice #fundSubscription
contract EVM2AnySubscriptionOnRampRouter_fundSubscription is EVM2EVMSubscriptionOnRampSetup {
  // Success

  // Assert that fundSubscription properly increases the balance of
  // a subscription.
  function testSuccess() public {
    uint256 fundingAmount = 550;
    assertEq(0, s_onRampRouter.getBalance(OWNER));

    vm.expectEmit(true, false, false, true);
    emit SubscriptionFunded(OWNER, fundingAmount);

    s_onRampRouter.fundSubscription(fundingAmount);

    assertEq(fundingAmount, s_onRampRouter.getBalance(OWNER));
  }

  // Reverts

  function testApproveTooLowReverts() public {
    assertEq(0, s_onRampRouter.getBalance(OWNER));
    vm.expectRevert("ERC20: transfer amount exceeds allowance");
    s_onRampRouter.fundSubscription(2**256 - 1);
    assertEq(0, s_onRampRouter.getBalance(OWNER));
  }

  function testFundsTooLowReverts() public {
    address mockAddress = address(9);
    assertEq(0, s_onRampRouter.getBalance(mockAddress));
    vm.expectRevert("ERC20: transfer amount exceeds balance");
    changePrank(mockAddress);
    s_onRampRouter.fundSubscription(2**256 - 1);
    assertEq(0, s_onRampRouter.getBalance(mockAddress));
  }
}

/// @notice #unfundSubscription
contract EVM2AnySubscriptionOnRampRouter_unfundSubscription is EVM2EVMSubscriptionOnRampSetup {
  uint256 immutable FUNDING_AMOUNT = 500;

  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();
    s_onRampRouter.fundSubscription(FUNDING_AMOUNT);
  }

  // Success

  function testSuccess() public {
    assertEq(FUNDING_AMOUNT, s_onRampRouter.getBalance(OWNER));

    vm.expectEmit(true, false, false, true);
    emit SubscriptionUnfunded(OWNER, FUNDING_AMOUNT);

    s_onRampRouter.unfundSubscription(FUNDING_AMOUNT);

    assertEq(0, s_onRampRouter.getBalance(OWNER));
  }

  // Reverts

  function testFundingTooLowReverts() public {
    assertEq(FUNDING_AMOUNT, s_onRampRouter.getBalance(OWNER));

    vm.expectRevert(stdError.arithmeticError);

    s_onRampRouter.unfundSubscription(FUNDING_AMOUNT * 2);

    assertEq(FUNDING_AMOUNT, s_onRampRouter.getBalance(OWNER));
  }
}

/// @notice #isChainSupported
contract EVM2AnySubscriptionOnRampRouter_isChainSupported is EVM2EVMSubscriptionOnRampSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_onRampRouter.isChainSupported(DEST_CHAIN_ID));
    assertFalse(s_onRampRouter.isChainSupported(DEST_CHAIN_ID + 1));
    assertFalse(s_onRampRouter.isChainSupported(0));
  }
}
