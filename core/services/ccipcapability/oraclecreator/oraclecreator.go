package oraclecreator

import (
	"github.com/google/uuid"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
)

type oracleCreator struct {
	ocrKeyBundles     map[chaintype.ChainType]ocr2key.KeyBundle
	transmitters      map[types.RelayID][]string
	relayers          map[types.RelayID]loop.Relayer
	peerWrapper       *ocrcommon.SingletonPeerWrapper
	externalJobID     uuid.UUID
	jobID             int32
	isNewlyCreatedJob bool
	relayConfigs      map[string]job.JSONConfig
	pluginConfig      job.JSONConfig
	db                ocr3types.Database
}

func New(
	ocrKeyBundles map[chaintype.ChainType]ocr2key.KeyBundle,
	transmitters map[types.RelayID][]string,
	relayers map[types.RelayID]loop.Relayer,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	externalJobID uuid.UUID,
	jobID int32,
	isNewlyCreatedJob bool,
	relayConfigs map[string]job.JSONConfig,
	pluginConfig job.JSONConfig,
	db ocr3types.Database,
) cctypes.OracleCreator {
	return &oracleCreator{
		ocrKeyBundles:     ocrKeyBundles,
		transmitters:      transmitters,
		relayers:          relayers,
		peerWrapper:       peerWrapper,
		externalJobID:     externalJobID,
		jobID:             jobID,
		isNewlyCreatedJob: isNewlyCreatedJob,
		relayConfigs:      relayConfigs,
		pluginConfig:      pluginConfig,
		db:                db,
	}
}

// CreateCommitOracle implements types.OracleCreator.
func (o *oracleCreator) CreateCommitOracle(config cctypes.OCRConfig) (cctypes.CCIPOracle, error) {
	panic("unimplemented")
}

// CreateExecOracle implements types.OracleCreator.
func (o *oracleCreator) CreateExecOracle(config cctypes.OCRConfig) (cctypes.CCIPOracle, error) {
	panic("unimplemented")
}
