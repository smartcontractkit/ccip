// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IEVM2AnyOnRamp} from "../../interfaces/IEVM2AnyOnRamp.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";
import {IRouterClient} from "../../interfaces/IRouterClient.sol";

import "../onRamp/EVM2EVMOnRampSetup.t.sol";
import "../helpers/receivers/SimpleMessageReceiver.sol";
import "../offRamp/EVM2EVMOffRampSetup.t.sol";

/// @notice #constructor
contract Router_constructor is EVM2EVMOnRampSetup {
  function testConstructorSuccess() public {
    assertEq("Router 1.0.0", s_sourceRouter.typeAndVersion());
    // owner
    assertEq(OWNER, s_sourceRouter.owner());
  }
}

/// @notice #ccipSend
contract Router_ccipSend is EVM2EVMOnRampSetup {
  event Burned(address indexed sender, uint256 amount);

  function testCCIPSendLinkFeeOneTokenSuccess_gas() public {
    vm.pauseGasMetering();
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    IERC20 sourceToken1 = IERC20(s_sourceTokens[1]);
    sourceToken1.approve(address(s_sourceRouter), 2**64);

    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].amount = 2**64;
    message.tokenAmounts[0].token = s_sourceTokens[1];

    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertGt(expectedFee, 0);

    uint256 balanceBefore = sourceToken1.balanceOf(OWNER);

    // Assert that the tokens are burned
    vm.expectEmit();
    emit Burned(address(s_onRamp), message.tokenAmounts[0].amount);

    Internal.EVM2EVMMessage memory msgEvent = _messageToEvent(message, 1, 1, expectedFee, OWNER);

    vm.expectEmit();
    emit CCIPSendRequested(msgEvent);

    vm.resumeGasMetering();
    bytes32 messageId = s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
    vm.pauseGasMetering();

    assertEq(msgEvent.messageId, messageId);
    // Assert the user balance is lowered by the tokenAmounts sent and the fee amount
    uint256 expectedBalance = balanceBefore - (message.tokenAmounts[0].amount);
    assertEq(expectedBalance, sourceToken1.balanceOf(OWNER));
    vm.resumeGasMetering();
  }

  function testCCIPSendLinkFeeNoTokenSuccess_gas() public {
    vm.pauseGasMetering();
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertGt(expectedFee, 0);

    Internal.EVM2EVMMessage memory msgEvent = _messageToEvent(message, 1, 1, expectedFee, OWNER);

    vm.expectEmit();
    emit CCIPSendRequested(msgEvent);

    vm.resumeGasMetering();
    bytes32 messageId = s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
    vm.pauseGasMetering();

    assertEq(msgEvent.messageId, messageId);
    vm.resumeGasMetering();
  }

  function testCCIPSendNativeFeeOneTokenSuccess_gas() public {
    vm.pauseGasMetering();
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    IERC20 sourceToken1 = IERC20(s_sourceTokens[1]);
    sourceToken1.approve(address(s_sourceRouter), 2**64);

    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].amount = 2**64;
    message.tokenAmounts[0].token = s_sourceTokens[1];
    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertGt(expectedFee, 0);

    uint256 balanceBefore = sourceToken1.balanceOf(OWNER);

    // Assert that the tokens are burned
    vm.expectEmit();
    emit Burned(address(s_onRamp), message.tokenAmounts[0].amount);

    // Native fees will be wrapped so we need to calculate the event with
    // the wrapped native feeCoin address.
    message.feeToken = s_sourceRouter.getWrappedNative();
    Internal.EVM2EVMMessage memory msgEvent = _messageToEvent(message, 1, 1, expectedFee, OWNER);
    // Set it to address(0) to indicate native
    message.feeToken = address(0);

    vm.expectEmit();
    emit CCIPSendRequested(msgEvent);

    vm.resumeGasMetering();
    bytes32 messageId = s_sourceRouter.ccipSend{value: expectedFee}(DEST_CHAIN_ID, message);
    vm.pauseGasMetering();

    assertEq(msgEvent.messageId, messageId);
    // Assert the user balance is lowered by the tokenAmounts sent and the fee amount
    uint256 expectedBalance = balanceBefore - (message.tokenAmounts[0].amount);
    assertEq(expectedBalance, sourceToken1.balanceOf(OWNER));
    vm.resumeGasMetering();
  }

  function testCCIPSendNativeFeeNoTokenSuccess_gas() public {
    vm.pauseGasMetering();
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertGt(expectedFee, 0);

    // Native fees will be wrapped so we need to calculate the event with
    // the wrapped native feeCoin address.
    message.feeToken = s_sourceRouter.getWrappedNative();
    Internal.EVM2EVMMessage memory msgEvent = _messageToEvent(message, 1, 1, expectedFee, OWNER);
    // Set it to address(0) to indicate native
    message.feeToken = address(0);

    vm.expectEmit();
    emit CCIPSendRequested(msgEvent);

    vm.resumeGasMetering();
    bytes32 messageId = s_sourceRouter.ccipSend{value: expectedFee}(DEST_CHAIN_ID, message);
    vm.pauseGasMetering();

    assertEq(msgEvent.messageId, messageId);
    // Assert the user balance is lowered by the tokenAmounts sent and the fee amount
    vm.resumeGasMetering();
  }

  function testNonLinkFeeTokenSuccess() public {
    EVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeTokenConfigArgs = new EVM2EVMOnRamp.FeeTokenConfigArgs[](1);
    feeTokenConfigArgs[0] = EVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceTokens[1],
      feeAmount: 2,
      multiplier: 108e16,
      destGasOverhead: 2
    });
    s_onRamp.setFeeConfig(feeTokenConfigArgs);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceTokens[1];
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceTokens[1];
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2**64);
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenSuccess() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    uint256 nativeQuote = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    vm.stopPrank();
    hoax(address(1), 100 ether);
    s_sourceRouter.ccipSend{value: nativeQuote}(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenOverpaySuccess() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    uint256 nativeQuote = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    vm.stopPrank();
    hoax(address(1), 100 ether);
    s_sourceRouter.ccipSend{value: nativeQuote + 1}(DEST_CHAIN_ID, message);
    // We expect the overpayment to be taken in full.
    assertEq(address(1).balance, 100 ether - (nativeQuote + 1));
    assertEq(address(s_sourceRouter).balance, 0);
  }

  function testWrappedNativeFeeTokenSuccess() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceRouter.getWrappedNative();
    uint256 nativeQuote = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    vm.stopPrank();
    hoax(address(1), 100 ether);
    // Now address(1) has nativeQuote wrapped.
    IWrappedNative(s_sourceRouter.getWrappedNative()).deposit{value: nativeQuote}();
    IWrappedNative(s_sourceRouter.getWrappedNative()).approve(address(s_sourceRouter), nativeQuote);
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testZeroFeeAndGasPriceSuccess() public {
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceTokens[1];
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));

    Internal.PriceUpdates memory priceUpdates = getSinglePriceUpdateStruct(s_sourceTokens[1], 2_000 ether);
    priceUpdates.destChainId = DEST_CHAIN_ID;
    priceUpdates.usdPerUnitGas = 0;

    s_priceRegistry.updatePrices(priceUpdates);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceTokens[1];
    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertEq(fee, 0);
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  // Reverts

  function testUnsupportedDestinationChainReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint64 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(IRouterClient.UnsupportedDestinationChain.selector, wrongChain));

    s_sourceRouter.ccipSend(wrongChain, message);
  }

  function testUnsupportedFeeTokenReverts(address wrongFeeToken) public {
    // We have three fee tokens set, all others should revert.
    vm.assume(address(s_sourceFeeToken) != wrongFeeToken);
    vm.assume(address(s_sourceRouter.getWrappedNative()) != wrongFeeToken);
    vm.assume(address(0) != wrongFeeToken);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = wrongFeeToken;

    vm.expectRevert(abi.encodeWithSelector(PriceRegistry.NotAFeeToken.selector, wrongFeeToken));

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 0);

    vm.expectRevert("ERC20: transfer amount exceeds allowance");

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testInvalidMsgValue() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    // Non-empty feeToken but with msg.value should revert
    vm.stopPrank();
    hoax(address(1), 1);
    vm.expectRevert(IRouterClient.InvalidMsgValue.selector);
    s_sourceRouter.ccipSend{value: 1}(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenZeroValue() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    // Include no value, should revert
    vm.expectRevert();
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testNativeFeeTokenInsufficientValue() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = address(0); // Raw native
    // Include insufficient, should also revert
    vm.stopPrank();

    s_onRamp.getFeeConfig(s_sourceRouter.getWrappedNative());

    hoax(address(1), 1);
    vm.expectRevert(IRouterClient.InsufficientFeeTokenAmount.selector);
    s_sourceRouter.ccipSend{value: 1}(DEST_CHAIN_ID, message);
  }
}

contract Router_applyRampUpdates is RouterSetup {
  SimpleMessageReceiver s_receiver;

  function setUp() public virtual override(RouterSetup) {
    RouterSetup.setUp();
    s_receiver = new SimpleMessageReceiver();
  }

  function generateSingleOffRampUpdate(uint64 sourceChainId, address offRamp)
    private
    pure
    returns (Router.OffRampUpdate memory)
  {
    address[] memory offRamps = new address[](1);
    offRamps[0] = offRamp;
    return Router.OffRampUpdate({sourceChainId: sourceChainId, offRamps: offRamps});
  }

  function generateDisableOffRampUpdate(uint64 sourceChainId) private pure returns (Router.OffRampUpdate memory) {
    address[] memory offRamps = new address[](0);
    return Router.OffRampUpdate({sourceChainId: sourceChainId, offRamps: offRamps});
  }

  function testRampUpdates(uint16 config) public {
    // Setup
    uint16 nApply = config & ((1 << 8) - 1);
    uint16 nOnRampUpdates = (config >> 8) & ((1 << 2) - 1);
    uint16 nOffRampUpdates = (config >> 10) & ((1 << 2) - 1);
    uint16 nOffRamps = (config >> 12) & ((1 << 2) - 1);
    uint16 disableOnRamp = (config >> 14) & ((1 << 2) - 1);
    vm.assume(nOffRampUpdates <= 2);
    vm.assume(nOnRampUpdates <= 2);
    vm.assume(nApply <= 2);
    vm.assume(nOffRamps <= 2);
    vm.assume(disableOnRamp <= 1);

    // Apply series of updates.
    for (uint256 j = 0; j < uint256(nApply); j++) {
      Router.OnRampUpdate[] memory onRampUpdates = new Router.OnRampUpdate[](nOnRampUpdates);
      Router.OffRampUpdate[] memory offRampUpdates = new Router.OffRampUpdate[](nOffRampUpdates);
      // For each application we use the same chainID range to ensure overwriting.
      for (uint256 i = 1; i < uint256(nOnRampUpdates) + 1; ++i) {
        if (uint256(disableOnRamp) == i - 1) {
          // Disable this chainID
          onRampUpdates[i - 1] = Router.OnRampUpdate({destChainId: uint64(i), onRamp: address(uint160(0))});
        } else {
          onRampUpdates[i - 1] = Router.OnRampUpdate({destChainId: uint64(i), onRamp: address(uint160(i))});
        }
      }
      for (uint256 i = 1; i < uint256(nOffRampUpdates) + 1; ++i) {
        address[] memory offRamps = new address[](nOffRamps);
        // If nOffRamps = 0, we are disabling this chainID.
        for (uint256 k = 0; k < uint256(nOffRamps); k++) {
          offRamps[k] = address(uint160(i));
        }
        offRampUpdates[i - 1] = Router.OffRampUpdate({sourceChainId: uint64(i), offRamps: offRamps});
      }
      s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);

      // Assert invariants
      for (uint256 i = 1; i < nOnRampUpdates + 1; ++i) {
        if (disableOnRamp == i - 1) {
          assertFalse(s_sourceRouter.isChainSupported(uint64(i)));
          assertEq(address(uint160(0)), s_sourceRouter.getOnRamp(uint64(i)));
        } else {
          assertTrue(s_sourceRouter.isChainSupported(uint64(i)));
          assertEq(address(uint160(i)), s_sourceRouter.getOnRamp(uint64(i)));
        }
      }
      for (uint256 i = 1; i < nOffRampUpdates + 1; ++i) {
        // Should be able to call route message from configured offramps.
        // This ensures the second map is properly maintained.
        vm.stopPrank();
        for (uint256 k = 0; k < nOffRamps; k++) {
          vm.prank(offRampUpdates[i - 1].offRamps[k]);
          s_sourceRouter.routeMessage(generateReceiverMessage(uint64(i)), false, 100_000, address(s_receiver));
        }
        vm.startPrank(OWNER);
        // Will be true even if we disable the ingress (empty list)
        assertEq(offRampUpdates[i - 1].offRamps, s_sourceRouter.getOffRamps(uint64(i)));
      }
      // Randomize number updates between applies
      nOnRampUpdates++;
      nOffRampUpdates++;
      disableOnRamp++;
      if (nOffRamps > 0) {
        nOffRamps--;
      }
    }
  }

  function testOffRampDisable() public {
    // Add ingress
    Router.OnRampUpdate[] memory onRampUpdates = new Router.OnRampUpdate[](0);
    Router.OffRampUpdate[] memory offRampUpdates = new Router.OffRampUpdate[](1);
    address offRamp = address(uint160(2));
    offRampUpdates[0] = generateSingleOffRampUpdate(SOURCE_CHAIN_ID, offRamp);
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
    assertEq(1, s_sourceRouter.getOffRamps(SOURCE_CHAIN_ID).length);
    assertEq(offRamp, s_sourceRouter.getOffRamps(SOURCE_CHAIN_ID)[0]);
    // Remove ingress
    offRampUpdates[0] = generateDisableOffRampUpdate(SOURCE_CHAIN_ID);
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
    assertEq(0, s_sourceRouter.getOffRamps(SOURCE_CHAIN_ID).length);

    // Disabled offramp should not be able to route.
    vm.expectRevert(IRouter.OnlyOffRamp.selector);
    changePrank(offRamp);
    s_sourceRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), false, 100_000, address(s_receiver));
    changePrank(OWNER);

    // Re-enabling should succeed
    offRampUpdates[0] = generateSingleOffRampUpdate(SOURCE_CHAIN_ID, offRamp);
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
    assertEq(1, s_sourceRouter.getOffRamps(SOURCE_CHAIN_ID).length);
    assertEq(offRamp, s_sourceRouter.getOffRamps(SOURCE_CHAIN_ID)[0]);
    changePrank(offRamp);
    s_sourceRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), false, 100_000, address(s_receiver));
  }

  function testOnRampDisable() public {
    // Add onRamp
    Router.OnRampUpdate[] memory onRampUpdates = new Router.OnRampUpdate[](1);
    Router.OffRampUpdate[] memory offRampUpdates = new Router.OffRampUpdate[](0);
    address onRamp = address(uint160(2));
    onRampUpdates[0] = Router.OnRampUpdate({destChainId: DEST_CHAIN_ID, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
    assertEq(onRamp, s_sourceRouter.getOnRamp(DEST_CHAIN_ID));
    assertTrue(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));

    // Disable onRamp
    onRampUpdates[0] = Router.OnRampUpdate({destChainId: DEST_CHAIN_ID, onRamp: address(0)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
    assertEq(address(0), s_sourceRouter.getOnRamp(DEST_CHAIN_ID));
    assertFalse(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));

    // Re-enable onRamp
    onRampUpdates[0] = Router.OnRampUpdate({destChainId: DEST_CHAIN_ID, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
    assertEq(onRamp, s_sourceRouter.getOnRamp(DEST_CHAIN_ID));
    assertTrue(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));
  }

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    Router.OnRampUpdate[] memory onRampUpdates = new Router.OnRampUpdate[](0);
    Router.OffRampUpdate[] memory offRampUpdates = new Router.OffRampUpdate[](0);
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);
  }
}

/// @notice #setWrappedNative
contract Router_setWrappedNative is EVM2EVMOnRampSetup {
  function testSetWrappedNativeSuccess(address wrappedNative) public {
    s_sourceRouter.setWrappedNative(wrappedNative);
    assertEq(wrappedNative, s_sourceRouter.getWrappedNative());
  }

  // Reverts
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.setWrappedNative(address(1));
  }
}

/// @notice #getSupportedTokens
contract Router_getSupportedTokens is EVM2EVMOnRampSetup {
  function testGetSupportedTokensSuccess() public {
    assertEq(s_sourceTokens, s_sourceRouter.getSupportedTokens(DEST_CHAIN_ID));
  }

  function testUnknownChainSuccess() public {
    address[] memory supportedTokens = s_sourceRouter.getSupportedTokens(DEST_CHAIN_ID + 10);
    assertEq(0, supportedTokens.length);
  }
}

/// @notice #routeMessage
contract Router_routeMessage is EVM2EVMOffRampSetup {
  function testManualExecSuccess() public {
    changePrank(address(s_offRamp));
    assertTrue(s_destRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), true, 100_000, address(s_receiver)));
    // Manuel execution cannot run out of gas
    assertTrue(s_destRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), true, 10, address(s_receiver)));
  }

  function testAutoExecSuccess() public {
    changePrank(address(s_offRamp));
    assertTrue(
      s_destRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), false, 100_000, address(s_receiver))
    );
    // Can run out of gas, should return false
    assertFalse(s_destRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), false, 1, address(s_receiver)));
  }

  // Reverts
  function testOnlyOffRampReverts() public {
    vm.expectRevert(IRouter.OnlyOffRamp.selector);
    s_destRouter.routeMessage(generateReceiverMessage(SOURCE_CHAIN_ID), true, 100_000, address(s_receiver));
  }
}

/// @notice #getFee
contract Router_getFee is EVM2EVMOnRampSetup {
  function testGetFeeSupportedChainSuccess() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertGt(expectedFee, 10e9);
  }

  // Reverts
  function testUnsupportedDestinationChainReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    vm.expectRevert(abi.encodeWithSelector(IRouterClient.UnsupportedDestinationChain.selector, 999));
    s_sourceRouter.getFee(999, message);
  }
}
