// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Types} from "../libraries/Types.sol";

/// @dev This is so that we can generate gethwrappers and easily encode/decode OCR3Config
/// in the offchain integration tests.
interface IOCR3ConfigEncoder {
  /// @dev Encodes an array of OCR3Config into a bytes array. For test usage only.
  function exposeOCR3Config(Types.OCR3Config[] calldata config) external view returns (bytes memory);
}
