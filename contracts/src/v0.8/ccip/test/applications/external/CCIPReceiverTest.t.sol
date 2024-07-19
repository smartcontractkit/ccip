// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIPBase} from "../../../applications/external/CCIPBase.sol";

import {CCIPReceiver} from "../../../applications/external/CCIPReceiver.sol";
import {CCIPReceiverReverting} from "../../helpers/receivers/CCIPReceiverReverting.sol";

import {Client} from "../../../libraries/Client.sol";
import {EVM2EVMOnRampSetup} from "../../onRamp/EVM2EVMOnRampSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CCIPReceiverTest is EVM2EVMOnRampSetup {
  event MessageFailed(bytes32 indexed messageId, bytes reason);
  event MessageSucceeded(bytes32 indexed messageId);
  event MessageRecovered(bytes32 indexed messageId);

  CCIPReceiverReverting internal s_receiver;
  uint64 internal sourceChainSelector = 7331;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_receiver = new CCIPReceiverReverting(address(s_destRouter));

    CCIPBase.ChainUpdate[] memory chainUpdates = new CCIPBase.ChainUpdate[](1);
    chainUpdates[0] = CCIPBase.ChainUpdate({
      chainSelector: sourceChainSelector,
      allowed: true,
      recipient: abi.encode(address(s_receiver)),
      extraArgsBytes: ""
    });
    s_receiver.applyChainUpdates(chainUpdates);

    CCIPBase.ApprovedSenderUpdate[] memory senderUpdates = new CCIPBase.ApprovedSenderUpdate[](1);
    senderUpdates[0] =
      CCIPBase.ApprovedSenderUpdate({destChainSelector: sourceChainSelector, sender: abi.encode(address(s_receiver))});

    s_receiver.updateApprovedSenders(senderUpdates, new CCIPBase.ApprovedSenderUpdate[](0));
  }

  function test_Recovery_with_intentional_Revert() public {
    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // Make sure the contract call reverts so we can test recovery.
    s_receiver.setSimRevert(true);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectEmit();
    emit MessageFailed(messageId, abi.encodeWithSelector(CCIPReceiverReverting.ErrorCase.selector));

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );

    address tokenReceiver = OWNER;
    uint256 tokenReceiverBalancePre = IERC20(token).balanceOf(tokenReceiver);
    uint256 receiverBalancePre = IERC20(token).balanceOf(address(s_receiver));

    // Recovery can only be done by the owner.
    vm.startPrank(OWNER);

    vm.expectEmit();
    emit CCIPReceiver.MessageAbandoned(messageId, OWNER);

    s_receiver.abandonFailedMessage(messageId, OWNER);

    // Assert the tokens have successfully been rescued from the contract.
    assertEq(
      IERC20(token).balanceOf(tokenReceiver), tokenReceiverBalancePre + amount, "tokens not sent to tokenReceiver"
    );
    assertEq(
      IERC20(token).balanceOf(address(s_receiver)), receiverBalancePre - amount, "tokens not subtracted from receiver"
    );
  }

  function test_Recovery_from_invalid_sender_Success() public {
    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectEmit();
    emit MessageFailed(
      messageId, abi.encodeWithSelector(bytes4(CCIPBase.InvalidSender.selector), abi.encode(address(1)))
    );

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        sender: abi.encode(address(1)),
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );

    vm.stopPrank();

    // Check that the message was stored properly by comparing each of the fields.
    // There's no way to check that a function internally will revert from a top-level test, so we need to check state differences
    Client.Any2EVMMessage memory failedMessage = s_receiver.getMessageContents(messageId);
    assertEq(failedMessage.sender, abi.encode(address(1)));
    assertEq(failedMessage.sourceChainSelector, sourceChainSelector);
    assertEq(failedMessage.destTokenAmounts[0].token, token);
    assertEq(failedMessage.destTokenAmounts[0].amount, amount);

    // Check that message status is failed
    assertTrue(s_receiver.isFailedMessage(messageId), "Message should be marked as failed");

    uint256 tokenBalanceBefore = IERC20(token).balanceOf(OWNER);

    vm.startPrank(OWNER);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_receiver), OWNER, amount);
    s_receiver.withdrawTokens(token, OWNER, amount);

    assertEq(IERC20(token).balanceOf(OWNER), tokenBalanceBefore + amount);
    assertGt(IERC20(token).balanceOf(OWNER), 0);
  }

  function test_retryFailedMessage_Success() public {
    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectEmit();
    emit MessageFailed(
      messageId, abi.encodeWithSelector(bytes4(CCIPBase.InvalidSender.selector), abi.encode(address(1)))
    );

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        sender: abi.encode(address(1)),
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );

    vm.stopPrank();

    // Check that the message was stored properly by comparing each of the fields.
    // There's no way to check that a function internally will revert from a top-level test, so we need to check state differences
    Client.Any2EVMMessage memory failedMessage = s_receiver.getMessageContents(messageId);
    assertEq(failedMessage.sender, abi.encode(address(1)));
    assertEq(failedMessage.sourceChainSelector, sourceChainSelector);
    assertEq(failedMessage.destTokenAmounts[0].token, token);
    assertEq(failedMessage.destTokenAmounts[0].amount, amount);

    // Check that message status is failed
    assertTrue(s_receiver.isFailedMessage(messageId), "Message should be marked as failed");

    vm.startPrank(OWNER);

    // The message failed initially because the sender was not approved. Now we approve it and retry processing. Because retryFailedMessage() calls processMessage normally, it should execute successfully now.
    CCIPBase.ApprovedSenderUpdate[] memory senderUpdates = new CCIPBase.ApprovedSenderUpdate[](1);

    senderUpdates[0] =
      CCIPBase.ApprovedSenderUpdate({destChainSelector: sourceChainSelector, sender: abi.encode(address(1))});

    s_receiver.updateApprovedSenders(senderUpdates, new CCIPBase.ApprovedSenderUpdate[](0));

    vm.expectEmit();
    emit CCIPReceiver.MessageRecovered(messageId);

    s_receiver.retryFailedMessage(messageId);
    assertFalse(s_receiver.isFailedMessage(messageId), "Message should be marked as resolved");
  }

  function test_HappyPath_Success() public {
    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectEmit();
    emit MessageSucceeded(messageId);

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        sender: abi.encode(address(s_receiver)), // correct sender
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );
  }

  function test_apply_InvalidChainUpdate_Revert() public {
    CCIPBase.ChainUpdate[] memory chainUpdates = new CCIPBase.ChainUpdate[](1);
    chainUpdates[0] =
      CCIPBase.ChainUpdate({chainSelector: sourceChainSelector, allowed: true, recipient: "", extraArgsBytes: ""});

    // Revert because the recipient of an allowed chain is the zero address, which is prohibited
    vm.expectRevert(abi.encodeWithSelector(CCIPBase.ZeroAddressNotAllowed.selector));
    s_receiver.applyChainUpdates(chainUpdates);
  }

  function test_disableChain_andRevert_onccipReceive_Revert() public {
    CCIPBase.ChainUpdate[] memory chainUpdates = new CCIPBase.ChainUpdate[](1);
    chainUpdates[0] = CCIPBase.ChainUpdate({
      chainSelector: sourceChainSelector,
      allowed: false,
      recipient: abi.encode(address(s_receiver)),
      extraArgsBytes: ""
    });

    vm.expectEmit();
    emit CCIPBase.ChainRemoved(sourceChainSelector);

    s_receiver.applyChainUpdates(chainUpdates);

    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectRevert(abi.encodeWithSelector(CCIPBase.InvalidChain.selector, sourceChainSelector));

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        sender: abi.encode(address(1)),
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );
  }

  function test_modifyRouter_Success() public {
    vm.expectRevert(abi.encodeWithSelector(CCIPBase.ZeroAddressNotAllowed.selector));
    s_receiver.updateRouter(address(0));

    address newRouter = address(0x1234);

    vm.expectEmit();
    emit CCIPBase.CCIPRouterModified(address(s_destRouter), newRouter);

    s_receiver.updateRouter(newRouter);

    assertEq(s_receiver.getRouter(), newRouter, "Router Address not set correctly to the new router");
  }

  function test_removeSender_from_approvedList_and_revert_Success() public {
    CCIPBase.ApprovedSenderUpdate[] memory senderUpdates = new CCIPBase.ApprovedSenderUpdate[](1);
    senderUpdates[0] =
      CCIPBase.ApprovedSenderUpdate({destChainSelector: sourceChainSelector, sender: abi.encode(address(s_receiver))});

    s_receiver.updateApprovedSenders(new CCIPBase.ApprovedSenderUpdate[](0), senderUpdates);

    // assertFalse(s_receiver.s_approvedSenders(sourceChainSelector, abi.encode(address(s_receiver))));
    assertFalse(s_receiver.isApprovedSender(sourceChainSelector, abi.encode(address(s_receiver))));

    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectEmit();
    emit MessageFailed(
      messageId, abi.encodeWithSelector(bytes4(CCIPBase.InvalidSender.selector), abi.encode(address(s_receiver)))
    );

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );
  }

  function test_withdraw_nativeToken_to_owner_Success() public {
    uint256 amount = 100 ether;
    deal(address(s_receiver), amount);

    uint256 balanceBefore = OWNER.balance;

    vm.startPrank(OWNER);

    s_receiver.withdrawTokens(address(0), payable(OWNER), amount);

    assertEq(OWNER.balance, balanceBefore + amount);
  }

  function test_retryFailedMessage_which_has_not_already_failed_Revert() public {
    bytes32 messageId = keccak256("RANDOM_DATA");

    vm.expectRevert(abi.encodeWithSelector(CCIPReceiver.MessageNotFailed.selector, messageId));

    s_receiver.retryFailedMessage(messageId);
  }

  function test_abandonFailedMessage_which_has_not_already_failed_Revert() public {
    bytes32 messageId = keccak256("RANDOM_DATA");

    vm.expectRevert(abi.encodeWithSelector(CCIPReceiver.MessageNotFailed.selector, messageId));

    s_receiver.abandonFailedMessage(messageId, OWNER);
  }
}
