package services

import (
	"context"
	"increase"
	"increase/options"
	"testing"
)

func TestGroupsGetDetails(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Groups.GetDetails(context.TODO())
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
