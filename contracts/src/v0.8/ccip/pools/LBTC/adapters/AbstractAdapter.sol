// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IAdapter} from "./IAdapter.sol";
import {Context} from "../../../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/Context.sol";
import {IBridge, ILBTC} from "../IBridge.sol";
/**
 * @title Abstract bridge adapter
 * @author Lombard.finance
 * @notice Implements basic communication with Bridge contract.
 * Should be extended with business logic of bridging protocols (e.g. CCIP, LayerZero).
 */
abstract contract AbstractAdapter is IAdapter, Context {
    error Adapter_ZeroAddress();

    error NotBridge();

    event BridgeChanged(IBridge indexed oldBridge, IBridge indexed newBridge);

    IBridge public override bridge;

    constructor(IBridge bridge_) {
        bridge = bridge_;
    }

    function lbtc() public view returns (ILBTC) {
        return bridge.lbtc();
    }

    /// MODIFIERS ///

    modifier onlyBridge() {
        _onlyBridge();
        _;
    }

    /// ONLY OWNER FUNCTIONS ///

    /**
     * @notice Change the bridge address
     * @param bridge_ New bridge address
     */
    function changeBridge(IBridge bridge_) external {
        _onlyOwner();
        _notZero(address(bridge_));

        IBridge oldBridge = bridge;
        bridge = bridge_;
        emit BridgeChanged(oldBridge, bridge_);
    }

    /// PRIVATE FUNCTIONS ///

    function _onlyOwner() internal view virtual;

    function _onlyBridge() internal view {
        if (_msgSender() != address(bridge)) {
            revert NotBridge();
        }
    }

    function _notZero(address addr) internal pure {
        if (addr == address(0)) {
            revert Adapter_ZeroAddress();
        }
    }

    /**
     * @dev Called when data is received.
     */
    function _receive(bytes32 fromChain, bytes memory payload) internal {
        bridge.receivePayload(fromChain, payload);
    }

    /**
     * @notice Sends a payload from the source to destination chain.
     * @param _toChain Destination chain's.
     * @param _payload The payload to send.
     * @param _refundAddress Address where refund fee
     */
    function _deposit(
        bytes32 _toChain,
        bytes memory _payload,
        address _refundAddress
    ) internal virtual {}

    function deposit(
        address _fromAddress,
        bytes32 _toChain,
        bytes32 /* _toContract */,
        bytes32 /* _toAddress */,
        uint256 /* _amount */,
        bytes memory _payload
    ) external payable virtual override {
        _deposit(_toChain, _payload, _fromAddress);
    }
}
