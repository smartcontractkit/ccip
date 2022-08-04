// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../interfaces/BlobVerifierInterface.sol";
import "../interfaces/offRamp/Any2EVMOffRampInterface.sol";
import "../interfaces/onRamp/BaseOnRampInterface.sol";
import "../interfaces/subscription/SubscriptionInterface.sol";

contract StructFactory {
  // addresses
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant STRANGER = 0x1111111111111111111111111111111111111111;
  address internal constant ON_RAMP_ADDRESS = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;

  // message info
  uint256 internal constant SOURCE_CHAIN_ID = 1;
  uint256 internal constant DEST_CHAIN_ID = 2;
  uint256 internal constant GAS_LIMIT = 100_000;

  // timing
  uint256 internal constant BLOCK_TIME = 1234567890;
  uint256 internal constant HEARTBEAT = 1e18;

  // offRamp
  uint256 internal constant POOL_BALANCE = 5000;
  uint64 internal constant EXECUTION_DELAY_SECONDS = 0;
  uint64 internal constant MAX_DATA_SIZE = 500;
  uint64 internal constant MAX_TOKENS_LENGTH = 5;
  uint32 internal constant PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS = 500;

  function offRampConfig() internal pure returns (BaseOffRampInterface.OffRampConfig memory) {
    return
      BaseOffRampInterface.OffRampConfig({
        executionDelaySeconds: EXECUTION_DELAY_SECONDS,
        maxDataSize: MAX_DATA_SIZE,
        maxTokensLength: MAX_TOKENS_LENGTH,
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS
      });
  }

  // onRamp
  uint64 internal constant RELAYING_FEE_JUELS = 1e18;

  function onRampConfig() internal pure returns (BaseOnRampInterface.OnRampConfig memory) {
    return
      BaseOnRampInterface.OnRampConfig({
        relayingFeeJuels: RELAYING_FEE_JUELS,
        maxDataSize: MAX_DATA_SIZE,
        maxTokensLength: MAX_TOKENS_LENGTH
      });
  }

  // blob verifier
  function blobVerifierConfig() internal pure returns (BlobVerifierInterface.BlobVerifierConfig memory) {
    address[] memory onRamps = new address[](3);
    onRamps[0] = ON_RAMP_ADDRESS;
    onRamps[1] = 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc;
    onRamps[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    uint64[] memory minSequenceNumbers = new uint64[](3);
    minSequenceNumbers[0] = 1;
    minSequenceNumbers[1] = 1;
    minSequenceNumbers[2] = 1;
    return BlobVerifierInterface.BlobVerifierConfig({onRamps: onRamps, minSeqNrByOnRamp: minSequenceNumbers});
  }

  // subscription
  uint32 internal constant SET_SUBSCRIPTION_SENDER_DELAY = 2 * 60 * 60;
  uint32 internal constant WITHDRAWAL_DELAY = 2 * 60 * 60;

  function subscriptionConfig(IERC20 feeToken) internal pure returns (SubscriptionInterface.SubscriptionConfig memory) {
    return SubscriptionInterface.SubscriptionConfig(SET_SUBSCRIPTION_SENDER_DELAY, WITHDRAWAL_DELAY, feeToken);
  }
}
