package merclib

import (
	"testing"
	"time"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
)

func TestVerifySingle(t *testing.T) {
	t.Run("happy path", func(tt *testing.T) {
		price := assets.Ether(1).ToInt()
		feedID := testutils.RandomFeedIDV3()
		configDigest := testutils.Random32Byte()
		f := 1
		obsTs := uint32(time.Now().UTC().Unix())
		expiresTs := obsTs + 100
		var ocr2Keys []ocr2key.KeyBundle
		for i := 0; i < 2*f+1; i++ {
			kb, err := ocr2key.New(chaintype.EVM)
			require.NoError(tt, err)
			ocr2Keys = append(ocr2Keys, kb)
		}
		signedReport, err := generateReport(feedID, configDigest, obsTs, expiresTs, price, ocr2Keys, f)
		require.NoError(tt, err)
		err = verifySingle(signedReport, f, func() []ocrtypes.OnchainPublicKey {
			var ocr2PubKeys []ocrtypes.OnchainPublicKey
			for _, kb := range ocr2Keys {
				ocr2PubKeys = append(ocr2PubKeys, kb.PublicKey())
			}
			return ocr2PubKeys
		}())
		require.NoError(tt, err)
	})
	t.Run("not enough signatures", func(tt *testing.T) {
		price := assets.Ether(1).ToInt()
		feedID := testutils.RandomFeedIDV3()
		configDigest := testutils.Random32Byte()
		f := 1
		obsTs := uint32(time.Now().UTC().Unix())
		expiresTs := obsTs + 100
		var ocr2Keys []ocr2key.KeyBundle
		for i := 0; i < 2*f+1; i++ {
			kb, err := ocr2key.New(chaintype.EVM)
			require.NoError(tt, err)
			ocr2Keys = append(ocr2Keys, kb)
		}
		signedReport, err := generateReport(feedID, configDigest, obsTs, expiresTs, price, ocr2Keys, f)
		require.NoError(tt, err)
		// verification expects f + 2 signatures but only f + 1 are provided
		err = verifySingle(signedReport, f+1, func() []ocrtypes.OnchainPublicKey {
			var ocr2PubKeys []ocrtypes.OnchainPublicKey
			for _, kb := range ocr2Keys {
				ocr2PubKeys = append(ocr2PubKeys, kb.PublicKey())
			}
			return ocr2PubKeys
		}())
		require.Error(tt, err)
		require.ErrorContainsf(tt, err, "expected 3 signatures, got 2", "err: %v", err)
	})
}
