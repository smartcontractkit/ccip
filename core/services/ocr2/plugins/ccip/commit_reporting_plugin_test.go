package ccip

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

func TestCommitReportSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded commit report size", prop.ForAll(func(root []byte, min, max uint64) bool {
		var root32 [32]byte
		copy(root32[:], root)
		rep, err := EncodeCommitReport(&commit_store.ICommitStoreCommitReport{MerkleRoot: root32, Interval: commit_store.ICommitStoreInterval{Min: min, Max: max}})
		require.NoError(t, err)
		return len(rep) <= MaxCommitReportLength
	}, gen.SliceOfN(32, gen.UInt8()), gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}

func TestCommitReportEncoding(t *testing.T) {
	// Set up a user.
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destChainId := uint64(1337)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(destChainId)))
	require.NoError(t, err)
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18))}},
		ethconfig.Defaults.Miner.GasCeil)

	// Deploy link token.
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)

	// Deploy link token pool.
	destPoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	_, err = lock_release_token_pool.NewLockReleaseTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)

	// Deploy AFN.
	afnAddress, _, _, err := afn_contract.DeployAFNContract(
		destUser,
		destChain,
		[]common.Address{destUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)

	feeManagerAddress, _, _, err := fee_manager.DeployFeeManager(destUser, destChain, []fee_manager.InternalFeeUpdate{}, []common.Address{}, uint32(time.Hour.Seconds()))
	require.NoError(t, err)

	// Deploy commitStore.
	onRampAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		destUser,  // user
		destChain, // client
		commit_store_helper.ICommitStoreCommitStoreConfig{
			ChainId:       destChainId,
			SourceChainId: 1337,
			OnRamp:        onRampAddress,
			FeeManager:    feeManagerAddress,
		},
		afnAddress, // AFN address
	)
	require.NoError(t, err)
	commitStore, err := commit_store_helper.NewCommitStoreHelper(commitStoreAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	feeManager, err := fee_manager.NewFeeManager(feeManagerAddress, destChain)
	require.NoError(t, err)

	_, err = feeManager.SetFeeUpdater(destUser, commitStoreAddress)
	require.NoError(t, err)
	destChain.Commit()

	// Send a report.
	mctx := hasher.NewKeccakCtx()
	tree, err := merklemulti.NewTree(mctx, [][32]byte{mctx.Hash([]byte{0xaa})})
	require.NoError(t, err)
	report := commit_store.ICommitStoreCommitReport{
		FeeUpdates: []commit_store.InternalFeeUpdate{
			{
				SourceFeeToken:              common.HexToAddress("0x2"),
				DestChainId:                 destChainId,
				FeeTokenBaseUnitsPerUnitGas: big.NewInt(1252352352),
			},
		},
		MerkleRoot: tree.Root(),
		Interval:   commit_store.ICommitStoreInterval{Min: 1, Max: 10},
	}
	out, err := EncodeCommitReport(&report)
	require.NoError(t, err)
	decodedReport, err := DecodeCommitReport(out)
	require.NoError(t, err)
	require.Equal(t, &report, decodedReport)

	tx, err := commitStore.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	// Ensure root exists.
	ts, err := commitStore.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	require.NotEqual(t, ts.String(), "0")
}

func TestCalculateMedianSourceGasPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		gasPrices []int64
		want      *big.Int
	}{
		{"basic", []int64{1, 2, 3, 4, 5, 6, 7}, big.NewInt(4)},
		{"not enough obs", []int64{1, 2, 3, 0, 0, 0, 0}, nil},
		{"round up", []int64{4, 6322364, 6322364, 323, 2, 722}, big.NewInt(722)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateMedianSourceGasPrice(numbersToObservations(tt.gasPrices))
			assert.Equal(t, tt.want, got)
		})
	}
}

func numbersToObservations(gasPrices []int64) (obs []CommitObservation) {
	for _, gasPrice := range gasPrices {
		if gasPrice == 0 {
			obs = append(obs, CommitObservation{
				commit_store.ICommitStoreInterval{},
				make(map[common.Address]*big.Int),
				nil,
			})
		} else {
			obs = append(obs, CommitObservation{
				commit_store.ICommitStoreInterval{},
				make(map[common.Address]*big.Int),
				big.NewInt(gasPrice),
			})
		}
	}
	return obs
}

func TestCalculateIntervalConsensus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                    string
		intervals               []commit_store.ICommitStoreInterval
		f                       int
		nextMinSeqNumForOffRamp uint64
		wantMin                 uint64
		wantMax                 uint64
		wantErr                 bool
	}{
		{"no obs", []commit_store.ICommitStoreInterval{{Min: 0, Max: 0}}, 0, 100, 0, 0, false},
		{"basic", []commit_store.ICommitStoreInterval{
			{Min: 9, Max: 14},
			{Min: 10, Max: 12},
			{Min: 10, Max: 14},
		}, 1, 10, 10, 14, false},
		{"not enough intervals", []commit_store.ICommitStoreInterval{}, 1, 0, 0, 0, true},
		{"wrong next min", []commit_store.ICommitStoreInterval{
			{Min: 9, Max: 14},
			{Min: 10, Max: 12},
			{Min: 10, Max: 14},
		}, 1, 11, 0, 0, true},
		{"min > max", []commit_store.ICommitStoreInterval{
			{Min: 9, Max: 5},
			{Min: 10, Max: 4},
			{Min: 10, Max: 6},
		}, 1, 10, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateIntervalConsensus(tt.intervals, tt.f, func() (uint64, error) { return tt.nextMinSeqNumForOffRamp, nil })
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.wantMin, got.Min)
			assert.Equal(t, tt.wantMax, got.Max)
		})
	}
}
