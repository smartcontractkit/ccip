package ccip_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_helper"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

func TestRelayReportEncoding(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18))}},
		ethconfig.Defaults.Miner.GasCeil)
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
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
	afnAddress, _, _, err := afn_contract.DeployAFNContract(
		destUser,
		destChain,
		[]common.Address{destUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)

	offRampAddress, _, _, err := offramp_helper.DeployOffRampHelper(
		destUser,                               // user
		destChain,                              // client
		big.NewInt(1337),                       // source chain id
		big.NewInt(1338),                       // dest chain id
		[]common.Address{destLinkTokenAddress}, // source tokens, as it doesn't matter for this test we use dest link address
		[]common.Address{destPoolAddress},      // dest pool addresses
		[]common.Address{destPoolAddress},      // Feeds
		afnAddress,                             // AFN address
		big.NewInt(86400),                      // max timeout without AFN signal  86400 seconds = one day
		0,                                      // executionDelaySeconds
		1000,                                   // maxTokensLength
	)
	require.NoError(t, err)
	offRamp, err := offramp_helper.NewOffRampHelper(offRampAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	mctx := merklemulti.NewKeccakCtx()
	tree := merklemulti.NewTree(mctx, [][32]byte{mctx.HashLeaf([]byte{0xaa})})
	root := tree.Root()
	report := offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: 1,
		MaxSequenceNumber: 10,
	}
	out, err := ccip.EncodeRelayReport(&report)
	require.NoError(t, err)
	decodedReport, err := ccip.DecodeRelayReport(out)
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
	require.Equal(t, rep.MerkleRoot, root, fmt.Sprintf("Got %v want %v", hexutil.Encode(root[:]), hexutil.Encode(rep.MerkleRoot[:])))
	exists, err := offRamp.GetMerkleRoot(nil, rep.MerkleRoot)
	require.NoError(t, err)
	require.True(t, exists.Int64() > 0)
}
