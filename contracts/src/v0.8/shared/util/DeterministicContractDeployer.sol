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

    if(contractAddress == address(0)) {
      revert DeploymentFailed();
    }

    return contractAddress;
  }

    function _predictAddressOfUndeployedContract(
    bytes memory initCode,
    bytes32 salt,
    address deployer
  ) internal pure returns (address) {
    
    // according to evm.codes, the below formula is used to predict the address of a contract
    // address = keccak256(0xff + sender_address + salt + keccak256(initialisation_code))[12:]
    bytes32 bytesValue = keccak256(abi.encodePacked(hex"ff", deployer, salt, keccak256(initCode)));

    return address(uint160(uint256(bytesValue)));
  }

}