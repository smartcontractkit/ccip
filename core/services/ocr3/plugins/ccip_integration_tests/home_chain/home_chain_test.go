package home_chain

import (
	"context"
	_ "embed"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	capcfg "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	helpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip_integration_tests"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	logger2 "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/stretchr/testify/assert"
)

const (
	chainID = 1337
	chainA  = ccipocr3.ChainSelector(1)
	chainB  = ccipocr3.ChainSelector(2)
	chainC  = ccipocr3.ChainSelector(3)
)

func TestHomeChainReader(t *testing.T) {
	// Generate a new key pair for the simulated account
	//privateKey, err := crypto.GenerateKey()
	//assert.NoError(t, err)
	//// Set up the genesis account with balance
	//blnc, ok := big.NewInt(0).SetString("999999999999999999999999999999999999", 10)
	//assert.True(t, ok)
	//alloc := map[common.Address]core.GenesisAccount{crypto.PubkeyToAddress(privateKey.PublicKey): {Balance: blnc}}
	//simulatedBackend := backends.NewSimulatedBackend(alloc, 0)
	//// Create a transactor
	//
	//auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	//assert.NoError(t, err)
	//auth.GasLimit = uint64(0)
	//
	//// Deploy the contract
	//address, tx, _, err := capcfg.DeployCCIPCapabilityConfiguration(auth, simulatedBackend, common.Address{})
	//assert.NoError(t, err)
	//simulatedBackend.Commit()
	//t.Logf("contract deployed: addr=%s tx=%s", address.Hex(), tx.Hash())
	//
	//// Setup contract client
	//contract, err := capcfg.NewCCIPCapabilityConfiguration(address, simulatedBackend)
	//assert.NoError(t, err)

	d := helpers.SetupTestWithCapability[capcfg.CCIPCapabilityConfiguration](t, context.Background(), capcfg.DeployCCIPCapabilityConfiguration, capcfg.NewCCIPCapabilityConfiguration, common.Address{})
	// Initialize chainReader
	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			"CCIPCapabilityConfiguration": {
				ContractABI: capcfg.CCIPCapabilityConfigurationMetaData.ABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					"ChainConfigSet": {
						ChainSpecificName:       "ChainConfigSet",
						ReadType:                evmtypes.Event,
						ConfidenceConfirmations: map[string]int{"0.0": 0, "1.0": 0},
					},
					"getAllChainConfigs": {
						ChainSpecificName: "getAllChainConfigs",
					},
				},
			},
		},
	}

	lggr := logger2.NullLogger
	db := pgtest.NewSqlxDB(t)
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Millisecond,
		FinalityDepth:            1,
		BackfillBatchSize:        1,
		RpcBatchSize:             1,
		KeepFinalizedBlocksDepth: 10000,
	}
	cl := client.NewSimulatedBackendClient(t, d.SimulatedBE, big.NewInt(chainID))
	lp := logpoller.NewLogPoller(logpoller.NewORM(big.NewInt(chainID), db, lggr), cl, lggr, lpOpts)
	assert.NoError(t, lp.Start(context.Background()))

	cr, err := evm.NewChainReaderService(context.Background(), lggr, lp, cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(context.Background(), []types2.BoundContract{
		{
			Address: d.ContractAddr.String(),
			Name:    "CCIPCapabilityConfiguration",
			Pending: false,
		},
	})
	assert.NoError(t, err)

	err = cr.Start(context.Background())
	assert.NoError(t, err)
	for {
		if err := cr.Ready(); err == nil {
			break
		}
	}

	d.SimulatedBE.Commit()

	// Apply chain configs to the contract
	inputConfig := setupConfigInfo()
	_, err = d.Contract.ApplyChainConfigUpdates(d.Auth, nil, inputConfig)
	d.SimulatedBE.Commit()
	assert.NoError(t, err)

	// Now read the contract using chain reader
	var allConfigs []capcfg.CCIPCapabilityConfigurationChainConfigInfo
	err = cr.GetLatestValue(
		context.Background(),
		"CCIPCapabilityConfiguration",
		"getAllChainConfigs",
		map[string]interface{}{},
		&allConfigs,
	)
	assert.NoError(t, err)
	assert.Equal(t, inputConfig, allConfigs)
}

func setupConfigInfo() []capcfg.CCIPCapabilityConfigurationChainConfigInfo {
	return []capcfg.CCIPCapabilityConfigurationChainConfigInfo{
		{
			ChainSelector: chainID,
			ChainConfig: capcfg.CCIPCapabilityConfigurationChainConfig{
				Readers: [][32]byte{},
				FChain:  2,
				Config:  []byte{1, 2, 3},
			},
		},
	}
}
