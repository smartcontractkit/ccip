// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMNV2} from "../../interfaces/IRMNV2.sol";
import {Internal} from "../../libraries/Internal.sol";
import {RMNRemote, RMN_V1_6_ANY2EVM_REPORT} from "../../rmn/RMNRemote.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {Vm} from "forge-std/Vm.sol";

import "forge-std/console.sol";

contract RMNRemoteSetup is BaseTest {
  RMNRemote public s_rmnRemote;
  address public OFF_RAMP_ADDRESS;

  RMNRemote.Signer[] public s_signers;
  Vm.Wallet[] public s_signerWallets;

  function setUp() public virtual override {
    super.setUp();
    s_rmnRemote = new RMNRemote(1);
    OFF_RAMP_ADDRESS = makeAddr("OFF RAMP");

    _setupSigners(10);
  }

  /// @notice sets up a list of signers with strictly increasing onchain public keys
  /// @dev signers do not have to be in order when configured, but they do when generating signatures
  /// rather than sort signers every time, we do it once here and store the sorted list
  function _setupSigners(uint256 numSigners) internal {
    for (uint256 i = 0; i < numSigners; i++) {
      s_signerWallets.push(vm.createWallet(_randomNum()));
    }

    _sort(s_signerWallets);

    for (uint256 i = 0; i < numSigners; i++) {
      s_signers.push(RMNRemote.Signer({onchainPublicKey: s_signerWallets[i].addr, nodeIndex: uint64(i)}));
    }
  }

  /// @notice generates n destLaneUpdates and matching valid signatures and populates them into
  /// the provided storage arrays
  /// @dev if tests are running out of gas, try reducing the number of sigs generated
  /// @dev important note here that ONLY v=27 sigs are valid in the RMN contract. Because there is
  /// very little control over how these sigs are generated in foundry, we have to "get lucky" with the
  /// payload / signature combination. Therefore, we generate a payload and sigs together here in 1 function.
  /// If we can't generate valid (v=27 for all signers) sigs we re-generate the payload and try again.
  /// Warning: this is very annoying and clunky code. Tweak at your own risk.
  function _generatePayloadAndSigs(
    uint256 numUpdates,
    uint256 numSigs,
    Internal.MerkleRoot[] storage destLaneUpdates,
    IRMNV2.Signature[] storage signatures
  ) internal {
    require(numUpdates > 0, "need at least 1 dest lane update");
    require(numSigs <= s_signerWallets.length, "cannot generate more sigs than signers");

    // remove any existing updates and sigs
    for (uint256 i = 0; i < destLaneUpdates.length; i++) {
      destLaneUpdates.pop();
    }
    for (uint256 i = 0; i < signatures.length; i++) {
      signatures.pop();
    }

    for (uint256 i = 0; i < numUpdates; i++) {
      destLaneUpdates.push(_generateRandomDestLaneUpdate());
    }

    while (true) {
      bool allSigsValid = true;
      for (uint256 i = 0; i < numSigs; i++) {
        (bool isValid, IRMNV2.Signature memory sig) = _signDestLaneUpdate(destLaneUpdates, s_signerWallets[i]);
        signatures.push(sig);
        allSigsValid = allSigsValid && isValid;
        if (!allSigsValid) {
          break;
        }
      }
      // if all sigs are valid, don't change anything!!
      if (allSigsValid) {
        break;
      }
      // try again with a different payload if not all sigs are valid
      destLaneUpdates.pop();
      destLaneUpdates.push(_generateRandomDestLaneUpdate());
      // clear existing sigs
      while (signatures.length > 0) {
        signatures.pop();
      }
    }
  }

  /// @notice generates a random dest lane update
  function _generateRandomDestLaneUpdate() private returns (Internal.MerkleRoot memory) {
    uint64 minSeqNum = uint32(_randomNum());
    uint64 maxSeqNum = minSeqNum + 100;
    return Internal.MerkleRoot({
      sourceChainSelector: uint64(_randomNum()),
      onRampAddress: abi.encode(_randomAddress()),
      minSeqNr: minSeqNum,
      maxSeqNr: maxSeqNum,
      merkleRoot: _randomBytes32()
    });
  }

  /// @notice signs the provided payload with the provided wallet
  /// @return valid true only if the v component of the signature == 27
  /// @return sig the signature
  function _signDestLaneUpdate(
    Internal.MerkleRoot[] memory destLaneUpdates,
    Vm.Wallet memory wallet
  ) private returns (bool valid, IRMNV2.Signature memory) {
    bytes32 digest = keccak256(
      abi.encode(
        RMN_V1_6_ANY2EVM_REPORT,
        RMNRemote.Report({
          destChainId: block.chainid,
          destChainSelector: s_rmnRemote.getChainSelector(),
          rmnRemoteContractAddress: address(s_rmnRemote),
          offrampAddress: OFF_RAMP_ADDRESS,
          rmnHomeContractConfigDigest: s_rmnRemote.getVersionedConfig().config.rmnHomeContractConfigDigest,
          destLaneUpdates: destLaneUpdates
        })
      )
    );
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(wallet, digest);
    return (v == 27, IRMNV2.Signature({r: r, s: s})); // only v==27 sigs are valid in RMN contract
  }

  /// @notice bubble sort on a storage array of wallets
  function _sort(Vm.Wallet[] storage wallets) private {
    bool swapped;
    for (uint256 i = 1; i < wallets.length; i++) {
      swapped = false;
      for (uint256 j = 0; j < wallets.length - i; j++) {
        Vm.Wallet memory next = wallets[j + 1];
        Vm.Wallet memory actual = wallets[j];
        if (next.addr < actual.addr) {
          wallets[j] = next;
          wallets[j + 1] = actual;
          swapped = true;
        }
      }
      if (!swapped) {
        return;
      }
    }
  }
}
