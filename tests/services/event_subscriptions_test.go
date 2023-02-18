package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestEventSubscriptionsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.New(context.TODO(), &types.CreateAnEventSubscriptionParameters{URL: increase.P("https://website.com/webhooks")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventSubscriptionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.Get(context.TODO(), "event_subscription_001dzz0r20rcdxgb013zqb8m04g")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventSubscriptionsUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.Update(
		context.TODO(),
		"event_subscription_001dzz0r20rcdxgb013zqb8m04g",
		&types.UpdateAnEventSubscriptionParameters{},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventSubscriptionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.List(context.TODO(), &types.EventSubscriptionListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
