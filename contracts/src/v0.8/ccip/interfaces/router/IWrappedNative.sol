pragma solidity ^0.8.0;

import {IERC20} from "../../../vendor/IERC20.sol";

interface IWrappedNative is IERC20 {
  function deposit() external payable;
}
