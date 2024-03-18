package evm

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CCIPSubjectUUID(t *testing.T) {
	assert.Equal(t, chainToUUID(big.NewInt(1)), chainToUUID(big.NewInt(1)))
	assert.NotEqual(t, chainToUUID(big.NewInt(1)), chainToUUID(big.NewInt(2)))
}
