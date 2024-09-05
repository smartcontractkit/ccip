// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIPBase} from "../../../applications/external/CCIPBase.sol";
import {CCIPReceiver} from "../../../applications/external/CCIPReceiver.sol";
import {CCIPReceiverWithACK} from "../../../applications/external/CCIPReceiverWithACK.sol";

import {Client} from "../../../libraries/Client.sol";
import {EVM2EVMOnRampSetup} from "../../onRamp/EVM2EVMOnRampSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CCIPReceiverWithAckTest is EVM2EVMOnRampSetup {
  event MessageFailed(bytes32 indexed messageId, bytes reason);
  event MessageSucceeded(bytes32 indexed messageId);
  event MessageRecovered(bytes32 indexed messageId);
  event MessageSent(bytes32 indexed incomingmessageId, bytes32 indexed ackmessageId);
  event MessageAckSent(bytes32 incomingMessageId);
  event MessageAckReceived(bytes32);

  CCIPReceiverWithACK internal s_receiver;
  uint64 internal destChainSelector = DEST_CHAIN_SELECTOR;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_receiver = new CCIPReceiverWithACK(address(s_sourceRouter), IERC20(s_sourceFeeToken));

    CCIPBase.ChainUpdate[] memory chainUpdates = new CCIPBase.ChainUpdate[](1);
    chainUpdates[0] = CCIPBase.ChainUpdate({
      chainSelector: destChainSelector,
      allowed: true,
      recipient: abi.encode(address(s_receiver)),
      extraArgsBytes: ""
    });
    s_receiver.applyChainUpdates(chainUpdates);

    CCIPBase.ApprovedSenderUpdate[] memory senderUpdates = new CCIPBase.ApprovedSenderUpdate[](1);
    senderUpdates[0] =
      CCIPBase.ApprovedSenderUpdate({destChainSelector: destChainSelector, sender: abi.encode(address(s_receiver))});

    s_receiver.updateApprovedSenders(senderUpdates, new CCIPBase.ApprovedSenderUpdate[](0));
  }

  function test_ccipReceive_and_respond_with_ack_Success() public {
    bytes32 messageId = keccak256("messageId");
    bytes32 ackMessageId = 0x07d90483b3ed7831c5402af6402e21ba3740a15e9d0837f7c7effb1cbffb39f7;
    address token = address(s_sourceFeeToken);
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), 1e24);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    CCIPReceiverWithACK.MessagePayload memory payload = CCIPReceiverWithACK.MessagePayload({
      version: "",
      data: "FAKE_DATA",
      messageType: CCIPReceiverWithACK.MessageType.OUTGOING
    });

    Client.EVM2AnyMessage memory ackMessage = Client.EVM2AnyMessage({
      receiver: abi.encode(address(s_receiver)),
      data: abi.encode(s_receiver.ACK_MESSAGE_HEADER(), messageId),
      tokenAmounts: destTokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });

    uint256 feeTokenAmount = s_sourceRouter.getFee(destChainSelector, ackMessage);

    uint256 receiverBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(address(s_receiver));

    vm.expectEmit(true, false, true, false);
    emit MessageSent(messageId, ackMessageId);

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    // Check that fee token is properly subtracted from balance to pay for ack message
    assertEq(IERC20(s_sourceFeeToken).balanceOf(address(s_receiver)), receiverBalanceBefore - feeTokenAmount);
  }

  function test_ccipReceive_ack_message_Success() public {
    bytes32 messageId = keccak256("messageId");
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    CCIPReceiverWithACK.MessagePayload memory payload = CCIPReceiverWithACK.MessagePayload({
      version: "",
      data: abi.encode(s_receiver.ACK_MESSAGE_HEADER(), messageId),
      messageType: CCIPReceiverWithACK.MessageType.ACK
    });

    vm.expectEmit();
    emit MessageAckReceived(messageId);

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    assertEq(
      uint256(s_receiver.s_messageStatus(messageId)),
      uint256(CCIPReceiverWithACK.MessageStatus.ACKNOWLEDGED),
      "Ack message was not properly received"
    );
  }

  function test_ccipReceiver_ack_with_invalidAckMessageHeaderBytes_Revert() public {
    bytes32 messageId = keccak256("messageId");
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    // Payload with incorrect ack message header should revert
    CCIPReceiverWithACK.MessagePayload memory payload = CCIPReceiverWithACK.MessagePayload({
      version: "",
      data: abi.encode("RANDOM_BYTES", messageId),
      messageType: CCIPReceiverWithACK.MessageType.ACK
    });

    // Expect the processing to revert from invalid Ack Message Header
    vm.expectEmit();
    emit MessageFailed(messageId, abi.encodeWithSelector(bytes4(CCIPReceiverWithACK.InvalidAckMessageHeader.selector)));

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    // Check that message status is failed
    assertTrue(s_receiver.isFailedMessage(messageId), "Message should be marked as failed");
  }

  function test_modifyFeeToken_Success() public {
    // WETH is used as a placeholder for any ERC20 token
    address WETH = s_sourceRouter.getWrappedNative();

    vm.expectEmit();
    emit IERC20.Approval(address(s_receiver), address(s_sourceRouter), 0);

    vm.expectEmit();
    emit CCIPReceiverWithACK.FeeTokenUpdated(s_sourceFeeToken, WETH);

    s_receiver.updateFeeToken(WETH);

    IERC20 newFeeToken = IERC20(s_receiver.getFeeToken());
    assertEq(address(newFeeToken), WETH);
    assertEq(newFeeToken.allowance(address(s_receiver), address(s_sourceRouter)), type(uint256).max);
    assertEq(IERC20(s_sourceFeeToken).allowance(address(s_receiver), address(s_sourceRouter)), 0);
  }

  function test_feeTokenApproval_in_constructor_Success() public {
    CCIPReceiverWithACK newReceiver = new CCIPReceiverWithACK(address(s_sourceRouter), IERC20(s_sourceFeeToken));

    assertEq(IERC20(s_sourceFeeToken).allowance(address(newReceiver), address(s_sourceRouter)), type(uint256).max);
  }

  function test_attemptACK_message_which_has_already_been_acknowledged_Revert() public {
    bytes32 messageId = keccak256("messageId");
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    CCIPReceiverWithACK.MessagePayload memory payload = CCIPReceiverWithACK.MessagePayload({
      version: "",
      data: abi.encode(s_receiver.ACK_MESSAGE_HEADER(), messageId),
      messageType: CCIPReceiverWithACK.MessageType.ACK
    });

    vm.expectEmit();
    emit MessageAckReceived(messageId);

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    // Assert that the message was received and ACK'ED the first time
    assertEq(
      uint256(s_receiver.s_messageStatus(messageId)),
      uint256(CCIPReceiverWithACK.MessageStatus.ACKNOWLEDGED),
      "Ack message was not properly received"
    );

    vm.expectEmit();
    emit CCIPReceiver.MessageFailed(
      messageId, abi.encodeWithSelector(CCIPReceiverWithACK.MessageAlreadyAcknowledged.selector, messageId)
    );

    s_receiver.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_receiver)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );
  }
}
