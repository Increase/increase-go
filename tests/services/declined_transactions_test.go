package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestDeclinedTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DeclinedTransactions.Get(context.TODO(), "declined_transaction_17jbn0yyhvkt4v4ooym8")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestDeclinedTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DeclinedTransactions.List(context.TODO(), &types.DeclinedTransactionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AccountID: increase.P("string"), RouteID: increase.P("string"), CreatedAt: &types.DeclinedTransactionsListParamsCreatedAt{After: increase.P("2019-12-27T18:11:19.117Z"), Before: increase.P("2019-12-27T18:11:19.117Z"), OnOrAfter: increase.P("2019-12-27T18:11:19.117Z"), OnOrBefore: increase.P("2019-12-27T18:11:19.117Z")}})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
