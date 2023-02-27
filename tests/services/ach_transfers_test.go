package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestACHTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.ACHTransfers.NewInbound(context.TODO(), &types.SimulateAnACHTransferToYourAccountParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(int64(1000)), CompanyDescriptiveDate: increase.P("x"), CompanyDiscretionaryData: increase.P("x"), CompanyEntryDescription: increase.P("x"), CompanyName: increase.P("x"), CompanyID: increase.P("x")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHTransfersReturnWithOptionalParams(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.ACHTransfers.Return(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
		&types.ReturnASandboxACHTransferParameters{Reason: increase.P(types.ReturnASandboxACHTransferParametersReasonInsufficientFund)},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHTransfersSubmit(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.ACHTransfers.Submit(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
