// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// CardService contains methods and other services that help with interacting with
// the increase API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCardService] method instead.
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
	opts = append(r.Options[:], opts...)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, cardID string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", cardID)
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
func (r *CardService) GetSensitiveDetails(ctx context.Context, cardID string, opts ...option.RequestOption) (res *CardDetails, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/details", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 string `json:"last4,required"`
	// This indicates if payments can be made with the card.
	Status CardStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type CardType `json:"type,required"`
	JSON cardJSON
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	BillingAddress  apijson.Field
	CreatedAt       apijson.Field
	Description     apijson.Field
	DigitalWallet   apijson.Field
	ExpirationMonth apijson.Field
	ExpirationYear  apijson.Field
	Last4           apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	State string `json:"state,required,nullable"`
	JSON  cardBillingAddressJSON
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

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
type CardDigitalWallet struct {
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email,required,nullable"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string `json:"phone,required,nullable"`
	JSON  cardDigitalWalletJSON
}

// cardDigitalWalletJSON contains the JSON metadata for the struct
// [CardDigitalWallet]
type cardDigitalWalletJSON struct {
	CardProfileID apijson.Field
	Email         apijson.Field
	Phone         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This indicates if payments can be made with the card.
type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

// A constant representing the object's type. For this resource it will always be
// `card`.
type CardType string

const (
	CardTypeCard CardType = "card"
)

// An object containing the sensitive details (card number, cvc, etc) for a Card.
type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID string `json:"card_id,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// The card number.
	PrimaryAccountNumber string `json:"primary_account_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type CardDetailsType `json:"type,required"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode string `json:"verification_code,required"`
	JSON             cardDetailsJSON
}

// cardDetailsJSON contains the JSON metadata for the struct [CardDetails]
type cardDetailsJSON struct {
	CardID               apijson.Field
	ExpirationMonth      apijson.Field
	ExpirationYear       apijson.Field
	PrimaryAccountNumber apijson.Field
	Type                 apijson.Field
	VerificationCode     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `card_details`.
type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

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
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID param.Field[string] `json:"card_profile_id"`
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email param.Field[string] `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
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
	// The card profile assigned to this card. Card profiles may also be assigned at
	// the program level.
	CardProfileID param.Field[string] `json:"card_profile_id"`
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email param.Field[string] `json:"email"`
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

type CardListParams struct {
	// Filter Cards to ones belonging to the specified Account.
	AccountID param.Field[string]                  `query:"account_id"`
	CreatedAt param.Field[CardListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
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
