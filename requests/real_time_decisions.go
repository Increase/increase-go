package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
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
