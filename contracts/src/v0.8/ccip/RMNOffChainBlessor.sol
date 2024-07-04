// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {EnumerableSet} from "../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";
import {IAny2EVMMultiOffRamp} from "./interfaces/IAny2EVMMultiOffRamp.sol";
import {IRMNOffChainBlessor} from "./interfaces/IRMNOffChainBlessor.sol";

import {RMNBase} from "./RMNBase.sol";

/// @dev This contract is owned by RMN, if changing, please notify the RMN maintainers.
// solhint-disable chainlink-solidity/explicit-returns
contract RMNOffChainBlessor is IRMNOffChainBlessor, RMNBase {
  using EnumerableSet for EnumerableSet.AddressSet;

  error MustRecoverFromCurse();

  // STATIC CONFIG
  string public constant override typeAndVersion = "RMN 1.6.0-dev";

  constructor(Config memory config) RMNBase(config) {
    {
      assert(type(uint200).max >> (MAX_NUM_VOTERS - 1) >= 1);
    }
  }

  function _taggedRootHash(TaggedRoot memory taggedRoot) internal pure returns (bytes32) {
    return keccak256(abi.encode(taggedRoot.sourceChainSelector, taggedRoot.commitStore, taggedRoot.root));
  }

  /// @notice Performs signature verification for a merkle root + interval
  /// @param root The merkle root struct used in the OffRamp commit function
  function isBlessed(IAny2EVMMultiOffRamp.MerkleRoot calldata root) external view returns (bool) {
    // If we have an active curse, something is really wrong. Let's err on the
    // side of caution and not accept further blessings during this time of
    // uncertainty.
    // TODO: confirm if we should revert or return false here
    if (isCursed(bytes16(bytes8(root.sourceChainSelector)))) revert MustRecoverFromCurse();

    if (s_permaBlessedCommitStores.contains(msg.sender)) return true;

    uint128 signed; // Bitmap to track signers that have already signed
    uint32 configVersion = s_versionedConfig.configVersion;
    uint16 blessWeightThreshold = s_versionedConfig.config.blessWeightThreshold;
    uint16 accumulatedWeight;

    for (uint256 i = 0; i < root.rmnSignatures.length; ++i) {
      bytes memory signature = root.rmnSignatures[i];
      bytes32 r;
      bytes32 s;
      uint8 v;
      /// @solidity memory-safe-assembly
      assembly {
        r := mload(add(signature, 0x20))
        s := mload(add(signature, 0x40))
        v := byte(0, mload(add(signature, 0x60)))
      }

      // Safe from ECDSA malleability here since we check for duplicate signers.
      // Note: current assumption is that RMN nodes sign merkle root + interval
      address signer =
        ecrecover(keccak256(abi.encode(root.merkleRoot, root.interval.min, root.interval.max)), v + 27, r, s);
      BlesserRecord memory blesserRecord = s_blesserRecords[signer];

      // TODO: confirm if we should revert or return false here
      if (blesserRecord.configVersion != configVersion) revert UnauthorizedVoter(signer);

      // TODO: confirm if we should revert or skip here
      if (!_bitmapGet(signed, blesserRecord.index)) {
        _bitmapSet(signed, blesserRecord.index);
        accumulatedWeight += blesserRecord.weight;
      }

      if (accumulatedWeight >= blessWeightThreshold) return true;
    }

    return false;
  }
}
