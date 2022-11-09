package ccip

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/toll_sender_dapp"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
)

var HundredCoins = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100))

type RateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type AFNConfig struct {
	AFNWeightsByParticipants map[string]*big.Int // mapping : AFN participant address => weight
	ThresholdForBlessing     *big.Int
	ThresholdForBadSignal    *big.Int
}

// NativeTokenPool represents a NativeTokenPool address
type NativeTokenPool struct {
	client     *blockchain.EthereumClient
	instance   *native_token_pool.NativeTokenPool
	EthAddress common.Address
}

func (pool *NativeTokenPool) Address() string {
	return pool.EthAddress.Hex()
}

func (pool *NativeTokenPool) LockOrBurnToken(linkToken contracts.LinkToken, amount *big.Int) error {
	log.Info().
		Str("Link Token", linkToken.Address()).
		Str("Token Pool", pool.Address()).
		Msg("Initiating transferring of token to token pool")
	err := linkToken.Transfer(pool.Address(), amount)
	if err != nil {
		return err
	}
	err = pool.client.WaitForEvents()
	if err != nil {
		return err
	}
	opts, err := pool.client.TransactionOpts(pool.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Msg("Initiating locking Tokens in pool")
	tx, err := pool.instance.LockOrBurn(opts, amount)
	if err != nil {
		return err
	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Msg("Pool is filled with tokens")
	return pool.client.ProcessTransaction(tx)
}

func (pool *NativeTokenPool) SetOnRamp(onRamp common.Address) error {
	opts, err := pool.client.TransactionOpts(pool.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Msg("Setting on ramp for onramp router")
	tx, err := pool.instance.SetOnRamp(opts, onRamp, true)
	if err != nil {
		return err
	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Str("OnRamp", onRamp.Hex()).
		Msg("OnRamp is set")
	return pool.client.ProcessTransaction(tx)
}

func (pool *NativeTokenPool) SetOffRamp(offRamp common.Address) error {
	opts, err := pool.client.TransactionOpts(pool.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Msg("Setting off ramp for Token Pool")
	tx, err := pool.instance.SetOffRamp(opts, offRamp, true)
	if err != nil {
		return err
	}
	log.Info().
		Str("Token Pool", pool.Address()).
		Str("OffRamp", offRamp.Hex()).
		Msg("OffRamp is set")
	return pool.client.ProcessTransaction(tx)
}

type AFN struct {
	client     *blockchain.EthereumClient
	instance   *mock_afn_contract.MockAFNContract
	EthAddress common.Address
}

func (afn *AFN) Address() string {
	return afn.EthAddress.Hex()
}

type TollOnRampRouter struct {
	client     *blockchain.EthereumClient
	instance   *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter
	EthAddress common.Address
}

func (router *TollOnRampRouter) Address() string {
	return router.EthAddress.Hex()
}

func (router *TollOnRampRouter) SetOnRamp(chainID *big.Int, onRamp common.Address) error {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("TollOnRampRouter", router.Address()).
		Msg("Setting on ramp for onramp router")
	tx, err := router.instance.SetOnRamp(opts, chainID, onRamp)
	if err != nil {
		return err
	}
	log.Info().
		Str("TollOnRamp", onRamp.Hex()).
		Msg("TollOnRampRouter router is configured")
	return router.client.ProcessTransaction(tx)
}

func (router *TollOnRampRouter) CCIPSend(destChainId *big.Int, msg evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage) (*types.Transaction, error) {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return nil, err
	}
	tx, err := router.instance.CcipSend(opts, destChainId, msg)
	if err != nil {
		return nil, err
	}
	return tx, router.client.ProcessTransaction(tx)
}

type SubOnRampRouter struct {
	client     *blockchain.EthereumClient
	instance   *evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter
	EthAddress common.Address
}

func (router *SubOnRampRouter) Address() string {
	return router.EthAddress.Hex()
}

func (router *SubOnRampRouter) SetOnRamp(chainID *big.Int, onRamp common.Address) error {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("SubOnRampRouter", router.Address()).
		Msg("Setting on ramp for onramp router")
	tx, err := router.instance.SetOnRamp(opts, chainID, onRamp)
	if err != nil {
		return err
	}
	log.Info().
		Str("SubOnRamp", onRamp.Hex()).
		Msg("SubOnRampRouter is configured")
	return router.client.ProcessTransaction(tx)
}

func (router *SubOnRampRouter) SetFee(fee *big.Int) error {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("fee", fee.String()).
		Msg("Setting fee on SubOnRampRouter")
	tx, err := router.instance.SetFee(opts, fee)
	if err != nil {
		return err
	}
	return router.client.ProcessTransaction(tx)
}

func (router *SubOnRampRouter) CCIPSend(destChainId *big.Int, msg evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = 5000000
	tx, err := router.instance.CcipSend(opts, destChainId, msg)
	if err != nil {
		return nil, err
	}
	return tx, router.client.ProcessTransaction(tx)
}

func (router *SubOnRampRouter) FundSubscription(fee *big.Int) error {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("fee", fee.String()).
		Msg("funding FundSubscription SubOnRampRouter")
	tx, err := router.instance.FundSubscription(opts, fee)
	if err != nil {
		return err
	}
	return router.client.ProcessTransaction(tx)
}

func (router *SubOnRampRouter) GetBalance(sender common.Address) (*big.Int, error) {
	log.Info().
		Str("sender", sender.Hex()).
		Msg("getting subscription SubOnRampRouter")
	return router.instance.GetBalance(nil, sender)
}

type TollOnRamp struct {
	client     *blockchain.EthereumClient
	instance   *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp
	EthAddress common.Address
}

func (onRamp *TollOnRamp) Address() string {
	return onRamp.EthAddress.Hex()
}

func (onRamp *TollOnRamp) FilterCCIPSendRequested(
	currentBlock uint64,
) (*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequestedIterator, error) {
	filter := bind.FilterOpts{Start: currentBlock}
	return onRamp.instance.FilterCCIPSendRequested(&filter)
}

func (onRamp *TollOnRamp) SetFeeConfig(tokens []common.Address, fees []*big.Int) error {
	opts, err := onRamp.client.TransactionOpts(onRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("TollOnRamp", onRamp.Address()).
		Msg("Setting Fee config TollOnRamp")
	tx, err := onRamp.instance.SetFeeConfig(opts, evm_2_evm_toll_onramp.EVM2EVMTollOnRampInterfaceFeeConfig{
		Fees:      fees,
		FeeTokens: tokens,
	})
	if err != nil {
		return err
	}
	return onRamp.client.ProcessTransaction(tx)
}

func (onRamp *TollOnRamp) SetTokenPrices(tokens []common.Address, prices []*big.Int) error {
	opts, err := onRamp.client.TransactionOpts(onRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("TollOnRamp", onRamp.Address()).
		Msg("Setting TollOnRamp token prices")
	tx, err := onRamp.instance.SetPrices(opts, tokens, prices)
	if err != nil {
		return err
	}
	return onRamp.client.ProcessTransaction(tx)
}

type SubOnRamp struct {
	client     *blockchain.EthereumClient
	instance   *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp
	EthAddress common.Address
}

func (onRamp *SubOnRamp) Address() string {
	return onRamp.EthAddress.Hex()
}

func (onRamp *SubOnRamp) FilterCCIPSendRequested(
	currentBlock uint64,
) (*evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator, error) {
	filter := bind.FilterOpts{Start: currentBlock}
	return onRamp.instance.FilterCCIPSendRequested(&filter)
}

func (onRamp *SubOnRamp) SetTokenPrices(tokens []common.Address, prices []*big.Int) error {
	opts, err := onRamp.client.TransactionOpts(onRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("SubOnRamp", onRamp.Address()).
		Msg("Setting SubOnRamp token prices")
	tx, err := onRamp.instance.SetPrices(opts, tokens, prices)
	if err != nil {
		return err
	}
	return onRamp.client.ProcessTransaction(tx)
}

type BlobVerifier struct {
	client     *blockchain.EthereumClient
	instance   *blob_verifier.BlobVerifier
	EthAddress common.Address
}

func (bv *BlobVerifier) Address() string {
	return bv.EthAddress.Hex()
}

// SetOCRConfig sets the offchain reporting protocol configuration
func (b *BlobVerifier) SetOCRConfig(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {
	log.Info().Str("Contract Address", b.Address()).Msg("Configuring OCR config for BlobVerifier Contract")
	// Set Config
	opts, err := b.client.TransactionOpts(b.client.GetDefaultWallet())
	if err != nil {
		return err
	}

	log.Info().
		Interface("signerAddresses", signers).
		Interface("transmitterAddresses", transmitters).
		Msg("Configuring BlobVerifier")
	tx, err := b.instance.SetConfig0(
		opts,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)

	if err != nil {
		return err
	}
	return b.client.ProcessTransaction(tx)
}

func (b *BlobVerifier) FilterReportAccepted(currentBlock uint64) (*blob_verifier.BlobVerifierReportAcceptedIterator, error) {
	return b.instance.FilterReportAccepted(&bind.FilterOpts{Start: currentBlock})
}

func (b *BlobVerifier) GetNextSeqNumber(onRamp common.Address) (uint64, error) {
	return b.instance.GetExpectedNextSequenceNumber(nil, onRamp)
}

type MessageReceiver struct {
	client     *blockchain.EthereumClient
	instance   *simple_message_receiver.SimpleMessageReceiver
	EthAddress common.Address
}

type TollOffRamp struct {
	client     *blockchain.EthereumClient
	instance   *any_2_evm_toll_offramp.EVM2EVMTollOffRamp
	EthAddress common.Address
}

func (offRamp *TollOffRamp) Address() string {
	return offRamp.EthAddress.Hex()
}

// SetConfig sets the offchain reporting protocol configuration
func (offRamp *TollOffRamp) SetConfig(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {
	log.Info().Str("Contract Address", offRamp.Address()).Msg("Configuring TollOffRamp Contract")
	// Set Config
	opts, err := offRamp.client.TransactionOpts(offRamp.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Interface("signerAddresses", signers).
		Interface("transmitterAddresses", transmitters).
		Msg("Configuring TollOffRamp")
	tx, err := offRamp.instance.SetConfig0(
		opts,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)

	if err != nil {
		return err
	}
	return offRamp.client.ProcessTransaction(tx)
}

func (offRamp *TollOffRamp) SetRouter(offRampRouterAddress common.Address) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("TollOffRamp", offRamp.Address()).
		Msg("Setting router for offramp")
	tx, err := offRamp.instance.SetRouter(opts, offRampRouterAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("offRampRouterAddress", offRampRouterAddress.Hex()).
		Msg("TollOffRamp router is configured")
	return offRamp.client.ProcessTransaction(tx)
}

func (offRamp *TollOffRamp) FilterExecutionStateChanged(seqNumber []uint64, currentBlockOnDest uint64) (
	*any_2_evm_toll_offramp.EVM2EVMTollOffRampExecutionStateChangedIterator, error,
) {
	return offRamp.instance.FilterExecutionStateChanged(&bind.FilterOpts{Start: currentBlockOnDest}, seqNumber)
}

func (offRamp *TollOffRamp) SetTokenPrices(tokens []common.Address, prices []*big.Int) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("TollOffRamp", offRamp.Address()).
		Msg("Setting TollOffRamp token prices")
	tx, err := offRamp.instance.SetPrices(opts, tokens, prices)
	if err != nil {
		return err
	}
	return offRamp.client.ProcessTransaction(tx)
}

type SubOffRamp struct {
	client     *blockchain.EthereumClient
	instance   *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp
	EthAddress common.Address
}

func (offRamp *SubOffRamp) Address() string {
	return offRamp.EthAddress.Hex()
}

// SetConfig sets the offchain reporting protocol configuration
func (offRamp *SubOffRamp) SetConfig(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {
	log.Info().Str("Contract Address", offRamp.Address()).Msg("Configuring SubOffRamp Contract")
	// Set Config
	opts, err := offRamp.client.TransactionOpts(offRamp.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Interface("signerAddresses", signers).
		Interface("transmitterAddresses", transmitters).
		Msg("Configuring SubOffRamp")
	tx, err := offRamp.instance.SetConfig0(
		opts,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)

	if err != nil {
		return err
	}
	return offRamp.client.ProcessTransaction(tx)
}

func (offRamp *SubOffRamp) SetRouter(offRampRouterAddress common.Address) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("SubOffRamp", offRamp.Address()).
		Msg("Setting router for SubOffRamp")
	tx, err := offRamp.instance.SetRouter(opts, offRampRouterAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("SubOffRampRouterAddress", offRampRouterAddress.Hex()).
		Msg("SubOffRamp router is configured")
	return offRamp.client.ProcessTransaction(tx)
}

func (offRamp *SubOffRamp) FilterExecutionStateChanged(seqNumber []uint64, currentBlock uint64) (
	*any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRampExecutionStateChangedIterator, error,
) {
	return offRamp.instance.FilterExecutionStateChanged(&bind.FilterOpts{Start: currentBlock}, seqNumber)
}

func (offRamp *SubOffRamp) SetTokenPrices(tokens []common.Address, prices []*big.Int) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("SubOffRamp", offRamp.Address()).
		Msg("Setting SubOffRamp token prices")
	tx, err := offRamp.instance.SetPrices(opts, tokens, prices)
	if err != nil {
		return err
	}
	return offRamp.client.ProcessTransaction(tx)
}

type TollOffRampRouter struct {
	client     *blockchain.EthereumClient
	instance   *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter
	EthAddress common.Address
}

func (orr *TollOffRampRouter) Address() string {
	return orr.EthAddress.Hex()
}

type SubOffRampRouter struct {
	client     *blockchain.EthereumClient
	instance   *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter
	EthAddress common.Address
}

func (orr *SubOffRampRouter) Address() string {
	return orr.EthAddress.Hex()
}

func (orr *SubOffRampRouter) CreateSubscription(senders []common.Address, receiver common.Address, strict bool, balance *big.Int) error {
	opts, err := orr.client.TransactionOpts(orr.client.DefaultWallet)
	if err != nil {
		return err
	}
	sub := any_2_evm_subscription_offramp_router.SubscriptionInterfaceOffRampSubscription{
		Senders:          senders,
		Receiver:         receiver,
		StrictSequencing: strict,
		Balance:          balance,
	}
	log.Info().
		Str("SubOffRampRouter", orr.Address()).
		Interface("Subscription", sub).
		Msg("creating subscription")
	tx, err := orr.instance.CreateSubscription(opts, sub)
	if err != nil {
		return err
	}
	return orr.client.ProcessTransaction(tx)
}

func (orr *SubOffRampRouter) GetSubscription(receiver common.Address) (*big.Int, error) {
	log.Info().
		Str("SubOffRampRouter", orr.Address()).
		Interface("receiver", receiver.Hex()).
		Msg("getting subscription")
	sub, err := orr.instance.GetSubscription(nil, receiver)
	if err != nil {
		return nil, err
	}
	return sub.Balance, nil
}

type ReceiverDapp struct {
	client     *blockchain.EthereumClient
	instance   *maybe_revert_message_receiver.MaybeRevertMessageReceiver
	EthAddress common.Address
}

func (rDapp *ReceiverDapp) Address() string {
	return rDapp.EthAddress.Hex()
}

type TollSender struct {
	client     *blockchain.EthereumClient
	instance   *toll_sender_dapp.TollSenderDapp
	EthAddress common.Address
}

func (sender *TollSender) Address() string {
	return sender.EthAddress.Hex()
}

func (sender *TollSender) SendTokens(
	receiverAddr common.Address,
	tokens []toll_sender_dapp.CCIPEVMTokenAndAmount,
) (*types.Transaction, error) {
	opts, err := sender.client.TransactionOpts(sender.client.GetDefaultWallet())
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("Receiver Address", receiverAddr.String()).
		Interface("Tokens", tokens).
		Msg("Sending tokens")
	tx, err := sender.instance.SendTokens(opts, receiverAddr, tokens)
	if err != nil {
		return nil, err
	}
	return tx, sender.client.ProcessTransaction(tx)
}
