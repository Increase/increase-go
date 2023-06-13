// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestInboundACHTransferReturnNew(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.InboundACHTransferReturns.New(context.TODO(), increase.InboundACHTransferReturnNewParams{
		Reason:        increase.F(increase.InboundACHTransferReturnNewParamsReasonAuthorizationRevokedByCustomer),
		TransactionID: increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundACHTransferReturnGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.InboundACHTransferReturns.Get(
		context.TODO(),
		"inbound_ach_transfer_return_fhcxk5huskwhmt7iz0gk",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundACHTransferReturnListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.InboundACHTransferReturns.List(context.TODO(), increase.InboundACHTransferReturnListParams{
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
