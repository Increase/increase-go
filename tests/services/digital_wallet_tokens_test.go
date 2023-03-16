package services

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestDigitalWalletTokensGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DigitalWalletTokens.Get(
		context.TODO(),
		"digital_wallet_token_izi62go3h51p369jrie0",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestDigitalWalletTokensListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.DigitalWalletTokens.List(context.TODO(), &types.DigitalWalletTokenListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), CardID: increase.P("string"), CreatedAt: increase.P(types.DigitalWalletTokenListParamsCreatedAt{After: increase.P(time.Now()), Before: increase.P(time.Now()), OnOrAfter: increase.P(time.Now()), OnOrBefore: increase.P(time.Now())})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
