package ccipevm

// tokensABI is the ABI of the Client.EVMTokenAmount struct defined in Client.sol.
const tokensABI = `
[
    {
        "components":
        [
            {
                "name": "token",
                "type": "address"
            },
            {
                "name": "amount",
                "type": "uint256"
            }
        ],
        "type": "tuple[]"
    }
]
`

// metaDataHashABI is the ABI of the input to the metadata keccak256 hash calculation.
/*
keccak256(
	abi.encode(
		ANY_2_EVM_MESSAGE_HASH,
		original.header.sourceChainSelector,
		original.header.destChainSelector,
		onRamp
	)
)
*/
const metaDataHashABI = `
[
	{
		"name": "ANY_2_EVM_MESSAGE_HASH",
		"type": "bytes32"
	},
	{
		"name": "headerSourceChainSelector",
		"type": "uint64"
	},
	{
		"name": "headerDestinationChainSelector",
		"type": "uint64"
	},
	{
		"name": "onRamp",
		"type": "bytes"
	}
]
`

// fixedSizeValuesABI is the input to the inner keccak256 hash that hashes
// fixed size fields in the Any2EVMRampMessage.
/*
keccak256(
	abi.encode(
		original.header.messageId,
		original.sender,
		original.receiver,
		original.header.sequenceNumber,
		original.gasLimit,
		original.header.nonce
	)
)
*/
const fixedSizeValuesABI = `
[
    {
        "name": "headerMessageId",
        "type": "bytes32"
    },
	{
        "name": "sender",
        "type": "bytes"
    },
    {
        "name": "receiver",
        "type": "address"
    },
    {
        "name": "headerSequenceNumber",
        "type": "uint64"
    },
    {
        "name": "gasLimit",
        "type": "uint256"
    },
    {
        "name": "headerNonce",
        "type": "uint64"
    }
]
`

// finalHashInputABI is the input to the final keccak256 hash in Internal._hash(Any2EVMRampMessage).
const finalHashInputABI = `
[
    {
        "name": "leafDomainSeparator",
        "type": "bytes32"
    },
    {
        "name": "implicitMetaDataHash",
        "type": "bytes32"
    },
    {
        "name": "fixedSizeValuesHash",
        "type": "bytes32"
    },
    {
        "name": "dataHash",
        "type": "bytes32"
    },
    {
        "name": "tokenAmountsHash",
        "type": "bytes32"
    },
    {
        "name": "sourceTokenDataHash",
        "type": "bytes32"
    }
]
`
