package merclib

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/llo-feeds/generated/verifier"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/llo-feeds/generated/verifier_proxy"
)

var _ ReportVerifier = &offlineVerifier{}

type offlineVerifier struct {
	pubkeyGetter ConfigInfoGetter
}

// NewOfflineVerifier returns a new ReportVerifier instance
// that verifies reports offline by fetching the ocr onchain public
// keys from the given ConfigInfoGetter.
// NOTE: verification here is EVM-only.
func NewOfflineVerifier(pubkeyGetter ConfigInfoGetter) *offlineVerifier {
	return &offlineVerifier{pubkeyGetter: pubkeyGetter}
}

func (o *offlineVerifier) VerifyReports(ctx context.Context, signedReports [][]byte) error {
	for _, signedReport := range signedReports {
		// decode into raw report constituents
		rwc, err := DecodeFullReportAndReportData(signedReport)
		if err != nil {
			return fmt.Errorf("decoding full report: %w", err)
		}

		// reportContext consists of:
		// reportContext[0]: ConfigDigest
		// reportContext[1]: 27 byte padding, 4-byte epoch and 1-byte round
		// reportContext[2]: ExtraHash
		configDigest := rwc.FullReport.ReportContext[0]
		onchainPubKeys, f, err := o.pubkeyGetter.GetOnchainPublicKeysAndF(configDigest, rwc.FeedId)
		if err != nil {
			return fmt.Errorf("getting onchain public keys and f: %w", err)
		}

		if err := verifySingle(rwc.FullReport, f, onchainPubKeys); err != nil {
			return err
		}
	}
	return nil
}

// verifySingle verifies a single signed report
// This is based off of the Verifier.sol contract:
// function verify(
//
//	  bytes calldata signedReport,
//	  address sender
//	) external override returns (bytes memory verifierResponse)
func verifySingle(fullReport *FullReport, f int, pubkeys []ocrtypes.OnchainPublicKey) error {
	if err := validateReport(fullReport.RawRs, fullReport.RawSs, f); err != nil {
		return fmt.Errorf("validating report: %w", err)
	}

	var signedCount int
	for i := range fullReport.RawRs {
		combinedSig := combineSignature(fullReport.RawRs[i], fullReport.RawSs[i], fullReport.RawVs[i])
		pubkey := ecrecover(
			fullReport.ReportContext,
			fullReport.ReportBlob,
			combinedSig)
		if pubkey == nil {
			continue
		}
		for _, ocr2Pubkey := range pubkeys {
			if bytes.Equal(pubkey, ocr2Pubkey) {
				signedCount++
				break
			}
		}
	}

	if signedCount != (f + 1) {
		return fmt.Errorf("expected %d signatures from OCR keys %+v, got %d", f+1, formatOCRKeys(pubkeys), signedCount)
	}

	return nil
}

func validateReport(rs [][32]byte, ss [][32]byte, f int) error {
	expectedNumSignatures := f + 1
	if len(rs) != expectedNumSignatures {
		return fmt.Errorf("expected %d signatures, got %d", expectedNumSignatures, len(rs))
	}
	if len(rs) != len(ss) {
		return fmt.Errorf("mismatched rs and ss: len(rs)=%d, len(ss)=%d", len(rs), len(ss))
	}
	return nil
}

func reportToSigData(rawReportContext [3][32]byte, report ocrtypes.Report) []byte {
	sigData := crypto.Keccak256(report)
	sigData = append(sigData, rawReportContext[0][:]...)
	sigData = append(sigData, rawReportContext[1][:]...)
	sigData = append(sigData, rawReportContext[2][:]...)
	return crypto.Keccak256(sigData)
}

func ecrecover(
	reportCtx [3][32]byte,
	report ocrtypes.Report,
	signature []byte) []byte {
	hash := reportToSigData(reportCtx, report)
	authorPubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return nil
	}
	authorAddress := crypto.PubkeyToAddress(*authorPubkey)
	return authorAddress.Bytes()
}

func combineSignature(rs [32]byte, ss [32]byte, v byte) (sig []byte) {
	sig = append(sig, rs[:]...)
	sig = append(sig, ss[:]...)
	sig = append(sig, v)
	return sig
}

func formatOCRKeys(pubkeys []ocrtypes.OnchainPublicKey) (r []string) {
	for _, key := range pubkeys {
		r = append(r, hexutil.Encode(key))
	}
	return r
}

// ConfigInfoGetter is an interface for fetching onchain public keys and f
type ConfigInfoGetter interface {
	GetOnchainPublicKeysAndF(configDigest, feedID [32]byte) (onchainPubKeys []ocrtypes.OnchainPublicKey, f int, err error)
}

var _ ConfigInfoGetter = &cachedConfigInfoGetter{}

// cachedConfigInfoGetter is an implementation of ConfigInfoGetter that caches the
// result of GetOnchainPublicKeys for a time period prior to refreshing.
type cachedConfigInfoGetter struct {
	// each configDigest maps to a set of onchain public keys
	cache map[[32]byte]struct {
		pubkeys []ocrtypes.OnchainPublicKey
		f       int
	}
	lastFetch     time.Time
	refresh       time.Duration
	client        evmclient.Client
	verifierProxy verifier_proxy.VerifierProxyInterface
}

func NewCachedConfigInfoGetter(
	client evmclient.Client,
	refresh time.Duration,
	verifierProxyAddress common.Address) (*cachedConfigInfoGetter, error) {
	verifierProxy, err := verifier_proxy.NewVerifierProxy(verifierProxyAddress, client)
	if err != nil {
		return nil, err
	}

	return &cachedConfigInfoGetter{
		client:        client,
		refresh:       refresh,
		verifierProxy: verifierProxy,
		cache: make(map[[32]byte]struct {
			pubkeys []ocrtypes.OnchainPublicKey
			f       int
		}),
	}, nil
}

func (c *cachedConfigInfoGetter) GetOnchainPublicKeysAndF(configDigest, feedID [32]byte) ([]ocrtypes.OnchainPublicKey, int, error) {
	pair, ok := c.cache[configDigest]
	if !ok || time.Since(c.lastFetch) > c.refresh {
		pubkeys, f, err := c.getConfigDigestInfo(configDigest, feedID)
		if err != nil {
			return nil, -1, err
		}
		c.cache[configDigest] = struct {
			pubkeys []ocrtypes.OnchainPublicKey
			f       int
		}{
			pubkeys: pubkeys,
			f:       f,
		}
		c.lastFetch = time.Now()
		return pubkeys, f, nil
	}
	return pair.pubkeys, pair.f, nil
}

func (c *cachedConfigInfoGetter) getConfigDigestInfo(configDigest, feedID [32]byte) ([]ocrtypes.OnchainPublicKey, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	verifierAddress, err := c.verifierProxy.GetVerifier(&bind.CallOpts{Context: ctx}, configDigest)
	if err != nil {
		return nil, -1, fmt.Errorf("getting verifier address for config digest %x: %w", configDigest, err)
	}
	tVerifier, err := verifier.NewVerifier(verifierAddress, c.client)
	if err != nil {
		return nil, -1, fmt.Errorf("creating verifier: %w", err)
	}
	configDetails, err := tVerifier.LatestConfigDetails(&bind.CallOpts{Context: ctx}, feedID)
	if err != nil {
		return nil, -1, fmt.Errorf("getting config details for config digest %x on verifier %s: %w", configDigest, verifierAddress.Hex(), err)
	}

	startEnd := uint64(configDetails.BlockNumber)
	iter, err := tVerifier.FilterConfigSet(&bind.FilterOpts{
		Context: ctx,
		Start:   startEnd,
		End:     &startEnd,
	}, nil)
	if err != nil {
		return nil, -1, fmt.Errorf("filtering configset events on verifier %s, startEnd: %d: %w", verifierAddress.Hex(), startEnd, err)
	}

	var (
		pubkeys []ocrtypes.OnchainPublicKey
		f       int
	)
	for iter.Next() {
		// shouldn't happen, but check anyways
		if iter.Event.ConfigDigest != configDigest {
			continue
		}
		pubkeys = func() []ocrtypes.OnchainPublicKey {
			var pubkeys []ocrtypes.OnchainPublicKey
			for _, key := range iter.Event.Signers {
				pubkeys = append(pubkeys, key.Bytes())
			}
			return pubkeys
		}()
		f = int(iter.Event.F)
	}
	if len(pubkeys) == 0 {
		return nil, -1, fmt.Errorf("no configset events found for config digest %x", configDigest)
	}
	return pubkeys, f, nil
}
