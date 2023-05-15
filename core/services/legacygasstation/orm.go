package legacygasstation

import (
	"github.com/smartcontractkit/sqlx"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var _ ORM = &orm{}

type orm struct {
	q pg.Q
}

// NewORM creates an ORM scoped to chainID.
// TODO: implement pruning logic if needed
func NewORM(db *sqlx.DB, lggr logger.Logger, cfg pg.QConfig) ORM {
	namedLogger := lggr.Named("LegacyGasStation")
	q := pg.NewQ(db, namedLogger, cfg)
	return &orm{
		q: q,
	}
}

func (o *orm) SelectBySourceChainIDAndStatus(sourceChainID uint64, status types.Status, qopts ...pg.QOpt) (txs []types.LegacyGaslessTx, err error) {
	q := o.q.WithOpts(qopts...)
	err = q.Select(&txs, `
		SELECT * FROM legacy_gasless_txs 
			WHERE legacy_gasless_txs.source_chain_id = $1 
			AND legacy_gasless_txs.tx_status = $2
		`, sourceChainID, status.String())
	return
}

func (o *orm) SelectByDestChainIDAndStatus(destChainID uint64, status types.Status, qopts ...pg.QOpt) (txs []types.LegacyGaslessTx, err error) {
	q := o.q.WithOpts(qopts...)
	err = q.Select(&txs, `
		SELECT * FROM legacy_gasless_txs
			WHERE legacy_gasless_txs.destination_chain_id = $1 
			AND legacy_gasless_txs.tx_status = $2
		`, destChainID, status.String())
	return
}

// InsertLegacyGaslessTx is idempotent
func (o *orm) InsertLegacyGaslessTx(tx types.LegacyGaslessTx, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	err := q.ExecQ(`INSERT INTO legacy_gasless_txs (legacy_gasless_tx_id, forwarder_address, from_address, target_address, receiver_address, nonce, amount, source_chain_id, destination_chain_id, valid_until_time, tx_signature, tx_status, token_name, token_version, eth_tx_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, NOW(), NOW())`,
		tx.ID,
		tx.Forwarder,
		tx.From,
		tx.Target,
		tx.Receiver,
		tx.Nonce,
		tx.Amount,
		tx.SourceChainID,
		tx.DestinationChainID,
		tx.ValidUntilTime,
		tx.Signature[:],
		tx.Status.String(),
		tx.TokenName,
		tx.TokenVersion,
		tx.EthTxID,
	)
	return err
}

// UpdateLegacyGaslessTx updates legacy gasless transaction with status, ccip message ID (optional), failure reason (optional)
func (o *orm) UpdateLegacyGaslessTx(tx types.LegacyGaslessTx, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	_, err := q.Exec(`UPDATE legacy_gasless_txs SET 
	tx_status = $2,
	ccip_message_id = $3,
	failure_reason = $4,
	updated_at = NOW()
	WHERE legacy_gasless_tx_id = $1`,
		tx.ID,
		tx.Status.String(),
		tx.CCIPMessageID,
		tx.FailureReason,
	)
	return err
}

func (o *orm) SelectEthTxsBySourceChainIDAndState(sourceChainID uint64, state txmgrtypes.TxState, qopts ...pg.QOpt) (ethTxs []txmgr.DbEthTx, err error) {
	q := o.q.WithOpts(qopts...)
	err = q.Select(&ethTxs, `SELECT eth_txes.* FROM legacy_gasless_txs
	INNER JOIN eth_txes ON eth_txes.id = legacy_gasless_txs.eth_tx_id
	WHERE legacy_gasless_txs.source_chain_id = $1 
	AND eth_txes.state = $2
	`, sourceChainID, state)
	return
}
