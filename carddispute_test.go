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

func TestCardDisputeNew(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardDisputes.New(context.TODO(), increase.CardDisputeNewParams{
		DisputedTransactionID: increase.F("string"),
		Explanation:           increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardDisputes.Get(context.TODO(), "card_dispute_h9sc95nbl1cgltpp7men")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardDisputes.List(context.TODO(), increase.CardDisputeListParams{
		CreatedAt: increase.F(increase.CardDisputeListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}),
		Cursor:    increase.F("string"),
		Limit:     increase.F(int64(0)),
		Status:    increase.F(increase.CardDisputeListParamsStatus{In: increase.F([]increase.CardDisputeListParamsStatusIn{increase.CardDisputeListParamsStatusInPendingReviewing, increase.CardDisputeListParamsStatusInPendingReviewing, increase.CardDisputeListParamsStatusInPendingReviewing})}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
