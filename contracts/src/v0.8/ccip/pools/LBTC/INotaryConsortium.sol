// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ECDSA} from "../../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/cryptography/ECDSA.sol";

interface INotaryConsortium {
    /// @dev Error thrown when signature payload is already used
    error PayloadAlreadyUsed();

    /// @dev Error thrown when signatures length is not equal to signers length
    error LengthMismatch();

    /// @dev Error thrown when there are not enough signatures
    error NotEnoughSignatures();

    /// @dev Error thrown when ECDSA signature verification fails
    error SignatureVerificationFailed(uint256 sigIndx, ECDSA.RecoverError err);

    /// @dev Error thrown when signature verification fails
    error WrongSignatureReceived(bytes sig);

    /// @dev Error thrown when unexpected action is used
    error UnexpectedAction(bytes4 action);

    /// @dev Event emitted when the validator set is updated
    event ValidatorSetUpdated(
        uint256 indexed epoch,
        address[] validators,
        uint256[] weights,
        uint256 threshold
    );

    /// @dev Error thrown when validator set already set
    error ValSetAlreadySet();

    /// @dev Error thrown when no validator set is set
    error NoValidatorSet();

    /// @dev Error thrown when invalid epoch is provided
    error InvalidEpoch();

    function checkProof(
        bytes32 _payloadHash,
        bytes calldata _proof
    ) external view;
}
