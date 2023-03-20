package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestCheckTransfersDeposit(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CheckTransfers.Deposit(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCheckTransfersMail(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CheckTransfers.Mail(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCheckTransfersReturn(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CheckTransfers.Return(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
		&types.ReturnASandboxCheckTransferParameters{Reason: increase.P(types.ReturnASandboxCheckTransferParametersReasonMailDeliveryFailure)},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
