package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestAccountNumbersCreate(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.AccountNumbers.Create(context.TODO(), &types.CreateAnAccountNumberParameters{AccountID: increase.P("account_in71c4amph0vgo2qllky"), Name: increase.P("Rent payments")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountNumbersRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.AccountNumbers.Retrieve(context.TODO(), "account_number_v18nkfqm6afpsrvy82b2")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountNumbersUpdateWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.AccountNumbers.Update(
		context.TODO(),
		"account_number_v18nkfqm6afpsrvy82b2",
		&types.UpdateAnAccountNumberParameters{},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestAccountNumbersListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.AccountNumbers.List(context.TODO(), &types.ListAccountNumbersQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
