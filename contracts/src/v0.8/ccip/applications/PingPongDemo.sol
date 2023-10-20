// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClientExtended} from "./interfaces/IRouterClientExtended.sol";
import {IEVM2AnyOnRampExtended} from "./interfaces/IEVM2AnyOnRampExtended.sol";
import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {Client} from "../libraries/Client.sol";
import {CCIPReceiver} from "./CCIPReceiver.sol";
import {EVM2EVMOnRamp} from "../onRamp/EVM2EVMOnRamp.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";

/// @title PingPongDemo - A simple ping-pong contract for demonstrating cross-chain communication
contract PingPongDemo is CCIPReceiver, OwnerIsCreator {
  event Ping(uint256 pingPongCount);
  event Pong(uint256 pingPongCount);

  // The chain ID of the counterpart ping pong contract
  uint64 private s_counterpartChainSelector;
  // The contract address of the counterpart ping pong contract
  address private s_counterpartAddress;

  // Pause ping-ponging
  bool private s_isPaused;
  IERC20 private s_feeToken;

  // number of ping-pongs till a call to the funding method 'fundPingPong' is made
  // note that 0 disables the funding.
  uint256 private s_fundingRounds = 5;

  constructor(address router, IERC20 feeToken) CCIPReceiver(router) {
    s_isPaused = false;
    s_feeToken = feeToken;
    s_feeToken.approve(address(router), 2 ** 256 - 1);
  }

  function setCounterpart(uint64 counterpartChainSelector, address counterpartAddress) external onlyOwner {
    s_counterpartChainSelector = counterpartChainSelector;
    s_counterpartAddress = counterpartAddress;
  }

  function startPingPong() external onlyOwner {
    s_isPaused = false;
    _respond(1);
  }

  function _respond(uint256 pingPongCount) private {
    if (pingPongCount & 1 == 1) {
      emit Ping(pingPongCount);
    } else {
      emit Pong(pingPongCount);
    }

    bytes memory data = abi.encode(pingPongCount);
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(s_counterpartAddress),
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000})),
      feeToken: address(s_feeToken)
    });
    IRouterClientExtended(getRouter()).ccipSend(s_counterpartChainSelector, message);

    if (s_fundingRounds > 0 && pingPongCount % s_fundingRounds == 0) {
      fundPingPong();
    }
  }

  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
    uint256 pingPongCount = abi.decode(message.data, (uint256));
    if (!s_isPaused) {
      _respond(pingPongCount + 1);
    }
  }

  /// @notice A function that is responsible for funding this contract.
  /// The contract can only be funded if it is set as a nop in the target onRamp.
  /// In case your contract is not a nop you can prevent this function from being called by setting s_fundingRounds=0.
  function fundPingPong() public {
    address onRampAddress = IRouterClientExtended(getRouter()).getOnRamp(s_counterpartChainSelector);

    // onRamp does not have anything to pay
    if (IEVM2AnyOnRampExtended(onRampAddress).getNopFeesJuels() == 0) {
      return;
    }

    // not enough link to fund the ping pong
    if (IEVM2AnyOnRampExtended(onRampAddress).linkAvailableForPayment() < 0) {
      return;
    }

    bool isNop = false;
    (EVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights, ) = IEVM2AnyOnRampExtended(onRampAddress).getNops();
    for (uint256 i = nopsAndWeights.length; i > 0; --i) {
      EVM2EVMOnRamp.NopAndWeight memory nopAndWeight = nopsAndWeights[i - 1];
      if (nopAndWeight.nop == address(this)) {
        isNop = true;
        break;
      }
    }

    if (isNop) {
      IEVM2AnyOnRampExtended(onRampAddress).payNops();
    }
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  function getCounterpartChainSelector() external view returns (uint64) {
    return s_counterpartChainSelector;
  }

  function setCounterpartChainSelector(uint64 chainSelector) external onlyOwner {
    s_counterpartChainSelector = chainSelector;
  }

  function getCounterpartAddress() external view returns (address) {
    return s_counterpartAddress;
  }

  function setCounterpartAddress(address addr) external onlyOwner {
    s_counterpartAddress = addr;
  }

  function isPaused() external view returns (bool) {
    return s_isPaused;
  }

  function setPaused(bool pause) external onlyOwner {
    s_isPaused = pause;
  }
}
