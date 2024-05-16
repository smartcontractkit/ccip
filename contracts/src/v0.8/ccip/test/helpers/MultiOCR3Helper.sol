// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {MultiOCR3Base} from "../../ocr/MultiOCR3Base.sol";

contract MultiOCR3Helper is MultiOCR3Base {
  function transmit(
    uint8 ocrPluginType,
    bytes32[3] calldata reportContext,
    bytes calldata report,
    bytes32[] calldata rs,
    bytes32[] calldata ss,
    bytes32 rawVs
  ) external {
    _transmit(ocrPluginType, reportContext, report, rs, ss, rawVs);
  }

  function typeAndVersion() public pure override returns (string memory) {
    return "MultiOCR3BaseHelper 1.0.0";
  }
}
