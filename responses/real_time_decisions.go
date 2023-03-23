package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt time.Time `json:"timeout_at,required" format:"date-time"`
	// The status of the Real-Time Decision.
	Status RealTimeDecisionStatus `json:"status,required"`
	// The category of the Real-Time Decision.
	Category RealTimeDecisionCategory `json:"category,required"`
	// Fields related to a card authorization.
	CardAuthorization RealTimeDecisionCardAuthorization `json:"card_authorization,required,nullable"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken RealTimeDecisionDigitalWalletToken `json:"digital_wallet_token,required,nullable"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication RealTimeDecisionDigitalWalletAuthentication `json:"digital_wallet_authentication,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type RealTimeDecisionType `json:"type,required"`
	JSON RealTimeDecisionJSON
}

type RealTimeDecisionJSON struct {
	ID                          pjson.Metadata
	CreatedAt                   pjson.Metadata
	TimeoutAt                   pjson.Metadata
	Status                      pjson.Metadata
	Category                    pjson.Metadata
	CardAuthorization           pjson.Metadata
	DigitalWalletToken          pjson.Metadata
	DigitalWalletAuthentication pjson.Metadata
	Type                        pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RealTimeDecision using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RealTimeDecision) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimeDecisionStatus string

const (
	RealTimeDecisionStatusPending   RealTimeDecisionStatus = "pending"
	RealTimeDecisionStatusResponded RealTimeDecisionStatus = "responded"
	RealTimeDecisionStatusTimedOut  RealTimeDecisionStatus = "timed_out"
)

type RealTimeDecisionCategory string

const (
	RealTimeDecisionCategoryCardAuthorizationRequested           RealTimeDecisionCategory = "card_authorization_requested"
	RealTimeDecisionCategoryDigitalWalletTokenRequested          RealTimeDecisionCategory = "digital_wallet_token_requested"
	RealTimeDecisionCategoryDigitalWalletAuthenticationRequested RealTimeDecisionCategory = "digital_wallet_authentication_requested"
)

type RealTimeDecisionCardAuthorization struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network RealTimeDecisionCardAuthorizationNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails RealTimeDecisionCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Whether or not the authorization was approved.
	Decision RealTimeDecisionCardAuthorizationDecision `json:"decision,required,nullable"`
	// The identifier of the Card that is being authorized.
	CardID string `json:"card_id,required"`
	// The identifier of the Account the authorization will debit.
	AccountID string `json:"account_id,required"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount int64 `json:"settlement_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency string `json:"settlement_currency,required"`
	JSON               RealTimeDecisionCardAuthorizationJSON
}

type RealTimeDecisionCardAuthorizationJSON struct {
	MerchantAcceptorID   pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	Network              pjson.Metadata
	NetworkDetails       pjson.Metadata
	Decision             pjson.Metadata
	CardID               pjson.Metadata
	AccountID            pjson.Metadata
	PresentmentAmount    pjson.Metadata
	PresentmentCurrency  pjson.Metadata
	SettlementAmount     pjson.Metadata
	SettlementCurrency   pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionCardAuthorization using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimeDecisionCardAuthorizationNetwork string

const (
	RealTimeDecisionCardAuthorizationNetworkVisa RealTimeDecisionCardAuthorizationNetwork = "visa"
)

type RealTimeDecisionCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa RealTimeDecisionCardAuthorizationNetworkDetailsVisa `json:"visa,required"`
	JSON RealTimeDecisionCardAuthorizationNetworkDetailsJSON
}

type RealTimeDecisionCardAuthorizationNetworkDetailsJSON struct {
	Visa pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionCardAuthorizationNetworkDetails using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimeDecisionCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    RealTimeDecisionCardAuthorizationNetworkDetailsVisaJSON
}

type RealTimeDecisionCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator pjson.Metadata
	PointOfServiceEntryMode     pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionCardAuthorizationNetworkDetailsVisa using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type RealTimeDecisionCardAuthorizationDecision string

const (
	RealTimeDecisionCardAuthorizationDecisionApprove RealTimeDecisionCardAuthorizationDecision = "approve"
	RealTimeDecisionCardAuthorizationDecisionDecline RealTimeDecisionCardAuthorizationDecision = "decline"
)

type RealTimeDecisionDigitalWalletToken struct {
	// Whether or not the provisioning request was approved. This will be null until
	// the real time decision is responded to.
	Decision RealTimeDecisionDigitalWalletTokenDecision `json:"decision,required,nullable"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet `json:"digital_wallet,required"`
	// The identifier of the Card Profile that was set via the real time decision. This
	// will be null until the real time decision is responded to or if the real time
	// decision did not set a card profile.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	JSON          RealTimeDecisionDigitalWalletTokenJSON
}

type RealTimeDecisionDigitalWalletTokenJSON struct {
	Decision      pjson.Metadata
	CardID        pjson.Metadata
	DigitalWallet pjson.Metadata
	CardProfileID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionDigitalWalletToken using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimeDecisionDigitalWalletTokenDecision string

const (
	RealTimeDecisionDigitalWalletTokenDecisionApprove RealTimeDecisionDigitalWalletTokenDecision = "approve"
	RealTimeDecisionDigitalWalletTokenDecisionDecline RealTimeDecisionDigitalWalletTokenDecision = "decline"
)

type RealTimeDecisionDigitalWalletTokenDigitalWallet string

const (
	RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay  RealTimeDecisionDigitalWalletTokenDigitalWallet = "apple_pay"
	RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay RealTimeDecisionDigitalWalletTokenDigitalWallet = "google_pay"
)

type RealTimeDecisionDigitalWalletAuthentication struct {
	// Whether your application successfully delivered the one-time passcode.
	Result RealTimeDecisionDigitalWalletAuthenticationResult `json:"result,required,nullable"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletAuthenticationDigitalWallet `json:"digital_wallet,required"`
	// The channel to send the card user their one-time passcode.
	Channel RealTimeDecisionDigitalWalletAuthenticationChannel `json:"channel,required"`
	// The one-time passcode to send the card user.
	OneTimePasscode string `json:"one_time_passcode,required"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone string `json:"phone,required,nullable"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email string `json:"email,required,nullable"`
	JSON  RealTimeDecisionDigitalWalletAuthenticationJSON
}

type RealTimeDecisionDigitalWalletAuthenticationJSON struct {
	Result          pjson.Metadata
	CardID          pjson.Metadata
	DigitalWallet   pjson.Metadata
	Channel         pjson.Metadata
	OneTimePasscode pjson.Metadata
	Phone           pjson.Metadata
	Email           pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionDigitalWalletAuthentication using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionDigitalWalletAuthentication) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimeDecisionDigitalWalletAuthenticationResult string

const (
	RealTimeDecisionDigitalWalletAuthenticationResultSuccess RealTimeDecisionDigitalWalletAuthenticationResult = "success"
	RealTimeDecisionDigitalWalletAuthenticationResultFailure RealTimeDecisionDigitalWalletAuthenticationResult = "failure"
)

type RealTimeDecisionDigitalWalletAuthenticationDigitalWallet string

const (
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay  RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "apple_pay"
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "google_pay"
)

type RealTimeDecisionDigitalWalletAuthenticationChannel string

const (
	RealTimeDecisionDigitalWalletAuthenticationChannelSMS   RealTimeDecisionDigitalWalletAuthenticationChannel = "sms"
	RealTimeDecisionDigitalWalletAuthenticationChannelEmail RealTimeDecisionDigitalWalletAuthenticationChannel = "email"
)

type RealTimeDecisionType string

const (
	RealTimeDecisionTypeRealTimeDecision RealTimeDecisionType = "real_time_decision"
)
