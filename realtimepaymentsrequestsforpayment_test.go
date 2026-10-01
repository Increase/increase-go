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

func TestRealTimePaymentsRequestsForPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestsForPayment.New(context.TODO(), increase.RealTimePaymentsRequestsForPaymentNewParams{
		AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:          increase.F(int64(100)),
		Debtor: increase.F(increase.RealTimePaymentsRequestsForPaymentNewParamsDebtor{
			Address: increase.F(increase.RealTimePaymentsRequestsForPaymentNewParamsDebtorAddress{
				Country:        increase.F("US"),
				AddressLine2:   increase.F("x"),
				BuildingNumber: increase.F("x"),
				City:           increase.F("x"),
				PostalCode:     increase.F("x"),
				State:          increase.F("xx"),
				StreetName:     increase.F("Liberty Street"),
			}),
			Name: increase.F("Ian Crease"),
		}),
		DebtorAccountNumber:               increase.F("987654321"),
		DebtorRoutingNumber:               increase.F("101050001"),
		ExpiresAt:                         increase.F(time.Now()),
		RequestedExecutionAt:              increase.F(time.Now()),
		UnstructuredRemittanceInformation: increase.F("Invoice 29582"),
		CreditorName:                      increase.F("National Phonograph Company"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsRequestsForPaymentGet(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestsForPayment.Get(context.TODO(), "real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsRequestsForPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestsForPayment.List(context.TODO(), increase.RealTimePaymentsRequestsForPaymentListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.RealTimePaymentsRequestsForPaymentListParamsCreatedAt{
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

func TestRealTimePaymentsRequestsForPaymentCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestsForPayment.Cancel(
		context.TODO(),
		"real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7",
		increase.RealTimePaymentsRequestsForPaymentCancelParams{
			AdditionalInformation: increase.F("x"),
			Reason:                increase.F(increase.RealTimePaymentsRequestsForPaymentCancelParamsReasonRequestedByCustomer),
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
