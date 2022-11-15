// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMSubscriptionOnRampSetup.t.sol";
import "../../../interfaces/onRamp/BaseOnRampRouterInterface.sol";

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
    emit CCIPSendRequested(this._messageToEvent(message, 1, 1));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testTokensSuccess() public {
    IERC20 sourceToken0 = IERC20(s_sourceTokens[0]);
    IERC20 sourceToken1 = IERC20(s_sourceTokens[1]);
    uint256 poolBalance0Before = sourceToken0.balanceOf(address(s_sourcePools[0]));
    uint256 poolBalance1Before = sourceToken1.balanceOf(address(s_sourcePools[1]));

    uint256 userBalance0Before = sourceToken0.balanceOf(OWNER);
    uint256 userBalance1Before = sourceToken1.balanceOf(OWNER);

    sourceToken0.approve(address(s_onRampRouter), i_tokenAmount0);
    sourceToken1.approve(address(s_onRampRouter), i_tokenAmount1);
    CCIP.EVM2AnySubscriptionMessage memory message = _generateTokenMessage();

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(this._messageToEvent(message, 1, 1));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
    // Assert the user balance is lowered by the tokensAndAmounts sent
    assertEq(userBalance0Before - message.tokensAndAmounts[0].amount, sourceToken0.balanceOf(OWNER));
    assertEq(userBalance1Before - message.tokensAndAmounts[1].amount, sourceToken1.balanceOf(OWNER));
    // Asserts the tokens are all sent to the proper pools
    assertEq(poolBalance0Before + i_tokenAmount0, sourceToken0.balanceOf(address(s_sourcePools[0])));
    // BurnMintTokenPool burns the token upon lock, so poolBalance should not change
    assertEq(poolBalance1Before, sourceToken1.balanceOf(address(s_sourcePools[1])));
  }

  function testChargeSubscriptionFundingSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    uint96 newFee = 100;
    s_onRampRouter.setFee(newFee);
    s_onRampRouter.fundSubscription(newFee);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(this._messageToEvent(message, 1, 1));

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
    EVM2EVMSubscriptionOnRampInterface onramp = EVM2EVMSubscriptionOnRampInterface(address(1));
    uint256 chainId = 1337;
    EVM2EVMSubscriptionOnRampInterface before = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_onRampRouter.isChainSupported(chainId));

    vm.expectEmit(false, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_onRampRouter.setOnRamp(chainId, onramp);
    EVM2EVMSubscriptionOnRampInterface afterSet = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_onRampRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        EVM2AnySubscriptionOnRampRouterInterface.OnRampAlreadySet.selector,
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
    s_onRampRouter.setOnRamp(1337, EVM2EVMSubscriptionOnRampInterface(address(1)));
  }
}

/// @notice #removeOnRamp
contract EVM2AnySubscriptionOnRampRouter_removeOnRamp is EVM2EVMSubscriptionOnRampSetup {
  // Success

  function testSuccess() public {
    EVM2EVMSubscriptionOnRampInterface onramp = EVM2EVMSubscriptionOnRampInterface(address(1));
    uint256 chainId = 1337;
    s_onRampRouter.setOnRamp(chainId, onramp);

    vm.expectEmit(false, true, false, true);
    emit OnRampRemoved(chainId, onramp);

    s_onRampRouter.removeOnRamp(chainId, onramp);

    EVM2EVMSubscriptionOnRampInterface afterSet = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(0), address(afterSet));
    assertEq(false, s_onRampRouter.isChainSupported(chainId));
  }

  // Reverts

  function testWrongAddressReverts() public {
    EVM2EVMSubscriptionOnRampInterface onramp = EVM2EVMSubscriptionOnRampInterface(address(1));
    uint256 chainId = 1337;
    s_onRampRouter.setOnRamp(chainId, onramp);

    address wrongRamp = address(10000);

    vm.expectRevert(
      abi.encodeWithSelector(EVM2AnySubscriptionOnRampRouterInterface.WrongOnRamp.selector, wrongRamp, address(onramp))
    );

    s_onRampRouter.removeOnRamp(chainId, EVM2EVMSubscriptionOnRampInterface(wrongRamp));
  }

  // Asserts that it can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRampRouter.removeOnRamp(1337, EVM2EVMSubscriptionOnRampInterface(address(1)));
  }
}

/// @notice #setFee
contract EVM2AnySubscriptionOnRampRouter_setFee is EVM2EVMSubscriptionOnRampSetup {
  // Success

  // Asserts that setFee sets the commit fee.
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
    vm.expectRevert(EVM2AnySubscriptionOnRampRouterInterface.OnlyCallableByFeeAdmin.selector);
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

    vm.expectEmit(false, false, false, true);
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
  uint256 internal immutable i_fundingAmount = 500;

  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();
    s_onRampRouter.fundSubscription(i_fundingAmount);
  }

  // Success

  function testSuccess() public {
    assertEq(i_fundingAmount, s_onRampRouter.getBalance(OWNER));

    vm.expectEmit(false, false, false, true);
    emit SubscriptionUnfunded(OWNER, i_fundingAmount);

    s_onRampRouter.unfundSubscription(i_fundingAmount);

    assertEq(0, s_onRampRouter.getBalance(OWNER));
  }

  // Reverts

  function testFundingTooLowReverts() public {
    assertEq(i_fundingAmount, s_onRampRouter.getBalance(OWNER));

    vm.expectRevert(stdError.arithmeticError);

    s_onRampRouter.unfundSubscription(i_fundingAmount * 2);

    assertEq(i_fundingAmount, s_onRampRouter.getBalance(OWNER));
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
