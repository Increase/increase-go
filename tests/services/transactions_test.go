package services

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.Get(
		context.TODO(),
		"transaction_uyrp7fld2ium70oa7oi",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.List(context.TODO(), &types.TransactionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AccountID: increase.P("string"), RouteID: increase.P("string"), CreatedAt: increase.P(types.TransactionListParamsCreatedAt{After: increase.P(time.Now()), Before: increase.P(time.Now()), OnOrAfter: increase.P(time.Now()), OnOrBefore: increase.P(time.Now())}), Category: increase.P(types.TransactionListParamsCategory{In: increase.P([]types.TransactionListParamsCategoryIn{types.TransactionListParamsCategoryInAccountTransferIntention, types.TransactionListParamsCategoryInAccountTransferIntention, types.TransactionListParamsCategoryInAccountTransferIntention})})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
