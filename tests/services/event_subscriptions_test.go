package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestEventSubscriptionsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.New(context.TODO(), &types.CreateAnEventSubscriptionParameters{URL: increase.P("https://website.com/webhooks"), SharedSecret: increase.P("x"), SelectedEventCategory: increase.P(types.CreateAnEventSubscriptionParametersSelectedEventCategoryAccountCreated)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventSubscriptionsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.Get(
		context.TODO(),
		"event_subscription_001dzz0r20rcdxgb013zqb8m04g",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventSubscriptionsUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.Update(
		context.TODO(),
		"event_subscription_001dzz0r20rcdxgb013zqb8m04g",
		&types.UpdateAnEventSubscriptionParameters{Status: increase.P(types.UpdateAnEventSubscriptionParametersStatusActive)},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventSubscriptionsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.List(context.TODO(), &types.EventSubscriptionListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
