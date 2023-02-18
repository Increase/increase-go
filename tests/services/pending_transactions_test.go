package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestPendingTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.Get(context.TODO(), "pending_transaction_k1sfetcau2qbvjbzgju4")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestPendingTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.List(context.TODO(), &types.PendingTransactionListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
