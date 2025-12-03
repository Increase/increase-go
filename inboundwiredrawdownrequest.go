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

// InboundWireDrawdownRequestService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundWireDrawdownRequestService] method instead.
type InboundWireDrawdownRequestService struct {
	Options []option.RequestOption
}

// NewInboundWireDrawdownRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInboundWireDrawdownRequestService(opts ...option.RequestOption) (r *InboundWireDrawdownRequestService) {
	r = &InboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Wire Drawdown Request
func (r *InboundWireDrawdownRequestService) Get(ctx context.Context, inboundWireDrawdownRequestID string, opts ...option.RequestOption) (res *InboundWireDrawdownRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundWireDrawdownRequestID == "" {
		err = errors.New("missing required inbound_wire_drawdown_request_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_wire_drawdown_requests/%s", inboundWireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) List(ctx context.Context, query InboundWireDrawdownRequestListParams, opts ...option.RequestOption) (res *pagination.Page[InboundWireDrawdownRequest], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_wire_drawdown_requests"
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

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) ListAutoPaging(ctx context.Context, query InboundWireDrawdownRequestListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundWireDrawdownRequest] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Inbound wire drawdown requests are requests from someone else to send them a
// wire. For more information, see our
// [Wire Drawdown Requests documentation](/documentation/wire-drawdown-requests).
type InboundWireDrawdownRequest struct {
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the inbound wire drawdown requested was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The creditor's account number.
	CreditorAccountNumber string `json:"creditor_account_number,required"`
	// A free-form address field set by the sender.
	CreditorAddressLine1 string `json:"creditor_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	CreditorAddressLine2 string `json:"creditor_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	CreditorAddressLine3 string `json:"creditor_address_line3,required,nullable"`
	// A name set by the sender.
	CreditorName string `json:"creditor_name,required,nullable"`
	// The creditor's routing number.
	CreditorRoutingNumber string `json:"creditor_routing_number,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// A free-form address field set by the sender.
	DebtorAddressLine1 string `json:"debtor_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	DebtorAddressLine2 string `json:"debtor_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	DebtorAddressLine3 string `json:"debtor_address_line3,required,nullable"`
	// A name set by the sender.
	DebtorName string `json:"debtor_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the drawdown
	// request.
	EndToEndIdentification string `json:"end_to_end_identification,required,nullable"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the originating bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The sending bank's identifier for the drawdown request.
	InstructionIdentification string `json:"instruction_identification,required,nullable"`
	// The Account Number from which the recipient of this request is being requested
	// to send funds.
	RecipientAccountNumberID string `json:"recipient_account_number_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_drawdown_request`.
	Type InboundWireDrawdownRequestType `json:"type,required"`
	// The Unique End-to-end Transaction Reference
	// ([UETR](https://www.swift.com/payments/what-unique-end-end-transaction-reference-uetr))
	// of the drawdown request.
	UniqueEndToEndTransactionReference string `json:"unique_end_to_end_transaction_reference,required,nullable"`
	// A free-form message set by the sender.
	UnstructuredRemittanceInformation string                         `json:"unstructured_remittance_information,required,nullable"`
	ExtraFields                       map[string]interface{}         `json:"-,extras"`
	JSON                              inboundWireDrawdownRequestJSON `json:"-"`
}

// inboundWireDrawdownRequestJSON contains the JSON metadata for the struct
// [InboundWireDrawdownRequest]
type inboundWireDrawdownRequestJSON struct {
	ID                                 apijson.Field
	Amount                             apijson.Field
	CreatedAt                          apijson.Field
	CreditorAccountNumber              apijson.Field
	CreditorAddressLine1               apijson.Field
	CreditorAddressLine2               apijson.Field
	CreditorAddressLine3               apijson.Field
	CreditorName                       apijson.Field
	CreditorRoutingNumber              apijson.Field
	Currency                           apijson.Field
	DebtorAddressLine1                 apijson.Field
	DebtorAddressLine2                 apijson.Field
	DebtorAddressLine3                 apijson.Field
	DebtorName                         apijson.Field
	EndToEndIdentification             apijson.Field
	InputMessageAccountabilityData     apijson.Field
	InstructionIdentification          apijson.Field
	RecipientAccountNumberID           apijson.Field
	Type                               apijson.Field
	UniqueEndToEndTransactionReference apijson.Field
	UnstructuredRemittanceInformation  apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundWireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundWireDrawdownRequestJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_drawdown_request`.
type InboundWireDrawdownRequestType string

const (
	InboundWireDrawdownRequestTypeInboundWireDrawdownRequest InboundWireDrawdownRequestType = "inbound_wire_drawdown_request"
)

func (r InboundWireDrawdownRequestType) IsKnown() bool {
	switch r {
	case InboundWireDrawdownRequestTypeInboundWireDrawdownRequest:
		return true
	}
	return false
}

type InboundWireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundWireDrawdownRequestListParams]'s query parameters as
// `url.Values`.
func (r InboundWireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
