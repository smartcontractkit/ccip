package config

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

// GetDestChain returns the destination chain for the given job spec.
func GetDestChain(spec *job.OCR2OracleSpec, chainSet evm.LegacyChainContainer) (evm.Chain, int64, error) {
	chainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return *new(evm.Chain), 0, errors.New("chainID must be provided in relay config")
	}
	destChainID := int64(chainIDInterface.(float64))
	destChain, err := chainSet.Get(strconv.FormatInt(destChainID, 10))
	if err != nil {
		return *new(evm.Chain), 0, errors.Wrap(err, "Chain not found in chainset")
	}
	return destChain, destChain.ID().Int64(), nil
}
