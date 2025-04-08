// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestAccountNumberNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountNumbers.New(context.TODO(), increase.AccountNumberNewParams{
		AccountID: increase.F("account_in71c4amph0vgo2qllky"),
		Name:      increase.F("Rent payments"),
		InboundACH: increase.F(increase.AccountNumberNewParamsInboundACH{
			DebitStatus: increase.F(increase.AccountNumberNewParamsInboundACHDebitStatusAllowed),
		}),
		InboundChecks: increase.F(increase.AccountNumberNewParamsInboundChecks{
			Status: increase.F(increase.AccountNumberNewParamsInboundChecksStatusAllowed),
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

func TestAccountNumberGet(t *testing.T) {
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
	_, err := client.AccountNumbers.Get(context.TODO(), "account_number_v18nkfqm6afpsrvy82b2")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountNumberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountNumbers.Update(
		context.TODO(),
		"account_number_v18nkfqm6afpsrvy82b2",
		increase.AccountNumberUpdateParams{
			InboundACH: increase.F(increase.AccountNumberUpdateParamsInboundACH{
				DebitStatus: increase.F(increase.AccountNumberUpdateParamsInboundACHDebitStatusBlocked),
			}),
			InboundChecks: increase.F(increase.AccountNumberUpdateParamsInboundChecks{
				Status: increase.F(increase.AccountNumberUpdateParamsInboundChecksStatusAllowed),
			}),
			Name:   increase.F("x"),
			Status: increase.F(increase.AccountNumberUpdateParamsStatusDisabled),
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

func TestAccountNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountNumbers.List(context.TODO(), increase.AccountNumberListParams{
		AccountID: increase.F("account_id"),
		ACHDebitStatus: increase.F(increase.AccountNumberListParamsACHDebitStatus{
			In: increase.F([]increase.AccountNumberListParamsACHDebitStatusIn{increase.AccountNumberListParamsACHDebitStatusInAllowed}),
		}),
		CreatedAt: increase.F(increase.AccountNumberListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.AccountNumberListParamsStatus{
			In: increase.F([]increase.AccountNumberListParamsStatusIn{increase.AccountNumberListParamsStatusInActive}),
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
