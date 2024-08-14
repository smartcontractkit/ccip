package ccipdeployment

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	nodev1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/node/v1"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"

	"github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/validate"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
)

type OCRConfig struct {
	OffchainPublicKey ocrtypes.OffchainPublicKey
	// For EVM-chains, this an *address*.
	OnchainPublicKey          ocrtypes.OnchainPublicKey
	PeerID                    p2pkey.PeerID
	TransmitAccount           ocrtypes.Account
	ConfigEncryptionPublicKey types.ConfigEncryptionPublicKey
	IsBootstrap               bool
	MultiAddr                 string // TODO: type
}

type Nodes []Node

func (n Nodes) PeerIDs(chainSel uint64) [][32]byte {
	var peerIDs [][32]byte
	for _, node := range n {
		cfg := node.selToOCRConfig[chainSel]
		// NOTE: Assume same peerID for all chains.
		// Might make sense to change proto as peerID is 1-1 with node?
		peerIDs = append(peerIDs, cfg.PeerID)
	}
	return peerIDs
}

func (n Nodes) BootstrapPeerIDs(chainSel uint64) [][32]byte {
	var peerIDs [][32]byte
	for _, node := range n {
		cfg := node.selToOCRConfig[chainSel]
		if !cfg.IsBootstrap {
			continue
		}
		peerIDs = append(peerIDs, cfg.PeerID)
	}
	return peerIDs
}

// OffchainPublicKey types.OffchainPublicKey
// // For EVM-chains, this an *address*.
// OnchainPublicKey types.OnchainPublicKey
// PeerID           string
// TransmitAccount  types.Account
type Node struct {
	selToOCRConfig map[uint64]OCRConfig
}

func MustPeerIDFromString(s string) p2pkey.PeerID {
	p := p2pkey.PeerID{}
	if err := p.UnmarshalString(s); err != nil {
		panic(err)
	}
	return p
}

// Gathers all the node info through JD required to be able to set
// OCR config for example.
func NodeInfo(nodeIDs []string, oc deployment.OffchainClient) (Nodes, error) {
	var nodes []Node
	for _, node := range nodeIDs {
		// TODO: Filter should accept multiple nodes
		nodeChainConfigs, err := oc.ListNodeChainConfigs(context.Background(), &nodev1.ListNodeChainConfigsRequest{Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
			NodeId: node,
		}})
		if err != nil {
			return nil, err
		}
		selToOCRConfig := make(map[uint64]OCRConfig)
		for _, chainConfig := range nodeChainConfigs.ChainConfigs {
			if chainConfig.Chain.Type == nodev1.ChainType_CHAIN_TYPE_SOLANA {
				// Note supported for CCIP yet.
				continue
			}
			evmChainID, err := strconv.Atoi(chainConfig.Chain.Id)
			if err != nil {
				return nil, err
			}
			sel, err := chainsel.SelectorFromChainId(uint64(evmChainID))
			if err != nil {
				return nil, err
			}
			b := common.Hex2Bytes(chainConfig.Ocr2Config.OcrKeyBundle.OffchainPublicKey)
			var opk ocrtypes.OffchainPublicKey
			copy(opk[:], b)

			b = common.Hex2Bytes(chainConfig.Ocr2Config.OcrKeyBundle.ConfigPublicKey)
			var cpk types.ConfigEncryptionPublicKey
			copy(cpk[:], b)

			selToOCRConfig[sel] = OCRConfig{
				OffchainPublicKey:         opk,
				OnchainPublicKey:          common.HexToAddress(chainConfig.Ocr2Config.OcrKeyBundle.OnchainSigningAddress).Bytes(),
				PeerID:                    MustPeerIDFromString(chainConfig.Ocr2Config.P2PKeyBundle.PeerId),
				TransmitAccount:           ocrtypes.Account(chainConfig.AccountAddress),
				ConfigEncryptionPublicKey: cpk,
				IsBootstrap:               chainConfig.Ocr2Config.IsBootstrap,
				MultiAddr:                 chainConfig.Ocr2Config.Multiaddr,
			}
		}
		nodes = append(nodes, Node{
			selToOCRConfig: selToOCRConfig,
		})
	}
	return nodes, nil
}

// In our case, the only address needed is the cap registry which is actually an env var.
// and will pre-exist for our deployment. So the job specs only depend on the environment operators.
func NewCCIPJobSpecs(nodeIds []string, oc deployment.OffchainClient) (map[string][]string, error) {
	// Generate a set of brand new job specs for CCIP for a specific environment
	// (including NOPs) and new addresses.
	// We want to assign one CCIP capability job to each node. And node with
	// an addr we'll list as bootstrapper.
	// Find the bootstrap nodes
	bootstrapMp := make(map[string]struct{})
	for _, node := range nodeIds {
		// TODO: Filter should accept multiple nodes
		nodeChainConfigs, err := oc.ListNodeChainConfigs(context.Background(), &nodev1.ListNodeChainConfigsRequest{Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
			NodeId: node,
		}})
		if err != nil {
			return nil, err
		}
		for _, chainConfig := range nodeChainConfigs.ChainConfigs {
			if chainConfig.Ocr2Config.IsBootstrap {
				bootstrapMp[fmt.Sprintf("%s@%s",
					// p2p_12D3... -> 12D3...
					chainConfig.Ocr2Config.P2PKeyBundle.PeerId[4:], chainConfig.Ocr2Config.Multiaddr)] = struct{}{}
			}
		}
	}
	var bootstraps []string
	for b := range bootstrapMp {
		bootstraps = append(bootstraps, b)
	}
	nodesToJobSpecs := make(map[string][]string)
	for _, node := range nodeIds {
		// TODO: Filter should accept multiple.
		nodeChainConfigs, err := oc.ListNodeChainConfigs(context.Background(), &nodev1.ListNodeChainConfigsRequest{Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
			NodeId: node,
		}})
		if err != nil {
			return nil, err
		}
		spec, err := validate.NewCCIPSpecToml(validate.SpecArgs{
			P2PV2Bootstrappers:     bootstraps,
			CapabilityVersion:      CapabilityVersion,
			CapabilityLabelledName: CapabilityLabelledName,
			OCRKeyBundleIDs: map[string]string{
				// TODO: Validate that that all EVM chains are using the same keybundle.
				relay.NetworkEVM: nodeChainConfigs.ChainConfigs[0].Ocr2Config.OcrKeyBundle.BundleId,
			},
			// TODO: validate that all EVM chains are using the same keybundle
			P2PKeyID:     nodeChainConfigs.ChainConfigs[0].Ocr2Config.P2PKeyBundle.PeerId,
			RelayConfigs: nil,
			PluginConfig: map[string]any{},
		})
		if err != nil {
			return nil, err
		}
		nodesToJobSpecs[node] = append(nodesToJobSpecs[node], spec)
	}
	return nodesToJobSpecs, nil
}
