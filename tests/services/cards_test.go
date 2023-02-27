package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestCardsAuthorizeWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Cards.Authorize(context.TODO(), &types.SimulateAnAuthorizationOnACardParameters{Amount: increase.P(int64(1000)), CardID: increase.P("card_oubs0hwk5rn6knuecxg2"), DigitalWalletTokenID: increase.P("string")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCardsSettlementWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Cards.Settlement(context.TODO(), &types.SimulateSettlingACardAuthorizationParameters{CardID: increase.P("card_oubs0hwk5rn6knuecxg2"), PendingTransactionID: increase.P("pending_transaction_k1sfetcau2qbvjbzgju4"), Amount: increase.P(int64(1))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
