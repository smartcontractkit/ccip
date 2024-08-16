package deployment

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	chainselectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type ChainConfig struct {
	ChainId uint64
	// TODO : use a slice of rpc urls for failing over to the available rpcs
	WsRpc       string
	HttpRpc     string
	DeployerKey *bind.TransactOpts
}

type RegistryConfig struct {
	EVMChainID uint64
	Contract   common.Address
}

func NewChainConfig(chainId uint64, wsRpc, httpRpc string, deployerKey *bind.TransactOpts) ChainConfig {
	return ChainConfig{
		ChainId:     chainId,
		WsRpc:       wsRpc,
		HttpRpc:     httpRpc,
		DeployerKey: deployerKey,
	}
}

func NewChains(logger logger.Logger, configs []ChainConfig) (map[uint64]Chain, error) {
	chains := make(map[uint64]Chain)
	for _, chainCfg := range configs {
		selector, err := chainselectors.SelectorFromChainId(chainCfg.ChainId)
		if err != nil {
			return nil, fmt.Errorf("failed to get selector from chain id %d: %w", chainCfg.ChainId, err)
		}
		// TODO : better client handling
		ec, err := ethclient.Dial(chainCfg.WsRpc)
		if err != nil {
			return nil, fmt.Errorf("failed to dial ws rpc %s: %w", chainCfg.WsRpc, err)
		}
		chains[selector] = Chain{
			Selector:    selector,
			Client:      ec,
			DeployerKey: chainCfg.DeployerKey,
			Confirm: func(tx common.Hash) error {
				return WaitForMined(logger, ec, tx, true)
			},
		}
	}
	return chains, nil
}
