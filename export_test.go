// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestExportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.New(context.TODO(), increase.ExportNewParams{
		Category: increase.F(increase.ExportNewParamsCategoryTransactionCsv),
		AccountStatementOfx: increase.F(increase.ExportNewParamsAccountStatementOfx{
			AccountID: increase.F("string"),
			CreatedAt: increase.F(increase.ExportNewParamsAccountStatementOfxCreatedAt{
				After:      increase.F(time.Now()),
				Before:     increase.F(time.Now()),
				OnOrAfter:  increase.F(time.Now()),
				OnOrBefore: increase.F(time.Now()),
			}),
		}),
		BalanceCsv: increase.F(increase.ExportNewParamsBalanceCsv{
			AccountID: increase.F("string"),
			CreatedAt: increase.F(increase.ExportNewParamsBalanceCsvCreatedAt{
				After:      increase.F(time.Now()),
				Before:     increase.F(time.Now()),
				OnOrAfter:  increase.F(time.Now()),
				OnOrBefore: increase.F(time.Now()),
			}),
		}),
		TransactionCsv: increase.F(increase.ExportNewParamsTransactionCsv{
			AccountID: increase.F("account_in71c4amph0vgo2qllky"),
			CreatedAt: increase.F(increase.ExportNewParamsTransactionCsvCreatedAt{
				After:      increase.F(time.Now()),
				Before:     increase.F(time.Now()),
				OnOrAfter:  increase.F(time.Now()),
				OnOrBefore: increase.F(time.Now()),
			}),
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

func TestExportGet(t *testing.T) {
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
	_, err := client.Exports.Get(context.TODO(), "export_8s4m48qz3bclzje0zwh9")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportListWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.List(context.TODO(), increase.ExportListParams{
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
