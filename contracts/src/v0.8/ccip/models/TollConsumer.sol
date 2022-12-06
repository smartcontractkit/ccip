pragma solidity ^0.8.0;

import {Common} from "./Common.sol";

library TollConsumer {
  /// @notice The Toll message type for EVM chains.
  struct EVM2AnyTollMessage {
    bytes receiver;
    bytes data;
    Common.EVMTokenAndAmount[] tokensAndAmounts;
    Common.EVMTokenAndAmount feeTokenAndAmount;
    bytes extraArgs;
  }

  // bytes4(keccak256("CCIP EVMExtraArgsV1")); // TODO: needs toll
  bytes4 public constant EVM_EXTRA_ARGS_V1_TAG = 0x97a657c9;
  struct EVMExtraArgsV1 {
    uint256 gasLimit;
    bool strict;
  }

  // TODO: call this serialize?
  function _argsToBytes(EVMExtraArgsV1 memory extraArgs) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(EVM_EXTRA_ARGS_V1_TAG, extraArgs);
  }
}
