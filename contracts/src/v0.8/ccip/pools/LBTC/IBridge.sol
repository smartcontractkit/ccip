// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ILBTC} from "../LBTC/ILBTC.sol";
import "./adapters/IAdapter.sol";
import {IConsortiumConsumer, INotaryConsortium} from "./IConsortiumConsumer.sol";

interface IBridge is IConsortiumConsumer {
    /// @notice Emitted when the destination is unknown.
    error UnknownDestination();

    /// @notice Emitted when the zero address is used.
    error Bridge_ZeroAddress();

    error Bridge_ZeroAmount();

    /// @notice Emitted adapter is not set for destination without consortium
    error BadConfiguration();

    /// @notice Emitted when the destination is already known.
    error KnownDestination();

    /// @notice Emitted when the zero contract hash is used.
    error ZeroContractHash();

    /// @notice Emitted when the chain id is invalid.
    error ZeroChainId();

    /// @notice Emitted when the destination is not valid.
    error NotValidDestination();

    /// @notice Emitted when amount is below commission
    error AmountLessThanCommission(uint256 commission);

    /// @notice Emitted when the origin contract is unknown.
    error UnknownOriginContract(bytes32 fromChain, bytes32 fromContract);

    /// @notice Emitted when the unexpected action is used.
    error UnexpectedAction(bytes4 action);

    error UnknownAdapter(address);

    error PayloadAlreadyUsed(bytes32);

    /// @notice Emitted no payload submitted by adapter
    error AdapterNotConfirmed();

    /// @notice Emitted no payload submitted by consortium
    error ConsortiumNotConfirmed();

    /// @notice Emitted when the deposit absolute commission is changed.
    event DepositAbsoluteCommissionChanged(
        uint64 newValue,
        bytes32 indexed chain
    );

    /// @notice Emitted when the deposit relative commission is changed.
    event DepositRelativeCommissionChanged(
        uint16 newValue,
        bytes32 indexed chain
    );

    /// @notice Emitted when a bridge destination is added.
    event BridgeDestinationAdded(
        bytes32 indexed chain,
        bytes32 indexed contractAddress
    );

    /// @notice Emitted when a bridge destination is removed.
    event BridgeDestinationRemoved(bytes32 indexed chain);

    /// @notice Emitted when the adapter is changed.
    event AdapterChanged(address previousAdapter, IAdapter newAdapter);

    /// @notice Emitted when the is a deposit in the bridge
    event DepositToBridge(
        address indexed fromAddress,
        bytes32 indexed toAddress,
        bytes32 indexed payloadHash,
        bytes payload
    );

    /// @notice Emitted when a withdraw is made from the bridge
    event WithdrawFromBridge(
        address indexed recipient,
        bytes32 indexed payloadHash,
        bytes payload,
        uint64 amount
    );

    event PayloadReceived(
        address indexed recipient,
        bytes32 indexed payloadHash,
        address indexed adapter
    );

    event PayloadNotarized(
        address indexed recipient,
        bytes32 indexed payloadHash
    );

    /// @notice Emitted when the treasury is changed.
    event TreasuryChanged(address previousTreasury, address newTreasury);

    function lbtc() external view returns (ILBTC);
    function receivePayload(bytes32 fromChain, bytes calldata payload) external;
    function deposit(
        bytes32 toChain,
        bytes32 toAddress,
        uint64 amount
    ) external payable returns (uint256, bytes memory);
    function authNotary(bytes calldata payload, bytes calldata proof) external;
    function withdraw(bytes calldata payload) external returns (uint64);
}
