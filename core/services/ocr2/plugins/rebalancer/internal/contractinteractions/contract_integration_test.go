package contractinteractions_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/testonly_liquidity_manager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/test-go/testify/require"
)

var (
	testOnlyLMABI = evmtypes.MustGetABI(testonly_liquidity_manager.TestOnlyLiquidityManagerABI)
)

const (
	l1ChainID uint64 = 1337
	l2ChainID uint64 = 1338
)

type testUniverse struct {
	backend *backends.SimulatedBackend
	chainID uint64

	wethAddress common.Address
	wethToken   *weth9.WETH9

	armAddress  common.Address
	armContract *mock_arm_contract.MockARMContract

	armProxyAddress common.Address
	armProxy        *arm_proxy_contract.ARMProxyContract

	lockReleasePoolAddr common.Address
	lockReleasePool     *lock_release_token_pool.LockReleaseTokenPool

	lmAddress common.Address
	lm        *testonly_liquidity_manager.TestOnlyLiquidityManager

	bridgeAdapterAddress common.Address
	bridgeAdapter        *mock_l1_bridge_adapter.MockL1BridgeAdapter
}

func Test_LiquidityContractInteractions(t *testing.T) {
	owner := testutils.MustNewSimTransactor(t)

	const (
		ownerStartBalance    int64 = 100_000
		bridgeAdapterBalance int64 = 5
		l1TokenPoolBalance   int64 = 5
	)

	// deploy contracts on both chains
	l1Contracts := deployRebalancerContractUniverse(t, owner, l1ChainID, ownerStartBalance)
	l2Contracts := deployRebalancerContractUniverse(t, owner, l2ChainID, ownerStartBalance)

	connectLMs(t, owner, l1Contracts, l2Contracts)

	transferBalances(t, owner, l1Contracts, l2Contracts, bridgeAdapterBalance, l1TokenPoolBalance)

	// create abi-encoded report for l1 lm
	// send 10 gwei from L1 to L2 using the bridge adapter
	l1SendReport := testonly_liquidity_manager.ILiquidityManagerLiquidityInstructions{
		SendLiquidityParams: []testonly_liquidity_manager.ILiquidityManagerSendLiquidityParams{
			{
				RemoteChainSelector: l2ChainID,
				Amount:              assets.GWei(10).ToInt(),
			},
		},
	}
	encodedSendReport, err := testOnlyLMABI.Pack("publicEncodeReport", l1SendReport)
	require.NoError(t, err, "failed to encode report")
	encodedSendReport = encodedSendReport[4:] // remove the function selector
	// send the report to the L1 LM
	_, err = l1Contracts.lm.PublicReport(owner, encodedSendReport, 1)
	require.NoError(t, err, "failed to report to l1 lm")
	l1Contracts.backend.Commit()
	// check that the L1 token pool balance has decreased
	// due to the withdrawal of liquidity
	l1LockReleasePoolWethBalance, err := l1Contracts.wethToken.BalanceOf(&bind.CallOpts{Context: testutils.Context(t)}, l1Contracts.lockReleasePoolAddr)
	require.NoError(t, err, "failed to get weth balance of l1 lock release pool")
	require.Equal(t, assets.Ether(l1TokenPoolBalance).Sub(assets.GWei(10)).ToInt(), l1LockReleasePoolWethBalance)

	// create abi-encoded report for l2 lm
	// receive the 10 gwei from L1 using the bridge adapter
	l2ReceiveReport := testonly_liquidity_manager.ILiquidityManagerLiquidityInstructions{
		ReceiveLiquidityParams: []testonly_liquidity_manager.ILiquidityManagerReceiveLiquidityParams{
			{
				RemoteChainSelector: l1ChainID,
				Amount:              assets.GWei(10).ToInt(),
			},
		},
	}
	encodedReceiveReport, err := testOnlyLMABI.Pack("publicEncodeReport", l2ReceiveReport)
	require.NoError(t, err, "failed to encode report")
	encodedReceiveReport = encodedReceiveReport[4:] // remove the function selector
	// send the report to the L2 LM
	_, err = l2Contracts.lm.PublicReport(owner, encodedReceiveReport, 1)
	require.NoError(t, err, "failed to report to l2 lm")
	l2Contracts.backend.Commit()
	// check that the L2 lock release pool balance has increased
	// due to liquidity injection
	l2LockReleaseWethBalance, err := l2Contracts.wethToken.BalanceOf(&bind.CallOpts{Context: testutils.Context(t)}, l2Contracts.lockReleasePoolAddr)
	require.NoError(t, err, "failed to get weth balance of l2 lm")
	require.Equal(t, assets.GWei(10).ToInt(), l2LockReleaseWethBalance)
}

func transferBalances(
	t *testing.T,
	owner *bind.TransactOpts,
	l1Contracts testUniverse,
	l2Contracts testUniverse,
	bridgeAdapterBalance int64,
	l1TokenPoolBalance int64,
) {
	// move some weth to the bridge adapters
	// so that they can transfer it to the liquidity manager
	// when it calls finalizeWithdrawal
	_, err := l2Contracts.wethToken.Transfer(owner, l2Contracts.bridgeAdapterAddress, assets.Ether(bridgeAdapterBalance).ToInt())
	require.NoError(t, err, "failed to transfer weth to bridge adapter on l2")
	l2Contracts.backend.Commit()

	// move some weth to the L1 lock/release pool
	// the LM will pull from this pool in order to send to the L2
	// check balances
	_, err = l1Contracts.wethToken.Transfer(owner, l1Contracts.lockReleasePoolAddr, assets.Ether(l1TokenPoolBalance).ToInt())
	require.NoError(t, err, "failed to transfer weth to lock/release pool on l1")
	l1Contracts.backend.Commit()

	// check balane of bridge adapter
	l2BridgeAdapterWethBalance, err := l2Contracts.wethToken.BalanceOf(&bind.CallOpts{Context: testutils.Context(t)}, l2Contracts.bridgeAdapterAddress)
	require.NoError(t, err, "failed to get weth balance of l2 bridge adapter")
	require.Equal(t, assets.Ether(bridgeAdapterBalance).ToInt(), l2BridgeAdapterWethBalance)

	// check balance of token pool
	l1LockReleasePoolWethBalance, err := l1Contracts.wethToken.BalanceOf(&bind.CallOpts{Context: testutils.Context(t)}, l1Contracts.lockReleasePoolAddr)
	require.NoError(t, err, "failed to get weth balance of l1 lock release pool")
	require.Equal(t, assets.Ether(5).ToInt(), l1LockReleasePoolWethBalance)

	// check balance of token pool through the LM, should be identical
	// to the balance retreived above
	liq, err := l1Contracts.lm.GetLiquidity(&bind.CallOpts{Context: testutils.Context(t)})
	require.NoError(t, err)
	require.Equal(t, l1LockReleasePoolWethBalance, liq)
}

func connectLMs(
	t *testing.T,
	owner *bind.TransactOpts,
	l1Contracts testUniverse,
	l2Contracts testUniverse,
) {
	// connect the liquidity managers together
	_, err := l1Contracts.lm.SetCrossChainLiquidityManager(owner, testonly_liquidity_manager.ILiquidityManagerCrossChainLiquidityManagerArgs{
		RemoteLiquidityManager: l2Contracts.lmAddress,
		LocalBridge:            l1Contracts.bridgeAdapterAddress,
		RemoteToken:            l2Contracts.wethAddress,
		RemoteChainSelector:    l2ChainID,
		Enabled:                true,
	})
	require.NoError(t, err, "failed to set cross chain lm on l1")
	l1Contracts.backend.Commit()
	_, err = l2Contracts.lm.SetCrossChainLiquidityManager(owner, testonly_liquidity_manager.ILiquidityManagerCrossChainLiquidityManagerArgs{
		RemoteLiquidityManager: l1Contracts.lmAddress,
		LocalBridge:            l2Contracts.bridgeAdapterAddress,
		RemoteToken:            l1Contracts.wethAddress,
		RemoteChainSelector:    l1ChainID,
		Enabled:                true,
	})
	require.NoError(t, err, "failed to set cross chain lm on l2")
	l2Contracts.backend.Commit()

	// check that the cross chain lm was set correctly on both lms
	l1RemoteLM, err := l1Contracts.lm.GetAllCrossChainLiquidityMangers(&bind.CallOpts{Context: testutils.Context(t)})
	require.NoError(t, err, "failed to get cross chain lm on l1")
	require.Len(t, l1RemoteLM, 1)
	require.Equal(t, l2Contracts.lmAddress, l1RemoteLM[0].RemoteLiquidityManager)
	require.Equal(t, l2Contracts.bridgeAdapterAddress, l1RemoteLM[0].LocalBridge)
	require.Equal(t, l2Contracts.wethAddress, l1RemoteLM[0].RemoteToken)
	require.Equal(t, l2ChainID, l1RemoteLM[0].RemoteChainSelector)
	require.True(t, l1RemoteLM[0].Enabled)

	l2RemoteLM, err := l2Contracts.lm.GetAllCrossChainLiquidityMangers(&bind.CallOpts{Context: testutils.Context(t)})
	require.NoError(t, err, "failed to get cross chain lm on l2")
	require.Len(t, l2RemoteLM, 1)
	require.Equal(t, l1Contracts.lmAddress, l2RemoteLM[0].RemoteLiquidityManager)
	require.Equal(t, l1Contracts.bridgeAdapterAddress, l2RemoteLM[0].LocalBridge)
	require.Equal(t, l1Contracts.wethAddress, l2RemoteLM[0].RemoteToken)
	require.Equal(t, l1ChainID, l2RemoteLM[0].RemoteChainSelector)
	require.True(t, l2RemoteLM[0].Enabled)
}

func deployRebalancerContractUniverse(
	t *testing.T,
	owner *bind.TransactOpts,
	chainID uint64,
	startBalance int64,
) testUniverse {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.From: {
			Balance: assets.Ether(startBalance).ToInt(),
		},
	}, 30e6)

	// Deploy wrapped ether contract
	// will act as the ERC-20 being bridged
	wethAddress, _, _, err := weth9.DeployWETH9(owner, backend)
	require.NoError(t, err, "failed to deploy WETH9 contract")
	backend.Commit()
	wethToken, err := weth9.NewWETH9(wethAddress, backend)
	require.NoError(t, err, "failed to create WETH9 wrapper")

	// deposit some eth into the weth contract
	_, err = wethToken.Deposit(&bind.TransactOpts{
		From:    owner.From,
		Signer:  owner.Signer,
		Value:   assets.Ether(100).ToInt(),
		Context: testutils.Context(t),
	})
	require.NoError(t, err, "failed to deposit eth into weth contract")

	// deploy arm and arm proxy.
	// required by the token pool
	armAddress, _, _, err := mock_arm_contract.DeployMockARMContract(owner, backend)
	require.NoError(t, err, "failed to deploy MockARMContract contract")
	backend.Commit()
	arm, err := mock_arm_contract.NewMockARMContract(armAddress, backend)
	require.NoError(t, err)
	armProxyAddress, _, _, err := arm_proxy_contract.DeployARMProxyContract(owner, backend, armAddress)
	require.NoError(t, err, "failed to deploy ARMProxyContract contract")
	backend.Commit()
	armProxy, err := arm_proxy_contract.NewARMProxyContract(armProxyAddress, backend)
	require.NoError(t, err)

	// deploy lock/release pool targeting the weth9 contract
	lockReleasePoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		owner, backend, wethAddress, []common.Address{}, armProxyAddress, true)
	require.NoError(t, err, "failed to deploy LockReleaseTokenPool contract")
	backend.Commit()
	lockReleasePool, err := lock_release_token_pool.NewLockReleaseTokenPool(lockReleasePoolAddress, backend)
	require.NoError(t, err)

	// deploy the liquidity manager and set the liquidity container to be the lock release pool
	lmAddress, _, _, err := testonly_liquidity_manager.DeployTestOnlyLiquidityManager(owner, backend, wethAddress, chainID, lockReleasePoolAddress)
	require.NoError(t, err, "failed to deploy testonly lm on l1")
	backend.Commit()
	lm, err := testonly_liquidity_manager.NewTestOnlyLiquidityManager(lmAddress, backend)
	require.NoError(t, err)
	liqContainer, err := lm.GetLocalLiquidityContainer(&bind.CallOpts{Context: testutils.Context(t)})
	require.NoError(t, err)
	require.Equal(t, lockReleasePoolAddress, liqContainer)

	// set the liquidity manager of the lock release pool to be the previously deployed LM
	_, err = lockReleasePool.SetLiquidityManager(owner, lmAddress)
	require.NoError(t, err, "failed to set liquidity manager on lock/release pool")
	backend.Commit()
	actualLM, err := lockReleasePool.GetLiquidityManager(&bind.CallOpts{Context: testutils.Context(t)})
	require.NoError(t, err)
	require.Equal(t, lmAddress, actualLM)

	// deploy the bridge adapter to point to the weth contract address
	bridgeAdapterAddress, _, _, err := mock_l1_bridge_adapter.DeployMockL1BridgeAdapter(owner, backend, wethAddress)
	require.NoError(t, err, "failed to deploy mock l1 bridge adapter")
	backend.Commit()
	bridgeAdapter, err := mock_l1_bridge_adapter.NewMockL1BridgeAdapter(bridgeAdapterAddress, backend)
	require.NoError(t, err)

	return testUniverse{
		backend:              backend,
		chainID:              chainID,
		wethAddress:          wethAddress,
		wethToken:            wethToken,
		armAddress:           armAddress,
		armContract:          arm,
		armProxyAddress:      armProxyAddress,
		armProxy:             armProxy,
		lockReleasePoolAddr:  lockReleasePoolAddress,
		lockReleasePool:      lockReleasePool,
		lmAddress:            lmAddress,
		lm:                   lm,
		bridgeAdapterAddress: bridgeAdapterAddress,
		bridgeAdapter:        bridgeAdapter,
	}
}
