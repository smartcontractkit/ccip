package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/sqlx"
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
	namedLogger := lggr.Named("JobORM")
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
	b.WriteString(`SELECT * FROM ccip_requests WHERE true`)
	if sourceChainId != nil {
		b.WriteString(fmt.Sprintf(" AND source_chain_id = '%s'", sourceChainId.String()))
	}
	if destChainId != nil {
		b.WriteString(fmt.Sprintf(" AND dest_chain_id = '%s'", destChainId.String()))
	}
	if minSeqNum != nil {
		b.WriteString(fmt.Sprintf(" AND seq_num >= CAST(%s AS NUMERIC(78,0))", minSeqNum.String()))
	}
	if maxSeqNum != nil {
		b.WriteString(fmt.Sprintf(" AND seq_num <= CAST(%s AS NUMERIC(78,0))", maxSeqNum.String()))
	}
	if status != "" {
		b.WriteString(fmt.Sprintf(" AND status = '%s'", status))
	}
	if executor != nil {
		b.WriteString(fmt.Sprintf(` AND executor = '\x%v'`, executor.String()[2:]))
	}
	if options != nil {
		b.WriteString(fmt.Sprintf(` AND options = '\x%v'`, hexutil.Encode(options)[2:]))
	}
	b.WriteString(` ORDER BY seq_num ASC`)
	err = q.Select(&reqs, b.String())
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
	ctx, cancel := pg.DefaultQueryCtx()
	defer cancel()
	res, err := q.ExecContext(ctx, sql, status, minSeqNum.String(), maxSeqNum.String(), sourceChainId.String(), destChainId.String())
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
	b.WriteString(fmt.Sprintf(`(CAST('%s' AS NUMERIC(78,0))`, seqNums[0].String()))
	for _, n := range seqNums[1:] {
		b.WriteString(fmt.Sprintf(`,CAST('%s' AS NUMERIC(78,0))`, n.String()))
	}
	b.WriteString(`)`)
	sql := fmt.Sprintf(`UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE seq_num IN %s 
		  AND source_chain_id = $2 
		  AND dest_chain_id = $3 
		RETURNING seq_num`, b.String())

	ctx, cancel := pg.DefaultQueryCtx()
	defer cancel()
	res, err := q.ExecContext(ctx, sql, status, sourceChainId.String(), destChainId.String())
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
	sql := fmt.Sprintf(`UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE now() > (updated_at + interval '%d seconds') 
			AND source_chain_id = $2
			AND dest_chain_id = $3
			AND status = $4`, expiryTimeoutSeconds)
	return q.ExecQ(sql, toStatus, sourceChainId.String(), destChainId.String(), fromStatus)
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
