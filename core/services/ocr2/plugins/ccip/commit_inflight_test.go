package ccip

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestCommitInflight(t *testing.T) {
	lggr := logger.TestLogger(t)
	c := newInflightCommitReportsContainer(time.Hour)

	// Initially should be empty
	assert.Nil(t, c.getLatestInflightGasPriceUpdate())
	assert.Equal(t, uint64(0), c.maxInflightSeqNr())

	// Add a single report inflight
	root1 := utils.Keccak256Fixed(hexutil.MustDecode("0xaa"))
	require.NoError(t, c.add(lggr, commit_store.CommitStoreCommitReport{Interval: commit_store.CommitStoreInterval{Min: 1, Max: 2}, MerkleRoot: root1}))
	assert.Nil(t, c.getLatestInflightGasPriceUpdate())
	assert.Equal(t, uint64(2), c.maxInflightSeqNr())

	// Add another price report
	root2 := utils.Keccak256Fixed(hexutil.MustDecode("0xab"))
	require.NoError(t, c.add(lggr, commit_store.CommitStoreCommitReport{Interval: commit_store.CommitStoreInterval{Min: 3, Max: 4}, MerkleRoot: root2}))
	assert.Nil(t, c.getLatestInflightGasPriceUpdate())
	assert.Equal(t, uint64(4), c.maxInflightSeqNr())

	// Add gas price updates
	require.NoError(t, c.add(lggr, commit_store.CommitStoreCommitReport{PriceUpdates: commit_store.InternalPriceUpdates{
		DestChainSelector: uint64(1),
		UsdPerUnitGas:     big.NewInt(1),
	}}))
	latest := c.getLatestInflightGasPriceUpdate()
	assert.Equal(t, big.NewInt(1), latest.value)
	assert.Equal(t, uint64(4), c.maxInflightSeqNr())

	// Add a token price update
	token := common.HexToAddress("0xa")
	require.NoError(t, c.add(lggr, commit_store.CommitStoreCommitReport{PriceUpdates: commit_store.InternalPriceUpdates{
		TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{
			{
				SourceToken: token,
				UsdPerToken: big.NewInt(10),
			},
		},
	}}))
	// Apply cache price to existing
	latestInflightTokenPriceUpdates := c.latestInflightTokenPriceUpdates()
	require.Equal(t, len(latestInflightTokenPriceUpdates), 1)
	assert.Equal(t, big.NewInt(10), latestInflightTokenPriceUpdates[token].value)
}
