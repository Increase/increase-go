// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

func TestDigitalCardProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.New(context.TODO(), increase.DigitalCardProfileNewParams{
		AppIconFileID:         increase.F("file_8zxqkwlh43wo144u8yec"),
		BackgroundImageFileID: increase.F("file_1ai913suu1zfn1pdetru"),
		CardDescription:       increase.F("MyBank Signature Card"),
		Description:           increase.F("My Card Profile"),
		IssuerName:            increase.F("MyBank"),
		ContactEmail:          increase.F("user@example.com"),
		ContactPhone:          increase.F("+18885551212"),
		ContactWebsite:        increase.F("https://example.com"),
		TextColor: increase.F(increase.DigitalCardProfileNewParamsTextColor{
			Blue:  increase.F(int64(59)),
			Green: increase.F(int64(43)),
			Red:   increase.F(int64(26)),
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

func TestDigitalCardProfileGet(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.Get(context.TODO(), "digital_card_profile_s3puplu90f04xhcwkiga")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDigitalCardProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.List(context.TODO(), increase.DigitalCardProfileListParams{
		Cursor:         increase.F("string"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.DigitalCardProfileListParamsStatus{
			In: increase.F([]increase.DigitalCardProfileListParamsStatusIn{increase.DigitalCardProfileListParamsStatusInPending, increase.DigitalCardProfileListParamsStatusInRejected, increase.DigitalCardProfileListParamsStatusInActive}),
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

func TestDigitalCardProfileArchive(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.Archive(context.TODO(), "digital_card_profile_s3puplu90f04xhcwkiga")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDigitalCardProfileCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.Clone(
		context.TODO(),
		"digital_card_profile_s3puplu90f04xhcwkiga",
		increase.DigitalCardProfileCloneParams{
			AppIconFileID:         increase.F("string"),
			BackgroundImageFileID: increase.F("file_1ai913suu1zfn1pdetru"),
			CardDescription:       increase.F("x"),
			ContactEmail:          increase.F("x"),
			ContactPhone:          increase.F("x"),
			ContactWebsite:        increase.F("string"),
			Description:           increase.F("x"),
			IssuerName:            increase.F("x"),
			TextColor: increase.F(increase.DigitalCardProfileCloneParamsTextColor{
				Blue:  increase.F(int64(0)),
				Green: increase.F(int64(0)),
				Red:   increase.F(int64(0)),
			}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
