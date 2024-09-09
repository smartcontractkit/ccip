// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMNV2} from "../../interfaces/IRMNV2.sol";
import {Internal} from "../../libraries/Internal.sol";
import {RMNRemote, RMN_V1_6_ANY2EVM_REPORT} from "../../rmn/RMNRemote.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {Vm} from "forge-std/Vm.sol";

contract RMNRemoteTest is BaseTest {
  RMNRemote public s_rmnRemote;
  address public OFF_RAMP_ADDRESS;

  Vm.Wallet signer1;
  Vm.Wallet signer2;
  Vm.Wallet signer3;
  RMNRemote.Signer[] public s_signers;
  Vm.Wallet[] public s_signerWallets;

  function setUp() public virtual override {
    super.setUp();
    s_rmnRemote = new RMNRemote(1);
    OFF_RAMP_ADDRESS = makeAddr("OFF RAMP");

    signer1 = vm.createWallet("signer wallet 1");
    signer2 = vm.createWallet("signer wallet 2!");
    signer3 = vm.createWallet("signer wallet 3!!"); // "!" are added experimentally to ensure the order of signers
    require(signer1.addr < signer2.addr && signer2.addr < signer3.addr, "signers not in order");
    s_signers.push(RMNRemote.Signer({onchainPublicKey: signer1.addr, nodeIndex: 0}));
    s_signers.push(RMNRemote.Signer({onchainPublicKey: signer2.addr, nodeIndex: 1}));
    s_signers.push(RMNRemote.Signer({onchainPublicKey: signer3.addr, nodeIndex: 2}));
    s_signerWallets.push(signer1);
    s_signerWallets.push(signer2);
    s_signerWallets.push(signer3);
  }

  /// @notice generates n destLaneUpdates and matching valid signatures and populates them into
  /// the provided storage arrays
  /// @dev important note here that ONLY v=27 sigs are valid in the RMN contract. Because there is
  /// very little control over how these sigs are generated in foundry, we have to "get lucky" with the
  /// payload / signature combination. Therefore, we generate a payload and sigs together here in 1 function.
  /// If we can't generate valid (v=27 for all signers) sigs we just tweak the payload and try again.
  /// Warning: this is very annoying and clunky code. Tweak at your own risk.
  function _generatePayloadAndSigs(
    uint256 n,
    Internal.MerkleRoot[] storage destLaneUpdates,
    IRMNV2.Signature[] storage signatures
  ) internal {
    require(n > 0, "need at least 1 dest lane update");
    require(destLaneUpdates.length == 0, "storage array should be empty");
    require(signatures.length == 0, "storage array should be empty");

    for (uint256 i = 0; i < n; i++) {
      destLaneUpdates.push(_randomDestLaneUpdate());
    }

    while (true) {
      bool allValid = true;
      for (uint256 i = 0; i < s_signerWallets.length; i++) {
        (bool valid, IRMNV2.Signature memory sig) = _generateSig(destLaneUpdates, s_signerWallets[i]);
        signatures.push(sig);
        allValid = allValid && valid;
      }
      // if all sigs are valid, don't change anything!!
      if (allValid) {
        break;
      }
      // try again with a different payload if not all sigs are valid
      destLaneUpdates.pop();
      destLaneUpdates.push(_randomDestLaneUpdate());
      // clear existing sigs
      while (signatures.length > 0) {
        signatures.pop();
      }
    }
  }

  function _randomDestLaneUpdate() internal returns (Internal.MerkleRoot memory) {
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

  function _generateSig(
    Internal.MerkleRoot[] memory destLaneUpdates,
    Vm.Wallet memory wallet
  ) internal returns (bool valid, IRMNV2.Signature memory) {
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
}

contract RMNRemote_constructor is RMNRemoteTest {
  function test_constructor_success() public {
    assertEq(s_rmnRemote.getChainSelector(), 1);
  }
}

contract RMNRemote_setConfig is RMNRemoteTest {
  function test_setConfig_minSignersIs0_success() public {
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    RMNRemote.VersionedConfig memory versionedConfig = s_rmnRemote.getVersionedConfig();
    assertEq(versionedConfig.config.minSigners, 0);
  }

  function test_setConfig_versionIncreases_success() public {
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});

    vm.expectEmit();
    emit RMNRemote.ConfigSet(RMNRemote.VersionedConfig({version: 1, config: config}));
    s_rmnRemote.setConfig(config);
    assertEq(s_rmnRemote.getVersionedConfig().version, 1);

    vm.expectEmit();
    emit RMNRemote.ConfigSet(RMNRemote.VersionedConfig({version: 2, config: config}));
    s_rmnRemote.setConfig(config);
    assertEq(s_rmnRemote.getVersionedConfig().version, 2);

    vm.expectEmit();
    emit RMNRemote.ConfigSet(RMNRemote.VersionedConfig({version: 3, config: config}));
    s_rmnRemote.setConfig(config);
    assertEq(s_rmnRemote.getVersionedConfig().version, 3);
  }

  function test_setConfig_addSigner_removeSigner_success() public {
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    RMNRemote.VersionedConfig memory versionedConfig = s_rmnRemote.getVersionedConfig();
    // add a signer
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(4), nodeIndex: 3}));
    config = RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    versionedConfig = s_rmnRemote.getVersionedConfig();
    assertEq(versionedConfig.config.signers.length, 4);
    assertEq(versionedConfig.config.signers[3].onchainPublicKey, address(4));
    // remove signers
    s_signers.pop();
    s_signers.pop();
    config = RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    versionedConfig = s_rmnRemote.getVersionedConfig();
    assertEq(versionedConfig.config.signers.length, 2);
  }

  function test_setConfig_invalidSignerOrder_reverts() public {
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(4), nodeIndex: 4}));
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(5), nodeIndex: 3}));
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});

    vm.expectRevert(RMNRemote.InvalidSignerOrder.selector);
    s_rmnRemote.setConfig(config);
  }

  function test_setConfig_minSignersTooHigh_reverts() public {
    RMNRemote.Config memory config = RMNRemote.Config({
      rmnHomeContractConfigDigest: _randomBytes32(),
      signers: s_signers,
      minSigners: uint64(s_signers.length + 1)
    });

    vm.expectRevert(RMNRemote.MinSignersTooHigh.selector);
    s_rmnRemote.setConfig(config);
  }

  function test_setConfig_duplicateOnChainPublicKey_reverts() public {
    s_signers.push(RMNRemote.Signer({onchainPublicKey: s_signerWallets[0].addr, nodeIndex: 3}));
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});

    vm.expectRevert(RMNRemote.DuplicateOnchainPublicKey.selector);
    s_rmnRemote.setConfig(config);
  }
}

contract RMNRemote_verify_withConfigNotSet is RMNRemoteTest {
  function test_verify_reverts() public {
    Internal.MerkleRoot[] memory destLaneUpdates = new Internal.MerkleRoot[](0);
    IRMNV2.Signature[] memory signatures = new IRMNV2.Signature[](0);

    vm.expectRevert(RMNRemote.ConfigNotSet.selector);
    s_rmnRemote.verify(destLaneUpdates, signatures);
  }
}

contract RMNRemote_verify_withConfigSet is RMNRemoteTest {
  Internal.MerkleRoot[] s_destLaneUpdates;
  IRMNV2.Signature[] s_signatures;

  function setUp() public override {
    super.setUp();
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 2});
    s_rmnRemote.setConfig(config);
    _generatePayloadAndSigs(2, s_destLaneUpdates, s_signatures);
  }

  function test_verify_success() public {
    vm.stopPrank();
    vm.prank(OFF_RAMP_ADDRESS);
    s_rmnRemote.verify(s_destLaneUpdates, s_signatures);
  }

  function test_verify_minSignersIsZero_success() public {
    s_rmnRemote.setConfig(
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0})
    );

    s_rmnRemote.verify(s_destLaneUpdates, new IRMNV2.Signature[](0));
  }
}
