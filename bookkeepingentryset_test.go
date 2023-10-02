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
			AccountID: increase.F("string"),
			Amount:    increase.F(int64(0)),
		}, {
			AccountID: increase.F("string"),
			Amount:    increase.F(int64(0)),
		}, {
			AccountID: increase.F("string"),
			Amount:    increase.F(int64(0)),
		}}),
		Date:          increase.F(time.Now()),
		TransactionID: increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
