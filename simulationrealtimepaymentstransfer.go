// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationRealTimePaymentsTransferService contains methods and other services
// that help with interacting with the increase API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSimulationRealTimePaymentsTransferService] method instead.
type SimulationRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewSimulationRealTimePaymentsTransferService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationRealTimePaymentsTransferService(opts ...option.RequestOption) (r *SimulationRealTimePaymentsTransferService) {
	r = &SimulationRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Simulates submission of a Real-Time Payments transfer and handling the response
// from the destination financial institution. This transfer must first have a
// `status` of `pending_submission`.
func (r *SimulationRealTimePaymentsTransferService) Complete(ctx context.Context, realTimePaymentsTransferID string, body SimulationRealTimePaymentsTransferCompleteParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/real_time_payments_transfers/%s/complete", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an inbound Real-Time Payments transfer to your account. Real-Time
// Payments are a beta feature.
func (r *SimulationRealTimePaymentsTransferService) NewInbound(ctx context.Context, body SimulationRealTimePaymentsTransferNewInboundParams, opts ...option.RequestOption) (res *InboundRealTimePaymentsTransferSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of an inbound Real-Time Payments Transfer simulation.
type InboundRealTimePaymentsTransferSimulationResult struct {
	// If the Real-Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// If the Real-Time Payments Transfer attempt succeeds, this will contain the
	// resulting [Transaction](#transactions) object. The Transaction's `source` will
	// be of `category: inbound_real_time_payments_transfer_confirmation`.
	Transaction InboundRealTimePaymentsTransferSimulationResultTransaction `json:"transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer_simulation_result`.
	Type InboundRealTimePaymentsTransferSimulationResultType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultJSON
}

// inboundRealTimePaymentsTransferSimulationResultJSON contains the JSON metadata
// for the struct [InboundRealTimePaymentsTransferSimulationResult]
type inboundRealTimePaymentsTransferSimulationResultJSON struct {
	DeclinedTransaction apijson.Field
	Transaction         apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the Real-Time Payments Transfer attempt fails, this will contain the
// resulting [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of
// `category: inbound_real_time_payments_transfer_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction struct {
	// The Declined Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transaction's Account.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency `json:"currency,required"`
	// This is the description the vendor provides.
	Description string `json:"description,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionJSON contains
// the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transaction's Account.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "USD"
)

// The type of the route this Declined Transaction came through.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType string

const (
	// An Account Number.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteTypeAccountNumber InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType = "account_number"
	// A Card.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteTypeCard InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource struct {
	// An ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory `json:"category,required"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// An International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceJSON struct {
	ACHDecline                             apijson.Field
	CardDecline                            apijson.Field
	Category                               apijson.Field
	CheckDecline                           apijson.Field
	InboundRealTimePaymentsTransferDecline apijson.Field
	InternationalACHDecline                apijson.Field
	WireDecline                            apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline struct {
	// The ACH Decline's identifier.
	ID string `json:"id,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The descriptive date of the transfer.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date,required,nullable"`
	// The additional information included with the transfer.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	// The identifier of the company that initiated the transfer.
	OriginatorCompanyID string `json:"originator_company_id,required"`
	// The name of the company that initiated the transfer.
	OriginatorCompanyName string `json:"originator_company_name,required"`
	// Why the ACH transfer was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	// The id of the receiver of the transfer.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the receiver of the transfer.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// The trace number of the transfer.
	TraceNumber string `json:"trace_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_decline`.
	Type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON struct {
	ID                                 apijson.Field
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	Reason                             apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	Type                               apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH transfer was declined.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason string

const (
	// The account number is canceled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	// The transaction would cause an Increase limit to be exceeded.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonBreachesLimit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused. This is a reasonable default reason for decline of
	// credits.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// A rare return reason. The return this message refers to was a duplicate.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonDuplicateReturn InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// Your account is inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// A rare return reason. The return this message refers to was misrouted.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonMisroutedReturn InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// The originating financial institution made a mistake and this return corrects
	// it.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "return_of_erroneous_or_reversing_debit"
	// The account number that was debited does not exist.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonNoACHRoute InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// The originating financial institution asked for this transfer to be returned.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonOriginatorRequest InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	// The user initiated the decline.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonUserInitiated InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "user_initiated"
)

// A constant representing the object's type. For this resource it will always be
// `ach_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineType string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineTypeACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineType = "ach_decline"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline struct {
	// The Card Decline identifier.
	ID string `json:"id,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Fields specific to the `network`.
	NetworkDetails InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// Fields related to verification of cardholder-provided values.
	Verification InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerification `json:"verification,required"`
	JSON         inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantState        apijson.Field
	NetworkDetails       apijson.Field
	PhysicalCardID       apijson.Field
	RealTimeDecisionID   apijson.Field
	Reason               apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

// Fields specific to the `network`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The payment network used to process this card authorization.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsCategory string

const (
	// Visa
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
)

// Fields specific to the `visa` network.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

// The method used to enter the cardholder's primary account number and card
// expiration date.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

// Why the transaction was declined.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCardNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	// The Physical Card was not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "physical_card_not_active"
	// The account's entity was not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	// The account was inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by a Limit.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonBreachesLimit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	// Your application declined the transaction via webhook.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookDeclined InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	// Your application webhook did not respond without the required timeout.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	// Declined by stand-in processing.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	// The card read had an invalid CVV, dCVV, or authorization request cryptogram.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	// The original card authorization for this incremental authorization does not
	// exist.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
	// The transaction was suspected to be fraudulent. Please reach out to
	// support@increase.com for more information.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonSuspectedFraud InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "suspected_fraud"
)

// Fields related to verification of cardholder-provided values.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerification]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The result of verifying the Card Verification Code.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNotChecked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultMatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNoMatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "no_match"
)

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddress struct {
	// Line 1 of the address on file for the cardholder.
	ActualLine1 string `json:"actual_line1,required,nullable"`
	// The postal code of the address on file for the cardholder.
	ActualPostalCode string `json:"actual_postal_code,required,nullable"`
	// The cardholder address line 1 provided for verification in the authorization
	// request.
	ProvidedLine1 string `json:"provided_line1,required,nullable"`
	// The postal code provided for verification in the authorization request.
	ProvidedPostalCode string `json:"provided_postal_code,required,nullable"`
	// The address verification result returned to the card network.
	Result InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult `json:"result,required"`
	JSON   inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddress]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The address verification result returned to the card network.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "no_match"
)

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory string

const (
	// ACH Decline: details will be under the `ach_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "ach_decline"
	// Card Decline: details will be under the `card_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCardDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "card_decline"
	// Check Decline: details will be under the `check_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCheckDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "check_decline"
	// Inbound Real-Time Payments Transfer Decline: details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// International ACH Decline: details will be under the `international_ach_decline`
	// object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInternationalACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "international_ach_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryWireDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "wire_decline"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryOther InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "other"
)

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// A computer-readable number printed on the MICR line of business checks, usually
	// the check number. This is useful for positive pay checks, but can be unreliably
	// transmitted by the bank of first deposit.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Field
	AuxiliaryOnUs apijson.Field
	Reason        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the check was declined.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason string

const (
	// The account number is disabled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The account number is canceled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The transaction would cause a limit to be exceeded.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonBreachesLimit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The account's entity is not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Stop payment requested for this check.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// The check was a duplicate deposit.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The check was not authorized.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonNotAuthorized InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The amount the receiving bank is attempting to deposit does not match the amount
	// on the check.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonAmountMismatch InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "amount_mismatch"
	// The check attempting to be deposited does not belong to Increase.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonNotOurItem InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "not_our_item"
)

// An Inbound Real-Time Payments Transfer Decline object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real-Time Payments
	// transfer.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Why the transfer was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	Reason                    apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real-Time Payments
// transfer.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

// Why the transfer was declined.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	// The account number is canceled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	// The account number is disabled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	// Your account is restricted.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	// Your account is inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	// The account's entity is not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	// Your account is not enabled to receive Real-Time Payments transfers.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

// An International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the destination country.
	DestinationCountryCode string `json:"destination_country_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// destination bank account.
	DestinationCurrencyCode string `json:"destination_currency_code,required"`
	// A description of how the foreign exchange rate was calculated.
	ForeignExchangeIndicator InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// originating bank account.
	OriginatingCurrencyCode string `json:"originating_currency_code,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the originating branch country.
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	// An identifier for the originating bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	OriginatingDepositoryFinancialInstitutionID string `json:"originating_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `originating_depository_financial_institution_id` field for this Transaction.
	OriginatingDepositoryFinancialInstitutionIDQualifier InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
	// The name of the originating bank. Sometimes this will refer to an American bank
	// and obscure the correspondent foreign bank.
	OriginatingDepositoryFinancialInstitutionName string `json:"originating_depository_financial_institution_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorCity string `json:"originator_city,required"`
	// A description field set by the originator.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// A portion of the originator address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the originator country.
	OriginatorCountry string `json:"originator_country,required"`
	// An identifier for the originating company. This is generally stable across
	// multiple ACH transfers.
	OriginatorIdentification string `json:"originator_identification,required"`
	// Either the name of the originator or an intermediary money transmitter.
	OriginatorName string `json:"originator_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorPostalCode string `json:"originator_postal_code,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStateOrProvince string `json:"originator_state_or_province,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStreetAddress string `json:"originator_street_address,required"`
	// A description field set by the originator.
	PaymentRelatedInformation string `json:"payment_related_information,required,nullable"`
	// A description field set by the originator.
	PaymentRelatedInformation2 string `json:"payment_related_information2,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverCity string `json:"receiver_city,required"`
	// A portion of the receiver address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the receiver country.
	ReceiverCountry string `json:"receiver_country,required"`
	// An identification number the originator uses for the receiver.
	ReceiverIdentificationNumber string `json:"receiver_identification_number,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverPostalCode string `json:"receiver_postal_code,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStateOrProvince string `json:"receiver_state_or_province,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStreetAddress string `json:"receiver_street_address,required"`
	// The name of the receiver of the transfer. This is not verified by Increase.
	ReceivingCompanyOrIndividualName string `json:"receiving_company_or_individual_name,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the receiving bank country.
	ReceivingDepositoryFinancialInstitutionCountry string `json:"receiving_depository_financial_institution_country,required"`
	// An identifier for the receiving bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	ReceivingDepositoryFinancialInstitutionID string `json:"receiving_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `receiving_depository_financial_institution_id` field for this Transaction.
	ReceivingDepositoryFinancialInstitutionIDQualifier InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach#returns).
	TraceNumber string `json:"trace_number,required"`
	JSON        inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineJSON struct {
	Amount                                                 apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReference                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	OriginatorStreetAddress                                apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A description of how the foreign exchange rate was calculated.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToVariable InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorVariableToFixed InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToFixed InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "fixed_to_fixed"
)

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeRate InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorBlank InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "blank"
)

// The type of transfer. Set by the originator.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAnnuity InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBusinessOrCommercial InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeDeposit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeLoan InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMiscellaneous InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMortgage InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePension InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRemittance InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRentOrLease InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSalaryOrPayroll InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTax InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAccountsReceivable InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBackOfficeConversion InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMachineTransfer InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfPurchase InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfSale InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRepresentedCheck InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSharedNetworkTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTelphoneInitiated InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeInternetInitiated InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "internet_initiated"
)

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierBicCode InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierIban InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierBicCode InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierIban InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// A name set by the sender.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	BeneficiaryReference string `json:"beneficiary_reference,required,nullable"`
	// An Increase-constructed description of the declined transaction.
	Description string `json:"description,required"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the originating bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator of the wire, set by the sending bank.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	// Why the wire transfer was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	JSON   inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineJSON struct {
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorRoutingNumber                 apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	Reason                                  apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the wire transfer was declined.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason string

const (
	// The account number is canceled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason = "account_number_canceled"
	// The account number is disabled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason = "account_number_disabled"
	// The account's entity is not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason = "entity_not_active"
	// Your account is inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason = "group_locked"
	// The beneficiary account number does not exist.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReasonNoAccountNumber InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason = "no_account_number"
	// The transaction is not allowed per Increase's terms.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDeclineReason = "transaction_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionTypeDeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType = "declined_transaction"
)

// If the Real-Time Payments Transfer attempt succeeds, this will contain the
// resulting [Transaction](#transactions) object. The Transaction's `source` will
// be of `category: inbound_real_time_payments_transfer_confirmation`.
type InboundRealTimePaymentsTransferSimulationResultTransaction struct {
	// The Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transaction's
	// Account.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionCurrency `json:"currency,required"`
	// An informational message describing this transaction. Use the fields in `source`
	// to get more detailed information. This field appears as the line-item on the
	// statement.
	Description string `json:"description,required"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Transaction came through.
	RouteType InboundRealTimePaymentsTransferSimulationResultTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source InboundRealTimePaymentsTransferSimulationResultTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultTransactionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionJSON contains the JSON
// metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransaction]
type inboundRealTimePaymentsTransferSimulationResultTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transaction's
// Account.
type InboundRealTimePaymentsTransferSimulationResultTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "USD"
)

// The type of the route this Transaction came through.
type InboundRealTimePaymentsTransferSimulationResultTransactionRouteType string

const (
	// An Account Number.
	InboundRealTimePaymentsTransferSimulationResultTransactionRouteTypeAccountNumber InboundRealTimePaymentsTransferSimulationResultTransactionRouteType = "account_number"
	// A Card.
	InboundRealTimePaymentsTransferSimulationResultTransactionRouteTypeCard InboundRealTimePaymentsTransferSimulationResultTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
type InboundRealTimePaymentsTransferSimulationResultTransactionSource struct {
	// An Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// An ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// An ACH Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// An ACH Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`.
	CheckTransferDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDeposit `json:"check_transfer_deposit,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// An Inbound ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// An Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// An Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// An Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// An Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// An Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// An Internal Source object. This field will be present in the JSON response if
	// and only if `category` is equal to `internal_source`.
	InternalSource InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Real-Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  inboundRealTimePaymentsTransferSimulationResultTransactionSourceJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceJSON contains
// the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSource]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceJSON struct {
	AccountTransferIntention                    apijson.Field
	ACHTransferIntention                        apijson.Field
	ACHTransferRejection                        apijson.Field
	ACHTransferReturn                           apijson.Field
	CardDisputeAcceptance                       apijson.Field
	CardRefund                                  apijson.Field
	CardRevenuePayment                          apijson.Field
	CardSettlement                              apijson.Field
	Category                                    apijson.Field
	CheckDepositAcceptance                      apijson.Field
	CheckDepositReturn                          apijson.Field
	CheckTransferDeposit                        apijson.Field
	CheckTransferIntention                      apijson.Field
	CheckTransferStopPaymentRequest             apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
	InboundCheck                                apijson.Field
	InboundInternationalACHTransfer             apijson.Field
	InboundRealTimePaymentsTransferConfirmation apijson.Field
	InboundWireDrawdownPayment                  apijson.Field
	InboundWireDrawdownPaymentReversal          apijson.Field
	InboundWireReversal                         apijson.Field
	InboundWireTransfer                         apijson.Field
	InterestPayment                             apijson.Field
	InternalSource                              apijson.Field
	RealTimePaymentsTransferAcknowledgement     apijson.Field
	SampleFunds                                 apijson.Field
	WireTransferIntention                       apijson.Field
	WireTransferRejection                       apijson.Field
	raw                                         string
	ExtraFields                                 map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	Description          apijson.Field
	DestinationAccountID apijson.Field
	SourceAccountID      apijson.Field
	TransferID           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "USD"
)

// An ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_intention`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention struct {
	// The account number for the destination account.
	AccountNumber string `json:"account_number,required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number,required"`
	// A description set when the ACH Transfer was created.
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntentionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntentionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntentionJSON struct {
	AccountNumber       apijson.Field
	Amount              apijson.Field
	RoutingNumber       apijson.Field
	StatementDescriptor apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_rejection`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejectionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejectionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_return`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned. This reason code is sent by the receiving
	// bank back to Increase.
	ReturnReasonCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Transaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnJSON struct {
	CreatedAt           apijson.Field
	RawReturnReasonCode apijson.Field
	ReturnReasonCode    apijson.Field
	TransactionID       apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH Transfer was returned. This reason code is sent by the receiving
// bank back to Increase.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the receiving account. Sometimes abbreviated to
	// NSF.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoAccount InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed at the receiving bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account at the receiving bank was frozen per the Office of Foreign
	// Assets Control.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer at the receiving bank reversed the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer at the receiving bank reversed the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer at the receiving institution informed their bank that
	// they have revoked authorization for a previously authorized transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution asked for this transfer to be
	// returned. The receiving bank is complying with the request.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation in the ACH network.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAddendaError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFieldError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating financial institution made a
	// mistake and this return corrects it.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH return's trace number does not match an
	// originated ACH.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "untimely_return"
)

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptanceJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptanceJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	CardDisputeID apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	PurchaseDetails      apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "USD"
)

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetails]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsJSON struct {
	CarRental                   apijson.Field
	CustomerReferenceIdentifier apijson.Field
	LocalTaxAmount              apijson.Field
	LocalTaxCurrency            apijson.Field
	Lodging                     apijson.Field
	NationalTaxAmount           apijson.Field
	NationalTaxCurrency         apijson.Field
	PurchaseIdentifier          apijson.Field
	PurchaseIdentifierFormat    apijson.Field
	Travel                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to car rentals.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRental struct {
	// Code indicating the vehicle's class.
	CarClassCode string `json:"car_class_code,required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date,required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency,required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented,required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency,required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency,required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency,required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name,required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRental]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalJSON struct {
	CarClassCode                 apijson.Field
	CheckoutDate                 apijson.Field
	DailyRentalRateAmount        apijson.Field
	DailyRentalRateCurrency      apijson.Field
	DaysRented                   apijson.Field
	ExtraCharges                 apijson.Field
	FuelChargesAmount            apijson.Field
	FuelChargesCurrency          apijson.Field
	InsuranceChargesAmount       apijson.Field
	InsuranceChargesCurrency     apijson.Field
	NoShowIndicator              apijson.Field
	OneWayDropOffChargesAmount   apijson.Field
	OneWayDropOffChargesCurrency apijson.Field
	RenterName                   apijson.Field
	WeeklyRentalRateAmount       apijson.Field
	WeeklyRentalRateCurrency     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (gas, late fee, etc.) being billed.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesGas InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

// Fields specific to lodging.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency,required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency,required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency,required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights,required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency,required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string `json:"total_tax_currency,required,nullable"`
	JSON             inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodging]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingJSON struct {
	CheckInDate                 apijson.Field
	DailyRoomRateAmount         apijson.Field
	DailyRoomRateCurrency       apijson.Field
	ExtraCharges                apijson.Field
	FolioCashAdvancesAmount     apijson.Field
	FolioCashAdvancesCurrency   apijson.Field
	FoodBeverageChargesAmount   apijson.Field
	FoodBeverageChargesCurrency apijson.Field
	NoShowIndicator             apijson.Field
	PrepaidExpensesAmount       apijson.Field
	PrepaidExpensesCurrency     apijson.Field
	RoomNights                  apijson.Field
	TotalRoomTaxAmount          apijson.Field
	TotalRoomTaxCurrency        apijson.Field
	TotalTaxAmount              apijson.Field
	TotalTaxCurrency            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (phone, late check-out, etc.) being billed.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesRestaurant InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesGiftShop InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesMiniBar InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesTelephone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesLaundry InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
)

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

// The format of the purchase identifier.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

// Fields specific to travel.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravel]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelJSON struct {
	Ancillary                     apijson.Field
	ComputerizedReservationSystem apijson.Field
	CreditReasonIndicator         apijson.Field
	DepartureDate                 apijson.Field
	OriginationCityAirportCode    apijson.Field
	PassengerName                 apijson.Field
	RestrictedTicketIndicator     apijson.Field
	TicketChangeIndicator         apijson.Field
	TicketNumber                  apijson.Field
	TravelAgencyCode              apijson.Field
	TravelAgencyName              apijson.Field
	TripLegs                      apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Ancillary purchases in addition to the airfare.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string `json:"ticket_document_number,required,nullable"`
	JSON                 inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillary]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the reason for a credit to the cardholder.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string `json:"sub_category,required,nullable"`
	JSON        inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the ancillary service.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

// Indicates the reason for a credit to the cardholder.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

// Indicates whether this ticket is non-refundable.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

// Indicates why a ticket was changed.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLeg struct {
	// Carrier code (e.g., United Airlines, Jet Blue, etc.).
	CarrierCode string `json:"carrier_code,required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code,required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code,required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number,required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class,required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLeg]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether a stopover is allowed on this ticket.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundTypeCardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType = "card_refund"
)

// A Card Revenue Payment object. This field will be present in the JSON response
// if and only if `category` is equal to `card_revenue_payment`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	PeriodEnd             apijson.Field
	PeriodStart           apijson.Field
	TransactedOnAccountID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "USD"
)

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement struct {
	// The Card Settlement identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The Card Authorization that was created prior to this Card Settlement, if one
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardAuthorization    apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "USD"
)

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetails]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsJSON struct {
	CarRental                   apijson.Field
	CustomerReferenceIdentifier apijson.Field
	LocalTaxAmount              apijson.Field
	LocalTaxCurrency            apijson.Field
	Lodging                     apijson.Field
	NationalTaxAmount           apijson.Field
	NationalTaxCurrency         apijson.Field
	PurchaseIdentifier          apijson.Field
	PurchaseIdentifierFormat    apijson.Field
	Travel                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to car rentals.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRental struct {
	// Code indicating the vehicle's class.
	CarClassCode string `json:"car_class_code,required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date,required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency,required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented,required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency,required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency,required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency,required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name,required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRental]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON struct {
	CarClassCode                 apijson.Field
	CheckoutDate                 apijson.Field
	DailyRentalRateAmount        apijson.Field
	DailyRentalRateCurrency      apijson.Field
	DaysRented                   apijson.Field
	ExtraCharges                 apijson.Field
	FuelChargesAmount            apijson.Field
	FuelChargesCurrency          apijson.Field
	InsuranceChargesAmount       apijson.Field
	InsuranceChargesCurrency     apijson.Field
	NoShowIndicator              apijson.Field
	OneWayDropOffChargesAmount   apijson.Field
	OneWayDropOffChargesCurrency apijson.Field
	RenterName                   apijson.Field
	WeeklyRentalRateAmount       apijson.Field
	WeeklyRentalRateCurrency     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (gas, late fee, etc.) being billed.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesGas InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

// Fields specific to lodging.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency,required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency,required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency,required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights,required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency,required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string `json:"total_tax_currency,required,nullable"`
	JSON             inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodging]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingJSON struct {
	CheckInDate                 apijson.Field
	DailyRoomRateAmount         apijson.Field
	DailyRoomRateCurrency       apijson.Field
	ExtraCharges                apijson.Field
	FolioCashAdvancesAmount     apijson.Field
	FolioCashAdvancesCurrency   apijson.Field
	FoodBeverageChargesAmount   apijson.Field
	FoodBeverageChargesCurrency apijson.Field
	NoShowIndicator             apijson.Field
	PrepaidExpensesAmount       apijson.Field
	PrepaidExpensesCurrency     apijson.Field
	RoomNights                  apijson.Field
	TotalRoomTaxAmount          apijson.Field
	TotalRoomTaxCurrency        apijson.Field
	TotalTaxAmount              apijson.Field
	TotalTaxCurrency            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (phone, late check-out, etc.) being billed.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesTelephone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesLaundry InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
)

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

// The format of the purchase identifier.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

// Fields specific to travel.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravel]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelJSON struct {
	Ancillary                     apijson.Field
	ComputerizedReservationSystem apijson.Field
	CreditReasonIndicator         apijson.Field
	DepartureDate                 apijson.Field
	OriginationCityAirportCode    apijson.Field
	PassengerName                 apijson.Field
	RestrictedTicketIndicator     apijson.Field
	TicketChangeIndicator         apijson.Field
	TicketNumber                  apijson.Field
	TravelAgencyCode              apijson.Field
	TravelAgencyName              apijson.Field
	TripLegs                      apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Ancillary purchases in addition to the airfare.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string `json:"ticket_document_number,required,nullable"`
	JSON                 inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillary]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the reason for a credit to the cardholder.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string `json:"sub_category,required,nullable"`
	JSON        inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the ancillary service.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

// Indicates the reason for a credit to the cardholder.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

// Indicates whether this ticket is non-refundable.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

// Indicates why a ticket was changed.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg struct {
	// Carrier code (e.g., United Airlines, Jet Blue, etc.).
	CarrierCode string `json:"carrier_code,required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code,required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code,required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number,required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class,required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether a stopover is allowed on this ticket.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementTypeCardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType = "card_settlement"
)

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory string

const (
	// Account Transfer Intention: details will be under the
	// `account_transfer_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryAccountTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "account_transfer_intention"
	// ACH Transfer Intention: details will be under the `ach_transfer_intention`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_intention"
	// ACH Transfer Rejection: details will be under the `ach_transfer_rejection`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_rejection"
	// ACH Transfer Return: details will be under the `ach_transfer_return` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_return"
	// Card Dispute Acceptance: details will be under the `card_dispute_acceptance`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardDisputeAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_dispute_acceptance"
	// Card Refund: details will be under the `card_refund` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_refund"
	// Card Settlement: details will be under the `card_settlement` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_settlement"
	// Card Revenue Payment: details will be under the `card_revenue_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRevenuePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_revenue_payment"
	// Check Deposit Acceptance: details will be under the `check_deposit_acceptance`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	// Check Deposit Return: details will be under the `check_deposit_return` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_return"
	// Check Transfer Deposit: details will be under the `check_transfer_deposit`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_deposit"
	// Check Transfer Intention: details will be under the `check_transfer_intention`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_intention"
	// Check Transfer Stop Payment Request: details will be under the
	// `check_transfer_stop_payment_request` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	// Fee Payment: details will be under the `fee_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryFeePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "fee_payment"
	// Inbound ACH Transfer Intention: details will be under the `inbound_ach_transfer`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer"
	// Inbound ACH Transfer Return Intention: details will be under the
	// `inbound_ach_transfer_return_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransferReturnIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// Inbound Check: details will be under the `inbound_check` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_check"
	// Inbound International ACH Transfer: details will be under the
	// `inbound_international_ach_transfer` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundInternationalACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_international_ach_transfer"
	// Inbound Real-Time Payments Transfer Confirmation: details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// Inbound Wire Drawdown Payment Reversal: details will be under the
	// `inbound_wire_drawdown_payment_reversal` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPaymentReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// Inbound Wire Drawdown Payment: details will be under the
	// `inbound_wire_drawdown_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment"
	// Inbound Wire Reversal: details will be under the `inbound_wire_reversal` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_reversal"
	// Inbound Wire Transfer: details will be under the `inbound_wire_transfer` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_transfer"
	// Interest Payment: details will be under the `interest_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInterestPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "interest_payment"
	// Internal Source: details will be under the `internal_source` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInternalSource InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "internal_source"
	// Real-Time Payments Transfer Acknowledgement: details will be under the
	// `real_time_payments_transfer_acknowledgement` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// Sample Funds: details will be under the `sample_funds` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategorySampleFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "sample_funds"
	// Wire Transfer Intention: details will be under the `wire_transfer_intention`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_intention"
	// Wire Transfer Rejection: details will be under the `wire_transfer_rejection`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_rejection"
	// The Transaction was made for an undocumented or deprecated reason.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryOther InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "other"
)

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance struct {
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string `json:"serial_number,required,nullable"`
	JSON         inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceJSON struct {
	AccountNumber  apijson.Field
	Amount         apijson.Field
	AuxiliaryOnUs  apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	RoutingNumber  apijson.Field
	SerialNumber   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// Why this check was returned by the bank holding the account it was drawn
	// against.
	ReturnReason InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string `json:"transaction_id,required"`
	JSON          inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	ReturnReason   apijson.Field
	ReturnedAt     apijson.Field
	TransactionID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "USD"
)

// Why this check was returned by the bank holding the account it was drawn
// against.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason string

const (
	// The check doesn't allow ACH conversion.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	// The account is closed.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonClosedAccount InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	// The check has already been deposited.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	// Insufficient funds
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	// No account was found matching the check details.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonNoAccount InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "no_account"
	// The check was not authorized.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonNotAuthorized InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	// The check is too old.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonStaleDated InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	// The payment has been stopped by the account holder.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonStopPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	// The reason for the return is unknown.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnknownReason InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	// The image doesn't match the details submitted.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	// The image could not be read.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnreadableImage InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
	// The check endorsement was irregular.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "endorsement_irregular"
)

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDeposit struct {
	// The identifier of the API File object containing an image of the back of the
	// deposited check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number,required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The identifier of the Check Transfer object that was deposited.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDeposit]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositJSON struct {
	BackImageFileID                 apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	DepositedAt                     apijson.Field
	FrontImageFileID                apijson.Field
	TransactionID                   apijson.Field
	TransferID                      apijson.Field
	Type                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositTypeCheckTransferDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositType = "check_transfer_deposit"
)

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention struct {
	// The city of the check's destination.
	AddressCity string `json:"address_city,required,nullable"`
	// The street address of the check's destination.
	AddressLine1 string `json:"address_line1,required,nullable"`
	// The second line of the address of the check's destination.
	AddressLine2 string `json:"address_line2,required,nullable"`
	// The state of the check's destination.
	AddressState string `json:"address_state,required,nullable"`
	// The postal code of the check's destination.
	AddressZip string `json:"address_zip,required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required,nullable"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionJSON struct {
	AddressCity   apijson.Field
	AddressLine1  apijson.Field
	AddressLine2  apijson.Field
	AddressState  apijson.Field
	AddressZip    apijson.Field
	Amount        apijson.Field
	Currency      apijson.Field
	RecipientName apijson.Field
	TransferID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "USD"
)

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The reason why this transfer was stopped.
	Reason InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why this transfer was stopped.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReasonMailDeliveryFailed InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was canceled by an Increase operator who will provide details
	// out-of-band.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReasonRejectedByIncrease InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason = "rejected_by_increase"
	// The check was stopped for another reason.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReasonUnknown InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason = "unknown"
)

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency `json:"currency,required"`
	// The start of this payment's fee period, usually the first day of a month.
	FeePeriodStart time.Time `json:"fee_period_start,required" format:"date"`
	JSON           inboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON struct {
	Amount         apijson.Field
	Currency       apijson.Field
	FeePeriodStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "USD"
)

// An Inbound ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_ach_transfer`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The description of the date of the transfer, usually in the format `YYMMDD`.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date,required,nullable"`
	// Data set by the originator.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	// An informational description of the transfer.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// An identifier for the originating company. This is generally, but not always, a
	// stable identifier across multiple transfers.
	OriginatorCompanyID string `json:"originator_company_id,required"`
	// A name set by the originator to identify themselves.
	OriginatorCompanyName string `json:"originator_company_name,required"`
	// The originator's identifier for the transfer receipient.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the transfer recipient. This value is informational and not verified
	// by Increase.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach#returns).
	TraceNumber string `json:"trace_number,required"`
	// The inbound ach transfer's identifier.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransferJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransferJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransferJSON struct {
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	TransferID                         apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number,required,nullable"`
	// The front image of the check. This is a black and white TIFF image file.
	CheckFrontImageFileID string `json:"check_front_image_file_id,required,nullable"`
	// The number of the check. This field is set by the depositing bank and can be
	// unreliable.
	CheckNumber string `json:"check_number,required,nullable"`
	// The rear image of the check. This is a black and white TIFF image file.
	CheckRearImageFileID string `json:"check_rear_image_file_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency `json:"currency,required"`
	JSON     inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON struct {
	Amount                          apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	CheckFrontImageFileID           apijson.Field
	CheckNumber                     apijson.Field
	CheckRearImageFileID            apijson.Field
	Currency                        apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "USD"
)

// An Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the destination country.
	DestinationCountryCode string `json:"destination_country_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// destination bank account.
	DestinationCurrencyCode string `json:"destination_currency_code,required"`
	// A description of how the foreign exchange rate was calculated.
	ForeignExchangeIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// originating bank account.
	OriginatingCurrencyCode string `json:"originating_currency_code,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the originating branch country.
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	// An identifier for the originating bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	OriginatingDepositoryFinancialInstitutionID string `json:"originating_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `originating_depository_financial_institution_id` field for this Transaction.
	OriginatingDepositoryFinancialInstitutionIDQualifier InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
	// The name of the originating bank. Sometimes this will refer to an American bank
	// and obscure the correspondent foreign bank.
	OriginatingDepositoryFinancialInstitutionName string `json:"originating_depository_financial_institution_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorCity string `json:"originator_city,required"`
	// A description field set by the originator.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// A portion of the originator address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the originator country.
	OriginatorCountry string `json:"originator_country,required"`
	// An identifier for the originating company. This is generally stable across
	// multiple ACH transfers.
	OriginatorIdentification string `json:"originator_identification,required"`
	// Either the name of the originator or an intermediary money transmitter.
	OriginatorName string `json:"originator_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorPostalCode string `json:"originator_postal_code,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStateOrProvince string `json:"originator_state_or_province,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStreetAddress string `json:"originator_street_address,required"`
	// A description field set by the originator.
	PaymentRelatedInformation string `json:"payment_related_information,required,nullable"`
	// A description field set by the originator.
	PaymentRelatedInformation2 string `json:"payment_related_information2,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverCity string `json:"receiver_city,required"`
	// A portion of the receiver address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the receiver country.
	ReceiverCountry string `json:"receiver_country,required"`
	// An identification number the originator uses for the receiver.
	ReceiverIdentificationNumber string `json:"receiver_identification_number,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverPostalCode string `json:"receiver_postal_code,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStateOrProvince string `json:"receiver_state_or_province,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStreetAddress string `json:"receiver_street_address,required"`
	// The name of the receiver of the transfer. This is not verified by Increase.
	ReceivingCompanyOrIndividualName string `json:"receiving_company_or_individual_name,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the receiving bank country.
	ReceivingDepositoryFinancialInstitutionCountry string `json:"receiving_depository_financial_institution_country,required"`
	// An identifier for the receiving bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	ReceivingDepositoryFinancialInstitutionID string `json:"receiving_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `receiving_depository_financial_institution_id` field for this Transaction.
	ReceivingDepositoryFinancialInstitutionIDQualifier InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach#returns).
	TraceNumber string `json:"trace_number,required"`
	JSON        inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferJSON struct {
	Amount                                                 apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReference                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	OriginatorStreetAddress                                apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A description of how the foreign exchange rate was calculated.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToVariable InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorVariableToFixed InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToFixed InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "fixed_to_fixed"
)

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeRate InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorBlank InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "blank"
)

// The type of transfer. Set by the originator.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAnnuity InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBusinessOrCommercial InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeLoan InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMiscellaneous InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMortgage InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePension InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRemittance InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRentOrLease InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSalaryOrPayroll InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTax InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAccountsReceivable InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBackOfficeConversion InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMachineTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfPurchase InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfSale InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRepresentedCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSharedNetworkTransaction InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTelphoneInitiated InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeInternetInitiated InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "internet_initiated"
)

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierBicCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierIban InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierBicCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierIban InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real-Time Payments transfer.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// A name set by the sender.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	BeneficiaryReference string `json:"beneficiary_reference,required,nullable"`
	// An Increase-constructed description of the transfer.
	Description string `json:"description,required"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the receiving bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator of the wire, set by the sending bank.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// An Increase-created concatenation of the Originator-to-Beneficiary lines.
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON struct {
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorRoutingNumber                 apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	JSON                       inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
	Amount                                        apijson.Field
	Description                                   apijson.Field
	InputCycleDate                                apijson.Field
	InputMessageAccountabilityData                apijson.Field
	InputSequenceNumber                           apijson.Field
	InputSource                                   apijson.Field
	OriginatorRoutingNumber                       apijson.Field
	PreviousMessageInputCycleDate                 apijson.Field
	PreviousMessageInputMessageAccountabilityData apijson.Field
	PreviousMessageInputSequenceNumber            apijson.Field
	PreviousMessageInputSource                    apijson.Field
	raw                                           string
	ExtraFields                                   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal struct {
	// The amount that was reversed in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire, set by the reversing bank.
	Description string `json:"description,required"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The Fedwire cycle date for the wire reversal. The "Fedwire day" begins at 9:00
	// PM Eastern Time on the evening before the `cycle date`.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// The Fedwire cycle date for the wire transfer that is being reversed by this
	// message.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string `json:"wire_transfer_id,required"`
	JSON           inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversalJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversalJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversalJSON struct {
	Amount                                                apijson.Field
	CreatedAt                                             apijson.Field
	Description                                           apijson.Field
	FinancialInstitutionToFinancialInstitutionInformation apijson.Field
	InputCycleDate                                        apijson.Field
	InputMessageAccountabilityData                        apijson.Field
	InputSequenceNumber                                   apijson.Field
	InputSource                                           apijson.Field
	OriginatorRoutingNumber                               apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	TransactionID                                         apijson.Field
	WireTransferID                                        apijson.Field
	raw                                                   string
	ExtraFields                                           map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer struct {
	// The amount in USD cents.
	Amount int64 `json:"amount,required"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// A name set by the sender.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	BeneficiaryReference string `json:"beneficiary_reference,required,nullable"`
	// An Increase-constructed description of the transfer.
	Description string `json:"description,required"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the originating bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator of the wire, set by the sending bank.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// An Increase-created concatenation of the Originator-to-Beneficiary lines.
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransferJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransferJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransferJSON struct {
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorRoutingNumber                 apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment struct {
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	JSON        inboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentJSON struct {
	AccruedOnAccountID apijson.Field
	Amount             apijson.Field
	Currency           apijson.Field
	PeriodEnd          apijson.Field
	PeriodStart        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "USD"
)

// An Internal Source object. This field will be present in the JSON response if
// and only if `category` is equal to `internal_source`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency `json:"currency,required"`
	// An Internal Source is a transaction between you and Increase. This describes the
	// reason for the transaction.
	Reason InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason `json:"reason,required"`
	JSON   inboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency string

const (
	// Canadian Dollar (CAD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "CAD"
	// Swiss Franc (CHF)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "CHF"
	// Euro (EUR)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "EUR"
	// British Pound (GBP)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "GBP"
	// Japanese Yen (JPY)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "JPY"
	// US Dollar (USD)
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "USD"
)

// An Internal Source is a transaction between you and Increase. This describes the
// reason for the transaction.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason string

const (
	// Account closure
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonAccountClosure InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonBankMigration InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonCashback InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "cashback"
	// Check adjustment
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonCheckAdjustment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "check_adjustment"
	// Collection receivable
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonCollectionReceivable InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "collection_receivable"
	// Empyreal adjustment
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonEmpyrealAdjustment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "empyreal_adjustment"
	// Error
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonError InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "error"
	// Error correction
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonErrorCorrection InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "error_correction"
	// Fees
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonFees InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "fees"
	// Interest
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonInterest InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "interest"
	// Negative balance forgiveness
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonNegativeBalanceForgiveness InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	// Sample funds
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonSampleFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "sample_funds"
	// Sample funds return
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonSampleFundsReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "sample_funds_return"
)

// A Real-Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	RemittanceInformation    apijson.Field
	TransferID               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFundsJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFundsJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFundsJSON struct {
	Originator  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention struct {
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntentionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntentionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntentionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection struct {
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejectionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejectionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
type InboundRealTimePaymentsTransferSimulationResultTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionTypeTransaction InboundRealTimePaymentsTransferSimulationResultTransactionType = "transaction"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_transfer_simulation_result`.
type InboundRealTimePaymentsTransferSimulationResultType string

const (
	InboundRealTimePaymentsTransferSimulationResultTypeInboundRealTimePaymentsTransferSimulationResult InboundRealTimePaymentsTransferSimulationResultType = "inbound_real_time_payments_transfer_simulation_result"
)

type SimulationRealTimePaymentsTransferCompleteParams struct {
	// If set, the simulation will reject the transfer.
	Rejection param.Field[SimulationRealTimePaymentsTransferCompleteParamsRejection] `json:"rejection"`
}

func (r SimulationRealTimePaymentsTransferCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the simulation will reject the transfer.
type SimulationRealTimePaymentsTransferCompleteParamsRejection struct {
	// The reason code that the simulated rejection will have.
	RejectReasonCode param.Field[SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode] `json:"reject_reason_code,required"`
}

func (r SimulationRealTimePaymentsTransferCompleteParamsRejection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason code that the simulated rejection will have.
type SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode string

const (
	// The destination account is closed. Corresponds to the Real-Time Payments reason
	// code `AC04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_closed"
	// The destination account is currently blocked from receiving transactions.
	// Corresponds to the Real-Time Payments reason code `AC06`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_blocked"
	// The destination account is ineligible to receive Real-Time Payments transfers.
	// Corresponds to the Real-Time Payments reason code `AC14`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_type"
	// The destination account does not exist. Corresponds to the Real-Time Payments
	// reason code `AC03`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_number"
	// The destination routing number is invalid. Corresponds to the Real-Time Payments
	// reason code `RC04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	// The destination account holder is deceased. Corresponds to the Real-Time
	// Payments reason code `MD07`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "end_customer_deceased"
	// The reason is provided as narrative information in the additional information
	// field.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeNarrative SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "narrative"
	// Real-Time Payments transfers are not allowed to the destination account.
	// Corresponds to the Real-Time Payments reason code `AG01`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_forbidden"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `AG03`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_type_not_supported"
	// The amount of the transfer is different than expected by the recipient.
	// Corresponds to the Real-Time Payments reason code `AM09`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnexpectedAmount SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unexpected_amount"
	// The amount is higher than the recipient is authorized to send or receive.
	// Corresponds to the Real-Time Payments reason code `AM14`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	// The creditor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_address"
	// The specified creditor is unknown. Corresponds to the Real-Time Payments reason
	// code `BE06`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnknownEndCustomer SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unknown_end_customer"
	// The debtor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE07`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_debtor_address"
	// There was a timeout processing the transfer. Corresponds to the Real-Time
	// Payments reason code `DS24`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTimeout SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "timeout"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `NOAT`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnsupportedMessageForRecipient SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unsupported_message_for_recipient"
	// The destination financial institution is currently not connected to Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9912`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRecipientConnectionNotAvailable SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "recipient_connection_not_available"
	// Real-Time Payments is currently unavailable. Corresponds to the Real-Time
	// Payments reason code `9948`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRealTimePaymentsSuspended SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "real_time_payments_suspended"
	// The destination financial institution is currently signed off of Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9910`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInstructedAgentSignedOff SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "instructed_agent_signed_off"
	// The transfer was rejected due to an internal Increase issue. We have been
	// notified.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeProcessingError SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "processing_error"
	// Some other error or issue has occurred.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeOther SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "other"
)

type SimulationRealTimePaymentsTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound Real-Time Payments Transfer is
	// for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in USD cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber param.Field[string] `json:"debtor_account_number"`
	// The name provided by the sender of the transfer.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber param.Field[string] `json:"debtor_routing_number"`
	// Additional information included with the transfer.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// The identifier of a pending Request for Payment that this transfer will fulfill.
	RequestForPaymentID param.Field[string] `json:"request_for_payment_id"`
}

func (r SimulationRealTimePaymentsTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
