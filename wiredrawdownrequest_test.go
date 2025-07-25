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

func TestWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WireDrawdownRequests.New(context.TODO(), increase.WireDrawdownRequestNewParams{
		AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:          increase.F(int64(10000)),
		CreditorAddress: increase.F(increase.WireDrawdownRequestNewParamsCreditorAddress{
			City:       increase.F("New York"),
			Country:    increase.F("US"),
			Line1:      increase.F("33 Liberty Street"),
			Line2:      increase.F("x"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
		}),
		CreditorName:        increase.F("National Phonograph Company"),
		DebtorAccountNumber: increase.F("987654321"),
		DebtorAddress: increase.F(increase.WireDrawdownRequestNewParamsDebtorAddress{
			City:       increase.F("New York"),
			Country:    increase.F("US"),
			Line1:      increase.F("33 Liberty Street"),
			Line2:      increase.F("x"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
		}),
		DebtorName:                        increase.F("Ian Crease"),
		DebtorRoutingNumber:               increase.F("101050001"),
		UnstructuredRemittanceInformation: increase.F("Invoice 29582"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestGet(t *testing.T) {
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
	_, err := client.WireDrawdownRequests.Get(context.TODO(), "wire_drawdown_request_q6lmocus3glo0lr2bfv3")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.WireDrawdownRequests.List(context.TODO(), increase.WireDrawdownRequestListParams{
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.WireDrawdownRequestListParamsStatus{
			In: increase.F([]increase.WireDrawdownRequestListParamsStatusIn{increase.WireDrawdownRequestListParamsStatusInPendingSubmission}),
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
