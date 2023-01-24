package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestPendingTransactionsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.PendingTransactions.Retrieve(context.TODO(), "pending_transaction_k1sfetcau2qbvjbzgju4")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestPendingTransactionsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.PendingTransactions.List(context.TODO(), &types.ListPendingTransactionsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
