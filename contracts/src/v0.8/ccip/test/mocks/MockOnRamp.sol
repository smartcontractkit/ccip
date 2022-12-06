// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/onRamp/EVM2EVMTollOnRampInterface.sol";
import "../../interfaces/pools/PoolInterface.sol";
import "../../models/Common.sol";
import "../../models/TollConsumer.sol";

contract MockOnRamp is EVM2EVMTollOnRampInterface {
  uint256 public immutable i_chainId;
  PoolInterface public immutable i_pool;
  uint256 public immutable i_destinationChainId;
  uint256 public immutable i_fee;

  bytes public messageReceiver;
  bytes public messageData;
  bytes public messageTokens;

  event GetRequiredFee(IERC20 token);
  event GetTokenPool(IERC20 token);

  constructor(
    uint256 chainId,
    PoolInterface pool,
    uint256 destinationChainId,
    uint256 fee
  ) {
    i_chainId = chainId;
    i_pool = pool;
    i_destinationChainId = destinationChainId;
    i_fee = fee;
  }

  function forwardFromRouter(TollConsumer.EVM2AnyTollMessage memory message, address)
    external
    override
    returns (uint64)
  {
    messageReceiver = message.receiver;
    messageData = message.data;
    messageTokens = abi.encode(message.tokensAndAmounts);
    return 0;
  }

  function getMessagePayload()
    external
    view
    returns (
      bytes memory receiver,
      bytes memory data,
      Common.EVMTokenAndAmount[] memory tokensAndAmounts
    )
  {
    receiver = messageReceiver;
    data = messageData;
    tokensAndAmounts = abi.decode(messageTokens, (Common.EVMTokenAndAmount[]));
  }

  function getRequiredFee(IERC20 token) external override returns (uint256) {
    emit GetRequiredFee(token);
    return i_fee;
  }

  function getExpectedNextSequenceNumber() external pure returns (uint64) {
    return 1;
  }

  function getPoolBySourceToken(IERC20) external view returns (PoolInterface) {
    return i_pool;
  }

  function setRouter(address) external override {}

  function getRouter() external pure override returns (address) {
    return address(0);
  }

  function setConfig(OnRampConfig calldata) external override {}

  function getConfig() external pure override returns (OnRampConfig memory config) {
    config = OnRampConfig({commitFeeJuels: 0, maxDataSize: 0, maxTokensLength: 0, maxGasLimit: 0});
  }

  function setAllowlistEnabled(bool) external override {}

  function getAllowlistEnabled() external pure override returns (bool) {
    return true;
  }

  function setAllowlist(address[] calldata allowlist) external override {}

  function getAllowlist() external pure override returns (address[] memory addresses) {
    addresses = new address[](0);
  }

  function setFeeConfig(FeeConfig calldata feeConfig) external override {}

  function getPoolTokens() public pure returns (IERC20[] memory) {
    return new IERC20[](0);
  }
}
