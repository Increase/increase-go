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
	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, cardID string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Cards
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *CardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create an iframe URL for a Card to display the card details. More details about
// styling and usage can be found in the
// [documentation](/documentation/embedded-card-component).
func (r *CardService) NewDetailsIframe(ctx context.Context, cardID string, body CardNewDetailsIframeParams, opts ...option.RequestOption) (res *CardIframeURL, err error) {
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
func (r *CardService) Details(ctx context.Context, cardID string, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s/details", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Card's PIN
func (r *CardService) UpdatePin(ctx context.Context, cardID string, body CardUpdatePinParams, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("cards/%s/update_pin", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cards are commercial credit cards. They'll immediately work for online purchases
// after you create them. All cards maintain a credit limit of 100% of the
// Account’s available balance at the time of transaction. Funds are deducted from
// the Account upon transaction settlement.
type Card struct {
	// The card identifier.
	ID string `json:"id,required"`
	// The identifier for the account this card belongs to.
	AccountID string `json:"account_id,required"`
	// The Card's billing address.
	BillingAddress CardBillingAddress `json:"billing_address,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The card's description for display purposes.
	Description string `json:"description,required,nullable"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet CardDigitalWallet `json:"digital_wallet,required,nullable"`
	// The identifier for the entity associated with this card.
	EntityID string `json:"entity_id,required,nullable"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 string `json:"last4,required"`
	// This indicates if payments can be made with the card.
	Status CardStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type        CardType               `json:"type,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        cardJSON               `json:"-"`
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	BillingAddress  apijson.Field
	CreatedAt       apijson.Field
	Description     apijson.Field
	DigitalWallet   apijson.Field
	EntityID        apijson.Field
	ExpirationMonth apijson.Field
	ExpirationYear  apijson.Field
	IdempotencyKey  apijson.Field
	Last4           apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardJSON) RawJSON() string {
	return r.raw
}

// The Card's billing address.
type CardBillingAddress struct {
	// The city of the billing address.
	City string `json:"city,required,nullable"`
	// The first line of the billing address.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the billing address.
	Line2 string `json:"line2,required,nullable"`
	// The postal code of the billing address.
	PostalCode string `json:"postal_code,required,nullable"`
	// The US state of the billing address.
	State string                 `json:"state,required,nullable"`
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
	DigitalCardProfileID string `json:"digital_card_profile_id,required,nullable"`
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email,required,nullable"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string                `json:"phone,required,nullable"`
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

// A list of Card objects.
type CardListResponse struct {
	// The contents of the list.
	Data []Card `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                 `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        cardListResponseJSON   `json:"-"`
}

// cardListResponseJSON contains the JSON metadata for the struct
// [CardListResponse]
type cardListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardListResponseJSON) RawJSON() string {
	return r.raw
}

type CardNewParams struct {
	// The Account the card should belong to.
	AccountID param.Field[string] `json:"account_id,required"`
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

// The card's billing address.
type CardNewParamsBillingAddress struct {
	// The city of the billing address.
	City param.Field[string] `json:"city,required"`
	// The first line of the billing address.
	Line1 param.Field[string] `json:"line1,required"`
	// The postal code of the billing address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state of the billing address.
	State param.Field[string] `json:"state,required"`
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

// The card's updated billing address.
type CardUpdateParamsBillingAddress struct {
	// The city of the billing address.
	City param.Field[string] `json:"city,required"`
	// The first line of the billing address.
	Line1 param.Field[string] `json:"line1,required"`
	// The postal code of the billing address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state of the billing address.
	State param.Field[string] `json:"state,required"`
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
	Pin param.Field[string] `json:"pin,required"`
}

func (r CardUpdatePinParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
