package ccip

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/toll_sender_dapp"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
)

var HundredCoins = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100))

type NativeTokenConfig struct {
	ReleaseConfig struct {
		Rate     *big.Int
		Capacity *big.Int
	}
	LockConfig struct {
		Rate     *big.Int
		Capacity *big.Int
	}
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
		Msg("Setting off ramp for onramp router")
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
	instance   *afn_contract.AFNContract
	EthAddress common.Address
}

func (afn *AFN) Address() string {
	return afn.EthAddress.Hex()
}

type OnRampRouter struct {
	client     *blockchain.EthereumClient
	instance   *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter
	EthAddress common.Address
}

func (router *OnRampRouter) Address() string {
	return router.EthAddress.Hex()
}

func (router *OnRampRouter) SetOnRamp(chainID *big.Int, onRamp common.Address) error {
	opts, err := router.client.TransactionOpts(router.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("Router", router.Address()).
		Msg("Setting on ramp for onramp router")
	tx, err := router.instance.SetOnRamp(opts, chainID, onRamp)
	if err != nil {
		return err
	}
	log.Info().
		Str("onramp", onRamp.Hex()).
		Msg("onramp router is configured")
	return router.client.ProcessTransaction(tx)
}

func (router *OnRampRouter) CCIPSend(destChainId *big.Int, msg evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage) (*types.Transaction, error) {
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

type OnRamp struct {
	client     *blockchain.EthereumClient
	instance   *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp
	EthAddress common.Address
}

func (onRamp *OnRamp) Address() string {
	return onRamp.EthAddress.Hex()
}

func (onRamp *OnRamp) FilterCCIPSendRequested(
	currentBlock uint64,
) (*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequestedIterator, error) {
	filter := bind.FilterOpts{Start: currentBlock}
	return onRamp.instance.FilterCCIPSendRequested(&filter)
}

func (onRamp *OnRamp) SetFeeConfig(tokens []common.Address, fees []*big.Int) error {
	opts, err := onRamp.client.TransactionOpts(onRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("Router", onRamp.Address()).
		Msg("Setting Fee config onramp")
	tx, err := onRamp.instance.SetFeeConfig(opts, evm_2_evm_toll_onramp.Any2EVMTollOnRampInterfaceFeeConfig{
		Fees:      fees,
		FeeTokens: tokens,
	})
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

type OffRamp struct {
	client     *blockchain.EthereumClient
	instance   *any_2_evm_toll_offramp.EVM2EVMTollOffRamp
	EthAddress common.Address
}

func (offRamp *OffRamp) Address() string {
	return offRamp.EthAddress.Hex()
}

// SetConfig sets the offchain reporting protocol configuration
func (o *OffRamp) SetConfig(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {
	log.Info().Str("Contract Address", o.Address()).Msg("Configuring Offramp Contract")
	// Set Config
	opts, err := o.client.TransactionOpts(o.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Interface("signerAddresses", signers).
		Interface("transmitterAddresses", transmitters).
		Msg("Configuring OffRamp")
	tx, err := o.instance.SetConfig(
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
	return o.client.ProcessTransaction(tx)
}

func (offRamp *OffRamp) SetRouter(offRampRouterAddress common.Address) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.DefaultWallet)
	if err != nil {
		return err
	}
	log.Info().
		Str("Offramp", offRamp.Address()).
		Msg("Setting router for offramp")
	tx, err := offRamp.instance.SetRouter(opts, offRampRouterAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("offRampRouterAddress", offRampRouterAddress.Hex()).
		Msg("offRamp router is configured")
	return offRamp.client.ProcessTransaction(tx)
}

func (offRamp *OffRamp) FilterExecutionStateChanged(seqNumber []uint64) (
	*any_2_evm_toll_offramp.EVM2EVMTollOffRampExecutionStateChangedIterator, error,
) {
	return offRamp.instance.FilterExecutionStateChanged(nil, seqNumber)
}

type OffRampRouter struct {
	client     *blockchain.EthereumClient
	instance   *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter
	EthAddress common.Address
}

func (orr *OffRampRouter) Address() string {
	return orr.EthAddress.Hex()
}

type ReceiverDapp struct {
	client     *blockchain.EthereumClient
	instance   *receiver_dapp.ReceiverDapp
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
	tokens []common.Address,
	amount []*big.Int,
) (*types.Transaction, error) {
	opts, err := sender.client.TransactionOpts(sender.client.GetDefaultWallet())
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("Receiver Address", receiverAddr.String()).
		Interface("Tokens", tokens).
		Interface("Amounts", amount).
		Msg("Sending tokens")
	tx, err := sender.instance.SendTokens(opts, receiverAddr, tokens, amount)
	if err != nil {
		return nil, err
	}
	return tx, sender.client.ProcessTransaction(tx)
}
