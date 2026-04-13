// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCardPurchaseSupplementService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardPurchaseSupplementService] method instead.
type SimulationCardPurchaseSupplementService struct {
	Options []option.RequestOption
}

// NewSimulationCardPurchaseSupplementService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationCardPurchaseSupplementService(opts ...option.RequestOption) (r *SimulationCardPurchaseSupplementService) {
	r = &SimulationCardPurchaseSupplementService{}
	r.Options = opts
	return
}

// Simulates the creation of a Card Purchase Supplement (Level 3 data) for a card
// settlement. This happens asynchronously in production when Visa sends enhanced
// transaction data about a purchase.
func (r *SimulationCardPurchaseSupplementService) New(ctx context.Context, body SimulationCardPurchaseSupplementNewParams, opts ...option.RequestOption) (res *CardPurchaseSupplement, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_purchase_supplements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SimulationCardPurchaseSupplementNewParams struct {
	// The identifier of the Transaction to create a Card Purchase Supplement for. The
	// Transaction must have a source of type `card_settlement`.
	TransactionID param.Field[string] `json:"transaction_id" api:"required"`
	// Invoice-level information about the payment.
	Invoice param.Field[SimulationCardPurchaseSupplementNewParamsInvoice] `json:"invoice"`
	// Line item information, such as individual products purchased.
	LineItems param.Field[[]SimulationCardPurchaseSupplementNewParamsLineItem] `json:"line_items"`
}

func (r SimulationCardPurchaseSupplementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invoice-level information about the payment.
type SimulationCardPurchaseSupplementNewParamsInvoice struct {
	// Discount given to cardholder.
	DiscountAmount param.Field[int64] `json:"discount_amount"`
	// Amount of duty taxes.
	DutyTaxAmount param.Field[int64] `json:"duty_tax_amount"`
	// Date the order was taken.
	OrderDate param.Field[time.Time] `json:"order_date" format:"date"`
	// The shipping cost.
	ShippingAmount param.Field[int64] `json:"shipping_amount"`
	// Country code of the shipping destination.
	ShippingDestinationCountryCode param.Field[string] `json:"shipping_destination_country_code"`
	// Postal code of the shipping destination.
	ShippingDestinationPostalCode param.Field[string] `json:"shipping_destination_postal_code"`
	// Postal code of the location being shipped from.
	ShippingSourcePostalCode param.Field[string] `json:"shipping_source_postal_code"`
	// Taxes paid for freight and shipping.
	ShippingTaxAmount param.Field[int64] `json:"shipping_tax_amount"`
	// Tax rate for freight and shipping.
	ShippingTaxRate param.Field[string] `json:"shipping_tax_rate"`
	// Value added tax invoice reference number.
	UniqueValueAddedTaxInvoiceReference param.Field[string] `json:"unique_value_added_tax_invoice_reference"`
}

func (r SimulationCardPurchaseSupplementNewParamsInvoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardPurchaseSupplementNewParamsLineItem struct {
	// Discount amount for this specific line item.
	DiscountAmount param.Field[int64] `json:"discount_amount"`
	// Code used to categorize the purchase item.
	ItemCommodityCode param.Field[string] `json:"item_commodity_code"`
	// Description of the purchase item.
	ItemDescriptor param.Field[string] `json:"item_descriptor"`
	// The number of units of the product being purchased.
	ItemQuantity param.Field[string] `json:"item_quantity"`
	// Code used to categorize the product being purchased.
	ProductCode param.Field[string] `json:"product_code"`
	// Sales tax amount for this line item.
	SalesTaxAmount param.Field[int64] `json:"sales_tax_amount"`
	// Sales tax rate for this line item.
	SalesTaxRate param.Field[string] `json:"sales_tax_rate"`
	// Total amount of all line items.
	TotalAmount param.Field[int64] `json:"total_amount"`
	// Cost of line item per unit of measure, in major units.
	UnitCost param.Field[string] `json:"unit_cost"`
	// Code indicating unit of measure (gallons, etc.).
	UnitOfMeasureCode param.Field[string] `json:"unit_of_measure_code"`
}

func (r SimulationCardPurchaseSupplementNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
