// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "../BaseTest.t.sol";
import "../mocks/MockERC20.sol";
import "../mocks/MockAFN.sol";
import "../mocks/MockPool.sol";
import "../helpers/BlobVerifierHelper.sol";
import "../../../tests/MockV3Aggregator.sol";
import "../../utils/CCIP.sol";
import "../../offRamp/toll/Any2EVMTollOffRampRouter.sol";

contract BlobVerifierSetup is BaseTest {
  BlobVerifierHelper s_blobVerifier;
  BlobVerifierInterface.BlobVerifierConfig s_config;

  function setUp() public virtual override {
    BaseTest.setUp();

    address[] memory onRamps = new address[](3);
    onRamps[0] = s_onRampAddress;
    onRamps[1] = 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc;
    onRamps[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    uint64[] memory minSequenceNumbers = new uint64[](3);
    minSequenceNumbers[0] = 1;
    minSequenceNumbers[1] = 1;
    minSequenceNumbers[2] = 1;
    s_config = BlobVerifierInterface.BlobVerifierConfig({
      sourceChainId: s_sourceChainId,
      onRamps: onRamps,
      minSeqNrByOnRamp: minSequenceNumbers
    });
    s_blobVerifier = new BlobVerifierHelper(s_destChainId, s_afn, 1e18, s_config);
  }
}
