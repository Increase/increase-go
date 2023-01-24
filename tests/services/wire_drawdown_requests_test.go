package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestWireDrawdownRequestsCreateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireDrawdownRequests.Create(context.TODO(), &types.CreateAWireDrawdownRequestParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(10000), MessageToRecipient: increase.P("Invoice 29582"), RecipientAccountNumber: increase.P("987654321"), RecipientRoutingNumber: increase.P("101050001")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestWireDrawdownRequestsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireDrawdownRequests.Retrieve(context.TODO(), "wire_drawdown_request_q6lmocus3glo0lr2bfv3")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestWireDrawdownRequestsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireDrawdownRequests.List(context.TODO(), &types.ListWireDrawdownRequestsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
