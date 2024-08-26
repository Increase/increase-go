// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
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

// InboundRealTimePaymentsTransferService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundRealTimePaymentsTransferService] method instead.
type InboundRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewInboundRealTimePaymentsTransferService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInboundRealTimePaymentsTransferService(opts ...option.RequestOption) (r *InboundRealTimePaymentsTransferService) {
	r = &InboundRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Real-Time Payments Transfer
func (r *InboundRealTimePaymentsTransferService) Get(ctx context.Context, inboundRealTimePaymentsTransferID string, opts ...option.RequestOption) (res *InboundRealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if inboundRealTimePaymentsTransferID == "" {
		err = errors.New("missing required inbound_real_time_payments_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_real_time_payments_transfers/%s", inboundRealTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Real-Time Payments Transfers
func (r *InboundRealTimePaymentsTransferService) List(ctx context.Context, query InboundRealTimePaymentsTransferListParams, opts ...option.RequestOption) (res *pagination.Page[InboundRealTimePaymentsTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_real_time_payments_transfers"
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

// List Inbound Real-Time Payments Transfers
func (r *InboundRealTimePaymentsTransferService) ListAutoPaging(ctx context.Context, query InboundRealTimePaymentsTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundRealTimePaymentsTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// An Inbound Real-Time Payments Transfer is a Real-Time Payments transfer
// initiated outside of Increase to your account.
type InboundRealTimePaymentsTransfer struct {
	// The inbound Real-Time Payments transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer was sent.
	AccountID string `json:"account_id,required"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id,required"`
	// The amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency InboundRealTimePaymentsTransferCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The lifecycle status of the transfer.
	Status InboundRealTimePaymentsTransferStatus `json:"status,required"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer`.
	Type InboundRealTimePaymentsTransferType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferJSON `json:"-"`
}

// inboundRealTimePaymentsTransferJSON contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransfer]
type inboundRealTimePaymentsTransferJSON struct {
	ID                        apijson.Field
	AccountID                 apijson.Field
	AccountNumberID           apijson.Field
	Amount                    apijson.Field
	CreatedAt                 apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	RemittanceInformation     apijson.Field
	Status                    apijson.Field
	TransactionIdentification apijson.Field
	Type                      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsTransferJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real-Time Payments transfer.
type InboundRealTimePaymentsTransferCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferCurrencyCad InboundRealTimePaymentsTransferCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferCurrencyChf InboundRealTimePaymentsTransferCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferCurrencyEur InboundRealTimePaymentsTransferCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferCurrencyGbp InboundRealTimePaymentsTransferCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferCurrencyJpy InboundRealTimePaymentsTransferCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferCurrencyUsd InboundRealTimePaymentsTransferCurrency = "USD"
)

func (r InboundRealTimePaymentsTransferCurrency) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferCurrencyCad, InboundRealTimePaymentsTransferCurrencyChf, InboundRealTimePaymentsTransferCurrencyEur, InboundRealTimePaymentsTransferCurrencyGbp, InboundRealTimePaymentsTransferCurrencyJpy, InboundRealTimePaymentsTransferCurrencyUsd:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type InboundRealTimePaymentsTransferStatus string

const (
	// The transfer is pending confirmation.
	InboundRealTimePaymentsTransferStatusPendingConfirmation InboundRealTimePaymentsTransferStatus = "pending_confirmation"
	// Your webhook failed to respond to the transfer in time.
	InboundRealTimePaymentsTransferStatusTimedOut InboundRealTimePaymentsTransferStatus = "timed_out"
	// The transfer has been received successfully and is confirmed.
	InboundRealTimePaymentsTransferStatusConfirmed InboundRealTimePaymentsTransferStatus = "confirmed"
)

func (r InboundRealTimePaymentsTransferStatus) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferStatusPendingConfirmation, InboundRealTimePaymentsTransferStatusTimedOut, InboundRealTimePaymentsTransferStatusConfirmed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_transfer`.
type InboundRealTimePaymentsTransferType string

const (
	InboundRealTimePaymentsTransferTypeInboundRealTimePaymentsTransfer InboundRealTimePaymentsTransferType = "inbound_real_time_payments_transfer"
)

func (r InboundRealTimePaymentsTransferType) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferTypeInboundRealTimePaymentsTransfer:
		return true
	}
	return false
}

type InboundRealTimePaymentsTransferListParams struct {
	// Filter Inbound Real-Time Payments Transfers to those belonging to the specified
	// Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound Real-Time Payments Transfers to ones belonging to the specified
	// Account Number.
	AccountNumberID param.Field[string]                                             `query:"account_number_id"`
	CreatedAt       param.Field[InboundRealTimePaymentsTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundRealTimePaymentsTransferListParams]'s query
// parameters as `url.Values`.
func (r InboundRealTimePaymentsTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundRealTimePaymentsTransferListParamsCreatedAt struct {
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

// URLQuery serializes [InboundRealTimePaymentsTransferListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r InboundRealTimePaymentsTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
