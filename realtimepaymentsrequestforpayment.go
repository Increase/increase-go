// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// RealTimePaymentsRequestForPaymentService contains methods and other services
// that help with interacting with the increase API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRealTimePaymentsRequestForPaymentService] method instead.
type RealTimePaymentsRequestForPaymentService struct {
	Options []option.RequestOption
}

// NewRealTimePaymentsRequestForPaymentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRealTimePaymentsRequestForPaymentService(opts ...option.RequestOption) (r *RealTimePaymentsRequestForPaymentService) {
	r = &RealTimePaymentsRequestForPaymentService{}
	r.Options = opts
	return
}

// Create a Real-Time Payments Request for Payment
func (r *RealTimePaymentsRequestForPaymentService) New(ctx context.Context, body RealTimePaymentsRequestForPaymentNewParams, opts ...option.RequestOption) (res *RealTimePaymentsRequestForPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "real_time_payments_request_for_payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Real-Time Payments Request for Payment
func (r *RealTimePaymentsRequestForPaymentService) Get(ctx context.Context, requestForPaymentID string, opts ...option.RequestOption) (res *RealTimePaymentsRequestForPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_payments_request_for_payments/%s", requestForPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Real-Time Payments Request for Payments
func (r *RealTimePaymentsRequestForPaymentService) List(ctx context.Context, query RealTimePaymentsRequestForPaymentListParams, opts ...option.RequestOption) (res *shared.Page[RealTimePaymentsRequestForPayment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "real_time_payments_request_for_payments"
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

// List Real-Time Payments Request for Payments
func (r *RealTimePaymentsRequestForPaymentService) ListAutoPaging(ctx context.Context, query RealTimePaymentsRequestForPaymentListParams, opts ...option.RequestOption) *shared.PageAutoPager[RealTimePaymentsRequestForPayment] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Real-Time Payments transfers move funds, within seconds, between your Increase
// account and any other account on the Real-Time Payments network. A request for
// payment is a request to the receiver to send funds to your account. The
// permitted uses of Requests For Payment are limited by the Real-Time Payments
// network to business-to-business payments and transfers between two accounts at
// different banks owned by the same individual. Please contact
// [support@increase.com](mailto:support@increase.com) to enable this API for your
// team.
type RealTimePaymentsRequestForPayment struct {
	// The Real-Time Payments Request for Payment's identifier.
	ID string `json:"id,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the request for payment was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For real-time payments transfers this is always equal to `USD`.
	Currency RealTimePaymentsRequestForPaymentCurrency `json:"currency,required"`
	// The name of the recipient the sender is requesting a transfer from.
	DebtorName string `json:"debtor_name,required"`
	// The Account Number in which a successful transfer will arrive.
	DestinationAccountNumberID string `json:"destination_account_number_id,required"`
	// The expiration time for this request, in UTC. The requestee will not be able to
	// pay after this date.
	ExpiresAt time.Time `json:"expires_at,required" format:"date"`
	// The transaction that fulfilled this request.
	FulfillmentTransactionID string `json:"fulfillment_transaction_id,required,nullable"`
	// If the request for payment is refused by the destination financial institution
	// or the receiving customer, this will contain supplemental details.
	Refusal RealTimePaymentsRequestForPaymentRefusal `json:"refusal,required,nullable"`
	// If the request for payment is rejected by Real-Time Payments or the destination
	// financial institution, this will contain supplemental details.
	Rejection RealTimePaymentsRequestForPaymentRejection `json:"rejection,required,nullable"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The account number the request is sent to.
	SourceAccountNumber string `json:"source_account_number,required"`
	// The receiver's American Bankers' Association (ABA) Routing Transit Number (RTN).
	SourceRoutingNumber string `json:"source_routing_number,required"`
	// The lifecycle status of the request for payment.
	Status RealTimePaymentsRequestForPaymentStatus `json:"status,required"`
	// After the request for payment is submitted to Real-Time Payments, this will
	// contain supplemental details.
	Submission RealTimePaymentsRequestForPaymentSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_payments_request_for_payment`.
	Type RealTimePaymentsRequestForPaymentType `json:"type,required"`
	JSON realTimePaymentsRequestForPaymentJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentJSON contains the JSON metadata for the struct
// [RealTimePaymentsRequestForPayment]
type realTimePaymentsRequestForPaymentJSON struct {
	ID                         apijson.Field
	Amount                     apijson.Field
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	DebtorName                 apijson.Field
	DestinationAccountNumberID apijson.Field
	ExpiresAt                  apijson.Field
	FulfillmentTransactionID   apijson.Field
	Refusal                    apijson.Field
	Rejection                  apijson.Field
	RemittanceInformation      apijson.Field
	SourceAccountNumber        apijson.Field
	SourceRoutingNumber        apijson.Field
	Status                     apijson.Field
	Submission                 apijson.Field
	Type                       apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For real-time payments transfers this is always equal to `USD`.
type RealTimePaymentsRequestForPaymentCurrency string

const (
	// Canadian Dollar (CAD)
	RealTimePaymentsRequestForPaymentCurrencyCad RealTimePaymentsRequestForPaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	RealTimePaymentsRequestForPaymentCurrencyChf RealTimePaymentsRequestForPaymentCurrency = "CHF"
	// Euro (EUR)
	RealTimePaymentsRequestForPaymentCurrencyEur RealTimePaymentsRequestForPaymentCurrency = "EUR"
	// British Pound (GBP)
	RealTimePaymentsRequestForPaymentCurrencyGbp RealTimePaymentsRequestForPaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	RealTimePaymentsRequestForPaymentCurrencyJpy RealTimePaymentsRequestForPaymentCurrency = "JPY"
	// US Dollar (USD)
	RealTimePaymentsRequestForPaymentCurrencyUsd RealTimePaymentsRequestForPaymentCurrency = "USD"
)

// If the request for payment is refused by the destination financial institution
// or the receiving customer, this will contain supplemental details.
type RealTimePaymentsRequestForPaymentRefusal struct {
	// The reason the request for payment was refused as provided by the recipient bank
	// or the customer.
	RefusalReasonCode RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode `json:"refusal_reason_code,required"`
	JSON              realTimePaymentsRequestForPaymentRefusalJSON              `json:"-"`
}

// realTimePaymentsRequestForPaymentRefusalJSON contains the JSON metadata for the
// struct [RealTimePaymentsRequestForPaymentRefusal]
type realTimePaymentsRequestForPaymentRefusalJSON struct {
	RefusalReasonCode apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentRefusal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason the request for payment was refused as provided by the recipient bank
// or the customer.
type RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode string

const (
	// The destination account is currently blocked from receiving transactions.
	// Corresponds to the Real-Time Payments reason code `AC06`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeAccountBlocked RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "account_blocked"
	// Real-Time Payments transfers are not allowed to the destination account.
	// Corresponds to the Real-Time Payments reason code `AG01`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeTransactionForbidden RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "transaction_forbidden"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `AG03`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeTransactionTypeNotSupported RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "transaction_type_not_supported"
	// The amount of the transfer is different than expected by the recipient.
	// Corresponds to the Real-Time Payments reason code `AM09`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeUnexpectedAmount RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "unexpected_amount"
	// The amount is higher than the recipient is authorized to send or receive.
	// Corresponds to the Real-Time Payments reason code `AM14`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeAmountExceedsBankLimits RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "amount_exceeds_bank_limits"
	// The debtor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE07`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeInvalidDebtorAddress RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "invalid_debtor_address"
	// The creditor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE04`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeInvalidCreditorAddress RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "invalid_creditor_address"
	// Creditor identifier incorrect. Corresponds to the Real-Time Payments reason code
	// `CH11`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeCreditorIdentifierIncorrect RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "creditor_identifier_incorrect"
	// The customer refused the request. Corresponds to the Real-Time Payments reason
	// code `CUST`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeRequestedByCustomer RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "requested_by_customer"
	// The order was rejected. Corresponds to the Real-Time Payments reason code
	// `DS04`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeOrderRejected RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "order_rejected"
	// The destination account holder is deceased. Corresponds to the Real-Time
	// Payments reason code `MD07`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeEndCustomerDeceased RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "end_customer_deceased"
	// The customer has opted out of receiving requests for payments from this
	// creditor. Corresponds to the Real-Time Payments reason code `SL12`.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeCustomerHasOptedOut RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "customer_has_opted_out"
	// Some other error or issue has occurred.
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeOther RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "other"
)

// If the request for payment is rejected by Real-Time Payments or the destination
// financial institution, this will contain supplemental details.
type RealTimePaymentsRequestForPaymentRejection struct {
	// The reason the request for payment was rejected as provided by the recipient
	// bank or the Real-Time Payments network.
	RejectReasonCode RealTimePaymentsRequestForPaymentRejectionRejectReasonCode `json:"reject_reason_code,required"`
	JSON             realTimePaymentsRequestForPaymentRejectionJSON             `json:"-"`
}

// realTimePaymentsRequestForPaymentRejectionJSON contains the JSON metadata for
// the struct [RealTimePaymentsRequestForPaymentRejection]
type realTimePaymentsRequestForPaymentRejectionJSON struct {
	RejectReasonCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason the request for payment was rejected as provided by the recipient
// bank or the Real-Time Payments network.
type RealTimePaymentsRequestForPaymentRejectionRejectReasonCode string

const (
	// The destination account is closed. Corresponds to the Real-Time Payments reason
	// code `AC04`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAccountClosed RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "account_closed"
	// The destination account is currently blocked from receiving transactions.
	// Corresponds to the Real-Time Payments reason code `AC06`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAccountBlocked RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "account_blocked"
	// The destination account is ineligible to receive Real-Time Payments transfers.
	// Corresponds to the Real-Time Payments reason code `AC14`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAccountType RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_account_type"
	// The destination account does not exist. Corresponds to the Real-Time Payments
	// reason code `AC03`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAccountNumber RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_account_number"
	// The destination routing number is invalid. Corresponds to the Real-Time Payments
	// reason code `RC04`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	// The destination account holder is deceased. Corresponds to the Real-Time
	// Payments reason code `MD07`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeEndCustomerDeceased RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "end_customer_deceased"
	// The reason is provided as narrative information in the additional information
	// field.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeNarrative RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "narrative"
	// Real-Time Payments transfers are not allowed to the destination account.
	// Corresponds to the Real-Time Payments reason code `AG01`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTransactionForbidden RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "transaction_forbidden"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `AG03`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTransactionTypeNotSupported RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "transaction_type_not_supported"
	// The amount of the transfer is different than expected by the recipient.
	// Corresponds to the Real-Time Payments reason code `AM09`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnexpectedAmount RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "unexpected_amount"
	// The amount is higher than the recipient is authorized to send or receive.
	// Corresponds to the Real-Time Payments reason code `AM14`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAmountExceedsBankLimits RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	// The creditor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE04`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAddress RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_address"
	// The specified creditor is unknown. Corresponds to the Real-Time Payments reason
	// code `BE06`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnknownEndCustomer RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "unknown_end_customer"
	// The debtor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE07`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidDebtorAddress RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_debtor_address"
	// There was a timeout processing the transfer. Corresponds to the Real-Time
	// Payments reason code `DS24`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTimeout RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "timeout"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `NOAT`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnsupportedMessageForRecipient RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "unsupported_message_for_recipient"
	// The destination financial institution is currently not connected to Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9912`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeRecipientConnectionNotAvailable RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "recipient_connection_not_available"
	// Real-Time Payments is currently unavailable. Corresponds to the Real-Time
	// Payments reason code `9948`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeRealTimePaymentsSuspended RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "real_time_payments_suspended"
	// The destination financial institution is currently signed off of Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9910`.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInstructedAgentSignedOff RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "instructed_agent_signed_off"
	// The transfer was rejected due to an internal Increase issue. We have been
	// notified.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeProcessingError RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "processing_error"
	// Some other error or issue has occurred.
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeOther RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "other"
)

// The lifecycle status of the request for payment.
type RealTimePaymentsRequestForPaymentStatus string

const (
	// The request for payment is queued to be submitted to Real-Time Payments.
	RealTimePaymentsRequestForPaymentStatusPendingSubmission RealTimePaymentsRequestForPaymentStatus = "pending_submission"
	// The request for payment has been submitted and is pending a response from
	// Real-Time Payments.
	RealTimePaymentsRequestForPaymentStatusPendingResponse RealTimePaymentsRequestForPaymentStatus = "pending_response"
	// The request for payment was rejected by the network or the recipient.
	RealTimePaymentsRequestForPaymentStatusRejected RealTimePaymentsRequestForPaymentStatus = "rejected"
	// The request for payment was accepted by the recipient but has not yet been paid.
	RealTimePaymentsRequestForPaymentStatusAccepted RealTimePaymentsRequestForPaymentStatus = "accepted"
	// The request for payment was refused by the recipient.
	RealTimePaymentsRequestForPaymentStatusRefused RealTimePaymentsRequestForPaymentStatus = "refused"
	// The request for payment was fulfilled by the receiver.
	RealTimePaymentsRequestForPaymentStatusFulfilled RealTimePaymentsRequestForPaymentStatus = "fulfilled"
)

// After the request for payment is submitted to Real-Time Payments, this will
// contain supplemental details.
type RealTimePaymentsRequestForPaymentSubmission struct {
	// The Real-Time Payments payment information identification of the request.
	PaymentInformationIdentification string                                          `json:"payment_information_identification,required"`
	JSON                             realTimePaymentsRequestForPaymentSubmissionJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentSubmissionJSON contains the JSON metadata for
// the struct [RealTimePaymentsRequestForPaymentSubmission]
type realTimePaymentsRequestForPaymentSubmissionJSON struct {
	PaymentInformationIdentification apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `real_time_payments_request_for_payment`.
type RealTimePaymentsRequestForPaymentType string

const (
	RealTimePaymentsRequestForPaymentTypeRealTimePaymentsRequestForPayment RealTimePaymentsRequestForPaymentType = "real_time_payments_request_for_payment"
)

type RealTimePaymentsRequestForPaymentNewParams struct {
	// The requested amount in USD cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// Details of the person being requested to pay.
	Debtor param.Field[RealTimePaymentsRequestForPaymentNewParamsDebtor] `json:"debtor,required"`
	// The identifier of the Account Number where the funds will land.
	DestinationAccountNumberID param.Field[string] `json:"destination_account_number_id,required"`
	// The expiration time for this request, in UTC. The requestee will not be able to
	// pay after this date.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date"`
	// Unstructured information that will show on the requestee's bank statement.
	RemittanceInformation param.Field[string] `json:"remittance_information,required"`
	// The account number the funds will be requested from.
	SourceAccountNumber param.Field[string] `json:"source_account_number,required"`
	// The requestee's American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	SourceRoutingNumber param.Field[string] `json:"source_routing_number,required"`
}

func (r RealTimePaymentsRequestForPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the person being requested to pay.
type RealTimePaymentsRequestForPaymentNewParamsDebtor struct {
	// Address of the debtor.
	Address param.Field[RealTimePaymentsRequestForPaymentNewParamsDebtorAddress] `json:"address,required"`
	// The name of the debtor.
	Name param.Field[string] `json:"name,required"`
}

func (r RealTimePaymentsRequestForPaymentNewParamsDebtor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Address of the debtor.
type RealTimePaymentsRequestForPaymentNewParamsDebtorAddress struct {
	// The ISO 3166, Alpha-2 country code.
	Country param.Field[string] `json:"country,required"`
	// The town or city.
	City param.Field[string] `json:"city"`
	// The postal code or zip.
	PostCode param.Field[string] `json:"post_code"`
	// The street name without the street number.
	StreetName param.Field[string] `json:"street_name"`
}

func (r RealTimePaymentsRequestForPaymentNewParamsDebtorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RealTimePaymentsRequestForPaymentListParams struct {
	// Filter Real-Time Payments Request for Payments to those destined to the
	// specified Account.
	AccountID param.Field[string]                                               `query:"account_id"`
	CreatedAt param.Field[RealTimePaymentsRequestForPaymentListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RealTimePaymentsRequestForPaymentListParams]'s query
// parameters as `url.Values`.
func (r RealTimePaymentsRequestForPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsRequestForPaymentListParamsCreatedAt struct {
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

// URLQuery serializes [RealTimePaymentsRequestForPaymentListParamsCreatedAt]'s
// query parameters as `url.Values`.
func (r RealTimePaymentsRequestForPaymentListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
