package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestCardProfileNew(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.New(context.TODO(), requests.CardProfileNewParams{Description: increase.F("My Card Profile"), DigitalWallets: increase.F(requests.CardProfileNewParamsDigitalWallets{TextColor: increase.F(requests.CardProfileNewParamsDigitalWalletsTextColor{Red: increase.F(int64(26)), Green: increase.F(int64(43)), Blue: increase.F(int64(59))}), IssuerName: increase.F("MyBank"), CardDescription: increase.F("MyBank Signature Card"), ContactWebsite: increase.F("https://example.com"), ContactEmail: increase.F("user@example.com"), ContactPhone: increase.F("+18885551212"), BackgroundImageFileID: increase.F("file_1ai913suu1zfn1pdetru"), AppIconFileID: increase.F("file_8zxqkwlh43wo144u8yec")})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfileGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.Get(
		context.TODO(),
		"card_profile_cox5y73lob2eqly18piy",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfileListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.List(context.TODO(), requests.CardProfileListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), Status: increase.F(requests.CardProfileListParamsStatus{In: increase.F([]requests.CardProfileListParamsStatusIn{requests.CardProfileListParamsStatusInPending, requests.CardProfileListParamsStatusInPending, requests.CardProfileListParamsStatusInPending})})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
