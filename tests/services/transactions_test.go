package services

import (
	"context"
	"testing"

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
	_, err := c.Transactions.List(context.TODO(), &types.TransactionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AccountID: increase.P("string"), RouteID: increase.P("string"), CreatedAt: increase.P(types.TransactionsListParamsCreatedAt{After: increase.P("2019-12-27T18:11:19.117Z"), Before: increase.P("2019-12-27T18:11:19.117Z"), OnOrAfter: increase.P("2019-12-27T18:11:19.117Z"), OnOrBefore: increase.P("2019-12-27T18:11:19.117Z")}), Category: increase.P(types.TransactionsListParamsCategory{In: increase.P([]types.TransactionsListParamsCategoryIn{types.TransactionsListParamsCategoryInAccountTransferIntention, types.TransactionsListParamsCategoryInAccountTransferIntention, types.TransactionsListParamsCategoryInAccountTransferIntention})})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
