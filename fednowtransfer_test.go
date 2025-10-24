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

func TestFednowTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.FednowTransfers.New(context.TODO(), increase.FednowTransferNewParams{
		AccountID:                         increase.F("account_in71c4amph0vgo2qllky"),
		Amount:                            increase.F(int64(100)),
		CreditorName:                      increase.F("Ian Crease"),
		DebtorName:                        increase.F("National Phonograph Company"),
		SourceAccountNumberID:             increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		UnstructuredRemittanceInformation: increase.F("Invoice 29582"),
		AccountNumber:                     increase.F("987654321"),
		CreditorAddress: increase.F(increase.FednowTransferNewParamsCreditorAddress{
			City:       increase.F("New York"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
			Line1:      increase.F("33 Liberty Street"),
		}),
		DebtorAddress: increase.F(increase.FednowTransferNewParamsDebtorAddress{
			City:       increase.F("x"),
			PostalCode: increase.F("x"),
			State:      increase.F("x"),
			Line1:      increase.F("x"),
		}),
		ExternalAccountID: increase.F("external_account_id"),
		RequireApproval:   increase.F(true),
		RoutingNumber:     increase.F("101050001"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFednowTransferGet(t *testing.T) {
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
	_, err := client.FednowTransfers.Get(context.TODO(), "fednow_transfer_4i0mptrdu1mueg1196bg")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFednowTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.FednowTransfers.List(context.TODO(), increase.FednowTransferListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.FednowTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:            increase.F("cursor"),
		ExternalAccountID: increase.F("external_account_id"),
		IdempotencyKey:    increase.F("x"),
		Limit:             increase.F(int64(1)),
		Status: increase.F(increase.FednowTransferListParamsStatus{
			In: increase.F([]increase.FednowTransferListParamsStatusIn{increase.FednowTransferListParamsStatusInPendingReviewing}),
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

func TestFednowTransferApprove(t *testing.T) {
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
	_, err := client.FednowTransfers.Approve(context.TODO(), "fednow_transfer_4i0mptrdu1mueg1196bg")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFednowTransferCancel(t *testing.T) {
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
	_, err := client.FednowTransfers.Cancel(context.TODO(), "fednow_transfer_4i0mptrdu1mueg1196bg")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
