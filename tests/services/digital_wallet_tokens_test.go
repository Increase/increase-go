package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestDigitalWalletTokensGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.DigitalWalletTokens.Get(context.TODO(), "digital_wallet_token_izi62go3h51p369jrie0")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestDigitalWalletTokensListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.DigitalWalletTokens.List(context.TODO(), &types.ListDigitalWalletTokensQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
