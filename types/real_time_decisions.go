package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID *string `pjson:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt *string `pjson:"created_at"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt *string `pjson:"timeout_at"`
	// The status of the Real-Time Decision.
	Status *RealTimeDecisionStatus `pjson:"status"`
	// The category of the Real-Time Decision.
	Category *RealTimeDecisionCategory `pjson:"category"`
	// Fields related to a card authorization.
	CardAuthorization *RealTimeDecisionCardAuthorization `pjson:"card_authorization"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken *RealTimeDecisionDigitalWalletToken `pjson:"digital_wallet_token"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication *RealTimeDecisionDigitalWalletAuthentication `pjson:"digital_wallet_authentication"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type       *RealTimeDecisionType  `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into RealTimeDecision using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RealTimeDecision) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RealTimeDecision into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RealTimeDecision) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Real-Time Decision identifier.
func (r *RealTimeDecision) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Real-Time Decision was created.
func (r *RealTimeDecision) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// your application can no longer respond to the Real-Time Decision.
func (r *RealTimeDecision) GetTimeoutAt() (TimeoutAt string) {
	if r != nil && r.TimeoutAt != nil {
		TimeoutAt = *r.TimeoutAt
	}
	return
}

// The status of the Real-Time Decision.
func (r *RealTimeDecision) GetStatus() (Status RealTimeDecisionStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The category of the Real-Time Decision.
func (r *RealTimeDecision) GetCategory() (Category RealTimeDecisionCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// Fields related to a card authorization.
func (r *RealTimeDecision) GetCardAuthorization() (CardAuthorization RealTimeDecisionCardAuthorization) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// Fields related to a digital wallet token provisioning attempt.
func (r *RealTimeDecision) GetDigitalWalletToken() (DigitalWalletToken RealTimeDecisionDigitalWalletToken) {
	if r != nil && r.DigitalWalletToken != nil {
		DigitalWalletToken = *r.DigitalWalletToken
	}
	return
}

// Fields related to a digital wallet authentication attempt.
func (r *RealTimeDecision) GetDigitalWalletAuthentication() (DigitalWalletAuthentication RealTimeDecisionDigitalWalletAuthentication) {
	if r != nil && r.DigitalWalletAuthentication != nil {
		DigitalWalletAuthentication = *r.DigitalWalletAuthentication
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `real_time_decision`.
func (r *RealTimeDecision) GetType() (Type RealTimeDecisionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r RealTimeDecision) String() (result string) {
	return fmt.Sprintf("&RealTimeDecision{ID:%s CreatedAt:%s TimeoutAt:%s Status:%s Category:%s CardAuthorization:%s DigitalWalletToken:%s DigitalWalletAuthentication:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.CreatedAt), core.FmtP(r.TimeoutAt), core.FmtP(r.Status), core.FmtP(r.Category), r.CardAuthorization, r.DigitalWalletToken, r.DigitalWalletAuthentication, core.FmtP(r.Type))
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
	MerchantAcceptorID *string `pjson:"merchant_acceptor_id"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor *string `pjson:"merchant_descriptor"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode *string `pjson:"merchant_category_code"`
	// The city the merchant resides in.
	MerchantCity *string `pjson:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry *string `pjson:"merchant_country"`
	// The payment network used to process this card authorization
	Network *RealTimeDecisionCardAuthorizationNetwork `pjson:"network"`
	// Fields specific to the `network`
	NetworkDetails *RealTimeDecisionCardAuthorizationNetworkDetails `pjson:"network_details"`
	// Whether or not the authorization was approved.
	Decision *RealTimeDecisionCardAuthorizationDecision `pjson:"decision"`
	// The identifier of the Card that is being authorized.
	CardID *string `pjson:"card_id"`
	// The identifier of the Account the authorization will debit.
	AccountID *string `pjson:"account_id"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount *int64 `pjson:"presentment_amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency *string `pjson:"presentment_currency"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount *int64 `pjson:"settlement_amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency *string                `pjson:"settlement_currency"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionCardAuthorization using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RealTimeDecisionCardAuthorization into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *RealTimeDecisionCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *RealTimeDecisionCardAuthorization) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *RealTimeDecisionCardAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *RealTimeDecisionCardAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The city the merchant resides in.
func (r *RealTimeDecisionCardAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *RealTimeDecisionCardAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

// The payment network used to process this card authorization
func (r *RealTimeDecisionCardAuthorization) GetNetwork() (Network RealTimeDecisionCardAuthorizationNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// Fields specific to the `network`
func (r *RealTimeDecisionCardAuthorization) GetNetworkDetails() (NetworkDetails RealTimeDecisionCardAuthorizationNetworkDetails) {
	if r != nil && r.NetworkDetails != nil {
		NetworkDetails = *r.NetworkDetails
	}
	return
}

// Whether or not the authorization was approved.
func (r *RealTimeDecisionCardAuthorization) GetDecision() (Decision RealTimeDecisionCardAuthorizationDecision) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
}

// The identifier of the Card that is being authorized.
func (r *RealTimeDecisionCardAuthorization) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The identifier of the Account the authorization will debit.
func (r *RealTimeDecisionCardAuthorization) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The amount of the attempted authorization in the currency the card user sees at
// the time of purchase, in the minor unit of that currency. For dollars, for
// example, this is cents.
func (r *RealTimeDecisionCardAuthorization) GetPresentmentAmount() (PresentmentAmount int64) {
	if r != nil && r.PresentmentAmount != nil {
		PresentmentAmount = *r.PresentmentAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
// user sees at the time of purchase.
func (r *RealTimeDecisionCardAuthorization) GetPresentmentCurrency() (PresentmentCurrency string) {
	if r != nil && r.PresentmentCurrency != nil {
		PresentmentCurrency = *r.PresentmentCurrency
	}
	return
}

// The amount of the attempted authorization in the currency it will be settled in.
// This currency is the same as that of the Account the card belongs to.
func (r *RealTimeDecisionCardAuthorization) GetSettlementAmount() (SettlementAmount int64) {
	if r != nil && r.SettlementAmount != nil {
		SettlementAmount = *r.SettlementAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
// transaction will be settled in.
func (r *RealTimeDecisionCardAuthorization) GetSettlementCurrency() (SettlementCurrency string) {
	if r != nil && r.SettlementCurrency != nil {
		SettlementCurrency = *r.SettlementCurrency
	}
	return
}

func (r RealTimeDecisionCardAuthorization) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionCardAuthorization{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Decision:%s CardID:%s AccountID:%s PresentmentAmount:%s PresentmentCurrency:%s SettlementAmount:%s SettlementCurrency:%s}", core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.Network), r.NetworkDetails, core.FmtP(r.Decision), core.FmtP(r.CardID), core.FmtP(r.AccountID), core.FmtP(r.PresentmentAmount), core.FmtP(r.PresentmentCurrency), core.FmtP(r.SettlementAmount), core.FmtP(r.SettlementCurrency))
}

type RealTimeDecisionCardAuthorizationNetwork string

const (
	RealTimeDecisionCardAuthorizationNetworkVisa RealTimeDecisionCardAuthorizationNetwork = "visa"
)

type RealTimeDecisionCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa       *RealTimeDecisionCardAuthorizationNetworkDetailsVisa `pjson:"visa"`
	jsonFields map[string]interface{}                               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionCardAuthorizationNetworkDetails using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RealTimeDecisionCardAuthorizationNetworkDetails into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *RealTimeDecisionCardAuthorizationNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Fields specific to the `visa` network
func (r *RealTimeDecisionCardAuthorizationNetworkDetails) GetVisa() (Visa RealTimeDecisionCardAuthorizationNetworkDetailsVisa) {
	if r != nil && r.Visa != nil {
		Visa = *r.Visa
	}
	return
}

func (r RealTimeDecisionCardAuthorizationNetworkDetails) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionCardAuthorizationNetworkDetails{Visa:%s}", r.Visa)
}

type RealTimeDecisionCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator *RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `pjson:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode *PointOfServiceEntryMode `pjson:"point_of_service_entry_mode"`
	jsonFields              map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionCardAuthorizationNetworkDetailsVisa using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RealTimeDecisionCardAuthorizationNetworkDetailsVisa into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) GetElectronicCommerceIndicator() (ElectronicCommerceIndicator RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) {
	if r != nil && r.ElectronicCommerceIndicator != nil {
		ElectronicCommerceIndicator = *r.ElectronicCommerceIndicator
	}
	return
}

// The method used to enter the cardholder's primary account number and card
// expiration date
func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) GetPointOfServiceEntryMode() (PointOfServiceEntryMode PointOfServiceEntryMode) {
	if r != nil && r.PointOfServiceEntryMode != nil {
		PointOfServiceEntryMode = *r.PointOfServiceEntryMode
	}
	return
}

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionCardAuthorizationNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", core.FmtP(r.ElectronicCommerceIndicator), core.FmtP(r.PointOfServiceEntryMode))
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
	Decision *RealTimeDecisionDigitalWalletTokenDecision `pjson:"decision"`
	// The identifier of the Card that is being tokenized.
	CardID *string `pjson:"card_id"`
	// The digital wallet app being used.
	DigitalWallet *RealTimeDecisionDigitalWalletTokenDigitalWallet `pjson:"digital_wallet"`
	// The identifier of the Card Profile that was set via the real time decision. This
	// will be null until the real time decision is responded to or if the real time
	// decision did not set a card profile.
	CardProfileID *string                `pjson:"card_profile_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionDigitalWalletToken using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RealTimeDecisionDigitalWalletToken into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *RealTimeDecisionDigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Whether or not the provisioning request was approved. This will be null until
// the real time decision is responded to.
func (r *RealTimeDecisionDigitalWalletToken) GetDecision() (Decision RealTimeDecisionDigitalWalletTokenDecision) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
}

// The identifier of the Card that is being tokenized.
func (r *RealTimeDecisionDigitalWalletToken) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The digital wallet app being used.
func (r *RealTimeDecisionDigitalWalletToken) GetDigitalWallet() (DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

// The identifier of the Card Profile that was set via the real time decision. This
// will be null until the real time decision is responded to or if the real time
// decision did not set a card profile.
func (r *RealTimeDecisionDigitalWalletToken) GetCardProfileID() (CardProfileID string) {
	if r != nil && r.CardProfileID != nil {
		CardProfileID = *r.CardProfileID
	}
	return
}

func (r RealTimeDecisionDigitalWalletToken) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionDigitalWalletToken{Decision:%s CardID:%s DigitalWallet:%s CardProfileID:%s}", core.FmtP(r.Decision), core.FmtP(r.CardID), core.FmtP(r.DigitalWallet), core.FmtP(r.CardProfileID))
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
	Result *RealTimeDecisionDigitalWalletAuthenticationResult `pjson:"result"`
	// The identifier of the Card that is being tokenized.
	CardID *string `pjson:"card_id"`
	// The digital wallet app being used.
	DigitalWallet *RealTimeDecisionDigitalWalletAuthenticationDigitalWallet `pjson:"digital_wallet"`
	// The channel to send the card user their one-time passcode.
	Channel *RealTimeDecisionDigitalWalletAuthenticationChannel `pjson:"channel"`
	// The one-time passcode to send the card user.
	OneTimePasscode *string `pjson:"one_time_passcode"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone *string `pjson:"phone"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email      *string                `pjson:"email"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimeDecisionDigitalWalletAuthentication using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimeDecisionDigitalWalletAuthentication) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RealTimeDecisionDigitalWalletAuthentication into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *RealTimeDecisionDigitalWalletAuthentication) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Whether your application successfully delivered the one-time passcode.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetResult() (Result RealTimeDecisionDigitalWalletAuthenticationResult) {
	if r != nil && r.Result != nil {
		Result = *r.Result
	}
	return
}

// The identifier of the Card that is being tokenized.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The digital wallet app being used.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetDigitalWallet() (DigitalWallet RealTimeDecisionDigitalWalletAuthenticationDigitalWallet) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

// The channel to send the card user their one-time passcode.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetChannel() (Channel RealTimeDecisionDigitalWalletAuthenticationChannel) {
	if r != nil && r.Channel != nil {
		Channel = *r.Channel
	}
	return
}

// The one-time passcode to send the card user.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetOneTimePasscode() (OneTimePasscode string) {
	if r != nil && r.OneTimePasscode != nil {
		OneTimePasscode = *r.OneTimePasscode
	}
	return
}

// The phone number to send the one-time passcode to if `channel` is equal to
// `sms`.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The email to send the one-time passcode to if `channel` is equal to `email`.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}

func (r RealTimeDecisionDigitalWalletAuthentication) String() (result string) {
	return fmt.Sprintf("&RealTimeDecisionDigitalWalletAuthentication{Result:%s CardID:%s DigitalWallet:%s Channel:%s OneTimePasscode:%s Phone:%s Email:%s}", core.FmtP(r.Result), core.FmtP(r.CardID), core.FmtP(r.DigitalWallet), core.FmtP(r.Channel), core.FmtP(r.OneTimePasscode), core.FmtP(r.Phone), core.FmtP(r.Email))
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
	CardAuthorization *ActionARealTimeDecisionParametersCardAuthorization `pjson:"card_authorization"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken *ActionARealTimeDecisionParametersDigitalWalletToken `pjson:"digital_wallet_token"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication *ActionARealTimeDecisionParametersDigitalWalletAuthentication `pjson:"digital_wallet_authentication"`
	jsonFields                  map[string]interface{}                                        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ActionARealTimeDecisionParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ActionARealTimeDecisionParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ActionARealTimeDecisionParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ActionARealTimeDecisionParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If the Real-Time Decision relates to a card authorization attempt, this object
// contains your response to the authorization.
func (r *ActionARealTimeDecisionParameters) GetCardAuthorization() (CardAuthorization ActionARealTimeDecisionParametersCardAuthorization) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// If the Real-Time Decision relates to a digital wallet token provisioning
// attempt, this object contains your response to the attempt.
func (r *ActionARealTimeDecisionParameters) GetDigitalWalletToken() (DigitalWalletToken ActionARealTimeDecisionParametersDigitalWalletToken) {
	if r != nil && r.DigitalWalletToken != nil {
		DigitalWalletToken = *r.DigitalWalletToken
	}
	return
}

// If the Real-Time Decision relates to a digital wallet authentication attempt,
// this object contains your response to the authentication.
func (r *ActionARealTimeDecisionParameters) GetDigitalWalletAuthentication() (DigitalWalletAuthentication ActionARealTimeDecisionParametersDigitalWalletAuthentication) {
	if r != nil && r.DigitalWalletAuthentication != nil {
		DigitalWalletAuthentication = *r.DigitalWalletAuthentication
	}
	return
}

func (r ActionARealTimeDecisionParameters) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParameters{CardAuthorization:%s DigitalWalletToken:%s DigitalWalletAuthentication:%s}", r.CardAuthorization, r.DigitalWalletToken, r.DigitalWalletAuthentication)
}

type ActionARealTimeDecisionParametersCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision   *ActionARealTimeDecisionParametersCardAuthorizationDecision `pjson:"decision"`
	jsonFields map[string]interface{}                                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ActionARealTimeDecisionParametersCardAuthorization using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ActionARealTimeDecisionParametersCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ActionARealTimeDecisionParametersCardAuthorization into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ActionARealTimeDecisionParametersCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Whether the card authorization should be approved or declined.
func (r *ActionARealTimeDecisionParametersCardAuthorization) GetDecision() (Decision ActionARealTimeDecisionParametersCardAuthorizationDecision) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
}

func (r ActionARealTimeDecisionParametersCardAuthorization) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersCardAuthorization{Decision:%s}", core.FmtP(r.Decision))
}

type ActionARealTimeDecisionParametersCardAuthorizationDecision string

const (
	ActionARealTimeDecisionParametersCardAuthorizationDecisionApprove ActionARealTimeDecisionParametersCardAuthorizationDecision = "approve"
	ActionARealTimeDecisionParametersCardAuthorizationDecisionDecline ActionARealTimeDecisionParametersCardAuthorizationDecision = "decline"
)

type ActionARealTimeDecisionParametersDigitalWalletToken struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval *ActionARealTimeDecisionParametersDigitalWalletTokenApproval `pjson:"approval"`
	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline    *ActionARealTimeDecisionParametersDigitalWalletTokenDecline `pjson:"decline"`
	jsonFields map[string]interface{}                                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ActionARealTimeDecisionParametersDigitalWalletToken using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ActionARealTimeDecisionParametersDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ActionARealTimeDecisionParametersDigitalWalletToken into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ActionARealTimeDecisionParametersDigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If your application approves the provisioning attempt, this contains metadata
// about the digital wallet token that will be generated.
func (r *ActionARealTimeDecisionParametersDigitalWalletToken) GetApproval() (Approval ActionARealTimeDecisionParametersDigitalWalletTokenApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your application declines the provisioning attempt, this contains details
// about the decline.
func (r *ActionARealTimeDecisionParametersDigitalWalletToken) GetDecline() (Decline ActionARealTimeDecisionParametersDigitalWalletTokenDecline) {
	if r != nil && r.Decline != nil {
		Decline = *r.Decline
	}
	return
}

func (r ActionARealTimeDecisionParametersDigitalWalletToken) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletToken{Approval:%s Decline:%s}", r.Approval, r.Decline)
}

type ActionARealTimeDecisionParametersDigitalWalletTokenApproval struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID *string `pjson:"card_profile_id"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `pjson:"phone"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email      *string                `pjson:"email"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ActionARealTimeDecisionParametersDigitalWalletTokenApproval using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ActionARealTimeDecisionParametersDigitalWalletTokenApproval into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenApproval) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Card Profile to assign to the Digital Wallet token.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenApproval) GetCardProfileID() (CardProfileID string) {
	if r != nil && r.CardProfileID != nil {
		CardProfileID = *r.CardProfileID
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenApproval) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// An email address that can be used to verify the cardholder via one-time
// passcode.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenApproval) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}

func (r ActionARealTimeDecisionParametersDigitalWalletTokenApproval) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletTokenApproval{CardProfileID:%s Phone:%s Email:%s}", core.FmtP(r.CardProfileID), core.FmtP(r.Phone), core.FmtP(r.Email))
}

type ActionARealTimeDecisionParametersDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason     *string                `pjson:"reason"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ActionARealTimeDecisionParametersDigitalWalletTokenDecline using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ActionARealTimeDecisionParametersDigitalWalletTokenDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Why the tokenization attempt was declined. This is for logging purposes only and
// is not displayed to the end-user.
func (r *ActionARealTimeDecisionParametersDigitalWalletTokenDecline) GetReason() (Reason string) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r ActionARealTimeDecisionParametersDigitalWalletTokenDecline) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletTokenDecline{Reason:%s}", core.FmtP(r.Reason))
}

type ActionARealTimeDecisionParametersDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result     *ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult `pjson:"result"`
	jsonFields map[string]interface{}                                              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ActionARealTimeDecisionParametersDigitalWalletAuthentication using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ActionARealTimeDecisionParametersDigitalWalletAuthentication) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ActionARealTimeDecisionParametersDigitalWalletAuthentication into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ActionARealTimeDecisionParametersDigitalWalletAuthentication) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Whether your application was able to deliver the one-time passcode.
func (r *ActionARealTimeDecisionParametersDigitalWalletAuthentication) GetResult() (Result ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult) {
	if r != nil && r.Result != nil {
		Result = *r.Result
	}
	return
}

func (r ActionARealTimeDecisionParametersDigitalWalletAuthentication) String() (result string) {
	return fmt.Sprintf("&ActionARealTimeDecisionParametersDigitalWalletAuthentication{Result:%s}", core.FmtP(r.Result))
}

type ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult string

const (
	ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultSuccess ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult = "success"
	ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultFailure ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult = "failure"
)
