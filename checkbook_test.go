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

func TestCheckbookNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Checkbooks.New(context.TODO(), increase.CheckbookNewParams{
		AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		MailingAddress: increase.F(increase.CheckbookNewParamsMailingAddress{
			City:       increase.F("New York"),
			Line1:      increase.F("33 Liberty Street"),
			Name:       increase.F("Ian Crease"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
			Line2:      increase.F("line2"),
		}),
		Payer: increase.F([]increase.CheckbookNewParamsPayer{{
			Contents: increase.F("National Phonograph Company"),
		}, {
			Contents: increase.F("33 Liberty Street"),
		}, {
			Contents: increase.F("New York, NY 10045"),
		}}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckbookGet(t *testing.T) {
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
	_, err := client.Checkbooks.Get(context.TODO(), "checkbook_7iudxhxum6vlkcjxmi8y")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckbookListWithOptionalParams(t *testing.T) {
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
	_, err := client.Checkbooks.List(context.TODO(), increase.CheckbookListParams{
		AccountNumberID: increase.F("account_number_id"),
		Cursor:          increase.F("cursor"),
		IdempotencyKey:  increase.F("x"),
		Limit:           increase.F(int64(1)),
		Status: increase.F(increase.CheckbookListParamsStatus{
			In: increase.F([]increase.CheckbookListParamsStatusIn{increase.CheckbookListParamsStatusInPendingSubmission}),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckbookStop(t *testing.T) {
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
	_, err := client.Checkbooks.Stop(context.TODO(), "checkbook_7iudxhxum6vlkcjxmi8y")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
