package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestDigitalWalletTokensGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DigitalWalletTokens.Get(context.TODO(), "digital_wallet_token_izi62go3h51p369jrie0")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestDigitalWalletTokensListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DigitalWalletTokens.List(context.TODO(), &types.DigitalWalletTokenListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
