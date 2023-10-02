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

func TestCheckTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckTransfers.New(context.TODO(), increase.CheckTransferNewParams{
		AccountID:         increase.F("string"),
		Amount:            increase.F(int64(1)),
		FulfillmentMethod: increase.F(increase.CheckTransferNewParamsFulfillmentMethodPhysicalCheck),
		PhysicalCheck: increase.F(increase.CheckTransferNewParamsPhysicalCheck{
			Memo:          increase.F("x"),
			Note:          increase.F("x"),
			RecipientName: increase.F("x"),
			MailingAddress: increase.F(increase.CheckTransferNewParamsPhysicalCheckMailingAddress{
				Name:       increase.F("x"),
				Line1:      increase.F("x"),
				Line2:      increase.F("x"),
				City:       increase.F("x"),
				State:      increase.F("x"),
				PostalCode: increase.F("x"),
			}),
			ReturnAddress: increase.F(increase.CheckTransferNewParamsPhysicalCheckReturnAddress{
				Name:       increase.F("x"),
				Line1:      increase.F("x"),
				Line2:      increase.F("x"),
				City:       increase.F("x"),
				State:      increase.F("x"),
				PostalCode: increase.F("x"),
			}),
		}),
		RequireApproval:       increase.F(true),
		SourceAccountNumberID: increase.F("string"),
		UniqueIdentifier:      increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferGet(t *testing.T) {
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
	_, err := client.CheckTransfers.Get(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckTransfers.List(context.TODO(), increase.CheckTransferListParams{
		AccountID: increase.F("string"),
		CreatedAt: increase.F(increase.CheckTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:           increase.F("string"),
		Limit:            increase.F(int64(0)),
		UniqueIdentifier: increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferApprove(t *testing.T) {
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
	_, err := client.CheckTransfers.Approve(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferCancel(t *testing.T) {
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
	_, err := client.CheckTransfers.Cancel(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferStopPaymentWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	t.Skip("Prism doesn't accept no request body being sent but returns 415 if it is sent")
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.CheckTransfers.StopPayment(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
		increase.CheckTransferStopPaymentParams{
			Reason: increase.F(increase.CheckTransferStopPaymentParamsReasonMailDeliveryFailed),
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
