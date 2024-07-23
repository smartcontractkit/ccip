// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.19;

/// @notice a contract which acts as a forwarder that forwards the input from
/// any caller to a a target contract.
contract CallProxy {
    event TargetSet(address target);

    address immutable i_target;

    constructor(address target) {
        i_target = target;
        emit TargetSet(target);
    }

    fallback() external payable {
        address target = i_target;
        assembly {
            // This code destroys Solidity's memory layout.
            // That's fine, because we never return to Solidity anyways,
            // we either return or revert out of the callframe at the end.
            calldatacopy(0, 0, calldatasize())
            let success := call(gas(), target, callvalue(), 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            if success { return(0, returndatasize()) }
            revert(0, returndatasize())
        }
    }
}