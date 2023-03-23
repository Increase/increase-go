package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestPendingTransactionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.Get(
		context.TODO(),
		"pending_transaction_k1sfetcau2qbvjbzgju4",
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

func TestPendingTransactionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.List(context.TODO(), &requests.PendingTransactionListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0)), AccountID: fields.F("string"), RouteID: fields.F("string"), SourceID: fields.F("string"), Status: fields.F(requests.PendingTransactionListParamsStatus{In: fields.F([]requests.PendingTransactionListParamsStatusIn{requests.PendingTransactionListParamsStatusInPending, requests.PendingTransactionListParamsStatusInPending, requests.PendingTransactionListParamsStatusInPending})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
