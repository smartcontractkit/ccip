package evm

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CCIPSubjectUUID(t *testing.T) {
	// We want the function to be
	// (1) an actual function (i.e., deterministic)
	assert.Equal(t, chainToUUID(ccipCommit, big.NewInt(1)), chainToUUID(ccipCommit, big.NewInt(1)))
	// (2) injective (produce different results for different inputs)
	assert.NotEqual(t, chainToUUID(ccipCommit, big.NewInt(1)), chainToUUID(ccipCommit, big.NewInt(2)))
	assert.NotEqual(t, chainToUUID(ccipExec, big.NewInt(1)), chainToUUID(ccipExec, big.NewInt(2)))
	assert.NotEqual(t, chainToUUID(ccipExec, big.NewInt(1)), chainToUUID(ccipCommit, big.NewInt(1)))
	// (3) stable across runs
	assert.Equal(t, "dcb373d4-4bb5-56b3-933f-3e0002bf5064", chainToUUID(ccipCommit, big.NewInt(1)).String())
	assert.Equal(t, "34abf7b1-4103-5294-93b5-0fc5716e11db", chainToUUID(ccipExec, big.NewInt(1)).String())
}
