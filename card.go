package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type CardService struct {
	Options []option.RequestOption
}

func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	return
}

// Create a Card
func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, card_id string, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, card_id string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Cards
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *shared.Page[Card], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *CardService) ListAutoPaging(ctx context.Context, query CardListParams, opts ...option.RequestOption) *shared.PageAutoPager[Card] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve sensitive details for a Card
func (r *CardService) GetSensitiveDetails(ctx context.Context, card_id string, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/details", card_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Card struct {
	// The card identifier.
	ID string `json:"id,required"`
	// The identifier for the account this card belongs to.
	AccountID string `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The card's description for display purposes.
	Description string `json:"description,required,nullable"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 string `json:"last4,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// This indicates if payments can be made with the card.
	Status CardStatus `json:"status,required"`
	// The Card's billing address.
	BillingAddress CardBillingAddress `json:"billing_address,required"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet CardDigitalWallet `json:"digital_wallet,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type CardType `json:"type,required"`
	JSON CardJSON
}

type CardJSON struct {
	ID              apijson.Metadata
	AccountID       apijson.Metadata
	CreatedAt       apijson.Metadata
	Description     apijson.Metadata
	Last4           apijson.Metadata
	ExpirationMonth apijson.Metadata
	ExpirationYear  apijson.Metadata
	Status          apijson.Metadata
	BillingAddress  apijson.Metadata
	DigitalWallet   apijson.Metadata
	Type            apijson.Metadata
	raw             string
	Extras          map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Card using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

type CardBillingAddress struct {
	// The first line of the billing address.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the billing address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the billing address.
	City string `json:"city,required,nullable"`
	// The US state of the billing address.
	State string `json:"state,required,nullable"`
	// The postal code of the billing address.
	PostalCode string `json:"postal_code,required,nullable"`
	JSON       CardBillingAddressJSON
}

type CardBillingAddressJSON struct {
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	City       apijson.Metadata
	State      apijson.Metadata
	PostalCode apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardBillingAddress using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email,required,nullable"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string `json:"phone,required,nullable"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	JSON          CardDigitalWalletJSON
}

type CardDigitalWalletJSON struct {
	Email         apijson.Metadata
	Phone         apijson.Metadata
	CardProfileID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDigitalWallet using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardType string

const (
	CardTypeCard CardType = "card"
)

type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID string `json:"card_id,required"`
	// The card number.
	PrimaryAccountNumber string `json:"primary_account_number,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode string `json:"verification_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type CardDetailsType `json:"type,required"`
	JSON CardDetailsJSON
}

type CardDetailsJSON struct {
	CardID               apijson.Metadata
	PrimaryAccountNumber apijson.Metadata
	ExpirationMonth      apijson.Metadata
	ExpirationYear       apijson.Metadata
	VerificationCode     apijson.Metadata
	Type                 apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDetails using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

type CardNewParams struct {
	// The Account the card should belong to.
	AccountID field.Field[string] `json:"account_id,required"`
	// The description you choose to give the card.
	Description field.Field[string] `json:"description"`
	// The card's billing address.
	BillingAddress field.Field[CardNewParamsBillingAddress] `json:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. To add the card to a digital wallet, you may supply an email or phone
	// number with this request. Otherwise, subscribe and then action a Real Time
	// Decision with the category `digital_wallet_token_requested` or
	// `digital_wallet_authentication_requested`.
	DigitalWallet field.Field[CardNewParamsDigitalWallet] `json:"digital_wallet"`
}

// MarshalJSON serializes CardNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardNewParamsBillingAddress struct {
	// The first line of the billing address.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the billing address.
	Line2 field.Field[string] `json:"line2"`
	// The city of the billing address.
	City field.Field[string] `json:"city,required"`
	// The US state of the billing address.
	State field.Field[string] `json:"state,required"`
	// The postal code of the billing address.
	PostalCode field.Field[string] `json:"postal_code,required"`
}

type CardNewParamsDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email field.Field[string] `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone field.Field[string] `json:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID field.Field[string] `json:"card_profile_id"`
}

type CardUpdateParams struct {
	// The description you choose to give the card.
	Description field.Field[string] `json:"description"`
	// The status to update the Card with.
	Status field.Field[CardUpdateParamsStatus] `json:"status"`
	// The card's updated billing address.
	BillingAddress field.Field[CardUpdateParamsBillingAddress] `json:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet field.Field[CardUpdateParamsDigitalWallet] `json:"digital_wallet"`
}

// MarshalJSON serializes CardUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r CardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardUpdateParamsStatus string

const (
	CardUpdateParamsStatusActive   CardUpdateParamsStatus = "active"
	CardUpdateParamsStatusDisabled CardUpdateParamsStatus = "disabled"
	CardUpdateParamsStatusCanceled CardUpdateParamsStatus = "canceled"
)

type CardUpdateParamsBillingAddress struct {
	// The first line of the billing address.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the billing address.
	Line2 field.Field[string] `json:"line2"`
	// The city of the billing address.
	City field.Field[string] `json:"city,required"`
	// The US state of the billing address.
	State field.Field[string] `json:"state,required"`
	// The postal code of the billing address.
	PostalCode field.Field[string] `json:"postal_code,required"`
}

type CardUpdateParamsDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email field.Field[string] `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone field.Field[string] `json:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID field.Field[string] `json:"card_profile_id"`
}

type CardListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Cards to ones belonging to the specified Account.
	AccountID field.Field[string]                  `query:"account_id"`
	CreatedAt field.Field[CardListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CardListParams into a url.Values of the query parameters
// associated with this value
func (r CardListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CardListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes CardListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r CardListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CardListResponse struct {
	// The contents of the list.
	Data []Card `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CardListResponseJSON
}

type CardListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
