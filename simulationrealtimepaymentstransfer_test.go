// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestSimulationRealTimePaymentsTransferCompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.RealTimePaymentsTransfers.Complete(
		context.TODO(),
		"real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		increase.SimulationRealTimePaymentsTransferCompleteParams{
			Rejection: increase.F(increase.SimulationRealTimePaymentsTransferCompleteParamsRejection{
				RejectReasonCode: increase.F(increase.SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed),
			}),
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

func TestSimulationRealTimePaymentsTransferNewInboundWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), increase.SimulationRealTimePaymentsTransferNewInboundParams{
		AccountNumberID:       increase.F("string"),
		Amount:                increase.F(int64(1)),
		DebtorAccountNumber:   increase.F("x"),
		DebtorName:            increase.F("x"),
		DebtorRoutingNumber:   increase.F("xxxxxxxxx"),
		RemittanceInformation: increase.F("x"),
		RequestForPaymentID:   increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
