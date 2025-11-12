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

func TestInboundMailItemGet(t *testing.T) {
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
	_, err := client.InboundMailItems.Get(context.TODO(), "inbound_mail_item_q6rrg7mmqpplx80zceev")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundMailItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.InboundMailItems.List(context.TODO(), increase.InboundMailItemListParams{
		CreatedAt: increase.F(increase.InboundMailItemListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:    increase.F("cursor"),
		Limit:     increase.F(int64(1)),
		LockboxID: increase.F("lockbox_id"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundMailItemAction(t *testing.T) {
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
	_, err := client.InboundMailItems.Action(
		context.TODO(),
		"inbound_mail_item_q6rrg7mmqpplx80zceev",
		increase.InboundMailItemActionParams{
			Checks: increase.F([]increase.InboundMailItemActionParamsCheck{{
				Action:  increase.F(increase.InboundMailItemActionParamsChecksActionDeposit),
				Account: increase.F("account_in71c4amph0vgo2qllky"),
			}, {
				Action:  increase.F(increase.InboundMailItemActionParamsChecksActionIgnore),
				Account: increase.F("account"),
			}}),
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
