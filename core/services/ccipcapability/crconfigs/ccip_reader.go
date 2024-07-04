package crconfigs

import (
	"encoding/json"

	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

func MustCCIPReaderConfig() []byte {
	rawConfig := CCIPReaderConfigRaw()
	encoded, err := json.Marshal(rawConfig)
	if err != nil {
		panic(err)
	}

	return encoded
}

func CCIPReaderConfigRaw() evmrelaytypes.ChainReaderConfig {
	// TODO: Implement
	return evmrelaytypes.ChainReaderConfig{}
}
