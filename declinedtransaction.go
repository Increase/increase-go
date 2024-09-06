// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/pagination"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// DeclinedTransactionService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeclinedTransactionService] method instead.
type DeclinedTransactionService struct {
	Options []option.RequestOption
}

// NewDeclinedTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeclinedTransactionService(opts ...option.RequestOption) (r *DeclinedTransactionService) {
	r = &DeclinedTransactionService{}
	r.Options = opts
	return
}

// Retrieve a Declined Transaction
func (r *DeclinedTransactionService) Get(ctx context.Context, declinedTransactionID string, opts ...option.RequestOption) (res *DeclinedTransaction, err error) {
	opts = append(r.Options[:], opts...)
	if declinedTransactionID == "" {
		err = errors.New("missing required declined_transaction_id parameter")
		return
	}
	path := fmt.Sprintf("declined_transactions/%s", declinedTransactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Declined Transactions
func (r *DeclinedTransactionService) List(ctx context.Context, query DeclinedTransactionListParams, opts ...option.RequestOption) (res *pagination.Page[DeclinedTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "declined_transactions"
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

// List Declined Transactions
func (r *DeclinedTransactionService) ListAutoPaging(ctx context.Context, query DeclinedTransactionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[DeclinedTransaction] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Declined Transactions are refused additions and removals of money from your bank
// account. For example, Declined Transactions are caused when your Account has an
// insufficient balance or your Limits are triggered.
type DeclinedTransaction struct {
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
	Currency DeclinedTransactionCurrency `json:"currency,required"`
	// This is the description the vendor provides.
	Description string `json:"description,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType DeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source DeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type DeclinedTransactionType `json:"type,required"`
	JSON declinedTransactionJSON `json:"-"`
}

// declinedTransactionJSON contains the JSON metadata for the struct
// [DeclinedTransaction]
type declinedTransactionJSON struct {
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

func (r *DeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transaction's Account.
type DeclinedTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	DeclinedTransactionCurrencyCad DeclinedTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	DeclinedTransactionCurrencyChf DeclinedTransactionCurrency = "CHF"
	// Euro (EUR)
	DeclinedTransactionCurrencyEur DeclinedTransactionCurrency = "EUR"
	// British Pound (GBP)
	DeclinedTransactionCurrencyGbp DeclinedTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	DeclinedTransactionCurrencyJpy DeclinedTransactionCurrency = "JPY"
	// US Dollar (USD)
	DeclinedTransactionCurrencyUsd DeclinedTransactionCurrency = "USD"
)

func (r DeclinedTransactionCurrency) IsKnown() bool {
	switch r {
	case DeclinedTransactionCurrencyCad, DeclinedTransactionCurrencyChf, DeclinedTransactionCurrencyEur, DeclinedTransactionCurrencyGbp, DeclinedTransactionCurrencyJpy, DeclinedTransactionCurrencyUsd:
		return true
	}
	return false
}

// The type of the route this Declined Transaction came through.
type DeclinedTransactionRouteType string

const (
	// An Account Number.
	DeclinedTransactionRouteTypeAccountNumber DeclinedTransactionRouteType = "account_number"
	// A Card.
	DeclinedTransactionRouteTypeCard DeclinedTransactionRouteType = "card"
	// A Lockbox.
	DeclinedTransactionRouteTypeLockbox DeclinedTransactionRouteType = "lockbox"
)

func (r DeclinedTransactionRouteType) IsKnown() bool {
	switch r {
	case DeclinedTransactionRouteTypeAccountNumber, DeclinedTransactionRouteTypeCard, DeclinedTransactionRouteTypeLockbox:
		return true
	}
	return false
}

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
type DeclinedTransactionSource struct {
	// An ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline DeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline DeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category DeclinedTransactionSourceCategory `json:"category,required"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline DeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Check Deposit Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_rejection`.
	CheckDepositRejection DeclinedTransactionSourceCheckDepositRejection `json:"check_deposit_rejection,required,nullable"`
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline DeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        declinedTransactionSourceJSON        `json:"-"`
}

// declinedTransactionSourceJSON contains the JSON metadata for the struct
// [DeclinedTransactionSource]
type declinedTransactionSourceJSON struct {
	ACHDecline                             apijson.Field
	CardDecline                            apijson.Field
	Category                               apijson.Field
	CheckDecline                           apijson.Field
	CheckDepositRejection                  apijson.Field
	InboundRealTimePaymentsTransferDecline apijson.Field
	WireDecline                            apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *DeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceJSON) RawJSON() string {
	return r.raw
}

// An ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type DeclinedTransactionSourceACHDecline struct {
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
	Reason DeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	// The id of the receiver of the transfer.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the receiver of the transfer.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// The trace number of the transfer.
	TraceNumber string `json:"trace_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_decline`.
	Type DeclinedTransactionSourceACHDeclineType `json:"type,required"`
	JSON declinedTransactionSourceACHDeclineJSON `json:"-"`
}

// declinedTransactionSourceACHDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceACHDecline]
type declinedTransactionSourceACHDeclineJSON struct {
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

func (r *DeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceACHDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the ACH transfer was declined.
type DeclinedTransactionSourceACHDeclineReason string

const (
	// The account number is canceled.
	DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled DeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled DeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	// The transaction would cause an Increase limit to be exceeded.
	DeclinedTransactionSourceACHDeclineReasonBreachesLimit DeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused. This is a reasonable default reason for decline of
	// credits.
	DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver DeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// A rare return reason. The return this message refers to was a duplicate.
	DeclinedTransactionSourceACHDeclineReasonDuplicateReturn DeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	DeclinedTransactionSourceACHDeclineReasonEntityNotActive DeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// There was an error with one of the required fields.
	DeclinedTransactionSourceACHDeclineReasonFieldError DeclinedTransactionSourceACHDeclineReason = "field_error"
	// Your account is inactive.
	DeclinedTransactionSourceACHDeclineReasonGroupLocked DeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	DeclinedTransactionSourceACHDeclineReasonInsufficientFunds DeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// A rare return reason. The return this message refers to was misrouted.
	DeclinedTransactionSourceACHDeclineReasonMisroutedReturn DeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// The originating financial institution made a mistake and this return corrects
	// it.
	DeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit DeclinedTransactionSourceACHDeclineReason = "return_of_erroneous_or_reversing_debit"
	// The account number that was debited does not exist.
	DeclinedTransactionSourceACHDeclineReasonNoACHRoute DeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// The originating financial institution asked for this transfer to be returned.
	DeclinedTransactionSourceACHDeclineReasonOriginatorRequest DeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	DeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed DeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	// Your integration declined this transfer via the API.
	DeclinedTransactionSourceACHDeclineReasonUserInitiated DeclinedTransactionSourceACHDeclineReason = "user_initiated"
)

func (r DeclinedTransactionSourceACHDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled, DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled, DeclinedTransactionSourceACHDeclineReasonBreachesLimit, DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver, DeclinedTransactionSourceACHDeclineReasonDuplicateReturn, DeclinedTransactionSourceACHDeclineReasonEntityNotActive, DeclinedTransactionSourceACHDeclineReasonFieldError, DeclinedTransactionSourceACHDeclineReasonGroupLocked, DeclinedTransactionSourceACHDeclineReasonInsufficientFunds, DeclinedTransactionSourceACHDeclineReasonMisroutedReturn, DeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit, DeclinedTransactionSourceACHDeclineReasonNoACHRoute, DeclinedTransactionSourceACHDeclineReasonOriginatorRequest, DeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed, DeclinedTransactionSourceACHDeclineReasonUserInitiated:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `ach_decline`.
type DeclinedTransactionSourceACHDeclineType string

const (
	DeclinedTransactionSourceACHDeclineTypeACHDecline DeclinedTransactionSourceACHDeclineType = "ach_decline"
)

func (r DeclinedTransactionSourceACHDeclineType) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceACHDeclineTypeACHDecline:
		return true
	}
	return false
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type DeclinedTransactionSourceCardDecline struct {
	// The Card Decline identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner DeclinedTransactionSourceCardDeclineActioner `json:"actioner,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency DeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// The identifier of the declined transaction created for this Card Decline.
	DeclinedTransactionID string `json:"declined_transaction_id,required"`
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
	// The merchant's postal code. For US merchants this is either a 5-digit or 9-digit
	// ZIP code, where the first 5 and last 4 are separated by a dash.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Fields specific to the `network`.
	NetworkDetails DeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers DeclinedTransactionSourceCardDeclineNetworkIdentifiers `json:"network_identifiers,required"`
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
	ProcessingCategory DeclinedTransactionSourceCardDeclineProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason DeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// Fields related to verification of cardholder-provided values.
	Verification DeclinedTransactionSourceCardDeclineVerification `json:"verification,required"`
	JSON         declinedTransactionSourceCardDeclineJSON         `json:"-"`
}

// declinedTransactionSourceCardDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceCardDecline]
type declinedTransactionSourceCardDeclineJSON struct {
	ID                    apijson.Field
	Actioner              apijson.Field
	Amount                apijson.Field
	CardPaymentID         apijson.Field
	Currency              apijson.Field
	DeclinedTransactionID apijson.Field
	DigitalWalletTokenID  apijson.Field
	MerchantAcceptorID    apijson.Field
	MerchantCategoryCode  apijson.Field
	MerchantCity          apijson.Field
	MerchantCountry       apijson.Field
	MerchantDescriptor    apijson.Field
	MerchantPostalCode    apijson.Field
	MerchantState         apijson.Field
	NetworkDetails        apijson.Field
	NetworkIdentifiers    apijson.Field
	NetworkRiskScore      apijson.Field
	PhysicalCardID        apijson.Field
	PresentmentAmount     apijson.Field
	PresentmentCurrency   apijson.Field
	ProcessingCategory    apijson.Field
	RealTimeDecisionID    apijson.Field
	Reason                apijson.Field
	Verification          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type DeclinedTransactionSourceCardDeclineActioner string

const (
	// This object was actioned by the user through a real-time decision.
	DeclinedTransactionSourceCardDeclineActionerUser DeclinedTransactionSourceCardDeclineActioner = "user"
	// This object was actioned by Increase without user intervention.
	DeclinedTransactionSourceCardDeclineActionerIncrease DeclinedTransactionSourceCardDeclineActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	DeclinedTransactionSourceCardDeclineActionerNetwork DeclinedTransactionSourceCardDeclineActioner = "network"
)

func (r DeclinedTransactionSourceCardDeclineActioner) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineActionerUser, DeclinedTransactionSourceCardDeclineActionerIncrease, DeclinedTransactionSourceCardDeclineActionerNetwork:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type DeclinedTransactionSourceCardDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	DeclinedTransactionSourceCardDeclineCurrencyCad DeclinedTransactionSourceCardDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	DeclinedTransactionSourceCardDeclineCurrencyChf DeclinedTransactionSourceCardDeclineCurrency = "CHF"
	// Euro (EUR)
	DeclinedTransactionSourceCardDeclineCurrencyEur DeclinedTransactionSourceCardDeclineCurrency = "EUR"
	// British Pound (GBP)
	DeclinedTransactionSourceCardDeclineCurrencyGbp DeclinedTransactionSourceCardDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	DeclinedTransactionSourceCardDeclineCurrencyJpy DeclinedTransactionSourceCardDeclineCurrency = "JPY"
	// US Dollar (USD)
	DeclinedTransactionSourceCardDeclineCurrencyUsd DeclinedTransactionSourceCardDeclineCurrency = "USD"
)

func (r DeclinedTransactionSourceCardDeclineCurrency) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineCurrencyCad, DeclinedTransactionSourceCardDeclineCurrencyChf, DeclinedTransactionSourceCardDeclineCurrencyEur, DeclinedTransactionSourceCardDeclineCurrencyGbp, DeclinedTransactionSourceCardDeclineCurrencyJpy, DeclinedTransactionSourceCardDeclineCurrencyUsd:
		return true
	}
	return false
}

// Fields specific to the `network`.
type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category DeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa DeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON declinedTransactionSourceCardDeclineNetworkDetailsJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineNetworkDetailsJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineNetworkDetails]
type declinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type DeclinedTransactionSourceCardDeclineNetworkDetailsCategory string

const (
	// Visa
	DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa DeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
)

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type DeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON                    `json:"-"`
}

// declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineNetworkDetailsVisa]
type declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type DeclinedTransactionSourceCardDeclineNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                     `json:"transaction_id,required,nullable"`
	JSON          declinedTransactionSourceCardDeclineNetworkIdentifiersJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineNetworkIdentifiersJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineNetworkIdentifiers]
type declinedTransactionSourceCardDeclineNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type DeclinedTransactionSourceCardDeclineProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	DeclinedTransactionSourceCardDeclineProcessingCategoryAccountFunding DeclinedTransactionSourceCardDeclineProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	DeclinedTransactionSourceCardDeclineProcessingCategoryAutomaticFuelDispenser DeclinedTransactionSourceCardDeclineProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	DeclinedTransactionSourceCardDeclineProcessingCategoryBillPayment DeclinedTransactionSourceCardDeclineProcessingCategory = "bill_payment"
	// A regular purchase.
	DeclinedTransactionSourceCardDeclineProcessingCategoryPurchase DeclinedTransactionSourceCardDeclineProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	DeclinedTransactionSourceCardDeclineProcessingCategoryQuasiCash DeclinedTransactionSourceCardDeclineProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	DeclinedTransactionSourceCardDeclineProcessingCategoryRefund DeclinedTransactionSourceCardDeclineProcessingCategory = "refund"
)

func (r DeclinedTransactionSourceCardDeclineProcessingCategory) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineProcessingCategoryAccountFunding, DeclinedTransactionSourceCardDeclineProcessingCategoryAutomaticFuelDispenser, DeclinedTransactionSourceCardDeclineProcessingCategoryBillPayment, DeclinedTransactionSourceCardDeclineProcessingCategoryPurchase, DeclinedTransactionSourceCardDeclineProcessingCategoryQuasiCash, DeclinedTransactionSourceCardDeclineProcessingCategoryRefund:
		return true
	}
	return false
}

// Why the transaction was declined.
type DeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	DeclinedTransactionSourceCardDeclineReasonCardNotActive DeclinedTransactionSourceCardDeclineReason = "card_not_active"
	// The Physical Card was not active.
	DeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive DeclinedTransactionSourceCardDeclineReason = "physical_card_not_active"
	// The account's entity was not active.
	DeclinedTransactionSourceCardDeclineReasonEntityNotActive DeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	// The account was inactive.
	DeclinedTransactionSourceCardDeclineReasonGroupLocked DeclinedTransactionSourceCardDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	DeclinedTransactionSourceCardDeclineReasonInsufficientFunds DeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	DeclinedTransactionSourceCardDeclineReasonCvv2Mismatch DeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	// The given expiration date did not match the card's value. Only applies when a
	// CVV2 is present.
	DeclinedTransactionSourceCardDeclineReasonCardExpirationMismatch DeclinedTransactionSourceCardDeclineReason = "card_expiration_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	DeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed DeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by a Limit.
	DeclinedTransactionSourceCardDeclineReasonBreachesLimit DeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	// Your application declined the transaction via webhook.
	DeclinedTransactionSourceCardDeclineReasonWebhookDeclined DeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	// Your application webhook did not respond without the required timeout.
	DeclinedTransactionSourceCardDeclineReasonWebhookTimedOut DeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	// Declined by stand-in processing.
	DeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing DeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	// The card read had an invalid CVV, dCVV, or authorization request cryptogram.
	DeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard DeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	// The original card authorization for this incremental authorization does not
	// exist.
	DeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization DeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
	// The transaction was suspected to be fraudulent. Please reach out to
	// support@increase.com for more information.
	DeclinedTransactionSourceCardDeclineReasonSuspectedFraud DeclinedTransactionSourceCardDeclineReason = "suspected_fraud"
)

func (r DeclinedTransactionSourceCardDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineReasonCardNotActive, DeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive, DeclinedTransactionSourceCardDeclineReasonEntityNotActive, DeclinedTransactionSourceCardDeclineReasonGroupLocked, DeclinedTransactionSourceCardDeclineReasonInsufficientFunds, DeclinedTransactionSourceCardDeclineReasonCvv2Mismatch, DeclinedTransactionSourceCardDeclineReasonCardExpirationMismatch, DeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed, DeclinedTransactionSourceCardDeclineReasonBreachesLimit, DeclinedTransactionSourceCardDeclineReasonWebhookDeclined, DeclinedTransactionSourceCardDeclineReasonWebhookTimedOut, DeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing, DeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard, DeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization, DeclinedTransactionSourceCardDeclineReasonSuspectedFraud:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type DeclinedTransactionSourceCardDeclineVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode DeclinedTransactionSourceCardDeclineVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress DeclinedTransactionSourceCardDeclineVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              declinedTransactionSourceCardDeclineVerificationJSON              `json:"-"`
}

// declinedTransactionSourceCardDeclineVerificationJSON contains the JSON metadata
// for the struct [DeclinedTransactionSourceCardDeclineVerification]
type declinedTransactionSourceCardDeclineVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type DeclinedTransactionSourceCardDeclineVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   declinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON   `json:"-"`
}

// declinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON
// contains the JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineVerificationCardVerificationCode]
type declinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNotChecked DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultMatch DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNoMatch DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "no_match"
)

func (r DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNotChecked, DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultMatch, DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type DeclinedTransactionSourceCardDeclineVerificationCardholderAddress struct {
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
	Result DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult `json:"result,required"`
	JSON   declinedTransactionSourceCardDeclineVerificationCardholderAddressJSON   `json:"-"`
}

// declinedTransactionSourceCardDeclineVerificationCardholderAddressJSON contains
// the JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineVerificationCardholderAddress]
type declinedTransactionSourceCardDeclineVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "no_match"
)

func (r DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type DeclinedTransactionSourceCategory string

const (
	// ACH Decline: details will be under the `ach_decline` object.
	DeclinedTransactionSourceCategoryACHDecline DeclinedTransactionSourceCategory = "ach_decline"
	// Card Decline: details will be under the `card_decline` object.
	DeclinedTransactionSourceCategoryCardDecline DeclinedTransactionSourceCategory = "card_decline"
	// Check Decline: details will be under the `check_decline` object.
	DeclinedTransactionSourceCategoryCheckDecline DeclinedTransactionSourceCategory = "check_decline"
	// Inbound Real-Time Payments Transfer Decline: details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	DeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline DeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	DeclinedTransactionSourceCategoryWireDecline DeclinedTransactionSourceCategory = "wire_decline"
	// Check Deposit Rejection: details will be under the `check_deposit_rejection`
	// object.
	DeclinedTransactionSourceCategoryCheckDepositRejection DeclinedTransactionSourceCategory = "check_deposit_rejection"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	DeclinedTransactionSourceCategoryOther DeclinedTransactionSourceCategory = "other"
)

func (r DeclinedTransactionSourceCategory) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCategoryACHDecline, DeclinedTransactionSourceCategoryCardDecline, DeclinedTransactionSourceCategoryCheckDecline, DeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline, DeclinedTransactionSourceCategoryWireDecline, DeclinedTransactionSourceCategoryCheckDepositRejection, DeclinedTransactionSourceCategoryOther:
		return true
	}
	return false
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type DeclinedTransactionSourceCheckDecline struct {
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
	Reason DeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   declinedTransactionSourceCheckDeclineJSON   `json:"-"`
}

// declinedTransactionSourceCheckDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceCheckDecline]
type declinedTransactionSourceCheckDeclineJSON struct {
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

func (r *DeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCheckDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the check was declined.
type DeclinedTransactionSourceCheckDeclineReason string

const (
	// The account number is disabled.
	DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled DeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The account number is canceled.
	DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled DeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The deposited check was altered or fictitious.
	DeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious DeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
	// The transaction would cause a limit to be exceeded.
	DeclinedTransactionSourceCheckDeclineReasonBreachesLimit DeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The check was not endorsed by the payee.
	DeclinedTransactionSourceCheckDeclineReasonEndorsementIrregular DeclinedTransactionSourceCheckDeclineReason = "endorsement_irregular"
	// The account's entity is not active.
	DeclinedTransactionSourceCheckDeclineReasonEntityNotActive DeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	DeclinedTransactionSourceCheckDeclineReasonGroupLocked DeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	DeclinedTransactionSourceCheckDeclineReasonInsufficientFunds DeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Stop payment requested for this check.
	DeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested DeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// The check was a duplicate deposit.
	DeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment DeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The check was not authorized.
	DeclinedTransactionSourceCheckDeclineReasonNotAuthorized DeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The amount the receiving bank is attempting to deposit does not match the amount
	// on the check.
	DeclinedTransactionSourceCheckDeclineReasonAmountMismatch DeclinedTransactionSourceCheckDeclineReason = "amount_mismatch"
	// The check attempting to be deposited does not belong to Increase.
	DeclinedTransactionSourceCheckDeclineReasonNotOurItem DeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	// The account number on the check does not exist at Increase.
	DeclinedTransactionSourceCheckDeclineReasonNoAccountNumberFound DeclinedTransactionSourceCheckDeclineReason = "no_account_number_found"
	// The check is not readable. Please refer to the image.
	DeclinedTransactionSourceCheckDeclineReasonReferToImage DeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	// The check cannot be processed. This is rare: please contact support.
	DeclinedTransactionSourceCheckDeclineReasonUnableToProcess DeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	// Your integration declined this check via the API.
	DeclinedTransactionSourceCheckDeclineReasonUserInitiated DeclinedTransactionSourceCheckDeclineReason = "user_initiated"
)

func (r DeclinedTransactionSourceCheckDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled, DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled, DeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious, DeclinedTransactionSourceCheckDeclineReasonBreachesLimit, DeclinedTransactionSourceCheckDeclineReasonEndorsementIrregular, DeclinedTransactionSourceCheckDeclineReasonEntityNotActive, DeclinedTransactionSourceCheckDeclineReasonGroupLocked, DeclinedTransactionSourceCheckDeclineReasonInsufficientFunds, DeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested, DeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment, DeclinedTransactionSourceCheckDeclineReasonNotAuthorized, DeclinedTransactionSourceCheckDeclineReasonAmountMismatch, DeclinedTransactionSourceCheckDeclineReasonNotOurItem, DeclinedTransactionSourceCheckDeclineReasonNoAccountNumberFound, DeclinedTransactionSourceCheckDeclineReasonReferToImage, DeclinedTransactionSourceCheckDeclineReasonUnableToProcess, DeclinedTransactionSourceCheckDeclineReasonUserInitiated:
		return true
	}
	return false
}

// A Check Deposit Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_rejection`.
type DeclinedTransactionSourceCheckDepositRejection struct {
	// The rejected amount in the minor unit of check's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was rejected.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency DeclinedTransactionSourceCheckDepositRejectionCurrency `json:"currency,required"`
	// The identifier of the associated declined transaction.
	DeclinedTransactionID string `json:"declined_transaction_id,required"`
	// Why the check deposit was rejected.
	Reason DeclinedTransactionSourceCheckDepositRejectionReason `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt time.Time                                          `json:"rejected_at,required" format:"date-time"`
	JSON       declinedTransactionSourceCheckDepositRejectionJSON `json:"-"`
}

// declinedTransactionSourceCheckDepositRejectionJSON contains the JSON metadata
// for the struct [DeclinedTransactionSourceCheckDepositRejection]
type declinedTransactionSourceCheckDepositRejectionJSON struct {
	Amount                apijson.Field
	CheckDepositID        apijson.Field
	Currency              apijson.Field
	DeclinedTransactionID apijson.Field
	Reason                apijson.Field
	RejectedAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCheckDepositRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCheckDepositRejectionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type DeclinedTransactionSourceCheckDepositRejectionCurrency string

const (
	// Canadian Dollar (CAD)
	DeclinedTransactionSourceCheckDepositRejectionCurrencyCad DeclinedTransactionSourceCheckDepositRejectionCurrency = "CAD"
	// Swiss Franc (CHF)
	DeclinedTransactionSourceCheckDepositRejectionCurrencyChf DeclinedTransactionSourceCheckDepositRejectionCurrency = "CHF"
	// Euro (EUR)
	DeclinedTransactionSourceCheckDepositRejectionCurrencyEur DeclinedTransactionSourceCheckDepositRejectionCurrency = "EUR"
	// British Pound (GBP)
	DeclinedTransactionSourceCheckDepositRejectionCurrencyGbp DeclinedTransactionSourceCheckDepositRejectionCurrency = "GBP"
	// Japanese Yen (JPY)
	DeclinedTransactionSourceCheckDepositRejectionCurrencyJpy DeclinedTransactionSourceCheckDepositRejectionCurrency = "JPY"
	// US Dollar (USD)
	DeclinedTransactionSourceCheckDepositRejectionCurrencyUsd DeclinedTransactionSourceCheckDepositRejectionCurrency = "USD"
)

func (r DeclinedTransactionSourceCheckDepositRejectionCurrency) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCheckDepositRejectionCurrencyCad, DeclinedTransactionSourceCheckDepositRejectionCurrencyChf, DeclinedTransactionSourceCheckDepositRejectionCurrencyEur, DeclinedTransactionSourceCheckDepositRejectionCurrencyGbp, DeclinedTransactionSourceCheckDepositRejectionCurrencyJpy, DeclinedTransactionSourceCheckDepositRejectionCurrencyUsd:
		return true
	}
	return false
}

// Why the check deposit was rejected.
type DeclinedTransactionSourceCheckDepositRejectionReason string

const (
	// The check's image is incomplete.
	DeclinedTransactionSourceCheckDepositRejectionReasonIncompleteImage DeclinedTransactionSourceCheckDepositRejectionReason = "incomplete_image"
	// This is a duplicate check submission.
	DeclinedTransactionSourceCheckDepositRejectionReasonDuplicate DeclinedTransactionSourceCheckDepositRejectionReason = "duplicate"
	// This check has poor image quality.
	DeclinedTransactionSourceCheckDepositRejectionReasonPoorImageQuality DeclinedTransactionSourceCheckDepositRejectionReason = "poor_image_quality"
	// The check was deposited with the incorrect amount.
	DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectAmount DeclinedTransactionSourceCheckDepositRejectionReason = "incorrect_amount"
	// The check is made out to someone other than the account holder.
	DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectRecipient DeclinedTransactionSourceCheckDepositRejectionReason = "incorrect_recipient"
	// This check was not eligible for mobile deposit.
	DeclinedTransactionSourceCheckDepositRejectionReasonNotEligibleForMobileDeposit DeclinedTransactionSourceCheckDepositRejectionReason = "not_eligible_for_mobile_deposit"
	// This check is missing at least one required field.
	DeclinedTransactionSourceCheckDepositRejectionReasonMissingRequiredDataElements DeclinedTransactionSourceCheckDepositRejectionReason = "missing_required_data_elements"
	// This check is suspected to be fraudulent.
	DeclinedTransactionSourceCheckDepositRejectionReasonSuspectedFraud DeclinedTransactionSourceCheckDepositRejectionReason = "suspected_fraud"
	// This check's deposit window has expired.
	DeclinedTransactionSourceCheckDepositRejectionReasonDepositWindowExpired DeclinedTransactionSourceCheckDepositRejectionReason = "deposit_window_expired"
	// The check was rejected for an unknown reason.
	DeclinedTransactionSourceCheckDepositRejectionReasonUnknown DeclinedTransactionSourceCheckDepositRejectionReason = "unknown"
)

func (r DeclinedTransactionSourceCheckDepositRejectionReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCheckDepositRejectionReasonIncompleteImage, DeclinedTransactionSourceCheckDepositRejectionReasonDuplicate, DeclinedTransactionSourceCheckDepositRejectionReasonPoorImageQuality, DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectAmount, DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectRecipient, DeclinedTransactionSourceCheckDepositRejectionReasonNotEligibleForMobileDeposit, DeclinedTransactionSourceCheckDepositRejectionReasonMissingRequiredDataElements, DeclinedTransactionSourceCheckDepositRejectionReasonSuspectedFraud, DeclinedTransactionSourceCheckDepositRejectionReasonDepositWindowExpired, DeclinedTransactionSourceCheckDepositRejectionReasonUnknown:
		return true
	}
	return false
}

// An Inbound Real-Time Payments Transfer Decline object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real-Time Payments
	// transfer.
	Currency DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Why the transfer was declined.
	Reason DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string                                                              `json:"transfer_id,required"`
	JSON       declinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON `json:"-"`
}

// declinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline]
type declinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	Reason                    apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	TransferID                apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real-Time Payments
// transfer.
type DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	// Euro (EUR)
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	// British Pound (GBP)
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	// US Dollar (USD)
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

func (r DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd:
		return true
	}
	return false
}

// Why the transfer was declined.
type DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	// The account number is canceled.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	// The account number is disabled.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	// Your account is restricted.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	// Your account is inactive.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	// The account's entity is not active.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	// Your account is not enabled to receive Real-Time Payments transfers.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

func (r DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled:
		return true
	}
	return false
}

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type DeclinedTransactionSourceWireDecline struct {
	// The identifier of the Inbound Wire Transfer that was declined.
	InboundWireTransferID string `json:"inbound_wire_transfer_id,required"`
	// Why the wire transfer was declined.
	Reason DeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	JSON   declinedTransactionSourceWireDeclineJSON   `json:"-"`
}

// declinedTransactionSourceWireDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceWireDecline]
type declinedTransactionSourceWireDeclineJSON struct {
	InboundWireTransferID apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceWireDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the wire transfer was declined.
type DeclinedTransactionSourceWireDeclineReason string

const (
	// The account number is canceled.
	DeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled DeclinedTransactionSourceWireDeclineReason = "account_number_canceled"
	// The account number is disabled.
	DeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled DeclinedTransactionSourceWireDeclineReason = "account_number_disabled"
	// The account's entity is not active.
	DeclinedTransactionSourceWireDeclineReasonEntityNotActive DeclinedTransactionSourceWireDeclineReason = "entity_not_active"
	// Your account is inactive.
	DeclinedTransactionSourceWireDeclineReasonGroupLocked DeclinedTransactionSourceWireDeclineReason = "group_locked"
	// The beneficiary account number does not exist.
	DeclinedTransactionSourceWireDeclineReasonNoAccountNumber DeclinedTransactionSourceWireDeclineReason = "no_account_number"
	// The transaction is not allowed per Increase's terms.
	DeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed DeclinedTransactionSourceWireDeclineReason = "transaction_not_allowed"
)

func (r DeclinedTransactionSourceWireDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled, DeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled, DeclinedTransactionSourceWireDeclineReasonEntityNotActive, DeclinedTransactionSourceWireDeclineReasonGroupLocked, DeclinedTransactionSourceWireDeclineReasonNoAccountNumber, DeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
type DeclinedTransactionType string

const (
	DeclinedTransactionTypeDeclinedTransaction DeclinedTransactionType = "declined_transaction"
)

func (r DeclinedTransactionType) IsKnown() bool {
	switch r {
	case DeclinedTransactionTypeDeclinedTransaction:
		return true
	}
	return false
}

type DeclinedTransactionListParams struct {
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID param.Field[string]                                 `query:"account_id"`
	Category  param.Field[DeclinedTransactionListParamsCategory]  `query:"category"`
	CreatedAt param.Field[DeclinedTransactionListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID param.Field[string] `query:"route_id"`
}

// URLQuery serializes [DeclinedTransactionListParams]'s query parameters as
// `url.Values`.
func (r DeclinedTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeclinedTransactionListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]DeclinedTransactionListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [DeclinedTransactionListParamsCategory]'s query parameters
// as `url.Values`.
func (r DeclinedTransactionListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeclinedTransactionListParamsCategoryIn string

const (
	// ACH Decline: details will be under the `ach_decline` object.
	DeclinedTransactionListParamsCategoryInACHDecline DeclinedTransactionListParamsCategoryIn = "ach_decline"
	// Card Decline: details will be under the `card_decline` object.
	DeclinedTransactionListParamsCategoryInCardDecline DeclinedTransactionListParamsCategoryIn = "card_decline"
	// Check Decline: details will be under the `check_decline` object.
	DeclinedTransactionListParamsCategoryInCheckDecline DeclinedTransactionListParamsCategoryIn = "check_decline"
	// Inbound Real-Time Payments Transfer Decline: details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	DeclinedTransactionListParamsCategoryInInboundRealTimePaymentsTransferDecline DeclinedTransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	DeclinedTransactionListParamsCategoryInWireDecline DeclinedTransactionListParamsCategoryIn = "wire_decline"
	// Check Deposit Rejection: details will be under the `check_deposit_rejection`
	// object.
	DeclinedTransactionListParamsCategoryInCheckDepositRejection DeclinedTransactionListParamsCategoryIn = "check_deposit_rejection"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	DeclinedTransactionListParamsCategoryInOther DeclinedTransactionListParamsCategoryIn = "other"
)

func (r DeclinedTransactionListParamsCategoryIn) IsKnown() bool {
	switch r {
	case DeclinedTransactionListParamsCategoryInACHDecline, DeclinedTransactionListParamsCategoryInCardDecline, DeclinedTransactionListParamsCategoryInCheckDecline, DeclinedTransactionListParamsCategoryInInboundRealTimePaymentsTransferDecline, DeclinedTransactionListParamsCategoryInWireDecline, DeclinedTransactionListParamsCategoryInCheckDepositRejection, DeclinedTransactionListParamsCategoryInOther:
		return true
	}
	return false
}

type DeclinedTransactionListParamsCreatedAt struct {
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

// URLQuery serializes [DeclinedTransactionListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r DeclinedTransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
