// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestProofOfAuthorizationRequestSubmissionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ProofOfAuthorizationRequestSubmissions.New(context.TODO(), increase.ProofOfAuthorizationRequestSubmissionNewParams{
		AuthorizationTerms:                            increase.F("I agree to the terms of service."),
		AuthorizedAt:                                  increase.F(time.Now()),
		AuthorizerEmail:                               increase.F("user@example.com"),
		AuthorizerName:                                increase.F("Ian Crease"),
		CustomerHasBeenOffboarded:                     increase.F(true),
		ProofOfAuthorizationRequestID:                 increase.F("proof_of_authorization_request_iwp8no25h3rjvil6ad3b"),
		ValidatedAccountOwnershipViaCredential:        increase.F(true),
		ValidatedAccountOwnershipWithAccountStatement: increase.F(true),
		ValidatedAccountOwnershipWithMicrodeposit:     increase.F(true),
		AuthorizerCompany:                             increase.F("National Phonograph Company"),
		AuthorizerIPAddress:                           increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProofOfAuthorizationRequestSubmissionGet(t *testing.T) {
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
	_, err := client.ProofOfAuthorizationRequestSubmissions.Get(context.TODO(), "proof_of_authorization_request_submission_uqhqroiley7n0097vizn")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProofOfAuthorizationRequestSubmissionListWithOptionalParams(t *testing.T) {
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
	_, err := client.ProofOfAuthorizationRequestSubmissions.List(context.TODO(), increase.ProofOfAuthorizationRequestSubmissionListParams{
		Cursor:                        increase.F("cursor"),
		IdempotencyKey:                increase.F("x"),
		Limit:                         increase.F(int64(1)),
		ProofOfAuthorizationRequestID: increase.F("proof_of_authorization_request_id"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
