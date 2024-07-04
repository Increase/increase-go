// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

func TestSimulationInboundInternationalACHTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundInternationalACHTransfers.New(context.TODO(), increase.SimulationInboundInternationalACHTransferNewParams{
		AccountNumberID:                   increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                            increase.F(int64(1000)),
		ForeignPaymentAmount:              increase.F(int64(10650)),
		OriginatingCurrencyCode:           increase.F("NOK"),
		OriginatorCompanyEntryDescription: increase.F("x"),
		OriginatorName:                    increase.F("x"),
		ReceiverIdentificationNumber:      increase.F("x"),
		ReceivingCompanyOrIndividualName:  increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
