package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestRealTimePaymentsTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), &types.SimulateARealTimePaymentsTransferToYourAccountParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(1000)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
