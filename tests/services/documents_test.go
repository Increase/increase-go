package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestDocumentsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Documents.New(context.TODO(), &types.SimulateATaxDocumentBeingCreatedParameters{AccountID: increase.P("account_in71c4amph0vgo2qllky")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
