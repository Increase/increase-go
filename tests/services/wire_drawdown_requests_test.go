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

func TestWireDrawdownRequestsNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.New(context.TODO(), &requests.CreateAWireDrawdownRequestParameters{AccountNumberID: fields.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: fields.F(int64(10000)), MessageToRecipient: fields.F("Invoice 29582"), RecipientAccountNumber: fields.F("987654321"), RecipientRoutingNumber: fields.F("101050001"), RecipientName: fields.F("Ian Crease"), RecipientAddressLine1: fields.F("33 Liberty Street"), RecipientAddressLine2: fields.F("New York, NY, 10045"), RecipientAddressLine3: fields.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.Get(
		context.TODO(),
		"wire_drawdown_request_q6lmocus3glo0lr2bfv3",
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

func TestWireDrawdownRequestsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.List(context.TODO(), &requests.WireDrawdownRequestListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
