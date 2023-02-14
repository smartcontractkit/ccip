// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "./Common.sol";

// End consumer library.
library Client {
  struct EVM2AnyMessage {
    bytes receiver;
    bytes data;
    Common.EVMTokenAndAmount[] tokensAndAmounts;
    address feeToken;
    bytes extraArgs;
  }

  // bytes4(keccak256("CCIP EVMExtraArgsV1"));
  bytes4 public constant EVM_EXTRA_ARGS_V1_TAG = 0x97a657c9;
  struct EVMExtraArgsV1 {
    uint256 gasLimit;
    bool strict;
  }

  function _argsToBytes(EVMExtraArgsV1 memory extraArgs) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(EVM_EXTRA_ARGS_V1_TAG, extraArgs);
  }
}
