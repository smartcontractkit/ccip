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
  event Funded();

  function test_FundingSuccess() public {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);
    changePrank(address(s_sourceRouter));

    for (uint256 pingPongNumber = 0; pingPongNumber <= 2 * s_roundTripsBeforeFunding; ++pingPongNumber) {
      Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
        messageId: bytes32("a"),
        sourceChainSelector: DEST_CHAIN_ID,
        sender: abi.encode(i_pongContract),
        data: abi.encode(pingPongNumber),
        destTokenAmounts: tokenAmounts
      });

      if (pingPongNumber == 2 * s_roundTripsBeforeFunding - 1) {
        vm.expectEmit();
        emit Funded();
      }

      s_pingPong.ccipReceive(message);
    }
  }

  function test_FundingFailure() public {
    EVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = new EVM2EVMOnRamp.NopAndWeight[](0);
    s_onRamp.setNops(nopsAndWeights);

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);
    changePrank(address(s_sourceRouter));

    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: bytes32("a"),
      sourceChainSelector: DEST_CHAIN_ID,
      sender: abi.encode(i_pongContract),
      data: abi.encode(2 * s_roundTripsBeforeFunding - 1),
      destTokenAmounts: tokenAmounts
    });

    vm.expectRevert(); // because pingPong is not set as a nop
    s_pingPong.ccipReceive(message);
  }
}
