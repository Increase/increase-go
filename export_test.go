// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestExportNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Exports.New(context.TODO(), increase.ExportNewParams{
		Category: increase.F(increase.ExportNewParamsCategoryTransactionCsv),
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
			AccountID: increase.F("string"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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
