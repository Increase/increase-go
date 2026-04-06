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

// CardService contains methods and other services that help with interacting with
// the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	Options []option.RequestOption
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	return
}

// Create a Card
func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Card
func (r *CardService) Update(ctx context.Context, cardID string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Cards
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *pagination.Page[Card], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cards"
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

// List Cards
func (r *CardService) ListAutoPaging(ctx context.Context, query CardListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Card] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create an iframe URL for a Card to display the card details. More details about
// styling and usage can be found in the
// [documentation](/documentation/embedded-card-component).
func (r *CardService) NewDetailsIframe(ctx context.Context, cardID string, body CardNewDetailsIframeParams, opts ...option.RequestOption) (res *CardIframeURL, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/create_details_iframe", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sensitive details for a Card include the primary account number, expiry, card
// verification code, and PIN.
func (r *CardService) Details(ctx context.Context, cardID string, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/details", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Card's PIN
func (r *CardService) UpdatePin(ctx context.Context, cardID string, body CardUpdatePinParams, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/update_pin", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cards are commercial credit cards. They'll immediately work for online purchases
// after you create them. All cards maintain a credit limit of 100% of the
// Account’s available balance at the time of transaction. Funds are deducted from
// the Account upon transaction settlement.
type Card struct {
	// The card identifier.
	ID string `json:"id" api:"required"`
	// The identifier for the account this card belongs to.
	AccountID string `json:"account_id" api:"required"`
	// Controls that restrict how this card can be used.
	AuthorizationControls CardAuthorizationControls `json:"authorization_controls" api:"required,nullable"`
	// The Card's billing address.
	BillingAddress CardBillingAddress `json:"billing_address" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The card's description for display purposes.
	Description string `json:"description" api:"required,nullable"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet CardDigitalWallet `json:"digital_wallet" api:"required,nullable"`
	// The identifier for the entity associated with this card.
	EntityID string `json:"entity_id" api:"required,nullable"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month" api:"required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year" api:"required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 string `json:"last4" api:"required"`
	// This indicates if payments can be made with the card.
	Status CardStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type        CardType               `json:"type" api:"required"`
	ExtraFields map[string]interface{} `json:"-" api:"extrafields"`
	JSON        cardJSON               `json:"-"`
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	ID                    apijson.Field
	AccountID             apijson.Field
	AuthorizationControls apijson.Field
	BillingAddress        apijson.Field
	CreatedAt             apijson.Field
	Description           apijson.Field
	DigitalWallet         apijson.Field
	EntityID              apijson.Field
	ExpirationMonth       apijson.Field
	ExpirationYear        apijson.Field
	IdempotencyKey        apijson.Field
	Last4                 apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardJSON) RawJSON() string {
	return r.raw
}

// Controls that restrict how this card can be used.
type CardAuthorizationControls struct {
	// Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations
	// on this card.
	MerchantAcceptorIdentifier CardAuthorizationControlsMerchantAcceptorIdentifier `json:"merchant_acceptor_identifier" api:"required,nullable"`
	// Restricts which Merchant Category Codes are allowed or blocked for
	// authorizations on this card.
	MerchantCategoryCode CardAuthorizationControlsMerchantCategoryCode `json:"merchant_category_code" api:"required,nullable"`
	// Restricts which merchant countries are allowed or blocked for authorizations on
	// this card.
	MerchantCountry CardAuthorizationControlsMerchantCountry `json:"merchant_country" api:"required,nullable"`
	// Controls how many times this card can be used.
	Usage CardAuthorizationControlsUsage `json:"usage" api:"required,nullable"`
	JSON  cardAuthorizationControlsJSON  `json:"-"`
}

// cardAuthorizationControlsJSON contains the JSON metadata for the struct
// [CardAuthorizationControls]
type cardAuthorizationControlsJSON struct {
	MerchantAcceptorIdentifier apijson.Field
	MerchantCategoryCode       apijson.Field
	MerchantCountry            apijson.Field
	Usage                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardAuthorizationControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsJSON) RawJSON() string {
	return r.raw
}

// Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations
// on this card.
type CardAuthorizationControlsMerchantAcceptorIdentifier struct {
	// The Merchant Acceptor IDs that are allowed for authorizations on this card.
	Allowed []CardAuthorizationControlsMerchantAcceptorIdentifierAllowed `json:"allowed" api:"required,nullable"`
	// The Merchant Acceptor IDs that are blocked for authorizations on this card.
	Blocked []CardAuthorizationControlsMerchantAcceptorIdentifierBlocked `json:"blocked" api:"required,nullable"`
	JSON    cardAuthorizationControlsMerchantAcceptorIdentifierJSON      `json:"-"`
}

// cardAuthorizationControlsMerchantAcceptorIdentifierJSON contains the JSON
// metadata for the struct [CardAuthorizationControlsMerchantAcceptorIdentifier]
type cardAuthorizationControlsMerchantAcceptorIdentifierJSON struct {
	Allowed     apijson.Field
	Blocked     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantAcceptorIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantAcceptorIdentifierJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsMerchantAcceptorIdentifierAllowed struct {
	// The Merchant Acceptor ID.
	Identifier string                                                         `json:"identifier" api:"required"`
	JSON       cardAuthorizationControlsMerchantAcceptorIdentifierAllowedJSON `json:"-"`
}

// cardAuthorizationControlsMerchantAcceptorIdentifierAllowedJSON contains the JSON
// metadata for the struct
// [CardAuthorizationControlsMerchantAcceptorIdentifierAllowed]
type cardAuthorizationControlsMerchantAcceptorIdentifierAllowedJSON struct {
	Identifier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantAcceptorIdentifierAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantAcceptorIdentifierAllowedJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsMerchantAcceptorIdentifierBlocked struct {
	// The Merchant Acceptor ID.
	Identifier string                                                         `json:"identifier" api:"required"`
	JSON       cardAuthorizationControlsMerchantAcceptorIdentifierBlockedJSON `json:"-"`
}

// cardAuthorizationControlsMerchantAcceptorIdentifierBlockedJSON contains the JSON
// metadata for the struct
// [CardAuthorizationControlsMerchantAcceptorIdentifierBlocked]
type cardAuthorizationControlsMerchantAcceptorIdentifierBlockedJSON struct {
	Identifier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantAcceptorIdentifierBlocked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantAcceptorIdentifierBlockedJSON) RawJSON() string {
	return r.raw
}

// Restricts which Merchant Category Codes are allowed or blocked for
// authorizations on this card.
type CardAuthorizationControlsMerchantCategoryCode struct {
	// The Merchant Category Codes that are allowed for authorizations on this card.
	Allowed []CardAuthorizationControlsMerchantCategoryCodeAllowed `json:"allowed" api:"required,nullable"`
	// The Merchant Category Codes that are blocked for authorizations on this card.
	Blocked []CardAuthorizationControlsMerchantCategoryCodeBlocked `json:"blocked" api:"required,nullable"`
	JSON    cardAuthorizationControlsMerchantCategoryCodeJSON      `json:"-"`
}

// cardAuthorizationControlsMerchantCategoryCodeJSON contains the JSON metadata for
// the struct [CardAuthorizationControlsMerchantCategoryCode]
type cardAuthorizationControlsMerchantCategoryCodeJSON struct {
	Allowed     apijson.Field
	Blocked     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantCategoryCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantCategoryCodeJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsMerchantCategoryCodeAllowed struct {
	// The Merchant Category Code (MCC).
	Code string                                                   `json:"code" api:"required"`
	JSON cardAuthorizationControlsMerchantCategoryCodeAllowedJSON `json:"-"`
}

// cardAuthorizationControlsMerchantCategoryCodeAllowedJSON contains the JSON
// metadata for the struct [CardAuthorizationControlsMerchantCategoryCodeAllowed]
type cardAuthorizationControlsMerchantCategoryCodeAllowedJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantCategoryCodeAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantCategoryCodeAllowedJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsMerchantCategoryCodeBlocked struct {
	// The Merchant Category Code (MCC).
	Code string                                                   `json:"code" api:"required"`
	JSON cardAuthorizationControlsMerchantCategoryCodeBlockedJSON `json:"-"`
}

// cardAuthorizationControlsMerchantCategoryCodeBlockedJSON contains the JSON
// metadata for the struct [CardAuthorizationControlsMerchantCategoryCodeBlocked]
type cardAuthorizationControlsMerchantCategoryCodeBlockedJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantCategoryCodeBlocked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantCategoryCodeBlockedJSON) RawJSON() string {
	return r.raw
}

// Restricts which merchant countries are allowed or blocked for authorizations on
// this card.
type CardAuthorizationControlsMerchantCountry struct {
	// The merchant countries that are allowed for authorizations on this card.
	Allowed []CardAuthorizationControlsMerchantCountryAllowed `json:"allowed" api:"required,nullable"`
	// The merchant countries that are blocked for authorizations on this card.
	Blocked []CardAuthorizationControlsMerchantCountryBlocked `json:"blocked" api:"required,nullable"`
	JSON    cardAuthorizationControlsMerchantCountryJSON      `json:"-"`
}

// cardAuthorizationControlsMerchantCountryJSON contains the JSON metadata for the
// struct [CardAuthorizationControlsMerchantCountry]
type cardAuthorizationControlsMerchantCountryJSON struct {
	Allowed     apijson.Field
	Blocked     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantCountry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantCountryJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsMerchantCountryAllowed struct {
	// The ISO 3166-1 alpha-2 country code.
	Country string                                              `json:"country" api:"required"`
	JSON    cardAuthorizationControlsMerchantCountryAllowedJSON `json:"-"`
}

// cardAuthorizationControlsMerchantCountryAllowedJSON contains the JSON metadata
// for the struct [CardAuthorizationControlsMerchantCountryAllowed]
type cardAuthorizationControlsMerchantCountryAllowedJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantCountryAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantCountryAllowedJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsMerchantCountryBlocked struct {
	// The ISO 3166-1 alpha-2 country code.
	Country string                                              `json:"country" api:"required"`
	JSON    cardAuthorizationControlsMerchantCountryBlockedJSON `json:"-"`
}

// cardAuthorizationControlsMerchantCountryBlockedJSON contains the JSON metadata
// for the struct [CardAuthorizationControlsMerchantCountryBlocked]
type cardAuthorizationControlsMerchantCountryBlockedJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsMerchantCountryBlocked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsMerchantCountryBlockedJSON) RawJSON() string {
	return r.raw
}

// Controls how many times this card can be used.
type CardAuthorizationControlsUsage struct {
	// Whether the card is for a single use or multiple uses.
	Category CardAuthorizationControlsUsageCategory `json:"category" api:"required"`
	// Controls for multi-use cards. Required if and only if `category` is `multi_use`.
	MultiUse CardAuthorizationControlsUsageMultiUse `json:"multi_use" api:"required,nullable"`
	// Controls for single-use cards. Required if and only if `category` is
	// `single_use`.
	SingleUse CardAuthorizationControlsUsageSingleUse `json:"single_use" api:"required,nullable"`
	JSON      cardAuthorizationControlsUsageJSON      `json:"-"`
}

// cardAuthorizationControlsUsageJSON contains the JSON metadata for the struct
// [CardAuthorizationControlsUsage]
type cardAuthorizationControlsUsageJSON struct {
	Category    apijson.Field
	MultiUse    apijson.Field
	SingleUse   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsUsageJSON) RawJSON() string {
	return r.raw
}

// Whether the card is for a single use or multiple uses.
type CardAuthorizationControlsUsageCategory string

const (
	CardAuthorizationControlsUsageCategorySingleUse CardAuthorizationControlsUsageCategory = "single_use"
	CardAuthorizationControlsUsageCategoryMultiUse  CardAuthorizationControlsUsageCategory = "multi_use"
)

func (r CardAuthorizationControlsUsageCategory) IsKnown() bool {
	switch r {
	case CardAuthorizationControlsUsageCategorySingleUse, CardAuthorizationControlsUsageCategoryMultiUse:
		return true
	}
	return false
}

// Controls for multi-use cards. Required if and only if `category` is `multi_use`.
type CardAuthorizationControlsUsageMultiUse struct {
	// Spending limits for this card. The most restrictive limit applies if multiple
	// limits match.
	SpendingLimits []CardAuthorizationControlsUsageMultiUseSpendingLimit `json:"spending_limits" api:"required,nullable"`
	JSON           cardAuthorizationControlsUsageMultiUseJSON            `json:"-"`
}

// cardAuthorizationControlsUsageMultiUseJSON contains the JSON metadata for the
// struct [CardAuthorizationControlsUsageMultiUse]
type cardAuthorizationControlsUsageMultiUseJSON struct {
	SpendingLimits apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardAuthorizationControlsUsageMultiUse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsUsageMultiUseJSON) RawJSON() string {
	return r.raw
}

type CardAuthorizationControlsUsageMultiUseSpendingLimit struct {
	// The interval at which the spending limit is enforced.
	Interval CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval `json:"interval" api:"required"`
	// The Merchant Category Codes (MCCs) this spending limit applies to. If not set,
	// the limit applies to all transactions.
	MerchantCategoryCodes []CardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode `json:"merchant_category_codes" api:"required,nullable"`
	// The maximum settlement amount permitted in the given interval.
	SettlementAmount int64                                                   `json:"settlement_amount" api:"required"`
	JSON             cardAuthorizationControlsUsageMultiUseSpendingLimitJSON `json:"-"`
}

// cardAuthorizationControlsUsageMultiUseSpendingLimitJSON contains the JSON
// metadata for the struct [CardAuthorizationControlsUsageMultiUseSpendingLimit]
type cardAuthorizationControlsUsageMultiUseSpendingLimitJSON struct {
	Interval              apijson.Field
	MerchantCategoryCodes apijson.Field
	SettlementAmount      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardAuthorizationControlsUsageMultiUseSpendingLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsUsageMultiUseSpendingLimitJSON) RawJSON() string {
	return r.raw
}

// The interval at which the spending limit is enforced.
type CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval string

const (
	CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalAllTime        CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "all_time"
	CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerTransaction CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_transaction"
	CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerDay         CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_day"
	CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerWeek        CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_week"
	CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerMonth       CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_month"
)

func (r CardAuthorizationControlsUsageMultiUseSpendingLimitsInterval) IsKnown() bool {
	switch r {
	case CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalAllTime, CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerTransaction, CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerDay, CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerWeek, CardAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerMonth:
		return true
	}
	return false
}

type CardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode struct {
	// The Merchant Category Code (MCC).
	Code string                                                                       `json:"code" api:"required"`
	JSON cardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCodeJSON `json:"-"`
}

// cardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCodeJSON
// contains the JSON metadata for the struct
// [CardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode]
type cardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCodeJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCodeJSON) RawJSON() string {
	return r.raw
}

// Controls for single-use cards. Required if and only if `category` is
// `single_use`.
type CardAuthorizationControlsUsageSingleUse struct {
	// The settlement amount constraint for this single-use card.
	SettlementAmount CardAuthorizationControlsUsageSingleUseSettlementAmount `json:"settlement_amount" api:"required"`
	JSON             cardAuthorizationControlsUsageSingleUseJSON             `json:"-"`
}

// cardAuthorizationControlsUsageSingleUseJSON contains the JSON metadata for the
// struct [CardAuthorizationControlsUsageSingleUse]
type cardAuthorizationControlsUsageSingleUseJSON struct {
	SettlementAmount apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CardAuthorizationControlsUsageSingleUse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsUsageSingleUseJSON) RawJSON() string {
	return r.raw
}

// The settlement amount constraint for this single-use card.
type CardAuthorizationControlsUsageSingleUseSettlementAmount struct {
	// The operator used to compare the settlement amount.
	Comparison CardAuthorizationControlsUsageSingleUseSettlementAmountComparison `json:"comparison" api:"required"`
	// The settlement amount value.
	Value int64                                                       `json:"value" api:"required"`
	JSON  cardAuthorizationControlsUsageSingleUseSettlementAmountJSON `json:"-"`
}

// cardAuthorizationControlsUsageSingleUseSettlementAmountJSON contains the JSON
// metadata for the struct
// [CardAuthorizationControlsUsageSingleUseSettlementAmount]
type cardAuthorizationControlsUsageSingleUseSettlementAmountJSON struct {
	Comparison  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationControlsUsageSingleUseSettlementAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardAuthorizationControlsUsageSingleUseSettlementAmountJSON) RawJSON() string {
	return r.raw
}

// The operator used to compare the settlement amount.
type CardAuthorizationControlsUsageSingleUseSettlementAmountComparison string

const (
	CardAuthorizationControlsUsageSingleUseSettlementAmountComparisonEquals           CardAuthorizationControlsUsageSingleUseSettlementAmountComparison = "equals"
	CardAuthorizationControlsUsageSingleUseSettlementAmountComparisonLessThanOrEquals CardAuthorizationControlsUsageSingleUseSettlementAmountComparison = "less_than_or_equals"
)

func (r CardAuthorizationControlsUsageSingleUseSettlementAmountComparison) IsKnown() bool {
	switch r {
	case CardAuthorizationControlsUsageSingleUseSettlementAmountComparisonEquals, CardAuthorizationControlsUsageSingleUseSettlementAmountComparisonLessThanOrEquals:
		return true
	}
	return false
}

// The Card's billing address.
type CardBillingAddress struct {
	// The city of the billing address.
	City string `json:"city" api:"required,nullable"`
	// The first line of the billing address.
	Line1 string `json:"line1" api:"required,nullable"`
	// The second line of the billing address.
	Line2 string `json:"line2" api:"required,nullable"`
	// The postal code of the billing address.
	PostalCode string `json:"postal_code" api:"required,nullable"`
	// The US state of the billing address.
	State string                 `json:"state" api:"required,nullable"`
	JSON  cardBillingAddressJSON `json:"-"`
}

// cardBillingAddressJSON contains the JSON metadata for the struct
// [CardBillingAddress]
type cardBillingAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardBillingAddressJSON) RawJSON() string {
	return r.raw
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
type CardDigitalWallet struct {
	// The digital card profile assigned to this digital card. Card profiles may also
	// be assigned at the program level.
	DigitalCardProfileID string `json:"digital_card_profile_id" api:"required,nullable"`
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email" api:"required,nullable"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string                `json:"phone" api:"required,nullable"`
	JSON  cardDigitalWalletJSON `json:"-"`
}

// cardDigitalWalletJSON contains the JSON metadata for the struct
// [CardDigitalWallet]
type cardDigitalWalletJSON struct {
	DigitalCardProfileID apijson.Field
	Email                apijson.Field
	Phone                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDigitalWalletJSON) RawJSON() string {
	return r.raw
}

// This indicates if payments can be made with the card.
type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

func (r CardStatus) IsKnown() bool {
	switch r {
	case CardStatusActive, CardStatusDisabled, CardStatusCanceled:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card`.
type CardType string

const (
	CardTypeCard CardType = "card"
)

func (r CardType) IsKnown() bool {
	switch r {
	case CardTypeCard:
		return true
	}
	return false
}

// An object containing the sensitive details (card number, CVC, PIN, etc) for a
// Card. These details are not included in the Card object. If you'd prefer to
// never access these details directly, you can use the
// [embedded iframe](/documentation/embedded-card-component) to display the
// information to your users.
type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID string `json:"card_id" api:"required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month" api:"required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year" api:"required"`
	// The 4-digit PIN for the card, for use with ATMs.
	Pin string `json:"pin" api:"required"`
	// The card number.
	PrimaryAccountNumber string `json:"primary_account_number" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type CardDetailsType `json:"type" api:"required"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode string          `json:"verification_code" api:"required"`
	JSON             cardDetailsJSON `json:"-"`
}

// cardDetailsJSON contains the JSON metadata for the struct [CardDetails]
type cardDetailsJSON struct {
	CardID               apijson.Field
	ExpirationMonth      apijson.Field
	ExpirationYear       apijson.Field
	Pin                  apijson.Field
	PrimaryAccountNumber apijson.Field
	Type                 apijson.Field
	VerificationCode     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDetailsJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_details`.
type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

func (r CardDetailsType) IsKnown() bool {
	switch r {
	case CardDetailsTypeCardDetails:
		return true
	}
	return false
}

// An object containing the iframe URL for a Card.
type CardIframeURL struct {
	// The time the iframe URL will expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The iframe URL for the Card. Treat this as an opaque URL.
	IframeURL string `json:"iframe_url" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_iframe_url`.
	Type CardIframeURLType `json:"type" api:"required"`
	JSON cardIframeURLJSON `json:"-"`
}

// cardIframeURLJSON contains the JSON metadata for the struct [CardIframeURL]
type cardIframeURLJSON struct {
	ExpiresAt   apijson.Field
	IframeURL   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardIframeURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardIframeURLJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_iframe_url`.
type CardIframeURLType string

const (
	CardIframeURLTypeCardIframeURL CardIframeURLType = "card_iframe_url"
)

func (r CardIframeURLType) IsKnown() bool {
	switch r {
	case CardIframeURLTypeCardIframeURL:
		return true
	}
	return false
}

type CardNewParams struct {
	// The Account the card should belong to.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Controls that restrict how this card can be used.
	AuthorizationControls param.Field[CardNewParamsAuthorizationControls] `json:"authorization_controls"`
	// The card's billing address.
	BillingAddress param.Field[CardNewParamsBillingAddress] `json:"billing_address"`
	// The description you choose to give the card.
	Description param.Field[string] `json:"description"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. To add the card to a digital wallet, you may supply an email or phone
	// number with this request. Otherwise, subscribe and then action a Real Time
	// Decision with the category `digital_wallet_token_requested` or
	// `digital_wallet_authentication_requested`.
	DigitalWallet param.Field[CardNewParamsDigitalWallet] `json:"digital_wallet"`
	// The Entity the card belongs to. You only need to supply this in rare situations
	// when the card is not for the Account holder.
	EntityID param.Field[string] `json:"entity_id"`
}

func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls that restrict how this card can be used.
type CardNewParamsAuthorizationControls struct {
	// Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations
	// on this card.
	MerchantAcceptorIdentifier param.Field[CardNewParamsAuthorizationControlsMerchantAcceptorIdentifier] `json:"merchant_acceptor_identifier"`
	// Restricts which Merchant Category Codes are allowed or blocked for
	// authorizations on this card.
	MerchantCategoryCode param.Field[CardNewParamsAuthorizationControlsMerchantCategoryCode] `json:"merchant_category_code"`
	// Restricts which merchant countries are allowed or blocked for authorizations on
	// this card.
	MerchantCountry param.Field[CardNewParamsAuthorizationControlsMerchantCountry] `json:"merchant_country"`
	// Controls how many times this card can be used.
	Usage param.Field[CardNewParamsAuthorizationControlsUsage] `json:"usage"`
}

func (r CardNewParamsAuthorizationControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations
// on this card.
type CardNewParamsAuthorizationControlsMerchantAcceptorIdentifier struct {
	// The Merchant Acceptor IDs that are allowed for authorizations on this card.
	// Authorizations with Merchant Acceptor IDs not in this list will be declined.
	Allowed param.Field[[]CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed] `json:"allowed"`
	// The Merchant Acceptor IDs that are blocked for authorizations on this card.
	// Authorizations with Merchant Acceptor IDs in this list will be declined.
	Blocked param.Field[[]CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked] `json:"blocked"`
}

func (r CardNewParamsAuthorizationControlsMerchantAcceptorIdentifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed struct {
	// The Merchant Acceptor ID.
	Identifier param.Field[string] `json:"identifier" api:"required"`
}

func (r CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked struct {
	// The Merchant Acceptor ID.
	Identifier param.Field[string] `json:"identifier" api:"required"`
}

func (r CardNewParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Restricts which Merchant Category Codes are allowed or blocked for
// authorizations on this card.
type CardNewParamsAuthorizationControlsMerchantCategoryCode struct {
	// The Merchant Category Codes that are allowed for authorizations on this card.
	// Authorizations with Merchant Category Codes not in this list will be declined.
	Allowed param.Field[[]CardNewParamsAuthorizationControlsMerchantCategoryCodeAllowed] `json:"allowed"`
	// The Merchant Category Codes that are blocked for authorizations on this card.
	// Authorizations with Merchant Category Codes in this list will be declined.
	Blocked param.Field[[]CardNewParamsAuthorizationControlsMerchantCategoryCodeBlocked] `json:"blocked"`
}

func (r CardNewParamsAuthorizationControlsMerchantCategoryCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsMerchantCategoryCodeAllowed struct {
	// The Merchant Category Code.
	Code param.Field[string] `json:"code" api:"required"`
}

func (r CardNewParamsAuthorizationControlsMerchantCategoryCodeAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsMerchantCategoryCodeBlocked struct {
	// The Merchant Category Code.
	Code param.Field[string] `json:"code" api:"required"`
}

func (r CardNewParamsAuthorizationControlsMerchantCategoryCodeBlocked) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Restricts which merchant countries are allowed or blocked for authorizations on
// this card.
type CardNewParamsAuthorizationControlsMerchantCountry struct {
	// The merchant countries that are allowed for authorizations on this card.
	// Authorizations with merchant countries not in this list will be declined.
	Allowed param.Field[[]CardNewParamsAuthorizationControlsMerchantCountryAllowed] `json:"allowed"`
	// The merchant countries that are blocked for authorizations on this card.
	// Authorizations with merchant countries in this list will be declined.
	Blocked param.Field[[]CardNewParamsAuthorizationControlsMerchantCountryBlocked] `json:"blocked"`
}

func (r CardNewParamsAuthorizationControlsMerchantCountry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsMerchantCountryAllowed struct {
	// The ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country" api:"required"`
}

func (r CardNewParamsAuthorizationControlsMerchantCountryAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsMerchantCountryBlocked struct {
	// The ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country" api:"required"`
}

func (r CardNewParamsAuthorizationControlsMerchantCountryBlocked) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how many times this card can be used.
type CardNewParamsAuthorizationControlsUsage struct {
	// Whether the card is for a single use or multiple uses.
	Category param.Field[CardNewParamsAuthorizationControlsUsageCategory] `json:"category" api:"required"`
	// Controls for multi-use cards. Required if and only if `category` is `multi_use`.
	MultiUse param.Field[CardNewParamsAuthorizationControlsUsageMultiUse] `json:"multi_use"`
	// Controls for single-use cards. Required if and only if `category` is
	// `single_use`.
	SingleUse param.Field[CardNewParamsAuthorizationControlsUsageSingleUse] `json:"single_use"`
}

func (r CardNewParamsAuthorizationControlsUsage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card is for a single use or multiple uses.
type CardNewParamsAuthorizationControlsUsageCategory string

const (
	CardNewParamsAuthorizationControlsUsageCategorySingleUse CardNewParamsAuthorizationControlsUsageCategory = "single_use"
	CardNewParamsAuthorizationControlsUsageCategoryMultiUse  CardNewParamsAuthorizationControlsUsageCategory = "multi_use"
)

func (r CardNewParamsAuthorizationControlsUsageCategory) IsKnown() bool {
	switch r {
	case CardNewParamsAuthorizationControlsUsageCategorySingleUse, CardNewParamsAuthorizationControlsUsageCategoryMultiUse:
		return true
	}
	return false
}

// Controls for multi-use cards. Required if and only if `category` is `multi_use`.
type CardNewParamsAuthorizationControlsUsageMultiUse struct {
	// Spending limits for this card. The most restrictive limit applies if multiple
	// limits match.
	SpendingLimits param.Field[[]CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimit] `json:"spending_limits"`
}

func (r CardNewParamsAuthorizationControlsUsageMultiUse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimit struct {
	// The interval at which the spending limit is enforced.
	Interval param.Field[CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval] `json:"interval" api:"required"`
	// The maximum settlement amount permitted in the given interval.
	SettlementAmount param.Field[int64] `json:"settlement_amount" api:"required"`
	// The Merchant Category Codes this spending limit applies to. If not set, the
	// limit applies to all transactions.
	MerchantCategoryCodes param.Field[[]CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode] `json:"merchant_category_codes"`
}

func (r CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interval at which the spending limit is enforced.
type CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval string

const (
	CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalAllTime        CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "all_time"
	CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerTransaction CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_transaction"
	CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerDay         CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_day"
	CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerWeek        CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_week"
	CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerMonth       CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_month"
)

func (r CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval) IsKnown() bool {
	switch r {
	case CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalAllTime, CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerTransaction, CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerDay, CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerWeek, CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerMonth:
		return true
	}
	return false
}

type CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode struct {
	// The Merchant Category Code.
	Code param.Field[string] `json:"code" api:"required"`
}

func (r CardNewParamsAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls for single-use cards. Required if and only if `category` is
// `single_use`.
type CardNewParamsAuthorizationControlsUsageSingleUse struct {
	// The settlement amount constraint for this single-use card.
	SettlementAmount param.Field[CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmount] `json:"settlement_amount" api:"required"`
}

func (r CardNewParamsAuthorizationControlsUsageSingleUse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settlement amount constraint for this single-use card.
type CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmount struct {
	// The operator used to compare the settlement amount.
	Comparison param.Field[CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison] `json:"comparison" api:"required"`
	// The settlement amount value.
	Value param.Field[int64] `json:"value" api:"required"`
}

func (r CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator used to compare the settlement amount.
type CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison string

const (
	CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonEquals           CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison = "equals"
	CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonLessThanOrEquals CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison = "less_than_or_equals"
)

func (r CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison) IsKnown() bool {
	switch r {
	case CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonEquals, CardNewParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonLessThanOrEquals:
		return true
	}
	return false
}

// The card's billing address.
type CardNewParamsBillingAddress struct {
	// The city of the billing address.
	City param.Field[string] `json:"city" api:"required"`
	// The first line of the billing address.
	Line1 param.Field[string] `json:"line1" api:"required"`
	// The postal code of the billing address.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// The US state of the billing address.
	State param.Field[string] `json:"state" api:"required"`
	// The second line of the billing address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CardNewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The contact information used in the two-factor steps for digital wallet card
// creation. To add the card to a digital wallet, you may supply an email or phone
// number with this request. Otherwise, subscribe and then action a Real Time
// Decision with the category `digital_wallet_token_requested` or
// `digital_wallet_authentication_requested`.
type CardNewParamsDigitalWallet struct {
	// The digital card profile assigned to this digital card.
	DigitalCardProfileID param.Field[string] `json:"digital_card_profile_id"`
	// An email address that can be used to contact and verify the cardholder via
	// one-time passcode over email.
	Email param.Field[string] `json:"email" format:"email"`
	// A phone number that can be used to contact and verify the cardholder via
	// one-time passcode over SMS.
	Phone param.Field[string] `json:"phone"`
}

func (r CardNewParamsDigitalWallet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParams struct {
	// Controls that restrict how this card can be used.
	AuthorizationControls param.Field[CardUpdateParamsAuthorizationControls] `json:"authorization_controls"`
	// The card's updated billing address.
	BillingAddress param.Field[CardUpdateParamsBillingAddress] `json:"billing_address"`
	// The description you choose to give the card.
	Description param.Field[string] `json:"description"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet param.Field[CardUpdateParamsDigitalWallet] `json:"digital_wallet"`
	// The Entity the card belongs to. You only need to supply this in rare situations
	// when the card is not for the Account holder.
	EntityID param.Field[string] `json:"entity_id"`
	// The status to update the Card with.
	Status param.Field[CardUpdateParamsStatus] `json:"status"`
}

func (r CardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls that restrict how this card can be used.
type CardUpdateParamsAuthorizationControls struct {
	// Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations
	// on this card.
	MerchantAcceptorIdentifier param.Field[CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifier] `json:"merchant_acceptor_identifier"`
	// Restricts which Merchant Category Codes are allowed or blocked for
	// authorizations on this card.
	MerchantCategoryCode param.Field[CardUpdateParamsAuthorizationControlsMerchantCategoryCode] `json:"merchant_category_code"`
	// Restricts which merchant countries are allowed or blocked for authorizations on
	// this card.
	MerchantCountry param.Field[CardUpdateParamsAuthorizationControlsMerchantCountry] `json:"merchant_country"`
	// Controls how many times this card can be used.
	Usage param.Field[CardUpdateParamsAuthorizationControlsUsage] `json:"usage"`
}

func (r CardUpdateParamsAuthorizationControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations
// on this card.
type CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifier struct {
	// The Merchant Acceptor IDs that are allowed for authorizations on this card.
	// Authorizations with Merchant Acceptor IDs not in this list will be declined.
	Allowed param.Field[[]CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed] `json:"allowed"`
	// The Merchant Acceptor IDs that are blocked for authorizations on this card.
	// Authorizations with Merchant Acceptor IDs in this list will be declined.
	Blocked param.Field[[]CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked] `json:"blocked"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed struct {
	// The Merchant Acceptor ID.
	Identifier param.Field[string] `json:"identifier" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked struct {
	// The Merchant Acceptor ID.
	Identifier param.Field[string] `json:"identifier" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantAcceptorIdentifierBlocked) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Restricts which Merchant Category Codes are allowed or blocked for
// authorizations on this card.
type CardUpdateParamsAuthorizationControlsMerchantCategoryCode struct {
	// The Merchant Category Codes that are allowed for authorizations on this card.
	// Authorizations with Merchant Category Codes not in this list will be declined.
	Allowed param.Field[[]CardUpdateParamsAuthorizationControlsMerchantCategoryCodeAllowed] `json:"allowed"`
	// The Merchant Category Codes that are blocked for authorizations on this card.
	// Authorizations with Merchant Category Codes in this list will be declined.
	Blocked param.Field[[]CardUpdateParamsAuthorizationControlsMerchantCategoryCodeBlocked] `json:"blocked"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantCategoryCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsMerchantCategoryCodeAllowed struct {
	// The Merchant Category Code.
	Code param.Field[string] `json:"code" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantCategoryCodeAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsMerchantCategoryCodeBlocked struct {
	// The Merchant Category Code.
	Code param.Field[string] `json:"code" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantCategoryCodeBlocked) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Restricts which merchant countries are allowed or blocked for authorizations on
// this card.
type CardUpdateParamsAuthorizationControlsMerchantCountry struct {
	// The merchant countries that are allowed for authorizations on this card.
	// Authorizations with merchant countries not in this list will be declined.
	Allowed param.Field[[]CardUpdateParamsAuthorizationControlsMerchantCountryAllowed] `json:"allowed"`
	// The merchant countries that are blocked for authorizations on this card.
	// Authorizations with merchant countries in this list will be declined.
	Blocked param.Field[[]CardUpdateParamsAuthorizationControlsMerchantCountryBlocked] `json:"blocked"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantCountry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsMerchantCountryAllowed struct {
	// The ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantCountryAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsMerchantCountryBlocked struct {
	// The ISO 3166-1 alpha-2 country code.
	Country param.Field[string] `json:"country" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsMerchantCountryBlocked) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how many times this card can be used.
type CardUpdateParamsAuthorizationControlsUsage struct {
	// Whether the card is for a single use or multiple uses.
	Category param.Field[CardUpdateParamsAuthorizationControlsUsageCategory] `json:"category" api:"required"`
	// Controls for multi-use cards. Required if and only if `category` is `multi_use`.
	MultiUse param.Field[CardUpdateParamsAuthorizationControlsUsageMultiUse] `json:"multi_use"`
	// Controls for single-use cards. Required if and only if `category` is
	// `single_use`.
	SingleUse param.Field[CardUpdateParamsAuthorizationControlsUsageSingleUse] `json:"single_use"`
}

func (r CardUpdateParamsAuthorizationControlsUsage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card is for a single use or multiple uses.
type CardUpdateParamsAuthorizationControlsUsageCategory string

const (
	CardUpdateParamsAuthorizationControlsUsageCategorySingleUse CardUpdateParamsAuthorizationControlsUsageCategory = "single_use"
	CardUpdateParamsAuthorizationControlsUsageCategoryMultiUse  CardUpdateParamsAuthorizationControlsUsageCategory = "multi_use"
)

func (r CardUpdateParamsAuthorizationControlsUsageCategory) IsKnown() bool {
	switch r {
	case CardUpdateParamsAuthorizationControlsUsageCategorySingleUse, CardUpdateParamsAuthorizationControlsUsageCategoryMultiUse:
		return true
	}
	return false
}

// Controls for multi-use cards. Required if and only if `category` is `multi_use`.
type CardUpdateParamsAuthorizationControlsUsageMultiUse struct {
	// Spending limits for this card. The most restrictive limit applies if multiple
	// limits match.
	SpendingLimits param.Field[[]CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimit] `json:"spending_limits"`
}

func (r CardUpdateParamsAuthorizationControlsUsageMultiUse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimit struct {
	// The interval at which the spending limit is enforced.
	Interval param.Field[CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval] `json:"interval" api:"required"`
	// The maximum settlement amount permitted in the given interval.
	SettlementAmount param.Field[int64] `json:"settlement_amount" api:"required"`
	// The Merchant Category Codes this spending limit applies to. If not set, the
	// limit applies to all transactions.
	MerchantCategoryCodes param.Field[[]CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode] `json:"merchant_category_codes"`
}

func (r CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interval at which the spending limit is enforced.
type CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval string

const (
	CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalAllTime        CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "all_time"
	CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerTransaction CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_transaction"
	CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerDay         CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_day"
	CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerWeek        CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_week"
	CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerMonth       CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval = "per_month"
)

func (r CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsInterval) IsKnown() bool {
	switch r {
	case CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalAllTime, CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerTransaction, CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerDay, CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerWeek, CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsIntervalPerMonth:
		return true
	}
	return false
}

type CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode struct {
	// The Merchant Category Code.
	Code param.Field[string] `json:"code" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsUsageMultiUseSpendingLimitsMerchantCategoryCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls for single-use cards. Required if and only if `category` is
// `single_use`.
type CardUpdateParamsAuthorizationControlsUsageSingleUse struct {
	// The settlement amount constraint for this single-use card.
	SettlementAmount param.Field[CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmount] `json:"settlement_amount" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsUsageSingleUse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settlement amount constraint for this single-use card.
type CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmount struct {
	// The operator used to compare the settlement amount.
	Comparison param.Field[CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison] `json:"comparison" api:"required"`
	// The settlement amount value.
	Value param.Field[int64] `json:"value" api:"required"`
}

func (r CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator used to compare the settlement amount.
type CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison string

const (
	CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonEquals           CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison = "equals"
	CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonLessThanOrEquals CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison = "less_than_or_equals"
)

func (r CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparison) IsKnown() bool {
	switch r {
	case CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonEquals, CardUpdateParamsAuthorizationControlsUsageSingleUseSettlementAmountComparisonLessThanOrEquals:
		return true
	}
	return false
}

// The card's updated billing address.
type CardUpdateParamsBillingAddress struct {
	// The city of the billing address.
	City param.Field[string] `json:"city" api:"required"`
	// The first line of the billing address.
	Line1 param.Field[string] `json:"line1" api:"required"`
	// The postal code of the billing address.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// The US state of the billing address.
	State param.Field[string] `json:"state" api:"required"`
	// The second line of the billing address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CardUpdateParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
type CardUpdateParamsDigitalWallet struct {
	// The digital card profile assigned to this digital card.
	DigitalCardProfileID param.Field[string] `json:"digital_card_profile_id"`
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email param.Field[string] `json:"email" format:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone param.Field[string] `json:"phone"`
}

func (r CardUpdateParamsDigitalWallet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status to update the Card with.
type CardUpdateParamsStatus string

const (
	CardUpdateParamsStatusActive   CardUpdateParamsStatus = "active"
	CardUpdateParamsStatusDisabled CardUpdateParamsStatus = "disabled"
	CardUpdateParamsStatusCanceled CardUpdateParamsStatus = "canceled"
)

func (r CardUpdateParamsStatus) IsKnown() bool {
	switch r {
	case CardUpdateParamsStatusActive, CardUpdateParamsStatusDisabled, CardUpdateParamsStatusCanceled:
		return true
	}
	return false
}

type CardListParams struct {
	// Filter Cards to ones belonging to the specified Account.
	AccountID param.Field[string]                  `query:"account_id"`
	CreatedAt param.Field[CardListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                `query:"limit"`
	Status param.Field[CardListParamsStatus] `query:"status"`
}

// URLQuery serializes [CardListParams]'s query parameters as `url.Values`.
func (r CardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardListParamsCreatedAt struct {
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

// URLQuery serializes [CardListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CardListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardListParamsStatus struct {
	// Filter Cards by status. For GET requests, this should be encoded as a
	// comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]CardListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CardListParamsStatus]'s query parameters as `url.Values`.
func (r CardListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardListParamsStatusIn string

const (
	CardListParamsStatusInActive   CardListParamsStatusIn = "active"
	CardListParamsStatusInDisabled CardListParamsStatusIn = "disabled"
	CardListParamsStatusInCanceled CardListParamsStatusIn = "canceled"
)

func (r CardListParamsStatusIn) IsKnown() bool {
	switch r {
	case CardListParamsStatusInActive, CardListParamsStatusInDisabled, CardListParamsStatusInCanceled:
		return true
	}
	return false
}

type CardNewDetailsIframeParams struct {
	// The identifier of the Physical Card to create an iframe for. This will inform
	// the appearance of the card rendered in the iframe.
	PhysicalCardID param.Field[string] `json:"physical_card_id"`
}

func (r CardNewDetailsIframeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdatePinParams struct {
	// The 4-digit PIN for the card, for use with ATMs.
	Pin param.Field[string] `json:"pin" api:"required"`
}

func (r CardUpdatePinParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
