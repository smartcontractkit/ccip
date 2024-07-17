// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package manychainmultisig

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ManyChainMultiSigConfig is an auto generated low-level Go binding around an user-defined struct.
type ManyChainMultiSigConfig struct {
	Signers      []ManyChainMultiSigSigner
	GroupQuorums [32]uint8
	GroupParents [32]uint8
}

// ManyChainMultiSigOp is an auto generated low-level Go binding around an user-defined struct.
type ManyChainMultiSigOp struct {
	ChainId  *big.Int
	MultiSig common.Address
	Nonce    *big.Int
	To       common.Address
	Value    *big.Int
	Data     []byte
}

// ManyChainMultiSigRootMetadata is an auto generated low-level Go binding around an user-defined struct.
type ManyChainMultiSigRootMetadata struct {
	ChainId              *big.Int
	MultiSig             common.Address
	PreOpCount           *big.Int
	PostOpCount          *big.Int
	OverridePreviousRoot bool
}

// ManyChainMultiSigSignature is an auto generated low-level Go binding around an user-defined struct.
type ManyChainMultiSigSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// ManyChainMultiSigSigner is an auto generated low-level Go binding around an user-defined struct.
type ManyChainMultiSigSigner struct {
	Addr  common.Address
	Index uint8
	Group uint8
}

// ManychainmultisigMetaData contains all meta data concerning the Manychainmultisig contract.
var ManychainmultisigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_NUM_SIGNERS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NUM_GROUPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"op\",\"type\":\"tuple\",\"internalType\":\"structManyChainMultiSig.Op\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiSig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structManyChainMultiSig.Config\",\"components\":[{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"structManyChainMultiSig.Signer[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"group\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"groupQuorums\",\"type\":\"uint8[32]\",\"internalType\":\"uint8[32]\"},{\"name\":\"groupParents\",\"type\":\"uint8[32]\",\"internalType\":\"uint8[32]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOpCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validUntil\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootMetadata\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structManyChainMultiSig.RootMetadata\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiSig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"preOpCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"postOpCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"overridePreviousRoot\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"signerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signerGroups\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"groupQuorums\",\"type\":\"uint8[32]\",\"internalType\":\"uint8[32]\"},{\"name\":\"groupParents\",\"type\":\"uint8[32]\",\"internalType\":\"uint8[32]\"},{\"name\":\"clearRoot\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validUntil\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structManyChainMultiSig.RootMetadata\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiSig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"preOpCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"postOpCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"overridePreviousRoot\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"metadataProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structManyChainMultiSig.Signature[]\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structManyChainMultiSig.Config\",\"components\":[{\"name\":\"signers\",\"type\":\"tuple[]\",\"internalType\":\"structManyChainMultiSig.Signer[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"group\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"groupQuorums\",\"type\":\"uint8[32]\",\"internalType\":\"uint8[32]\"},{\"name\":\"groupParents\",\"type\":\"uint8[32]\",\"internalType\":\"uint8[32]\"}]},{\"name\":\"isRootCleared\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validUntil\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structManyChainMultiSig.RootMetadata\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiSig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"preOpCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"postOpCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"overridePreviousRoot\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OpExecuted\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallReverted\",\"inputs\":[{\"name\":\"error\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"GroupTreeNotWellFormed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientSigners\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MissingConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBoundsGroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBoundsGroupQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBoundsNumOfSigners\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingOps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostOpCountReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofCannotBeVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignedHashAlreadySeen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignerGroupsLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignerInDisabledGroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignersAddressesMustBeStrictlyIncreasing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidUntilHasAlreadyPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongMultiSig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPostOpCount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPreOpCount\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001d3362000023565b62000091565b600180546001600160a01b03191690556200003e8162000041565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61264380620000a16000396000f3fe6080604052600436106100e15760003560e01c8063846c67ef1161007f578063b759d68511610059578063b759d68514610345578063c3f909d414610358578063e30c39781461037a578063f2fde38b1461039857600080fd5b8063846c67ef146102de5780638da5cb5b146102fe578063a76f55981461033057600080fd5b80636b45fb3e116100bb5780636b45fb3e146101a3578063715018a61461029257806379ba5097146102a95780637cc38b28146102be57600080fd5b80635a2519ef146100ed5780635ca1e16514610119578063627e8a3b1461016f57600080fd5b366100e857005b600080fd5b3480156100f957600080fd5b50610102602081565b60405160ff90911681526020015b60405180910390f35b34801561012557600080fd5b506040805160608101825260075480825260085463ffffffff81166020808501829052600160201b90920464ffffffffff1693850193909352835191825281019190915201610110565b34801561017b57600080fd5b50600854600160201b900464ffffffffff1660405164ffffffffff9091168152602001610110565b3480156101af57600080fd5b5061023b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a0810182526009548152600a546001600160a01b0381166020830152600160a01b810464ffffffffff90811693830193909352600160c81b81049092166060820152600160f01b90910460ff161515608082015290565b6040516101109190815181526020808301516001600160a01b03169082015260408083015164ffffffffff908116918301919091526060808401519091169082015260809182015115159181019190915260a00190565b34801561029e57600080fd5b506102a76103b8565b005b3480156102b557600080fd5b506102a76103cc565b3480156102ca57600080fd5b506102a76102d9366004611b5b565b61044b565b3480156102ea57600080fd5b506102a76102f9366004611c5f565b610a2f565b34801561030a57600080fd5b506000546001600160a01b03165b6040516001600160a01b039091168152602001610110565b34801561033c57600080fd5b5061010260c881565b6102a7610353366004611d0a565b6111a2565b34801561036457600080fd5b5061036d61144d565b6040516101109190611dab565b34801561038657600080fd5b506001546001600160a01b0316610318565b3480156103a457600080fd5b506102a76103b3366004611e5f565b61158f565b6103c0611600565b6103ca600061165a565b565b60015433906001600160a01b0316811461043f5760405162461bcd60e51b815260206004820152602960248201527f4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206044820152683732bb9037bbb732b960b91b60648201526084015b60405180910390fd5b6104488161165a565b50565b60006104bb888860405160200161047292919091825263ffffffff16602082015260400190565b604051602081830303815290604052805190602001207f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b60008181526006602052604090205490915060ff16156104ee576040516348c2688b60e01b815260040160405180910390fd5b6040805160608101825260008082526020820181905291810182905290610513611a19565b60005b858110156106de573687878381811061053157610531611e7c565b6060029190910191506000905061055e8761054f6020850185611e92565b84602001358560400135611673565b9050806001600160a01b0316856001600160a01b03161061059257604051630946dd8160e31b815260040160405180910390fd5b6001600160a01b038082166000818152600260209081526040918290208251606081018452905494851680825260ff600160a01b8704811693830193909352600160a81b90950490911691810191909152975091955085911461060857604051632057875960e21b815260040160405180910390fd5b60408601515b848160ff166020811061062357610623611e7c565b6020020180519061063382611ecb565b60ff9081169091526004915082166020811061065157610651611e7c565b602091828204019190069054906101000a900460ff1660ff16858260ff166020811061067f5761067f611e7c565b602002015160ff16036106c85760ff8116156106c857600560ff8216602081106106ab576106ab611e7c565b602081049091015460ff601f9092166101000a900416905061060e565b50505080806106d690611eea565b915050610516565b5060045460ff1660000361070557604051635530c2e560e11b815260040160405180910390fd5b600454815160ff91821691161015610730576040516361774dcf60e11b815260040160405180910390fd5b505050428763ffffffff16101561075a5760405163582bd22960e11b815260040160405180910390fd5b60007fe6b82be989101b4eb519770114b997b97b3c8707515286748a871717f0e4ea1c8760405160200161078f929190611f82565b6040516020818303038152906040528051906020012090506107e78686808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508d925085915061169b9050565b6108035760405162948a8760e61b815260040160405180910390fd5b5046863514610824576040516217e1ef60ea1b815260040160405180910390fd5b6108346040870160208801611e5f565b6001600160a01b0316306001600160a01b03161461086557604051639a84601560e01b815260040160405180910390fd5b600854600a5464ffffffffff600160201b909204821691600160c81b9091041681148015906108a1575061089f60a0880160808901611f96565b155b156108bf57604051633230825b60e01b815260040160405180910390fd5b6108cf6060880160408901611fb3565b64ffffffffff168164ffffffffff16146108fc5760405163a255a76360e01b815260040160405180910390fd5b61090c6080880160608901611fb3565b64ffffffffff166109236060890160408a01611fb3565b64ffffffffff161115610949576040516318c26a5f60e31b815260040160405180910390fd5b600082815260066020908152604091829020805460ff191660011790558151606080820184528c825263ffffffff8c1692820192909252918281019161099491908b01908b01611fb3565b64ffffffffff9081169091528151600755602082015160088054604090940151909216600160201b0268ffffffffffffffffff1990931663ffffffff909116179190911790558660096109e78282611fd0565b905050887f7ea643ae44677f24e0d6f40168893712daaf729b0a38fe7702d21cb544c841018989604051610a1c929190612067565b60405180910390a2505050505050505050565b610a37611600565b851580610a44575060c886115b15610a6257604051633c3b072960e21b815260040160405180910390fd5b858414610a8257604051630f1f305360e41b815260040160405180910390fd5b610a8a611a19565b60005b85811015610b42576020878783818110610aa957610aa9611e7c565b9050602002016020810190610abe9190611e92565b60ff1610610adf57604051635cd7472960e11b815260040160405180910390fd5b81878783818110610af257610af2611e7c565b9050602002016020810190610b079190611e92565b60ff1660208110610b1a57610b1a611e7c565b60200201805190610b2a82611ecb565b60ff1690525080610b3a81611eea565b915050610a8d565b5060005b6020811015610d3457600081610b5e60016020612081565b60ff16610b6b919061209a565b90508015801590610ba3575080858260208110610b8a57610b8a611e7c565b602002016020810190610b9d9190611e92565b60ff1610155b80610bdd575080158015610bdd5750848160208110610bc457610bc4611e7c565b602002016020810190610bd79190611e92565b60ff1615155b15610bfe576040516001627ce2ed60e11b0319815260040160405180910390fd5b6000868260208110610c1257610c12611e7c565b602002016020810190610c259190611e92565b60ff161590508015610c6e57838260208110610c4357610c43611e7c565b602002015160ff1615610c6957604051638db4e75d60e01b815260040160405180910390fd5b610d1f565b868260208110610c8057610c80611e7c565b602002016020810190610c939190611e92565b60ff16848360208110610ca857610ca8611e7c565b602002015160ff161015610ccf57604051635d8009b760e11b815260040160405180910390fd5b83868360208110610ce257610ce2611e7c565b602002016020810190610cf59190611e92565b60ff1660208110610d0857610d08611e7c565b60200201805190610d1882611ecb565b60ff169052505b50508080610d2c90611eea565b915050610b46565b505060006003600001805480602002602001604051908101604052809291908181526020016000905b82821015610db457600084815260209081902060408051606081018252918501546001600160a01b038116835260ff600160a01b8204811684860152600160a81b9091041690820152825260019092019101610d5d565b50505050905060005b8151811015610e58576000828281518110610dda57610dda611e7c565b602090810291909101810151516001600160a01b03811660009081526002909252604090912080546001600160b01b0319169055600380549192509080610e2357610e236120ad565b600082815260209020810160001990810180546001600160b01b03191690550190555080610e5081611eea565b915050610dbd565b5060035415610e6957610e696120c3565b610e766004856020611a38565b50610e846005846020611a38565b506000805b8881101561108c57898982818110610ea357610ea3611e7c565b9050602002016020810190610eb89190611e5f565b6001600160a01b0316826001600160a01b031610610ee957604051630946dd8160e31b815260040160405180910390fd5b600060405180606001604052808c8c85818110610f0857610f08611e7c565b9050602002016020810190610f1d9190611e5f565b6001600160a01b031681526020018360ff1681526020018a8a85818110610f4657610f46611e7c565b9050602002016020810190610f5b9190611e92565b60ff169052905080600260008d8d86818110610f7957610f79611e7c565b9050602002016020810190610f8e9190611e5f565b6001600160a01b0390811682526020808301939093526040918201600090812085518154878701519786015160ff908116600160a81b90810260ff60a81b199a8316600160a01b9081026001600160a81b0319958616968a1696909617959095178b161790945560038054600181018255955289517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9095018054988b0151978b01518216909402961690910295169190921617929092179092169190911790558a8a8381811061106157611061611e7c565b90506020020160208101906110769190611e5f565b925050808061108490611eea565b915050610e89565b50821561115d5760088054604080516060808201835260008083526020808401829052600160201b80870464ffffffffff1694860185905260079290925568ffffffffffffffffff1990951690830217909455815160a08101835246808252309482018590529281018290529384018190526001608090940193909352600955600a8054600160f01b6001600160c81b031990911664ffffffffff60a01b1990931692909217600160a01b84021765ffffffffffff60c81b1916600160c81b90930260ff60f01b1916929092171790555b7f0a4974ad206b9c736f9ab2feac1c9b1d043fe4ef377c70ae45659f2ef089f03e60038460405161118f929190612361565b60405180910390a1505050505050505050565b60408051606081018252600754815260085463ffffffff8116602083015264ffffffffff600160201b9091048116928201839052600a549192600160c81b9092041611611202576040516315b6266360e31b815260040160405180910390fd5b83354614611222576040516217e1ef60ea1b815260040160405180910390fd5b306112336040860160208701611e5f565b6001600160a01b03161461125a57604051639a84601560e01b815260040160405180910390fd5b806020015163ffffffff16421115611285576040516309ba674360e41b815260040160405180910390fd5b806040015164ffffffffff168460400160208101906112a49190611fb3565b64ffffffffff16146112c95760405163d9c6386f60e01b815260040160405180910390fd5b60007f08d275622006c4ca82d03f498e90163cafd53c663a48470c3b52ac8bfbd9f52c856040516020016112fe92919061242e565b604051602081830303815290604052805190602001209050611356848480806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250508551915084905061169b565b6113725760405162948a8760e61b815260040160405180910390fd5b60408201516113829060016124f7565b6008805464ffffffffff92909216600160201b0268ffffffffff00000000199092169190911790556113d46113bd6080870160608801611e5f565b60808701356113cf60a089018961251c565b6116b3565b6113e46060860160408701611fb3565b64ffffffffff167f87d58fdd48be753fb9ef4ec8a5895086c401506da8b4d752abc90602c3e62d1d61141c6080880160608901611e5f565b61142960a089018961251c565b896080013560405161143e9493929190612563565b60405180910390a25050505050565b611455611ace565b604080516003805460806020820284018101909452606083018181529293919284929091849160009085015b828210156114d857600084815260209081902060408051606081018252918501546001600160a01b038116835260ff600160a01b8204811684860152600160a81b9091041690820152825260019092019101611481565b50505090825250604080516104008101918290526020928301929091600185019190826000855b825461010083900a900460ff168152602060019283018181049485019490930390920291018084116114ff575050509284525050604080516104008101918290526020938401939092506002850191826000855b825461010083900a900460ff16815260206001928301818104948501949093039092029101808411611553579050505050505081525050905090565b611597611600565b600180546001600160a01b0383166001600160a01b031990911681179091556115c86000546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000546001600160a01b031633146103ca5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610436565b600180546001600160a01b031916905561044881611740565b600080600061168487878787611790565b9150915061169181611854565b5095945050505050565b6000826116a8858461199e565b1490505b9392505050565b600080856001600160a01b03168585856040516116d1929190612599565b60006040518083038185875af1925050503d806000811461170e576040519150601f19603f3d011682016040523d82523d6000602084013e611713565b606091505b50915091508161173857806040516370de1b4b60e01b815260040161043691906125a9565b505050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156117c7575060009050600361184b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561181b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166118445760006001925092505061184b565b9150600090505b94509492505050565b6000816004811115611868576118686125f7565b036118705750565b6001816004811115611884576118846125f7565b036118d15760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610436565b60028160048111156118e5576118e56125f7565b036119325760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610436565b6003816004811115611946576119466125f7565b036104485760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610436565b600081815b84518110156119e3576119cf828683815181106119c2576119c2611e7c565b60200260200101516119ed565b9150806119db81611eea565b9150506119a3565b5090505b92915050565b6000818310611a095760008281526020849052604090206116ac565b5060009182526020526040902090565b6040518061040001604052806020906020820280368337509192915050565b600183019183908215611abe5791602002820160005b83821115611a8f57833560ff1683826101000a81548160ff021916908360ff1602179055509260200192600101602081600001049283019260010302611a4e565b8015611abc5782816101000a81549060ff0219169055600101602081600001049283019260010302611a8f565b505b50611aca929150611afa565b5090565b604051806060016040528060608152602001611ae8611a19565b8152602001611af5611a19565b905290565b5b80821115611aca5760008155600101611afb565b60008083601f840112611b2157600080fd5b50813567ffffffffffffffff811115611b3957600080fd5b6020830191508360208260051b8501011115611b5457600080fd5b9250929050565b6000806000806000806000878903610120811215611b7857600080fd5b88359750602089013563ffffffff81168114611b9357600080fd5b965060a0603f1982011215611ba757600080fd5b5060408801945060e088013567ffffffffffffffff80821115611bc957600080fd5b611bd58b838c01611b0f565b90965094506101008a0135915080821115611bef57600080fd5b818a0191508a601f830112611c0357600080fd5b813581811115611c1257600080fd5b8b6020606083028501011115611c2757600080fd5b60208301945080935050505092959891949750929550565b8061040081018310156119e757600080fd5b801515811461044857600080fd5b6000806000806000806000610860888a031215611c7b57600080fd5b873567ffffffffffffffff80821115611c9357600080fd5b611c9f8b838c01611b0f565b909950975060208a0135915080821115611cb857600080fd5b50611cc58a828b01611b0f565b9096509450611cd990508960408a01611c3f565b9250611ce9896104408a01611c3f565b9150610840880135611cfa81611c51565b8091505092959891949750929550565b600080600060408486031215611d1f57600080fd5b833567ffffffffffffffff80821115611d3757600080fd5b9085019060c08288031215611d4b57600080fd5b90935060208501359080821115611d6157600080fd5b50611d6e86828701611b0f565b9497909650939450505050565b8060005b6020808210611d8e5750611da5565b825160ff1685529384019390910190600101611d7f565b50505050565b6020808252825161082083830152805161084084018190526000929182019083906108608601905b80831015611e1857835180516001600160a01b031683528581015160ff9081168785015260409182015116908301529284019260019290920191606090910190611dd3565b509286015192611e2b6040870185611d7b565b60408701519350611e40610440870185611d7b565b9695505050505050565b6001600160a01b038116811461044857600080fd5b600060208284031215611e7157600080fd5b81356116ac81611e4a565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611ea457600080fd5b813560ff811681146116ac57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff8103611ee157611ee1611eb5565b60010192915050565b600060018201611efc57611efc611eb5565b5060010190565b64ffffffffff8116811461044857600080fd5b803582526020810135611f2881611e4a565b6001600160a01b031660208301526040810135611f4481611f03565b64ffffffffff9081166040840152606082013590611f6182611f03565b1660608301526080810135611f7581611c51565b8015156080840152505050565b82815260c081016116ac6020830184611f16565b600060208284031215611fa857600080fd5b81356116ac81611c51565b600060208284031215611fc557600080fd5b81356116ac81611f03565b81358155600181016020830135611fe681611e4a565b81546040850135611ff681611f03565b606086013561200481611f03565b608087013561201281611c51565b60c89190911b64ffffffffff60c81b1660a09290921b64ffffffffff60a01b166001600160f81b0319939093166001600160a01b039490941693909317919091171790151560f01b60ff60f01b161790555050565b63ffffffff8316815260c081016116ac6020830184611f16565b60ff82811682821603908111156119e7576119e7611eb5565b818103818111156119e7576119e7611eb5565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fd5b8060005b602080601f8301106120ef5750611da5565b825460ff8082168752600882901c8116838801526040612118818901838560101c1660ff169052565b606061212d818a01848660181c1660ff169052565b6080612141818b018587891c1660ff169052565b60a09550612158868b01858760281c1660ff169052565b60c061216d818c01868860301c1660ff169052565b60e0612182818d01878960381c1660ff169052565b60ff87861c8716166101008d01526121a56101208d01878960481c1660ff169052565b6121ba6101408d01878960501c1660ff169052565b6121cf6101608d01878960581c1660ff169052565b60ff87851c8716166101808d01526121f26101a08d01878960681c1660ff169052565b6122076101c08d01878960701c1660ff169052565b61221c6101e08d01878960781c1660ff169052565b60ff87841c8716166102008d015261223f6102208d01878960881c1660ff169052565b6122546102408d01878960901c1660ff169052565b6122696102608d01878960981c1660ff169052565b60ff87891c8716166102808d015261228c6102a08d01878960a81c1660ff169052565b6122a16102c08d01878960b01c1660ff169052565b6122b66102e08d01878960b81c1660ff169052565b60ff87831c8716166103008d01526122d96103208d01878960c81c1660ff169052565b6122ee6103408d01878960d01c1660ff169052565b6123036103608d01878960d81c1660ff169052565b60ff87821c8716166103808d0152505050505061232b6103a08801828460e81c1660ff169052565b6123406103c08801828460f01c1660ff169052565b5060f81c6103e08601525061040090930192600191909101906020016120dd565b600060408083526108608301610820828501528086548083526108808601915087600052602092508260002060005b828110156123d05781546001600160a01b038116855260ff60a082901c81168787015260a89190911c168685015260609093019260019182019101612390565b5050506123e360608601600189016120d9565b6123f46104608601600289016120d9565b941515930192909252509092915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b82815260406020820152813560408201526000602083013561244f81611e4a565b6001600160a01b03908116606084015260408401359061246e82611f03565b64ffffffffff821660808501526060850135915061248b82611e4a565b1660a083810191909152608084013560c084015283013536849003601e190181126124b557600080fd5b830160208101903567ffffffffffffffff8111156124d257600080fd5b8036038213156124e157600080fd5b60c060e0850152611e4061010085018284612405565b64ffffffffff81811683821601908082111561251557612515611eb5565b5092915050565b6000808335601e1984360301811261253357600080fd5b83018035915067ffffffffffffffff82111561254e57600080fd5b602001915036819003821315611b5457600080fd5b6001600160a01b03851681526060602082018190526000906125889083018587612405565b905082604083015295945050505050565b8183823760009101908152919050565b600060208083528351808285015260005b818110156125d6578581018301518582016040015282016125ba565b506000604082860101526040601f19601f8301168501019250505092915050565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220789e10e5d6588dc4b26dba62aaa06971ed6ac6cbd43499f78a57e1d8415d671c64736f6c63430008130033",
}

// ManychainmultisigABI is the input ABI used to generate the binding from.
// Deprecated: Use ManychainmultisigMetaData.ABI instead.
var ManychainmultisigABI = ManychainmultisigMetaData.ABI

// ManychainmultisigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ManychainmultisigMetaData.Bin instead.
var ManychainmultisigBin = ManychainmultisigMetaData.Bin

// DeployManychainmultisig deploys a new Ethereum contract, binding an instance of Manychainmultisig to it.
func DeployManychainmultisig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Manychainmultisig, error) {
	parsed, err := ManychainmultisigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ManychainmultisigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Manychainmultisig{ManychainmultisigCaller: ManychainmultisigCaller{contract: contract}, ManychainmultisigTransactor: ManychainmultisigTransactor{contract: contract}, ManychainmultisigFilterer: ManychainmultisigFilterer{contract: contract}}, nil
}

// Manychainmultisig is an auto generated Go binding around an Ethereum contract.
type Manychainmultisig struct {
	ManychainmultisigCaller     // Read-only binding to the contract
	ManychainmultisigTransactor // Write-only binding to the contract
	ManychainmultisigFilterer   // Log filterer for contract events
}

// ManychainmultisigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManychainmultisigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManychainmultisigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManychainmultisigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManychainmultisigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManychainmultisigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManychainmultisigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManychainmultisigSession struct {
	Contract     *Manychainmultisig // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManychainmultisigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManychainmultisigCallerSession struct {
	Contract *ManychainmultisigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ManychainmultisigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManychainmultisigTransactorSession struct {
	Contract     *ManychainmultisigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ManychainmultisigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManychainmultisigRaw struct {
	Contract *Manychainmultisig // Generic contract binding to access the raw methods on
}

// ManychainmultisigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManychainmultisigCallerRaw struct {
	Contract *ManychainmultisigCaller // Generic read-only contract binding to access the raw methods on
}

// ManychainmultisigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManychainmultisigTransactorRaw struct {
	Contract *ManychainmultisigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManychainmultisig creates a new instance of Manychainmultisig, bound to a specific deployed contract.
func NewManychainmultisig(address common.Address, backend bind.ContractBackend) (*Manychainmultisig, error) {
	contract, err := bindManychainmultisig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manychainmultisig{ManychainmultisigCaller: ManychainmultisigCaller{contract: contract}, ManychainmultisigTransactor: ManychainmultisigTransactor{contract: contract}, ManychainmultisigFilterer: ManychainmultisigFilterer{contract: contract}}, nil
}

// NewManychainmultisigCaller creates a new read-only instance of Manychainmultisig, bound to a specific deployed contract.
func NewManychainmultisigCaller(address common.Address, caller bind.ContractCaller) (*ManychainmultisigCaller, error) {
	contract, err := bindManychainmultisig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigCaller{contract: contract}, nil
}

// NewManychainmultisigTransactor creates a new write-only instance of Manychainmultisig, bound to a specific deployed contract.
func NewManychainmultisigTransactor(address common.Address, transactor bind.ContractTransactor) (*ManychainmultisigTransactor, error) {
	contract, err := bindManychainmultisig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigTransactor{contract: contract}, nil
}

// NewManychainmultisigFilterer creates a new log filterer instance of Manychainmultisig, bound to a specific deployed contract.
func NewManychainmultisigFilterer(address common.Address, filterer bind.ContractFilterer) (*ManychainmultisigFilterer, error) {
	contract, err := bindManychainmultisig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigFilterer{contract: contract}, nil
}

// bindManychainmultisig binds a generic wrapper to an already deployed contract.
func bindManychainmultisig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ManychainmultisigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manychainmultisig *ManychainmultisigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manychainmultisig.Contract.ManychainmultisigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manychainmultisig *ManychainmultisigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.ManychainmultisigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manychainmultisig *ManychainmultisigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.ManychainmultisigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manychainmultisig *ManychainmultisigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manychainmultisig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manychainmultisig *ManychainmultisigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manychainmultisig *ManychainmultisigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.contract.Transact(opts, method, params...)
}

// MAXNUMSIGNERS is a free data retrieval call binding the contract method 0xa76f5598.
//
// Solidity: function MAX_NUM_SIGNERS() view returns(uint8)
func (_Manychainmultisig *ManychainmultisigCaller) MAXNUMSIGNERS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "MAX_NUM_SIGNERS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXNUMSIGNERS is a free data retrieval call binding the contract method 0xa76f5598.
//
// Solidity: function MAX_NUM_SIGNERS() view returns(uint8)
func (_Manychainmultisig *ManychainmultisigSession) MAXNUMSIGNERS() (uint8, error) {
	return _Manychainmultisig.Contract.MAXNUMSIGNERS(&_Manychainmultisig.CallOpts)
}

// MAXNUMSIGNERS is a free data retrieval call binding the contract method 0xa76f5598.
//
// Solidity: function MAX_NUM_SIGNERS() view returns(uint8)
func (_Manychainmultisig *ManychainmultisigCallerSession) MAXNUMSIGNERS() (uint8, error) {
	return _Manychainmultisig.Contract.MAXNUMSIGNERS(&_Manychainmultisig.CallOpts)
}

// NUMGROUPS is a free data retrieval call binding the contract method 0x5a2519ef.
//
// Solidity: function NUM_GROUPS() view returns(uint8)
func (_Manychainmultisig *ManychainmultisigCaller) NUMGROUPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "NUM_GROUPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NUMGROUPS is a free data retrieval call binding the contract method 0x5a2519ef.
//
// Solidity: function NUM_GROUPS() view returns(uint8)
func (_Manychainmultisig *ManychainmultisigSession) NUMGROUPS() (uint8, error) {
	return _Manychainmultisig.Contract.NUMGROUPS(&_Manychainmultisig.CallOpts)
}

// NUMGROUPS is a free data retrieval call binding the contract method 0x5a2519ef.
//
// Solidity: function NUM_GROUPS() view returns(uint8)
func (_Manychainmultisig *ManychainmultisigCallerSession) NUMGROUPS() (uint8, error) {
	return _Manychainmultisig.Contract.NUMGROUPS(&_Manychainmultisig.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(((address,uint8,uint8)[],uint8[32],uint8[32]))
func (_Manychainmultisig *ManychainmultisigCaller) GetConfig(opts *bind.CallOpts) (ManyChainMultiSigConfig, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(ManyChainMultiSigConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ManyChainMultiSigConfig)).(*ManyChainMultiSigConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(((address,uint8,uint8)[],uint8[32],uint8[32]))
func (_Manychainmultisig *ManychainmultisigSession) GetConfig() (ManyChainMultiSigConfig, error) {
	return _Manychainmultisig.Contract.GetConfig(&_Manychainmultisig.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(((address,uint8,uint8)[],uint8[32],uint8[32]))
func (_Manychainmultisig *ManychainmultisigCallerSession) GetConfig() (ManyChainMultiSigConfig, error) {
	return _Manychainmultisig.Contract.GetConfig(&_Manychainmultisig.CallOpts)
}

// GetOpCount is a free data retrieval call binding the contract method 0x627e8a3b.
//
// Solidity: function getOpCount() view returns(uint40)
func (_Manychainmultisig *ManychainmultisigCaller) GetOpCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "getOpCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOpCount is a free data retrieval call binding the contract method 0x627e8a3b.
//
// Solidity: function getOpCount() view returns(uint40)
func (_Manychainmultisig *ManychainmultisigSession) GetOpCount() (*big.Int, error) {
	return _Manychainmultisig.Contract.GetOpCount(&_Manychainmultisig.CallOpts)
}

// GetOpCount is a free data retrieval call binding the contract method 0x627e8a3b.
//
// Solidity: function getOpCount() view returns(uint40)
func (_Manychainmultisig *ManychainmultisigCallerSession) GetOpCount() (*big.Int, error) {
	return _Manychainmultisig.Contract.GetOpCount(&_Manychainmultisig.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32 root, uint32 validUntil)
func (_Manychainmultisig *ManychainmultisigCaller) GetRoot(opts *bind.CallOpts) (struct {
	Root       [32]byte
	ValidUntil uint32
}, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "getRoot")

	outstruct := new(struct {
		Root       [32]byte
		ValidUntil uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ValidUntil = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32 root, uint32 validUntil)
func (_Manychainmultisig *ManychainmultisigSession) GetRoot() (struct {
	Root       [32]byte
	ValidUntil uint32
}, error) {
	return _Manychainmultisig.Contract.GetRoot(&_Manychainmultisig.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32 root, uint32 validUntil)
func (_Manychainmultisig *ManychainmultisigCallerSession) GetRoot() (struct {
	Root       [32]byte
	ValidUntil uint32
}, error) {
	return _Manychainmultisig.Contract.GetRoot(&_Manychainmultisig.CallOpts)
}

// GetRootMetadata is a free data retrieval call binding the contract method 0x6b45fb3e.
//
// Solidity: function getRootMetadata() view returns((uint256,address,uint40,uint40,bool))
func (_Manychainmultisig *ManychainmultisigCaller) GetRootMetadata(opts *bind.CallOpts) (ManyChainMultiSigRootMetadata, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "getRootMetadata")

	if err != nil {
		return *new(ManyChainMultiSigRootMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(ManyChainMultiSigRootMetadata)).(*ManyChainMultiSigRootMetadata)

	return out0, err

}

// GetRootMetadata is a free data retrieval call binding the contract method 0x6b45fb3e.
//
// Solidity: function getRootMetadata() view returns((uint256,address,uint40,uint40,bool))
func (_Manychainmultisig *ManychainmultisigSession) GetRootMetadata() (ManyChainMultiSigRootMetadata, error) {
	return _Manychainmultisig.Contract.GetRootMetadata(&_Manychainmultisig.CallOpts)
}

// GetRootMetadata is a free data retrieval call binding the contract method 0x6b45fb3e.
//
// Solidity: function getRootMetadata() view returns((uint256,address,uint40,uint40,bool))
func (_Manychainmultisig *ManychainmultisigCallerSession) GetRootMetadata() (ManyChainMultiSigRootMetadata, error) {
	return _Manychainmultisig.Contract.GetRootMetadata(&_Manychainmultisig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manychainmultisig *ManychainmultisigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manychainmultisig *ManychainmultisigSession) Owner() (common.Address, error) {
	return _Manychainmultisig.Contract.Owner(&_Manychainmultisig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manychainmultisig *ManychainmultisigCallerSession) Owner() (common.Address, error) {
	return _Manychainmultisig.Contract.Owner(&_Manychainmultisig.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Manychainmultisig *ManychainmultisigCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manychainmultisig.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Manychainmultisig *ManychainmultisigSession) PendingOwner() (common.Address, error) {
	return _Manychainmultisig.Contract.PendingOwner(&_Manychainmultisig.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Manychainmultisig *ManychainmultisigCallerSession) PendingOwner() (common.Address, error) {
	return _Manychainmultisig.Contract.PendingOwner(&_Manychainmultisig.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Manychainmultisig *ManychainmultisigTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manychainmultisig.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Manychainmultisig *ManychainmultisigSession) AcceptOwnership() (*types.Transaction, error) {
	return _Manychainmultisig.Contract.AcceptOwnership(&_Manychainmultisig.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Manychainmultisig.Contract.AcceptOwnership(&_Manychainmultisig.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb759d685.
//
// Solidity: function execute((uint256,address,uint40,address,uint256,bytes) op, bytes32[] proof) payable returns()
func (_Manychainmultisig *ManychainmultisigTransactor) Execute(opts *bind.TransactOpts, op ManyChainMultiSigOp, proof [][32]byte) (*types.Transaction, error) {
	return _Manychainmultisig.contract.Transact(opts, "execute", op, proof)
}

// Execute is a paid mutator transaction binding the contract method 0xb759d685.
//
// Solidity: function execute((uint256,address,uint40,address,uint256,bytes) op, bytes32[] proof) payable returns()
func (_Manychainmultisig *ManychainmultisigSession) Execute(op ManyChainMultiSigOp, proof [][32]byte) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.Execute(&_Manychainmultisig.TransactOpts, op, proof)
}

// Execute is a paid mutator transaction binding the contract method 0xb759d685.
//
// Solidity: function execute((uint256,address,uint40,address,uint256,bytes) op, bytes32[] proof) payable returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) Execute(op ManyChainMultiSigOp, proof [][32]byte) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.Execute(&_Manychainmultisig.TransactOpts, op, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manychainmultisig *ManychainmultisigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manychainmultisig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manychainmultisig *ManychainmultisigSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manychainmultisig.Contract.RenounceOwnership(&_Manychainmultisig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manychainmultisig.Contract.RenounceOwnership(&_Manychainmultisig.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x846c67ef.
//
// Solidity: function setConfig(address[] signerAddresses, uint8[] signerGroups, uint8[32] groupQuorums, uint8[32] groupParents, bool clearRoot) returns()
func (_Manychainmultisig *ManychainmultisigTransactor) SetConfig(opts *bind.TransactOpts, signerAddresses []common.Address, signerGroups []uint8, groupQuorums [32]uint8, groupParents [32]uint8, clearRoot bool) (*types.Transaction, error) {
	return _Manychainmultisig.contract.Transact(opts, "setConfig", signerAddresses, signerGroups, groupQuorums, groupParents, clearRoot)
}

// SetConfig is a paid mutator transaction binding the contract method 0x846c67ef.
//
// Solidity: function setConfig(address[] signerAddresses, uint8[] signerGroups, uint8[32] groupQuorums, uint8[32] groupParents, bool clearRoot) returns()
func (_Manychainmultisig *ManychainmultisigSession) SetConfig(signerAddresses []common.Address, signerGroups []uint8, groupQuorums [32]uint8, groupParents [32]uint8, clearRoot bool) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.SetConfig(&_Manychainmultisig.TransactOpts, signerAddresses, signerGroups, groupQuorums, groupParents, clearRoot)
}

// SetConfig is a paid mutator transaction binding the contract method 0x846c67ef.
//
// Solidity: function setConfig(address[] signerAddresses, uint8[] signerGroups, uint8[32] groupQuorums, uint8[32] groupParents, bool clearRoot) returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) SetConfig(signerAddresses []common.Address, signerGroups []uint8, groupQuorums [32]uint8, groupParents [32]uint8, clearRoot bool) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.SetConfig(&_Manychainmultisig.TransactOpts, signerAddresses, signerGroups, groupQuorums, groupParents, clearRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0x7cc38b28.
//
// Solidity: function setRoot(bytes32 root, uint32 validUntil, (uint256,address,uint40,uint40,bool) metadata, bytes32[] metadataProof, (uint8,bytes32,bytes32)[] signatures) returns()
func (_Manychainmultisig *ManychainmultisigTransactor) SetRoot(opts *bind.TransactOpts, root [32]byte, validUntil uint32, metadata ManyChainMultiSigRootMetadata, metadataProof [][32]byte, signatures []ManyChainMultiSigSignature) (*types.Transaction, error) {
	return _Manychainmultisig.contract.Transact(opts, "setRoot", root, validUntil, metadata, metadataProof, signatures)
}

// SetRoot is a paid mutator transaction binding the contract method 0x7cc38b28.
//
// Solidity: function setRoot(bytes32 root, uint32 validUntil, (uint256,address,uint40,uint40,bool) metadata, bytes32[] metadataProof, (uint8,bytes32,bytes32)[] signatures) returns()
func (_Manychainmultisig *ManychainmultisigSession) SetRoot(root [32]byte, validUntil uint32, metadata ManyChainMultiSigRootMetadata, metadataProof [][32]byte, signatures []ManyChainMultiSigSignature) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.SetRoot(&_Manychainmultisig.TransactOpts, root, validUntil, metadata, metadataProof, signatures)
}

// SetRoot is a paid mutator transaction binding the contract method 0x7cc38b28.
//
// Solidity: function setRoot(bytes32 root, uint32 validUntil, (uint256,address,uint40,uint40,bool) metadata, bytes32[] metadataProof, (uint8,bytes32,bytes32)[] signatures) returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) SetRoot(root [32]byte, validUntil uint32, metadata ManyChainMultiSigRootMetadata, metadataProof [][32]byte, signatures []ManyChainMultiSigSignature) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.SetRoot(&_Manychainmultisig.TransactOpts, root, validUntil, metadata, metadataProof, signatures)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manychainmultisig *ManychainmultisigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Manychainmultisig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manychainmultisig *ManychainmultisigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.TransferOwnership(&_Manychainmultisig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manychainmultisig.Contract.TransferOwnership(&_Manychainmultisig.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Manychainmultisig *ManychainmultisigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manychainmultisig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Manychainmultisig *ManychainmultisigSession) Receive() (*types.Transaction, error) {
	return _Manychainmultisig.Contract.Receive(&_Manychainmultisig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Manychainmultisig *ManychainmultisigTransactorSession) Receive() (*types.Transaction, error) {
	return _Manychainmultisig.Contract.Receive(&_Manychainmultisig.TransactOpts)
}

// ManychainmultisigConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the Manychainmultisig contract.
type ManychainmultisigConfigSetIterator struct {
	Event *ManychainmultisigConfigSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ManychainmultisigConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManychainmultisigConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ManychainmultisigConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ManychainmultisigConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManychainmultisigConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManychainmultisigConfigSet represents a ConfigSet event raised by the Manychainmultisig contract.
type ManychainmultisigConfigSet struct {
	Config        ManyChainMultiSigConfig
	IsRootCleared bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x0a4974ad206b9c736f9ab2feac1c9b1d043fe4ef377c70ae45659f2ef089f03e.
//
// Solidity: event ConfigSet(((address,uint8,uint8)[],uint8[32],uint8[32]) config, bool isRootCleared)
func (_Manychainmultisig *ManychainmultisigFilterer) FilterConfigSet(opts *bind.FilterOpts) (*ManychainmultisigConfigSetIterator, error) {

	logs, sub, err := _Manychainmultisig.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigConfigSetIterator{contract: _Manychainmultisig.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x0a4974ad206b9c736f9ab2feac1c9b1d043fe4ef377c70ae45659f2ef089f03e.
//
// Solidity: event ConfigSet(((address,uint8,uint8)[],uint8[32],uint8[32]) config, bool isRootCleared)
func (_Manychainmultisig *ManychainmultisigFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ManychainmultisigConfigSet) (event.Subscription, error) {

	logs, sub, err := _Manychainmultisig.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManychainmultisigConfigSet)
				if err := _Manychainmultisig.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigSet is a log parse operation binding the contract event 0x0a4974ad206b9c736f9ab2feac1c9b1d043fe4ef377c70ae45659f2ef089f03e.
//
// Solidity: event ConfigSet(((address,uint8,uint8)[],uint8[32],uint8[32]) config, bool isRootCleared)
func (_Manychainmultisig *ManychainmultisigFilterer) ParseConfigSet(log types.Log) (*ManychainmultisigConfigSet, error) {
	event := new(ManychainmultisigConfigSet)
	if err := _Manychainmultisig.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManychainmultisigNewRootIterator is returned from FilterNewRoot and is used to iterate over the raw logs and unpacked data for NewRoot events raised by the Manychainmultisig contract.
type ManychainmultisigNewRootIterator struct {
	Event *ManychainmultisigNewRoot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ManychainmultisigNewRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManychainmultisigNewRoot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ManychainmultisigNewRoot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ManychainmultisigNewRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManychainmultisigNewRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManychainmultisigNewRoot represents a NewRoot event raised by the Manychainmultisig contract.
type ManychainmultisigNewRoot struct {
	Root       [32]byte
	ValidUntil uint32
	Metadata   ManyChainMultiSigRootMetadata
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRoot is a free log retrieval operation binding the contract event 0x7ea643ae44677f24e0d6f40168893712daaf729b0a38fe7702d21cb544c84101.
//
// Solidity: event NewRoot(bytes32 indexed root, uint32 validUntil, (uint256,address,uint40,uint40,bool) metadata)
func (_Manychainmultisig *ManychainmultisigFilterer) FilterNewRoot(opts *bind.FilterOpts, root [][32]byte) (*ManychainmultisigNewRootIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Manychainmultisig.contract.FilterLogs(opts, "NewRoot", rootRule)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigNewRootIterator{contract: _Manychainmultisig.contract, event: "NewRoot", logs: logs, sub: sub}, nil
}

// WatchNewRoot is a free log subscription operation binding the contract event 0x7ea643ae44677f24e0d6f40168893712daaf729b0a38fe7702d21cb544c84101.
//
// Solidity: event NewRoot(bytes32 indexed root, uint32 validUntil, (uint256,address,uint40,uint40,bool) metadata)
func (_Manychainmultisig *ManychainmultisigFilterer) WatchNewRoot(opts *bind.WatchOpts, sink chan<- *ManychainmultisigNewRoot, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Manychainmultisig.contract.WatchLogs(opts, "NewRoot", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManychainmultisigNewRoot)
				if err := _Manychainmultisig.contract.UnpackLog(event, "NewRoot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRoot is a log parse operation binding the contract event 0x7ea643ae44677f24e0d6f40168893712daaf729b0a38fe7702d21cb544c84101.
//
// Solidity: event NewRoot(bytes32 indexed root, uint32 validUntil, (uint256,address,uint40,uint40,bool) metadata)
func (_Manychainmultisig *ManychainmultisigFilterer) ParseNewRoot(log types.Log) (*ManychainmultisigNewRoot, error) {
	event := new(ManychainmultisigNewRoot)
	if err := _Manychainmultisig.contract.UnpackLog(event, "NewRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManychainmultisigOpExecutedIterator is returned from FilterOpExecuted and is used to iterate over the raw logs and unpacked data for OpExecuted events raised by the Manychainmultisig contract.
type ManychainmultisigOpExecutedIterator struct {
	Event *ManychainmultisigOpExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ManychainmultisigOpExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManychainmultisigOpExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ManychainmultisigOpExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ManychainmultisigOpExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManychainmultisigOpExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManychainmultisigOpExecuted represents a OpExecuted event raised by the Manychainmultisig contract.
type ManychainmultisigOpExecuted struct {
	Nonce *big.Int
	To    common.Address
	Data  []byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOpExecuted is a free log retrieval operation binding the contract event 0x87d58fdd48be753fb9ef4ec8a5895086c401506da8b4d752abc90602c3e62d1d.
//
// Solidity: event OpExecuted(uint40 indexed nonce, address to, bytes data, uint256 value)
func (_Manychainmultisig *ManychainmultisigFilterer) FilterOpExecuted(opts *bind.FilterOpts, nonce []*big.Int) (*ManychainmultisigOpExecutedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Manychainmultisig.contract.FilterLogs(opts, "OpExecuted", nonceRule)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigOpExecutedIterator{contract: _Manychainmultisig.contract, event: "OpExecuted", logs: logs, sub: sub}, nil
}

// WatchOpExecuted is a free log subscription operation binding the contract event 0x87d58fdd48be753fb9ef4ec8a5895086c401506da8b4d752abc90602c3e62d1d.
//
// Solidity: event OpExecuted(uint40 indexed nonce, address to, bytes data, uint256 value)
func (_Manychainmultisig *ManychainmultisigFilterer) WatchOpExecuted(opts *bind.WatchOpts, sink chan<- *ManychainmultisigOpExecuted, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Manychainmultisig.contract.WatchLogs(opts, "OpExecuted", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManychainmultisigOpExecuted)
				if err := _Manychainmultisig.contract.UnpackLog(event, "OpExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpExecuted is a log parse operation binding the contract event 0x87d58fdd48be753fb9ef4ec8a5895086c401506da8b4d752abc90602c3e62d1d.
//
// Solidity: event OpExecuted(uint40 indexed nonce, address to, bytes data, uint256 value)
func (_Manychainmultisig *ManychainmultisigFilterer) ParseOpExecuted(log types.Log) (*ManychainmultisigOpExecuted, error) {
	event := new(ManychainmultisigOpExecuted)
	if err := _Manychainmultisig.contract.UnpackLog(event, "OpExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManychainmultisigOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Manychainmultisig contract.
type ManychainmultisigOwnershipTransferStartedIterator struct {
	Event *ManychainmultisigOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ManychainmultisigOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManychainmultisigOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ManychainmultisigOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ManychainmultisigOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManychainmultisigOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManychainmultisigOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Manychainmultisig contract.
type ManychainmultisigOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Manychainmultisig *ManychainmultisigFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ManychainmultisigOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manychainmultisig.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigOwnershipTransferStartedIterator{contract: _Manychainmultisig.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Manychainmultisig *ManychainmultisigFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ManychainmultisigOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manychainmultisig.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManychainmultisigOwnershipTransferStarted)
				if err := _Manychainmultisig.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Manychainmultisig *ManychainmultisigFilterer) ParseOwnershipTransferStarted(log types.Log) (*ManychainmultisigOwnershipTransferStarted, error) {
	event := new(ManychainmultisigOwnershipTransferStarted)
	if err := _Manychainmultisig.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManychainmultisigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Manychainmultisig contract.
type ManychainmultisigOwnershipTransferredIterator struct {
	Event *ManychainmultisigOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ManychainmultisigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManychainmultisigOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ManychainmultisigOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ManychainmultisigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManychainmultisigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManychainmultisigOwnershipTransferred represents a OwnershipTransferred event raised by the Manychainmultisig contract.
type ManychainmultisigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manychainmultisig *ManychainmultisigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ManychainmultisigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manychainmultisig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManychainmultisigOwnershipTransferredIterator{contract: _Manychainmultisig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manychainmultisig *ManychainmultisigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ManychainmultisigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manychainmultisig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManychainmultisigOwnershipTransferred)
				if err := _Manychainmultisig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manychainmultisig *ManychainmultisigFilterer) ParseOwnershipTransferred(log types.Log) (*ManychainmultisigOwnershipTransferred, error) {
	event := new(ManychainmultisigOwnershipTransferred)
	if err := _Manychainmultisig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
