package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address

type NetworkID uint64

type Transfer struct {
	From   NetworkID
	To     NetworkID
	Amount *big.Int
}

type PendingTransfer struct {
	Transfer
	Status TransferStatus
}

func (p PendingTransfer) Hash() ([32]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return [32]byte{}, fmt.Errorf("marshal: %w", err)
	}
	return sha256.Sum256(b), nil
}

func NewPendingTransfer(tr Transfer) PendingTransfer {
	return PendingTransfer{
		Transfer: tr,
		Status:   TransferStatusNotReady,
	}
}

type TransferStatus string

const (
	TransferStatusNotReady  = "not-ready"
	TransferStatusReady     = "ready"
	TransferStatusFinalized = "finalized"
	TransferStatusExecuted  = "executed"
)

func (t Transfer) String() string {
	return fmt.Sprintf("%v->%v %s", t.From, t.To, t.Amount.String())
}

type ReportMetadata struct {
	Transfers               []Transfer
	LiquidityManagerAddress Address
	NetworkID               NetworkID
}

func NewReportMetadata(transfers []Transfer, lmAddr Address, networkID NetworkID) ReportMetadata {
	return ReportMetadata{
		Transfers:               transfers,
		LiquidityManagerAddress: lmAddr,
		NetworkID:               networkID,
	}
}

func (r ReportMetadata) Encode() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		panic(fmt.Errorf("report meta %#v encoding unexpected internal error: %w", r, err))
	}
	return b
}

func DecodeReportMetadata(b []byte) (ReportMetadata, error) {
	var meta ReportMetadata
	err := json.Unmarshal(b, &meta)
	return meta, err
}
