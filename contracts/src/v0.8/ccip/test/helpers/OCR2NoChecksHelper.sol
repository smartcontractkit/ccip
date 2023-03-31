// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OCR2BaseNoChecks} from "../../ocr/OCR2BaseNoChecks.sol";

contract OCR2NoChecksHelper is OCR2BaseNoChecks {
  function configDigestFromConfigData(
    uint256 chainId,
    address contractAddress,
    uint64 configCount,
    address[] memory signers,
    address[] memory transmitters,
    uint8 f,
    bytes memory onchainConfig,
    uint64 offchainConfigVersion,
    bytes memory offchainConfig
  ) public pure returns (bytes32) {
    return
      _configDigestFromConfigData(
        chainId,
        contractAddress,
        configCount,
        signers,
        transmitters,
        f,
        onchainConfig,
        offchainConfigVersion,
        offchainConfig
      );
  }

  function _report(bytes memory report) internal override {}

  function typeAndVersion() public pure override returns (string memory) {
    return "OCR2BaseHelper 1.0.0";
  }
}
