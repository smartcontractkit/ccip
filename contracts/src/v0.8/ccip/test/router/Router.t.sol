// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IEVM2AnyOnRamp} from "../../interfaces/onRamp/IEVM2AnyOnRamp.sol";
import {IRouter} from "../../interfaces/router/IRouter.sol";
import {IWrappedNative} from "../../interfaces/router/IWrappedNative.sol";
import {IRouterClient} from "../../interfaces/router/IRouterClient.sol";
import {IOwnable} from "../../interfaces/IOwnable.sol";
import {ITypeAndVersion} from "../../interfaces/ITypeAndVersion.sol";

import "../onRamp/EVM2EVMOnRampSetup.t.sol";

/// @notice #constructor
contract Router_constructor is EVM2EVMOnRampSetup {
  // Success

  function testSuccess() public {
    assertEq("Router 1.0.0", ITypeAndVersion(address(s_sourceRouter)).typeAndVersion());
    // owner
    assertEq(OWNER, IOwnable(address(s_sourceRouter)).owner());
  }
}

/// @notice #ccipSend
contract Router_ccipSend is EVM2EVMOnRampSetup {
  event Burned(address indexed sender, uint256 amount);

  // Success

  function testCCIPSendOneTokenSuccess_gas() public {
    vm.pauseGasMetering();
    address sourceToken1Address = s_sourceTokens[1];
    IERC20 sourceToken1 = IERC20(sourceToken1Address);
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();

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

    Internal.EVM2EVMMessage memory msgEvent = _messageToEvent(message, 1, 1, expectedFee);

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
    Internal.FeeUpdate[] memory feeUpdates = new Internal.FeeUpdate[](1);
    feeUpdates[0] = Internal.FeeUpdate({
      sourceFeeToken: s_sourceTokens[1],
      destChainId: DEST_CHAIN_ID,
      feeTokenBaseUnitsPerUnitGas: 1000
    });
    s_IFeeManager.updateFees(feeUpdates);
    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeTokenConfigArgs = new IEVM2EVMOnRamp.FeeTokenConfigArgs[](1);
    feeTokenConfigArgs[0] = IEVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceTokens[1],
      feeAmount: 2,
      multiplier: 108e16,
      destGasOverhead: 2
    });
    s_onRamp.setFeeConfig(feeTokenConfigArgs);

    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceTokens[1];
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2**64);
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenSuccess() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    uint256 nativeQuote = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    vm.stopPrank();
    hoax(address(1), 100 ether);
    s_sourceRouter.ccipSend{value: nativeQuote}(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenOverpaySuccess() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    uint256 nativeQuote = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    vm.stopPrank();
    hoax(address(1), 100 ether);
    s_sourceRouter.ccipSend{value: 1e18}(DEST_CHAIN_ID, message);
    // We expect the overpayment to be taken in full.
    assertEq(address(1).balance, 99 ether);
  }

  function testWrappedNativeFeeTokenSuccess() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceRouter.getWrappedNative();
    uint256 nativeQuote = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    vm.stopPrank();
    hoax(address(1), 100 ether);
    // Now address(1) has nativeQuote wrapped.
    IWrappedNative(s_sourceRouter.getWrappedNative()).deposit{value: nativeQuote}();
    IWrappedNative(s_sourceRouter.getWrappedNative()).approve(address(s_sourceRouter), nativeQuote);
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  // Reverts

  function testUnsupportedDestinationChainReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint64 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(IRouterClient.UnsupportedDestinationChain.selector, wrongChain));

    s_sourceRouter.ccipSend(wrongChain, message);
  }

  function testUnsupportedFeeTokenReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    address wrongFeeToken = address(1);
    message.feeToken = wrongFeeToken;

    vm.expectRevert(
      abi.encodeWithSelector(IFeeManager.TokenOrChainNotSupported.selector, wrongFeeToken, DEST_CHAIN_ID)
    );

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 0);

    vm.expectRevert("ERC20: transfer amount exceeds allowance");

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testInvalidMsgValue() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    // Non-empty feeToken but with msg.value should revert
    vm.stopPrank();
    hoax(address(1), 1);
    vm.expectRevert(IRouterClient.InvalidMsgValue.selector);
    s_sourceRouter.ccipSend{value: 1}(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenZeroValue() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    // Include no value, should revert
    vm.expectRevert();
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenInsufficientValue() public {
    Consumer.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    // Include insufficient, should also revert
    vm.stopPrank();
    hoax(address(1), 1);
    vm.expectRevert(IRouterClient.InsufficientFeeTokenAmount.selector);
    s_sourceRouter.ccipSend{value: 1}(DEST_CHAIN_ID, message);
  }
}

/// @notice #setOnRamp
contract Router_setOnRamp is EVM2EVMOnRampSetup {
  event OnRampSet(uint64 indexed chainId, IEVM2AnyOnRamp indexed onRamp);

  // Success

  // Asserts that setOnRamp changes the configured onramp. Also tests getOnRamp
  // and isChainSupported.
  function testSuccess() public {
    IEVM2AnyOnRamp onramp = IEVM2AnyOnRamp(address(1));
    uint64 chainId = 1337;
    IEVM2AnyOnRamp before = s_sourceRouter.getOnRamp(chainId);
    assertEq(address(0), address(before));
    assertFalse(s_sourceRouter.isChainSupported(chainId));

    vm.expectEmit(true, true, false, true);
    emit OnRampSet(chainId, onramp);

    s_sourceRouter.setOnRamp(chainId, onramp);
    IEVM2AnyOnRamp afterSet = s_sourceRouter.getOnRamp(chainId);
    assertEq(address(onramp), address(afterSet));
    assertTrue(s_sourceRouter.isChainSupported(chainId));
  }

  // Reverts

  // Asserts that setOnRamp reverts when the config was already set to
  // the same onRamp.
  function testAlreadySetReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IRouter.OnRampAlreadySet.selector, DEST_CHAIN_ID, s_onRamp));
    s_sourceRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);
  }

  // Asserts that setOnRamp can only be called by the owner.
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.setOnRamp(1337, IEVM2AnyOnRamp(address(1)));
  }
}

/// @notice #setWrappedNative
contract Router_setWrappedNative is EVM2EVMOnRampSetup {
  // Success
  function testSuccess() public {
    s_sourceRouter.setWrappedNative(address(1));
  }

  // Reverts
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.setWrappedNative(address(1));
  }
}

/// @notice #isChainSupported
contract Router_isChainSupported is EVM2EVMOnRampSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));
    assertFalse(s_sourceRouter.isChainSupported(DEST_CHAIN_ID + 1));
    assertFalse(s_sourceRouter.isChainSupported(0));
  }
}

/// @notice #getSupportedTokens
contract Router_getSupportedTokens is EVM2EVMOnRampSetup {
  // Success

  function testGetSupportedTokensSuccess() public {
    assertEq(s_sourceTokens, s_sourceRouter.getSupportedTokens(DEST_CHAIN_ID));
  }

  function testUnknownChainSuccess() public {
    address[] memory supportedTokens = s_sourceRouter.getSupportedTokens(DEST_CHAIN_ID + 10);
    assertEq(0, supportedTokens.length);
  }
}

/// @notice #addOffRamp
contract Router_addOffRamp is EVM2EVMOnRampSetup {
  address internal s_newOffRamp;

  event OffRampAdded(address indexed offRamp);

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_newOffRamp = address(1);
  }

  // Success

  function testSuccess() public {
    assertFalse(s_sourceRouter.isOffRamp(s_newOffRamp));
    uint256 lengthBefore = s_sourceRouter.getOffRamps().length;

    vm.expectEmit(true, false, false, true);
    emit OffRampAdded(s_newOffRamp);
    s_sourceRouter.addOffRamp(s_newOffRamp);

    assertTrue(s_sourceRouter.isOffRamp(s_newOffRamp));
    assertEq(lengthBefore + 1, s_sourceRouter.getOffRamps().length);
  }

  // Reverts

  function testOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.addOffRamp(s_newOffRamp);
  }

  function testAlreadyConfiguredReverts() public {
    address existingOffRamp = s_offRamps[0];
    vm.expectRevert(abi.encodeWithSelector(IRouter.AlreadyConfigured.selector, existingOffRamp));
    s_sourceRouter.addOffRamp(existingOffRamp);
  }

  function testZeroAddressReverts() public {
    vm.expectRevert(IRouter.InvalidAddress.selector);
    s_sourceRouter.addOffRamp(address(0));
  }
}

/// @notice #removeOffRamp
contract Router_removeOffRamp is EVM2EVMOnRampSetup {
  event OffRampRemoved(address indexed offRamp);

  // Success

  function testSuccess() public {
    uint256 lengthBefore = s_sourceRouter.getOffRamps().length;

    vm.expectEmit(true, false, false, true);
    emit OffRampRemoved(s_offRamps[0]);
    s_sourceRouter.removeOffRamp(s_offRamps[0]);

    assertFalse(s_sourceRouter.isOffRamp(s_offRamps[0]));
    assertEq(lengthBefore - 1, s_sourceRouter.getOffRamps().length);
  }

  // Reverts

  function testOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.removeOffRamp(s_offRamps[0]);
  }

  function testNoOffRampsReverts() public {
    s_sourceRouter.removeOffRamp(s_offRamps[0]);
    s_sourceRouter.removeOffRamp(s_offRamps[1]);

    assertEq(0, s_sourceRouter.getOffRamps().length);

    vm.expectRevert(IRouter.NoOffRampsConfigured.selector);
    s_sourceRouter.removeOffRamp(s_offRamps[0]);
  }

  function testOffRampNotAllowedReverts() public {
    address newRamp = address(1234678);
    vm.expectRevert(abi.encodeWithSelector(IRouter.OffRampNotAllowed.selector, newRamp));
    s_sourceRouter.removeOffRamp(newRamp);
  }
}

/// @notice #getOffRamps
contract Router_getOffRamps is EVM2EVMOnRampSetup {
  // Success
  function testGetOffRampsSuccess() public {
    address[] memory offRamps = s_sourceRouter.getOffRamps();
    assertEq(2, offRamps.length);
    assertEq(address(s_offRamps[0]), address(offRamps[0]));
    assertEq(address(s_offRamps[1]), address(offRamps[1]));
  }
}

/// @notice #isOffRamp
contract Router_isOffRamp is EVM2EVMOnRampSetup {
  // Success
  function testIsOffRampSuccess() public {
    assertTrue(s_sourceRouter.isOffRamp(s_offRamps[0]));
    assertTrue(s_sourceRouter.isOffRamp(s_offRamps[1]));
    assertFalse(s_sourceRouter.isOffRamp(address(1)));
  }
}
