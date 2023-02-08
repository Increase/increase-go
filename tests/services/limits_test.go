package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestLimitsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.New(context.TODO(), &types.CreateALimitParameters{Metric: increase.P(types.CreateALimitParametersMetricCount), ModelID: increase.P("account_in71c4amph0vgo2qllky"), Value: increase.P(1234)})
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestLimitsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.Get(context.TODO(), "limit_fku42k0qtc8ulsuas38q")
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestLimitsUpdate(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.Update(
		context.TODO(),
		"limit_fku42k0qtc8ulsuas38q",
		&types.UpdateALimitParameters{Status: increase.P(types.UpdateALimitParametersStatusInactive)},
	)
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestLimitsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.List(context.TODO(), &types.LimitListParams{})
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}
