package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestDigitalWalletTokenRequestsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.DigitalWalletTokenRequests.New(context.TODO(), &types.SimulateDigitalWalletProvisioningForACardParameters{CardID: increase.P("card_oubs0hwk5rn6knuecxg2")})
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}
