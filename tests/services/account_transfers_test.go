package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
)

func TestAccountTransfersComplete(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.AccountTransfers.Complete(
		context.TODO(),
		"account_transfer_7k9qe1ysdgqztnt63l7n",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
