package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type RealTimeDecisionActionParams struct {
	// If the Real-Time Decision relates to a card authorization attempt, this object
	// contains your response to the authorization.
	CardAuthorization field.Field[RealTimeDecisionActionParamsCardAuthorization] `json:"card_authorization"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken field.Field[RealTimeDecisionActionParamsDigitalWalletToken] `json:"digital_wallet_token"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication field.Field[RealTimeDecisionActionParamsDigitalWalletAuthentication] `json:"digital_wallet_authentication"`
}

// MarshalJSON serializes RealTimeDecisionActionParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RealTimeDecisionActionParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type RealTimeDecisionActionParamsCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision field.Field[RealTimeDecisionActionParamsCardAuthorizationDecision] `json:"decision,required"`
}

type RealTimeDecisionActionParamsCardAuthorizationDecision string

const (
	RealTimeDecisionActionParamsCardAuthorizationDecisionApprove RealTimeDecisionActionParamsCardAuthorizationDecision = "approve"
	RealTimeDecisionActionParamsCardAuthorizationDecisionDecline RealTimeDecisionActionParamsCardAuthorizationDecision = "decline"
)

type RealTimeDecisionActionParamsDigitalWalletToken struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval field.Field[RealTimeDecisionActionParamsDigitalWalletTokenApproval] `json:"approval"`
	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline field.Field[RealTimeDecisionActionParamsDigitalWalletTokenDecline] `json:"decline"`
}

type RealTimeDecisionActionParamsDigitalWalletTokenApproval struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID field.Field[string] `json:"card_profile_id,required"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone field.Field[string] `json:"phone"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email field.Field[string] `json:"email"`
}

type RealTimeDecisionActionParamsDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason field.Field[string] `json:"reason"`
}

type RealTimeDecisionActionParamsDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result field.Field[RealTimeDecisionActionParamsDigitalWalletAuthenticationResult] `json:"result,required"`
}

type RealTimeDecisionActionParamsDigitalWalletAuthenticationResult string

const (
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "success"
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "failure"
)
