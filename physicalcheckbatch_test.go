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

func TestPhysicalCheckBatchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCheckBatches.New(context.TODO(), increase.PhysicalCheckBatchNewParams{
		MailingAddress: increase.F(increase.PhysicalCheckBatchNewParamsMailingAddress{
			City:       increase.F("New York"),
			Line1:      increase.F("33 Liberty Street"),
			Name:       increase.F("Ian Crease"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
			Line2:      increase.F("line2"),
			Phone:      increase.F("x"),
		}),
		ReturnAddress: increase.F(increase.PhysicalCheckBatchNewParamsReturnAddress{
			City:       increase.F("New York"),
			Line1:      increase.F("33 Liberty Street"),
			Name:       increase.F("National Phonograph Company"),
			PostalCode: increase.F("10045"),
			State:      increase.F("NY"),
			Line2:      increase.F("line2"),
			Phone:      increase.F("x"),
		}),
		ShippingMethod: increase.F(increase.PhysicalCheckBatchNewParamsShippingMethodUspsFirstClass),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCheckBatchCancel(t *testing.T) {
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
	_, err := client.PhysicalCheckBatches.Cancel(context.TODO(), "physical_check_batch_yzdwjhdbw0in6191whce")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCheckBatchComplete(t *testing.T) {
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
	_, err := client.PhysicalCheckBatches.Complete(context.TODO(), "physical_check_batch_yzdwjhdbw0in6191whce")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
