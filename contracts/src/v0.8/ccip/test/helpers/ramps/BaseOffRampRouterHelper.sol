// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../../offRamp/BaseOffRampRouter.sol";

// Needed because BaseOffRampRouter is an abstract contract
contract BaseOffRampRouterHelper is BaseOffRampRouter {
  constructor(IBaseOffRamp[] memory offRamps) BaseOffRampRouter(offRamps) {}
}
