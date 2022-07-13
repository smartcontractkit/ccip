// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import "../helpers/BlobVerifierHelper.sol";
import "../../utils/CCIP.sol";

contract BlobVerifierSetup is BaseTest {
  event BlobVerifierConfigSet(BlobVerifierInterface.BlobVerifierConfig config);

  BlobVerifierHelper s_blobVerifier;
  BlobVerifierInterface.BlobVerifierConfig s_config;

  function setUp() public virtual override {
    BaseTest.setUp();

    address[] memory onRamps = new address[](3);
    onRamps[0] = ON_RAMP_ADDRESS;
    onRamps[1] = 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc;
    onRamps[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    uint64[] memory minSequenceNumbers = new uint64[](3);
    minSequenceNumbers[0] = 1;
    minSequenceNumbers[1] = 1;
    minSequenceNumbers[2] = 1;
    s_config = BlobVerifierInterface.BlobVerifierConfig({onRamps: onRamps, minSeqNrByOnRamp: minSequenceNumbers});
    s_blobVerifier = new BlobVerifierHelper(DEST_CHAIN_ID, SOURCE_CHAIN_ID, s_afn, HEARTBEAT, s_config);
  }
}
