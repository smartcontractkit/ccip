package ccip_test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestORM(t *testing.T) {
	// Use a real db so we can do timestamp testing.
	_, db := heavyweight.FullTestDB(t, "orm_test", true, false)
	orm := ccip.NewORM(db)
	source := big.NewInt(1)
	dest := big.NewInt(2)

	// Check we can read/write requests.
	req := ccip.Request{
		SourceChainID: source.String(),
		DestChainID:   dest.String(),
		SeqNum:        *utils.NewBigI(10),
		Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Receiver:      common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Data:          []byte("hello"),
		Tokens:        pq.StringArray{},
		Amounts:       pq.StringArray{},
		Options:       []byte{},
	}
	err := orm.SaveRequest(&req)
	require.NoError(t, err)
	reqRead, err := orm.Requests(source, dest, req.SeqNum.ToInt(), req.SeqNum.ToInt(), "", nil, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(reqRead))
	assert.True(t, reqRead[0].UpdatedAt != time.Time{})
	assert.True(t, reqRead[0].CreatedAt != time.Time{})
	assert.Equal(t, req.Data, reqRead[0].Data)

	// Check we can update the request status.
	err = orm.UpdateRequestStatus(source, dest, req.SeqNum.ToInt(), req.SeqNum.ToInt(), ccip.RequestStatusRelayPending)
	require.NoError(t, err)
	// Updating an non-existent reqID should error
	err = orm.UpdateRequestStatus(source, dest, big.NewInt(1337), big.NewInt(1337), ccip.RequestStatusUnstarted)
	require.Error(t, err)
	reqReadAfterUpdate, err := orm.Requests(source, dest, req.SeqNum.ToInt(), req.SeqNum.ToInt(), "", nil, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(reqReadAfterUpdate))
	assert.Equal(t, ccip.RequestStatusRelayPending, reqReadAfterUpdate[0].Status)
	assert.True(t, reqReadAfterUpdate[0].UpdatedAt.After(reqRead[0].UpdatedAt), fmt.Sprintf("before %v after %v", reqRead[0].UpdatedAt, reqReadAfterUpdate[0].UpdatedAt))
	assert.Equal(t, reqReadAfterUpdate[0].CreatedAt, reqRead[0].CreatedAt)

	// Check we can read/write relay reports.
	var aroot = [32]byte{0x01}
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{
		Root:      aroot[:],
		MinSeqNum: *utils.NewBig(big.NewInt(1)),
		MaxSeqNum: *utils.NewBig(big.NewInt(2)),
	}))
	r, err := orm.RelayReport(big.NewInt(1))
	require.NoError(t, err)
	assert.Equal(t, byte(0x01), r.Root[0])
	require.NoError(t, err)

	// Check we can filter by status and executor with multiple requests present.
	executor := common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB5")
	reqForOracleExecution := ccip.Request{
		SourceChainID: source.String(),
		DestChainID:   dest.String(),
		SeqNum:        *utils.NewBigI(11),
		Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Receiver:      common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Data:          []byte("hello"),
		Tokens:        pq.StringArray{},
		Amounts:       pq.StringArray{},
		Executor:      executor,
		Options:       []byte{},
	}
	require.NoError(t, orm.SaveRequest(&reqForOracleExecution))
	require.NoError(t, orm.UpdateRequestStatus(source, dest, big.NewInt(11), big.NewInt(11), ccip.RequestStatusRelayConfirmed))
	reqsForOracle, err := orm.Requests(source, dest, nil, nil, ccip.RequestStatusRelayConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqsForOracle, 1)
	reqsForOracle, err = orm.Requests(source, dest, nil, nil, ccip.RequestStatusRelayConfirmed, &executor, nil)
	require.NoError(t, err)
	require.Len(t, reqsForOracle, 1)

	// Check we can update the status with specific seq nums, as opposed to a range.
	reqsBefore, err := orm.Requests(source, dest, big.NewInt(10), big.NewInt(11), ccip.RequestStatusRelayConfirmed, nil, nil)
	require.NoError(t, err)
	require.NoError(t, orm.UpdateRequestSetStatus(source, dest, []*big.Int{big.NewInt(10), big.NewInt(11)}, ccip.RequestStatusExecutionConfirmed))
	reqs, err := orm.Requests(source, dest, nil, nil, ccip.RequestStatusExecutionConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqs, 2)
	assert.True(t, reqs[0].UpdatedAt.After(reqsBefore[0].UpdatedAt), fmt.Sprintf("before %v after %v", reqRead[0].UpdatedAt, reqReadAfterUpdate[0].UpdatedAt))

	// Check that we can reset the status of expired requests.
	res, err := db.Exec(`UPDATE ccip_requests SET updated_at = $1`, time.Now().Add(-2*time.Second))
	require.NoError(t, err)
	n, err := res.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(2), n)
	// Now they should be recognized as being 1s old, so we can reset them with a timeout of 1s.
	require.NoError(t, orm.ResetExpiredRequests(source, dest, 1, ccip.RequestStatusExecutionConfirmed, ccip.RequestStatusRelayConfirmed))
	// Should all be relay confirmed now.
	reqs, err = orm.Requests(source, dest, nil, nil, ccip.RequestStatusRelayConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqs, 2)
}
