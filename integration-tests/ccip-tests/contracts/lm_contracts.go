package contracts

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/ccip/integration-tests/wrappers"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/arbitrum_l2_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/mock_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/mock_l2_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_l2_bridge_adapter"
)

type ArmProxy struct {
	client     blockchain.EVMClient
	Instance   *arm_proxy_contract.ARMProxyContract
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployArmProxy(arm common.Address) (*ArmProxy, error) {
	address, _, instance, err := e.evmClient.DeployContract("ARMProxy", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return arm_proxy_contract.DeployARMProxyContract(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			arm,
		)
	})
	if err != nil {
		return nil, err
	}
	return &ArmProxy{
		client:     e.evmClient,
		Instance:   instance.(*arm_proxy_contract.ARMProxyContract),
		EthAddress: address,
	}, err
}

func (pool *TokenPool) SetRouter(routerAddr common.Address) error {
	log.Info().
		Str("Token Pool", pool.Address()).
		Msg("Setting router on pool")
	opts, err := pool.client.TransactionOpts(pool.client.GetDefaultWallet())
	if err != nil {
		return fmt.Errorf("failed to get transaction opts: %w", err)
	}
	tx, err := pool.Instance.Latest.LockReleasePool.SetRouter(opts, routerAddr)
	if err != nil {
		return fmt.Errorf("failed to set router: %w", err)

	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Str("Router", routerAddr.String()).
		Msg("Router set on pool")
	return pool.client.ProcessTransaction(tx)
}

func (pool *TokenPool) GetRouter() (common.Address, error) {
	routerAddr, err := pool.Instance.Latest.LockReleasePool.GetRouter(nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get router address: %w", err)
	}
	return routerAddr, nil
}

func (pool *TokenPool) SetRebalancer(rebalancerAddress common.Address) error {
	log.Info().
		Str("Token Pool", pool.Address()).
		Msg("Setting rebalancer on pool")
	opts, err := pool.client.TransactionOpts(pool.client.GetDefaultWallet())
	if err != nil {
		return fmt.Errorf("failed to get transaction opts: %w", err)
	}
	tx, err := pool.Instance.Latest.LockReleasePool.SetRebalancer(opts, rebalancerAddress)
	if err != nil {
		return fmt.Errorf("failed to set router: %w", err)

	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Str("Rebalancer", rebalancerAddress.String()).
		Msg("Rebalancer set on pool")
	return pool.client.ProcessTransaction(tx)
}

func (pool *TokenPool) GetRebalancer() (common.Address, error) {
	rebalancerAddr, err := pool.Instance.Latest.LockReleasePool.GetRebalancer(nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get rebalancer address: %w", err)
	}
	return rebalancerAddr, nil
}

type LiquidityManager struct {
	client     blockchain.EVMClient
	Instance   *liquiditymanager.LiquidityManager
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployLiquidityManager(
	token common.Address,
	localChainSelector uint64,
	localLiquidityContainer common.Address,
	minimumLiquidity *big.Int,
) (*LiquidityManager, error) {
	address, _, instance, err := e.evmClient.DeployContract("LiquidityManager", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return liquiditymanager.DeployLiquidityManager(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			token,
			localChainSelector,
			localLiquidityContainer,
			minimumLiquidity,
		)
	})
	if err != nil {
		return nil, err
	}
	return &LiquidityManager{
		client:     e.evmClient,
		Instance:   instance.(*liquiditymanager.LiquidityManager),
		EthAddress: address,
	}, err
}

func (v *LiquidityManager) SetCrossChainRebalancer(
	crossChainRebalancerArgs liquiditymanager.ILiquidityManagerCrossChainRebalancerArgs,
) error {
	log.Info().
		Str("Liquidity Manager", v.EthAddress.String()).
		Msg("Setting crosschain rebalancer on liquidity manager")
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return fmt.Errorf("failed to get transaction opts: %w", err)
	}
	tx, err := v.Instance.SetCrossChainRebalancer(opts, crossChainRebalancerArgs)
	if err != nil {
		return fmt.Errorf("failed to set cross chain rebalancer: %w", err)

	}
	log.Info().
		Str("Liquidity Manager", v.EthAddress.String()).
		Interface("Rebalance Argsr", crossChainRebalancerArgs).
		Msg("Crosschain Rebalancer set on liquidity manager")
	return v.client.ProcessTransaction(tx)
}

func (v *LiquidityManager) SetOCR3Config(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {

	log.Info().
		Str("Liquidity Manager", v.EthAddress.String()).
		Msg("Setting ocr3 config on liquidity manager")
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return fmt.Errorf("failed to get transaction opts: %w", err)
	}
	tx, err := v.Instance.SetOCR3Config(
		opts,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig)
	if err != nil {
		return fmt.Errorf("failed to set cross chain rebalancer: %w", err)

	}
	log.Info().
		Str("Liquidity Manager", v.EthAddress.String()).
		Msg("Set OCR3Config on LM")
	return v.client.ProcessTransaction(tx)
}

type ArbitrumL1BridgeAdapter struct {
	client     blockchain.EVMClient
	Instance   *arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapter
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployArbitrumL1BridgeAdapter(
	l1GatewayRouter common.Address,
	l1Outbox common.Address,
) (*ArbitrumL1BridgeAdapter, error) {
	address, _, instance, err := e.evmClient.DeployContract("ArbitrumL1BridgeAdapter", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return arbitrum_l1_bridge_adapter.DeployArbitrumL1BridgeAdapter(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			l1GatewayRouter,
			l1Outbox,
		)
	})
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapter{
		client:     e.evmClient,
		Instance:   instance.(*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapter),
		EthAddress: address,
	}, err
}

type ArbitrumL2BridgeAdapter struct {
	client     blockchain.EVMClient
	Instance   *arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapter
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployArbitrumL2BridgeAdapter(l2GatewayRouter common.Address) (*ArbitrumL2BridgeAdapter, error) {
	address, _, instance, err := e.evmClient.DeployContract("ArbitrumL2BridgeAdapter", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return arbitrum_l2_bridge_adapter.DeployArbitrumL2BridgeAdapter(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			l2GatewayRouter,
		)
	})
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapter{
		client:     e.evmClient,
		Instance:   instance.(*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapter),
		EthAddress: address,
	}, err
}

type OptimismL1BridgeAdapter struct {
	client     blockchain.EVMClient
	Instance   *optimism_l1_bridge_adapter.OptimismL1BridgeAdapter
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployOptimismL1BridgeAdapter(
	l1Bridge common.Address,
	wrappedNative common.Address,
	optimismPortal common.Address,
) (*OptimismL1BridgeAdapter, error) {
	address, _, instance, err := e.evmClient.DeployContract("OptimismL1BridgeAdapter", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return optimism_l1_bridge_adapter.DeployOptimismL1BridgeAdapter(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			l1Bridge,
			wrappedNative,
			optimismPortal,
		)
	})
	if err != nil {
		return nil, err
	}
	return &OptimismL1BridgeAdapter{
		client:     e.evmClient,
		Instance:   instance.(*optimism_l1_bridge_adapter.OptimismL1BridgeAdapter),
		EthAddress: address,
	}, err
}

type OptimismL2BridgeAdapter struct {
	client     blockchain.EVMClient
	Instance   *optimism_l2_bridge_adapter.OptimismL2BridgeAdapter
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployOptimismL2BridgeAdapter(l2GatewayRouter common.Address) (*OptimismL2BridgeAdapter, error) {
	address, _, instance, err := e.evmClient.DeployContract("OptimismL2BridgeAdapter", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return optimism_l2_bridge_adapter.DeployOptimismL2BridgeAdapter(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			l2GatewayRouter,
		)
	})
	if err != nil {
		return nil, err
	}
	return &OptimismL2BridgeAdapter{
		client:     e.evmClient,
		Instance:   instance.(*optimism_l2_bridge_adapter.OptimismL2BridgeAdapter),
		EthAddress: address,
	}, err
}

type MockL1BridgeAdapter struct {
	client     blockchain.EVMClient
	Instance   *mock_l1_bridge_adapter.MockL1BridgeAdapter
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployMockL1BridgeAdapter(tokenAddr common.Address, holdNative bool) (*MockL1BridgeAdapter, error) {
	address, _, instance, err := e.evmClient.DeployContract("MockL1BridgeAdapter", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return mock_l1_bridge_adapter.DeployMockL1BridgeAdapter(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
			tokenAddr,
			holdNative,
		)
	})
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapter{
		client:     e.evmClient,
		Instance:   instance.(*mock_l1_bridge_adapter.MockL1BridgeAdapter),
		EthAddress: address,
	}, err
}

type MockL2BridgeAdapter struct {
	client     blockchain.EVMClient
	Instance   *mock_l2_bridge_adapter.MockL2BridgeAdapter
	EthAddress *common.Address
}

func (e *CCIPContractsDeployer) DeployMockL2BridgeAdapter() (*MockL2BridgeAdapter, error) {
	address, _, instance, err := e.evmClient.DeployContract("MockL2BridgeAdapter", func(
		auth *bind.TransactOpts,
		_ bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return mock_l2_bridge_adapter.DeployMockL2BridgeAdapter(
			auth,
			wrappers.MustNewWrappedContractBackend(e.evmClient, nil),
		)
	})
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapter{
		client:     e.evmClient,
		Instance:   instance.(*mock_l2_bridge_adapter.MockL2BridgeAdapter),
		EthAddress: address,
	}, err
}
