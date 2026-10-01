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

// RealTimePaymentsRequestsForPaymentService contains methods and other services
// that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRealTimePaymentsRequestsForPaymentService] method instead.
type RealTimePaymentsRequestsForPaymentService struct {
	Options []option.RequestOption
}

// NewRealTimePaymentsRequestsForPaymentService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRealTimePaymentsRequestsForPaymentService(opts ...option.RequestOption) (r *RealTimePaymentsRequestsForPaymentService) {
	r = &RealTimePaymentsRequestsForPaymentService{}
	r.Options = opts
	return
}

// Create a Real-Time Payments Request for Payment
func (r *RealTimePaymentsRequestsForPaymentService) New(ctx context.Context, body RealTimePaymentsRequestsForPaymentNewParams, opts ...option.RequestOption) (res *RealTimePaymentsRequestForPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "real_time_payments_requests_for_payment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Real-Time Payments Request for Payment
func (r *RealTimePaymentsRequestsForPaymentService) Get(ctx context.Context, realTimePaymentsRequestForPaymentID string, opts ...option.RequestOption) (res *RealTimePaymentsRequestForPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimePaymentsRequestForPaymentID == "" {
		err = errors.New("missing required real_time_payments_request_for_payment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("real_time_payments_requests_for_payment/%s", realTimePaymentsRequestForPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Real-Time Payments Requests for Payment
func (r *RealTimePaymentsRequestsForPaymentService) List(ctx context.Context, query RealTimePaymentsRequestsForPaymentListParams, opts ...option.RequestOption) (res *pagination.Page[RealTimePaymentsRequestForPayment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "real_time_payments_requests_for_payment"
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

// List Real-Time Payments Requests for Payment
func (r *RealTimePaymentsRequestsForPaymentService) ListAutoPaging(ctx context.Context, query RealTimePaymentsRequestsForPaymentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[RealTimePaymentsRequestForPayment] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Cancels a Real-Time Payments Request for Payment that is still awaiting payment.
func (r *RealTimePaymentsRequestsForPaymentService) Cancel(ctx context.Context, realTimePaymentsRequestForPaymentID string, body RealTimePaymentsRequestsForPaymentCancelParams, opts ...option.RequestOption) (res *RealTimePaymentsRequestForPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimePaymentsRequestForPaymentID == "" {
		err = errors.New("missing required real_time_payments_request_for_payment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("real_time_payments_requests_for_payment/%s/cancel", realTimePaymentsRequestForPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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
	ID string `json:"id" api:"required"`
	// The Account in which a successful transfer will arrive.
	AccountID string `json:"account_id" api:"required"`
	// The Account Number in which a successful transfer will arrive.
	AccountNumberID string `json:"account_number_id" api:"required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// If a cancellation has been requested, this will contain supplemental details.
	// The request for payment moves to `canceled` once the recipient bank acknowledges
	// the cancellation.
	Cancellation RealTimePaymentsRequestForPaymentCancellation `json:"cancellation" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the request for payment was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the creditor requesting the payment.
	CreditorName string `json:"creditor_name" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For real-time payments transfers this is always equal to `USD`.
	Currency RealTimePaymentsRequestForPaymentCurrency `json:"currency" api:"required"`
	// Details of the person being requested to pay.
	Debtor RealTimePaymentsRequestForPaymentDebtor `json:"debtor" api:"required"`
	// The debtor's account number, which the request is sent to.
	DebtorAccountNumber string `json:"debtor_account_number" api:"required"`
	// The debtor's American Bankers' Association (ABA) Routing Transit Number (RTN).
	DebtorRoutingNumber string `json:"debtor_routing_number" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time after which
	// the request for payment is no longer valid. After this time the debtor's bank
	// should no longer allow the debtor to pay it.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The identifier of the Inbound Real-Time Payments Transfer that fulfilled this
	// request.
	FulfillmentInboundRealTimePaymentsTransferID string `json:"fulfillment_inbound_real_time_payments_transfer_id" api:"required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// If the request for payment is refused by the destination financial institution
	// or the receiving customer, this will contain supplemental details.
	Refusal RealTimePaymentsRequestForPaymentRefusal `json:"refusal" api:"required,nullable"`
	// If the request for payment is rejected by Real-Time Payments or the destination
	// financial institution, this will contain supplemental details.
	Rejection RealTimePaymentsRequestForPaymentRejection `json:"rejection" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time by which
	// the payment was requested to be made.
	RequestedExecutionAt time.Time `json:"requested_execution_at" api:"required,nullable" format:"date-time"`
	// The lifecycle status of the request for payment.
	Status RealTimePaymentsRequestForPaymentStatus `json:"status" api:"required"`
	// After the request for payment is submitted to Real-Time Payments, this will
	// contain supplemental details.
	Submission RealTimePaymentsRequestForPaymentSubmission `json:"submission" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_payments_request_for_payment`.
	Type RealTimePaymentsRequestForPaymentType `json:"type" api:"required"`
	// Unstructured information that will show on the recipient's bank statement.
	UnstructuredRemittanceInformation string                                `json:"unstructured_remittance_information" api:"required"`
	JSON                              realTimePaymentsRequestForPaymentJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentJSON contains the JSON metadata for the struct
// [RealTimePaymentsRequestForPayment]
type realTimePaymentsRequestForPaymentJSON struct {
	ID                                           apijson.Field
	AccountID                                    apijson.Field
	AccountNumberID                              apijson.Field
	Amount                                       apijson.Field
	Cancellation                                 apijson.Field
	CreatedAt                                    apijson.Field
	CreditorName                                 apijson.Field
	Currency                                     apijson.Field
	Debtor                                       apijson.Field
	DebtorAccountNumber                          apijson.Field
	DebtorRoutingNumber                          apijson.Field
	ExpiresAt                                    apijson.Field
	FulfillmentInboundRealTimePaymentsTransferID apijson.Field
	IdempotencyKey                               apijson.Field
	Refusal                                      apijson.Field
	Rejection                                    apijson.Field
	RequestedExecutionAt                         apijson.Field
	Status                                       apijson.Field
	Submission                                   apijson.Field
	Type                                         apijson.Field
	UnstructuredRemittanceInformation            apijson.Field
	raw                                          string
	ExtraFields                                  map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsRequestForPaymentJSON) RawJSON() string {
	return r.raw
}

// If a cancellation has been requested, this will contain supplemental details.
// The request for payment moves to `canceled` once the recipient bank acknowledges
// the cancellation.
type RealTimePaymentsRequestForPaymentCancellation struct {
	// Additional information about the cancellation, sent on to the recipient bank.
	AdditionalInformation string `json:"additional_information" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the cancellation was requested.
	CanceledAt time.Time `json:"canceled_at" api:"required" format:"date-time"`
	// The reason the request for payment was canceled.
	Reason RealTimePaymentsRequestForPaymentCancellationReason `json:"reason" api:"required"`
	JSON   realTimePaymentsRequestForPaymentCancellationJSON   `json:"-"`
}

// realTimePaymentsRequestForPaymentCancellationJSON contains the JSON metadata for
// the struct [RealTimePaymentsRequestForPaymentCancellation]
type realTimePaymentsRequestForPaymentCancellationJSON struct {
	AdditionalInformation apijson.Field
	CanceledAt            apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsRequestForPaymentCancellationJSON) RawJSON() string {
	return r.raw
}

// The reason the request for payment was canceled.
type RealTimePaymentsRequestForPaymentCancellationReason string

const (
	RealTimePaymentsRequestForPaymentCancellationReasonRequestedByCustomer RealTimePaymentsRequestForPaymentCancellationReason = "requested_by_customer"
	RealTimePaymentsRequestForPaymentCancellationReasonPaidByOtherMeans    RealTimePaymentsRequestForPaymentCancellationReason = "paid_by_other_means"
	RealTimePaymentsRequestForPaymentCancellationReasonDuplicate           RealTimePaymentsRequestForPaymentCancellationReason = "duplicate"
	RealTimePaymentsRequestForPaymentCancellationReasonWrongAmount         RealTimePaymentsRequestForPaymentCancellationReason = "wrong_amount"
)

func (r RealTimePaymentsRequestForPaymentCancellationReason) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestForPaymentCancellationReasonRequestedByCustomer, RealTimePaymentsRequestForPaymentCancellationReasonPaidByOtherMeans, RealTimePaymentsRequestForPaymentCancellationReasonDuplicate, RealTimePaymentsRequestForPaymentCancellationReasonWrongAmount:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For real-time payments transfers this is always equal to `USD`.
type RealTimePaymentsRequestForPaymentCurrency string

const (
	RealTimePaymentsRequestForPaymentCurrencyUsd RealTimePaymentsRequestForPaymentCurrency = "USD"
)

func (r RealTimePaymentsRequestForPaymentCurrency) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestForPaymentCurrencyUsd:
		return true
	}
	return false
}

// Details of the person being requested to pay.
type RealTimePaymentsRequestForPaymentDebtor struct {
	// Address of the debtor.
	Address RealTimePaymentsRequestForPaymentDebtorAddress `json:"address" api:"required"`
	// The name of the debtor.
	Name string                                      `json:"name" api:"required"`
	JSON realTimePaymentsRequestForPaymentDebtorJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentDebtorJSON contains the JSON metadata for the
// struct [RealTimePaymentsRequestForPaymentDebtor]
type realTimePaymentsRequestForPaymentDebtorJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentDebtor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsRequestForPaymentDebtorJSON) RawJSON() string {
	return r.raw
}

// Address of the debtor.
type RealTimePaymentsRequestForPaymentDebtorAddress struct {
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
	StreetName string                                             `json:"street_name" api:"required,nullable"`
	JSON       realTimePaymentsRequestForPaymentDebtorAddressJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentDebtorAddressJSON contains the JSON metadata
// for the struct [RealTimePaymentsRequestForPaymentDebtorAddress]
type realTimePaymentsRequestForPaymentDebtorAddressJSON struct {
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

func (r *RealTimePaymentsRequestForPaymentDebtorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsRequestForPaymentDebtorAddressJSON) RawJSON() string {
	return r.raw
}

// If the request for payment is refused by the destination financial institution
// or the receiving customer, this will contain supplemental details.
type RealTimePaymentsRequestForPaymentRefusal struct {
	// Additional information about the refusal provided by the recipient bank or the
	// customer. This is typically present when the `refusal_reason_code` is `other`.
	RefusalReasonAdditionalInformation string `json:"refusal_reason_additional_information" api:"required,nullable"`
	// The reason the request for payment was refused as provided by the recipient bank
	// or the customer.
	RefusalReasonCode RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode `json:"refusal_reason_code" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the request for payment was refused.
	RefusedAt time.Time                                    `json:"refused_at" api:"required,nullable" format:"date-time"`
	JSON      realTimePaymentsRequestForPaymentRefusalJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentRefusalJSON contains the JSON metadata for the
// struct [RealTimePaymentsRequestForPaymentRefusal]
type realTimePaymentsRequestForPaymentRefusalJSON struct {
	RefusalReasonAdditionalInformation apijson.Field
	RefusalReasonCode                  apijson.Field
	RefusedAt                          apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentRefusal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsRequestForPaymentRefusalJSON) RawJSON() string {
	return r.raw
}

// The reason the request for payment was refused as provided by the recipient bank
// or the customer.
type RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode string

const (
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeAccountBlocked              RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "account_blocked"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeTransactionForbidden        RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "transaction_forbidden"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeTransactionTypeNotSupported RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "transaction_type_not_supported"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeUnexpectedAmount            RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "unexpected_amount"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeAmountExceedsBankLimits     RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "amount_exceeds_bank_limits"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeInvalidDebtorAddress        RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "invalid_debtor_address"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeInvalidCreditorAddress      RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "invalid_creditor_address"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeCreditorIdentifierIncorrect RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "creditor_identifier_incorrect"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeRequestedByCustomer         RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "requested_by_customer"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeOrderRejected               RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "order_rejected"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeEndCustomerDeceased         RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "end_customer_deceased"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeCustomerHasOptedOut         RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "customer_has_opted_out"
	RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeOther                       RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode = "other"
)

func (r RealTimePaymentsRequestForPaymentRefusalRefusalReasonCode) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeAccountBlocked, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeTransactionForbidden, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeTransactionTypeNotSupported, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeUnexpectedAmount, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeAmountExceedsBankLimits, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeInvalidDebtorAddress, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeInvalidCreditorAddress, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeCreditorIdentifierIncorrect, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeRequestedByCustomer, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeOrderRejected, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeEndCustomerDeceased, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeCustomerHasOptedOut, RealTimePaymentsRequestForPaymentRefusalRefusalReasonCodeOther:
		return true
	}
	return false
}

// If the request for payment is rejected by Real-Time Payments or the destination
// financial institution, this will contain supplemental details.
type RealTimePaymentsRequestForPaymentRejection struct {
	// Additional information about the rejection provided by the recipient bank or the
	// Real-Time Payments network. This is typically present when the
	// `reject_reason_code` is `narrative`.
	RejectReasonAdditionalInformation string `json:"reject_reason_additional_information" api:"required,nullable"`
	// The reason the request for payment was rejected as provided by the recipient
	// bank or the Real-Time Payments network.
	RejectReasonCode RealTimePaymentsRequestForPaymentRejectionRejectReasonCode `json:"reject_reason_code" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the request for payment was rejected.
	RejectedAt time.Time                                      `json:"rejected_at" api:"required,nullable" format:"date-time"`
	JSON       realTimePaymentsRequestForPaymentRejectionJSON `json:"-"`
}

// realTimePaymentsRequestForPaymentRejectionJSON contains the JSON metadata for
// the struct [RealTimePaymentsRequestForPaymentRejection]
type realTimePaymentsRequestForPaymentRejectionJSON struct {
	RejectReasonAdditionalInformation apijson.Field
	RejectReasonCode                  apijson.Field
	RejectedAt                        apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *RealTimePaymentsRequestForPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsRequestForPaymentRejectionJSON) RawJSON() string {
	return r.raw
}

// The reason the request for payment was rejected as provided by the recipient
// bank or the Real-Time Payments network.
type RealTimePaymentsRequestForPaymentRejectionRejectReasonCode string

const (
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAccountClosed                                 RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "account_closed"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAccountBlocked                                RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "account_blocked"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAccountType                    RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_account_type"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAccountNumber                  RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_account_number"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeEndCustomerDeceased                           RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "end_customer_deceased"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeNarrative                                     RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "narrative"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTransactionForbidden                          RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "transaction_forbidden"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTransactionTypeNotSupported                   RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "transaction_type_not_supported"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnexpectedAmount                              RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "unexpected_amount"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAmountExceedsBankLimits                       RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAddress                        RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_creditor_address"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnknownEndCustomer                            RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "unknown_end_customer"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidDebtorAddress                          RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "invalid_debtor_address"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTimeout                                       RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "timeout"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnsupportedMessageForRecipient                RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "unsupported_message_for_recipient"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeRecipientConnectionNotAvailable               RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "recipient_connection_not_available"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeRealTimePaymentsSuspended                     RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "real_time_payments_suspended"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInstructedAgentSignedOff                      RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "instructed_agent_signed_off"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeProcessingError                               RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "processing_error"
	RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeOther                                         RealTimePaymentsRequestForPaymentRejectionRejectReasonCode = "other"
)

func (r RealTimePaymentsRequestForPaymentRejectionRejectReasonCode) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAccountClosed, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAccountBlocked, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAccountType, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAccountNumber, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeEndCustomerDeceased, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeNarrative, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTransactionForbidden, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTransactionTypeNotSupported, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnexpectedAmount, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeAmountExceedsBankLimits, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidCreditorAddress, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnknownEndCustomer, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInvalidDebtorAddress, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeTimeout, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeUnsupportedMessageForRecipient, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeRecipientConnectionNotAvailable, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeRealTimePaymentsSuspended, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeInstructedAgentSignedOff, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeProcessingError, RealTimePaymentsRequestForPaymentRejectionRejectReasonCodeOther:
		return true
	}
	return false
}

// The lifecycle status of the request for payment.
type RealTimePaymentsRequestForPaymentStatus string

const (
	RealTimePaymentsRequestForPaymentStatusPendingSubmission RealTimePaymentsRequestForPaymentStatus = "pending_submission"
	RealTimePaymentsRequestForPaymentStatusPendingResponse   RealTimePaymentsRequestForPaymentStatus = "pending_response"
	RealTimePaymentsRequestForPaymentStatusRejected          RealTimePaymentsRequestForPaymentStatus = "rejected"
	RealTimePaymentsRequestForPaymentStatusAccepted          RealTimePaymentsRequestForPaymentStatus = "accepted"
	RealTimePaymentsRequestForPaymentStatusRefused           RealTimePaymentsRequestForPaymentStatus = "refused"
	RealTimePaymentsRequestForPaymentStatusFulfilled         RealTimePaymentsRequestForPaymentStatus = "fulfilled"
	RealTimePaymentsRequestForPaymentStatusCanceled          RealTimePaymentsRequestForPaymentStatus = "canceled"
)

func (r RealTimePaymentsRequestForPaymentStatus) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestForPaymentStatusPendingSubmission, RealTimePaymentsRequestForPaymentStatusPendingResponse, RealTimePaymentsRequestForPaymentStatusRejected, RealTimePaymentsRequestForPaymentStatusAccepted, RealTimePaymentsRequestForPaymentStatusRefused, RealTimePaymentsRequestForPaymentStatusFulfilled, RealTimePaymentsRequestForPaymentStatusCanceled:
		return true
	}
	return false
}

// After the request for payment is submitted to Real-Time Payments, this will
// contain supplemental details.
type RealTimePaymentsRequestForPaymentSubmission struct {
	// The Real-Time Payments payment information identification of the request.
	PaymentInformationIdentification string                                          `json:"payment_information_identification" api:"required"`
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

func (r realTimePaymentsRequestForPaymentSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `real_time_payments_request_for_payment`.
type RealTimePaymentsRequestForPaymentType string

const (
	RealTimePaymentsRequestForPaymentTypeRealTimePaymentsRequestForPayment RealTimePaymentsRequestForPaymentType = "real_time_payments_request_for_payment"
)

func (r RealTimePaymentsRequestForPaymentType) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestForPaymentTypeRealTimePaymentsRequestForPayment:
		return true
	}
	return false
}

type RealTimePaymentsRequestsForPaymentNewParams struct {
	// The identifier of the Account Number where the funds will land.
	AccountNumberID param.Field[string] `json:"account_number_id" api:"required"`
	// The requested amount in USD cents. Must be positive.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// Details of the person being requested to pay.
	Debtor param.Field[RealTimePaymentsRequestsForPaymentNewParamsDebtor] `json:"debtor" api:"required"`
	// The debtor's account number, which the funds will be requested from.
	DebtorAccountNumber param.Field[string] `json:"debtor_account_number" api:"required"`
	// The debtor's American Bankers' Association (ABA) Routing Transit Number (RTN).
	DebtorRoutingNumber param.Field[string] `json:"debtor_routing_number" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time after which
	// the request for payment is no longer valid. After this time the debtor's bank
	// should no longer allow the debtor to pay it. Must not be before
	// `requested_execution_at`.
	ExpiresAt param.Field[time.Time] `json:"expires_at" api:"required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time by which
	// you are requesting the payment to be made.
	RequestedExecutionAt param.Field[time.Time] `json:"requested_execution_at" api:"required" format:"date-time"`
	// Unstructured information that will show on the recipient's bank statement.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information" api:"required"`
	// The name of the creditor requesting the payment. If not provided, defaults to
	// the name of the account's entity.
	CreditorName param.Field[string] `json:"creditor_name"`
}

func (r RealTimePaymentsRequestsForPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the person being requested to pay.
type RealTimePaymentsRequestsForPaymentNewParamsDebtor struct {
	// Address of the debtor.
	Address param.Field[RealTimePaymentsRequestsForPaymentNewParamsDebtorAddress] `json:"address" api:"required"`
	// The name of the debtor.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r RealTimePaymentsRequestsForPaymentNewParamsDebtor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Address of the debtor.
type RealTimePaymentsRequestsForPaymentNewParamsDebtorAddress struct {
	// The ISO 3166, Alpha-2 country code.
	//
	// Defaults to `US`.
	Country param.Field[string] `json:"country" api:"required"`
	// A second address line, such as an apartment or suite number. The first address
	// line is separated into `building_number` and `street_name`.
	AddressLine2 param.Field[string] `json:"address_line2"`
	// The number identifying the position of the building on the street.
	BuildingNumber param.Field[string] `json:"building_number"`
	// The town or city.
	City param.Field[string] `json:"city"`
	// The postal code or zip.
	PostalCode param.Field[string] `json:"postal_code"`
	// The US state component of the address.
	State param.Field[string] `json:"state"`
	// The street name without the street number.
	StreetName param.Field[string] `json:"street_name"`
}

func (r RealTimePaymentsRequestsForPaymentNewParamsDebtorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RealTimePaymentsRequestsForPaymentListParams struct {
	// Filter Real-Time Payments Requests for Payment to those destined to the
	// specified Account.
	AccountID param.Field[string]                                                `query:"account_id"`
	CreatedAt param.Field[RealTimePaymentsRequestsForPaymentListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RealTimePaymentsRequestsForPaymentListParams]'s query
// parameters as `url.Values`.
func (r RealTimePaymentsRequestsForPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsRequestsForPaymentListParamsCreatedAt struct {
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

// URLQuery serializes [RealTimePaymentsRequestsForPaymentListParamsCreatedAt]'s
// query parameters as `url.Values`.
func (r RealTimePaymentsRequestsForPaymentListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsRequestsForPaymentCancelParams struct {
	// Additional information about the cancellation to pass on to the recipient bank.
	AdditionalInformation param.Field[string] `json:"additional_information"`
	// The reason the request for payment is being canceled. Defaults to
	// `requested_by_customer`.
	Reason param.Field[RealTimePaymentsRequestsForPaymentCancelParamsReason] `json:"reason"`
}

func (r RealTimePaymentsRequestsForPaymentCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason the request for payment is being canceled. Defaults to
// `requested_by_customer`.
type RealTimePaymentsRequestsForPaymentCancelParamsReason string

const (
	RealTimePaymentsRequestsForPaymentCancelParamsReasonRequestedByCustomer RealTimePaymentsRequestsForPaymentCancelParamsReason = "requested_by_customer"
	RealTimePaymentsRequestsForPaymentCancelParamsReasonPaidByOtherMeans    RealTimePaymentsRequestsForPaymentCancelParamsReason = "paid_by_other_means"
	RealTimePaymentsRequestsForPaymentCancelParamsReasonDuplicate           RealTimePaymentsRequestsForPaymentCancelParamsReason = "duplicate"
	RealTimePaymentsRequestsForPaymentCancelParamsReasonWrongAmount         RealTimePaymentsRequestsForPaymentCancelParamsReason = "wrong_amount"
)

func (r RealTimePaymentsRequestsForPaymentCancelParamsReason) IsKnown() bool {
	switch r {
	case RealTimePaymentsRequestsForPaymentCancelParamsReasonRequestedByCustomer, RealTimePaymentsRequestsForPaymentCancelParamsReasonPaidByOtherMeans, RealTimePaymentsRequestsForPaymentCancelParamsReasonDuplicate, RealTimePaymentsRequestsForPaymentCancelParamsReasonWrongAmount:
		return true
	}
	return false
}
