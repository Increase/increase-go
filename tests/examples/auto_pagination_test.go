package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestAutoPagination(t *testing.T) {
	client := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	iter := client.Accounts.ListAutoPager(context.TODO(), &requests.AccountListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		account := iter.Current()
		fmt.Printf("%+v\n", account)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
