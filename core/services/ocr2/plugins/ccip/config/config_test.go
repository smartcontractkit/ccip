package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	c := RelayPluginConfig{}
	require.NoError(t, json.Unmarshal([]byte(`{"sourceChainID":1}`), &c))
	t.Log(c)
}
