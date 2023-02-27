package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestInboundWireDrawdownRequestsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.InboundWireDrawdownRequests.New(context.TODO(), &types.SimulateAnInboundWireDrawdownRequestBeingCreatedParameters{RecipientAccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), OriginatorAccountNumber: increase.P("987654321"), OriginatorRoutingNumber: increase.P("101050001"), BeneficiaryAccountNumber: increase.P("987654321"), BeneficiaryRoutingNumber: increase.P("101050001"), Amount: increase.P(int64(10000)), Currency: increase.P("USD"), MessageToRecipient: increase.P("Invoice 29582"), OriginatorToBeneficiaryInformationLine1: increase.P("x"), OriginatorToBeneficiaryInformationLine2: increase.P("x"), OriginatorToBeneficiaryInformationLine3: increase.P("x"), OriginatorToBeneficiaryInformationLine4: increase.P("x"), OriginatorName: increase.P("Ian Crease"), OriginatorAddressLine1: increase.P("33 Liberty Street"), OriginatorAddressLine2: increase.P("New York, NY, 10045"), OriginatorAddressLine3: increase.P("x"), BeneficiaryName: increase.P("Ian Crease"), BeneficiaryAddressLine1: increase.P("33 Liberty Street"), BeneficiaryAddressLine2: increase.P("New York, NY, 10045"), BeneficiaryAddressLine3: increase.P("x")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
