// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {PoolCollector} from "../../pools/PoolCollector.sol";
import {BaseOnRampRouterInterface} from "../../interfaces/onRamp/BaseOnRampRouterInterface.sol";
import {GERouterInterface} from "../../interfaces/router/GERouterInterface.sol";
import "../onRamp/ge/EVM2EVMGEOnRampSetup.t.sol";

/// @notice #constructor
contract GERouter_constructor is EVM2EVMGEOnRampSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("GERouter 1.0.0", s_onRampRouter.typeAndVersion());

    // owner
    assertEq(OWNER, s_onRampRouter.owner());
  }
}

/// @notice #ccipSend
contract GERouter_ccipSend is EVM2EVMGEOnRampSetup {
  event Burned(address indexed sender, uint256 amount);

  // Success

  function testCCIPSendSuccess() public {
    address sourceToken1Address = s_sourceTokens[1];
    IERC20 sourceToken1 = IERC20(sourceToken1Address);
    CCIP.EVM2AnyGEMessage memory message = _generateEmptyMessage();

    sourceToken1.approve(address(s_onRampRouter), 2**64);

    message.tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    message.tokensAndAmounts[0].amount = 2**64;
    message.tokensAndAmounts[0].token = sourceToken1Address;
    message.feeToken = s_sourceTokens[0];

    uint256 expectedFee = s_onRampRouter.getFee(DEST_CHAIN_ID, message);

    uint256 balanceBefore = sourceToken1.balanceOf(OWNER);

    // Assert that the tokens are burned
    vm.expectEmit(false, false, false, true);
    emit Burned(address(s_onRamp), message.tokensAndAmounts[0].amount);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1, expectedFee));

    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, message));
    // Assert the user balance is lowered by the tokensAndAmounts sent and the fee amount
    uint256 expectedBalance = balanceBefore - (message.tokensAndAmounts[0].amount);
    assertEq(expectedBalance, sourceToken1.balanceOf(OWNER));
  }

  function testShouldIncrementSeqNumSuccess() public {
    assertEq(1, s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage()));
    assertEq(2, s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage()));
    assertEq(3, s_onRampRouter.ccipSend(DEST_CHAIN_ID, _generateEmptyMessage()));
  }

  // Reverts

  function testUnsupportedDestinationChainReverts() public {
    CCIP.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    uint256 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampRouterInterface.UnsupportedDestinationChain.selector, wrongChain));

    s_onRampRouter.ccipSend(wrongChain, message);
  }

  function testUnsupportedFeeTokenReverts() public {
    CCIP.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    address wrongFeeToken = address(1);
    message.feeToken = wrongFeeToken;

    vm.expectRevert(
      abi.encodeWithSelector(
        DynamicFeeCalculatorInterface.MismatchedFeeToken.selector,
        s_sourceTokens[0],
        wrongFeeToken
      )
    );

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    CCIP.EVM2AnyGEMessage memory message = _generateEmptyMessage();
    IERC20(s_sourceTokens[0]).approve(address(s_onRampRouter), 0);

    vm.expectRevert("ERC20: transfer amount exceeds allowance");

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);
  }
}

/// @notice #setOnRamp
contract GERouter_setOnRamp is EVM2EVMGEOnRampSetup {
  event OnRampSet(uint256 indexed chainId, EVM2EVMGEOnRampInterface indexed onRamp);

  // Success

  // Asserts that setOnRamp changes the configured onramp. Also tests getOnRamp
  // and isChainSupported.
  function testSuccess() public {
    EVM2EVMGEOnRampInterface onramp = EVM2EVMGEOnRampInterface(address(1));
    uint256 chainId = 1337;
    EVM2EVMGEOnRampInterface before = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_onRampRouter.isChainSupported(chainId));

    vm.expectEmit(true, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_onRampRouter.setOnRamp(chainId, onramp);
    EVM2EVMGEOnRampInterface afterSet = s_onRampRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_onRampRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(abi.encodeWithSelector(GERouterInterface.OnRampAlreadySet.selector, DEST_CHAIN_ID, s_onRamp));
    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);
  }

  // Asserts that setOnRamp can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRampRouter.setOnRamp(1337, EVM2EVMGEOnRampInterface(address(1)));
  }
}

/// @notice #isChainSupported
contract GERouter_isChainSupported is EVM2EVMGEOnRampSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_onRampRouter.isChainSupported(DEST_CHAIN_ID));
    assertFalse(s_onRampRouter.isChainSupported(DEST_CHAIN_ID + 1));
    assertFalse(s_onRampRouter.isChainSupported(0));
  }
}
