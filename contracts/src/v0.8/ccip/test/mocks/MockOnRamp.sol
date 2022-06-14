// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../onRamp/interfaces/TollOnRampInterface.sol";

contract MockOnRamp is TollOnRampInterface {
  uint256 public immutable CHAIN_ID;
  PoolInterface public immutable POOL;
  uint256 public immutable DESTINATION_CHAIN_ID;
  uint256 public immutable FEE;

  CCIP.EVM2AnyTollMessage public mp;

  event GetRequiredFee(IERC20 token);
  event GetTokenPool(IERC20 token);

  constructor(
    uint256 chainId,
    PoolInterface pool,
    uint256 destinationChainId,
    uint256 fee
  ) {
    CHAIN_ID = chainId;
    POOL = pool;
    DESTINATION_CHAIN_ID = destinationChainId;
    FEE = fee;
  }

  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address originalSender)
    external
    override
    returns (uint64)
  {
    mp = message;
    return 0;
  }

  function getMessagePayload()
    external
    view
    returns (
      address receiver,
      bytes memory data,
      IERC20[] memory tokens,
      uint256[] memory amounts
    )
  {
    receiver = mp.receiver;
    data = mp.data;
    tokens = mp.tokens;
    amounts = mp.amounts;
  }

  function getRequiredFee(IERC20 token) external override returns (uint256) {
    emit GetRequiredFee(token);
    return FEE;
  }

  function getSequenceNumber() external view returns (uint64) {
    return 1;
  }

  function getTokenPool(IERC20 token) external returns (PoolInterface) {
    return POOL;
  }
}
