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

type CCIPMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Payload        CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Tokens             []common.Address
	Amounts            []*big.Int
	DestinationChainId *big.Int
	Receiver           common.Address
	Executor           common.Address
	Data               []byte
	Options            []byte
}

type OnRampInterfaceOnRampConfig struct {
	Router           common.Address
	RelayingFeeJuels uint64
	MaxDataSize      uint64
	MaxTokensLength  uint64
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"destinationChainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMistmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMistmatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowlistEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowlistSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CrossChainSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"full\",\"type\":\"bool\"}],\"name\":\"NewTokenBucketConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationChains\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"getSequenceNumberOfDestinationChain\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"requestCrossChainSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620042b6380380620042b6833981016040819052620000349162000904565b6000805460ff191681558790869082908990879087903390819081620000a15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000db57620000db8162000506565b5050506001600160a01b0382161580620000f3575080155b156200011257604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001585760405162d8548360e71b815260040160405180910390fd5b81516200016d906005906020850190620005b7565b5060005b82518110156200025157600082828151811062000192576200019262000a20565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001dc57620001dc62000a20565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002488162000a36565b91505062000171565b5050508051825114620002775760405163ee9d106b60e01b815260040160405180910390fd5b81516200028c906008906020850190620005b7565b5060005b825181101562000359576000828281518110620002b157620002b162000a20565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060076000868581518110620002fb57620002fb62000a20565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555080620003508162000a36565b91505062000290565b505050608089905287516200037690600a9060208b019062000621565b5060005b8851811015620003ef576001600960008b84815181106200039f576200039f62000a20565b6020026020010151815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b031602179055508080620003e69062000a36565b9150506200037a565b508351156200041c57600d805460ff1916600117905583516200041a90600f906020870190620005b7565b505b60005b84518110156200048b576001600e600087848151811062000444576200044462000a20565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580620004828162000a36565b9150506200041f565b508051600b805460208401516001600160a01b039093166001600160e01b031990911617600160a01b6001600160401b03938416021790556040820151600c80546060909401519183166001600160801b0319909416939093176801000000000000000091909216021790555062000a5e9650505050505050565b336001600160a01b03821603620005605760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000098565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200060f579160200282015b828111156200060f57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620005d8565b506200061d9291506200065f565b5090565b8280548282559060005260206000209081019282156200060f579160200282015b828111156200060f57825182559160200191906001019062000642565b5b808211156200061d576000815560010162000660565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006b757620006b762000676565b604052919050565b60006001600160401b03821115620006db57620006db62000676565b5060051b60200190565b600082601f830112620006f757600080fd5b81516020620007106200070a83620006bf565b6200068c565b82815260059290921b840181019181810190868411156200073057600080fd5b8286015b848110156200074d578051835291830191830162000734565b509695505050505050565b6001600160a01b03811681146200076e57600080fd5b50565b600082601f8301126200078357600080fd5b81516020620007966200070a83620006bf565b82815260059290921b84018101918181019086841115620007b657600080fd5b8286015b848110156200074d578051620007d08162000758565b8352918301918301620007ba565b600082601f830112620007f057600080fd5b81516020620008036200070a83620006bf565b82815260059290921b840181019181810190868411156200082357600080fd5b8286015b848110156200074d5780516200083d8162000758565b835291830191830162000827565b8051620008588162000758565b919050565b80516001600160401b03811681146200085857600080fd5b6000608082840312156200088857600080fd5b604051608081016001600160401b0381118282101715620008ad57620008ad62000676565b80604052508091508251620008c28162000758565b8152620008d2602084016200085d565b6020820152620008e5604084016200085d565b6040820152620008f8606084016200085d565b60608201525092915050565b60008060008060008060008060006101808a8c0312156200092457600080fd5b895160208b01519099506001600160401b03808211156200094457600080fd5b620009528d838e01620006e5565b995060408c01519150808211156200096957600080fd5b620009778d838e0162000771565b985060608c01519150808211156200098e57600080fd5b6200099c8d838e0162000771565b975060808c0151915080821115620009b357600080fd5b620009c18d838e01620007de565b965060a08c0151915080821115620009d857600080fd5b50620009e78c828d0162000771565b945050620009f860c08b016200084b565b925060e08a0151915062000a118b6101008c0162000875565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b60006001820162000a5757634e487b7160e01b600052601160045260246000fd5b5060010190565b60805161383562000a816000396000818161042b015261188e01526138356000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806379ba50971161010f578063b6608c3b116100a2578063d7644ba211610071578063d7644ba2146105fb578063eb511dd41461060e578063f2fde38b14610621578063f78faa321461063457600080fd5b8063b6608c3b146104a3578063bbe4f6db146104b6578063c3f909d4146104ef578063c5eff3d0146105e657600080fd5b806389c06568116100de57806389c065681461045b5780638da5cb5b14610463578063abc343a714610486578063b034909c1461049b57600080fd5b806379ba50971461040157806381be8fa4146104095780638456cb591461041e57806385e1f4d01461042657600080fd5b8063552b818b116101875780635b16ebb7116101565780635b16ebb7146103975780635c975abb146103d0578063625f9e19146103db578063744b92e2146103ee57600080fd5b8063552b818b1461033b578063567c814b1461034e5780635853c6271461037157806359e96b5b1461038457600080fd5b80632222dd42116101c35780632222dd42146102ef5780632b898c251461030d5780632df836c0146103205780633f4ba83a1461033357600080fd5b8063108ee5fc146101f557806316b8e7311461020a578063181f5a771461026d57806318797167146102ac575b600080fd5b610208610203366004612cd4565b61063f565b005b610243610218366004612cd4565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b604080518082018252600c81527f4f6e52616d7020302e302e310000000000000000000000000000000000000000602082015290516102649190612d67565b6102d66102ba366004612d7a565b60009081526009602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610264565b60025473ffffffffffffffffffffffffffffffffffffffff16610243565b61020861031b366004612d93565b61071b565b61020861032e366004612dcc565b610aeb565b610208610b3d565b610208610349366004612de4565b610b4f565b61036161035c366004612d7a565b610d40565b6040519015158152602001610264565b61020861037f366004612d93565b610e87565b610208610392366004612e69565b611096565b6103616103a5366004612cd4565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610361565b6102d66103e93660046130d2565b611114565b6102086103fc366004612d93565b611980565b610208611d75565b610411611e97565b604051610264919061323a565b610208611f06565b61044d7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610264565b610411611f16565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610243565b61048e611f83565b604051610264919061327d565b60035461044d565b6102086104b1366004612d7a565b611fda565b6102436104c4366004612cd4565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b61058b6040805160808101825260008082526020820181905291810182905260608101919091525060408051608081018252600b5473ffffffffffffffffffffffffffffffffffffffff8116825267ffffffffffffffff7401000000000000000000000000000000000000000090910481166020830152600c548082169383019390935268010000000000000000909204909116606082015290565b6040516102649190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b6105ee61205a565b6040516102649190613290565b6102086106093660046132f8565b6120c7565b61020861061c366004612d93565b61212e565b61020861062f366004612cd4565b61236e565b600d5460ff16610361565b610647612382565b73ffffffffffffffffffffffffffffffffffffffff8116610694576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610723612382565b6008546000819003610761576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906107fc576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610865576040517f6c17b98700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006008610874600185613344565b815481106108845761088461335b565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff16815481106108d6576108d661335b565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166008610905600186613344565b815481106109155761091561335b565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff16815481106109835761098361335b565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610a2557610a2561338a565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b610af3612382565b80600b610b0082826133cf565b9050507feac62265bdcb30e1e7a4822fecd5035bf208f242c899453ca9a3cdb5eb44225b81604051610b3291906134e9565b60405180910390a150565b610b45612382565b610b4d612408565b565b610b57612382565b6000600f805480602002602001604051908101604052809291908181526020018280548015610bbc57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610b91575b5050505050905060005b8151811015610c56576000600e6000848481518110610be757610be761335b565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610c4e81613565565b915050610bc6565b50610c63600f8484612c15565b5060005b82811015610d01576001600e6000868685818110610c8757610c8761335b565b9050602002016020810190610c9c9190612cd4565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610cf981613565565b915050610c67565b507f27f242de1bc4ed72c4329591ffff7d223b5f025e3514a07e05afec6d4eb889cf8383604051610d3392919061359d565b60405180910390a1505050565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610db0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd491906135ed565b158015610e815750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610e4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e70919061360a565b60200151610e7e9084613344565b11155b92915050565b610e8f612382565b73ffffffffffffffffffffffffffffffffffffffff82161580610ec6575073ffffffffffffffffffffffffffffffffffffffff8116155b15610efd576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015610f99576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb6013059101610d33565b61109e612382565b6110bf73ffffffffffffffffffffffffffffffffffffffff841683836124e9565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001610d33565b6000805460ff1615611187576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121891906135ed565b1561124e576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156112be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e2919061360a565b90506003548160200151426112f79190613344565b111561132f576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff8416156113a557600b5473ffffffffffffffffffffffffffffffffffffffff8281169116146113a0576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113a9565b8093505b600d5460ff1680156113e1575073ffffffffffffffffffffffffffffffffffffffff84166000908152600e602052604090205460ff16155b15611430576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161117e565b604080860151600090815260096020529081205467ffffffffffffffff16908190036114905785604001516040517f45abe4ae00000000000000000000000000000000000000000000000000000000815260040161117e91815260200190565b600c5460a08701515167ffffffffffffffff90911610156114f857600c5460a0870151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482015260440161117e565b600c548651516801000000000000000090910467ffffffffffffffff161080611528575060208601515186515114155b1561155f576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600086600001516000815181106115785761157861335b565b60200260200101519050600061158d826125c2565b9050801561163d578088602001516000815181106115ad576115ad61335b565b602002602001018181516115c19190613344565b9052506115e673ffffffffffffffffffffffffffffffffffffffff8316853084612700565b6040805173ffffffffffffffffffffffffffffffffffffffff861681523060208201529081018290527f945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d9060600160405180910390a15b60005b885151811015611881576000896000015182815181106116625761166261335b565b60200260200101519050600061169d8273ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff8116611704576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161117e565b60008b60200151848151811061171c5761171c61335b565b6020908102919091010151905061174b73ffffffffffffffffffffffffffffffffffffffff8416893084612700565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526024820183905284169063095ea7b3906044016020604051808303816000875af11580156117c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e491906135ed565b506040517feb54b3bf0000000000000000000000000000000000000000000000000000000081523060048201526024810182905273ffffffffffffffffffffffffffffffffffffffff83169063eb54b3bf90604401600060405180830381600087803b15801561185357600080fd5b505af1158015611867573d6000803e3d6000fd5b50505050505050808061187990613565565b915050611640565b50604080516080810182527f0000000000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516602082015273ffffffffffffffffffffffffffffffffffffffff891691810191909152606081018990526118ef846001613666565b60408a8101516000908152600960205281902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff9390931692909217909155517fbc2337a0942981c229faae38f5d31c14d85be0310ab538fe20e8c92b0585d10190611968908390613692565b60405180910390a16020015198975050505050505050565b611988612382565b60055460008190036119c6576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611a61576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611aca576040517fd428911900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005611ad9600185613344565b81548110611ae957611ae961335b565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110611b3b57611b3b61335b565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005611b6a600186613344565b81548110611b7a57611b7a61335b565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110611be857611be861335b565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480611c8a57611c8a61338a565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610adc565b60015473ffffffffffffffffffffffffffffffffffffffff163314611df6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161117e565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606008805480602002602001604051908101604052809291908181526020018280548015611efc57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611ed1575b5050505050905090565b611f0e612382565b610b4d612764565b60606005805480602002602001604051908101604052809291908181526020018280548015611efc5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611ed1575050505050905090565b6060600a805480602002602001604051908101604052809291908181526020018280548015611efc57602002820191906000526020600020905b815481526020019060010190808311611fbd575050505050905090565b611fe2612382565b8060000361201c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251910161070f565b6060600f805480602002602001604051908101604052809291908181526020018280548015611efc5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611ed1575050505050905090565b6120cf612382565b600d80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fa1bf86c493917580dec207969ef59976f0c378f10ece581237f19acfbd858f1c90602001610b32565b612136612382565b73ffffffffffffffffffffffffffffffffffffffff8216158061216d575073ffffffffffffffffffffffffffffffffffffffff8116155b156121a4576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612240576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101610d33565b612376612382565b61237f81612824565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610b4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161117e565b60005460ff16612474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161117e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526125bd9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261291f565b505050565b6000806125f48373ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff811661265b576040517feef7849700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161117e565b8073ffffffffffffffffffffffffffffffffffffffff166350d25bcd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156126a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ca91906137b6565b600b546126f9919074010000000000000000000000000000000000000000900467ffffffffffffffff166137cf565b9392505050565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261275e9085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161253b565b50505050565b60005460ff16156127d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161117e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586124bf3390565b3373ffffffffffffffffffffffffffffffffffffffff8216036128a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161117e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000612981826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612a2b9092919063ffffffff16565b8051909150156125bd578080602001905181019061299f91906135ed565b6125bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161117e565b6060612a3a8484600085612a42565b949350505050565b606082471015612ad4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161117e565b843b612b3c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161117e565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612b65919061380c565b60006040518083038185875af1925050503d8060008114612ba2576040519150601f19603f3d011682016040523d82523d6000602084013e612ba7565b606091505b5091509150612bb7828286612bc2565b979650505050505050565b60608315612bd15750816126f9565b825115612be15782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117e9190612d67565b828054828255906000526020600020908101928215612c8d579160200282015b82811115612c8d5781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190612c35565b50612c99929150612c9d565b5090565b5b80821115612c995760008155600101612c9e565b73ffffffffffffffffffffffffffffffffffffffff8116811461237f57600080fd5b600060208284031215612ce657600080fd5b81356126f981612cb2565b60005b83811015612d0c578181015183820152602001612cf4565b8381111561275e5750506000910152565b60008151808452612d35816020860160208601612cf1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006126f96020830184612d1d565b600060208284031215612d8c57600080fd5b5035919050565b60008060408385031215612da657600080fd5b8235612db181612cb2565b91506020830135612dc181612cb2565b809150509250929050565b600060808284031215612dde57600080fd5b50919050565b60008060208385031215612df757600080fd5b823567ffffffffffffffff80821115612e0f57600080fd5b818501915085601f830112612e2357600080fd5b813581811115612e3257600080fd5b8660208260051b8501011115612e4757600080fd5b60209290920196919550909350505050565b8035612e6481612cb2565b919050565b600080600060608486031215612e7e57600080fd5b8335612e8981612cb2565b92506020840135612e9981612cb2565b929592945050506040919091013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612efc57612efc612eaa565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612f4957612f49612eaa565b604052919050565b600067ffffffffffffffff821115612f6b57612f6b612eaa565b5060051b60200190565b600082601f830112612f8657600080fd5b81356020612f9b612f9683612f51565b612f02565b82815260059290921b84018101918181019086841115612fba57600080fd5b8286015b84811015612fde578035612fd181612cb2565b8352918301918301612fbe565b509695505050505050565b600082601f830112612ffa57600080fd5b8135602061300a612f9683612f51565b82815260059290921b8401810191818101908684111561302957600080fd5b8286015b84811015612fde578035835291830191830161302d565b600082601f83011261305557600080fd5b813567ffffffffffffffff81111561306f5761306f612eaa565b6130a060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612f02565b8181528460208386010111156130b557600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156130e557600080fd5b823567ffffffffffffffff808211156130fd57600080fd5b9084019060e0828703121561311157600080fd5b613119612ed9565b82358281111561312857600080fd5b61313488828601612f75565b82525060208301358281111561314957600080fd5b61315588828601612fe9565b6020830152506040830135604082015261317160608401612e59565b606082015261318260808401612e59565b608082015260a08301358281111561319957600080fd5b6131a588828601613044565b60a08301525060c0830135828111156131bd57600080fd5b6131c988828601613044565b60c08301525093506131e091505060208401612e59565b90509250929050565b600081518084526020808501945080840160005b8381101561322f57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016131fd565b509495945050505050565b6020815260006126f960208301846131e9565b600081518084526020808501945080840160005b8381101561322f57815187529582019590820190600101613261565b6020815260006126f9602083018461324d565b6020808252825182820181905260009190848201906040850190845b818110156132de57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016132ac565b50909695505050505050565b801515811461237f57600080fd5b60006020828403121561330a57600080fd5b81356126f9816132ea565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561335657613356613315565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b67ffffffffffffffff8116811461237f57600080fd5b81356133da81612cb2565b73ffffffffffffffffffffffffffffffffffffffff811690508154817fffffffffffffffffffffffff00000000000000000000000000000000000000008216178355602084013561342a816133b9565b7bffffffffffffffff00000000000000000000000000000000000000008160a01b16837fffffffff00000000000000000000000000000000000000000000000000000000841617178455505050600181016040830135613489816133b9565b81546060850135613499816133b9565b6fffffffffffffffff00000000000000008160401b1667ffffffffffffffff84167fffffffffffffffffffffffffffffffff00000000000000000000000000000000841617178455505050505050565b6080810182356134f881612cb2565b73ffffffffffffffffffffffffffffffffffffffff168252602083013561351e816133b9565b67ffffffffffffffff908116602084015260408401359061353e826133b9565b9081166040840152606084013590613555826133b9565b8082166060850152505092915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361359657613596613315565b5060010190565b60208082528181018390526000908460408401835b86811015612fde5782356135c581612cb2565b73ffffffffffffffffffffffffffffffffffffffff16825291830191908301906001016135b2565b6000602082840312156135ff57600080fd5b81516126f9816132ea565b60006060828403121561361c57600080fd5b6040516060810181811067ffffffffffffffff8211171561363f5761363f612eaa565b80604052508251815260208301516020820152604083015160408201528091505092915050565b600067ffffffffffffffff80831681851680830382111561368957613689613315565b01949350505050565b602081528151602082015267ffffffffffffffff60208301511660408201526000604083015173ffffffffffffffffffffffffffffffffffffffff808216606085015260608501519150608080850152815160e060a08601526136f96101808601826131e9565b905060208301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60808784030160c0880152613735838361324d565b604086015160e0890152606086015194909416610100880152608085015173ffffffffffffffffffffffffffffffffffffffff1661012088015260a0850151878503820161014089015293925061378c8385612d1d565b935060c0850151945080878503016101608801525050506137ad8183612d1d565b95945050505050565b6000602082840312156137c857600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561380757613807613315565b500290565b6000825161381e818460208701612cf1565b919091019291505056fea164736f6c634300080d000a",
}

var OnRampABI = OnRampMetaData.ABI

var OnRampBin = OnRampMetaData.Bin

func DeployOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainIds []*big.Int, tokens []common.Address, pools []common.Address, feeds []common.Address, allowlist []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, config OnRampInterfaceOnRampConfig) (common.Address, *types.Transaction, *OnRamp, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampBin), backend, chainId, destinationChainIds, tokens, pools, feeds, allowlist, afn, maxTimeWithoutAFNSignal, config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRamp{OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
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
	parsed, err := abi.JSON(strings.NewReader(OnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

func (_OnRamp *OnRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) CHAINID() (*big.Int, error) {
	return _OnRamp.Contract.CHAINID(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) CHAINID() (*big.Int, error) {
	return _OnRamp.Contract.CHAINID(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAFN() (common.Address, error) {
	return _OnRamp.Contract.GetAFN(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAFN() (common.Address, error) {
	return _OnRamp.Contract.GetAFN(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAllowlist() ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowlist(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _OnRamp.Contract.GetAllowlist(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) GetAllowlistEnabled() (bool, error) {
	return _OnRamp.Contract.GetAllowlistEnabled(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _OnRamp.Contract.GetAllowlistEnabled(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetConfig(opts *bind.CallOpts) (OnRampInterfaceOnRampConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(OnRampInterfaceOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampInterfaceOnRampConfig)).(*OnRampInterfaceOnRampConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetConfig() (OnRampInterfaceOnRampConfig, error) {
	return _OnRamp.Contract.GetConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetConfig() (OnRampInterfaceOnRampConfig, error) {
	return _OnRamp.Contract.GetConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetDestinationChains(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDestinationChains")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetDestinationChains() ([]*big.Int, error) {
	return _OnRamp.Contract.GetDestinationChains(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetDestinationChains() ([]*big.Int, error) {
	return _OnRamp.Contract.GetDestinationChains(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFeed(token common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetFeed(&_OnRamp.CallOpts, token)
}

func (_OnRamp *OnRampCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetFeed(&_OnRamp.CallOpts, token)
}

func (_OnRamp *OnRampCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFeedTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetFeedTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetFeedTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _OnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPool(&_OnRamp.CallOpts, sourceToken)
}

func (_OnRamp *OnRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPool(&_OnRamp.CallOpts, sourceToken)
}

func (_OnRamp *OnRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPoolTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetPoolTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _OnRamp.Contract.GetPoolTokens(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetSequenceNumberOfDestinationChain(opts *bind.CallOpts, destinationChainId *big.Int) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getSequenceNumberOfDestinationChain", destinationChainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetSequenceNumberOfDestinationChain(destinationChainId *big.Int) (uint64, error) {
	return _OnRamp.Contract.GetSequenceNumberOfDestinationChain(&_OnRamp.CallOpts, destinationChainId)
}

func (_OnRamp *OnRampCallerSession) GetSequenceNumberOfDestinationChain(destinationChainId *big.Int) (uint64, error) {
	return _OnRamp.Contract.GetSequenceNumberOfDestinationChain(&_OnRamp.CallOpts, destinationChainId)
}

func (_OnRamp *OnRampCaller) IsHealthy(opts *bind.CallOpts, now *big.Int) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "isHealthy", now)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) IsHealthy(now *big.Int) (bool, error) {
	return _OnRamp.Contract.IsHealthy(&_OnRamp.CallOpts, now)
}

func (_OnRamp *OnRampCallerSession) IsHealthy(now *big.Int) (bool, error) {
	return _OnRamp.Contract.IsHealthy(&_OnRamp.CallOpts, now)
}

func (_OnRamp *OnRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) IsPool(addr common.Address) (bool, error) {
	return _OnRamp.Contract.IsPool(&_OnRamp.CallOpts, addr)
}

func (_OnRamp *OnRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _OnRamp.Contract.IsPool(&_OnRamp.CallOpts, addr)
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

func (_OnRamp *OnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRamp *OnRampSession) Paused() (bool, error) {
	return _OnRamp.Contract.Paused(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) Paused() (bool, error) {
	return _OnRamp.Contract.Paused(&_OnRamp.CallOpts)
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

func (_OnRamp *OnRampTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "addFeed", token, feed)
}

func (_OnRamp *OnRampSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_OnRamp *OnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddPool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.AddPool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "pause")
}

func (_OnRamp *OnRampSession) Pause() (*types.Transaction, error) {
	return _OnRamp.Contract.Pause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _OnRamp.Contract.Pause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "removeFeed", token, feed)
}

func (_OnRamp *OnRampSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemoveFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemoveFeed(&_OnRamp.TransactOpts, token, feed)
}

func (_OnRamp *OnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_OnRamp *OnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemovePool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RemovePool(&_OnRamp.TransactOpts, token, pool)
}

func (_OnRamp *OnRampTransactor) RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "requestCrossChainSend", payload, originalSender)
}

func (_OnRamp *OnRampSession) RequestCrossChainSend(payload CCIPMessagePayload, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RequestCrossChainSend(&_OnRamp.TransactOpts, payload, originalSender)
}

func (_OnRamp *OnRampTransactorSession) RequestCrossChainSend(payload CCIPMessagePayload, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.RequestCrossChainSend(&_OnRamp.TransactOpts, payload, originalSender)
}

func (_OnRamp *OnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_OnRamp *OnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAFN(&_OnRamp.TransactOpts, afn)
}

func (_OnRamp *OnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAFN(&_OnRamp.TransactOpts, afn)
}

func (_OnRamp *OnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_OnRamp *OnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlist(&_OnRamp.TransactOpts, allowlist)
}

func (_OnRamp *OnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlist(&_OnRamp.TransactOpts, allowlist)
}

func (_OnRamp *OnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_OnRamp *OnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlistEnabled(&_OnRamp.TransactOpts, enabled)
}

func (_OnRamp *OnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _OnRamp.Contract.SetAllowlistEnabled(&_OnRamp.TransactOpts, enabled)
}

func (_OnRamp *OnRampTransactor) SetConfig(opts *bind.TransactOpts, config OnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setConfig", config)
}

func (_OnRamp *OnRampSession) SetConfig(config OnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetConfig(&_OnRamp.TransactOpts, config)
}

func (_OnRamp *OnRampTransactorSession) SetConfig(config OnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetConfig(&_OnRamp.TransactOpts, config)
}

func (_OnRamp *OnRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_OnRamp *OnRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.TransactOpts, newTime)
}

func (_OnRamp *OnRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_OnRamp.TransactOpts, newTime)
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

func (_OnRamp *OnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "unpause")
}

func (_OnRamp *OnRampSession) Unpause() (*types.Transaction, error) {
	return _OnRamp.Contract.Unpause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _OnRamp.Contract.Unpause(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_OnRamp *OnRampSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawAccumulatedFees(&_OnRamp.TransactOpts, feeToken, recipient, amount)
}

func (_OnRamp *OnRampTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawAccumulatedFees(&_OnRamp.TransactOpts, feeToken, recipient, amount)
}

type OnRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *OnRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(OnRampAFNMaxHeartbeatTimeSet)
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

func (it *OnRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OnRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAFNMaxHeartbeatTimeSetIterator{contract: _OnRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAFNMaxHeartbeatTimeSet)
				if err := _OnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OnRampAFNMaxHeartbeatTimeSet, error) {
	event := new(OnRampAFNMaxHeartbeatTimeSet)
	if err := _OnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAFNSetIterator struct {
	Event *OnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAFNSet)
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
		it.Event = new(OnRampAFNSet)
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

func (it *OnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_OnRamp *OnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*OnRampAFNSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAFNSetIterator{contract: _OnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAFNSet)
				if err := _OnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAFNSet(log types.Log) (*OnRampAFNSet, error) {
	event := new(OnRampAFNSet)
	if err := _OnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowlistEnabledSetIterator struct {
	Event *OnRampAllowlistEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowlistEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowlistEnabledSet)
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
		it.Event = new(OnRampAllowlistEnabledSet)
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

func (it *OnRampAllowlistEnabledSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowlistEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowlistEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowlistEnabledSet(opts *bind.FilterOpts) (*OnRampAllowlistEnabledSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowlistEnabledSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAllowlistEnabledSetIterator{contract: _OnRamp.contract, event: "AllowlistEnabledSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowlistEnabledSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistEnabledSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowlistEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowlistEnabledSet)
				if err := _OnRamp.contract.UnpackLog(event, "AllowlistEnabledSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowlistEnabledSet(log types.Log) (*OnRampAllowlistEnabledSet, error) {
	event := new(OnRampAllowlistEnabledSet)
	if err := _OnRamp.contract.UnpackLog(event, "AllowlistEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampAllowlistSetIterator struct {
	Event *OnRampAllowlistSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampAllowlistSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampAllowlistSet)
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
		it.Event = new(OnRampAllowlistSet)
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

func (it *OnRampAllowlistSetIterator) Error() error {
	return it.fail
}

func (it *OnRampAllowlistSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampAllowlistSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_OnRamp *OnRampFilterer) FilterAllowlistSet(opts *bind.FilterOpts) (*OnRampAllowlistSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "AllowlistSet")
	if err != nil {
		return nil, err
	}
	return &OnRampAllowlistSetIterator{contract: _OnRamp.contract, event: "AllowlistSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchAllowlistSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "AllowlistSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampAllowlistSet)
				if err := _OnRamp.contract.UnpackLog(event, "AllowlistSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseAllowlistSet(log types.Log) (*OnRampAllowlistSet, error) {
	event := new(OnRampAllowlistSet)
	if err := _OnRamp.contract.UnpackLog(event, "AllowlistSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampCrossChainSendRequestedIterator struct {
	Event *OnRampCrossChainSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampCrossChainSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampCrossChainSendRequested)
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
		it.Event = new(OnRampCrossChainSendRequested)
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

func (it *OnRampCrossChainSendRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampCrossChainSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampCrossChainSendRequested struct {
	Message CCIPMessage
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterCrossChainSendRequested(opts *bind.FilterOpts) (*OnRampCrossChainSendRequestedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CrossChainSendRequested")
	if err != nil {
		return nil, err
	}
	return &OnRampCrossChainSendRequestedIterator{contract: _OnRamp.contract, event: "CrossChainSendRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCrossChainSendRequested(opts *bind.WatchOpts, sink chan<- *OnRampCrossChainSendRequested) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CrossChainSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampCrossChainSendRequested)
				if err := _OnRamp.contract.UnpackLog(event, "CrossChainSendRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseCrossChainSendRequested(log types.Log) (*OnRampCrossChainSendRequested, error) {
	event := new(OnRampCrossChainSendRequested)
	if err := _OnRamp.contract.UnpackLog(event, "CrossChainSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeeChargedIterator struct {
	Event *OnRampFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeeCharged)
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
		it.Event = new(OnRampFeeCharged)
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

func (it *OnRampFeeChargedIterator) Error() error {
	return it.fail
}

func (it *OnRampFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*OnRampFeeChargedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &OnRampFeeChargedIterator{contract: _OnRamp.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *OnRampFeeCharged) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeeCharged)
				if err := _OnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeeCharged(log types.Log) (*OnRampFeeCharged, error) {
	event := new(OnRampFeeCharged)
	if err := _OnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeedAddedIterator struct {
	Event *OnRampFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeedAdded)
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
		it.Event = new(OnRampFeedAdded)
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

func (it *OnRampFeedAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*OnRampFeedAddedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &OnRampFeedAddedIterator{contract: _OnRamp.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OnRampFeedAdded) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeedAdded)
				if err := _OnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeedAdded(log types.Log) (*OnRampFeedAdded, error) {
	event := new(OnRampFeedAdded)
	if err := _OnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeedRemovedIterator struct {
	Event *OnRampFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeedRemoved)
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
		it.Event = new(OnRampFeedRemoved)
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

func (it *OnRampFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*OnRampFeedRemovedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &OnRampFeedRemovedIterator{contract: _OnRamp.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OnRampFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeedRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeedRemoved(log types.Log) (*OnRampFeedRemoved, error) {
	event := new(OnRampFeedRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeesWithdrawnIterator struct {
	Event *OnRampFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeesWithdrawn)
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
		it.Event = new(OnRampFeesWithdrawn)
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

func (it *OnRampFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*OnRampFeesWithdrawnIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OnRampFeesWithdrawnIterator{contract: _OnRamp.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeesWithdrawn)
				if err := _OnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeesWithdrawn(log types.Log) (*OnRampFeesWithdrawn, error) {
	event := new(OnRampFeesWithdrawn)
	if err := _OnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampNewTokenBucketConstructedIterator struct {
	Event *OnRampNewTokenBucketConstructed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampNewTokenBucketConstructedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampNewTokenBucketConstructed)
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
		it.Event = new(OnRampNewTokenBucketConstructed)
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

func (it *OnRampNewTokenBucketConstructedIterator) Error() error {
	return it.fail
}

func (it *OnRampNewTokenBucketConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampNewTokenBucketConstructed struct {
	Rate     *big.Int
	Capacity *big.Int
	Full     bool
	Raw      types.Log
}

func (_OnRamp *OnRampFilterer) FilterNewTokenBucketConstructed(opts *bind.FilterOpts) (*OnRampNewTokenBucketConstructedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "NewTokenBucketConstructed")
	if err != nil {
		return nil, err
	}
	return &OnRampNewTokenBucketConstructedIterator{contract: _OnRamp.contract, event: "NewTokenBucketConstructed", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchNewTokenBucketConstructed(opts *bind.WatchOpts, sink chan<- *OnRampNewTokenBucketConstructed) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "NewTokenBucketConstructed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampNewTokenBucketConstructed)
				if err := _OnRamp.contract.UnpackLog(event, "NewTokenBucketConstructed", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseNewTokenBucketConstructed(log types.Log) (*OnRampNewTokenBucketConstructed, error) {
	event := new(OnRampNewTokenBucketConstructed)
	if err := _OnRamp.contract.UnpackLog(event, "NewTokenBucketConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOnRampConfigSetIterator struct {
	Event *OnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOnRampConfigSet)
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
		it.Event = new(OnRampOnRampConfigSet)
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

func (it *OnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOnRampConfigSet struct {
	Config OnRampInterfaceOnRampConfig
	Raw    types.Log
}

func (_OnRamp *OnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*OnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampOnRampConfigSetIterator{contract: _OnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOnRampConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOnRampConfigSet(log types.Log) (*OnRampOnRampConfigSet, error) {
	event := new(OnRampOnRampConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

type OnRampPausedIterator struct {
	Event *OnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampPaused)
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
		it.Event = new(OnRampPaused)
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

func (it *OnRampPausedIterator) Error() error {
	return it.fail
}

func (it *OnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*OnRampPausedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OnRampPausedIterator{contract: _OnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OnRampPaused) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampPaused)
				if err := _OnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParsePaused(log types.Log) (*OnRampPaused, error) {
	event := new(OnRampPaused)
	if err := _OnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampPoolAddedIterator struct {
	Event *OnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampPoolAdded)
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
		it.Event = new(OnRampPoolAdded)
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

func (it *OnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *OnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*OnRampPoolAddedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &OnRampPoolAddedIterator{contract: _OnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampPoolAdded)
				if err := _OnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParsePoolAdded(log types.Log) (*OnRampPoolAdded, error) {
	event := new(OnRampPoolAdded)
	if err := _OnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampPoolRemovedIterator struct {
	Event *OnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampPoolRemoved)
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
		it.Event = new(OnRampPoolRemoved)
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

func (it *OnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *OnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_OnRamp *OnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*OnRampPoolRemovedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &OnRampPoolRemovedIterator{contract: _OnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampPoolRemoved)
				if err := _OnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParsePoolRemoved(log types.Log) (*OnRampPoolRemoved, error) {
	event := new(OnRampPoolRemoved)
	if err := _OnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterSetIterator struct {
	Event *OnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterSet)
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
		it.Event = new(OnRampRouterSet)
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

func (it *OnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_OnRamp *OnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*OnRampRouterSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &OnRampRouterSetIterator{contract: _OnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterSet)
				if err := _OnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseRouterSet(log types.Log) (*OnRampRouterSet, error) {
	event := new(OnRampRouterSet)
	if err := _OnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampUnpausedIterator struct {
	Event *OnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampUnpaused)
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
		it.Event = new(OnRampUnpaused)
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

func (it *OnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *OnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_OnRamp *OnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OnRampUnpausedIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OnRampUnpausedIterator{contract: _OnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampUnpaused)
				if err := _OnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseUnpaused(log types.Log) (*OnRampUnpaused, error) {
	event := new(OnRampUnpaused)
	if err := _OnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OnRamp *OnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _OnRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _OnRamp.abi.Events["AFNSet"].ID:
		return _OnRamp.ParseAFNSet(log)
	case _OnRamp.abi.Events["AllowlistEnabledSet"].ID:
		return _OnRamp.ParseAllowlistEnabledSet(log)
	case _OnRamp.abi.Events["AllowlistSet"].ID:
		return _OnRamp.ParseAllowlistSet(log)
	case _OnRamp.abi.Events["CrossChainSendRequested"].ID:
		return _OnRamp.ParseCrossChainSendRequested(log)
	case _OnRamp.abi.Events["FeeCharged"].ID:
		return _OnRamp.ParseFeeCharged(log)
	case _OnRamp.abi.Events["FeedAdded"].ID:
		return _OnRamp.ParseFeedAdded(log)
	case _OnRamp.abi.Events["FeedRemoved"].ID:
		return _OnRamp.ParseFeedRemoved(log)
	case _OnRamp.abi.Events["FeesWithdrawn"].ID:
		return _OnRamp.ParseFeesWithdrawn(log)
	case _OnRamp.abi.Events["NewTokenBucketConstructed"].ID:
		return _OnRamp.ParseNewTokenBucketConstructed(log)
	case _OnRamp.abi.Events["OnRampConfigSet"].ID:
		return _OnRamp.ParseOnRampConfigSet(log)
	case _OnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _OnRamp.ParseOwnershipTransferRequested(log)
	case _OnRamp.abi.Events["OwnershipTransferred"].ID:
		return _OnRamp.ParseOwnershipTransferred(log)
	case _OnRamp.abi.Events["Paused"].ID:
		return _OnRamp.ParsePaused(log)
	case _OnRamp.abi.Events["PoolAdded"].ID:
		return _OnRamp.ParsePoolAdded(log)
	case _OnRamp.abi.Events["PoolRemoved"].ID:
		return _OnRamp.ParsePoolRemoved(log)
	case _OnRamp.abi.Events["RouterSet"].ID:
		return _OnRamp.ParseRouterSet(log)
	case _OnRamp.abi.Events["Unpaused"].ID:
		return _OnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OnRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (OnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (OnRampAllowlistEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xa1bf86c493917580dec207969ef59976f0c378f10ece581237f19acfbd858f1c")
}

func (OnRampAllowlistSet) Topic() common.Hash {
	return common.HexToHash("0x27f242de1bc4ed72c4329591ffff7d223b5f025e3514a07e05afec6d4eb889cf")
}

func (OnRampCrossChainSendRequested) Topic() common.Hash {
	return common.HexToHash("0xbc2337a0942981c229faae38f5d31c14d85be0310ab538fe20e8c92b0585d101")
}

func (OnRampFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (OnRampFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (OnRampFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (OnRampFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (OnRampNewTokenBucketConstructed) Topic() common.Hash {
	return common.HexToHash("0xfaf3310019e551542b5c6014c1ae13e2a8d3943d7611d779c4df9b36c111924f")
}

func (OnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xeac62265bdcb30e1e7a4822fecd5035bf208f242c899453ca9a3cdb5eb44225b")
}

func (OnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (OnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (OnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (OnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (OnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_OnRamp *OnRamp) Address() common.Address {
	return _OnRamp.address
}

type OnRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (OnRampInterfaceOnRampConfig, error)

	GetDestinationChains(opts *bind.CallOpts) ([]*big.Int, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetSequenceNumberOfDestinationChain(opts *bind.CallOpts, destinationChainId *big.Int) (uint64, error)

	IsHealthy(opts *bind.CallOpts, now *big.Int) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload, originalSender common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config OnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*OnRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*OnRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*OnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *OnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*OnRampAFNSet, error)

	FilterAllowlistEnabledSet(opts *bind.FilterOpts) (*OnRampAllowlistEnabledSetIterator, error)

	WatchAllowlistEnabledSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistEnabledSet) (event.Subscription, error)

	ParseAllowlistEnabledSet(log types.Log) (*OnRampAllowlistEnabledSet, error)

	FilterAllowlistSet(opts *bind.FilterOpts) (*OnRampAllowlistSetIterator, error)

	WatchAllowlistSet(opts *bind.WatchOpts, sink chan<- *OnRampAllowlistSet) (event.Subscription, error)

	ParseAllowlistSet(log types.Log) (*OnRampAllowlistSet, error)

	FilterCrossChainSendRequested(opts *bind.FilterOpts) (*OnRampCrossChainSendRequestedIterator, error)

	WatchCrossChainSendRequested(opts *bind.WatchOpts, sink chan<- *OnRampCrossChainSendRequested) (event.Subscription, error)

	ParseCrossChainSendRequested(log types.Log) (*OnRampCrossChainSendRequested, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*OnRampFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *OnRampFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*OnRampFeeCharged, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*OnRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *OnRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*OnRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*OnRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *OnRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*OnRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*OnRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*OnRampFeesWithdrawn, error)

	FilterNewTokenBucketConstructed(opts *bind.FilterOpts) (*OnRampNewTokenBucketConstructedIterator, error)

	WatchNewTokenBucketConstructed(opts *bind.WatchOpts, sink chan<- *OnRampNewTokenBucketConstructed) (event.Subscription, error)

	ParseNewTokenBucketConstructed(log types.Log) (*OnRampNewTokenBucketConstructed, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*OnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*OnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*OnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *OnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*OnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*OnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *OnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*OnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*OnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *OnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*OnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*OnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*OnRampRouterSet, error)

	FilterUnpaused(opts *bind.FilterOpts) (*OnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*OnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
