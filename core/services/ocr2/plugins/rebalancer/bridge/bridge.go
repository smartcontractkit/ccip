package bridge

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

// Bridge provides a way to get pending transfers from one chain to another
// for transfers that are using the native bridge for the (source, dest) chain pair.
// For example, if ethereum is the source, and arbitrum is the destination, the bridge
// would be able to get pending transfers from ethereum to arbitrum via the standard arbitrum
// bridge.
type Bridge interface {
	// GetTransfers returns all of the pending transfers from the source chain to the destination chain
	// for the given local and remote token addresses.
	GetTransfers(ctx context.Context, localToken, remoteToken models.Address) ([]models.PendingTransfer, error)

	// GetBridgeSpecificPayload returns the bridge specific payload for the given transfer.
	// This payload must always correctly ABI-encoded.
	// Note that this payload is not directly provided to the bridge but the bridge adapter
	// contracts. The bridge adapter may slightly alter the payload before sending it to the bridge.
	// For example, for an L1 to L2 transfer using Arbitrum's bridge, this will return the
	// fees required for the transfer to succeed reliably.
	// For an L2 -> L1 finalization transaction, this will return the finalization payload.
	GetBridgeSpecificPayload(ctx context.Context, transfer models.Transfer) ([]byte, error)

	Close(ctx context.Context) error

	// LocalChainSelector returns the local chain selector of the bridge
	// This is where tokens are being transferred from
	LocalChainSelector() models.NetworkSelector

	// RemoteChainSelector returns the destination chain selector of the bridge
	// This is where tokens are being transferred to
	RemoteChainSelector() models.NetworkSelector
}

type Container interface {
	GetBridge(source, dest models.NetworkSelector) (Bridge, error)
}

type container struct {
	bridges map[models.NetworkSelector]map[models.NetworkSelector]Bridge
}

func NewContainer() Container {
	return &container{
		bridges: make(map[models.NetworkSelector]map[models.NetworkSelector]Bridge),
	}
}

func (c *container) GetBridge(source, dest models.NetworkSelector) (Bridge, error) {
	if source == dest {
		return nil, fmt.Errorf("no bridge between the same network and itself: %d", source)
	}

	bridgesFromSource, ok := c.bridges[dest]
	if !ok {
		return nil, nil
	}

	bridgeToDest, ok := bridgesFromSource[dest]
	if !ok {
		return nil, nil
	}

	return bridgeToDest, nil
}

func (c *container) AddBridge(from, to models.NetworkSelector, bridge Bridge) error {
	_, ok := c.bridges[from]
	if !ok {
		c.bridges[from] = make(map[models.NetworkSelector]Bridge)
	}

	// check if bridge is already set
	if _, ok := c.bridges[from][to]; ok {
		return fmt.Errorf("bridge already set from %d to %d", from, to)
	}

	c.bridges[from][to] = bridge
	return nil
}
