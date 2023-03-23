package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestRealTimeDecisionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Get(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimeDecisionsActionWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		&requests.ActionARealTimeDecisionParameters{CardAuthorization: fields.F(requests.ActionARealTimeDecisionParametersCardAuthorization{Decision: fields.F(requests.ActionARealTimeDecisionParametersCardAuthorizationDecisionApprove)}), DigitalWalletToken: fields.F(requests.ActionARealTimeDecisionParametersDigitalWalletToken{Approval: fields.F(requests.ActionARealTimeDecisionParametersDigitalWalletTokenApproval{CardProfileID: fields.F("string"), Phone: fields.F("x"), Email: fields.F("x")}), Decline: fields.F(requests.ActionARealTimeDecisionParametersDigitalWalletTokenDecline{Reason: fields.F("x")})}), DigitalWalletAuthentication: fields.F(requests.ActionARealTimeDecisionParametersDigitalWalletAuthentication{Result: fields.F(requests.ActionARealTimeDecisionParametersDigitalWalletAuthenticationResultSuccess)})},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
