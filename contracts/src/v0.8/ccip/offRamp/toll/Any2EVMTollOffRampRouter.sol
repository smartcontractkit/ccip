// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {BaseOffRampRouter, BaseOffRampInterface} from "../BaseOffRampRouter.sol";

contract Any2EVMTollOffRampRouter is BaseOffRampRouter, TypeAndVersionInterface {
  string public constant override typeAndVersion = "Any2EVMTollOffRampRouter 1.0.0";

  constructor(BaseOffRampInterface[] memory offRamps) BaseOffRampRouter(offRamps) {}
}
