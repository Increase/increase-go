// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// InboundFednowTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundFednowTransferService] method instead.
type InboundFednowTransferService struct {
	Options []option.RequestOption
}

// NewInboundFednowTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboundFednowTransferService(opts ...option.RequestOption) (r *InboundFednowTransferService) {
	r = &InboundFednowTransferService{}
	r.Options = opts
	return
}

// Retrieve an Inbound FedNow Transfer
func (r *InboundFednowTransferService) Get(ctx context.Context, inboundFednowTransferID string, opts ...option.RequestOption) (res *InboundFednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundFednowTransferID == "" {
		err = errors.New("missing required inbound_fednow_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_fednow_transfers/%s", inboundFednowTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound FedNow Transfers
func (r *InboundFednowTransferService) List(ctx context.Context, query InboundFednowTransferListParams, opts ...option.RequestOption) (res *InboundFednowTransferListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inbound_fednow_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An Inbound FedNow Transfer is a FedNow transfer initiated outside of Increase to
// your account.
type InboundFednowTransfer struct {
	// The inbound FedNow transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer was sent.
	AccountID string `json:"account_id,required"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id,required"`
	// The amount in USD cents.
	Amount int64 `json:"amount,required"`
	// If your transfer is confirmed, this will contain details of the confirmation.
	Confirmation InboundFednowTransferConfirmation `json:"confirmation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a FedNow transfer.
	Currency InboundFednowTransferCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// If your transfer is declined, this will contain details of the decline.
	Decline InboundFednowTransferDecline `json:"decline,required,nullable"`
	// The lifecycle status of the transfer.
	Status InboundFednowTransferStatus `json:"status,required"`
	// The identifier of the Transaction object created when the transfer was
	// confirmed.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_fednow_transfer`.
	Type InboundFednowTransferType `json:"type,required"`
	// Additional information included with the transfer.
	UnstructuredRemittanceInformation string                    `json:"unstructured_remittance_information,required,nullable"`
	JSON                              inboundFednowTransferJSON `json:"-"`
}

// inboundFednowTransferJSON contains the JSON metadata for the struct
// [InboundFednowTransfer]
type inboundFednowTransferJSON struct {
	ID                                apijson.Field
	AccountID                         apijson.Field
	AccountNumberID                   apijson.Field
	Amount                            apijson.Field
	Confirmation                      apijson.Field
	CreatedAt                         apijson.Field
	CreditorName                      apijson.Field
	Currency                          apijson.Field
	DebtorAccountNumber               apijson.Field
	DebtorName                        apijson.Field
	DebtorRoutingNumber               apijson.Field
	Decline                           apijson.Field
	Status                            apijson.Field
	TransactionID                     apijson.Field
	Type                              apijson.Field
	UnstructuredRemittanceInformation apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *InboundFednowTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundFednowTransferJSON) RawJSON() string {
	return r.raw
}

// If your transfer is confirmed, this will contain details of the confirmation.
type InboundFednowTransferConfirmation struct {
	// The identifier of the FedNow Transfer that led to this Transaction.
	TransferID  string                                `json:"transfer_id,required"`
	ExtraFields map[string]interface{}                `json:"-,extras"`
	JSON        inboundFednowTransferConfirmationJSON `json:"-"`
}

// inboundFednowTransferConfirmationJSON contains the JSON metadata for the struct
// [InboundFednowTransferConfirmation]
type inboundFednowTransferConfirmationJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundFednowTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundFednowTransferConfirmationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a FedNow transfer.
type InboundFednowTransferCurrency string

const (
	InboundFednowTransferCurrencyUsd InboundFednowTransferCurrency = "USD"
)

func (r InboundFednowTransferCurrency) IsKnown() bool {
	switch r {
	case InboundFednowTransferCurrencyUsd:
		return true
	}
	return false
}

// If your transfer is declined, this will contain details of the decline.
type InboundFednowTransferDecline struct {
	// Why the transfer was declined.
	Reason InboundFednowTransferDeclineReason `json:"reason,required"`
	// The identifier of the FedNow Transfer that led to this declined transaction.
	TransferID  string                           `json:"transfer_id,required"`
	ExtraFields map[string]interface{}           `json:"-,extras"`
	JSON        inboundFednowTransferDeclineJSON `json:"-"`
}

// inboundFednowTransferDeclineJSON contains the JSON metadata for the struct
// [InboundFednowTransferDecline]
type inboundFednowTransferDeclineJSON struct {
	Reason      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundFednowTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundFednowTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the transfer was declined.
type InboundFednowTransferDeclineReason string

const (
	InboundFednowTransferDeclineReasonAccountNumberCanceled InboundFednowTransferDeclineReason = "account_number_canceled"
	InboundFednowTransferDeclineReasonAccountNumberDisabled InboundFednowTransferDeclineReason = "account_number_disabled"
	InboundFednowTransferDeclineReasonAccountRestricted     InboundFednowTransferDeclineReason = "account_restricted"
	InboundFednowTransferDeclineReasonGroupLocked           InboundFednowTransferDeclineReason = "group_locked"
	InboundFednowTransferDeclineReasonEntityNotActive       InboundFednowTransferDeclineReason = "entity_not_active"
	InboundFednowTransferDeclineReasonFednowNotEnabled      InboundFednowTransferDeclineReason = "fednow_not_enabled"
)

func (r InboundFednowTransferDeclineReason) IsKnown() bool {
	switch r {
	case InboundFednowTransferDeclineReasonAccountNumberCanceled, InboundFednowTransferDeclineReasonAccountNumberDisabled, InboundFednowTransferDeclineReasonAccountRestricted, InboundFednowTransferDeclineReasonGroupLocked, InboundFednowTransferDeclineReasonEntityNotActive, InboundFednowTransferDeclineReasonFednowNotEnabled:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type InboundFednowTransferStatus string

const (
	InboundFednowTransferStatusPendingConfirming InboundFednowTransferStatus = "pending_confirming"
	InboundFednowTransferStatusTimedOut          InboundFednowTransferStatus = "timed_out"
	InboundFednowTransferStatusConfirmed         InboundFednowTransferStatus = "confirmed"
	InboundFednowTransferStatusDeclined          InboundFednowTransferStatus = "declined"
	InboundFednowTransferStatusRequiresAttention InboundFednowTransferStatus = "requires_attention"
)

func (r InboundFednowTransferStatus) IsKnown() bool {
	switch r {
	case InboundFednowTransferStatusPendingConfirming, InboundFednowTransferStatusTimedOut, InboundFednowTransferStatusConfirmed, InboundFednowTransferStatusDeclined, InboundFednowTransferStatusRequiresAttention:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_fednow_transfer`.
type InboundFednowTransferType string

const (
	InboundFednowTransferTypeInboundFednowTransfer InboundFednowTransferType = "inbound_fednow_transfer"
)

func (r InboundFednowTransferType) IsKnown() bool {
	switch r {
	case InboundFednowTransferTypeInboundFednowTransfer:
		return true
	}
	return false
}

// A list of Inbound FedNow Transfer objects.
type InboundFednowTransferListResponse struct {
	// The contents of the list.
	Data []InboundFednowTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                                `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{}                `json:"-,extras"`
	JSON        inboundFednowTransferListResponseJSON `json:"-"`
}

// inboundFednowTransferListResponseJSON contains the JSON metadata for the struct
// [InboundFednowTransferListResponse]
type inboundFednowTransferListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundFednowTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundFednowTransferListResponseJSON) RawJSON() string {
	return r.raw
}

type InboundFednowTransferListParams struct {
	// Filter Inbound FedNow Transfers to those belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound FedNow Transfers to ones belonging to the specified Account
	// Number.
	AccountNumberID param.Field[string]                                   `query:"account_number_id"`
	CreatedAt       param.Field[InboundFednowTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundFednowTransferListParams]'s query parameters as
// `url.Values`.
func (r InboundFednowTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundFednowTransferListParamsCreatedAt struct {
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

// URLQuery serializes [InboundFednowTransferListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r InboundFednowTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
