package secrets

import (
	"os"
	"strconv"
)

func GetRPC(chainID uint64) string {
	envVariable := "RPC_" + strconv.FormatUint(chainID, 10)
	rpc := os.Getenv(envVariable)
	if rpc != "" {
		return rpc
	}
	panic("RPC not found. Please set the environment variable for chain " + strconv.FormatUint(chainID, 10) + " e.g. RPC_420=https://rpc.420.com")
}
