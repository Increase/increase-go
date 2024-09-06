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

func TestRealTimePaymentsRequestForPaymentNew(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestForPayments.New(context.TODO(), increase.RealTimePaymentsRequestForPaymentNewParams{
		Amount: increase.F(int64(100)),
		Debtor: increase.F(increase.RealTimePaymentsRequestForPaymentNewParamsDebtor{
			Address: increase.F(increase.RealTimePaymentsRequestForPaymentNewParamsDebtorAddress{
				Country:    increase.F("US"),
				City:       increase.F("x"),
				PostCode:   increase.F("x"),
				StreetName: increase.F("Liberty Street"),
			}),
			Name: increase.F("Ian Crease"),
		}),
		DestinationAccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		ExpiresAt:                  increase.F(time.Now()),
		RemittanceInformation:      increase.F("Invoice 29582"),
		SourceAccountNumber:        increase.F("987654321"),
		SourceRoutingNumber:        increase.F("101050001"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsRequestForPaymentGet(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestForPayments.Get(context.TODO(), "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsRequestForPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestForPayments.List(context.TODO(), increase.RealTimePaymentsRequestForPaymentListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.RealTimePaymentsRequestForPaymentListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
