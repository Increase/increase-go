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

// InboundCheckDepositService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewInboundCheckDepositService]
// method instead.
type InboundCheckDepositService struct {
	Options []option.RequestOption
}

// NewInboundCheckDepositService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboundCheckDepositService(opts ...option.RequestOption) (r *InboundCheckDepositService) {
	r = &InboundCheckDepositService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Check Deposit
func (r *InboundCheckDepositService) Get(ctx context.Context, inboundCheckDepositID string, opts ...option.RequestOption) (res *InboundCheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_check_deposits/%s", inboundCheckDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Check Deposits
func (r *InboundCheckDepositService) List(ctx context.Context, query InboundCheckDepositListParams, opts ...option.RequestOption) (res *pagination.Page[InboundCheckDeposit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_check_deposits"
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

// List Inbound Check Deposits
func (r *InboundCheckDepositService) ListAutoPaging(ctx context.Context, query InboundCheckDepositListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundCheckDeposit] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Inbound Check Deposits are records of third-parties attempting to deposit checks
// against your account.
type InboundCheckDeposit struct {
	// The deposit's identifier.
	ID string `json:"id,required"`
	// The Account the check is being deposited against.
	AccountID string `json:"account_id,required"`
	// The Account Number the check is being deposited against.
	AccountNumberID string `json:"account_number_id,required"`
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number,required,nullable"`
	// The check number printed on the check being deposited.
	CheckNumber string `json:"check_number,required,nullable"`
	// If this deposit is for an existing Check Transfer, the identifier of that Check
	// Transfer.
	CheckTransferID string `json:"check_transfer_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the deposit was attempted.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency InboundCheckDepositCurrency `json:"currency,required"`
	// If the deposit attempt has been rejected, the identifier of the Declined
	// Transaction object created as a result of the failed deposit.
	DeclinedTransactionID string `json:"declined_transaction_id,required,nullable"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The status of the Inbound Check Deposit.
	Status InboundCheckDepositStatus `json:"status,required"`
	// If the deposit attempt has been accepted, the identifier of the Transaction
	// object created as a result of the successful deposit.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_check_deposit`.
	Type InboundCheckDepositType `json:"type,required"`
	JSON inboundCheckDepositJSON `json:"-"`
}

// inboundCheckDepositJSON contains the JSON metadata for the struct
// [InboundCheckDeposit]
type inboundCheckDepositJSON struct {
	ID                              apijson.Field
	AccountID                       apijson.Field
	AccountNumberID                 apijson.Field
	Amount                          apijson.Field
	BackImageFileID                 apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	CheckNumber                     apijson.Field
	CheckTransferID                 apijson.Field
	CreatedAt                       apijson.Field
	Currency                        apijson.Field
	DeclinedTransactionID           apijson.Field
	FrontImageFileID                apijson.Field
	Status                          apijson.Field
	TransactionID                   apijson.Field
	Type                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *InboundCheckDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundCheckDepositJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
type InboundCheckDepositCurrency string

const (
	// Canadian Dollar (CAD)
	InboundCheckDepositCurrencyCad InboundCheckDepositCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundCheckDepositCurrencyChf InboundCheckDepositCurrency = "CHF"
	// Euro (EUR)
	InboundCheckDepositCurrencyEur InboundCheckDepositCurrency = "EUR"
	// British Pound (GBP)
	InboundCheckDepositCurrencyGbp InboundCheckDepositCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundCheckDepositCurrencyJpy InboundCheckDepositCurrency = "JPY"
	// US Dollar (USD)
	InboundCheckDepositCurrencyUsd InboundCheckDepositCurrency = "USD"
)

func (r InboundCheckDepositCurrency) IsKnown() bool {
	switch r {
	case InboundCheckDepositCurrencyCad, InboundCheckDepositCurrencyChf, InboundCheckDepositCurrencyEur, InboundCheckDepositCurrencyGbp, InboundCheckDepositCurrencyJpy, InboundCheckDepositCurrencyUsd:
		return true
	}
	return false
}

// The status of the Inbound Check Deposit.
type InboundCheckDepositStatus string

const (
	// The Inbound Check Deposit is pending.
	InboundCheckDepositStatusPending InboundCheckDepositStatus = "pending"
	// The Inbound Check Deposit was accepted.
	InboundCheckDepositStatusAccepted InboundCheckDepositStatus = "accepted"
	// The Inbound Check Deposit was rejected.
	InboundCheckDepositStatusRejected InboundCheckDepositStatus = "rejected"
)

func (r InboundCheckDepositStatus) IsKnown() bool {
	switch r {
	case InboundCheckDepositStatusPending, InboundCheckDepositStatusAccepted, InboundCheckDepositStatusRejected:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_check_deposit`.
type InboundCheckDepositType string

const (
	InboundCheckDepositTypeInboundCheckDeposit InboundCheckDepositType = "inbound_check_deposit"
)

func (r InboundCheckDepositType) IsKnown() bool {
	switch r {
	case InboundCheckDepositTypeInboundCheckDeposit:
		return true
	}
	return false
}

type InboundCheckDepositListParams struct {
	// Filter Inbound Check Deposits to those belonging to the specified Account.
	AccountID param.Field[string]                                 `query:"account_id"`
	CreatedAt param.Field[InboundCheckDepositListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundCheckDepositListParams]'s query parameters as
// `url.Values`.
func (r InboundCheckDepositListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundCheckDepositListParamsCreatedAt struct {
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

// URLQuery serializes [InboundCheckDepositListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r InboundCheckDepositListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
