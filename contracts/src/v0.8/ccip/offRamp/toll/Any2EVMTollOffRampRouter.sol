// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../../vendor/SafeERC20.sol";
import "../../access/OwnerIsCreator.sol";
import "../BaseOffRampRouter.sol";

contract Any2EVMTollOffRampRouter is BaseOffRampRouter, TypeAndVersionInterface {
  string public constant override typeAndVersion = "Any2EVMTollOffRampRouter 1.0.0";

  constructor(BaseOffRampInterface[] memory offRamps) BaseOffRampRouter(offRamps) {}
}
