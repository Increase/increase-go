package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestAccountNumbersNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountNumbers.New(context.TODO(), &types.CreateAnAccountNumberParameters{AccountID: increase.P("account_in71c4amph0vgo2qllky"), Name: increase.P("Rent payments")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountNumbersGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountNumbers.Get(
		context.TODO(),
		"account_number_v18nkfqm6afpsrvy82b2",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountNumbersUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountNumbers.Update(
		context.TODO(),
		"account_number_v18nkfqm6afpsrvy82b2",
		&types.UpdateAnAccountNumberParameters{Name: increase.P("x"), Status: increase.P(types.UpdateAnAccountNumberParametersStatusActive)},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestAccountNumbersListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.AccountNumbers.List(context.TODO(), &types.AccountNumberListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), Status: increase.P(types.AccountNumberListParamsStatusActive), AccountID: increase.P("string")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
