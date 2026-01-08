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

func TestSwiftTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SwiftTransfers.New(context.TODO(), increase.SwiftTransferNewParams{
		AccountID:              increase.F("account_in71c4amph0vgo2qllky"),
		AccountNumber:          increase.F("987654321"),
		BankIdentificationCode: increase.F("ECBFDEFFTPP"),
		CreditorAddress: increase.F(increase.SwiftTransferNewParamsCreditorAddress{
			City:       increase.F("Frankfurt"),
			Country:    increase.F("DE"),
			Line1:      increase.F("Sonnemannstrasse 20"),
			Line2:      increase.F("x"),
			PostalCode: increase.F("60314"),
			State:      increase.F("x"),
		}),
		CreditorName: increase.F("Ian Crease"),
		DebtorAddress: increase.F(increase.SwiftTransferNewParamsDebtorAddress{
			City:       increase.F("New York"),
			Country:    increase.F("US"),
			Line1:      increase.F("33 Liberty Street"),
			Line2:      increase.F("x"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
		}),
		DebtorName:                        increase.F("National Phonograph Company"),
		InstructedAmount:                  increase.F(int64(100)),
		InstructedCurrency:                increase.F(increase.SwiftTransferNewParamsInstructedCurrencyUsd),
		SourceAccountNumberID:             increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		UnstructuredRemittanceInformation: increase.F("New Swift transfer"),
		RequireApproval:                   increase.F(true),
		RoutingNumber:                     increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSwiftTransferGet(t *testing.T) {
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
	_, err := client.SwiftTransfers.Get(context.TODO(), "swift_transfer_29h21xkng03788zwd3fh")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSwiftTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.SwiftTransfers.List(context.TODO(), increase.SwiftTransferListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.SwiftTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.SwiftTransferListParamsStatus{
			In: increase.F([]increase.SwiftTransferListParamsStatusIn{increase.SwiftTransferListParamsStatusInPendingApproval}),
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

func TestSwiftTransferApprove(t *testing.T) {
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
	_, err := client.SwiftTransfers.Approve(context.TODO(), "swift_transfer_29h21xkng03788zwd3fh")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSwiftTransferCancel(t *testing.T) {
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
	_, err := client.SwiftTransfers.Cancel(context.TODO(), "swift_transfer_29h21xkng03788zwd3fh")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
