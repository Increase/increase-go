package services

import (
	"context"
	"testing"

	client "increase"
)

func TestGroupsRetrieveDetails(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Groups.RetrieveDetails(context.TODO())
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
