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

func TestCardPushTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CardPushTransfers.New(context.TODO(), increase.CardPushTransferNewParams{
		Amount:                        increase.F(int64(100)),
		BusinessApplicationIdentifier: increase.F(increase.CardPushTransferNewParamsBusinessApplicationIdentifierFundsDisbursement),
		CardTokenID:                   increase.F("outbound_card_token_zlt0ml6youq3q7vcdlg0"),
		MerchantCategoryCode:          increase.F("1234"),
		MerchantCityName:              increase.F("New York"),
		MerchantName:                  increase.F("Acme Corp"),
		MerchantNamePrefix:            increase.F("Acme"),
		MerchantPostalCode:            increase.F("10045"),
		MerchantState:                 increase.F("NY"),
		RecipientName:                 increase.F("Ian Crease"),
		SenderAddressCity:             increase.F("New York"),
		SenderAddressLine1:            increase.F("33 Liberty Street"),
		SenderAddressPostalCode:       increase.F("10045"),
		SenderAddressState:            increase.F("NY"),
		SenderName:                    increase.F("Ian Crease"),
		SourceAccountNumberID:         increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		RequireApproval:               increase.F(true),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardPushTransferGet(t *testing.T) {
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
	_, err := client.CardPushTransfers.Get(context.TODO(), "outbound_card_push_transfer_e0z9rdpamraczh4tvwye")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardPushTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.CardPushTransfers.List(context.TODO(), increase.CardPushTransferListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.CardPushTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.CardPushTransferListParamsStatus{
			In: increase.F([]increase.CardPushTransferListParamsStatusIn{increase.CardPushTransferListParamsStatusInPendingApproval}),
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

func TestCardPushTransferApprove(t *testing.T) {
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
	_, err := client.CardPushTransfers.Approve(context.TODO(), "outbound_card_push_transfer_e0z9rdpamraczh4tvwye")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardPushTransferCancel(t *testing.T) {
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
	_, err := client.CardPushTransfers.Cancel(context.TODO(), "outbound_card_push_transfer_e0z9rdpamraczh4tvwye")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
