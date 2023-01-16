// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/offRamp/IBaseOffRamp.sol";
import "../../interfaces/applications/IAny2EVMMessageReceiver.sol";
import "../../interfaces/ICommitStore.sol";
import "../../models/Toll.sol";
import "../../models/TollConsumer.sol";

contract MockOffRamp is IBaseOffRamp {
  IERC20 public s_token;

  function deliverMessageTo(IAny2EVMMessageReceiver recipient, Common.Any2EVMMessage calldata message) external {
    recipient.ccipReceive(message);
  }

  function i_sourceChainId() external view returns (uint256) {}

  function i_chainId() external view returns (uint256) {}

  function setRouter(IAny2EVMOffRampRouter router) external {}

  function getRouter() external pure override returns (IAny2EVMOffRampRouter) {
    return IAny2EVMOffRampRouter(address(0));
  }

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature/
   */
  function ccipReceive(Common.Any2EVMMessage calldata) external pure {
    revert();
  }

  function executeSingleMessage(Toll.EVM2EVMTollMessage memory message) external {}

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }

  /// @inheritdoc IBaseOffRamp
  function getExecutionState(uint64) public pure returns (Internal.MessageExecutionState) {
    return Internal.MessageExecutionState.SUCCESS;
  }

  /// @inheritdoc IBaseOffRamp
  function getCommitStore() public pure returns (ICommitStore) {
    return ICommitStore(address(1));
  }

  /// @inheritdoc IBaseOffRamp
  function setCommitStore(ICommitStore commitStore) public pure {}

  function getConfig() public pure returns (OffRampConfig memory config) {}

  function setConfig(OffRampConfig memory config) public {}
}
