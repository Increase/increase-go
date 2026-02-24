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

// CardTokenService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardTokenService] method instead.
type CardTokenService struct {
	Options []option.RequestOption
}

// NewCardTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardTokenService(opts ...option.RequestOption) (r *CardTokenService) {
	r = &CardTokenService{}
	r.Options = opts
	return
}

// Retrieve a Card Token
func (r *CardTokenService) Get(ctx context.Context, cardTokenID string, opts ...option.RequestOption) (res *CardToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardTokenID == "" {
		err = errors.New("missing required card_token_id parameter")
		return
	}
	path := fmt.Sprintf("card_tokens/%s", cardTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Tokens
func (r *CardTokenService) List(ctx context.Context, query CardTokenListParams, opts ...option.RequestOption) (res *pagination.Page[CardToken], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_tokens"
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

// List Card Tokens
func (r *CardTokenService) ListAutoPaging(ctx context.Context, query CardTokenListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CardToken] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// The capabilities of a Card Token describe whether the card can be used for
// specific operations, such as Card Push Transfers. The capabilities can change
// over time based on the issuing bank's configuration of the card range.
func (r *CardTokenService) Capabilities(ctx context.Context, cardTokenID string, opts ...option.RequestOption) (res *CardTokenCapabilities, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardTokenID == "" {
		err = errors.New("missing required card_token_id parameter")
		return
	}
	path := fmt.Sprintf("card_tokens/%s/capabilities", cardTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Card Tokens represent a tokenized card number that can be used for Card Push
// Transfers and Card Validations.
type CardToken struct {
	// The Card Token's identifier.
	ID string `json:"id" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the card token was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the card
	// expires.
	ExpirationDate time.Time `json:"expiration_date" api:"required" format:"date"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4" api:"required"`
	// The length of the card number.
	Length int64 `json:"length" api:"required"`
	// The prefix of the card number, usually 8 digits.
	Prefix string `json:"prefix" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_token`.
	Type CardTokenType `json:"type" api:"required"`
	JSON cardTokenJSON `json:"-"`
}

// cardTokenJSON contains the JSON metadata for the struct [CardToken]
type cardTokenJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	ExpirationDate apijson.Field
	Last4          apijson.Field
	Length         apijson.Field
	Prefix         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardTokenJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_token`.
type CardTokenType string

const (
	CardTokenTypeCardToken CardTokenType = "card_token"
)

func (r CardTokenType) IsKnown() bool {
	switch r {
	case CardTokenTypeCardToken:
		return true
	}
	return false
}

// The capabilities of a Card Token describe whether the card can be used for
// specific operations, such as Card Push Transfers. The capabilities can change
// over time based on the issuing bank's configuration of the card range.
type CardTokenCapabilities struct {
	// Each route represent a path e.g., a push transfer can take.
	Routes []CardTokenCapabilitiesRoute `json:"routes" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_token_capabilities`.
	Type CardTokenCapabilitiesType `json:"type" api:"required"`
	JSON cardTokenCapabilitiesJSON `json:"-"`
}

// cardTokenCapabilitiesJSON contains the JSON metadata for the struct
// [CardTokenCapabilities]
type cardTokenCapabilitiesJSON struct {
	Routes      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardTokenCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardTokenCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type CardTokenCapabilitiesRoute struct {
	// Whether you can push funds to the card using cross-border Card Push Transfers.
	CrossBorderPushTransfers CardTokenCapabilitiesRoutesCrossBorderPushTransfers `json:"cross_border_push_transfers" api:"required"`
	// Whether you can push funds to the card using domestic Card Push Transfers.
	DomesticPushTransfers CardTokenCapabilitiesRoutesDomesticPushTransfers `json:"domestic_push_transfers" api:"required"`
	// The card network route the capabilities apply to.
	Route CardTokenCapabilitiesRoutesRoute `json:"route" api:"required"`
	JSON  cardTokenCapabilitiesRouteJSON   `json:"-"`
}

// cardTokenCapabilitiesRouteJSON contains the JSON metadata for the struct
// [CardTokenCapabilitiesRoute]
type cardTokenCapabilitiesRouteJSON struct {
	CrossBorderPushTransfers apijson.Field
	DomesticPushTransfers    apijson.Field
	Route                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardTokenCapabilitiesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardTokenCapabilitiesRouteJSON) RawJSON() string {
	return r.raw
}

// Whether you can push funds to the card using cross-border Card Push Transfers.
type CardTokenCapabilitiesRoutesCrossBorderPushTransfers string

const (
	CardTokenCapabilitiesRoutesCrossBorderPushTransfersSupported    CardTokenCapabilitiesRoutesCrossBorderPushTransfers = "supported"
	CardTokenCapabilitiesRoutesCrossBorderPushTransfersNotSupported CardTokenCapabilitiesRoutesCrossBorderPushTransfers = "not_supported"
)

func (r CardTokenCapabilitiesRoutesCrossBorderPushTransfers) IsKnown() bool {
	switch r {
	case CardTokenCapabilitiesRoutesCrossBorderPushTransfersSupported, CardTokenCapabilitiesRoutesCrossBorderPushTransfersNotSupported:
		return true
	}
	return false
}

// Whether you can push funds to the card using domestic Card Push Transfers.
type CardTokenCapabilitiesRoutesDomesticPushTransfers string

const (
	CardTokenCapabilitiesRoutesDomesticPushTransfersSupported    CardTokenCapabilitiesRoutesDomesticPushTransfers = "supported"
	CardTokenCapabilitiesRoutesDomesticPushTransfersNotSupported CardTokenCapabilitiesRoutesDomesticPushTransfers = "not_supported"
)

func (r CardTokenCapabilitiesRoutesDomesticPushTransfers) IsKnown() bool {
	switch r {
	case CardTokenCapabilitiesRoutesDomesticPushTransfersSupported, CardTokenCapabilitiesRoutesDomesticPushTransfersNotSupported:
		return true
	}
	return false
}

// The card network route the capabilities apply to.
type CardTokenCapabilitiesRoutesRoute string

const (
	CardTokenCapabilitiesRoutesRouteVisa       CardTokenCapabilitiesRoutesRoute = "visa"
	CardTokenCapabilitiesRoutesRouteMastercard CardTokenCapabilitiesRoutesRoute = "mastercard"
)

func (r CardTokenCapabilitiesRoutesRoute) IsKnown() bool {
	switch r {
	case CardTokenCapabilitiesRoutesRouteVisa, CardTokenCapabilitiesRoutesRouteMastercard:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_token_capabilities`.
type CardTokenCapabilitiesType string

const (
	CardTokenCapabilitiesTypeCardTokenCapabilities CardTokenCapabilitiesType = "card_token_capabilities"
)

func (r CardTokenCapabilitiesType) IsKnown() bool {
	switch r {
	case CardTokenCapabilitiesTypeCardTokenCapabilities:
		return true
	}
	return false
}

type CardTokenListParams struct {
	CreatedAt param.Field[CardTokenListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CardTokenListParams]'s query parameters as `url.Values`.
func (r CardTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardTokenListParamsCreatedAt struct {
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

// URLQuery serializes [CardTokenListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CardTokenListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
