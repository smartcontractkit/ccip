// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {PingPongDemo} from "./PingPongDemo.sol";
import {IRouterWithOnRamps} from "./interfaces/IRouterWithOnRamps.sol";
import {Client} from "../libraries/Client.sol";
import {EVM2EVMOnRamp} from "../onRamp/EVM2EVMOnRamp.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";

contract SelfFundedPingPong is PingPongDemo {
  // Defines the increase in ping pong count before self-funding is attempted.
  // Set to 0 to disable auto-funding, auto-funding only works for ping-pongs that are set as NOPs in the onRamp.
  uint8 private s_countIncrBeforeFunding;

  event Funded();

  constructor(address router, IERC20 feeToken, uint8 roundTripsBeforeFunding) PingPongDemo(router, feeToken) {
    // PingPong count increases by 2 for each round trip.
    s_countIncrBeforeFunding = roundTripsBeforeFunding * 2;
  }

  function _respond(uint256 pingPongCount) internal override {
    if (pingPongCount & 1 == 1) {
      emit Ping(pingPongCount);
    } else {
      emit Pong(pingPongCount);
    }

    fundPingPong(pingPongCount);

    bytes memory data = abi.encode(pingPongCount);
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(getCounterpartAddress()),
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: "",
      feeToken: address(getFeeToken())
    });
    IRouterWithOnRamps(getRouter()).ccipSend(getCounterpartChainSelector(), message);
  }

  /// @notice A function that is responsible for funding this contract.
  /// The contract can only be funded if it is set as a nop in the target onRamp.
  /// In case your contract is not a nop you can prevent this function from being called by setting s_countIncrBeforeFunding=0.
  function fundPingPong(uint256 pingPongCount) public {
    // If selfFunding is disabled, or ping pong count has not reached s_countIncrPerFunding, do not attempt funding.
    if (s_countIncrBeforeFunding == 0 || pingPongCount < s_countIncrBeforeFunding) return;

    // Ping pong on one side will always be even, one side will always to odd.
    // Funding threshold is met if pingPongCount = (s_countIncrBeforeFunding * I) + (0 || 1)
    if (pingPongCount % s_countIncrBeforeFunding <= 1) {
      EVM2EVMOnRamp(IRouterWithOnRamps(getRouter()).getOnRamp(getCounterpartChainSelector())).payNops();
      emit Funded();
    }
  }

  function getCountIncrBeforeFunding() external view returns (uint8) {
    return s_countIncrBeforeFunding;
  }
}
