package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestOauthConnectionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.OauthConnections.Get(context.TODO(), "connection_dauknoksyr4wilz4e6my")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestOauthConnectionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.OauthConnections.List(context.TODO(), &types.OauthConnectionListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
