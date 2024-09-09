// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {RMNRemote} from "../../rmn/RMNRemote.sol";
import {BaseTest} from "../BaseTest.t.sol";

contract RMNRemoteTest is BaseTest {
  RMNRemote public s_rmnRemote;
  RMNRemote.Signer[] public s_signers;

  function setUp() public virtual override {
    super.setUp();
    s_rmnRemote = new RMNRemote(1);
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(1), nodeIndex: 0}));
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(2), nodeIndex: 1}));
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(3), nodeIndex: 2}));
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
    s_signers.push(RMNRemote.Signer({onchainPublicKey: address(2), nodeIndex: 3}));
    RMNRemote.Config memory config =
      RMNRemote.Config({rmnHomeContractConfigDigest: _randomBytes32(), signers: s_signers, minSigners: 0});

    vm.expectRevert(RMNRemote.DuplicateOnchainPublicKey.selector);
    s_rmnRemote.setConfig(config);
  }
}
