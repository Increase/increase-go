// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

func TestACHTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ACHTransfers.New(context.TODO(), increase.ACHTransferNewParams{
		AccountID:           increase.F("account_in71c4amph0vgo2qllky"),
		Amount:              increase.F(int64(100)),
		StatementDescriptor: increase.F("New ACH transfer"),
		AccountNumber:       increase.F("987654321"),
		Addenda: increase.F(increase.ACHTransferNewParamsAddenda{
			Category: increase.F(increase.ACHTransferNewParamsAddendaCategoryFreeform),
			Freeform: increase.F(increase.ACHTransferNewParamsAddendaFreeform{
				Entries: increase.F([]increase.ACHTransferNewParamsAddendaFreeformEntry{{
					PaymentRelatedInformation: increase.F("x"),
				}, {
					PaymentRelatedInformation: increase.F("x"),
				}, {
					PaymentRelatedInformation: increase.F("x"),
				}}),
			}),
			PaymentOrderRemittanceAdvice: increase.F(increase.ACHTransferNewParamsAddendaPaymentOrderRemittanceAdvice{
				Invoices: increase.F([]increase.ACHTransferNewParamsAddendaPaymentOrderRemittanceAdviceInvoice{{
					InvoiceNumber: increase.F("x"),
					PaidAmount:    increase.F(int64(0)),
				}, {
					InvoiceNumber: increase.F("x"),
					PaidAmount:    increase.F(int64(0)),
				}, {
					InvoiceNumber: increase.F("x"),
					PaidAmount:    increase.F(int64(0)),
				}}),
			}),
		}),
		CompanyDescriptiveDate:   increase.F("x"),
		CompanyDiscretionaryData: increase.F("x"),
		CompanyEntryDescription:  increase.F("x"),
		CompanyName:              increase.F("x"),
		DestinationAccountHolder: increase.F(increase.ACHTransferNewParamsDestinationAccountHolderBusiness),
		EffectiveDate:            increase.F(time.Now()),
		ExternalAccountID:        increase.F("string"),
		Funding:                  increase.F(increase.ACHTransferNewParamsFundingChecking),
		IndividualID:             increase.F("x"),
		IndividualName:           increase.F("x"),
		PreferredEffectiveDate: increase.F(increase.ACHTransferNewParamsPreferredEffectiveDate{
			Date:               increase.F(time.Now()),
			SettlementSchedule: increase.F(increase.ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleSameDay),
		}),
		RequireApproval:        increase.F(true),
		RoutingNumber:          increase.F("101050001"),
		StandardEntryClassCode: increase.F(increase.ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferGet(t *testing.T) {
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
	_, err := client.ACHTransfers.Get(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.ACHTransfers.List(context.TODO(), increase.ACHTransferListParams{
		AccountID: increase.F("string"),
		CreatedAt: increase.F(increase.ACHTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:            increase.F("string"),
		ExternalAccountID: increase.F("string"),
		IdempotencyKey:    increase.F("x"),
		Limit:             increase.F(int64(1)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferApprove(t *testing.T) {
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
	_, err := client.ACHTransfers.Approve(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferCancel(t *testing.T) {
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
	_, err := client.ACHTransfers.Cancel(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
