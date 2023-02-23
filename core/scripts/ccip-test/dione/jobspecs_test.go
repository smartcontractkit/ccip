package dione

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

func TestGetTokensPerFeeCoinPipeline(t *testing.T) {
	link := common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca")
	weth := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	var tt = []struct {
		tokens   map[rhea.Token]rhea.EVMBridgedToken
		expected string
	}{
		{
			map[rhea.Token]rhea.EVMBridgedToken{
				rhea.LINK: {Token: link},
			},
			fmt.Sprintf(`merge [type=merge left="{}" right="{\\\"%s\\\":\\\"1000000000000000000\\\"}"];`, link.Hex()),
		},
		{
			map[rhea.Token]rhea.EVMBridgedToken{
				rhea.LINK: {Token: link},
				rhea.WETH: {Token: weth},
			},
			fmt.Sprintf(`merge [type=merge left="{}" right="{\\\"%s\\\":\\\"1000000000000000000\\\",\\\"%s\\\":\\\"1000000000000000000\\\"}"];`, link.Hex(), weth.Hex()),
		},
	}

	for _, tc := range tt {
		tc := tc
		a := GetTokensPerFeeCoinPipeline(tc.tokens)
		assert.Equal(t, tc.expected, a)
	}
}
