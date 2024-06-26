package home_chain

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ccipreader "github.com/smartcontractkit/ccipocr3/pkg/reader"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	capcfg "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
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
					FnGetChainConfigs: {
						ChainSpecificName: FnGetChainConfigs,
					},
				},
			},
		},
	}

	transactor := testutils.MustNewSimTransactor(t)
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		transactor.From: {Balance: assets.Ether(1000).ToInt()},
	}, 30e6)
	//backend, transactor := helpers.SetupBackendWithAuth(t)

	capRegAddress, capRegContract, err := prepareCapabilityRegistry(t, backend, transactor)
	capConfAddress, capConfContract, err := prepareCCIPCapabilityConfig(t, backend, transactor, capRegAddress)
	addCapabilities(t, backend, transactor, capRegContract, capConfAddress)
	chainReader := *helpers.SetupChainReader(t, backend, capConfAddress, cfg, ContractName)

	// Apply chain configs to the capConfContract
	inputConfig := setupConfigInfo()
	//[]capcfg.CCIPCapabilityConfigurationChainConfigInfo{
	//	setupConfigInfo(chainA, []uint8{nodeAID, nodeBID, nodeCID}, fChainA, []byte{}),
	//	setupConfigInfo(chainB, []uint8{nodeAID, nodeBID}, fChainB, []byte{}),
	//	setupConfigInfo(chainC, []uint8{nodeCID}, fChainC, []byte{}),
	//}

	_, err = capConfContract.ApplyChainConfigUpdates(transactor, nil, inputConfig)
	assert.NoError(t, err)
	backend.Commit()

	var configs []capcfg.CCIPCapabilityConfigurationChainConfigInfo
	configs, err = capConfContract.GetAllChainConfigs(nil)
	assert.NoError(t, err)
	assert.Equal(t, inputConfig, configs)

	// Now read the capConfContract using chain reader into ccipreader type
	var ccipConfigResults []ccipreader.CCIPCapabilityConfigurationChainConfigInfo
	err = chainReader.GetLatestValue(
		context.Background(),
		ContractName,
		FnGetChainConfigs,
		map[string]interface{}{},
		&ccipConfigResults,
	)
	assert.NoError(t, err)
	for i, c := range ccipConfigResults {
		assert.Equal(t, inputConfig[i].ChainSelector, uint64(c.ChainSelector))
		assert.Equal(t, inputConfig[i].ChainConfig.Config, c.ChainConfig.Config)
		assert.Equal(t, inputConfig[i].ChainConfig.FChain, c.ChainConfig.FChain)
		for j, r := range inputConfig[i].ChainConfig.Readers {
			assert.Equal(t, r, c.ChainConfig.Readers[j])
		}

	}
	//homeChain := ccipreader.NewHomeChainReader(chainReader, logger2.NullLogger, 1*time.Second)
}

//func randomBytes(n int) []byte {
//	b := make([]byte, n)
//	_, err := rand.Read(b)
//	if err != nil {
//		panic(err)
//	}
//	return b
//}
//
//// Random32Byte returns a random [32]byte
//func Random32Byte() (b [32]byte) {
//	copy(b[:], randomBytes(32))
//	return b
//}

func prepareCCIPCapabilityConfig(t *testing.T, backend *backends.SimulatedBackend, transactor *bind.TransactOpts, capRegAddress common.Address) (common.Address, *capcfg.CCIPCapabilityConfiguration, error) {
	ccAddress, _, _, err := capcfg.DeployCCIPCapabilityConfiguration(transactor, backend, capRegAddress)
	assert.NoError(t, err)
	backend.Commit()

	contract, err := capcfg.NewCCIPCapabilityConfiguration(ccAddress, backend)
	assert.NoError(t, err)
	backend.Commit()

	return ccAddress, contract, nil
}
func prepareCapabilityRegistry(t *testing.T, backend *backends.SimulatedBackend, transactor *bind.TransactOpts) (common.Address, *capabilities_registry.CapabilitiesRegistry, error) {
	crAddress, _, _, err := capabilities_registry.DeployCapabilitiesRegistry(transactor, backend)
	assert.NoError(t, err)
	backend.Commit()

	capReg, err := capabilities_registry.NewCapabilitiesRegistry(crAddress, backend)
	assert.NoError(t, err)
	backend.Commit()

	return crAddress, capReg, nil
}

func addCapabilities(t *testing.T, backend *backends.SimulatedBackend, transactor *bind.TransactOpts, capReg *capabilities_registry.CapabilitiesRegistry, capConfAddress common.Address) {
	// add the CCIP capability to the registry
	_, err := capReg.AddCapabilities(transactor, []capabilities_registry.CapabilitiesRegistryCapability{
		{
			LabelledName:          "ccip",
			Version:               "v1.0",
			CapabilityType:        0,
			ResponseType:          0,
			ConfigurationContract: capConfAddress,
		},
	})
	assert.NoError(t, err, "failed to add capability to registry")
	backend.Commit()

	ccipCapabilityID, err := capReg.GetHashedCapabilityId(nil, "ccip", "v1.0")
	assert.NoError(t, err)

	// Add the p2p ids of the ccip nodes
	var p2pIDs [][32]byte
	for i := 0; i < 4; i++ {
		p2pID := p2pkey.MustNewV2XXXTestingOnly(big.NewInt(int64(i + 1))).PeerID()
		p2pIDs = append(p2pIDs, p2pID)
		_, err = capReg.AddNodeOperators(transactor, []capabilities_registry.CapabilitiesRegistryNodeOperator{
			{
				Admin: transactor.From,
				Name:  fmt.Sprintf("nop-%d", i),
			},
		})
		assert.NoError(t, err)
		backend.Commit()

		// get the node operator id from the event
		it, err := capReg.FilterNodeOperatorAdded(nil, nil, nil)
		assert.NoError(t, err)
		var nodeOperatorID uint32
		for it.Next() {
			if it.Event.Name == fmt.Sprintf("nop-%d", i) {
				nodeOperatorID = it.Event.NodeOperatorId
				break
			}
		}
		assert.NotZero(t, nodeOperatorID)

		_, err = capReg.AddNodes(transactor, []capabilities_registry.CapabilitiesRegistryNodeParams{
			{
				NodeOperatorId:      nodeOperatorID,
				Signer:              testutils.Random32Byte(),
				P2pId:               p2pID,
				HashedCapabilityIds: [][32]byte{ccipCapabilityID},
			},
		})
		assert.NoError(t, err)
		backend.Commit()

		// verify that the node was added successfully
		nodeInfo, err := capReg.GetNode(nil, p2pID)
		assert.NoError(t, err)

		assert.Equal(t, nodeOperatorID, nodeInfo.NodeOperatorId)
		assert.Equal(t, p2pID[:], nodeInfo.P2pId[:])
	}
}

func setupConfigInfo() []capcfg.CCIPCapabilityConfigurationChainConfigInfo {
	return []capcfg.CCIPCapabilityConfigurationChainConfigInfo{
		{
			ChainSelector: 1,
			ChainConfig: capcfg.CCIPCapabilityConfigurationChainConfig{
				//Readers: [][32]byte{[32]byte{1}},
				//Readers: [][32]byte{Random32Byte()},
				Readers: [][32]byte{},
				FChain:  11,
				Config:  []byte{1, 2, 3, 5, 6, 7},
			},
		},
		{
			ChainSelector: 2,
			ChainConfig: capcfg.CCIPCapabilityConfigurationChainConfig{
				Readers: [][32]byte{},
				FChain:  22,
				Config:  []byte{1, 2, 3, 5, 6, 7},
			},
		},
	}
}
