package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestAccountStatementsCreate(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.AccountStatements.Simulations.Create(context.TODO(), &types.SimulateAnAccountStatementBeingCreatedParameters{AccountID: increase.P("account_in71c4amph0vgo2qllky")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
