package load

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/smartcontractkit/wasp"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
)

type Dummy struct {
	Owner   string
	ChainId *big.Int
	client  *ethclient.Client
}

func (d *Dummy) Close() {
	d.client.Close()
}

func (d *Dummy) Call(_ *wasp.Generator) *wasp.Response {
	res := &wasp.Response{}
	ownerKey, err := crypto.HexToECDSA(d.Owner)
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}

	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, d.ChainId)
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	_, _, _, err = link_token_interface.DeployLinkToken(user, d.client)
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	return res
}

func NewDummy(url, ownerKey string, chainid *big.Int) (*Dummy, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Dummy{
		client:  client,
		Owner:   ownerKey,
		ChainId: chainid,
	}, nil
}

func CreateDummyTraffic(t *testing.T, url, ownerKey string, chainId *big.Int) error {
	d, err := NewDummy(url, ownerKey, chainId)
	if err != nil {
		return err
	}
	waspCfg := &wasp.Config{
		T:                     t,
		Schedule:              wasp.Plain(1, 30*time.Minute),
		LoadType:              wasp.RPS,
		RateLimitUnitDuration: 300 * time.Millisecond,
		CallResultBufLen:      10, // we keep the last 10 call results for each generator, as the detailed report is generated at the end of the test
		Gun:                   d,
	}
	loadRunner, err := wasp.NewGenerator(waspCfg)
	require.NoError(t, err)
	loadRunner.Run(false)
	return nil
}
