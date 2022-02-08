package ccip

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp_helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRelayReportEncoding(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		ethconfig.Defaults.Miner.GasCeil)
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := lock_unlock_pool.DeployLockUnlockPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	_, err = lock_unlock_pool.NewLockUnlockPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destAfn := deployAfn(t, destUser, destChain)

	offRampAddress, _, _, err := single_token_offramp_helper.DeploySingleTokenOffRampHelper(
		destUser,             // user
		destChain,            // client
		big.NewInt(1337),     // source chain id
		big.NewInt(1338),     // dest chain id
		destLinkTokenAddress, // link token address
		destPoolAddress,      // dest pool address
		big.NewInt(1),        // token bucket rate
		big.NewInt(1000),     // token bucket capacity
		destAfn,              // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp_helper.NewSingleTokenOffRampHelper(offRampAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	r, proof := GenerateMerkleProof(2, [][]byte{{0xaa}}, 0)
	rootLocal := GenerateMerkleRoot([]byte{0xaa}, proof)
	require.True(t, bytes.Equal(rootLocal[:], r[:]))
	t.Log(proof.PathForExecute(), proof.path)

	report := single_token_offramp.CCIPRelayReport{
		MerkleRoot:        r,
		MinSequenceNumber: big.NewInt(1),
		MaxSequenceNumber: big.NewInt(10),
	}
	out, err := EncodeRelayReport(&report)
	require.NoError(t, err)
	decodedReport, err := DecodeRelayReport(out)
	require.NoError(t, err)
	require.Equal(t, &report, decodedReport)

	tx, err := offRamp.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	rep, err := offRamp.GetLastReport(nil)
	require.NoError(t, err)
	// Verify it locally
	require.True(t, bytes.Equal(rep.MerkleRoot[:], rootLocal[:]), fmt.Sprintf("Got %v want %v", hexutil.Encode(rootLocal[:]), hexutil.Encode(rep.MerkleRoot[:])))
	exists, err := offRamp.GetMerkleRoot(nil, rep.MerkleRoot)
	require.NoError(t, err)
	require.True(t, exists.Int64() > 0)

	// Verify it onchain
	lh := HashLeaf([]byte{0xaa})
	// Should merely be doing H(lhash, 32 zero bytes) and obtaining the same hash
	root, err := offRamp.GenerateMerkleRoot(nil, proof.PathForExecute(), lh, proof.Index())
	require.NoError(t, err)

	t.Log("verifies", root, "path", proof.PathForExecute(), "Index", proof.Index(), "root", rep.MerkleRoot, "rootlocal", hashInternal(lh, proof.PathForExecute()[0]))
	require.Equal(t, rootLocal, root)
}
