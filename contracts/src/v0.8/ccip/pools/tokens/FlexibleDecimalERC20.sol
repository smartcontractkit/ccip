// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {ERC20} from "../../../vendor/ERC20.sol";

contract FlexibleDecimalERC20 is ERC20 {
    uint8 immutable i_decimals;

    constructor(string memory name, string memory symbol, uint8 decimals_) ERC20(name, symbol){
        i_decimals = decimals_;
    }

    function decimals() public view virtual override returns (uint8) {
        return i_decimals;
    }
}
