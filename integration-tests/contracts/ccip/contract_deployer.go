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

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
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

func (e *CCIPContractsDeployer) DeployAFNContract(
	weightByParticipants map[string]*big.Int,
	blessThreshold *big.Int,
	bagSignalThreshold *big.Int,
) (
	*AFN,
	error,
) {
	var weights []*big.Int
	var participants []common.Address
	for addr, weight := range weightByParticipants {
		weights = append(weights, weight)
		participants = append(participants, common.HexToAddress(addr))
	}
	address, _, instance, err := e.evmClient.DeployContract("AFN Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return afn_contract.DeployAFNContract(auth, backend, participants, weights, blessThreshold, bagSignalThreshold)
	})
	return &AFN{
		client:     e.evmClient,
		instance:   instance.(*afn_contract.AFNContract),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployBlobVerifier(
	sourceChainId, destChainId *big.Int,
	afn common.Address,
	bConfig blob_verifier.BlobVerifierInterfaceBlobVerifierConfig,
) (
	*BlobVerifier,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("BlobVerifier Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return blob_verifier.DeployBlobVerifier(auth, backend, destChainId, sourceChainId, afn, bConfig)
	})
	return &BlobVerifier{
		client:     e.evmClient,
		instance:   instance.(*blob_verifier.BlobVerifier),
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

func (e *CCIPContractsDeployer) DeployOffRamp(
	sourceChainId, destChainId *big.Int,
	blobVerifier, onRamp, afn common.Address,
	sourceToken, pools []common.Address,
	opts RateLimiterConfig) (
	*OffRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("OffRamp Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return any_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
			auth, backend, sourceChainId, destChainId,
			any_2_evm_toll_offramp.BaseOffRampInterfaceOffRampConfig{
				OnRampAddress:                           onRamp,
				ExecutionDelaySeconds:                   60,
				MaxDataSize:                             1e5,
				MaxTokensLength:                         15,
				PermissionLessExecutionThresholdSeconds: 60,
			},
			blobVerifier,
			afn,
			sourceToken,
			pools,
			any_2_evm_toll_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
				Rate:     opts.Rate,
				Capacity: opts.Capacity,
			},
			auth.From)
	})
	return &OffRamp{
		client:     e.evmClient,
		instance:   instance.(*any_2_evm_toll_offramp.EVM2EVMTollOffRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployOffRampRouter(
	offRamps []common.Address) (
	*OffRampRouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("OffRampRouter Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(auth, backend, offRamps)
	})
	return &OffRampRouter{
		client:     e.evmClient,
		instance:   instance.(*any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployReceiverDapp(
	offRampRouter common.Address) (
	*ReceiverDapp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("ReceiverDapp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return receiver_dapp.DeployReceiverDapp(auth, backend, offRampRouter)
	})
	return &ReceiverDapp{
		client:     e.evmClient,
		instance:   instance.(*receiver_dapp.ReceiverDapp),
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

func (e *CCIPContractsDeployer) DeployOnRampRouter() (
	*OnRampRouter,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("OnRampRouter", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(auth, backend)
	})
	if err != nil {
		return nil, err
	}
	return &OnRampRouter{
		client:     e.evmClient,
		instance:   instance.(*evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployOnRamp(
	chainId, destChainId *big.Int,
	tokens, pools, allowList []common.Address,
	afn, router common.Address,
	opts RateLimiterConfig,
) (
	*OnRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("OnRamp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		config := evm_2_evm_toll_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e6,
			MaxTokensLength:  5,
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
	return &OnRamp{
		client:     e.evmClient,
		instance:   instance.(*evm_2_evm_toll_onramp.EVM2EVMTollOnRamp),
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
		DeltaProgress:                           40 * time.Second,
		DeltaResend:                             10 * time.Second,
		DeltaRound:                              30 * time.Second,
		DeltaGrace:                              1 * time.Second,
		DeltaStage:                              60 * time.Second,
		RMax:                                    5,
		S:                                       s,
		F:                                       faultyNodes,
		Oracles:                                 []ocrConfigHelper2.OracleIdentityExtra{},
		MaxDurationQuery:                        5 * time.Second,
		MaxDurationObservation:                  5 * time.Second,
		MaxDurationReport:                       5 * time.Second,
		MaxDurationShouldAcceptFinalizedReport:  5 * time.Second,
		MaxDurationShouldTransmitAcceptedReport: 5 * time.Second,
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
		log.Info().Interface("OCR2 Key", ocr2Key).Msg("Key details delete later")
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
