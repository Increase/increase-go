package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestCardsAuthorizeWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Cards.Authorize(context.TODO(), &types.SimulateAnAuthorizationOnACardParameters{Amount: increase.P(1000)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCardsSettlementWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Cards.Settlement(context.TODO(), &types.SimulateSettlingACardAuthorizationParameters{CardID: increase.P("card_oubs0hwk5rn6knuecxg2"), PendingTransactionID: increase.P("pending_transaction_k1sfetcau2qbvjbzgju4")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
