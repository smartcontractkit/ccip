// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice This library contains various token pool functions to aid constructing the return data.
library Pool {
  // The tag used to signal support for the pool v1 standard
  // bytes4(keccak256("CCIP_POOL_V1"))
  bytes4 public constant CCIP_POOL_V1 = 0xaff2afbf;

  struct LockOrBurnInV1 {
    address originalSender;
    bytes receiver;
    uint256 amount;
    uint64 remoteChainSelector;
  }

  struct LockOrBurnOutV1 {
    bytes destPoolAddress;
    bytes destPoolData;
  }

  struct ReleaseOrMintInV1 {
    bytes originalSender;
    address receiver;
    uint256 amount;
    uint64 remoteChainSelector;
    bytes sourcePoolAddress;
    bytes sourcePoolData;
    bytes offchainTokenData;
  }

  struct ReleaseOrMintOutV1 {
    address localToken;
    uint256 destinationAmount;
  }

  function _encodeLockOrBurnInV1(
    address originalSender,
    bytes memory receiver,
    uint256 amount,
    uint64 remoteChainSelector
  ) internal pure returns (bytes memory) {
    return abi.encode(
      LockOrBurnInV1({
        originalSender: originalSender,
        receiver: receiver,
        amount: amount,
        remoteChainSelector: remoteChainSelector
      })
    );
  }

  function _decodeLockOrBurnInV1(bytes memory encodedData) internal pure returns (LockOrBurnInV1 memory) {
    return abi.decode(encodedData, (LockOrBurnInV1));
  }

  ///  @notice Generates the return dataV1 for the lockOrBurn pool call.
  ///  @param remotePoolAddress The address of the remote pool.
  ///  @param destPoolData The data to send to the remote pool.
  ///  @return The return data for the burnOrMint pool call.
  function _encodeLockOrBurnOutV1(
    bytes memory remotePoolAddress,
    bytes memory destPoolData
  ) internal pure returns (bytes memory) {
    return abi.encode(LockOrBurnOutV1({destPoolAddress: remotePoolAddress, destPoolData: destPoolData}));
  }

  /// @notice Decodes the PoolReturnDataV1 struct from the given data. Also checks if the tag is correct.
  /// @param encodedData The data to decode.
  /// @dev Can revert. Since this is only used on the sending side, this is acceptable.
  /// @return The decoded LockOrBurnOutV1 struct.
  function _decodeLockOrBurnOutV1(bytes memory encodedData) internal pure returns (LockOrBurnOutV1 memory) {
    return abi.decode(encodedData, (LockOrBurnOutV1));
  }

  function _encodeReleaseOrMintInV1(
    bytes memory originalSender,
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector,
    bytes memory sourcePoolAddress,
    bytes memory sourcePoolData,
    bytes memory offchainTokenData
  ) internal pure returns (bytes memory) {
    return abi.encode(
      ReleaseOrMintInV1({
        originalSender: originalSender,
        receiver: receiver,
        amount: amount,
        remoteChainSelector: remoteChainSelector,
        sourcePoolAddress: sourcePoolAddress,
        sourcePoolData: sourcePoolData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  function _decodeReleaseOrMintInV1(bytes memory encodedData) internal pure returns (ReleaseOrMintInV1 memory) {
    return abi.decode(encodedData, (ReleaseOrMintInV1));
  }

  function _encodeReleaseOrMintOutV1(
    address localToken,
    uint256 destinationAmount
  ) internal pure returns (bytes memory) {
    return abi.encode(ReleaseOrMintOutV1({localToken: localToken, destinationAmount: destinationAmount}));
  }

  function _decodeReleaseOrMintOutV1(bytes memory encodedData) internal pure returns (ReleaseOrMintOutV1 memory) {
    return abi.decode(encodedData, (ReleaseOrMintOutV1));
  }
}
