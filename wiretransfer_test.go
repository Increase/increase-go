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

func TestWireTransferNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.New(context.TODO(), increase.WireTransferNewParams{
		AccountID:               increase.F("string"),
		Amount:                  increase.F(int64(1)),
		BeneficiaryName:         increase.F("x"),
		MessageToRecipient:      increase.F("x"),
		AccountNumber:           increase.F("x"),
		BeneficiaryAddressLine1: increase.F("x"),
		BeneficiaryAddressLine2: increase.F("x"),
		BeneficiaryAddressLine3: increase.F("x"),
		ExternalAccountID:       increase.F("string"),
		RequireApproval:         increase.F(true),
		RoutingNumber:           increase.F("xxxxxxxxx"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.Get(
		context.TODO(),
		"wire_transfer_5akynk7dqsq25qwk9q2u",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.List(context.TODO(), increase.WireTransferListParams{
		AccountID:         increase.F("string"),
		CreatedAt:         increase.F(increase.WireTransferListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}),
		Cursor:            increase.F("string"),
		ExternalAccountID: increase.F("string"),
		Limit:             increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferApprove(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.Approve(
		context.TODO(),
		"wire_transfer_5akynk7dqsq25qwk9q2u",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferCancel(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.Cancel(
		context.TODO(),
		"wire_transfer_5akynk7dqsq25qwk9q2u",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferReverse(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism tests are broken")
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.Reverse(
		context.TODO(),
		"wire_transfer_5akynk7dqsq25qwk9q2u",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferSubmit(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism tests are broken")
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.Submit(
		context.TODO(),
		"wire_transfer_5akynk7dqsq25qwk9q2u",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
