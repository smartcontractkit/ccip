// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {BaseOffRampRouter, IBaseOffRamp} from "../BaseOffRampRouter.sol";

contract Any2EVMTollOffRampRouter is BaseOffRampRouter, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "Any2EVMTollOffRampRouter 1.0.0";

  constructor(IBaseOffRamp[] memory offRamps) BaseOffRampRouter(offRamps) {}
}
