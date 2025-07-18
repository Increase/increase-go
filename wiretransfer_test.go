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

func TestWireTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WireTransfers.New(context.TODO(), increase.WireTransferNewParams{
		AccountID:                    increase.F("account_in71c4amph0vgo2qllky"),
		Amount:                       increase.F(int64(100)),
		BeneficiaryName:              increase.F("Ian Crease"),
		MessageToRecipient:           increase.F("New account transfer"),
		AccountNumber:                increase.F("987654321"),
		BeneficiaryAddressLine1:      increase.F("33 Liberty Street"),
		BeneficiaryAddressLine2:      increase.F("New York"),
		BeneficiaryAddressLine3:      increase.F("NY 10045"),
		ExternalAccountID:            increase.F("external_account_id"),
		InboundWireDrawdownRequestID: increase.F("inbound_wire_drawdown_request_id"),
		OriginatorAddressLine1:       increase.F("x"),
		OriginatorAddressLine2:       increase.F("x"),
		OriginatorAddressLine3:       increase.F("x"),
		OriginatorName:               increase.F("x"),
		RequireApproval:              increase.F(true),
		RoutingNumber:                increase.F("101050001"),
		SourceAccountNumberID:        increase.F("source_account_number_id"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferGet(t *testing.T) {
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
	_, err := client.WireTransfers.Get(context.TODO(), "wire_transfer_5akynk7dqsq25qwk9q2u")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.WireTransfers.List(context.TODO(), increase.WireTransferListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.WireTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:            increase.F("cursor"),
		ExternalAccountID: increase.F("external_account_id"),
		IdempotencyKey:    increase.F("x"),
		Limit:             increase.F(int64(1)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferApprove(t *testing.T) {
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
	_, err := client.WireTransfers.Approve(context.TODO(), "wire_transfer_5akynk7dqsq25qwk9q2u")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferCancel(t *testing.T) {
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
	_, err := client.WireTransfers.Cancel(context.TODO(), "wire_transfer_5akynk7dqsq25qwk9q2u")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
