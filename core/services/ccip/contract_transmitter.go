package ccip

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/sqlx"
)

var (
	_ ocrtypes.ContractTransmitter = &OfframpTransmitter{}
	_ ocrtypes.ContractTransmitter = &ExecutionTransmitter{}
)

type ExecutionTransmitter struct {
	contractABI abi.ABI
	transmitter Transmitter
	contract    *message_executor.MessageExecutor
	lggr        logger.Logger
}

func NewExecutionTransmitter(
	contract *message_executor.MessageExecutor,
	contractABI abi.ABI,
	transmitter Transmitter,
	lggr logger.Logger,
) *ExecutionTransmitter {
	return &ExecutionTransmitter{
		contractABI: contractABI,
		transmitter: transmitter,
		contract:    contract,
		lggr:        lggr,
	}
}

func (oc *ExecutionTransmitter) Transmit(ctx context.Context, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signatures []ocrtypes.AttributedOnchainSignature) error {
	rs, ss, vs := splitSigs(signatures)
	rawReportCtx := evmutil.RawReportContext(reportCtx)
	oc.lggr.Infow("executor transmitting report", "report", hex.EncodeToString(report), "rawReportCtx", rawReportCtx, "contractAddress", oc.contract.Address())

	payload, err := oc.contractABI.Pack("transmit", rawReportCtx, []byte(report), rs, ss, vs)
	if err != nil {
		return errors.Wrap(err, "abi.Pack failed")
	}

	return errors.Wrap(oc.transmitter.CreateEthTransaction(ctx, oc.contract.Address(), payload, report), "failed to send Eth transaction")
}

func (oc *ExecutionTransmitter) LatestConfigDigestAndEpoch(ctx context.Context) (configDigest ocrtypes.ConfigDigest, epoch uint32, err error) {
	//! TODO: not efficient!
	it, err := oc.contract.FilterTransmited(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	})
	if err != nil {
		return ocrtypes.ConfigDigest{}, 0, err
	}
	defer it.Close()
	for it.Next() {
		fmt.Println("LatestConfigDigestAndEpoch:", it.Event)
		configDigest = it.Event.ConfigDigest
		epoch = it.Event.Epoch
	}

	if it.Error() != nil {
		return ocrtypes.ConfigDigest{}, 0, it.Error()
	}
	return configDigest, epoch, nil
}

func (oc *ExecutionTransmitter) FromAccount() ocrtypes.Account {
	return ocrtypes.Account(oc.transmitter.FromAddress().String())
}

type OfframpTransmitter struct {
	contractABI abi.ABI
	transmitter Transmitter
	contract    *single_token_offramp.SingleTokenOffRamp
	lggr        logger.Logger
}

func NewOfframpTransmitter(
	contract *single_token_offramp.SingleTokenOffRamp,
	contractABI abi.ABI,
	transmitter Transmitter,
	lggr logger.Logger,
) *OfframpTransmitter {
	return &OfframpTransmitter{
		contractABI: contractABI,
		transmitter: transmitter,
		contract:    contract,
		lggr:        lggr,
	}
}

func splitSigs(signatures []ocrtypes.AttributedOnchainSignature) (rs [][32]byte, ss [][32]byte, vs [32]byte) {
	for i, as := range signatures {
		r, s, v, err := evmutil.SplitSignature(as.Signature)
		if err != nil {
			panic("eventTransmit(ev): error in SplitSignature")
		}
		rs = append(rs, r)
		ss = append(ss, s)
		vs[i] = v
	}
	return
}

func (oc *OfframpTransmitter) Transmit(ctx context.Context, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signatures []ocrtypes.AttributedOnchainSignature) error {
	rs, ss, vs := splitSigs(signatures)
	rawReportCtx := evmutil.RawReportContext(reportCtx)
	oc.lggr.Debugw("Transmitting report", "report", hex.EncodeToString(report), "rawReportCtx", rawReportCtx, "contractAddress", oc.contract.Address())

	payload, err := oc.contractABI.Pack("transmit", rawReportCtx, []byte(report), rs, ss, vs)
	if err != nil {
		return errors.Wrap(err, "abi.Pack failed")
	}

	return errors.Wrap(oc.transmitter.CreateEthTransaction(ctx, oc.contract.Address(), payload, report), "failed to send Eth transaction")
}

func (oc *OfframpTransmitter) LatestConfigDigestAndEpoch(ctx context.Context) (configDigest ocrtypes.ConfigDigest, epoch uint32, err error) {
	//! TODO: not efficient!
	it, err := oc.contract.FilterTransmited(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	})
	if err != nil {
		return ocrtypes.ConfigDigest{}, 0, err
	}
	defer it.Close()
	for it.Next() {
		fmt.Println("LatestConfigDigestAndEpoch:", it.Event)
		configDigest = it.Event.ConfigDigest
		epoch = it.Event.Epoch
	}

	if it.Error() != nil {
		return ocrtypes.ConfigDigest{}, 0, it.Error()
	}
	return configDigest, epoch, nil
}

func (oc *OfframpTransmitter) FromAccount() ocrtypes.Account {
	return ocrtypes.Account(oc.transmitter.FromAddress().String())
}

type relayTransmitter struct {
	txm                        TxManager
	db                         *sqlx.DB
	fromAddress                gethCommon.Address
	gasLimit                   uint64
	strategy                   bulletprooftxmanager.TxStrategy
	ec                         eth.Client
	sourceChainID, destChainID *big.Int
}

type TxManager interface {
	CreateEthTransaction(newTx bulletprooftxmanager.NewTx, qs ...pg.QOpt) (etx bulletprooftxmanager.EthTx, err error)
}

type Transmitter interface {
	CreateEthTransaction(ctx context.Context, toAddress gethCommon.Address, payload []byte, report []byte) error
	FromAddress() gethCommon.Address
}

// NewTransmitter creates a new eth relayTransmitter
func NewRelayTransmitter(txm TxManager, db *sqlx.DB, sourceChainID, destChainID *big.Int, fromAddress gethCommon.Address, gasLimit uint64, strategy bulletprooftxmanager.TxStrategy, ec eth.Client) Transmitter {
	return &relayTransmitter{
		txm:           txm,
		db:            db,
		fromAddress:   fromAddress,
		gasLimit:      gasLimit,
		strategy:      strategy,
		ec:            ec,
		sourceChainID: sourceChainID,
		destChainID:   destChainID,
	}
}

func (t *relayTransmitter) CreateEthTransaction(ctx context.Context, toAddress gethCommon.Address, payload []byte, report []byte) error {
	twoGwei := big.NewInt(2_000_000_000)
	a := toAddress
	gasEstimate, err := t.ec.EstimateGas(ctx, ethereum.CallMsg{
		From:      t.fromAddress,
		To:        &a,
		Gas:       0,
		GasFeeCap: twoGwei,
		GasTipCap: twoGwei,
		Data:      payload,
	})
	if err != nil {
		return errors.Wrap(err, "failed to estimating gas cost for ccip execute transaction")
	}
	if gasEstimate > t.gasLimit {
		return errors.Wrap(err, fmt.Sprintf("gas estimate of %d exceeds gas limit set by node %d", gasEstimate, t.gasLimit))
	}
	// TODO: As soon as gorm is removed, these can two db ops need to be in the same transaction
	_, err = t.txm.CreateEthTransaction(bulletprooftxmanager.NewTx{
		FromAddress:    t.fromAddress,
		ToAddress:      toAddress,
		EncodedPayload: payload,
		GasLimit:       gasEstimate,
		Meta:           nil,
		Strategy:       t.strategy,
	})
	return errors.Wrap(err, "creating ETH ccip relay transaction")
}

func (t *relayTransmitter) FromAddress() gethCommon.Address {
	return t.fromAddress
}

type executeTransmitter struct {
	txm                        TxManager
	db                         *sqlx.DB
	fromAddress                gethCommon.Address
	gasLimit                   uint64
	strategy                   bulletprooftxmanager.TxStrategy
	ec                         eth.Client
	sourceChainID, destChainID *big.Int
}

// NewTransmitter creates a new eth relayTransmitter
func NewExecuteTransmitter(txm TxManager, db *sqlx.DB, sourceChainID, destChainID *big.Int, fromAddress gethCommon.Address, gasLimit uint64, strategy bulletprooftxmanager.TxStrategy, ec eth.Client) Transmitter {
	return &executeTransmitter{
		txm:           txm,
		db:            db,
		fromAddress:   fromAddress,
		gasLimit:      gasLimit,
		strategy:      strategy,
		ec:            ec,
		sourceChainID: sourceChainID,
		destChainID:   destChainID,
	}
}

func (t *executeTransmitter) CreateEthTransaction(ctx context.Context, toAddress gethCommon.Address, payload []byte, report []byte) error {
	twoGwei := big.NewInt(2_000_000_000)
	a := toAddress
	gasEstimate, err := t.ec.EstimateGas(ctx, ethereum.CallMsg{
		From:      t.fromAddress,
		To:        &a,
		Gas:       0,
		GasFeeCap: twoGwei,
		GasTipCap: twoGwei,
		Data:      payload,
	})
	if err != nil {
		return errors.Wrap(err, "failed to estimating gas cost for ccip execute transaction")
	}
	if gasEstimate > t.gasLimit {
		return errors.Wrap(err, fmt.Sprintf("gas estimate of %d exceeds gas limit set by node %d", gasEstimate, t.gasLimit))
	}
	// TODO: As soon as gorm is removed, these can two db ops need to be in the same transaction
	_, err = t.txm.CreateEthTransaction(bulletprooftxmanager.NewTx{
		FromAddress:    t.fromAddress,
		ToAddress:      toAddress,
		EncodedPayload: payload,
		GasLimit:       gasEstimate,
		Meta:           nil,
		Strategy:       t.strategy,
	})
	return errors.Wrap(err, "creating ETH ccip execute transaction")
}

func (t *executeTransmitter) FromAddress() gethCommon.Address {
	return t.fromAddress
}
