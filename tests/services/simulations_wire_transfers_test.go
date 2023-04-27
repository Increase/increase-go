package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestSimulationWireTransferNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.WireTransfers.NewInbound(context.TODO(), requests.SimulationWireTransferNewInboundParams{AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.F(int64(1000)), BeneficiaryAddressLine1: increase.F("x"), BeneficiaryAddressLine2: increase.F("x"), BeneficiaryAddressLine3: increase.F("x"), BeneficiaryName: increase.F("x"), BeneficiaryReference: increase.F("x"), OriginatorAddressLine1: increase.F("x"), OriginatorAddressLine2: increase.F("x"), OriginatorAddressLine3: increase.F("x"), OriginatorName: increase.F("x"), OriginatorToBeneficiaryInformationLine1: increase.F("x"), OriginatorToBeneficiaryInformationLine2: increase.F("x"), OriginatorToBeneficiaryInformationLine3: increase.F("x"), OriginatorToBeneficiaryInformationLine4: increase.F("x")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
