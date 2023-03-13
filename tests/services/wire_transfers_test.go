package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestWireTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.WireTransfers.NewInbound(context.TODO(), &types.SimulateAWireTransferToYourAccountParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(int64(1000)), BeneficiaryAddressLine1: increase.P("x"), BeneficiaryAddressLine2: increase.P("x"), BeneficiaryAddressLine3: increase.P("x"), BeneficiaryName: increase.P("x"), BeneficiaryReference: increase.P("x"), OriginatorAddressLine1: increase.P("x"), OriginatorAddressLine2: increase.P("x"), OriginatorAddressLine3: increase.P("x"), OriginatorName: increase.P("x"), OriginatorToBeneficiaryInformationLine1: increase.P("x"), OriginatorToBeneficiaryInformationLine2: increase.P("x"), OriginatorToBeneficiaryInformationLine3: increase.P("x"), OriginatorToBeneficiaryInformationLine4: increase.P("x")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
