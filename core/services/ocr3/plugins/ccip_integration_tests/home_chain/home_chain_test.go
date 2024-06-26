package home_chain

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ccipreader "github.com/smartcontractkit/ccipocr3/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	capcfg "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	logger2 "github.com/smartcontractkit/chainlink/v2/core/logger"
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
	//===============================================================
	transactor := testutils.MustNewSimTransactor(t)
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		transactor.From: {Balance: assets.Ether(1000).ToInt()},
	}, 30e6)
	//==============================Setup Contracts - with capabilities=================================
	capRegAddress, capRegContract, err := prepareCapabilityRegistry(t, backend, transactor)
	assert.NoError(t, err)
	capConfAddress, capConfContract, err := prepareCCIPCapabilityConfig(t, backend, transactor, capRegAddress)
	assert.NoError(t, err)
	p2pIDS := addCapabilities(t, backend, transactor, capRegContract, capConfAddress)
	//==============================Apply configs to Capability Contract=================================
	chainAConf := setupConfigInfo(chainA, p2pIDS, fChainA, []byte{})
	chainBConf := setupConfigInfo(chainB, p2pIDS[1:], fChainB, []byte{})
	chainCConf := setupConfigInfo(chainC, p2pIDS[2:], fChainC, []byte{})

	inputConfig := []capcfg.CCIPCapabilityConfigurationChainConfigInfo{
		chainAConf,
		chainBConf,
		chainCConf,
	}
	_, err = capConfContract.ApplyChainConfigUpdates(transactor, nil, inputConfig)
	assert.NoError(t, err)
	backend.Commit()
	//================================Setup HomeChainReader===============================
	chainReader := *helpers.SetupChainReader(t, backend, capConfAddress, cfg, ContractName)
	var ccipConfigResults []ccipreader.CCIPCapabilityConfigurationChainConfigInfo
	err = chainReader.GetLatestValue(
		context.Background(),
		ContractName,
		FnGetChainConfigs,
		map[string]interface{}{},
		&ccipConfigResults,
	)
	assert.NoError(t, err)
	homeChain := ccipreader.NewHomeChainReader(chainReader, logger2.NullLogger, 1*time.Second)
	err = homeChain.Start(context.Background())
	assert.NoError(t, err)
	//===============================================================
	expectedChainConfigs := map[cciptypes.ChainSelector]ccipreader.CCIPCapabilityConfigurationChainConfig{}
	for _, c := range inputConfig {
		expectedChainConfigs[cciptypes.ChainSelector(c.ChainSelector)] = ccipreader.CCIPCapabilityConfigurationChainConfig{
			FChain:         int(c.ChainConfig.FChain),
			SupportedNodes: toPeerIDs(c.ChainConfig.Readers),
		}
	}
	configs, err := homeChain.GetAllChainConfigs()
	assert.NoError(t, err)
	assert.Equal(t, expectedChainConfigs, configs)
}

func toPeerIDs(readers [][32]byte) mapset.Set[libocrtypes.PeerID] {
	peerIDs := mapset.NewSet[libocrtypes.PeerID]()
	for _, r := range readers {
		peerIDs.Add(r)
	}
	return peerIDs
}
func setupConfigInfo(chainSelector uint64, readers [][32]byte, fChain uint8, cfg []byte) capcfg.CCIPCapabilityConfigurationChainConfigInfo {
	return capcfg.CCIPCapabilityConfigurationChainConfigInfo{
		ChainSelector: chainSelector,
		ChainConfig: capcfg.CCIPCapabilityConfigurationChainConfig{
			Readers: readers,
			FChain:  fChain,
			Config:  cfg,
		},
	}
}
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

func addCapabilities(
	t *testing.T,
	backend *backends.SimulatedBackend,
	transactor *bind.TransactOpts,
	capReg *capabilities_registry.CapabilitiesRegistry,
	capConfAddress common.Address) [][32]byte {
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
	return p2pIDs
}
