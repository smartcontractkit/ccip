package testhelpers

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
)

var (
	// Source
	SourcePool       = "source pool"
	SourceSub        = "source sub"
	TollOnRampRouter = "toll onramp router"
	TollOnRamp       = "toll onramp"
	SubOnRamp        = "sub onramp"
	SubOnRampRouter  = "sub onramp router"

	// Dest
	TollOffRampRouter = "toll offramp router"
	TollOffRamp       = "toll offramp"
	SubOffRampRouter  = "sub offramp router"
	DestPool          = "dest pool"
	DestSub           = "dest sub"
	Receiver          = "receiver"
	Sender            = "sender"
)

type MaybeRevertReceiver struct {
	Receiver *maybe_revert_message_receiver.MaybeRevertMessageReceiver
	Strict   bool
}

type CCIPContracts struct {
	t                              *testing.T
	SourceChainID, DestChainID     *big.Int
	SourceUser, DestUser           *bind.TransactOpts
	SourceChain, DestChain         *backends.SimulatedBackend
	SourcePool, DestPool           *native_token_pool.NativeTokenPool
	SourceLinkToken, DestLinkToken *link_token_interface.LinkToken
	BlobVerifier                   *blob_verifier.BlobVerifier
	Receivers                      []MaybeRevertReceiver

	// Toll contracts
	TollOnRampFees    map[common.Address]*big.Int
	TollOnRampRouter  *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter
	TollOnRamp        *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp
	TollOffRampRouter *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter
	TollOffRamp       *any_2_evm_toll_offramp.EVM2EVMTollOffRamp

	// Sub contracts
	SubOnRampFee     *big.Int
	SubOnRampRouter  *evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter
	SubOnRamp        *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp
	SubOffRampRouter *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter
	SubOffRamp       *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp
}

type BalanceAssertion struct {
	Name     string
	Address  common.Address
	Expected string
	Getter   func(addr common.Address) *big.Int
	Within   string
}

type BalanceReq struct {
	Name   string
	Addr   common.Address
	Getter func(addr common.Address) *big.Int
}

func (c *CCIPContracts) GetSourceSubBalance(addr common.Address) *big.Int {
	bal, err := c.SubOnRampRouter.GetBalance(nil, addr)
	require.NoError(c.t, err)
	return bal
}

func (c *CCIPContracts) GetDestSubBalance(addr common.Address) *big.Int {
	sub, err := c.SubOffRampRouter.GetSubscription(nil, addr)
	require.NoError(c.t, err)
	return sub.Balance
}

func (c *CCIPContracts) GetSourceLinkBalance(addr common.Address) *big.Int {
	bal, err := c.SourceLinkToken.BalanceOf(nil, addr)
	require.NoError(c.t, err)
	return bal
}
func (c *CCIPContracts) GetDestLinkBalance(addr common.Address) *big.Int {
	bal, err := c.DestLinkToken.BalanceOf(nil, addr)
	require.NoError(c.t, err)
	return bal
}

func (c *CCIPContracts) AssertBalances(bas []BalanceAssertion) {
	for _, b := range bas {
		actual := b.Getter(b.Address)
		require.NotNil(c.t, actual, "%v getter return nil", b.Name)
		if b.Within == "" {
			assert.Equal(c.t, b.Expected, actual.String(), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
		} else {
			bi, _ := big.NewInt(0).SetString(b.Expected, 10)
			withinI, _ := big.NewInt(0).SetString(b.Within, 10)
			high := big.NewInt(0).Add(bi, withinI)
			low := big.NewInt(0).Sub(bi, withinI)
			assert.Equal(c.t, -1, actual.Cmp(high), "wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			assert.Equal(c.t, 1, actual.Cmp(low), "wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
		}
	}
}

func GetBalances(brs []BalanceReq) (map[string]*big.Int, error) {
	m := make(map[string]*big.Int)
	for _, br := range brs {
		m[br.Name] = br.Getter(br.Addr)
		if m[br.Name] == nil {
			return nil, fmt.Errorf("%v getter return nil", br.Name)
		}
	}
	return m, nil
}

func MustAddBigInt(a *big.Int, b string) *big.Int {
	bi, _ := big.NewInt(0).SetString(b, 10)
	return big.NewInt(0).Add(a, bi)
}

func MustSubBigInt(a *big.Int, b string) *big.Int {
	bi, _ := big.NewInt(0).SetString(b, 10)
	return big.NewInt(0).Sub(a, bi)
}

func SetupCCIPContracts(t *testing.T, sourceChainID, destChainID *big.Int) CCIPContracts {
	sourceChain, sourceUser := SetupChain(t)
	destChain, destUser := SetupChain(t)

	var hundredLink = big.NewInt(0)
	// 100 LINK
	hundredLink.SetString("100000000000000000000", 10)

	// Deploy link token and pool on source chain
	sourceLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceLinkToken, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)
	sourcePoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(sourceUser,
		sourceChain,
		sourceLinkTokenAddress)
	require.NoError(t, err)
	sourceChain.Commit()
	sourcePool, err := native_token_pool.NewNativeTokenPool(sourcePoolAddress, sourceChain)
	require.NoError(t, err)

	// Deploy link token and pool on destination chain
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	destLinkToken, err := link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Float the offramp pool
	o, err := destPool.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, destUser.From.String(), o.String())
	_, err = destLinkToken.Transfer(destUser, destPoolAddress, hundredLink)
	require.NoError(t, err)
	destChain.Commit()

	afnSourceAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		sourceUser,
		sourceChain,
	)
	require.NoError(t, err)
	sourceChain.Commit()

	// Create onramp router
	onRampRouterAddress, _, _, err := evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

	// Deploy onramp source chain
	onRampAddress, _, _, err := evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
		sourceUser,                               // user
		sourceChain,                              // client
		sourceChainID,                            // source chain id
		destChainID,                              // destinationChainIds
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		evm_2_evm_toll_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e12,
			MaxTokensLength:  5,
		},
		evm_2_evm_toll_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		onRampRouterAddress,
	)
	require.NoError(t, err)
	tollOnRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddress, sourceChain)
	require.NoError(t, err)
	tollOnRampFees := map[common.Address]*big.Int{
		sourceLinkTokenAddress: big.NewInt(1), // 1 juel relay fee.
	}
	_, err = tollOnRamp.SetFeeConfig(sourceUser, evm_2_evm_toll_onramp.EVM2EVMTollOnRampInterfaceFeeConfig{
		Fees:      []*big.Int{tollOnRampFees[sourceLinkTokenAddress]},
		FeeTokens: []common.Address{sourceLinkTokenAddress},
	})
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, onRampAddress, true)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = tollOnRamp.SetPrices(sourceUser, []common.Address{sourceLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)
	tollOnRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(onRampRouterAddress, sourceChain)
	require.NoError(t, err)
	_, err = tollOnRampRouter.SetOnRamp(sourceUser, destChainID, onRampAddress)
	require.NoError(t, err)
	sourceChain.Commit()

	afnDestAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		destUser,
		destChain,
	)
	require.NoError(t, err)
	destChain.Commit()

	// Deploy offramp dest chain
	blobVerifierAddress, _, _, err := blob_verifier.DeployBlobVerifier(
		destUser,    // user
		destChain,   // client
		destChainID, // dest chain id
		sourceChainID,
		afnDestAddress, // AFN address
		blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{},
	)
	require.NoError(t, err)
	blobVerifier, err := blob_verifier.NewBlobVerifier(blobVerifierAddress, destChain)
	require.NoError(t, err)
	// Set the pool to be the offramp
	destChain.Commit()
	offRampAddress, _, _, err := any_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		any_2_evm_toll_offramp.BaseOffRampInterfaceOffRampConfig{
			OnRampAddress:         onRampAddress,
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		blobVerifier.Address(),
		afnDestAddress,
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		any_2_evm_toll_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
	)
	require.NoError(t, err)
	tollOffRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	// Create offRampAddr router
	offRampRouterAddress, _, _, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	tollOffRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(offRampRouterAddress, destChain)
	require.NoError(t, err)
	_, err = tollOffRamp.SetRouter(destUser, offRampRouterAddress)
	require.NoError(t, err)
	_, err = tollOffRamp.SetPrices(destUser, []common.Address{destLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	// Deploy 2 revertable (one SS one non-SS)
	revertingMessageReceiver1Address, _, _, err := maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(destUser, destChain, false)
	require.NoError(t, err)
	revertingMessageReceiver1, _ := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(revertingMessageReceiver1Address, destChain)
	revertingMessageReceiver2Address, _, _, err := maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(destUser, destChain, false)
	require.NoError(t, err)
	revertingMessageReceiver2, _ := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(revertingMessageReceiver2Address, destChain)
	// Need to commit here, or we will hit the block gas limit when deploying the executor
	sourceChain.Commit()
	destChain.Commit()

	// Setup subscription contracts.
	subOnRampRouterAddress, _, _, err := evm_2_any_subscription_onramp_router.DeployEVM2AnySubscriptionOnRampRouter(
		sourceUser, sourceChain, evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouterInterfaceRouterConfig{
			Fee:      big.NewInt(0),
			FeeToken: sourceLinkTokenAddress,
			FeeAdmin: sourceUser.From,
		})
	require.NoError(t, err)
	subOnRampRouter, _ := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(subOnRampRouterAddress, sourceChain)
	subOnRampAddress, _, _, err := evm_2_evm_subscription_onramp.DeployEVM2EVMSubscriptionOnRamp(sourceUser, sourceChain, sourceChainID, destChainID,
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		evm_2_evm_subscription_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e12,
			MaxTokensLength:  5,
		},
		evm_2_evm_subscription_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		onRampRouterAddress,
	)
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, subOnRampAddress, true)
	require.NoError(t, err)
	subOnRamp, _ := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(subOnRampAddress, sourceChain)
	_, err = subOnRamp.SetRouter(sourceUser, subOnRampRouterAddress)
	require.NoError(t, err)
	_, err = subOnRampRouter.SetOnRamp(sourceUser, destChainID, subOnRampAddress)
	require.NoError(t, err)
	subOnRampFee := big.NewInt(1)
	_, err = subOnRampRouter.SetFee(sourceUser, subOnRampFee)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = subOnRamp.SetPrices(sourceUser, []common.Address{sourceLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	subOffRampRouterAddress, _, _, err := any_2_evm_subscription_offramp_router.DeployAny2EVMSubscriptionOffRampRouter(
		destUser, destChain, []common.Address{}, any_2_evm_subscription_offramp_router.SubscriptionInterfaceSubscriptionConfig{
			SetSubscriptionSenderDelay: 0,
			WithdrawalDelay:            0,
			FeeToken:                   destLinkTokenAddress,
		})
	require.NoError(t, err)
	destChain.Commit()
	subOffRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(subOffRampRouterAddress, destChain)
	require.NoError(t, err)
	subOffRampAddress, _, _, err := any_2_evm_subscription_offramp.DeployEVM2EVMSubscriptionOffRamp(destUser, destChain, sourceChainID, destChainID,
		any_2_evm_subscription_offramp.BaseOffRampInterfaceOffRampConfig{
			OnRampAddress:         subOnRampAddress,
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		blobVerifier.Address(),
		afnDestAddress,
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		any_2_evm_subscription_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
	)
	require.NoError(t, err)
	subOffRamp, _ := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(subOffRampAddress, destChain)
	_, err = destPool.SetOffRamp(destUser, subOffRampAddress, true)
	require.NoError(t, err)
	_, err = subOffRamp.SetRouter(destUser, subOffRampRouterAddress)
	require.NoError(t, err)
	_, err = subOffRampRouter.AddOffRamp(destUser, subOffRampAddress)
	require.NoError(t, err)
	destChain.Commit()
	_, err = subOffRamp.SetPrices(destUser, []common.Address{destLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	_, err = subOffRampRouter.GetSupportedTokensForExecutionFee(nil)
	require.NoError(t, err)
	_, err = subOffRampRouter.GetFeeToken(nil)
	require.NoError(t, err)
	// Enable onramps on blob verifier.
	_, err = blobVerifier.SetConfig(destUser, blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
		OnRamps:          []common.Address{onRampAddress, subOnRampAddress},
		MinSeqNrByOnRamp: []uint64{1, 1},
	})
	require.NoError(t, err)
	// Ensure we have at least finality blocks.
	for i := 0; i < 50; i++ {
		sourceChain.Commit()
		destChain.Commit()
	}

	return CCIPContracts{
		t:               t,
		SourceChainID:   sourceChainID,
		DestChainID:     destChainID,
		SourceUser:      sourceUser,
		DestUser:        destUser,
		SourceChain:     sourceChain,
		DestChain:       destChain,
		SourcePool:      sourcePool,
		DestPool:        destPool,
		SourceLinkToken: sourceLinkToken,
		DestLinkToken:   destLinkToken,
		BlobVerifier:    blobVerifier,
		Receivers:       []MaybeRevertReceiver{{Receiver: revertingMessageReceiver1, Strict: false}, {Receiver: revertingMessageReceiver2, Strict: true}},

		// Toll
		TollOnRampFees:    tollOnRampFees,
		TollOnRamp:        tollOnRamp,
		TollOnRampRouter:  tollOnRampRouter,
		TollOffRampRouter: tollOffRampRouter,
		TollOffRamp:       tollOffRamp,

		// Sub
		SubOnRampFee:     subOnRampFee,
		SubOnRamp:        subOnRamp,
		SubOnRampRouter:  subOnRampRouter,
		SubOffRampRouter: subOffRampRouter,
		SubOffRamp:       subOffRamp,
	}
}

func QueueSubRequest(
	t *testing.T,
	ccipContracts CCIPContracts,
	msgPayload string,
	tokens []common.Address,
	amounts []*big.Int,
	gasLimit *big.Int,
	receiver common.Address,
) *types.Transaction {
	msg := evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage{
		Receiver: receiver,
		Data:     []byte(msgPayload),
		Tokens:   tokens,
		Amounts:  amounts,
		GasLimit: gasLimit,
	}
	tx, err := ccipContracts.SubOnRampRouter.CcipSend(ccipContracts.SourceUser, ccipContracts.DestChainID, msg)
	require.NoError(t, err)
	return tx
}

func QueueRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, feeToken common.Address, feeTokenAmount *big.Int, gasLimit *big.Int) *types.Transaction {
	msg := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Receiver:       ccipContracts.Receivers[0].Receiver.Address(),
		Data:           []byte(msgPayload),
		Tokens:         tokens,
		Amounts:        amounts,
		FeeToken:       feeToken,
		FeeTokenAmount: feeTokenAmount,
		GasLimit:       gasLimit,
	}
	tx, err := ccipContracts.TollOnRampRouter.CcipSend(ccipContracts.SourceUser, ccipContracts.DestChainID, msg)
	require.NoError(t, err)
	return tx
}

func SendSubRequest(
	t *testing.T,
	ccipContracts CCIPContracts,
	msgPayload string,
	tokens []common.Address,
	amounts []*big.Int,
	gasLimit *big.Int,
	receiver common.Address,
) {
	tx := QueueSubRequest(t, ccipContracts, msgPayload, tokens, amounts, gasLimit, receiver)
	ConfirmTxs(t, []*types.Transaction{tx}, ccipContracts.SourceChain)
}

func SendRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, feeToken common.Address, feeTokenAmount *big.Int, gasLimit *big.Int) {
	tx := QueueRequest(t, ccipContracts, msgPayload, tokens, amounts, feeToken, feeTokenAmount, gasLimit)
	ConfirmTxs(t, []*types.Transaction{tx}, ccipContracts.SourceChain)
}

func SetupOnchainConfig(t *testing.T, ccipContracts CCIPContracts, oracles []confighelper.OracleIdentityExtra, reportingPluginConfig []byte) int64 {
	// Note We do NOT set the payees, payment is done in the OCR2Base implementation
	// Set the offramp offchainConfig.
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper.ContractSetConfigArgsForTests(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		reportingPluginConfig,
		50*time.Millisecond, // Max duration query
		1*time.Second,       // Max duration observation
		100*time.Millisecond,
		100*time.Millisecond,
		100*time.Millisecond,
		1, // faults
		nil,
	)

	require.NoError(t, err)
	lggr := logger.TestLogger(t)
	lggr.Infow("Setting Config on Oracle Contract",
		"signers", signers,
		"transmitters", transmitters,
		"threshold", threshold,
		"onchainConfig", onchainConfig,
		"encodedConfigVersion", offchainConfigVersion,
	)

	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	require.NoError(t, err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	require.NoError(t, err)

	blockBeforeConfig, err := ccipContracts.DestChain.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)
	// Set the DON on the offramp
	_, err = ccipContracts.BlobVerifier.SetConfig0(
		ccipContracts.DestUser,
		signerAddresses,
		transmitterAddresses,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	ccipContracts.DestChain.Commit()

	// Same DON on the toll offramp
	_, err = ccipContracts.TollOffRamp.SetConfig0(
		ccipContracts.DestUser,
		signerAddresses,
		transmitterAddresses,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	ccipContracts.DestChain.Commit()

	// Same DON on the sub offramp
	_, err = ccipContracts.SubOffRamp.SetConfig0(
		ccipContracts.DestUser,
		signerAddresses,
		transmitterAddresses,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	return blockBeforeConfig.Number().Int64()
}

func AssertTollExecSuccess(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log) {
	executionStateChanged, err := ccipContracts.TollOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != ccip.Success {
		t.Log("Execution failed")
		t.Fail()
	}
}

func AssertTollExecFailure(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log) {
	executionStateChanged, err := ccipContracts.TollOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != ccip.MessageStateFailure {
		t.Errorf("Execution Succeeded but expected failure")
	}
}

func AssertSubExecSuccess(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log) {
	executionStateChanged, err := ccipContracts.SubOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != ccip.Success {
		t.Log("Execution failed")
		t.Fail()
	}
}

func AssertSubExecFailure(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log) {
	executionStateChanged, err := ccipContracts.SubOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != ccip.MessageStateFailure {
		t.Errorf("Execution Succeeded but expected failure")
	}
}

func EventuallyReportRelayed(t *testing.T, ccipContracts CCIPContracts, onRamp common.Address, min, max int) {
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		minSeqNum, err := ccipContracts.BlobVerifier.GetExpectedNextSequenceNumber(nil, onRamp)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		ccipContracts.DestChain.Commit()
		t.Log("min seq num reported", minSeqNum)
		return minSeqNum > uint64(max)
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "report has not been relayed")
}

func EventuallyExecutionStateChangedToSuccess(t *testing.T, ccipContracts CCIPContracts, seqNum []uint64) {
	g := gomega.NewGomegaWithT(t)
	g.Eventually(func() ccip.MessageExecutionState {
		it, err := ccipContracts.SubOffRamp.FilterExecutionStateChanged(nil, seqNum)
		g.Expect(err).NotTo(gomega.HaveOccurred(), "Error filtering ExecutionStateChanged event")
		g.Expect(it.Next()).To(gomega.BeTrue(), "No ExecutionStateChanged event found")
		var state ccip.MessageExecutionState
		// if multiple events found return the last one
		for it.Next() {
			state = ccip.MessageExecutionState(it.Event.State)
		}
		return state
	}, testutils.WaitTimeout(t), time.Second).
		Should(gomega.Equal(ccip.Success), "ExecutionStateChanged Event")
}

func EventuallyRelayReportAccepted(t *testing.T, ccipContracts CCIPContracts, currentBlock uint64) blob_verifier.CCIPRelayReport {
	g := gomega.NewGomegaWithT(t)
	var report blob_verifier.CCIPRelayReport
	g.Eventually(func() []common.Address {
		it, err := ccipContracts.BlobVerifier.FilterReportAccepted(&bind.FilterOpts{Start: currentBlock})
		g.Expect(err).NotTo(gomega.HaveOccurred(), "Error filtering ReportAccepted event")
		g.Expect(it.Next()).To(gomega.BeTrue(), "No ReportAccepted event found")
		report = it.Event.Report
		if len(report.OnRamps) > 0 {
			t.Log("Report Accepted by blob verifier")
		}
		return report.OnRamps
	}, testutils.WaitTimeout(t), 1*time.Second).
		Should(gomega.ContainElement(ccipContracts.SubOnRamp.Address()), "report has not been relayed")
	return report
}

func ExecuteSubMessage(
	t *testing.T,
	ccipContracts CCIPContracts,
	req logpoller.Log,
	allReqs []logpoller.Log,
	report blob_verifier.CCIPRelayReport,
) uint64 {
	t.Log("Executing request manually")
	// Build full tree for report
	mctx := hasher.NewKeccakCtx()
	leafHasher := ccip.NewSubscriptionLeafHasher(ccipContracts.SourceChainID, ccipContracts.DestChainID, ccipContracts.SubOnRamp.Address(), mctx)

	var leafHashes [][32]byte
	for _, otherReq := range allReqs {
		hash, err := leafHasher.HashLeaf(otherReq.GetGethLog())
		require.NoError(t, err)
		leafHashes = append(leafHashes, hash)
	}
	intervalsByOnRamp := make(map[common.Address]blob_verifier.CCIPInterval)
	merkleRootsByOnRamp := make(map[common.Address][32]byte)
	for i, onRamp := range report.OnRamps {
		intervalsByOnRamp[onRamp] = report.Intervals[i]
		merkleRootsByOnRamp[onRamp] = report.MerkleRoots[i]
	}
	interval := intervalsByOnRamp[ccipContracts.SubOnRamp.Address()]
	decodedMsg, err := ccip.DecodeCCIPSubMessage(req.Data)
	require.NoError(t, err)
	innerIdx := int(decodedMsg.SequenceNumber - interval.Min)
	innerTree, err := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)
	innerProof := innerTree.Prove([]int{innerIdx})
	var onRampIdx int
	var outerTreeLeafs [][32]byte
	for i, onRamp := range report.OnRamps {
		if onRamp == ccipContracts.SubOnRamp.Address() {
			onRampIdx = i
		}
		outerTreeLeafs = append(outerTreeLeafs, merkleRootsByOnRamp[onRamp])
	}
	outerTree, err := merklemulti.NewTree(mctx, outerTreeLeafs)
	require.NoError(t, err)
	require.Equal(t, outerTree.Root(), report.RootOfRoots, "Roots donot match")

	outerProof := outerTree.Prove([]int{onRampIdx})

	offRampProof := any_2_evm_subscription_offramp.CCIPExecutionReport{
		SequenceNumbers:          []uint64{decodedMsg.SequenceNumber},
		TokenPerFeeCoinAddresses: []common.Address{ccipContracts.SourceLinkToken.Address()},
		TokenPerFeeCoin:          []*big.Int{big.NewInt(1)},
		EncodedMessages:          [][]byte{req.Data},
		InnerProofs:              innerProof.Hashes,
		InnerProofFlagBits:       ccip.ProofFlagsToBits(innerProof.SourceFlags),
		OuterProofs:              outerProof.Hashes,
		OuterProofFlagBits:       ccip.ProofFlagsToBits(outerProof.SourceFlags),
	}

	// Execute.
	tx, err := ccipContracts.SubOffRamp.Execute(ccipContracts.DestUser, offRampProof, true)
	require.NoError(t, err, "Executing manually")
	ccipContracts.DestChain.Commit()
	rec, err := ccipContracts.DestChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status, "manual execution failed")
	t.Log("Manual Execution completed")
	return decodedMsg.SequenceNumber
}
