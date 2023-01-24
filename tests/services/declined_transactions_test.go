package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestDeclinedTransactionsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.DeclinedTransactions.Retrieve(context.TODO(), "declined_transaction_17jbn0yyhvkt4v4ooym8")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestDeclinedTransactionsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.DeclinedTransactions.List(context.TODO(), &types.ListDeclinedTransactionsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
