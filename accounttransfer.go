// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// AccountTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTransferService] method instead.
type AccountTransferService struct {
	Options []option.RequestOption
}

// NewAccountTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTransferService(opts ...option.RequestOption) (r *AccountTransferService) {
	r = &AccountTransferService{}
	r.Options = opts
	return
}

// Create an Account Transfer
func (r *AccountTransferService) New(ctx context.Context, body AccountTransferNewParams, opts ...option.RequestOption) (res *AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Account Transfer
func (r *AccountTransferService) Get(ctx context.Context, accountTransferID string, opts ...option.RequestOption) (res *AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s", accountTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Account Transfers
func (r *AccountTransferService) List(ctx context.Context, query AccountTransferListParams, opts ...option.RequestOption) (res *pagination.Page[AccountTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_transfers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Account Transfers
func (r *AccountTransferService) ListAutoPaging(ctx context.Context, query AccountTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AccountTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve an Account Transfer
func (r *AccountTransferService) Approve(ctx context.Context, accountTransferID string, opts ...option.RequestOption) (res *AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s/approve", accountTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel an Account Transfer
func (r *AccountTransferService) Cancel(ctx context.Context, accountTransferID string, opts ...option.RequestOption) (res *AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s/cancel", accountTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Account transfers move funds between your own accounts at Increase.
type AccountTransfer struct {
	// The account transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The transfer amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval AccountTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation AccountTransferCancellation `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency AccountTransferCurrency `json:"currency,required"`
	// The description that will show on the transactions.
	Description string `json:"description,required"`
	// The destination account's identifier.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The ID for the transaction receiving the transfer.
	DestinationTransactionID string `json:"destination_transaction_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The transfer's network.
	Network AccountTransferNetwork `json:"network,required"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The lifecycle status of the transfer.
	Status AccountTransferStatus `json:"status,required"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `account_transfer`.
	Type AccountTransferType `json:"type,required"`
	JSON accountTransferJSON `json:"-"`
}

// accountTransferJSON contains the JSON metadata for the struct [AccountTransfer]
type accountTransferJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	Amount                   apijson.Field
	Approval                 apijson.Field
	Cancellation             apijson.Field
	CreatedAt                apijson.Field
	Currency                 apijson.Field
	Description              apijson.Field
	DestinationAccountID     apijson.Field
	DestinationTransactionID apijson.Field
	IdempotencyKey           apijson.Field
	Network                  apijson.Field
	PendingTransactionID     apijson.Field
	Status                   apijson.Field
	TransactionID            apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransferJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type AccountTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                      `json:"approved_by,required,nullable"`
	JSON       accountTransferApprovalJSON `json:"-"`
}

// accountTransferApprovalJSON contains the JSON metadata for the struct
// [AccountTransferApproval]
type accountTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type AccountTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                          `json:"canceled_by,required,nullable"`
	JSON       accountTransferCancellationJSON `json:"-"`
}

// accountTransferCancellationJSON contains the JSON metadata for the struct
// [AccountTransferCancellation]
type accountTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type AccountTransferCurrency string

const (
	// Canadian Dollar (CAD)
	AccountTransferCurrencyCad AccountTransferCurrency = "CAD"
	// Swiss Franc (CHF)
	AccountTransferCurrencyChf AccountTransferCurrency = "CHF"
	// Euro (EUR)
	AccountTransferCurrencyEur AccountTransferCurrency = "EUR"
	// British Pound (GBP)
	AccountTransferCurrencyGbp AccountTransferCurrency = "GBP"
	// Japanese Yen (JPY)
	AccountTransferCurrencyJpy AccountTransferCurrency = "JPY"
	// US Dollar (USD)
	AccountTransferCurrencyUsd AccountTransferCurrency = "USD"
)

func (r AccountTransferCurrency) IsKnown() bool {
	switch r {
	case AccountTransferCurrencyCad, AccountTransferCurrencyChf, AccountTransferCurrencyEur, AccountTransferCurrencyGbp, AccountTransferCurrencyJpy, AccountTransferCurrencyUsd:
		return true
	}
	return false
}

// The transfer's network.
type AccountTransferNetwork string

const (
	AccountTransferNetworkAccount AccountTransferNetwork = "account"
)

func (r AccountTransferNetwork) IsKnown() bool {
	switch r {
	case AccountTransferNetworkAccount:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type AccountTransferStatus string

const (
	// The transfer is pending approval.
	AccountTransferStatusPendingApproval AccountTransferStatus = "pending_approval"
	// The transfer has been canceled.
	AccountTransferStatusCanceled AccountTransferStatus = "canceled"
	// The transfer has been completed.
	AccountTransferStatusComplete AccountTransferStatus = "complete"
)

func (r AccountTransferStatus) IsKnown() bool {
	switch r {
	case AccountTransferStatusPendingApproval, AccountTransferStatusCanceled, AccountTransferStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `account_transfer`.
type AccountTransferType string

const (
	AccountTransferTypeAccountTransfer AccountTransferType = "account_transfer"
)

func (r AccountTransferType) IsKnown() bool {
	switch r {
	case AccountTransferTypeAccountTransfer:
		return true
	}
	return false
}

type AccountTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The description you choose to give the transfer.
	Description param.Field[string] `json:"description,required"`
	// The identifier for the account that will receive the transfer.
	DestinationAccountID param.Field[string] `json:"destination_account_id,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
}

func (r AccountTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTransferListParams struct {
	// Filter Account Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                             `query:"account_id"`
	CreatedAt param.Field[AccountTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AccountTransferListParams]'s query parameters as
// `url.Values`.
func (r AccountTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [AccountTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r AccountTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
