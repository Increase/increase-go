package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestEventsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.Get(context.TODO(), "event_001dzz0r20rzr4zrhrr1364hy80")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEventsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Events.List(context.TODO(), &types.EventListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), AssociatedObjectID: increase.P("string"), CreatedAt: &types.EventsListParamsCreatedAt{After: increase.P("2019-12-27T18:11:19.117Z"), Before: increase.P("2019-12-27T18:11:19.117Z"), OnOrAfter: increase.P("2019-12-27T18:11:19.117Z"), OnOrBefore: increase.P("2019-12-27T18:11:19.117Z")}, Category: &types.EventsListParamsCategory{In: &[]types.EventsListParamsCategoryIn{types.EventsListParamsCategoryInAccountCreated, types.EventsListParamsCategoryInAccountCreated, types.EventsListParamsCategoryInAccountCreated}}})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
