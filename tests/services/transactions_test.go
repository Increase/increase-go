package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.Get(
		context.TODO(),
		"transaction_uyrp7fld2ium70oa7oi",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Transactions.List(context.TODO(), &requests.TransactionListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0)), AccountID: fields.F("string"), RouteID: fields.F("string"), CreatedAt: fields.F(requests.TransactionListParamsCreatedAt{After: fields.F(time.Now()), Before: fields.F(time.Now()), OnOrAfter: fields.F(time.Now()), OnOrBefore: fields.F(time.Now())}), Category: fields.F(requests.TransactionListParamsCategory{In: fields.F([]requests.TransactionListParamsCategoryIn{requests.TransactionListParamsCategoryInAccountTransferIntention, requests.TransactionListParamsCategoryInAccountTransferIntention, requests.TransactionListParamsCategoryInAccountTransferIntention})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
