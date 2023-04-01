package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateACardParameters struct {
	// The Account the card should belong to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The description you choose to give the card.
	Description fields.Field[string] `json:"description"`
	// The card's billing address.
	BillingAddress fields.Field[CreateACardParametersBillingAddress] `json:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet fields.Field[CreateACardParametersDigitalWallet] `json:"digital_wallet"`
}

// MarshalJSON serializes CreateACardParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACardParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardParameters{AccountID:%s Description:%s BillingAddress:%s DigitalWallet:%s}", r.AccountID, r.Description, r.BillingAddress, r.DigitalWallet)
}

type CreateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the billing address.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the billing address.
	City fields.Field[string] `json:"city,required"`
	// The US state of the billing address.
	State fields.Field[string] `json:"state,required"`
	// The postal code of the billing address.
	PostalCode fields.Field[string] `json:"postal_code,required"`
}

func (r CreateACardParametersBillingAddress) String() (result string) {
	return fmt.Sprintf("&CreateACardParametersBillingAddress{Line1:%s Line2:%s City:%s State:%s PostalCode:%s}", r.Line1, r.Line2, r.City, r.State, r.PostalCode)
}

type CreateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email fields.Field[string] `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone fields.Field[string] `json:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID fields.Field[string] `json:"card_profile_id"`
}

func (r CreateACardParametersDigitalWallet) String() (result string) {
	return fmt.Sprintf("&CreateACardParametersDigitalWallet{Email:%s Phone:%s CardProfileID:%s}", r.Email, r.Phone, r.CardProfileID)
}

type UpdateACardParameters struct {
	// The description you choose to give the card.
	Description fields.Field[string] `json:"description"`
	// The status to update the Card with.
	Status fields.Field[UpdateACardParametersStatus] `json:"status"`
	// The card's updated billing address.
	BillingAddress fields.Field[UpdateACardParametersBillingAddress] `json:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet fields.Field[UpdateACardParametersDigitalWallet] `json:"digital_wallet"`
}

// MarshalJSON serializes UpdateACardParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateACardParameters) String() (result string) {
	return fmt.Sprintf("&UpdateACardParameters{Description:%s Status:%s BillingAddress:%s DigitalWallet:%s}", r.Description, r.Status, r.BillingAddress, r.DigitalWallet)
}

type UpdateACardParametersStatus string

const (
	UpdateACardParametersStatusActive   UpdateACardParametersStatus = "active"
	UpdateACardParametersStatusDisabled UpdateACardParametersStatus = "disabled"
	UpdateACardParametersStatusCanceled UpdateACardParametersStatus = "canceled"
)

type UpdateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the billing address.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the billing address.
	City fields.Field[string] `json:"city,required"`
	// The US state of the billing address.
	State fields.Field[string] `json:"state,required"`
	// The postal code of the billing address.
	PostalCode fields.Field[string] `json:"postal_code,required"`
}

func (r UpdateACardParametersBillingAddress) String() (result string) {
	return fmt.Sprintf("&UpdateACardParametersBillingAddress{Line1:%s Line2:%s City:%s State:%s PostalCode:%s}", r.Line1, r.Line2, r.City, r.State, r.PostalCode)
}

type UpdateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email fields.Field[string] `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone fields.Field[string] `json:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID fields.Field[string] `json:"card_profile_id"`
}

func (r UpdateACardParametersDigitalWallet) String() (result string) {
	return fmt.Sprintf("&UpdateACardParametersDigitalWallet{Email:%s Phone:%s CardProfileID:%s}", r.Email, r.Phone, r.CardProfileID)
}

type CardListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Cards to ones belonging to the specified Account.
	AccountID fields.Field[string]                  `query:"account_id"`
	CreatedAt fields.Field[CardListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CardListParams into a url.Values of the query parameters
// associated with this value
func (r *CardListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardListParams) String() (result string) {
	return fmt.Sprintf("&CardListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type CardListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes CardListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *CardListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CardListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}
