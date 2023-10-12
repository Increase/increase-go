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
		Amount:                                  increase.F(int64(10000)),
		BeneficiaryAccountNumber:                increase.F("987654321"),
		BeneficiaryRoutingNumber:                increase.F("101050001"),
		Currency:                                increase.F("USD"),
		MessageToRecipient:                      increase.F("Invoice 29582"),
		OriginatorAccountNumber:                 increase.F("987654321"),
		OriginatorRoutingNumber:                 increase.F("101050001"),
		RecipientAccountNumberID:                increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		BeneficiaryAddressLine1:                 increase.F("33 Liberty Street"),
		BeneficiaryAddressLine2:                 increase.F("New York, NY, 10045"),
		BeneficiaryAddressLine3:                 increase.F("x"),
		BeneficiaryName:                         increase.F("Ian Crease"),
		OriginatorAddressLine1:                  increase.F("33 Liberty Street"),
		OriginatorAddressLine2:                  increase.F("New York, NY, 10045"),
		OriginatorAddressLine3:                  increase.F("x"),
		OriginatorName:                          increase.F("Ian Crease"),
		OriginatorToBeneficiaryInformationLine1: increase.F("x"),
		OriginatorToBeneficiaryInformationLine2: increase.F("x"),
		OriginatorToBeneficiaryInformationLine3: increase.F("x"),
		OriginatorToBeneficiaryInformationLine4: increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
