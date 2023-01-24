package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestOauthConnectionsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.OauthConnections.Retrieve(context.TODO(), "connection_dauknoksyr4wilz4e6my")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestOauthConnectionsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.OauthConnections.List(context.TODO(), &types.ListOauthConnectionsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
