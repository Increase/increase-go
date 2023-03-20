package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestAccountsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.New(context.TODO(), &types.CreateAnAccountParameters{EntityID: increase.P("string"), ProgramID: increase.P("string"), InformationalEntityID: increase.P("string"), Name: increase.P("New Account!")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Get(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountsUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Update(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
		&types.UpdateAnAccountParameters{Name: increase.P("My renamed account")},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.List(context.TODO(), &types.AccountListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), EntityID: increase.P("string"), Status: increase.P(types.AccountListParamsStatusOpen)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountsClose(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Close(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
