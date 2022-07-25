// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMTollOnRampSetup.t.sol";

/// @notice #constructor
contract EVM2AnyTollOnRampRouter_constructor is EVM2EVMTollOnRampSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2AnyTollOnRampRouter 1.0.0", s_onRampRouter.typeAndVersion());

    // owner
    assertEq(OWNER, s_onRampRouter.owner());
  }
}

/// @notice #ccipSend
contract EVM2AnyTollOnRampRouter_ccipSend is EVM2EVMTollOnRampSetup {
  // Success

  function testSuccess() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.amounts = new uint256[](1);
    message.amounts[0] = 2**64;
    message.tokens = new IERC20[](1);
    message.tokens[0] = s_sourceTokens[0];

    uint256 balanceBefore = s_sourceTokens[0].balanceOf(OWNER);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(messageToEvent(message, 1));

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
    // Assert the user balance is lowered by the tokens sent and the fee amount
    uint256 expectedBalance = balanceBefore - (message.amounts[0] + FEE_AMOUNT);
    assertEq(expectedBalance, s_sourceTokens[0].balanceOf(OWNER));
    // Asserts the tokens are sent to the pool
    assertEq(message.amounts[0], s_sourceTokens[0].balanceOf(address(s_sourcePools[0])));
    // Asserts the fee amount is left in the router
    assertEq(FEE_AMOUNT, s_sourceTokens[0].balanceOf(address(s_onRampRouter)));
  }

  function testExactApproveSuccess() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.amounts = new uint256[](1);
    // since the fee token is the same we should reduce the amount sent
    // when we want an exact approve.
    message.amounts[0] = 2**128 - FEE_AMOUNT;
    message.tokens = new IERC20[](1);
    message.tokens[0] = s_sourceTokens[0];

    uint256 balanceBefore = s_sourceTokens[0].balanceOf(OWNER);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(messageToEvent(message, 1));

    uint256 expectedBalance = balanceBefore - (message.amounts[0] + FEE_AMOUNT);

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
    assertEq(expectedBalance, s_sourceTokens[0].balanceOf(OWNER));
  }

  function testShouldIncrementSeqNumSuccess() public {
    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, getEmptyMessage()));
    assertEq(2, s_onRampRouter.ccipSend(DEST_CHAIN_ID, getEmptyMessage()));
    assertEq(3, s_onRampRouter.ccipSend(DEST_CHAIN_ID, getEmptyMessage()));
  }

  // Reverts

  function testUnsupportedDestinationChainReverts() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    uint256 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampRouterInterface.UnsupportedDestinationChain.selector, wrongChain));

    s_onRampRouter.ccipSend(wrongChain, message);
  }

  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.amounts = new uint256[](5);

    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testUnsupportedFeeTokenReverts() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    IERC20 wrongFeeToken = IERC20(address(1));
    message.feeToken = wrongFeeToken;

    vm.expectRevert(abi.encodeWithSelector(Any2EVMTollOnRampInterface.UnsupportedFeeToken.selector, wrongFeeToken));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.feeTokenAmount = 0;

    vm.expectRevert(PoolCollector.FeeTokenAmountTooLow.selector);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }
}

/// @notice #setOnRamp
contract EVM2AnyTollOnRampRouter_setOnRamp is EVM2EVMTollOnRampSetup {
  event OnRampSet(uint256 indexed chainId, Any2EVMTollOnRampInterface indexed onRamp);

  // Success

  // Asserts that setOnRamp changes the configured onramp. Also tests getOnRamp
  // and isChainSupported.
  function testSuccess() public {
    Any2EVMTollOnRampInterface onramp = Any2EVMTollOnRampInterface(address(1));
    uint256 chainId = 1337;
    Any2EVMTollOnRampInterface before = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_onRampRouter.isChainSupported(chainId));

    vm.expectEmit(true, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_onRampRouter.setOnRamp(chainId, onramp);
    Any2EVMTollOnRampInterface afterSet = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_onRampRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(Any2EVMTollOnRampRouterInterface.OnRampAlreadySet.selector, DEST_CHAIN_ID, s_onRamp)
    );
    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);
  }

  // Asserts that setOnRamp can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRampRouter.setOnRamp(1337, Any2EVMTollOnRampInterface(address(1)));
  }
}

/// @notice #isChainSupported
contract EVM2AnyTollOnRampRouter_isChainSupported is EVM2EVMTollOnRampSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_onRampRouter.isChainSupported(DEST_CHAIN_ID));
    assertFalse(s_onRampRouter.isChainSupported(DEST_CHAIN_ID + 1));
    assertFalse(s_onRampRouter.isChainSupported(0));
  }
}
