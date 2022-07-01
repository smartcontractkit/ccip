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

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(messageToEvent(message, 1));

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
  }

  function testExactApproveSuccess() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.amounts = new uint256[](1);
    // since the fee token is the same we should reduce the amount sent
    // when we want an exact approve.
    message.amounts[0] = 2**128 - FEE_AMOUNT;
    message.tokens = new IERC20[](1);
    message.tokens[0] = s_sourceTokens[0];

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
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
