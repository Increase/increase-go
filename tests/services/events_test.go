package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestEventsGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Events.Get(context.TODO(), "event_001dzz0r20rzr4zrhrr1364hy80")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestEventsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Events.List(context.TODO(), &types.EventListParams{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
