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

	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink/core/assets"
	forwarder_wrapper "github.com/smartcontractkit/chainlink/core/gethwrappers/generated/forwarder"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/meta_erc20"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/services/metatx"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/assert"
)

var (
	sourceChainID = uint64(1000)
	destChainID   = uint64(1337)
)

func TestMetaERC20CrossChain(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, sourceChainID, destChainID)

	// holder1Key sends tokens to holder2
	holder1Key := cltest.MustGenerateRandomKey(t)
	holder1Transactor, err := bind.NewKeyedTransactorWithChainID(holder1Key.ToEcdsaPrivKey(), ccipContracts.Source.Chain.Blockchain().Config().ChainID)
	require.NoError(t, err)

	// holder2Key receives tokens
	holder2Key := cltest.MustGenerateRandomKey(t)
	holder2Transactor, err := bind.NewKeyedTransactorWithChainID(holder2Key.ToEcdsaPrivKey(), ccipContracts.Dest.Chain.Blockchain().Config().ChainID)
	require.NoError(t, err)

	// relayKey is the relayer that submits signed meta-transaction to the forwarder contract on-chain
	relayKey := cltest.MustGenerateRandomKey(t)
	relayTransactor, err := bind.NewKeyedTransactorWithChainID(relayKey.ToEcdsaPrivKey(), ccipContracts.Source.Chain.Blockchain().Config().ChainID)
	require.NoError(t, err)

	var (
		holder1 = holder1Transactor
		holder2 = holder2Transactor
		relay   = relayTransactor
	)

	forwarderAddress, forwarder := setUpForwarder(t, ccipContracts.Source.User, ccipContracts.Source.Chain)

	sourceMetaWrapperAddress, sourceMetaWrapper, destMetaWrapperAddress, destMetaWrapper := setUpMetaERC20Contract(t,
		ccipContracts.Source.User,
		ccipContracts.Dest.User,
		ccipContracts.Source.Chain,
		ccipContracts.Dest.Chain,
		ccipContracts.Source.Router.Address(),
		ccipContracts.Dest.Router.Address(),
		ccipContracts.Source.LinkToken.Address(),
		forwarderAddress,
	)

	linkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
		require.NoError(t, err)
	}))
	tokensPerFeeCoinPipeline := fmt.Sprintf(`
// Price 1 
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="JuelsPerETH"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];`,
		linkEth.URL, ccipContracts.Dest.LinkToken.Address())
	defer linkEth.Close()

	nodes, _ := testhelpers.SetUpNodesAndJobs(t, ccipContracts, tokensPerFeeCoinPipeline)

	geCurrentSeqNum := 1

	t.Run("single cross-chain meta transfer", func(t *testing.T) {

		// transfer MetaERC20 from owner to holder1
		_, err = sourceMetaWrapper.Transfer(ccipContracts.Source.User, holder1.From, assets.Ether(1).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		holder1Bal, err := sourceMetaWrapper.BalanceOf(nil, holder1.From)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(1).ToInt(), holder1Bal)

		deadline := big.NewInt(int64(ccipContracts.Source.Chain.Blockchain().CurrentHeader().Time + uint64(time.Hour)))

		calldata, calldataHash := generateMetaTransferCalldata(t, destMetaWrapperAddress, holder2.From, assets.Ether(1).ToInt(), ccipContracts.Dest.ChainID)

		v, r, s, domainSeparatorHash, typeHash, forwarderNonce, err := metatx.SignMetaTransfer(
			*forwarder,
			holder1Key.ToEcdsaPrivKey(),
			holder1.From,
			sourceMetaWrapperAddress,
			destMetaWrapperAddress,
			holder2.From,
			calldataHash,
			deadline)
		require.NoError(t, err)

		var signature []byte
		signature = append(signature, r[:]...)
		signature = append(signature, s[:]...)
		signature = append(signature, v)

		forwardRequest := forwarder_wrapper.IForwarderForwardRequest{
			From:           holder1.From,
			To:             sourceMetaWrapperAddress,
			Value:          big.NewInt(0),
			Nonce:          forwarderNonce,
			Data:           calldata,
			ValidUntilTime: deadline,
		}

		nonce, err := ccipContracts.Source.Chain.NonceAt(testutils.Context(t), ccipContracts.Source.User.From, nil)
		require.NoError(t, err)

		// transfer eth to relayer for gas
		tx := types.NewTransaction(
			nonce, relay.From,
			assets.Ether(1).ToInt(),
			21000,
			assets.GWei(1).ToInt(),
			nil)
		signedTx, err := ccipContracts.Source.User.Signer(ccipContracts.Source.User.From, tx)
		require.NoError(t, err)

		err = ccipContracts.Source.Chain.SendTransaction(testutils.Context(t), signedTx)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		relayNativeBalance, err := ccipContracts.Source.Chain.BalanceAt(context.Background(), relay.From, nil)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(1).ToInt(), relayNativeBalance)

		// send meta transaction to forwarder
		_, err = forwarder.Execute(relay, forwardRequest, domainSeparatorHash, typeHash, []byte{}, signature)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		gomega.NewWithT(t).Eventually(func() bool {
			ccipContracts.Dest.Chain.Commit()
			holder2Balance, err := destMetaWrapper.BalanceOf(nil, holder2.From)
			require.NoError(t, err)
			return holder2Balance.Cmp(assets.Ether(1).ToInt()) == 0
		}, testutils.WaitTimeout(t), 5*time.Second).Should(gomega.BeTrue())

		eventSignatures := ccip.GetEventSignatures()

		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertExecState(t, ccipContracts, executionLogs[0], ccip.Success)
	})
}

func setUpForwarder(t *testing.T, owner *bind.TransactOpts, chain *backends.SimulatedBackend) (common.Address, *forwarder_wrapper.Forwarder) {
	// deploys EIP 2771 forwarder contract that verifies signatures from meta transaction and forwards the call to recipient contract (i.e MetaERC20 token)
	forwarderAddress, _, forwarder, err := forwarder_wrapper.DeployForwarder(owner, chain)
	require.NoError(t, err)
	chain.Commit()
	// registers EIP712-compliant domain separator for MetaERC20 token
	_, err = forwarder.RegisterDomainSeparator(owner, metatx.MetaERC20Name, metatx.MetaERC20Version)
	require.NoError(t, err)
	chain.Commit()

	return forwarderAddress, forwarder
}

func setUpMetaERC20Contract(t *testing.T,
	sourceOwner *bind.TransactOpts,
	destOwner *bind.TransactOpts,
	sourceChain *backends.SimulatedBackend,
	destChain *backends.SimulatedBackend,
	sourceCCIPRouterAddress,
	destCCIPRouterAddress,
	sourceLinkTokenAddress,
	sourceForwarderAddress common.Address,
) (common.Address, *meta_erc20.MetaERC20, common.Address, *meta_erc20.MetaERC20) {
	// deploys MetaERC20 token that enables meta transactions for same-chain and cross-chain token transfers
	sourceMetaWrapperAddress, _, sourceMetaWrapper, err := meta_erc20.DeployMetaERC20(
		sourceOwner, sourceChain, assets.Ether(int64(1e18)).ToInt(),
		sourceCCIPRouterAddress)
	require.NoError(t, err)
	sourceChain.Commit()

	// fee token is used to pay for CCIP transaction
	_, err = sourceMetaWrapper.SetFeeToken(sourceOwner, sourceLinkTokenAddress)
	require.NoError(t, err)
	sourceChain.Commit()

	// authorizes EIP 2771 forwarder to relay requests to MetaERC20 token
	_, err = sourceMetaWrapper.SetForwarder(sourceOwner, sourceForwarderAddress)
	require.NoError(t, err)
	sourceChain.Commit()

	link, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)

	_, err = link.Approve(sourceOwner, sourceMetaWrapperAddress, assets.Ether(100).ToInt())
	require.NoError(t, err)
	sourceChain.Commit()

	// fund MetaERC20 contract with LINK to pay for token fee.
	// Funding MetaERC20 contract approves CCIP router to use LINK owned by MetaERC20 contract
	_, err = sourceMetaWrapper.Fund(sourceOwner, assets.Ether(100).ToInt())
	require.NoError(t, err)
	sourceChain.Commit()

	sourceMetaWrapperLinkBal, err := link.BalanceOf(nil, sourceMetaWrapperAddress)
	require.NoError(t, err)
	require.Equal(t, assets.Ether(100).ToInt(), sourceMetaWrapperLinkBal)

	// initialize destination ERC20 with 0 total balance
	// during cross chain transfers, token will be burnt from source chain contract and minted on the destination chain contract
	destMetaWrapperAddress, _, destMetaWrapper, err := meta_erc20.DeployMetaERC20(
		destOwner, destChain, assets.Ether(int64(0)).ToInt(),
		destCCIPRouterAddress)
	require.NoError(t, err)
	destChain.Commit()

	return sourceMetaWrapperAddress, sourceMetaWrapper, destMetaWrapperAddress, destMetaWrapper
}

func generateMetaTransferCalldata(t *testing.T, destTokenAddress, destReceiverAddress common.Address, amount *big.Int, chainID uint64) ([]byte, [32]byte) {
	calldataDefinition := `
	[
		{
			"inputs": [{
				"internalType": "address",
				"name": "destinationTokenAddress",
				"type": "address"
			}, {
				"internalType": "address",
				"name": "recipientAddress",
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

	calldata, err := calldataAbi.Pack("metaTransfer", destTokenAddress, destReceiverAddress, amount, chainID)
	require.NoError(t, err)

	calldataHashRaw := crypto.Keccak256(calldata)

	var calldataHash [32]byte
	copy(calldataHash[:], calldataHashRaw[:])

	return calldata, calldataHash
}
