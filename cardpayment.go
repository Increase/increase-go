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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	// A Card Authentication object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authentication`. Card Authentications
	// are attempts to authenticate a transaction or a card with 3DS.
	CardAuthentication CardPaymentElementsCardAuthentication `json:"card_authentication,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`. Card Authorizations are
	// temporary holds placed on a customers funds with the intent to later clear a
	// transaction.
	CardAuthorization CardPaymentElementsCardAuthorization `json:"card_authorization,required,nullable"`
	// A Card Authorization Expiration object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_authorization_expiration`.
	// Card Authorization Expirations are cancellations of authorizations that were
	// never settled by the acquirer.
	CardAuthorizationExpiration CardPaymentElementsCardAuthorizationExpiration `json:"card_authorization_expiration,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline CardPaymentElementsCardDecline `json:"card_decline,required,nullable"`
	// A Card Fuel Confirmation object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_fuel_confirmation`. Card Fuel
	// Confirmations update the amount of a Card Authorization after a fuel pump
	// transaction is completed.
	CardFuelConfirmation CardPaymentElementsCardFuelConfirmation `json:"card_fuel_confirmation,required,nullable"`
	// A Card Increment object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_increment`. Card Increments increase the
	// pending amount of an authorized transaction.
	CardIncrement CardPaymentElementsCardIncrement `json:"card_increment,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`. Card Refunds move money back to
	// the cardholder. While they are usually connected to a Card Settlement an
	// acquirer can also refund money directly to a card without relation to a
	// transaction.
	CardRefund CardPaymentElementsCardRefund `json:"card_refund,required,nullable"`
	// A Card Reversal object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_reversal`. Card Reversals cancel parts of
	// or the entirety of an existing Card Authorization.
	CardReversal CardPaymentElementsCardReversal `json:"card_reversal,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`. Card Settlements are card
	// transactions that have cleared and settled. While a settlement is usually
	// preceded by an authorization, an acquirer can also directly clear a transaction
	// without first authorizing it.
	CardSettlement CardPaymentElementsCardSettlement `json:"card_settlement,required,nullable"`
	// An Inbound Card Validation object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_validation`. Inbound Card
	// Validations are requests from a merchant to verify that a card number and
	// optionally its address and/or Card Verification Value are valid.
	CardValidation CardPaymentElementsCardValidation `json:"card_validation,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category CardPaymentElementsCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the card payment element was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If the category of this Transaction source is equal to `other`, this field will
	// contain an empty object, otherwise it will contain null.
	Other interface{}            `json:"other,required,nullable"`
	JSON  cardPaymentElementJSON `json:"-"`
}

// cardPaymentElementJSON contains the JSON metadata for the struct
// [CardPaymentElement]
type cardPaymentElementJSON struct {
	CardAuthentication          apijson.Field
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
	Other                       apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardPaymentElement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementJSON) RawJSON() string {
	return r.raw
}

// A Card Authentication object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authentication`. Card Authentications
// are attempts to authenticate a transaction or a card with 3DS.
type CardPaymentElementsCardAuthentication struct {
	// The Card Authentication identifier.
	ID string `json:"id,required"`
	// The identifier of the Card.
	CardID string `json:"card_id,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The category of the card authentication attempt.
	Category CardPaymentElementsCardAuthenticationCategory `json:"category,required,nullable"`
	// Details about the challenge, if one was requested.
	Challenge CardPaymentElementsCardAuthenticationChallenge `json:"challenge,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Card
	// Authentication was attempted.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The reason why this authentication attempt was denied, if it was.
	DenyReason CardPaymentElementsCardAuthenticationDenyReason `json:"deny_reason,required,nullable"`
	// The device channel of the card authentication attempt.
	DeviceChannel CardPaymentElementsCardAuthenticationDeviceChannel `json:"device_channel,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required"`
	// The purchase amount in minor units.
	PurchaseAmount int64 `json:"purchase_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// authentication attempt's purchase currency.
	PurchaseCurrency string `json:"purchase_currency,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// authentication attempt.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// The status of the card authentication.
	Status CardPaymentElementsCardAuthenticationStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authentication`.
	Type CardPaymentElementsCardAuthenticationType `json:"type,required"`
	JSON cardPaymentElementsCardAuthenticationJSON `json:"-"`
}

// cardPaymentElementsCardAuthenticationJSON contains the JSON metadata for the
// struct [CardPaymentElementsCardAuthentication]
type cardPaymentElementsCardAuthenticationJSON struct {
	ID                   apijson.Field
	CardID               apijson.Field
	CardPaymentID        apijson.Field
	Category             apijson.Field
	Challenge            apijson.Field
	CreatedAt            apijson.Field
	DenyReason           apijson.Field
	DeviceChannel        apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	PurchaseAmount       apijson.Field
	PurchaseCurrency     apijson.Field
	RealTimeDecisionID   apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthenticationJSON) RawJSON() string {
	return r.raw
}

// The category of the card authentication attempt.
type CardPaymentElementsCardAuthenticationCategory string

const (
	CardPaymentElementsCardAuthenticationCategoryPaymentAuthentication    CardPaymentElementsCardAuthenticationCategory = "payment_authentication"
	CardPaymentElementsCardAuthenticationCategoryNonPaymentAuthentication CardPaymentElementsCardAuthenticationCategory = "non_payment_authentication"
)

func (r CardPaymentElementsCardAuthenticationCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationCategoryPaymentAuthentication, CardPaymentElementsCardAuthenticationCategoryNonPaymentAuthentication:
		return true
	}
	return false
}

// Details about the challenge, if one was requested.
type CardPaymentElementsCardAuthenticationChallenge struct {
	// Details about the challenge verification attempts, if any happened.
	Attempts []CardPaymentElementsCardAuthenticationChallengeAttempt `json:"attempts,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Card
	// Authentication Challenge was started.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The one-time code used for the Card Authentication Challenge.
	OneTimeCode string `json:"one_time_code,required"`
	// The method used to verify the Card Authentication Challenge.
	VerificationMethod CardPaymentElementsCardAuthenticationChallengeVerificationMethod `json:"verification_method,required"`
	// E.g., the email address or phone number used for the Card Authentication
	// Challenge.
	VerificationValue string                                             `json:"verification_value,required,nullable"`
	JSON              cardPaymentElementsCardAuthenticationChallengeJSON `json:"-"`
}

// cardPaymentElementsCardAuthenticationChallengeJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardAuthenticationChallenge]
type cardPaymentElementsCardAuthenticationChallengeJSON struct {
	Attempts           apijson.Field
	CreatedAt          apijson.Field
	OneTimeCode        apijson.Field
	VerificationMethod apijson.Field
	VerificationValue  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthenticationChallenge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthenticationChallengeJSON) RawJSON() string {
	return r.raw
}

type CardPaymentElementsCardAuthenticationChallengeAttempt struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time of the Card
	// Authentication Challenge Attempt.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The outcome of the Card Authentication Challenge Attempt.
	Outcome CardPaymentElementsCardAuthenticationChallengeAttemptsOutcome `json:"outcome,required"`
	JSON    cardPaymentElementsCardAuthenticationChallengeAttemptJSON     `json:"-"`
}

// cardPaymentElementsCardAuthenticationChallengeAttemptJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardAuthenticationChallengeAttempt]
type cardPaymentElementsCardAuthenticationChallengeAttemptJSON struct {
	CreatedAt   apijson.Field
	Outcome     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthenticationChallengeAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthenticationChallengeAttemptJSON) RawJSON() string {
	return r.raw
}

// The outcome of the Card Authentication Challenge Attempt.
type CardPaymentElementsCardAuthenticationChallengeAttemptsOutcome string

const (
	CardPaymentElementsCardAuthenticationChallengeAttemptsOutcomeSuccessful CardPaymentElementsCardAuthenticationChallengeAttemptsOutcome = "successful"
	CardPaymentElementsCardAuthenticationChallengeAttemptsOutcomeFailed     CardPaymentElementsCardAuthenticationChallengeAttemptsOutcome = "failed"
)

func (r CardPaymentElementsCardAuthenticationChallengeAttemptsOutcome) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationChallengeAttemptsOutcomeSuccessful, CardPaymentElementsCardAuthenticationChallengeAttemptsOutcomeFailed:
		return true
	}
	return false
}

// The method used to verify the Card Authentication Challenge.
type CardPaymentElementsCardAuthenticationChallengeVerificationMethod string

const (
	CardPaymentElementsCardAuthenticationChallengeVerificationMethodTextMessage   CardPaymentElementsCardAuthenticationChallengeVerificationMethod = "text_message"
	CardPaymentElementsCardAuthenticationChallengeVerificationMethodEmail         CardPaymentElementsCardAuthenticationChallengeVerificationMethod = "email"
	CardPaymentElementsCardAuthenticationChallengeVerificationMethodNoneAvailable CardPaymentElementsCardAuthenticationChallengeVerificationMethod = "none_available"
)

func (r CardPaymentElementsCardAuthenticationChallengeVerificationMethod) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationChallengeVerificationMethodTextMessage, CardPaymentElementsCardAuthenticationChallengeVerificationMethodEmail, CardPaymentElementsCardAuthenticationChallengeVerificationMethodNoneAvailable:
		return true
	}
	return false
}

// The reason why this authentication attempt was denied, if it was.
type CardPaymentElementsCardAuthenticationDenyReason string

const (
	CardPaymentElementsCardAuthenticationDenyReasonGroupLocked           CardPaymentElementsCardAuthenticationDenyReason = "group_locked"
	CardPaymentElementsCardAuthenticationDenyReasonCardNotActive         CardPaymentElementsCardAuthenticationDenyReason = "card_not_active"
	CardPaymentElementsCardAuthenticationDenyReasonEntityNotActive       CardPaymentElementsCardAuthenticationDenyReason = "entity_not_active"
	CardPaymentElementsCardAuthenticationDenyReasonTransactionNotAllowed CardPaymentElementsCardAuthenticationDenyReason = "transaction_not_allowed"
	CardPaymentElementsCardAuthenticationDenyReasonWebhookDenied         CardPaymentElementsCardAuthenticationDenyReason = "webhook_denied"
	CardPaymentElementsCardAuthenticationDenyReasonWebhookTimedOut       CardPaymentElementsCardAuthenticationDenyReason = "webhook_timed_out"
)

func (r CardPaymentElementsCardAuthenticationDenyReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationDenyReasonGroupLocked, CardPaymentElementsCardAuthenticationDenyReasonCardNotActive, CardPaymentElementsCardAuthenticationDenyReasonEntityNotActive, CardPaymentElementsCardAuthenticationDenyReasonTransactionNotAllowed, CardPaymentElementsCardAuthenticationDenyReasonWebhookDenied, CardPaymentElementsCardAuthenticationDenyReasonWebhookTimedOut:
		return true
	}
	return false
}

// The device channel of the card authentication attempt.
type CardPaymentElementsCardAuthenticationDeviceChannel string

const (
	CardPaymentElementsCardAuthenticationDeviceChannelApp                       CardPaymentElementsCardAuthenticationDeviceChannel = "app"
	CardPaymentElementsCardAuthenticationDeviceChannelBrowser                   CardPaymentElementsCardAuthenticationDeviceChannel = "browser"
	CardPaymentElementsCardAuthenticationDeviceChannelThreeDSRequestorInitiated CardPaymentElementsCardAuthenticationDeviceChannel = "three_ds_requestor_initiated"
)

func (r CardPaymentElementsCardAuthenticationDeviceChannel) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationDeviceChannelApp, CardPaymentElementsCardAuthenticationDeviceChannelBrowser, CardPaymentElementsCardAuthenticationDeviceChannelThreeDSRequestorInitiated:
		return true
	}
	return false
}

// The status of the card authentication.
type CardPaymentElementsCardAuthenticationStatus string

const (
	CardPaymentElementsCardAuthenticationStatusDenied                        CardPaymentElementsCardAuthenticationStatus = "denied"
	CardPaymentElementsCardAuthenticationStatusAuthenticatedWithChallenge    CardPaymentElementsCardAuthenticationStatus = "authenticated_with_challenge"
	CardPaymentElementsCardAuthenticationStatusAuthenticatedWithoutChallenge CardPaymentElementsCardAuthenticationStatus = "authenticated_without_challenge"
	CardPaymentElementsCardAuthenticationStatusAwaitingChallenge             CardPaymentElementsCardAuthenticationStatus = "awaiting_challenge"
	CardPaymentElementsCardAuthenticationStatusValidatingChallenge           CardPaymentElementsCardAuthenticationStatus = "validating_challenge"
	CardPaymentElementsCardAuthenticationStatusCanceled                      CardPaymentElementsCardAuthenticationStatus = "canceled"
	CardPaymentElementsCardAuthenticationStatusTimedOutAwaitingChallenge     CardPaymentElementsCardAuthenticationStatus = "timed_out_awaiting_challenge"
	CardPaymentElementsCardAuthenticationStatusErrored                       CardPaymentElementsCardAuthenticationStatus = "errored"
	CardPaymentElementsCardAuthenticationStatusExceededAttemptThreshold      CardPaymentElementsCardAuthenticationStatus = "exceeded_attempt_threshold"
)

func (r CardPaymentElementsCardAuthenticationStatus) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationStatusDenied, CardPaymentElementsCardAuthenticationStatusAuthenticatedWithChallenge, CardPaymentElementsCardAuthenticationStatusAuthenticatedWithoutChallenge, CardPaymentElementsCardAuthenticationStatusAwaitingChallenge, CardPaymentElementsCardAuthenticationStatusValidatingChallenge, CardPaymentElementsCardAuthenticationStatusCanceled, CardPaymentElementsCardAuthenticationStatusTimedOutAwaitingChallenge, CardPaymentElementsCardAuthenticationStatusErrored, CardPaymentElementsCardAuthenticationStatusExceededAttemptThreshold:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_authentication`.
type CardPaymentElementsCardAuthenticationType string

const (
	CardPaymentElementsCardAuthenticationTypeCardAuthentication CardPaymentElementsCardAuthenticationType = "card_authentication"
)

func (r CardPaymentElementsCardAuthenticationType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthenticationTypeCardAuthentication:
		return true
	}
	return false
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`. Card Authorizations are
// temporary holds placed on a customers funds with the intent to later clear a
// transaction.
type CardPaymentElementsCardAuthorization struct {
	// The Card Authorization identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardAuthorizationActioner `json:"actioner,required"`
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts CardPaymentElementsCardAuthorizationAdditionalAmounts `json:"additional_amounts,required"`
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
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
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
	AdditionalAmounts    apijson.Field
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
	TerminalID           apijson.Field
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
	CardPaymentElementsCardAuthorizationActionerUser     CardPaymentElementsCardAuthorizationActioner = "user"
	CardPaymentElementsCardAuthorizationActionerIncrease CardPaymentElementsCardAuthorizationActioner = "increase"
	CardPaymentElementsCardAuthorizationActionerNetwork  CardPaymentElementsCardAuthorizationActioner = "network"
)

func (r CardPaymentElementsCardAuthorizationActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationActionerUser, CardPaymentElementsCardAuthorizationActionerIncrease, CardPaymentElementsCardAuthorizationActionerNetwork:
		return true
	}
	return false
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type CardPaymentElementsCardAuthorizationAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic CardPaymentElementsCardAuthorizationAdditionalAmountsClinic `json:"clinic,required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental CardPaymentElementsCardAuthorizationAdditionalAmountsDental `json:"dental,required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription CardPaymentElementsCardAuthorizationAdditionalAmountsPrescription `json:"prescription,required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge CardPaymentElementsCardAuthorizationAdditionalAmountsSurcharge `json:"surcharge,required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative CardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulative `json:"total_cumulative,required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare CardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcare `json:"total_healthcare,required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit CardPaymentElementsCardAuthorizationAdditionalAmountsTransit `json:"transit,required,nullable"`
	// An unknown additional amount.
	Unknown CardPaymentElementsCardAuthorizationAdditionalAmountsUnknown `json:"unknown,required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision CardPaymentElementsCardAuthorizationAdditionalAmountsVision `json:"vision,required,nullable"`
	JSON   cardPaymentElementsCardAuthorizationAdditionalAmountsJSON   `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardAuthorizationAdditionalAmounts]
type cardPaymentElementsCardAuthorizationAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
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

func (r *CardPaymentElementsCardAuthorizationAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type CardPaymentElementsCardAuthorizationAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsClinicJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsClinicJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsClinic]
type cardPaymentElementsCardAuthorizationAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type CardPaymentElementsCardAuthorizationAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsDentalJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsDentalJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsDental]
type cardPaymentElementsCardAuthorizationAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type CardPaymentElementsCardAuthorizationAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsPrescriptionJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsPrescriptionJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsPrescription]
type cardPaymentElementsCardAuthorizationAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type CardPaymentElementsCardAuthorizationAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsSurchargeJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsSurchargeJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsSurcharge]
type cardPaymentElementsCardAuthorizationAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type CardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                   `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulativeJSON
// contains the JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulative]
type cardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type CardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                   `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcareJSON
// contains the JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcare]
type cardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type CardPaymentElementsCardAuthorizationAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                           `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsTransitJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsTransitJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsTransit]
type cardPaymentElementsCardAuthorizationAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type CardPaymentElementsCardAuthorizationAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                           `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsUnknownJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsUnknownJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsUnknown]
type cardPaymentElementsCardAuthorizationAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type CardPaymentElementsCardAuthorizationAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     cardPaymentElementsCardAuthorizationAdditionalAmountsVisionJSON `json:"-"`
}

// cardPaymentElementsCardAuthorizationAdditionalAmountsVisionJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardAuthorizationAdditionalAmountsVision]
type cardPaymentElementsCardAuthorizationAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardAuthorizationAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardAuthorizationAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CardPaymentElementsCardAuthorizationCurrency string

const (
	CardPaymentElementsCardAuthorizationCurrencyCad CardPaymentElementsCardAuthorizationCurrency = "CAD"
	CardPaymentElementsCardAuthorizationCurrencyChf CardPaymentElementsCardAuthorizationCurrency = "CHF"
	CardPaymentElementsCardAuthorizationCurrencyEur CardPaymentElementsCardAuthorizationCurrency = "EUR"
	CardPaymentElementsCardAuthorizationCurrencyGbp CardPaymentElementsCardAuthorizationCurrency = "GBP"
	CardPaymentElementsCardAuthorizationCurrencyJpy CardPaymentElementsCardAuthorizationCurrency = "JPY"
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
	CardPaymentElementsCardAuthorizationDirectionSettlement CardPaymentElementsCardAuthorizationDirection = "settlement"
	CardPaymentElementsCardAuthorizationDirectionRefund     CardPaymentElementsCardAuthorizationDirection = "refund"
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
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason,required,nullable"`
	JSON                    cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
}

// cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardAuthorizationNetworkDetailsVisa]
type cardPaymentElementsCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
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
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    CardPaymentElementsCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown                    CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual                     CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless                CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardPaymentElementsCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason string

const (
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonIssuerError                                              CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonOther                                                    CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonIssuerError, CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInternalVisaError, CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, CardPaymentElementsCardAuthorizationNetworkDetailsVisaStandInProcessingReasonOther:
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
	CardPaymentElementsCardAuthorizationProcessingCategoryAccountFunding         CardPaymentElementsCardAuthorizationProcessingCategory = "account_funding"
	CardPaymentElementsCardAuthorizationProcessingCategoryAutomaticFuelDispenser CardPaymentElementsCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	CardPaymentElementsCardAuthorizationProcessingCategoryBillPayment            CardPaymentElementsCardAuthorizationProcessingCategory = "bill_payment"
	CardPaymentElementsCardAuthorizationProcessingCategoryOriginalCredit         CardPaymentElementsCardAuthorizationProcessingCategory = "original_credit"
	CardPaymentElementsCardAuthorizationProcessingCategoryPurchase               CardPaymentElementsCardAuthorizationProcessingCategory = "purchase"
	CardPaymentElementsCardAuthorizationProcessingCategoryQuasiCash              CardPaymentElementsCardAuthorizationProcessingCategory = "quasi_cash"
	CardPaymentElementsCardAuthorizationProcessingCategoryRefund                 CardPaymentElementsCardAuthorizationProcessingCategory = "refund"
	CardPaymentElementsCardAuthorizationProcessingCategoryCashDisbursement       CardPaymentElementsCardAuthorizationProcessingCategory = "cash_disbursement"
	CardPaymentElementsCardAuthorizationProcessingCategoryUnknown                CardPaymentElementsCardAuthorizationProcessingCategory = "unknown"
)

func (r CardPaymentElementsCardAuthorizationProcessingCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationProcessingCategoryAccountFunding, CardPaymentElementsCardAuthorizationProcessingCategoryAutomaticFuelDispenser, CardPaymentElementsCardAuthorizationProcessingCategoryBillPayment, CardPaymentElementsCardAuthorizationProcessingCategoryOriginalCredit, CardPaymentElementsCardAuthorizationProcessingCategoryPurchase, CardPaymentElementsCardAuthorizationProcessingCategoryQuasiCash, CardPaymentElementsCardAuthorizationProcessingCategoryRefund, CardPaymentElementsCardAuthorizationProcessingCategoryCashDisbursement, CardPaymentElementsCardAuthorizationProcessingCategoryUnknown:
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
	CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultNotChecked CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultMatch      CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult = "match"
	CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResultNoMatch    CardPaymentElementsCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
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
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNotChecked                       CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultMatch                            CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "match"
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNoMatch                          CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "no_match"
	CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r CardPaymentElementsCardAuthorizationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNotChecked, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultNoMatch, CardPaymentElementsCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// A Card Authorization Expiration object. This field will be present in the JSON
// response if and only if `category` is equal to `card_authorization_expiration`.
// Card Authorization Expirations are cancellations of authorizations that were
// never settled by the acquirer.
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
	CardPaymentElementsCardAuthorizationExpirationCurrencyCad CardPaymentElementsCardAuthorizationExpirationCurrency = "CAD"
	CardPaymentElementsCardAuthorizationExpirationCurrencyChf CardPaymentElementsCardAuthorizationExpirationCurrency = "CHF"
	CardPaymentElementsCardAuthorizationExpirationCurrencyEur CardPaymentElementsCardAuthorizationExpirationCurrency = "EUR"
	CardPaymentElementsCardAuthorizationExpirationCurrencyGbp CardPaymentElementsCardAuthorizationExpirationCurrency = "GBP"
	CardPaymentElementsCardAuthorizationExpirationCurrencyJpy CardPaymentElementsCardAuthorizationExpirationCurrency = "JPY"
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
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts CardPaymentElementsCardDeclineAdditionalAmounts `json:"additional_amounts,required"`
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
	// The direction describes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction CardPaymentElementsCardDeclineDirection `json:"direction,required"`
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
	// This is present if a specific decline reason was given in the real-time
	// decision.
	RealTimeDecisionReason CardPaymentElementsCardDeclineRealTimeDecisionReason `json:"real_time_decision_reason,required,nullable"`
	// Why the transaction was declined.
	Reason CardPaymentElementsCardDeclineReason `json:"reason,required"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
	// Fields related to verification of cardholder-provided values.
	Verification CardPaymentElementsCardDeclineVerification `json:"verification,required"`
	JSON         cardPaymentElementsCardDeclineJSON         `json:"-"`
}

// cardPaymentElementsCardDeclineJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardDecline]
type cardPaymentElementsCardDeclineJSON struct {
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
	CardPaymentElementsCardDeclineActionerUser     CardPaymentElementsCardDeclineActioner = "user"
	CardPaymentElementsCardDeclineActionerIncrease CardPaymentElementsCardDeclineActioner = "increase"
	CardPaymentElementsCardDeclineActionerNetwork  CardPaymentElementsCardDeclineActioner = "network"
)

func (r CardPaymentElementsCardDeclineActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineActionerUser, CardPaymentElementsCardDeclineActionerIncrease, CardPaymentElementsCardDeclineActionerNetwork:
		return true
	}
	return false
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type CardPaymentElementsCardDeclineAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic CardPaymentElementsCardDeclineAdditionalAmountsClinic `json:"clinic,required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental CardPaymentElementsCardDeclineAdditionalAmountsDental `json:"dental,required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription CardPaymentElementsCardDeclineAdditionalAmountsPrescription `json:"prescription,required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge CardPaymentElementsCardDeclineAdditionalAmountsSurcharge `json:"surcharge,required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative CardPaymentElementsCardDeclineAdditionalAmountsTotalCumulative `json:"total_cumulative,required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare CardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcare `json:"total_healthcare,required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit CardPaymentElementsCardDeclineAdditionalAmountsTransit `json:"transit,required,nullable"`
	// An unknown additional amount.
	Unknown CardPaymentElementsCardDeclineAdditionalAmountsUnknown `json:"unknown,required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision CardPaymentElementsCardDeclineAdditionalAmountsVision `json:"vision,required,nullable"`
	JSON   cardPaymentElementsCardDeclineAdditionalAmountsJSON   `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardDeclineAdditionalAmounts]
type cardPaymentElementsCardDeclineAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
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

func (r *CardPaymentElementsCardDeclineAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type CardPaymentElementsCardDeclineAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                    `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsClinicJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsClinicJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardDeclineAdditionalAmountsClinic]
type cardPaymentElementsCardDeclineAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type CardPaymentElementsCardDeclineAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                    `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsDentalJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsDentalJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardDeclineAdditionalAmountsDental]
type cardPaymentElementsCardDeclineAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type CardPaymentElementsCardDeclineAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsPrescriptionJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsPrescriptionJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardDeclineAdditionalAmountsPrescription]
type cardPaymentElementsCardDeclineAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type CardPaymentElementsCardDeclineAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsSurchargeJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsSurchargeJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardDeclineAdditionalAmountsSurcharge]
type cardPaymentElementsCardDeclineAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type CardPaymentElementsCardDeclineAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsTotalCumulativeJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardDeclineAdditionalAmountsTotalCumulative]
type cardPaymentElementsCardDeclineAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type CardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcareJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcare]
type cardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type CardPaymentElementsCardDeclineAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                     `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsTransitJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsTransitJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardDeclineAdditionalAmountsTransit]
type cardPaymentElementsCardDeclineAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type CardPaymentElementsCardDeclineAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                     `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsUnknownJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsUnknownJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardDeclineAdditionalAmountsUnknown]
type cardPaymentElementsCardDeclineAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type CardPaymentElementsCardDeclineAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                    `json:"currency,required"`
	JSON     cardPaymentElementsCardDeclineAdditionalAmountsVisionJSON `json:"-"`
}

// cardPaymentElementsCardDeclineAdditionalAmountsVisionJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardDeclineAdditionalAmountsVision]
type cardPaymentElementsCardDeclineAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardDeclineAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardDeclineAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type CardPaymentElementsCardDeclineCurrency string

const (
	CardPaymentElementsCardDeclineCurrencyCad CardPaymentElementsCardDeclineCurrency = "CAD"
	CardPaymentElementsCardDeclineCurrencyChf CardPaymentElementsCardDeclineCurrency = "CHF"
	CardPaymentElementsCardDeclineCurrencyEur CardPaymentElementsCardDeclineCurrency = "EUR"
	CardPaymentElementsCardDeclineCurrencyGbp CardPaymentElementsCardDeclineCurrency = "GBP"
	CardPaymentElementsCardDeclineCurrencyJpy CardPaymentElementsCardDeclineCurrency = "JPY"
	CardPaymentElementsCardDeclineCurrencyUsd CardPaymentElementsCardDeclineCurrency = "USD"
)

func (r CardPaymentElementsCardDeclineCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineCurrencyCad, CardPaymentElementsCardDeclineCurrencyChf, CardPaymentElementsCardDeclineCurrencyEur, CardPaymentElementsCardDeclineCurrencyGbp, CardPaymentElementsCardDeclineCurrencyJpy, CardPaymentElementsCardDeclineCurrencyUsd:
		return true
	}
	return false
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type CardPaymentElementsCardDeclineDirection string

const (
	CardPaymentElementsCardDeclineDirectionSettlement CardPaymentElementsCardDeclineDirection = "settlement"
	CardPaymentElementsCardDeclineDirectionRefund     CardPaymentElementsCardDeclineDirection = "refund"
)

func (r CardPaymentElementsCardDeclineDirection) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineDirectionSettlement, CardPaymentElementsCardDeclineDirectionRefund:
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
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason,required,nullable"`
	JSON                    cardPaymentElementsCardDeclineNetworkDetailsVisaJSON                    `json:"-"`
}

// cardPaymentElementsCardDeclineNetworkDetailsVisaJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardDeclineNetworkDetailsVisa]
type cardPaymentElementsCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
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
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    CardPaymentElementsCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown                    CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual                     CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless                CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardPaymentElementsCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason string

const (
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonIssuerError                                              CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonOther                                                    CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonIssuerError, CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonInternalVisaError, CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, CardPaymentElementsCardDeclineNetworkDetailsVisaStandInProcessingReasonOther:
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
	CardPaymentElementsCardDeclineProcessingCategoryAccountFunding         CardPaymentElementsCardDeclineProcessingCategory = "account_funding"
	CardPaymentElementsCardDeclineProcessingCategoryAutomaticFuelDispenser CardPaymentElementsCardDeclineProcessingCategory = "automatic_fuel_dispenser"
	CardPaymentElementsCardDeclineProcessingCategoryBillPayment            CardPaymentElementsCardDeclineProcessingCategory = "bill_payment"
	CardPaymentElementsCardDeclineProcessingCategoryOriginalCredit         CardPaymentElementsCardDeclineProcessingCategory = "original_credit"
	CardPaymentElementsCardDeclineProcessingCategoryPurchase               CardPaymentElementsCardDeclineProcessingCategory = "purchase"
	CardPaymentElementsCardDeclineProcessingCategoryQuasiCash              CardPaymentElementsCardDeclineProcessingCategory = "quasi_cash"
	CardPaymentElementsCardDeclineProcessingCategoryRefund                 CardPaymentElementsCardDeclineProcessingCategory = "refund"
	CardPaymentElementsCardDeclineProcessingCategoryCashDisbursement       CardPaymentElementsCardDeclineProcessingCategory = "cash_disbursement"
	CardPaymentElementsCardDeclineProcessingCategoryUnknown                CardPaymentElementsCardDeclineProcessingCategory = "unknown"
)

func (r CardPaymentElementsCardDeclineProcessingCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineProcessingCategoryAccountFunding, CardPaymentElementsCardDeclineProcessingCategoryAutomaticFuelDispenser, CardPaymentElementsCardDeclineProcessingCategoryBillPayment, CardPaymentElementsCardDeclineProcessingCategoryOriginalCredit, CardPaymentElementsCardDeclineProcessingCategoryPurchase, CardPaymentElementsCardDeclineProcessingCategoryQuasiCash, CardPaymentElementsCardDeclineProcessingCategoryRefund, CardPaymentElementsCardDeclineProcessingCategoryCashDisbursement, CardPaymentElementsCardDeclineProcessingCategoryUnknown:
		return true
	}
	return false
}

// This is present if a specific decline reason was given in the real-time
// decision.
type CardPaymentElementsCardDeclineRealTimeDecisionReason string

const (
	CardPaymentElementsCardDeclineRealTimeDecisionReasonInsufficientFunds       CardPaymentElementsCardDeclineRealTimeDecisionReason = "insufficient_funds"
	CardPaymentElementsCardDeclineRealTimeDecisionReasonTransactionNeverAllowed CardPaymentElementsCardDeclineRealTimeDecisionReason = "transaction_never_allowed"
	CardPaymentElementsCardDeclineRealTimeDecisionReasonExceedsApprovalLimit    CardPaymentElementsCardDeclineRealTimeDecisionReason = "exceeds_approval_limit"
	CardPaymentElementsCardDeclineRealTimeDecisionReasonCardTemporarilyDisabled CardPaymentElementsCardDeclineRealTimeDecisionReason = "card_temporarily_disabled"
	CardPaymentElementsCardDeclineRealTimeDecisionReasonSuspectedFraud          CardPaymentElementsCardDeclineRealTimeDecisionReason = "suspected_fraud"
	CardPaymentElementsCardDeclineRealTimeDecisionReasonOther                   CardPaymentElementsCardDeclineRealTimeDecisionReason = "other"
)

func (r CardPaymentElementsCardDeclineRealTimeDecisionReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineRealTimeDecisionReasonInsufficientFunds, CardPaymentElementsCardDeclineRealTimeDecisionReasonTransactionNeverAllowed, CardPaymentElementsCardDeclineRealTimeDecisionReasonExceedsApprovalLimit, CardPaymentElementsCardDeclineRealTimeDecisionReasonCardTemporarilyDisabled, CardPaymentElementsCardDeclineRealTimeDecisionReasonSuspectedFraud, CardPaymentElementsCardDeclineRealTimeDecisionReasonOther:
		return true
	}
	return false
}

// Why the transaction was declined.
type CardPaymentElementsCardDeclineReason string

const (
	CardPaymentElementsCardDeclineReasonAccountClosed                CardPaymentElementsCardDeclineReason = "account_closed"
	CardPaymentElementsCardDeclineReasonCardNotActive                CardPaymentElementsCardDeclineReason = "card_not_active"
	CardPaymentElementsCardDeclineReasonCardCanceled                 CardPaymentElementsCardDeclineReason = "card_canceled"
	CardPaymentElementsCardDeclineReasonPhysicalCardNotActive        CardPaymentElementsCardDeclineReason = "physical_card_not_active"
	CardPaymentElementsCardDeclineReasonEntityNotActive              CardPaymentElementsCardDeclineReason = "entity_not_active"
	CardPaymentElementsCardDeclineReasonGroupLocked                  CardPaymentElementsCardDeclineReason = "group_locked"
	CardPaymentElementsCardDeclineReasonInsufficientFunds            CardPaymentElementsCardDeclineReason = "insufficient_funds"
	CardPaymentElementsCardDeclineReasonCvv2Mismatch                 CardPaymentElementsCardDeclineReason = "cvv2_mismatch"
	CardPaymentElementsCardDeclineReasonPinMismatch                  CardPaymentElementsCardDeclineReason = "pin_mismatch"
	CardPaymentElementsCardDeclineReasonCardExpirationMismatch       CardPaymentElementsCardDeclineReason = "card_expiration_mismatch"
	CardPaymentElementsCardDeclineReasonTransactionNotAllowed        CardPaymentElementsCardDeclineReason = "transaction_not_allowed"
	CardPaymentElementsCardDeclineReasonBreachesLimit                CardPaymentElementsCardDeclineReason = "breaches_limit"
	CardPaymentElementsCardDeclineReasonWebhookDeclined              CardPaymentElementsCardDeclineReason = "webhook_declined"
	CardPaymentElementsCardDeclineReasonWebhookTimedOut              CardPaymentElementsCardDeclineReason = "webhook_timed_out"
	CardPaymentElementsCardDeclineReasonDeclinedByStandInProcessing  CardPaymentElementsCardDeclineReason = "declined_by_stand_in_processing"
	CardPaymentElementsCardDeclineReasonInvalidPhysicalCard          CardPaymentElementsCardDeclineReason = "invalid_physical_card"
	CardPaymentElementsCardDeclineReasonMissingOriginalAuthorization CardPaymentElementsCardDeclineReason = "missing_original_authorization"
	CardPaymentElementsCardDeclineReasonFailed3DSAuthentication      CardPaymentElementsCardDeclineReason = "failed_3ds_authentication"
	CardPaymentElementsCardDeclineReasonSuspectedCardTesting         CardPaymentElementsCardDeclineReason = "suspected_card_testing"
	CardPaymentElementsCardDeclineReasonSuspectedFraud               CardPaymentElementsCardDeclineReason = "suspected_fraud"
)

func (r CardPaymentElementsCardDeclineReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineReasonAccountClosed, CardPaymentElementsCardDeclineReasonCardNotActive, CardPaymentElementsCardDeclineReasonCardCanceled, CardPaymentElementsCardDeclineReasonPhysicalCardNotActive, CardPaymentElementsCardDeclineReasonEntityNotActive, CardPaymentElementsCardDeclineReasonGroupLocked, CardPaymentElementsCardDeclineReasonInsufficientFunds, CardPaymentElementsCardDeclineReasonCvv2Mismatch, CardPaymentElementsCardDeclineReasonPinMismatch, CardPaymentElementsCardDeclineReasonCardExpirationMismatch, CardPaymentElementsCardDeclineReasonTransactionNotAllowed, CardPaymentElementsCardDeclineReasonBreachesLimit, CardPaymentElementsCardDeclineReasonWebhookDeclined, CardPaymentElementsCardDeclineReasonWebhookTimedOut, CardPaymentElementsCardDeclineReasonDeclinedByStandInProcessing, CardPaymentElementsCardDeclineReasonInvalidPhysicalCard, CardPaymentElementsCardDeclineReasonMissingOriginalAuthorization, CardPaymentElementsCardDeclineReasonFailed3DSAuthentication, CardPaymentElementsCardDeclineReasonSuspectedCardTesting, CardPaymentElementsCardDeclineReasonSuspectedFraud:
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
	CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultNotChecked CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult = "not_checked"
	CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultMatch      CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult = "match"
	CardPaymentElementsCardDeclineVerificationCardVerificationCodeResultNoMatch    CardPaymentElementsCardDeclineVerificationCardVerificationCodeResult = "no_match"
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
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultNotChecked                       CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "not_checked"
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultMatch                            CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "match"
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultNoMatch                          CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "no_match"
	CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardPaymentElementsCardDeclineVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r CardPaymentElementsCardDeclineVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardDeclineVerificationCardholderAddressResultNotChecked, CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultNoMatch, CardPaymentElementsCardDeclineVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// A Card Fuel Confirmation object. This field will be present in the JSON response
// if and only if `category` is equal to `card_fuel_confirmation`. Card Fuel
// Confirmations update the amount of a Card Authorization after a fuel pump
// transaction is completed.
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
	CardPaymentElementsCardFuelConfirmationCurrencyCad CardPaymentElementsCardFuelConfirmationCurrency = "CAD"
	CardPaymentElementsCardFuelConfirmationCurrencyChf CardPaymentElementsCardFuelConfirmationCurrency = "CHF"
	CardPaymentElementsCardFuelConfirmationCurrencyEur CardPaymentElementsCardFuelConfirmationCurrency = "EUR"
	CardPaymentElementsCardFuelConfirmationCurrencyGbp CardPaymentElementsCardFuelConfirmationCurrency = "GBP"
	CardPaymentElementsCardFuelConfirmationCurrencyJpy CardPaymentElementsCardFuelConfirmationCurrency = "JPY"
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
// only if `category` is equal to `card_increment`. Card Increments increase the
// pending amount of an authorized transaction.
type CardPaymentElementsCardIncrement struct {
	// The Card Increment identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardIncrementActioner `json:"actioner,required"`
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts CardPaymentElementsCardIncrementAdditionalAmounts `json:"additional_amounts,required"`
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
	// The amount of this increment in the minor unit of the transaction's presentment
	// currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
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
	AdditionalAmounts          apijson.Field
	Amount                     apijson.Field
	CardAuthorizationID        apijson.Field
	Currency                   apijson.Field
	Network                    apijson.Field
	NetworkIdentifiers         apijson.Field
	NetworkRiskScore           apijson.Field
	PendingTransactionID       apijson.Field
	PresentmentAmount          apijson.Field
	PresentmentCurrency        apijson.Field
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
	CardPaymentElementsCardIncrementActionerUser     CardPaymentElementsCardIncrementActioner = "user"
	CardPaymentElementsCardIncrementActionerIncrease CardPaymentElementsCardIncrementActioner = "increase"
	CardPaymentElementsCardIncrementActionerNetwork  CardPaymentElementsCardIncrementActioner = "network"
)

func (r CardPaymentElementsCardIncrementActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardIncrementActionerUser, CardPaymentElementsCardIncrementActionerIncrease, CardPaymentElementsCardIncrementActionerNetwork:
		return true
	}
	return false
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type CardPaymentElementsCardIncrementAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic CardPaymentElementsCardIncrementAdditionalAmountsClinic `json:"clinic,required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental CardPaymentElementsCardIncrementAdditionalAmountsDental `json:"dental,required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription CardPaymentElementsCardIncrementAdditionalAmountsPrescription `json:"prescription,required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge CardPaymentElementsCardIncrementAdditionalAmountsSurcharge `json:"surcharge,required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative CardPaymentElementsCardIncrementAdditionalAmountsTotalCumulative `json:"total_cumulative,required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare CardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcare `json:"total_healthcare,required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit CardPaymentElementsCardIncrementAdditionalAmountsTransit `json:"transit,required,nullable"`
	// An unknown additional amount.
	Unknown CardPaymentElementsCardIncrementAdditionalAmountsUnknown `json:"unknown,required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision CardPaymentElementsCardIncrementAdditionalAmountsVision `json:"vision,required,nullable"`
	JSON   cardPaymentElementsCardIncrementAdditionalAmountsJSON   `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsJSON contains the JSON metadata
// for the struct [CardPaymentElementsCardIncrementAdditionalAmounts]
type cardPaymentElementsCardIncrementAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
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

func (r *CardPaymentElementsCardIncrementAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type CardPaymentElementsCardIncrementAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                      `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsClinicJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsClinicJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsClinic]
type cardPaymentElementsCardIncrementAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type CardPaymentElementsCardIncrementAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                      `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsDentalJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsDentalJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsDental]
type cardPaymentElementsCardIncrementAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type CardPaymentElementsCardIncrementAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                            `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsPrescriptionJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsPrescriptionJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsPrescription]
type cardPaymentElementsCardIncrementAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type CardPaymentElementsCardIncrementAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                         `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsSurchargeJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsSurchargeJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsSurcharge]
type cardPaymentElementsCardIncrementAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type CardPaymentElementsCardIncrementAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                               `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsTotalCumulativeJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsTotalCumulative]
type cardPaymentElementsCardIncrementAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type CardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                               `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcareJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcare]
type cardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type CardPaymentElementsCardIncrementAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsTransitJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsTransitJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsTransit]
type cardPaymentElementsCardIncrementAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type CardPaymentElementsCardIncrementAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsUnknownJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsUnknownJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsUnknown]
type cardPaymentElementsCardIncrementAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type CardPaymentElementsCardIncrementAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                      `json:"currency,required"`
	JSON     cardPaymentElementsCardIncrementAdditionalAmountsVisionJSON `json:"-"`
}

// cardPaymentElementsCardIncrementAdditionalAmountsVisionJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardIncrementAdditionalAmountsVision]
type cardPaymentElementsCardIncrementAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardIncrementAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardIncrementAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the increment's
// currency.
type CardPaymentElementsCardIncrementCurrency string

const (
	CardPaymentElementsCardIncrementCurrencyCad CardPaymentElementsCardIncrementCurrency = "CAD"
	CardPaymentElementsCardIncrementCurrencyChf CardPaymentElementsCardIncrementCurrency = "CHF"
	CardPaymentElementsCardIncrementCurrencyEur CardPaymentElementsCardIncrementCurrency = "EUR"
	CardPaymentElementsCardIncrementCurrencyGbp CardPaymentElementsCardIncrementCurrency = "GBP"
	CardPaymentElementsCardIncrementCurrencyJpy CardPaymentElementsCardIncrementCurrency = "JPY"
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
// only if `category` is equal to `card_refund`. Card Refunds move money back to
// the cardholder. While they are usually connected to a Card Settlement an
// acquirer can also refund money directly to a card without relation to a
// transaction.
type CardPaymentElementsCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// Cashback debited for this transaction, if eligible. Cashback is paid out in
	// aggregate, monthly.
	Cashback CardPaymentElementsCardRefundCashback `json:"cashback,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency CardPaymentElementsCardRefundCurrency `json:"currency,required"`
	// Interchange assessed as a part of this transaciton.
	Interchange CardPaymentElementsCardRefundInterchange `json:"interchange,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required"`
	// The merchant's postal code. For US merchants this is always a 5-digit ZIP code.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
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
	Cashback             apijson.Field
	Currency             apijson.Field
	Interchange          apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantPostalCode   apijson.Field
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

// Cashback debited for this transaction, if eligible. Cashback is paid out in
// aggregate, monthly.
type CardPaymentElementsCardRefundCashback struct {
	// The cashback amount given as a string containing a decimal number. The amount is
	// a positive number if it will be credited to you (e.g., settlements) and a
	// negative number if it will be debited (e.g., refunds).
	Amount string `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
	Currency CardPaymentElementsCardRefundCashbackCurrency `json:"currency,required"`
	JSON     cardPaymentElementsCardRefundCashbackJSON     `json:"-"`
}

// cardPaymentElementsCardRefundCashbackJSON contains the JSON metadata for the
// struct [CardPaymentElementsCardRefundCashback]
type cardPaymentElementsCardRefundCashbackJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardRefundCashback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardRefundCashbackJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
type CardPaymentElementsCardRefundCashbackCurrency string

const (
	CardPaymentElementsCardRefundCashbackCurrencyCad CardPaymentElementsCardRefundCashbackCurrency = "CAD"
	CardPaymentElementsCardRefundCashbackCurrencyChf CardPaymentElementsCardRefundCashbackCurrency = "CHF"
	CardPaymentElementsCardRefundCashbackCurrencyEur CardPaymentElementsCardRefundCashbackCurrency = "EUR"
	CardPaymentElementsCardRefundCashbackCurrencyGbp CardPaymentElementsCardRefundCashbackCurrency = "GBP"
	CardPaymentElementsCardRefundCashbackCurrencyJpy CardPaymentElementsCardRefundCashbackCurrency = "JPY"
	CardPaymentElementsCardRefundCashbackCurrencyUsd CardPaymentElementsCardRefundCashbackCurrency = "USD"
)

func (r CardPaymentElementsCardRefundCashbackCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardRefundCashbackCurrencyCad, CardPaymentElementsCardRefundCashbackCurrencyChf, CardPaymentElementsCardRefundCashbackCurrencyEur, CardPaymentElementsCardRefundCashbackCurrencyGbp, CardPaymentElementsCardRefundCashbackCurrencyJpy, CardPaymentElementsCardRefundCashbackCurrencyUsd:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type CardPaymentElementsCardRefundCurrency string

const (
	CardPaymentElementsCardRefundCurrencyCad CardPaymentElementsCardRefundCurrency = "CAD"
	CardPaymentElementsCardRefundCurrencyChf CardPaymentElementsCardRefundCurrency = "CHF"
	CardPaymentElementsCardRefundCurrencyEur CardPaymentElementsCardRefundCurrency = "EUR"
	CardPaymentElementsCardRefundCurrencyGbp CardPaymentElementsCardRefundCurrency = "GBP"
	CardPaymentElementsCardRefundCurrencyJpy CardPaymentElementsCardRefundCurrency = "JPY"
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
	// The interchange amount given as a string containing a decimal number in major
	// units (so e.g., "3.14" for $3.14). The amount is a positive number if it is
	// credited to Increase (e.g., settlements) and a negative number if it is debited
	// (e.g., refunds).
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
	CardPaymentElementsCardRefundInterchangeCurrencyCad CardPaymentElementsCardRefundInterchangeCurrency = "CAD"
	CardPaymentElementsCardRefundInterchangeCurrencyChf CardPaymentElementsCardRefundInterchangeCurrency = "CHF"
	CardPaymentElementsCardRefundInterchangeCurrencyEur CardPaymentElementsCardRefundInterchangeCurrency = "EUR"
	CardPaymentElementsCardRefundInterchangeCurrencyGbp CardPaymentElementsCardRefundInterchangeCurrency = "GBP"
	CardPaymentElementsCardRefundInterchangeCurrencyJpy CardPaymentElementsCardRefundInterchangeCurrency = "JPY"
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
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge    CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesGas              CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage     CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn       CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee CardPaymentElementsCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
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
	CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable               CardPaymentElementsCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
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
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesRestaurant    CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesGiftShop      CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesMiniBar       CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesTelephone     CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesOther         CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraChargesLaundry       CardPaymentElementsCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
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
	CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow        CardPaymentElementsCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
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
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText              CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber           CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber      CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber         CardPaymentElementsCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
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
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit                                                        CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther                                                           CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
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
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone                  CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService        CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee            CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee             CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo                 CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset          CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer         CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard              CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport       CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge                CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical               CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage          CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther                 CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee    CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets                  CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees              CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby               CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee            CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore                 CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService         CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel   CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades              CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi                  CardPaymentElementsCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
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
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit                                                        CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation                                       CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther                                                           CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket                                    CardPaymentElementsCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
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
	CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions                CardPaymentElementsCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
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
	CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone                   CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket              CardPaymentElementsCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
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
	CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone               CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed    CardPaymentElementsCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
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
// only if `category` is equal to `card_reversal`. Card Reversals cancel parts of
// or the entirety of an existing Card Authorization.
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
	MerchantCategoryCode string `json:"merchant_category_code,required"`
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
	// The card network used to process this card authorization.
	Network CardPaymentElementsCardReversalNetwork `json:"network,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers CardPaymentElementsCardReversalNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Card Reversal.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the reversal's
	// presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The amount of this reversal in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	ReversalAmount int64 `json:"reversal_amount,required"`
	// The amount of this reversal in the minor unit of the transaction's presentment
	// currency. For dollars, for example, this is cents.
	ReversalPresentmentAmount int64 `json:"reversal_presentment_amount,required"`
	// Why this reversal was initiated.
	ReversalReason CardPaymentElementsCardReversalReversalReason `json:"reversal_reason,required,nullable"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_reversal`.
	Type CardPaymentElementsCardReversalType `json:"type,required"`
	// The amount left pending on the Card Authorization in the minor unit of the
	// transaction's currency. For dollars, for example, this is cents.
	UpdatedAuthorizationAmount int64 `json:"updated_authorization_amount,required"`
	// The amount left pending on the Card Authorization in the minor unit of the
	// transaction's presentment currency. For dollars, for example, this is cents.
	UpdatedAuthorizationPresentmentAmount int64                               `json:"updated_authorization_presentment_amount,required"`
	JSON                                  cardPaymentElementsCardReversalJSON `json:"-"`
}

// cardPaymentElementsCardReversalJSON contains the JSON metadata for the struct
// [CardPaymentElementsCardReversal]
type cardPaymentElementsCardReversalJSON struct {
	ID                                    apijson.Field
	CardAuthorizationID                   apijson.Field
	Currency                              apijson.Field
	MerchantAcceptorID                    apijson.Field
	MerchantCategoryCode                  apijson.Field
	MerchantCity                          apijson.Field
	MerchantCountry                       apijson.Field
	MerchantDescriptor                    apijson.Field
	MerchantPostalCode                    apijson.Field
	MerchantState                         apijson.Field
	Network                               apijson.Field
	NetworkIdentifiers                    apijson.Field
	PendingTransactionID                  apijson.Field
	PresentmentCurrency                   apijson.Field
	ReversalAmount                        apijson.Field
	ReversalPresentmentAmount             apijson.Field
	ReversalReason                        apijson.Field
	TerminalID                            apijson.Field
	Type                                  apijson.Field
	UpdatedAuthorizationAmount            apijson.Field
	UpdatedAuthorizationPresentmentAmount apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
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
	CardPaymentElementsCardReversalCurrencyCad CardPaymentElementsCardReversalCurrency = "CAD"
	CardPaymentElementsCardReversalCurrencyChf CardPaymentElementsCardReversalCurrency = "CHF"
	CardPaymentElementsCardReversalCurrencyEur CardPaymentElementsCardReversalCurrency = "EUR"
	CardPaymentElementsCardReversalCurrencyGbp CardPaymentElementsCardReversalCurrency = "GBP"
	CardPaymentElementsCardReversalCurrencyJpy CardPaymentElementsCardReversalCurrency = "JPY"
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
	CardPaymentElementsCardReversalReversalReasonReversedByCustomer          CardPaymentElementsCardReversalReversalReason = "reversed_by_customer"
	CardPaymentElementsCardReversalReversalReasonReversedByNetworkOrAcquirer CardPaymentElementsCardReversalReversalReason = "reversed_by_network_or_acquirer"
	CardPaymentElementsCardReversalReversalReasonReversedByPointOfSale       CardPaymentElementsCardReversalReversalReason = "reversed_by_point_of_sale"
	CardPaymentElementsCardReversalReversalReasonPartialReversal             CardPaymentElementsCardReversalReversalReason = "partial_reversal"
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
// only if `category` is equal to `card_settlement`. Card Settlements are card
// transactions that have cleared and settled. While a settlement is usually
// preceded by an authorization, an acquirer can also directly clear a transaction
// without first authorizing it.
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
	// Cashback earned on this transaction, if eligible. Cashback is paid out in
	// aggregate, monthly.
	Cashback CardPaymentElementsCardSettlementCashback `json:"cashback,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency CardPaymentElementsCardSettlementCurrency `json:"currency,required"`
	// Interchange assessed as a part of this transaction.
	Interchange CardPaymentElementsCardSettlementInterchange `json:"interchange,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required"`
	// The merchant's postal code. For US merchants this is always a 5-digit ZIP code.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The card network on which this transaction was processed.
	Network CardPaymentElementsCardSettlementNetwork `json:"network,required"`
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
	// Surcharge amount details, if applicable. The amount is positive if the surcharge
	// is added to to the overall transaction amount (surcharge), and negative if the
	// surcharge is deducted from the overall transaction amount (discount).
	Surcharge CardPaymentElementsCardSettlementSurcharge `json:"surcharge,required,nullable"`
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
	Cashback             apijson.Field
	Currency             apijson.Field
	Interchange          apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantPostalCode   apijson.Field
	MerchantState        apijson.Field
	Network              apijson.Field
	NetworkIdentifiers   apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
	Surcharge            apijson.Field
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

// Cashback earned on this transaction, if eligible. Cashback is paid out in
// aggregate, monthly.
type CardPaymentElementsCardSettlementCashback struct {
	// The cashback amount given as a string containing a decimal number. The amount is
	// a positive number if it will be credited to you (e.g., settlements) and a
	// negative number if it will be debited (e.g., refunds).
	Amount string `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
	Currency CardPaymentElementsCardSettlementCashbackCurrency `json:"currency,required"`
	JSON     cardPaymentElementsCardSettlementCashbackJSON     `json:"-"`
}

// cardPaymentElementsCardSettlementCashbackJSON contains the JSON metadata for the
// struct [CardPaymentElementsCardSettlementCashback]
type cardPaymentElementsCardSettlementCashbackJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementCashback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementCashbackJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
type CardPaymentElementsCardSettlementCashbackCurrency string

const (
	CardPaymentElementsCardSettlementCashbackCurrencyCad CardPaymentElementsCardSettlementCashbackCurrency = "CAD"
	CardPaymentElementsCardSettlementCashbackCurrencyChf CardPaymentElementsCardSettlementCashbackCurrency = "CHF"
	CardPaymentElementsCardSettlementCashbackCurrencyEur CardPaymentElementsCardSettlementCashbackCurrency = "EUR"
	CardPaymentElementsCardSettlementCashbackCurrencyGbp CardPaymentElementsCardSettlementCashbackCurrency = "GBP"
	CardPaymentElementsCardSettlementCashbackCurrencyJpy CardPaymentElementsCardSettlementCashbackCurrency = "JPY"
	CardPaymentElementsCardSettlementCashbackCurrencyUsd CardPaymentElementsCardSettlementCashbackCurrency = "USD"
)

func (r CardPaymentElementsCardSettlementCashbackCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementCashbackCurrencyCad, CardPaymentElementsCardSettlementCashbackCurrencyChf, CardPaymentElementsCardSettlementCashbackCurrencyEur, CardPaymentElementsCardSettlementCashbackCurrencyGbp, CardPaymentElementsCardSettlementCashbackCurrencyJpy, CardPaymentElementsCardSettlementCashbackCurrencyUsd:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type CardPaymentElementsCardSettlementCurrency string

const (
	CardPaymentElementsCardSettlementCurrencyCad CardPaymentElementsCardSettlementCurrency = "CAD"
	CardPaymentElementsCardSettlementCurrencyChf CardPaymentElementsCardSettlementCurrency = "CHF"
	CardPaymentElementsCardSettlementCurrencyEur CardPaymentElementsCardSettlementCurrency = "EUR"
	CardPaymentElementsCardSettlementCurrencyGbp CardPaymentElementsCardSettlementCurrency = "GBP"
	CardPaymentElementsCardSettlementCurrencyJpy CardPaymentElementsCardSettlementCurrency = "JPY"
	CardPaymentElementsCardSettlementCurrencyUsd CardPaymentElementsCardSettlementCurrency = "USD"
)

func (r CardPaymentElementsCardSettlementCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementCurrencyCad, CardPaymentElementsCardSettlementCurrencyChf, CardPaymentElementsCardSettlementCurrencyEur, CardPaymentElementsCardSettlementCurrencyGbp, CardPaymentElementsCardSettlementCurrencyJpy, CardPaymentElementsCardSettlementCurrencyUsd:
		return true
	}
	return false
}

// Interchange assessed as a part of this transaction.
type CardPaymentElementsCardSettlementInterchange struct {
	// The interchange amount given as a string containing a decimal number in major
	// units (so e.g., "3.14" for $3.14). The amount is a positive number if it is
	// credited to Increase (e.g., settlements) and a negative number if it is debited
	// (e.g., refunds).
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
	CardPaymentElementsCardSettlementInterchangeCurrencyCad CardPaymentElementsCardSettlementInterchangeCurrency = "CAD"
	CardPaymentElementsCardSettlementInterchangeCurrencyChf CardPaymentElementsCardSettlementInterchangeCurrency = "CHF"
	CardPaymentElementsCardSettlementInterchangeCurrencyEur CardPaymentElementsCardSettlementInterchangeCurrency = "EUR"
	CardPaymentElementsCardSettlementInterchangeCurrencyGbp CardPaymentElementsCardSettlementInterchangeCurrency = "GBP"
	CardPaymentElementsCardSettlementInterchangeCurrencyJpy CardPaymentElementsCardSettlementInterchangeCurrency = "JPY"
	CardPaymentElementsCardSettlementInterchangeCurrencyUsd CardPaymentElementsCardSettlementInterchangeCurrency = "USD"
)

func (r CardPaymentElementsCardSettlementInterchangeCurrency) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementInterchangeCurrencyCad, CardPaymentElementsCardSettlementInterchangeCurrencyChf, CardPaymentElementsCardSettlementInterchangeCurrencyEur, CardPaymentElementsCardSettlementInterchangeCurrencyGbp, CardPaymentElementsCardSettlementInterchangeCurrencyJpy, CardPaymentElementsCardSettlementInterchangeCurrencyUsd:
		return true
	}
	return false
}

// The card network on which this transaction was processed.
type CardPaymentElementsCardSettlementNetwork string

const (
	CardPaymentElementsCardSettlementNetworkVisa CardPaymentElementsCardSettlementNetwork = "visa"
)

func (r CardPaymentElementsCardSettlementNetwork) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementNetworkVisa:
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
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge    CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesGas              CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage     CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn       CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee CardPaymentElementsCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
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
	CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable               CardPaymentElementsCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
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
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant    CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop      CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar       CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesTelephone     CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesOther         CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraChargesLaundry       CardPaymentElementsCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
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
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow        CardPaymentElementsCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
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
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText              CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber           CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber      CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber         CardPaymentElementsCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
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
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit                                                        CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther                                                           CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
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
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone                  CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService        CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee            CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee             CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo                 CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset          CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer         CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard              CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport       CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge                CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical               CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage          CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther                 CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee    CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets                  CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees              CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby               CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee            CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore                 CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService         CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel   CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades              CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi                  CardPaymentElementsCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
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
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit                                                        CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation                                       CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther                                                           CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket                                    CardPaymentElementsCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
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
	CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions                CardPaymentElementsCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
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
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone                   CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket              CardPaymentElementsCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
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
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone               CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed    CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

func (r CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone, CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed, CardPaymentElementsCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed:
		return true
	}
	return false
}

// Surcharge amount details, if applicable. The amount is positive if the surcharge
// is added to to the overall transaction amount (surcharge), and negative if the
// surcharge is deducted from the overall transaction amount (discount).
type CardPaymentElementsCardSettlementSurcharge struct {
	// The surcharge amount in the minor unit of the transaction's settlement currency.
	Amount int64 `json:"amount,required"`
	// The surcharge amount in the minor unit of the transaction's presentment
	// currency.
	PresentmentAmount int64                                          `json:"presentment_amount,required"`
	JSON              cardPaymentElementsCardSettlementSurchargeJSON `json:"-"`
}

// cardPaymentElementsCardSettlementSurchargeJSON contains the JSON metadata for
// the struct [CardPaymentElementsCardSettlementSurcharge]
type cardPaymentElementsCardSettlementSurchargeJSON struct {
	Amount            apijson.Field
	PresentmentAmount apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CardPaymentElementsCardSettlementSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardSettlementSurchargeJSON) RawJSON() string {
	return r.raw
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

// An Inbound Card Validation object. This field will be present in the JSON
// response if and only if `category` is equal to `card_validation`. Inbound Card
// Validations are requests from a merchant to verify that a card number and
// optionally its address and/or Card Verification Value are valid.
type CardPaymentElementsCardValidation struct {
	// The Card Validation identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner CardPaymentElementsCardValidationActioner `json:"actioner,required"`
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts CardPaymentElementsCardValidationAdditionalAmounts `json:"additional_amounts,required"`
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
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_validation`.
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
	AdditionalAmounts    apijson.Field
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
	TerminalID           apijson.Field
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
	CardPaymentElementsCardValidationActionerUser     CardPaymentElementsCardValidationActioner = "user"
	CardPaymentElementsCardValidationActionerIncrease CardPaymentElementsCardValidationActioner = "increase"
	CardPaymentElementsCardValidationActionerNetwork  CardPaymentElementsCardValidationActioner = "network"
)

func (r CardPaymentElementsCardValidationActioner) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationActionerUser, CardPaymentElementsCardValidationActionerIncrease, CardPaymentElementsCardValidationActionerNetwork:
		return true
	}
	return false
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type CardPaymentElementsCardValidationAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic CardPaymentElementsCardValidationAdditionalAmountsClinic `json:"clinic,required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental CardPaymentElementsCardValidationAdditionalAmountsDental `json:"dental,required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription CardPaymentElementsCardValidationAdditionalAmountsPrescription `json:"prescription,required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge CardPaymentElementsCardValidationAdditionalAmountsSurcharge `json:"surcharge,required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative CardPaymentElementsCardValidationAdditionalAmountsTotalCumulative `json:"total_cumulative,required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare CardPaymentElementsCardValidationAdditionalAmountsTotalHealthcare `json:"total_healthcare,required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit CardPaymentElementsCardValidationAdditionalAmountsTransit `json:"transit,required,nullable"`
	// An unknown additional amount.
	Unknown CardPaymentElementsCardValidationAdditionalAmountsUnknown `json:"unknown,required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision CardPaymentElementsCardValidationAdditionalAmountsVision `json:"vision,required,nullable"`
	JSON   cardPaymentElementsCardValidationAdditionalAmountsJSON   `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardValidationAdditionalAmounts]
type cardPaymentElementsCardValidationAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
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

func (r *CardPaymentElementsCardValidationAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type CardPaymentElementsCardValidationAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsClinicJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsClinicJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsClinic]
type cardPaymentElementsCardValidationAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type CardPaymentElementsCardValidationAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsDentalJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsDentalJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsDental]
type cardPaymentElementsCardValidationAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type CardPaymentElementsCardValidationAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsPrescriptionJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsPrescriptionJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsPrescription]
type cardPaymentElementsCardValidationAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type CardPaymentElementsCardValidationAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsSurchargeJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsSurchargeJSON contains the
// JSON metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsSurcharge]
type cardPaymentElementsCardValidationAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type CardPaymentElementsCardValidationAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsTotalCumulativeJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsTotalCumulative]
type cardPaymentElementsCardValidationAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type CardPaymentElementsCardValidationAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsTotalHealthcareJSON contains
// the JSON metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsTotalHealthcare]
type cardPaymentElementsCardValidationAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type CardPaymentElementsCardValidationAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                        `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsTransitJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsTransitJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsTransit]
type cardPaymentElementsCardValidationAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type CardPaymentElementsCardValidationAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                        `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsUnknownJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsUnknownJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsUnknown]
type cardPaymentElementsCardValidationAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type CardPaymentElementsCardValidationAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     cardPaymentElementsCardValidationAdditionalAmountsVisionJSON `json:"-"`
}

// cardPaymentElementsCardValidationAdditionalAmountsVisionJSON contains the JSON
// metadata for the struct
// [CardPaymentElementsCardValidationAdditionalAmountsVision]
type cardPaymentElementsCardValidationAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPaymentElementsCardValidationAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPaymentElementsCardValidationAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CardPaymentElementsCardValidationCurrency string

const (
	CardPaymentElementsCardValidationCurrencyCad CardPaymentElementsCardValidationCurrency = "CAD"
	CardPaymentElementsCardValidationCurrencyChf CardPaymentElementsCardValidationCurrency = "CHF"
	CardPaymentElementsCardValidationCurrencyEur CardPaymentElementsCardValidationCurrency = "EUR"
	CardPaymentElementsCardValidationCurrencyGbp CardPaymentElementsCardValidationCurrency = "GBP"
	CardPaymentElementsCardValidationCurrencyJpy CardPaymentElementsCardValidationCurrency = "JPY"
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
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason,required,nullable"`
	JSON                    cardPaymentElementsCardValidationNetworkDetailsVisaJSON                    `json:"-"`
}

// cardPaymentElementsCardValidationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [CardPaymentElementsCardValidationNetworkDetailsVisa]
type cardPaymentElementsCardValidationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
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
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    CardPaymentElementsCardValidationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeUnknown                    CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeManual                     CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactless                CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeUnknown, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeManual, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactless, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, CardPaymentElementsCardValidationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason string

const (
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonIssuerError                                              CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonOther                                                    CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonIssuerError, CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonInternalVisaError, CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, CardPaymentElementsCardValidationNetworkDetailsVisaStandInProcessingReasonOther:
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
// `inbound_card_validation`.
type CardPaymentElementsCardValidationType string

const (
	CardPaymentElementsCardValidationTypeInboundCardValidation CardPaymentElementsCardValidationType = "inbound_card_validation"
)

func (r CardPaymentElementsCardValidationType) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationTypeInboundCardValidation:
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
	CardPaymentElementsCardValidationVerificationCardVerificationCodeResultNotChecked CardPaymentElementsCardValidationVerificationCardVerificationCodeResult = "not_checked"
	CardPaymentElementsCardValidationVerificationCardVerificationCodeResultMatch      CardPaymentElementsCardValidationVerificationCardVerificationCodeResult = "match"
	CardPaymentElementsCardValidationVerificationCardVerificationCodeResultNoMatch    CardPaymentElementsCardValidationVerificationCardVerificationCodeResult = "no_match"
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
	CardPaymentElementsCardValidationVerificationCardholderAddressResultNotChecked                       CardPaymentElementsCardValidationVerificationCardholderAddressResult = "not_checked"
	CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    CardPaymentElementsCardValidationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    CardPaymentElementsCardValidationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	CardPaymentElementsCardValidationVerificationCardholderAddressResultMatch                            CardPaymentElementsCardValidationVerificationCardholderAddressResult = "match"
	CardPaymentElementsCardValidationVerificationCardholderAddressResultNoMatch                          CardPaymentElementsCardValidationVerificationCardholderAddressResult = "no_match"
	CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked CardPaymentElementsCardValidationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r CardPaymentElementsCardValidationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case CardPaymentElementsCardValidationVerificationCardholderAddressResultNotChecked, CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultNoMatch, CardPaymentElementsCardValidationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type CardPaymentElementsCategory string

const (
	CardPaymentElementsCategoryCardAuthorization           CardPaymentElementsCategory = "card_authorization"
	CardPaymentElementsCategoryCardAuthentication          CardPaymentElementsCategory = "card_authentication"
	CardPaymentElementsCategoryCardValidation              CardPaymentElementsCategory = "card_validation"
	CardPaymentElementsCategoryCardDecline                 CardPaymentElementsCategory = "card_decline"
	CardPaymentElementsCategoryCardReversal                CardPaymentElementsCategory = "card_reversal"
	CardPaymentElementsCategoryCardAuthorizationExpiration CardPaymentElementsCategory = "card_authorization_expiration"
	CardPaymentElementsCategoryCardIncrement               CardPaymentElementsCategory = "card_increment"
	CardPaymentElementsCategoryCardSettlement              CardPaymentElementsCategory = "card_settlement"
	CardPaymentElementsCategoryCardRefund                  CardPaymentElementsCategory = "card_refund"
	CardPaymentElementsCategoryCardFuelConfirmation        CardPaymentElementsCategory = "card_fuel_confirmation"
	CardPaymentElementsCategoryOther                       CardPaymentElementsCategory = "other"
)

func (r CardPaymentElementsCategory) IsKnown() bool {
	switch r {
	case CardPaymentElementsCategoryCardAuthorization, CardPaymentElementsCategoryCardAuthentication, CardPaymentElementsCategoryCardValidation, CardPaymentElementsCategoryCardDecline, CardPaymentElementsCategoryCardReversal, CardPaymentElementsCategoryCardAuthorizationExpiration, CardPaymentElementsCategoryCardIncrement, CardPaymentElementsCategoryCardSettlement, CardPaymentElementsCategoryCardRefund, CardPaymentElementsCategoryCardFuelConfirmation, CardPaymentElementsCategoryOther:
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
	// The total refunded amount in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	RefundedAmount int64 `json:"refunded_amount,required"`
	// The total reversed amount in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	ReversedAmount int64 `json:"reversed_amount,required"`
	// The total settled amount in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	SettledAmount int64                `json:"settled_amount,required"`
	JSON          cardPaymentStateJSON `json:"-"`
}

// cardPaymentStateJSON contains the JSON metadata for the struct
// [CardPaymentState]
type cardPaymentStateJSON struct {
	AuthorizedAmount    apijson.Field
	FuelConfirmedAmount apijson.Field
	IncrementedAmount   apijson.Field
	RefundedAmount      apijson.Field
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
