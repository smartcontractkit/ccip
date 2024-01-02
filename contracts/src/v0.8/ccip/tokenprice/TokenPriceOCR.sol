// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {OCR3Base} from "../ocr/OCR3Base.sol";
import {IPriceRegistry} from "../interfaces/IPriceRegistry.sol";
import {Internal} from "../libraries/Internal.sol";

contract TokenPriceOCR is ITypeAndVersion, OCR3Base {
  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "TokenPriceOCR 1.0.0";

  // DYNAMIC CONFIG
  /// @dev The price registry to write price updates to.
  IPriceRegistry internal s_priceRegistry;

  struct Report {
    Internal.TokenPriceUpdate[] priceUpdates;
  }

  mapping(address token => Internal.TimestampedPackedUint224) internal s_tokenPriceUpdates;
  mapping(uint64 destChainSelector => Internal.TimestampedPackedUint224) internal s_gasPriceUpdates;

  constructor(IPriceRegistry priceRegistry) OCR3Base() {
    s_priceRegistry = priceRegistry;
  }

  function setPriceRegistry(IPriceRegistry priceRegistry) external onlyOwner {
    s_priceRegistry = priceRegistry;
  }

  function _report(bytes calldata report, uint64 /* sequenceNumber */) internal override {
    Report memory decodedReport = abi.decode(report, (Report));
    s_priceRegistry.updatePrices(Internal.PriceUpdates({
      tokenPriceUpdates: decodedReport.priceUpdates,
      gasPriceUpdates: new Internal.GasPriceUpdate[](0)
    }));
  }

  function exposeForEncoding(Report memory report) external pure returns (bytes memory) {
    return abi.encode(report);
  }

  function getGasPriceUpdate(uint64 destChainSelector) external view returns (uint224 value, uint32 timestamp) {
    return (s_gasPriceUpdates[destChainSelector].value, s_gasPriceUpdates[destChainSelector].timestamp);
  }

  function getTokenPriceUpdate(address token) external view returns (uint224 value, uint32 timestamp) {
    return (s_tokenPriceUpdates[token].value, s_tokenPriceUpdates[token].timestamp);
  }
}
