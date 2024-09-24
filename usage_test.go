// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"os"
	"testing"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestUsage(t *testing.T) {
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
	account, err := client.Accounts.New(context.TODO(), increase.AccountNewParams{
		Name:      increase.F("New Account!"),
		EntityID:  increase.F("entity_n8y8tnk2p9339ti393yi"),
		ProgramID: increase.F("program_i2v2os4mwza1oetokh8i"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", account.ID)
}
