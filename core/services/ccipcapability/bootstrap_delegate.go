package ccipcapability

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type BootstrapDelegate struct {
	lggr            logger.Logger
	registrarConfig plugins.RegistrarConfig
	pipelineRunner  pipeline.Runner
	relayGetter     RelayGetter
	capRegistry     cctypes.CapabilityRegistry
	keystore        keystore.Master
	ds              sqlutil.DataSource
	peerWrapper     *ocrcommon.SingletonPeerWrapper

	isNewlyCreatedJob bool
}

func NewBootstrapDelegate(
	lggr logger.Logger,
	registrarConfig plugins.RegistrarConfig,
	pipelineRunner pipeline.Runner,
	relayGetter RelayGetter,
	registrySyncer cctypes.CapabilityRegistry,
	keystore keystore.Master,
	ds sqlutil.DataSource,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
) *BootstrapDelegate {
	return &BootstrapDelegate{
		lggr:            lggr,
		registrarConfig: registrarConfig,
		pipelineRunner:  pipelineRunner,
		relayGetter:     relayGetter,
		capRegistry:     registrySyncer,
		ds:              ds,
		keystore:        keystore,
		peerWrapper:     peerWrapper,
	}
}

func (d *BootstrapDelegate) JobType() job.Type {
	return job.CCIPBootstrap
}

func (d *BootstrapDelegate) BeforeJobCreated(job.Job) {
	// This is only called first time the job is created
	d.isNewlyCreatedJob = true
}

func (d *BootstrapDelegate) ServicesForSpec(ctx context.Context, spec job.Job) (services []job.ServiceCtx, err error) {
	// TODO: implement.
	return []job.ServiceCtx{}, nil
}

func (d *BootstrapDelegate) AfterJobCreated(spec job.Job) {}

func (d *BootstrapDelegate) BeforeJobDeleted(spec job.Job) {}

func (d *BootstrapDelegate) OnDeleteJob(ctx context.Context, spec job.Job) error {
	// TODO: shut down needed services?
	return nil
}
