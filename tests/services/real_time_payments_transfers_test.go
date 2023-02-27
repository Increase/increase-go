package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestRealTimePaymentsTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), &types.SimulateARealTimePaymentsTransferToYourAccountParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(int64(1000)), RequestForPaymentID: increase.P("real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7"), DebtorName: increase.P("x"), DebtorAccountNumber: increase.P("x"), DebtorRoutingNumber: increase.P("xxxxxxxxx"), RemittanceInformation: increase.P("x")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
