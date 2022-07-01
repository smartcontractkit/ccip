package ccip

import (
	"context"
	"math/big"
	"testing"

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

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

func TestRelayReportSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded relay report size", prop.ForAll(func(root []byte, min, max uint64) bool {
		var root32 [32]byte
		copy(root32[:], root)
		rep, err := EncodeRelayReport(&blob_verifier.CCIPRelayReport{MerkleRoots: [][32]byte{root32}, Intervals: []blob_verifier.CCIPInterval{{Min: min, Max: max}}})
		require.NoError(t, err)
		return len(rep) <= MaxRelayReportLength
	}, gen.SliceOfN(32, gen.UInt8()), gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}

func TestRelayReportEncoding(t *testing.T) {
	// Set up a user.
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
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
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress,
		native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		}, native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		})
	require.NoError(t, err)
	destChain.Commit()
	_, err = native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
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

	// Deploy blob verifier.
	onRampAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	blobVerifierAddress, _, _, err := blob_verifier_helper.DeployBlobVerifierHelper(
		destUser,         // user
		destChain,        // client
		big.NewInt(1338), // dest chain id
		big.NewInt(1337),
		afnAddress,        // AFN address
		big.NewInt(86400), // max timeout without AFN signal  86400 seconds = one day
		blob_verifier_helper.BlobVerifierInterfaceBlobVerifierConfig{
			OnRamps:          []common.Address{onRampAddress},
			MinSeqNrByOnRamp: []uint64{1},
		},
	)
	require.NoError(t, err)
	blobVerifier, err := blob_verifier_helper.NewBlobVerifierHelper(blobVerifierAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Send a report.
	mctx := merklemulti.NewKeccakCtx()
	tree := merklemulti.NewTree(mctx, [][32]byte{mctx.HashLeaf([]byte{0xaa})})
	root := tree.Root()
	report := blob_verifier.CCIPRelayReport{
		OnRamps:     []common.Address{onRampAddress},
		MerkleRoots: [][32]byte{root},
		Intervals:   []blob_verifier.CCIPInterval{{Min: 1, Max: 10}},
		RootOfRoots: root,
	}
	out, err := EncodeRelayReport(&report)
	require.NoError(t, err)
	decodedReport, err := DecodeRelayReport(out)
	require.NoError(t, err)
	require.Equal(t, &report, decodedReport)

	tx, err := blobVerifier.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	// Ensure root exists.
	ts, err := blobVerifier.GetMerkleRoot(nil, root)
	require.NoError(t, err)
	require.NotEqual(t, ts.String(), "0")
}
