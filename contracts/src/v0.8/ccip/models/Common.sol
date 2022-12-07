pragma solidity ^0.8.0;

library Common {
  struct EVMTokenAndAmount {
    address token;
    uint256 amount;
  }
  // TODO: Do we want to maintain the guarantee that different message types
  // always map to the same receiving interface? I'm not so sure... refund details for example
  // might be different
  struct Any2EVMMessage {
    uint64 sourceChainId;
    bytes sender;
    bytes data;
    Common.EVMTokenAndAmount[] destTokensAndAmounts;
  }
}
