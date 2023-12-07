package merclib

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var _ ReportVerifier = &ethCallVerifier{}

type ethCallVerifier struct {
	ethClient evmclient.Client
	// verifierProxyAddress is the address of the verifier proxy contract
	verifierProxyAddress common.Address
	// wrappedNativeAddress is the address of the wrapped native token
	// on the same chain as the verifier proxy
	wrappedNativeAddress common.Address
	// fromAddress is the address that is used to call the verifier proxy contract
	// in practice this is the OCR2 onchain key as an ethereum address
	fromAddress common.Address
	lggr        logger.Logger
}

// NewEthCallVerifier returns a new ReportVerifier instance
// that uses the verifier proxy onchain contract to verify mercury reports
// by calling the verifyBulk method.
// NOTE: the fromAddress passed into the verifier must have 100% discount on the mercury fee manager contract,
// otherwise the verification will fail due to insufficient funds.
func NewEthCallVerifier(
	ethClient evmclient.Client,
	verifierProxyAddress,
	wrappedNativeAddress,
	fromAddress common.Address,
	lggr logger.Logger) *ethCallVerifier {
	return &ethCallVerifier{
		ethClient:            ethClient,
		verifierProxyAddress: verifierProxyAddress,
		wrappedNativeAddress: wrappedNativeAddress,
		fromAddress:          fromAddress,
		lggr:                 lggr,
	}
}

func (e *ethCallVerifier) VerifyReports(ctx context.Context, signedReports [][]byte) error {
	// payment options don't really matter since an eth_call won't pay anything in the end
	// NOTE: the fromAddress passed into the client must have 100% discount on the mercury fee manager contract.
	// need to abi-encode the address first before passing it in because the fee manager uses abi.decode()
	// and expects 32 bytes
	encodedAddr, err := utils.ABIEncode(`[{ "type": "address" }]`, e.wrappedNativeAddress)
	if err != nil {
		return fmt.Errorf("failed to abi encode address '%s': %w", e.wrappedNativeAddress.String(), err)
	}
	calldata, err := verifierProxyABI.Pack("verifyBulk", signedReports, encodedAddr)
	if err != nil {
		return fmt.Errorf("failed to pack verifyBulk: %w", err)
	}
	callMsg := ethereum.CallMsg{
		From:     e.fromAddress,
		Data:     calldata,
		To:       &e.verifierProxyAddress,
		GasPrice: nil,
		Gas:      0,
	}
	e.lggr.Debugw("calling verifier contract",
		"verifierProxyAddress", e.verifierProxyAddress.String(), "calldata", hexutil.Encode(calldata), "callMsg", callMsg)
	_, err = e.ethClient.CallContract(ctx, callMsg, nil)
	if err != nil {
		return fmt.Errorf("failed to call verifier contract at %s: %w, calldata: %s", e.verifierProxyAddress.String(), err, hexutil.Encode(calldata))
	}

	// simulation passing means verification succeeded
	return nil
}
