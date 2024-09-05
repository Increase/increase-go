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

func TestOAuthTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OAuthTokens.New(context.TODO(), increase.OAuthTokenNewParams{
		GrantType:       increase.F(increase.OAuthTokenNewParamsGrantTypeAuthorizationCode),
		ClientID:        increase.F("12345"),
		ClientSecret:    increase.F("supersecret"),
		Code:            increase.F("123"),
		ProductionToken: increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
