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

func TestCardDisputeNew(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardDisputes.New(context.TODO(), requests.CardDisputeNewParams{DisputedTransactionID: increase.F("transaction_uyrp7fld2ium70oa7oi"), Explanation: increase.F("Unauthorized recurring transaction.")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardDisputes.Get(
		context.TODO(),
		"card_dispute_h9sc95nbl1cgltpp7men",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardDisputes.List(context.TODO(), requests.CardDisputeListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), CreatedAt: increase.F(requests.CardDisputeListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}), Status: increase.F(requests.CardDisputeListParamsStatus{In: increase.F([]requests.CardDisputeListParamsStatusIn{requests.CardDisputeListParamsStatusInPendingReviewing, requests.CardDisputeListParamsStatusInPendingReviewing, requests.CardDisputeListParamsStatusInPendingReviewing})})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
