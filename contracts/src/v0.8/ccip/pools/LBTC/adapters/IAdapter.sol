// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IBridge} from "../IBridge.sol";

interface IAdapter {
    /// @notice Thrown when msg.value is not enough to pay CCIP fee.
    error NotEnoughToPayFee(uint256 fee);

    event ExecutionGasLimitSet(uint128 indexed prevVal, uint128 indexed newVal);

    function bridge() external view returns (IBridge);
    function getFee(
        bytes32 _toChain,
        bytes32 _toContract,
        bytes32 _toAddress,
        uint256 _amount,
        bytes memory _payload
    ) external view returns (uint256);
    function deposit(
        address _fromAddress,
        bytes32 _toChain,
        bytes32 _toContract,
        bytes32 _toAddress,
        uint256 _amount,
        bytes memory _payload
    ) external payable;
}
