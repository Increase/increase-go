// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// CardDetailService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardDetailService] method instead.
type CardDetailService struct {
	Options []option.RequestOption
}

// NewCardDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardDetailService(opts ...option.RequestOption) (r *CardDetailService) {
	r = &CardDetailService{}
	r.Options = opts
	return
}

// Update a Card's PIN
func (r *CardDetailService) Update(ctx context.Context, cardID string, body CardDetailUpdateParams, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s/details", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Create an iframe URL for a Card to display the card details. More details about
// styling and usage can be found in the
// [documentation](/documentation/embedded-card-component).
func (r *CardDetailService) NewDetailsIframe(ctx context.Context, cardID string, body CardDetailNewDetailsIframeParams, opts ...option.RequestOption) (res *CardIframeURL, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s/create_details_iframe", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sensitive details for a Card include the primary account number, expiry, card
// verification code, and PIN.
func (r *CardDetailService) Details(ctx context.Context, cardID string, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s/details", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// An object containing the sensitive details (card number, CVC, PIN, etc) for a
// Card. These details are not included in the Card object. If you'd prefer to
// never access these details directly, you can use the
// [embedded iframe](/documentation/embedded-card-component) to display the
// information to your users.
type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID string `json:"card_id,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// The 4-digit PIN for the card, for use with ATMs.
	Pin string `json:"pin,required"`
	// The card number.
	PrimaryAccountNumber string `json:"primary_account_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type CardDetailsType `json:"type,required"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode string          `json:"verification_code,required"`
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
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The iframe URL for the Card. Treat this as an opaque URL.
	IframeURL string `json:"iframe_url,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_iframe_url`.
	Type CardIframeURLType `json:"type,required"`
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

type CardDetailUpdateParams struct {
	// The 4-digit PIN for the card, for use with ATMs.
	Pin param.Field[string] `json:"pin,required"`
}

func (r CardDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardDetailNewDetailsIframeParams struct {
	// The identifier of the Physical Card to retrieve details for.
	PhysicalCardID param.Field[string] `json:"physical_card_id"`
}

func (r CardDetailNewDetailsIframeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
