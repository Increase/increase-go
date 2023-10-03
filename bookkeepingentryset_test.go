// File generated from our OpenAPI spec by Stainless.

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

func TestBookkeepingEntrySetNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.BookkeepingEntrySets.New(context.TODO(), increase.BookkeepingEntrySetNewParams{
		Entries: increase.F([]increase.BookkeepingEntrySetNewParamsEntry{{
			AccountID: increase.F("bookkeeping_account_9husfpw68pzmve9dvvc7"),
			Amount:    increase.F(int64(100)),
		}, {
			AccountID: increase.F("bookkeeping_account_t2obldz1rcu15zr54umg"),
			Amount:    increase.F(int64(-100)),
		}}),
		Date:          increase.F(time.Now()),
		TransactionID: increase.F("transaction_uyrp7fld2ium70oa7oi"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
