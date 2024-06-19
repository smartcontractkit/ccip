package launcher

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	mocktypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types/mocks"
)

func Test_commitExecDeployment_Shutdown(t *testing.T) {
	tests := []struct {
		name        string
		commitBlue  *mocktypes.CCIPOracle
		commitGreen *mocktypes.CCIPOracle
		execBlue    *mocktypes.CCIPOracle
		execGreen   *mocktypes.CCIPOracle
		expect      func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle)
		asserts     func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle)
		wantErr     bool
	}{
		{
			name:        "no errors, blue only",
			commitBlue:  mocktypes.NewCCIPOracle(t),
			commitGreen: nil,
			execBlue:    mocktypes.NewCCIPOracle(t),
			execGreen:   nil,
			expect: func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle) {
				commitBlue.On("Shutdown").Return(nil).Once()
				execBlue.On("Shutdown").Return(nil).Once()
			},
			asserts: func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle) {
				commitBlue.AssertExpectations(t)
				execBlue.AssertExpectations(t)
			},
			wantErr: false,
		},
		{
			name:        "no errors, blue and green",
			commitBlue:  mocktypes.NewCCIPOracle(t),
			commitGreen: mocktypes.NewCCIPOracle(t),
			execBlue:    mocktypes.NewCCIPOracle(t),
			execGreen:   mocktypes.NewCCIPOracle(t),
			expect: func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle) {
				commitBlue.On("Shutdown").Return(nil).Once()
				commitGreen.On("Shutdown").Return(nil).Once()
				execBlue.On("Shutdown").Return(nil).Once()
				execGreen.On("Shutdown").Return(nil).Once()
			},
			asserts: func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle) {
				commitBlue.AssertExpectations(t)
				commitGreen.AssertExpectations(t)
				execBlue.AssertExpectations(t)
				execGreen.AssertExpectations(t)
			},
			wantErr: false,
		},
		{
			name:        "error on commit blue",
			commitBlue:  mocktypes.NewCCIPOracle(t),
			commitGreen: nil,
			execBlue:    mocktypes.NewCCIPOracle(t),
			execGreen:   nil,
			expect: func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle) {
				commitBlue.On("Shutdown").Return(errors.New("failed")).Once()
				execBlue.On("Shutdown").Return(nil).Once()
			},
			asserts: func(t *testing.T, commitBlue, commitGreen, execBlue, execGreen *mocktypes.CCIPOracle) {
				commitBlue.AssertExpectations(t)
				execBlue.AssertExpectations(t)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ccipDeployment{
				commit: blueGreenDeployment{
					blue: tt.commitBlue,
				},
				exec: blueGreenDeployment{
					blue: tt.execBlue,
				},
			}
			if tt.commitGreen != nil {
				c.commit.green = tt.commitGreen
			}
			if tt.execGreen != nil {
				c.exec.green = tt.execGreen
			}
			tt.expect(t, tt.commitBlue, tt.commitGreen, tt.execBlue, tt.execGreen)
			defer tt.asserts(t, tt.commitBlue, tt.commitGreen, tt.execBlue, tt.execGreen)
			err := c.Shutdown()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_commitExecDeployment_HasGreenCommitInstance(t *testing.T) {
	type fields struct {
		commit blueGreenDeployment
		exec   blueGreenDeployment
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "only commit blue is present",
			fields: fields{
				commit: blueGreenDeployment{
					blue: mocktypes.NewCCIPOracle(t),
				},
			},
			want: false,
		},
		{
			name: "both commit blue and green are present",
			fields: fields{
				commit: blueGreenDeployment{
					blue:  mocktypes.NewCCIPOracle(t),
					green: mocktypes.NewCCIPOracle(t),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ccipDeployment{
				commit: tt.fields.commit,
				exec:   tt.fields.exec,
			}
			if got := c.HasGreenCommitInstance(); got != tt.want {
				t.Errorf("commitExecDeployment.HasGreenCommitInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commitExecDeployment_NumExecInstances(t *testing.T) {
	type fields struct {
		commit blueGreenDeployment
		exec   blueGreenDeployment
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "only exec blue is present",
			fields: fields{
				exec: blueGreenDeployment{
					blue: mocktypes.NewCCIPOracle(t),
				},
			},
			want: false,
		},
		{
			name: "both exec blue and green are present",
			fields: fields{
				exec: blueGreenDeployment{
					blue:  mocktypes.NewCCIPOracle(t),
					green: mocktypes.NewCCIPOracle(t),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ccipDeployment{
				commit: tt.fields.commit,
				exec:   tt.fields.exec,
			}
			if got := c.HasGreenExecInstance(); got != tt.want {
				t.Errorf("commitExecDeployment.NumExecInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}
