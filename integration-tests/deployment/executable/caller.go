package executable

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/errors"
)

type Caller struct {
	Callers map[string]*gethwrappers.ManyChainMultiSig
	Clients map[string]ContractDeployBackend
}

func NewCaller(mcms map[string]common.Address, clients map[string]ContractDeployBackend) (*Caller, error) {
	mcmsWrappers := make(map[string]*gethwrappers.ManyChainMultiSig)

	for chain, mcmAddress := range mcms {
		client, ok := clients[chain]
		if !ok {
			return nil, &errors.ErrMissingChainClient{
				ChainIdentifier: chain,
			}
		}

		mcms, err := gethwrappers.NewManyChainMultiSig(mcmAddress, client)
		if err != nil {
			return nil, err
		}

		mcmsWrappers[chain] = mcms
	}

	return &Caller{
		Callers: mcmsWrappers,
		Clients: clients,
	}, nil
}

func (m *Caller) GetCurrentOpCounts() (map[string]big.Int, error) {
	opCounts := make(map[string]big.Int)

	for chain, wrapper := range m.Callers {
		opCount, err := wrapper.GetOpCount(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}

		opCounts[chain] = *opCount
	}

	return opCounts, nil
}

func (m *Caller) GetConfigs() (map[string]gethwrappers.ManyChainMultiSigConfig, error) {
	configs := make(map[string]gethwrappers.ManyChainMultiSigConfig)

	for chain, wrapper := range m.Callers {
		config, err := wrapper.GetConfig(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}

		configs[chain] = config
	}

	return configs, nil
}
