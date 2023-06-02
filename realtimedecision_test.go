package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestRealTimeDecisionGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Get(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimeDecisionActionWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		increase.RealTimeDecisionActionParams{
			CardAuthorization:           increase.F(increase.RealTimeDecisionActionParamsCardAuthorization{Decision: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationDecisionApprove)}),
			DigitalWalletAuthentication: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthentication{Result: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess)}),
			DigitalWalletToken:          increase.F(increase.RealTimeDecisionActionParamsDigitalWalletToken{Approval: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletTokenApproval{CardProfileID: increase.F("string"), Phone: increase.F("x"), Email: increase.F("x")}), Decline: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletTokenDecline{Reason: increase.F("x")})}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
