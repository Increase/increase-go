package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestRoutingNumbersListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RoutingNumbers.List(context.TODO(), &types.RoutingNumberListParams{RoutingNumber: increase.P("xxxxxxxxx")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
