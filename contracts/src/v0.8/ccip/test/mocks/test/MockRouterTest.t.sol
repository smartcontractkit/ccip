pragma solidity ^0.8.0;

import {Client} from "../../../libraries/Client.sol";
import {MockCCIPRouter, IRouter, IRouterClient} from "../MockRouter.sol";

import {Test} from "forge-std/Test.sol";

contract MockRouterTest is Test {
  MockCCIPRouter public mockRouter;

  function setUp() public {
    mockRouter = new MockCCIPRouter();

    //Configure the Fee to 0.1 ether for native token fees
    mockRouter.setFee(0.1 ether);

    deal(address(this), 100 ether);
  }

  function test_ccipSendWithInvalidNativeTokens_Revert() public {
    //Message with Native Token as Fee
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(address(0x12345)),
      data: abi.encode("Hello World"),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: address(0),
      extraArgs: ""
    });

    uint64 mockChainSelector = 123456;

    //Should revert because did not include sufficient eth to pay for fees
    vm.expectRevert(IRouterClient.InsufficientFeeTokenAmount.selector);
    mockRouter.ccipSend(mockChainSelector, message);

    //ccipSend with sufficient native tokens for fees
    mockRouter.ccipSend{value: 0.1 ether}(mockChainSelector, message);

    message.feeToken = address(1); //Set feeToken to something other than native asset
    //Should revert because msg.value should be zero when feeToken is not native asset;
    vm.expectRevert(IRouterClient.InvalidMsgValue.selector);
    mockRouter.ccipSend{value: 0.1 ether}(mockChainSelector, message);

    //Test a non-zero feeToken with zero msg.value. Call should succeed.
    mockRouter.ccipSend(mockChainSelector, message);
  }
}
