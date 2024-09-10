// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMNV2} from "../../interfaces/IRMNV2.sol";
import {Internal} from "../../libraries/Internal.sol";
import {RMNRemote} from "../../rmn/RMNRemote.sol";
import {RMNRemoteSetup} from "./RMNRemoteSetup.t.sol";

contract RMNRemote_constructor is RMNRemoteSetup {
  function test_constructor_success() public {
    assertEq(s_rmnRemote.getChainSelector(), 1);
  }
}

contract RMNRemote_setConfig is RMNRemoteSetup {
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
    uint256 numSigners = s_signers.length;
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    RMNRemote.VersionedConfig memory versionedConfig = s_rmnRemote.getVersionedConfig();
    // add a signer
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(1), nodeIndex: uint64(numSigners)}));
    config = RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    versionedConfig = s_rmnRemote.getVersionedConfig();
    assertEq(versionedConfig.config.signers.length, numSigners + 1);
    assertEq(versionedConfig.config.signers[numSigners].onchainPublicKey, address(1));
    // remove signers
    s_signers.pop();
    s_signers.pop();
    config = RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});
    s_rmnRemote.setConfig(config);
    versionedConfig = s_rmnRemote.getVersionedConfig();
    assertEq(versionedConfig.config.signers.length, numSigners - 1);
  }

  function test_setConfig_invalidSignerOrder_reverts() public {
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(4), nodeIndex: 0}));
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
    s_signers.push(RMNRemote.Signer({onchainPublicKey: s_signerWallets[0].addr, nodeIndex: uint64(s_signers.length)}));
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});

    vm.expectRevert(RMNRemote.DuplicateOnchainPublicKey.selector);
    s_rmnRemote.setConfig(config);
  }
}

contract RMNRemote_verify_withConfigNotSet is RMNRemoteSetup {
  function test_verify_reverts() public {
    Internal.MerkleRoot[] memory destLaneUpdates = new Internal.MerkleRoot[](0);
    IRMNV2.Signature[] memory signatures = new IRMNV2.Signature[](0);

    vm.expectRevert(RMNRemote.ConfigNotSet.selector);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, destLaneUpdates, signatures);
  }
}

contract RMNRemote_verify_withConfigSet is RMNRemoteSetup {
  Internal.MerkleRoot[] s_destLaneUpdates;
  IRMNV2.Signature[] s_signatures;

  function setUp() public override {
    super.setUp();
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 2});
    s_rmnRemote.setConfig(config);
    _generatePayloadAndSigs(2, 2, s_destLaneUpdates, s_signatures);
  }

  function test_verify_success() public {
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, s_signatures);
  }

  function test_verify_minSignersIsZero_success() public {
    vm.stopPrank();
    vm.prank(OWNER);
    s_rmnRemote.setConfig(
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0})
    );

    vm.stopPrank();
    vm.prank(OFF_RAMP_ADDRESS);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, new IRMNV2.Signature[](0));
  }

  function test_verify_invalidSig_reverts() public {
    IRMNV2.Signature memory sig = s_signatures[s_signatures.length - 1];
    sig.r = _randomBytes32();
    s_signatures.pop();
    s_signatures.push(sig);

    vm.expectRevert(RMNRemote.InvalidSignature.selector);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, s_signatures);
  }

  function test_verify_outOfOrderSig_reverts() public {
    IRMNV2.Signature memory sig1 = s_signatures[s_signatures.length - 1];
    s_signatures.pop();
    IRMNV2.Signature memory sig2 = s_signatures[s_signatures.length - 1];
    s_signatures.pop();
    s_signatures.push(sig1);
    s_signatures.push(sig2);

    vm.expectRevert(RMNRemote.OutOfOrderSignatures.selector);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, s_signatures);
  }

  function test_verify_duplicateSignature_reverts() public {
    IRMNV2.Signature memory sig = s_signatures[s_signatures.length - 2];
    s_signatures.pop();
    s_signatures.push(sig);

    vm.expectRevert(RMNRemote.OutOfOrderSignatures.selector);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, s_signatures);
  }

  function test_verify_unknownSigner_reverts() public {
    _setupSigners(2); // create 2 new signers that aren't configured on RMNRemote
    _generatePayloadAndSigs(2, 2, s_destLaneUpdates, s_signatures);

    vm.expectRevert(RMNRemote.UnexpectedSigner.selector);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, s_signatures);
  }

  function test_verify_insufficientSignatures_reverts() public {
    _generatePayloadAndSigs(2, 1, s_destLaneUpdates, s_signatures); // 1 sig requested, but 2 required

    vm.expectRevert(RMNRemote.ThresholdNotMet.selector);
    s_rmnRemote.verify(OFF_RAMP_ADDRESS, s_destLaneUpdates, s_signatures);
  }
}
