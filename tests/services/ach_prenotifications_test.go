package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestACHPrenotificationsNewWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ACHPrenotifications.New(context.TODO(), &types.CreateAnACHPrenotificationParameters{AccountNumber: increase.P("987654321"), RoutingNumber: increase.P("101050001")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestACHPrenotificationsGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ACHPrenotifications.Get(context.TODO(), "ach_prenotification_ubjf9qqsxl3obbcn1u34")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestACHPrenotificationsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ACHPrenotifications.List(context.TODO(), &types.ACHPrenotificationListParams{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
