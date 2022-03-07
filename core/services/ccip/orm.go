package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/pg"
)

// ORM We intend to use the same table for all xchain requests.
// TODO: I think we may need to pass in string based chainIDs
// in the future when we support non-evm chains, for now keep the interface EVM
// The triplet (seqNum, source, dest) defined the Message.
//go:generate mockery --name ORM --output ./mocks/ --case=underscore
type ORM interface {
	// Note always returns them sorted by seqNum
	Requests(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus, executor *common.Address, options []byte, opt ...pg.QOpt) ([]*Request, error)
	SaveRequest(request *Request, qopts ...pg.QOpt) error
	UpdateRequestStatus(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus, qopts ...pg.QOpt) error
	UpdateRequestSetStatus(sourceChainId, destChainId *big.Int, seqNums []*big.Int, status RequestStatus, qopts ...pg.QOpt) error
	ResetExpiredRequests(sourceChainId, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus RequestStatus, toStatus RequestStatus, qopts ...pg.QOpt) error
	RelayReport(seqNum *big.Int, qopts ...pg.QOpt) (RelayReport, error)
	SaveRelayReport(report RelayReport, qopts ...pg.QOpt) error
}

type orm struct {
	db   *sqlx.DB
	q    pg.Q
	lggr logger.Logger
	cfg  pg.LogConfig
}

var _ORM = (*orm)(nil)

func NewORM(db *sqlx.DB, lggr logger.Logger, cfg pg.LogConfig) ORM {
	namedLogger := lggr.Named("CCIPORM")
	return &orm{
		db:   db,
		q:    pg.NewQ(db, namedLogger, cfg),
		lggr: namedLogger,
		cfg:  cfg,
	}
}

// Note that executor can be left unset in the request, meaning anyone can execute.
// A nil executor as an argument here however means "don't filter on executor" and so it will return requests with both unset and set executors.
func (o *orm) Requests(sourceChainId, destChainId *big.Int, minSeqNum, maxSeqNum *big.Int, status RequestStatus, executor *common.Address, options []byte, qopts ...pg.QOpt) (reqs []*Request, err error) {
	q := o.q.WithOpts(qopts...)
	var b strings.Builder
	var params []interface{}
	b.WriteString(`SELECT * FROM ccip_requests WHERE true`)
	if sourceChainId != nil {
		b.WriteString(" AND source_chain_id = ?")
		params = append(params, sourceChainId.String())
	}
	if destChainId != nil {
		b.WriteString(" AND dest_chain_id = ?")
		params = append(params, destChainId.String())
	}
	if minSeqNum != nil {
		b.WriteString(" AND seq_num >= CAST(? AS NUMERIC(78,0))")
		params = append(params, minSeqNum.String())
	}
	if maxSeqNum != nil {
		b.WriteString(" AND seq_num <= CAST(? AS NUMERIC(78,0))")
		params = append(params, maxSeqNum.String())
	}
	if status != "" {
		b.WriteString(" AND status = ?")
		params = append(params, status)
	}
	if executor != nil {
		b.WriteString(` AND executor = ?`)
		params = append(params, fmt.Sprintf(`\x%v`, executor.String()[2:]))
	}
	if options != nil {
		b.WriteString(` AND options = ?`)
		params = append(params, fmt.Sprintf(`\x%v`, hexutil.Encode(options)[2:]))
	}
	b.WriteString(` ORDER BY seq_num ASC`)
	stmt := sqlx.Rebind(sqlx.DOLLAR, b.String())

	err = q.Select(&reqs, stmt, params...)
	return
}

func (o *orm) UpdateRequestStatus(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	// We return seqNum here to error if it doesn't exist
	sql := `UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE seq_num >= CAST($2 AS NUMERIC(78,0))
		  AND seq_num <= CAST($3 AS NUMERIC(78,0))
		  AND source_chain_id = $4 
		  AND dest_chain_id = $5 
		RETURNING seq_num`
	res, err := q.Exec(sql, status, minSeqNum.String(), maxSeqNum.String(), sourceChainId.String(), destChainId.String())
	if err != nil {
		return err
	}
	seqRange := big.NewInt(0).Sub(maxSeqNum, minSeqNum)
	expectedUpdates := seqRange.Add(seqRange, big.NewInt(1)).Int64()
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n != expectedUpdates {
		return fmt.Errorf("did not update expected num rows, got %v want %v", n, expectedUpdates)
	}
	return nil
}

func (o *orm) UpdateRequestSetStatus(sourceChainId, destChainId *big.Int, seqNums []*big.Int, status RequestStatus, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	if len(seqNums) == 0 {
		return nil
	}
	var b strings.Builder
	var params []interface{}

	b.WriteString(`UPDATE ccip_requests SET status = ?, updated_at = now() 
						WHERE seq_num IN`)
	params = append(params, status)
	b.WriteString(`(CAST(? AS NUMERIC(78,0))`)
	params = append(params, seqNums[0].String())

	for _, n := range seqNums[1:] {
		b.WriteString(`,CAST(? AS NUMERIC(78,0))`)
		params = append(params, n.String())
	}
	b.WriteString(`) AND source_chain_id = ? AND dest_chain_id = ? RETURNING seq_num`)
	params = append(params, sourceChainId.String())
	params = append(params, destChainId.String())

	stmt := sqlx.Rebind(sqlx.DOLLAR, b.String())
	res, err := q.Exec(stmt, params...)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(n) != len(seqNums) {
		return fmt.Errorf("did not update expected num rows, got %v want %v", n, len(seqNums))
	}
	return err
}

func (o *orm) ResetExpiredRequests(sourceChainId, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus RequestStatus, toStatus RequestStatus, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	sql := `UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE now() > (updated_at + $2) 
			AND source_chain_id = $3
			AND dest_chain_id = $4
			AND status = $5`
	return q.ExecQ(sql, toStatus, fmt.Sprintf("%d seconds", expiryTimeoutSeconds), sourceChainId.String(), destChainId.String(), fromStatus)
}

// Note requests will only be added in an unstarted status
func (o *orm) SaveRequest(request *Request, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	sql := `INSERT INTO ccip_requests 
    (seq_num, source_chain_id, dest_chain_id, sender, receiver, data, tokens, amounts, executor, options, raw, status, created_at, updated_at) 
    VALUES (:seq_num, :source_chain_id, :dest_chain_id, :sender, :receiver, :data, :tokens, :amounts, :executor, :options, :raw, 'unstarted', now(), now())
   	ON CONFLICT DO NOTHING`

	return q.ExecQNamed(sql, request)
}

func (o *orm) RelayReport(seqNum *big.Int, qopts ...pg.QOpt) (report RelayReport, err error) {
	q := o.q.WithOpts(qopts...)
	sql := `SELECT * FROM ccip_relay_reports WHERE min_seq_num <= $1 and max_seq_num >= $1`

	if err = q.Get(&report, sql, seqNum.String()); err != nil {
		return RelayReport{}, err
	}
	return
}

func (o *orm) SaveRelayReport(report RelayReport, qopts ...pg.QOpt) error {
	q := o.q.WithOpts(qopts...)
	sql := `INSERT INTO ccip_relay_reports (root, min_seq_num, max_seq_num, created_at) VALUES ($1, $2, $3, now()) ON CONFLICT DO NOTHING`

	return q.ExecQ(sql, report.Root[:], report.MinSeqNum.String(), report.MaxSeqNum.String())
}
