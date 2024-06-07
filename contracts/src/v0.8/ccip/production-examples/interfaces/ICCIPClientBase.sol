pragma solidity ^0.8.0;

interface ICCIPClientBase {
    error InvalidRouter(address router);
    error InvalidChain(uint64 chainSelector);
    error InvalidSender(bytes sender);
    error InvalidRecipient(bytes recipient);

    struct approvedSenderUpdate {
        uint64 destChainSelector;
        bytes sender;
    }

    struct ChainInfo {
        bytes extraArgsBytes;
        bytes recipient;
    }

    function getRouter() external view returns (address);
}