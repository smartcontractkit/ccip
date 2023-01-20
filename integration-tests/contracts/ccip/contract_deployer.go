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

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
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

func (e *CCIPContractsDeployer) NewLinkTokenContract(addr common.Address) (contracts.LinkToken, error) {
	return e.EthDeployer.NewLinkTokenContract(addr)
}

func (e *CCIPContractsDeployer) NewNativeTokenPoolContract(addr common.Address) (
	*NativeTokenPool,
	error,
) {
	pool, err := native_token_pool.NewNativeTokenPool(addr, e.evmClient.Client)
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "Native Token Pool").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &NativeTokenPool{
		client:     e.evmClient,
		instance:   pool,
		EthAddress: addr,
	}, err
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

func (e *CCIPContractsDeployer) NewAFNContract(addr common.Address) (*AFN, error) {
	afn, err := mock_afn_contract.NewMockAFNContract(addr, e.evmClient.Client)
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "Mock AFN Contract").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &AFN{
		client:     e.evmClient,
		instance:   afn,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployCommitStore(
	sourceChainId, destChainId uint64,
	afn common.Address,
	bConfig commit_store.ICommitStoreCommitStoreConfig,
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
	sourceChainId, destChainId uint64,
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
		return evm_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
			auth, backend, sourceChainId, destChainId,
			evm_2_evm_toll_offramp.IBaseOffRampOffRampConfig{
				PermissionLessExecutionThresholdSeconds: 0,
				ExecutionDelaySeconds:                   0,
				MaxDataSize:                             1e5,
				MaxTokensLength:                         15,
			},
			onRamp,
			commitStore,
			afn,
			sourceToken,
			pools,
			evm_2_evm_toll_offramp.IAggregateRateLimiterRateLimiterConfig{
				Rate:     opts.Rate,
				Capacity: opts.Capacity,
			},
			auth.From)
	})
	return &TollOffRamp{
		client:     e.evmClient,
		instance:   instance.(*evm_2_evm_toll_offramp.EVM2EVMTollOffRamp),
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
	destChainId uint64,
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
	chainId, destChainId uint64,
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
		config := evm_2_evm_toll_onramp.IBaseOnRampOnRampConfig{
			CommitFeeJuels:  0,
			MaxDataSize:     1e6,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		}
		return evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
			auth, backend, chainId, destChainId, tokens, pools,
			allowList, afn, config,
			evm_2_evm_toll_onramp.IAggregateRateLimiterRateLimiterConfig{
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

func (e *CCIPContractsDeployer) DeployGERouter(
	offRamps []common.Address,
) (
	*GERouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("GERouter", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return ge_router.DeployGERouter(auth, backend, offRamps)
	})
	if err != nil {
		return nil, err
	}
	return &GERouter{
		client:     e.evmClient,
		Instance:   instance.(*ge_router.GERouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewGERouter(addr common.Address) (
	*GERouter,
	error,
) {
	r, err := ge_router.NewGERouter(addr, e.evmClient.Client)
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "GERouter").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &GERouter{
		client:     e.evmClient,
		Instance:   r,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployFeeManager(
	feeUpdates []fee_manager.GEFeeUpdate,
) (
	*FeeManager,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("FeeManager", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return fee_manager.DeployFeeManager(auth, backend, feeUpdates, nil, big.NewInt(1e18))
	})
	if err != nil {
		return nil, err
	}
	return &FeeManager{
		client:     e.evmClient,
		instance:   instance.(*fee_manager.FeeManager),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployGEOnRamp(
	sourceChainId, destChainId uint64,
	tokens, pools, allowList []common.Address,
	afn, router common.Address,
	opts RateLimiterConfig,
	feeConfig evm_2_evm_ge_onramp.IEVM2EVMGEOnRampDynamicFeeConfig,
) (
	*GEOnRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("GEOnRamp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_evm_ge_onramp.DeployEVM2EVMGEOnRamp(auth, backend, sourceChainId, destChainId, tokens, pools, allowList, afn,
			evm_2_evm_ge_onramp.IBaseOnRampOnRampConfig{
				CommitFeeJuels:  0,
				MaxDataSize:     1e5,
				MaxTokensLength: 5,
				MaxGasLimit:     ccip.GasLimitPerTx,
			},
			evm_2_evm_ge_onramp.IAggregateRateLimiterRateLimiterConfig{
				Capacity: opts.Capacity,
				Rate:     opts.Rate,
			},
			auth.From,
			router,
			feeConfig)
	})
	if err != nil {
		return nil, err
	}
	return &GEOnRamp{
		client:     e.evmClient,
		instance:   instance.(*evm_2_evm_ge_onramp.EVM2EVMGEOnRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployGEOffRamp(
	sourceChainId, destChainId uint64,
	commitStore, onRamp, afn, feetoken, destFeeManagerAddress common.Address,
	sourceToken, pools []common.Address,
	opts RateLimiterConfig, gasOverhead *big.Int) (
	*GEOffRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("GEOffRamp Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_evm_ge_offramp.DeployEVM2EVMGEOffRamp(
			auth, backend, sourceChainId, destChainId,
			evm_2_evm_ge_offramp.IEVM2EVMGEOffRampGEOffRampConfig{
				GasOverhead:                             gasOverhead,
				FeeManager:                              destFeeManagerAddress,
				PermissionLessExecutionThresholdSeconds: 0,
				ExecutionDelaySeconds:                   0,
				MaxDataSize:                             1e5,
				MaxTokensLength:                         15,
			},
			onRamp,
			commitStore,
			afn,
			sourceToken,
			pools,
			evm_2_evm_ge_offramp.IAggregateRateLimiterRateLimiterConfig{
				Rate:     opts.Rate,
				Capacity: opts.Capacity,
			},
			auth.From, feetoken)
	})
	return &GEOffRamp{
		client:     e.evmClient,
		instance:   instance.(*evm_2_evm_ge_offramp.EVM2EVMGEOffRamp),
		EthAddress: *address,
	}, err
}

func DefaultOffChainAggregatorV2Config(numberNodes int) contracts.OffChainAggregatorV2Config {
	if numberNodes <= 4 {
		log.Err(fmt.Errorf("insufficient number of nodes (%d) supplied for OCR, need at least 5", numberNodes)).
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
		DeltaStage:                              30 * time.Second,
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
