// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

import {CCIPRouter} from "../../applications/EtherSenderReceiver.sol";
import {EtherSenderReceiverHelper} from "./../helpers/EtherSenderReceiverHelper.sol";
import {Client} from "../../libraries/Client.sol";
import {WETH9} from "../WETH9.sol";

contract EtherSenderReceiverTest is Test {
  EtherSenderReceiverHelper internal s_etherSenderReceiver;
  WETH9 internal s_weth;
  WETH9 internal s_someOtherWeth;

  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant ROUTER = 0x0F3779ee3a832D10158073ae2F5e61ac7FBBF880;
  address internal constant XCHAIN_RECEIVER = 0xBd91b2073218AF872BF73b65e2e5950ea356d147;
  address internal constant FINAL_RECEIVER = 0xfC6dbA3917083b749285D96984cc70c0F7760756;

  function setUp() public {
    vm.startPrank(OWNER);

    s_someOtherWeth = new WETH9();
    s_weth = new WETH9();
    vm.mockCall(
      ROUTER,
      abi.encodeWithSelector(CCIPRouter.getWrappedNative.selector),
      abi.encode(address(s_weth)));
    s_etherSenderReceiver = new EtherSenderReceiverHelper(ROUTER);

    deal(OWNER, 1 ether);
  }
}

contract EtherSenderReceiverTest_validateFeeToken is EtherSenderReceiverTest {
  function test_validateFeeToken_valid() public {
    uint256 amount = 100;
    {
      // Case 1: feeToken is address(0), i.e native.
      Client.EVMTokenAmount[] memory tokenAmount = new Client.EVMTokenAmount[](1);
      tokenAmount[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmount,
        feeToken: address(0),
        extraArgs: ""
      });

      s_etherSenderReceiver.validateFeeToken{value: amount+1}(message);
    }
    {
      // Case 2: feeToken is a nonzero address.
      Client.EVMTokenAmount[] memory tokenAmount = new Client.EVMTokenAmount[](1);
      tokenAmount[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmount,
        feeToken: address(s_weth),
        extraArgs: ""
      });

      s_etherSenderReceiver.validateFeeToken{value: amount}(message);
    }
  }
}

contract EtherSenderReceiverTest_validateMessage is EtherSenderReceiverTest {
  error InvalidDestinationReceiver(bytes destReceiver);
  error InvalidTokenAmounts(uint256 gotAmounts);
  error InvalidWethAddress(address want, address got);
  error GasLimitTooLow(uint256 minLimit, uint256 gotLimit);

  function test_validateMessage_validMessage() public view {
    uint256 amount = 100;
    {
      // Case 1: Valid message.
      // extraArgs not specified.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      s_etherSenderReceiver.validateMessage(message);
    }
    {
      // Case 2: Valid message.
      // extraArgs is specified.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({
          gasLimit: 200_000
        }))
      });

      s_etherSenderReceiver.validateMessage(message);
    }
  }

  function test_validateMessage_invalid() public {
    uint256 amount = 100;
    {
      // Case 1: Bad data.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      // data should be able to get decoded to an address using abi.decode.
      // encodePacked will not pad to 32 byte boundaries.
      bytes memory badData = abi.encodePacked(FINAL_RECEIVER);
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: badData,
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      vm.expectRevert();
      s_etherSenderReceiver.validateMessage(message);
    }
    {
      // Case 2: Receiver is 0x0.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      // data should be able to get decoded to an address using abi.decode.
      // encodePacked will not pad to 32 byte boundaries.
      bytes memory badData = abi.encode(address(0));
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: badData,
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      vm.expectRevert(
        abi.encodeWithSelector(InvalidDestinationReceiver.selector, badData)
      );
      s_etherSenderReceiver.validateMessage(message);
    }
    {
      // Case 3: incorrect token amounts.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      tokenAmounts[1] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      vm.expectRevert(
        abi.encodeWithSelector(InvalidTokenAmounts.selector, uint256(2))
      );
      s_etherSenderReceiver.validateMessage(message);
    }
    {
      // Case 4: incorrect token address in tokenAmounts.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_someOtherWeth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      vm.expectRevert(
        abi.encodeWithSelector(InvalidWethAddress.selector, address(s_weth), address(s_someOtherWeth))
      );
      s_etherSenderReceiver.validateMessage(message);
    }
    {
      // Case 5: gas specified in extra args is too low.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(s_weth),
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(FINAL_RECEIVER),
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({
          gasLimit: 100
        }))
      });

      vm.expectRevert(
        abi.encodeWithSelector(GasLimitTooLow.selector, uint256(200_000), uint256(100))
      );
      s_etherSenderReceiver.validateMessage(message);
    }
  }
}

// TODO: write tests.
