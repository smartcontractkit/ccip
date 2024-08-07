package deployment

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/v2/core/deployment"
)

func GenerateAcceptOwnershipProposal(
	e deployment.Environment,
	chains []uint64,
	state CCIPOnChainState,
) (deployment.Proposal, error) {
	// TODO: Just onramp as an example
	var ops []deployment.ManyChainMultiSigOp
	for _, sel := range chains {
		e.Chains[sel].DeployerKey.NoSend = true
		txData, err := state.EvmOnRampsV160[sel].AcceptOwnership(e.Chains[sel].DeployerKey)
		if err != nil {
			return deployment.Proposal{}, err
		}
		evmID, err := chainsel.ChainIdFromSelector(sel)
		if err != nil {
			return deployment.Proposal{}, err
		}
		ops = append(ops, deployment.ManyChainMultiSigOp{
			ChainId:  big.NewInt(int64(evmID)),
			MultiSig: common.Address{},
			Nonce:    big.NewInt(0),
			To:       state.EvmOnRampsV160[sel].Address(),
			Value:    big.NewInt(0),
			Data:     txData.Data(),
		})
	}
	// TODO: Real valid until.
	return deployment.Proposal{ValidUntil: uint32(time.Now().Unix()), Ops: ops}, nil
}

func ApplyProposal(env deployment.Environment, p deployment.Proposal, state CCIPOnChainState) error {
	// TODO
	return nil
}
