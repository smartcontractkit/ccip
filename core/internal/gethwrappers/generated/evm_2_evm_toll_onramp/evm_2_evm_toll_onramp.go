// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_toll_onramp

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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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
)

type BaseOnRampInterfaceOnRampConfig struct {
	RelayingFeeJuels uint64
	MaxDataSize      uint64
	MaxTokensLength  uint64
}

type CCIPEVM2AnyTollMessage struct {
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

type CCIPEVM2EVMTollEvent struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

var EVM2EVMTollOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenConfigMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.EVM2EVMTollEvent\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2AnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162004287380380620042878339810160408190526200003491620007fc565b6000805460ff191681558a908a908a908a908a908a908a908a908a908a9085908990889082908b90899089903390819081620000b75760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000f157620000f181620004b5565b5050506001600160a01b038216158062000109575080155b156200012857604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200016e5760405162d8548360e71b815260040160405180910390fd5b81516200018390600590602085019062000566565b5060005b825181101562000265576000828281518110620001a857620001a86200090e565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001f257620001f26200090e565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff191660011790556200025d8162000924565b905062000187565b50505080518251146200028b5760405163ee9d106b60e01b815260040160405180910390fd5b8151620002a090600890602085019062000566565b5060005b82518110156200036b576000828281518110620002c557620002c56200090e565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600760008685815181106200030f576200030f6200090e565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b029290911691909117905550620003638162000924565b9050620002a4565b505081511590506200039b576009805460ff1916600117905580516200039990600a90602084019062000566565b505b60005b815181101562000408576001600b6000848481518110620003c357620003c36200090e565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620004008162000924565b90506200039e565b505060809990995260a09790975250508451600d805460208801516040909801516001600160401b03908116600160801b02600160801b600160c01b031999821668010000000000000000026001600160801b031990931691909416171796909616179094555050600e80546001600160a01b039094166001600160a01b0319909416939093179092555050600c80546001600160401b0319169055506200094c98505050505050505050565b336001600160a01b038216036200050f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ae565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620005be579160200282015b82811115620005be57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000587565b50620005cc929150620005d0565b5090565b5b80821115620005cc5760008155600101620005d1565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006285762000628620005e7565b604052919050565b60006001600160401b038211156200064c576200064c620005e7565b5060051b60200190565b6001600160a01b03811681146200066c57600080fd5b50565b600082601f8301126200068157600080fd5b815160206200069a620006948362000630565b620005fd565b82815260059290921b84018101918181019086841115620006ba57600080fd5b8286015b84811015620006e2578051620006d48162000656565b8352918301918301620006be565b509695505050505050565b600082601f830112620006ff57600080fd5b8151602062000712620006948362000630565b82815260059290921b840181019181810190868411156200073257600080fd5b8286015b84811015620006e25780516200074c8162000656565b835291830191830162000736565b8051620007678162000656565b919050565b80516001600160401b03811681146200076757600080fd5b6000606082840312156200079757600080fd5b604051606081016001600160401b0381118282101715620007bc57620007bc620005e7565b604052905080620007cd836200076c565b8152620007dd602084016200076c565b6020820152620007f0604084016200076c565b60408201525092915050565b6000806000806000806000806000806101808b8d0312156200081d57600080fd5b8a5160208c015160408d0151919b5099506001600160401b03808211156200084457600080fd5b620008528e838f016200066f565b995060608d01519150808211156200086957600080fd5b620008778e838f016200066f565b985060808d01519150808211156200088e57600080fd5b6200089c8e838f01620006ed565b975060a08d0151915080821115620008b357600080fd5b50620008c28d828e01620006ed565b955050620008d360c08c016200075a565b935060e08b01519250620008ec8c6101008d0162000784565b9150620008fd6101608c016200075a565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000600182016200094557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05161390e6200097960003960006103cd01526000818161050501526109bc015261390e6000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c806379ba509711610160578063b6608c3b116100d8578063d0d5de611161008c578063eb511dd411610071578063eb511dd4146106c7578063f2fde38b146106da578063f78faa32146106ed57600080fd5b8063d0d5de61146106a1578063d7644ba2146106b457600080fd5b8063c0d78655116100bd578063c0d78655146105d7578063c3f909d4146105ea578063c5eff3d01461068c57600080fd5b8063b6608c3b1461058b578063bbe4f6db1461059e57600080fd5b806389c065681161012f578063b034909c11610114578063b034909c14610552578063b0f479a11461055a578063b4069b311461057857600080fd5b806389c06568146105275780638da5cb5b1461052f57600080fd5b806379ba5097146104e857806381be8fa4146104f05780638456cb59146104f857806385e1f4d01461050057600080fd5b80634120fccd1161020e5780635b16ebb7116101c2578063671dc337116101a7578063671dc337146104ad578063681fba16146104c0578063744b92e2146104d557600080fd5b80635b16ebb7146104695780635c975abb146104a257600080fd5b8063567c814b116101f3578063567c814b146104205780635853c6271461044357806359e96b5b1461045657600080fd5b80634120fccd14610405578063552b818b1461040d57600080fd5b8063181f5a77116102655780632b898c251161024a5780632b898c25146103b55780632ea02369146103c85780633f4ba83a146103fd57600080fd5b8063181f5a771461034e5780632222dd421461039757600080fd5b806304c2a34a1461029757806305afe24a146102d4578063108ee5fc1461030057806316b8e73114610315575b600080fd5b6102aa6102a5366004612dee565b6106f8565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6102e76102e2366004613033565b610729565b60405167ffffffffffffffff90911681526020016102cb565b61031361030e366004612dee565b610b04565b005b6102aa610323366004612dee565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b61038a6040518060400160405280601781526020017f45564d3245564d546f6c6c4f6e52616d7020312e302e3000000000000000000081525081565b6040516102cb91906131ab565b60025473ffffffffffffffffffffffffffffffffffffffff166102aa565b6103136103c33660046131be565b610be0565b6103ef7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102cb565b610313610fb0565b6102e7610fc2565b61031361041b3660046131f7565b610fe2565b61043361042e36600461326c565b6111cf565b60405190151581526020016102cb565b6103136104513660046131be565b611315565b610313610464366004613285565b611524565b610433610477366004612dee565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610433565b6103136104bb3660046132c6565b6115a2565b6104c86115f4565b6040516102cb919061332f565b6103136104e33660046131be565b6116d3565b610313611ac8565b6104c8611bea565b610313611c59565b6103ef7f000000000000000000000000000000000000000000000000000000000000000081565b6104c8611c69565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102aa565b6003546103ef565b600e5473ffffffffffffffffffffffffffffffffffffffff166102aa565b6102aa610586366004612dee565b611cd6565b61031361059936600461326c565b611ddb565b6102aa6105ac366004612dee565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6103136105e5366004612dee565b611e5b565b61065960408051606081018252600080825260208201819052918101919091525060408051606081018252600d5467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000909104169181019190915290565b60408051825167ffffffffffffffff908116825260208085015182169083015292820151909216908201526060016102cb565b610694611ed6565b6040516102cb9190613342565b6103ef6106af366004612dee565b611f43565b6103136106c23660046133aa565b612062565b6103136106d53660046131be565b6120c9565b6103136106e8366004612dee565b612309565b60095460ff16610433565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600460205260408120549091165b92915050565b6000805460ff161561079c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610809573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082d91906133c7565b15610863576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156108d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f791906133e4565b905060035481602001514261090c919061346f565b1115610944576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e5473ffffffffffffffffffffffffffffffffffffffff163314610995576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109ae846020015151856040015186606001518661231d565b6040805161014081019091527f00000000000000000000000000000000000000000000000000000000000000008152600c805460009291602083019184906109ff9067ffffffffffffffff16613486565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001866000015173ffffffffffffffffffffffffffffffffffffffff168152602001866020015181526020018660400151815260200186606001518152602001866080015173ffffffffffffffffffffffffffffffffffffffff1681526020018660a0015181526020018660c0015181525090507fab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d48181604051610af091906134dd565b60405180910390a160200151949350505050565b610b0c61267e565b73ffffffffffffffffffffffffffffffffffffffff8116610b59576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610be861267e565b6008546000819003610c26576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610cc1576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610d2a576040517f9403a50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006008610d3960018561346f565b81548110610d4957610d49613609565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff1681548110610d9b57610d9b613609565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166008610dca60018661346f565b81548110610dda57610dda613609565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610e4857610e48613609565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610eea57610eea613638565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b610fb861267e565b610fc0612704565b565b600c54600090610fdd9067ffffffffffffffff166001613667565b905090565b610fea61267e565b6000600a80548060200260200160405190810160405280929190818152602001828054801561104f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611024575b5050505050905060005b81518110156110e7576000600b600084848151811061107a5761107a613609565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556110e081613693565b9050611059565b506110f4600a8484612d28565b5060005b82811015611190576001600b600086868581811061111857611118613609565b905060200201602081019061112d9190612dee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905561118981613693565b90506110f8565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda83836040516111c29291906136cb565b60405180910390a1505050565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa15801561123f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126391906133c7565b1580156107235750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156112db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ff91906133e4565b6020015161130d908461346f565b111592915050565b61131d61267e565b73ffffffffffffffffffffffffffffffffffffffff82161580611354575073ffffffffffffffffffffffffffffffffffffffff8116155b1561138b576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015611427576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb60130591016111c2565b61152c61267e565b61154d73ffffffffffffffffffffffffffffffffffffffff841683836127e5565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8906060016111c2565b6115aa61267e565b80600d6115b78282613731565b9050507fcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e816040516115e9919061381f565b60405180910390a150565b60055460609067ffffffffffffffff81111561161257611612612e0b565b60405190808252806020026020018201604052801561163b578160200160208202803683370190505b50905060005b6005548110156116cf576116886005828154811061166157611661613609565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16611cd6565b82828151811061169a5761169a613609565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526116c881613693565b9050611641565b5090565b6116db61267e565b6005546000819003611719576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906117b4576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff161461181d576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600561182c60018561346f565b8154811061183c5761183c613609565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff168154811061188e5761188e613609565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056118bd60018661346f565b815481106118cd576118cd613609565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff168154811061193b5761193b613609565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560058054806119dd576119dd613638565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610fa1565b60015473ffffffffffffffffffffffffffffffffffffffff163314611b49576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610793565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606008805480602002602001604051908101604052809291908181526020018280548015611c4f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611c24575b5050505050905090565b611c6161267e565b610fc0612877565b60606005805480602002602001604051908101604052809291908181526020018280548015611c4f5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611c24575050505050905090565b73ffffffffffffffffffffffffffffffffffffffff80821660009081526004602052604081205490911680611d37576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600460208181526040928390205483517f21df0da700000000000000000000000000000000000000000000000000000000815293519416936321df0da79380840193908290030181865afa158015611db0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dd49190613872565b9392505050565b611de361267e565b80600003611e1d576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610bd4565b611e6361267e565b600e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15906020016115e9565b6060600a805480602002602001604051908101604052809291908181526020018280548015611c4f5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611c24575050505050905090565b600080611f758373ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611fdc576040517feef7849700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610793565b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612027573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061204b919061388f565b600d54611dd4919067ffffffffffffffff166138a8565b61206a61267e565b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032906020016115e9565b6120d161267e565b73ffffffffffffffffffffffffffffffffffffffff82161580612108575073ffffffffffffffffffffffffffffffffffffffff8116155b1561213f576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156121db576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016111c2565b61231161267e565b61231a81612937565b50565b600e5473ffffffffffffffffffffffffffffffffffffffff1661236c576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166123b9576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5468010000000000000000900467ffffffffffffffff1684111561242d57600d546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910467ffffffffffffffff16600482015260248101859052604401610793565b8251600d54700100000000000000000000000000000000900467ffffffffffffffff1681118061245e575082518114155b15612495576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095460ff1680156124cd575073ffffffffffffffffffffffffffffffffffffffff82166000908152600b602052604090205460ff16155b1561251c576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610793565b60005b8181101561267657600085828151811061253b5761253b613609565b6020026020010151905060006125768273ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff81166125dd576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610793565b8073ffffffffffffffffffffffffffffffffffffffff1663503c285887858151811061260b5761260b613609565b60200260200101516040518263ffffffff1660e01b815260040161263191815260200190565b600060405180830381600087803b15801561264b57600080fd5b505af115801561265f573d6000803e3d6000fd5b5050505050508061266f90613693565b905061251f565b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610fc0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610793565b60005460ff16612770576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610793565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612872908490612a32565b505050565b60005460ff16156128e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610793565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586127bb3390565b3373ffffffffffffffffffffffffffffffffffffffff8216036129b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610793565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000612a94826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612b3e9092919063ffffffff16565b8051909150156128725780806020019051810190612ab291906133c7565b612872576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610793565b6060612b4d8484600085612b55565b949350505050565b606082471015612be7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610793565b843b612c4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610793565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612c7891906138e5565b60006040518083038185875af1925050503d8060008114612cb5576040519150601f19603f3d011682016040523d82523d6000602084013e612cba565b606091505b5091509150612cca828286612cd5565b979650505050505050565b60608315612ce4575081611dd4565b825115612cf45782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079391906131ab565b828054828255906000526020600020908101928215612da0579160200282015b82811115612da05781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190612d48565b506116cf9291505b808211156116cf5760008155600101612da8565b73ffffffffffffffffffffffffffffffffffffffff8116811461231a57600080fd5b8035612de981612dbc565b919050565b600060208284031215612e0057600080fd5b8135611dd481612dbc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612e5d57612e5d612e0b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612eaa57612eaa612e0b565b604052919050565b600082601f830112612ec357600080fd5b813567ffffffffffffffff811115612edd57612edd612e0b565b612f0e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612e63565b818152846020838601011115612f2357600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115612f5a57612f5a612e0b565b5060051b60200190565b600082601f830112612f7557600080fd5b81356020612f8a612f8583612f40565b612e63565b82815260059290921b84018101918181019086841115612fa957600080fd5b8286015b84811015612fcd578035612fc081612dbc565b8352918301918301612fad565b509695505050505050565b600082601f830112612fe957600080fd5b81356020612ff9612f8583612f40565b82815260059290921b8401810191818101908684111561301857600080fd5b8286015b84811015612fcd578035835291830191830161301c565b6000806040838503121561304657600080fd5b823567ffffffffffffffff8082111561305e57600080fd5b9084019060e0828703121561307257600080fd5b61307a612e3a565b61308383612dde565b815260208301358281111561309757600080fd5b6130a388828601612eb2565b6020830152506040830135828111156130bb57600080fd5b6130c788828601612f64565b6040830152506060830135828111156130df57600080fd5b6130eb88828601612fd8565b6060830152506130fd60808401612dde565b608082015260a083013560a082015260c083013560c082015280945050505061312860208401612dde565b90509250929050565b60005b8381101561314c578181015183820152602001613134565b8381111561315b576000848401525b50505050565b60008151808452613179816020860160208601613131565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611dd46020830184613161565b600080604083850312156131d157600080fd5b82356131dc81612dbc565b915060208301356131ec81612dbc565b809150509250929050565b6000806020838503121561320a57600080fd5b823567ffffffffffffffff8082111561322257600080fd5b818501915085601f83011261323657600080fd5b81358181111561324557600080fd5b8660208260051b850101111561325a57600080fd5b60209290920196919550909350505050565b60006020828403121561327e57600080fd5b5035919050565b60008060006060848603121561329a57600080fd5b83356132a581612dbc565b925060208401356132b581612dbc565b929592945050506040919091013590565b6000606082840312156132d857600080fd5b50919050565b600081518084526020808501945080840160005b8381101561332457815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016132f2565b509495945050505050565b602081526000611dd460208301846132de565b6020808252825182820181905260009190848201906040850190845b8181101561339057835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161335e565b50909695505050505050565b801515811461231a57600080fd5b6000602082840312156133bc57600080fd5b8135611dd48161339c565b6000602082840312156133d957600080fd5b8151611dd48161339c565b6000606082840312156133f657600080fd5b6040516060810181811067ffffffffffffffff8211171561341957613419612e0b565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561348157613481613440565b500390565b600067ffffffffffffffff8083168181036134a3576134a3613440565b6001019392505050565b600081518084526020808501945080840160005b83811015613324578151875295820195908201906001016134c1565b602081528151602082015260006020830151613505604084018267ffffffffffffffff169052565b50604083015173ffffffffffffffffffffffffffffffffffffffff8116606084015250606083015173ffffffffffffffffffffffffffffffffffffffff811660808401525060808301516101408060a0850152613566610160850183613161565b915060a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160c08701526135a284836132de565b935060c08701519150808685030160e0870152506135c083826134ad565b92505060e08501516101006135ec8187018373ffffffffffffffffffffffffffffffffffffffff169052565b860151610120868101919091529095015193019290925250919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600067ffffffffffffffff80831681851680830382111561368a5761368a613440565b01949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036136c4576136c4613440565b5060010190565b60208082528181018390526000908460408401835b86811015612fcd5782356136f381612dbc565b73ffffffffffffffffffffffffffffffffffffffff16825291830191908301906001016136e0565b67ffffffffffffffff8116811461231a57600080fd5b813561373c8161371b565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821617835560208401356137808161371b565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff0000000000000000000000000000000084161717845560408501356137cf8161371b565b77ffffffffffffffff000000000000000000000000000000008160801b16847fffffffffffffffff0000000000000000000000000000000000000000000000008516178317178555505050505050565b60608101823561382e8161371b565b67ffffffffffffffff908116835260208401359061384b8261371b565b90811660208401526040840135906138628261371b565b8082166040850152505092915050565b60006020828403121561388457600080fd5b8151611dd481612dbc565b6000602082840312156138a157600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156138e0576138e0613440565b500290565b600082516138f7818460208701613131565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMTollOnRampABI = EVM2EVMTollOnRampMetaData.ABI

var EVM2EVMTollOnRampBin = EVM2EVMTollOnRampMetaData.Bin

func DeployEVM2EVMTollOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainId *big.Int, tokens []common.Address, pools []common.Address, feeds []common.Address, allowlist []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, config BaseOnRampInterfaceOnRampConfig, router common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOnRamp, error) {
	parsed, err := EVM2EVMTollOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOnRampBin), backend, chainId, destinationChainId, tokens, pools, feeds, allowlist, afn, maxTimeWithoutAFNSignal, config, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOnRamp{EVM2EVMTollOnRampCaller: EVM2EVMTollOnRampCaller{contract: contract}, EVM2EVMTollOnRampTransactor: EVM2EVMTollOnRampTransactor{contract: contract}, EVM2EVMTollOnRampFilterer: EVM2EVMTollOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMTollOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOnRampCaller
	EVM2EVMTollOnRampTransactor
	EVM2EVMTollOnRampFilterer
}

type EVM2EVMTollOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOnRampSession struct {
	Contract     *EVM2EVMTollOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOnRampCallerSession struct {
	Contract *EVM2EVMTollOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOnRampTransactorSession struct {
	Contract     *EVM2EVMTollOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOnRampRaw struct {
	Contract *EVM2EVMTollOnRamp
}

type EVM2EVMTollOnRampCallerRaw struct {
	Contract *EVM2EVMTollOnRampCaller
}

type EVM2EVMTollOnRampTransactorRaw struct {
	Contract *EVM2EVMTollOnRampTransactor
}

func NewEVM2EVMTollOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRamp{address: address, abi: abi, EVM2EVMTollOnRampCaller: EVM2EVMTollOnRampCaller{contract: contract}, EVM2EVMTollOnRampTransactor: EVM2EVMTollOnRampTransactor{contract: contract}, EVM2EVMTollOnRampFilterer: EVM2EVMTollOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOnRampCaller, error) {
	contract, err := bindEVM2EVMTollOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMTollOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOnRampTransactor, error) {
	contract, err := bindEVM2EVMTollOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOnRampFilterer, error) {
	contract, err := bindEVM2EVMTollOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOnRamp.Contract.EVM2EVMTollOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.EVM2EVMTollOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.EVM2EVMTollOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.CHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.CHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.DESTINATIONCHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.DESTINATIONCHAINID(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAFN(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAFN(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlist(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlist(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOnRampInterfaceOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOnRampInterfaceOnRampConfig)).(*BaseOnRampInterfaceOnRampConfig)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMTollOnRamp.Contract.GetConfig(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMTollOnRamp.Contract.GetConfig(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationToken(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationToken(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMTollOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMTollOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetFeed(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeed(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeed(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetFeedTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeedTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetFeedTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPool(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPool(&_EVM2EVMTollOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPoolTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetPoolTokens(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getRequiredFee", feeToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetRequiredFee(feeToken common.Address) (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRequiredFee(&_EVM2EVMTollOnRamp.CallOpts, feeToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetRequiredFee(feeToken common.Address) (*big.Int, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRequiredFee(&_EVM2EVMTollOnRamp.CallOpts, feeToken)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRouter(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetRouter(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "getTokenPool", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetTokenPool(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.GetTokenPool(&_EVM2EVMTollOnRamp.CallOpts, token)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsHealthy(&_EVM2EVMTollOnRamp.CallOpts, timeNow)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsHealthy(&_EVM2EVMTollOnRamp.CallOpts, timeNow)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsPool(&_EVM2EVMTollOnRamp.CallOpts, addr)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.IsPool(&_EVM2EVMTollOnRamp.CallOpts, addr)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.Owner(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOnRamp.Contract.Owner(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Paused() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.Paused(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOnRamp.Contract.Paused(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOnRamp.Contract.TypeAndVersion(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOnRamp.Contract.TypeAndVersion(&_EVM2EVMTollOnRamp.CallOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AcceptOwnership(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AcceptOwnership(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "addFeed", token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddPool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.AddPool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "forwardFromRouter", message, originalSender)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) ForwardFromRouter(message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.ForwardFromRouter(&_EVM2EVMTollOnRamp.TransactOpts, message, originalSender)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) ForwardFromRouter(message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.ForwardFromRouter(&_EVM2EVMTollOnRamp.TransactOpts, message, originalSender)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Pause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Pause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "removeFeed", token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemoveFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemoveFeed(&_EVM2EVMTollOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemovePool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.RemovePool(&_EVM2EVMTollOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAFN(&_EVM2EVMTollOnRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAFN(&_EVM2EVMTollOnRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlist(&_EVM2EVMTollOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlist(&_EVM2EVMTollOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMTollOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMTollOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetConfig(&_EVM2EVMTollOnRamp.TransactOpts, config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetConfig(&_EVM2EVMTollOnRamp.TransactOpts, config)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOnRamp.TransactOpts, newTime)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOnRamp.TransactOpts, newTime)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetRouter(&_EVM2EVMTollOnRamp.TransactOpts, router)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.SetRouter(&_EVM2EVMTollOnRamp.TransactOpts, router)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.TransferOwnership(&_EVM2EVMTollOnRamp.TransactOpts, to)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.TransferOwnership(&_EVM2EVMTollOnRamp.TransactOpts, to)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Unpause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.Unpause(&_EVM2EVMTollOnRamp.TransactOpts)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.WithdrawAccumulatedFees(&_EVM2EVMTollOnRamp.TransactOpts, feeToken, recipient, amount)
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOnRamp.Contract.WithdrawAccumulatedFees(&_EVM2EVMTollOnRamp.TransactOpts, feeToken, recipient, amount)
}

type EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet)
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

func (it *EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet, error) {
	event := new(EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampAFNSetIterator struct {
	Event *EVM2EVMTollOnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAFNSet)
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
		it.Event = new(EVM2EVMTollOnRampAFNSet)
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

func (it *EVM2EVMTollOnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAFNSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAFNSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOnRampAFNSet, error) {
	event := new(EVM2EVMTollOnRampAFNSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampAllowListEnabledSetIterator struct {
	Event *EVM2EVMTollOnRampAllowListEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAllowListEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAllowListEnabledSet)
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
		it.Event = new(EVM2EVMTollOnRampAllowListEnabledSet)
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

func (it *EVM2EVMTollOnRampAllowListEnabledSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAllowListEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAllowListEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListEnabledSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAllowListEnabledSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AllowListEnabledSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListEnabledSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAllowListEnabledSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAllowListEnabledSet(log types.Log) (*EVM2EVMTollOnRampAllowListEnabledSet, error) {
	event := new(EVM2EVMTollOnRampAllowListEnabledSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampAllowListSetIterator struct {
	Event *EVM2EVMTollOnRampAllowListSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampAllowListSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampAllowListSet)
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
		it.Event = new(EVM2EVMTollOnRampAllowListSet)
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

func (it *EVM2EVMTollOnRampAllowListSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampAllowListSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampAllowListSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampAllowListSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseAllowListSet(log types.Log) (*EVM2EVMTollOnRampAllowListSet, error) {
	event := new(EVM2EVMTollOnRampAllowListSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMTollOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMTollOnRampCCIPSendRequested)
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

func (it *EVM2EVMTollOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampCCIPSendRequested struct {
	Message CCIPEVM2EVMTollEvent
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMTollOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampCCIPSendRequestedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampCCIPSendRequested)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMTollOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMTollOnRampCCIPSendRequested)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeeChargedIterator struct {
	Event *EVM2EVMTollOnRampFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeeCharged)
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
		it.Event = new(EVM2EVMTollOnRampFeeCharged)
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

func (it *EVM2EVMTollOnRampFeeChargedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeeChargedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeeChargedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeeCharged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeeCharged)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeeCharged(log types.Log) (*EVM2EVMTollOnRampFeeCharged, error) {
	event := new(EVM2EVMTollOnRampFeeCharged)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeedAddedIterator struct {
	Event *EVM2EVMTollOnRampFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeedAdded)
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
		it.Event = new(EVM2EVMTollOnRampFeedAdded)
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

func (it *EVM2EVMTollOnRampFeedAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeedAddedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeedAdded)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeedAdded(log types.Log) (*EVM2EVMTollOnRampFeedAdded, error) {
	event := new(EVM2EVMTollOnRampFeedAdded)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeedRemovedIterator struct {
	Event *EVM2EVMTollOnRampFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeedRemoved)
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
		it.Event = new(EVM2EVMTollOnRampFeedRemoved)
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

func (it *EVM2EVMTollOnRampFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeedRemovedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeedRemoved)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeedRemoved(log types.Log) (*EVM2EVMTollOnRampFeedRemoved, error) {
	event := new(EVM2EVMTollOnRampFeedRemoved)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampFeesWithdrawnIterator struct {
	Event *EVM2EVMTollOnRampFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampFeesWithdrawn)
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
		it.Event = new(EVM2EVMTollOnRampFeesWithdrawn)
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

func (it *EVM2EVMTollOnRampFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeesWithdrawnIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampFeesWithdrawnIterator{contract: _EVM2EVMTollOnRamp.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampFeesWithdrawn)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseFeesWithdrawn(log types.Log) (*EVM2EVMTollOnRampFeesWithdrawn, error) {
	event := new(EVM2EVMTollOnRampFeesWithdrawn)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampOnRampConfigSetIterator struct {
	Event *EVM2EVMTollOnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampOnRampConfigSet)
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
		it.Event = new(EVM2EVMTollOnRampOnRampConfigSet)
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

func (it *EVM2EVMTollOnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampOnRampConfigSet struct {
	Config BaseOnRampInterfaceOnRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampOnRampConfigSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampOnRampConfigSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseOnRampConfigSet(log types.Log) (*EVM2EVMTollOnRampOnRampConfigSet, error) {
	event := new(EVM2EVMTollOnRampOnRampConfigSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMTollOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMTollOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampOwnershipTransferRequested)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOnRampOwnershipTransferRequested)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMTollOnRampOwnershipTransferred)
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

func (it *EVM2EVMTollOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampOwnershipTransferredIterator{contract: _EVM2EVMTollOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampOwnershipTransferred)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMTollOnRampOwnershipTransferred)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampPausedIterator struct {
	Event *EVM2EVMTollOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampPaused)
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
		it.Event = new(EVM2EVMTollOnRampPaused)
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

func (it *EVM2EVMTollOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampPausedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampPaused)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOnRampPaused, error) {
	event := new(EVM2EVMTollOnRampPaused)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampPoolAddedIterator struct {
	Event *EVM2EVMTollOnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampPoolAdded)
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
		it.Event = new(EVM2EVMTollOnRampPoolAdded)
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

func (it *EVM2EVMTollOnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampPoolAddedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampPoolAdded)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOnRampPoolAdded, error) {
	event := new(EVM2EVMTollOnRampPoolAdded)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampPoolRemovedIterator struct {
	Event *EVM2EVMTollOnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampPoolRemoved)
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
		it.Event = new(EVM2EVMTollOnRampPoolRemoved)
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

func (it *EVM2EVMTollOnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampPoolRemovedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampPoolRemoved)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOnRampPoolRemoved, error) {
	event := new(EVM2EVMTollOnRampPoolRemoved)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampRouterSetIterator struct {
	Event *EVM2EVMTollOnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampRouterSet)
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
		it.Event = new(EVM2EVMTollOnRampRouterSet)
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

func (it *EVM2EVMTollOnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampRouterSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampRouterSetIterator{contract: _EVM2EVMTollOnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampRouterSet)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseRouterSet(log types.Log) (*EVM2EVMTollOnRampRouterSet, error) {
	event := new(EVM2EVMTollOnRampRouterSet)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOnRampUnpausedIterator struct {
	Event *EVM2EVMTollOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOnRampUnpaused)
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
		it.Event = new(EVM2EVMTollOnRampUnpaused)
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

func (it *EVM2EVMTollOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOnRampUnpausedIterator{contract: _EVM2EVMTollOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOnRampUnpaused)
				if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOnRampUnpaused, error) {
	event := new(EVM2EVMTollOnRampUnpaused)
	if err := _EVM2EVMTollOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOnRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAFNSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["AllowListSet"].ID:
		return _EVM2EVMTollOnRamp.ParseAllowListSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMTollOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeeCharged"].ID:
		return _EVM2EVMTollOnRamp.ParseFeeCharged(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeedAdded"].ID:
		return _EVM2EVMTollOnRamp.ParseFeedAdded(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeedRemoved"].ID:
		return _EVM2EVMTollOnRamp.ParseFeedRemoved(log)
	case _EVM2EVMTollOnRamp.abi.Events["FeesWithdrawn"].ID:
		return _EVM2EVMTollOnRamp.ParseFeesWithdrawn(log)
	case _EVM2EVMTollOnRamp.abi.Events["OnRampConfigSet"].ID:
		return _EVM2EVMTollOnRamp.ParseOnRampConfigSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOnRamp.abi.Events["Paused"].ID:
		return _EVM2EVMTollOnRamp.ParsePaused(log)
	case _EVM2EVMTollOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOnRamp.ParsePoolAdded(log)
	case _EVM2EVMTollOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOnRamp.ParsePoolRemoved(log)
	case _EVM2EVMTollOnRamp.abi.Events["RouterSet"].ID:
		return _EVM2EVMTollOnRamp.ParseRouterSet(log)
	case _EVM2EVMTollOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (EVM2EVMTollOnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOnRampAllowListEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032")
}

func (EVM2EVMTollOnRampAllowListSet) Topic() common.Hash {
	return common.HexToHash("0xf8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda")
}

func (EVM2EVMTollOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xab2ca9da6d303be28d1a5e854e3e170be286e07696245e77f8ea11f55367d481")
}

func (EVM2EVMTollOnRampFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (EVM2EVMTollOnRampFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (EVM2EVMTollOnRampFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (EVM2EVMTollOnRampFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (EVM2EVMTollOnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e")
}

func (EVM2EVMTollOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (EVM2EVMTollOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOnRamp *EVM2EVMTollOnRamp) Address() common.Address {
	return _EVM2EVMTollOnRamp.address
}

type EVM2EVMTollOnRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRequiredFee(opts *bind.CallOpts, feeToken common.Address) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error)

	IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnyTollMessage, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMTollOnRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOnRampAFNSet, error)

	FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListEnabledSetIterator, error)

	WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListEnabledSet) (event.Subscription, error)

	ParseAllowListEnabledSet(log types.Log) (*EVM2EVMTollOnRampAllowListEnabledSet, error)

	FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampAllowListSetIterator, error)

	WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampAllowListSet) (event.Subscription, error)

	ParseAllowListSet(log types.Log) (*EVM2EVMTollOnRampAllowListSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMTollOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMTollOnRampCCIPSendRequested, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2EVMTollOnRampFeeCharged, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*EVM2EVMTollOnRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*EVM2EVMTollOnRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2EVMTollOnRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*EVM2EVMTollOnRampFeesWithdrawn, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*EVM2EVMTollOnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMTollOnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*EVM2EVMTollOnRampRouterSet, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
