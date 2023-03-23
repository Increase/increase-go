package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestCardProfilesNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.New(context.TODO(), &requests.CreateACardProfileParameters{Description: fields.F("My Card Profile"), DigitalWallets: fields.F(requests.CreateACardProfileParametersDigitalWallets{TextColor: fields.F(requests.CreateACardProfileParametersDigitalWalletsTextColor{Red: fields.F(int64(26)), Green: fields.F(int64(43)), Blue: fields.F(int64(59))}), IssuerName: fields.F("MyBank"), CardDescription: fields.F("MyBank Signature Card"), ContactWebsite: fields.F("https://example.com"), ContactEmail: fields.F("user@example.com"), ContactPhone: fields.F("+18885551212"), BackgroundImageFileID: fields.F("file_1ai913suu1zfn1pdetru"), AppIconFileID: fields.F("file_8zxqkwlh43wo144u8yec")})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfilesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.Get(
		context.TODO(),
		"card_profile_cox5y73lob2eqly18piy",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfilesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CardProfiles.List(context.TODO(), &requests.CardProfileListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0)), Status: fields.F(requests.CardProfileListParamsStatus{In: fields.F([]requests.CardProfileListParamsStatusIn{requests.CardProfileListParamsStatusInPending, requests.CardProfileListParamsStatusInPending, requests.CardProfileListParamsStatusInPending})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
