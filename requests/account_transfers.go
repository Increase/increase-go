package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type AccountTransfer struct {
	// The account transfer's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The transfer amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The Account to which the transfer belongs.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[AccountTransferCurrency] `json:"currency,required"`
	// The destination account's identifier.
	DestinationAccountID fields.Field[string] `json:"destination_account_id,required"`
	// The ID for the transaction receiving the transfer.
	DestinationTransactionID fields.Field[string] `json:"destination_transaction_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The description that will show on the transactions.
	Description fields.Field[string] `json:"description,required"`
	// The transfer's network.
	Network fields.Field[AccountTransferNetwork] `json:"network,required"`
	// The lifecycle status of the transfer.
	Status fields.Field[AccountTransferStatus] `json:"status,required"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID fields.Field[string] `json:"template_id,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval fields.Field[AccountTransferApproval] `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation fields.Field[AccountTransferCancellation] `json:"cancellation,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `account_transfer`.
	Type fields.Field[AccountTransferType] `json:"type,required"`
}

// MarshalJSON serializes AccountTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountTransfer) String() (result string) {
	return fmt.Sprintf("&AccountTransfer{ID:%s Amount:%s AccountID:%s Currency:%s DestinationAccountID:%s DestinationTransactionID:%s CreatedAt:%s Description:%s Network:%s Status:%s TemplateID:%s TransactionID:%s Approval:%s Cancellation:%s Type:%s}", r.ID, r.Amount, r.AccountID, r.Currency, r.DestinationAccountID, r.DestinationTransactionID, r.CreatedAt, r.Description, r.Network, r.Status, r.TemplateID, r.TransactionID, r.Approval, r.Cancellation, r.Type)
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
	AccountTransferStatusPendingSubmission AccountTransferStatus = "pending_submission"
	AccountTransferStatusPendingApproval   AccountTransferStatus = "pending_approval"
	AccountTransferStatusCanceled          AccountTransferStatus = "canceled"
	AccountTransferStatusRequiresAttention AccountTransferStatus = "requires_attention"
	AccountTransferStatusFlaggedByOperator AccountTransferStatus = "flagged_by_operator"
	AccountTransferStatusComplete          AccountTransferStatus = "complete"
)

type AccountTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt fields.Field[time.Time] `json:"approved_at,required" format:"date-time"`
}

// MarshalJSON serializes AccountTransferApproval into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferApproval) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountTransferApproval) String() (result string) {
	return fmt.Sprintf("&AccountTransferApproval{ApprovedAt:%s}", r.ApprovedAt)
}

type AccountTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt fields.Field[time.Time] `json:"canceled_at,required" format:"date-time"`
}

// MarshalJSON serializes AccountTransferCancellation into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferCancellation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountTransferCancellation) String() (result string) {
	return fmt.Sprintf("&AccountTransferCancellation{CanceledAt:%s}", r.CanceledAt)
}

type AccountTransferType string

const (
	AccountTransferTypeAccountTransfer AccountTransferType = "account_transfer"
)

type CreateAnAccountTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description you choose to give the transfer.
	Description fields.Field[string] `json:"description,required"`
	// The identifier for the account that will receive the transfer.
	DestinationAccountID fields.Field[string] `json:"destination_account_id,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CreateAnAccountTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnAccountTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountTransferParameters{AccountID:%s Amount:%s Description:%s DestinationAccountID:%s RequireApproval:%s}", r.AccountID, r.Amount, r.Description, r.DestinationAccountID, r.RequireApproval)
}

type AccountTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Account Transfers to those that originated from the specified Account.
	AccountID fields.Field[string]                             `query:"account_id"`
	CreatedAt fields.Field[AccountTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes AccountTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountTransferListParams) String() (result string) {
	return fmt.Sprintf("&AccountTransferListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type AccountTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes AccountTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *AccountTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&AccountTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}
