// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIPBase} from "../../../applications/external/CCIPBase.sol";
import {CCIPClient} from "../../../applications/external/CCIPClient.sol";

import {Client} from "../../../libraries/Client.sol";
import {EVM2EVMOnRampSetup} from "../../onRamp/EVM2EVMOnRampSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

contract CCIPClientTest is EVM2EVMOnRampSetup {
  event MessageFailed(bytes32 indexed messageId, bytes reason);
  event MessageSucceeded(bytes32 indexed messageId);
  event MessageRecovered(bytes32 indexed messageId);

  CCIPClient internal s_sender;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_sender = new CCIPClient(address(s_sourceRouter), IERC20(s_sourceFeeToken), false);

    CCIPBase.ChainUpdate[] memory chainUpdates = new CCIPBase.ChainUpdate[](1);
    chainUpdates[0] = CCIPBase.ChainUpdate({
      chainSelector: DEST_CHAIN_SELECTOR,
      allowed: true,
      recipient: abi.encode(address(s_sender)),
      extraArgsBytes: ""
    });
    s_sender.applyChainUpdates(chainUpdates);

    CCIPBase.ApprovedSenderUpdate[] memory senderUpdates = new CCIPBase.ApprovedSenderUpdate[](1);
    senderUpdates[0] =
      CCIPBase.ApprovedSenderUpdate({destChainSelector: DEST_CHAIN_SELECTOR, sender: abi.encode(address(s_sender))});

    s_sender.updateApprovedSenders(senderUpdates, new CCIPBase.ApprovedSenderUpdate[](0));
  }

  function test_ccipSend_withNonNativeFeetoken_andDestTokens_Success() public {
    address token = address(s_sourceFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    IERC20(token).approve(address(s_sender), type(uint256).max);

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

    // Assert that tokens were transfered for bridging + fees
    assertEq(IERC20(token).balanceOf(OWNER), feeTokenBalanceBefore - amount - feeTokenAmount);
  }

  function test_ccipSend_withNonNativeFeetoken_andNoDestTokens_Success() public {
    address token = address(s_sourceFeeToken);
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);

    // Make sure we give the receiver contract enough tokens like CCIP would.
    IERC20(token).approve(address(s_sender), type(uint256).max);

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

    // Assert that tokens were transfered for bridging + fees
    assertEq(IERC20(token).balanceOf(OWNER), feeTokenBalanceBefore - feeTokenAmount);
  }

  function test_modifyFeeToken_Success() public {
    // WETH is used as a placeholder for any ERC20 token
    address WETH = s_sourceRouter.getWrappedNative();

    vm.expectEmit();
    emit IERC20.Approval(address(s_sender), address(s_sourceRouter), 0);

    vm.expectEmit();
    emit CCIPClient.FeeTokenUpdated(s_sourceFeeToken, WETH);

    s_sender.updateFeeToken(WETH);

    IERC20 newFeeToken = IERC20(s_sender.getFeeToken());
    assertEq(address(newFeeToken), WETH);
    assertEq(newFeeToken.allowance(address(s_sender), address(s_sourceRouter)), type(uint256).max);
    assertEq(IERC20(s_sourceFeeToken).allowance(address(s_sender), address(s_sourceRouter)), 0);
  }

  function test_ccipSend_with_NativeFeeToken_andDestTokens_Success() public {
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

  function test_HappyPath_Success() public {
    bytes32 messageId = keccak256("messageId");
    address token = address(s_destFeeToken);
    uint256 amount = 111333333777;
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    // Make sure we give the receiver contract enough tokens like CCIP would.
    deal(token, address(s_sender), amount);

    // The receiver contract will revert if the router is not the sender.
    vm.startPrank(address(s_sourceRouter));

    vm.expectEmit();
    emit MessageSucceeded(messageId);

    s_sender.ccipReceive(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: DEST_CHAIN_SELECTOR,
        sender: abi.encode(address(s_sender)), // correct sender
        data: "",
        destTokenAmounts: destTokenAmounts
      })
    );
  }
}
