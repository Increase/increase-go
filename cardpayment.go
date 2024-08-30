// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// CardPaymentService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardPaymentService] method instead.
type CardPaymentService struct {
	Options []option.RequestOption
}

// NewCardPaymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardPaymentService(opts ...option.RequestOption) (r *CardPaymentService) {
	r = &CardPaymentService{}
	r.Options = opts
	return
}

// Retrieve a Card Payment
func (r *CardPaymentService) Get(ctx context.Context, cardPaymentID string, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	if cardPaymentID == "" {
		err = errors.New("missing required card_payment_id parameter")
		return
	}
	path := fmt.Sprintf("card_payments/%s", cardPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Payments
func (r *CardPaymentService) List(ctx context.Context, query CardPaymentListParams, opts ...option.RequestOption) (res *pagination.Page[CardPayment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_payments"
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

// List Card Payments
func (r *CardPaymentService) ListAutoPaging(ctx context.Context, query CardPaymentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CardPayment] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Card Payments group together interactions related to a single card payment, such
// as an authorization and its corresponding settlement.
type CardPayment struct {
	// The Card Payment identifier.
	ID string `json:"id,required"`
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Card identifier for this payment.
	CardID string `json:"card_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Card
	// Payment was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Digital Wallet Token identifier for this payment.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The interactions related to this card payment.
	Elements []CardPaymentElement `json:"elements,required"`
	// The Physical Card identifier for this payment.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The summarized state of this card payment.
	State CardPaymentState `json:"state,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_payment`.
	Type CardPaymentType `json:"type,required"`
	JSON cardPaymentJSON `json:"-"`
}

// cardPaymentJSON contains the JSON metadata for the struct [CardPayment]
type cardPaymentJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	CardID               apijson.Field
	CreatedAt            apijson.Field
	DigitalWalletTokenID apijson.Field
	Elements             apijson.Field
	PhysicalCardID       apijson.Field
	State                apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentJSON) RawJSON() string {
	return r.raw
}

type CardPaymentElement struct {
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization CardPaymentElementsCardAuthorization `json:"card_authorization,required,nullable"`
	// A Card Authorization Expiration object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_authorization_expiration`.
	CardAuthorizationExpiration CardPaymentElementsCardAuthorizationExpiration `json:"card_authorization_expiration,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline CardPaymentElementsCardDecline `json:"card_decline,required,nullable"`
	// A Card Fuel Confirmation object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_fuel_confirmation`.
	CardFuelConfirmation CardPaymentElementsCardFuelConfirmation `json:"card_fuel_confirmation,required,nullable"`
	// A Card Increment object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_increment`.
	CardIncrement CardPaymentElementsCardIncrement `json:"card_increment,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund CardPaymentElementsCardRefund `json:"card_refund,required,nullable"`
	// A Card Reversal object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_reversal`.
	CardReversal CardPaymentElementsCardReversal `json:"card_reversal,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement CardPaymentElementsCardSettlement `json:"card_settlement,required,nullable"`
	// A Card Validation object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_validation`.
	CardValidation CardPaymentElementsCardValidation `json:"card_validation,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category CardPaymentElementsCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the card payment element was created.
	CreatedAt time.Time              `json:"created_at,required" format:"date-time"`
	JSON      cardPaymentElementJSON `json:"-"`
}

// cardPaymentElementJSON contains the JSON metadata for the struct
// [CardPaymentElement]
type cardPaymentElementJSON struct {
	CardAuthorization           apijson.Field
	CardAuthorizationExpiration apijson.Field
	CardDecline                 apijson.Field
	CardFuelConfirmation        apijson.Field
	CardIncrement               apijson.Field
	CardRefund                  apijson.Field
	CardReversal                apijson.Field
	CardSettlement              apijson.Field
	CardValidation              apijson.Field
	Category                    apijson.Field
	CreatedAt                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardPaymentElement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementJSON) RawJSON() string {
	return r.raw
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
type CardPaymentElementsCardAuthorization struct {
	// The Card Authorization identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardAuthorizationActioner `json:"actioner,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardPaymentElementsCardAuthorizationCurrency `json:"currency,required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The direction describes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction CardPaymentElementsCardAuthorizationDirection `json:"direction,required"`
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
	// The merchant's postal code. For US merchants this is either a 5-digit or 9-digit
	// ZIP code, where the first 5 and last 4 are separated by a dash.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Fields specific to the `network`.
	NetworkDetails CardPaymentElementsCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardAuthorizationNetworkIdentifiers `json:"network_identifiers,required"`
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
	ProcessingCategory CardPaymentElementsCardAuthorizationProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type CardPaymentElementsCardAuthorizationType `json:"type,required"`
	// Fields related to verification of cardholder-provided values.
	Verification CardPaymentElementsCardAuthorizationVerification `json:"verification,required"`
	JSON         cardPaymentElementsCardAuthorizationJSON         `json:"-"`
}

// cardPaymentElementsCardAuthorizationJSON contains the JSON metadata for the
// struct [CardPaymentElementsCardAuthorization]
type cardPaymentElementsCardAuthorizationJSON struct {
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
	MerchantPostalCode   apijson.Field
	MerchantState        apijson.Field
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

func (r *CardPaymentElementsCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type CardPaymentElementsCardAuthorizationActioner string

const (
	// This object was actioned by the user through a real-time decision.
	CardPaymentElementsCardAuthorizationActionerUser CardPaymentElementsCardAuthorizationActioner = "user"
	// This object was actioned by Increase without user intervention.
	CardPaymentElementsCardAuthorizationActionerIncrease CardPaymentElementsCardAuthorizationActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	CardPaymentElementsCardAuthorizationActionerNetwork CardPaymentElementsCardAuthorizationActioner = "network"
)

func (r CardPaymentElementsCardAuthorizationActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationActionerUser, CardPaymentElementsCardAuthorizationActionerIncrease, CardPaymentElementsCardAuthorizationActionerNetwork:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CardPaymentElementsCardAuthorizationCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardAuthorizationCurrencyCad CardPaymentElementsCardAuthorizationCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardAuthorizationCurrencyChf CardPaymentElementsCardAuthorizationCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardAuthorizationCurrencyEur CardPaymentElementsCardAuthorizationCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardAuthorizationCurrencyGbp CardPaymentElementsCardAuthorizationCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardAuthorizationCurrencyJpy CardPaymentElementsCardAuthorizationCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardAuthorizationCurrencyUsd CardPaymentElementsCardAuthorizationCurrency = "USD"
)

func (r CardPaymentElementsCardAuthorizationCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationCurrencyCad, CardPaymentElementsCardAuthorizationCurrencyChf, CardPaymentElementsCardAuthorizationCurrencyEur, CardPaymentElementsCardAuthorizationCurrencyGbp, CardPaymentElementsCardAuthorizationCurrencyJpy, CardPaymentElementsCardAuthorizationCurrencyUsd:
		return true
	}
	return false
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type CardPaymentElementsCardAuthorizationDirection string

const (
	// A regular card authorization where funds are debited from the cardholder.
	CardPaymentElementsCardAuthorizationDirectionSettlement CardPaymentElementsCardAuthorizationDirection = "settlement"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	CardPaymentElementsCardAuthorizationDirectionRefund CardPaymentElementsCardAuthorizationDirection = "refund"
)

func (r CardPaymentElementsCardAuthorizationDirection) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationDirectionSettlement, CardPaymentElementsCardAuthorizationDirectionRefund:
		return true
	}
	return false
}

// Fields specific to the `network`.
type CardPaymentElementsCardAuthorizationNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category CardPaymentElementsCardAuthorizationNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa CardPaymentElementsCardAuthorizationNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON cardPaymentElementsCardAuthorizationNetworkDetailsJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationNetworkDetailsJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardAuthorizationNetworkDetails]
type cardPaymentElementsCardAuthorizationNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type CardPaymentElementsCardAuthorizationNetworkDetailsCategory string

const (
	// Visa
	CardPaymentElementsCardAuthorizationNetworkDetailsCategoryVisa CardPaymentElementsCardAuthorizationNetworkDetailsCategory = "visa"
)

func (r CardPaymentElementsCardAuthorizationNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type CardPaymentElementsCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
}

// cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardAuthorizationNetworkDetailsVisa]
type cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardPaymentElementsCardAuthorizationNetworkIdentifiers struct {
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
	JSON          cardPaymentElementsCardAuthorizationNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationNetworkIdentifiersJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardAuthorizationNetworkIdentifiers]
type cardPaymentElementsCardAuthorizationNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type CardPaymentElementsCardAuthorizationProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	CardPaymentElementsCardAuthorizationProcessingCategoryAccountFunding CardPaymentElementsCardAuthorizationProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	CardPaymentElementsCardAuthorizationProcessingCategoryAutomaticFuelDispenser CardPaymentElementsCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	CardPaymentElementsCardAuthorizationProcessingCategoryBillPayment CardPaymentElementsCardAuthorizationProcessingCategory = "bill_payment"
	// A regular purchase.
	CardPaymentElementsCardAuthorizationProcessingCategoryPurchase CardPaymentElementsCardAuthorizationProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	CardPaymentElementsCardAuthorizationProcessingCategoryQuasiCash CardPaymentElementsCardAuthorizationProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	CardPaymentElementsCardAuthorizationProcessingCategoryRefund CardPaymentElementsCardAuthorizationProcessingCategory = "refund"
)

func (r CardPaymentElementsCardAuthorizationProcessingCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationProcessingCategoryAccountFunding, CardPaymentElementsCardAuthorizationProcessingCategoryAutomaticFuelDispenser, CardPaymentElementsCardAuthorizationProcessingCategoryBillPayment, CardPaymentElementsCardAuthorizationProcessingCategoryPurchase, CardPaymentElementsCardAuthorizationProcessingCategoryQuasiCash, CardPaymentElementsCardAuthorizationProcessingCategoryRefund:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_authorization`.
type CardPaymentElementsCardAuthorizationType string

const (
	CardPaymentElementsCardAuthorizationTypeCardAuthorization CardPaymentElementsCardAuthorizationType = "card_authorization"
)

func (r CardPaymentElementsCardAuthorizationType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationTypeCardAuthorization:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type CardPaymentElementsCardAuthorizationVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode CardPaymentElementsCardAuthorizationVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress CardPaymentElementsCardAuthorizationVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              cardPaymentElementsCardAuthorizationVerificationJSON              `json:"-"`
}

// cardPaymentElementsCardAuthorizationVerificationJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardAuthorizationVerification]
type cardPaymentElementsCardAuthorizationVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type CardPaymentElementsCardAuthorizationVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   cardPaymentElementsCardAuthorizationVerificationCardVerificationCodeJSON   `json:"-"`
}

// cardPaymentElementsCardAuthorizationVerificationCardVerificationCodeJSON
// contains the JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationVerificationCardVerificationCode]
type cardPaymentElementsCardAuthorizationVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultNotChecked CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultMatch CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultNoMatch CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
)

func (r CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultNotChecked, CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultMatch, CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type CardPaymentElementsCardAuthorizationVerificationCardholderAddress struct {
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
	Result CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult `json:"result,required"`
	JSON   cardPaymentElementsCardAuthorizationVerificationCardholderAddressJSON   `json:"-"`
}

// cardPaymentElementsCardAuthorizationVerificationCardholderAddressJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationVerificationCardholderAddress]
type cardPaymentElementsCardAuthorizationVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNotChecked CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultMatch CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNoMatch CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "no_match"
)

func (r CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNotChecked, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// A Card Authorization Expiration object. This field will be present in the JSON
// response if and only if `category` is equal to `card_authorization_expiration`.
type CardPaymentElementsCardAuthorizationExpiration struct {
	// The Card Authorization Expiration identifier.
	ID string `json:"id,required"`
	// The identifier for the Card Authorization this reverses.
	CardAuthorizationID string `json:"card_authorization_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the reversal's
	// currency.
	Currency CardPaymentElementsCardAuthorizationExpirationCurrency `json:"currency,required"`
	// The amount of this authorization expiration in the minor unit of the
	// transaction's currency. For dollars, for example, this is cents.
	ExpiredAmount int64 `json:"expired_amount,required"`
	// The card network used to process this card authorization.
	Network CardPaymentElementsCardAuthorizationExpirationNetwork `json:"network,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization_expiration`.
	Type CardPaymentElementsCardAuthorizationExpirationType `json:"type,required"`
	JSON cardPaymentElementsCardAuthorizationExpirationJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationExpirationJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardAuthorizationExpiration]
type cardPaymentElementsCardAuthorizationExpirationJSON struct {
	ID                  apijson.Field
	CardAuthorizationID apijson.Field
	Currency            apijson.Field
	ExpiredAmount       apijson.Field
	Network             apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationExpirationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the reversal's
// currency.
type CardPaymentElementsCardAuthorizationExpirationCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardAuthorizationExpirationCurrencyCad CardPaymentElementsCardAuthorizationExpirationCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardAuthorizationExpirationCurrencyChf CardPaymentElementsCardAuthorizationExpirationCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardAuthorizationExpirationCurrencyEur CardPaymentElementsCardAuthorizationExpirationCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardAuthorizationExpirationCurrencyGbp CardPaymentElementsCardAuthorizationExpirationCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardAuthorizationExpirationCurrencyJpy CardPaymentElementsCardAuthorizationExpirationCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardAuthorizationExpirationCurrencyUsd CardPaymentElementsCardAuthorizationExpirationCurrency = "USD"
)

func (r CardPaymentElementsCardAuthorizationExpirationCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationExpirationCurrencyCad, CardPaymentElementsCardAuthorizationExpirationCurrencyChf, CardPaymentElementsCardAuthorizationExpirationCurrencyEur, CardPaymentElementsCardAuthorizationExpirationCurrencyGbp, CardPaymentElementsCardAuthorizationExpirationCurrencyJpy, CardPaymentElementsCardAuthorizationExpirationCurrencyUsd:
		return true
	}
	return false
}

// The card network used to process this card authorization.
type CardPaymentElementsCardAuthorizationExpirationNetwork string

const (
	// Visa
	CardPaymentElementsCardAuthorizationExpirationNetworkVisa CardPaymentElementsCardAuthorizationExpirationNetwork = "visa"
)

func (r CardPaymentElementsCardAuthorizationExpirationNetwork) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationExpirationNetworkVisa:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_authorization_expiration`.
type CardPaymentElementsCardAuthorizationExpirationType string

const (
	CardPaymentElementsCardAuthorizationExpirationTypeCardAuthorizationExpiration CardPaymentElementsCardAuthorizationExpirationType = "card_authorization_expiration"
)

func (r CardPaymentElementsCardAuthorizationExpirationType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationExpirationTypeCardAuthorizationExpiration:
		return true
	}
	return false
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type CardPaymentElementsCardDecline struct {
	// The Card Decline identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardDeclineActioner `json:"actioner,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardPaymentElementsCardDeclineCurrency `json:"currency,required"`
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
	NetworkDetails CardPaymentElementsCardDeclineNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardDeclineNetworkIdentifiers `json:"network_identifiers,required"`
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
	ProcessingCategory CardPaymentElementsCardDeclineProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason CardPaymentElementsCardDeclineReason `json:"reason,required"`
	// Fields related to verification of cardholder-provided values.
	Verification CardPaymentElementsCardDeclineVerification `json:"verification,required"`
	JSON         cardPaymentElementsCardDeclineJSON         `json:"-"`
}

// cardPaymentElementsCardDeclineJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardDecline]
type cardPaymentElementsCardDeclineJSON struct {
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

func (r *CardPaymentElementsCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type CardPaymentElementsCardDeclineActioner string

const (
	// This object was actioned by the user through a real-time decision.
	CardPaymentElementsCardDeclineActionerUser CardPaymentElementsCardDeclineActioner = "user"
	// This object was actioned by Increase without user intervention.
	CardPaymentElementsCardDeclineActionerIncrease CardPaymentElementsCardDeclineActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	CardPaymentElementsCardDeclineActionerNetwork CardPaymentElementsCardDeclineActioner = "network"
)

func (r CardPaymentElementsCardDeclineActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineActionerUser, CardPaymentElementsCardDeclineActionerIncrease, CardPaymentElementsCardDeclineActionerNetwork:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type CardPaymentElementsCardDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardDeclineCurrencyCad CardPaymentElementsCardDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardDeclineCurrencyChf CardPaymentElementsCardDeclineCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardDeclineCurrencyEur CardPaymentElementsCardDeclineCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardDeclineCurrencyGbp CardPaymentElementsCardDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardDeclineCurrencyJpy CardPaymentElementsCardDeclineCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardDeclineCurrencyUsd CardPaymentElementsCardDeclineCurrency = "USD"
)

func (r CardPaymentElementsCardDeclineCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineCurrencyCad, CardPaymentElementsCardDeclineCurrencyChf, CardPaymentElementsCardDeclineCurrencyEur, CardPaymentElementsCardDeclineCurrencyGbp, CardPaymentElementsCardDeclineCurrencyJpy, CardPaymentElementsCardDeclineCurrencyUsd:
		return true
	}
	return false
}

// Fields specific to the `network`.
type CardPaymentElementsCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category CardPaymentElementsCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa CardPaymentElementsCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON cardPaymentElementsCardDeclineNetworkDetailsJSON `json:"-"`
}

// cardPaymentElementsCardDeclineNetworkDetailsJSON contains the JSON metadata for
// the struct [CardPaymentElementsCardDeclineNetworkDetails]
type cardPaymentElementsCardDeclineNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type CardPaymentElementsCardDeclineNetworkDetailsCategory string

const (
	// Visa
	CardPaymentElementsCardDeclineNetworkDetailsCategoryVisa CardPaymentElementsCardDeclineNetworkDetailsCategory = "visa"
)

func (r CardPaymentElementsCardDeclineNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type CardPaymentElementsCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardPaymentElementsCardDeclineNetworkDetailsVisaJSON                    `json:"-"`
}

// cardPaymentElementsCardDeclineNetworkDetailsVisaJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardDeclineNetworkDetailsVisa]
type cardPaymentElementsCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardPaymentElementsCardDeclineNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                               `json:"transaction_id,required,nullable"`
	JSON          cardPaymentElementsCardDeclineNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardDeclineNetworkIdentifiersJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardDeclineNetworkIdentifiers]
type cardPaymentElementsCardDeclineNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type CardPaymentElementsCardDeclineProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	CardPaymentElementsCardDeclineProcessingCategoryAccountFunding CardPaymentElementsCardDeclineProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	CardPaymentElementsCardDeclineProcessingCategoryAutomaticFuelDispenser CardPaymentElementsCardDeclineProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	CardPaymentElementsCardDeclineProcessingCategoryBillPayment CardPaymentElementsCardDeclineProcessingCategory = "bill_payment"
	// A regular purchase.
	CardPaymentElementsCardDeclineProcessingCategoryPurchase CardPaymentElementsCardDeclineProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	CardPaymentElementsCardDeclineProcessingCategoryQuasiCash CardPaymentElementsCardDeclineProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	CardPaymentElementsCardDeclineProcessingCategoryRefund CardPaymentElementsCardDeclineProcessingCategory = "refund"
)

func (r CardPaymentElementsCardDeclineProcessingCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineProcessingCategoryAccountFunding, CardPaymentElementsCardDeclineProcessingCategoryAutomaticFuelDispenser, CardPaymentElementsCardDeclineProcessingCategoryBillPayment, CardPaymentElementsCardDeclineProcessingCategoryPurchase, CardPaymentElementsCardDeclineProcessingCategoryQuasiCash, CardPaymentElementsCardDeclineProcessingCategoryRefund:
		return true
	}
	return false
}

// Why the transaction was declined.
type CardPaymentElementsCardDeclineReason string

const (
	// The Card was not active.
	CardPaymentElementsCardDeclineReasonCardNotActive CardPaymentElementsCardDeclineReason = "card_not_active"
	// The Physical Card was not active.
	CardPaymentElementsCardDeclineReasonPhysicalCardNotActive CardPaymentElementsCardDeclineReason = "physical_card_not_active"
	// The account's entity was not active.
	CardPaymentElementsCardDeclineReasonEntityNotActive CardPaymentElementsCardDeclineReason = "entity_not_active"
	// The account was inactive.
	CardPaymentElementsCardDeclineReasonGroupLocked CardPaymentElementsCardDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	CardPaymentElementsCardDeclineReasonInsufficientFunds CardPaymentElementsCardDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	CardPaymentElementsCardDeclineReasonCvv2Mismatch CardPaymentElementsCardDeclineReason = "cvv2_mismatch"
	// The given expiration date did not match the card's value. Only applies when a
	// CVV2 is present.
	CardPaymentElementsCardDeclineReasonCardExpirationMismatch CardPaymentElementsCardDeclineReason = "card_expiration_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	CardPaymentElementsCardDeclineReasonTransactionNotAllowed CardPaymentElementsCardDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by a Limit.
	CardPaymentElementsCardDeclineReasonBreachesLimit CardPaymentElementsCardDeclineReason = "breaches_limit"
	// Your application declined the transaction via webhook.
	CardPaymentElementsCardDeclineReasonWebhookDeclined CardPaymentElementsCardDeclineReason = "webhook_declined"
	// Your application webhook did not respond without the required timeout.
	CardPaymentElementsCardDeclineReasonWebhookTimedOut CardPaymentElementsCardDeclineReason = "webhook_timed_out"
	// Declined by stand-in processing.
	CardPaymentElementsCardDeclineReasonDeclinedByStandInProcessing CardPaymentElementsCardDeclineReason = "declined_by_stand_in_processing"
	// The card read had an invalid CVV, dCVV, or authorization request cryptogram.
	CardPaymentElementsCardDeclineReasonInvalidPhysicalCard CardPaymentElementsCardDeclineReason = "invalid_physical_card"
	// The original card authorization for this incremental authorization does not
	// exist.
	CardPaymentElementsCardDeclineReasonMissingOriginalAuthorization CardPaymentElementsCardDeclineReason = "missing_original_authorization"
	// The transaction was suspected to be fraudulent. Please reach out to
	// support@increase.com for more information.
	CardPaymentElementsCardDeclineReasonSuspectedFraud CardPaymentElementsCardDeclineReason = "suspected_fraud"
)

func (r CardPaymentElementsCardDeclineReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineReasonCardNotActive, CardPaymentElementsCardDeclineReasonPhysicalCardNotActive, CardPaymentElementsCardDeclineReasonEntityNotActive, CardPaymentElementsCardDeclineReasonGroupLocked, CardPaymentElementsCardDeclineReasonInsufficientFunds, CardPaymentElementsCardDeclineReasonCvv2Mismatch, CardPaymentElementsCardDeclineReasonCardExpirationMismatch, CardPaymentElementsCardDeclineReasonTransactionNotAllowed, CardPaymentElementsCardDeclineReasonBreachesLimit, CardPaymentElementsCardDeclineReasonWebhookDeclined, CardPaymentElementsCardDeclineReasonWebhookTimedOut, CardPaymentElementsCardDeclineReasonDeclinedByStandInProcessing, CardPaymentElementsCardDeclineReasonInvalidPhysicalCard, CardPaymentElementsCardDeclineReasonMissingOriginalAuthorization, CardPaymentElementsCardDeclineReasonSuspectedFraud:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type CardPaymentElementsCardDeclineVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode CardPaymentElementsCardDeclineVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress CardPaymentElementsCardDeclineVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              cardPaymentElementsCardDeclineVerificationJSON              `json:"-"`
}

// cardPaymentElementsCardDeclineVerificationJSON contains the JSON metadata for
// the struct [CardPaymentElementsCardDeclineVerification]
type cardPaymentElementsCardDeclineVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type CardPaymentElementsCardDeclineVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   cardPaymentElementsCardDeclineVerificationCardVerificationCodeJSON   `json:"-"`
}

// cardPaymentElementsCardDeclineVerificationCardVerificationCodeJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardDeclineVerificationCardVerificationCode]
type cardPaymentElementsCardDeclineVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultNotChecked CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultMatch CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultNoMatch CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult = "no_match"
)

func (r CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultNotChecked, CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultMatch, CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type CardPaymentElementsCardDeclineVerificationCardholderAddress struct {
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
	Result CardPaymentElementsCardDeclineVerificationCardholderAddressResult `json:"result,required"`
	JSON   cardPaymentElementsCardDeclineVerificationCardholderAddressJSON   `json:"-"`
}

// cardPaymentElementsCardDeclineVerificationCardholderAddressJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardDeclineVerificationCardholderAddress]
type cardPaymentElementsCardDeclineVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type CardPaymentElementsCardDeclineVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultNotChecked CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultMatch CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultNoMatch CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "no_match"
)

func (r CardPaymentElementsCardDeclineVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineVerificationCardholderAddressResultNotChecked, CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// A Card Fuel Confirmation object. This field will be present in the JSON response
// if and only if `category` is equal to `card_fuel_confirmation`.
type CardPaymentElementsCardFuelConfirmation struct {
	// The Card Fuel Confirmation identifier.
	ID string `json:"id,required"`
	// The identifier for the Card Authorization this updates.
	CardAuthorizationID string `json:"card_authorization_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the increment's
	// currency.
	Currency CardPaymentElementsCardFuelConfirmationCurrency `json:"currency,required"`
	// The card network used to process this card authorization.
	Network CardPaymentElementsCardFuelConfirmationNetwork `json:"network,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardFuelConfirmationNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Card Fuel
	// Confirmation.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_fuel_confirmation`.
	Type CardPaymentElementsCardFuelConfirmationType `json:"type,required"`
	// The updated authorization amount after this fuel confirmation, in the minor unit
	// of the transaction's currency. For dollars, for example, this is cents.
	UpdatedAuthorizationAmount int64                                       `json:"updated_authorization_amount,required"`
	JSON                       cardPaymentElementsCardFuelConfirmationJSON `json:"-"`
}

// cardPaymentElementsCardFuelConfirmationJSON contains the JSON metadata for the
// struct [CardPaymentElementsCardFuelConfirmation]
type cardPaymentElementsCardFuelConfirmationJSON struct {
	ID                         apijson.Field
	CardAuthorizationID        apijson.Field
	Currency                   apijson.Field
	Network                    apijson.Field
	NetworkIdentifiers         apijson.Field
	PendingTransactionID       apijson.Field
	Type                       apijson.Field
	UpdatedAuthorizationAmount apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardPaymentElementsCardFuelConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardFuelConfirmationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the increment's
// currency.
type CardPaymentElementsCardFuelConfirmationCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardFuelConfirmationCurrencyCad CardPaymentElementsCardFuelConfirmationCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardFuelConfirmationCurrencyChf CardPaymentElementsCardFuelConfirmationCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardFuelConfirmationCurrencyEur CardPaymentElementsCardFuelConfirmationCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardFuelConfirmationCurrencyGbp CardPaymentElementsCardFuelConfirmationCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardFuelConfirmationCurrencyJpy CardPaymentElementsCardFuelConfirmationCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardFuelConfirmationCurrencyUsd CardPaymentElementsCardFuelConfirmationCurrency = "USD"
)

func (r CardPaymentElementsCardFuelConfirmationCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardFuelConfirmationCurrencyCad, CardPaymentElementsCardFuelConfirmationCurrencyChf, CardPaymentElementsCardFuelConfirmationCurrencyEur, CardPaymentElementsCardFuelConfirmationCurrencyGbp, CardPaymentElementsCardFuelConfirmationCurrencyJpy, CardPaymentElementsCardFuelConfirmationCurrencyUsd:
		return true
	}
	return false
}

// The card network used to process this card authorization.
type CardPaymentElementsCardFuelConfirmationNetwork string

const (
	// Visa
	CardPaymentElementsCardFuelConfirmationNetworkVisa CardPaymentElementsCardFuelConfirmationNetwork = "visa"
)

func (r CardPaymentElementsCardFuelConfirmationNetwork) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardFuelConfirmationNetworkVisa:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardPaymentElementsCardFuelConfirmationNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                        `json:"transaction_id,required,nullable"`
	JSON          cardPaymentElementsCardFuelConfirmationNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardFuelConfirmationNetworkIdentifiersJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardFuelConfirmationNetworkIdentifiers]
type cardPaymentElementsCardFuelConfirmationNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPaymentElementsCardFuelConfirmationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardFuelConfirmationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_fuel_confirmation`.
type CardPaymentElementsCardFuelConfirmationType string

const (
	CardPaymentElementsCardFuelConfirmationTypeCardFuelConfirmation CardPaymentElementsCardFuelConfirmationType = "card_fuel_confirmation"
)

func (r CardPaymentElementsCardFuelConfirmationType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardFuelConfirmationTypeCardFuelConfirmation:
		return true
	}
	return false
}

// A Card Increment object. This field will be present in the JSON response if and
// only if `category` is equal to `card_increment`.
type CardPaymentElementsCardIncrement struct {
	// The Card Increment identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardIncrementActioner `json:"actioner,required"`
	// The amount of this increment in the minor unit of the transaction's currency.
	// For dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier for the Card Authorization this increments.
	CardAuthorizationID string `json:"card_authorization_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the increment's
	// currency.
	Currency CardPaymentElementsCardIncrementCurrency `json:"currency,required"`
	// The card network used to process this card authorization.
	Network CardPaymentElementsCardIncrementNetwork `json:"network,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardIncrementNetworkIdentifiers `json:"network_identifiers,required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score,required,nullable"`
	// The identifier of the Pending Transaction associated with this Card Increment.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// incremental authorization.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_increment`.
	Type CardPaymentElementsCardIncrementType `json:"type,required"`
	// The updated authorization amount after this increment, in the minor unit of the
	// transaction's currency. For dollars, for example, this is cents.
	UpdatedAuthorizationAmount int64                                `json:"updated_authorization_amount,required"`
	JSON                       cardPaymentElementsCardIncrementJSON `json:"-"`
}

// cardPaymentElementsCardIncrementJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardIncrement]
type cardPaymentElementsCardIncrementJSON struct {
	ID                         apijson.Field
	Actioner                   apijson.Field
	Amount                     apijson.Field
	CardAuthorizationID        apijson.Field
	Currency                   apijson.Field
	Network                    apijson.Field
	NetworkIdentifiers         apijson.Field
	NetworkRiskScore           apijson.Field
	PendingTransactionID       apijson.Field
	RealTimeDecisionID         apijson.Field
	Type                       apijson.Field
	UpdatedAuthorizationAmount apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type CardPaymentElementsCardIncrementActioner string

const (
	// This object was actioned by the user through a real-time decision.
	CardPaymentElementsCardIncrementActionerUser CardPaymentElementsCardIncrementActioner = "user"
	// This object was actioned by Increase without user intervention.
	CardPaymentElementsCardIncrementActionerIncrease CardPaymentElementsCardIncrementActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	CardPaymentElementsCardIncrementActionerNetwork CardPaymentElementsCardIncrementActioner = "network"
)

func (r CardPaymentElementsCardIncrementActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardIncrementActionerUser, CardPaymentElementsCardIncrementActionerIncrease, CardPaymentElementsCardIncrementActionerNetwork:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the increment's
// currency.
type CardPaymentElementsCardIncrementCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardIncrementCurrencyCad CardPaymentElementsCardIncrementCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardIncrementCurrencyChf CardPaymentElementsCardIncrementCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardIncrementCurrencyEur CardPaymentElementsCardIncrementCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardIncrementCurrencyGbp CardPaymentElementsCardIncrementCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardIncrementCurrencyJpy CardPaymentElementsCardIncrementCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardIncrementCurrencyUsd CardPaymentElementsCardIncrementCurrency = "USD"
)

func (r CardPaymentElementsCardIncrementCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardIncrementCurrencyCad, CardPaymentElementsCardIncrementCurrencyChf, CardPaymentElementsCardIncrementCurrencyEur, CardPaymentElementsCardIncrementCurrencyGbp, CardPaymentElementsCardIncrementCurrencyJpy, CardPaymentElementsCardIncrementCurrencyUsd:
		return true
	}
	return false
}

// The card network used to process this card authorization.
type CardPaymentElementsCardIncrementNetwork string

const (
	// Visa
	CardPaymentElementsCardIncrementNetworkVisa CardPaymentElementsCardIncrementNetwork = "visa"
)

func (r CardPaymentElementsCardIncrementNetwork) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardIncrementNetworkVisa:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardPaymentElementsCardIncrementNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                 `json:"transaction_id,required,nullable"`
	JSON          cardPaymentElementsCardIncrementNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardIncrementNetworkIdentifiersJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardIncrementNetworkIdentifiers]
type cardPaymentElementsCardIncrementNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_increment`.
type CardPaymentElementsCardIncrementType string

const (
	CardPaymentElementsCardIncrementTypeCardIncrement CardPaymentElementsCardIncrementType = "card_increment"
)

func (r CardPaymentElementsCardIncrementType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardIncrementTypeCardIncrement:
		return true
	}
	return false
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
type CardPaymentElementsCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency CardPaymentElementsCardRefundCurrency `json:"currency,required"`
	// Interchange assessed as a part of this transaciton.
	Interchange CardPaymentElementsCardRefundInterchange `json:"interchange,required,nullable"`
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
	// Network-specific identifiers for this refund.
	NetworkIdentifiers CardPaymentElementsCardRefundNetworkIdentifiers `json:"network_identifiers,required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails CardPaymentElementsCardRefundPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type CardPaymentElementsCardRefundType `json:"type,required"`
	JSON cardPaymentElementsCardRefundJSON `json:"-"`
}

// cardPaymentElementsCardRefundJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardRefund]
type cardPaymentElementsCardRefundJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	Interchange          apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	NetworkIdentifiers   apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type CardPaymentElementsCardRefundCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardRefundCurrencyCad CardPaymentElementsCardRefundCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardRefundCurrencyChf CardPaymentElementsCardRefundCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardRefundCurrencyEur CardPaymentElementsCardRefundCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardRefundCurrencyGbp CardPaymentElementsCardRefundCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardRefundCurrencyJpy CardPaymentElementsCardRefundCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardRefundCurrencyUsd CardPaymentElementsCardRefundCurrency = "USD"
)

func (r CardPaymentElementsCardRefundCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundCurrencyCad, CardPaymentElementsCardRefundCurrencyChf, CardPaymentElementsCardRefundCurrencyEur, CardPaymentElementsCardRefundCurrencyGbp, CardPaymentElementsCardRefundCurrencyJpy, CardPaymentElementsCardRefundCurrencyUsd:
		return true
	}
	return false
}

// Interchange assessed as a part of this transaciton.
type CardPaymentElementsCardRefundInterchange struct {
	// The interchange amount given as a string containing a decimal number. The amount
	// is a positive number if it is credited to Increase (e.g., settlements) and a
	// negative number if it is debited (e.g., refunds).
	Amount string `json:"amount,required"`
	// The card network specific interchange code.
	Code string `json:"code,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
	// reimbursement.
	Currency CardPaymentElementsCardRefundInterchangeCurrency `json:"currency,required"`
	JSON     cardPaymentElementsCardRefundInterchangeJSON     `json:"-"`
}

// cardPaymentElementsCardRefundInterchangeJSON contains the JSON metadata for the
// struct [CardPaymentElementsCardRefundInterchange]
type cardPaymentElementsCardRefundInterchangeJSON struct {
	Amount      apijson.Field
	Code        apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefundInterchange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundInterchangeJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
// reimbursement.
type CardPaymentElementsCardRefundInterchangeCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardRefundInterchangeCurrencyCad CardPaymentElementsCardRefundInterchangeCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardRefundInterchangeCurrencyChf CardPaymentElementsCardRefundInterchangeCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardRefundInterchangeCurrencyEur CardPaymentElementsCardRefundInterchangeCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardRefundInterchangeCurrencyGbp CardPaymentElementsCardRefundInterchangeCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardRefundInterchangeCurrencyJpy CardPaymentElementsCardRefundInterchangeCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardRefundInterchangeCurrencyUsd CardPaymentElementsCardRefundInterchangeCurrency = "USD"
)

func (r CardPaymentElementsCardRefundInterchangeCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundInterchangeCurrencyCad, CardPaymentElementsCardRefundInterchangeCurrencyChf, CardPaymentElementsCardRefundInterchangeCurrencyEur, CardPaymentElementsCardRefundInterchangeCurrencyGbp, CardPaymentElementsCardRefundInterchangeCurrencyJpy, CardPaymentElementsCardRefundInterchangeCurrencyUsd:
		return true
	}
	return false
}

// Network-specific identifiers for this refund.
type CardPaymentElementsCardRefundNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id,required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                              `json:"transaction_id,required,nullable"`
	JSON          cardPaymentElementsCardRefundNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardRefundNetworkIdentifiersJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardRefundNetworkIdentifiers]
type cardPaymentElementsCardRefundNetworkIdentifiersJSON struct {
	AcquirerBusinessID      apijson.Field
	AcquirerReferenceNumber apijson.Field
	TransactionID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefundNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type CardPaymentElementsCardRefundPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental CardPaymentElementsCardRefundPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging CardPaymentElementsCardRefundPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel CardPaymentElementsCardRefundPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   cardPaymentElementsCardRefundPurchaseDetailsJSON   `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsJSON contains the JSON metadata for
// the struct [CardPaymentElementsCardRefundPurchaseDetails]
type cardPaymentElementsCardRefundPurchaseDetailsJSON struct {
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

func (r *CardPaymentElementsCardRefundPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsJSON) RawJSON() string {
	return r.raw
}

// Fields specific to car rentals.
type CardPaymentElementsCardRefundPurchaseDetailsCarRental struct {
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
	ExtraCharges CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	WeeklyRentalRateCurrency string                                                    `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     cardPaymentElementsCardRefundPurchaseDetailsCarRentalJSON `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsCarRentalJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardRefundPurchaseDetailsCarRental]
type cardPaymentElementsCardRefundPurchaseDetailsCarRentalJSON struct {
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

func (r *CardPaymentElementsCardRefundPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsCarRentalJSON) RawJSON() string {
	return r.raw
}

// Additional charges (gas, late fee, etc.) being billed.
type CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesGas CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge, CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesGas, CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage, CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn, CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee, CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation:
		return true
	}
	return false
}

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable, CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle:
		return true
	}
	return false
}

// Fields specific to lodging.
type CardPaymentElementsCardRefundPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	TotalTaxCurrency string                                                  `json:"total_tax_currency,required,nullable"`
	JSON             cardPaymentElementsCardRefundPurchaseDetailsLodgingJSON `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsLodgingJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardRefundPurchaseDetailsLodging]
type cardPaymentElementsCardRefundPurchaseDetailsLodgingJSON struct {
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

func (r *CardPaymentElementsCardRefundPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsLodgingJSON) RawJSON() string {
	return r.raw
}

// Additional charges (phone, late check-out, etc.) being billed.
type CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesRestaurant CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesGiftShop CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesMiniBar CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesTelephone CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesOther CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesLaundry CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge, CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesRestaurant, CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesGiftShop, CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesMiniBar, CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesTelephone, CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesOther, CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesLaundry:
		return true
	}
	return false
}

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable, CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow:
		return true
	}
	return false
}

// The format of the purchase identifier.
type CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText, CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber, CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber, CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber, CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber:
		return true
	}
	return false
}

// Fields specific to travel.
type CardPaymentElementsCardRefundPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary CardPaymentElementsCardRefundPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []CardPaymentElementsCardRefundPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     cardPaymentElementsCardRefundPurchaseDetailsTravelJSON      `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsTravelJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardRefundPurchaseDetailsTravel]
type cardPaymentElementsCardRefundPurchaseDetailsTravelJSON struct {
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

func (r *CardPaymentElementsCardRefundPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsTravelJSON) RawJSON() string {
	return r.raw
}

// Ancillary purchases in addition to the airfare.
type CardPaymentElementsCardRefundPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string                                                          `json:"ticket_document_number,required,nullable"`
	JSON                 cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryJSON `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardRefundPurchaseDetailsTravelAncillary]
type cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefundPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryJSON) RawJSON() string {
	return r.raw
}

// Indicates the reason for a credit to the cardholder.
type CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther:
		return true
	}
	return false
}

type CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                                 `json:"sub_category,required,nullable"`
	JSON        cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServiceJSON `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServiceJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryService]
type cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServiceJSON) RawJSON() string {
	return r.raw
}

// Category of the ancillary service.
type CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades, CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi:
		return true
	}
	return false
}

// Indicates the reason for a credit to the cardholder.
type CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit, CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation, CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther, CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket:
		return true
	}
	return false
}

// Indicates whether this ticket is non-refundable.
type CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions, CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket:
		return true
	}
	return false
}

// Indicates why a ticket was changed.
type CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone, CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket, CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket:
		return true
	}
	return false
}

type CardPaymentElementsCardRefundPurchaseDetailsTravelTripLeg struct {
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
	StopOverCode CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         cardPaymentElementsCardRefundPurchaseDetailsTravelTripLegJSON          `json:"-"`
}

// cardPaymentElementsCardRefundPurchaseDetailsTravelTripLegJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardRefundPurchaseDetailsTravelTripLeg]
type cardPaymentElementsCardRefundPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefundPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundPurchaseDetailsTravelTripLegJSON) RawJSON() string {
	return r.raw
}

// Indicates whether a stopover is allowed on this ticket.
type CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

func (r CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone, CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed, CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
type CardPaymentElementsCardRefundType string

const (
	CardPaymentElementsCardRefundTypeCardRefund CardPaymentElementsCardRefundType = "card_refund"
)

func (r CardPaymentElementsCardRefundType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundTypeCardRefund:
		return true
	}
	return false
}

// A Card Reversal object. This field will be present in the JSON response if and
// only if `category` is equal to `card_reversal`.
type CardPaymentElementsCardReversal struct {
	// The Card Reversal identifier.
	ID string `json:"id,required"`
	// The identifier for the Card Authorization this reverses.
	CardAuthorizationID string `json:"card_authorization_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the reversal's
	// currency.
	Currency CardPaymentElementsCardReversalCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
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
	// The card network used to process this card authorization.
	Network CardPaymentElementsCardReversalNetwork `json:"network,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardReversalNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Card Reversal.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount of this reversal in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	ReversalAmount int64 `json:"reversal_amount,required"`
	// Why this reversal was initiated.
	ReversalReason CardPaymentElementsCardReversalReversalReason `json:"reversal_reason,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_reversal`.
	Type CardPaymentElementsCardReversalType `json:"type,required"`
	// The amount left pending on the Card Authorization in the minor unit of the
	// transaction's currency. For dollars, for example, this is cents.
	UpdatedAuthorizationAmount int64                               `json:"updated_authorization_amount,required"`
	JSON                       cardPaymentElementsCardReversalJSON `json:"-"`
}

// cardPaymentElementsCardReversalJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardReversal]
type cardPaymentElementsCardReversalJSON struct {
	ID                         apijson.Field
	CardAuthorizationID        apijson.Field
	Currency                   apijson.Field
	MerchantAcceptorID         apijson.Field
	MerchantCategoryCode       apijson.Field
	MerchantCity               apijson.Field
	MerchantCountry            apijson.Field
	MerchantDescriptor         apijson.Field
	MerchantPostalCode         apijson.Field
	MerchantState              apijson.Field
	Network                    apijson.Field
	NetworkIdentifiers         apijson.Field
	PendingTransactionID       apijson.Field
	ReversalAmount             apijson.Field
	ReversalReason             apijson.Field
	Type                       apijson.Field
	UpdatedAuthorizationAmount apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardPaymentElementsCardReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardReversalJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the reversal's
// currency.
type CardPaymentElementsCardReversalCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardReversalCurrencyCad CardPaymentElementsCardReversalCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardReversalCurrencyChf CardPaymentElementsCardReversalCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardReversalCurrencyEur CardPaymentElementsCardReversalCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardReversalCurrencyGbp CardPaymentElementsCardReversalCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardReversalCurrencyJpy CardPaymentElementsCardReversalCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardReversalCurrencyUsd CardPaymentElementsCardReversalCurrency = "USD"
)

func (r CardPaymentElementsCardReversalCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardReversalCurrencyCad, CardPaymentElementsCardReversalCurrencyChf, CardPaymentElementsCardReversalCurrencyEur, CardPaymentElementsCardReversalCurrencyGbp, CardPaymentElementsCardReversalCurrencyJpy, CardPaymentElementsCardReversalCurrencyUsd:
		return true
	}
	return false
}

// The card network used to process this card authorization.
type CardPaymentElementsCardReversalNetwork string

const (
	// Visa
	CardPaymentElementsCardReversalNetworkVisa CardPaymentElementsCardReversalNetwork = "visa"
)

func (r CardPaymentElementsCardReversalNetwork) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardReversalNetworkVisa:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardPaymentElementsCardReversalNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                `json:"transaction_id,required,nullable"`
	JSON          cardPaymentElementsCardReversalNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardReversalNetworkIdentifiersJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardReversalNetworkIdentifiers]
type cardPaymentElementsCardReversalNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPaymentElementsCardReversalNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardReversalNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// Why this reversal was initiated.
type CardPaymentElementsCardReversalReversalReason string

const (
	// The Card Reversal was initiated at the customer's request.
	CardPaymentElementsCardReversalReversalReasonReversedByCustomer CardPaymentElementsCardReversalReversalReason = "reversed_by_customer"
	// The Card Reversal was initiated by the network or acquirer.
	CardPaymentElementsCardReversalReversalReasonReversedByNetworkOrAcquirer CardPaymentElementsCardReversalReversalReason = "reversed_by_network_or_acquirer"
	// The Card Reversal was initiated by the point of sale device.
	CardPaymentElementsCardReversalReversalReasonReversedByPointOfSale CardPaymentElementsCardReversalReversalReason = "reversed_by_point_of_sale"
	// The Card Reversal was a partial reversal, for any reason.
	CardPaymentElementsCardReversalReversalReasonPartialReversal CardPaymentElementsCardReversalReversalReason = "partial_reversal"
)

func (r CardPaymentElementsCardReversalReversalReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardReversalReversalReasonReversedByCustomer, CardPaymentElementsCardReversalReversalReasonReversedByNetworkOrAcquirer, CardPaymentElementsCardReversalReversalReasonReversedByPointOfSale, CardPaymentElementsCardReversalReversalReasonPartialReversal:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_reversal`.
type CardPaymentElementsCardReversalType string

const (
	CardPaymentElementsCardReversalTypeCardReversal CardPaymentElementsCardReversalType = "card_reversal"
)

func (r CardPaymentElementsCardReversalType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardReversalTypeCardReversal:
		return true
	}
	return false
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
type CardPaymentElementsCardSettlement struct {
	// The Card Settlement identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The Card Authorization that was created prior to this Card Settlement, if one
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency CardPaymentElementsCardSettlementCurrency `json:"currency,required"`
	// Interchange assessed as a part of this transaciton.
	Interchange CardPaymentElementsCardSettlementInterchange `json:"interchange,required,nullable"`
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
	// Network-specific identifiers for this refund.
	NetworkIdentifiers CardPaymentElementsCardSettlementNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails CardPaymentElementsCardSettlementPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type CardPaymentElementsCardSettlementType `json:"type,required"`
	JSON cardPaymentElementsCardSettlementJSON `json:"-"`
}

// cardPaymentElementsCardSettlementJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardSettlement]
type cardPaymentElementsCardSettlementJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardAuthorization    apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	Interchange          apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	NetworkIdentifiers   apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type CardPaymentElementsCardSettlementCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardSettlementCurrencyCad CardPaymentElementsCardSettlementCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardSettlementCurrencyChf CardPaymentElementsCardSettlementCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardSettlementCurrencyEur CardPaymentElementsCardSettlementCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardSettlementCurrencyGbp CardPaymentElementsCardSettlementCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardSettlementCurrencyJpy CardPaymentElementsCardSettlementCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardSettlementCurrencyUsd CardPaymentElementsCardSettlementCurrency = "USD"
)

func (r CardPaymentElementsCardSettlementCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementCurrencyCad, CardPaymentElementsCardSettlementCurrencyChf, CardPaymentElementsCardSettlementCurrencyEur, CardPaymentElementsCardSettlementCurrencyGbp, CardPaymentElementsCardSettlementCurrencyJpy, CardPaymentElementsCardSettlementCurrencyUsd:
		return true
	}
	return false
}

// Interchange assessed as a part of this transaciton.
type CardPaymentElementsCardSettlementInterchange struct {
	// The interchange amount given as a string containing a decimal number. The amount
	// is a positive number if it is credited to Increase (e.g., settlements) and a
	// negative number if it is debited (e.g., refunds).
	Amount string `json:"amount,required"`
	// The card network specific interchange code.
	Code string `json:"code,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
	// reimbursement.
	Currency CardPaymentElementsCardSettlementInterchangeCurrency `json:"currency,required"`
	JSON     cardPaymentElementsCardSettlementInterchangeJSON     `json:"-"`
}

// cardPaymentElementsCardSettlementInterchangeJSON contains the JSON metadata for
// the struct [CardPaymentElementsCardSettlementInterchange]
type cardPaymentElementsCardSettlementInterchangeJSON struct {
	Amount      apijson.Field
	Code        apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementInterchange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementInterchangeJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
// reimbursement.
type CardPaymentElementsCardSettlementInterchangeCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardSettlementInterchangeCurrencyCad CardPaymentElementsCardSettlementInterchangeCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardSettlementInterchangeCurrencyChf CardPaymentElementsCardSettlementInterchangeCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardSettlementInterchangeCurrencyEur CardPaymentElementsCardSettlementInterchangeCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardSettlementInterchangeCurrencyGbp CardPaymentElementsCardSettlementInterchangeCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardSettlementInterchangeCurrencyJpy CardPaymentElementsCardSettlementInterchangeCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardSettlementInterchangeCurrencyUsd CardPaymentElementsCardSettlementInterchangeCurrency = "USD"
)

func (r CardPaymentElementsCardSettlementInterchangeCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementInterchangeCurrencyCad, CardPaymentElementsCardSettlementInterchangeCurrencyChf, CardPaymentElementsCardSettlementInterchangeCurrencyEur, CardPaymentElementsCardSettlementInterchangeCurrencyGbp, CardPaymentElementsCardSettlementInterchangeCurrencyJpy, CardPaymentElementsCardSettlementInterchangeCurrencyUsd:
		return true
	}
	return false
}

// Network-specific identifiers for this refund.
type CardPaymentElementsCardSettlementNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id,required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                  `json:"transaction_id,required,nullable"`
	JSON          cardPaymentElementsCardSettlementNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardSettlementNetworkIdentifiersJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardSettlementNetworkIdentifiers]
type cardPaymentElementsCardSettlementNetworkIdentifiersJSON struct {
	AcquirerBusinessID      apijson.Field
	AcquirerReferenceNumber apijson.Field
	TransactionID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type CardPaymentElementsCardSettlementPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental CardPaymentElementsCardSettlementPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging CardPaymentElementsCardSettlementPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel CardPaymentElementsCardSettlementPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   cardPaymentElementsCardSettlementPurchaseDetailsJSON   `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardSettlementPurchaseDetails]
type cardPaymentElementsCardSettlementPurchaseDetailsJSON struct {
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

func (r *CardPaymentElementsCardSettlementPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsJSON) RawJSON() string {
	return r.raw
}

// Fields specific to car rentals.
type CardPaymentElementsCardSettlementPurchaseDetailsCarRental struct {
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
	ExtraCharges CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	WeeklyRentalRateCurrency string                                                        `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     cardPaymentElementsCardSettlementPurchaseDetailsCarRentalJSON `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsCarRentalJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardSettlementPurchaseDetailsCarRental]
type cardPaymentElementsCardSettlementPurchaseDetailsCarRentalJSON struct {
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

func (r *CardPaymentElementsCardSettlementPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsCarRentalJSON) RawJSON() string {
	return r.raw
}

// Additional charges (gas, late fee, etc.) being billed.
type CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesGas CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge, CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesGas, CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage, CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn, CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee, CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation:
		return true
	}
	return false
}

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable, CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle:
		return true
	}
	return false
}

// Fields specific to lodging.
type CardPaymentElementsCardSettlementPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	TotalTaxCurrency string                                                      `json:"total_tax_currency,required,nullable"`
	JSON             cardPaymentElementsCardSettlementPurchaseDetailsLodgingJSON `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsLodgingJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardSettlementPurchaseDetailsLodging]
type cardPaymentElementsCardSettlementPurchaseDetailsLodgingJSON struct {
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

func (r *CardPaymentElementsCardSettlementPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsLodgingJSON) RawJSON() string {
	return r.raw
}

// Additional charges (phone, late check-out, etc.) being billed.
type CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesTelephone CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesOther CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesLaundry CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge, CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant, CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop, CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar, CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesTelephone, CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesOther, CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesLaundry:
		return true
	}
	return false
}

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable, CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow:
		return true
	}
	return false
}

// The format of the purchase identifier.
type CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText, CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber, CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber, CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber, CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber:
		return true
	}
	return false
}

// Fields specific to travel.
type CardPaymentElementsCardSettlementPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     cardPaymentElementsCardSettlementPurchaseDetailsTravelJSON      `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsTravelJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardSettlementPurchaseDetailsTravel]
type cardPaymentElementsCardSettlementPurchaseDetailsTravelJSON struct {
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

func (r *CardPaymentElementsCardSettlementPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsTravelJSON) RawJSON() string {
	return r.raw
}

// Ancillary purchases in addition to the airfare.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string                                                              `json:"ticket_document_number,required,nullable"`
	JSON                 cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryJSON `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillary]
type cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryJSON) RawJSON() string {
	return r.raw
}

// Indicates the reason for a credit to the cardholder.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther:
		return true
	}
	return false
}

type CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                                     `json:"sub_category,required,nullable"`
	JSON        cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServiceJSON `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryService]
type cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServiceJSON) RawJSON() string {
	return r.raw
}

// Category of the ancillary service.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades, CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi:
		return true
	}
	return false
}

// Indicates the reason for a credit to the cardholder.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit, CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation, CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther, CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket:
		return true
	}
	return false
}

// Indicates whether this ticket is non-refundable.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions, CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket:
		return true
	}
	return false
}

// Indicates why a ticket was changed.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone, CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket, CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket:
		return true
	}
	return false
}

type CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLeg struct {
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
	StopOverCode CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         cardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegJSON          `json:"-"`
}

// cardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLeg]
type cardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegJSON) RawJSON() string {
	return r.raw
}

// Indicates whether a stopover is allowed on this ticket.
type CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone, CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed, CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type CardPaymentElementsCardSettlementType string

const (
	CardPaymentElementsCardSettlementTypeCardSettlement CardPaymentElementsCardSettlementType = "card_settlement"
)

func (r CardPaymentElementsCardSettlementType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementTypeCardSettlement:
		return true
	}
	return false
}

// A Card Validation object. This field will be present in the JSON response if and
// only if `category` is equal to `card_validation`.
type CardPaymentElementsCardValidation struct {
	// The Card Validation identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardValidationActioner `json:"actioner,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardPaymentElementsCardValidationCurrency `json:"currency,required"`
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
	NetworkDetails CardPaymentElementsCardValidationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardValidationNetworkIdentifiers `json:"network_identifiers,required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score,required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_validation`.
	Type CardPaymentElementsCardValidationType `json:"type,required"`
	// Fields related to verification of cardholder-provided values.
	Verification CardPaymentElementsCardValidationVerification `json:"verification,required"`
	JSON         cardPaymentElementsCardValidationJSON         `json:"-"`
}

// cardPaymentElementsCardValidationJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardValidation]
type cardPaymentElementsCardValidationJSON struct {
	ID                   apijson.Field
	Actioner             apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantPostalCode   apijson.Field
	MerchantState        apijson.Field
	NetworkDetails       apijson.Field
	NetworkIdentifiers   apijson.Field
	NetworkRiskScore     apijson.Field
	PhysicalCardID       apijson.Field
	RealTimeDecisionID   apijson.Field
	Type                 apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type CardPaymentElementsCardValidationActioner string

const (
	// This object was actioned by the user through a real-time decision.
	CardPaymentElementsCardValidationActionerUser CardPaymentElementsCardValidationActioner = "user"
	// This object was actioned by Increase without user intervention.
	CardPaymentElementsCardValidationActionerIncrease CardPaymentElementsCardValidationActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	CardPaymentElementsCardValidationActionerNetwork CardPaymentElementsCardValidationActioner = "network"
)

func (r CardPaymentElementsCardValidationActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationActionerUser, CardPaymentElementsCardValidationActionerIncrease, CardPaymentElementsCardValidationActionerNetwork:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CardPaymentElementsCardValidationCurrency string

const (
	// Canadian Dollar (CAD)
	CardPaymentElementsCardValidationCurrencyCad CardPaymentElementsCardValidationCurrency = "CAD"
	// Swiss Franc (CHF)
	CardPaymentElementsCardValidationCurrencyChf CardPaymentElementsCardValidationCurrency = "CHF"
	// Euro (EUR)
	CardPaymentElementsCardValidationCurrencyEur CardPaymentElementsCardValidationCurrency = "EUR"
	// British Pound (GBP)
	CardPaymentElementsCardValidationCurrencyGbp CardPaymentElementsCardValidationCurrency = "GBP"
	// Japanese Yen (JPY)
	CardPaymentElementsCardValidationCurrencyJpy CardPaymentElementsCardValidationCurrency = "JPY"
	// US Dollar (USD)
	CardPaymentElementsCardValidationCurrencyUsd CardPaymentElementsCardValidationCurrency = "USD"
)

func (r CardPaymentElementsCardValidationCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationCurrencyCad, CardPaymentElementsCardValidationCurrencyChf, CardPaymentElementsCardValidationCurrencyEur, CardPaymentElementsCardValidationCurrencyGbp, CardPaymentElementsCardValidationCurrencyJpy, CardPaymentElementsCardValidationCurrencyUsd:
		return true
	}
	return false
}

// Fields specific to the `network`.
type CardPaymentElementsCardValidationNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category CardPaymentElementsCardValidationNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa CardPaymentElementsCardValidationNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON cardPaymentElementsCardValidationNetworkDetailsJSON `json:"-"`
}

// cardPaymentElementsCardValidationNetworkDetailsJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardValidationNetworkDetails]
type cardPaymentElementsCardValidationNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type CardPaymentElementsCardValidationNetworkDetailsCategory string

const (
	// Visa
	CardPaymentElementsCardValidationNetworkDetailsCategoryVisa CardPaymentElementsCardValidationNetworkDetailsCategory = "visa"
)

func (r CardPaymentElementsCardValidationNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type CardPaymentElementsCardValidationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardPaymentElementsCardValidationNetworkDetailsVisaJSON                    `json:"-"`
}

// cardPaymentElementsCardValidationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardValidationNetworkDetailsVisa]
type cardPaymentElementsCardValidationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorRecurring CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorInstallment CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorRecurring, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorInstallment, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeUnknown CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeManual CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactless CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeManual, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactless, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type CardPaymentElementsCardValidationNetworkIdentifiers struct {
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
	JSON          cardPaymentElementsCardValidationNetworkIdentifiersJSON `json:"-"`
}

// cardPaymentElementsCardValidationNetworkIdentifiersJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardValidationNetworkIdentifiers]
type cardPaymentElementsCardValidationNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_validation`.
type CardPaymentElementsCardValidationType string

const (
	CardPaymentElementsCardValidationTypeCardValidation CardPaymentElementsCardValidationType = "card_validation"
)

func (r CardPaymentElementsCardValidationType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationTypeCardValidation:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type CardPaymentElementsCardValidationVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode CardPaymentElementsCardValidationVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress CardPaymentElementsCardValidationVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              cardPaymentElementsCardValidationVerificationJSON              `json:"-"`
}

// cardPaymentElementsCardValidationVerificationJSON contains the JSON metadata for
// the struct [CardPaymentElementsCardValidationVerification]
type cardPaymentElementsCardValidationVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type CardPaymentElementsCardValidationVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result CardPaymentElementsCardValidationVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   cardPaymentElementsCardValidationVerificationCardVerificationCodeJSON   `json:"-"`
}

// cardPaymentElementsCardValidationVerificationCardVerificationCodeJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardValidationVerificationCardVerificationCode]
type cardPaymentElementsCardValidationVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type CardPaymentElementsCardValidationVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	CardPaymentElementsCardValidationVerificationCardVerificationCodeResultNotChecked CardPaymentElementsCardValidationVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	CardPaymentElementsCardValidationVerificationCardVerificationCodeResultMatch CardPaymentElementsCardValidationVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	CardPaymentElementsCardValidationVerificationCardVerificationCodeResultNoMatch CardPaymentElementsCardValidationVerificationCardVerificationCodeResult = "no_match"
)

func (r CardPaymentElementsCardValidationVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationVerificationCardVerificationCodeResultNotChecked, CardPaymentElementsCardValidationVerificationCardVerificationCodeResultMatch, CardPaymentElementsCardValidationVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type CardPaymentElementsCardValidationVerificationCardholderAddress struct {
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
	Result CardPaymentElementsCardValidationVerificationCardholderAddressResult `json:"result,required"`
	JSON   cardPaymentElementsCardValidationVerificationCardholderAddressJSON   `json:"-"`
}

// cardPaymentElementsCardValidationVerificationCardholderAddressJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardValidationVerificationCardholderAddress]
type cardPaymentElementsCardValidationVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type CardPaymentElementsCardValidationVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	CardPaymentElementsCardValidationVerificationCardholderAddressResultNotChecked CardPaymentElementsCardValidationVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardPaymentElementsCardValidationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch CardPaymentElementsCardValidationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch CardPaymentElementsCardValidationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	CardPaymentElementsCardValidationVerificationCardholderAddressResultMatch CardPaymentElementsCardValidationVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	CardPaymentElementsCardValidationVerificationCardholderAddressResultNoMatch CardPaymentElementsCardValidationVerificationCardholderAddressResult = "no_match"
)

func (r CardPaymentElementsCardValidationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationVerificationCardholderAddressResultNotChecked, CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type CardPaymentElementsCategory string

const (
	// Card Authorization: details will be under the `card_authorization` object.
	CardPaymentElementsCategoryCardAuthorization CardPaymentElementsCategory = "card_authorization"
	// Card Validation: details will be under the `card_validation` object.
	CardPaymentElementsCategoryCardValidation CardPaymentElementsCategory = "card_validation"
	// Card Decline: details will be under the `card_decline` object.
	CardPaymentElementsCategoryCardDecline CardPaymentElementsCategory = "card_decline"
	// Card Reversal: details will be under the `card_reversal` object.
	CardPaymentElementsCategoryCardReversal CardPaymentElementsCategory = "card_reversal"
	// Card Authorization Expiration: details will be under the
	// `card_authorization_expiration` object.
	CardPaymentElementsCategoryCardAuthorizationExpiration CardPaymentElementsCategory = "card_authorization_expiration"
	// Card Increment: details will be under the `card_increment` object.
	CardPaymentElementsCategoryCardIncrement CardPaymentElementsCategory = "card_increment"
	// Card Settlement: details will be under the `card_settlement` object.
	CardPaymentElementsCategoryCardSettlement CardPaymentElementsCategory = "card_settlement"
	// Card Refund: details will be under the `card_refund` object.
	CardPaymentElementsCategoryCardRefund CardPaymentElementsCategory = "card_refund"
	// Card Fuel Confirmation: details will be under the `card_fuel_confirmation`
	// object.
	CardPaymentElementsCategoryCardFuelConfirmation CardPaymentElementsCategory = "card_fuel_confirmation"
	// Unknown card payment element.
	CardPaymentElementsCategoryOther CardPaymentElementsCategory = "other"
)

func (r CardPaymentElementsCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCategoryCardAuthorization, CardPaymentElementsCategoryCardValidation, CardPaymentElementsCategoryCardDecline, CardPaymentElementsCategoryCardReversal, CardPaymentElementsCategoryCardAuthorizationExpiration, CardPaymentElementsCategoryCardIncrement, CardPaymentElementsCategoryCardSettlement, CardPaymentElementsCategoryCardRefund, CardPaymentElementsCategoryCardFuelConfirmation, CardPaymentElementsCategoryOther:
		return true
	}
	return false
}

// The summarized state of this card payment.
type CardPaymentState struct {
	// The total authorized amount in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	AuthorizedAmount int64 `json:"authorized_amount,required"`
	// The total amount from fuel confirmations in the minor unit of the transaction's
	// currency. For dollars, for example, this is cents.
	FuelConfirmedAmount int64 `json:"fuel_confirmed_amount,required"`
	// The total incrementally updated authorized amount in the minor unit of the
	// transaction's currency. For dollars, for example, this is cents.
	IncrementedAmount int64 `json:"incremented_amount,required"`
	// The total reversed amount in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	ReversedAmount int64 `json:"reversed_amount,required"`
	// The total settled or refunded amount in the minor unit of the transaction's
	// currency. For dollars, for example, this is cents.
	SettledAmount int64                `json:"settled_amount,required"`
	JSON          cardPaymentStateJSON `json:"-"`
}

// cardPaymentStateJSON contains the JSON metadata for the struct
// [CardPaymentState]
type cardPaymentStateJSON struct {
	AuthorizedAmount    apijson.Field
	FuelConfirmedAmount apijson.Field
	IncrementedAmount   apijson.Field
	ReversedAmount      apijson.Field
	SettledAmount       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardPaymentState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentStateJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_payment`.
type CardPaymentType string

const (
	CardPaymentTypeCardPayment CardPaymentType = "card_payment"
)

func (r CardPaymentType) IsKnown() bool {
	switch r {
	case CardPaymentTypeCardPayment:
		return true
	}
	return false
}

type CardPaymentListParams struct {
	// Filter Card Payments to ones belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Card Payments to ones belonging to the specified Card.
	CardID    param.Field[string]                         `query:"card_id"`
	CreatedAt param.Field[CardPaymentListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CardPaymentListParams]'s query parameters as `url.Values`.
func (r CardPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardPaymentListParamsCreatedAt struct {
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

// URLQuery serializes [CardPaymentListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CardPaymentListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
