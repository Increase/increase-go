package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestRealTimeDecisionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Get(context.TODO(), "real_time_decision_j76n2e810ezcg3zh5qtn")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestRealTimeDecisionsActionWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		&types.ActionARealTimeDecisionParameters{CardAuthorization: &types.ActionARealTimeDecisionParametersCardAuthorization{Decision: increase.P(types.ActionARealTimeDecisionParametersCardAuthorizationDecisionApprove)}, DigitalWalletToken: &types.ActionARealTimeDecisionParametersDigitalWalletToken{Approval: &types.ActionARealTimeDecisionParametersDigitalWalletTokenApproval{CardProfileID: increase.P("string"), Phone: increase.P("x"), Email: increase.P("x")}, Decline: &types.ActionARealTimeDecisionParametersDigitalWalletTokenDecline{Reason: increase.P("x")}}, DigitalWalletAuthentication: &types.ActionARealTimeDecisionParametersDigitalWalletAuthentication{Result: increase.P(types.ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultSuccess)}},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
