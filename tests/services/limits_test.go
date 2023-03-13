package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestLimitsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.New(context.TODO(), &types.CreateALimitParameters{Metric: increase.P(types.CreateALimitParametersMetricCount), Interval: increase.P(types.CreateALimitParametersIntervalTransaction), ModelID: increase.P("account_in71c4amph0vgo2qllky"), Value: increase.P(int64(1234))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestLimitsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.Get(
		context.TODO(),
		"limit_fku42k0qtc8ulsuas38q",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestLimitsUpdate(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.Update(
		context.TODO(),
		"limit_fku42k0qtc8ulsuas38q",
		&types.UpdateALimitParameters{Status: increase.P(types.UpdateALimitParametersStatusInactive)},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestLimitsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Limits.List(context.TODO(), &types.LimitListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), ModelID: increase.P("x"), Status: increase.P("x")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
