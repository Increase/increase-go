// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestCardProfileNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardProfiles.New(context.TODO(), increase.CardProfileNewParams{
		Description: increase.F("My Card Profile"),
		DigitalWallets: increase.F(increase.CardProfileNewParamsDigitalWallets{
			TextColor: increase.F(increase.CardProfileNewParamsDigitalWalletsTextColor{
				Red:   increase.F(int64(26)),
				Green: increase.F(int64(43)),
				Blue:  increase.F(int64(59)),
			}),
			IssuerName:            increase.F("MyBank"),
			CardDescription:       increase.F("MyBank Signature Card"),
			ContactWebsite:        increase.F("https://example.com"),
			ContactEmail:          increase.F("user@example.com"),
			ContactPhone:          increase.F("+18885551212"),
			BackgroundImageFileID: increase.F("file_1ai913suu1zfn1pdetru"),
			AppIconFileID:         increase.F("file_8zxqkwlh43wo144u8yec"),
		}),
		PhysicalCards: increase.F(increase.CardProfileNewParamsPhysicalCards{
			ContactPhone:       increase.F("x"),
			FrontImageFileID:   increase.F("string"),
			CarrierImageFileID: increase.F("string"),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfileGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardProfiles.Get(context.TODO(), "card_profile_cox5y73lob2eqly18piy")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfileListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardProfiles.List(context.TODO(), increase.CardProfileListParams{
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(1)),
		PhysicalCardsStatus: increase.F(increase.CardProfileListParamsPhysicalCardsStatus{
			In: increase.F([]increase.CardProfileListParamsPhysicalCardsStatusIn{increase.CardProfileListParamsPhysicalCardsStatusInNotEligible, increase.CardProfileListParamsPhysicalCardsStatusInRejected, increase.CardProfileListParamsPhysicalCardsStatusInPendingCreating}),
		}),
		Status: increase.F(increase.CardProfileListParamsStatus{
			In: increase.F([]increase.CardProfileListParamsStatusIn{increase.CardProfileListParamsStatusInPending, increase.CardProfileListParamsStatusInRejected, increase.CardProfileListParamsStatusInActive}),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfileArchive(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardProfiles.Archive(context.TODO(), "card_profile_cox5y73lob2eqly18piy")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
