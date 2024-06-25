package home_chain

import (
	"context"
	_ "embed"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	capcfg "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	helpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip_integration_tests"
	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/stretchr/testify/assert"
)

var (
	chainA       uint64 = 1
	nodeAID      uint8  = 1
	fChainA      uint8  = 1
	oracleAID           = commontypes.OracleID(nodeAID)
	peerAID             = "12D3KooWPjceQrSwdWXPyLLeABRXmuqt69Rg3sBYbU1Nft9HyQ6X"
	p2pOracleAID        = libocrtypes.PeerID{byte(nodeAID)}

	chainB       uint64 = 2
	nodeBID      uint8  = 2
	fChainB      uint8  = 2
	oracleBID           = commontypes.OracleID(nodeBID)
	p2pOracleBID        = libocrtypes.PeerID{byte(nodeBID)}

	chainC       uint64 = 3
	nodeCID      uint8  = 3
	fChainC      uint8  = 3
	oracleCID           = commontypes.OracleID(nodeCID)
	p2pOracleCID        = libocrtypes.PeerID{byte(nodeCID)}
)

func TestHomeChainReader(t *testing.T) {
	deployFunc := func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *capcfg.CCIPCapabilityConfiguration, error) {
		return capcfg.DeployCCIPCapabilityConfiguration(auth, backend, common.Address{})
	}
	const (
		ContractName      = "CCIPCapabilityConfiguration"
		FnGetChainConfigs = "getAllChainConfigs"
	)
	// Initialize chainReader
	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			ContractName: {
				ContractABI: capcfg.CCIPCapabilityConfigurationMetaData.ABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					"ChainConfigSet": {
						ChainSpecificName:       "ChainConfigSet",
						ReadType:                evmtypes.Event,
						ConfidenceConfirmations: map[string]int{"0.0": 0, "1.0": 0},
					},
					FnGetChainConfigs: {
						ChainSpecificName: FnGetChainConfigs,
					},
				},
			},
		},
	}

	d := helpers.SetupChainReaderTest[capcfg.CCIPCapabilityConfiguration](t, context.Background(), deployFunc, capcfg.NewCCIPCapabilityConfiguration, cfg, ContractName)
	chainReader := *d.ChainReader

	// Apply chain configs to the contract
	inputConfig := setupConfigInfo()
	//[]capcfg.CCIPCapabilityConfigurationChainConfigInfo{
	//	setupConfigInfo(chainA, []uint8{nodeAID, nodeBID, nodeCID}, fChainA, []byte{}),
	//	setupConfigInfo(chainB, []uint8{nodeAID, nodeBID}, fChainB, []byte{}),
	//	setupConfigInfo(chainC, []uint8{nodeCID}, fChainC, []byte{}),
	//setupConfigInfo(chainA, []byte{'a', 'b', 'c'}, fChainA, []byte{}),
	//setupConfigInfo(chainB, []byte{'a', 'b'}, fChainB, []byte{}),
	//setupConfigInfo(chainC, []byte{}, fChainC, []byte{}),
	//}

	//d.Auth.GasLimit = 50000
	_, err := d.Contract.ApplyChainConfigUpdates(d.Auth, nil, inputConfig)
	d.SimulatedBE.Commit()
	assert.NoError(t, err)

	// Now read the contract using chain reader
	var allConfigs []capcfg.CCIPCapabilityConfigurationChainConfigInfo
	err = chainReader.GetLatestValue(
		context.Background(),
		ContractName,
		FnGetChainConfigs,
		map[string]interface{}{},
		&allConfigs,
	)
	assert.NoError(t, err)
	assert.Equal(t, inputConfig, allConfigs)
}

func setupConfigInfo() []capcfg.CCIPCapabilityConfigurationChainConfigInfo {
	return []capcfg.CCIPCapabilityConfigurationChainConfigInfo{
		{
			ChainSelector: 1,
			ChainConfig: capcfg.CCIPCapabilityConfigurationChainConfig{
				//Readers: [][32]byte{{1}, {2}, {3}},
				Readers: [][32]byte{},
				FChain:  2,
				Config:  []byte{1, 2, 3},
			},
		},
	}
}
