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

func TestCheckDepositNew(t *testing.T) {
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
	_, err := client.CheckDeposits.New(context.TODO(), increase.CheckDepositNewParams{
		AccountID:        increase.F("account_in71c4amph0vgo2qllky"),
		Amount:           increase.F(int64(1000)),
		BackImageFileID:  increase.F("file_26khfk98mzfz90a11oqx"),
		Currency:         increase.F("USD"),
		FrontImageFileID: increase.F("file_hkv175ovmc2tb2v2zbrm"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckDepositGet(t *testing.T) {
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
	_, err := client.CheckDeposits.Get(context.TODO(), "check_deposit_instruction_q2shv7x9qhevfm71kor8")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckDepositListWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckDeposits.List(context.TODO(), increase.CheckDepositListParams{
		AccountID: increase.F("string"),
		CreatedAt: increase.F(increase.CheckDepositListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(1)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
