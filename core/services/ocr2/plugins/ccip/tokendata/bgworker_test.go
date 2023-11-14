package tokendata_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestBackgroundWorker(t *testing.T) {
	ctx := testutils.Context(t)

	const numTokens = 100
	const numWorkers = 20
	const numMessages = 40
	const maxReaderLatencyMS = 200
	const percentOfTokensWithoutTokenData = 10

	tokens := make([]common.Address, numTokens)
	readers := make(map[common.Address]*tokendata.MockReader, numTokens)
	tokenDataReaders := make(map[common.Address]tokendata.Reader, numTokens)
	tokenData := make(map[common.Address][]byte)
	delays := make(map[common.Address]time.Duration)

	for i := range tokens {
		tokens[i] = utils.RandomAddress()
		readers[tokens[i]] = tokendata.NewMockReader(t)
		if rand.Intn(100) >= percentOfTokensWithoutTokenData {
			tokenDataReaders[tokens[i]] = readers[tokens[i]]
			tokenData[tokens[i]] = []byte(fmt.Sprintf("...token %x data...", tokens[i]))
		}

		// specify a random latency for the reader implementation
		readerLatency := rand.Intn(maxReaderLatencyMS)
		delays[tokens[i]] = time.Duration(readerLatency) * time.Millisecond
	}
	w := tokendata.NewBackgroundWorker(ctx, tokenDataReaders, numWorkers)

	msgs := make([]internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, numMessages)
	for i := range msgs {
		tk := tokens[i%len(tokens)]

		msgs[i] = internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{
			EVM2EVMMessage: internal.EVM2EVMMessage{
				SequenceNumber: uint64(i + 1),
				TokenAmounts:   []internal.TokenAmount{{Token: tk}},
			},
		}

		reader := readers[tk]
		reader.On("ReadTokenData", mock.Anything, msgs[i]).Run(func(args mock.Arguments) {
			time.Sleep(delays[tk])
		}).Return(tokenData[tk], nil).Maybe()
	}

	w.AddJobsFromMsgs(ctx, msgs)
	// processing of the messages should have started at this point

	tStart := time.Now()
	for _, msg := range msgs {
		b, err := w.GetMsgTokenData(ctx, msg)
		assert.NoError(t, err)
		assert.Equal(t, tokenData[msg.TokenAmounts[0].Token], b[0])
	}
	t.Logf("initial get: %v", time.Since(tStart))

	tStart = time.Now()
	for _, msg := range msgs {
		b, err := w.GetMsgTokenData(ctx, msg)
		assert.NoError(t, err)
		assert.Equal(t, tokenData[msg.TokenAmounts[0].Token], b[0])
	}
	t.Logf("get second time: %v", time.Since(tStart))

	w.AddJobsFromMsgs(ctx, msgs)
	// add same msgs again
	tStart = time.Now()
	for _, msg := range msgs {
		b, err := w.GetMsgTokenData(ctx, msg)
		assert.NoError(t, err)
		assert.Equal(t, tokenData[msg.TokenAmounts[0].Token], b[0])
	}
	t.Logf("get after adding same msgs: %v", time.Since(tStart))
}
