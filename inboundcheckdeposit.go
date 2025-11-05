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
	"github.com/Increase/increase-go/packages/pagination"
)

// InboundCheckDepositService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundCheckDepositService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	if inboundCheckDepositID == "" {
		err = errors.New("missing required inbound_check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_check_deposits/%s", inboundCheckDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Check Deposits
func (r *InboundCheckDepositService) List(ctx context.Context, query InboundCheckDepositListParams, opts ...option.RequestOption) (res *pagination.Page[InboundCheckDeposit], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Decline an Inbound Check Deposit
func (r *InboundCheckDepositService) Decline(ctx context.Context, inboundCheckDepositID string, opts ...option.RequestOption) (res *InboundCheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundCheckDepositID == "" {
		err = errors.New("missing required inbound_check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_check_deposits/%s/decline", inboundCheckDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Return an Inbound Check Deposit
func (r *InboundCheckDepositService) Return(ctx context.Context, inboundCheckDepositID string, body InboundCheckDepositReturnParams, opts ...option.RequestOption) (res *InboundCheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundCheckDepositID == "" {
		err = errors.New("missing required inbound_check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_check_deposits/%s/return", inboundCheckDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Inbound Check Deposits are records of third-parties attempting to deposit checks
// against your account.
type InboundCheckDeposit struct {
	// The deposit's identifier.
	ID string `json:"id,required"`
	// If the Inbound Check Deposit was accepted, the
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which this
	// took place.
	AcceptedAt time.Time `json:"accepted_at,required,nullable" format:"date-time"`
	// The Account the check is being deposited against.
	AccountID string `json:"account_id,required"`
	// The Account Number the check is being deposited against.
	AccountNumberID string `json:"account_number_id,required,nullable"`
	// If the deposit or the return was adjusted by the sending institution, this will
	// contain details of the adjustments.
	Adjustments []InboundCheckDepositAdjustment `json:"adjustments,required"`
	// The deposited amount in USD cents.
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
	// If the Inbound Check Deposit was declined, the
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which this
	// took place.
	DeclinedAt time.Time `json:"declined_at,required,nullable" format:"date-time"`
	// If the deposit attempt has been rejected, the identifier of the Declined
	// Transaction object created as a result of the failed deposit.
	DeclinedTransactionID string `json:"declined_transaction_id,required,nullable"`
	// If you requested a return of this deposit, this will contain details of the
	// return.
	DepositReturn InboundCheckDepositDepositReturn `json:"deposit_return,required,nullable"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// Whether the details on the check match the recipient name of the check transfer.
	// This is an optional feature, contact sales to enable.
	PayeeNameAnalysis InboundCheckDepositPayeeNameAnalysis `json:"payee_name_analysis,required"`
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
	AcceptedAt                      apijson.Field
	AccountID                       apijson.Field
	AccountNumberID                 apijson.Field
	Adjustments                     apijson.Field
	Amount                          apijson.Field
	BackImageFileID                 apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	CheckNumber                     apijson.Field
	CheckTransferID                 apijson.Field
	CreatedAt                       apijson.Field
	Currency                        apijson.Field
	DeclinedAt                      apijson.Field
	DeclinedTransactionID           apijson.Field
	DepositReturn                   apijson.Field
	FrontImageFileID                apijson.Field
	PayeeNameAnalysis               apijson.Field
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

type InboundCheckDepositAdjustment struct {
	// The time at which the return adjustment was received.
	AdjustedAt time.Time `json:"adjusted_at,required" format:"date-time"`
	// The amount of the adjustment.
	Amount int64 `json:"amount,required"`
	// The reason for the adjustment.
	Reason InboundCheckDepositAdjustmentsReason `json:"reason,required"`
	// The id of the transaction for the adjustment.
	TransactionID string                            `json:"transaction_id,required"`
	JSON          inboundCheckDepositAdjustmentJSON `json:"-"`
}

// inboundCheckDepositAdjustmentJSON contains the JSON metadata for the struct
// [InboundCheckDepositAdjustment]
type inboundCheckDepositAdjustmentJSON struct {
	AdjustedAt    apijson.Field
	Amount        apijson.Field
	Reason        apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundCheckDepositAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundCheckDepositAdjustmentJSON) RawJSON() string {
	return r.raw
}

// The reason for the adjustment.
type InboundCheckDepositAdjustmentsReason string

const (
	InboundCheckDepositAdjustmentsReasonLateReturn        InboundCheckDepositAdjustmentsReason = "late_return"
	InboundCheckDepositAdjustmentsReasonWrongPayeeCredit  InboundCheckDepositAdjustmentsReason = "wrong_payee_credit"
	InboundCheckDepositAdjustmentsReasonAdjustedAmount    InboundCheckDepositAdjustmentsReason = "adjusted_amount"
	InboundCheckDepositAdjustmentsReasonNonConformingItem InboundCheckDepositAdjustmentsReason = "non_conforming_item"
)

func (r InboundCheckDepositAdjustmentsReason) IsKnown() bool {
	switch r {
	case InboundCheckDepositAdjustmentsReasonLateReturn, InboundCheckDepositAdjustmentsReasonWrongPayeeCredit, InboundCheckDepositAdjustmentsReasonAdjustedAmount, InboundCheckDepositAdjustmentsReasonNonConformingItem:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
type InboundCheckDepositCurrency string

const (
	InboundCheckDepositCurrencyUsd InboundCheckDepositCurrency = "USD"
)

func (r InboundCheckDepositCurrency) IsKnown() bool {
	switch r {
	case InboundCheckDepositCurrencyUsd:
		return true
	}
	return false
}

// If you requested a return of this deposit, this will contain details of the
// return.
type InboundCheckDepositDepositReturn struct {
	// The reason the deposit was returned.
	Reason InboundCheckDepositDepositReturnReason `json:"reason,required"`
	// The time at which the deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The id of the transaction for the returned deposit.
	TransactionID string                               `json:"transaction_id,required"`
	JSON          inboundCheckDepositDepositReturnJSON `json:"-"`
}

// inboundCheckDepositDepositReturnJSON contains the JSON metadata for the struct
// [InboundCheckDepositDepositReturn]
type inboundCheckDepositDepositReturnJSON struct {
	Reason        apijson.Field
	ReturnedAt    apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundCheckDepositDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundCheckDepositDepositReturnJSON) RawJSON() string {
	return r.raw
}

// The reason the deposit was returned.
type InboundCheckDepositDepositReturnReason string

const (
	InboundCheckDepositDepositReturnReasonAlteredOrFictitious  InboundCheckDepositDepositReturnReason = "altered_or_fictitious"
	InboundCheckDepositDepositReturnReasonNotAuthorized        InboundCheckDepositDepositReturnReason = "not_authorized"
	InboundCheckDepositDepositReturnReasonDuplicatePresentment InboundCheckDepositDepositReturnReason = "duplicate_presentment"
	InboundCheckDepositDepositReturnReasonEndorsementMissing   InboundCheckDepositDepositReturnReason = "endorsement_missing"
	InboundCheckDepositDepositReturnReasonEndorsementIrregular InboundCheckDepositDepositReturnReason = "endorsement_irregular"
)

func (r InboundCheckDepositDepositReturnReason) IsKnown() bool {
	switch r {
	case InboundCheckDepositDepositReturnReasonAlteredOrFictitious, InboundCheckDepositDepositReturnReasonNotAuthorized, InboundCheckDepositDepositReturnReasonDuplicatePresentment, InboundCheckDepositDepositReturnReasonEndorsementMissing, InboundCheckDepositDepositReturnReasonEndorsementIrregular:
		return true
	}
	return false
}

// Whether the details on the check match the recipient name of the check transfer.
// This is an optional feature, contact sales to enable.
type InboundCheckDepositPayeeNameAnalysis string

const (
	InboundCheckDepositPayeeNameAnalysisNameMatches  InboundCheckDepositPayeeNameAnalysis = "name_matches"
	InboundCheckDepositPayeeNameAnalysisDoesNotMatch InboundCheckDepositPayeeNameAnalysis = "does_not_match"
	InboundCheckDepositPayeeNameAnalysisNotEvaluated InboundCheckDepositPayeeNameAnalysis = "not_evaluated"
)

func (r InboundCheckDepositPayeeNameAnalysis) IsKnown() bool {
	switch r {
	case InboundCheckDepositPayeeNameAnalysisNameMatches, InboundCheckDepositPayeeNameAnalysisDoesNotMatch, InboundCheckDepositPayeeNameAnalysisNotEvaluated:
		return true
	}
	return false
}

// The status of the Inbound Check Deposit.
type InboundCheckDepositStatus string

const (
	InboundCheckDepositStatusPending           InboundCheckDepositStatus = "pending"
	InboundCheckDepositStatusAccepted          InboundCheckDepositStatus = "accepted"
	InboundCheckDepositStatusDeclined          InboundCheckDepositStatus = "declined"
	InboundCheckDepositStatusReturned          InboundCheckDepositStatus = "returned"
	InboundCheckDepositStatusRequiresAttention InboundCheckDepositStatus = "requires_attention"
)

func (r InboundCheckDepositStatus) IsKnown() bool {
	switch r {
	case InboundCheckDepositStatusPending, InboundCheckDepositStatusAccepted, InboundCheckDepositStatusDeclined, InboundCheckDepositStatusReturned, InboundCheckDepositStatusRequiresAttention:
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
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound Check Deposits to those belonging to the specified Check
	// Transfer.
	CheckTransferID param.Field[string]                                 `query:"check_transfer_id"`
	CreatedAt       param.Field[InboundCheckDepositListParamsCreatedAt] `query:"created_at"`
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

type InboundCheckDepositReturnParams struct {
	// The reason to return the Inbound Check Deposit.
	Reason param.Field[InboundCheckDepositReturnParamsReason] `json:"reason,required"`
}

func (r InboundCheckDepositReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason to return the Inbound Check Deposit.
type InboundCheckDepositReturnParamsReason string

const (
	InboundCheckDepositReturnParamsReasonAlteredOrFictitious  InboundCheckDepositReturnParamsReason = "altered_or_fictitious"
	InboundCheckDepositReturnParamsReasonNotAuthorized        InboundCheckDepositReturnParamsReason = "not_authorized"
	InboundCheckDepositReturnParamsReasonDuplicatePresentment InboundCheckDepositReturnParamsReason = "duplicate_presentment"
	InboundCheckDepositReturnParamsReasonEndorsementMissing   InboundCheckDepositReturnParamsReason = "endorsement_missing"
	InboundCheckDepositReturnParamsReasonEndorsementIrregular InboundCheckDepositReturnParamsReason = "endorsement_irregular"
)

func (r InboundCheckDepositReturnParamsReason) IsKnown() bool {
	switch r {
	case InboundCheckDepositReturnParamsReasonAlteredOrFictitious, InboundCheckDepositReturnParamsReasonNotAuthorized, InboundCheckDepositReturnParamsReasonDuplicatePresentment, InboundCheckDepositReturnParamsReasonEndorsementMissing, InboundCheckDepositReturnParamsReasonEndorsementIrregular:
		return true
	}
	return false
}
