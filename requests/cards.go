package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
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
	return query.Marshal(r)
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
	return query.Marshal(r)
}
