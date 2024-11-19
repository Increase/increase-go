// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
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
		Category: increase.F(increase.ExportNewParamsCategoryAccountStatementOfx),
		AccountStatementOfx: increase.F(increase.ExportNewParamsAccountStatementOfx{
			AccountID: increase.F("account_id"),
			CreatedAt: increase.F(increase.ExportNewParamsAccountStatementOfxCreatedAt{
				After:      increase.F(time.Now()),
				Before:     increase.F(time.Now()),
				OnOrAfter:  increase.F(time.Now()),
				OnOrBefore: increase.F(time.Now()),
			}),
		}),
		BalanceCsv: increase.F(increase.ExportNewParamsBalanceCsv{
			AccountID: increase.F("account_id"),
			CreatedAt: increase.F(increase.ExportNewParamsBalanceCsvCreatedAt{
				After:      increase.F(time.Now()),
				Before:     increase.F(time.Now()),
				OnOrAfter:  increase.F(time.Now()),
				OnOrBefore: increase.F(time.Now()),
			}),
			ProgramID: increase.F("program_id"),
		}),
		BookkeepingAccountBalanceCsv: increase.F(increase.ExportNewParamsBookkeepingAccountBalanceCsv{
			BookkeepingAccountID: increase.F("bookkeeping_account_id"),
			CreatedAt: increase.F(increase.ExportNewParamsBookkeepingAccountBalanceCsvCreatedAt{
				After:      increase.F(time.Now()),
				Before:     increase.F(time.Now()),
				OnOrAfter:  increase.F(time.Now()),
				OnOrBefore: increase.F(time.Now()),
			}),
		}),
		EntityCsv: increase.F(increase.ExportNewParamsEntityCsv{
			Status: increase.F(increase.ExportNewParamsEntityCsvStatus{
				In: increase.F([]increase.ExportNewParamsEntityCsvStatusIn{increase.ExportNewParamsEntityCsvStatusInActive}),
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
			ProgramID: increase.F("program_id"),
		}),
		VendorCsv: increase.F[any](map[string]interface{}{}),
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
		Category: increase.F(increase.ExportListParamsCategory{
			In: increase.F([]increase.ExportListParamsCategoryIn{increase.ExportListParamsCategoryInAccountStatementOfx}),
		}),
		CreatedAt: increase.F(increase.ExportListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.ExportListParamsStatus{
			In: increase.F([]increase.ExportListParamsStatusIn{increase.ExportListParamsStatusInPending}),
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
