package ccip

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink/core/services/postgres"

	"github.com/smartcontractkit/sqlx"
)

// ORM We intend to use the same table for all xchain requests.
// TODO: I think we may need to pass in string based chainIDs
// in the future when we support non-evm chains, for now keep the interface EVM
// The triplet (seqNum, source, dest) defined the Message.
//go:generate mockery --name ORM --output ./mocks/ --case=underscore
type ORM interface {
	// Note always returns them sorted by seqNum
	Requests(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus, executor *common.Address, options []byte) ([]*Request, error)
	SaveRequest(request *Request) error
	UpdateRequestStatus(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus) error
	UpdateRequestSetStatus(sourceChainId, destChainId *big.Int, seqNums []*big.Int, status RequestStatus) error
	ResetExpiredRequests(sourceChainId, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus RequestStatus, toStatus RequestStatus) error
	RelayReport(seqNum *big.Int) (RelayReport, error)
	SaveRelayReport(report RelayReport) error
}

type orm struct {
	db *sqlx.DB
}

var _ORM = (*orm)(nil)

func NewORM(db *sqlx.DB) ORM {
	return &orm{db}
}

// Note that executor can be left unset in the request, meaning anyone can execute.
// A nil executor as an argument here however means "don't filter on executor" and so it will return requests with both unset and set executors.
func (o *orm) Requests(sourceChainId, destChainId *big.Int, minSeqNum, maxSeqNum *big.Int, status RequestStatus, executor *common.Address, options []byte) (reqs []*Request, err error) {
	q := `SELECT * FROM ccip_requests WHERE true`
	if sourceChainId != nil {
		q += fmt.Sprintf(" AND source_chain_id = '%s'", sourceChainId.String())
	}
	if destChainId != nil {
		q += fmt.Sprintf(" AND dest_chain_id = '%s'", destChainId.String())
	}
	if minSeqNum != nil {
		q += fmt.Sprintf(" AND seq_num >= CAST(%s AS NUMERIC(78,0))", minSeqNum.String())
	}
	if maxSeqNum != nil {
		q += fmt.Sprintf(" AND seq_num <= CAST(%s AS NUMERIC(78,0))", maxSeqNum.String())
	}
	if status != "" {
		q += fmt.Sprintf(" AND status = '%s'", status)
	}
	if executor != nil {
		q += fmt.Sprintf(` AND executor = '\x%v'`, executor.String()[2:])
	}
	if options != nil {
		q += fmt.Sprintf(` AND options = '\x%v'`, hexutil.Encode(options)[2:])
	}
	q += ` ORDER BY seq_num ASC`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	err = o.db.SelectContext(ctx, &reqs, q)
	return
}

func (o *orm) UpdateRequestStatus(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus) error {
	// We return seqNum here to error if it doesn't exist
	q := `UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE seq_num >= CAST($2 AS NUMERIC(78,0))
		  AND seq_num <= CAST($3 AS NUMERIC(78,0))
		  AND source_chain_id = $4 
		  AND dest_chain_id = $5 
		RETURNING seq_num`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	res, err := o.db.ExecContext(ctx, q, status, minSeqNum.String(), maxSeqNum.String(), sourceChainId.String(), destChainId.String())
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

func (o *orm) UpdateRequestSetStatus(sourceChainId, destChainId *big.Int, seqNums []*big.Int, status RequestStatus) error {
	if len(seqNums) == 0 {
		return nil
	}
	seqNumsSet := fmt.Sprintf(`(CAST('%s' AS NUMERIC(78,0))`, seqNums[0].String())
	for _, n := range seqNums[1:] {
		seqNumsSet += fmt.Sprintf(`,CAST('%s' AS NUMERIC(78,0))`, n.String())
	}
	seqNumsSet += `)`
	q := fmt.Sprintf(`UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE seq_num IN %s 
		  AND source_chain_id = $2 
		  AND dest_chain_id = $3 
		RETURNING seq_num`, seqNumsSet)
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	res, err := o.db.ExecContext(ctx, q, status, sourceChainId.String(), destChainId.String())
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

func (o *orm) ResetExpiredRequests(sourceChainId, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus RequestStatus, toStatus RequestStatus) error {
	q := fmt.Sprintf(`UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE now() > (updated_at + interval '%d seconds') 
			AND source_chain_id = $2
			AND dest_chain_id = $3
			AND status = $4`, expiryTimeoutSeconds)
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	_, err := o.db.ExecContext(ctx, q, toStatus, sourceChainId.String(), destChainId.String(), fromStatus)
	return err
}

// Note requests will only be added in an unstarted status
func (o *orm) SaveRequest(request *Request) error {
	q := `INSERT INTO ccip_requests 
    (seq_num, source_chain_id, dest_chain_id, sender, receiver, data, tokens, amounts, executor, options, raw, status, created_at, updated_at) 
    VALUES (:seq_num, :source_chain_id, :dest_chain_id, :sender, :receiver, :data, :tokens, :amounts, :executor, :options, :raw, 'unstarted', now(), now())
   	ON CONFLICT DO NOTHING `
	stmt, err := o.db.PrepareNamed(q)
	if err != nil {
		return err
	}
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	_, err = stmt.ExecContext(ctx, request)
	return err
}

func (o *orm) RelayReport(seqNum *big.Int) (report RelayReport, err error) {
	q := `SELECT * FROM ccip_relay_reports WHERE min_seq_num <= $1 and max_seq_num >= $1`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	err = o.db.GetContext(ctx, &report, q, seqNum.String())
	return
}

func (o *orm) SaveRelayReport(report RelayReport) error {
	q := `INSERT INTO ccip_relay_reports (root, min_seq_num, max_seq_num, created_at) VALUES ($1, $2, $3, now()) ON CONFLICT DO NOTHING`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	_, err := o.db.ExecContext(ctx, q, report.Root[:], report.MinSeqNum.String(), report.MaxSeqNum.String())
	return err
}
