package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestCardRefundsNew(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Simulations.CardRefunds.New(context.TODO(), &types.SimulateARefundOnACardParameters{TransactionID: increase.P("transaction_uyrp7fld2ium70oa7oi")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
