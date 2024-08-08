package deployment

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	deployment2 "github.com/smartcontractkit/ccip/integration-tests/deployment"
	chainsel "github.com/smartcontractkit/chain-selectors"
)

func GenerateAcceptOwnershipProposal(
	e deployment2.Environment,
	chains []uint64,
	state CCIPOnChainState,
) (deployment2.Proposal, error) {
	// TODO: Just onramp as an example
	var ops []deployment2.ManyChainMultiSigOp
	for _, sel := range chains {
		e.Chains[sel].DeployerKey.NoSend = true
		txData, err := state.EvmOnRampsV160[sel].AcceptOwnership(e.Chains[sel].DeployerKey)
		if err != nil {
			return deployment2.Proposal{}, err
		}
		evmID, err := chainsel.ChainIdFromSelector(sel)
		if err != nil {
			return deployment2.Proposal{}, err
		}
		ops = append(ops, deployment2.ManyChainMultiSigOp{
			ChainId:  big.NewInt(int64(evmID)),
			MultiSig: common.Address{},
			Nonce:    big.NewInt(0),
			To:       state.EvmOnRampsV160[sel].Address(),
			Value:    big.NewInt(0),
			Data:     txData.Data(),
		})
	}
	// TODO: Real valid until.
	return deployment2.Proposal{ValidUntil: uint32(time.Now().Unix()), Ops: ops}, nil
}

func ApplyProposal(env deployment2.Environment, p deployment2.Proposal, state CCIPOnChainState) error {
	// TODO
	return nil
}
