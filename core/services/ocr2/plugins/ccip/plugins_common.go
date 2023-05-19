package ccip

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/observability"
)

const (
	MaxTokensPerMessage  = 5
	MaxMessagesPerBatch  = 256     // merkle proof bits need to fit in a uint256
	MaxQueryLength       = 0       // empty for both plugins
	MaxObservationLength = 250_000 // plugins's Observation should make sure to cap to this limit
)

var ErrCommitStoreIsDown = errors.New("commitStore is down")

func LoadOnRamp(onRampAddress common.Address, client client.Client) (evm_2_evm_onramp.EVM2EVMOnRampInterface, error) {
	err := ccipconfig.VerifyTypeAndVersion(onRampAddress, client, ccipconfig.EVM2EVMOnRamp)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid onRamp contract")
	}
	return observability.NewObservedEVM2EVMnRamp(onRampAddress, client)
}

func LoadOffRamp(offRampAddress common.Address, client client.Client) (evm_2_evm_offramp.EVM2EVMOffRampInterface, error) {
	err := ccipconfig.VerifyTypeAndVersion(offRampAddress, client, ccipconfig.EVM2EVMOffRamp)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid offRamp contract")
	}
	return observability.NewObservedEVM2EVMOffRamp(offRampAddress, client)
}

func LoadCommitStore(commitStoreAddress common.Address, client client.Client) (commit_store.CommitStoreInterface, error) {
	err := ccipconfig.VerifyTypeAndVersion(commitStoreAddress, client, ccipconfig.CommitStore)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid commitStore contract")
	}
	return observability.NewObservedCommitStore(commitStoreAddress, client)
}

func contiguousReqs(lggr logger.Logger, min, max uint64, seqNrs []uint64) bool {
	if int(max-min+1) != len(seqNrs) {
		return false
	}
	for i, j := min, 0; i <= max && j < len(seqNrs); i, j = i+1, j+1 {
		if seqNrs[j] != i {
			lggr.Errorw("unexpected gap in seq nums", "seq", i)
			return false
		}
	}
	return true
}

func calculateUsdPerUnitGas(sourceGasPrice *big.Int, usdPerFeeCoin *big.Int) *big.Int {
	// (wei / gas) * (usd / eth) * (1 eth / 1e18 wei)  = usd/gas
	tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
	return tmp.Div(tmp, big.NewInt(1e18))
}

// Extracts the hashed leaves from a given set of logs
func leavesFromIntervals(
	lggr logger.Logger,
	seqParser func(logpoller.Log) (uint64, error),
	interval commit_store.CommitStoreInterval,
	hasher hasher.LeafHasherInterface[[32]byte],
	logs []logpoller.Log,
) ([][32]byte, error) {
	var seqNrs []uint64
	for _, log := range logs {
		seqNr, err2 := seqParser(log)
		if err2 != nil {
			return nil, err2
		}
		seqNrs = append(seqNrs, seqNr)
	}
	if !contiguousReqs(lggr, interval.Min, interval.Max, seqNrs) {
		return nil, errors.Errorf("do not have full range [%v, %v] have %v", interval.Min, interval.Max, seqNrs)
	}
	var leaves [][32]byte
	for _, log := range logs {
		hash, err2 := hasher.HashLeaf(log.ToGethLog())
		if err2 != nil {
			return nil, err2
		}
		leaves = append(leaves, hash)
	}

	return leaves, nil
}

// Checks whether the commit store is down by doing an onchain check for
// Paused and AFN status
func isCommitStoreDownNow(ctx context.Context, lggr logger.Logger, commitStore commit_store.CommitStoreInterface) bool {
	paused, err := commitStore.Paused(&bind.CallOpts{Context: ctx})
	if err != nil {
		// Air on side of caution by halting if we cannot read the state?
		lggr.Errorw("Unable to read CommitStore paused", "err", err)
		return true
	}
	healthy, err := commitStore.IsAFNHealthy(&bind.CallOpts{Context: ctx})
	if err != nil {
		lggr.Errorw("Unable to read CommitStore AFN state", "err", err)
		return true
	}
	return paused || !healthy
}

func copyMap[M ~map[K]V, K comparable, V any](m M) M {
	cpy := make(M)
	for k, v := range m {
		cpy[k] = v
	}
	return cpy
}

func max[T constraints.Ordered](first T, rest ...T) T {
	max := first
	for _, v := range rest {
		if v > max {
			max = v
		}
	}
	return max
}

func median(vals []*big.Int) *big.Int {
	valsCopy := make([]*big.Int, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j]) == -1
	})
	return valsCopy[len(valsCopy)/2]
}

// deviation_parts_per_billion = ((x2 - x1) / x1) * 1e9
func deviates(x1, x2 *big.Int, ppb int64) bool {
	// if x1 == 0, deviates if x2 != x1, to avoid the relative division by 0 error
	if x1.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	diff := big.NewInt(0).Sub(x1, x2)
	diff.Mul(diff, big.NewInt(1e9))
	diff.Div(diff, x1)
	return diff.CmpAbs(big.NewInt(ppb)) > 0
}
