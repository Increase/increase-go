package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestDeclinedTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DeclinedTransactions.Get(context.TODO(), "declined_transaction_17jbn0yyhvkt4v4ooym8")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestDeclinedTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DeclinedTransactions.List(context.TODO(), &types.DeclinedTransactionListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
