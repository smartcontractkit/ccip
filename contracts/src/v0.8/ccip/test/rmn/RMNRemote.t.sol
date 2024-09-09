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
  }
}
