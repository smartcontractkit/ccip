// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Test} from "forge-std/Test.sol";

import {CCIPRouter} from "../../applications/EtherSenderReceiver.sol";
import {EtherSenderReceiverHelper} from "./../helpers/EtherSenderReceiverHelper.sol";
import {Client} from "../../libraries/Client.sol";
import {IRouterClient} from "../../interfaces/IRouterClient.sol";
import {WETH9} from "../WETH9.sol";

import {ERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/ERC20.sol";

contract EtherSenderReceiverTest is Test {
  EtherSenderReceiverHelper internal s_etherSenderReceiver;
  WETH9 internal s_weth;
  WETH9 internal s_someOtherWeth;
  ERC20 internal s_linkToken;

  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant ROUTER = 0x0F3779ee3a832D10158073ae2F5e61ac7FBBF880;
  address internal constant XCHAIN_RECEIVER = 0xBd91b2073218AF872BF73b65e2e5950ea356d147;

  function setUp() public {
    vm.startPrank(OWNER);

    s_linkToken = new ERC20("Chainlink Token", "LINK");
    s_someOtherWeth = new WETH9();
    s_weth = new WETH9();
    vm.mockCall(ROUTER, abi.encodeWithSelector(CCIPRouter.getWrappedNative.selector), abi.encode(address(s_weth)));
    s_etherSenderReceiver = new EtherSenderReceiverHelper(ROUTER);

    deal(OWNER, 100 ether);
    deal(address(s_linkToken), OWNER, 100 ether);

    // deposit some eth into the weth contract.
    s_weth.deposit{value: 10 ether}();
    uint256 wethSupply = s_weth.totalSupply();
    assertEq(wethSupply, 10 ether, "total weth supply must be 10 ether");
  }
}

contract EtherSenderReceiverTest_constructor is EtherSenderReceiverTest {
  function test_constructor() public {
    assertEq(s_etherSenderReceiver.getRouter(), ROUTER, "router must be set correctly");
    uint256 allowance = s_weth.allowance(address(s_etherSenderReceiver), ROUTER);
    assertEq(allowance, type(uint256).max, "allowance must be set infinite");
  }
}

contract EtherSenderReceiverTest_validateFeeToken is EtherSenderReceiverTest {
  function test_validateFeeToken_valid() public {
    uint256 amount = 100;
    {
      // Case 1: feeToken is address(0), i.e native.
      Client.EVMTokenAmount[] memory tokenAmount = new Client.EVMTokenAmount[](1);
      tokenAmount[0] = Client.EVMTokenAmount({token: address(s_weth), amount: amount});
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmount,
        feeToken: address(0),
        extraArgs: ""
      });

      s_etherSenderReceiver.validateFeeToken{value: amount + 1}(message);
    }
    {
      // Case 2: feeToken is a nonzero address.
      Client.EVMTokenAmount[] memory tokenAmount = new Client.EVMTokenAmount[](1);
      tokenAmount[0] = Client.EVMTokenAmount({token: address(s_weth), amount: amount});
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
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

  function test_validatedMessage_validMessage() public {
    uint256 amount = 100;
    {
      // Case 1: data not specified, is overwritten to be msg.sender.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);
      assertEq(validatedMessage.receiver, abi.encode(XCHAIN_RECEIVER), "receiver must be XCHAIN_RECEIVER");
      assertEq(validatedMessage.data, abi.encode(OWNER), "data must be msg.sender");
      assertEq(validatedMessage.tokenAmounts[0].token, address(s_weth), "token must be weth");
      assertEq(validatedMessage.tokenAmounts[0].amount, amount, "amount must be correct");
      assertEq(validatedMessage.feeToken, address(0), "feeToken must be 0");
      assertEq(validatedMessage.extraArgs, bytes(""), "extraArgs must be empty");
    }
    {
      // Case 2: data specified, is still overwritten to be msg.sender.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: abi.encode(address(42)),
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);
      assertEq(validatedMessage.receiver, abi.encode(XCHAIN_RECEIVER), "receiver must be XCHAIN_RECEIVER");
      assertEq(validatedMessage.data, abi.encode(OWNER), "data must be msg.sender");
      assertEq(validatedMessage.tokenAmounts[0].token, address(s_weth), "token must be weth");
      assertEq(validatedMessage.tokenAmounts[0].amount, amount, "amount must be correct");
      assertEq(validatedMessage.feeToken, address(0), "feeToken must be 0");
      assertEq(validatedMessage.extraArgs, bytes(""), "extraArgs must be empty");
    }
    {
      // Case 3: token specified incorrectly, is overwritten to be weth.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(42), // incorrect token.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);
      assertEq(validatedMessage.receiver, abi.encode(XCHAIN_RECEIVER), "receiver must be XCHAIN_RECEIVER");
      assertEq(validatedMessage.data, abi.encode(OWNER), "data must be msg.sender");
      assertEq(validatedMessage.tokenAmounts[0].token, address(s_weth), "token must be weth");
      assertEq(validatedMessage.tokenAmounts[0].amount, amount, "amount must be correct");
      assertEq(validatedMessage.feeToken, address(0), "feeToken must be 0");
      assertEq(validatedMessage.extraArgs, bytes(""), "extraArgs must be empty");
    }
    {
      // Case 4: extraArgs specified.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000}))
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);
      assertEq(validatedMessage.receiver, abi.encode(XCHAIN_RECEIVER), "receiver must be XCHAIN_RECEIVER");
      assertEq(validatedMessage.data, abi.encode(OWNER), "data must be msg.sender");
      assertEq(validatedMessage.tokenAmounts[0].token, address(s_weth), "token must be weth");
      assertEq(validatedMessage.tokenAmounts[0].amount, amount, "amount must be correct");
      assertEq(validatedMessage.feeToken, address(0), "feeToken must be 0");
      assertEq(
        validatedMessage.extraArgs,
        Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000})),
        "extraArgs must be correct"
      );
    }
  }

  function test_validatedMessage_invalid() public {
    uint256 amount = 100;
    {
      // Case 1: incorrect token amounts.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
      tokenAmounts[0] = Client.EVMTokenAmount({token: address(0), amount: amount});
      tokenAmounts[1] = Client.EVMTokenAmount({token: address(0), amount: amount});
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      vm.expectRevert(abi.encodeWithSelector(InvalidTokenAmounts.selector, uint256(2)));
      s_etherSenderReceiver.validatedMessage(message);
    }
    {
      // Case 2: gas specified in extra args is too low.
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({token: address(0), amount: amount});
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 100}))
      });

      vm.expectRevert(abi.encodeWithSelector(GasLimitTooLow.selector, uint256(200_000), uint256(100)));
      s_etherSenderReceiver.validatedMessage(message);
    }
  }
}

contract EtherSenderReceiverTest_getFee is EtherSenderReceiverTest {}

contract EtherSenderReceiverTest_ccipReceive is EtherSenderReceiverTest {}

contract EtherSenderReceiverTest_ccipSend is EtherSenderReceiverTest {
  error InsufficientFee(uint256 gotFee, uint256 fee);

  uint256 internal constant amount = 100;
  uint64 internal constant destinationChainSelector = 424242;
  uint256 internal constant feeWei = 121212;
  uint256 internal constant feeJuels = 232323;

  function test_ccipSend_reverts_insufficientFee_weth() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(s_weth),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeWei)
      );

      s_weth.approve(address(s_etherSenderReceiver), feeWei - 1);

      vm.expectRevert("SafeERC20: low-level call failed");
      s_etherSenderReceiver.ccipSend{value: amount}(destinationChainSelector, message);
    }
  }

  function test_ccipSend_reverts_insufficientFee_feeToken() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(s_linkToken),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeJuels)
      );

      s_linkToken.approve(address(s_etherSenderReceiver), feeJuels - 1);

      vm.expectRevert("ERC20: insufficient allowance");
      s_etherSenderReceiver.ccipSend{value: amount}(destinationChainSelector, message);
    }
  }

  function test_ccipSend_reverts_insufficientFee_native() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeWei)
      );

      vm.expectRevert(abi.encodeWithSelector(InsufficientFee.selector, feeWei - 1, feeWei));
      s_etherSenderReceiver.ccipSend{value: amount + feeWei - 1}(destinationChainSelector, message);
    }
  }

  function test_ccipSend_success_nativeExcess() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      bytes32 expectedMsgId = keccak256(abi.encode("ccip send"));
      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeWei)
      );

      // we assert that the correct value is sent to the router call, which should be
      // the msg.value - feeWei.
      vm.mockCall(
        ROUTER,
        feeWei + 1,
        abi.encodeWithSelector(IRouterClient.ccipSend.selector, destinationChainSelector, validatedMessage),
        abi.encode(expectedMsgId)
      );

      bytes32 actualMsgId = s_etherSenderReceiver.ccipSend{value: amount + feeWei + 1}(
        destinationChainSelector,
        message
      );
      assertEq(actualMsgId, expectedMsgId, "message id must be correct");
    }
  }

  function test_ccipSend_success_native() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      bytes32 expectedMsgId = keccak256(abi.encode("ccip send"));
      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeWei)
      );
      vm.mockCall(
        ROUTER,
        feeWei,
        abi.encodeWithSelector(IRouterClient.ccipSend.selector, destinationChainSelector, validatedMessage),
        abi.encode(expectedMsgId)
      );

      bytes32 actualMsgId = s_etherSenderReceiver.ccipSend{value: amount + feeWei}(destinationChainSelector, message);
      assertEq(actualMsgId, expectedMsgId, "message id must be correct");
    }
  }

  function test_ccipSend_success_feeToken() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(s_linkToken),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      bytes32 expectedMsgId = keccak256(abi.encode("ccip send"));
      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeJuels)
      );
      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.ccipSend.selector, destinationChainSelector, validatedMessage),
        abi.encode(expectedMsgId)
      );

      s_linkToken.approve(address(s_etherSenderReceiver), feeJuels);

      bytes32 actualMsgId = s_etherSenderReceiver.ccipSend{value: amount}(destinationChainSelector, message);
      assertEq(actualMsgId, expectedMsgId, "message id must be correct");
      uint256 routerAllowance = s_linkToken.allowance(address(s_etherSenderReceiver), ROUTER);
      assertEq(routerAllowance, feeJuels, "router allowance must be feeJuels");
    }
  }

  function test_ccipSend_success_weth() public {
    {
      Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
      tokenAmounts[0] = Client.EVMTokenAmount({
        token: address(0), // callers may not specify this.
        amount: amount
      });
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(XCHAIN_RECEIVER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(s_weth),
        extraArgs: ""
      });

      Client.EVM2AnyMessage memory validatedMessage = s_etherSenderReceiver.validatedMessage(message);

      bytes32 expectedMsgId = keccak256(abi.encode("ccip send"));
      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.getFee.selector, destinationChainSelector, validatedMessage),
        abi.encode(feeWei)
      );
      vm.mockCall(
        ROUTER,
        abi.encodeWithSelector(IRouterClient.ccipSend.selector, destinationChainSelector, validatedMessage),
        abi.encode(expectedMsgId)
      );

      s_weth.approve(address(s_etherSenderReceiver), feeWei);

      bytes32 actualMsgId = s_etherSenderReceiver.ccipSend{value: amount}(destinationChainSelector, message);
      assertEq(actualMsgId, expectedMsgId, "message id must be correct");
      uint256 routerAllowance = s_weth.allowance(address(s_etherSenderReceiver), ROUTER);
      assertEq(routerAllowance, type(uint256).max, "router allowance must be max for weth");
    }
  }
}
