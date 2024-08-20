// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_multi_onramp

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

type EVM2EVMMultiOnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
}

type EVM2EVMMultiOnRampDynamicConfig struct {
	PriceRegistry    common.Address
	MessageValidator common.Address
	FeeAggregator    common.Address
	AllowListAdmin   common.Address
}

type EVM2EVMMultiOnRampStaticConfig struct {
	ChainSelector      uint64
	RmnProxy           common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

type InternalEVM2AnyRampMessage struct {
	Header         InternalRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	TokenAmounts   []InternalRampTokenAmount
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type InternalRampTokenAmount struct {
	SourcePoolAddress []byte
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
}

var EVM2EVMMultiOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowList\",\"type\":\"address[]\"}],\"name\":\"AllowListAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"name\":\"AllowListAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"AllowListDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"AllowListEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowList\",\"type\":\"address[]\"}],\"name\":\"AllowListRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"}],\"name\":\"DestChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeTokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"destinationChainSelectors\",\"type\":\"uint64[]\"}],\"name\":\"disableAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"destinationChainSelectors\",\"type\":\"uint64[]\"}],\"name\":\"enableAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"}],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"}],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPoolV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"name\":\"setAllowListAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200425a3803806200425a83398101604081905262000035916200066a565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000186565b505083516001600160401b031615905080620000e6575060208301516001600160a01b0316155b80620000fd575060408301516001600160a01b0316155b8062000114575060608301516001600160a01b0316155b1562000133576040516306b7c75960e31b815260040160405180910390fd5b82516001600160401b031660805260208301516001600160a01b0390811660a0526040840151811660c05260608401511660e052620001728262000231565b6200017d8162000390565b505050620007a4565b336001600160a01b03821603620001e05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0316158062000254575060408101516001600160a01b0316155b1562000273576040516306b7c75960e31b815260040160405180910390fd5b8051600280546001600160a01b03199081166001600160a01b0393841617909155602080840180516003805485169186169190911790556040808601805160048054871691881691909117905560608088018051600580549098169089161790965582516080808201855280516001600160401b031680835260a080518b16848a0190815260c080518d16868a0190815260e080518f169789019788528a5195865292518e169b85019b909b5299518c169783019790975292518a169381019390935289518916908301529351871693810193909352518516928201929092529151909216918101919091527f23a1adf8ad7fad6091a4803227af2cee848c01a7c812404cade7c25636925e32906101000160405180910390a150565b60005b8151811015620004d4576000828281518110620003b457620003b46200078e565b602002602001015190506000838381518110620003d557620003d56200078e565b6020026020010151600001519050806001600160401b03166000036200041a5760405163c35aa79d60e01b81526001600160401b038216600482015260240162000083565b6020828101516001600160401b038381166000818152600685526040908190208054600160481b600160e81b0319811669010000000000000000006001600160a01b03978816810291821793849055845192871691909616178152938104909416948301949094526801000000000000000090920460ff16151592810192909252907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef59060600160405180910390a2505060010162000393565b5050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620005135762000513620004d8565b60405290565b604080519081016001600160401b0381118282101715620005135762000513620004d8565b604051601f8201601f191681016001600160401b0381118282101715620005695762000569620004d8565b604052919050565b80516001600160401b03811681146200058957600080fd5b919050565b6001600160a01b0381168114620005a457600080fd5b50565b600082601f830112620005b957600080fd5b815160206001600160401b03821115620005d757620005d7620004d8565b620005e7818360051b016200053e565b82815260069290921b840181019181810190868411156200060757600080fd5b8286015b848110156200065f5760408189031215620006265760008081fd5b6200063062000519565b6200063b8262000571565b8152848201516200064c816200058e565b818601528352918301916040016200060b565b509695505050505050565b60008060008385036101208112156200068257600080fd5b60808112156200069157600080fd5b6200069b620004ee565b620006a68662000571565b81526020860151620006b8816200058e565b60208201526040860151620006cd816200058e565b60408201526060860151620006e2816200058e565b606082015293506080607f1982011215620006fc57600080fd5b5062000707620004ee565b608085015162000717816200058e565b815260a085015162000729816200058e565b602082015260c08501516200073e816200058e565b604082015260e085015162000753816200058e565b60608201526101008501519092506001600160401b038111156200077657600080fd5b6200078486828701620005a7565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b60805160a05160c05160e051613a3d6200081d6000396000818161026d015281816109cf0152611ccc015260008181610231015281816113120152611ca50152600081816101f5015281816106020152611c7b0152600081816101c501528181611238015281816116c80152611c4e0152613a3d6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80638da5cb5b116100d8578063b627f8ec1161008c578063f0f512f411610066578063f0f512f414610551578063f2fde38b14610564578063fbca3b741461057757600080fd5b8063b627f8ec146104de578063d77d5ed0146104f1578063df0aa9e91461053e57600080fd5b8063991e7e6e116100bd578063991e7e6e14610462578063a6f3ab6c14610482578063b2bd28691461049557600080fd5b80638da5cb5b146104185780639041be3d1461043657600080fd5b806348a98aa41161012f5780637437ff9f116101145780637437ff9f1461037d57806379ba5097146103fd57806380c162f71461040557600080fd5b806348a98aa414610332578063567261b61461036a57600080fd5b8063181f5a7711610160578063181f5a77146102c057806320487ded146103095780633a0199401461032a57600080fd5b80630242cf601461017c57806306285c6914610191575b600080fd5b61018f61018a366004612a18565b61058a565b005b6102aa60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b6040516102b79190612adb565b60405180910390f35b6102fc6040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b6040516102b79190612ba0565b61031c610317366004612bcb565b61059e565b6040519081526020016102b7565b61018f610757565b610345610340366004612c1b565b610987565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102b7565b61018f610378366004612ca0565b610a3c565b6103f0604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260025473ffffffffffffffffffffffffffffffffffffffff908116825260035481166020830152600454811692820192909252600554909116606082015290565b6040516102b79190612ce2565b61018f610bb9565b61018f610413366004612ca0565b610cb6565b60005473ffffffffffffffffffffffffffffffffffffffff16610345565b610449610444366004612d2b565b610e33565b60405167ffffffffffffffff90911681526020016102b7565b610475610470366004612d2b565b610e5c565b6040516102b79190612d48565b61018f610490366004612db2565b610e84565b6104ce6104a3366004612d2b565b67ffffffffffffffff1660009081526006602052604090205468010000000000000000900460ff1690565b60405190151581526020016102b7565b61018f6104ec366004612e37565b610e95565b6103456104ff366004612d2b565b67ffffffffffffffff166000908152600660205260409020546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1690565b61031c61054c366004612e54565b610f0c565b61018f61055f366004612ec0565b61177c565b61018f610572366004612e37565b611868565b610475610585366004612d2b565b611879565b6105926118ad565b61059b81611930565b50565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b16600482015260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690632cbc26bb90602401602060405180830381865afa158015610649573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066d9190612f53565b156106b5576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6002546040517fd8694ccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d8694ccd9061070d908690869060040161307b565b602060405180830381865afa15801561072a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074e91906131c4565b90505b92915050565b600254604080517fcdc73d51000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cdc73d5191600480830192869291908290030181865afa1580156107c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261080c91908101906131dd565b60045490915073ffffffffffffffffffffffffffffffffffffffff1660005b82518110156109825760008382815181106108485761084861326c565b60209081029190910101516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa1580156108c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e791906131c4565b905080156109785761091073ffffffffffffffffffffffffffffffffffffffff83168583611aac565b8173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e8360405161096f91815260200190565b60405180910390a35b505060010161082b565b505050565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bbe4f6db90602401602060405180830381865afa158015610a18573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074e919061329b565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590610a7c575060055473ffffffffffffffffffffffffffffffffffffffff163314155b15610ab3576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b67ffffffffffffffff81168211156109825760006006600085858567ffffffffffffffff16818110610aea57610aea61326c565b9050602002016020810190610aff9190612d2b565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160086101000a81548160ff0219169083151502179055507f50013bb4071fed8f5c37e8a9ee6ecc157fe969ed7e303069ecbbba1e89e1206483838367ffffffffffffffff16818110610b7857610b7861326c565b9050602002016020810190610b8d9190612d2b565b60405167ffffffffffffffff909116815260200160405180910390a1610bb2816132e7565b9050610ab6565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016106ac565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590610cf6575060055473ffffffffffffffffffffffffffffffffffffffff163314155b15610d2d576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b67ffffffffffffffff81168211156109825760016006600085858567ffffffffffffffff16818110610d6457610d6461326c565b9050602002016020810190610d799190612d2b565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160086101000a81548160ff0219169083151502179055507f4d8b8e1b49157081f6c6ec663e74f97d33537bf46858985e5eaefd364410be6083838367ffffffffffffffff16818110610df257610df261326c565b9050602002016020810190610e079190612d2b565b60405167ffffffffffffffff909116815260200160405180910390a1610e2c816132e7565b9050610d30565b67ffffffffffffffff80821660009081526006602052604081205490916107519116600161330e565b67ffffffffffffffff8116600090815260066020526040902060609061075190600101611b39565b610e8c6118ad565b61059b81611b4d565b610e9d6118ad565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e90600090a250565b67ffffffffffffffff84166000908152600660205260408120805468010000000000000000900460ff168015610f4a5750610f4a6001820184611d2f565b15610f99576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016106ac565b73ffffffffffffffffffffffffffffffffffffffff8316610fe6576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611043576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035473ffffffffffffffffffffffffffffffffffffffff1680156110e9576040517fe0a0e50600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063e0a0e506906110b6908a908a9060040161307b565b600060405180830381600087803b1580156110d057600080fd5b505af11580156110e4573d6000803e3d6000fd5b505050505b506002546000908190819073ffffffffffffffffffffffffffffffffffffffff1663c4276bfc8a61112060808c0160608d01612e37565b8a61112e60808e018e61332f565b6040518663ffffffff1660e01b815260040161114e959493929190613394565b600060405180830381865afa15801561116b573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526111b1919081019061345c565b919450925090506111c86080890160608a01612e37565b73ffffffffffffffffffffffffffffffffffffffff167f075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f8460405161120f91815260200190565b60405180910390a2604080516101a081019091526000610100820181815267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166101208501528c811661014085015287549293928392916101608401918a91879161128491166132e7565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff16815260200186611384576040517fea458c0c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8f16600482015273ffffffffffffffffffffffffffffffffffffffff8c811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063ea458c0c906044016020604051808303816000875af115801561135b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137f91906134b3565b611387565b60005b67ffffffffffffffff1681525081526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018a80602001906113c5919061332f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020016114098b8061332f565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161145060808c018c61332f565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161149a60808c0160608d01612e37565b73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018a80604001906114cb91906134d0565b905067ffffffffffffffff8111156114e5576114e5612915565b60405190808252806020026020018201604052801561154157816020015b61152e6040518060800160405280606081526020016060815260200160608152602001600081525090565b8152602001906001900390816115035790505b509052905060005b61155660408b018b6134d0565b9050811015611605576115dc61156f60408c018c6134d0565b8381811061157f5761157f61326c565b9050604002018036038101906115959190613538565b8c6115a08d8061332f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508e9250611d5e915050565b8260e0015182815181106115f2576115f261326c565b6020908102919091010152600101611549565b5060025460e082015173ffffffffffffffffffffffffffffffffffffffff9091169063cc88924c908c9061163c60408e018e6134d0565b6040518563ffffffff1660e01b815260040161165b9493929190613634565b60006040518083038186803b15801561167357600080fd5b505afa158015611687573d6000803e3d6000fd5b505050506080808201839052604080517f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908c166060820152309181019190915261172490829060a00160405160208183030381529060405280519060200120612068565b81515260405167ffffffffffffffff8b16907f0f07cd31e53232da9125e517f09550fdde74bf43d6a0a76ebd41674dafe2ab299061176390849061366a565b60405180910390a251519450505050505b949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906117bc575060055473ffffffffffffffffffffffffffffffffffffffff163314155b156117f3576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611861858585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505060408051602080890282810182019093528882529093508892508791829185019084908082843760009201919091525061216892505050565b5050505050565b6118706118ad565b61059b81612324565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005473ffffffffffffffffffffffffffffffffffffffff16331461192e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106ac565b565b60005b8151811015611aa85760008282815181106119505761195061326c565b60200260200101519050600083838151811061196e5761196e61326c565b60200260200101516000015190508067ffffffffffffffff166000036119cc576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016106ac565b60208281015167ffffffffffffffff83811660008181526006855260409081902080547fffffff0000000000000000000000000000000000000000ffffffffffffffffff8116690100000000000000000073ffffffffffffffffffffffffffffffffffffffff978816810291821793849055845192871691909616178152938104909416948301949094526801000000000000000090920460ff16151592810192909252907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef59060600160405180910390a25050600101611933565b5050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610982908490612419565b60606000611b4683612525565b9392505050565b805173ffffffffffffffffffffffffffffffffffffffff161580611b895750604081015173ffffffffffffffffffffffffffffffffffffffff16155b15611bc0576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600280547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff93841617909155602080840151600380548416918516919091179055604080850151600480548516918616919091179055606080860151600580549095169086161790935580516080810182527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681527f00000000000000000000000000000000000000000000000000000000000000008516928101929092527f00000000000000000000000000000000000000000000000000000000000000008416828201527f00000000000000000000000000000000000000000000000000000000000000009093169181019190915290517f23a1adf8ad7fad6091a4803227af2cee848c01a7c812404cade7c25636925e3291611d249184906137b8565b60405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561074e565b611d896040518060800160405280606081526020016060815260200160608152602001600081525090565b8460200151600003611dc7576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611dd7858760000151610987565b905073ffffffffffffffffffffffffffffffffffffffff81161580611ea757506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf00000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015611e81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea59190612f53565b155b15611ef95785516040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016106ac565b60008173ffffffffffffffffffffffffffffffffffffffff16639a4575b96040518060a001604052808881526020018967ffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018a6020015181526020018a6000015173ffffffffffffffffffffffffffffffffffffffff168152506040518263ffffffff1660e01b8152600401611f989190613857565b6000604051808303816000875af1158015611fb7573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611ffd91908101906138cd565b604080516080810190915273ffffffffffffffffffffffffffffffffffffffff841660a08201529091508060c0810160405160208183030381529060405281526020018260000151815260200182602001518152602001886020015181525092505050949350505050565b60008060001b82846020015185606001518660000151606001518760000151608001518860a001518960c001516040516020016120aa9695949392919061395e565b604051602081830303815290604052805190602001208560400151805190602001208660e001516040516020016120e191906139bf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201206080808c0151805190840120928501989098529183019590955260608201939093529384015260a083015260c082015260e00160405160208183030381529060405280519060200120905092915050565b67ffffffffffffffff83166000908152600660205260409020805468010000000000000000900460ff166121c8576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b835181101561220c576122038482815181106121e9576121e961326c565b60200260200101518360010161258190919063ffffffff16565b506001016121cb565b50825115612257578367ffffffffffffffff167f0d5e755ea090d8b16e0b3fed043532ed762c7e31a1f3884ac561cab59c7dbf1a8460405161224e9190612d48565b60405180910390a25b60005b82518110156122d35760008382815181106122775761227761326c565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036122bb57506122cb565b6122c860018401826125a3565b50505b60010161225a565b5081511561231e578367ffffffffffffffff167fe56852833d8b7c1f0dd03b91d1cff239d83ba81fbc1f10422d4cc4ae89199226836040516123159190612d48565b60405180910390a25b50505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036123a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106ac565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061247b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166125c59092919063ffffffff16565b80519091501561098257808060200190518101906124999190612f53565b610982576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016106ac565b60608160000180548060200260200160405190810160405280929190818152602001828054801561257557602002820191906000526020600020905b815481526020019060010190808311612561575b50505050509050919050565b600061074e8373ffffffffffffffffffffffffffffffffffffffff84166125d4565b600061074e8373ffffffffffffffffffffffffffffffffffffffff84166126ce565b6060611774848460008561271d565b600081815260018301602052604081205480156126bd5760006125f86001836139d2565b855490915060009061260c906001906139d2565b905080821461267157600086600001828154811061262c5761262c61326c565b906000526020600020015490508087600001848154811061264f5761264f61326c565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612682576126826139e5565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610751565b6000915050610751565b5092915050565b600081815260018301602052604081205461271557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610751565b506000610751565b6060824710156127af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016106ac565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516127d89190613a14565b60006040518083038185875af1925050503d8060008114612815576040519150601f19603f3d011682016040523d82523d6000602084013e61281a565b606091505b509150915061282b87838387612836565b979650505050505050565b606083156128cc5782516000036128c55773ffffffffffffffffffffffffffffffffffffffff85163b6128c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016106ac565b5081611774565b61177483838151156128e15781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ac9190612ba0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561296757612967612915565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156129b4576129b4612915565b604052919050565b600067ffffffffffffffff8211156129d6576129d6612915565b5060051b60200190565b67ffffffffffffffff8116811461059b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461059b57600080fd5b60006020808385031215612a2b57600080fd5b823567ffffffffffffffff811115612a4257600080fd5b8301601f81018513612a5357600080fd5b8035612a66612a61826129bc565b61296d565b81815260069190911b82018301908381019087831115612a8557600080fd5b928401925b8284101561282b5760408489031215612aa35760008081fd5b612aab612944565b8435612ab6816129e0565b815284860135612ac5816129f6565b8187015282526040939093019290840190612a8a565b60808101610751828467ffffffffffffffff8151168252602081015173ffffffffffffffffffffffffffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b60005b83811015612b4d578181015183820152602001612b35565b50506000910152565b60008151808452612b6e816020860160208601612b32565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061074e6020830184612b56565b600060a08284031215612bc557600080fd5b50919050565b60008060408385031215612bde57600080fd5b8235612be9816129e0565b9150602083013567ffffffffffffffff811115612c0557600080fd5b612c1185828601612bb3565b9150509250929050565b60008060408385031215612c2e57600080fd5b8235612c39816129e0565b91506020830135612c49816129f6565b809150509250929050565b60008083601f840112612c6657600080fd5b50813567ffffffffffffffff811115612c7e57600080fd5b6020830191508360208260051b8501011115612c9957600080fd5b9250929050565b60008060208385031215612cb357600080fd5b823567ffffffffffffffff811115612cca57600080fd5b612cd685828601612c54565b90969095509350505050565b608081016107518284805173ffffffffffffffffffffffffffffffffffffffff908116835260208083015182169084015260408083015182169084015260609182015116910152565b600060208284031215612d3d57600080fd5b8135611b46816129e0565b6020808252825182820181905260009190848201906040850190845b81811015612d9657835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612d64565b50909695505050505050565b8035612dad816129f6565b919050565b600060808284031215612dc457600080fd5b6040516080810181811067ffffffffffffffff82111715612de757612de7612915565b6040528235612df5816129f6565b81526020830135612e05816129f6565b60208201526040830135612e18816129f6565b60408201526060830135612e2b816129f6565b60608201529392505050565b600060208284031215612e4957600080fd5b8135611b46816129f6565b60008060008060808587031215612e6a57600080fd5b8435612e75816129e0565b9350602085013567ffffffffffffffff811115612e9157600080fd5b612e9d87828801612bb3565b935050604085013591506060850135612eb5816129f6565b939692955090935050565b600080600080600060608688031215612ed857600080fd5b8535612ee3816129e0565b9450602086013567ffffffffffffffff80821115612f0057600080fd5b612f0c89838a01612c54565b90965094506040880135915080821115612f2557600080fd5b50612f3288828901612c54565b969995985093965092949392505050565b80518015158114612dad57600080fd5b600060208284031215612f6557600080fd5b61074e82612f43565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612fa357600080fd5b830160208101925035905067ffffffffffffffff811115612fc357600080fd5b803603821315612c9957600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b8581101561307057813561303e816129f6565b73ffffffffffffffffffffffffffffffffffffffff16875281830135838801526040968701969091019060010161302b565b509495945050505050565b600067ffffffffffffffff80851683526040602084015261309c8485612f6e565b60a060408601526130b160e086018284612fd2565b9150506130c16020860186612f6e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808785030160608801526130f7848385612fd2565b9350604088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301831261313057600080fd5b6020928801928301923591508482111561314957600080fd5b8160061b360383131561315b57600080fd5b8087850301608088015261317084838561301b565b945061317e60608901612da2565b73ffffffffffffffffffffffffffffffffffffffff811660a089015293506131a96080890189612f6e565b94509250808786030160c0880152505061282b838383612fd2565b6000602082840312156131d657600080fd5b5051919050565b600060208083850312156131f057600080fd5b825167ffffffffffffffff81111561320757600080fd5b8301601f8101851361321857600080fd5b8051613226612a61826129bc565b81815260059190911b8201830190838101908783111561324557600080fd5b928401925b8284101561282b57835161325d816129f6565b8252928401929084019061324a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156132ad57600080fd5b8151611b46816129f6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818103613304576133046132b8565b6001019392505050565b67ffffffffffffffff8181168382160190808211156126c7576126c76132b8565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261336457600080fd5b83018035915067ffffffffffffffff82111561337f57600080fd5b602001915036819003821315612c9957600080fd5b67ffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015260806060820152600061282b608083018486612fd2565b600082601f8301126133eb57600080fd5b815167ffffffffffffffff81111561340557613405612915565b61343660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161296d565b81815284602083860101111561344b57600080fd5b611774826020830160208701612b32565b60008060006060848603121561347157600080fd5b8351925061348160208501612f43565b9150604084015167ffffffffffffffff81111561349d57600080fd5b6134a9868287016133da565b9150509250925092565b6000602082840312156134c557600080fd5b8151611b46816129e0565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261350557600080fd5b83018035915067ffffffffffffffff82111561352057600080fd5b6020019150600681901b3603821315612c9957600080fd5b60006040828403121561354a57600080fd5b613552612944565b823561355d816129f6565b81526020928301359281019290925250919050565b600082825180855260208086019550808260051b84010181860160005b84811015613627577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030189528151608081518186526135d382870182612b56565b91505085820151858203878701526135eb8282612b56565b915050604080830151868303828801526136058382612b56565b606094850151979094019690965250509884019892509083019060010161358f565b5090979650505050505050565b67ffffffffffffffff851681526060602082015260006136576060830186613572565b828103604084015261282b81858761301b565b602081526136bb60208201835180518252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015280608084015116608085015250505050565b600060208301516136e460c084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060408301516101808060e08501526137016101a0850183612b56565b915060608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030161010087015261373e8483612b56565b935060808701519150808685030161012087015261375c8483612b56565b935060a0870151915061378861014087018373ffffffffffffffffffffffffffffffffffffffff169052565b60c087015161016087015260e08701519150808685030183870152506137ae8382613572565b9695505050505050565b6101008101613810828567ffffffffffffffff8151168252602081015173ffffffffffffffffffffffffffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b825173ffffffffffffffffffffffffffffffffffffffff90811660808401526020840151811660a08401526040840151811660c084015260608401511660e0830152611b46565b602081526000825160a0602084015261387360c0840182612b56565b905067ffffffffffffffff6020850151166040840152604084015173ffffffffffffffffffffffffffffffffffffffff8082166060860152606086015160808601528060808701511660a086015250508091505092915050565b6000602082840312156138df57600080fd5b815167ffffffffffffffff808211156138f757600080fd5b908301906040828603121561390b57600080fd5b613913612944565b82518281111561392257600080fd5b61392e878286016133da565b82525060208301518281111561394357600080fd5b61394f878286016133da565b60208301525095945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808916835260c0602084015261398e60c0840189612b56565b67ffffffffffffffff97881660408501529590961660608301525091909316608082015260a0019190915292915050565b60208152600061074e6020830184613572565b81810381811115610751576107516132b8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251613a26818460208701612b32565b919091019291505056fea164736f6c6343000818000a",
}

var EVM2EVMMultiOnRampABI = EVM2EVMMultiOnRampMetaData.ABI

var EVM2EVMMultiOnRampBin = EVM2EVMMultiOnRampMetaData.Bin

func DeployEVM2EVMMultiOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOnRampStaticConfig, dynamicConfig EVM2EVMMultiOnRampDynamicConfig, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (common.Address, *types.Transaction, *EVM2EVMMultiOnRamp, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMMultiOnRamp{address: address, abi: *parsed, EVM2EVMMultiOnRampCaller: EVM2EVMMultiOnRampCaller{contract: contract}, EVM2EVMMultiOnRampTransactor: EVM2EVMMultiOnRampTransactor{contract: contract}, EVM2EVMMultiOnRampFilterer: EVM2EVMMultiOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMMultiOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMMultiOnRampCaller
	EVM2EVMMultiOnRampTransactor
	EVM2EVMMultiOnRampFilterer
}

type EVM2EVMMultiOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOnRampSession struct {
	Contract     *EVM2EVMMultiOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMMultiOnRampCallerSession struct {
	Contract *EVM2EVMMultiOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMMultiOnRampTransactorSession struct {
	Contract     *EVM2EVMMultiOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMMultiOnRampRaw struct {
	Contract *EVM2EVMMultiOnRamp
}

type EVM2EVMMultiOnRampCallerRaw struct {
	Contract *EVM2EVMMultiOnRampCaller
}

type EVM2EVMMultiOnRampTransactorRaw struct {
	Contract *EVM2EVMMultiOnRampTransactor
}

func NewEVM2EVMMultiOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMMultiOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMMultiOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMMultiOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRamp{address: address, abi: abi, EVM2EVMMultiOnRampCaller: EVM2EVMMultiOnRampCaller{contract: contract}, EVM2EVMMultiOnRampTransactor: EVM2EVMMultiOnRampTransactor{contract: contract}, EVM2EVMMultiOnRampFilterer: EVM2EVMMultiOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMMultiOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMMultiOnRampCaller, error) {
	contract, err := bindEVM2EVMMultiOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMMultiOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMMultiOnRampTransactor, error) {
	contract, err := bindEVM2EVMMultiOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMMultiOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMMultiOnRampFilterer, error) {
	contract, err := bindEVM2EVMMultiOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMMultiOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMMultiOnRamp.Contract.EVM2EVMMultiOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.EVM2EVMMultiOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.EVM2EVMMultiOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMMultiOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetAllowList(opts *bind.CallOpts, destinationChainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getAllowList", destinationChainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetAllowList(destinationChainSelector uint64) ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetAllowList(&_EVM2EVMMultiOnRamp.CallOpts, destinationChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetAllowList(destinationChainSelector uint64) ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetAllowList(&_EVM2EVMMultiOnRamp.CallOpts, destinationChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetAllowListEnabled(opts *bind.CallOpts, destinationChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getAllowListEnabled", destinationChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetAllowListEnabled(destinationChainSelector uint64) (bool, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetAllowListEnabled(&_EVM2EVMMultiOnRamp.CallOpts, destinationChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetAllowListEnabled(destinationChainSelector uint64) (bool, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetAllowListEnabled(&_EVM2EVMMultiOnRamp.CallOpts, destinationChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMMultiOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampDynamicConfig)).(*EVM2EVMMultiOnRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetDynamicConfig() (EVM2EVMMultiOnRampDynamicConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetDynamicConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetDynamicConfig() (EVM2EVMMultiOnRampDynamicConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetDynamicConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFee(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, message)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFee(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, message)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMMultiOnRamp.CallOpts, arg0, sourceToken)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMMultiOnRamp.CallOpts, arg0, sourceToken)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetRouter(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getRouter", destChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetRouter(destChainSelector uint64) (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetRouter(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetRouter(destChainSelector uint64) (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetRouter(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMMultiOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampStaticConfig)).(*EVM2EVMMultiOnRampStaticConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetStaticConfig() (EVM2EVMMultiOnRampStaticConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetStaticConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetStaticConfig() (EVM2EVMMultiOnRampStaticConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetStaticConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSupportedTokens(&_EVM2EVMMultiOnRamp.CallOpts, arg0)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSupportedTokens(&_EVM2EVMMultiOnRamp.CallOpts, arg0)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.Owner(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.Owner(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMMultiOnRamp.Contract.TypeAndVersion(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMMultiOnRamp.Contract.TypeAndVersion(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.AcceptOwnership(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.AcceptOwnership(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyAllowListUpdates", destinationChainSelector, removes, adds)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyAllowListUpdates(destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyAllowListUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destinationChainSelector, removes, adds)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyAllowListUpdates(destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyAllowListUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destinationChainSelector, removes, adds)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) DisableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "disableAllowList", destinationChainSelectors)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) DisableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.DisableAllowList(&_EVM2EVMMultiOnRamp.TransactOpts, destinationChainSelectors)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) DisableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.DisableAllowList(&_EVM2EVMMultiOnRamp.TransactOpts, destinationChainSelectors)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) EnableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "enableAllowList", destinationChainSelectors)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) EnableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.EnableAllowList(&_EVM2EVMMultiOnRamp.TransactOpts, destinationChainSelectors)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) EnableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.EnableAllowList(&_EVM2EVMMultiOnRamp.TransactOpts, destinationChainSelectors)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ForwardFromRouter(&_EVM2EVMMultiOnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ForwardFromRouter(&_EVM2EVMMultiOnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetAllowListAdmin(opts *bind.TransactOpts, allowListAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setAllowListAdmin", allowListAdmin)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetAllowListAdmin(allowListAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetAllowListAdmin(&_EVM2EVMMultiOnRamp.TransactOpts, allowListAdmin)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetAllowListAdmin(allowListAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetAllowListAdmin(&_EVM2EVMMultiOnRamp.TransactOpts, allowListAdmin)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetDynamicConfig(dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetDynamicConfig(&_EVM2EVMMultiOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetDynamicConfig(dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetDynamicConfig(&_EVM2EVMMultiOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.TransferOwnership(&_EVM2EVMMultiOnRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.TransferOwnership(&_EVM2EVMMultiOnRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "withdrawFeeTokens")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) WithdrawFeeTokens() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.WithdrawFeeTokens(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) WithdrawFeeTokens() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.WithdrawFeeTokens(&_EVM2EVMMultiOnRamp.TransactOpts)
}

type EVM2EVMMultiOnRampAllowListAddedIterator struct {
	Event *EVM2EVMMultiOnRampAllowListAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampAllowListAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampAllowListAdded)
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
		it.Event = new(EVM2EVMMultiOnRampAllowListAdded)
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

func (it *EVM2EVMMultiOnRampAllowListAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampAllowListAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampAllowListAdded struct {
	DestChainSelector uint64
	AllowList         []common.Address
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterAllowListAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampAllowListAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "AllowListAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampAllowListAddedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "AllowListAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchAllowListAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "AllowListAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampAllowListAdded)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListAdded", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseAllowListAdded(log types.Log) (*EVM2EVMMultiOnRampAllowListAdded, error) {
	event := new(EVM2EVMMultiOnRampAllowListAdded)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampAllowListAdminSetIterator struct {
	Event *EVM2EVMMultiOnRampAllowListAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampAllowListAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampAllowListAdminSet)
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
		it.Event = new(EVM2EVMMultiOnRampAllowListAdminSet)
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

func (it *EVM2EVMMultiOnRampAllowListAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampAllowListAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampAllowListAdminSet struct {
	AllowListAdmin common.Address
	Raw            types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterAllowListAdminSet(opts *bind.FilterOpts, allowListAdmin []common.Address) (*EVM2EVMMultiOnRampAllowListAdminSetIterator, error) {

	var allowListAdminRule []interface{}
	for _, allowListAdminItem := range allowListAdmin {
		allowListAdminRule = append(allowListAdminRule, allowListAdminItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "AllowListAdminSet", allowListAdminRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampAllowListAdminSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "AllowListAdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListAdminSet, allowListAdmin []common.Address) (event.Subscription, error) {

	var allowListAdminRule []interface{}
	for _, allowListAdminItem := range allowListAdmin {
		allowListAdminRule = append(allowListAdminRule, allowListAdminItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "AllowListAdminSet", allowListAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampAllowListAdminSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseAllowListAdminSet(log types.Log) (*EVM2EVMMultiOnRampAllowListAdminSet, error) {
	event := new(EVM2EVMMultiOnRampAllowListAdminSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampAllowListDisabledIterator struct {
	Event *EVM2EVMMultiOnRampAllowListDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampAllowListDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampAllowListDisabled)
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
		it.Event = new(EVM2EVMMultiOnRampAllowListDisabled)
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

func (it *EVM2EVMMultiOnRampAllowListDisabledIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampAllowListDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampAllowListDisabled struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterAllowListDisabled(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAllowListDisabledIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "AllowListDisabled")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampAllowListDisabledIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "AllowListDisabled", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchAllowListDisabled(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListDisabled) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "AllowListDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampAllowListDisabled)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListDisabled", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseAllowListDisabled(log types.Log) (*EVM2EVMMultiOnRampAllowListDisabled, error) {
	event := new(EVM2EVMMultiOnRampAllowListDisabled)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampAllowListEnabledIterator struct {
	Event *EVM2EVMMultiOnRampAllowListEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampAllowListEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampAllowListEnabled)
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
		it.Event = new(EVM2EVMMultiOnRampAllowListEnabled)
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

func (it *EVM2EVMMultiOnRampAllowListEnabledIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampAllowListEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampAllowListEnabled struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterAllowListEnabled(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAllowListEnabledIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "AllowListEnabled")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampAllowListEnabledIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "AllowListEnabled", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchAllowListEnabled(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListEnabled) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "AllowListEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampAllowListEnabled)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListEnabled", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseAllowListEnabled(log types.Log) (*EVM2EVMMultiOnRampAllowListEnabled, error) {
	event := new(EVM2EVMMultiOnRampAllowListEnabled)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampAllowListRemovedIterator struct {
	Event *EVM2EVMMultiOnRampAllowListRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampAllowListRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampAllowListRemoved)
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
		it.Event = new(EVM2EVMMultiOnRampAllowListRemoved)
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

func (it *EVM2EVMMultiOnRampAllowListRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampAllowListRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampAllowListRemoved struct {
	DestChainSelector uint64
	AllowList         []common.Address
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterAllowListRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampAllowListRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "AllowListRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampAllowListRemovedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "AllowListRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchAllowListRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "AllowListRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampAllowListRemoved)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListRemoved", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseAllowListRemoved(log types.Log) (*EVM2EVMMultiOnRampAllowListRemoved, error) {
	event := new(EVM2EVMMultiOnRampAllowListRemoved)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AllowListRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMMultiOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMMultiOnRampCCIPSendRequested)
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

func (it *EVM2EVMMultiOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampCCIPSendRequested struct {
	DestChainSelector uint64
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "CCIPSendRequested", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampCCIPSendRequestedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "CCIPSendRequested", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampCCIPSendRequested)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMMultiOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMMultiOnRampCCIPSendRequested)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampConfigSetIterator struct {
	Event *EVM2EVMMultiOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampConfigSet)
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
		it.Event = new(EVM2EVMMultiOnRampConfigSet)
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

func (it *EVM2EVMMultiOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampConfigSet struct {
	StaticConfig  EVM2EVMMultiOnRampStaticConfig
	DynamicConfig EVM2EVMMultiOnRampDynamicConfig
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampConfigSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampConfigSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMMultiOnRampConfigSet, error) {
	event := new(EVM2EVMMultiOnRampConfigSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampDestChainConfigSetIterator struct {
	Event *EVM2EVMMultiOnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampDestChainConfigSet)
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
		it.Event = new(EVM2EVMMultiOnRampDestChainConfigSet)
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

func (it *EVM2EVMMultiOnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	AllowListEnabled  bool
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampDestChainConfigSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampDestChainConfigSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseDestChainConfigSet(log types.Log) (*EVM2EVMMultiOnRampDestChainConfigSet, error) {
	event := new(EVM2EVMMultiOnRampDestChainConfigSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampFeePaidIterator struct {
	Event *EVM2EVMMultiOnRampFeePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampFeePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampFeePaid)
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
		it.Event = new(EVM2EVMMultiOnRampFeePaid)
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

func (it *EVM2EVMMultiOnRampFeePaidIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampFeePaid struct {
	FeeToken      common.Address
	FeeValueJuels *big.Int
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterFeePaid(opts *bind.FilterOpts, feeToken []common.Address) (*EVM2EVMMultiOnRampFeePaidIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "FeePaid", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFeePaidIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeePaid, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "FeePaid", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampFeePaid)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeePaid", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseFeePaid(log types.Log) (*EVM2EVMMultiOnRampFeePaid, error) {
	event := new(EVM2EVMMultiOnRampFeePaid)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampFeeTokenWithdrawnIterator struct {
	Event *EVM2EVMMultiOnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
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
		it.Event = new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
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

func (it *EVM2EVMMultiOnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*EVM2EVMMultiOnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFeeTokenWithdrawnIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*EVM2EVMMultiOnRampFeeTokenWithdrawn, error) {
	event := new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMMultiOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMMultiOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMMultiOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampOwnershipTransferRequested)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMMultiOnRampOwnershipTransferRequested)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMMultiOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMMultiOnRampOwnershipTransferred)
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

func (it *EVM2EVMMultiOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampOwnershipTransferredIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampOwnershipTransferred)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMMultiOnRampOwnershipTransferred)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOnRamp.abi.Events["AllowListAdded"].ID:
		return _EVM2EVMMultiOnRamp.ParseAllowListAdded(log)
	case _EVM2EVMMultiOnRamp.abi.Events["AllowListAdminSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseAllowListAdminSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["AllowListDisabled"].ID:
		return _EVM2EVMMultiOnRamp.ParseAllowListDisabled(log)
	case _EVM2EVMMultiOnRamp.abi.Events["AllowListEnabled"].ID:
		return _EVM2EVMMultiOnRamp.ParseAllowListEnabled(log)
	case _EVM2EVMMultiOnRamp.abi.Events["AllowListRemoved"].ID:
		return _EVM2EVMMultiOnRamp.ParseAllowListRemoved(log)
	case _EVM2EVMMultiOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["FeePaid"].ID:
		return _EVM2EVMMultiOnRamp.ParseFeePaid(log)
	case _EVM2EVMMultiOnRamp.abi.Events["FeeTokenWithdrawn"].ID:
		return _EVM2EVMMultiOnRamp.ParseFeeTokenWithdrawn(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMMultiOnRampAllowListAdded) Topic() common.Hash {
	return common.HexToHash("0xe56852833d8b7c1f0dd03b91d1cff239d83ba81fbc1f10422d4cc4ae89199226")
}

func (EVM2EVMMultiOnRampAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (EVM2EVMMultiOnRampAllowListDisabled) Topic() common.Hash {
	return common.HexToHash("0x50013bb4071fed8f5c37e8a9ee6ecc157fe969ed7e303069ecbbba1e89e12064")
}

func (EVM2EVMMultiOnRampAllowListEnabled) Topic() common.Hash {
	return common.HexToHash("0x4d8b8e1b49157081f6c6ec663e74f97d33537bf46858985e5eaefd364410be60")
}

func (EVM2EVMMultiOnRampAllowListRemoved) Topic() common.Hash {
	return common.HexToHash("0x0d5e755ea090d8b16e0b3fed043532ed762c7e31a1f3884ac561cab59c7dbf1a")
}

func (EVM2EVMMultiOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0x0f07cd31e53232da9125e517f09550fdde74bf43d6a0a76ebd41674dafe2ab29")
}

func (EVM2EVMMultiOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x23a1adf8ad7fad6091a4803227af2cee848c01a7c812404cade7c25636925e32")
}

func (EVM2EVMMultiOnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5")
}

func (EVM2EVMMultiOnRampFeePaid) Topic() common.Hash {
	return common.HexToHash("0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f")
}

func (EVM2EVMMultiOnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (EVM2EVMMultiOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMMultiOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) Address() common.Address {
	return _EVM2EVMMultiOnRamp.address
}

type EVM2EVMMultiOnRampInterface interface {
	GetAllowList(opts *bind.CallOpts, destinationChainSelector uint64) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts, destinationChainSelector uint64) (bool, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetRouter(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error)

	DisableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error)

	EnableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetAllowListAdmin(opts *bind.TransactOpts, allowListAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAllowListAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampAllowListAddedIterator, error)

	WatchAllowListAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListAdded(log types.Log) (*EVM2EVMMultiOnRampAllowListAdded, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowListAdmin []common.Address) (*EVM2EVMMultiOnRampAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListAdminSet, allowListAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*EVM2EVMMultiOnRampAllowListAdminSet, error)

	FilterAllowListDisabled(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAllowListDisabledIterator, error)

	WatchAllowListDisabled(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListDisabled) (event.Subscription, error)

	ParseAllowListDisabled(log types.Log) (*EVM2EVMMultiOnRampAllowListDisabled, error)

	FilterAllowListEnabled(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAllowListEnabledIterator, error)

	WatchAllowListEnabled(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListEnabled) (event.Subscription, error)

	ParseAllowListEnabled(log types.Log) (*EVM2EVMMultiOnRampAllowListEnabled, error)

	FilterAllowListRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampAllowListRemovedIterator, error)

	WatchAllowListRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAllowListRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListRemoved(log types.Log) (*EVM2EVMMultiOnRampAllowListRemoved, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMMultiOnRampCCIPSendRequested, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMMultiOnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*EVM2EVMMultiOnRampDestChainConfigSet, error)

	FilterFeePaid(opts *bind.FilterOpts, feeToken []common.Address) (*EVM2EVMMultiOnRampFeePaidIterator, error)

	WatchFeePaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeePaid, feeToken []common.Address) (event.Subscription, error)

	ParseFeePaid(log types.Log) (*EVM2EVMMultiOnRampFeePaid, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*EVM2EVMMultiOnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*EVM2EVMMultiOnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
