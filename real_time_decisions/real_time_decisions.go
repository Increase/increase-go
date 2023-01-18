package real_time_decisions

import "context"
import "increase/core"
import "fmt"

type RealTimeDecisionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewRealTimeDecisionService(requester core.Requester) (r *RealTimeDecisionService) {
	r = &RealTimeDecisionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt string `json:"created_at"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt string `json:"timeout_at"`
	// The status of the Real-Time Decision.
	Status RealTimeDecisionStatus `json:"status"`
	// The category of the Real-Time Decision.
	Category RealTimeDecisionCategory `json:"category"`
	// Fields related to a card authorization.
	CardAuthorization *RealTimeDecisionCardAuthorization `json:"card_authorization"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken *RealTimeDecisionDigitalWalletToken `json:"digital_wallet_token"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication *RealTimeDecisionDigitalWalletAuthentication `json:"digital_wallet_authentication"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type RealTimeDecisionType `json:"type"`
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

//
type RealTimeDecisionCardAuthorization struct {
	// Whether or not the authorization was approved.
	Decision *RealTimeDecisionCardAuthorizationDecision `json:"decision"`
	// The identifier of the Card that is being authorized.
	CardID string `json:"card_id"`
	// The identifier of the Account the authorization will debit.
	AccountID string `json:"account_id"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount int `json:"presentment_amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency string `json:"presentment_currency"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount int `json:"settlement_amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency string `json:"settlement_currency"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id"`
	// The city the merchant resides in.
	MerchantCity *string `json:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country"`
}

// Whether or not the authorization was approved.
func (r *RealTimeDecisionCardAuthorization) GetDecision() (Decision RealTimeDecisionCardAuthorizationDecision) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
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

type RealTimeDecisionCardAuthorizationDecision string

const (
	RealTimeDecisionCardAuthorizationDecisionApprove RealTimeDecisionCardAuthorizationDecision = "approve"
	RealTimeDecisionCardAuthorizationDecisionDecline RealTimeDecisionCardAuthorizationDecision = "decline"
)

//
type RealTimeDecisionDigitalWalletToken struct {
	// Whether or not the provisioning request was approved.
	Decision *RealTimeDecisionDigitalWalletTokenDecision `json:"decision"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet `json:"digital_wallet"`
}

// Whether or not the provisioning request was approved.
func (r *RealTimeDecisionDigitalWalletToken) GetDecision() (Decision RealTimeDecisionDigitalWalletTokenDecision) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
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

//
type RealTimeDecisionDigitalWalletAuthentication struct {
	// Whether your application successfully delivered the one-time passcode.
	Result *RealTimeDecisionDigitalWalletAuthenticationResult `json:"result"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletAuthenticationDigitalWallet `json:"digital_wallet"`
	// The channel to send the card user their one-time passcode.
	Channel RealTimeDecisionDigitalWalletAuthenticationChannel `json:"channel"`
	// The one-time passcode to send the card user.
	OneTimePasscode string `json:"one_time_passcode"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone *string `json:"phone"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email *string `json:"email"`
}

// Whether your application successfully delivered the one-time passcode.
func (r *RealTimeDecisionDigitalWalletAuthentication) GetResult() (Result RealTimeDecisionDigitalWalletAuthenticationResult) {
	if r != nil && r.Result != nil {
		Result = *r.Result
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
	CardAuthorization ActionARealTimeDecisionParametersCardAuthorization `json:"card_authorization,omitempty"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken ActionARealTimeDecisionParametersDigitalWalletToken `json:"digital_wallet_token,omitempty"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication ActionARealTimeDecisionParametersDigitalWalletAuthentication `json:"digital_wallet_authentication,omitempty"`
}

//
type ActionARealTimeDecisionParametersCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision ActionARealTimeDecisionParametersCardAuthorizationDecision `json:"decision"`
}

type ActionARealTimeDecisionParametersCardAuthorizationDecision string

const (
	ActionARealTimeDecisionParametersCardAuthorizationDecisionApprove ActionARealTimeDecisionParametersCardAuthorizationDecision = "approve"
	ActionARealTimeDecisionParametersCardAuthorizationDecisionDecline ActionARealTimeDecisionParametersCardAuthorizationDecision = "decline"
)

//
type ActionARealTimeDecisionParametersDigitalWalletToken struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval ActionARealTimeDecisionParametersDigitalWalletTokenApproval `json:"approval,omitempty"`
	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline ActionARealTimeDecisionParametersDigitalWalletTokenDecline `json:"decline,omitempty"`
}

//
type ActionARealTimeDecisionParametersDigitalWalletTokenApproval struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID string `json:"card_profile_id"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string `json:"phone,omitempty"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email string `json:"email,omitempty"`
}

//
type ActionARealTimeDecisionParametersDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason string `json:"reason,omitempty"`
}

//
type ActionARealTimeDecisionParametersDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult `json:"result"`
}

type ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult string

const (
	ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultSuccess ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult = "success"
	ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultFailure ActionARealTimeDecisionParametersDigitalWalletAuthenticationResult = "failure"
)

func (r *RealTimeDecisionService) Retrieve(ctx context.Context, real_time_decision_id string, opts ...*core.RequestOpts) (res *RealTimeDecision, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/real_time_decisions/%s", real_time_decision_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *RealTimeDecisionService) Action(ctx context.Context, real_time_decision_id string, body *ActionARealTimeDecisionParameters, opts ...*core.RequestOpts) (res *RealTimeDecision, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/real_time_decisions/%s/action", real_time_decision_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}
