// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestAutoPagination(t *testing.T) {
	client := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	iter := client.Accounts.ListAutoPaging(context.TODO(), increase.AccountListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		account := iter.Current()
		fmt.Printf("%+v\n", account)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
