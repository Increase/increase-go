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

func TestCardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.New(context.TODO(), increase.CardNewParams{
		AccountID: increase.F("account_in71c4amph0vgo2qllky"),
		AuthorizationControls: increase.F(increase.CardNewParamsAuthorizationControls{
			MaximumAuthorizationCount: increase.F(increase.CardNewParamsAuthorizationControlsMaximumAuthorizationCount{
				AllTime: increase.F(int64(0)),
			}),
			MerchantAcceptorIdentifier: increase.F(increase.CardNewParamsAuthorizationControlsMerchantAcceptorIdentifier{
				Allowed: increase.F([]increase.CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed{{
					Identifier: increase.F("x"),
				}}),
				Blocked: increase.F([]increase.CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked{{
					Identifier: increase.F("x"),
				}}),
			}),
			MerchantCategoryCode: increase.F(increase.CardNewParamsAuthorizationControlsMerchantCategoryCode{
				Allowed: increase.F([]increase.CardNewParamsAuthorizationControlsMerchantCategoryCodeAllowed{{
					Code: increase.F("xxxx"),
				}}),
				Blocked: increase.F([]increase.CardNewParamsAuthorizationControlsMerchantCategoryCodeBlocked{{
					Code: increase.F("xxxx"),
				}}),
			}),
			MerchantCountry: increase.F(increase.CardNewParamsAuthorizationControlsMerchantCountry{
				Allowed: increase.F([]increase.CardNewParamsAuthorizationControlsMerchantCountryAllowed{{
					Country: increase.F("xx"),
				}}),
				Blocked: increase.F([]increase.CardNewParamsAuthorizationControlsMerchantCountryBlocked{{
					Country: increase.F("xx"),
				}}),
			}),
			SpendingLimits: increase.F([]increase.CardNewParamsAuthorizationControlsSpendingLimit{{
				Interval:         increase.F(increase.CardNewParamsAuthorizationControlsSpendingLimitsIntervalAllTime),
				SettlementAmount: increase.F(int64(0)),
				MerchantCategoryCodes: increase.F([]increase.CardNewParamsAuthorizationControlsSpendingLimitsMerchantCategoryCode{{
					Code: increase.F("x"),
				}}),
			}}),
		}),
		BillingAddress: increase.F(increase.CardNewParamsBillingAddress{
			City:       increase.F("x"),
			Line1:      increase.F("x"),
			PostalCode: increase.F("x"),
			State:      increase.F("x"),
			Line2:      increase.F("x"),
		}),
		Description: increase.F("Card for Ian Crease"),
		DigitalWallet: increase.F(increase.CardNewParamsDigitalWallet{
			DigitalCardProfileID: increase.F("digital_card_profile_id"),
			Email:                increase.F("dev@stainless.com"),
			Phone:                increase.F("x"),
		}),
		EntityID: increase.F("entity_id"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGet(t *testing.T) {
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
	_, err := client.Cards.Get(context.TODO(), "card_oubs0hwk5rn6knuecxg2")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.Update(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
		increase.CardUpdateParams{
			AuthorizationControls: increase.F(increase.CardUpdateParamsAuthorizationControls{
				MaximumAuthorizationCount: increase.F(increase.CardUpdateParamsAuthorizationControlsMaximumAuthorizationCount{
					AllTime: increase.F(int64(0)),
				}),
				MerchantAcceptorIdentifier: increase.F(increase.CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifier{
					Allowed: increase.F([]increase.CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed{{
						Identifier: increase.F("x"),
					}}),
					Blocked: increase.F([]increase.CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked{{
						Identifier: increase.F("x"),
					}}),
				}),
				MerchantCategoryCode: increase.F(increase.CardUpdateParamsAuthorizationControlsMerchantCategoryCode{
					Allowed: increase.F([]increase.CardUpdateParamsAuthorizationControlsMerchantCategoryCodeAllowed{{
						Code: increase.F("xxxx"),
					}}),
					Blocked: increase.F([]increase.CardUpdateParamsAuthorizationControlsMerchantCategoryCodeBlocked{{
						Code: increase.F("xxxx"),
					}}),
				}),
				MerchantCountry: increase.F(increase.CardUpdateParamsAuthorizationControlsMerchantCountry{
					Allowed: increase.F([]increase.CardUpdateParamsAuthorizationControlsMerchantCountryAllowed{{
						Country: increase.F("xx"),
					}}),
					Blocked: increase.F([]increase.CardUpdateParamsAuthorizationControlsMerchantCountryBlocked{{
						Country: increase.F("xx"),
					}}),
				}),
				SpendingLimits: increase.F([]increase.CardUpdateParamsAuthorizationControlsSpendingLimit{{
					Interval:         increase.F(increase.CardUpdateParamsAuthorizationControlsSpendingLimitsIntervalAllTime),
					SettlementAmount: increase.F(int64(0)),
					MerchantCategoryCodes: increase.F([]increase.CardUpdateParamsAuthorizationControlsSpendingLimitsMerchantCategoryCode{{
						Code: increase.F("x"),
					}}),
				}}),
			}),
			BillingAddress: increase.F(increase.CardUpdateParamsBillingAddress{
				City:       increase.F("x"),
				Line1:      increase.F("x"),
				PostalCode: increase.F("x"),
				State:      increase.F("x"),
				Line2:      increase.F("x"),
			}),
			Description: increase.F("New description"),
			DigitalWallet: increase.F(increase.CardUpdateParamsDigitalWallet{
				DigitalCardProfileID: increase.F("digital_card_profile_id"),
				Email:                increase.F("dev@stainless.com"),
				Phone:                increase.F("x"),
			}),
			EntityID: increase.F("entity_id"),
			Status:   increase.F(increase.CardUpdateParamsStatusActive),
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

func TestCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.List(context.TODO(), increase.CardListParams{
		AccountID: increase.F("account_id"),
		CreatedAt: increase.F(increase.CardListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.CardListParamsStatus{
			In: increase.F([]increase.CardListParamsStatusIn{increase.CardListParamsStatusInActive}),
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

func TestCardNewDetailsIframeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.NewDetailsIframe(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
		increase.CardNewDetailsIframeParams{
			PhysicalCardID: increase.F("physical_card_id"),
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

func TestCardDetails(t *testing.T) {
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
	_, err := client.Cards.Details(context.TODO(), "card_oubs0hwk5rn6knuecxg2")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdatePin(t *testing.T) {
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
	_, err := client.Cards.UpdatePin(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
		increase.CardUpdatePinParams{
			Pin: increase.F("1234"),
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
