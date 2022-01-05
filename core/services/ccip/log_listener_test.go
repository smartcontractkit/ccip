package ccip

import (
	"bytes"
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/lib/pq"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/testdata/testspecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type lc struct {
}

func (l lc) BlockBackfillDepth() uint64 {
	return 1
}

func (l lc) BlockBackfillSkip() bool {
	return false
}

func (l lc) EvmFinalityDepth() uint32 {
	return 50
}

func (l lc) EvmLogBackfillBatchSize() uint32 {
	return 1
}

func TestLogListener_SavesRequests(t *testing.T) {
	// Deploy contract
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		user.From: {Balance: big.NewInt(1000000000000000000)}},
		ethconfig.Defaults.Miner.GasCeil)
	linkTokenAddress, _, linkToken, err := link_token_interface.DeployLinkToken(user, backend)
	require.NoError(t, err)
	poolAddress, _, pool, err := lock_unlock_pool.DeployLockUnlockPool(user, backend, linkTokenAddress)
	require.NoError(t, err)
	afn := deployAfn(t, user, backend)

	onRampAddress, _, _, err := single_token_onramp.DeploySingleTokenOnRamp(
		user,               // user
		backend,            // client
		big.NewInt(2),      // source chain id
		linkTokenAddress,   // source token
		poolAddress,        // source pool
		big.NewInt(1),      // dest chain id
		linkTokenAddress,   // remoteToken
		[]common.Address{}, // allow list
		false,              // enableAllowList
		big.NewInt(1),      // token bucket rate
		big.NewInt(1000),   // token bucket capacity
		afn,                // AFN
		// 86400 seconds = one day
		big.NewInt(86400), //maxTimeWithoutAFNSignal
	)
	require.NoError(t, err)
	onRamp, err := single_token_onramp.NewSingleTokenOnRamp(onRampAddress, backend)
	require.NoError(t, err)
	_, err = pool.SetOnRamp(user, onRampAddress, true)
	require.NoError(t, err)
	_, err = linkToken.Approve(user, poolAddress, big.NewInt(100))
	require.NoError(t, err)
	offRampAddress, _, _, err := single_token_offramp.DeploySingleTokenOffRamp(
		user,             // user
		backend,          // client
		big.NewInt(1),    // source chain id
		big.NewInt(2),    // dest chain id
		linkTokenAddress, // link token address
		poolAddress,      // dest pool address
		big.NewInt(1),    // token bucket rate
		big.NewInt(1000), // token bucket capacity
		afn,              // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(offRampAddress, backend)
	require.NoError(t, err)
	backend.Commit()

	// Start the log broadcaster/log listener
	// and add a CCIP job.
	db := pgtest.NewGormDB(t)
	ethClient := eth.NewClientFromSim(backend, big.NewInt(1337))
	lggr := logger.TestLogger(t)
	lorm := log.NewORM(db, *big.NewInt(1337))
	r, err := lorm.FindConsumedLogs(0, 100)
	require.NoError(t, err)
	t.Log(r)
	lb := log.NewBroadcaster(lorm, ethClient, lc{}, lggr, nil)
	require.NoError(t, lb.Start())
	jobORM := job.NewORM(db, nil, pipeline.NewORM(db), nil, lggr)
	ccipSpec, err := ValidatedCCIPSpec(testspecs.GenerateCCIPSpec(testspecs.CCIPSpecParams{}).Toml())
	require.NoError(t, err)
	jb, err := jobORM.CreateJob(context.Background(), &ccipSpec, ccipSpec.Pipeline)
	require.NoError(t, err)
	ccipConfig := OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   0,
	}
	logListener := NewLogListener(logger.Default, lb, lb, onRamp, offRamp, ccipConfig, db, jb.ID)
	t.Log("Ramp address", onRampAddress, onRamp.Address())
	require.NoError(t, logListener.Start())

	// Update the ccip config on chain and assert that the log listener uses the new config values
	newCcipConfig := OffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   5,
	}
	updateOffchainConfig(t, newCcipConfig, offRamp, user)
	backend.Commit()

	// Send blocks until that request is saved.
	head, err := backend.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)
	startHead := head.Number.Int64()
	var reqs []*Request
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		lb.OnNewLongestChain(context.Background(), eth.Head{Hash: head.Hash(), Number: startHead})
		startHead++
		reqs, err = logListener.orm.Requests(big.NewInt(2), big.NewInt(1), big.NewInt(0), nil, RequestStatusUnstarted, nil, nil)
		require.NoError(t, err)
		t.Logf("log %+v\n", reqs)
		return logListener.offchainConfig.DestIncomingConfirmations == newCcipConfig.DestIncomingConfirmations &&
			logListener.offchainConfig.SourceIncomingConfirmations == newCcipConfig.SourceIncomingConfirmations
	}, 3*time.Second, 100*time.Millisecond).Should(gomega.BeTrue())

	//Send a request.
	executor := common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4")
	msg := single_token_onramp.CCIPMessagePayload{
		Receiver: linkTokenAddress,
		Data:     []byte("hello xchain world"),
		Tokens:   []common.Address{linkTokenAddress},
		Amounts:  []*big.Int{big.NewInt(100)},
		Executor: executor,
		Options:  nil,
	}
	_, err = onRamp.RequestCrossChainSend(user, msg)
	require.NoError(t, err)
	backend.Commit()

	// Send blocks until that request is saved.
	head, err = backend.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)
	startHead = head.Number.Int64()
	reqs = []*Request{}
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		lb.OnNewLongestChain(context.Background(), eth.Head{Hash: head.Hash(), Number: startHead})
		startHead++
		reqs, err = logListener.orm.Requests(big.NewInt(2), big.NewInt(1), big.NewInt(0), nil, RequestStatusUnstarted, nil, nil)
		require.NoError(t, err)
		t.Logf("log %+v\n", reqs)
		return len(reqs) == 1
	}, 3*time.Second, 100*time.Millisecond).Should(gomega.BeTrue())

	// Assert the xchain request was saved correctly.
	assert.Equal(t, "100", reqs[0].Amounts[0])
	assert.Equal(t, msg.Data, reqs[0].Data)
	assert.Equal(t, pq.StringArray{linkTokenAddress.String()}, reqs[0].Tokens)
	assert.Equal(t, msg.Receiver, reqs[0].Receiver)
	assert.Equal(t, msg.Executor.String(), reqs[0].Executor.String())
	assert.Equal(t, []byte{}, reqs[0].Options)
	// We expect the raw request bytes to be the abi.encoded CCIP Message
	b, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{single_token_onramp.CCIPMessage{
		SequenceNumber:     big.NewInt(1),
		SourceChainId:      big.NewInt(2),
		DestinationChainId: big.NewInt(1),
		Sender:             user.From,
		Payload:            msg,
	}})
	require.NoError(t, err)
	require.True(t, bytes.Equal(reqs[0].Raw, b))
	// Round trip should be the same bytes
	cmsg, err := abihelpers.DecodeCCIPMessage(b)
	require.NoError(t, err)
	b2, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{cmsg})
	require.NoError(t, err)
	require.True(t, bytes.Equal(b2, b))

	require.NoError(t, lb.Close())
	require.NoError(t, logListener.Close())
	require.NoError(t, jobORM.DeleteJob(context.Background(), jb.ID))
}

func updateOffchainConfig(t *testing.T, reportingPluginConfig OffchainConfig, offRamp *single_token_offramp.SingleTokenOffRamp, user *bind.TransactOpts) {
	encoded, err := reportingPluginConfig.Encode()
	require.NoError(t, err)

	var oracles = []confighelper2.OracleIdentityExtra{
		{
			// Node 1
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0xf4e7b2426718b11d8df7008d688d48c8926768d3").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0x016D97857a21A501a0C10b526011516000cE4586"),
				OffchainPublicKey: hexutil.MustDecode("0x510bdd47650e70f3006b24261944d5c3685bc1b8194e5e209beea02916189952"),
				PeerID:            "12D3KooWENNxGhdSx7wXWRXcrZ2uKrY8FEagUCntS6Jw55gXqrTX",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0xb2b25ce373a833e3fa7f23538a6ace837673e4ef890db7f7e02830e8d5b6d009"),
		},
		{
			// Node 2
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x33a96c0976DD8c10Cc3e9709Ed25f2CF7d7d970E").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0xcca943C692b27b47a43cB532b2354591BD8a7E9b"),
				OffchainPublicKey: hexutil.MustDecode("0x705cec8e7df7ca42fb8465a60e68ff4e02afd90e17dfef2b01e1166c8dd0cb96"),
				PeerID:            "12D3KooWJtEHwtgkC96umAg2C3Gc8oWpqqT81z6RQXEhkFZK1P21",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x0661dc7f751df3c97b1303a78d310d09d7cf32c24df5404136c6275a0385d172"),
		},
		{
			// Node 3
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x19dec24A8748c117b102Bb29418F36c45E8C94f1").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0x2fD8930F52bD73Eb01C78b375E8449D6c107170c"),
				OffchainPublicKey: hexutil.MustDecode("0xccc929da9f3185f018c357a14d427cb9c982e981e3d4e20c391cbfb13d9fbb81"),
				PeerID:            "12D3KooWEC7dxiVkSRTCbFV72R4MSn2EZhDtnH7sH5mtYZifzqCW",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x3c21f181098f39d854cc77a4189b3a56b37bee7fec2386abe04e1e36b9177d15"),
		},
		{
			// Node 4
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x257ca0ff00204861bbeb626d70a733ece8dc71fa").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0x338820995b4772fAafCEd3bF56824D4b7a6996De"),
				OffchainPublicKey: hexutil.MustDecode("0x2b6fe2d95b217e93da7192bc495828bd5a7c8fc5e7deee919a21c19bc4b951c7"),
				PeerID:            "12D3KooWAyafDntpPKSnGeT4ybu7onfDtUAe54LNzaJGKGnfBx6c",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0xd14d160383b80e13dff1130fcdaed3afd54eabbb1f1c1136d3ea6b77e802744b"),
		},
	}
	// Change the offramp config
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgs(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		encoded,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		1, // faults
		nil,
	)
	_, err = offRamp.SetConfig(user, signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err)
}

func stringTo32Bytes(s string) [32]byte {
	var b [32]byte
	copy(b[:], hexutil.MustDecode(s))
	return b
}

func deployAfn(t *testing.T, user *bind.TransactOpts, chain *backends.SimulatedBackend) common.Address {
	afnSourceAddress, _, _, err := afn_contract.DeployAFNContract(
		user,
		chain,
		[]common.Address{user.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	chain.Commit()
	return afnSourceAddress
}
