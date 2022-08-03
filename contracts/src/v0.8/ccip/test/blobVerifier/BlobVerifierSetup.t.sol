// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import "../helpers/BlobVerifierHelper.sol";
import "../../utils/CCIP.sol";

contract BlobVerifierSetup is BaseTest {
  event BlobVerifierConfigSet(BlobVerifierInterface.BlobVerifierConfig config);

  BlobVerifierHelper s_blobVerifier;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_blobVerifier = new BlobVerifierHelper(DEST_CHAIN_ID, SOURCE_CHAIN_ID, s_afn, HEARTBEAT, blobVerifierConfig());
  }
}
