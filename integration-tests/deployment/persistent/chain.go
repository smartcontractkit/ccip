package persistent

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"reflect"
	"testing"
	"unsafe"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/seth"

	ctf_config "github.com/smartcontractkit/chainlink-testing-framework/config"
	ctf_test_env "github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	seth_utils "github.com/smartcontractkit/chainlink-testing-framework/utils/seth"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/docker/test_env"
	tc "github.com/smartcontractkit/chainlink/integration-tests/testconfig"
)

func NewPersistentChains(t *testing.T, testConfig tc.TestConfig) (map[uint64]deployment.Chain, error) {
	chains := make(map[uint64]deployment.Chain)
	var networks []*ctf_config.EthereumNetworkConfig

	for _, evmNetwork := range testConfig.Network.EVMNetworks {
		chainConfig := ctf_config.GetDefaultChainConfig()
		chainConfig.ChainID = int(evmNetwork.ChainID)

		ethBuilder := ctf_test_env.NewEthereumNetworkBuilder()
		network, err := ethBuilder.
			WithEthereumVersion(ctf_config.EthereumVersion_Eth1).
			WithExecutionLayer(ctf_config.ExecutionLayer_Geth).
			WithEthereumChainConfig(chainConfig).
			Build()

		if err != nil {
			return chains, err
		}

		networks = append(networks, &network.EthereumNetworkConfig)
	}

	env, err := test_env.NewCLTestEnvBuilder().
		WithTestInstance(t).
		WithTestConfig(&testConfig).
		WithPrivateEthereumNetworks(networks).
		WithStandardCleanup().
		Build()
	if err != nil {
		return chains, err
	}

	for _, evmNetwork := range env.EVMNetworks {
		sethClient, err := seth_utils.GetChainClient(testConfig, *evmNetwork)
		if err != nil {
			return chains, err
		}

		asUint := uint64(evmNetwork.ChainID)
		chain, err := buildChain(sethClient, asUint)
		if err != nil {
			return make(map[uint64]deployment.Chain), nil
		}
		chains[asUint] = chain
	}

	return chains, nil
}

func buildChain(sethClient *seth.Client, chainId uint64) (deployment.Chain, error) {
	shouldRetryOnErrFn := func(err error) bool {
		// some retry logic here
		return true
	}

	prepareReplacementTransactionFn := func(sethClient *seth.Client, tx *types.Transaction) (*types.Transaction, error) {
		// some replacement tx creation logic here
		// maybe adjusting base fee aggressively if it's too low for transaction to even be included in the block
		return tx, nil
	}

	sel, err := chainsel.SelectorFromChainId(chainId)
	if err != nil {
		return deployment.Chain{}, err
	}

	return deployment.Chain{
		Selector: sel,
		Client:   sethClient.Client,
		DeployerKey: func() *bind.TransactOpts {
			opts := sethClient.NewTXOpts()
			opts.Nonce = nil
			return opts
		}(),
		Confirm: func(txHash common.Hash) error {
			ctx, cancelFn := context.WithTimeout(context.Background(), sethClient.Cfg.Network.TxnTimeout.Duration())
			tx, _, err := sethClient.Client.TransactionByHash(ctx, txHash)
			cancelFn()
			if err != nil {
				return err
			}
			// we pass `nil` as the second argument, because we do not have transaction submission error here
			// Decode() can also bump gas if transaction takes too much time to confirm
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

			// we pass `nil` as the first argument, because otherwise `Decode()` would wait for transaction to be mined (and we prefer to do that in Confirm() function),
			// and we don't want that, we are only interested in decoding submission error, if any
			_, decodeErr := sethClient.Decode(nil, retryErr)
			return tx, decodeErr
		},
		LoadAbi: func(contract interface{}) error {
			val := reflect.ValueOf(contract).Elem()
			abiField := val.FieldByName("abi")

			if abiField.IsValid() && abiField.CanSet() {
				return errors.New("abi field is not settable")
			} else {
				// Make the field accessible by using reflection
				abiField = reflect.NewAt(abiField.Type(), unsafe.Pointer(abiField.UnsafeAddr())).Elem()
				contractName := fmt.Sprintf("%T", contract)

				switch abiField.Interface().(type) {
				case abi.ABI:
					sethClient.ContractStore.AddABI(contractName, abiField.Interface().(abi.ABI))
				default:
					return fmt.Errorf("abi field is not the expected abi.ABI, but %T", abiField.Interface())
				}
			}

			return nil
		},
	}, nil
}
