// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 */

contract Storage {
    uint256 number;


    event storedNumber(
        address indexed _from,
        uint256 indexed _oldNumber,
        uint256 indexed _number
    );

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        uint256 old = number;
        number = num;
        emit storedNumber(msg.sender, old, num);

    }

    /**
     * @dev Return value
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256) {
        return number;

    }
}