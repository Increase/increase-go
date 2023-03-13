package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestCardRefundsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CardRefunds.New(context.TODO(), &types.SimulateARefundOnACardParameters{TransactionID: increase.P("transaction_uyrp7fld2ium70oa7oi")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
