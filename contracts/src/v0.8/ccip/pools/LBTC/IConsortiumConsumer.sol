// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {INotaryConsortium} from "./INotaryConsortium.sol";

/**
 * @title Consortium Consumer interface
 * @author Lombard.Finance
 * @notice Common interface for contracts who verify signatures with `INotaryConsortium`
 */
interface IConsortiumConsumer {
    event ConsortiumChanged(
        INotaryConsortium indexed prevVal,
        INotaryConsortium indexed newVal
    );

    function changeConsortium(INotaryConsortium newVal) external;
    function consortium() external view returns (INotaryConsortium);
}
