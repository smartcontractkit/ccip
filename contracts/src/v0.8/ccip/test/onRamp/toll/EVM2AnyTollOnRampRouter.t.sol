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
    address sourceToken0Address = s_sourceTokens[0];
    IERC20 sourceToken0 = IERC20(sourceToken0Address);
    CCIP.EVM2AnyTollMessage memory message = _generateEmptyMessage();

    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].amount = 2**64;
    message.tokensAndAmounts[0].token = sourceToken0Address;
    message.feeTokenAndAmount.token = sourceToken0Address;
    message.feeTokenAndAmount.amount = RELAYING_FEE_JUELS;

    uint256 balanceBefore = sourceToken0.balanceOf(OWNER);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1));

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
    // Assert the user balance is lowered by the tokensAndAmounts sent and the fee amount
    uint256 expectedBalance = balanceBefore - (message.tokensAndAmounts[0].amount + RELAYING_FEE_JUELS);
    assertEq(expectedBalance, sourceToken0.balanceOf(OWNER));
    // Asserts the tokensAndAmounts are sent to the pool
    assertEq(message.tokensAndAmounts[0].amount, sourceToken0.balanceOf(address(s_sourcePools[0])));
    // Asserts the fee amount is left in the router
    assertEq(RELAYING_FEE_JUELS, sourceToken0.balanceOf(address(s_onRampRouter)));
  }

  function testExactApproveSuccess() public {
    address sourceToken0Address = s_sourceTokens[0];
    IERC20 sourceToken0 = IERC20(sourceToken0Address);
    CCIP.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    // since the fee token is the same we should reduce the amount sent
    // when we want an exact approve.
    message.tokensAndAmounts[0].amount = 2**64 - RELAYING_FEE_JUELS;
    message.tokensAndAmounts[0].token = sourceToken0Address;

    uint256 balanceBefore = sourceToken0.balanceOf(OWNER);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1));

    uint256 expectedBalance = balanceBefore - (message.tokensAndAmounts[0].amount + RELAYING_FEE_JUELS);

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
    assertEq(expectedBalance, sourceToken0.balanceOf(OWNER));
  }

  function testShouldIncrementSeqNumSuccess() public {
    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage()));
    assertEq(2, s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage()));
    assertEq(3, s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage()));
  }

  // Reverts

  function testUnsupportedDestinationChainReverts() public {
    CCIP.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    uint256 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampRouterInterface.UnsupportedDestinationChain.selector, wrongChain));

    s_onRampRouter.ccipSend(wrongChain, message);
  }

  function testUnsupportedFeeTokenReverts() public {
    CCIP.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    address wrongFeeToken = address(1);
    message.feeTokenAndAmount = CCIP.EVMTokenAndAmount({token: wrongFeeToken, amount: 0});

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, wrongFeeToken));

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    CCIP.EVM2AnyTollMessage memory message = _generateEmptyMessage();
    message.feeTokenAndAmount.amount = 0;

    vm.expectRevert(PoolCollector.FeeTokenAmountTooLow.selector);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }
}

/// @notice #setOnRamp
contract EVM2AnyTollOnRampRouter_setOnRamp is EVM2EVMTollOnRampSetup {
  event OnRampSet(uint256 indexed chainId, EVM2EVMTollOnRampInterface indexed onRamp);

  // Success

  // Asserts that setOnRamp changes the configured onramp. Also tests getOnRamp
  // and isChainSupported.
  function testSuccess() public {
    EVM2EVMTollOnRampInterface onramp = EVM2EVMTollOnRampInterface(address(1));
    uint256 chainId = 1337;
    EVM2EVMTollOnRampInterface before = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_onRampRouter.isChainSupported(chainId));

    vm.expectEmit(true, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_onRampRouter.setOnRamp(chainId, onramp);
    EVM2EVMTollOnRampInterface afterSet = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_onRampRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(EVM2AnyTollOnRampRouterInterface.OnRampAlreadySet.selector, DEST_CHAIN_ID, s_onRamp)
    );
    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);
  }

  // Asserts that setOnRamp can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRampRouter.setOnRamp(1337, EVM2EVMTollOnRampInterface(address(1)));
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
