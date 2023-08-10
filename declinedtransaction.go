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

// DeclinedTransactionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeclinedTransactionService]
// method instead.
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
	path := fmt.Sprintf("declined_transactions/%s", declinedTransactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Declined Transactions
func (r *DeclinedTransactionService) List(ctx context.Context, query DeclinedTransactionListParams, opts ...option.RequestOption) (res *shared.Page[DeclinedTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *DeclinedTransactionService) ListAutoPaging(ctx context.Context, query DeclinedTransactionListParams, opts ...option.RequestOption) *shared.PageAutoPager[DeclinedTransaction] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
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
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
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
	JSON declinedTransactionJSON
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

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
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

// The type of the route this Declined Transaction came through.
type DeclinedTransactionRouteType string

const (
	// An Account Number.
	DeclinedTransactionRouteTypeAccountNumber DeclinedTransactionRouteType = "account_number"
	// A Card.
	DeclinedTransactionRouteTypeCard DeclinedTransactionRouteType = "card"
)

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
	// An Inbound Real Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// An International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline DeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline DeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        declinedTransactionSourceJSON
}

// declinedTransactionSourceJSON contains the JSON metadata for the struct
// [DeclinedTransactionSource]
type declinedTransactionSourceJSON struct {
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

func (r *DeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type DeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	// Why the ACH transfer was declined.
	Reason           DeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                    `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                    `json:"receiver_name,required,nullable"`
	TraceNumber      string                                    `json:"trace_number,required"`
	JSON             declinedTransactionSourceACHDeclineJSON
}

// declinedTransactionSourceACHDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceACHDecline]
type declinedTransactionSourceACHDeclineJSON struct {
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

func (r *DeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH transfer was declined.
type DeclinedTransactionSourceACHDeclineReason string

const (
	// The account number is canceled.
	DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled DeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled DeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	DeclinedTransactionSourceACHDeclineReasonBreachesLimit DeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused.
	DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver DeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// Other.
	DeclinedTransactionSourceACHDeclineReasonDuplicateReturn DeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	DeclinedTransactionSourceACHDeclineReasonEntityNotActive DeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// Your account is inactive.
	DeclinedTransactionSourceACHDeclineReasonGroupLocked DeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	DeclinedTransactionSourceACHDeclineReasonInsufficientFunds DeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// Other.
	DeclinedTransactionSourceACHDeclineReasonMisroutedReturn DeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// Other.
	DeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit DeclinedTransactionSourceACHDeclineReason = "return_of_erroneous_or_reversing_debit"
	// The account number that was debited does not exist.
	DeclinedTransactionSourceACHDeclineReasonNoACHRoute DeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// Other.
	DeclinedTransactionSourceACHDeclineReasonOriginatorRequest DeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	DeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed DeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type DeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency DeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
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
	// Fields specific to the `network`
	NetworkDetails DeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason DeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	JSON   declinedTransactionSourceCardDeclineJSON
}

// declinedTransactionSourceCardDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceCardDecline]
type declinedTransactionSourceCardDeclineJSON struct {
	Amount               apijson.Field
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
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Fields specific to the `network`
type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization
	Category DeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network
	Visa DeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON declinedTransactionSourceCardDeclineNetworkDetailsJSON
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

// The payment network used to process this card authorization
type DeclinedTransactionSourceCardDeclineNetworkDetailsCategory string

const (
	// Visa
	DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa DeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
)

// Fields specific to the `visa` network
type DeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    declinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
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

// Why the transaction was declined.
type DeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	DeclinedTransactionSourceCardDeclineReasonCardNotActive DeclinedTransactionSourceCardDeclineReason = "card_not_active"
	// The account's entity was not active.
	DeclinedTransactionSourceCardDeclineReasonEntityNotActive DeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	// The account was inactive.
	DeclinedTransactionSourceCardDeclineReasonGroupLocked DeclinedTransactionSourceCardDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	DeclinedTransactionSourceCardDeclineReasonInsufficientFunds DeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	DeclinedTransactionSourceCardDeclineReasonCvv2Mismatch DeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	DeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed DeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by an internal limit for new Increase accounts.
	DeclinedTransactionSourceCardDeclineReasonBreachesInternalLimit DeclinedTransactionSourceCardDeclineReason = "breaches_internal_limit"
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
)

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
	// Inbound Real Time Payments Transfer Decline: details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	DeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline DeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// International ACH Decline: details will be under the `international_ach_decline`
	// object.
	DeclinedTransactionSourceCategoryInternationalACHDecline DeclinedTransactionSourceCategory = "international_ach_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	DeclinedTransactionSourceCategoryWireDecline DeclinedTransactionSourceCategory = "wire_decline"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	DeclinedTransactionSourceCategoryOther DeclinedTransactionSourceCategory = "other"
)

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type DeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason DeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   declinedTransactionSourceCheckDeclineJSON
}

// declinedTransactionSourceCheckDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceCheckDecline]
type declinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Field
	AuxiliaryOnUs apijson.Field
	Reason        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the check was declined.
type DeclinedTransactionSourceCheckDeclineReason string

const (
	// The account number is canceled.
	DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled DeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled DeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	DeclinedTransactionSourceCheckDeclineReasonBreachesLimit DeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The account's entity is not active.
	DeclinedTransactionSourceCheckDeclineReasonEntityNotActive DeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	DeclinedTransactionSourceCheckDeclineReasonGroupLocked DeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	DeclinedTransactionSourceCheckDeclineReasonInsufficientFunds DeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Unable to locate account.
	DeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount DeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	// Routing number on the check is not ours.
	DeclinedTransactionSourceCheckDeclineReasonNotOurItem DeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	// Unable to process.
	DeclinedTransactionSourceCheckDeclineReasonUnableToProcess DeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	// Refer to image.
	DeclinedTransactionSourceCheckDeclineReasonReferToImage DeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	// Stop payment requested for this check.
	DeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested DeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// Check was returned to sender.
	DeclinedTransactionSourceCheckDeclineReasonReturned DeclinedTransactionSourceCheckDeclineReason = "returned"
	// The check was a duplicate deposit.
	DeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment DeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The transaction is not allowed.
	DeclinedTransactionSourceCheckDeclineReasonNotAuthorized DeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The check was altered or fictitious.
	DeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious DeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
)

// An Inbound Real Time Payments Transfer Decline object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
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
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      declinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
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
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
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
	// Your account is not enabled to receive Real Time Payments transfers.
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

// An International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type DeclinedTransactionSourceInternationalACHDecline struct {
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
	JSON                                                   declinedTransactionSourceInternationalACHDeclineJSON
}

// declinedTransactionSourceInternationalACHDeclineJSON contains the JSON metadata
// for the struct [DeclinedTransactionSourceInternationalACHDecline]
type declinedTransactionSourceInternationalACHDeclineJSON struct {
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

func (r *DeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type DeclinedTransactionSourceWireDecline struct {
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
	Reason DeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	JSON   declinedTransactionSourceWireDeclineJSON
}

// declinedTransactionSourceWireDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceWireDecline]
type declinedTransactionSourceWireDeclineJSON struct {
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

func (r *DeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
type DeclinedTransactionType string

const (
	DeclinedTransactionTypeDeclinedTransaction DeclinedTransactionType = "declined_transaction"
)

type DeclinedTransactionListParams struct {
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID param.Field[string]                                 `query:"account_id"`
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
