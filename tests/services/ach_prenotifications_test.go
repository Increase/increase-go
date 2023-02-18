package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestACHPrenotificationsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), &types.CreateAnACHPrenotificationParameters{AccountNumber: increase.P("987654321"), RoutingNumber: increase.P("101050001")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHPrenotificationsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.Get(context.TODO(), "ach_prenotification_ubjf9qqsxl3obbcn1u34")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHPrenotificationsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), &types.ACHPrenotificationListParams{})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
