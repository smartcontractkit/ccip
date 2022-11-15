package ccip

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ocrConfigHelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"golang.org/x/crypto/curve25519"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/toll_sender_dapp"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
)

// CCIPContractsDeployer provides the implementations for deploying CCIP ETH contracts
type CCIPContractsDeployer struct {
	evmClient   *blockchain.EthereumClient
	EthDeployer *contracts.EthereumContractDeployer
}

// NewCCIPContractsDeployer returns an instance of a contract deployer for CCIP
func NewCCIPContractsDeployer(bcClient blockchain.EVMClient) (*CCIPContractsDeployer, error) {
	return &CCIPContractsDeployer{
		evmClient:   bcClient.Get().(*blockchain.EthereumClient),
		EthDeployer: contracts.NewEthereumContractDeployer(bcClient),
	}, nil
}

func (e *CCIPContractsDeployer) DeployLinkTokenContract() (contracts.LinkToken, error) {
	return e.EthDeployer.DeployLinkTokenContract()
}

func (e *CCIPContractsDeployer) DeployNativeTokenPoolContract(linkAddr string) (
	*NativeTokenPool,
	error,
) {
	log.Debug().Str("token", linkAddr).Msg("Deploying native token pool")
	token := common.HexToAddress(linkAddr)
	address, _, instance, err := e.evmClient.DeployContract("Native Token Pool", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return native_token_pool.DeployNativeTokenPool(auth, backend, token)
	})

	if err != nil {
		return nil, err
	}
	return &NativeTokenPool{
		client:     e.evmClient,
		instance:   instance.(*native_token_pool.NativeTokenPool),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployAFNContract() (*AFN, error) {
	address, _, instance, err := e.evmClient.DeployContract("Mock AFN Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return mock_afn_contract.DeployMockAFNContract(auth, backend)
	})
	return &AFN{
		client:     e.evmClient,
		instance:   instance.(*mock_afn_contract.MockAFNContract),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployCommitStore(
	sourceChainId, destChainId *big.Int,
	afn common.Address,
	bConfig commit_store.CommitStoreInterfaceCommitStoreConfig,
) (
	*CommitStore,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("CommitStore Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return commit_store.DeployCommitStore(auth, backend, destChainId, sourceChainId, afn, bConfig)
	})
	return &CommitStore{
		client:     e.evmClient,
		instance:   instance.(*commit_store.CommitStore),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeploySimpleMessageReceiver() (
	*MessageReceiver,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("SimpleMessageReceiver Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return simple_message_receiver.DeploySimpleMessageReceiver(auth, backend)
	})
	return &MessageReceiver{
		client:     e.evmClient,
		instance:   instance.(*simple_message_receiver.SimpleMessageReceiver),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployTollOffRamp(
	sourceChainId, destChainId *big.Int,
	commitStore, onRamp, afn common.Address,
	sourceToken, pools []common.Address,
	opts RateLimiterConfig) (
	*TollOffRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("Toll OffRamp Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return any_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
			auth, backend, sourceChainId, destChainId,
			any_2_evm_toll_offramp.BaseOffRampInterfaceOffRampConfig{
				OnRampAddress:         onRamp,
				ExecutionDelaySeconds: 0,
				MaxDataSize:           1e12,
				MaxTokensLength:       15,
			},
			commitStore,
			afn,
			sourceToken,
			pools,
			any_2_evm_toll_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
				Rate:     opts.Rate,
				Capacity: opts.Capacity,
			},
			auth.From)
	})
	return &TollOffRamp{
		client:     e.evmClient,
		instance:   instance.(*any_2_evm_toll_offramp.EVM2EVMTollOffRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployTollOffRampRouter(
	offRamps []common.Address) (
	*TollOffRampRouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("Toll OffRampRouter Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(auth, backend, offRamps)
	})
	return &TollOffRampRouter{
		client:     e.evmClient,
		instance:   instance.(*any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeploySubOffRampRouter(
	offRamps []common.Address, feeToken common.Address) (
	*SubOffRampRouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("Sub OffRampRouter Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return any_2_evm_subscription_offramp_router.DeployAny2EVMSubscriptionOffRampRouter(auth, backend, offRamps,
			any_2_evm_subscription_offramp_router.SubscriptionInterfaceSubscriptionConfig{
				SetSubscriptionSenderDelay: 0,
				WithdrawalDelay:            0,
				FeeToken:                   feeToken,
			})
	})
	return &SubOffRampRouter{
		client:     e.evmClient,
		instance:   instance.(*any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeploySubOffRamp(
	sourceChainId, destChainId *big.Int,
	commitStore, onRamp, afn common.Address,
	sourceToken, pools []common.Address,
	opts RateLimiterConfig,
	offRampConfig any_2_evm_subscription_offramp.BaseOffRampInterfaceOffRampConfig,
) (*SubOffRamp, error) {
	offRampConfig.OnRampAddress = onRamp
	address, _, instance, err := e.evmClient.DeployContract("Sub OffRamp Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return any_2_evm_subscription_offramp.DeployEVM2EVMSubscriptionOffRamp(
			auth, backend, sourceChainId, destChainId,
			offRampConfig,
			commitStore,
			afn,
			sourceToken,
			pools,
			any_2_evm_subscription_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
				Rate:     opts.Rate,
				Capacity: opts.Capacity,
			},
			auth.From)
	})
	return &SubOffRamp{
		client:     e.evmClient,
		instance:   instance.(*any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployReceiverDapp(toRevert bool) (
	*ReceiverDapp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("ReceiverDapp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(auth, backend, toRevert)
	})
	return &ReceiverDapp{
		client:     e.evmClient,
		instance:   instance.(*maybe_revert_message_receiver.MaybeRevertMessageReceiver),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployTollSenderDapp(
	onRampRouter, receiver common.Address,
	destChainId *big.Int,
) (
	*TollSender,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("TollSenderDapp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return toll_sender_dapp.DeployTollSenderDapp(auth, backend, onRampRouter, destChainId, receiver)
	})
	return &TollSender{
		client:     e.evmClient,
		instance:   instance.(*toll_sender_dapp.TollSenderDapp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployTollOnRampRouter() (
	*TollOnRampRouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("TollOnRampRouter", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(auth, backend)
	})
	if err != nil {
		return nil, err
	}
	return &TollOnRampRouter{
		client:     e.evmClient,
		instance:   instance.(*evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployTollOnRamp(
	chainId, destChainId *big.Int,
	tokens, pools, allowList []common.Address,
	afn, router common.Address,
	opts RateLimiterConfig,
) (
	*TollOnRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("TollOnRamp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		config := evm_2_evm_toll_onramp.BaseOnRampInterfaceOnRampConfig{
			CommitFeeJuels:  0,
			MaxDataSize:     1e6,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		}
		return evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
			auth, backend, chainId, destChainId, tokens, pools,
			allowList, afn, config,
			evm_2_evm_toll_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
				Rate:     opts.Rate,
				Capacity: opts.Rate,
			},
			auth.From,
			router)
	})
	if err != nil {
		return nil, err
	}
	return &TollOnRamp{
		client:     e.evmClient,
		instance:   instance.(*evm_2_evm_toll_onramp.EVM2EVMTollOnRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeploySubOnRampRouter(feeToken common.Address) (
	*SubOnRampRouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("SubOnRampRouter", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_any_subscription_onramp_router.DeployEVM2AnySubscriptionOnRampRouter(auth, backend,
			evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouterInterfaceRouterConfig{
				Fee:      big.NewInt(0),
				FeeToken: feeToken,
				FeeAdmin: auth.From,
			},
		)
	})
	if err != nil {
		return nil, err
	}
	return &SubOnRampRouter{
		client:     e.evmClient,
		instance:   instance.(*evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeploySubOnRamp(
	chainId, destChainId *big.Int,
	tokens, pools, allowList []common.Address,
	afn, router common.Address,
	onRampConfig evm_2_evm_subscription_onramp.BaseOnRampInterfaceOnRampConfig,
	rateLimiterConfig evm_2_evm_subscription_onramp.AggregateRateLimiterInterfaceRateLimiterConfig,
) (*SubOnRamp, error) {
	address, _, instance, err := e.evmClient.DeployContract("SubOnRamp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_evm_subscription_onramp.DeployEVM2EVMSubscriptionOnRamp(
			auth, backend, chainId, destChainId, tokens, pools,
			allowList, afn, onRampConfig,
			rateLimiterConfig,
			auth.From,
			router)
	})
	if err != nil {
		return nil, err
	}
	return &SubOnRamp{
		client:     e.evmClient,
		instance:   instance.(*evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp),
		EthAddress: *address,
	}, err
}

func DefaultOffChainAggregatorV2Config(numberNodes int) contracts.OffChainAggregatorV2Config {
	if numberNodes <= 4 {
		log.Err(fmt.Errorf("Insufficient number of nodes (%d) supplied for OCR, need at least 5", numberNodes)).
			Int("Number Chainlink Nodes", numberNodes).
			Msg("You likely need more chainlink nodes to properly configure OCR, try 5 or more.")
	}
	s := make([]int, 0)
	for i := 0; i < numberNodes; i++ {
		s = append(s, 1)
	}
	faultyNodes := 0
	if numberNodes > 1 {
		faultyNodes = numberNodes/3 - 1
	}
	if faultyNodes == 0 {
		faultyNodes = 1
	}
	return contracts.OffChainAggregatorV2Config{
		DeltaProgress:                           10 * time.Second,
		DeltaResend:                             2 * time.Second,
		DeltaRound:                              5 * time.Second,
		DeltaGrace:                              500 * time.Millisecond,
		DeltaStage:                              10 * time.Second,
		RMax:                                    3,
		S:                                       s,
		F:                                       faultyNodes,
		Oracles:                                 []ocrConfigHelper2.OracleIdentityExtra{},
		MaxDurationQuery:                        1 * time.Second,
		MaxDurationObservation:                  1 * time.Second,
		MaxDurationReport:                       1 * time.Second,
		MaxDurationShouldAcceptFinalizedReport:  1 * time.Second,
		MaxDurationShouldTransmitAcceptedReport: 1 * time.Second,
		OnchainConfig:                           []byte{},
	}
}

func stripKeyPrefix(key string) string {
	chunks := strings.Split(key, "_")
	if len(chunks) == 3 {
		return chunks[2]
	}
	return key
}

func NewOffChainAggregatorV2Config(
	nodes []*client.CLNodesWithKeys,
) (
	signers []common.Address,
	transmitters []common.Address,
	f_ uint8,
	onchainConfig_ []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
	err error,
) {
	oracleIdentities := make([]ocrConfigHelper2.OracleIdentityExtra, 0)
	ocrConfig := DefaultOffChainAggregatorV2Config(len(nodes))
	var onChainKeys []ocrtypes2.OnchainPublicKey
	for i, nodeWithKeys := range nodes {
		ocr2Key := nodeWithKeys.KeysBundle.OCR2Key.Data
		offChainPubKeyTemp, err := hex.DecodeString(stripKeyPrefix(ocr2Key.Attributes.OffChainPublicKey))
		if err != nil {
			return nil, nil, 0, nil, 0, nil, err
		}
		formattedOnChainPubKey := stripKeyPrefix(ocr2Key.Attributes.OnChainPublicKey)
		cfgPubKeyTemp, err := hex.DecodeString(stripKeyPrefix(ocr2Key.Attributes.ConfigPublicKey))
		if err != nil {
			return nil, nil, 0, nil, 0, nil, err
		}
		cfgPubKeyBytes := [ed25519.PublicKeySize]byte{}
		copy(cfgPubKeyBytes[:], cfgPubKeyTemp)
		offChainPubKey := [curve25519.PointSize]byte{}
		copy(offChainPubKey[:], offChainPubKeyTemp)
		ethAddress := nodeWithKeys.KeysBundle.EthAddress
		p2pKeys := nodeWithKeys.KeysBundle.P2PKeys
		peerID := p2pKeys.Data[0].Attributes.PeerID
		oracleIdentities = append(oracleIdentities, ocrConfigHelper2.OracleIdentityExtra{
			OracleIdentity: ocrConfigHelper2.OracleIdentity{
				OffchainPublicKey: offChainPubKey,
				OnchainPublicKey:  common.HexToAddress(formattedOnChainPubKey).Bytes(),
				PeerID:            peerID,
				TransmitAccount:   ocrtypes2.Account(ethAddress),
			},
			ConfigEncryptionPublicKey: cfgPubKeyBytes,
		})
		onChainKeys = append(onChainKeys, oracleIdentities[i].OnchainPublicKey)
		transmitters = append(transmitters, common.HexToAddress(ethAddress))
	}
	signers, err = ocrcommon.OnchainPublicKeyToAddress(onChainKeys)
	if err != nil {
		return nil, nil, 0, nil, 0, nil, err
	}
	ocrConfig.Oracles = oracleIdentities
	ocrConfig.ReportingPluginConfig, err = ccip.OffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   1,
	}.Encode()
	if err != nil {
		return nil, nil, 0, nil, 0, nil, err
	}
	_, _, f_, onchainConfig_, offchainConfigVersion, offchainConfig, err = ocrConfigHelper2.ContractSetConfigArgsForTests(
		ocrConfig.DeltaProgress,
		ocrConfig.DeltaResend,
		ocrConfig.DeltaRound,
		ocrConfig.DeltaGrace,
		ocrConfig.DeltaStage,
		ocrConfig.RMax,
		ocrConfig.S,
		ocrConfig.Oracles,
		ocrConfig.ReportingPluginConfig,
		ocrConfig.MaxDurationQuery,
		ocrConfig.MaxDurationObservation,
		ocrConfig.MaxDurationReport,
		ocrConfig.MaxDurationShouldAcceptFinalizedReport,
		ocrConfig.MaxDurationShouldTransmitAcceptedReport,
		ocrConfig.F,
		ocrConfig.OnchainConfig,
	)
	return
}
