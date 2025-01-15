package generated

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
)

// AbigenLog is an interface for abigen generated log topics
type AbigenLog interface {
	Topic() common.Hash
}

func IsZKSync(backend bind.ContractBackend) bool {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return false
	}
	zkclient := zkSyncClient.NewClient(client.Client())
	// we dont care about the address
	// we only care if the method is available
	_, err := zkclient.MainContractAddress(context.Background())
	return err == nil
}

type Transaction struct {
	*types.Transaction
	HashZks common.Hash
}

func (tx *Transaction) Hash() common.Hash {
	return tx.HashZks
}

func ConvertZkTxToEthTx(resp zktypes.TransactionResponse) *Transaction {
	// make this legacy fee ?
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := Transaction{Transaction: tx, HashZks: resp.Hash}
	return &customTransaction
}

func getZKAuthFromEthAuth(auth *bind.TransactOpts) *zkSyncAccounts.TransactOpts {
	return &zkSyncAccounts.TransactOpts{
		Nonce:     auth.Nonce,
		Value:     auth.Value,
		GasPrice:  auth.GasPrice,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
		GasLimit:  auth.GasLimit,
	}
}

func DeployContract(auth *bind.TransactOpts, contractAbi *abi.ABI, contractBytes []byte, backend bind.ContractBackend, params ...interface{}) (common.Address, *Transaction, *bind.BoundContract, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, fmt.Errorf("backend is not an *ethclient.Client")
	}
	zkclient := zkSyncClient.NewClient(client.Client())

	walletValue := auth.Context.Value("wallet")
	wallet, ok := walletValue.(*zkSyncAccounts.Wallet)
	if !ok || wallet == nil {
		return common.Address{}, nil, nil, fmt.Errorf("wallet not found in context or invalid type")
	}

	constructor, _ := contractAbi.Pack("", params...)

	hash, err := wallet.DeployWithCreate(getZKAuthFromEthAuth(auth), zkSyncAccounts.CreateTransaction{
		Bytecode: contractBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("Error deploying contract: %w", err)
	}

	receipt, err := zkclient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("Error waiting for contract deployment: %w", err)
	}

	tx, _, err := zkclient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("Error getting transaction by hash: %w", err)
	}

	address := receipt.ContractAddress
	contractBind := bind.NewBoundContract(address, *contractAbi, backend, backend, backend)
	return address, ConvertZkTxToEthTx(*tx), contractBind, nil
}
