//go:build solana
// +build solana

package deployment

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func confirmTx(c *rpc.Client, ctx context.Context, sig solana.Signature) error {
	var confirmed bool
	for i := 0; i < 30; i++ {
		block, err := c.GetConfirmedTransactionWithOpts(ctx, sig, &rpc.GetTransactionOpts{
			// Must be finalized for state to actually change.
			Commitment: rpc.CommitmentConfirmed,
		})
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("Confirmed!", block)
		confirmed = true
		break
	}
	if !confirmed {
		return fmt.Errorf("transaction not confirmed")
	}
	return nil
}

func TestSolanaCreateAccount(t *testing.T) {
	t.Skip()
	// Note that uploading a program is fairly complex:
	// 2 options
	// - Replicate https://github.com/solana-labs/solana/blob/7409d9d2687fba21078a745842c25df805cdf105/cli/src/program.rs#L2086
	// using solana-go. Not particularly difficult but a little bit of work.
	// - Shell out to solana go for deployment.
	ctx := context.Background()

	// solana-test-validator
	client := rpc.New("http://127.0.0.1:8899")

	// Fund a deployer
	deployer := solana.NewWallet()
	airdrop, err := client.RequestAirdrop(ctx, deployer.PublicKey(), 10*solana.LAMPORTS_PER_SOL, rpc.CommitmentConfirmed)
	require.NoError(t, err)
	require.NoError(t, confirmTx(client, ctx, airdrop))
	b, err := client.GetBalance(ctx, deployer.PublicKey(), rpc.CommitmentFinalized)
	require.NoError(t, err)
	assert.Equal(t, 10*solana.LAMPORTS_PER_SOL, b.Value)

	// Create
	programAccount := solana.NewWallet()
	createInst, err := system.NewCreateAccountInstruction(
		solana.LAMPORTS_PER_SOL, // If you don't fund the account it won't exist
		9,
		solana.SystemProgramID,
		deployer.PublicKey(),
		programAccount.PublicKey()).ValidateAndBuild()
	require.NoError(t, err)

	bh, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	require.NoError(t, err)
	programAccountTx, err := solana.NewTransaction(
		[]solana.Instruction{
			createInst, // Create the program account
		},
		bh.Value.Blockhash,
	)
	require.NoError(t, err)
	fmt.Println(programAccountTx)
	sig, err := programAccountTx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(deployer.PublicKey()) {
			return &deployer.PrivateKey
		} else if key.Equals(programAccount.PublicKey()) {
			return &programAccount.PrivateKey
		}
		panic("unauthorized sign")
	})
	require.NoError(t, err)
	programAccountTx.Signatures = sig
	require.NoError(t, programAccountTx.VerifySignatures())
	txSig, err := client.SendTransaction(ctx, programAccountTx)
	require.NoError(t, err)
	fmt.Printf("Program deployment transaction signature: %s\n", txSig.String())
	// Program account will not exist without this
	require.NoError(t, confirmTx(client, ctx, txSig))
	acct, err := client.GetAccountInfo(ctx, programAccount.PublicKey())
	require.NoError(t, err)
	fmt.Println(acct.Value.Data)
}
