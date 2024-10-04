// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type InternalEVM2AnyRampMessage struct {
	Header         InternalRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenAmounts   []InternalEVM2AnyTokenTransfer
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type OnRampAllowListConfigArgs struct {
	DestChainSelector         uint64
	AllowListEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	AllowListEnabled  bool
}

type OnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	MessageInterceptor     common.Address
	FeeAggregator          common.Address
	AllowListAdmin         common.Address
}

type OnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRMNV2\",\"name\":\"rmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"messageInterceptor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"}],\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidAllowListRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"name\":\"AllowListAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"AllowListSendersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"AllowListSendersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourcePoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"}],\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPMessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRMNV2\",\"name\":\"rmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"messageInterceptor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"}],\"name\":\"DestChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeTokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\"}],\"internalType\":\"structOnRamp.AllowListConfigArgs[]\",\"name\":\"allowListConfigArgsItems\",\"type\":\"tuple[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"}],\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getAllowedSendersList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"messageInterceptor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPoolV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRMNV2\",\"name\":\"rmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"messageInterceptor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620042c8380380620042c883398101604081905262000035916200076e565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000186565b505083516001600160401b031615905080620000e6575060208301516001600160a01b0316155b80620000fd575060408301516001600160a01b0316155b8062000114575060608301516001600160a01b0316155b1562000133576040516306b7c75960e31b815260040160405180910390fd5b82516001600160401b031660805260208301516001600160a01b0390811660a0526040840151811660c05260608401511660e052620001728262000231565b6200017d81620003e0565b5050506200086f565b336001600160a01b03821603620001e05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0316158062000254575060608101516001600160a01b0316155b8062000261575080602001515b1562000280576040516306b7c75960e31b815260040160405180910390fd5b8051600280546020808501511515600160a01b026001600160a81b03199092166001600160a01b039485161791909117909155604080840151600380549185166001600160a01b0319928316179055606080860151600480549187169184169190911790556080808701516005805491881691909416179092558251808301845291516001600160401b0316825260a05185169382019390935260c05184168183015260e05190931691830191909152517fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f191620003d591849082516001600160401b031681526020808401516001600160a01b03908116828401526040858101518216818501526060958601518216868501528451821660808086019190915292850151151560a0850152840151811660c084015293830151841660e0830152909101519091166101008201526101200190565b60405180910390a150565b60005b81518110156200053057600082828151811062000404576200040462000859565b60200260200101519050600083838151811062000425576200042562000859565b6020026020010151600001519050806001600160401b03166000036200046a5760405163c35aa79d60e01b81526001600160401b038216600482015260240162000083565b6001600160401b0381811660008181526006602090815260409182902086820151815488850151600160401b600160e81b031990911669010000000000000000006001600160a01b0390931692830260ff60401b19161768010000000000000000911515820217808455855197811688529387019190915260ff920491909116151591840191909152917fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef59060600160405180910390a2505050806001019050620003e3565b5050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156200056f576200056f62000534565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620005a057620005a062000534565b604052919050565b80516001600160401b0381168114620005c057600080fd5b919050565b6001600160a01b0381168114620005db57600080fd5b50565b80518015158114620005c057600080fd5b600060a082840312156200060257600080fd5b60405160a081016001600160401b038111828210171562000627576200062762000534565b806040525080915082516200063c81620005c5565b81526200064c60208401620005de565b602082015260408301516200066181620005c5565b604082015260608301516200067681620005c5565b606082015260808301516200068b81620005c5565b6080919091015292915050565b600082601f830112620006aa57600080fd5b815160206001600160401b03821115620006c857620006c862000534565b620006d8818360051b0162000575565b82815260609283028501820192828201919087851115620006f857600080fd5b8387015b85811015620007615781818a031215620007165760008081fd5b620007206200054a565b6200072b82620005a8565b8152858201516200073c81620005c5565b8187015260406200074f838201620005de565b908201528452928401928101620006fc565b5090979650505050505050565b60008060008385036101408112156200078657600080fd5b60808112156200079557600080fd5b50604051608081016001600160401b038082118383101715620007bc57620007bc62000534565b81604052620007cb87620005a8565b835260208701519150620007df82620005c5565b81602084015260408701519150620007f782620005c5565b816040840152606087015191506200080f82620005c5565b816060840152829550620008278860808901620005ef565b94506101208701519250808311156200083f57600080fd5b50506200084f8682870162000698565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b60805160a05160c05160e0516139e0620008e86000396000818161021701528181610cfd0152611bf50152600081816101db015281816115c40152611bce01526000818161019f015281816105f40152611ba401526000818161016f0152818161114b015281816116e10152611b8001526139e06000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80637437ff9f116100b2578063972b461211610081578063df0aa9e911610066578063df0aa9e914610557578063f2fde38b1461056a578063fbca3b741461057d57600080fd5b8063972b4612146104ea578063d77d5ed01461050a57600080fd5b80637437ff9f146103ee57806379ba5097146104985780638da5cb5b146104a05780639041be3d146104be57600080fd5b806327e936f1116101095780633a019940116100ee5780633a0199401461030f57806348a98aa4146103175780636def4ce71461034f57600080fd5b806327e936f1146102e957806334adf494146102fc57600080fd5b806306285c691461013b578063181f5a771461026a57806320487ded146102b35780632716072b146102d4575b600080fd5b61025460408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b60405161026191906126eb565b60405180910390f35b6102a66040518060400160405280601081526020017f4f6e52616d7020312e362e302d6465760000000000000000000000000000000081525081565b60405161026191906127b0565b6102c66102c13660046127f1565b610590565b604051908152602001610261565b6102e76102e236600461295f565b610749565b005b6102e76102f7366004612a4d565b61075d565b6102e761030a366004612ae5565b61076e565b6102e7610a8a565b61032a610325366004612b5a565b610cb5565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610261565b6103b261035d366004612b93565b67ffffffffffffffff9081166000908152600660205260409020549081169168010000000000000000820460ff16916901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1690565b6040805167ffffffffffffffff9094168452911515602084015273ffffffffffffffffffffffffffffffffffffffff1690820152606001610261565b61048b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a08101825260025473ffffffffffffffffffffffffffffffffffffffff80821683527401000000000000000000000000000000000000000090910460ff161515602083015260035481169282019290925260045482166060820152600554909116608082015290565b6040516102619190612bb0565b6102e7610d6a565b60005473ffffffffffffffffffffffffffffffffffffffff1661032a565b6104d16104cc366004612b93565b610e67565b60405167ffffffffffffffff9091168152602001610261565b6104fd6104f8366004612b93565b610e90565b6040516102619190612c09565b61032a610518366004612b93565b67ffffffffffffffff166000908152600660205260409020546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1690565b6102c6610565366004612c63565b610eb8565b6102e7610578366004612ccf565b6117bc565b6104fd61058b366004612b93565b6117cd565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b16600482015260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690632cbc26bb90602401602060405180830381865afa15801561063b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065f9190612cec565b156106a7576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6002546040517fd8694ccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d8694ccd906106ff9086908690600401612e1d565b602060405180830381865afa15801561071c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107409190612f66565b90505b92915050565b610751611801565b61075a81611884565b50565b610765611801565b61075a81611a27565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107de5760055473ffffffffffffffffffffffffffffffffffffffff1633146107de576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610a855760008383838181106107fd576107fd612f7f565b905060200281019061080f9190612fae565b6108189061305f565b805167ffffffffffffffff1660009081526006602090815260409182902090830151815490151568010000000000000000027fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff909116178155908201515191925090156109d8578160200151156109975760005b826040015151811015610947576000836040015182815181106108b1576108b1612f7f565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109305783516040517f463258ff00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161069e565b61093d6001840182611c57565b505060010161088c565b50816000015167ffffffffffffffff167f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281836040015160405161098a9190612c09565b60405180910390a26109d8565b81516040517f463258ff00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161069e565b60005b826060015151811015610a2457610a1b83606001518281518110610a0157610a01612f7f565b602002602001015183600101611c7990919063ffffffff16565b506001016109db565b5060608201515115610a7b57816000015167ffffffffffffffff167fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d4215868360600151604051610a729190612c09565b60405180910390a25b50506001016107e1565b505050565b600254604080517fcdc73d51000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cdc73d5191600480830192869291908290030181865afa158015610af9573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610b3f9190810190613110565b60045490915073ffffffffffffffffffffffffffffffffffffffff1660005b8251811015610a85576000838281518110610b7b57610b7b612f7f565b60209081029190910101516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa158015610bf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1a9190612f66565b90508015610cab57610c4373ffffffffffffffffffffffffffffffffffffffff83168583611c9b565b8173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e83604051610ca291815260200190565b60405180910390a35b5050600101610b5e565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bbe4f6db90602401602060405180830381865afa158015610d46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610740919061319f565b60015473ffffffffffffffffffffffffffffffffffffffff163314610deb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161069e565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff8082166000908152600660205260408120549091610743911660016131eb565b67ffffffffffffffff8116600090815260066020526040902060609061074390600101611d28565b60025460009074010000000000000000000000000000000000000000900460ff1615610f10576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905567ffffffffffffffff8516600090815260066020526040902073ffffffffffffffffffffffffffffffffffffffff8316610fb5576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805468010000000000000000900460ff161561102657610fd86001820184611d3c565b611026576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161069e565b80546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611083576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035473ffffffffffffffffffffffffffffffffffffffff168015611129576040517fe0a0e50600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063e0a0e506906110f6908a908a90600401612e1d565b600060405180830381600087803b15801561111057600080fd5b505af1158015611124573d6000803e3d6000fd5b505050505b50604080516101c081019091526000610120820181815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116610140850152898116610160850152845492939283929161018084019187918791611197911661320c565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff168152602001600067ffffffffffffffff1681525081526020018573ffffffffffffffffffffffffffffffffffffffff16815260200187806020019061120b9190613233565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161124f8880613233565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250938552505060408051602081810183529381529284019290925250016112ab6080890160608a01612ccf565b73ffffffffffffffffffffffffffffffffffffffff168152602001868152602001600081526020018780604001906112e39190613298565b905067ffffffffffffffff8111156112fd576112fd612841565b60405190808252806020026020018201604052801561137657816020015b6113636040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160008152602001606081525090565b81526020019060019003908161131b5790505b5090529050600061138a6040880188613298565b808060200260200160405190810160405280939291908181526020016000905b828210156113d6576113c760408302860136819003810190613300565b815260200190600101906113aa565b5050505050905060005b6113ed6040890189613298565b90508110156114875761145d82828151811061140b5761140b612f7f565b60209081029190910101518a6114218b80613233565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250611d6b915050565b836101000151828151811061147457611474612f7f565b60209081029190910101526001016113e0565b50600254600090606090819073ffffffffffffffffffffffffffffffffffffffff1663430d138c8c6114be60808e018e8601612ccf565b8c8e80608001906114cf9190613233565b8b61010001518b6040518863ffffffff1660e01b81526004016114f8979695949392919061341d565b600060405180830381865afa158015611515573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261155b919081019061354f565b60e08901939093529094509250905082611636576040517fea458c0c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c16600482015273ffffffffffffffffffffffffffffffffffffffff89811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063ea458c0c906044016020604051808303816000875af115801561160d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116319190613642565b611639565b60005b855167ffffffffffffffff909116608091820152850182905260005b856101000151518110156116ab5781818151811061167557611675612f7f565b6020026020010151866101000151828151811061169457611694612f7f565b602090810291909101015160800152600101611655565b50604080517f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908c16606082015230608082015261173b90869060a00160405160208183030381529060405280519060200120612082565b85515260405167ffffffffffffffff8c16907f8cd775d4a25bd349439a70817fd110144d6ab229ae1b9f54a1e5777d2041bfed9061177a90889061365f565b60405180910390a25050600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff169055505051519150505b949350505050565b6117c4611801565b61075a816121d4565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005473ffffffffffffffffffffffffffffffffffffffff163314611882576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161069e565b565b60005b8151811015611a235760008282815181106118a4576118a4612f7f565b6020026020010151905060008383815181106118c2576118c2612f7f565b60200260200101516000015190508067ffffffffffffffff16600003611920576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8216600482015260240161069e565b67ffffffffffffffff818116600081815260066020908152604091829020868201518154888501517fffffff000000000000000000000000000000000000000000ffffffffffffffff909116690100000000000000000073ffffffffffffffffffffffffffffffffffffffff9093169283027fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff161768010000000000000000911515820217808455855197811688529387019190915260ff920491909116151591840191909152917fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef59060600160405180910390a2505050806001019050611887565b5050565b805173ffffffffffffffffffffffffffffffffffffffff161580611a635750606081015173ffffffffffffffffffffffffffffffffffffffff16155b80611a6f575080602001515b15611aa6576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160028054602080850151151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00000000000000000000000000000000000000000090921673ffffffffffffffffffffffffffffffffffffffff9485161791909117909155604080840151600380549185167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055606080860151600480549187169184169190911790556080808701516005805491881691909416179092558251918201835267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001682527f00000000000000000000000000000000000000000000000000000000000000008516938201939093527f00000000000000000000000000000000000000000000000000000000000000008416818301527f000000000000000000000000000000000000000000000000000000000000000090931691830191909152517fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f191611c4c9184906137b7565b60405180910390a150565b60006107408373ffffffffffffffffffffffffffffffffffffffff84166122c9565b60006107408373ffffffffffffffffffffffffffffffffffffffff8416612318565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610a85908490612412565b60606000611d358361251e565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610740565b611db36040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160008152602001606081525090565b8460200151600003611df1576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611e01858760000151610cb5565b905073ffffffffffffffffffffffffffffffffffffffff81161580611ed157506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf00000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015611eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecf9190612cec565b155b15611f235785516040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260240161069e565b60008173ffffffffffffffffffffffffffffffffffffffff16639a4575b96040518060a001604052808881526020018967ffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018a6020015181526020018a6000015173ffffffffffffffffffffffffffffffffffffffff168152506040518263ffffffff1660e01b8152600401611fc29190613865565b6000604051808303816000875af1158015611fe1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261202791908101906138db565b6040805160a08101825273ffffffffffffffffffffffffffffffffffffffff90941684528151602080860191909152918201518482015288820151606085015280519182019052600081526080830152509050949350505050565b60008060001b8284602001518560000151606001518660000151608001518760a001518860c0015160405160200161210095949392919073ffffffffffffffffffffffffffffffffffffffff958616815267ffffffffffffffff94851660208201529290931660408301529092166060830152608082015260a00190565b60405160208183030381529060405280519060200120856060015180519060200120866040015180519060200120876101000151604051602001612144919061396c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201206080808d0151805190840120928501999099529183019690965260608201949094529485019190915260a084015260c083015260e08201526101000160405160208183030381529060405280519060200120905092915050565b3373ffffffffffffffffffffffffffffffffffffffff821603612253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161069e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600081815260018301602052604081205461231057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610743565b506000610743565b6000818152600183016020526040812054801561240157600061233c60018361397f565b85549091506000906123509060019061397f565b90508082146123b557600086600001828154811061237057612370612f7f565b906000526020600020015490508087600001848154811061239357612393612f7f565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806123c6576123c6613992565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610743565b6000915050610743565b5092915050565b6000612474826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661257a9092919063ffffffff16565b805190915015610a8557808060200190518101906124929190612cec565b610a85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161069e565b60608160000180548060200260200160405190810160405280929190818152602001828054801561256e57602002820191906000526020600020905b81548152602001906001019080831161255a575b50505050509050919050565b60606117b48484600085856000808673ffffffffffffffffffffffffffffffffffffffff1685876040516125ae91906139c1565b60006040518083038185875af1925050503d80600081146125eb576040519150601f19603f3d011682016040523d82523d6000602084013e6125f0565b606091505b50915091506126018783838761260c565b979650505050505050565b606083156126a257825160000361269b5773ffffffffffffffffffffffffffffffffffffffff85163b61269b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161069e565b50816117b4565b6117b483838151156126b75781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069e91906127b0565b60808101610743828467ffffffffffffffff8151168252602081015173ffffffffffffffffffffffffffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b60005b8381101561275d578181015183820152602001612745565b50506000910152565b6000815180845261277e816020860160208601612742565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006107406020830184612766565b67ffffffffffffffff8116811461075a57600080fd5b600060a082840312156127eb57600080fd5b50919050565b6000806040838503121561280457600080fd5b823561280f816127c3565b9150602083013567ffffffffffffffff81111561282b57600080fd5b612837858286016127d9565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561289357612893612841565b60405290565b6040805190810167ffffffffffffffff8111828210171561289357612893612841565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561290357612903612841565b604052919050565b600067ffffffffffffffff82111561292557612925612841565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff8116811461075a57600080fd5b801515811461075a57600080fd5b6000602080838503121561297257600080fd5b823567ffffffffffffffff81111561298957600080fd5b8301601f8101851361299a57600080fd5b80356129ad6129a88261290b565b6128bc565b818152606091820283018401918482019190888411156129cc57600080fd5b938501935b83851015612a315780858a0312156129e95760008081fd5b6129f1612870565b85356129fc816127c3565b815285870135612a0b8161292f565b81880152604086810135612a1e81612951565b90820152835293840193918501916129d1565b50979650505050505050565b8035612a488161292f565b919050565b600060a08284031215612a5f57600080fd5b60405160a0810181811067ffffffffffffffff82111715612a8257612a82612841565b6040528235612a908161292f565b81526020830135612aa081612951565b60208201526040830135612ab38161292f565b60408201526060830135612ac68161292f565b60608201526080830135612ad98161292f565b60808201529392505050565b60008060208385031215612af857600080fd5b823567ffffffffffffffff80821115612b1057600080fd5b818501915085601f830112612b2457600080fd5b813581811115612b3357600080fd5b8660208260051b8501011115612b4857600080fd5b60209290920196919550909350505050565b60008060408385031215612b6d57600080fd5b8235612b78816127c3565b91506020830135612b888161292f565b809150509250929050565b600060208284031215612ba557600080fd5b8135611d35816127c3565b60a08101610743828473ffffffffffffffffffffffffffffffffffffffff808251168352602082015115156020840152806040830151166040840152806060830151166060840152806080830151166080840152505050565b6020808252825182820181905260009190848201906040850190845b81811015612c5757835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612c25565b50909695505050505050565b60008060008060808587031215612c7957600080fd5b8435612c84816127c3565b9350602085013567ffffffffffffffff811115612ca057600080fd5b612cac878288016127d9565b935050604085013591506060850135612cc48161292f565b939692955090935050565b600060208284031215612ce157600080fd5b8135611d358161292f565b600060208284031215612cfe57600080fd5b8151611d3581612951565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612d3e57600080fd5b830160208101925035905067ffffffffffffffff811115612d5e57600080fd5b803603821315612d6d57600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b85811015612e12578135612de08161292f565b73ffffffffffffffffffffffffffffffffffffffff168752818301358388015260409687019690910190600101612dcd565b509495945050505050565b600067ffffffffffffffff808516835260406020840152612e3e8485612d09565b60a06040860152612e5360e086018284612d74565b915050612e636020860186612d09565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878503016060880152612e99848385612d74565b9350604088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018312612ed257600080fd5b60209288019283019235915084821115612eeb57600080fd5b8160061b3603831315612efd57600080fd5b80878503016080880152612f12848385612dbd565b9450612f2060608901612a3d565b73ffffffffffffffffffffffffffffffffffffffff811660a08901529350612f4b6080890189612d09565b94509250808786030160c08801525050612601838383612d74565b600060208284031215612f7857600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112612fe257600080fd5b9190910192915050565b600082601f830112612ffd57600080fd5b8135602061300d6129a88361290b565b8083825260208201915060208460051b87010193508684111561302f57600080fd5b602086015b848110156130545780356130478161292f565b8352918301918301613034565b509695505050505050565b60006080823603121561307157600080fd5b6040516080810167ffffffffffffffff828210818311171561309557613095612841565b81604052843591506130a6826127c3565b9082526020840135906130b882612951565b81602084015260408501359150808211156130d257600080fd5b6130de36838701612fec565b604084015260608501359150808211156130f757600080fd5b5061310436828601612fec565b60608301525092915050565b6000602080838503121561312357600080fd5b825167ffffffffffffffff81111561313a57600080fd5b8301601f8101851361314b57600080fd5b80516131596129a88261290b565b81815260059190911b8201830190838101908783111561317857600080fd5b928401925b828410156126015783516131908161292f565b8252928401929084019061317d565b6000602082840312156131b157600080fd5b8151611d358161292f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff81811683821601908082111561240b5761240b6131bc565b600067ffffffffffffffff808316818103613229576132296131bc565b6001019392505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261326857600080fd5b83018035915067ffffffffffffffff82111561328357600080fd5b602001915036819003821315612d6d57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126132cd57600080fd5b83018035915067ffffffffffffffff8211156132e857600080fd5b6020019150600681901b3603821315612d6d57600080fd5b60006040828403121561331257600080fd5b61331a612899565b82356133258161292f565b81526020928301359281019290925250919050565b600082825180855260208086019550808260051b84010181860160005b84811015613410577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018952815160a073ffffffffffffffffffffffffffffffffffffffff82511685528582015181878701526133b982870182612766565b915050604080830151868303828801526133d38382612766565b925050506060808301518187015250608080830151925085820381870152506133fc8183612766565b9a86019a9450505090830190600101613357565b5090979650505050505050565b67ffffffffffffffff881681526000602073ffffffffffffffffffffffffffffffffffffffff808a1682850152604089604086015260c0606086015261346760c08601898b612d74565b8581036080870152613479818961333a565b86810360a0880152875180825285890192509085019060005b818110156134b9578351805187168452870151878401529286019291840191600101613492565b50909e9d5050505050505050505050505050565b600082601f8301126134de57600080fd5b815167ffffffffffffffff8111156134f8576134f8612841565b61352960207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016128bc565b81815284602083860101111561353e57600080fd5b6117b4826020830160208701612742565b6000806000806080858703121561356557600080fd5b8451935060208086015161357881612951565b604087015190945067ffffffffffffffff8082111561359657600080fd5b6135a289838a016134cd565b945060608801519150808211156135b857600080fd5b818801915088601f8301126135cc57600080fd5b81516135da6129a88261290b565b81815260059190911b8301840190848101908b8311156135f957600080fd5b8585015b83811015613631578051858111156136155760008081fd5b6136238e89838a01016134cd565b8452509186019186016135fd565b50989b979a50959850505050505050565b60006020828403121561365457600080fd5b8151611d35816127c3565b602081526136b060208201835180518252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015280608084015116608085015250505050565b600060208301516136d960c084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060408301516101a08060e08501526136f66101c0850183612766565b915060608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06101008187860301818801526137348584612766565b94506080880151925081878603016101208801526137528584612766565b945060a0880151925061377e61014088018473ffffffffffffffffffffffffffffffffffffffff169052565b60c088015161016088015260e08801516101808801528701518685039091018387015290506137ad838261333a565b9695505050505050565b610120810161380f828567ffffffffffffffff8151168252602081015173ffffffffffffffffffffffffffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b825173ffffffffffffffffffffffffffffffffffffffff9081166080848101919091526020850151151560a08501526040850151821660c08501526060850151821660e085015284015116610100830152611d35565b602081526000825160a0602084015261388160c0840182612766565b905067ffffffffffffffff6020850151166040840152604084015173ffffffffffffffffffffffffffffffffffffffff8082166060860152606086015160808601528060808701511660a086015250508091505092915050565b6000602082840312156138ed57600080fd5b815167ffffffffffffffff8082111561390557600080fd5b908301906040828603121561391957600080fd5b613921612899565b82518281111561393057600080fd5b61393c878286016134cd565b82525060208301518281111561395157600080fd5b61395d878286016134cd565b60208301525095945050505050565b602081526000610740602083018461333a565b81810381811115610743576107436131bc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251612fe281846020870161274256fea164736f6c6343000818000a",
}

var OnRampABI = OnRampMetaData.ABI

var OnRampBin = OnRampMetaData.Bin

func DeployOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OnRampStaticConfig, dynamicConfig OnRampDynamicConfig, destChainConfigArgs []OnRampDestChainConfigArgs) (common.Address, *types.Transaction, *OnRamp, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRamp{address: address, abi: *parsed, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

type OnRamp struct {
	address common.Address
	abi     abi.ABI
	OnRampCaller
	OnRampTransactor
	OnRampFilterer
}

type OnRampCaller struct {
	contract *bind.BoundContract
}

type OnRampTransactor struct {
	contract *bind.BoundContract
}

type OnRampFilterer struct {
	contract *bind.BoundContract
}

type OnRampSession struct {
	Contract     *OnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampCallerSession struct {
	Contract *OnRampCaller
	CallOpts bind.CallOpts
}

type OnRampTransactorSession struct {
	Contract     *OnRampTransactor
	TransactOpts bind.TransactOpts
}

type OnRampRaw struct {
	Contract *OnRamp
}

type OnRampCallerRaw struct {
	Contract *OnRampCaller
}

type OnRampTransactorRaw struct {
	Contract *OnRampTransactor
}

func NewOnRamp(address common.Address, backend bind.ContractBackend) (*OnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRamp{address: address, abi: abi, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

func NewOnRampCaller(address common.Address, caller bind.ContractCaller) (*OnRampCaller, error) {
	contract, err := bindOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampCaller{contract: contract}, nil
}

func NewOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampTransactor, error) {
	contract, err := bindOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampTransactor{contract: contract}, nil
}

func NewOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampFilterer, error) {
	contract, err := bindOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampFilterer{contract: contract}, nil
}

func bindOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OnRamp *OnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.OnRampCaller.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transfer(opts)
}

func (_OnRamp *OnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transfer(opts)
}

func (_OnRamp *OnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCaller) GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowedSendersList", destChainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAllowedSendersList(destChainSelector uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowedSendersList(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetAllowedSendersList(destChainSelector uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowedSendersList(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AllowListEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_OnRamp *OnRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDynamicConfig)).(*OnRampDynamicConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRamp.Contract.GetDynamicConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRamp.Contract.GetDynamicConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextSequenceNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextSequenceNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRamp.Contract.GetFee(&_OnRamp.CallOpts, destChainSelector, message)
}

func (_OnRamp *OnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRamp.Contract.GetFee(&_OnRamp.CallOpts, destChainSelector, message)
}

func (_OnRamp *OnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPoolBySourceToken(&_OnRamp.CallOpts, arg0, sourceToken)
}

func (_OnRamp *OnRampCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPoolBySourceToken(&_OnRamp.CallOpts, arg0, sourceToken)
}

func (_OnRamp *OnRampCaller) GetRouter(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getRouter", destChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetRouter(destChainSelector uint64) (common.Address, error) {
	return _OnRamp.Contract.GetRouter(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetRouter(destChainSelector uint64) (common.Address, error) {
	return _OnRamp.Contract.GetRouter(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampStaticConfig)).(*OnRampStaticConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRamp.Contract.GetStaticConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRamp.Contract.GetStaticConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetSupportedTokens(&_OnRamp.CallOpts, arg0)
}

func (_OnRamp *OnRampCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetSupportedTokens(&_OnRamp.CallOpts, arg0)
}

func (_OnRamp *OnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRamp *OnRampSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OnRamp *OnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, allowListConfigArgsItems []OnRampAllowListConfigArgs) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "applyAllowListUpdates", allowListConfigArgsItems)
}

func (_OnRamp *OnRampSession) ApplyAllowListUpdates(allowListConfigArgsItems []OnRampAllowListConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyAllowListUpdates(&_OnRamp.TransactOpts, allowListConfigArgsItems)
}

func (_OnRamp *OnRampTransactorSession) ApplyAllowListUpdates(allowListConfigArgsItems []OnRampAllowListConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyAllowListUpdates(&_OnRamp.TransactOpts, allowListConfigArgsItems)
}

func (_OnRamp *OnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_OnRamp *OnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyDestChainConfigUpdates(&_OnRamp.TransactOpts, destChainConfigArgs)
}

func (_OnRamp *OnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyDestChainConfigUpdates(&_OnRamp.TransactOpts, destChainConfigArgs)
}

func (_OnRamp *OnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OnRamp *OnRampSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetDynamicConfig(&_OnRamp.TransactOpts, dynamicConfig)
}

func (_OnRamp *OnRampTransactorSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetDynamicConfig(&_OnRamp.TransactOpts, dynamicConfig)
}

func (_OnRamp *OnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRamp *OnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "withdrawFeeTokens")
}

func (_OnRamp *OnRampSession) WithdrawFeeTokens() (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawFeeTokens(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) WithdrawFeeTokens() (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawFeeTokens(&_OnRamp.TransactOpts)
}

type OnRampAllowListAdminSetIterator struct {
	Event *OnRampAllowListAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListAdminSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampAllowListAdminSet)
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

func (it *OnRampAllowListAdminSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListAdminSet struct {
	AllowListAdmin common.Address
	Raw            types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListAdminSet(opts *bind.FilterOpts, allowListAdmin []common.Address) (*OnRampAllowListAdminSetIterator, error) {

	var allowListAdminRule []interface{}
	for _, allowListAdminItem := range allowListAdmin {
		allowListAdminRule = append(allowListAdminRule, allowListAdminItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListAdminSet", allowListAdminRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListAdminSetIterator{contract: _OnRamp.contract, event: "AllowListAdminSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdminSet, allowListAdmin []common.Address) (event.Subscription, error) {

	var allowListAdminRule []interface{}
	for _, allowListAdminItem := range allowListAdmin {
		allowListAdminRule = append(allowListAdminRule, allowListAdminItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListAdminSet", allowListAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListAdminSet)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListAdminSet(log types.Log) (*OnRampAllowListAdminSet, error) {
	event := new(OnRampAllowListAdminSet)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowListSendersAddedIterator struct {
	Event *OnRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListSendersAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampAllowListSendersAdded)
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

func (it *OnRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListSendersAddedIterator{contract: _OnRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListSendersAdded)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListSendersAdded(log types.Log) (*OnRampAllowListSendersAdded, error) {
	event := new(OnRampAllowListSendersAdded)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowListSendersRemovedIterator struct {
	Event *OnRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListSendersRemoved)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampAllowListSendersRemoved)
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

func (it *OnRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListSendersRemovedIterator{contract: _OnRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListSendersRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*OnRampAllowListSendersRemoved, error) {
	event := new(OnRampAllowListSendersRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampCCIPMessageSentIterator struct {
	Event *OnRampCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampCCIPMessageSent)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampCCIPMessageSent)
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

func (it *OnRampCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *OnRampCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampCCIPMessageSent)
				if err := _OnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error) {
	event := new(OnRampCCIPMessageSent)
	if err := _OnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampConfigSetIterator struct {
	Event *OnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampConfigSet)
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

func (it *OnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampConfigSet struct {
	StaticConfig  OnRampStaticConfig
	DynamicConfig OnRampDynamicConfig
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampConfigSetIterator{contract: _OnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseConfigSet(log types.Log) (*OnRampConfigSet, error) {
	event := new(OnRampConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampDestChainConfigSetIterator struct {
	Event *OnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampDestChainConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampDestChainConfigSet)
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

func (it *OnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	AllowListEnabled  bool
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampDestChainConfigSetIterator{contract: _OnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampDestChainConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error) {
	event := new(OnRampDestChainConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeeTokenWithdrawnIterator struct {
	Event *OnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeeTokenWithdrawn)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampFeeTokenWithdrawn)
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

func (it *OnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampFeeTokenWithdrawnIterator{contract: _OnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeeTokenWithdrawn)
				if err := _OnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error) {
	event := new(OnRampFeeTokenWithdrawn)
	if err := _OnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferRequestedIterator struct {
	Event *OnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampOwnershipTransferRequested)
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

func (it *OnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferRequestedIterator{contract: _OnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferRequested)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error) {
	event := new(OnRampOwnershipTransferRequested)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferredIterator struct {
	Event *OnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(OnRampOwnershipTransferred)
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

func (it *OnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferredIterator{contract: _OnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferred)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error) {
	event := new(OnRampOwnershipTransferred)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	SequenceNumber   uint64
	AllowListEnabled bool
	Router           common.Address
}

func (_OnRamp *OnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRamp.abi.Events["AllowListAdminSet"].ID:
		return _OnRamp.ParseAllowListAdminSet(log)
	case _OnRamp.abi.Events["AllowListSendersAdded"].ID:
		return _OnRamp.ParseAllowListSendersAdded(log)
	case _OnRamp.abi.Events["AllowListSendersRemoved"].ID:
		return _OnRamp.ParseAllowListSendersRemoved(log)
	case _OnRamp.abi.Events["CCIPMessageSent"].ID:
		return _OnRamp.ParseCCIPMessageSent(log)
	case _OnRamp.abi.Events["ConfigSet"].ID:
		return _OnRamp.ParseConfigSet(log)
	case _OnRamp.abi.Events["DestChainConfigSet"].ID:
		return _OnRamp.ParseDestChainConfigSet(log)
	case _OnRamp.abi.Events["FeeTokenWithdrawn"].ID:
		return _OnRamp.ParseFeeTokenWithdrawn(log)
	case _OnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _OnRamp.ParseOwnershipTransferRequested(log)
	case _OnRamp.abi.Events["OwnershipTransferred"].ID:
		return _OnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OnRampAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (OnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (OnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (OnRampCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f1")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5")
}

func (OnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (OnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRamp *OnRamp) Address() common.Address {
	return _OnRamp.address
}

type OnRampInterface interface {
	GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) ([]common.Address, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetRouter(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, allowListConfigArgsItems []OnRampAllowListConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowListAdmin []common.Address) (*OnRampAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdminSet, allowListAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*OnRampAllowListAdminSet, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*OnRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *OnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*OnRampAllowListSendersRemoved, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
