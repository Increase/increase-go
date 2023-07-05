// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// SimulationCardService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationCardService] method
// instead.
type SimulationCardService struct {
	Options []option.RequestOption
}

// NewSimulationCardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimulationCardService(opts ...option.RequestOption) (r *SimulationCardService) {
	r = &SimulationCardService{}
	r.Options = opts
	return
}

// Simulates a purchase authorization on a [Card](#cards). Depending on the balance
// available to the card and the `amount` submitted, the authorization activity
// will result in a [Pending Transaction](#pending-transactions) of type
// `card_authorization` or a [Declined Transaction](#declined-transactions) of type
// `card_decline`. You can pass either a Card id or a
// [Digital Wallet Token](#digital-wallet-tokens) id to simulate the two different
// ways purchases can be made.
func (r *SimulationCardService) Authorize(ctx context.Context, body SimulationCardAuthorizeParams, opts ...option.RequestOption) (res *CardAuthorizationSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the settlement of an authorization by a card acquirer. After a card
// authorization is created, the merchant will eventually send a settlement. This
// simulates that event, which may occur many days after the purchase in
// production. The amount settled can be different from the amount originally
// authorized, for example, when adding a tip to a restaurant bill.
func (r *SimulationCardService) Settlement(ctx context.Context, body SimulationCardSettlementParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_settlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of a Card Authorization simulation.
type CardAuthorizationSimulation struct {
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction CardAuthorizationSimulationDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction CardAuthorizationSimulationPendingTransaction `json:"pending_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type CardAuthorizationSimulationType `json:"type,required"`
	JSON cardAuthorizationSimulationJSON
}

// cardAuthorizationSimulationJSON contains the JSON metadata for the struct
// [CardAuthorizationSimulation]
type cardAuthorizationSimulationJSON struct {
	DeclinedTransaction apijson.Field
	PendingTransaction  apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardAuthorizationSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the authorization attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: card_decline`.
type CardAuthorizationSimulationDeclinedTransaction struct {
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
	Currency CardAuthorizationSimulationDeclinedTransactionCurrency `json:"currency,required"`
	// This is the description the vendor provides.
	Description string `json:"description,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType CardAuthorizationSimulationDeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source CardAuthorizationSimulationDeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type CardAuthorizationSimulationDeclinedTransactionType `json:"type,required"`
	JSON cardAuthorizationSimulationDeclinedTransactionJSON
}

// cardAuthorizationSimulationDeclinedTransactionJSON contains the JSON metadata
// for the struct [CardAuthorizationSimulationDeclinedTransaction]
type cardAuthorizationSimulationDeclinedTransactionJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
type CardAuthorizationSimulationDeclinedTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationDeclinedTransactionCurrencyCad CardAuthorizationSimulationDeclinedTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationDeclinedTransactionCurrencyChf CardAuthorizationSimulationDeclinedTransactionCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationDeclinedTransactionCurrencyEur CardAuthorizationSimulationDeclinedTransactionCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationDeclinedTransactionCurrencyGbp CardAuthorizationSimulationDeclinedTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationDeclinedTransactionCurrencyJpy CardAuthorizationSimulationDeclinedTransactionCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationDeclinedTransactionCurrencyUsd CardAuthorizationSimulationDeclinedTransactionCurrency = "USD"
)

// The type of the route this Declined Transaction came through.
type CardAuthorizationSimulationDeclinedTransactionRouteType string

const (
	// An Account Number.
	CardAuthorizationSimulationDeclinedTransactionRouteTypeAccountNumber CardAuthorizationSimulationDeclinedTransactionRouteType = "account_number"
	// A Card.
	CardAuthorizationSimulationDeclinedTransactionRouteTypeCard CardAuthorizationSimulationDeclinedTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
type CardAuthorizationSimulationDeclinedTransactionSource struct {
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline CardAuthorizationSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline CardAuthorizationSimulationDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category CardAuthorizationSimulationDeclinedTransactionSourceCategory `json:"category,required"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline CardAuthorizationSimulationDeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        cardAuthorizationSimulationDeclinedTransactionSourceJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceJSON contains the JSON
// metadata for the struct [CardAuthorizationSimulationDeclinedTransactionSource]
type cardAuthorizationSimulationDeclinedTransactionSourceJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	// Why the ACH transfer was declined.
	Reason           CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                                               `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                                               `json:"receiver_name,required,nullable"`
	TraceNumber      string                                                               `json:"trace_number,required"`
	JSON             cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceACHDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH transfer was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	// The account number is canceled.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// Other.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// Your account is inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// Other.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonMisroutedReturn CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// The account number that was debited does not exist.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// Other.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
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
	Network CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

// The payment network used to process this card authorization
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	// Visa
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkVisa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

// Fields specific to the `network`
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to the `visa` network
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

// Why the transaction was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	// The account's entity was not active.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	// The account was inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by an internal limit for new Increase accounts.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesInternalLimit CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_internal_limit"
	// The transaction was blocked by a Limit.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	// Your application declined the transaction via webhook.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	// Your application webhook did not respond without the required timeout.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	// Declined by stand-in processing.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	// The card read had an invalid CVV, dCVV, or authorization request cryptogram.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	// The original card authorization for this incremental authorization does not
	// exist.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
)

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
type CardAuthorizationSimulationDeclinedTransactionSourceCategory string

const (
	// The Declined Transaction was created by a ACH Decline object. Details will be
	// under the `ach_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryACHDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "ach_decline"
	// The Declined Transaction was created by a Card Decline object. Details will be
	// under the `card_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "card_decline"
	// The Declined Transaction was created by a Check Decline object. Details will be
	// under the `check_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "check_decline"
	// The Declined Transaction was created by a Inbound Real Time Payments Transfer
	// Decline object. Details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// The Declined Transaction was created by a International ACH Decline object.
	// Details will be under the `international_ach_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInternationalACHDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	// The Declined Transaction was created by a Wire Decline object. Details will be
	// under the `wire_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryWireDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "wire_decline"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryOther CardAuthorizationSimulationDeclinedTransactionSourceCategory = "other"
)

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON contains
// the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Field
	AuxiliaryOnUs apijson.Field
	Reason        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the check was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	// The account number is canceled.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The account's entity is not active.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Unable to locate account.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	// Routing number on the check is not ours.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotOurItem CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	// Unable to process.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	// Refer to image.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	// Stop payment requested for this check.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// Check was returned to sender.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReturned CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "returned"
	// The check was a duplicate deposit.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The transaction is not allowed.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotAuthorized CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The check was altered or fictitious.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
)

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Why the transfer was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

// Why the transfer was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	// The account number is canceled.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	// The account number is disabled.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	// Your account is restricted.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	// Your account is inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	// The account's entity is not active.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	// Your account is not enabled to receive Real Time Payments transfers.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline struct {
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
	JSON                                                   cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceWireDecline struct {
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
	Reason CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceWireDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON struct {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the wire transfer was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason string

const (
	// The account number is canceled.
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "account_number_canceled"
	// The account number is disabled.
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "account_number_disabled"
	// The account's entity is not active.
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "entity_not_active"
	// Your account is inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "group_locked"
	// The beneficiary account number does not exist.
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonNoAccountNumber CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "no_account_number"
	// The transaction is not allowed per Increase's terms.
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "transaction_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
type CardAuthorizationSimulationDeclinedTransactionType string

const (
	CardAuthorizationSimulationDeclinedTransactionTypeDeclinedTransaction CardAuthorizationSimulationDeclinedTransactionType = "declined_transaction"
)

// If the authorization attempt succeeds, this will contain the resulting Pending
// Transaction object. The Pending Transaction's `source` will be of
// `category: card_authorization`.
type CardAuthorizationSimulationPendingTransaction struct {
	// The Pending Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the account this Pending Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction was completed.
	CompletedAt time.Time `json:"completed_at,required,nullable" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency CardAuthorizationSimulationPendingTransactionCurrency `json:"currency,required"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description string `json:"description,required"`
	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Pending Transaction came through.
	RouteType CardAuthorizationSimulationPendingTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source CardAuthorizationSimulationPendingTransactionSource `json:"source,required"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status CardAuthorizationSimulationPendingTransactionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type CardAuthorizationSimulationPendingTransactionType `json:"type,required"`
	JSON cardAuthorizationSimulationPendingTransactionJSON
}

// cardAuthorizationSimulationPendingTransactionJSON contains the JSON metadata for
// the struct [CardAuthorizationSimulationPendingTransaction]
type cardAuthorizationSimulationPendingTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CompletedAt apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
// Transaction's currency. This will match the currency on the Pending
// Transcation's Account.
type CardAuthorizationSimulationPendingTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationPendingTransactionCurrencyCad CardAuthorizationSimulationPendingTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationPendingTransactionCurrencyChf CardAuthorizationSimulationPendingTransactionCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationPendingTransactionCurrencyEur CardAuthorizationSimulationPendingTransactionCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationPendingTransactionCurrencyGbp CardAuthorizationSimulationPendingTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationPendingTransactionCurrencyJpy CardAuthorizationSimulationPendingTransactionCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationPendingTransactionCurrencyUsd CardAuthorizationSimulationPendingTransactionCurrency = "USD"
)

// The type of the route this Pending Transaction came through.
type CardAuthorizationSimulationPendingTransactionRouteType string

const (
	// An Account Number.
	CardAuthorizationSimulationPendingTransactionRouteTypeAccountNumber CardAuthorizationSimulationPendingTransactionRouteType = "account_number"
	// A Card.
	CardAuthorizationSimulationPendingTransactionRouteTypeCard CardAuthorizationSimulationPendingTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
type CardAuthorizationSimulationPendingTransactionSource struct {
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction,required,nullable"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorization `json:"card_authorization,required,nullable"`
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category CardAuthorizationSimulationPendingTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction,required,nullable"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// A Real Time Payments Transfer Instruction object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_instruction`.
	RealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction `json:"real_time_payments_transfer_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction,required,nullable"`
	JSON                    cardAuthorizationSimulationPendingTransactionSourceJSON
}

// cardAuthorizationSimulationPendingTransactionSourceJSON contains the JSON
// metadata for the struct [CardAuthorizationSimulationPendingTransactionSource]
type cardAuthorizationSimulationPendingTransactionSourceJSON struct {
	AccountTransferInstruction          apijson.Field
	ACHTransferInstruction              apijson.Field
	CardAuthorization                   apijson.Field
	Category                            apijson.Field
	CheckDepositInstruction             apijson.Field
	CheckTransferInstruction            apijson.Field
	InboundFundsHold                    apijson.Field
	RealTimePaymentsTransferInstruction apijson.Field
	WireTransferInstruction             apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorization struct {
	// The Card Authorization identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency `json:"currency,required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) when this authorization
	// will expire and the pending transaction will be released.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
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
	// The payment network used to process this card authorization
	Network CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails `json:"network_details,required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType `json:"type,required"`
	JSON cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorization]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	ExpiresAt            apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	Network              apijson.Field
	NetworkDetails       apijson.Field
	PendingTransactionID apijson.Field
	RealTimeDecisionID   apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "USD"
)

// The payment network used to process this card authorization
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork string

const (
	// Visa
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkVisa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork = "visa"
)

// Fields specific to the `network`
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa `json:"visa,required"`
	JSON cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON struct {
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to the `visa` network
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

// A constant representing the object's type. For this resource it will always be
// `card_authorization`.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationTypeCardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType = "card_authorization"
)

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
type CardAuthorizationSimulationPendingTransactionSourceCategory string

const (
	// The Pending Transaction was created by a Account Transfer Instruction object.
	// Details will be under the `account_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryAccountTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "account_transfer_instruction"
	// The Pending Transaction was created by a ACH Transfer Instruction object.
	// Details will be under the `ach_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryACHTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "ach_transfer_instruction"
	// The Pending Transaction was created by a Card Authorization object. Details will
	// be under the `card_authorization` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardAuthorization CardAuthorizationSimulationPendingTransactionSourceCategory = "card_authorization"
	// The Pending Transaction was created by a Check Deposit Instruction object.
	// Details will be under the `check_deposit_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckDepositInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "check_deposit_instruction"
	// The Pending Transaction was created by a Check Transfer Instruction object.
	// Details will be under the `check_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "check_transfer_instruction"
	// The Pending Transaction was created by a Inbound Funds Hold object. Details will
	// be under the `inbound_funds_hold` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryInboundFundsHold CardAuthorizationSimulationPendingTransactionSourceCategory = "inbound_funds_hold"
	// The Pending Transaction was created by a Real Time Payments Transfer Instruction
	// object. Details will be under the `real_time_payments_transfer_instruction`
	// object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryRealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	// The Pending Transaction was created by a Wire Transfer Instruction object.
	// Details will be under the `wire_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_transfer_instruction"
	// The Pending Transaction was made for an undocumented or deprecated reason.
	CardAuthorizationSimulationPendingTransactionSourceCategoryOther CardAuthorizationSimulationPendingTransactionSourceCategory = "other"
)

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Deposit.
	CheckDepositID string `json:"check_deposit_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency `json:"currency,required"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID string `json:"front_image_file_id,required"`
	JSON             cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction]
type cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON struct {
	Amount           apijson.Field
	BackImageFileID  apijson.Field
	CheckDepositID   apijson.Field
	Currency         apijson.Field
	FrontImageFileID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

// A Inbound Funds Hold object. This field will be present in the JSON response if
// and only if `category` is equal to `inbound_funds_hold`.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold struct {
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency `json:"currency,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus `json:"status,required"`
	JSON   cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON
}

// cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON contains
// the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold]
type cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON struct {
	Amount                  apijson.Field
	AutomaticallyReleasesAt apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyCad CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyChf CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyEur CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "USD"
)

// The status of the hold.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus string

const (
	// Funds are still being held.
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusHeld CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "held"
	// Funds have been released.
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusComplete CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "complete"
)

// A Real Time Payments Transfer Instruction object. This field will be present in
// the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Real Time Payments Transfer that led to this Pending
	// Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
	AccountNumber string `json:"account_number,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	RoutingNumber      string `json:"routing_number,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Pending Transaction has been confirmed and has an associated
// Transaction.
type CardAuthorizationSimulationPendingTransactionStatus string

const (
	// The Pending Transaction is still awaiting confirmation.
	CardAuthorizationSimulationPendingTransactionStatusPending CardAuthorizationSimulationPendingTransactionStatus = "pending"
	// The Pending Transaction is confirmed. An associated Transaction exists for this
	// object. The Pending Transaction will no longer count against your balance and
	// can generally be hidden from UIs, etc.
	CardAuthorizationSimulationPendingTransactionStatusComplete CardAuthorizationSimulationPendingTransactionStatus = "complete"
)

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
type CardAuthorizationSimulationPendingTransactionType string

const (
	CardAuthorizationSimulationPendingTransactionTypePendingTransaction CardAuthorizationSimulationPendingTransactionType = "pending_transaction"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_card_authorization_simulation_result`.
type CardAuthorizationSimulationType string

const (
	CardAuthorizationSimulationTypeInboundCardAuthorizationSimulationResult CardAuthorizationSimulationType = "inbound_card_authorization_simulation_result"
)

type SimulationCardAuthorizeParams struct {
	// The authorization amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID param.Field[string] `json:"digital_wallet_token_id"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID param.Field[string] `json:"event_subscription_id"`
}

func (r SimulationCardAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardSettlementParams struct {
	// The identifier of the Card to create a settlement on.
	CardID param.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID param.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount param.Field[int64] `json:"amount"`
}

func (r SimulationCardSettlementParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
