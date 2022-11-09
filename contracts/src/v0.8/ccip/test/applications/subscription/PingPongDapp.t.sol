// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../onRamp/subscription/EVM2EVMSubscriptionOnRampSetup.t.sol";
import "../../../applications/PingPongDemo.sol";

// setup
contract PingPongDappSetup is EVM2EVMSubscriptionOnRampSetup {
  event Ping(uint256 pingPongs);
  event Pong(uint256 pingPongs);

  PingPongDemo s_pingPong;
  IERC20 s_feeToken;

  Any2EVMOffRampRouterInterface internal s_receivingRouter;
  address immutable i_pongContract = address(10);

  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();

    s_receivingRouter = Any2EVMOffRampRouterInterface(address(100));

    s_feeToken = IERC20(s_sourceTokens[0]);
    s_pingPong = new PingPongDemo(
      CCIPRouterInterface(address(s_receivingRouter)),
      CCIPRouterInterface(address(s_onRampRouter))
    );
    s_pingPong.setCounterpart(DEST_CHAIN_ID, i_pongContract);
  }
}

/// @notice #startPingPong
contract PingPong_startPingPong is PingPongDappSetup {
  event ConfigPropagated(uint256 chainId, address contractAddress);

  // Success
  function testSuccess() public {
    uint256 pingPongNumber = 1;

    bytes memory data = abi.encode(pingPongNumber);
    CCIP.EVM2EVMSubscriptionMessage memory subscriptionMsg = CCIP.EVM2EVMSubscriptionMessage({
      sequenceNumber: 1,
      sourceChainId: SOURCE_CHAIN_ID,
      sender: address(s_pingPong),
      receiver: i_pongContract,
      nonce: 1,
      data: data,
      tokensAndAmounts: new CCIP.EVMTokenAndAmount[](0),
      gasLimit: 2e5
    });

    vm.expectEmit(false, false, false, true);
    emit Ping(pingPongNumber);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(subscriptionMsg);

    s_pingPong.startPingPong();
  }
}

/// @notice #ccipReceive
contract PingPong_ccipReceive is PingPongDappSetup {
  // Success

  function testSuccess() public {
    EVMTokenAndAmount[] memory tokensAndAmounts = new EVMTokenAndAmount[](0);

    uint256 pingPongNumber = 5;

    CCIPReceiverInterface.ReceivedMessage memory message = CCIPReceiverInterface.ReceivedMessage({
      sourceChainId: DEST_CHAIN_ID,
      sender: abi.encode(i_pongContract),
      data: abi.encode(pingPongNumber),
      tokensAndAmounts: tokensAndAmounts
    });

    changePrank(address(s_receivingRouter));

    vm.expectEmit(false, false, false, true);
    emit Pong(pingPongNumber + 1);

    s_pingPong.ccipReceive(message);
  }
  // Revert
}
