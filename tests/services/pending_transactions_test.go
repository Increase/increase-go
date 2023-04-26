package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestPendingTransactionGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.Get(
		context.TODO(),
		"pending_transaction_k1sfetcau2qbvjbzgju4",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPendingTransactionListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.List(context.TODO(), &requests.PendingTransactionListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), AccountID: increase.F("string"), RouteID: increase.F("string"), SourceID: increase.F("string"), Status: increase.F(requests.PendingTransactionListParamsStatus{In: increase.F([]requests.PendingTransactionListParamsStatusIn{requests.PendingTransactionListParamsStatusInPending, requests.PendingTransactionListParamsStatusInPending, requests.PendingTransactionListParamsStatusInPending})}), CreatedAt: increase.F(requests.PendingTransactionListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
