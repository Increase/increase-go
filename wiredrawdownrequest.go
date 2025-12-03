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
	opts = slices.Concat(r.Options, opts)
	path := "wire_drawdown_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Wire Drawdown Request
func (r *WireDrawdownRequestService) Get(ctx context.Context, wireDrawdownRequestID string, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if wireDrawdownRequestID == "" {
		err = errors.New("missing required wire_drawdown_request_id parameter")
		return
	}
	path := fmt.Sprintf("wire_drawdown_requests/%s", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query WireDrawdownRequestListParams, opts ...option.RequestOption) (res *WireDrawdownRequestListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wire_drawdown_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Wire drawdown requests enable you to request that someone else send you a wire.
// Because there is nuance to making sure your counterparty's bank processes these
// correctly, we ask that you reach out to
// [support@increase.com](mailto:support@increase.com) to enable this feature so we
// can help you plan your integration. For more information, see our
// [Wire Drawdown Requests documentation](/documentation/wire-drawdown-requests).
type WireDrawdownRequest struct {
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The Account Number to which the debtor—the recipient of this request—is being
	// requested to send funds.
	AccountNumberID string `json:"account_number_id,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the wire drawdown request was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The creditor's address.
	CreditorAddress WireDrawdownRequestCreditorAddress `json:"creditor_address,required"`
	// The creditor's name.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// The debtor's account number.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The debtor's address.
	DebtorAddress WireDrawdownRequestDebtorAddress `json:"debtor_address,required"`
	// The debtor's external account identifier.
	DebtorExternalAccountID string `json:"debtor_external_account_id,required,nullable"`
	// The debtor's name.
	DebtorName string `json:"debtor_name,required"`
	// The debtor's routing number.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentInboundWireTransferID string `json:"fulfillment_inbound_wire_transfer_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The lifecycle status of the drawdown request.
	Status WireDrawdownRequestStatus `json:"status,required"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission WireDrawdownRequestSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type WireDrawdownRequestType `json:"type,required"`
	// Remittance information the debtor will see as part of the drawdown request.
	UnstructuredRemittanceInformation string                  `json:"unstructured_remittance_information,required"`
	JSON                              wireDrawdownRequestJSON `json:"-"`
}

// wireDrawdownRequestJSON contains the JSON metadata for the struct
// [WireDrawdownRequest]
type wireDrawdownRequestJSON struct {
	ID                                apijson.Field
	AccountNumberID                   apijson.Field
	Amount                            apijson.Field
	CreatedAt                         apijson.Field
	CreditorAddress                   apijson.Field
	CreditorName                      apijson.Field
	Currency                          apijson.Field
	DebtorAccountNumber               apijson.Field
	DebtorAddress                     apijson.Field
	DebtorExternalAccountID           apijson.Field
	DebtorName                        apijson.Field
	DebtorRoutingNumber               apijson.Field
	FulfillmentInboundWireTransferID  apijson.Field
	IdempotencyKey                    apijson.Field
	Status                            apijson.Field
	Submission                        apijson.Field
	Type                              apijson.Field
	UnstructuredRemittanceInformation apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *WireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireDrawdownRequestJSON) RawJSON() string {
	return r.raw
}

// The creditor's address.
type WireDrawdownRequestCreditorAddress struct {
	// The city, district, town, or village of the address.
	City string `json:"city,required"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country string `json:"country,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The ZIP code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// The address state.
	State string                                 `json:"state,required,nullable"`
	JSON  wireDrawdownRequestCreditorAddressJSON `json:"-"`
}

// wireDrawdownRequestCreditorAddressJSON contains the JSON metadata for the struct
// [WireDrawdownRequestCreditorAddress]
type wireDrawdownRequestCreditorAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireDrawdownRequestCreditorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireDrawdownRequestCreditorAddressJSON) RawJSON() string {
	return r.raw
}

// The debtor's address.
type WireDrawdownRequestDebtorAddress struct {
	// The city, district, town, or village of the address.
	City string `json:"city,required"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country string `json:"country,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The ZIP code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// The address state.
	State string                               `json:"state,required,nullable"`
	JSON  wireDrawdownRequestDebtorAddressJSON `json:"-"`
}

// wireDrawdownRequestDebtorAddressJSON contains the JSON metadata for the struct
// [WireDrawdownRequestDebtorAddress]
type wireDrawdownRequestDebtorAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireDrawdownRequestDebtorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireDrawdownRequestDebtorAddressJSON) RawJSON() string {
	return r.raw
}

// The lifecycle status of the drawdown request.
type WireDrawdownRequestStatus string

const (
	WireDrawdownRequestStatusPendingSubmission WireDrawdownRequestStatus = "pending_submission"
	WireDrawdownRequestStatusFulfilled         WireDrawdownRequestStatus = "fulfilled"
	WireDrawdownRequestStatusPendingResponse   WireDrawdownRequestStatus = "pending_response"
	WireDrawdownRequestStatusRefused           WireDrawdownRequestStatus = "refused"
)

func (r WireDrawdownRequestStatus) IsKnown() bool {
	switch r {
	case WireDrawdownRequestStatusPendingSubmission, WireDrawdownRequestStatusFulfilled, WireDrawdownRequestStatusPendingResponse, WireDrawdownRequestStatusRefused:
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

// A list of Wire Drawdown Request objects.
type WireDrawdownRequestListResponse struct {
	// The contents of the list.
	Data []WireDrawdownRequest `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                              `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{}              `json:"-,extras"`
	JSON        wireDrawdownRequestListResponseJSON `json:"-"`
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

func (r wireDrawdownRequestListResponseJSON) RawJSON() string {
	return r.raw
}

type WireDrawdownRequestNewParams struct {
	// The Account Number to which the debtor should send funds.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The amount requested from the debtor, in USD cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The creditor's address.
	CreditorAddress param.Field[WireDrawdownRequestNewParamsCreditorAddress] `json:"creditor_address,required"`
	// The creditor's name.
	CreditorName param.Field[string] `json:"creditor_name,required"`
	// The debtor's address.
	DebtorAddress param.Field[WireDrawdownRequestNewParamsDebtorAddress] `json:"debtor_address,required"`
	// The debtor's name.
	DebtorName param.Field[string] `json:"debtor_name,required"`
	// Remittance information the debtor will see as part of the request.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information,required"`
	// The debtor's account number.
	DebtorAccountNumber param.Field[string] `json:"debtor_account_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `debtor_account_number` and `debtor_routing_number` must be absent.
	DebtorExternalAccountID param.Field[string] `json:"debtor_external_account_id"`
	// The debtor's routing number.
	DebtorRoutingNumber param.Field[string] `json:"debtor_routing_number"`
}

func (r WireDrawdownRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The creditor's address.
type WireDrawdownRequestNewParamsCreditorAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The ZIP code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// The address state.
	State param.Field[string] `json:"state"`
}

func (r WireDrawdownRequestNewParamsCreditorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The debtor's address.
type WireDrawdownRequestNewParamsDebtorAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The ZIP code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// The address state.
	State param.Field[string] `json:"state"`
}

func (r WireDrawdownRequestNewParamsDebtorAddress) MarshalJSON() (data []byte, err error) {
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
	Limit  param.Field[int64]                               `query:"limit"`
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

type WireDrawdownRequestListParamsStatus struct {
	// Filter Wire Drawdown Requests for those with the specified status. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]WireDrawdownRequestListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [WireDrawdownRequestListParamsStatus]'s query parameters as
// `url.Values`.
func (r WireDrawdownRequestListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WireDrawdownRequestListParamsStatusIn string

const (
	WireDrawdownRequestListParamsStatusInPendingSubmission WireDrawdownRequestListParamsStatusIn = "pending_submission"
	WireDrawdownRequestListParamsStatusInFulfilled         WireDrawdownRequestListParamsStatusIn = "fulfilled"
	WireDrawdownRequestListParamsStatusInPendingResponse   WireDrawdownRequestListParamsStatusIn = "pending_response"
	WireDrawdownRequestListParamsStatusInRefused           WireDrawdownRequestListParamsStatusIn = "refused"
)

func (r WireDrawdownRequestListParamsStatusIn) IsKnown() bool {
	switch r {
	case WireDrawdownRequestListParamsStatusInPendingSubmission, WireDrawdownRequestListParamsStatusInFulfilled, WireDrawdownRequestListParamsStatusInPendingResponse, WireDrawdownRequestListParamsStatusInRefused:
		return true
	}
	return false
}
