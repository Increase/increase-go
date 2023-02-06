package services

import (
	"context"
	"testing"

	client "increase"
)

func TestAccountTransfersComplete(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Simulations.AccountTransfers.Complete(context.TODO(), "account_transfer_7k9qe1ysdgqztnt63l7n")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
