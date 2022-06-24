// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/TollOffRampInterface.sol";

contract MockOffRamp is TollOffRampInterface {
  IERC20 public s_token;

  function deliverMessageTo(CrossChainMessageReceiverInterface recipient, CCIP.Any2EVMTollMessage calldata message)
    external
  {
    recipient.ccipReceive(message);
  }

  function SOURCE_CHAIN_ID() external view returns (uint256) {}

  function CHAIN_ID() external view returns (uint256) {}

  function setRouter(TollOffRampRouterInterface router) external {}

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature/
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external override {
    revert();
  }

  function execute(CCIP.ExecutionReport memory report, bool needFee)
    external
    override
    returns (CCIP.ExecutionResult[] memory)
  {
    CCIP.ExecutionResult[] memory results = new CCIP.ExecutionResult[](0);
    return results;
  }

  function executeSingleMessage(CCIP.Any2EVMTollMessage memory message) external {}

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }
}
