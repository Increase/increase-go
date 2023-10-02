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
		option.WithAPIKey("APIKey"),
	)
	_, err := client.CardProfiles.New(context.TODO(), increase.CardProfileNewParams{
		Description: increase.F("x"),
		DigitalWallets: increase.F(increase.CardProfileNewParamsDigitalWallets{
			TextColor: increase.F(increase.CardProfileNewParamsDigitalWalletsTextColor{
				Red:   increase.F(int64(0)),
				Green: increase.F(int64(0)),
				Blue:  increase.F(int64(0)),
			}),
			IssuerName:            increase.F("x"),
			CardDescription:       increase.F("x"),
			ContactWebsite:        increase.F("string"),
			ContactEmail:          increase.F("x"),
			ContactPhone:          increase.F("x"),
			BackgroundImageFileID: increase.F("string"),
			AppIconFileID:         increase.F("string"),
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
		option.WithAPIKey("APIKey"),
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
		option.WithAPIKey("APIKey"),
	)
	_, err := client.CardProfiles.List(context.TODO(), increase.CardProfileListParams{
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
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
		option.WithAPIKey("APIKey"),
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
