// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestPhysicalCardProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.New(context.TODO(), increase.PhysicalCardProfileNewParams{
		CarrierImageFileID:    increase.F("file_h6v7mtipe119os47ehlu"),
		ContactPhone:          increase.F("+16505046304"),
		Description:           increase.F("My Card Profile"),
		FrontImageFileID:      increase.F("file_o6aex13wm1jcc36sgcj1"),
		ProgramID:             increase.F("program_i2v2os4mwza1oetokh9i"),
		CardStockReference:    increase.F("x"),
		CarrierStockReference: increase.F("x"),
		FrontText: increase.F(increase.PhysicalCardProfileNewParamsFrontText{
			Line1: increase.F("x"),
			Line2: increase.F("x"),
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

func TestPhysicalCardProfileGet(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.Get(context.TODO(), "physical_card_profile_m534d5rn9qyy9ufqxoec")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.List(context.TODO(), increase.PhysicalCardProfileListParams{
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.PhysicalCardProfileListParamsStatus{
			In: increase.F([]increase.PhysicalCardProfileListParamsStatusIn{increase.PhysicalCardProfileListParamsStatusInPendingCreating}),
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

func TestPhysicalCardProfileArchive(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.Archive(context.TODO(), "physical_card_profile_m534d5rn9qyy9ufqxoec")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardProfileCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.Clone(
		context.TODO(),
		"physical_card_profile_m534d5rn9qyy9ufqxoec",
		increase.PhysicalCardProfileCloneParams{
			CarrierImageFileID: increase.F("carrier_image_file_id"),
			ContactPhone:       increase.F("x"),
			Description:        increase.F("x"),
			FrontImageFileID:   increase.F("file_o6aex13wm1jcc36sgcj1"),
			FrontText: increase.F(increase.PhysicalCardProfileCloneParamsFrontText{
				Line1: increase.F("x"),
				Line2: increase.F("x"),
			}),
			ProgramID: increase.F("program_id"),
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
