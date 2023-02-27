package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestOauthConnectionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.OauthConnections.Get(context.TODO(), "connection_dauknoksyr4wilz4e6my")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestOauthConnectionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.OauthConnections.List(context.TODO(), &types.OauthConnectionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
