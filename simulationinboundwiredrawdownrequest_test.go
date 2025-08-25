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

func TestSimulationInboundWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundWireDrawdownRequests.New(context.TODO(), increase.SimulationInboundWireDrawdownRequestNewParams{
		Amount:                             increase.F(int64(10000)),
		CreditorAccountNumber:              increase.F("987654321"),
		CreditorRoutingNumber:              increase.F("101050001"),
		Currency:                           increase.F("USD"),
		RecipientAccountNumberID:           increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		CreditorAddressLine1:               increase.F("33 Liberty Street"),
		CreditorAddressLine2:               increase.F("New York, NY, 10045"),
		CreditorAddressLine3:               increase.F("x"),
		CreditorName:                       increase.F("Ian Crease"),
		DebtorAccountNumber:                increase.F("987654321"),
		DebtorAddressLine1:                 increase.F("33 Liberty Street"),
		DebtorAddressLine2:                 increase.F("New York, NY, 10045"),
		DebtorAddressLine3:                 increase.F("x"),
		DebtorName:                         increase.F("Ian Crease"),
		DebtorRoutingNumber:                increase.F("101050001"),
		EndToEndIdentification:             increase.F("x"),
		InstructionIdentification:          increase.F("x"),
		UniqueEndToEndTransactionReference: increase.F("x"),
		UnstructuredRemittanceInformation:  increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
