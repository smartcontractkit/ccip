// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OCR2Base} from "../../ocr/OCR2Base.sol";

contract OCR2Helper is OCR2Base {
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

  function _payTransmitter(uint256 initialGas, address transmitter) internal override {}

  function typeAndVersion() public pure override returns (string memory) {
    return "OCR2BaseHelper 1.0.0";
  }

  function _beforeSetOCR2Config(uint8 f, bytes memory onchainConfig) internal override {}
}
