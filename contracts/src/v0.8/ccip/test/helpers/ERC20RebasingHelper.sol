// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/ERC20.sol";

contract ERC20RebasingHelper is ERC20 {
  uint16 public s_multiplierPercentage = 100;

  constructor() ERC20("Rebasing", "REB") {}

  function mint(address to, uint256 amount) external {
    _mint(to, amount * s_multiplierPercentage / 100);
  }

  function setMultiplierPercentage(uint16 multiplierPercentage) external {
    s_multiplierPercentage = multiplierPercentage;
  }
}
