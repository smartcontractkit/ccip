package ccip

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
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

// LockReleaseTokenPool represents a LockReleaseTokenPool address
type LockReleaseTokenPool struct {
	client     blockchain.EVMClient
	instance   *lock_release_token_pool.LockReleaseTokenPool
	EthAddress common.Address
}

func (pool *LockReleaseTokenPool) Address() string {
	return pool.EthAddress.Hex()
}

func (pool *LockReleaseTokenPool) LockOrBurnToken(linkToken contracts.LinkToken, amount *big.Int) error {
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
	opts, err := pool.client.TransactionOpts(pool.client.GetDefaultWallet())
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

func (pool *LockReleaseTokenPool) SetOnRamp(onRamp common.Address) error {
	opts, err := pool.client.TransactionOpts(pool.client.GetDefaultWallet())
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

func (pool *LockReleaseTokenPool) SetOffRamp(offRamp common.Address) error {
	opts, err := pool.client.TransactionOpts(pool.client.GetDefaultWallet())
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
	client     blockchain.EVMClient
	instance   *mock_afn_contract.MockAFNContract
	EthAddress common.Address
}

func (afn *AFN) Address() string {
	return afn.EthAddress.Hex()
}

type CommitStore struct {
	client     blockchain.EVMClient
	instance   *commit_store.CommitStore
	EthAddress common.Address
}

func (bv *CommitStore) Address() string {
	return bv.EthAddress.Hex()
}

// SetOCR2Config sets the offchain reporting protocol configuration
func (b *CommitStore) SetOCR2Config(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {
	log.Info().Str("Contract Address", b.Address()).Msg("Configuring OCR config for CommitStore Contract")
	// Set Config
	opts, err := b.client.TransactionOpts(b.client.GetDefaultWallet())
	if err != nil {
		return err
	}

	log.Info().
		Interface("signerAddresses", signers).
		Interface("transmitterAddresses", transmitters).
		Msg("Configuring CommitStore")
	tx, err := b.instance.SetOCR2Config(
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

func (b *CommitStore) FilterReportAccepted(currentBlock uint64) (*commit_store.CommitStoreReportAcceptedIterator, error) {
	return b.instance.FilterReportAccepted(&bind.FilterOpts{Start: currentBlock})
}

func (b *CommitStore) GetNextSeqNumber(onRamp common.Address) (uint64, error) {
	return b.instance.GetExpectedNextSequenceNumber(nil, onRamp)
}

type MessageReceiver struct {
	client     blockchain.EVMClient
	instance   *simple_message_receiver.SimpleMessageReceiver
	EthAddress common.Address
}

type ReceiverDapp struct {
	client     blockchain.EVMClient
	instance   *maybe_revert_message_receiver.MaybeRevertMessageReceiver
	EthAddress common.Address
}

func (rDapp *ReceiverDapp) Address() string {
	return rDapp.EthAddress.Hex()
}

type FeeManager struct {
	client     blockchain.EVMClient
	instance   *fee_manager.FeeManager
	EthAddress common.Address
}

func (c *FeeManager) Address() string {
	return c.EthAddress.Hex()
}

func (c *FeeManager) SetFeeUpdater(addr common.Address) error {
	opts, err := c.client.TransactionOpts(c.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := c.instance.SetFeeUpdater(opts, addr)
	if err != nil {
		return err
	}
	log.Info().
		Str("updater", addr.Hex()).
		Msg("FeeManager updater set")
	return c.client.ProcessTransaction(tx)
}

func (c *FeeManager) UpdateFees(feeUpdates []fee_manager.InternalFeeUpdate) error {
	opts, err := c.client.TransactionOpts(c.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := c.instance.UpdateFees(opts, feeUpdates)
	if err != nil {
		return err
	}
	log.Info().
		Msg("FeeManager fee updated")
	return c.client.ProcessTransaction(tx)
}

type Router struct {
	client     blockchain.EVMClient
	Instance   *router.Router
	EthAddress common.Address
}

func (router *Router) Copy() *Router {
	r := *router.Instance
	return &Router{
		client:     router.client,
		Instance:   &r,
		EthAddress: router.EthAddress,
	}
}

func (router *Router) Address() string {
	return router.EthAddress.Hex()
}

func (router *Router) SetOnRamp(chainID uint64, onRamp common.Address) error {
	opts, err := router.client.TransactionOpts(router.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Str("Router", router.Address()).
		Msg("Setting on ramp for router")
	tx, err := router.Instance.SetOnRamp(opts, chainID, onRamp)
	if err != nil {
		return err
	}
	log.Info().
		Str("onRamp", onRamp.Hex()).
		Msg("Router is configured")
	return router.client.ProcessTransaction(tx)
}

func (router *Router) CCIPSend(destChainId uint64, msg router.ConsumerEVM2AnyMessage) (*types.Transaction, error) {
	opts, err := router.client.TransactionOpts(router.client.GetDefaultWallet())
	if err != nil {
		return nil, err
	}
	tx, err := router.Instance.CcipSend(opts, destChainId, msg)
	if err != nil {
		return nil, err
	}
	return tx, router.client.ProcessTransaction(tx)
}

func (router *Router) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	opts, err := router.client.TransactionOpts(router.client.GetDefaultWallet())
	if err != nil {
		return nil, err
	}
	tx, err := router.Instance.AddOffRamp(opts, offRamp)
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("offRamp", offRamp.Hex()).
		Msg("offRamp is added to Router")
	return tx, router.client.ProcessTransaction(tx)
}

func (router *Router) GetFee(destinationChainId uint64, message router.ConsumerEVM2AnyMessage) (*big.Int, error) {
	return router.Instance.GetFee(nil, destinationChainId, message)
}

type OnRamp struct {
	client     blockchain.EVMClient
	instance   *evm_2_evm_onramp.EVM2EVMOnRamp
	EthAddress common.Address
}

func (onRamp *OnRamp) Address() string {
	return onRamp.EthAddress.Hex()
}

func (onRamp *OnRamp) FilterCCIPSendRequested(
	currentBlock uint64,
) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequestedIterator, error) {
	filter := bind.FilterOpts{Start: currentBlock}
	return onRamp.instance.FilterCCIPSendRequested(&filter)
}

func (onRamp *OnRamp) SetTokenPrices(tokens []common.Address, prices []*big.Int) error {
	opts, err := onRamp.client.TransactionOpts(onRamp.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Str("OnRamp", onRamp.Address()).
		Msg("Setting OnRamp token prices")
	tx, err := onRamp.instance.SetPrices(opts, tokens, prices)
	if err != nil {
		return err
	}
	return onRamp.client.ProcessTransaction(tx)
}

type OffRamp struct {
	client     blockchain.EVMClient
	instance   *evm_2_evm_offramp.EVM2EVMOffRamp
	EthAddress common.Address
}

func (offRamp *OffRamp) Address() string {
	return offRamp.EthAddress.Hex()
}

// SetOCR2Config sets the offchain reporting protocol configuration
func (offRamp *OffRamp) SetOCR2Config(
	signers []common.Address,
	transmitters []common.Address,
	f uint8,
	onchainConfig []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
) error {
	log.Info().Str("Contract Address", offRamp.Address()).Msg("Configuring OffRamp Contract")
	// Set Config
	opts, err := offRamp.client.TransactionOpts(offRamp.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Interface("signerAddresses", signers).
		Interface("transmitterAddresses", transmitters).
		Msg("Configuring OffRamp")
	tx, err := offRamp.instance.SetOCR2Config(
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

func (offRamp *OffRamp) SetRouter(offRampRouterAddress common.Address) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Str("OffRamp", offRamp.Address()).
		Msg("Setting router for offramp")
	tx, err := offRamp.instance.SetRouter(opts, offRampRouterAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("offRampRouterAddress", offRampRouterAddress.Hex()).
		Msg("OffRamp router is configured")
	return offRamp.client.ProcessTransaction(tx)
}

func (offRamp *OffRamp) FilterExecutionStateChanged(seqNumber []uint64, messageId [][32]byte, currentBlockOnDest uint64) (
	*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChangedIterator, error,
) {
	return offRamp.instance.FilterExecutionStateChanged(&bind.FilterOpts{Start: currentBlockOnDest}, seqNumber, messageId)
}

func (offRamp *OffRamp) SetTokenPrices(tokens []common.Address, prices []*big.Int) error {
	opts, err := offRamp.client.TransactionOpts(offRamp.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	log.Info().
		Str("OffRamp", offRamp.Address()).
		Msg("Setting OffRamp token prices")
	tx, err := offRamp.instance.SetPrices(opts, tokens, prices)
	if err != nil {
		return err
	}
	return offRamp.client.ProcessTransaction(tx)
}
