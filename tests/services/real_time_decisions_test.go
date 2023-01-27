package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestRealTimeDecisionsGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.RealTimeDecisions.Get(context.TODO(), "real_time_decision_j76n2e810ezcg3zh5qtn")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestRealTimeDecisionsActionWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		&types.ActionARealTimeDecisionParameters{},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
