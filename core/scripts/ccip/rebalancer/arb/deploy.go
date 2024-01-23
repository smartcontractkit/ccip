package arb

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/bridgeutil"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l2_bridge_adapter"
	"go.uber.org/zap"
)

var (
	// Arbitrum Contracts
	// See https://docs.arbitrum.io/for-devs/useful-addresses
	ArbitrumContracts map[uint64]map[string]common.Address
)

func init() {
	ArbitrumContracts = map[uint64]map[string]common.Address{
		bridgeutil.SepoliaChainID: {
			"L1GatewayRouter": common.HexToAddress("0xcE18836b233C83325Cc8848CA4487e94C6288264"),
			"L1Outbox":        common.HexToAddress("0x65f07C7D521164a4d5DaC6eB8Fac8DA067A3B78F"),
			"Rollup":          common.HexToAddress("0xd80810638dbDF9081b72C1B33c65375e807281C8"),
			"WETH":            common.HexToAddress("0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9"),
		},
		bridgeutil.ArbitrumSepoliaChainID: {
			"L2GatewayRouter": common.HexToAddress("0x9fDD1C4E4AA24EEc1d913FABea925594a20d43C7"),
			"NodeInterface":   common.HexToAddress("0x00000000000000000000000000000000000000C8"),
			"WETH":            common.HexToAddress("0x980B62Da83eFf3D4576C647993b0c1D7faf17c73"),
		},
	}
}

func DeployAdapters(e multienv.Env) {
	DeployL1Adapter(e)
	DeployL2Adapter(e)
}

func DeployL1Adapter(e multienv.Env) {
	for _, chainID := range []uint64{bridgeutil.MainnetChainID, bridgeutil.SepoliaChainID} {
		if _, ok := e.Clients[chainID]; !ok {
			zap.L().Info("skipping L1 adapter deploy for chain id because client not found", zap.Uint64("chainID", chainID))
			continue
		}
		deployL1Adapter(e, chainID)
	}
}

func deployL1Adapter(e multienv.Env, chainID uint64) {
	zap.L().Info("deploying arbitrum L1 bridge adapter",
		zap.String("L1GatewayRouter", ArbitrumContracts[chainID]["L1GatewayRouter"].Hex()),
		zap.String("L1Outbox", ArbitrumContracts[chainID]["L1Outbox"].Hex()))
	l1BridgeAdapterAddress, tx, _, err := arbitrum_l1_bridge_adapter.DeployArbitrumL1BridgeAdapter(
		e.Transactors[chainID],
		e.Clients[chainID],
		ArbitrumContracts[chainID]["L1GatewayRouter"],
		ArbitrumContracts[chainID]["L1Outbox"],
	)
	helpers.PanicErr(err)
	helpers.ConfirmContractDeployed(context.Background(), e.Clients[chainID], tx, int64(chainID))
	zap.L().Sugar().Infof("L1 bridge adapter address: %s", l1BridgeAdapterAddress.Hex())
}

func DeployL2Adapter(e multienv.Env) {
	for _, chainID := range []uint64{
		bridgeutil.ArbitrumSepoliaChainID,
		bridgeutil.ArbitrumOneChainID,
		bridgeutil.ArbitrumNovaChainID,
	} {
		if _, ok := e.Clients[chainID]; !ok {
			zap.L().Sugar().Infoln("skipping arbitrum L2 adapter deploy for chain id", chainID, "because client not found")
			continue
		}
		deployL2Adapter(e, chainID)
	}
}

func deployL2Adapter(e multienv.Env, chainID uint64) {
	zap.L().Info("deploying arbitrum L2 bridge adapter",
		zap.String("L2GatewayRouter", ArbitrumContracts[chainID]["L2GatewayRouter"].Hex()))
	l2BridgeAdapterAddress, tx, _, err := arbitrum_l2_bridge_adapter.DeployArbitrumL2BridgeAdapter(
		e.Transactors[chainID],
		e.Clients[chainID],
		ArbitrumContracts[chainID]["L2GatewayRouter"],
	)
	helpers.PanicErr(err)
	helpers.ConfirmContractDeployed(context.Background(), e.Clients[chainID], tx, int64(chainID))
	zap.L().Sugar().Infoln("L2 bridge adapter address:", l2BridgeAdapterAddress.Hex())
}
