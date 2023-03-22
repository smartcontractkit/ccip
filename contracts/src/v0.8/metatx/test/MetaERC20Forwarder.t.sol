

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import { MetaERC20 } from "../MetaERC20.sol";
import "forge-std/Test.sol";

contract MetaERC20Forwarder { 
    function setUp() public {
    }

    function testPrintEncodedHash() public {
        address from = 0xf017f32B661213AfAdF54B0b13534BBe7C797f37;
        address to = 0x40d08f49Ecfc3C20511B3F799701E53A1598F5B8;
        string memory requestType = string(abi.encodePacked("ForwardRequest(address from,address to,uint256 value,uint256 nonce,bytes data,uint256 validUntilTime)"));
        bytes32 requestTypeHash = keccak256(bytes(requestType));
        console.log("fromAddress");
        console.log(from); 
        console.log(to);
        uint256 amount = 1 ether;
        uint256 nonce = 0;
        bytes memory encoded = abi.encodePacked(
            requestTypeHash,
            uint256(uint160(from)),
            uint256(uint160(to)),
            amount,
            nonce
            //keccak256(req.data),
            //req.validUntilTime,
            //suffixData
        );
        console.logBytes(encoded);
        bytes32 encodedHash = keccak256(encoded);
        console.logBytes32(encodedHash);

        bytes memory encodedCalldata = abi.encodeWithSignature("metaTransfer(address,address,uint256,uint64)", from, to, amount, uint64(1337));
        console.logBytes(encodedCalldata);
        bytes32 calldataHash = keccak256(encodedCalldata);
        console.logBytes32(calldataHash);
    }
}
