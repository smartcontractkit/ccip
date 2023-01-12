// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IEVM2AnyGEOnRamp} from "../../interfaces/onRamp/IEVM2AnyGEOnRamp.sol";
import {IBaseOnRampRouter} from "../../interfaces/onRamp/IBaseOnRampRouter.sol";
import {IGERouter} from "../../interfaces/router/IGERouter.sol";

import {PoolCollector} from "../../pools/PoolCollector.sol";
import "../onRamp/ge/EVM2EVMGEOnRampSetup.t.sol";

/// @notice #constructor
contract GERouter_constructor is EVM2EVMGEOnRampSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("GERouter 1.0.0", s_sourceRouter.typeAndVersion());

    // owner
    assertEq(OWNER, s_sourceRouter.owner());
  }
}

/// @notice #ccipSend
contract GERouter_ccipSend is EVM2EVMGEOnRampSetup {
  event Burned(address indexed sender, uint256 amount);

  // Success

  function testCCIPSendOneTokenSuccess_gas() public {
    vm.pauseGasMetering();
    address sourceToken1Address = s_sourceTokens[1];
    IERC20 sourceToken1 = IERC20(sourceToken1Address);
    GEConsumer.EVM2AnyGEMessage memory message = _generateEmptyMessage();

    sourceToken1.approve(address(s_sourceRouter), 2**64);

    message.tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].amount = 2**64;
    message.tokensAndAmounts[0].token = sourceToken1Address;
    message.feeToken = s_sourceTokens[0];

    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);

    uint256 balanceBefore = sourceToken1.balanceOf(OWNER);

    // Assert that the tokens are burned
    vm.expectEmit(false, false, false, true);
    emit Burned(address(s_onRamp), message.tokensAndAmounts[0].amount);

    GE.EVM2EVMGEMessage memory msgEvent = _messageToEvent(message, 1, 1, expectedFee);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(msgEvent);

    vm.resumeGasMetering();
    bytes32 messageId = s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
    vm.pauseGasMetering();

    assertEq(msgEvent.messageId, messageId);
    // Assert the user balance is lowered by the tokensAndAmounts sent and the fee amount
    uint256 expectedBalance = balanceBefore - (message.tokensAndAmounts[0].amount);
    assertEq(expectedBalance, sourceToken1.balanceOf(OWNER));
    vm.resumeGasMetering();
  }

  function testNonLinkFeeTokenSuccess() public {
    GE.FeeUpdate[] memory feeUpdates = new GE.FeeUpdate[](1);
    feeUpdates[0] = GE.FeeUpdate({token: s_sourceTokens[1], chainId: DEST_CHAIN_ID, linkPerUnitGas: 1000});
    s_IFeeManager.updateFees(feeUpdates);

    GEConsumer.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceTokens[1];

    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2**64);

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  // Reverts

  function testUnsupportedDestinationChainReverts() public {
    GEConsumer.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    uint64 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(IBaseOnRampRouter.UnsupportedDestinationChain.selector, wrongChain));

    s_sourceRouter.ccipSend(wrongChain, message);
  }

  function testUnsupportedFeeTokenReverts() public {
    GEConsumer.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    address wrongFeeToken = address(1);
    message.feeToken = wrongFeeToken;

    vm.expectRevert(
      abi.encodeWithSelector(IFeeManager.TokenOrChainNotSupported.selector, wrongFeeToken, DEST_CHAIN_ID)
    );

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    GEConsumer.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 0);

    vm.expectRevert("ERC20: transfer amount exceeds allowance");

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }
}

/// @notice #setOnRamp
contract GERouter_setOnRamp is EVM2EVMGEOnRampSetup {
  event OnRampSet(uint64 indexed chainId, IEVM2AnyGEOnRamp indexed onRamp);

  // Success

  // Asserts that setOnRamp changes the configured onramp. Also tests getOnRamp
  // and isChainSupported.
  function testSuccess() public {
    IEVM2AnyGEOnRamp onramp = IEVM2AnyGEOnRamp(address(1));
    uint64 chainId = 1337;
    IEVM2AnyGEOnRamp before = s_sourceRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_sourceRouter.isChainSupported(chainId));

    vm.expectEmit(true, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_sourceRouter.setOnRamp(chainId, onramp);
    IEVM2AnyGEOnRamp afterSet = s_sourceRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_sourceRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IGERouter.OnRampAlreadySet.selector, DEST_CHAIN_ID, s_onRamp));
    s_sourceRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);
  }

  // Asserts that setOnRamp can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.setOnRamp(1337, IEVM2AnyGEOnRamp(address(1)));
  }
}

/// @notice #isChainSupported
contract GERouter_isChainSupported is EVM2EVMGEOnRampSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));
    assertFalse(s_sourceRouter.isChainSupported(DEST_CHAIN_ID + 1));
    assertFalse(s_sourceRouter.isChainSupported(0));
  }
}
