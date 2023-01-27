package services

import (
	"context"
	"testing"

	client "increase"
)

func TestGroupsGetDetails(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Groups.GetDetails(context.TODO())
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
