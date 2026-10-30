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

// RealTimePaymentsTransferService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRealTimePaymentsTransferService] method instead.
type RealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewRealTimePaymentsTransferService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRealTimePaymentsTransferService(opts ...option.RequestOption) (r *RealTimePaymentsTransferService) {
	r = &RealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Create a Real-Time Payments Transfer
func (r *RealTimePaymentsTransferService) New(ctx context.Context, body RealTimePaymentsTransferNewParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Real-Time Payments Transfer
func (r *RealTimePaymentsTransferService) Get(ctx context.Context, realTimePaymentsTransferID string, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimePaymentsTransferID == "" {
		err = errors.New("missing required real_time_payments_transfer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("real_time_payments_transfers/%s", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Real-Time Payments Transfers
func (r *RealTimePaymentsTransferService) List(ctx context.Context, query RealTimePaymentsTransferListParams, opts ...option.RequestOption) (res *pagination.Page[RealTimePaymentsTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "real_time_payments_transfers"
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

// List Real-Time Payments Transfers
func (r *RealTimePaymentsTransferService) ListAutoPaging(ctx context.Context, query RealTimePaymentsTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[RealTimePaymentsTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves a Real-Time Payments Transfer in a pending_approval state.
func (r *RealTimePaymentsTransferService) Approve(ctx context.Context, realTimePaymentsTransferID string, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimePaymentsTransferID == "" {
		err = errors.New("missing required real_time_payments_transfer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("real_time_payments_transfers/%s/approve", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Cancels a Real-Time Payments Transfer in a pending_approval state.
func (r *RealTimePaymentsTransferService) Cancel(ctx context.Context, realTimePaymentsTransferID string, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimePaymentsTransferID == "" {
		err = errors.New("missing required real_time_payments_transfer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("real_time_payments_transfers/%s/cancel", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type RealTimePaymentsTransferNewParams struct {
	// The transfer amount in USD cents. For Real-Time Payments transfers, must be
	// positive.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The name of the transfer's recipient.
	CreditorName param.Field[string] `json:"creditor_name" api:"required"`
	// The identifier of the Account Number from which to send the transfer.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id" api:"required"`
	// Unstructured information that will show on the recipient's bank statement.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information" api:"required"`
	// The destination account number.
	AccountNumber param.Field[string] `json:"account_number"`
	// The name of the transfer's sender. If not provided, defaults to the name of the
	// account's entity.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// The ID of an Inbound Real-Time Payments Request for Payment in response to which
	// this transfer is being sent.
	InboundRealTimePaymentsRequestForPaymentID param.Field[string] `json:"inbound_real_time_payments_request_for_payment_id"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	RoutingNumber param.Field[string] `json:"routing_number"`
	// The name of the ultimate recipient of the transfer. Set this if the creditor is
	// an intermediary receiving the payment for someone else.
	UltimateCreditorName param.Field[string] `json:"ultimate_creditor_name"`
	// The name of the ultimate sender of the transfer. Set this if the funds are being
	// sent on behalf of someone who is not the account holder at Increase.
	UltimateDebtorName param.Field[string] `json:"ultimate_debtor_name"`
}

func (r RealTimePaymentsTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RealTimePaymentsTransferListParams struct {
	// Filter Real-Time Payments Transfers to those belonging to the specified Account.
	AccountID param.Field[string]                                      `query:"account_id"`
	CreatedAt param.Field[RealTimePaymentsTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Real-Time Payments Transfers to those made to the specified External
	// Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit  param.Field[int64]                                    `query:"limit"`
	Status param.Field[RealTimePaymentsTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [RealTimePaymentsTransferListParams]'s query parameters as
// `url.Values`.
func (r RealTimePaymentsTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsTransferListParamsCreatedAt struct {
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

// URLQuery serializes [RealTimePaymentsTransferListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r RealTimePaymentsTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsTransferListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]RealTimePaymentsTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [RealTimePaymentsTransferListParamsStatus]'s query
// parameters as `url.Values`.
func (r RealTimePaymentsTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsTransferListParamsStatusIn string

const (
	RealTimePaymentsTransferListParamsStatusInPendingApproval   RealTimePaymentsTransferListParamsStatusIn = "pending_approval"
	RealTimePaymentsTransferListParamsStatusInCanceled          RealTimePaymentsTransferListParamsStatusIn = "canceled"
	RealTimePaymentsTransferListParamsStatusInPendingSubmission RealTimePaymentsTransferListParamsStatusIn = "pending_submission"
	RealTimePaymentsTransferListParamsStatusInPendingReviewing  RealTimePaymentsTransferListParamsStatusIn = "pending_reviewing"
	RealTimePaymentsTransferListParamsStatusInRequiresAttention RealTimePaymentsTransferListParamsStatusIn = "requires_attention"
	RealTimePaymentsTransferListParamsStatusInRejected          RealTimePaymentsTransferListParamsStatusIn = "rejected"
	RealTimePaymentsTransferListParamsStatusInSubmitted         RealTimePaymentsTransferListParamsStatusIn = "submitted"
	RealTimePaymentsTransferListParamsStatusInComplete          RealTimePaymentsTransferListParamsStatusIn = "complete"
)

func (r RealTimePaymentsTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferListParamsStatusInPendingApproval, RealTimePaymentsTransferListParamsStatusInCanceled, RealTimePaymentsTransferListParamsStatusInPendingSubmission, RealTimePaymentsTransferListParamsStatusInPendingReviewing, RealTimePaymentsTransferListParamsStatusInRequiresAttention, RealTimePaymentsTransferListParamsStatusInRejected, RealTimePaymentsTransferListParamsStatusInSubmitted, RealTimePaymentsTransferListParamsStatusInComplete:
		return true
	}
	return false
}
