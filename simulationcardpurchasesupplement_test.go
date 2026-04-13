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

func TestSimulationCardPurchaseSupplementNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.CardPurchaseSupplements.New(context.TODO(), increase.SimulationCardPurchaseSupplementNewParams{
		TransactionID: increase.F("transaction_uyrp7fld2ium70oa7oi"),
		Invoice: increase.F(increase.SimulationCardPurchaseSupplementNewParamsInvoice{
			DiscountAmount:                      increase.F(int64(100)),
			DutyTaxAmount:                       increase.F(int64(200)),
			OrderDate:                           increase.F(time.Now()),
			ShippingAmount:                      increase.F(int64(300)),
			ShippingDestinationCountryCode:      increase.F("US"),
			ShippingDestinationPostalCode:       increase.F("10045"),
			ShippingSourcePostalCode:            increase.F("10045"),
			ShippingTaxAmount:                   increase.F(int64(400)),
			ShippingTaxRate:                     increase.F("0.2"),
			UniqueValueAddedTaxInvoiceReference: increase.F("12302"),
		}),
		LineItems: increase.F([]increase.SimulationCardPurchaseSupplementNewParamsLineItem{{
			DiscountAmount:    increase.F(int64(0)),
			ItemCommodityCode: increase.F("001"),
			ItemDescriptor:    increase.F("Coffee"),
			ItemQuantity:      increase.F("1"),
			ProductCode:       increase.F("101"),
			SalesTaxAmount:    increase.F(int64(0)),
			SalesTaxRate:      increase.F("-16699"),
			TotalAmount:       increase.F(int64(500)),
			UnitCost:          increase.F("5"),
			UnitOfMeasureCode: increase.F("NMB"),
		}}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
