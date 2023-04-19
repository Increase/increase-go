package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type AccountTransfer struct {
	// The account transfer's identifier.
	ID string `json:"id,required"`
	// The transfer amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency AccountTransferCurrency `json:"currency,required"`
	// The destination account's identifier.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The ID for the transaction receiving the transfer.
	DestinationTransactionID string `json:"destination_transaction_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description that will show on the transactions.
	Description string `json:"description,required"`
	// The transfer's network.
	Network AccountTransferNetwork `json:"network,required"`
	// The lifecycle status of the transfer.
	Status AccountTransferStatus `json:"status,required"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval AccountTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation AccountTransferCancellation `json:"cancellation,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `account_transfer`.
	Type AccountTransferType `json:"type,required"`
	JSON AccountTransferJSON
}

type AccountTransferJSON struct {
	ID                       pjson.Metadata
	Amount                   pjson.Metadata
	AccountID                pjson.Metadata
	Currency                 pjson.Metadata
	DestinationAccountID     pjson.Metadata
	DestinationTransactionID pjson.Metadata
	CreatedAt                pjson.Metadata
	Description              pjson.Metadata
	Network                  pjson.Metadata
	Status                   pjson.Metadata
	TransactionID            pjson.Metadata
	Approval                 pjson.Metadata
	Cancellation             pjson.Metadata
	Type                     pjson.Metadata
	Raw                      []byte
	Extras                   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountTransferCurrency string

const (
	AccountTransferCurrencyCad AccountTransferCurrency = "CAD"
	AccountTransferCurrencyChf AccountTransferCurrency = "CHF"
	AccountTransferCurrencyEur AccountTransferCurrency = "EUR"
	AccountTransferCurrencyGbp AccountTransferCurrency = "GBP"
	AccountTransferCurrencyJpy AccountTransferCurrency = "JPY"
	AccountTransferCurrencyUsd AccountTransferCurrency = "USD"
)

type AccountTransferNetwork string

const (
	AccountTransferNetworkAccount AccountTransferNetwork = "account"
)

type AccountTransferStatus string

const (
	AccountTransferStatusPendingApproval AccountTransferStatus = "pending_approval"
	AccountTransferStatusCanceled        AccountTransferStatus = "canceled"
	AccountTransferStatusComplete        AccountTransferStatus = "complete"
)

type AccountTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       AccountTransferApprovalJSON
}

type AccountTransferApprovalJSON struct {
	ApprovedAt pjson.Metadata
	ApprovedBy pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferApproval using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       AccountTransferCancellationJSON
}

type AccountTransferCancellationJSON struct {
	CanceledAt pjson.Metadata
	CanceledBy pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferCancellation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountTransferType string

const (
	AccountTransferTypeAccountTransfer AccountTransferType = "account_transfer"
)

type AccountTransferListResponse struct {
	// The contents of the list.
	Data []AccountTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       AccountTransferListResponseJSON
}

type AccountTransferListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
