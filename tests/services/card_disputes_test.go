package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestCardDisputesActionWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Simulations.CardDisputes.Action(
		context.TODO(),
		"card_dispute_h9sc95nbl1cgltpp7men",
		&types.SimulatesAdvancingTheStateOfACardDisputeParameters{Status: increase.P(types.SimulatesAdvancingTheStateOfACardDisputeParametersStatusAccepted)},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
