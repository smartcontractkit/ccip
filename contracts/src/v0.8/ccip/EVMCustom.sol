// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// https://docs.moonbeam.network/builders/pallets-precompiles/precompiles/eth-mainnet/#hashing-with-sha256
contract EVMCustomTest {
    
    //----------------- SHA-256 TEST -----------------//

    bytes32 public expected256Hash =
    0x7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069;

    function calculate256Hash() internal pure returns (bytes32) {
        string memory word = "Hello World!";
        bytes32 hash = sha256(bytes(word));

        return hash;
    }

    function check256Hash() public view returns (bool) {
        return (calculate256Hash() == expected256Hash);
    }

    //----------------- RIPEMD160 TEST -----------------//

    bytes20 public expectedHash = hex"8476ee4631b9b30ac2754b0ee0c47e161d3f724c";

    function calculate160Hash() internal pure returns (bytes20) {
        string memory word = "Hello World!";
        bytes20 hash = ripemd160(bytes(word));

        return hash;
    }

    function checkRipeMd160Hash() public view returns (bool) {
        return (calculate160Hash() == expectedHash);
    }

    //----------------- GLOBAL VARIABLE TEST -----------------//

    function getBlockVariables() public view returns (uint, uint, bytes32)
    {
        return (block.number, block.timestamp, blockhash(block.number - 1));
    }

    //----------------- ECRECOVER TEST -----------------//

    address addressTest = 0x12Cb274aAD8251C875c0bf6872b67d9983E53fDd;
    bytes32 msgHash =
    0xc51dac836bc7841a01c4b631fa620904fc8724d7f9f1d3c420f0e02adf229d50;
    uint8 v = 0x1b;
    bytes32 r =
    0x44287513919034a471a7dc2b2ed121f95984ae23b20f9637ba8dff471b6719ef;
    bytes32 s =
    0x7d7dc30309a3baffbfd9342b97d0e804092c0aeb5821319aa732bc09146eafb4;

    function verifyECrecover() public view returns (bool) {
        // Use ECRECOVER to verify address
        return (ecrecover(msgHash, v, r, s) == (addressTest));
    }

    //----------------- BIG MOD EXP TEST -----------------//

    uint public checkResult;

    function verify(uint _base, uint _exp, uint _modulus) public {
        checkResult = modExp(_base, _exp, _modulus);
    }

    function modExp(
        uint256 _b,
        uint256 _e,
        uint256 _m
    ) public returns (uint256 result) {
        assembly {
        // Free memory pointer
            let pointer := mload(0x40)
        // Define length of base, exponent and modulus. 0x20 == 32 bytes
            mstore(pointer, 0x20)
            mstore(add(pointer, 0x20), 0x20)
            mstore(add(pointer, 0x40), 0x20)
        // Define variables base, exponent and modulus
            mstore(add(pointer, 0x60), _b)
            mstore(add(pointer, 0x80), _e)
            mstore(add(pointer, 0xa0), _m)
        // Store the result
            let value := mload(0xc0)
        // Call the precompiled contract 0x05 = bigModExp
            if iszero(call(not(0), 0x05, 0, pointer, 0xc0, value, 0x20)) {
                revert(0, 0)
            }
            result := mload(value)
        }
    }

    //----------------- GAS PRICE OPCODE TEST -----------------//

    uint64 gasprice;
    uint64 dummy; // dummy variable to trick G++ provider to estimate correct gasLimit

    function setGasPrice(uint64 t) external {
        dummy = t + 1;
        uint64 gp;
        assembly {
            gp := gasprice()
        }
        gasprice = gp;
    }

    function getGasPrice() external view returns (uint64) {
        return gasprice;
    }

    //----------------- DATA COPY TEST -----------------//

    bytes public memoryStored;

    function callDatacopy(bytes memory data) public returns (bytes memory) {
        bytes memory result = new bytes(data.length);
        assembly {
            let len := mload(data)
            if iszero(
                call(
                    gas(),
                    0x04,
                    0,
                    add(data, 0x20),
                    len,
                    add(result, 0x20),
                    len
                )
            ) {
                invalid()
            }
        }

        memoryStored = result;

        return result;
    }
}