package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	chainselectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
)

type Address common.Address
type NetworkSelector uint64

func (n NetworkSelector) Type() NetworkType {
	chainID, err := chainselectors.ChainIdFromSelector(uint64(n))
	isEvm := err == nil && chainID >= 0
	if isEvm {
		return NetworkTypeEvm
	}

	return NetworkTypeUnknown
}

type NetworkType string

var (
	NetworkTypeUnknown = NetworkType("Unknown")
	NetworkTypeEvm     = NetworkType(relay.EVM)
)

type Transfer struct {
	From   NetworkSelector
	To     NetworkSelector
	Amount *big.Int
	Date   time.Time
	//ID     uint64
}

func NewTransfer(from, to NetworkSelector, date time.Time, amount *big.Int) Transfer {
	return Transfer{
		From:   from,
		To:     to,
		Date:   date,
		Amount: amount,
	}
}

func (t Transfer) Equals(other Transfer) bool {
	return t.From == other.From &&
		t.To == other.To &&
		t.Amount.Cmp(other.Amount) == 0 &&
		t.Date.Equal(other.Date)
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
	NetworkID               NetworkSelector
}

func NewReportMetadata(transfers []Transfer, lmAddr Address, networkID NetworkSelector) ReportMetadata {
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
