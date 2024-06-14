package launcher

import (
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
)

func Test_diff(t *testing.T) {
	type args struct {
		capabilityVersion      string
		capabilityLabelledName string
		oldState               cctypes.RegistryState
		newState               cctypes.RegistryState
	}
	tests := []struct {
		name    string
		args    args
		want    diffResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := diff(tt.args.capabilityVersion, tt.args.capabilityLabelledName, tt.args.oldState, tt.args.newState)
			if (err != nil) != tt.wantErr {
				t.Errorf("diff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareDONs(t *testing.T) {
	type args struct {
		currCCIPDONs map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
		newCCIPDONs  map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
	}
	tests := []struct {
		name        string
		args        args
		wantAdded   map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
		wantRemoved map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
		wantUpdated map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdded, gotRemoved, gotUpdated, err := compareDONs(tt.args.currCCIPDONs, tt.args.newCCIPDONs)
			if (err != nil) != tt.wantErr {
				t.Errorf("compareDONs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAdded, tt.wantAdded) {
				t.Errorf("compareDONs() gotAdded = %v, want %v", gotAdded, tt.wantAdded)
			}
			if !reflect.DeepEqual(gotRemoved, tt.wantRemoved) {
				t.Errorf("compareDONs() gotRemoved = %v, want %v", gotRemoved, tt.wantRemoved)
			}
			if !reflect.DeepEqual(gotUpdated, tt.wantUpdated) {
				t.Errorf("compareDONs() gotUpdated = %v, want %v", gotUpdated, tt.wantUpdated)
			}
		})
	}
}

func Test_filterCCIPDONs(t *testing.T) {
	type args struct {
		ccipCapability keystone_capability_registry.CapabilityRegistryCapability
		state          cctypes.RegistryState
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterCCIPDONs(tt.args.ccipCapability, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterCCIPDONs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterCCIPDONs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCapabilityPresence(t *testing.T) {
	type args struct {
		capabilityVersion      string
		capabilityLabelledName string
		state                  cctypes.RegistryState
	}
	tests := []struct {
		name    string
		args    args
		want    keystone_capability_registry.CapabilityRegistryCapability
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkCapabilityPresence(tt.args.capabilityVersion, tt.args.capabilityLabelledName, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkCapabilityPresence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkCapabilityPresence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashedCapabilityId(t *testing.T) {
	type args struct {
		capabilityVersion      string
		capabilityLabelledName string
	}
	tests := []struct {
		name  string
		args  args
		wantR [32]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := hashedCapabilityId(tt.args.capabilityVersion, tt.args.capabilityLabelledName); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("hashedCapabilityId() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_isMemberOfDON(t *testing.T) {
	type args struct {
		don   keystone_capability_registry.CapabilityRegistryDONInfo
		p2pID p2pkey.KeyV2
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMemberOfDON(tt.args.don, tt.args.p2pID); got != tt.want {
				t.Errorf("isMemberOfDON() = %v, want %v", got, tt.want)
			}
		})
	}
}
