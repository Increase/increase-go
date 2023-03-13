package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestDigitalWalletTokenRequestsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.DigitalWalletTokenRequests.New(context.TODO(), &types.SimulateDigitalWalletProvisioningForACardParameters{CardID: increase.P("card_oubs0hwk5rn6knuecxg2")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
