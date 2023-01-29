package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestAccountsNewWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Accounts.New(context.TODO(), &types.CreateAnAccountParameters{Name: increase.P("New Account!")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountsGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Accounts.Get(context.TODO(), "account_in71c4amph0vgo2qllky")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountsUpdateWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Accounts.Update(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
		&types.UpdateAnAccountParameters{},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Accounts.List(context.TODO(), &types.AccountListParams{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountsClose(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Accounts.Close(context.TODO(), "account_in71c4amph0vgo2qllky")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
