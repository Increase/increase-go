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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	Type        DeclinedTransactionType `json:"type,required"`
	ExtraFields map[string]interface{}  `json:"-,extras"`
	JSON        declinedTransactionJSON `json:"-"`
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
	DeclinedTransactionCurrencyCad DeclinedTransactionCurrency = "CAD"
	DeclinedTransactionCurrencyChf DeclinedTransactionCurrency = "CHF"
	DeclinedTransactionCurrencyEur DeclinedTransactionCurrency = "EUR"
	DeclinedTransactionCurrencyGbp DeclinedTransactionCurrency = "GBP"
	DeclinedTransactionCurrencyJpy DeclinedTransactionCurrency = "JPY"
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
	DeclinedTransactionRouteTypeAccountNumber DeclinedTransactionRouteType = "account_number"
	DeclinedTransactionRouteTypeCard          DeclinedTransactionRouteType = "card"
	DeclinedTransactionRouteTypeLockbox       DeclinedTransactionRouteType = "lockbox"
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
	// An Inbound FedNow Transfer Decline object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_fednow_transfer_decline`.
	InboundFednowTransferDecline DeclinedTransactionSourceInboundFednowTransferDecline `json:"inbound_fednow_transfer_decline,required,nullable"`
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// If the category of this Transaction source is equal to `other`, this field will
	// contain an empty object, otherwise it will contain null.
	Other DeclinedTransactionSourceOther `json:"other,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline DeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	ExtraFields map[string]interface{}               `json:"-,extras"`
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
	InboundFednowTransferDecline           apijson.Field
	InboundRealTimePaymentsTransferDecline apijson.Field
	Other                                  apijson.Field
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
	// The declined amount in USD cents.
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
	Type        DeclinedTransactionSourceACHDeclineType `json:"type,required"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
	JSON        declinedTransactionSourceACHDeclineJSON `json:"-"`
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
	DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled                                            DeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled                                            DeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	DeclinedTransactionSourceACHDeclineReasonBreachesLimit                                               DeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	DeclinedTransactionSourceACHDeclineReasonEntityNotActive                                             DeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	DeclinedTransactionSourceACHDeclineReasonGroupLocked                                                 DeclinedTransactionSourceACHDeclineReason = "group_locked"
	DeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed                                       DeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	DeclinedTransactionSourceACHDeclineReasonUserInitiated                                               DeclinedTransactionSourceACHDeclineReason = "user_initiated"
	DeclinedTransactionSourceACHDeclineReasonInsufficientFunds                                           DeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceACHDeclineReasonReturnedPerOdfiRequest                                      DeclinedTransactionSourceACHDeclineReason = "returned_per_odfi_request"
	DeclinedTransactionSourceACHDeclineReasonAuthorizationRevokedByCustomer                              DeclinedTransactionSourceACHDeclineReason = "authorization_revoked_by_customer"
	DeclinedTransactionSourceACHDeclineReasonPaymentStopped                                              DeclinedTransactionSourceACHDeclineReason = "payment_stopped"
	DeclinedTransactionSourceACHDeclineReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   DeclinedTransactionSourceACHDeclineReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	DeclinedTransactionSourceACHDeclineReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity DeclinedTransactionSourceACHDeclineReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	DeclinedTransactionSourceACHDeclineReasonBeneficiaryOrAccountHolderDeceased                          DeclinedTransactionSourceACHDeclineReason = "beneficiary_or_account_holder_deceased"
	DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver                                DeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	DeclinedTransactionSourceACHDeclineReasonDuplicateEntry                                              DeclinedTransactionSourceACHDeclineReason = "duplicate_entry"
	DeclinedTransactionSourceACHDeclineReasonCorporateCustomerAdvisedNotAuthorized                       DeclinedTransactionSourceACHDeclineReason = "corporate_customer_advised_not_authorized"
)

func (r DeclinedTransactionSourceACHDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled, DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled, DeclinedTransactionSourceACHDeclineReasonBreachesLimit, DeclinedTransactionSourceACHDeclineReasonEntityNotActive, DeclinedTransactionSourceACHDeclineReasonGroupLocked, DeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed, DeclinedTransactionSourceACHDeclineReasonUserInitiated, DeclinedTransactionSourceACHDeclineReasonInsufficientFunds, DeclinedTransactionSourceACHDeclineReasonReturnedPerOdfiRequest, DeclinedTransactionSourceACHDeclineReasonAuthorizationRevokedByCustomer, DeclinedTransactionSourceACHDeclineReasonPaymentStopped, DeclinedTransactionSourceACHDeclineReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, DeclinedTransactionSourceACHDeclineReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, DeclinedTransactionSourceACHDeclineReasonBeneficiaryOrAccountHolderDeceased, DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver, DeclinedTransactionSourceACHDeclineReasonDuplicateEntry, DeclinedTransactionSourceACHDeclineReasonCorporateCustomerAdvisedNotAuthorized:
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
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts DeclinedTransactionSourceCardDeclineAdditionalAmounts `json:"additional_amounts,required"`
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
	// The direction describes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction DeclinedTransactionSourceCardDeclineDirection `json:"direction,required"`
	// The identifier of the card authorization this request attempted to incrementally
	// authorize.
	IncrementedCardAuthorizationID string `json:"incremented_card_authorization_id,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
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
	// Authorization risk score, from 0 to 99, where 99 is the riskiest. For Pulse the
	// score is from 0 to 999, where 999 is the riskiest.
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
	// This is present if a specific decline reason was given in the real-time
	// decision.
	RealTimeDecisionReason DeclinedTransactionSourceCardDeclineRealTimeDecisionReason `json:"real_time_decision_reason,required,nullable"`
	// Why the transaction was declined.
	Reason DeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
	// Fields related to verification of cardholder-provided values.
	Verification DeclinedTransactionSourceCardDeclineVerification `json:"verification,required"`
	ExtraFields  map[string]interface{}                           `json:"-,extras"`
	JSON         declinedTransactionSourceCardDeclineJSON         `json:"-"`
}

// declinedTransactionSourceCardDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceCardDecline]
type declinedTransactionSourceCardDeclineJSON struct {
	ID                             apijson.Field
	Actioner                       apijson.Field
	AdditionalAmounts              apijson.Field
	Amount                         apijson.Field
	CardPaymentID                  apijson.Field
	Currency                       apijson.Field
	DeclinedTransactionID          apijson.Field
	DigitalWalletTokenID           apijson.Field
	Direction                      apijson.Field
	IncrementedCardAuthorizationID apijson.Field
	MerchantAcceptorID             apijson.Field
	MerchantCategoryCode           apijson.Field
	MerchantCity                   apijson.Field
	MerchantCountry                apijson.Field
	MerchantDescriptor             apijson.Field
	MerchantPostalCode             apijson.Field
	MerchantState                  apijson.Field
	NetworkDetails                 apijson.Field
	NetworkIdentifiers             apijson.Field
	NetworkRiskScore               apijson.Field
	PhysicalCardID                 apijson.Field
	PresentmentAmount              apijson.Field
	PresentmentCurrency            apijson.Field
	ProcessingCategory             apijson.Field
	RealTimeDecisionID             apijson.Field
	RealTimeDecisionReason         apijson.Field
	Reason                         apijson.Field
	TerminalID                     apijson.Field
	Verification                   apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
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
	DeclinedTransactionSourceCardDeclineActionerUser     DeclinedTransactionSourceCardDeclineActioner = "user"
	DeclinedTransactionSourceCardDeclineActionerIncrease DeclinedTransactionSourceCardDeclineActioner = "increase"
	DeclinedTransactionSourceCardDeclineActionerNetwork  DeclinedTransactionSourceCardDeclineActioner = "network"
)

func (r DeclinedTransactionSourceCardDeclineActioner) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineActionerUser, DeclinedTransactionSourceCardDeclineActionerIncrease, DeclinedTransactionSourceCardDeclineActionerNetwork:
		return true
	}
	return false
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type DeclinedTransactionSourceCardDeclineAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic DeclinedTransactionSourceCardDeclineAdditionalAmountsClinic `json:"clinic,required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental DeclinedTransactionSourceCardDeclineAdditionalAmountsDental `json:"dental,required,nullable"`
	// The original pre-authorized amount.
	Original DeclinedTransactionSourceCardDeclineAdditionalAmountsOriginal `json:"original,required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription DeclinedTransactionSourceCardDeclineAdditionalAmountsPrescription `json:"prescription,required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge DeclinedTransactionSourceCardDeclineAdditionalAmountsSurcharge `json:"surcharge,required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulative `json:"total_cumulative,required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcare `json:"total_healthcare,required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit DeclinedTransactionSourceCardDeclineAdditionalAmountsTransit `json:"transit,required,nullable"`
	// An unknown additional amount.
	Unknown DeclinedTransactionSourceCardDeclineAdditionalAmountsUnknown `json:"unknown,required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision DeclinedTransactionSourceCardDeclineAdditionalAmountsVision `json:"vision,required,nullable"`
	JSON   declinedTransactionSourceCardDeclineAdditionalAmountsJSON   `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineAdditionalAmounts]
type declinedTransactionSourceCardDeclineAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
	Original        apijson.Field
	Prescription    apijson.Field
	Surcharge       apijson.Field
	TotalCumulative apijson.Field
	TotalHealthcare apijson.Field
	Transit         apijson.Field
	Unknown         apijson.Field
	Vision          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsClinicJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsClinicJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsClinic]
type declinedTransactionSourceCardDeclineAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsDentalJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsDentalJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsDental]
type declinedTransactionSourceCardDeclineAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The original pre-authorized amount.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsOriginal struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                            `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsOriginalJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsOriginalJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsOriginal]
type declinedTransactionSourceCardDeclineAdditionalAmountsOriginalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsOriginal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsOriginalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsPrescriptionJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsPrescriptionJSON contains
// the JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsPrescription]
type declinedTransactionSourceCardDeclineAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsSurchargeJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsSurchargeJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsSurcharge]
type declinedTransactionSourceCardDeclineAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                   `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulativeJSON
// contains the JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulative]
type declinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                   `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcareJSON
// contains the JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcare]
type declinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                           `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsTransitJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsTransitJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsTransit]
type declinedTransactionSourceCardDeclineAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                           `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsUnknownJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsUnknownJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsUnknown]
type declinedTransactionSourceCardDeclineAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type DeclinedTransactionSourceCardDeclineAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     declinedTransactionSourceCardDeclineAdditionalAmountsVisionJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineAdditionalAmountsVisionJSON contains the
// JSON metadata for the struct
// [DeclinedTransactionSourceCardDeclineAdditionalAmountsVision]
type declinedTransactionSourceCardDeclineAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type DeclinedTransactionSourceCardDeclineCurrency string

const (
	DeclinedTransactionSourceCardDeclineCurrencyCad DeclinedTransactionSourceCardDeclineCurrency = "CAD"
	DeclinedTransactionSourceCardDeclineCurrencyChf DeclinedTransactionSourceCardDeclineCurrency = "CHF"
	DeclinedTransactionSourceCardDeclineCurrencyEur DeclinedTransactionSourceCardDeclineCurrency = "EUR"
	DeclinedTransactionSourceCardDeclineCurrencyGbp DeclinedTransactionSourceCardDeclineCurrency = "GBP"
	DeclinedTransactionSourceCardDeclineCurrencyJpy DeclinedTransactionSourceCardDeclineCurrency = "JPY"
	DeclinedTransactionSourceCardDeclineCurrencyUsd DeclinedTransactionSourceCardDeclineCurrency = "USD"
)

func (r DeclinedTransactionSourceCardDeclineCurrency) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineCurrencyCad, DeclinedTransactionSourceCardDeclineCurrencyChf, DeclinedTransactionSourceCardDeclineCurrencyEur, DeclinedTransactionSourceCardDeclineCurrencyGbp, DeclinedTransactionSourceCardDeclineCurrencyJpy, DeclinedTransactionSourceCardDeclineCurrencyUsd:
		return true
	}
	return false
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type DeclinedTransactionSourceCardDeclineDirection string

const (
	DeclinedTransactionSourceCardDeclineDirectionSettlement DeclinedTransactionSourceCardDeclineDirection = "settlement"
	DeclinedTransactionSourceCardDeclineDirectionRefund     DeclinedTransactionSourceCardDeclineDirection = "refund"
)

func (r DeclinedTransactionSourceCardDeclineDirection) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineDirectionSettlement, DeclinedTransactionSourceCardDeclineDirectionRefund:
		return true
	}
	return false
}

// Fields specific to the `network`.
type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category DeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `pulse` network.
	Pulse DeclinedTransactionSourceCardDeclineNetworkDetailsPulse `json:"pulse,required,nullable"`
	// Fields specific to the `visa` network.
	Visa DeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON declinedTransactionSourceCardDeclineNetworkDetailsJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineNetworkDetailsJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineNetworkDetails]
type declinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Category    apijson.Field
	Pulse       apijson.Field
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
	DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa  DeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
	DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryPulse DeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "pulse"
)

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa, DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryPulse:
		return true
	}
	return false
}

// Fields specific to the `pulse` network.
type DeclinedTransactionSourceCardDeclineNetworkDetailsPulse struct {
	JSON declinedTransactionSourceCardDeclineNetworkDetailsPulseJSON `json:"-"`
}

// declinedTransactionSourceCardDeclineNetworkDetailsPulseJSON contains the JSON
// metadata for the struct
// [DeclinedTransactionSourceCardDeclineNetworkDetailsPulse]
type declinedTransactionSourceCardDeclineNetworkDetailsPulseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsPulse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceCardDeclineNetworkDetailsPulseJSON) RawJSON() string {
	return r.raw
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
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason,required,nullable"`
	JSON                    declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON                    `json:"-"`
}

// declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineNetworkDetailsVisa]
type declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
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
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown                    DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual                     DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless                DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason string

const (
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonIssuerError                                              DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonOther                                                    DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonIssuerError, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonInternalVisaError, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, DeclinedTransactionSourceCardDeclineNetworkDetailsVisaStandInProcessingReasonOther:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type DeclinedTransactionSourceCardDeclineNetworkIdentifiers struct {
	// The randomly generated 6-character Authorization Identification Response code
	// sent back to the acquirer in an approved response.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response,required,nullable"`
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
	AuthorizationIdentificationResponse apijson.Field
	RetrievalReferenceNumber            apijson.Field
	TraceNumber                         apijson.Field
	TransactionID                       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
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
	DeclinedTransactionSourceCardDeclineProcessingCategoryAccountFunding         DeclinedTransactionSourceCardDeclineProcessingCategory = "account_funding"
	DeclinedTransactionSourceCardDeclineProcessingCategoryAutomaticFuelDispenser DeclinedTransactionSourceCardDeclineProcessingCategory = "automatic_fuel_dispenser"
	DeclinedTransactionSourceCardDeclineProcessingCategoryBillPayment            DeclinedTransactionSourceCardDeclineProcessingCategory = "bill_payment"
	DeclinedTransactionSourceCardDeclineProcessingCategoryOriginalCredit         DeclinedTransactionSourceCardDeclineProcessingCategory = "original_credit"
	DeclinedTransactionSourceCardDeclineProcessingCategoryPurchase               DeclinedTransactionSourceCardDeclineProcessingCategory = "purchase"
	DeclinedTransactionSourceCardDeclineProcessingCategoryQuasiCash              DeclinedTransactionSourceCardDeclineProcessingCategory = "quasi_cash"
	DeclinedTransactionSourceCardDeclineProcessingCategoryRefund                 DeclinedTransactionSourceCardDeclineProcessingCategory = "refund"
	DeclinedTransactionSourceCardDeclineProcessingCategoryCashDisbursement       DeclinedTransactionSourceCardDeclineProcessingCategory = "cash_disbursement"
	DeclinedTransactionSourceCardDeclineProcessingCategoryUnknown                DeclinedTransactionSourceCardDeclineProcessingCategory = "unknown"
)

func (r DeclinedTransactionSourceCardDeclineProcessingCategory) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineProcessingCategoryAccountFunding, DeclinedTransactionSourceCardDeclineProcessingCategoryAutomaticFuelDispenser, DeclinedTransactionSourceCardDeclineProcessingCategoryBillPayment, DeclinedTransactionSourceCardDeclineProcessingCategoryOriginalCredit, DeclinedTransactionSourceCardDeclineProcessingCategoryPurchase, DeclinedTransactionSourceCardDeclineProcessingCategoryQuasiCash, DeclinedTransactionSourceCardDeclineProcessingCategoryRefund, DeclinedTransactionSourceCardDeclineProcessingCategoryCashDisbursement, DeclinedTransactionSourceCardDeclineProcessingCategoryUnknown:
		return true
	}
	return false
}

// This is present if a specific decline reason was given in the real-time
// decision.
type DeclinedTransactionSourceCardDeclineRealTimeDecisionReason string

const (
	DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonInsufficientFunds       DeclinedTransactionSourceCardDeclineRealTimeDecisionReason = "insufficient_funds"
	DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonTransactionNeverAllowed DeclinedTransactionSourceCardDeclineRealTimeDecisionReason = "transaction_never_allowed"
	DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonExceedsApprovalLimit    DeclinedTransactionSourceCardDeclineRealTimeDecisionReason = "exceeds_approval_limit"
	DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonCardTemporarilyDisabled DeclinedTransactionSourceCardDeclineRealTimeDecisionReason = "card_temporarily_disabled"
	DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonSuspectedFraud          DeclinedTransactionSourceCardDeclineRealTimeDecisionReason = "suspected_fraud"
	DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonOther                   DeclinedTransactionSourceCardDeclineRealTimeDecisionReason = "other"
)

func (r DeclinedTransactionSourceCardDeclineRealTimeDecisionReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonInsufficientFunds, DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonTransactionNeverAllowed, DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonExceedsApprovalLimit, DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonCardTemporarilyDisabled, DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonSuspectedFraud, DeclinedTransactionSourceCardDeclineRealTimeDecisionReasonOther:
		return true
	}
	return false
}

// Why the transaction was declined.
type DeclinedTransactionSourceCardDeclineReason string

const (
	DeclinedTransactionSourceCardDeclineReasonAccountClosed                DeclinedTransactionSourceCardDeclineReason = "account_closed"
	DeclinedTransactionSourceCardDeclineReasonCardNotActive                DeclinedTransactionSourceCardDeclineReason = "card_not_active"
	DeclinedTransactionSourceCardDeclineReasonCardCanceled                 DeclinedTransactionSourceCardDeclineReason = "card_canceled"
	DeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive        DeclinedTransactionSourceCardDeclineReason = "physical_card_not_active"
	DeclinedTransactionSourceCardDeclineReasonEntityNotActive              DeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	DeclinedTransactionSourceCardDeclineReasonGroupLocked                  DeclinedTransactionSourceCardDeclineReason = "group_locked"
	DeclinedTransactionSourceCardDeclineReasonInsufficientFunds            DeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                 DeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	DeclinedTransactionSourceCardDeclineReasonPinMismatch                  DeclinedTransactionSourceCardDeclineReason = "pin_mismatch"
	DeclinedTransactionSourceCardDeclineReasonCardExpirationMismatch       DeclinedTransactionSourceCardDeclineReason = "card_expiration_mismatch"
	DeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed        DeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	DeclinedTransactionSourceCardDeclineReasonBreachesLimit                DeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	DeclinedTransactionSourceCardDeclineReasonWebhookDeclined              DeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	DeclinedTransactionSourceCardDeclineReasonWebhookTimedOut              DeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	DeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing  DeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	DeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard          DeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	DeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization DeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
	DeclinedTransactionSourceCardDeclineReasonFailed3DSAuthentication      DeclinedTransactionSourceCardDeclineReason = "failed_3ds_authentication"
	DeclinedTransactionSourceCardDeclineReasonSuspectedCardTesting         DeclinedTransactionSourceCardDeclineReason = "suspected_card_testing"
	DeclinedTransactionSourceCardDeclineReasonSuspectedFraud               DeclinedTransactionSourceCardDeclineReason = "suspected_fraud"
)

func (r DeclinedTransactionSourceCardDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineReasonAccountClosed, DeclinedTransactionSourceCardDeclineReasonCardNotActive, DeclinedTransactionSourceCardDeclineReasonCardCanceled, DeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive, DeclinedTransactionSourceCardDeclineReasonEntityNotActive, DeclinedTransactionSourceCardDeclineReasonGroupLocked, DeclinedTransactionSourceCardDeclineReasonInsufficientFunds, DeclinedTransactionSourceCardDeclineReasonCvv2Mismatch, DeclinedTransactionSourceCardDeclineReasonPinMismatch, DeclinedTransactionSourceCardDeclineReasonCardExpirationMismatch, DeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed, DeclinedTransactionSourceCardDeclineReasonBreachesLimit, DeclinedTransactionSourceCardDeclineReasonWebhookDeclined, DeclinedTransactionSourceCardDeclineReasonWebhookTimedOut, DeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing, DeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard, DeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization, DeclinedTransactionSourceCardDeclineReasonFailed3DSAuthentication, DeclinedTransactionSourceCardDeclineReasonSuspectedCardTesting, DeclinedTransactionSourceCardDeclineReasonSuspectedFraud:
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
	DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNotChecked DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "not_checked"
	DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultMatch      DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "match"
	DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResultNoMatch    DeclinedTransactionSourceCardDeclineVerificationCardVerificationCodeResult = "no_match"
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
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked                       DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "not_checked"
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch                            DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "match"
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch                          DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "no_match"
	DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNotChecked, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultNoMatch, DeclinedTransactionSourceCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type DeclinedTransactionSourceCategory string

const (
	DeclinedTransactionSourceCategoryACHDecline                             DeclinedTransactionSourceCategory = "ach_decline"
	DeclinedTransactionSourceCategoryCardDecline                            DeclinedTransactionSourceCategory = "card_decline"
	DeclinedTransactionSourceCategoryCheckDecline                           DeclinedTransactionSourceCategory = "check_decline"
	DeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline DeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	DeclinedTransactionSourceCategoryInboundFednowTransferDecline           DeclinedTransactionSourceCategory = "inbound_fednow_transfer_decline"
	DeclinedTransactionSourceCategoryWireDecline                            DeclinedTransactionSourceCategory = "wire_decline"
	DeclinedTransactionSourceCategoryCheckDepositRejection                  DeclinedTransactionSourceCategory = "check_deposit_rejection"
	DeclinedTransactionSourceCategoryOther                                  DeclinedTransactionSourceCategory = "other"
)

func (r DeclinedTransactionSourceCategory) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCategoryACHDecline, DeclinedTransactionSourceCategoryCardDecline, DeclinedTransactionSourceCategoryCheckDecline, DeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline, DeclinedTransactionSourceCategoryInboundFednowTransferDecline, DeclinedTransactionSourceCategoryWireDecline, DeclinedTransactionSourceCategoryCheckDepositRejection, DeclinedTransactionSourceCategoryOther:
		return true
	}
	return false
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type DeclinedTransactionSourceCheckDecline struct {
	// The declined amount in USD cents.
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
	Reason      DeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	ExtraFields map[string]interface{}                      `json:"-,extras"`
	JSON        declinedTransactionSourceCheckDeclineJSON   `json:"-"`
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
	DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled     DeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled     DeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	DeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious  DeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
	DeclinedTransactionSourceCheckDeclineReasonBreachesLimit        DeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	DeclinedTransactionSourceCheckDeclineReasonEndorsementIrregular DeclinedTransactionSourceCheckDeclineReason = "endorsement_irregular"
	DeclinedTransactionSourceCheckDeclineReasonEntityNotActive      DeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	DeclinedTransactionSourceCheckDeclineReasonGroupLocked          DeclinedTransactionSourceCheckDeclineReason = "group_locked"
	DeclinedTransactionSourceCheckDeclineReasonInsufficientFunds    DeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested DeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	DeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment DeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	DeclinedTransactionSourceCheckDeclineReasonNotAuthorized        DeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	DeclinedTransactionSourceCheckDeclineReasonAmountMismatch       DeclinedTransactionSourceCheckDeclineReason = "amount_mismatch"
	DeclinedTransactionSourceCheckDeclineReasonNotOurItem           DeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	DeclinedTransactionSourceCheckDeclineReasonNoAccountNumberFound DeclinedTransactionSourceCheckDeclineReason = "no_account_number_found"
	DeclinedTransactionSourceCheckDeclineReasonReferToImage         DeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	DeclinedTransactionSourceCheckDeclineReasonUnableToProcess      DeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	DeclinedTransactionSourceCheckDeclineReasonUnusableImage        DeclinedTransactionSourceCheckDeclineReason = "unusable_image"
	DeclinedTransactionSourceCheckDeclineReasonUserInitiated        DeclinedTransactionSourceCheckDeclineReason = "user_initiated"
)

func (r DeclinedTransactionSourceCheckDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled, DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled, DeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious, DeclinedTransactionSourceCheckDeclineReasonBreachesLimit, DeclinedTransactionSourceCheckDeclineReasonEndorsementIrregular, DeclinedTransactionSourceCheckDeclineReasonEntityNotActive, DeclinedTransactionSourceCheckDeclineReasonGroupLocked, DeclinedTransactionSourceCheckDeclineReasonInsufficientFunds, DeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested, DeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment, DeclinedTransactionSourceCheckDeclineReasonNotAuthorized, DeclinedTransactionSourceCheckDeclineReasonAmountMismatch, DeclinedTransactionSourceCheckDeclineReasonNotOurItem, DeclinedTransactionSourceCheckDeclineReasonNoAccountNumberFound, DeclinedTransactionSourceCheckDeclineReasonReferToImage, DeclinedTransactionSourceCheckDeclineReasonUnableToProcess, DeclinedTransactionSourceCheckDeclineReasonUnusableImage, DeclinedTransactionSourceCheckDeclineReasonUserInitiated:
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
	RejectedAt  time.Time                                          `json:"rejected_at,required" format:"date-time"`
	ExtraFields map[string]interface{}                             `json:"-,extras"`
	JSON        declinedTransactionSourceCheckDepositRejectionJSON `json:"-"`
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
	DeclinedTransactionSourceCheckDepositRejectionCurrencyCad DeclinedTransactionSourceCheckDepositRejectionCurrency = "CAD"
	DeclinedTransactionSourceCheckDepositRejectionCurrencyChf DeclinedTransactionSourceCheckDepositRejectionCurrency = "CHF"
	DeclinedTransactionSourceCheckDepositRejectionCurrencyEur DeclinedTransactionSourceCheckDepositRejectionCurrency = "EUR"
	DeclinedTransactionSourceCheckDepositRejectionCurrencyGbp DeclinedTransactionSourceCheckDepositRejectionCurrency = "GBP"
	DeclinedTransactionSourceCheckDepositRejectionCurrencyJpy DeclinedTransactionSourceCheckDepositRejectionCurrency = "JPY"
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
	DeclinedTransactionSourceCheckDepositRejectionReasonIncompleteImage             DeclinedTransactionSourceCheckDepositRejectionReason = "incomplete_image"
	DeclinedTransactionSourceCheckDepositRejectionReasonDuplicate                   DeclinedTransactionSourceCheckDepositRejectionReason = "duplicate"
	DeclinedTransactionSourceCheckDepositRejectionReasonPoorImageQuality            DeclinedTransactionSourceCheckDepositRejectionReason = "poor_image_quality"
	DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectAmount             DeclinedTransactionSourceCheckDepositRejectionReason = "incorrect_amount"
	DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectRecipient          DeclinedTransactionSourceCheckDepositRejectionReason = "incorrect_recipient"
	DeclinedTransactionSourceCheckDepositRejectionReasonNotEligibleForMobileDeposit DeclinedTransactionSourceCheckDepositRejectionReason = "not_eligible_for_mobile_deposit"
	DeclinedTransactionSourceCheckDepositRejectionReasonMissingRequiredDataElements DeclinedTransactionSourceCheckDepositRejectionReason = "missing_required_data_elements"
	DeclinedTransactionSourceCheckDepositRejectionReasonSuspectedFraud              DeclinedTransactionSourceCheckDepositRejectionReason = "suspected_fraud"
	DeclinedTransactionSourceCheckDepositRejectionReasonDepositWindowExpired        DeclinedTransactionSourceCheckDepositRejectionReason = "deposit_window_expired"
	DeclinedTransactionSourceCheckDepositRejectionReasonRequestedByUser             DeclinedTransactionSourceCheckDepositRejectionReason = "requested_by_user"
	DeclinedTransactionSourceCheckDepositRejectionReasonUnknown                     DeclinedTransactionSourceCheckDepositRejectionReason = "unknown"
)

func (r DeclinedTransactionSourceCheckDepositRejectionReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceCheckDepositRejectionReasonIncompleteImage, DeclinedTransactionSourceCheckDepositRejectionReasonDuplicate, DeclinedTransactionSourceCheckDepositRejectionReasonPoorImageQuality, DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectAmount, DeclinedTransactionSourceCheckDepositRejectionReasonIncorrectRecipient, DeclinedTransactionSourceCheckDepositRejectionReasonNotEligibleForMobileDeposit, DeclinedTransactionSourceCheckDepositRejectionReasonMissingRequiredDataElements, DeclinedTransactionSourceCheckDepositRejectionReasonSuspectedFraud, DeclinedTransactionSourceCheckDepositRejectionReasonDepositWindowExpired, DeclinedTransactionSourceCheckDepositRejectionReasonRequestedByUser, DeclinedTransactionSourceCheckDepositRejectionReasonUnknown:
		return true
	}
	return false
}

// An Inbound FedNow Transfer Decline object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_fednow_transfer_decline`.
type DeclinedTransactionSourceInboundFednowTransferDecline struct {
	// Why the transfer was declined.
	Reason DeclinedTransactionSourceInboundFednowTransferDeclineReason `json:"reason,required"`
	// The identifier of the FedNow Transfer that led to this declined transaction.
	TransferID  string                                                    `json:"transfer_id,required"`
	ExtraFields map[string]interface{}                                    `json:"-,extras"`
	JSON        declinedTransactionSourceInboundFednowTransferDeclineJSON `json:"-"`
}

// declinedTransactionSourceInboundFednowTransferDeclineJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceInboundFednowTransferDecline]
type declinedTransactionSourceInboundFednowTransferDeclineJSON struct {
	Reason      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceInboundFednowTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceInboundFednowTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// Why the transfer was declined.
type DeclinedTransactionSourceInboundFednowTransferDeclineReason string

const (
	DeclinedTransactionSourceInboundFednowTransferDeclineReasonAccountNumberCanceled DeclinedTransactionSourceInboundFednowTransferDeclineReason = "account_number_canceled"
	DeclinedTransactionSourceInboundFednowTransferDeclineReasonAccountNumberDisabled DeclinedTransactionSourceInboundFednowTransferDeclineReason = "account_number_disabled"
	DeclinedTransactionSourceInboundFednowTransferDeclineReasonAccountRestricted     DeclinedTransactionSourceInboundFednowTransferDeclineReason = "account_restricted"
	DeclinedTransactionSourceInboundFednowTransferDeclineReasonGroupLocked           DeclinedTransactionSourceInboundFednowTransferDeclineReason = "group_locked"
	DeclinedTransactionSourceInboundFednowTransferDeclineReasonEntityNotActive       DeclinedTransactionSourceInboundFednowTransferDeclineReason = "entity_not_active"
	DeclinedTransactionSourceInboundFednowTransferDeclineReasonFednowNotEnabled      DeclinedTransactionSourceInboundFednowTransferDeclineReason = "fednow_not_enabled"
)

func (r DeclinedTransactionSourceInboundFednowTransferDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceInboundFednowTransferDeclineReasonAccountNumberCanceled, DeclinedTransactionSourceInboundFednowTransferDeclineReasonAccountNumberDisabled, DeclinedTransactionSourceInboundFednowTransferDeclineReasonAccountRestricted, DeclinedTransactionSourceInboundFednowTransferDeclineReasonGroupLocked, DeclinedTransactionSourceInboundFednowTransferDeclineReasonEntityNotActive, DeclinedTransactionSourceInboundFednowTransferDeclineReasonFednowNotEnabled:
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
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
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
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted          DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

func (r DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) IsKnown() bool {
	switch r {
	case DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive, DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled:
		return true
	}
	return false
}

// If the category of this Transaction source is equal to `other`, this field will
// contain an empty object, otherwise it will contain null.
type DeclinedTransactionSourceOther struct {
	JSON declinedTransactionSourceOtherJSON `json:"-"`
}

// declinedTransactionSourceOtherJSON contains the JSON metadata for the struct
// [DeclinedTransactionSourceOther]
type declinedTransactionSourceOtherJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeclinedTransactionSourceOther) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r declinedTransactionSourceOtherJSON) RawJSON() string {
	return r.raw
}

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type DeclinedTransactionSourceWireDecline struct {
	// The identifier of the Inbound Wire Transfer that was declined.
	InboundWireTransferID string `json:"inbound_wire_transfer_id,required"`
	// Why the wire transfer was declined.
	Reason      DeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	ExtraFields map[string]interface{}                     `json:"-,extras"`
	JSON        declinedTransactionSourceWireDeclineJSON   `json:"-"`
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
	DeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled DeclinedTransactionSourceWireDeclineReason = "account_number_canceled"
	DeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled DeclinedTransactionSourceWireDeclineReason = "account_number_disabled"
	DeclinedTransactionSourceWireDeclineReasonEntityNotActive       DeclinedTransactionSourceWireDeclineReason = "entity_not_active"
	DeclinedTransactionSourceWireDeclineReasonGroupLocked           DeclinedTransactionSourceWireDeclineReason = "group_locked"
	DeclinedTransactionSourceWireDeclineReasonNoAccountNumber       DeclinedTransactionSourceWireDeclineReason = "no_account_number"
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
	DeclinedTransactionListParamsCategoryInACHDecline                             DeclinedTransactionListParamsCategoryIn = "ach_decline"
	DeclinedTransactionListParamsCategoryInCardDecline                            DeclinedTransactionListParamsCategoryIn = "card_decline"
	DeclinedTransactionListParamsCategoryInCheckDecline                           DeclinedTransactionListParamsCategoryIn = "check_decline"
	DeclinedTransactionListParamsCategoryInInboundRealTimePaymentsTransferDecline DeclinedTransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_decline"
	DeclinedTransactionListParamsCategoryInInboundFednowTransferDecline           DeclinedTransactionListParamsCategoryIn = "inbound_fednow_transfer_decline"
	DeclinedTransactionListParamsCategoryInWireDecline                            DeclinedTransactionListParamsCategoryIn = "wire_decline"
	DeclinedTransactionListParamsCategoryInCheckDepositRejection                  DeclinedTransactionListParamsCategoryIn = "check_deposit_rejection"
	DeclinedTransactionListParamsCategoryInOther                                  DeclinedTransactionListParamsCategoryIn = "other"
)

func (r DeclinedTransactionListParamsCategoryIn) IsKnown() bool {
	switch r {
	case DeclinedTransactionListParamsCategoryInACHDecline, DeclinedTransactionListParamsCategoryInCardDecline, DeclinedTransactionListParamsCategoryInCheckDecline, DeclinedTransactionListParamsCategoryInInboundRealTimePaymentsTransferDecline, DeclinedTransactionListParamsCategoryInInboundFednowTransferDecline, DeclinedTransactionListParamsCategoryInWireDecline, DeclinedTransactionListParamsCategoryInCheckDepositRejection, DeclinedTransactionListParamsCategoryInOther:
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
