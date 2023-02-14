// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library Common {
  struct EVMTokenAndAmount {
    address token;
    uint256 amount;
  }

  struct Any2EVMMessage {
    uint64 sourceChainId;
    bytes sender;
    bytes data;
    Common.EVMTokenAndAmount[] destTokensAndAmounts;
  }
}
