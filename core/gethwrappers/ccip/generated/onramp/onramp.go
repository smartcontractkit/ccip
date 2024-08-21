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

type OnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
}

type OnRampDynamicConfig struct {
	PriceRegistry    common.Address
	MessageValidator common.Address
	FeeAggregator    common.Address
	AllowListAdmin   common.Address
}

type OnRampStaticConfig struct {
	ChainSelector      uint64
	RmnProxy           common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowList\",\"type\":\"address[]\"}],\"name\":\"AllowListAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"name\":\"AllowListAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"AllowListDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"AllowListEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowList\",\"type\":\"address[]\"}],\"name\":\"AllowListRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowListEnabled\",\"type\":\"bool\"}],\"name\":\"DestChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeTokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"destinationChainSelectors\",\"type\":\"uint64[]\"}],\"name\":\"disableAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"destinationChainSelectors\",\"type\":\"uint64[]\"}],\"name\":\"enableAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"}],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPoolV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"}],\"name\":\"isAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"name\":\"setAllowListAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowListAdmin\",\"type\":\"address\"}],\"internalType\":\"structOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200418d3803806200418d83398101604081905262000035916200066a565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000186565b505083516001600160401b031615905080620000e6575060208301516001600160a01b0316155b80620000fd575060408301516001600160a01b0316155b8062000114575060608301516001600160a01b0316155b1562000133576040516306b7c75960e31b815260040160405180910390fd5b82516001600160401b031660805260208301516001600160a01b0390811660a0526040840151811660c05260608401511660e052620001728262000231565b6200017d8162000390565b505050620007a4565b336001600160a01b03821603620001e05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0316158062000254575060408101516001600160a01b0316155b1562000273576040516306b7c75960e31b815260040160405180910390fd5b8051600280546001600160a01b03199081166001600160a01b0393841617909155602080840180516003805485169186169190911790556040808601805160048054871691881691909117905560608088018051600580549098169089161790965582516080808201855280516001600160401b031680835260a080518b16848a0190815260c080518d16868a0190815260e080518f169789019788528a5195865292518e169b85019b909b5299518c169783019790975292518a169381019390935289518916908301529351871693810193909352518516928201929092529151909216918101919091527f23a1adf8ad7fad6091a4803227af2cee848c01a7c812404cade7c25636925e32906101000160405180910390a150565b60005b8151811015620004d4576000828281518110620003b457620003b46200078e565b602002602001015190506000838381518110620003d557620003d56200078e565b6020026020010151600001519050806001600160401b03166000036200041a5760405163c35aa79d60e01b81526001600160401b038216600482015260240162000083565b6020828101516001600160401b038381166000818152600685526040908190208054600160481b600160e81b0319811669010000000000000000006001600160a01b03978816810291821793849055845192871691909616178152938104909416948301949094526801000000000000000090920460ff16151592810192909252907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef59060600160405180910390a2505060010162000393565b5050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620005135762000513620004d8565b60405290565b604080519081016001600160401b0381118282101715620005135762000513620004d8565b604051601f8201601f191681016001600160401b0381118282101715620005695762000569620004d8565b604052919050565b80516001600160401b03811681146200058957600080fd5b919050565b6001600160a01b0381168114620005a457600080fd5b50565b600082601f830112620005b957600080fd5b815160206001600160401b03821115620005d757620005d7620004d8565b620005e7818360051b016200053e565b82815260069290921b840181019181810190868411156200060757600080fd5b8286015b848110156200065f5760408189031215620006265760008081fd5b6200063062000519565b6200063b8262000571565b8152848201516200064c816200058e565b818601528352918301916040016200060b565b509695505050505050565b60008060008385036101208112156200068257600080fd5b60808112156200069157600080fd5b6200069b620004ee565b620006a68662000571565b81526020860151620006b8816200058e565b60208201526040860151620006cd816200058e565b60408201526060860151620006e2816200058e565b606082015293506080607f1982011215620006fc57600080fd5b5062000707620004ee565b608085015162000717816200058e565b815260a085015162000729816200058e565b602082015260c08501516200073e816200058e565b604082015260e085015162000753816200058e565b60608201526101008501519092506001600160401b038111156200077657600080fd5b6200078486828701620005a7565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b60805160a05160c05160e0516139706200081d6000396000818161026d015281816109cf0152611df6015260008181610231015281816113010152611dcf0152600081816101f5015281816106020152611da50152600081816101c501528181611227015281816116b70152611d7801526139706000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80638da5cb5b116100d8578063b627f8ec1161008c578063f0f512f411610066578063f0f512f414610551578063f2fde38b14610564578063fbca3b741461057757600080fd5b8063b627f8ec146104de578063d77d5ed0146104f1578063df0aa9e91461053e57600080fd5b8063991e7e6e116100bd578063991e7e6e146104625780639a6281f714610482578063a6f3ab6c146104cb57600080fd5b80638da5cb5b146104185780639041be3d1461043657600080fd5b806348a98aa41161012f5780637437ff9f116101145780637437ff9f1461037d57806379ba5097146103fd57806380c162f71461040557600080fd5b806348a98aa414610332578063567261b61461036a57600080fd5b8063181f5a7711610160578063181f5a77146102c057806320487ded146103095780633a0199401461032a57600080fd5b80630242cf601461017c57806306285c6914610191575b600080fd5b61018f61018a3660046128f0565b61058a565b005b6102aa60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b6040516102b791906129b3565b60405180910390f35b6102fc6040518060400160405280601081526020017f4f6e52616d7020312e362e302d6465760000000000000000000000000000000081525081565b6040516102b79190612a78565b61031c610317366004612aa3565b61059e565b6040519081526020016102b7565b61018f610757565b610345610340366004612af3565b610987565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102b7565b61018f610378366004612b78565b610a3c565b6103f0604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260025473ffffffffffffffffffffffffffffffffffffffff908116825260035481166020830152600454811692820192909252600554909116606082015290565b6040516102b79190612bba565b61018f610bb2565b61018f610413366004612b78565b610caf565b60005473ffffffffffffffffffffffffffffffffffffffff16610345565b610449610444366004612c03565b610e25565b60405167ffffffffffffffff90911681526020016102b7565b610475610470366004612c03565b610e4e565b6040516102b79190612c20565b6104bb610490366004612c03565b67ffffffffffffffff1660009081526006602052604090205468010000000000000000900460ff1690565b60405190151581526020016102b7565b61018f6104d9366004612c8a565b610e76565b61018f6104ec366004612d0f565b610e87565b6103456104ff366004612c03565b67ffffffffffffffff166000908152600660205260409020546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1690565b61031c61054c366004612d2c565b610efe565b61018f61055f366004612d98565b61176b565b61018f610572366004612d0f565b611992565b610475610585366004612c03565b6119a3565b6105926119d7565b61059b81611a5a565b50565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b16600482015260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690632cbc26bb90602401602060405180830381865afa158015610649573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066d9190612e2b565b156106b5576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6002546040517fd8694ccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d8694ccd9061070d9086908690600401612f53565b602060405180830381865afa15801561072a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074e919061309c565b90505b92915050565b600254604080517fcdc73d51000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cdc73d5191600480830192869291908290030181865afa1580156107c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261080c91908101906130b5565b60045490915073ffffffffffffffffffffffffffffffffffffffff1660005b825181101561098257600083828151811061084857610848613144565b60209081029190910101516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa1580156108c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e7919061309c565b905080156109785761091073ffffffffffffffffffffffffffffffffffffffff83168583611bd6565b8173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e8360405161096f91815260200190565b60405180910390a35b505060010161082b565b505050565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bbe4f6db90602401602060405180830381865afa158015610a18573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074e9190613173565b60005473ffffffffffffffffffffffffffffffffffffffff163314610aac5760055473ffffffffffffffffffffffffffffffffffffffff163314610aac576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b67ffffffffffffffff81168211156109825760006006600085858567ffffffffffffffff16818110610ae357610ae3613144565b9050602002016020810190610af89190612c03565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160086101000a81548160ff0219169083151502179055507f50013bb4071fed8f5c37e8a9ee6ecc157fe969ed7e303069ecbbba1e89e1206483838367ffffffffffffffff16818110610b7157610b71613144565b9050602002016020810190610b869190612c03565b60405167ffffffffffffffff909116815260200160405180910390a1610bab816131bf565b9050610aaf565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016106ac565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff163314610d1f5760055473ffffffffffffffffffffffffffffffffffffffff163314610d1f576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b67ffffffffffffffff81168211156109825760016006600085858567ffffffffffffffff16818110610d5657610d56613144565b9050602002016020810190610d6b9190612c03565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160086101000a81548160ff0219169083151502179055507f4d8b8e1b49157081f6c6ec663e74f97d33537bf46858985e5eaefd364410be6083838367ffffffffffffffff16818110610de457610de4613144565b9050602002016020810190610df99190612c03565b60405167ffffffffffffffff909116815260200160405180910390a1610e1e816131bf565b9050610d22565b67ffffffffffffffff8082166000908152600660205260408120549091610751911660016131e6565b67ffffffffffffffff8116600090815260066020526040902060609061075190600101611c63565b610e7e6119d7565b61059b81611c77565b610e8f6119d7565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e90600090a250565b67ffffffffffffffff84166000908152600660205260408120805468010000000000000000900460ff1615610f8857610f3a6001820184611e59565b610f88576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016106ac565b73ffffffffffffffffffffffffffffffffffffffff8316610fd5576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611032576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035473ffffffffffffffffffffffffffffffffffffffff1680156110d8576040517fe0a0e50600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063e0a0e506906110a5908a908a90600401612f53565b600060405180830381600087803b1580156110bf57600080fd5b505af11580156110d3573d6000803e3d6000fd5b505050505b506002546000908190819073ffffffffffffffffffffffffffffffffffffffff1663c4276bfc8a61110f60808c0160608d01612d0f565b8a61111d60808e018e613207565b6040518663ffffffff1660e01b815260040161113d95949392919061326c565b600060405180830381865afa15801561115a573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526111a09190810190613334565b919450925090506111b76080890160608a01612d0f565b73ffffffffffffffffffffffffffffffffffffffff167f075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f846040516111fe91815260200190565b60405180910390a2604080516101a081019091526000610100820181815267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166101208501528c811661014085015287549293928392916101608401918a91879161127391166131bf565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff16815260200186611373576040517fea458c0c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8f16600482015273ffffffffffffffffffffffffffffffffffffffff8c811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063ea458c0c906044016020604051808303816000875af115801561134a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061136e919061338b565b611376565b60005b67ffffffffffffffff1681525081526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018a80602001906113b49190613207565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020016113f88b80613207565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161143f60808c018c613207565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161148960808c0160608d01612d0f565b73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018a80604001906114ba91906133a8565b905067ffffffffffffffff8111156114d4576114d46127ed565b60405190808252806020026020018201604052801561153057816020015b61151d6040518060800160405280606081526020016060815260200160608152602001600081525090565b8152602001906001900390816114f25790505b509052905060005b61154560408b018b6133a8565b90508110156115f4576115cb61155e60408c018c6133a8565b8381811061156e5761156e613144565b9050604002018036038101906115849190613410565b8c61158f8d80613207565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508e9250611e88915050565b8260e0015182815181106115e1576115e1613144565b6020908102919091010152600101611538565b5060025460e082015173ffffffffffffffffffffffffffffffffffffffff9091169063cc88924c908c9061162b60408e018e6133a8565b6040518563ffffffff1660e01b815260040161164a949392919061350c565b60006040518083038186803b15801561166257600080fd5b505afa158015611676573d6000803e3d6000fd5b505050506080808201839052604080517f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908c166060820152309181019190915261171390829060a00160405160208183030381529060405280519060200120612192565b81515260405167ffffffffffffffff8b16907f0f07cd31e53232da9125e517f09550fdde74bf43d6a0a76ebd41674dafe2ab2990611752908490613542565b60405180910390a251519450505050505b949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146117db5760055473ffffffffffffffffffffffffffffffffffffffff1633146117db576040517f905d7d9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff85166000908152600660205260409020805468010000000000000000900460ff1661183b576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b848110156118835761187a86868381811061185b5761185b613144565b90506020020160208101906118709190612d0f565b6001840190612292565b5060010161183e565b5083156118cf578567ffffffffffffffff167f0d5e755ea090d8b16e0b3fed043532ed762c7e31a1f3884ac561cab59c7dbf1a86866040516118c6929190613690565b60405180910390a25b60005b8281101561193e5760008484838181106118ee576118ee613144565b90506020020160208101906119039190612d0f565b905073ffffffffffffffffffffffffffffffffffffffff81166119265750611936565b61193360018401826122b4565b50505b6001016118d2565b50811561198a578567ffffffffffffffff167fe56852833d8b7c1f0dd03b91d1cff239d83ba81fbc1f10422d4cc4ae891992268484604051611981929190613690565b60405180910390a25b505050505050565b61199a6119d7565b61059b816122d6565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005473ffffffffffffffffffffffffffffffffffffffff163314611a58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106ac565b565b60005b8151811015611bd2576000828281518110611a7a57611a7a613144565b602002602001015190506000838381518110611a9857611a98613144565b60200260200101516000015190508067ffffffffffffffff16600003611af6576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016106ac565b60208281015167ffffffffffffffff83811660008181526006855260409081902080547fffffff0000000000000000000000000000000000000000ffffffffffffffffff8116690100000000000000000073ffffffffffffffffffffffffffffffffffffffff978816810291821793849055845192871691909616178152938104909416948301949094526801000000000000000090920460ff16151592810192909252907fd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef59060600160405180910390a25050600101611a5d565b5050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526109829084906123cb565b60606000611c70836124d7565b9392505050565b805173ffffffffffffffffffffffffffffffffffffffff161580611cb35750604081015173ffffffffffffffffffffffffffffffffffffffff16155b15611cea576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600280547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff93841617909155602080840151600380548416918516919091179055604080850151600480548516918616919091179055606080860151600580549095169086161790935580516080810182527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681527f00000000000000000000000000000000000000000000000000000000000000008516928101929092527f00000000000000000000000000000000000000000000000000000000000000008416828201527f00000000000000000000000000000000000000000000000000000000000000009093169181019190915290517f23a1adf8ad7fad6091a4803227af2cee848c01a7c812404cade7c25636925e3291611e4e9184906136eb565b60405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561074e565b611eb36040518060800160405280606081526020016060815260200160608152602001600081525090565b8460200151600003611ef1576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611f01858760000151610987565b905073ffffffffffffffffffffffffffffffffffffffff81161580611fd157506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf00000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015611fab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fcf9190612e2b565b155b156120235785516040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016106ac565b60008173ffffffffffffffffffffffffffffffffffffffff16639a4575b96040518060a001604052808881526020018967ffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018a6020015181526020018a6000015173ffffffffffffffffffffffffffffffffffffffff168152506040518263ffffffff1660e01b81526004016120c2919061378a565b6000604051808303816000875af11580156120e1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526121279190810190613800565b604080516080810190915273ffffffffffffffffffffffffffffffffffffffff841660a08201529091508060c0810160405160208183030381529060405281526020018260000151815260200182602001518152602001886020015181525092505050949350505050565b60008060001b82846020015185606001518660000151606001518760000151608001518860a001518960c001516040516020016121d496959493929190613891565b604051602081830303815290604052805190602001208560400151805190602001208660e0015160405160200161220b91906138f2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201206080808c0151805190840120928501989098529183019590955260608201939093529384015260a083015260c082015260e00160405160208183030381529060405280519060200120905092915050565b600061074e8373ffffffffffffffffffffffffffffffffffffffff8416612533565b600061074e8373ffffffffffffffffffffffffffffffffffffffff841661262d565b3373ffffffffffffffffffffffffffffffffffffffff821603612355576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106ac565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061242d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661267c9092919063ffffffff16565b805190915015610982578080602001905181019061244b9190612e2b565b610982576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016106ac565b60608160000180548060200260200160405190810160405280929190818152602001828054801561252757602002820191906000526020600020905b815481526020019060010190808311612513575b50505050509050919050565b6000818152600183016020526040812054801561261c576000612557600183613905565b855490915060009061256b90600190613905565b90508082146125d057600086600001828154811061258b5761258b613144565b90600052602060002001549050808760000184815481106125ae576125ae613144565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806125e1576125e1613918565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610751565b6000915050610751565b5092915050565b600081815260018301602052604081205461267457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610751565b506000610751565b60606117638484600085856000808673ffffffffffffffffffffffffffffffffffffffff1685876040516126b09190613947565b60006040518083038185875af1925050503d80600081146126ed576040519150601f19603f3d011682016040523d82523d6000602084013e6126f2565b606091505b50915091506127038783838761270e565b979650505050505050565b606083156127a457825160000361279d5773ffffffffffffffffffffffffffffffffffffffff85163b61279d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016106ac565b5081611763565b61176383838151156127b95781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ac9190612a78565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561283f5761283f6127ed565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561288c5761288c6127ed565b604052919050565b600067ffffffffffffffff8211156128ae576128ae6127ed565b5060051b60200190565b67ffffffffffffffff8116811461059b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461059b57600080fd5b6000602080838503121561290357600080fd5b823567ffffffffffffffff81111561291a57600080fd5b8301601f8101851361292b57600080fd5b803561293e61293982612894565b612845565b81815260069190911b8201830190838101908783111561295d57600080fd5b928401925b82841015612703576040848903121561297b5760008081fd5b61298361281c565b843561298e816128b8565b81528486013561299d816128ce565b8187015282526040939093019290840190612962565b60808101610751828467ffffffffffffffff8151168252602081015173ffffffffffffffffffffffffffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b60005b83811015612a25578181015183820152602001612a0d565b50506000910152565b60008151808452612a46816020860160208601612a0a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061074e6020830184612a2e565b600060a08284031215612a9d57600080fd5b50919050565b60008060408385031215612ab657600080fd5b8235612ac1816128b8565b9150602083013567ffffffffffffffff811115612add57600080fd5b612ae985828601612a8b565b9150509250929050565b60008060408385031215612b0657600080fd5b8235612b11816128b8565b91506020830135612b21816128ce565b809150509250929050565b60008083601f840112612b3e57600080fd5b50813567ffffffffffffffff811115612b5657600080fd5b6020830191508360208260051b8501011115612b7157600080fd5b9250929050565b60008060208385031215612b8b57600080fd5b823567ffffffffffffffff811115612ba257600080fd5b612bae85828601612b2c565b90969095509350505050565b608081016107518284805173ffffffffffffffffffffffffffffffffffffffff908116835260208083015182169084015260408083015182169084015260609182015116910152565b600060208284031215612c1557600080fd5b8135611c70816128b8565b6020808252825182820181905260009190848201906040850190845b81811015612c6e57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612c3c565b50909695505050505050565b8035612c85816128ce565b919050565b600060808284031215612c9c57600080fd5b6040516080810181811067ffffffffffffffff82111715612cbf57612cbf6127ed565b6040528235612ccd816128ce565b81526020830135612cdd816128ce565b60208201526040830135612cf0816128ce565b60408201526060830135612d03816128ce565b60608201529392505050565b600060208284031215612d2157600080fd5b8135611c70816128ce565b60008060008060808587031215612d4257600080fd5b8435612d4d816128b8565b9350602085013567ffffffffffffffff811115612d6957600080fd5b612d7587828801612a8b565b935050604085013591506060850135612d8d816128ce565b939692955090935050565b600080600080600060608688031215612db057600080fd5b8535612dbb816128b8565b9450602086013567ffffffffffffffff80821115612dd857600080fd5b612de489838a01612b2c565b90965094506040880135915080821115612dfd57600080fd5b50612e0a88828901612b2c565b969995985093965092949392505050565b80518015158114612c8557600080fd5b600060208284031215612e3d57600080fd5b61074e82612e1b565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612e7b57600080fd5b830160208101925035905067ffffffffffffffff811115612e9b57600080fd5b803603821315612b7157600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b85811015612f48578135612f16816128ce565b73ffffffffffffffffffffffffffffffffffffffff168752818301358388015260409687019690910190600101612f03565b509495945050505050565b600067ffffffffffffffff808516835260406020840152612f748485612e46565b60a06040860152612f8960e086018284612eaa565b915050612f996020860186612e46565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878503016060880152612fcf848385612eaa565b9350604088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301831261300857600080fd5b6020928801928301923591508482111561302157600080fd5b8160061b360383131561303357600080fd5b80878503016080880152613048848385612ef3565b945061305660608901612c7a565b73ffffffffffffffffffffffffffffffffffffffff811660a089015293506130816080890189612e46565b94509250808786030160c08801525050612703838383612eaa565b6000602082840312156130ae57600080fd5b5051919050565b600060208083850312156130c857600080fd5b825167ffffffffffffffff8111156130df57600080fd5b8301601f810185136130f057600080fd5b80516130fe61293982612894565b81815260059190911b8201830190838101908783111561311d57600080fd5b928401925b82841015612703578351613135816128ce565b82529284019290840190613122565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561318557600080fd5b8151611c70816128ce565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff8083168181036131dc576131dc613190565b6001019392505050565b67ffffffffffffffff81811683821601908082111561262657612626613190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261323c57600080fd5b83018035915067ffffffffffffffff82111561325757600080fd5b602001915036819003821315612b7157600080fd5b67ffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff85166020820152836040820152608060608201526000612703608083018486612eaa565b600082601f8301126132c357600080fd5b815167ffffffffffffffff8111156132dd576132dd6127ed565b61330e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612845565b81815284602083860101111561332357600080fd5b611763826020830160208701612a0a565b60008060006060848603121561334957600080fd5b8351925061335960208501612e1b565b9150604084015167ffffffffffffffff81111561337557600080fd5b613381868287016132b2565b9150509250925092565b60006020828403121561339d57600080fd5b8151611c70816128b8565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126133dd57600080fd5b83018035915067ffffffffffffffff8211156133f857600080fd5b6020019150600681901b3603821315612b7157600080fd5b60006040828403121561342257600080fd5b61342a61281c565b8235613435816128ce565b81526020928301359281019290925250919050565b600082825180855260208086019550808260051b84010181860160005b848110156134ff577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030189528151608081518186526134ab82870182612a2e565b91505085820151858203878701526134c38282612a2e565b915050604080830151868303828801526134dd8382612a2e565b6060948501519790940196909652505098840198925090830190600101613467565b5090979650505050505050565b67ffffffffffffffff8516815260606020820152600061352f606083018661344a565b8281036040840152612703818587612ef3565b6020815261359360208201835180518252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015280608084015116608085015250505050565b600060208301516135bc60c084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060408301516101808060e08501526135d96101a0850183612a2e565b915060608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080868503016101008701526136168483612a2e565b93506080870151915080868503016101208701526136348483612a2e565b935060a0870151915061366061014087018373ffffffffffffffffffffffffffffffffffffffff169052565b60c087015161016087015260e0870151915080868503018387015250613686838261344a565b9695505050505050565b60208082528181018390526000908460408401835b868110156136e05782356136b8816128ce565b73ffffffffffffffffffffffffffffffffffffffff16825291830191908301906001016136a5565b509695505050505050565b6101008101613743828567ffffffffffffffff8151168252602081015173ffffffffffffffffffffffffffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b825173ffffffffffffffffffffffffffffffffffffffff90811660808401526020840151811660a08401526040840151811660c084015260608401511660e0830152611c70565b602081526000825160a060208401526137a660c0840182612a2e565b905067ffffffffffffffff6020850151166040840152604084015173ffffffffffffffffffffffffffffffffffffffff8082166060860152606086015160808601528060808701511660a086015250508091505092915050565b60006020828403121561381257600080fd5b815167ffffffffffffffff8082111561382a57600080fd5b908301906040828603121561383e57600080fd5b61384661281c565b82518281111561385557600080fd5b613861878286016132b2565b82525060208301518281111561387657600080fd5b613882878286016132b2565b60208301525095945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808916835260c060208401526138c160c0840189612a2e565b67ffffffffffffffff97881660408501529590961660608301525091909316608082015260a0019190915292915050565b60208152600061074e602083018461344a565b8181038181111561075157610751613190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251613959818460208701612a0a565b919091019291505056fea164736f6c6343000818000a",
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

func (_OnRamp *OnRampCaller) GetAllowList(opts *bind.CallOpts, destinationChainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowList", destinationChainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAllowList(destinationChainSelector uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowList(&_OnRamp.CallOpts, destinationChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetAllowList(destinationChainSelector uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowList(&_OnRamp.CallOpts, destinationChainSelector)
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

func (_OnRamp *OnRampCaller) IsAllowListEnabled(opts *bind.CallOpts, destinationChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "isAllowListEnabled", destinationChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) IsAllowListEnabled(destinationChainSelector uint64) (bool, error) {
	return _OnRamp.Contract.IsAllowListEnabled(&_OnRamp.CallOpts, destinationChainSelector)
}

func (_OnRamp *OnRampCallerSession) IsAllowListEnabled(destinationChainSelector uint64) (bool, error) {
	return _OnRamp.Contract.IsAllowListEnabled(&_OnRamp.CallOpts, destinationChainSelector)
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

func (_OnRamp *OnRampTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "applyAllowListUpdates", destinationChainSelector, removes, adds)
}

func (_OnRamp *OnRampSession) ApplyAllowListUpdates(destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyAllowListUpdates(&_OnRamp.TransactOpts, destinationChainSelector, removes, adds)
}

func (_OnRamp *OnRampTransactorSession) ApplyAllowListUpdates(destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyAllowListUpdates(&_OnRamp.TransactOpts, destinationChainSelector, removes, adds)
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

func (_OnRamp *OnRampTransactor) DisableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "disableAllowList", destinationChainSelectors)
}

func (_OnRamp *OnRampSession) DisableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _OnRamp.Contract.DisableAllowList(&_OnRamp.TransactOpts, destinationChainSelectors)
}

func (_OnRamp *OnRampTransactorSession) DisableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _OnRamp.Contract.DisableAllowList(&_OnRamp.TransactOpts, destinationChainSelectors)
}

func (_OnRamp *OnRampTransactor) EnableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "enableAllowList", destinationChainSelectors)
}

func (_OnRamp *OnRampSession) EnableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _OnRamp.Contract.EnableAllowList(&_OnRamp.TransactOpts, destinationChainSelectors)
}

func (_OnRamp *OnRampTransactorSession) EnableAllowList(destinationChainSelectors []uint64) (*types.Transaction, error) {
	return _OnRamp.Contract.EnableAllowList(&_OnRamp.TransactOpts, destinationChainSelectors)
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

func (_OnRamp *OnRampTransactor) SetAllowListAdmin(opts *bind.TransactOpts, allowListAdmin common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAllowListAdmin", allowListAdmin)
}

func (_OnRamp *OnRampSession) SetAllowListAdmin(allowListAdmin common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowListAdmin(&_OnRamp.TransactOpts, allowListAdmin)
}

func (_OnRamp *OnRampTransactorSession) SetAllowListAdmin(allowListAdmin common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowListAdmin(&_OnRamp.TransactOpts, allowListAdmin)
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

type OnRampAllowListAddedIterator struct {
	Event *OnRampAllowListAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListAdded)
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
		it.Event = new(OnRampAllowListAdded)
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

func (it *OnRampAllowListAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListAdded struct {
	DestChainSelector uint64
	AllowList         []common.Address
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListAddedIterator{contract: _OnRamp.contract, event: "AllowListAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListAdded(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListAdded)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListAdded(log types.Log) (*OnRampAllowListAdded, error) {
	event := new(OnRampAllowListAdded)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

type OnRampAllowListDisabledIterator struct {
	Event *OnRampAllowListDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListDisabled)
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
		it.Event = new(OnRampAllowListDisabled)
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

func (it *OnRampAllowListDisabledIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListDisabled struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListDisabled(opts *bind.FilterOpts) (*OnRampAllowListDisabledIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListDisabled")
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListDisabledIterator{contract: _OnRamp.contract, event: "AllowListDisabled", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListDisabled(opts *bind.WatchOpts, sink chan<- *OnRampAllowListDisabled) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListDisabled)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListDisabled", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListDisabled(log types.Log) (*OnRampAllowListDisabled, error) {
	event := new(OnRampAllowListDisabled)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowListEnabledIterator struct {
	Event *OnRampAllowListEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListEnabled)
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
		it.Event = new(OnRampAllowListEnabled)
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

func (it *OnRampAllowListEnabledIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListEnabled struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListEnabled(opts *bind.FilterOpts) (*OnRampAllowListEnabledIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListEnabled")
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListEnabledIterator{contract: _OnRamp.contract, event: "AllowListEnabled", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListEnabled(opts *bind.WatchOpts, sink chan<- *OnRampAllowListEnabled) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListEnabled)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListEnabled", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListEnabled(log types.Log) (*OnRampAllowListEnabled, error) {
	event := new(OnRampAllowListEnabled)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowListRemovedIterator struct {
	Event *OnRampAllowListRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowListRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowListRemoved)
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
		it.Event = new(OnRampAllowListRemoved)
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

func (it *OnRampAllowListRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowListRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowListRemoved struct {
	DestChainSelector uint64
	AllowList         []common.Address
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowListRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowListRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampAllowListRemovedIterator{contract: _OnRamp.contract, event: "AllowListRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowListRemoved(opts *bind.WatchOpts, sink chan<- *OnRampAllowListRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowListRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowListRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "AllowListRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowListRemoved(log types.Log) (*OnRampAllowListRemoved, error) {
	event := new(OnRampAllowListRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "AllowListRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampCCIPSendRequestedIterator struct {
	Event *OnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampCCIPSendRequested)
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
		it.Event = new(OnRampCCIPSendRequested)
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

func (it *OnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampCCIPSendRequested struct {
	DestChainSelector uint64
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampCCIPSendRequestedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPSendRequested", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPSendRequestedIterator{contract: _OnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *OnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPSendRequested", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampCCIPSendRequested)
				if err := _OnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseCCIPSendRequested(log types.Log) (*OnRampCCIPSendRequested, error) {
	event := new(OnRampCCIPSendRequested)
	if err := _OnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

type OnRampFeePaidIterator struct {
	Event *OnRampFeePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeePaid)
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
		it.Event = new(OnRampFeePaid)
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

func (it *OnRampFeePaidIterator) Error() error {
	return it.fail
}

func (it *OnRampFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeePaid struct {
	FeeToken      common.Address
	FeeValueJuels *big.Int
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeePaid(opts *bind.FilterOpts, feeToken []common.Address) (*OnRampFeePaidIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeePaid", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampFeePaidIterator{contract: _OnRamp.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *OnRampFeePaid, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeePaid", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeePaid)
				if err := _OnRamp.contract.UnpackLog(event, "FeePaid", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeePaid(log types.Log) (*OnRampFeePaid, error) {
	event := new(OnRampFeePaid)
	if err := _OnRamp.contract.UnpackLog(event, "FeePaid", log); err != nil {
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

func (_OnRamp *OnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRamp.abi.Events["AllowListAdded"].ID:
		return _OnRamp.ParseAllowListAdded(log)
	case _OnRamp.abi.Events["AllowListAdminSet"].ID:
		return _OnRamp.ParseAllowListAdminSet(log)
	case _OnRamp.abi.Events["AllowListDisabled"].ID:
		return _OnRamp.ParseAllowListDisabled(log)
	case _OnRamp.abi.Events["AllowListEnabled"].ID:
		return _OnRamp.ParseAllowListEnabled(log)
	case _OnRamp.abi.Events["AllowListRemoved"].ID:
		return _OnRamp.ParseAllowListRemoved(log)
	case _OnRamp.abi.Events["CCIPSendRequested"].ID:
		return _OnRamp.ParseCCIPSendRequested(log)
	case _OnRamp.abi.Events["ConfigSet"].ID:
		return _OnRamp.ParseConfigSet(log)
	case _OnRamp.abi.Events["DestChainConfigSet"].ID:
		return _OnRamp.ParseDestChainConfigSet(log)
	case _OnRamp.abi.Events["FeePaid"].ID:
		return _OnRamp.ParseFeePaid(log)
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

func (OnRampAllowListAdded) Topic() common.Hash {
	return common.HexToHash("0xe56852833d8b7c1f0dd03b91d1cff239d83ba81fbc1f10422d4cc4ae89199226")
}

func (OnRampAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (OnRampAllowListDisabled) Topic() common.Hash {
	return common.HexToHash("0x50013bb4071fed8f5c37e8a9ee6ecc157fe969ed7e303069ecbbba1e89e12064")
}

func (OnRampAllowListEnabled) Topic() common.Hash {
	return common.HexToHash("0x4d8b8e1b49157081f6c6ec663e74f97d33537bf46858985e5eaefd364410be60")
}

func (OnRampAllowListRemoved) Topic() common.Hash {
	return common.HexToHash("0x0d5e755ea090d8b16e0b3fed043532ed762c7e31a1f3884ac561cab59c7dbf1a")
}

func (OnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0x0f07cd31e53232da9125e517f09550fdde74bf43d6a0a76ebd41674dafe2ab29")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x23a1adf8ad7fad6091a4803227af2cee848c01a7c812404cade7c25636925e32")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd5ad72bc37dc7a80a8b9b9df20500046fd7341adb1be2258a540466fdd7dcef5")
}

func (OnRampFeePaid) Topic() common.Hash {
	return common.HexToHash("0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f")
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
	GetAllowList(opts *bind.CallOpts, destinationChainSelector uint64) ([]common.Address, error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetRouter(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	IsAllowListEnabled(opts *bind.CallOpts, destinationChainSelector uint64) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, destinationChainSelector uint64, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	DisableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error)

	EnableAllowList(opts *bind.TransactOpts, destinationChainSelectors []uint64) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetAllowListAdmin(opts *bind.TransactOpts, allowListAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAllowListAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListAddedIterator, error)

	WatchAllowListAdded(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListAdded(log types.Log) (*OnRampAllowListAdded, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowListAdmin []common.Address) (*OnRampAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowListAdminSet, allowListAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*OnRampAllowListAdminSet, error)

	FilterAllowListDisabled(opts *bind.FilterOpts) (*OnRampAllowListDisabledIterator, error)

	WatchAllowListDisabled(opts *bind.WatchOpts, sink chan<- *OnRampAllowListDisabled) (event.Subscription, error)

	ParseAllowListDisabled(log types.Log) (*OnRampAllowListDisabled, error)

	FilterAllowListEnabled(opts *bind.FilterOpts) (*OnRampAllowListEnabledIterator, error)

	WatchAllowListEnabled(opts *bind.WatchOpts, sink chan<- *OnRampAllowListEnabled) (event.Subscription, error)

	ParseAllowListEnabled(log types.Log) (*OnRampAllowListEnabled, error)

	FilterAllowListRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampAllowListRemovedIterator, error)

	WatchAllowListRemoved(opts *bind.WatchOpts, sink chan<- *OnRampAllowListRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListRemoved(log types.Log) (*OnRampAllowListRemoved, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *OnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*OnRampCCIPSendRequested, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error)

	FilterFeePaid(opts *bind.FilterOpts, feeToken []common.Address) (*OnRampFeePaidIterator, error)

	WatchFeePaid(opts *bind.WatchOpts, sink chan<- *OnRampFeePaid, feeToken []common.Address) (event.Subscription, error)

	ParseFeePaid(log types.Log) (*OnRampFeePaid, error)

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
