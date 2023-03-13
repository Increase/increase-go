package services

import (
	"context"
	"testing"

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
	_, err := c.DigitalWalletTokens.List(context.TODO(), &types.DigitalWalletTokenListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), CardID: increase.P("string"), CreatedAt: increase.P(types.DigitalWalletTokensListParamsCreatedAt{After: increase.P("2019-12-27T18:11:19.117Z"), Before: increase.P("2019-12-27T18:11:19.117Z"), OnOrAfter: increase.P("2019-12-27T18:11:19.117Z"), OnOrBefore: increase.P("2019-12-27T18:11:19.117Z")})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
