package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestDigitalWalletTokenRequestsNew(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Simulations.DigitalWalletTokenRequests.New(context.TODO(), &types.SimulateDigitalWalletProvisioningForACardParameters{CardID: increase.P("card_oubs0hwk5rn6knuecxg2")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
