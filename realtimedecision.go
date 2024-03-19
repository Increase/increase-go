// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// RealTimeDecisionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRealTimeDecisionService] method
// instead.
type RealTimeDecisionService struct {
	Options []option.RequestOption
}

// NewRealTimeDecisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRealTimeDecisionService(opts ...option.RequestOption) (r *RealTimeDecisionService) {
	r = &RealTimeDecisionService{}
	r.Options = opts
	return
}

// Retrieve a Real-Time Decision
func (r *RealTimeDecisionService) Get(ctx context.Context, realTimeDecisionID string, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s", realTimeDecisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, realTimeDecisionID string, body RealTimeDecisionActionParams, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s/action", realTimeDecisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Real Time Decisions are created when your application needs to take action in
// real-time to some event such as a card authorization. For more information, see
// our
// [Real-Time Decisions guide](https://increase.com/documentation/real-time-decisions).
type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID string `json:"id,required"`
	// Fields related to a card authorization.
	CardAuthorization RealTimeDecisionCardAuthorization `json:"card_authorization,required,nullable"`
	// The category of the Real-Time Decision.
	Category RealTimeDecisionCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication RealTimeDecisionDigitalWalletAuthentication `json:"digital_wallet_authentication,required,nullable"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken RealTimeDecisionDigitalWalletToken `json:"digital_wallet_token,required,nullable"`
	// The status of the Real-Time Decision.
	Status RealTimeDecisionStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt time.Time `json:"timeout_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type RealTimeDecisionType `json:"type,required"`
	JSON realTimeDecisionJSON `json:"-"`
}

// realTimeDecisionJSON contains the JSON metadata for the struct
// [RealTimeDecision]
type realTimeDecisionJSON struct {
	ID                          apijson.Field
	CardAuthorization           apijson.Field
	Category                    apijson.Field
	CreatedAt                   apijson.Field
	DigitalWalletAuthentication apijson.Field
	DigitalWalletToken          apijson.Field
	Status                      apijson.Field
	TimeoutAt                   apijson.Field
	Type                        apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionJSON) RawJSON() string {
	return r.raw
}

// Fields related to a card authorization.
type RealTimeDecisionCardAuthorization struct {
	// The identifier of the Account the authorization will debit.
	AccountID string `json:"account_id,required"`
	// The identifier of the Card that is being authorized.
	CardID string `json:"card_id,required"`
	// Whether or not the authorization was approved.
	Decision RealTimeDecisionCardAuthorizationDecision `json:"decision,required,nullable"`
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
	// Fields specific to the `network`.
	NetworkDetails RealTimeDecisionCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers RealTimeDecisionCardAuthorizationNetworkIdentifiers `json:"network_identifiers,required"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory RealTimeDecisionCardAuthorizationProcessingCategory `json:"processing_category,required"`
	// Fields specific to the type of request, such as an incremental authorization.
	RequestDetails RealTimeDecisionCardAuthorizationRequestDetails `json:"request_details,required"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount int64 `json:"settlement_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency string `json:"settlement_currency,required"`
	// Fields related to verification of cardholder-provided values.
	Verification RealTimeDecisionCardAuthorizationVerification `json:"verification,required"`
	JSON         realTimeDecisionCardAuthorizationJSON         `json:"-"`
}

// realTimeDecisionCardAuthorizationJSON contains the JSON metadata for the struct
// [RealTimeDecisionCardAuthorization]
type realTimeDecisionCardAuthorizationJSON struct {
	AccountID            apijson.Field
	CardID               apijson.Field
	Decision             apijson.Field
	DigitalWalletTokenID apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	NetworkDetails       apijson.Field
	NetworkIdentifiers   apijson.Field
	PhysicalCardID       apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	ProcessingCategory   apijson.Field
	RequestDetails       apijson.Field
	SettlementAmount     apijson.Field
	SettlementCurrency   apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Whether or not the authorization was approved.
type RealTimeDecisionCardAuthorizationDecision string

const (
	// Approve the authorization.
	RealTimeDecisionCardAuthorizationDecisionApprove RealTimeDecisionCardAuthorizationDecision = "approve"
	// Decline the authorization.
	RealTimeDecisionCardAuthorizationDecisionDecline RealTimeDecisionCardAuthorizationDecision = "decline"
)

func (r RealTimeDecisionCardAuthorizationDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationDecisionApprove, RealTimeDecisionCardAuthorizationDecisionDecline:
		return true
	}
	return false
}

// Fields specific to the `network`.
type RealTimeDecisionCardAuthorizationNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category RealTimeDecisionCardAuthorizationNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa RealTimeDecisionCardAuthorizationNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON realTimeDecisionCardAuthorizationNetworkDetailsJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationNetworkDetailsJSON contains the JSON metadata
// for the struct [RealTimeDecisionCardAuthorizationNetworkDetails]
type realTimeDecisionCardAuthorizationNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type RealTimeDecisionCardAuthorizationNetworkDetailsCategory string

const (
	// Visa
	RealTimeDecisionCardAuthorizationNetworkDetailsCategoryVisa RealTimeDecisionCardAuthorizationNetworkDetailsCategory = "visa"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
}

// realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [RealTimeDecisionCardAuthorizationNetworkDetailsVisa]
type realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type RealTimeDecisionCardAuthorizationNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                  `json:"transaction_id,required,nullable"`
	JSON          realTimeDecisionCardAuthorizationNetworkIdentifiersJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationNetworkIdentifiersJSON contains the JSON
// metadata for the struct [RealTimeDecisionCardAuthorizationNetworkIdentifiers]
type realTimeDecisionCardAuthorizationNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type RealTimeDecisionCardAuthorizationProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	RealTimeDecisionCardAuthorizationProcessingCategoryAccountFunding RealTimeDecisionCardAuthorizationProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	RealTimeDecisionCardAuthorizationProcessingCategoryAutomaticFuelDispenser RealTimeDecisionCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	RealTimeDecisionCardAuthorizationProcessingCategoryBillPayment RealTimeDecisionCardAuthorizationProcessingCategory = "bill_payment"
	// A regular purchase.
	RealTimeDecisionCardAuthorizationProcessingCategoryPurchase RealTimeDecisionCardAuthorizationProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	RealTimeDecisionCardAuthorizationProcessingCategoryQuasiCash RealTimeDecisionCardAuthorizationProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	RealTimeDecisionCardAuthorizationProcessingCategoryRefund RealTimeDecisionCardAuthorizationProcessingCategory = "refund"
)

func (r RealTimeDecisionCardAuthorizationProcessingCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationProcessingCategoryAccountFunding, RealTimeDecisionCardAuthorizationProcessingCategoryAutomaticFuelDispenser, RealTimeDecisionCardAuthorizationProcessingCategoryBillPayment, RealTimeDecisionCardAuthorizationProcessingCategoryPurchase, RealTimeDecisionCardAuthorizationProcessingCategoryQuasiCash, RealTimeDecisionCardAuthorizationProcessingCategoryRefund:
		return true
	}
	return false
}

// Fields specific to the type of request, such as an incremental authorization.
type RealTimeDecisionCardAuthorizationRequestDetails struct {
	// The type of this request (e.g., an initial authorization or an incremental
	// authorization).
	Category RealTimeDecisionCardAuthorizationRequestDetailsCategory `json:"category,required"`
	// Fields specific to the category `incremental_authorization`.
	IncrementalAuthorization RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization `json:"incremental_authorization,required,nullable"`
	// Fields specific to the category `initial_authorization`.
	InitialAuthorization interface{}                                         `json:"initial_authorization,required,nullable"`
	JSON                 realTimeDecisionCardAuthorizationRequestDetailsJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationRequestDetailsJSON contains the JSON metadata
// for the struct [RealTimeDecisionCardAuthorizationRequestDetails]
type realTimeDecisionCardAuthorizationRequestDetailsJSON struct {
	Category                 apijson.Field
	IncrementalAuthorization apijson.Field
	InitialAuthorization     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationRequestDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationRequestDetailsJSON) RawJSON() string {
	return r.raw
}

// The type of this request (e.g., an initial authorization or an incremental
// authorization).
type RealTimeDecisionCardAuthorizationRequestDetailsCategory string

const (
	// A regular, standalone authorization.
	RealTimeDecisionCardAuthorizationRequestDetailsCategoryInitialAuthorization RealTimeDecisionCardAuthorizationRequestDetailsCategory = "initial_authorization"
	// An incremental request to increase the amount of an existing authorization.
	RealTimeDecisionCardAuthorizationRequestDetailsCategoryIncrementalAuthorization RealTimeDecisionCardAuthorizationRequestDetailsCategory = "incremental_authorization"
)

func (r RealTimeDecisionCardAuthorizationRequestDetailsCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationRequestDetailsCategoryInitialAuthorization, RealTimeDecisionCardAuthorizationRequestDetailsCategoryIncrementalAuthorization:
		return true
	}
	return false
}

// Fields specific to the category `incremental_authorization`.
type RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization struct {
	// The card payment for this authorization and increment.
	CardPaymentID string `json:"card_payment_id,required"`
	// The identifier of the card authorization this request is attempting to
	// increment.
	OriginalCardAuthorizationID string                                                                      `json:"original_card_authorization_id,required"`
	JSON                        realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON
// contains the JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization]
type realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON struct {
	CardPaymentID               apijson.Field
	OriginalCardAuthorizationID apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of cardholder-provided values.
type RealTimeDecisionCardAuthorizationVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode RealTimeDecisionCardAuthorizationVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress RealTimeDecisionCardAuthorizationVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              realTimeDecisionCardAuthorizationVerificationJSON              `json:"-"`
}

// realTimeDecisionCardAuthorizationVerificationJSON contains the JSON metadata for
// the struct [RealTimeDecisionCardAuthorizationVerification]
type realTimeDecisionCardAuthorizationVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type RealTimeDecisionCardAuthorizationVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON   `json:"-"`
}

// realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON contains
// the JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationVerificationCardVerificationCode]
type realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNotChecked RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultMatch RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNoMatch RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
)

func (r RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNotChecked, RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultMatch, RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type RealTimeDecisionCardAuthorizationVerificationCardholderAddress struct {
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
	Result RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult `json:"result,required"`
	JSON   realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON   `json:"-"`
}

// realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON contains the
// JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationVerificationCardholderAddress]
type realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNotChecked RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultMatch RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNoMatch RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "no_match"
)

func (r RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNotChecked, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// The category of the Real-Time Decision.
type RealTimeDecisionCategory string

const (
	// A card is being authorized.
	RealTimeDecisionCategoryCardAuthorizationRequested RealTimeDecisionCategory = "card_authorization_requested"
	// A card is being loaded into a digital wallet.
	RealTimeDecisionCategoryDigitalWalletTokenRequested RealTimeDecisionCategory = "digital_wallet_token_requested"
	// A card is being loaded into a digital wallet and requires cardholder
	// authentication.
	RealTimeDecisionCategoryDigitalWalletAuthenticationRequested RealTimeDecisionCategory = "digital_wallet_authentication_requested"
)

func (r RealTimeDecisionCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCategoryCardAuthorizationRequested, RealTimeDecisionCategoryDigitalWalletTokenRequested, RealTimeDecisionCategoryDigitalWalletAuthenticationRequested:
		return true
	}
	return false
}

// Fields related to a digital wallet authentication attempt.
type RealTimeDecisionDigitalWalletAuthentication struct {
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The channel to send the card user their one-time passcode.
	Channel RealTimeDecisionDigitalWalletAuthenticationChannel `json:"channel,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletAuthenticationDigitalWallet `json:"digital_wallet,required"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email string `json:"email,required,nullable"`
	// The one-time passcode to send the card user.
	OneTimePasscode string `json:"one_time_passcode,required"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone string `json:"phone,required,nullable"`
	// Whether your application successfully delivered the one-time passcode.
	Result RealTimeDecisionDigitalWalletAuthenticationResult `json:"result,required,nullable"`
	JSON   realTimeDecisionDigitalWalletAuthenticationJSON   `json:"-"`
}

// realTimeDecisionDigitalWalletAuthenticationJSON contains the JSON metadata for
// the struct [RealTimeDecisionDigitalWalletAuthentication]
type realTimeDecisionDigitalWalletAuthenticationJSON struct {
	CardID          apijson.Field
	Channel         apijson.Field
	DigitalWallet   apijson.Field
	Email           apijson.Field
	OneTimePasscode apijson.Field
	Phone           apijson.Field
	Result          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionDigitalWalletAuthenticationJSON) RawJSON() string {
	return r.raw
}

// The channel to send the card user their one-time passcode.
type RealTimeDecisionDigitalWalletAuthenticationChannel string

const (
	// Send one-time passcodes over SMS.
	RealTimeDecisionDigitalWalletAuthenticationChannelSMS RealTimeDecisionDigitalWalletAuthenticationChannel = "sms"
	// Send one-time passcodes over email.
	RealTimeDecisionDigitalWalletAuthenticationChannelEmail RealTimeDecisionDigitalWalletAuthenticationChannel = "email"
)

func (r RealTimeDecisionDigitalWalletAuthenticationChannel) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletAuthenticationChannelSMS, RealTimeDecisionDigitalWalletAuthenticationChannelEmail:
		return true
	}
	return false
}

// The digital wallet app being used.
type RealTimeDecisionDigitalWalletAuthenticationDigitalWallet string

const (
	// Apple Pay
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "apple_pay"
	// Google Pay
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "google_pay"
	// Unknown
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletUnknown RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "unknown"
)

func (r RealTimeDecisionDigitalWalletAuthenticationDigitalWallet) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay, RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay, RealTimeDecisionDigitalWalletAuthenticationDigitalWalletUnknown:
		return true
	}
	return false
}

// Whether your application successfully delivered the one-time passcode.
type RealTimeDecisionDigitalWalletAuthenticationResult string

const (
	// Your application successfully delivered the one-time passcode to the cardholder.
	RealTimeDecisionDigitalWalletAuthenticationResultSuccess RealTimeDecisionDigitalWalletAuthenticationResult = "success"
	// Your application failed to deliver the one-time passcode to the cardholder.
	RealTimeDecisionDigitalWalletAuthenticationResultFailure RealTimeDecisionDigitalWalletAuthenticationResult = "failure"
)

func (r RealTimeDecisionDigitalWalletAuthenticationResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletAuthenticationResultSuccess, RealTimeDecisionDigitalWalletAuthenticationResultFailure:
		return true
	}
	return false
}

// Fields related to a digital wallet token provisioning attempt.
type RealTimeDecisionDigitalWalletToken struct {
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The identifier of the Card Profile that was set via the real time decision. This
	// will be null until the real time decision is responded to or if the real time
	// decision did not set a card profile.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	// Whether or not the provisioning request was approved. This will be null until
	// the real time decision is responded to.
	Decision RealTimeDecisionDigitalWalletTokenDecision `json:"decision,required,nullable"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet `json:"digital_wallet,required"`
	JSON          realTimeDecisionDigitalWalletTokenJSON          `json:"-"`
}

// realTimeDecisionDigitalWalletTokenJSON contains the JSON metadata for the struct
// [RealTimeDecisionDigitalWalletToken]
type realTimeDecisionDigitalWalletTokenJSON struct {
	CardID        apijson.Field
	CardProfileID apijson.Field
	Decision      apijson.Field
	DigitalWallet apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionDigitalWalletTokenJSON) RawJSON() string {
	return r.raw
}

// Whether or not the provisioning request was approved. This will be null until
// the real time decision is responded to.
type RealTimeDecisionDigitalWalletTokenDecision string

const (
	// Approve the provisioning request.
	RealTimeDecisionDigitalWalletTokenDecisionApprove RealTimeDecisionDigitalWalletTokenDecision = "approve"
	// Decline the provisioning request.
	RealTimeDecisionDigitalWalletTokenDecisionDecline RealTimeDecisionDigitalWalletTokenDecision = "decline"
)

func (r RealTimeDecisionDigitalWalletTokenDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletTokenDecisionApprove, RealTimeDecisionDigitalWalletTokenDecisionDecline:
		return true
	}
	return false
}

// The digital wallet app being used.
type RealTimeDecisionDigitalWalletTokenDigitalWallet string

const (
	// Apple Pay
	RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay RealTimeDecisionDigitalWalletTokenDigitalWallet = "apple_pay"
	// Google Pay
	RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay RealTimeDecisionDigitalWalletTokenDigitalWallet = "google_pay"
	// Unknown
	RealTimeDecisionDigitalWalletTokenDigitalWalletUnknown RealTimeDecisionDigitalWalletTokenDigitalWallet = "unknown"
)

func (r RealTimeDecisionDigitalWalletTokenDigitalWallet) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay, RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay, RealTimeDecisionDigitalWalletTokenDigitalWalletUnknown:
		return true
	}
	return false
}

// The status of the Real-Time Decision.
type RealTimeDecisionStatus string

const (
	// The decision is pending action via real-time webhook.
	RealTimeDecisionStatusPending RealTimeDecisionStatus = "pending"
	// Your webhook actioned the real-time decision.
	RealTimeDecisionStatusResponded RealTimeDecisionStatus = "responded"
	// Your webhook failed to respond to the authorization in time.
	RealTimeDecisionStatusTimedOut RealTimeDecisionStatus = "timed_out"
)

func (r RealTimeDecisionStatus) IsKnown() bool {
	switch r {
	case RealTimeDecisionStatusPending, RealTimeDecisionStatusResponded, RealTimeDecisionStatusTimedOut:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `real_time_decision`.
type RealTimeDecisionType string

const (
	RealTimeDecisionTypeRealTimeDecision RealTimeDecisionType = "real_time_decision"
)

func (r RealTimeDecisionType) IsKnown() bool {
	switch r {
	case RealTimeDecisionTypeRealTimeDecision:
		return true
	}
	return false
}

type RealTimeDecisionActionParams struct {
	// If the Real-Time Decision relates to a card authorization attempt, this object
	// contains your response to the authorization.
	CardAuthorization param.Field[RealTimeDecisionActionParamsCardAuthorization] `json:"card_authorization"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication param.Field[RealTimeDecisionActionParamsDigitalWalletAuthentication] `json:"digital_wallet_authentication"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken param.Field[RealTimeDecisionActionParamsDigitalWalletToken] `json:"digital_wallet_token"`
}

func (r RealTimeDecisionActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Real-Time Decision relates to a card authorization attempt, this object
// contains your response to the authorization.
type RealTimeDecisionActionParamsCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision param.Field[RealTimeDecisionActionParamsCardAuthorizationDecision] `json:"decision,required"`
}

func (r RealTimeDecisionActionParamsCardAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card authorization should be approved or declined.
type RealTimeDecisionActionParamsCardAuthorizationDecision string

const (
	// Approve the authorization.
	RealTimeDecisionActionParamsCardAuthorizationDecisionApprove RealTimeDecisionActionParamsCardAuthorizationDecision = "approve"
	// Decline the authorization.
	RealTimeDecisionActionParamsCardAuthorizationDecisionDecline RealTimeDecisionActionParamsCardAuthorizationDecision = "decline"
)

func (r RealTimeDecisionActionParamsCardAuthorizationDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthorizationDecisionApprove, RealTimeDecisionActionParamsCardAuthorizationDecisionDecline:
		return true
	}
	return false
}

// If the Real-Time Decision relates to a digital wallet authentication attempt,
// this object contains your response to the authentication.
type RealTimeDecisionActionParamsDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result param.Field[RealTimeDecisionActionParamsDigitalWalletAuthenticationResult] `json:"result,required"`
}

func (r RealTimeDecisionActionParamsDigitalWalletAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether your application was able to deliver the one-time passcode.
type RealTimeDecisionActionParamsDigitalWalletAuthenticationResult string

const (
	// Your application successfully delivered the one-time passcode to the cardholder.
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "success"
	// Your application failed to deliver the one-time passcode to the cardholder.
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "failure"
)

func (r RealTimeDecisionActionParamsDigitalWalletAuthenticationResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess, RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure:
		return true
	}
	return false
}

// If the Real-Time Decision relates to a digital wallet token provisioning
// attempt, this object contains your response to the attempt.
type RealTimeDecisionActionParamsDigitalWalletToken struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval param.Field[RealTimeDecisionActionParamsDigitalWalletTokenApproval] `json:"approval"`
	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline param.Field[RealTimeDecisionActionParamsDigitalWalletTokenDecline] `json:"decline"`
}

func (r RealTimeDecisionActionParamsDigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If your application approves the provisioning attempt, this contains metadata
// about the digital wallet token that will be generated.
type RealTimeDecisionActionParamsDigitalWalletTokenApproval struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID param.Field[string] `json:"card_profile_id"`
	// The identifier of the Digital Card Profile to assign to the Digital Wallet
	// token.
	DigitalCardProfileID param.Field[string] `json:"digital_card_profile_id"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email param.Field[string] `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone param.Field[string] `json:"phone"`
}

func (r RealTimeDecisionActionParamsDigitalWalletTokenApproval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If your application declines the provisioning attempt, this contains details
// about the decline.
type RealTimeDecisionActionParamsDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason param.Field[string] `json:"reason"`
}

func (r RealTimeDecisionActionParamsDigitalWalletTokenDecline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
