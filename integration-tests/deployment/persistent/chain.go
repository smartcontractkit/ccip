package persistent

import (
	"context"
	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/ccip/integration-tests/deployment"
	"github.com/smartcontractkit/seth"
)

func NewPersistentChain(selector uint64) (deployment.Chain, error) {
	cfg, err := seth.ReadConfig()
	if err != nil {
		return deployment.Chain{}, err
	}
	sethClient, err := seth.NewClientWithConfig(cfg)
	if err != nil {
		return deployment.Chain{}, err
	}

	shouldRetryOnErrFn := func(err error) bool {
		// some more complex logic here
		return true
	}

	prepareReplacementTransactionFn := func(sethClient *seth.Client, tx *types.Transaction) (*types.Transaction, error) {
		// some more complex logic here
		return tx, nil
	}

	return deployment.Chain{
		Selector:    selector,
		Client:      sethClient.Client,
		DeployerKey: sethClient.NewTXOpts(),
		Confirm: func(txHash common.Hash) error {
			ctx, cancelFn := context.WithTimeout(context.Background(), sethClient.Cfg.Network.TxnTimeout.Duration())
			tx, _, err := sethClient.Client.TransactionByHash(ctx, txHash)
			cancelFn()
			if err != nil {
				return err
			}
			_, decodedErr := sethClient.Decode(tx, nil)
			if decodedErr != nil {
				return decodedErr
			}
			return nil
		},
		RetrySubmit: func(tx *types.Transaction, err error) (*types.Transaction, error) {
			if err == nil {
				return tx, nil
			}

			retryErr := retry.Do(
				func() error {
					ctx, cancel := context.WithTimeout(context.Background(), sethClient.Cfg.Network.TxnTimeout.Duration())
					defer cancel()

					return sethClient.Client.SendTransaction(ctx, tx)
				}, retry.OnRetry(func(i uint, retryErr error) {
					replacementTx, replacementErr := prepareReplacementTransactionFn(sethClient, tx)
					if replacementErr != nil {
						return
					}
					tx = replacementTx
				}),
				retry.DelayType(retry.FixedDelay),
				retry.Attempts(10),
				retry.RetryIf(shouldRetryOnErrFn),
			)

			_, decodeErr := sethClient.Decode(nil, retryErr)
			return tx, decodeErr
		},
	}, nil
}
