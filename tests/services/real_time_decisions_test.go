package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestRealTimeDecisionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Get(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestRealTimeDecisionsActionWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		&types.ActionARealTimeDecisionParameters{CardAuthorization: increase.P(types.ActionARealTimeDecisionParametersCardAuthorization{Decision: increase.P(types.ActionARealTimeDecisionParametersCardAuthorizationDecisionApprove)}), DigitalWalletToken: increase.P(types.ActionARealTimeDecisionParametersDigitalWalletToken{Approval: increase.P(types.ActionARealTimeDecisionParametersDigitalWalletTokenApproval{CardProfileID: increase.P("string"), Phone: increase.P("x"), Email: increase.P("x")}), Decline: increase.P(types.ActionARealTimeDecisionParametersDigitalWalletTokenDecline{Reason: increase.P("x")})}), DigitalWalletAuthentication: increase.P(types.ActionARealTimeDecisionParametersDigitalWalletAuthentication{Result: increase.P(types.ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultSuccess)})},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
