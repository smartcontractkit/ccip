// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../onRamp/interfaces/Any2EVMTollOnRampInterface.sol";

contract MockOnRamp is Any2EVMTollOnRampInterface {
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

  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address) external override returns (uint64) {
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

  function getExpectedNextSequenceNumber() external pure returns (uint64) {
    return 1;
  }

  function getTokenPool(IERC20) external view returns (PoolInterface) {
    return POOL;
  }

  function setRouter(address) external override {}

  function getRouter() external pure override returns (address) {
    return address(0);
  }

  function setConfig(OnRampConfig calldata) external override {}

  function getConfig() external pure override returns (OnRampConfig memory config) {
    config = OnRampConfig({relayingFeeJuels: 0, maxDataSize: 0, maxTokensLength: 0});
  }

  function setAllowlistEnabled(bool) external override {}

  function getAllowlistEnabled() external pure override returns (bool) {
    return true;
  }

  function setAllowlist(address[] calldata allowlist) external override {}

  function getAllowlist() external pure override returns (address[] memory addresses) {
    addresses = new address[](0);
  }
}
