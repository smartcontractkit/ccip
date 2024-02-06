package contracts
import (
	"context"
	"fmt"
	"math/big"
	"testing"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
)
func TestDeployLink(t *testing.T) {
	fmt.Println("test deploy link")
	client, err := ethclient.Dial("https://chain-beta-rpc.main.sand.cldev.sh")
	require.NoError(t, err)
	ownerKey, err := crypto.HexToECDSA("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63")
	require.NoError(t, err)
	fmt.Println("ownerKey", ownerKey)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, big.NewInt(int64(2337)))
	require.NoError(t, err)
	tokenAddr, tx, _, err := link_token_interface.DeployLinkToken(user, client)
	fmt.Println("tokenAddr", tokenAddr)
	fmt.Println("tx", tx)
	fmt.Println("err", err)
	require.NoError(t, err)
	mined, err := bind.WaitMined(context.Background(), client, tx)
	require.NoError(t, err)
	fmt.Println("mined", mined)
	// _ = client
	// _ = ownerKey
	// _ = user

	//     client, err := ethclient.Dial("< your rpc url here>")
//     require.NoError(t, err)
}