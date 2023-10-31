package contractutil

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
)

// deprecated
func LoadCommitStore(commitStoreAddress common.Address, client client.Client) (commit_store.CommitStoreInterface, semver.Version, error) {
	version, err := ccipconfig.VerifyTypeAndVersion(commitStoreAddress, client, ccipconfig.CommitStore)
	if err != nil {
		return nil, semver.Version{}, errors.Wrap(err, "Invalid commitStore contract")
	}

	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, client)
	return commitStore, version, err
}
