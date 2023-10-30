// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../../applications/SelfFundedPingPong.sol";
import "../onRamp/EVM2EVMOnRampSetup.t.sol";
import "../../libraries/Client.sol";

contract SelfFundedPingPongDappSetup is EVM2EVMOnRampSetup {
  event Ping(uint256 pingPongs);
  event Pong(uint256 pingPongs);

  SelfFundedPingPong internal s_pingPong;
  IERC20 internal s_feeToken;
  uint8 internal s_roundTripsBeforeFunding = 3;

  address internal immutable i_pongContract = address(10);

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_feeToken = IERC20(s_sourceTokens[0]);
    s_pingPong = new SelfFundedPingPong(address(s_sourceRouter), s_feeToken, s_roundTripsBeforeFunding);
    s_pingPong.setCounterpart(DEST_CHAIN_ID, i_pongContract);

    uint256 fundingAmount = 5e18;

    // set ping pong as an onRamp nop to make sure that funding runs
    EVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = new EVM2EVMOnRamp.NopAndWeight[](1);
    nopsAndWeights[0] = EVM2EVMOnRamp.NopAndWeight({nop: address(s_pingPong), weight: 1});
    s_onRamp.setNops(nopsAndWeights);

    // Fund the contract with LINK tokens
    s_feeToken.transfer(address(s_pingPong), fundingAmount);
  }
}

/// @notice #ccipReceive
contract SelfFundedPingPong_funding is SelfFundedPingPongDappSetup {
  function testFunding() public {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);
    changePrank(address(s_sourceRouter));

    uint256 pingPongBalance = IERC20(s_feeToken).balanceOf(address(s_pingPong));

    for (uint256 pingPongNumber = 0; pingPongNumber < 2 * s_roundTripsBeforeFunding; ++pingPongNumber) {
      Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
        messageId: bytes32("a"),
        sourceChainSelector: DEST_CHAIN_ID,
        sender: abi.encode(i_pongContract),
        data: abi.encode(pingPongNumber),
        destTokenAmounts: tokenAmounts
      });
      s_pingPong.ccipReceive(message);

      uint256 currentPingPongBalance = IERC20(s_feeToken).balanceOf(address(s_pingPong));
      if (pingPongNumber == 2 * s_roundTripsBeforeFunding - 1) {
        require(currentPingPongBalance > pingPongBalance, "ping pong was funded");
      } else {
        require(currentPingPongBalance < pingPongBalance, "funding is not made yet");
      }
      pingPongBalance = currentPingPongBalance;
    }
  }
}
