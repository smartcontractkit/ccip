package v1_3_0

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var _ ccipdata.CommitStoreReader = &CommitStore{}

type CommitStore struct {
	*v1_2_0.CommitStore
}

// Do not change the JSON format of this struct without consulting with
// the RDD people first.
type CommitOffchainConfig struct {
	GasPriceHeartBeat        models.Duration
	DAGasPriceDeviationPPB   uint32
	ExecGasPriceDeviationPPB uint32
	TokenPriceHeartBeat      models.Duration
	TokenPriceDeviationPPB   uint32
	SourceMaxGasPrice        uint64
	InflightCacheExpiry      models.Duration
}

func (c CommitOffchainConfig) Validate() error {
	if c.GasPriceHeartBeat.Duration() == 0 {
		return errors.New("must set GasPriceHeartBeat")
	}
	if c.ExecGasPriceDeviationPPB == 0 {
		return errors.New("must set ExecGasPriceDeviationPPB")
	}
	if c.TokenPriceHeartBeat.Duration() == 0 {
		return errors.New("must set TokenPriceHeartBeat")
	}
	if c.TokenPriceDeviationPPB == 0 {
		return errors.New("must set TokenPriceDeviationPPB")
	}
	if c.SourceMaxGasPrice == 0 {
		return errors.New("must set SourceMaxGasPrice")
	}
	if c.InflightCacheExpiry.Duration() == 0 {
		return errors.New("must set InflightCacheExpiry")
	}
	// DAGasPriceDeviationPPB is not validated because it can be 0 on non-rollups

	return nil
}

func NewCommitStore(lggr logger.Logger, addr common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator) (*CommitStore, error) {
	commitStoreV120, err := v1_2_0.NewCommitStore(lggr, addr, ec, lp, estimator)
	if err != nil {
		return nil, err
	}

	return &CommitStore{
		CommitStore: commitStoreV120,
	}, nil
}
