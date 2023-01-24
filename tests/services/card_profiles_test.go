package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestCardProfilesCreate(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.CardProfiles.Create(context.TODO(), &types.CreateACardProfileParameters{Description: increase.P("My Card Profile"), DigitalWallets: &types.CreateACardProfileParametersDigitalWallets{TextColor: &types.CreateACardProfileParametersDigitalWalletsTextColor{Red: increase.P(26), Green: increase.P(43), Blue: increase.P(59)}, IssuerName: increase.P("MyBank"), CardDescription: increase.P("MyBank Signature Card"), ContactWebsite: increase.P("https://example.com"), ContactEmail: increase.P("user@example.com"), ContactPhone: increase.P("+18885551212"), BackgroundImageFileID: increase.P("file_1ai913suu1zfn1pdetru"), AppIconFileID: increase.P("file_8zxqkwlh43wo144u8yec")}})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestCardProfilesRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.CardProfiles.Retrieve(context.TODO(), "card_profile_cox5y73lob2eqly18piy")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestCardProfilesListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.CardProfiles.List(context.TODO(), &types.ListCardProfilesQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
