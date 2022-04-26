// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../interfaces/OnRampInterface.sol";

contract MockOnRamp is OnRampInterface {
  uint256 public immutable CHAIN_ID;
  IERC20 public immutable TOKEN;
  IERC20 public immutable DESTINATION_TOKEN;
  PoolInterface public immutable POOL;
  uint256 public immutable DESTINATION_CHAIN_ID;

  CCIP.MessagePayload public mp;

  constructor(
    uint256 chainId,
    IERC20 token,
    IERC20 destinationToken,
    PoolInterface pool,
    uint256 destinationChainId
  ) {
    CHAIN_ID = chainId;
    TOKEN = token;
    DESTINATION_TOKEN = destinationToken;
    POOL = pool;
    DESTINATION_CHAIN_ID = destinationChainId;
  }

  function requestCrossChainSend(CCIP.MessagePayload calldata payload, address originalSender)
    external
    override
    returns (uint64)
  {
    mp = payload;
    return 0;
  }

  function getMessagePayload()
    external
    view
    returns (
      address receiver,
      bytes memory data,
      IERC20[] memory tokens,
      uint256[] memory amounts,
      bytes memory options
    )
  {
    receiver = mp.receiver;
    data = mp.data;
    tokens = mp.tokens;
    amounts = mp.amounts;
    options = mp.options;
  }
}
