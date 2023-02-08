package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestCardRefundsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CardRefunds.New(context.TODO(), &types.SimulateARefundOnACardParameters{TransactionID: increase.P("transaction_uyrp7fld2ium70oa7oi")})
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}
