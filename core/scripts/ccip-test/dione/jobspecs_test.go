package dione

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

func TestGetTokenPricesUSDPipeline(t *testing.T) {
	srcWeth := rhea.EVMBridgedToken{
		Token: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		Price: new(big.Int).Mul(big.NewInt(1500), big.NewInt(1e18)),
	}
	dstLink := rhea.EVMBridgedToken{
		Token: common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca"),
		Price: new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)),
	}
	dstWeth := rhea.EVMBridgedToken{
		Token: common.HexToAddress("0x4200000000000000000000000000000000000006"),
		Price: new(big.Int).Mul(big.NewInt(1500), big.NewInt(1e18)),
	}
	var tt = []struct {
		pipelineTokens []rhea.EVMBridgedToken
		expected       string
	}{
		{
			[]rhea.EVMBridgedToken{dstLink, srcWeth},
			fmt.Sprintf(`merge [type=merge left="{}" right="{\\\"%s\\\":\\\"10000000000000000000\\\",\\\"%s\\\":\\\"1500000000000000000000\\\"}"];`,
				dstLink.Token.Hex(), srcWeth.Token.Hex()),
		},
		{
			[]rhea.EVMBridgedToken{dstLink, dstWeth, srcWeth},
			fmt.Sprintf(`merge [type=merge left="{}" right="{\\\"%s\\\":\\\"10000000000000000000\\\",\\\"%s\\\":\\\"1500000000000000000000\\\",\\\"%s\\\":\\\"1500000000000000000000\\\"}"];`,
				dstLink.Token.Hex(), dstWeth.Token.Hex(), srcWeth.Token.Hex()),
		},
	}

	for _, tc := range tt {
		tc := tc
		a := GetTokenPricesUSDPipeline(tc.pipelineTokens)
		assert.Equal(t, tc.expected, a)
	}
}
