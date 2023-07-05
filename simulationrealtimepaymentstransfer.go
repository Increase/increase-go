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
	"github.com/increase/increase-go/internal/shared"
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

// Simulates submission of a Real Time Payments transfer and handling the response
// from the destination financial institution. This transfer must first have a
// `status` of `pending_submission`.
func (r *SimulationRealTimePaymentsTransferService) Complete(ctx context.Context, realTimePaymentsTransferID string, body SimulationRealTimePaymentsTransferCompleteParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/real_time_payments_transfers/%s/complete", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an inbound Real Time Payments transfer to your account. Real Time
// Payments are a beta feature.
func (r *SimulationRealTimePaymentsTransferService) NewInbound(ctx context.Context, body SimulationRealTimePaymentsTransferNewInboundParams, opts ...option.RequestOption) (res *InboundRealTimePaymentsTransferSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of an inbound Real Time Payments Transfer simulation.
type InboundRealTimePaymentsTransferSimulationResult struct {
	// If the Real Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// If the Real Time Payments Transfer attempt succeeds, this will contain the
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

// If the Real Time Payments Transfer attempt fails, this will contain the
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
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
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
// Transcation's Account.
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
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory `json:"category,required"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
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

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	// Why the ACH transfer was declined.
	Reason           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                                                                   `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                                                                   `json:"receiver_name,required,nullable"`
	TraceNumber      string                                                                                   `json:"trace_number,required"`
	JSON             inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON struct {
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	Reason                             apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
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
	// The transaction would cause a limit to be exceeded.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonBreachesLimit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// Other.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonDuplicateReturn InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// Your account is inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// Other.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonMisroutedReturn InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// The account number that was debited does not exist.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonNoACHRoute InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// Other.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonOriginatorRequest InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
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
	// The payment network used to process this card authorization
	Network InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	JSON   inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantState        apijson.Field
	Network              apijson.Field
	NetworkDetails       apijson.Field
	RealTimeDecisionID   apijson.Field
	Reason               apijson.Field
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

// The payment network used to process this card authorization
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork string

const (
	// Visa
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkVisa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

// Fields specific to the `network`
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

// inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails]
type inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to the `visa` network
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
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

// Why the transaction was declined.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCardNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "card_not_active"
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
	// The transaction was blocked by an internal limit for new Increase accounts.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonBreachesInternalLimit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "breaches_internal_limit"
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
)

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory string

const (
	// The Declined Transaction was created by a ACH Decline object. Details will be
	// under the `ach_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "ach_decline"
	// The Declined Transaction was created by a Card Decline object. Details will be
	// under the `card_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCardDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "card_decline"
	// The Declined Transaction was created by a Check Decline object. Details will be
	// under the `check_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCheckDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "check_decline"
	// The Declined Transaction was created by a Inbound Real Time Payments Transfer
	// Decline object. Details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// The Declined Transaction was created by a International ACH Decline object.
	// Details will be under the `international_ach_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInternationalACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "international_ach_decline"
	// The Declined Transaction was created by a Wire Decline object. Details will be
	// under the `wire_decline` object.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryWireDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "wire_decline"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryOther InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "other"
)

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
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
	// The account number is canceled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonBreachesLimit InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The account's entity is not active.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonEntityNotActive InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonGroupLocked InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Unable to locate account.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	// Routing number on the check is not ours.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonNotOurItem InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	// Unable to process.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonUnableToProcess InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	// Refer to image.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonReferToImage InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	// Stop payment requested for this check.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// Check was returned to sender.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonReturned InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "returned"
	// The check was a duplicate deposit.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The transaction is not allowed.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonNotAuthorized InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The check was altered or fictitious.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
)

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
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
	// The Real Time Payments network identification of the declined transfer.
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
// transfer's currency. This will always be "USD" for a Real Time Payments
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
	// Your account is not enabled to receive Real Time Payments transfers.
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                                                 int64  `json:"amount,required"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   inboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineJSON
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

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceWireDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                                  int64  `json:"amount,required"`
	BeneficiaryAddressLine1                 string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    string `json:"beneficiary_reference,required,nullable"`
	Description                             string `json:"description,required"`
	InputMessageAccountabilityData          string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  string `json:"originator_address_line3,required,nullable"`
	OriginatorName                          string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
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

// If the Real Time Payments Transfer attempt succeeds, this will contain the
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
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionCurrency `json:"currency,required"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
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
// Transaction's currency. This will match the currency on the Transcation's
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
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
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
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
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
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
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
	CheckTransferRejection                      apijson.Field
	CheckTransferReturn                         apijson.Field
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

// A Account Transfer Intention object. This field will be present in the JSON
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

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention struct {
	AccountNumber string `json:"account_number,required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	RoutingNumber       string `json:"routing_number,required"`
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

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
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

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Tranasaction associated with this return.
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

// Why the ACH Transfer was returned.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the source account.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoAccount InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account was frozen per the Office of Foreign Assets Control.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer reversed the transfer.
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
	// Code R10. The customer reversed the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer who initiated the transfer revoked authorization.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution reversed the transfer.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation.
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
	// Code R62. A rare return reason. The originating bank made a mistake earlier and
	// this return corrects it.
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
	// Code R27. A rare return reason. An ACH Return's trace number does not match an
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
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
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
	// The Card Authorization that was created prior to this Card Settlement, if on
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
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

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementTypeCardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType = "card_settlement"
)

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory string

const (
	// The Transaction was created by a Account Transfer Intention object. Details will
	// be under the `account_transfer_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryAccountTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "account_transfer_intention"
	// The Transaction was created by a ACH Transfer Intention object. Details will be
	// under the `ach_transfer_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_intention"
	// The Transaction was created by a ACH Transfer Rejection object. Details will be
	// under the `ach_transfer_rejection` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_rejection"
	// The Transaction was created by a ACH Transfer Return object. Details will be
	// under the `ach_transfer_return` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_return"
	// The Transaction was created by a Card Dispute Acceptance object. Details will be
	// under the `card_dispute_acceptance` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardDisputeAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_dispute_acceptance"
	// The Transaction was created by a Card Refund object. Details will be under the
	// `card_refund` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_refund"
	// The Transaction was created by a Card Revenue Payment object. Details will be
	// under the `card_revenue_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRevenuePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_revenue_payment"
	// The Transaction was created by a Card Settlement object. Details will be under
	// the `card_settlement` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_settlement"
	// The Transaction was created by a Check Deposit Acceptance object. Details will
	// be under the `check_deposit_acceptance` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	// The Transaction was created by a Check Deposit Return object. Details will be
	// under the `check_deposit_return` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_return"
	// The Transaction was created by a Check Transfer Deposit object. Details will be
	// under the `check_transfer_deposit` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_deposit"
	// The Transaction was created by a Check Transfer Intention object. Details will
	// be under the `check_transfer_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_intention"
	// The Transaction was created by a Check Transfer Rejection object. Details will
	// be under the `check_transfer_rejection` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_rejection"
	// The Transaction was created by a Check Transfer Return object. Details will be
	// under the `check_transfer_return` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_return"
	// The Transaction was created by a Check Transfer Stop Payment Request object.
	// Details will be under the `check_transfer_stop_payment_request` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	// The Transaction was created by a Fee Payment object. Details will be under the
	// `fee_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryFeePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "fee_payment"
	// The Transaction was created by a Inbound ACH Transfer object. Details will be
	// under the `inbound_ach_transfer` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer"
	// The Transaction was created by a Inbound ACH Transfer Return Intention object.
	// Details will be under the `inbound_ach_transfer_return_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransferReturnIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// The Transaction was created by a Inbound Check object. Details will be under the
	// `inbound_check` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_check"
	// The Transaction was created by a Inbound International ACH Transfer object.
	// Details will be under the `inbound_international_ach_transfer` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundInternationalACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_international_ach_transfer"
	// The Transaction was created by a Inbound Real Time Payments Transfer
	// Confirmation object. Details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// The Transaction was created by a Inbound Wire Drawdown Payment object. Details
	// will be under the `inbound_wire_drawdown_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment"
	// The Transaction was created by a Inbound Wire Drawdown Payment Reversal object.
	// Details will be under the `inbound_wire_drawdown_payment_reversal` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPaymentReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// The Transaction was created by a Inbound Wire Reversal object. Details will be
	// under the `inbound_wire_reversal` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_reversal"
	// The Transaction was created by a Inbound Wire Transfer object. Details will be
	// under the `inbound_wire_transfer` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_transfer"
	// The Transaction was created by a Interest Payment object. Details will be under
	// the `interest_payment` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInterestPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "interest_payment"
	// The Transaction was created by a Internal Source object. Details will be under
	// the `internal_source` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInternalSource InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "internal_source"
	// The Transaction was created by a Real Time Payments Transfer Acknowledgement
	// object. Details will be under the `real_time_payments_transfer_acknowledgement`
	// object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// The Transaction was created by a Sample Funds object. Details will be under the
	// `sample_funds` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategorySampleFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "sample_funds"
	// The Transaction was created by a Wire Transfer Intention object. Details will be
	// under the `wire_transfer_intention` object.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_intention"
	// The Transaction was created by a Wire Transfer Rejection object. Details will be
	// under the `wire_transfer_rejection` object.
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
	Currency     InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency     `json:"currency,required"`
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
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDeposit]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferDepositJSON struct {
	BackImageFileID  apijson.Field
	DepositedAt      apijson.Field
	FrontImageFileID apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
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

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejectionJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejectionJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Check Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_return`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn struct {
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The identifier of the returned Check Transfer.
	TransferID string `json:"transfer_id,required"`
	JSON       inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnJSON struct {
	FileID        apijson.Field
	Reason        apijson.Field
	ReturnedAt    apijson.Field
	TransactionID apijson.Field
	TransferID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why the check was returned.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason string

const (
	// Mail delivery failed and the check was returned to sender.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonMailDeliveryFailure InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	// The check arrived and the recipient refused to deposit it.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonRefusedByRecipient InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
	// The check was fraudulently deposited and the transfer was returned to the Bank
	// of First Deposit.
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonReturnedNotAuthorized InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "returned_not_authorized"
)

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
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
	RequestedAt   apijson.Field
	TransactionID apijson.Field
	TransferID    apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	JSON     inboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyEntryDescription  string `json:"originator_company_entry_description,required"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	ReceiverIDNumber                   string `json:"receiver_id_number,required,nullable"`
	ReceiverName                       string `json:"receiver_name,required,nullable"`
	TraceNumber                        string `json:"trace_number,required"`
	JSON                               inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransferJSON
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
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                int64  `json:"amount,required"`
	CheckFrontImageFileID string `json:"check_front_image_file_id,required,nullable"`
	CheckNumber           string `json:"check_number,required,nullable"`
	CheckRearImageFileID  string `json:"check_rear_image_file_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency `json:"currency,required"`
	JSON     inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Field
	CheckFrontImageFileID apijson.Field
	CheckNumber           apijson.Field
	CheckRearImageFileID  apijson.Field
	Currency              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                                                 int64  `json:"amount,required"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferJSON
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

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real Time Payments network identification of the transfer
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
// currency. This will always be "USD" for a Real Time Payments transfer.
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

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                             int64  `json:"amount,required"`
	BeneficiaryAddressLine1            string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2            string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3            string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                    string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference               string `json:"beneficiary_reference,required,nullable"`
	Description                        string `json:"description,required"`
	InputMessageAccountabilityData     string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1             string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2             string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3             string `json:"originator_address_line3,required,nullable"`
	OriginatorName                     string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	JSON                               inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
}

// inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
// contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment]
type inboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON struct {
	Amount                             apijson.Field
	BeneficiaryAddressLine1            apijson.Field
	BeneficiaryAddressLine2            apijson.Field
	BeneficiaryAddressLine3            apijson.Field
	BeneficiaryName                    apijson.Field
	BeneficiaryReference               apijson.Field
	Description                        apijson.Field
	InputMessageAccountabilityData     apijson.Field
	OriginatorAddressLine1             apijson.Field
	OriginatorAddressLine2             apijson.Field
	OriginatorAddressLine3             apijson.Field
	OriginatorName                     apijson.Field
	OriginatorToBeneficiaryInformation apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
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

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
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
	TransactionID string `json:"transaction_id,required,nullable"`
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

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                                  int64  `json:"amount,required"`
	BeneficiaryAddressLine1                 string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    string `json:"beneficiary_reference,required,nullable"`
	Description                             string `json:"description,required"`
	InputMessageAccountabilityData          string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  string `json:"originator_address_line3,required,nullable"`
	OriginatorName                          string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformation      string `json:"originator_to_beneficiary_information,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
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

// A Interest Payment object. This field will be present in the JSON response if
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

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     inboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceJSON
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

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason string

const (
	// Account closure
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonAccountClosure InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonBankMigration InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonCashback InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "cashback"
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

// A Real Time Payments Transfer Acknowledgement object. This field will be present
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
	// The identifier of the Real Time Payments Transfer that led to this Transaction.
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
	TransferID    string `json:"transfer_id,required"`
	JSON          inboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntentionJSON
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
	// The destination account is closed. Corresponds to the Real Time Payments reason
	// code `AC04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_closed"
	// The destination account is currently blocked from receiving transactions.
	// Corresponds to the Real Time Payments reason code `AC06`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_blocked"
	// The destination account is ineligible to receive Real Time Payments transfers.
	// Corresponds to the Real Time Payments reason code `AC14`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_type"
	// The destination account does not exist. Corresponds to the Real Time Payments
	// reason code `AC03`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_number"
	// The destination routing number is invalid. Corresponds to the Real Time Payments
	// reason code `RC04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	// The destination account holder is deceased. Corresponds to the Real Time
	// Payments reason code `MD07`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "end_customer_deceased"
	// The reason is provided as narrative information in the additional information
	// field.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeNarrative SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "narrative"
	// Real Time Payments transfers are not allowed to the destination account.
	// Corresponds to the Real Time Payments reason code `AG01`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_forbidden"
	// Real Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real Time Payments reason code `AG03`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_type_not_supported"
	// The amount of the transfer is different than expected by the recipient.
	// Corresponds to the Real Time Payments reason code `AM09`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnexpectedAmount SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unexpected_amount"
	// The amount is higher than the recipient is authorized to send or receive.
	// Corresponds to the Real Time Payments reason code `AM14`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	// The creditor's address is required, but missing or invalid. Corresponds to the
	// Real Time Payments reason code `BE04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_address"
	// The specified creditor is unknown. Corresponds to the Real Time Payments reason
	// code `BE06`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnknownEndCustomer SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unknown_end_customer"
	// The debtor's address is required, but missing or invalid. Corresponds to the
	// Real Time Payments reason code `BE07`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_debtor_address"
	// There was a timeout processing the transfer. Corresponds to the Real Time
	// Payments reason code `DS24`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTimeout SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "timeout"
	// Real Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real Time Payments reason code `NOAT`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnsupportedMessageForRecipient SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unsupported_message_for_recipient"
	// The destination financial institution is currently not connected to Real Time
	// Payments. Corresponds to the Real Time Payments reason code `9912`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRecipientConnectionNotAvailable SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "recipient_connection_not_available"
	// Real Time Payments is currently unavailable. Corresponds to the Real Time
	// Payments reason code `9948`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRealTimePaymentsSuspended SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "real_time_payments_suspended"
	// The destination financial institution is currently signed off of Real Time
	// Payments. Corresponds to the Real Time Payments reason code `9910`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInstructedAgentSignedOff SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "instructed_agent_signed_off"
	// The transfer was rejected due to an internal Increase issue. We have been
	// notified.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeProcessingError SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "processing_error"
	// Some other error or issue has occurred.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeOther SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "other"
)

type SimulationRealTimePaymentsTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound Real Time Payments Transfer is
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
