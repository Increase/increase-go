// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// CardPurchaseSupplementService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardPurchaseSupplementService] method instead.
type CardPurchaseSupplementService struct {
	Options []option.RequestOption
}

// NewCardPurchaseSupplementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCardPurchaseSupplementService(opts ...option.RequestOption) (r *CardPurchaseSupplementService) {
	r = &CardPurchaseSupplementService{}
	r.Options = opts
	return
}

// Retrieve a Card Purchase Supplement
func (r *CardPurchaseSupplementService) Get(ctx context.Context, cardPurchaseSupplementID string, opts ...option.RequestOption) (res *CardPurchaseSupplement, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardPurchaseSupplementID == "" {
		err = errors.New("missing required card_purchase_supplement_id parameter")
		return
	}
	path := fmt.Sprintf("card_purchase_supplements/%s", cardPurchaseSupplementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Purchase Supplements
func (r *CardPurchaseSupplementService) List(ctx context.Context, query CardPurchaseSupplementListParams, opts ...option.RequestOption) (res *pagination.Page[CardPurchaseSupplement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_purchase_supplements"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Card Purchase Supplements
func (r *CardPurchaseSupplementService) ListAutoPaging(ctx context.Context, query CardPurchaseSupplementListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CardPurchaseSupplement] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Additional information about a card purchase (e.g., settlement or refund), such
// as level 3 line item data.
type CardPurchaseSupplement struct {
	// The Card Purchase Supplement identifier.
	ID string `json:"id,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// Invoice-level information about the payment.
	Invoice CardPurchaseSupplementInvoice `json:"invoice,required,nullable"`
	// Line item information, such as individual products purchased.
	LineItems []CardPurchaseSupplementLineItem `json:"line_items,required,nullable"`
	// The ID of the transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_purchase_supplement`.
	Type CardPurchaseSupplementType `json:"type,required"`
	JSON cardPurchaseSupplementJSON `json:"-"`
}

// cardPurchaseSupplementJSON contains the JSON metadata for the struct
// [CardPurchaseSupplement]
type cardPurchaseSupplementJSON struct {
	ID            apijson.Field
	CardPaymentID apijson.Field
	Invoice       apijson.Field
	LineItems     apijson.Field
	TransactionID apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardPurchaseSupplement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPurchaseSupplementJSON) RawJSON() string {
	return r.raw
}

// Invoice-level information about the payment.
type CardPurchaseSupplementInvoice struct {
	// Discount given to cardholder.
	DiscountAmount int64 `json:"discount_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the discount.
	DiscountCurrency string `json:"discount_currency,required,nullable"`
	// Indicates how the merchant applied the discount.
	DiscountTreatmentCode CardPurchaseSupplementInvoiceDiscountTreatmentCode `json:"discount_treatment_code,required,nullable"`
	// Amount of duty taxes.
	DutyTaxAmount int64 `json:"duty_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the duty tax.
	DutyTaxCurrency string `json:"duty_tax_currency,required,nullable"`
	// Date the order was taken.
	OrderDate time.Time `json:"order_date,required,nullable" format:"date"`
	// The shipping cost.
	ShippingAmount int64 `json:"shipping_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the shipping
	// cost.
	ShippingCurrency string `json:"shipping_currency,required,nullable"`
	// Country code of the shipping destination.
	ShippingDestinationCountryCode string `json:"shipping_destination_country_code,required,nullable"`
	// Postal code of the shipping destination.
	ShippingDestinationPostalCode string `json:"shipping_destination_postal_code,required,nullable"`
	// Postal code of the location being shipped from.
	ShippingSourcePostalCode string `json:"shipping_source_postal_code,required,nullable"`
	// Taxes paid for freight and shipping.
	ShippingTaxAmount int64 `json:"shipping_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the shipping
	// tax.
	ShippingTaxCurrency string `json:"shipping_tax_currency,required,nullable"`
	// Tax rate for freight and shipping.
	ShippingTaxRate string `json:"shipping_tax_rate,required,nullable"`
	// Indicates how the merchant applied taxes.
	TaxTreatments CardPurchaseSupplementInvoiceTaxTreatments `json:"tax_treatments,required,nullable"`
	// Value added tax invoice reference number.
	UniqueValueAddedTaxInvoiceReference string                            `json:"unique_value_added_tax_invoice_reference,required,nullable"`
	JSON                                cardPurchaseSupplementInvoiceJSON `json:"-"`
}

// cardPurchaseSupplementInvoiceJSON contains the JSON metadata for the struct
// [CardPurchaseSupplementInvoice]
type cardPurchaseSupplementInvoiceJSON struct {
	DiscountAmount                      apijson.Field
	DiscountCurrency                    apijson.Field
	DiscountTreatmentCode               apijson.Field
	DutyTaxAmount                       apijson.Field
	DutyTaxCurrency                     apijson.Field
	OrderDate                           apijson.Field
	ShippingAmount                      apijson.Field
	ShippingCurrency                    apijson.Field
	ShippingDestinationCountryCode      apijson.Field
	ShippingDestinationPostalCode       apijson.Field
	ShippingSourcePostalCode            apijson.Field
	ShippingTaxAmount                   apijson.Field
	ShippingTaxCurrency                 apijson.Field
	ShippingTaxRate                     apijson.Field
	TaxTreatments                       apijson.Field
	UniqueValueAddedTaxInvoiceReference apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *CardPurchaseSupplementInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPurchaseSupplementInvoiceJSON) RawJSON() string {
	return r.raw
}

// Indicates how the merchant applied the discount.
type CardPurchaseSupplementInvoiceDiscountTreatmentCode string

const (
	CardPurchaseSupplementInvoiceDiscountTreatmentCodeNoInvoiceLevelDiscountProvided          CardPurchaseSupplementInvoiceDiscountTreatmentCode = "no_invoice_level_discount_provided"
	CardPurchaseSupplementInvoiceDiscountTreatmentCodeTaxCalculatedOnPostDiscountInvoiceTotal CardPurchaseSupplementInvoiceDiscountTreatmentCode = "tax_calculated_on_post_discount_invoice_total"
	CardPurchaseSupplementInvoiceDiscountTreatmentCodeTaxCalculatedOnPreDiscountInvoiceTotal  CardPurchaseSupplementInvoiceDiscountTreatmentCode = "tax_calculated_on_pre_discount_invoice_total"
)

func (r CardPurchaseSupplementInvoiceDiscountTreatmentCode) IsKnown() bool {
	switch r {
	case CardPurchaseSupplementInvoiceDiscountTreatmentCodeNoInvoiceLevelDiscountProvided, CardPurchaseSupplementInvoiceDiscountTreatmentCodeTaxCalculatedOnPostDiscountInvoiceTotal, CardPurchaseSupplementInvoiceDiscountTreatmentCodeTaxCalculatedOnPreDiscountInvoiceTotal:
		return true
	}
	return false
}

// Indicates how the merchant applied taxes.
type CardPurchaseSupplementInvoiceTaxTreatments string

const (
	CardPurchaseSupplementInvoiceTaxTreatmentsNoTaxApplies            CardPurchaseSupplementInvoiceTaxTreatments = "no_tax_applies"
	CardPurchaseSupplementInvoiceTaxTreatmentsNetPriceLineItemLevel   CardPurchaseSupplementInvoiceTaxTreatments = "net_price_line_item_level"
	CardPurchaseSupplementInvoiceTaxTreatmentsNetPriceInvoiceLevel    CardPurchaseSupplementInvoiceTaxTreatments = "net_price_invoice_level"
	CardPurchaseSupplementInvoiceTaxTreatmentsGrossPriceLineItemLevel CardPurchaseSupplementInvoiceTaxTreatments = "gross_price_line_item_level"
	CardPurchaseSupplementInvoiceTaxTreatmentsGrossPriceInvoiceLevel  CardPurchaseSupplementInvoiceTaxTreatments = "gross_price_invoice_level"
)

func (r CardPurchaseSupplementInvoiceTaxTreatments) IsKnown() bool {
	switch r {
	case CardPurchaseSupplementInvoiceTaxTreatmentsNoTaxApplies, CardPurchaseSupplementInvoiceTaxTreatmentsNetPriceLineItemLevel, CardPurchaseSupplementInvoiceTaxTreatmentsNetPriceInvoiceLevel, CardPurchaseSupplementInvoiceTaxTreatmentsGrossPriceLineItemLevel, CardPurchaseSupplementInvoiceTaxTreatmentsGrossPriceInvoiceLevel:
		return true
	}
	return false
}

type CardPurchaseSupplementLineItem struct {
	// The Card Purchase Supplement Line Item identifier.
	ID string `json:"id,required"`
	// Indicates the type of line item.
	DetailIndicator CardPurchaseSupplementLineItemsDetailIndicator `json:"detail_indicator,required,nullable"`
	// Discount amount for this specific line item.
	DiscountAmount int64 `json:"discount_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the discount.
	DiscountCurrency string `json:"discount_currency,required,nullable"`
	// Indicates how the merchant applied the discount for this specific line item.
	DiscountTreatmentCode CardPurchaseSupplementLineItemsDiscountTreatmentCode `json:"discount_treatment_code,required,nullable"`
	// Code used to categorize the purchase item.
	ItemCommodityCode string `json:"item_commodity_code,required,nullable"`
	// Description of the purchase item.
	ItemDescriptor string `json:"item_descriptor,required,nullable"`
	// The number of units of the product being purchased.
	ItemQuantity string `json:"item_quantity,required,nullable"`
	// Code used to categorize the product being purchased.
	ProductCode string `json:"product_code,required,nullable"`
	// Sales tax amount for this line item.
	SalesTaxAmount int64 `json:"sales_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the sales tax
	// assessed.
	SalesTaxCurrency string `json:"sales_tax_currency,required,nullable"`
	// Sales tax rate for this line item.
	SalesTaxRate string `json:"sales_tax_rate,required,nullable"`
	// Total amount of all line items.
	TotalAmount int64 `json:"total_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total
	// amount.
	TotalAmountCurrency string `json:"total_amount_currency,required,nullable"`
	// Cost of line item per unit of measure, in major units.
	UnitCost string `json:"unit_cost,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the unit cost.
	UnitCostCurrency string `json:"unit_cost_currency,required,nullable"`
	// Code indicating unit of measure (gallons, etc.).
	UnitOfMeasureCode string                             `json:"unit_of_measure_code,required,nullable"`
	JSON              cardPurchaseSupplementLineItemJSON `json:"-"`
}

// cardPurchaseSupplementLineItemJSON contains the JSON metadata for the struct
// [CardPurchaseSupplementLineItem]
type cardPurchaseSupplementLineItemJSON struct {
	ID                    apijson.Field
	DetailIndicator       apijson.Field
	DiscountAmount        apijson.Field
	DiscountCurrency      apijson.Field
	DiscountTreatmentCode apijson.Field
	ItemCommodityCode     apijson.Field
	ItemDescriptor        apijson.Field
	ItemQuantity          apijson.Field
	ProductCode           apijson.Field
	SalesTaxAmount        apijson.Field
	SalesTaxCurrency      apijson.Field
	SalesTaxRate          apijson.Field
	TotalAmount           apijson.Field
	TotalAmountCurrency   apijson.Field
	UnitCost              apijson.Field
	UnitCostCurrency      apijson.Field
	UnitOfMeasureCode     apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardPurchaseSupplementLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPurchaseSupplementLineItemJSON) RawJSON() string {
	return r.raw
}

// Indicates the type of line item.
type CardPurchaseSupplementLineItemsDetailIndicator string

const (
	CardPurchaseSupplementLineItemsDetailIndicatorNormal  CardPurchaseSupplementLineItemsDetailIndicator = "normal"
	CardPurchaseSupplementLineItemsDetailIndicatorCredit  CardPurchaseSupplementLineItemsDetailIndicator = "credit"
	CardPurchaseSupplementLineItemsDetailIndicatorPayment CardPurchaseSupplementLineItemsDetailIndicator = "payment"
)

func (r CardPurchaseSupplementLineItemsDetailIndicator) IsKnown() bool {
	switch r {
	case CardPurchaseSupplementLineItemsDetailIndicatorNormal, CardPurchaseSupplementLineItemsDetailIndicatorCredit, CardPurchaseSupplementLineItemsDetailIndicatorPayment:
		return true
	}
	return false
}

// Indicates how the merchant applied the discount for this specific line item.
type CardPurchaseSupplementLineItemsDiscountTreatmentCode string

const (
	CardPurchaseSupplementLineItemsDiscountTreatmentCodeNoLineItemLevelDiscountProvided          CardPurchaseSupplementLineItemsDiscountTreatmentCode = "no_line_item_level_discount_provided"
	CardPurchaseSupplementLineItemsDiscountTreatmentCodeTaxCalculatedOnPostDiscountLineItemTotal CardPurchaseSupplementLineItemsDiscountTreatmentCode = "tax_calculated_on_post_discount_line_item_total"
	CardPurchaseSupplementLineItemsDiscountTreatmentCodeTaxCalculatedOnPreDiscountLineItemTotal  CardPurchaseSupplementLineItemsDiscountTreatmentCode = "tax_calculated_on_pre_discount_line_item_total"
)

func (r CardPurchaseSupplementLineItemsDiscountTreatmentCode) IsKnown() bool {
	switch r {
	case CardPurchaseSupplementLineItemsDiscountTreatmentCodeNoLineItemLevelDiscountProvided, CardPurchaseSupplementLineItemsDiscountTreatmentCodeTaxCalculatedOnPostDiscountLineItemTotal, CardPurchaseSupplementLineItemsDiscountTreatmentCodeTaxCalculatedOnPreDiscountLineItemTotal:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_purchase_supplement`.
type CardPurchaseSupplementType string

const (
	CardPurchaseSupplementTypeCardPurchaseSupplement CardPurchaseSupplementType = "card_purchase_supplement"
)

func (r CardPurchaseSupplementType) IsKnown() bool {
	switch r {
	case CardPurchaseSupplementTypeCardPurchaseSupplement:
		return true
	}
	return false
}

type CardPurchaseSupplementListParams struct {
	// Filter Card Purchase Supplements to ones belonging to the specified Card
	// Payment.
	CardPaymentID param.Field[string]                                    `query:"card_payment_id"`
	CreatedAt     param.Field[CardPurchaseSupplementListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CardPurchaseSupplementListParams]'s query parameters as
// `url.Values`.
func (r CardPurchaseSupplementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardPurchaseSupplementListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [CardPurchaseSupplementListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r CardPurchaseSupplementListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
