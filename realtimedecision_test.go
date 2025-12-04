// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
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
			CardAuthentication: increase.F(increase.RealTimeDecisionActionParamsCardAuthentication{
				Decision: increase.F(increase.RealTimeDecisionActionParamsCardAuthenticationDecisionApprove),
			}),
			CardAuthenticationChallenge: increase.F(increase.RealTimeDecisionActionParamsCardAuthenticationChallenge{
				Result: increase.F(increase.RealTimeDecisionActionParamsCardAuthenticationChallengeResultSuccess),
			}),
			CardAuthorization: increase.F(increase.RealTimeDecisionActionParamsCardAuthorization{
				Decision: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationDecisionApprove),
				Approval: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationApproval{
					CardholderAddressVerificationResult: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResult{
						Line1:      increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1Match),
						PostalCode: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCodeNoMatch),
					}),
					PartialAmount: increase.F(int64(1)),
				}),
				Decline: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationDecline{
					Reason: increase.F(increase.RealTimeDecisionActionParamsCardAuthorizationDeclineReasonInsufficientFunds),
				}),
			}),
			DigitalWalletAuthentication: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthentication{
				Result: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess),
				Success: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletAuthenticationSuccess{
					Email: increase.F("dev@stainless.com"),
					Phone: increase.F("x"),
				}),
			}),
			DigitalWalletToken: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletToken{
				Approval: increase.F(increase.RealTimeDecisionActionParamsDigitalWalletTokenApproval{
					Email: increase.F("dev@stainless.com"),
					Phone: increase.F("x"),
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
