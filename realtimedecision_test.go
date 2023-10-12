// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestRealTimeDecisionGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.RealTimeDecisions.Get(context.TODO(), "real_time_decision_j76n2e810ezcg3zh5qtn")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimeDecisionActionWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		increase.RealTimeDecisionActionParams{
			CardAuthorization: increase.F(increase.RealTimeDecisionActionParamsCardAuthorization{
				Decision: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationDecisionApprove),
			}),
			DigitalWalletAuthentication: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthentication{
				Result: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess),
			}),
			DigitalWalletToken: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletToken{
				Approval: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletTokenApproval{
					CardProfileID: increase.F("string"),
					Phone:         increase.F("x"),
					Email:         increase.F("x"),
				}),
				Decline: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletTokenDecline{
					Reason: increase.F("x"),
				}),
			}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
