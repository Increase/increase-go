package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestPendingTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.Get(context.TODO(), "pending_transaction_k1sfetcau2qbvjbzgju4")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestPendingTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.List(context.TODO(), &types.PendingTransactionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AccountID: increase.P("string"), RouteID: increase.P("string"), SourceID: increase.P("string"), Status: &types.PendingTransactionsListParamsStatus{In: &[]types.PendingTransactionsListParamsStatusIn{types.PendingTransactionsListParamsStatusInPending, types.PendingTransactionsListParamsStatusInPending, types.PendingTransactionsListParamsStatusInPending}}})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
