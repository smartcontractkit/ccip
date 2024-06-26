package home_chain

import (
	"context"
	_ "embed"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ccipreader "github.com/smartcontractkit/ccipocr3/pkg/reader"
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

	simulatedBackend, auth := helpers.SetupBackendWithAuth(t)
	// Deploy the contract using the provided deployFunc
	address, tx, _, err := capcfg.DeployCCIPCapabilityConfiguration(auth, simulatedBackend, common.Address{})
	assert.NoError(t, err)
	simulatedBackend.Commit()
	t.Logf("contract deployed: addr=%s tx=%s", address.Hex(), tx.Hash())

	// Setup contract client using the provided newFunc
	contract, err := capcfg.NewCCIPCapabilityConfiguration(address, simulatedBackend)
	assert.NoError(t, err)

	chainReader := *helpers.SetupChainReader(t, simulatedBackend, address, cfg, ContractName)

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

	auth.GasLimit = 3518659
	_, err = contract.ApplyChainConfigUpdates(auth, nil, inputConfig)
	assert.NoError(t, err)
	simulatedBackend.Commit()

	var configs []capcfg.CCIPCapabilityConfigurationChainConfigInfo
	configs, err = contract.GetAllChainConfigs(nil)
	assert.NoError(t, err)
	assert.Equal(t, inputConfig, configs)

	// Now read the contract using chain reader into ccipreader type
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
				Readers: [][32]byte{
					//[32]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				FChain: 22,
				Config: []byte{1, 2, 3, 5, 6, 7},
			},
		},
	}
}
