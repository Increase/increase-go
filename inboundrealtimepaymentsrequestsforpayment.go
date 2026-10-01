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

// InboundRealTimePaymentsRequestsForPaymentService contains methods and other
// services that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundRealTimePaymentsRequestsForPaymentService] method instead.
type InboundRealTimePaymentsRequestsForPaymentService struct {
	Options []option.RequestOption
}

// NewInboundRealTimePaymentsRequestsForPaymentService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewInboundRealTimePaymentsRequestsForPaymentService(opts ...option.RequestOption) (r *InboundRealTimePaymentsRequestsForPaymentService) {
	r = &InboundRealTimePaymentsRequestsForPaymentService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Real-Time Payments Request for Payment
func (r *InboundRealTimePaymentsRequestsForPaymentService) Get(ctx context.Context, inboundRealTimePaymentsRequestForPaymentID string, opts ...option.RequestOption) (res *InboundRealTimePaymentsRequestForPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundRealTimePaymentsRequestForPaymentID == "" {
		err = errors.New("missing required inbound_real_time_payments_request_for_payment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbound_real_time_payments_requests_for_payment/%s", inboundRealTimePaymentsRequestForPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Inbound Real-Time Payments Requests for Payment
func (r *InboundRealTimePaymentsRequestsForPaymentService) List(ctx context.Context, query InboundRealTimePaymentsRequestsForPaymentListParams, opts ...option.RequestOption) (res *pagination.Page[InboundRealTimePaymentsRequestForPayment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_real_time_payments_requests_for_payment"
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

// List Inbound Real-Time Payments Requests for Payment
func (r *InboundRealTimePaymentsRequestsForPaymentService) ListAutoPaging(ctx context.Context, query InboundRealTimePaymentsRequestsForPaymentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundRealTimePaymentsRequestForPayment] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// An Inbound Real-Time Payments Request for Payment is a request initiated outside
// of Increase for one of your accounts to send a Real-Time Payments transfer.
type InboundRealTimePaymentsRequestForPayment struct {
	// The inbound Real-Time Payments request for payment's identifier.
	ID string `json:"id" api:"required"`
	// The Account the request for payment is for.
	AccountID string `json:"account_id" api:"required"`
	// The identifier of the Account Number the request for payment is for.
	AccountNumberID string `json:"account_number_id" api:"required"`
	// The requested amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the request for payment was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Details of the party requesting payment.
	Creditor InboundRealTimePaymentsRequestForPaymentCreditor `json:"creditor" api:"required"`
	// The creditor's account number.
	CreditorAccountNumber string `json:"creditor_account_number" api:"required"`
	// The creditor's American Bankers' Association (ABA) Routing Transit Number (RTN).
	CreditorRoutingNumber string `json:"creditor_routing_number" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the requested
	// currency. This will always be "USD" for a Real-Time Payments request for
	// payment.
	Currency InboundRealTimePaymentsRequestForPaymentCurrency `json:"currency" api:"required"`
	// The name of the account holder the payment is requested from, as provided by the
	// creditor.
	DebtorName string `json:"debtor_name" api:"required"`
	// A free-form reference string set by the creditor, to help identify the request
	// for payment.
	EndToEndIdentification string `json:"end_to_end_identification" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time after which
	// the request for payment is no longer valid and should no longer be paid.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The identifier of the Real-Time Payments Transfer that fulfilled this request
	// for payment. This is set once a transfer sent in response to the request for
	// payment has been acknowledged by the Real-Time Payments network.
	FulfillmentRealTimePaymentsTransferID string `json:"fulfillment_real_time_payments_transfer_id" api:"required,nullable"`
	// An identifier for the party that issued the invoice, for requests for payment
	// sent on behalf of another party.
	InvoicerIdentification string `json:"invoicer_identification" api:"required,nullable"`
	// The Real-Time Payments network identification of the request for payment.
	PaymentInformationIdentification string `json:"payment_information_identification" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time by which
	// the creditor requests the payment to be made.
	RequestedExecutionAt time.Time `json:"requested_execution_at" api:"required,nullable" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_request_for_payment`.
	Type InboundRealTimePaymentsRequestForPaymentType `json:"type" api:"required"`
	// Unstructured information included with the request for payment.
	UnstructuredRemittanceInformation string                                       `json:"unstructured_remittance_information" api:"required,nullable"`
	JSON                              inboundRealTimePaymentsRequestForPaymentJSON `json:"-"`
}

// inboundRealTimePaymentsRequestForPaymentJSON contains the JSON metadata for the
// struct [InboundRealTimePaymentsRequestForPayment]
type inboundRealTimePaymentsRequestForPaymentJSON struct {
	ID                                    apijson.Field
	AccountID                             apijson.Field
	AccountNumberID                       apijson.Field
	Amount                                apijson.Field
	CreatedAt                             apijson.Field
	Creditor                              apijson.Field
	CreditorAccountNumber                 apijson.Field
	CreditorRoutingNumber                 apijson.Field
	Currency                              apijson.Field
	DebtorName                            apijson.Field
	EndToEndIdentification                apijson.Field
	ExpiresAt                             apijson.Field
	FulfillmentRealTimePaymentsTransferID apijson.Field
	InvoicerIdentification                apijson.Field
	PaymentInformationIdentification      apijson.Field
	RequestedExecutionAt                  apijson.Field
	Type                                  apijson.Field
	UnstructuredRemittanceInformation     apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *InboundRealTimePaymentsRequestForPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsRequestForPaymentJSON) RawJSON() string {
	return r.raw
}

// Details of the party requesting payment.
type InboundRealTimePaymentsRequestForPaymentCreditor struct {
	// The name of the account that would receive the payment, as provided by the
	// creditor.
	AccountName string `json:"account_name" api:"required,nullable"`
	// Address of the creditor.
	Address InboundRealTimePaymentsRequestForPaymentCreditorAddress `json:"address" api:"required"`
	// The name of the creditor.
	Name string                                               `json:"name" api:"required"`
	JSON inboundRealTimePaymentsRequestForPaymentCreditorJSON `json:"-"`
}

// inboundRealTimePaymentsRequestForPaymentCreditorJSON contains the JSON metadata
// for the struct [InboundRealTimePaymentsRequestForPaymentCreditor]
type inboundRealTimePaymentsRequestForPaymentCreditorJSON struct {
	AccountName apijson.Field
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsRequestForPaymentCreditor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsRequestForPaymentCreditorJSON) RawJSON() string {
	return r.raw
}

// Address of the creditor.
type InboundRealTimePaymentsRequestForPaymentCreditorAddress struct {
	// A second address line, such as an apartment or suite number. The first address
	// line is separated into `building_number` and `street_name`.
	AddressLine2 string `json:"address_line2" api:"required,nullable"`
	// The number identifying the position of the building on the street.
	BuildingNumber string `json:"building_number" api:"required,nullable"`
	// The town or city.
	City string `json:"city" api:"required,nullable"`
	// The ISO 3166, Alpha-2 country code.
	Country string `json:"country" api:"required,nullable"`
	// The postal code or zip.
	PostalCode string `json:"postal_code" api:"required,nullable"`
	// The US state component of the address.
	State string `json:"state" api:"required,nullable"`
	// The street name without the street number.
	StreetName string                                                      `json:"street_name" api:"required,nullable"`
	JSON       inboundRealTimePaymentsRequestForPaymentCreditorAddressJSON `json:"-"`
}

// inboundRealTimePaymentsRequestForPaymentCreditorAddressJSON contains the JSON
// metadata for the struct
// [InboundRealTimePaymentsRequestForPaymentCreditorAddress]
type inboundRealTimePaymentsRequestForPaymentCreditorAddressJSON struct {
	AddressLine2   apijson.Field
	BuildingNumber apijson.Field
	City           apijson.Field
	Country        apijson.Field
	PostalCode     apijson.Field
	State          apijson.Field
	StreetName     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InboundRealTimePaymentsRequestForPaymentCreditorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsRequestForPaymentCreditorAddressJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the requested
// currency. This will always be "USD" for a Real-Time Payments request for
// payment.
type InboundRealTimePaymentsRequestForPaymentCurrency string

const (
	InboundRealTimePaymentsRequestForPaymentCurrencyUsd InboundRealTimePaymentsRequestForPaymentCurrency = "USD"
)

func (r InboundRealTimePaymentsRequestForPaymentCurrency) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsRequestForPaymentCurrencyUsd:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_request_for_payment`.
type InboundRealTimePaymentsRequestForPaymentType string

const (
	InboundRealTimePaymentsRequestForPaymentTypeInboundRealTimePaymentsRequestForPayment InboundRealTimePaymentsRequestForPaymentType = "inbound_real_time_payments_request_for_payment"
)

func (r InboundRealTimePaymentsRequestForPaymentType) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsRequestForPaymentTypeInboundRealTimePaymentsRequestForPayment:
		return true
	}
	return false
}

type InboundRealTimePaymentsRequestsForPaymentListParams struct {
	// Filter Inbound Real-Time Payments Requests for Payment to those belonging to the
	// specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound Real-Time Payments Requests for Payment to ones belonging to the
	// specified Account Number.
	AccountNumberID param.Field[string]                                                       `query:"account_number_id"`
	CreatedAt       param.Field[InboundRealTimePaymentsRequestsForPaymentListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundRealTimePaymentsRequestsForPaymentListParams]'s
// query parameters as `url.Values`.
func (r InboundRealTimePaymentsRequestsForPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundRealTimePaymentsRequestsForPaymentListParamsCreatedAt struct {
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

// URLQuery serializes
// [InboundRealTimePaymentsRequestsForPaymentListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r InboundRealTimePaymentsRequestsForPaymentListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
