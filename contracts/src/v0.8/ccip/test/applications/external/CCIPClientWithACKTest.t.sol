// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIPClientWithACK} from "../../../applications/external/CCIPClientWithACK.sol";

import {CCIPBase} from "../../../applications/external/CCIPBase.sol";
import {CCIPReceiverWithACK} from "../../../applications/external/CCIPReceiverWithACK.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {Client} from "../../../libraries/Client.sol";
import {EVM2EVMOnRampSetup} from "../../onRamp/EVM2EVMOnRampSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CCIPClientWithACKTest is EVM2EVMOnRampSetup {
  event MessageFailed(bytes32 indexed messageId, bytes reason);
  event MessageSucceeded(bytes32 indexed messageId);
  event MessageRecovered(bytes32 indexed messageId);
  event MessageSent(bytes32 indexed, bytes32 indexed);
  event MessageAckSent(bytes32 incomingMessageId);
  event MessageAckReceived(bytes32);

  CCIPClientWithACK internal s_sender;
  uint64 internal destChainSelector = DEST_CHAIN_SELECTOR;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_sender = new CCIPClientWithACK(address(s_sourceRouter), IERC20(s_sourceFeeToken), false);

    CCIPBase.ChainUpdate[] memory chainUpdates = new CCIPBase.ChainUpdate[](1);
    chainUpdates[0] = CCIPBase.ChainUpdate({
      chainSelector: destChainSelector,
      allowed: true,
      recipient: abi.encode(address(s_sender)),
      extraArgsBytes: ""
    });

    s_sender.applyChainUpdates(chainUpdates);

    CCIPBase.ApprovedSenderUpdate[] memory senderUpdates = new CCIPBase.ApprovedSenderUpdate[](1);
    senderUpdates[0] =
      CCIPBase.ApprovedSenderUpdate({destChainSelector: destChainSelector, sender: abi.encode(address(s_sender))});

    s_sender.updateApprovedSenders(senderUpdates, new CCIPBase.ApprovedSenderUpdate[](0));
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

    s_sender.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_sender)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    // Check that message status is failed
    assertTrue(s_sender.isFailedMessage(messageId), "Message Should be marked as failed");
  }

  function test_ccipReceiveAndSendAck_Success() public {
    bytes32 messageId = keccak256("messageId");
    bytes32 ackMessageId = 0x07d90483b3ed7831c5402af6402e21ba3740a15e9d0837f7c7effb1cbffb39f7;
    address token = address(s_sourceFeeToken);
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_sender), 1e24);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    CCIPReceiverWithACK.MessagePayload memory payload = CCIPReceiverWithACK.MessagePayload({
      version: "",
      data: "FAKE_DATA",
      messageType: CCIPReceiverWithACK.MessageType.OUTGOING
    });

    Client.EVM2AnyMessage memory ackMessage = Client.EVM2AnyMessage({
      receiver: abi.encode(address(s_sender)),
      data: abi.encode(s_sender.ACK_MESSAGE_HEADER(), messageId),
      tokenAmounts: destTokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });

    uint256 feeTokenAmount = s_sourceRouter.getFee(destChainSelector, ackMessage);

    uint256 receiverBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(address(s_sender));

    // Check the messageId since we can control that, but not ackMessageId since its generated at execution time
    vm.expectEmit(true, false, true, false, address(s_sender));
    emit MessageSent(messageId, ackMessageId);

    s_sender.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_sender)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    // Check that fee token is properly subtracted from balance to pay for ack message
    assertEq(IERC20(s_sourceFeeToken).balanceOf(address(s_sender)), receiverBalanceBefore - feeTokenAmount);
  }

  function test_ccipSend_withNonNativeFeetoken_andNoDestTokens_Success() public {
    address token = address(s_sourceFeeToken);
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // Make sure we give the receiver contract enough tokens like CCIP would.
    IERC20(token).approve(address(s_sender), type(uint256).max);

    assertFalse(s_sender.usePreFundedFeeTokens(), "Not Using pre-funded fee tokens");

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(address(s_sender)),
      data: "",
      tokenAmounts: destTokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });

    uint256 feeTokenAmount = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    uint256 feeTokenBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(OWNER);

    s_sender.ccipSend({destChainSelector: DEST_CHAIN_SELECTOR, tokenAmounts: destTokenAmounts, data: ""});
    // feeToken: address(s_sourceFeeToken)

    // Assert that tokens were transfered for bridging + fees
    assertEq(IERC20(token).balanceOf(OWNER), feeTokenBalanceBefore - feeTokenAmount);
  }

  function test_ccipSend_with_NativeFeeToken_andNoDestTokens_Success() public {
    address token = address(s_sourceFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    s_sender.updateFeeToken(address(0));

    // Make sure we give the receiver contract enough tokens like CCIP would.
    IERC20(token).approve(address(s_sender), type(uint256).max);

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(address(s_sender)),
      data: "",
      tokenAmounts: destTokenAmounts,
      extraArgs: "",
      feeToken: address(s_sourceFeeToken)
    });

    uint256 feeTokenAmount = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    uint256 tokenBalanceBefore = IERC20(token).balanceOf(OWNER);
    uint256 nativeFeeTokenBalanceBefore = OWNER.balance;

    s_sender.ccipSend{value: feeTokenAmount}({
      destChainSelector: DEST_CHAIN_SELECTOR,
      tokenAmounts: destTokenAmounts,
      data: ""
    });

    // Assert that native fees are paid successfully and tokens are transferred
    assertEq(IERC20(token).balanceOf(OWNER), tokenBalanceBefore - amount, "Tokens were not successfully delivered");
    assertEq(
      OWNER.balance, nativeFeeTokenBalanceBefore - feeTokenAmount, "Native fee tokens were not successfully forwarded"
    );
  }

  function test_ccipSendAndReceiveAck_in_return_Success() public {
    address token = address(s_sourceFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    IERC20(token).approve(address(s_sender), type(uint256).max);

    bytes32 messageId =
      s_sender.ccipSend({destChainSelector: DEST_CHAIN_SELECTOR, tokenAmounts: destTokenAmounts, data: ""});
    // feeToken: address(s_sourceFeeToken)

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    CCIPReceiverWithACK.MessagePayload memory payload = CCIPReceiverWithACK.MessagePayload({
      version: "",
      data: abi.encode(s_sender.ACK_MESSAGE_HEADER(), messageId),
      messageType: CCIPReceiverWithACK.MessageType.ACK
    });

    vm.expectEmit();
    emit MessageAckReceived(messageId);

    s_sender.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: destChainSelector,
        sender: abi.encode(address(s_sender)),
        data: abi.encode(payload),
        destTokenAmounts: destTokenAmounts
      })
    );

    assertEq(
      uint256(s_sender.s_messageStatus(messageId)),
      uint256(CCIPReceiverWithACK.MessageStatus.ACKNOWLEDGED),
      "Ack message was not properly received"
    );
  }

  function test_send_tokens_that_are_not_feeToken_Success() public {
    address token = s_sourceTokens[1];
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    IERC20(token).approve(address(s_sender), type(uint256).max);
    IERC20(s_sourceFeeToken).approve(address(s_sender), type(uint256).max);
    deal(token, address(this), 1e24);

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(address(s_sender)),
      data: "",
      tokenAmounts: destTokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });

    uint256 feeTokenAmount = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    uint256 tokenBalanceBefore = IERC20(token).balanceOf(OWNER);
    uint256 feeTokenBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(OWNER);

    s_sender.ccipSend({destChainSelector: DEST_CHAIN_SELECTOR, tokenAmounts: destTokenAmounts, data: ""});

    // Assert that tokens were transfered for bridging + fees
    assertEq(IERC20(token).balanceOf(OWNER), tokenBalanceBefore - amount);
    assertEq(IERC20(s_sourceFeeToken).balanceOf(OWNER), feeTokenBalanceBefore - feeTokenAmount);
  }
}
