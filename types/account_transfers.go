package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type AccountTransfer struct {
	// The account transfer's identifier.
	ID *string `pjson:"id"`
	// The transfer amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The Account to which the transfer belongs.
	AccountID *string `pjson:"account_id"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *AccountTransferCurrency `pjson:"currency"`
	// The destination account's identifier.
	DestinationAccountID *string `pjson:"destination_account_id"`
	// The ID for the transaction receiving the transfer.
	DestinationTransactionID *string `pjson:"destination_transaction_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// The description that will show on the transactions.
	Description *string `pjson:"description"`
	// The transfer's network.
	Network *AccountTransferNetwork `pjson:"network"`
	// The lifecycle status of the transfer.
	Status *AccountTransferStatus `pjson:"status"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `pjson:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionID *string `pjson:"transaction_id"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *AccountTransferApproval `pjson:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *AccountTransferCancellation `pjson:"cancellation"`
	// A constant representing the object's type. For this resource it will always be
	// `account_transfer`.
	Type       *AccountTransferType   `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The account transfer's identifier.
func (r AccountTransfer) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The transfer amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r AccountTransfer) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The Account to which the transfer belongs.
func (r AccountTransfer) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r AccountTransfer) GetCurrency() (Currency AccountTransferCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The destination account's identifier.
func (r AccountTransfer) GetDestinationAccountID() (DestinationAccountID string) {
	if r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// The ID for the transaction receiving the transfer.
func (r AccountTransfer) GetDestinationTransactionID() (DestinationTransactionID string) {
	if r.DestinationTransactionID != nil {
		DestinationTransactionID = *r.DestinationTransactionID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r AccountTransfer) GetCreatedAt() (CreatedAt string) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The description that will show on the transactions.
func (r AccountTransfer) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The transfer's network.
func (r AccountTransfer) GetNetwork() (Network AccountTransferNetwork) {
	if r.Network != nil {
		Network = *r.Network
	}
	return
}

// The lifecycle status of the transfer.
func (r AccountTransfer) GetStatus() (Status AccountTransferStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r AccountTransfer) GetTemplateID() (TemplateID string) {
	if r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction funding the transfer.
func (r AccountTransfer) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r AccountTransfer) GetApproval() (Approval AccountTransferApproval) {
	if r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r AccountTransfer) GetCancellation() (Cancellation AccountTransferCancellation) {
	if r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_transfer`.
func (r AccountTransfer) GetType() (Type AccountTransferType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r AccountTransfer) String() (result string) {
	return fmt.Sprintf("&AccountTransfer{ID:%s Amount:%s AccountID:%s Currency:%s DestinationAccountID:%s DestinationTransactionID:%s CreatedAt:%s Description:%s Network:%s Status:%s TemplateID:%s TransactionID:%s Approval:%s Cancellation:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.Amount), core.FmtP(r.AccountID), core.FmtP(r.Currency), core.FmtP(r.DestinationAccountID), core.FmtP(r.DestinationTransactionID), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.Network), core.FmtP(r.Status), core.FmtP(r.TemplateID), core.FmtP(r.TransactionID), r.Approval, r.Cancellation, core.FmtP(r.Type))
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
	ApprovedAt *string                `pjson:"approved_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferApproval using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountTransferApproval into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferApproval) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r AccountTransferApproval) GetApprovedAt() (ApprovedAt string) {
	if r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}

func (r AccountTransferApproval) String() (result string) {
	return fmt.Sprintf("&AccountTransferApproval{ApprovedAt:%s}", core.FmtP(r.ApprovedAt))
}

type AccountTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string                `pjson:"canceled_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferCancellation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountTransferCancellation into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferCancellation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r AccountTransferCancellation) GetCanceledAt() (CanceledAt string) {
	if r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}

func (r AccountTransferCancellation) String() (result string) {
	return fmt.Sprintf("&AccountTransferCancellation{CanceledAt:%s}", core.FmtP(r.CanceledAt))
}

type AccountTransferType string

const (
	AccountTransferTypeAccountTransfer AccountTransferType = "account_transfer"
)

type CreateAnAccountTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID *string `pjson:"account_id"`
	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The description you choose to give the transfer.
	Description *string `pjson:"description"`
	// The identifier for the account that will receive the transfer.
	DestinationAccountID *string `pjson:"destination_account_id"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval *bool                  `pjson:"require_approval"`
	jsonFields      map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnAccountTransferParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CreateAnAccountTransferParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnAccountTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the account that will send the transfer.
func (r CreateAnAccountTransferParameters) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The transfer amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r CreateAnAccountTransferParameters) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description you choose to give the transfer.
func (r CreateAnAccountTransferParameters) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier for the account that will receive the transfer.
func (r CreateAnAccountTransferParameters) GetDestinationAccountID() (DestinationAccountID string) {
	if r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// Whether the transfer requires explicit approval via the dashboard or API.
func (r CreateAnAccountTransferParameters) GetRequireApproval() (RequireApproval bool) {
	if r.RequireApproval != nil {
		RequireApproval = *r.RequireApproval
	}
	return
}

func (r CreateAnAccountTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountTransferParameters{AccountID:%s Amount:%s Description:%s DestinationAccountID:%s RequireApproval:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.DestinationAccountID), core.FmtP(r.RequireApproval))
}

type AccountTransferListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Account Transfers to those that originated from the specified Account.
	AccountID  *string                              `query:"account_id"`
	CreatedAt  *AccountTransfersListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountTransferListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountTransferListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r AccountTransferListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r AccountTransferListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Account Transfers to those that originated from the specified Account.
func (r AccountTransferListParams) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r AccountTransferListParams) GetCreatedAt() (CreatedAt AccountTransfersListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r AccountTransferListParams) String() (result string) {
	return fmt.Sprintf("&AccountTransferListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), r.CreatedAt)
}

type AccountTransfersListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// AccountTransfersListParamsCreatedAt using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AccountTransfersListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountTransfersListParamsCreatedAt into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *AccountTransfersListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountTransfersListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *AccountTransfersListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r AccountTransfersListParamsCreatedAt) GetAfter() (After string) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r AccountTransfersListParamsCreatedAt) GetBefore() (Before string) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r AccountTransfersListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r AccountTransfersListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r AccountTransfersListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&AccountTransfersListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type AccountTransferList struct {
	// The contents of the list.
	Data *[]AccountTransfer `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountTransferList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountTransferList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountTransferList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountTransferList into a url.Values of the query
// parameters associated with this value
func (r *AccountTransferList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r AccountTransferList) GetData() (Data []AccountTransfer) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r AccountTransferList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r AccountTransferList) String() (result string) {
	return fmt.Sprintf("&AccountTransferList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type AccountTransfersPage struct {
	*pagination.Page[AccountTransfer]
}

func (r *AccountTransfersPage) AccountTransfer() *AccountTransfer {
	return r.Current()
}

func (r *AccountTransfersPage) NextPage() (*AccountTransfersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &AccountTransfersPage{page}, nil
	}
}
