package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestAccountStatementsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.AccountStatements.New(context.TODO(), &types.SimulateAnAccountStatementBeingCreatedParameters{AccountID: increase.P("account_in71c4amph0vgo2qllky")})
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}
