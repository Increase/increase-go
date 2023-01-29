package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestRoutingNumbersListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.RoutingNumbers.List(context.TODO(), &types.RoutingNumberListParams{RoutingNumber: increase.P("xxxxxxxxx")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
