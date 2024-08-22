package executable

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTransactionCounts(t *testing.T) {
	transactions := []ChainOperation{
		{ChainIdentifier: "1"},
		{ChainIdentifier: "1"},
		{ChainIdentifier: "2"},
	}

	expected := map[string]uint64{
		"1": 2,
		"2": 1,
	}

	result := calculateTransactionCounts(transactions)
	assert.Equal(t, expected, result)
}

func TestBuildRootMetadatas_Success(t *testing.T) {
	chainMetadata := map[string]ExecutableMCMSChainMetadata{
		"1": {MCMAddress: common.HexToAddress("0x1"), NonceOffset: 0},
		"2": {MCMAddress: common.HexToAddress("0x2"), NonceOffset: 1},
	}
	txCounts := map[string]uint64{
		"1": 2,
		"2": 1,
	}
	currentOpCounts := map[string]big.Int{
		"1": *big.NewInt(0),
		"2": *big.NewInt(2),
	}

	expected := map[string]gethwrappers.ManyChainMultiSigRootMetadata{
		"1": {
			ChainId:              big.NewInt(1),
			MultiSig:             common.HexToAddress("0x1"),
			PreOpCount:           big.NewInt(0),
			PostOpCount:          big.NewInt(2),
			OverridePreviousRoot: true,
		},
		"2": {
			ChainId:              big.NewInt(2),
			MultiSig:             common.HexToAddress("0x2"),
			PreOpCount:           big.NewInt(3),
			PostOpCount:          big.NewInt(4),
			OverridePreviousRoot: true,
		},
	}

	result, err := buildRootMetadatas(chainMetadata, txCounts, currentOpCounts, true)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestBuildRootMetadatas_InvalidChainID(t *testing.T) {
	chainMetadata := map[string]ExecutableMCMSChainMetadata{
		"invalid": {MCMAddress: common.HexToAddress("0x1"), NonceOffset: 0},
	}
	txCounts := map[string]uint64{
		"invalid": 1,
	}
	currentOpCounts := map[string]big.Int{
		"invalid": *big.NewInt(0),
	}

	result, err := buildRootMetadatas(chainMetadata, txCounts, currentOpCounts, true)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.IsType(t, &errors.ErrInvalidChainID{}, err)
}

func TestBuildOperations(t *testing.T) {
	transactions := []ChainOperation{
		{ChainIdentifier: "1", Operation: Operation{
			To: common.HexToAddress("0x1"), Data: "0x", Value: 1,
		}},
		{ChainIdentifier: "1", Operation: Operation{
			To: common.HexToAddress("0x2"), Data: "0x", Value: 2,
		}},
		{ChainIdentifier: "2", Operation: Operation{
			To: common.HexToAddress("0x3"), Data: "0x", Value: 3,
		}},
	}
	rootMetadatas := map[string]gethwrappers.ManyChainMultiSigRootMetadata{
		"1": {
			ChainId:    big.NewInt(1),
			MultiSig:   common.HexToAddress("0x1"),
			PreOpCount: big.NewInt(0),
		},
		"2": {
			ChainId:    big.NewInt(2),
			MultiSig:   common.HexToAddress("0x2"),
			PreOpCount: big.NewInt(0),
		},
	}
	txCounts := map[string]uint64{
		"1": 2,
		"2": 1,
	}

	expected := map[string][]gethwrappers.ManyChainMultiSigOp{
		"1": {
			{
				ChainId:  big.NewInt(1),
				MultiSig: common.HexToAddress("0x1"),
				Nonce:    big.NewInt(0),
				To:       common.HexToAddress("0x1"),
				Data:     common.FromHex("0x"),
				Value:    big.NewInt(1),
			},
			{
				ChainId:  big.NewInt(1),
				MultiSig: common.HexToAddress("0x1"),
				Nonce:    big.NewInt(1),
				To:       common.HexToAddress("0x2"),
				Data:     common.FromHex("0x"),
				Value:    big.NewInt(2),
			},
		},
		"2": {
			{
				ChainId:  big.NewInt(2),
				MultiSig: common.HexToAddress("0x2"),
				Nonce:    big.NewInt(0),
				To:       common.HexToAddress("0x3"),
				Data:     common.FromHex("0x"),
				Value:    big.NewInt(3),
			},
		},
	}

	result, err := buildOperations(transactions, rootMetadatas, txCounts)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSortedChainIdentifiers(t *testing.T) {
	chainMetadata := map[string]ExecutableMCMSChainMetadata{
		"2": {},
		"1": {},
		"3": {},
	}

	expected := []string{"1", "2", "3"}

	result := sortedChainIdentifiers(chainMetadata)
	assert.Equal(t, expected, result)
}

func TestBuildMerkleTree(t *testing.T) {
	chainIdentifiers := []string{"1", "2"}
	ops := map[string][]gethwrappers.ManyChainMultiSigOp{
		"1": {
			{
				ChainId:  big.NewInt(1),
				MultiSig: common.HexToAddress("0x1"),
				Nonce:    big.NewInt(0),
				To:       common.HexToAddress("0x1"),
				Data:     common.FromHex("0x"),
				Value:    big.NewInt(1),
			},
		},
		"2": {
			{
				ChainId:  big.NewInt(2),
				MultiSig: common.HexToAddress("0x2"),
				Nonce:    big.NewInt(0),
				To:       common.HexToAddress("0x2"),
				Data:     common.FromHex("0x"),
				Value:    big.NewInt(2),
			},
		},
	}
	rootMetadatas := map[string]gethwrappers.ManyChainMultiSigRootMetadata{
		"1": {
			ChainId:              big.NewInt(1),
			MultiSig:             common.HexToAddress("0x1"),
			PreOpCount:           big.NewInt(0),
			PostOpCount:          big.NewInt(1),
			OverridePreviousRoot: false,
		},
		"2": {
			ChainId:              big.NewInt(2),
			MultiSig:             common.HexToAddress("0x2"),
			PreOpCount:           big.NewInt(0),
			PostOpCount:          big.NewInt(1),
			OverridePreviousRoot: false,
		},
	}

	tree, err := buildMerkleTree(chainIdentifiers, rootMetadatas, ops)
	assert.NoError(t, err)
	assert.NotNil(t, tree)
	assert.NotEmpty(t, tree.Root)
}

func TestMetadataEncoder(t *testing.T) {
	rootMetadata := gethwrappers.ManyChainMultiSigRootMetadata{
		ChainId:              big.NewInt(1),
		MultiSig:             common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		PreOpCount:           big.NewInt(0),
		PostOpCount:          big.NewInt(1),
		OverridePreviousRoot: true,
	}

	hash, err := metadataEncoder(rootMetadata)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToHash("0x69b33f7bd0adfd9d8fc5381091e2970ff92e4fc425986a6d5b92f5e0b66f9e0b"), hash)
}

func TestTxEncoder(t *testing.T) {
	op := gethwrappers.ManyChainMultiSigOp{
		ChainId:  big.NewInt(1),
		MultiSig: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		Nonce:    big.NewInt(1),
		To:       common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef"),
		Value:    big.NewInt(1000),
		Data:     []byte("data"),
	}

	hash, err := txEncoder(op)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToHash("0x7ee35d1bbff5302395a2ca0b7485ac16c17160798d1ac1bd96f72fa2995b62ff"), hash)
}
