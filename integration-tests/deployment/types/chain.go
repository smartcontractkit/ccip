package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainselectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/ccip/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type ChainConfig struct {
	ChainId     int64
	WsRpc       string
	HttpRpc     string
	DeployerKey *bind.TransactOpts
}

type RegistryConfig struct {
	EVMChainID uint64
	Contract   common.Address
}

func NewChainConfig(chainId int64, wsRpc, httpRpc string, deployerKey *bind.TransactOpts) ChainConfig {
	return ChainConfig{
		ChainId:     chainId,
		WsRpc:       wsRpc,
		HttpRpc:     httpRpc,
		DeployerKey: deployerKey,
	}
}

func NewChains(logger logger.Logger, configs []ChainConfig) map[uint64]deployment.Chain {
	selector, err := chainselectors.SelectorFromChainId()
	if err != nil {
		return deployment.Chain{}
	}

}
