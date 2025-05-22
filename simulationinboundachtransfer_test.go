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

func TestSimulationInboundACHTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundACHTransfers.New(context.TODO(), increase.SimulationInboundACHTransferNewParams{
		AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:          increase.F(int64(1000)),
		Addenda: increase.F(increase.SimulationInboundACHTransferNewParamsAddenda{
			Category: increase.F(increase.SimulationInboundACHTransferNewParamsAddendaCategoryFreeform),
			Freeform: increase.F(increase.SimulationInboundACHTransferNewParamsAddendaFreeform{
				Entries: increase.F([]increase.SimulationInboundACHTransferNewParamsAddendaFreeformEntry{{
					PaymentRelatedInformation: increase.F("x"),
				}}),
			}),
		}),
		CompanyDescriptiveDate:   increase.F("x"),
		CompanyDiscretionaryData: increase.F("x"),
		CompanyEntryDescription:  increase.F("x"),
		CompanyID:                increase.F("x"),
		CompanyName:              increase.F("x"),
		ReceiverIDNumber:         increase.F("x"),
		ReceiverName:             increase.F("x"),
		ResolveAt:                increase.F(time.Now()),
		StandardEntryClassCode:   increase.F(increase.SimulationInboundACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
