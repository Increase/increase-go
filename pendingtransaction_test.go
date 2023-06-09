// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestPendingTransactionGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.PendingTransactions.List(context.TODO(), increase.PendingTransactionListParams{
		AccountID: increase.F("string"),
		CreatedAt: increase.F(increase.PendingTransactionListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}),
		Cursor:    increase.F("string"),
		Limit:     increase.F(int64(0)),
		RouteID:   increase.F("string"),
		SourceID:  increase.F("string"),
		Status:    increase.F(increase.PendingTransactionListParamsStatus{In: increase.F([]increase.PendingTransactionListParamsStatusIn{increase.PendingTransactionListParamsStatusInPending, increase.PendingTransactionListParamsStatusInPending, increase.PendingTransactionListParamsStatusInPending})}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
