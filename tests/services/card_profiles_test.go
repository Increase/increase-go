package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestCardProfilesNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.New(context.TODO(), &types.CreateACardProfileParameters{Description: increase.P("My Card Profile"), DigitalWallets: increase.P(types.CreateACardProfileParametersDigitalWallets{TextColor: increase.P(types.CreateACardProfileParametersDigitalWalletsTextColor{Red: increase.P(int64(26)), Green: increase.P(int64(43)), Blue: increase.P(int64(59))}), IssuerName: increase.P("MyBank"), CardDescription: increase.P("MyBank Signature Card"), ContactWebsite: increase.P("https://example.com"), ContactEmail: increase.P("user@example.com"), ContactPhone: increase.P("+18885551212"), BackgroundImageFileID: increase.P("file_1ai913suu1zfn1pdetru"), AppIconFileID: increase.P("file_8zxqkwlh43wo144u8yec")})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCardProfilesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.Get(
		context.TODO(),
		"card_profile_cox5y73lob2eqly18piy",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCardProfilesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.List(context.TODO(), &types.CardProfileListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), Status: increase.P(types.CardProfileListParamsStatus{In: increase.P([]types.CardProfileListParamsStatusIn{types.CardProfileListParamsStatusInPending, types.CardProfileListParamsStatusInPending, types.CardProfileListParamsStatusInPending})})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
