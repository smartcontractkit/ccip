package launcher

import (
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
)

func Test_launcher_Close(t *testing.T) {
	type fields struct {
		capabilityVersion      string
		capabilityLabelledName string
		p2pID                  p2pkey.KeyV2
		capRegistry            cctypes.CapabilityRegistry
		lggr                   logger.Logger
		homeChainReader        cctypes.HomeChainReader
		stopChan               chan struct{}
		regState               cctypes.RegistryState
		oracleCreator          cctypes.OracleCreator
		dons                   map[uint32]*ccipDeployment
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &launcher{
				capabilityVersion:      tt.fields.capabilityVersion,
				capabilityLabelledName: tt.fields.capabilityLabelledName,
				p2pID:                  tt.fields.p2pID,
				capRegistry:            tt.fields.capRegistry,
				lggr:                   tt.fields.lggr,
				homeChainReader:        tt.fields.homeChainReader,
				stopChan:               tt.fields.stopChan,
				regState:               tt.fields.regState,
				oracleCreator:          tt.fields.oracleCreator,
				dons:                   tt.fields.dons,
			}
			if err := l.Close(); (err != nil) != tt.wantErr {
				t.Errorf("launcher.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_launcher_tick(t *testing.T) {
	type fields struct {
		capabilityVersion      string
		capabilityLabelledName string
		p2pID                  p2pkey.KeyV2
		capRegistry            cctypes.CapabilityRegistry
		lggr                   logger.Logger
		homeChainReader        cctypes.HomeChainReader
		stopChan               chan struct{}
		regState               cctypes.RegistryState
		oracleCreator          cctypes.OracleCreator
		dons                   map[uint32]*ccipDeployment
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &launcher{
				capabilityVersion:      tt.fields.capabilityVersion,
				capabilityLabelledName: tt.fields.capabilityLabelledName,
				p2pID:                  tt.fields.p2pID,
				capRegistry:            tt.fields.capRegistry,
				lggr:                   tt.fields.lggr,
				homeChainReader:        tt.fields.homeChainReader,
				stopChan:               tt.fields.stopChan,
				regState:               tt.fields.regState,
				oracleCreator:          tt.fields.oracleCreator,
				dons:                   tt.fields.dons,
			}
			if err := l.tick(); (err != nil) != tt.wantErr {
				t.Errorf("launcher.tick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_launcher_processDiff(t *testing.T) {
	type fields struct {
		capabilityVersion      string
		capabilityLabelledName string
		p2pID                  p2pkey.KeyV2
		capRegistry            cctypes.CapabilityRegistry
		lggr                   logger.Logger
		homeChainReader        cctypes.HomeChainReader
		stopChan               chan struct{}
		regState               cctypes.RegistryState
		oracleCreator          cctypes.OracleCreator
		dons                   map[uint32]*ccipDeployment
	}
	type args struct {
		diff diffResult
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &launcher{
				capabilityVersion:      tt.fields.capabilityVersion,
				capabilityLabelledName: tt.fields.capabilityLabelledName,
				p2pID:                  tt.fields.p2pID,
				capRegistry:            tt.fields.capRegistry,
				lggr:                   tt.fields.lggr,
				homeChainReader:        tt.fields.homeChainReader,
				stopChan:               tt.fields.stopChan,
				regState:               tt.fields.regState,
				oracleCreator:          tt.fields.oracleCreator,
				dons:                   tt.fields.dons,
			}
			if err := l.processDiff(tt.args.diff); (err != nil) != tt.wantErr {
				t.Errorf("launcher.processDiff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_launcher_removeDON(t *testing.T) {
	type fields struct {
		capabilityVersion      string
		capabilityLabelledName string
		p2pID                  p2pkey.KeyV2
		capRegistry            cctypes.CapabilityRegistry
		lggr                   logger.Logger
		homeChainReader        cctypes.HomeChainReader
		stopChan               chan struct{}
		regState               cctypes.RegistryState
		oracleCreator          cctypes.OracleCreator
		dons                   map[uint32]*ccipDeployment
	}
	type args struct {
		id uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &launcher{
				capabilityVersion:      tt.fields.capabilityVersion,
				capabilityLabelledName: tt.fields.capabilityLabelledName,
				p2pID:                  tt.fields.p2pID,
				capRegistry:            tt.fields.capRegistry,
				lggr:                   tt.fields.lggr,
				homeChainReader:        tt.fields.homeChainReader,
				stopChan:               tt.fields.stopChan,
				regState:               tt.fields.regState,
				oracleCreator:          tt.fields.oracleCreator,
				dons:                   tt.fields.dons,
			}
			if err := l.removeDON(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("launcher.removeDON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_launcher_updateDON(t *testing.T) {
	type fields struct {
		capabilityVersion      string
		capabilityLabelledName string
		p2pID                  p2pkey.KeyV2
		capRegistry            cctypes.CapabilityRegistry
		lggr                   logger.Logger
		homeChainReader        cctypes.HomeChainReader
		stopChan               chan struct{}
		regState               cctypes.RegistryState
		oracleCreator          cctypes.OracleCreator
		dons                   map[uint32]*ccipDeployment
	}
	type args struct {
		don keystone_capability_registry.CapabilityRegistryDONInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &launcher{
				capabilityVersion:      tt.fields.capabilityVersion,
				capabilityLabelledName: tt.fields.capabilityLabelledName,
				p2pID:                  tt.fields.p2pID,
				capRegistry:            tt.fields.capRegistry,
				lggr:                   tt.fields.lggr,
				homeChainReader:        tt.fields.homeChainReader,
				stopChan:               tt.fields.stopChan,
				regState:               tt.fields.regState,
				oracleCreator:          tt.fields.oracleCreator,
				dons:                   tt.fields.dons,
			}
			if err := l.updateDON(tt.args.don); (err != nil) != tt.wantErr {
				t.Errorf("launcher.updateDON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_launcher_addDON(t *testing.T) {
	type fields struct {
		capabilityVersion      string
		capabilityLabelledName string
		p2pID                  p2pkey.KeyV2
		capRegistry            cctypes.CapabilityRegistry
		lggr                   logger.Logger
		homeChainReader        cctypes.HomeChainReader
		stopChan               chan struct{}
		regState               cctypes.RegistryState
		oracleCreator          cctypes.OracleCreator
		dons                   map[uint32]*ccipDeployment
	}
	type args struct {
		don keystone_capability_registry.CapabilityRegistryDONInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &launcher{
				capabilityVersion:      tt.fields.capabilityVersion,
				capabilityLabelledName: tt.fields.capabilityLabelledName,
				p2pID:                  tt.fields.p2pID,
				capRegistry:            tt.fields.capRegistry,
				lggr:                   tt.fields.lggr,
				homeChainReader:        tt.fields.homeChainReader,
				stopChan:               tt.fields.stopChan,
				regState:               tt.fields.regState,
				oracleCreator:          tt.fields.oracleCreator,
				dons:                   tt.fields.dons,
			}
			if err := l.addDON(tt.args.don); (err != nil) != tt.wantErr {
				t.Errorf("launcher.addDON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
