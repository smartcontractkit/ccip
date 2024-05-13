pragma solidity ^0.8.0;

import {Client} from "../../../libraries/Client.sol";
import {IRouter, IRouterClient, MockCCIPRouter} from "../MockRouter.sol";

import {Test} from "forge-std/Test.sol";

contract MockRouterTest is Test {
  MockCCIPRouter public mockRouter;

  uint64 public constant mockChainSelector = 123456;

  Client.EVM2AnyMessage public message;

  function setUp() public {
    mockRouter = new MockCCIPRouter();

    //Configure the Fee to 0.1 ether for native token fees
    mockRouter.setFee(0.1 ether);

    deal(address(this), 100 ether);

    message.receiver = abi.encode(address(0x12345));
    message.data = abi.encode("Hello World");
  }

  function test_ccipSendWithInsufficientNativeTokens_Revert() public {
    //Should revert because did not include sufficient eth to pay for fees
    vm.expectRevert(IRouterClient.InsufficientFeeTokenAmount.selector);
    mockRouter.ccipSend(mockChainSelector, message);
  }

  function test_ccipSendWithSufficientNativeFeeTokens_Success() public {
    //ccipSend with sufficient native tokens for fees
    mockRouter.ccipSend{value: 0.1 ether}(mockChainSelector, message);
  }

  function test_ccipSendWithInvalidMsgValue_Revert() public {
    message.feeToken = address(1); //Set to non native-token fees

    vm.expectRevert(IRouterClient.InvalidMsgValue.selector);
    mockRouter.ccipSend{value: 0.1 ether}(mockChainSelector, message);
  }

  function test_ccipSendWithValidMsgValueAndNonNativeFeeToken_Success() public {
    message.feeToken = address(1); //Set to non native-token fees

    mockRouter.ccipSend(mockChainSelector, message);
  }
}
