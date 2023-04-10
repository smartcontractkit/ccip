package metatx_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/cross_chain_erc20_extension"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	forwarder_wrapper "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/forwarder"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/wrapped_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/metatx"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
)

func TestMetaERC20SameChain(t *testing.T) {
	chainID := uint64(1337)

	// deploys and owns contract
	_, contractOwner := generateKeyAndTransactor(t, chainID)
	// holder1Key sends tokens to holder2
	holder1Key, holder1 := generateKeyAndTransactor(t, chainID)
	// holder2Key receives tokens
	_, holder2 := generateKeyAndTransactor(t, chainID)
	// relayKey is the relayer that submits signed meta-transaction to the forwarder contract on-chain
	_, relay := generateKeyAndTransactor(t, chainID)

	chain := backends.NewSimulatedBackend(core.GenesisAlloc{
		contractOwner.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(1e18)),
		}}, ethconfig.Defaults.Miner.GasCeil)

	// deploys forwarder that verifies meta transaction signature and forwards requests to token
	forwarderAddress, forwarder := setUpForwarder(t, contractOwner, chain)

	totalTokens := big.NewInt(1e9)
	tokenAddress, token := setUpCrossChainERC20(t, contractOwner, chain, forwarderAddress, common.HexToAddress("0x0"), totalTokens, false)

	amount := assets.Ether(1).ToInt()

	// fund MetaERC20 contract with native ETH
	transferNative(t, contractOwner, tokenAddress, 50_000, amount, chain)

	sourceTokenEthBal, err := chain.BalanceAt(testutils.Context(t), tokenAddress, nil)
	require.NoError(t, err)
	require.Equal(t, amount, sourceTokenEthBal)

	t.Run("single same-chain meta transfer", func(t *testing.T) {
		// transfer MetaERC20 from contract owner to holder1
		transferToken(t, token, contractOwner, holder1, amount, chain)

		deadline := big.NewInt(int64(chain.Blockchain().CurrentHeader().Time + uint64(time.Hour)))

		calldata, calldataHash := generateMetaTransferCalldata(t, holder2.From, amount, chainID)

		signature, domainSeparatorHash, typeHash, forwarderNonce, err := metatx.SignMetaTransfer(*forwarder,
			holder1Key.ToEcdsaPrivKey(),
			holder1.From,
			tokenAddress,
			holder2.From,
			calldataHash,
			deadline)
		require.NoError(t, err)

		forwardRequest := forwarder_wrapper.IForwarderForwardRequest{
			From:           holder1.From,
			Target:         tokenAddress,
			Nonce:          forwarderNonce,
			Data:           calldata,
			ValidUntilTime: deadline,
		}

		transferNative(t, contractOwner, relay.From, 21_000, amount, chain)

		holder1BalanceBefore, err := token.BalanceOf(nil, holder1.From)
		require.NoError(t, err)

		// send meta transaction to forwarder
		_, err = forwarder.Execute(relay, forwardRequest, domainSeparatorHash, typeHash, nil, signature)
		require.NoError(t, err)
		chain.Commit()

		holder2Balance, err := token.BalanceOf(nil, holder2.From)
		require.NoError(t, err)
		require.Equal(t, holder2Balance, amount)

		holder1Balance, err := token.BalanceOf(nil, holder1.From)
		require.NoError(t, err)
		require.Equal(t, holder1Balance, holder1BalanceBefore.Sub(holder1BalanceBefore, amount))

		totalSupplyAfter, err := token.TotalSupply(nil)
		require.NoError(t, err)
		require.Equal(t, totalSupplyAfter, big.NewInt(0).Mul(totalTokens, big.NewInt(1e18)))
	})
}

func TestMetaERC20CrossChain(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, testhelpers.SourceChainID, testhelpers.DestChainID)

	// holder1Key sends tokens to holder2
	holder1Key, holder1 := generateKeyAndTransactor(t, ccipContracts.Source.Chain.Blockchain().Config().ChainID.Uint64())
	// holder2Key receives tokens
	_, holder2 := generateKeyAndTransactor(t, ccipContracts.Dest.Chain.Blockchain().Config().ChainID.Uint64())
	// relayKey is the relayer that submits signed meta-transaction to the forwarder contract on-chain
	_, relay := generateKeyAndTransactor(t, ccipContracts.Source.Chain.Blockchain().Config().ChainID.Uint64())

	forwarderAddress, forwarder := setUpForwarder(t, ccipContracts.Source.User, ccipContracts.Source.Chain)

	totalTokens := big.NewInt(1e9)
	sourceTokenAddress, sourceToken := setUpCrossChainERC20(t, ccipContracts.Source.User, ccipContracts.Source.Chain, forwarderAddress, ccipContracts.Source.Router.Address(), totalTokens, true)

	wrappedDestTokenPoolAddress, _, wrappedDestTokenPool, err := wrapped_token_pool.DeployWrappedTokenPool(ccipContracts.Dest.User, ccipContracts.Dest.Chain, "WrappedBankToken", "WBANK", 18, wrapped_token_pool.RateLimiterConfig{
		Capacity:  testhelpers.HundredLink,
		Rate:      big.NewInt(1e18),
		IsEnabled: true,
	})
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()

	sourcePoolAddress, _, sourcePool, err := lock_release_token_pool.DeployLockReleaseTokenPool(ccipContracts.Source.User, ccipContracts.Source.Chain, sourceTokenAddress, lock_release_token_pool.RateLimiterConfig{
		Capacity:  testhelpers.HundredLink,
		Rate:      big.NewInt(1e18),
		IsEnabled: true,
	})
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()

	// transfer tokens to source token pool to allow the pool to burn the source token
	//_, err = sourceToken.Transfer(c.Source.User, sourcePoolAddress, assets.Ether(10).ToInt())
	//require.NoError(c.t, err)
	//c.Source.Chain.Commit()
	//sourcePoolBal, err := sourceToken.BalanceOf(nil, sourcePoolAddress)
	//require.NoError(c.t, err)
	//require.Equal(c.t, assets.Ether(10).ToInt(), sourcePoolBal)

	// set onRamp as valid caller for source pool
	_, err = sourcePool.ApplyRampUpdates(ccipContracts.Source.User, []lock_release_token_pool.IPoolRampUpdate{
		{
			Ramp:    ccipContracts.Source.OnRamp.Address(),
			Allowed: true,
		},
	}, nil)
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()

	_, err = wrappedDestTokenPool.ApplyRampUpdates(ccipContracts.Dest.User, nil, []wrapped_token_pool.IPoolRampUpdate{
		{
			Ramp:    ccipContracts.Dest.OffRamp.Address(),
			Allowed: true,
		},
	})
	require.NoError(t, err)
	ccipContracts.Dest.Chain.Commit()

	wrappedNativeAddress, err := ccipContracts.Source.Router.GetWrappedNative(nil)
	require.NoError(t, err)

	// native token is used as fee token
	_, err = ccipContracts.Source.PriceRegistry.UpdatePrices(ccipContracts.Source.User, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: wrappedNativeAddress,
				UsdPerToken: big.NewInt(1e18), // 1usd
			},
		},
		DestChainId:   ccipContracts.Dest.ChainID,
		UsdPerUnitGas: big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9,
	})
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()
	_, err = ccipContracts.Source.PriceRegistry.ApplyFeeTokensUpdates(ccipContracts.Source.User, []common.Address{wrappedNativeAddress}, nil)
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()

	// add new token pool created above
	_, err = ccipContracts.Source.OnRamp.ApplyPoolUpdates(ccipContracts.Source.User, nil, []evm_2_evm_onramp.InternalPoolUpdate{
		{
			Token: sourceTokenAddress,
			Pool:  sourcePoolAddress,
		},
	})
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()

	// set token limit
	_, err = ccipContracts.Source.OnRamp.SetPrices(ccipContracts.Source.User, []common.Address{sourceTokenAddress}, []*big.Int{big.NewInt(5)})
	require.NoError(t, err)
	ccipContracts.Source.Chain.Commit()

	_, err = ccipContracts.Dest.OffRamp.ApplyPoolUpdates(ccipContracts.Dest.User, nil, []evm_2_evm_offramp.InternalPoolUpdate{
		{
			Token: sourceTokenAddress,
			Pool:  wrappedDestTokenPoolAddress,
		},
	})
	require.NoError(t, err)
	ccipContracts.Dest.Chain.Commit()

	_, err = ccipContracts.Dest.OffRamp.SetPrices(ccipContracts.Dest.User, []common.Address{wrappedDestTokenPoolAddress}, []*big.Int{big.NewInt(5)})
	require.NoError(t, err)
	ccipContracts.Dest.Chain.Commit()

	amount := assets.Ether(1).ToInt()
	transferNative(t, ccipContracts.Source.User, sourceTokenAddress, 50_000, amount, ccipContracts.Source.Chain)

	linkUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerLink": "8000000000000000000"}`))
		require.NoError(t, err)
	}))
	ethUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerETH": "2000000000000000000000"}`))
		require.NoError(t, err)
	}))
	metaERC20USD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerMetaERC20": "5000000000000000000"}`))
		require.NoError(t, err)
	}))
	wrapped, err := ccipContracts.Source.Router.GetWrappedNative(nil)
	require.NoError(t, err)
	tokenPricesUSDPipeline := fmt.Sprintf(`
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="UsdPerLink"];
link->link_parse;
eth [type=http method=GET url="%s"];
eth_parse [type=jsonparse path="UsdPerETH"];
eth->eth_parse;
metaERC20 [type=http method=GET url="%s"];
metaERC20_parse [type=jsonparse path="UsdPerMetaERC20"];
metaERC20->metaERC20_parse
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse), \\\"%s\\\":$(eth_parse), \\\"%s\\\":$(metaERC20_parse)}"];`,
		linkUSD.URL, ethUSD.URL, metaERC20USD.URL, ccipContracts.Dest.LinkToken.Address(), wrapped, sourceTokenAddress)
	defer linkUSD.Close()
	defer ethUSD.Close()
	defer metaERC20USD.Close()

	nodes, _ := testhelpers.SetUpNodesAndJobs(t, &ccipContracts, tokenPricesUSDPipeline)

	geCurrentSeqNum := 1

	t.Run("single cross-chain meta transfer", func(t *testing.T) {
		// transfer MetaERC20 from owner to holder1
		transferToken(t, sourceToken, ccipContracts.Source.User, holder1, amount, ccipContracts.Source.Chain)

		deadline := big.NewInt(int64(ccipContracts.Source.Chain.Blockchain().CurrentHeader().Time + uint64(time.Hour)))

		calldata, calldataHash := generateMetaTransferCalldata(t, holder2.From, amount, ccipContracts.Dest.ChainID)

		signature, domainSeparatorHash, typeHash, forwarderNonce, err := metatx.SignMetaTransfer(
			*forwarder,
			holder1Key.ToEcdsaPrivKey(),
			holder1.From,
			sourceTokenAddress,
			holder2.From,
			calldataHash,
			deadline)
		require.NoError(t, err)

		forwardRequest := forwarder_wrapper.IForwarderForwardRequest{
			From:           holder1.From,
			Target:         sourceTokenAddress,
			Nonce:          forwarderNonce,
			Data:           calldata,
			ValidUntilTime: deadline,
		}

		transferNative(t, ccipContracts.Source.User, relay.From, 21_000, amount, ccipContracts.Source.Chain)

		// send meta transaction to forwarder
		_, err = forwarder.Execute(relay, forwardRequest, domainSeparatorHash, typeHash, []byte{}, signature)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		gomega.NewWithT(t).Eventually(func() bool {
			ccipContracts.Dest.Chain.Commit()
			holder2Balance, err := wrappedDestTokenPool.BalanceOf(nil, holder2.From)
			require.NoError(t, err)
			return holder2Balance.Cmp(amount) == 0
		}, testutils.WaitTimeout(t), 5*time.Second).Should(gomega.BeTrue())

		eventSignatures := ccip.GetEventSignatures()

		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertExecState(t, ccipContracts, executionLogs[0], ccip.Success)

		//source token is locked in the token pool
		lockedTokenBal, err := sourceToken.BalanceOf(nil, sourcePoolAddress)
		require.NoError(t, err)
		require.Equal(t, lockedTokenBal, amount)

		// source total supply should stay the same
		sourceTotalSupply, err := sourceToken.TotalSupply(nil)
		require.NoError(t, err)
		require.Equal(t, sourceTotalSupply, big.NewInt(0).Mul(totalTokens, big.NewInt(1e18)))

		// new wrapped tokens minted on dest token
		destTotalSupply, err := wrappedDestTokenPool.TotalSupply(nil)
		require.NoError(t, err)
		require.Equal(t, destTotalSupply, amount)
	})
}

func setUpForwarder(t *testing.T, owner *bind.TransactOpts, chain *backends.SimulatedBackend) (common.Address, *forwarder_wrapper.Forwarder) {
	// deploys EIP 2771 forwarder contract that verifies signatures from meta transaction and forwards the call to recipient contract (i.e MetaERC20 token)
	forwarderAddress, _, forwarder, err := forwarder_wrapper.DeployForwarder(owner, chain)
	require.NoError(t, err)
	chain.Commit()
	// registers EIP712-compliant domain separator for MetaERC20 token
	_, err = forwarder.RegisterDomainSeparator(owner, metatx.CrossChainERC20ExtensionName, metatx.CrossChainERC20ExtensionVersion)
	require.NoError(t, err)
	chain.Commit()

	return forwarderAddress, forwarder
}

func generateMetaTransferCalldata(t *testing.T, receiver common.Address, amount *big.Int, chainID uint64) ([]byte, [32]byte) {
	calldataDefinition := `
	[
		{
			"inputs": [{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			}, {
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}, {
				"internalType": "uint64",
				"name": "destinationChainId",
				"type": "uint64"
			}],
			"name": "metaTransfer",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]
	`

	calldataAbi, err := abi.JSON(strings.NewReader(calldataDefinition))
	require.NoError(t, err)

	calldata, err := calldataAbi.Pack("metaTransfer", receiver, amount, chainID)
	require.NoError(t, err)

	calldataHashRaw := crypto.Keccak256(calldata)

	var calldataHash [32]byte
	copy(calldataHash[:], calldataHashRaw[:])

	return calldata, calldataHash
}

func generateKeyAndTransactor(t *testing.T, chainID uint64) (key ethkey.KeyV2, transactor *bind.TransactOpts) {
	key = cltest.MustGenerateRandomKey(t)
	transactor, err := bind.NewKeyedTransactorWithChainID(key.ToEcdsaPrivKey(), big.NewInt(0).SetUint64(chainID))
	require.NoError(t, err)
	return
}

func setUpCrossChainERC20(t *testing.T, owner *bind.TransactOpts, chain *backends.SimulatedBackend, forwarderAddress, routerAddress common.Address, totalSupply *big.Int, isCrossChainTransfer bool) (common.Address, *cross_chain_erc20_extension.CrossChainERC20Extension) {
	// deploys MetaERC20 token that enables meta transactions for same-chain and cross-chain token transfers
	tokenAddress, _, token, err := cross_chain_erc20_extension.DeployCrossChainERC20Extension(
		owner, chain, "BankToken", "BANK", big.NewInt(0).Mul(totalSupply, big.NewInt(1e18)), forwarderAddress, routerAddress, isCrossChainTransfer)
	require.NoError(t, err)
	chain.Commit()
	return tokenAddress, token
}

func transferToken(t *testing.T, token *cross_chain_erc20_extension.CrossChainERC20Extension, sender, receiver *bind.TransactOpts, amount *big.Int, chain *backends.SimulatedBackend) {
	senderBalanceBefore, err := token.BalanceOf(nil, sender.From)
	require.NoError(t, err)
	chain.Commit()

	_, err = token.Transfer(sender, receiver.From, amount)
	require.NoError(t, err)
	chain.Commit()

	receiverBal, err := token.BalanceOf(nil, receiver.From)
	require.NoError(t, err)
	require.Equal(t, amount, receiverBal)

	senderBal, err := token.BalanceOf(nil, sender.From)
	require.NoError(t, err)
	require.Equal(t, senderBalanceBefore.Sub(senderBalanceBefore, amount), senderBal)
}

func transferNative(t *testing.T, sender *bind.TransactOpts, receiverAddress common.Address, gasLimit uint64, amount *big.Int, chain *backends.SimulatedBackend) {
	nonce, err := chain.NonceAt(testutils.Context(t), sender.From, nil)
	require.NoError(t, err)
	tx := types.NewTransaction(
		nonce, receiverAddress,
		amount,
		gasLimit,
		assets.GWei(1).ToInt(),
		nil)
	signedTx, err := sender.Signer(sender.From, tx)
	require.NoError(t, err)
	err = chain.SendTransaction(testutils.Context(t), signedTx)
	require.NoError(t, err)
	chain.Commit()

	receiverBalance, err := chain.BalanceAt(context.Background(), receiverAddress, nil)
	require.NoError(t, err)
	require.Equal(t, amount, receiverBalance)
}
