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
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
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
	JSON declinedTransactionSourceACHDeclineJSON
}

// declinedTransactionSourceACHDeclineJSON contains the JSON metadata for the
// struct [DeclinedTransactionSourceACHDecline]
type declinedTransactionSourceACHDeclineJSON struct {
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
	// The transaction would cause an Increase limit to be exceeded.
	DeclinedTransactionSourceACHDeclineReasonBreachesLimit DeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused. This is a reasonable default reason for decline of
	// credits.
	DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver DeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// A rare return reason. The return this message refers to was a duplicate.
	DeclinedTransactionSourceACHDeclineReasonDuplicateReturn DeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	DeclinedTransactionSourceACHDeclineReasonEntityNotActive DeclinedTransactionSourceACHDeclineReason = "entity_not_active"
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
	// The user initiated the decline.
	DeclinedTransactionSourceACHDeclineReasonUserInitiated DeclinedTransactionSourceACHDeclineReason = "user_initiated"
)

// A constant representing the object's type. For this resource it will always be
// `ach_decline`.
type DeclinedTransactionSourceACHDeclineType string

const (
	DeclinedTransactionSourceACHDeclineTypeACHDecline DeclinedTransactionSourceACHDeclineType = "ach_decline"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type DeclinedTransactionSourceCardDecline struct {
	// The Card Decline identifier.
	ID string `json:"id,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress DeclinedTransactionSourceCardDeclineCardholderAddress `json:"cardholder_address,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency DeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
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
	ID                   apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	CardholderAddress    apijson.Field
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

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type DeclinedTransactionSourceCardDeclineCardholderAddress struct {
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
	VerificationResult DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult `json:"verification_result,required"`
	JSON               declinedTransactionSourceCardDeclineCardholderAddressJSON
}

// declinedTransactionSourceCardDeclineCardholderAddressJSON contains the JSON
// metadata for the struct [DeclinedTransactionSourceCardDeclineCardholderAddress]
type declinedTransactionSourceCardDeclineCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	VerificationResult apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DeclinedTransactionSourceCardDeclineCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The address verification result returned to the card network.
type DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult string

const (
	// No adress was provided in the authorization request.
	DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResultNotChecked DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult = "not_checked"
	// Postal code matches, but the street address was not verified
	DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResultPostalCodeMatchAddressNotChecked DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match
	DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResultPostalCodeMatchAddressNoMatch DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches
	DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResultPostalCodeNoMatchAddressMatch DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult = "postal_code_no_match_address_match"
	// Postal code and street address match
	DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResultMatch DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult = "match"
	// Postal code and street address do not match
	DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResultNoMatch DeclinedTransactionSourceCardDeclineCardholderAddressVerificationResult = "no_match"
)

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

// Fields specific to the `network`.
type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category DeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
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

// The payment network used to process this card authorization.
type DeclinedTransactionSourceCardDeclineNetworkDetailsCategory string

const (
	// Visa
	DeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa DeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
)

// Fields specific to the `visa` network.
type DeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode DeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
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
	Amount int64 `json:"amount,required"`
	// A computer-readable number printed on the MICR line of business checks, usually
	// the check number. This is useful for positive pay checks, but can be unreliably
	// transmitted by the bank of first deposit.
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
	// The account number is disabled.
	DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled DeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The account number is canceled.
	DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled DeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The transaction would cause a limit to be exceeded.
	DeclinedTransactionSourceCheckDeclineReasonBreachesLimit DeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
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
)

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

// An International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type DeclinedTransactionSourceInternationalACHDecline struct {
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
	ForeignExchangeIndicator DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
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
	OriginatingDepositoryFinancialInstitutionIDQualifier DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
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
	ReceivingDepositoryFinancialInstitutionIDQualifier DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach#returns).
	TraceNumber string `json:"trace_number,required"`
	JSON        declinedTransactionSourceInternationalACHDeclineJSON
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

// A description of how the foreign exchange rate was calculated.
type DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToVariable DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorVariableToFixed DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicatorFixedToFixed DeclinedTransactionSourceInternationalACHDeclineForeignExchangeIndicator = "fixed_to_fixed"
)

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeRate DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicatorBlank DeclinedTransactionSourceInternationalACHDeclineForeignExchangeReferenceIndicator = "blank"
)

// The type of transfer. Set by the originator.
type DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAnnuity DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBusinessOrCommercial DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeDeposit DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeLoan DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMiscellaneous DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMortgage DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePension DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRemittance DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRentOrLease DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSalaryOrPayroll DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTax DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeAccountsReceivable DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeBackOfficeConversion DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeMachineTransfer DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfPurchase DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodePointOfSale DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeRepresentedCheck DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeSharedNetworkTransaction DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeTelphoneInitiated DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCodeInternetInitiated DeclinedTransactionSourceInternationalACHDeclineInternationalTransactionTypeCode = "internet_initiated"
)

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierBicCode DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifierIban DeclinedTransactionSourceInternationalACHDeclineOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierBicCode DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifierIban DeclinedTransactionSourceInternationalACHDeclineReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type DeclinedTransactionSourceWireDecline struct {
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
	OriginatorRoutingNumber                 apijson.Field
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
	// International ACH Decline: details will be under the `international_ach_decline`
	// object.
	DeclinedTransactionListParamsCategoryInInternationalACHDecline DeclinedTransactionListParamsCategoryIn = "international_ach_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	DeclinedTransactionListParamsCategoryInWireDecline DeclinedTransactionListParamsCategoryIn = "wire_decline"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	DeclinedTransactionListParamsCategoryInOther DeclinedTransactionListParamsCategoryIn = "other"
)

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
