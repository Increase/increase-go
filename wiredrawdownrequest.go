// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// WireDrawdownRequestService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWireDrawdownRequestService]
// method instead.
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
	path := fmt.Sprintf("wire_drawdown_requests/%s", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query WireDrawdownRequestListParams, opts ...option.RequestOption) (res *shared.Page[WireDrawdownRequest], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *WireDrawdownRequestService) ListAutoPaging(ctx context.Context, query WireDrawdownRequestListParams, opts ...option.RequestOption) *shared.PageAutoPager[WireDrawdownRequest] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Wire drawdown requests enable you to request that someone else send you a wire.
// This feature is in beta; reach out to
// [support@increase.com](mailto:support@increase.com) to learn more.
type WireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type WireDrawdownRequestType `json:"type,required"`
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberID string `json:"account_number_id,required"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber string `json:"recipient_account_number,required"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber string `json:"recipient_routing_number,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The drawdown request's recipient's name.
	RecipientName string `json:"recipient_name,required,nullable"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 string `json:"recipient_address_line1,required,nullable"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 string `json:"recipient_address_line2,required,nullable"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 string `json:"recipient_address_line3,required,nullable"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission WireDrawdownRequestSubmission `json:"submission,required,nullable"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionID string `json:"fulfillment_transaction_id,required,nullable"`
	// The lifecycle status of the drawdown request.
	Status WireDrawdownRequestStatus `json:"status,required"`
	JSON   wireDrawdownRequestJSON
}

// wireDrawdownRequestJSON contains the JSON metadata for the struct
// [WireDrawdownRequest]
type wireDrawdownRequestJSON struct {
	Type                     apijson.Field
	ID                       apijson.Field
	AccountNumberID          apijson.Field
	RecipientAccountNumber   apijson.Field
	RecipientRoutingNumber   apijson.Field
	Amount                   apijson.Field
	Currency                 apijson.Field
	MessageToRecipient       apijson.Field
	RecipientName            apijson.Field
	RecipientAddressLine1    apijson.Field
	RecipientAddressLine2    apijson.Field
	RecipientAddressLine3    apijson.Field
	Submission               apijson.Field
	FulfillmentTransactionID apijson.Field
	Status                   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *WireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireDrawdownRequestType string

const (
	WireDrawdownRequestTypeWireDrawdownRequest WireDrawdownRequestType = "wire_drawdown_request"
)

// After the drawdown request is submitted to Fedwire, this will contain
// supplemental details.
type WireDrawdownRequestSubmission struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	JSON                           wireDrawdownRequestSubmissionJSON
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

type WireDrawdownRequestStatus string

const (
	WireDrawdownRequestStatusPendingSubmission WireDrawdownRequestStatus = "pending_submission"
	WireDrawdownRequestStatusPendingResponse   WireDrawdownRequestStatus = "pending_response"
	WireDrawdownRequestStatusFulfilled         WireDrawdownRequestStatus = "fulfilled"
	WireDrawdownRequestStatusRefused           WireDrawdownRequestStatus = "refused"
)

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
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WireDrawdownRequestListParams]'s query parameters as
// `url.Values`.
func (r WireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of Wire Drawdown Request objects
type WireDrawdownRequestListResponse struct {
	// The contents of the list.
	Data []WireDrawdownRequest `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       wireDrawdownRequestListResponseJSON
}

// wireDrawdownRequestListResponseJSON contains the JSON metadata for the struct
// [WireDrawdownRequestListResponse]
type wireDrawdownRequestListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireDrawdownRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
