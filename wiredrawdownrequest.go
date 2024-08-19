// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// WireDrawdownRequestService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWireDrawdownRequestService] method instead.
type WireDrawdownRequestService struct {
	Options []option.RequestOption
}

// NewWireDrawdownRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWireDrawdownRequestService(opts ...option.RequestOption) (r *WireDrawdownRequestService) {
	r = &WireDrawdownRequestService{}
	r.Options = opts
	return
}

// Create a Wire Drawdown Request
func (r *WireDrawdownRequestService) New(ctx context.Context, body WireDrawdownRequestNewParams, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "wire_drawdown_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Wire Drawdown Request
func (r *WireDrawdownRequestService) Get(ctx context.Context, wireDrawdownRequestID string, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	if wireDrawdownRequestID == "" {
		err = errors.New("missing required wire_drawdown_request_id parameter")
		return
	}
	path := fmt.Sprintf("wire_drawdown_requests/%s", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query WireDrawdownRequestListParams, opts ...option.RequestOption) (res *pagination.Page[WireDrawdownRequest], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "wire_drawdown_requests"
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

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) ListAutoPaging(ctx context.Context, query WireDrawdownRequestListParams, opts ...option.RequestOption) *pagination.PageAutoPager[WireDrawdownRequest] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Wire drawdown requests enable you to request that someone else send you a wire.
// This feature is in beta; reach out to
// [support@increase.com](mailto:support@increase.com) to learn more.
type WireDrawdownRequest struct {
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberID string `json:"account_number_id,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionID string `json:"fulfillment_transaction_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The originator's address line 1.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The originator's address line 2.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The originator's address line 3.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator's name.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber string `json:"recipient_account_number,required"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 string `json:"recipient_address_line1,required,nullable"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 string `json:"recipient_address_line2,required,nullable"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 string `json:"recipient_address_line3,required,nullable"`
	// The drawdown request's recipient's name.
	RecipientName string `json:"recipient_name,required,nullable"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber string `json:"recipient_routing_number,required"`
	// The lifecycle status of the drawdown request.
	Status WireDrawdownRequestStatus `json:"status,required"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission WireDrawdownRequestSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type WireDrawdownRequestType `json:"type,required"`
	JSON wireDrawdownRequestJSON `json:"-"`
}

// wireDrawdownRequestJSON contains the JSON metadata for the struct
// [WireDrawdownRequest]
type wireDrawdownRequestJSON struct {
	ID                       apijson.Field
	AccountNumberID          apijson.Field
	Amount                   apijson.Field
	Currency                 apijson.Field
	FulfillmentTransactionID apijson.Field
	IdempotencyKey           apijson.Field
	MessageToRecipient       apijson.Field
	OriginatorAddressLine1   apijson.Field
	OriginatorAddressLine2   apijson.Field
	OriginatorAddressLine3   apijson.Field
	OriginatorName           apijson.Field
	RecipientAccountNumber   apijson.Field
	RecipientAddressLine1    apijson.Field
	RecipientAddressLine2    apijson.Field
	RecipientAddressLine3    apijson.Field
	RecipientName            apijson.Field
	RecipientRoutingNumber   apijson.Field
	Status                   apijson.Field
	Submission               apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *WireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireDrawdownRequestJSON) RawJSON() string {
	return r.raw
}

// The lifecycle status of the drawdown request.
type WireDrawdownRequestStatus string

const (
	// The drawdown request is queued to be submitted to Fedwire.
	WireDrawdownRequestStatusPendingSubmission WireDrawdownRequestStatus = "pending_submission"
	// The drawdown request has been sent and the recipient should respond in some way.
	WireDrawdownRequestStatusPendingResponse WireDrawdownRequestStatus = "pending_response"
	// The drawdown request has been fulfilled by the recipient.
	WireDrawdownRequestStatusFulfilled WireDrawdownRequestStatus = "fulfilled"
	// The drawdown request has been refused by the recipient.
	WireDrawdownRequestStatusRefused WireDrawdownRequestStatus = "refused"
)

func (r WireDrawdownRequestStatus) IsKnown() bool {
	switch r {
	case WireDrawdownRequestStatusPendingSubmission, WireDrawdownRequestStatusPendingResponse, WireDrawdownRequestStatusFulfilled, WireDrawdownRequestStatusRefused:
		return true
	}
	return false
}

// After the drawdown request is submitted to Fedwire, this will contain
// supplemental details.
type WireDrawdownRequestSubmission struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData string                            `json:"input_message_accountability_data,required"`
	JSON                           wireDrawdownRequestSubmissionJSON `json:"-"`
}

// wireDrawdownRequestSubmissionJSON contains the JSON metadata for the struct
// [WireDrawdownRequestSubmission]
type wireDrawdownRequestSubmissionJSON struct {
	InputMessageAccountabilityData apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *WireDrawdownRequestSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireDrawdownRequestSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `wire_drawdown_request`.
type WireDrawdownRequestType string

const (
	WireDrawdownRequestTypeWireDrawdownRequest WireDrawdownRequestType = "wire_drawdown_request"
)

func (r WireDrawdownRequestType) IsKnown() bool {
	switch r {
	case WireDrawdownRequestTypeWireDrawdownRequest:
		return true
	}
	return false
}

type WireDrawdownRequestNewParams struct {
	// The Account Number to which the recipient should send funds.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The amount requested from the recipient, in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// A message the recipient will see as part of the request.
	MessageToRecipient param.Field[string] `json:"message_to_recipient,required"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber param.Field[string] `json:"recipient_account_number,required"`
	// The drawdown request's recipient's name.
	RecipientName param.Field[string] `json:"recipient_name,required"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber param.Field[string] `json:"recipient_routing_number,required"`
	// The drawdown request originator's address line 1. This is only necessary if
	// you're requesting a payment to a commingled account. Otherwise, we'll use the
	// associated entity's details.
	OriginatorAddressLine1 param.Field[string] `json:"originator_address_line1"`
	// The drawdown request originator's address line 2. This is only necessary if
	// you're requesting a payment to a commingled account. Otherwise, we'll use the
	// associated entity's details.
	OriginatorAddressLine2 param.Field[string] `json:"originator_address_line2"`
	// The drawdown request originator's address line 3. This is only necessary if
	// you're requesting a payment to a commingled account. Otherwise, we'll use the
	// associated entity's details.
	OriginatorAddressLine3 param.Field[string] `json:"originator_address_line3"`
	// The drawdown request originator's name. This is only necessary if you're
	// requesting a payment to a commingled account. Otherwise, we'll use the
	// associated entity's details.
	OriginatorName param.Field[string] `json:"originator_name"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 param.Field[string] `json:"recipient_address_line1"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 param.Field[string] `json:"recipient_address_line2"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 param.Field[string] `json:"recipient_address_line3"`
}

func (r WireDrawdownRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WireDrawdownRequestListParams struct {
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
	// Filter Wire Drawdown Requests for those with the specified status.
	Status param.Field[WireDrawdownRequestListParamsStatus] `query:"status"`
}

// URLQuery serializes [WireDrawdownRequestListParams]'s query parameters as
// `url.Values`.
func (r WireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter Wire Drawdown Requests for those with the specified status.
type WireDrawdownRequestListParamsStatus string

const (
	// The drawdown request is queued to be submitted to Fedwire.
	WireDrawdownRequestListParamsStatusPendingSubmission WireDrawdownRequestListParamsStatus = "pending_submission"
	// The drawdown request has been sent and the recipient should respond in some way.
	WireDrawdownRequestListParamsStatusPendingResponse WireDrawdownRequestListParamsStatus = "pending_response"
	// The drawdown request has been fulfilled by the recipient.
	WireDrawdownRequestListParamsStatusFulfilled WireDrawdownRequestListParamsStatus = "fulfilled"
	// The drawdown request has been refused by the recipient.
	WireDrawdownRequestListParamsStatusRefused WireDrawdownRequestListParamsStatus = "refused"
)

func (r WireDrawdownRequestListParamsStatus) IsKnown() bool {
	switch r {
	case WireDrawdownRequestListParamsStatusPendingSubmission, WireDrawdownRequestListParamsStatusPendingResponse, WireDrawdownRequestListParamsStatusFulfilled, WireDrawdownRequestListParamsStatusRefused:
		return true
	}
	return false
}
