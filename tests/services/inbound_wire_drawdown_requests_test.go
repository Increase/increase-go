package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestInboundWireDrawdownRequestsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.InboundWireDrawdownRequests.New(context.TODO(), &types.SimulateAnInboundWireDrawdownRequestBeingCreatedParameters{RecipientAccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), OriginatorAccountNumber: increase.P("987654321"), OriginatorRoutingNumber: increase.P("101050001"), BeneficiaryAccountNumber: increase.P("987654321"), BeneficiaryRoutingNumber: increase.P("101050001"), Amount: increase.P(10000), Currency: increase.P("USD"), MessageToRecipient: increase.P("Invoice 29582")})
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}
