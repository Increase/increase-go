package increase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestUsage(t *testing.T) {
	client := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	account, err := client.Accounts.New(context.TODO(), increase.AccountNewParams{
		Name: increase.F("My First Increase Account"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", account)
}
