// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIPReceiver} from "../../../applications/external/CCIPReceiver.sol";
import {ICCIPClientBase} from "../../../interfaces/ICCIPClientBase.sol";

import {Client} from "../../../libraries/Client.sol";
import {EVM2EVMOnRampSetup} from "../../onRamp/EVM2EVMOnRampSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CCIPReceiverTest is EVM2EVMOnRampSetup {
  event MessageFailed(bytes32 indexed messageId, bytes reason);
  event MessageSucceeded(bytes32 indexed messageId);
  event MessageRecovered(bytes32 indexed messageId);

  CCIPReceiver internal s_receiver;
  uint64 internal sourceChainSelector = 7331;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_receiver = new CCIPReceiver(address(s_destRouter));
    s_receiver.enableChain(sourceChainSelector, abi.encode(address(1)), "");

    ICCIPClientBase.approvedSenderUpdate[] memory senderUpdates = new ICCIPClientBase.approvedSenderUpdate[](1);
    senderUpdates[0] = ICCIPClientBase.approvedSenderUpdate({
      destChainSelector: sourceChainSelector,
      sender: abi.encode(address(s_receiver))
    });

    s_receiver.updateApprovedSenders(senderUpdates, new ICCIPClientBase.approvedSenderUpdate[](0));
  }

  function test_Recovery_with_intentional_revert() public {
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
    emit MessageFailed(messageId, abi.encodeWithSelector(CCIPReceiver.ErrorCase.selector));

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
    emit MessageRecovered(messageId);

    s_receiver.retryFailedMessage(messageId);

    // Assert the tokens have successfully been rescued from the contract.
    assertEq(
      IERC20(token).balanceOf(tokenReceiver), tokenReceiverBalancePre + amount, "tokens not sent to tokenReceiver"
    );
    assertEq(
      IERC20(token).balanceOf(address(s_receiver)), receiverBalancePre - amount, "tokens not subtracted from receiver"
    );
  }

  function test_Recovery_from_invalid_sender() public {
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
      messageId, abi.encodeWithSelector(bytes4(ICCIPClientBase.InvalidSender.selector), abi.encode(address(1)))
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
    assertEq(s_receiver.getMessageStatus(messageId), 1);

    uint256 tokenBalanceBefore = IERC20(token).balanceOf(OWNER);

    vm.startPrank(OWNER);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_receiver), OWNER, amount);
    s_receiver.withdrawTokens(token, OWNER, amount);

    assertEq(IERC20(token).balanceOf(OWNER), tokenBalanceBefore + amount);
    assertGt(IERC20(token).balanceOf(OWNER), 0);
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

  function test_disableChain_andRevert_onccipReceive_REVERT() public {
    s_receiver.disableChain(sourceChainSelector);

    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_receiver), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_destRouter));

    vm.expectRevert(abi.encodeWithSelector(ICCIPClientBase.InvalidChain.selector, sourceChainSelector));

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

  function test_removeSender_from_approvedList_and_revert() public {
    ICCIPClientBase.approvedSenderUpdate[] memory senderUpdates = new ICCIPClientBase.approvedSenderUpdate[](1);
    senderUpdates[0] = ICCIPClientBase.approvedSenderUpdate({
      destChainSelector: sourceChainSelector,
      sender: abi.encode(address(s_receiver))
    });

    s_receiver.updateApprovedSenders(new ICCIPClientBase.approvedSenderUpdate[](0), senderUpdates);

    assertFalse(s_receiver.s_approvedSenders(sourceChainSelector, abi.encode(address(s_receiver))));

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
      messageId, abi.encodeWithSelector(bytes4(ICCIPClientBase.InvalidSender.selector), abi.encode(address(s_receiver)))
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

  function test_withdraw_nativeToken_to_owner() public {
    uint256 amount = 100 ether;
    deal(address(s_receiver), amount);

    uint256 balanceBefore = OWNER.balance;

    vm.startPrank(OWNER);

    s_receiver.withdrawNativeToken(payable(OWNER), amount);

    assertEq(OWNER.balance, balanceBefore + amount);
  }
}
