package requests

import (
	"fmt"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID fields.Field[string] `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt fields.Field[time.Time] `json:"timeout_at,required" format:"date-time"`
	// The status of the Real-Time Decision.
	Status fields.Field[RealTimeDecisionStatus] `json:"status,required"`
	// The category of the Real-Time Decision.
	Category fields.Field[RealTimeDecisionCategory] `json:"category,required"`
	// Fields related to a card authorization.
	CardAuthorization fields.Field[RealTimeDecisionCardAuthorization] `json:"card_authorization,required,nullable"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken fields.Field[RealTimeDecisionDigitalWalletToken] `json:"digital_wallet_token,required,nullable"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication fields.Field[RealTimeDecisionDigitalWalletAuthentication] `json:"digital_wallet_authentication,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type fields.Field[RealTimeDecisionType] `json:"type,required"`
}

// MarshalJSON serializes RealTimeDecision into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RealTimeDecision) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RealTimeDecision) String() (result string) {
	return fmt.Sprintf("&RealTimeDecision{ID:%s CreatedAt:%s TimeoutAt:%s Status:%s Category:%s CardAuthorization:%s DigitalWalletToken:%s DigitalWalletAuthentication:%s Type:%s}", r.ID, r.CreatedAt, r.TimeoutAt, r.Status, r.Category, r.CardAuthorization, r.DigitalWalletToken, r.DigitalWalletAuthentication, r.Type)
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
	MerchantAcceptorID fields.Field[string] `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor fields.Field[string] `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode fields.Field[string] `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity fields.Field[string] `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry fields.Field[string] `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network fields.Field[RealTimeDecisionCardAuthorizationNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[RealTimeDecisionCardAuthorizationNetworkDetails] `json:"network_details,required"`
	// Whether or not the authorization was approved.
	Decision fields.Field[RealTimeDecisionCardAuthorizationDecision] `json:"decision,required,nullable"`
	// The identifier of the Card that is being authorized.
	CardID fields.Field[string] `json:"card_id,required"`
	// The identifier of the Account the authorization will debit.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount fields.Field[int64] `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency fields.Field[string] `json:"presentment_currency,required"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount fields.Field[int64] `json:"settlement_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency fields.Field[string] `json:"settlement_currency,required"`
}

// MarshalJSON serializes RealTimeDecisionCardAuthorization into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *RealTimeDecisionCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RealTimeDecisionCardAuthorization) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionCardAuthorization{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Decision:%s CardID:%s AccountID:%s PresentmentAmount:%s PresentmentCurrency:%s SettlementAmount:%s SettlementCurrency:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Decision, r.CardID, r.AccountID, r.PresentmentAmount, r.PresentmentCurrency, r.SettlementAmount, r.SettlementCurrency)
}

type RealTimeDecisionCardAuthorizationNetwork string

const (
	RealTimeDecisionCardAuthorizationNetworkVisa RealTimeDecisionCardAuthorizationNetwork = "visa"
)

type RealTimeDecisionCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[RealTimeDecisionCardAuthorizationNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes RealTimeDecisionCardAuthorizationNetworkDetails into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *RealTimeDecisionCardAuthorizationNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RealTimeDecisionCardAuthorizationNetworkDetails) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionCardAuthorizationNetworkDetails{Visa:%s}", r.Visa)
}

type RealTimeDecisionCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes RealTimeDecisionCardAuthorizationNetworkDetailsVisa into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionCardAuthorizationNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
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
	Decision fields.Field[RealTimeDecisionDigitalWalletTokenDecision] `json:"decision,required,nullable"`
	// The identifier of the Card that is being tokenized.
	CardID fields.Field[string] `json:"card_id,required"`
	// The digital wallet app being used.
	DigitalWallet fields.Field[RealTimeDecisionDigitalWalletTokenDigitalWallet] `json:"digital_wallet,required"`
	// The identifier of the Card Profile that was set via the real time decision. This
	// will be null until the real time decision is responded to or if the real time
	// decision did not set a card profile.
	CardProfileID fields.Field[string] `json:"card_profile_id,required,nullable"`
}

// MarshalJSON serializes RealTimeDecisionDigitalWalletToken into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *RealTimeDecisionDigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RealTimeDecisionDigitalWalletToken) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionDigitalWalletToken{Decision:%s CardID:%s DigitalWallet:%s CardProfileID:%s}", r.Decision, r.CardID, r.DigitalWallet, r.CardProfileID)
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
	Result fields.Field[RealTimeDecisionDigitalWalletAuthenticationResult] `json:"result,required,nullable"`
	// The identifier of the Card that is being tokenized.
	CardID fields.Field[string] `json:"card_id,required"`
	// The digital wallet app being used.
	DigitalWallet fields.Field[RealTimeDecisionDigitalWalletAuthenticationDigitalWallet] `json:"digital_wallet,required"`
	// The channel to send the card user their one-time passcode.
	Channel fields.Field[RealTimeDecisionDigitalWalletAuthenticationChannel] `json:"channel,required"`
	// The one-time passcode to send the card user.
	OneTimePasscode fields.Field[string] `json:"one_time_passcode,required"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone fields.Field[string] `json:"phone,required,nullable"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email fields.Field[string] `json:"email,required,nullable"`
}

// MarshalJSON serializes RealTimeDecisionDigitalWalletAuthentication into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *RealTimeDecisionDigitalWalletAuthentication) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RealTimeDecisionDigitalWalletAuthentication) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionDigitalWalletAuthentication{Result:%s CardID:%s DigitalWallet:%s Channel:%s OneTimePasscode:%s Phone:%s Email:%s}", r.Result, r.CardID, r.DigitalWallet, r.Channel, r.OneTimePasscode, r.Phone, r.Email)
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

type ActionARealTimeDecisionParameters struct {
	// If the Real-Time Decision relates to a card authorization attempt, this object
	// contains your response to the authorization.
	CardAuthorization fields.Field[ActionARealTimeDecisionParametersCardAuthorization] `json:"card_authorization"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken fields.Field[ActionARealTimeDecisionParametersDigitalWalletToken] `json:"digital_wallet_token"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication fields.Field[ActionARealTimeDecisionParametersDigitalWalletAuthentication] `json:"digital_wallet_authentication"`
}

// MarshalJSON serializes ActionARealTimeDecisionParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ActionARealTimeDecisionParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ActionARealTimeDecisionParameters) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParameters{CardAuthorization:%s DigitalWalletToken:%s DigitalWalletAuthentication:%s}", r.CardAuthorization, r.DigitalWalletToken, r.DigitalWalletAuthentication)
}

type ActionARealTimeDecisionParametersCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision fields.Field[ActionARealTimeDecisionParametersCardAuthorizationDecision] `json:"decision,required"`
}

func (r ActionARealTimeDecisionParametersCardAuthorization) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersCardAuthorization{Decision:%s}", r.Decision)
}

type ActionARealTimeDecisionParametersCardAuthorizationDecision string

const (
	ActionARealTimeDecisionParametersCardAuthorizationDecisionApprove ActionARealTimeDecisionParametersCardAuthorizationDecision = "approve"
	ActionARealTimeDecisionParametersCardAuthorizationDecisionDecline ActionARealTimeDecisionParametersCardAuthorizationDecision = "decline"
)

type ActionARealTimeDecisionParametersDigitalWalletToken struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval fields.Field[ActionARealTimeDecisionParametersDigitalWalletTokenApproval] `json:"approval"`
	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline fields.Field[ActionARealTimeDecisionParametersDigitalWalletTokenDecline] `json:"decline"`
}

func (r ActionARealTimeDecisionParametersDigitalWalletToken) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletToken{Approval:%s Decline:%s}", r.Approval, r.Decline)
}

type ActionARealTimeDecisionParametersDigitalWalletTokenApproval struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID fields.Field[string] `json:"card_profile_id,required"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone fields.Field[string] `json:"phone"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email fields.Field[string] `json:"email"`
}

func (r ActionARealTimeDecisionParametersDigitalWalletTokenApproval) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletTokenApproval{CardProfileID:%s Phone:%s Email:%s}", r.CardProfileID, r.Phone, r.Email)
}

type ActionARealTimeDecisionParametersDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason fields.Field[string] `json:"reason"`
}

func (r ActionARealTimeDecisionParametersDigitalWalletTokenDecline) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletTokenDecline{Reason:%s}", r.Reason)
}

type ActionARealTimeDecisionParametersDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result fields.Field[ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult] `json:"result,required"`
}

func (r ActionARealTimeDecisionParametersDigitalWalletAuthentication) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletAuthentication{Result:%s}", r.Result)
}

type ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult string

const (
	ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultSuccess ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult = "success"
	ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultFailure ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult = "failure"
)
