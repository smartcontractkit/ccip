package ccip

import (
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateCCIPSpec(t *testing.T) {
	var tt = []struct {
		name         string
		toml         string
		setGlobalCfg func(t *testing.T, c *configtest.TestGeneralConfig)
		assertion    func(t *testing.T, os job.Job, err error)
	}{
		{
			name: "decodes valid ccip spec toml",
			toml: `
type               = "ccip-relay"
schemaVersion      = 1
contractAddress    = "0x613a38AC1659769640aaE063C651F48E0250454C"
p2pPeerID          = "12D3KooWHfYFQ8hGttAYbMCevQVESEQhzJAqFZokMVtom8bNxwGq"
p2pBootstrapPeers  = [
"/dns4/chain.link/tcp/1234/p2p/16Uiu2HAm58SP7UL8zsnpeuwHfytLocaqgnyaYKP8wu7qRdrixLju",
]
keyBundleID        = "73e8966a78ca09bb912e9565cfb79fbe8a6048fab1f0cf49b18047c3895e0447"
monitoringEndpoint = "chain.link:4321"
transmitterAddress = "0xF67D0290337bca0847005C7ffD1BC75BA9AAE6e4"
observationTimeout = "10s"
observationSource = """
ds1          [type=bridge name=voter_turnout];
ds1_parse    [type=jsonparse path="one,two"];
ds1_multiply [type=multiply times=1.23];
ds1 -> ds1_parse -> ds1_multiply -> answer1;
answer1      [type=median Index=0];
"""
`,
			assertion: func(t *testing.T, os job.Job, err error) {
				require.NoError(t, err)
				assert.Equal(t, 1, int(os.SchemaVersion))
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := ValidatedCCIPSpec(tc.toml)
			tc.assertion(t, s, err)
		})
	}
}

func TestValidateCCIPBootstrapSpec(t *testing.T) {
	var tt = []struct {
		name         string
		toml         string
		setGlobalCfg func(t *testing.T, c *configtest.TestGeneralConfig)
		assertion    func(t *testing.T, os job.Job, err error)
	}{
		{
			name: "decodes valid ccip bootstrap spec toml",
			toml: `
type               = "ccip-bootstrap"
schemaVersion      = 1
contractAddress    = "0x613a38AC1659769640aaE063C651F48E0250454C"
monitoringEndpoint = "chain.link:4321"
`,
			assertion: func(t *testing.T, os job.Job, err error) {
				require.NoError(t, err)
				assert.Equal(t, 1, int(os.SchemaVersion))
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := ValidatedCCIPBootstrapSpec(tc.toml)
			tc.assertion(t, s, err)
		})
	}
}
