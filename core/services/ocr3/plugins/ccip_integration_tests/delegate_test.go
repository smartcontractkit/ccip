package ccip_integration_tests

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	"github.com/stretchr/testify/require"
)

func TestDelegate_ChainReaderBinding(t *testing.T) {
	transactor := testutils.MustNewSimTransactor(t)
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		transactor.From: {Balance: assets.Ether(1000).ToInt()},
	}, 30e6)

	crAddress, _, _, err := capabilities_registry.DeployCapabilitiesRegistry(transactor, backend)
	require.NoError(t, err)
	backend.Commit()

	capReg, err := capabilities_registry.NewCapabilitiesRegistry(crAddress, backend)
	require.NoError(t, err)

	ccAddress, _, _, err := ccip_capability_configuration.DeployCCIPCapabilityConfiguration(transactor, backend, crAddress)
	require.NoError(t, err)
	backend.Commit()

	capConfig, err := ccip_capability_configuration.NewCCIPCapabilityConfiguration(ccAddress, backend)
	require.NoError(t, err)

	// add the CCIP capability to the registry
	_, err = capReg.AddCapabilities(transactor, []capabilities_registry.CapabilitiesRegistryCapability{
		{
			LabelledName:          "ccip",
			Version:               "v1.0",
			CapabilityType:        0,
			ResponseType:          0,
			ConfigurationContract: ccAddress,
		},
	})
	require.NoError(t, err, "failed to add capability to registry")
	backend.Commit()

	ccipCapabilityID, err := capReg.GetHashedCapabilityId(nil, "ccip", "v1.0")
	require.NoError(t, err)

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
		require.NoError(t, err)
		backend.Commit()

		// get the node operator id from the event
		it, err := capReg.FilterNodeOperatorAdded(nil, nil, nil)
		require.NoError(t, err)
		var nodeOperatorID uint32
		for it.Next() {
			if it.Event.Name == fmt.Sprintf("nop-%d", i) {
				nodeOperatorID = it.Event.NodeOperatorId
				break
			}
		}
		require.NotZero(t, nodeOperatorID)

		_, err = capReg.AddNodes(transactor, []capabilities_registry.CapabilitiesRegistryNodeParams{
			{
				NodeOperatorId:      nodeOperatorID,
				Signer:              testutils.Random32Byte(),
				P2pId:               p2pID,
				HashedCapabilityIds: [][32]byte{ccipCapabilityID},
			},
		})
		require.NoError(t, err)
		backend.Commit()

		// verify that the node was added successfully
		nodeInfo, err := capReg.GetNode(nil, p2pID)
		require.NoError(t, err)

		require.Equal(t, nodeOperatorID, nodeInfo.NodeOperatorId)
		require.Equal(t, p2pID[:], nodeInfo.P2pId[:])
	}

	owner, err := capConfig.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, transactor.From, owner)

	// add some chain configs to the ccip cap config contract
	_, err = capConfig.ApplyChainConfigUpdates(
		transactor,
		[]uint64{},
		[]ccip_capability_configuration.CCIPCapabilityConfigurationChainConfigInfo{
			{
				ChainSelector: 1,
				ChainConfig: ccip_capability_configuration.CCIPCapabilityConfigurationChainConfig{
					Readers: p2pIDs,
					FChain:  1,
					Config:  []byte("0xdeadbeef"),
				},
			},
		})
	require.NoError(t, err, "unable to apply chain config updates")
	backend.Commit()

	// spin up the chain reader deps
	lggr := logger.TestLogger(t)
	db := pgtest.NewSqlxDB(t)
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Millisecond,
		FinalityDepth:            0,
		BackfillBatchSize:        10,
		RpcBatchSize:             10,
		KeepFinalizedBlocksDepth: 100000,
	}
	sbc := client.NewSimulatedBackendClient(t, backend, big.NewInt(1337))
	lp := logpoller.NewLogPoller(logpoller.NewORM(big.NewInt(1337), db, lggr), sbc, lggr, lpOpts)
	require.NoError(t, lp.Start(testutils.Context(t)))

	//chainReaderConfig := homeChainReaderConfig()
	//cr, err := evm.NewChainReaderService(testutils.Context(t), lggr, lp, sbc, chainReaderConfig)
	//require.NoError(t, err)
	//
	//cr, err = bindReader(cr, crAddress.Hex(), "ccip", "v1.0")
	//require.NoError(t, err)
	//
	//require.NoError(t, cr.Start(testutils.Context(t)))
	//
	//var allChainConfigs []ccip_capability_configuration.CCIPCapabilityConfigurationChainConfigInfo
	//err = cr.GetLatestValue(testutils.Context(t), "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &allChainConfigs)
	//require.NoError(t, err, "failed to get all chain configs using chain reader")
	//require.Len(t, allChainConfigs, 1)
	//require.Equal(t, uint64(1), allChainConfigs[0].ChainSelector)
	//require.Equal(t, []byte("0xdeadbeef"), allChainConfigs[0].ChainConfig.Config)
	//require.Equal(t, p2pIDs, allChainConfigs[0].ChainConfig.Readers)
}
