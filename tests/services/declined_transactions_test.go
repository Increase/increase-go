package services

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestDeclinedTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DeclinedTransactions.Get(
		context.TODO(),
		"declined_transaction_17jbn0yyhvkt4v4ooym8",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestDeclinedTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DeclinedTransactions.List(context.TODO(), &types.DeclinedTransactionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AccountID: increase.P("string"), RouteID: increase.P("string"), CreatedAt: increase.P(types.DeclinedTransactionListParamsCreatedAt{After: increase.P(time.Now()), Before: increase.P(time.Now()), OnOrAfter: increase.P(time.Now()), OnOrBefore: increase.P(time.Now())})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
