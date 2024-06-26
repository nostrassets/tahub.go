package models

import (
	"time"
)

const (
	EntryTypeIncoming           = "incoming"
	EntryTypeOutgoing           = "outgoing"
	EntryTypeFee                = "fee"
	EntryTypeFeeReserve         = "fee_reserve"
	EntryTypeServiceFee         = "service_fee"
	EntryTypeServiceFeeReversal = "service_fee_reversal"
	EntryTypeFeeReserveReversal = "fee_reserve_reversal"
	EntryTypeOutgoingReversal   = "outgoing_reversal"

	BroadcastStatePending   = "pending"
	BroadcastStateBroadcast = "broadcast"
	TahubInternalOutpoint   = "tahub_internal_outpoint"
	TahubInternalComplete   = "tahub_internal_complete"
)

// TransactionEntry : Transaction Entries Model
type TransactionEntry struct {
	ID              int64             `bun:",pk,autoincrement"`
	UserID          int64             `bun:",notnull"`
	User            *User             `bun:"rel:belongs-to,join:user_id=id"`
	InvoiceID       int64             `bun:",nullzero"`
	Invoice         *Invoice          `bun:"rel:belongs-to,join:invoice_id=id"`
	ParentID        int64             `bun:",nullzero"`
	Parent          *TransactionEntry `bun:"rel:belongs-to"`
	TaAssetID       string			  `bun:",nullzero"`
	Asset		    *Asset            `bun:"rel:belongs-to,join:ta_asset_id=ta_asset_id"`
	CreditAccountID int64             `bun:",notnull"`
	FeeReserve      *TransactionEntry `bun:"rel:belongs-to"`
	ServiceFee      *TransactionEntry `bun:"rel:belongs-to"`
	CreditAccount   *Account          `bun:"rel:belongs-to,join:credit_account_id=id"`
	DebitAccountID  int64             `bun:",notnull"`
	DebitAccount    *Account          `bun:"rel:belongs-to,join:debit_account_id=id"`
	Amount          int64             `bun:",notnull"`
	CreatedAt       time.Time         `bun:",nullzero,notnull,default:current_timestamp"`
	EntryType       string
	Outpoint        string 
	BroadcastState  string
}
