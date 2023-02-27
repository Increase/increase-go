package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestWireDrawdownRequestsNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.New(context.TODO(), &types.CreateAWireDrawdownRequestParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(int64(10000)), MessageToRecipient: increase.P("Invoice 29582"), RecipientAccountNumber: increase.P("987654321"), RecipientRoutingNumber: increase.P("101050001"), RecipientName: increase.P("Ian Crease"), RecipientAddressLine1: increase.P("33 Liberty Street"), RecipientAddressLine2: increase.P("New York, NY, 10045"), RecipientAddressLine3: increase.P("x")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestWireDrawdownRequestsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.Get(context.TODO(), "wire_drawdown_request_q6lmocus3glo0lr2bfv3")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestWireDrawdownRequestsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.WireDrawdownRequests.List(context.TODO(), &types.WireDrawdownRequestListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
