package secrets

import "math/big"

var chainIdToRPC = map[string]string{}

func GetRPC(chainID *big.Int) string {
	if rpc, ok := chainIdToRPC[chainID.String()]; ok {
		return rpc
	}
	panic("RPC not found. Please check secrets.go for chainID " + chainID.String())
}
