package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestEventsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Get(context.TODO(), "event_001dzz0r20rzr4zrhrr1364hy80")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.List(context.TODO(), &types.EventListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
