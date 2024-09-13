pragma solidity ^0.8.0;

/// @title Deterministic Deployer Library
/// @notice Library for deterministic contract deployment.
library DeterministicContractDeployer {
  error DeploymentFailed();

  /// @notice Deploys a contract with a deterministically computed address.
  /// @param salt A value to modify the deployment address.
  /// @param initCode The bytecode of the contract to deploy.
  /// @return contractAddress The address of the deployed contract.
  function _deploy(bytes memory initCode, bytes32 salt) internal returns (address contractAddress) {
    assembly {
      // create2(value, memoryAddress, size, salt)
      contractAddress := create2(0, add(initCode, 0x20), mload(initCode), salt)
    }

    // When deploying with assembly, a revert will not occur if the deployment fails, so manual checks are needed
    if (contractAddress == address(0)) {
      revert DeploymentFailed();
    }

    return contractAddress;
  }

  /// @notice Predicts the address of a contract that would be deployed with create2 and the given parameters.
  /// @param initCode The bytecode of the contract to deploy, with constructor arguments already appended
  /// @param salt A value to used with create2 to result in a unique the deployment address.
  /// @param deployer The address of the account that will deploy the contract.
  /// @return address The predicted address of the contract.
  function _predictAddressOfUndeployedContract(
    bytes memory initCode,
    bytes32 salt,
    address deployer
  ) internal pure returns (address) {
    // Current EVM specs use following formula is used to decide contract addresses when deployed with create2
    // address = keccak256(0xff + sender_address + salt + keccak256(initialisation_code))[12:]
    bytes32 bytesValue = keccak256(abi.encodePacked(hex"ff", deployer, salt, keccak256(initCode)));

    // Return the left 20 bytes of the 32 byte hash, which is the address of the contract
    return address(uint160(uint256(bytesValue)));
  }
}
