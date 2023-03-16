package services

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestEventsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Get(
		context.TODO(),
		"event_001dzz0r20rzr4zrhrr1364hy80",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.List(context.TODO(), &types.EventListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AssociatedObjectID: increase.P("string"), CreatedAt: increase.P(types.EventListParamsCreatedAt{After: increase.P(time.Now()), Before: increase.P(time.Now()), OnOrAfter: increase.P(time.Now()), OnOrBefore: increase.P(time.Now())}), Category: increase.P(types.EventListParamsCategory{In: increase.P([]types.EventListParamsCategoryIn{types.EventListParamsCategoryInAccountCreated, types.EventListParamsCategoryInAccountCreated, types.EventListParamsCategoryInAccountCreated})})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
