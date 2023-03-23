package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type Card struct {
	// The card identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the account this card belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The card's description for display purposes.
	Description fields.Field[string] `json:"description,required,nullable"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 fields.Field[string] `json:"last4,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth fields.Field[int64] `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear fields.Field[int64] `json:"expiration_year,required"`
	// This indicates if payments can be made with the card.
	Status fields.Field[CardStatus] `json:"status,required"`
	// The Card's billing address.
	BillingAddress fields.Field[CardBillingAddress] `json:"billing_address,required"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet fields.Field[CardDigitalWallet] `json:"digital_wallet,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type fields.Field[CardType] `json:"type,required"`
}

// MarshalJSON serializes Card into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Card) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Card) String() (result string) {
	return fmt.Sprintf("&Card{ID:%s AccountID:%s CreatedAt:%s Description:%s Last4:%s ExpirationMonth:%s ExpirationYear:%s Status:%s BillingAddress:%s DigitalWallet:%s Type:%s}", r.ID, r.AccountID, r.CreatedAt, r.Description, r.Last4, r.ExpirationMonth, r.ExpirationYear, r.Status, r.BillingAddress, r.DigitalWallet, r.Type)
}

type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

type CardBillingAddress struct {
	// The first line of the billing address.
	Line1 fields.Field[string] `json:"line1,required,nullable"`
	// The second line of the billing address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the billing address.
	City fields.Field[string] `json:"city,required,nullable"`
	// The US state of the billing address.
	State fields.Field[string] `json:"state,required,nullable"`
	// The postal code of the billing address.
	PostalCode fields.Field[string] `json:"postal_code,required,nullable"`
}

// MarshalJSON serializes CardBillingAddress into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardBillingAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardBillingAddress) String() (result string) {
	return fmt.Sprintf("&CardBillingAddress{Line1:%s Line2:%s City:%s State:%s PostalCode:%s}", r.Line1, r.Line2, r.City, r.State, r.PostalCode)
}

type CardDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email fields.Field[string] `json:"email,required,nullable"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone fields.Field[string] `json:"phone,required,nullable"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID fields.Field[string] `json:"card_profile_id,required,nullable"`
}

// MarshalJSON serializes CardDigitalWallet into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDigitalWallet) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardDigitalWallet) String() (result string) {
	return fmt.Sprintf("&CardDigitalWallet{Email:%s Phone:%s CardProfileID:%s}", r.Email, r.Phone, r.CardProfileID)
}

type CardType string

const (
	CardTypeCard CardType = "card"
)

type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID fields.Field[string] `json:"card_id,required"`
	// The card number.
	PrimaryAccountNumber fields.Field[string] `json:"primary_account_number,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth fields.Field[int64] `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear fields.Field[int64] `json:"expiration_year,required"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode fields.Field[string] `json:"verification_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type fields.Field[CardDetailsType] `json:"type,required"`
}

// MarshalJSON serializes CardDetails into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardDetails) String() (result string) {
	return fmt.Sprintf("&CardDetails{CardID:%s PrimaryAccountNumber:%s ExpirationMonth:%s ExpirationYear:%s VerificationCode:%s Type:%s}", r.CardID, r.PrimaryAccountNumber, r.ExpirationMonth, r.ExpirationYear, r.VerificationCode, r.Type)
}

type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
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
