package hashlib

import (
	"strconv"

	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

// BytesOfBytesKeccak will compute a keccak256 hash of the provided bytes of bytes slice
func BytesOfBytesKeccak(b [][]byte) ([32]byte, error) {
	if len(b) == 0 {
		return [32]byte{}, nil
	}

	hashes := make([]byte, 0, len(b)*(32+2))
	for i := range b {
		h := utils.Keccak256Fixed(b[i])

		// prepending the array length to prevent collision
		hashes = append(hashes, []byte(strconv.FormatInt(int64(len(b[i])), 10))...)

		hashes = append(hashes, h[:]...)
	}

	return utils.Keccak256Fixed(hashes), nil
}
