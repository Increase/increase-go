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

func TestWireTransferNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.New(context.TODO(), &requests.WireTransferNewParams{AccountID: increase.F("account_in71c4amph0vgo2qllky"), AccountNumber: increase.F("987654321"), RoutingNumber: increase.F("101050001"), ExternalAccountID: increase.F("string"), Amount: increase.F(int64(100)), MessageToRecipient: increase.F("New account transfer"), BeneficiaryName: increase.F("Ian Crease"), BeneficiaryAddressLine1: increase.F("33 Liberty Street"), BeneficiaryAddressLine2: increase.F("New York"), BeneficiaryAddressLine3: increase.F("NY 10045"), RequireApproval: increase.F(true)})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireTransfers.List(context.TODO(), &requests.WireTransferListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), AccountID: increase.F("string"), ExternalAccountID: increase.F("string"), CreatedAt: increase.F(requests.WireTransferListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireTransferApprove(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
