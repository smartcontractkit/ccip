// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.18;

contract SimpleContract {
    event SimpleEvent(uint256 indexed value);
    uint256 public eventCount;

    function emitEvent() public {
        eventCount++;
        emit SimpleEvent(eventCount+1);
    }

    function getEventCount() public view returns (uint256) {
        return eventCount;
    }
}
