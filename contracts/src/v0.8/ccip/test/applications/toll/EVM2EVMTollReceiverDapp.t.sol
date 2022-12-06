// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../applications/ReceiverDapp.sol";
import "../../mocks/MockTollOffRampRouter.sol";

// setup
contract ReceiverDappSetup is TokenSetup {
  ReceiverDapp s_receiverDapp;
  MockTollOffRampRouter s_mockRouter;
  IERC20 s_feeToken;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_feeToken = IERC20(s_destFeeToken);

    s_mockRouter = new MockTollOffRampRouter();
    s_receiverDapp = new ReceiverDapp(s_mockRouter);

    IERC20(s_destTokens[0]).transfer(address(s_receiverDapp), 2**64);
    IERC20(s_destTokens[1]).transfer(address(s_receiverDapp), 2**64);
  }
}

/// @notice #constructor
contract EVM2EVMTollReceiverDapp_constructor is ReceiverDappSetup {
  // Success
  function testSuccess() public {
    // typeAndVersion
    assertEq("ReceiverDapp 1.0.0", s_receiverDapp.typeAndVersion());
  }
}

/// @notice #ccipReceive
contract EVM2EVMTollReceiverDapp_ccipReceive is ReceiverDappSetup {
  // Success

  function testSuccess() public {
    Common.Any2EVMMessage memory message;
    uint256 transferAmount = 5000;
    message.destTokensAndAmounts = getCastedDestinationEVMTokenAndAmountsWithZeroAmounts();
    message.destTokensAndAmounts[0].amount = transferAmount;
    message.data = abi.encode(OWNER, OWNER);

    uint256 startingBalanceOwner = s_feeToken.balanceOf(OWNER);
    uint256 startingBalanceContract = s_feeToken.balanceOf(address(s_receiverDapp));

    changePrank(address(s_mockRouter));

    s_receiverDapp.ccipReceive(message);

    assertEq(transferAmount, s_feeToken.balanceOf(OWNER) - startingBalanceOwner);
    assertEq(transferAmount, startingBalanceContract - s_feeToken.balanceOf(address(s_receiverDapp)));
  }

  // Revert

  function testInvalidDelivererReverts() public {
    vm.expectRevert(abi.encodeWithSelector(ReceiverDapp.InvalidDeliverer.selector, OWNER));
    Common.Any2EVMMessage memory message;

    s_receiverDapp.ccipReceive(message);
  }
}
