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
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.PendingTransactions.Get(context.TODO(), "pending_transaction_k1sfetcau2qbvjbzgju4")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPendingTransactionListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.PendingTransactions.List(context.TODO(), increase.PendingTransactionListParams{
		AccountID: increase.F("string"),
		Category: increase.F(increase.PendingTransactionListParamsCategory{
			In: increase.F([]increase.PendingTransactionListParamsCategoryIn{increase.PendingTransactionListParamsCategoryInAccountTransferInstruction, increase.PendingTransactionListParamsCategoryInACHTransferInstruction, increase.PendingTransactionListParamsCategoryInCardAuthorization}),
		}),
		CreatedAt: increase.F(increase.PendingTransactionListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:   increase.F("string"),
		Limit:    increase.F(int64(0)),
		RouteID:  increase.F("string"),
		SourceID: increase.F("string"),
		Status: increase.F(increase.PendingTransactionListParamsStatus{
			In: increase.F([]increase.PendingTransactionListParamsStatusIn{increase.PendingTransactionListParamsStatusInPending, increase.PendingTransactionListParamsStatusInComplete}),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
