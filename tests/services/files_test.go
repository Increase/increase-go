package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestFilesNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Files.New(context.TODO(), &types.CreateAFileParameters{File: nil, Purpose: increase.P(types.CreateAFileParametersPurposeCheckImageFront)})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestFilesGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Files.Get(context.TODO(), "file_makxrc67oh9l6sg7w9yc")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestFilesListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Files.List(context.TODO(), &types.ListFilesQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
