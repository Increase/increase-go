package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.New(context.TODO(), &requests.WireDrawdownRequestNewParams{AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.F(int64(10000)), MessageToRecipient: increase.F("Invoice 29582"), RecipientAccountNumber: increase.F("987654321"), RecipientRoutingNumber: increase.F("101050001"), RecipientName: increase.F("Ian Crease"), RecipientAddressLine1: increase.F("33 Liberty Street"), RecipientAddressLine2: increase.F("New York, NY, 10045"), RecipientAddressLine3: increase.F("x")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.Get(
		context.TODO(),
		"wire_drawdown_request_q6lmocus3glo0lr2bfv3",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.List(context.TODO(), &requests.WireDrawdownRequestListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0))})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
