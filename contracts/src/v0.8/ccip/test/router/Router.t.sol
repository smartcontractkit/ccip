// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IEVM2AnyOnRamp} from "../../interfaces/IEVM2AnyOnRamp.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";
import {IRouterClient} from "../../interfaces/IRouterClient.sol";

import "../onRamp/EVM2EVMOnRampSetup.t.sol";
import {MaybeRevertMessageReceiver} from "../helpers/receivers/MaybeRevertMessageReceiver.sol";
import "../offRamp/EVM2EVMOffRampSetup.t.sol";

/// @notice #constructor
contract Router_constructor is EVM2EVMOnRampSetup {
  function testConstructorSuccess() public {
    assertEq("Router 1.0.0", s_sourceRouter.typeAndVersion());
    // owner
    assertEq(OWNER, s_sourceRouter.owner());
  }
}

/// @notice #recoverTokens
contract Router_recoverTokens is EVM2EVMOnRampSetup {
  function testRecoverTokensSuccess() public {
    // Assert we can recover sourceToken
    IERC20 token = IERC20(s_sourceTokens[0]);
    uint256 balanceBefore = token.balanceOf(OWNER);
    token.transfer(address(s_sourceRouter), 1);
    assertEq(token.balanceOf(address(s_sourceRouter)), 1);
    s_sourceRouter.recoverTokens(address(token), OWNER, 1);
    assertEq(token.balanceOf(address(s_sourceRouter)), 0);
    assertEq(token.balanceOf(OWNER), balanceBefore);

    // Assert we can recover native
    balanceBefore = OWNER.balance;
    deal(address(s_sourceRouter), 10);
    assertEq(address(s_sourceRouter).balance, 10);
    s_sourceRouter.recoverTokens(address(0), OWNER, 10);
    assertEq(OWNER.balance, balanceBefore + 10);
    assertEq(address(s_sourceRouter).balance, 0);
  }

  function testRecoverTokensNonOwnerReverts() public {
    // Reverts if not owner
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.recoverTokens(address(0), STRANGER, 1);
  }

  function testRecoverTokensInvalidRecipientReverts() public {
    vm.expectRevert(abi.encodeWithSelector(Router.InvalidRecipientAddress.selector, address(0)));
    s_sourceRouter.recoverTokens(address(0), address(0), 1);
  }

  function testRecoverTokensNoFundsReverts() public {
    // Reverts if no funds present
    vm.expectRevert();
    s_sourceRouter.recoverTokens(address(0), OWNER, 10);
  }

  function testRecoverTokensValueReceiverReverts() public {
    MaybeRevertMessageReceiver revertingValueReceiver = new MaybeRevertMessageReceiver(true);
    deal(address(s_sourceRouter), 10);

    // Value receiver reverts
    vm.expectRevert(Router.FailedToSendValue.selector);
    s_sourceRouter.recoverTokens(address(0), address(revertingValueReceiver), 10);
  }
}

/// @notice #ccipSend
contract Router_ccipSend is EVM2EVMOnRampSetup {
  event Burned(address indexed sender, uint256 amount);

  function testCCIPSendLinkFeeOneTokenSuccess_gas() public {
    vm.pauseGasMetering();
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    IERC20 sourceToken1 = IERC20(s_sourceTokens[1]);
    sourceToken1.approve(address(s_sourceRouter), 2 ** 64);

    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].amount = 2 ** 64;
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
    sourceToken1.approve(address(s_sourceRouter), 2 ** 64);

    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0].amount = 2 ** 64;
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
      networkFeeUSD: 1,
      minTokenTransferFeeUSD: 1,
      maxTokenTransferFeeUSD: 5,
      gasMultiplier: 108e16,
      premiumMultiplier: 1e18,
      enabled: true
    });
    s_onRamp.setFeeTokenConfig(feeTokenConfigArgs);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceTokens[1];
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = s_sourceTokens[1];
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2 ** 64);
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

  // Since sending with zero fees is a legitimate use case for some destination
  // chains, e.g. private chains, we want to make sure that we can still send even
  // when the configured fee is 0.
  function testZeroFeeAndGasPriceSuccess() public {
    // Configure a new fee token that has zero gas and zero fees but is still
    // enabled and valid to pay with.
    address feeTokenWithZeroFeeAndGas = s_sourceTokens[1];

    // Set the new token as feeToken
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = feeTokenWithZeroFeeAndGas;
    s_priceRegistry.applyFeeTokensUpdates(feeTokens, new address[](0));

    // Update the price of the newly set feeToken
    Internal.PriceUpdates memory priceUpdates = getSinglePriceUpdateStruct(feeTokenWithZeroFeeAndGas, 2_000 ether);
    priceUpdates.destChainSelector = DEST_CHAIN_ID;
    priceUpdates.usdPerUnitGas = 0;

    s_priceRegistry.updatePrices(priceUpdates);

    // Set the feeToken args on the onRamp
    EVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeTokenConfigArgs = new EVM2EVMOnRamp.FeeTokenConfigArgs[](1);
    feeTokenConfigArgs[0] = EVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceTokens[1],
      networkFeeUSD: 0,
      minTokenTransferFeeUSD: 0,
      maxTokenTransferFeeUSD: 5,
      gasMultiplier: 108e16,
      premiumMultiplier: 1e18,
      enabled: true
    });

    s_onRamp.setFeeTokenConfig(feeTokenConfigArgs);

    // Send a message with the new feeToken
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = feeTokenWithZeroFeeAndGas;

    // Fee should be 0 and sending should not revert
    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_ID, message);
    assertEq(fee, 0);

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  // Reverts

  function testWhenNotHealthyReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    s_mockARM.voteToCurse(bytes32(0));
    vm.expectRevert(Router.BadARMSignal.selector);
    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testUnsupportedDestinationChainReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint64 wrongChain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(IRouterClient.UnsupportedDestinationChain.selector, wrongChain));

    s_sourceRouter.ccipSend(wrongChain, message);
  }

  function testFuzz_UnsupportedFeeTokenReverts(address wrongFeeToken) public {
    // We have three fee tokens set, all others should revert.
    vm.assume(address(s_sourceFeeToken) != wrongFeeToken);
    vm.assume(address(s_sourceRouter.getWrappedNative()) != wrongFeeToken);
    vm.assume(address(0) != wrongFeeToken);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.feeToken = wrongFeeToken;

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOnRamp.NotAFeeToken.selector, wrongFeeToken));

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFuzz_UnsupportedTokenReverts(address wrongToken) public {
    for (uint256 i = 0; i < s_sourceTokens.length; ++i) {
      vm.assume(address(s_sourceTokens[i]) != wrongToken);
    }
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: wrongToken, amount: 1});
    message.tokenAmounts = tokenAmounts;

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOnRamp.UnsupportedToken.selector, wrongToken));

    s_sourceRouter.ccipSend(DEST_CHAIN_ID, message);
  }

  function testFeeTokenAmountTooLowReverts() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 0);

    vm.expectRevert("ERC20: insufficient allowance");

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

    s_onRamp.getFeeTokenConfig(s_sourceRouter.getWrappedNative());

    hoax(address(1), 1);
    vm.expectRevert(IRouterClient.InsufficientFeeTokenAmount.selector);
    s_sourceRouter.ccipSend{value: 1}(DEST_CHAIN_ID, message);
  }
}

contract Router_applyRampUpdates is RouterSetup {
  event OffRampAdded(uint64 indexed sourceChainSelector, address offRamp);

  MaybeRevertMessageReceiver internal s_receiver;

  function setUp() public virtual override(RouterSetup) {
    RouterSetup.setUp();
    s_receiver = new MaybeRevertMessageReceiver(false);
  }

  function testOffRampDisable() public {
    // Add ingress
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](0);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    address offRamp = address(uint160(2));
    offRampUpdates[0] = Router.OffRamp(SOURCE_CHAIN_ID, offRamp);
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
    assertEq(1, s_sourceRouter.getOffRamps().length);
    Router.OffRamp[] memory gotOffRamps = s_sourceRouter.getOffRamps();
    assertEq(offRampUpdates[0].sourceChainSelector, gotOffRamps[0].sourceChainSelector);
    assertEq(offRampUpdates[0].offRamp, gotOffRamps[0].offRamp);
    // Remove ingress
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates, new Router.OffRamp[](0));
    assertEq(0, s_sourceRouter.getOffRamps().length);

    // Disabled offramp should not be able to route.
    vm.expectRevert(IRouter.OnlyOffRamp.selector);
    changePrank(offRamp);
    s_sourceRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      100_000,
      address(s_receiver)
    );
    changePrank(OWNER);

    // Re-enabling should succeed
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
    assertEq(1, s_sourceRouter.getOffRamps().length);
    gotOffRamps = s_sourceRouter.getOffRamps();
    assertEq(offRampUpdates[0].sourceChainSelector, gotOffRamps[0].sourceChainSelector);
    assertEq(offRampUpdates[0].offRamp, gotOffRamps[0].offRamp);
    changePrank(offRamp);
    s_sourceRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      100_000,
      address(s_receiver)
    );
  }

  function testOnRampDisable() public {
    // Add onRamp
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](0);
    address onRamp = address(uint160(2));
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_ID, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
    assertEq(onRamp, s_sourceRouter.getOnRamp(DEST_CHAIN_ID));
    assertTrue(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));

    // Disable onRamp
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_ID, onRamp: address(0)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));
    assertEq(address(0), s_sourceRouter.getOnRamp(DEST_CHAIN_ID));
    assertFalse(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));

    // Re-enable onRamp
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_ID, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));
    assertEq(onRamp, s_sourceRouter.getOnRamp(DEST_CHAIN_ID));
    assertTrue(s_sourceRouter.isChainSupported(DEST_CHAIN_ID));
  }

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](0);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](0);
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates, offRampUpdates);
  }

  function testOffRampMismatchReverts() public {
    address offRamp = address(uint160(2));

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](0);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp(DEST_CHAIN_ID, offRamp);

    vm.expectEmit();
    emit OffRampAdded(DEST_CHAIN_ID, offRamp);
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    offRampUpdates[0] = Router.OffRamp(SOURCE_CHAIN_ID, offRamp);

    vm.expectRevert(Router.OffRampMismatch.selector);
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates, offRampUpdates);
  }
}

/// @notice #setWrappedNative
contract Router_setWrappedNative is EVM2EVMOnRampSetup {
  function testFuzz_SetWrappedNativeSuccess(address wrappedNative) public {
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
  event MessageExecuted(bytes32 messageId, uint64 sourceChainSelector, address offRamp, bytes32 calldataHash);

  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
    changePrank(address(s_offRamp));
  }

  function testManualExecSuccess() public {
    Client.Any2EVMMessage memory message = generateReceiverMessage(SOURCE_CHAIN_ID);
    // Manuel execution cannot run out of gas

    (bool success, bytes memory retData) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      generateManualGasLimit(message.data.length),
      address(s_receiver)
    );
    assertTrue(success);
    assertEq("", retData);
  }

  function testExecutionEventSuccess() public {
    Client.Any2EVMMessage memory message = generateReceiverMessage(SOURCE_CHAIN_ID);
    // Should revert with reason
    bytes memory realError1 = new bytes(2);
    realError1[0] = 0xbe;
    realError1[1] = 0xef;
    s_reverting_receiver.setErr(realError1);

    vm.expectEmit();
    emit MessageExecuted(
      message.messageId,
      message.sourceChainSelector,
      address(s_offRamp),
      keccak256(abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message))
    );

    (bool success, bytes memory retData) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      generateManualGasLimit(message.data.length),
      address(s_reverting_receiver)
    );

    assertFalse(success);
    assertEq(abi.encodeWithSelector(MaybeRevertMessageReceiver.CustomError.selector, realError1), retData);

    // Reason is truncated
    // Over the MAX_RET_BYTES limit (including offset and length word since we have a dynamic values), should be ignored
    bytes memory realError2 = new bytes(32 * 2 + 1);
    realError2[32 * 2 - 1] = 0xAA;
    realError2[32 * 2] = 0xFF;
    s_reverting_receiver.setErr(realError2);

    vm.expectEmit();
    emit MessageExecuted(
      message.messageId,
      message.sourceChainSelector,
      address(s_offRamp),
      keccak256(abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message))
    );

    (success, retData) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      generateManualGasLimit(message.data.length),
      address(s_reverting_receiver)
    );

    assertFalse(success);
    assertEq(
      abi.encodeWithSelector(
        MaybeRevertMessageReceiver.CustomError.selector,
        uint256(32),
        uint256(realError2.length),
        uint256(0),
        uint256(0xAA)
      ),
      retData
    );

    // Should emit success
    vm.expectEmit();
    emit MessageExecuted(
      message.messageId,
      message.sourceChainSelector,
      address(s_offRamp),
      keccak256(abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message))
    );

    (success, retData) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      generateManualGasLimit(message.data.length),
      address(s_receiver)
    );

    assertTrue(success);
    assertEq("", retData);
  }

  function testFuzz_ExecutionEventSuccess(bytes calldata error) public {
    Client.Any2EVMMessage memory message = generateReceiverMessage(SOURCE_CHAIN_ID);
    s_reverting_receiver.setErr(error);

    bytes memory expectedRetData;

    if (error.length >= 33) {
      uint256 cutOff = error.length > 64 ? 64 : error.length;
      vm.expectEmit();
      emit MessageExecuted(
        message.messageId,
        message.sourceChainSelector,
        address(s_offRamp),
        keccak256(abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message))
      );
      expectedRetData = abi.encodeWithSelector(
        MaybeRevertMessageReceiver.CustomError.selector,
        uint256(32),
        uint256(error.length),
        bytes32(error[:32]),
        bytes32(error[32:cutOff])
      );
    } else {
      vm.expectEmit();
      emit MessageExecuted(
        message.messageId,
        message.sourceChainSelector,
        address(s_offRamp),
        keccak256(abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message))
      );
      expectedRetData = abi.encodeWithSelector(MaybeRevertMessageReceiver.CustomError.selector, error);
    }

    (bool success, bytes memory retData) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      generateManualGasLimit(message.data.length),
      address(s_reverting_receiver)
    );

    assertFalse(success);
    assertEq(expectedRetData, retData);
  }

  function testAutoExecSuccess() public {
    (bool success, ) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      100_000,
      address(s_receiver)
    );

    assertTrue(success);

    (success, ) = s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      1,
      address(s_receiver)
    );

    // Can run out of gas, should return false
    assertFalse(success);
  }

  // Reverts
  function testOnlyOffRampReverts() public {
    changePrank(STRANGER);

    vm.expectRevert(IRouter.OnlyOffRamp.selector);
    s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      100_000,
      address(s_receiver)
    );
  }

  function testWhenNotHealthyReverts() public {
    s_mockARM.voteToCurse(bytes32(0));
    vm.expectRevert(Router.BadARMSignal.selector);
    s_destRouter.routeMessage(
      generateReceiverMessage(SOURCE_CHAIN_ID),
      GAS_FOR_CALL_EXACT_CHECK,
      100_000,
      address(s_receiver)
    );
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
