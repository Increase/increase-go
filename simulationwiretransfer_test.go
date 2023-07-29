// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestSimulationWireTransferNewInboundWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Simulations.WireTransfers.NewInbound(context.TODO(), increase.SimulationWireTransferNewInboundParams{
		AccountNumberID:                         increase.F("string"),
		Amount:                                  increase.F(int64(1)),
		BeneficiaryAddressLine1:                 increase.F("x"),
		BeneficiaryAddressLine2:                 increase.F("x"),
		BeneficiaryAddressLine3:                 increase.F("x"),
		BeneficiaryName:                         increase.F("x"),
		BeneficiaryReference:                    increase.F("x"),
		OriginatorAddressLine1:                  increase.F("x"),
		OriginatorAddressLine2:                  increase.F("x"),
		OriginatorAddressLine3:                  increase.F("x"),
		OriginatorName:                          increase.F("x"),
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
