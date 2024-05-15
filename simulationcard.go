// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationCardService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardService] method instead.
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
	JSON cardAuthorizationSimulationJSON `json:"-"`
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

func (r cardAuthorizationSimulationJSON) RawJSON() string {
	return r.raw
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
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transaction's Account.
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
	JSON cardAuthorizationSimulationDeclinedTransactionJSON `json:"-"`
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

func (r cardAuthorizationSimulationDeclinedTransactionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transaction's Account.
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

func (r CardAuthorizationSimulationDeclinedTransactionCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionCurrencyCad, CardAuthorizationSimulationDeclinedTransactionCurrencyChf, CardAuthorizationSimulationDeclinedTransactionCurrencyEur, CardAuthorizationSimulationDeclinedTransactionCurrencyGbp, CardAuthorizationSimulationDeclinedTransactionCurrencyJpy, CardAuthorizationSimulationDeclinedTransactionCurrencyUsd:
		return true
	}
	return false
}

// The type of the route this Declined Transaction came through.
type CardAuthorizationSimulationDeclinedTransactionRouteType string

const (
	// An Account Number.
	CardAuthorizationSimulationDeclinedTransactionRouteTypeAccountNumber CardAuthorizationSimulationDeclinedTransactionRouteType = "account_number"
	// A Card.
	CardAuthorizationSimulationDeclinedTransactionRouteTypeCard CardAuthorizationSimulationDeclinedTransactionRouteType = "card"
)

func (r CardAuthorizationSimulationDeclinedTransactionRouteType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionRouteTypeAccountNumber, CardAuthorizationSimulationDeclinedTransactionRouteTypeCard:
		return true
	}
	return false
}

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
type CardAuthorizationSimulationDeclinedTransactionSource struct {
	// An ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline CardAuthorizationSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline CardAuthorizationSimulationDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category CardAuthorizationSimulationDeclinedTransactionSourceCategory `json:"category,required"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Check Deposit Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_rejection`.
	CheckDepositRejection CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejection `json:"check_deposit_rejection,required,nullable"`
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// An International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline CardAuthorizationSimulationDeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        cardAuthorizationSimulationDeclinedTransactionSourceJSON        `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceJSON contains the JSON
// metadata for the struct [CardAuthorizationSimulationDeclinedTransactionSource]
type cardAuthorizationSimulationDeclinedTransactionSourceJSON struct {
	ACHDecline                             apijson.Field
	CardDecline                            apijson.Field
	Category                               apijson.Field
	CheckDecline                           apijson.Field
	CheckDepositRejection                  apijson.Field
	InboundRealTimePaymentsTransferDecline apijson.Field
	InternationalACHDecline                apijson.Field
	WireDecline                            apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceJSON) RawJSON() string {
	return r.raw
}

// An ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceACHDecline struct {
	// The ACH Decline's identifier.
	ID string `json:"id,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Inbound ACH Transfer object associated with this decline.
	InboundACHTransferID string `json:"inbound_ach_transfer_id,required"`
	// The descriptive date of the transfer.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date,required,nullable"`
	// The additional information included with the transfer.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	// The identifier of the company that initiated the transfer.
	OriginatorCompanyID string `json:"originator_company_id,required"`
	// The name of the company that initiated the transfer.
	OriginatorCompanyName string `json:"originator_company_name,required"`
	// Why the ACH transfer was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	// The id of the receiver of the transfer.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the receiver of the transfer.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// The trace number of the transfer.
	TraceNumber string `json:"trace_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_decline`.
	Type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineType `json:"type,required"`
	JSON cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceACHDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON struct {
	ID                                 apijson.Field
	Amount                             apijson.Field
	InboundACHTransferID               apijson.Field
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the ACH transfer was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	// The account number is canceled.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	// The transaction would cause an Increase limit to be exceeded.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused. This is a reasonable default reason for decline of
	// credits.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// A rare return reason. The return this message refers to was a duplicate.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// Your account is inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// A rare return reason. The return this message refers to was misrouted.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonMisroutedReturn CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// The originating financial institution made a mistake and this return corrects
	// it.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "return_of_erroneous_or_reversing_debit"
	// The account number that was debited does not exist.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// The originating financial institution asked for this transfer to be returned.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	// Your integration declined this transfer via the API.
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonUserInitiated CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "user_initiated"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonMisroutedReturn, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed, CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonUserInitiated:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineType string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineTypeACHDecline CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineType = "ach_decline"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineTypeACHDecline:
		return true
	}
	return false
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDecline struct {
	// The Card Decline identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActioner `json:"actioner,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
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
	NetworkDetails CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiers `json:"network_identifiers,required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score,required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The declined amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// Fields related to verification of cardholder-provided values.
	Verification CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerification `json:"verification,required"`
	JSON         cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON         `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON struct {
	ID                   apijson.Field
	Actioner             apijson.Field
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
	NetworkIdentifiers   apijson.Field
	NetworkRiskScore     apijson.Field
	PhysicalCardID       apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	ProcessingCategory   apijson.Field
	RealTimeDecisionID   apijson.Field
	Reason               apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActioner string

const (
	// This object was actioned by the user through a real-time decision.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActionerUser CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActioner = "user"
	// This object was actioned by Increase without user intervention.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActionerIncrease CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActionerNetwork CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActioner = "network"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActioner) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActionerUser, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActionerIncrease, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineActionerNetwork:
		return true
	}
	return false
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

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyCad, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyChf, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyEur, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd:
		return true
	}
	return false
}

// Fields specific to the `network`.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory string

const (
	// Visa
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON                    `json:"-"`
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

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                                                `json:"transaction_id,required,nullable"`
	JSON          cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiersJSON `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiersJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiers]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryAccountFunding CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryAutomaticFuelDispenser CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryBillPayment CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory = "bill_payment"
	// A regular purchase.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryPurchase CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryQuasiCash CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryRefund CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory = "refund"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryAccountFunding, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryAutomaticFuelDispenser, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryBillPayment, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryPurchase, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryQuasiCash, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineProcessingCategoryRefund:
		return true
	}
	return false
}

// Why the transaction was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	// The Physical Card was not active.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "physical_card_not_active"
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
	// The transaction was suspected to be fraudulent. Please reach out to
	// support@increase.com for more information.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonSuspectedFraud CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "suspected_fraud"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonSuspectedFraud:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationJSON              `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerification]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON   `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNotChecked CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultMatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNoMatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "no_match"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNotChecked, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultMatch, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddress struct {
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
	Result CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult `json:"result,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON   `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddress]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "no_match"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch, CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type CardAuthorizationSimulationDeclinedTransactionSourceCategory string

const (
	// ACH Decline: details will be under the `ach_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryACHDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "ach_decline"
	// Card Decline: details will be under the `card_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "card_decline"
	// Check Decline: details will be under the `check_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "check_decline"
	// Inbound Real-Time Payments Transfer Decline: details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// International ACH Decline: details will be under the `international_ach_decline`
	// object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInternationalACHDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryWireDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "wire_decline"
	// Check Deposit Rejection: details will be under the `check_deposit_rejection`
	// object.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDepositRejection CardAuthorizationSimulationDeclinedTransactionSourceCategory = "check_deposit_rejection"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryOther CardAuthorizationSimulationDeclinedTransactionSourceCategory = "other"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCategoryACHDecline, CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardDecline, CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDecline, CardAuthorizationSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline, CardAuthorizationSimulationDeclinedTransactionSourceCategoryInternationalACHDecline, CardAuthorizationSimulationDeclinedTransactionSourceCategoryWireDecline, CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDepositRejection, CardAuthorizationSimulationDeclinedTransactionSourceCategoryOther:
		return true
	}
	return false
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// A computer-readable number printed on the MICR line of business checks, usually
	// the check number. This is useful for positive pay checks, but can be unreliably
	// transmitted by the bank of first deposit.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The identifier of the API File object containing an image of the back of the
	// declined check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Transfer object associated with this decline.
	CheckTransferID string `json:"check_transfer_id,required,nullable"`
	// The identifier of the API File object containing an image of the front of the
	// declined check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Inbound Check Deposit object associated with this decline.
	InboundCheckDepositID string `json:"inbound_check_deposit_id,required,nullable"`
	// Why the check was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON   `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON contains
// the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount                apijson.Field
	AuxiliaryOnUs         apijson.Field
	BackImageFileID       apijson.Field
	CheckTransferID       apijson.Field
	FrontImageFileID      apijson.Field
	InboundCheckDepositID apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the check was declined.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	// The account number is disabled.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The account number is canceled.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The deposited check was altered or fictitious.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
	// The transaction would cause a limit to be exceeded.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The account's entity is not active.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Stop payment requested for this check.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// The check was a duplicate deposit.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The check was not authorized.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotAuthorized CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The amount the receiving bank is attempting to deposit does not match the amount
	// on the check.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonAmountMismatch CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "amount_mismatch"
	// The check attempting to be deposited does not belong to Increase.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotOurItem CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	// The account number on the check does not exist at Increase.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNoAccountNumberFound CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "no_account_number_found"
	// The check is not readable. Please refer to the image.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	// The check cannot be processed. This is rare: please contact support.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	// Your integration declined this check via the API.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUserInitiated CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "user_initiated"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotAuthorized, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonAmountMismatch, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotOurItem, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNoAccountNumberFound, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess, CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUserInitiated:
		return true
	}
	return false
}

// A Check Deposit Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_rejection`.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejection struct {
	// The rejected amount in the minor unit of check's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was rejected.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency `json:"currency,required"`
	// Why the check deposit was rejected.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt time.Time                                                                     `json:"rejected_at,required" format:"date-time"`
	JSON       cardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionJSON `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejection]
type cardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	Reason         apijson.Field
	RejectedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency string

const (
	// Canadian Dollar (CAD)
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency = "CAD"
	// Swiss Franc (CHF)
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency = "CHF"
	// Euro (EUR)
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency = "EUR"
	// British Pound (GBP)
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency = "GBP"
	// Japanese Yen (JPY)
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency = "JPY"
	// US Dollar (USD)
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency = "USD"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyCad, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyChf, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyEur, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyGbp, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyJpy, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionCurrencyUsd:
		return true
	}
	return false
}

// Why the check deposit was rejected.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason string

const (
	// The check's image is incomplete.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonIncompleteImage CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "incomplete_image"
	// This is a duplicate check submission.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonDuplicate CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "duplicate"
	// This check has poor image quality.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonPoorImageQuality CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "poor_image_quality"
	// The check was deposited with the incorrect amount.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonIncorrectAmount CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "incorrect_amount"
	// The check is made out to someone other than the account holder.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonIncorrectRecipient CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "incorrect_recipient"
	// This check was not eligible for mobile deposit.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonNotEligibleForMobileDeposit CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "not_eligible_for_mobile_deposit"
	// This check is missing at least one required field.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonMissingRequiredDataElements CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "missing_required_data_elements"
	// This check is suspected to be fraudulent.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonSuspectedFraud CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "suspected_fraud"
	// This check's deposit window has expired.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonDepositWindowExpired CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "deposit_window_expired"
	// The check was rejected for an unknown reason.
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonUnknown CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason = "unknown"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReason) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonIncompleteImage, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonDuplicate, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonPoorImageQuality, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonIncorrectAmount, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonIncorrectRecipient, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonNotEligibleForMobileDeposit, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonMissingRequiredDataElements, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonSuspectedFraud, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonDepositWindowExpired, CardAuthorizationSimulationDeclinedTransactionSourceCheckDepositRejectionReasonUnknown:
		return true
	}
	return false
}

// An Inbound Real-Time Payments Transfer Decline object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real-Time Payments
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
	// The Real-Time Payments network identification of the declined transfer.
	TransactionIdentification string                                                                                         `json:"transaction_identification,required"`
	JSON                      cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON `json:"-"`
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

func (r cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real-Time Payments
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

func (r CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd:
		return true
	}
	return false
}

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
	// Your account is not enabled to receive Real-Time Payments transfers.
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive, CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled:
		return true
	}
	return false
}

// An International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline struct {
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
	ForeignExchangeIndicator CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
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
	OriginatingDepositoryFinancialInstitutionIDQualifier CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
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
	ReceivingDepositoryFinancialInstitutionIDQualifier CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string                                                                          `json:"trace_number,required"`
	JSON        cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON `json:"-"`
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

func (r cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON) RawJSON() string {
	return r.raw
}

// A description of how the foreign exchange rate was calculated.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToVariable CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorVariableToFixed CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToFixed CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "fixed_to_fixed"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToVariable, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorVariableToFixed, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToFixed:
		return true
	}
	return false
}

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeRate CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorBlank CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "blank"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeRate, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorBlank:
		return true
	}
	return false
}

// The type of transfer. Set by the originator.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAnnuity CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBusinessOrCommercial CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeDeposit CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeLoan CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMiscellaneous CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMortgage CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePension CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRemittance CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRentOrLease CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSalaryOrPayroll CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTax CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAccountsReceivable CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBackOfficeConversion CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMachineTransfer CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfPurchase CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfSale CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRepresentedCheck CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSharedNetworkTransaction CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTelphoneInitiated CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeInternetInitiated CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "internet_initiated"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAnnuity, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBusinessOrCommercial, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeDeposit, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeLoan, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMiscellaneous, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMortgage, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePension, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRemittance, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRentOrLease, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSalaryOrPayroll, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTax, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAccountsReceivable, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBackOfficeConversion, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMachineTransfer, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfPurchase, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfSale, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRepresentedCheck, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSharedNetworkTransaction, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTelphoneInitiated, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeInternetInitiated:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierBicCode CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierIban CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierBicCode, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierBicCode CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierIban CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierBicCode, CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceWireDecline struct {
	// The identifier of the Inbound Wire Transfer that was declined.
	InboundWireTransferID string `json:"inbound_wire_transfer_id,required"`
	// Why the wire transfer was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON   `json:"-"`
}

// cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceWireDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON struct {
	InboundWireTransferID apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled, CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled, CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonEntityNotActive, CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonGroupLocked, CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonNoAccountNumber, CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
type CardAuthorizationSimulationDeclinedTransactionType string

const (
	CardAuthorizationSimulationDeclinedTransactionTypeDeclinedTransaction CardAuthorizationSimulationDeclinedTransactionType = "declined_transaction"
)

func (r CardAuthorizationSimulationDeclinedTransactionType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationDeclinedTransactionTypeDeclinedTransaction:
		return true
	}
	return false
}

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
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transaction's Account.
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
	JSON cardAuthorizationSimulationPendingTransactionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
// Transaction's currency. This will match the currency on the Pending
// Transaction's Account.
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

func (r CardAuthorizationSimulationPendingTransactionCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionCurrencyCad, CardAuthorizationSimulationPendingTransactionCurrencyChf, CardAuthorizationSimulationPendingTransactionCurrencyEur, CardAuthorizationSimulationPendingTransactionCurrencyGbp, CardAuthorizationSimulationPendingTransactionCurrencyJpy, CardAuthorizationSimulationPendingTransactionCurrencyUsd:
		return true
	}
	return false
}

// The type of the route this Pending Transaction came through.
type CardAuthorizationSimulationPendingTransactionRouteType string

const (
	// An Account Number.
	CardAuthorizationSimulationPendingTransactionRouteTypeAccountNumber CardAuthorizationSimulationPendingTransactionRouteType = "account_number"
	// A Card.
	CardAuthorizationSimulationPendingTransactionRouteTypeCard CardAuthorizationSimulationPendingTransactionRouteType = "card"
)

func (r CardAuthorizationSimulationPendingTransactionRouteType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionRouteTypeAccountNumber, CardAuthorizationSimulationPendingTransactionRouteTypeCard:
		return true
	}
	return false
}

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
type CardAuthorizationSimulationPendingTransactionSource struct {
	// An Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction,required,nullable"`
	// An ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorization `json:"card_authorization,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category CardAuthorizationSimulationPendingTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction,required,nullable"`
	// An Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// A Real-Time Payments Transfer Instruction object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_instruction`.
	RealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction `json:"real_time_payments_transfer_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction,required,nullable"`
	JSON                    cardAuthorizationSimulationPendingTransactionSourceJSON                    `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceJSON) RawJSON() string {
	return r.raw
}

// An Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string                                                                            `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyCad, CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyChf, CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyEur, CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyGbp, CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyJpy, CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyUsd:
		return true
	}
	return false
}

// An ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID string                                                                        `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorization struct {
	// The Card Authorization identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActioner `json:"actioner,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency `json:"currency,required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The direction descibes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirection `json:"direction,required"`
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
	// Fields specific to the `network`.
	NetworkDetails CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiers `json:"network_identifiers,required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The pending amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType `json:"type,required"`
	// Fields related to verification of cardholder-provided values.
	Verification CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerification `json:"verification,required"`
	JSON         cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON         `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorization]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON struct {
	ID                   apijson.Field
	Actioner             apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	Direction            apijson.Field
	ExpiresAt            apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	NetworkDetails       apijson.Field
	NetworkIdentifiers   apijson.Field
	NetworkRiskScore     apijson.Field
	PendingTransactionID apijson.Field
	PhysicalCardID       apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	ProcessingCategory   apijson.Field
	RealTimeDecisionID   apijson.Field
	Type                 apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActioner string

const (
	// This object was actioned by the user through a real-time decision.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActionerUser CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActioner = "user"
	// This object was actioned by Increase without user intervention.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActionerIncrease CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActionerNetwork CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActioner = "network"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActioner) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActionerUser, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActionerIncrease, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationActionerNetwork:
		return true
	}
	return false
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

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyCad, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyChf, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyEur, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyGbp, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyJpy, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyUsd:
		return true
	}
	return false
}

// The direction descibes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirection string

const (
	// A regular card authorization where funds are debited from the cardholder.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirectionSettlement CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirection = "settlement"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirectionRefund CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirection = "refund"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirection) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirectionSettlement, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationDirectionRefund:
		return true
	}
	return false
}

// Fields specific to the `network`.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsCategory string

const (
	// Visa
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsCategoryVisa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsCategory = "visa"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                                                     `json:"transaction_id,required,nullable"`
	JSON          cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiers]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryAccountFunding CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryAutomaticFuelDispenser CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryBillPayment CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory = "bill_payment"
	// A regular purchase.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryPurchase CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryQuasiCash CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryRefund CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory = "refund"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryAccountFunding, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryAutomaticFuelDispenser, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryBillPayment, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryPurchase, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryQuasiCash, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationProcessingCategoryRefund:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_authorization`.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationTypeCardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType = "card_authorization"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationTypeCardAuthorization:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationJSON              `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerification]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON   `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCode]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNotChecked CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultMatch CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNoMatch CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNotChecked, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultMatch, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddress struct {
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
	Result CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult `json:"result,required"`
	JSON   cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON   `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddress]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNotChecked CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultMatch CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNoMatch CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "no_match"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNotChecked, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultMatch, CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type CardAuthorizationSimulationPendingTransactionSourceCategory string

const (
	// Account Transfer Instruction: details will be under the
	// `account_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryAccountTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "account_transfer_instruction"
	// ACH Transfer Instruction: details will be under the `ach_transfer_instruction`
	// object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryACHTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "ach_transfer_instruction"
	// Card Authorization: details will be under the `card_authorization` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardAuthorization CardAuthorizationSimulationPendingTransactionSourceCategory = "card_authorization"
	// Check Deposit Instruction: details will be under the `check_deposit_instruction`
	// object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckDepositInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "check_deposit_instruction"
	// Check Transfer Instruction: details will be under the
	// `check_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "check_transfer_instruction"
	// Inbound Funds Hold: details will be under the `inbound_funds_hold` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryInboundFundsHold CardAuthorizationSimulationPendingTransactionSourceCategory = "inbound_funds_hold"
	// Real-Time Payments Transfer Instruction: details will be under the
	// `real_time_payments_transfer_instruction` object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryRealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	// Wire Transfer Instruction: details will be under the `wire_transfer_instruction`
	// object.
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_transfer_instruction"
	// The Pending Transaction was made for an undocumented or deprecated reason.
	CardAuthorizationSimulationPendingTransactionSourceCategoryOther CardAuthorizationSimulationPendingTransactionSourceCategory = "other"
)

func (r CardAuthorizationSimulationPendingTransactionSourceCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCategoryAccountTransferInstruction, CardAuthorizationSimulationPendingTransactionSourceCategoryACHTransferInstruction, CardAuthorizationSimulationPendingTransactionSourceCategoryCardAuthorization, CardAuthorizationSimulationPendingTransactionSourceCategoryCheckDepositInstruction, CardAuthorizationSimulationPendingTransactionSourceCategoryCheckTransferInstruction, CardAuthorizationSimulationPendingTransactionSourceCategoryInboundFundsHold, CardAuthorizationSimulationPendingTransactionSourceCategoryRealTimePaymentsTransferInstruction, CardAuthorizationSimulationPendingTransactionSourceCategoryWireTransferInstruction, CardAuthorizationSimulationPendingTransactionSourceCategoryOther:
		return true
	}
	return false
}

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
	FrontImageFileID string                                                                         `json:"front_image_file_id,required"`
	JSON             cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyCad, CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyChf, CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyEur, CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyGbp, CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyJpy, CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyUsd:
		return true
	}
	return false
}

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
	TransferID string                                                                          `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyCad, CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyChf, CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyEur, CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyGbp, CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyJpy, CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyUsd:
		return true
	}
	return false
}

// An Inbound Funds Hold object. This field will be present in the JSON response if
// and only if `category` is equal to `inbound_funds_hold`.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold struct {
	// The Inbound Funds Hold identifier.
	ID string `json:"id,required"`
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
	// A constant representing the object's type. For this resource it will always be
	// `inbound_funds_hold`.
	Type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldType `json:"type,required"`
	JSON cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON `json:"-"`
}

// cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON contains
// the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold]
type cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON struct {
	ID                      apijson.Field
	Amount                  apijson.Field
	AutomaticallyReleasesAt apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyCad, CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyChf, CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyEur, CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyGbp, CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyJpy, CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyUsd:
		return true
	}
	return false
}

// The status of the hold.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus string

const (
	// Funds are still being held.
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusHeld CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "held"
	// Funds have been released.
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusComplete CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "complete"
)

func (r CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusHeld, CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_funds_hold`.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldType string

const (
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldTypeInboundFundsHold CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldType = "inbound_funds_hold"
)

func (r CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldTypeInboundFundsHold:
		return true
	}
	return false
}

// A Real-Time Payments Transfer Instruction object. This field will be present in
// the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Pending
	// Transaction.
	TransferID string                                                                                     `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
	// The account number for the destination account.
	AccountNumber string `json:"account_number,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number,required"`
	// The identifier of the Wire Transfer that led to this Pending Transaction.
	TransferID string                                                                         `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON `json:"-"`
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

func (r cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON) RawJSON() string {
	return r.raw
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

func (r CardAuthorizationSimulationPendingTransactionStatus) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionStatusPending, CardAuthorizationSimulationPendingTransactionStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
type CardAuthorizationSimulationPendingTransactionType string

const (
	CardAuthorizationSimulationPendingTransactionTypePendingTransaction CardAuthorizationSimulationPendingTransactionType = "pending_transaction"
)

func (r CardAuthorizationSimulationPendingTransactionType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationPendingTransactionTypePendingTransaction:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_card_authorization_simulation_result`.
type CardAuthorizationSimulationType string

const (
	CardAuthorizationSimulationTypeInboundCardAuthorizationSimulationResult CardAuthorizationSimulationType = "inbound_card_authorization_simulation_result"
)

func (r CardAuthorizationSimulationType) IsKnown() bool {
	switch r {
	case CardAuthorizationSimulationTypeInboundCardAuthorizationSimulationResult:
		return true
	}
	return false
}

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
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID param.Field[string] `json:"merchant_acceptor_id"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode param.Field[string] `json:"merchant_category_code"`
	// The city the merchant resides in.
	MerchantCity param.Field[string] `json:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry param.Field[string] `json:"merchant_country"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor param.Field[string] `json:"merchant_descriptor"`
	// The identifier of the Physical Card to be authorized.
	PhysicalCardID param.Field[string] `json:"physical_card_id"`
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
