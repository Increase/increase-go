package services

import (
	"context"
	"increase"
	"increase/options"
	"testing"
)

func TestAccountTransfersComplete(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.AccountTransfers.Complete(context.TODO(), "account_transfer_7k9qe1ysdgqztnt63l7n")
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}
