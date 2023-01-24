package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestLimitsCreateWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Limits.Create(context.TODO(), &types.CreateALimitParameters{Metric: increase.P(types.CreateALimitParametersMetricCount), ModelID: increase.P("account0"), Value: increase.P(1234)})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestLimitsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Limits.Retrieve(context.TODO(), "limit_fku42k0qtc8ulsuas38q")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestLimitsUpdate(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Limits.Update(
		context.TODO(),
		"limit_fku42k0qtc8ulsuas38q",
		&types.UpdateALimitParameters{Status: increase.P(types.UpdateALimitParametersStatusInactive)},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestLimitsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Limits.List(context.TODO(), &types.ListLimitsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
