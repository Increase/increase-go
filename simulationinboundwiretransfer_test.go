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

func TestSimulationInboundWireTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundWireTransfers.New(context.TODO(), increase.SimulationInboundWireTransferNewParams{
		AccountNumberID:                    increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                             increase.F(int64(1000)),
		CreditorAddressLine1:               increase.F("x"),
		CreditorAddressLine2:               increase.F("x"),
		CreditorAddressLine3:               increase.F("x"),
		CreditorName:                       increase.F("x"),
		DebtorAddressLine1:                 increase.F("x"),
		DebtorAddressLine2:                 increase.F("x"),
		DebtorAddressLine3:                 increase.F("x"),
		DebtorName:                         increase.F("x"),
		EndToEndIdentification:             increase.F("x"),
		InstructingAgentRoutingNumber:      increase.F("x"),
		InstructionIdentification:          increase.F("x"),
		UniqueEndToEndTransactionReference: increase.F("x"),
		UnstructuredRemittanceInformation:  increase.F("x"),
		WireDrawdownRequestID:              increase.F("wire_drawdown_request_id"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
