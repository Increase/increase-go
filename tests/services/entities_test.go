package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestEntitiesNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.New(context.TODO(), &types.CreateAnEntityParameters{Structure: increase.P(types.CreateAnEntityParametersStructureCorporation), Relationship: increase.P(types.CreateAnEntityParametersRelationshipAffiliated)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEntitiesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.Get(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEntitiesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.List(context.TODO(), &types.EntityListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
