package launcher

import (
	"math/big"
	"reflect"
	"testing"

	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/registrysyncer"
	"github.com/stretchr/testify/require"
)

func Test_diff(t *testing.T) {
	type args struct {
		capabilityVersion      string
		capabilityLabelledName string
		oldState               registrysyncer.State
		newState               registrysyncer.State
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
		currCCIPDONs map[registrysyncer.DonID]kcr.CapabilitiesRegistryDONInfo
		newCCIPDONs  map[registrysyncer.DonID]kcr.CapabilitiesRegistryDONInfo
	}
	tests := []struct {
		name        string
		args        args
		wantAdded   map[registrysyncer.DonID]kcr.CapabilitiesRegistryDONInfo
		wantRemoved map[registrysyncer.DonID]kcr.CapabilitiesRegistryDONInfo
		wantUpdated map[registrysyncer.DonID]kcr.CapabilitiesRegistryDONInfo
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr, err := compareDONs(tt.args.currCCIPDONs, tt.args.newCCIPDONs)
			if (err != nil) != tt.wantErr {
				t.Errorf("compareDONs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(dr.added, tt.wantAdded) {
				t.Errorf("compareDONs() gotAdded = %v, want %v", dr.added, tt.wantAdded)
			}
			if !reflect.DeepEqual(dr.removed, tt.wantRemoved) {
				t.Errorf("compareDONs() gotRemoved = %v, want %v", dr.removed, tt.wantRemoved)
			}
			if !reflect.DeepEqual(dr.updated, tt.wantUpdated) {
				t.Errorf("compareDONs() gotUpdated = %v, want %v", dr.updated, tt.wantUpdated)
			}
		})
	}
}

func Test_filterCCIPDONs(t *testing.T) {
	type args struct {
		ccipCapability kcr.CapabilitiesRegistryCapabilityInfo
		state          registrysyncer.State
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint32]kcr.CapabilitiesRegistryDONInfo
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
		state                  registrysyncer.State
	}
	tests := []struct {
		name    string
		args    args
		want    kcr.CapabilitiesRegistryCapabilityInfo
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
		don   kcr.CapabilitiesRegistryDONInfo
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

func Test_isMemberOfBootstrapSubcommittee(t *testing.T) {
	var bootstrapKeys [][32]byte
	for i := range [4]struct{}{} {
		bootstrapKeys = append(bootstrapKeys, p2pkey.MustNewV2XXXTestingOnly(big.NewInt(int64(i+1))).PeerID())
	}
	require.True(t, isMemberOfBootstrapSubcommittee(bootstrapKeys, p2pkey.MustNewV2XXXTestingOnly(big.NewInt(1))))
	require.False(t, isMemberOfBootstrapSubcommittee(bootstrapKeys, p2pkey.MustNewV2XXXTestingOnly(big.NewInt(5))))
}
