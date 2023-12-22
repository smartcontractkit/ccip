// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {OCR2Base} from "../ocr/OCR2Base.sol";
import {IPriceRegistry} from "../interfaces/IPriceRegistry.sol";
import {Internal} from "../libraries/Internal.sol";

contract TokenPriceOCR is ITypeAndVersion, OCR2Base {
  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "TokenPriceOCR 1.0.0";

  // DYNAMIC CONFIG
  /// @dev The price registry to write price updates to.
  IPriceRegistry internal s_priceRegistry;

  constructor(IPriceRegistry priceRegistry) OCR2Base(true) {
    s_priceRegistry = priceRegistry;
  }

  function _report(bytes calldata report, uint40 /* epochAndRound */) internal override {
    Internal.PriceUpdates memory priceUpdates = abi.decode(report, (Internal.PriceUpdates));
    s_priceRegistry.updatePrices(priceUpdates);
  }
}
