package config

import (
	"strconv"

	"github.com/pkg/errors"
	chainselectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

// GetDestChain returns the destination chain for the given job spec.
func GetDestChain(spec *job.OCR2OracleSpec, chainSet evm.LegacyChainContainer) (evm.Chain, int64, error) {
	chainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return nil, 0, errors.New("chainID must be provided in relay config")
	}
	destChainID := uint64(chainIDInterface.(float64))
	return getChainInChainset(destChainID, chainSet)
	//destChain, err := chainSet.Get(strconv.FormatInt(destChainID, 10))
	//if err != nil {
	//	return nil, 0, errors.Wrap(err, "chain not found in chainset")
	//}
	//return destChain, destChain.ID().Int64(), nil
}

// GetChain returns the chain for the given chain selector.
func GetChain(chainSelector uint64, chainSet evm.LegacyChainContainer) (evm.Chain, int64, error) {
	chainId, err := chainselectors.ChainIdFromSelector(chainSelector)
	if err != nil {
		return nil, 0, err
	}
	return getChainInChainset(chainId, chainSet)
}

func getChainInChainset(chainId uint64, chainSet evm.LegacyChainContainer) (evm.Chain, int64, error) {
	chain, err := chainSet.Get(strconv.FormatUint(chainId, 10))
	if err != nil {
		return nil, 0, errors.Wrap(err, "chain not found in chainset")
	}
	return chain, chain.ID().Int64(), nil
}
