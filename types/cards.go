package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type Card struct {
	// The card identifier.
	ID *string `pjson:"id"`
	// The identifier for the account this card belongs to.
	AccountID *string `pjson:"account_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt *string `pjson:"created_at"`
	// The card's description for display purposes.
	Description *string `pjson:"description"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 *string `pjson:"last4"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth *int64 `pjson:"expiration_month"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear *int64 `pjson:"expiration_year"`
	// This indicates if payments can be made with the card.
	Status *CardStatus `pjson:"status"`
	// The Card's billing address.
	BillingAddress *CardBillingAddress `pjson:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CardDigitalWallet `pjson:"digital_wallet"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type       *CardType              `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Card using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Card into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Card) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The card identifier.
func (r Card) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the account this card belongs to.
func (r Card) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card was created.
func (r Card) GetCreatedAt() (CreatedAt string) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The card's description for display purposes.
func (r Card) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The last 4 digits of the Card's Primary Account Number.
func (r Card) GetLast4() (Last4 string) {
	if r.Last4 != nil {
		Last4 = *r.Last4
	}
	return
}

// The month the card expires in M format (e.g., August is 8).
func (r Card) GetExpirationMonth() (ExpirationMonth int64) {
	if r.ExpirationMonth != nil {
		ExpirationMonth = *r.ExpirationMonth
	}
	return
}

// The year the card expires in YYYY format (e.g., 2025).
func (r Card) GetExpirationYear() (ExpirationYear int64) {
	if r.ExpirationYear != nil {
		ExpirationYear = *r.ExpirationYear
	}
	return
}

// This indicates if payments can be made with the card.
func (r Card) GetStatus() (Status CardStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// The Card's billing address.
func (r Card) GetBillingAddress() (BillingAddress CardBillingAddress) {
	if r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r Card) GetDigitalWallet() (DigitalWallet CardDigitalWallet) {
	if r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card`.
func (r Card) GetType() (Type CardType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Card) String() (result string) {
	return fmt.Sprintf("&Card{ID:%s AccountID:%s CreatedAt:%s Description:%s Last4:%s ExpirationMonth:%s ExpirationYear:%s Status:%s BillingAddress:%s DigitalWallet:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.AccountID), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.Last4), core.FmtP(r.ExpirationMonth), core.FmtP(r.ExpirationYear), core.FmtP(r.Status), r.BillingAddress, r.DigitalWallet, core.FmtP(r.Type))
}

type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

type CardBillingAddress struct {
	// The first line of the billing address.
	Line1 *string `pjson:"line1"`
	// The second line of the billing address.
	Line2 *string `pjson:"line2"`
	// The city of the billing address.
	City *string `pjson:"city"`
	// The US state of the billing address.
	State *string `pjson:"state"`
	// The postal code of the billing address.
	PostalCode *string                `pjson:"postal_code"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardBillingAddress using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardBillingAddress into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardBillingAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The first line of the billing address.
func (r CardBillingAddress) GetLine1() (Line1 string) {
	if r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r CardBillingAddress) GetLine2() (Line2 string) {
	if r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r CardBillingAddress) GetCity() (City string) {
	if r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r CardBillingAddress) GetState() (State string) {
	if r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r CardBillingAddress) GetPostalCode() (PostalCode string) {
	if r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}

func (r CardBillingAddress) String() (result string) {
	return fmt.Sprintf("&CardBillingAddress{Line1:%s Line2:%s City:%s State:%s PostalCode:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.PostalCode))
}

type CardDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `pjson:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `pjson:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID *string                `pjson:"card_profile_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDigitalWallet using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDigitalWallet into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDigitalWallet) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r CardDigitalWallet) GetEmail() (Email string) {
	if r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r CardDigitalWallet) GetPhone() (Phone string) {
	if r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r CardDigitalWallet) GetCardProfileID() (CardProfileID string) {
	if r.CardProfileID != nil {
		CardProfileID = *r.CardProfileID
	}
	return
}

func (r CardDigitalWallet) String() (result string) {
	return fmt.Sprintf("&CardDigitalWallet{Email:%s Phone:%s CardProfileID:%s}", core.FmtP(r.Email), core.FmtP(r.Phone), core.FmtP(r.CardProfileID))
}

type CardType string

const (
	CardTypeCard CardType = "card"
)

type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID *string `pjson:"card_id"`
	// The card number.
	PrimaryAccountNumber *string `pjson:"primary_account_number"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth *int64 `pjson:"expiration_month"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear *int64 `pjson:"expiration_year"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode *string `pjson:"verification_code"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type       *CardDetailsType       `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDetails using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDetails into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Card for which sensitive details have been returned.
func (r CardDetails) GetCardID() (CardID string) {
	if r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The card number.
func (r CardDetails) GetPrimaryAccountNumber() (PrimaryAccountNumber string) {
	if r.PrimaryAccountNumber != nil {
		PrimaryAccountNumber = *r.PrimaryAccountNumber
	}
	return
}

// The month the card expires in M format (e.g., August is 8).
func (r CardDetails) GetExpirationMonth() (ExpirationMonth int64) {
	if r.ExpirationMonth != nil {
		ExpirationMonth = *r.ExpirationMonth
	}
	return
}

// The year the card expires in YYYY format (e.g., 2025).
func (r CardDetails) GetExpirationYear() (ExpirationYear int64) {
	if r.ExpirationYear != nil {
		ExpirationYear = *r.ExpirationYear
	}
	return
}

// The three-digit verification code for the card. It's also known as the Card
// Verification Code (CVC), the Card Verification Value (CVV), or the Card
// Identification (CID).
func (r CardDetails) GetVerificationCode() (VerificationCode string) {
	if r.VerificationCode != nil {
		VerificationCode = *r.VerificationCode
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_details`.
func (r CardDetails) GetType() (Type CardDetailsType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CardDetails) String() (result string) {
	return fmt.Sprintf("&CardDetails{CardID:%s PrimaryAccountNumber:%s ExpirationMonth:%s ExpirationYear:%s VerificationCode:%s Type:%s}", core.FmtP(r.CardID), core.FmtP(r.PrimaryAccountNumber), core.FmtP(r.ExpirationMonth), core.FmtP(r.ExpirationYear), core.FmtP(r.VerificationCode), core.FmtP(r.Type))
}

type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

type CreateACardParameters struct {
	// The Account the card should belong to.
	AccountID *string `pjson:"account_id"`
	// The description you choose to give the card.
	Description *string `pjson:"description"`
	// The card's billing address.
	BillingAddress *CreateACardParametersBillingAddress `pjson:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CreateACardParametersDigitalWallet `pjson:"digital_wallet"`
	jsonFields    map[string]interface{}              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateACardParameters using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CreateACardParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account the card should belong to.
func (r CreateACardParameters) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The description you choose to give the card.
func (r CreateACardParameters) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The card's billing address.
func (r CreateACardParameters) GetBillingAddress() (BillingAddress CreateACardParametersBillingAddress) {
	if r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r CreateACardParameters) GetDigitalWallet() (DigitalWallet CreateACardParametersDigitalWallet) {
	if r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

func (r CreateACardParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardParameters{AccountID:%s Description:%s BillingAddress:%s DigitalWallet:%s}", core.FmtP(r.AccountID), core.FmtP(r.Description), r.BillingAddress, r.DigitalWallet)
}

type CreateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 *string `pjson:"line1"`
	// The second line of the billing address.
	Line2 *string `pjson:"line2"`
	// The city of the billing address.
	City *string `pjson:"city"`
	// The US state of the billing address.
	State *string `pjson:"state"`
	// The postal code of the billing address.
	PostalCode *string                `pjson:"postal_code"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateACardParametersBillingAddress using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateACardParametersBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardParametersBillingAddress into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateACardParametersBillingAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The first line of the billing address.
func (r CreateACardParametersBillingAddress) GetLine1() (Line1 string) {
	if r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r CreateACardParametersBillingAddress) GetLine2() (Line2 string) {
	if r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r CreateACardParametersBillingAddress) GetCity() (City string) {
	if r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r CreateACardParametersBillingAddress) GetState() (State string) {
	if r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r CreateACardParametersBillingAddress) GetPostalCode() (PostalCode string) {
	if r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}

func (r CreateACardParametersBillingAddress) String() (result string) {
	return fmt.Sprintf("&CreateACardParametersBillingAddress{Line1:%s Line2:%s City:%s State:%s PostalCode:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.PostalCode))
}

type CreateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `pjson:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `pjson:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID *string                `pjson:"card_profile_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateACardParametersDigitalWallet using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateACardParametersDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardParametersDigitalWallet into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateACardParametersDigitalWallet) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r CreateACardParametersDigitalWallet) GetEmail() (Email string) {
	if r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r CreateACardParametersDigitalWallet) GetPhone() (Phone string) {
	if r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r CreateACardParametersDigitalWallet) GetCardProfileID() (CardProfileID string) {
	if r.CardProfileID != nil {
		CardProfileID = *r.CardProfileID
	}
	return
}

func (r CreateACardParametersDigitalWallet) String() (result string) {
	return fmt.Sprintf("&CreateACardParametersDigitalWallet{Email:%s Phone:%s CardProfileID:%s}", core.FmtP(r.Email), core.FmtP(r.Phone), core.FmtP(r.CardProfileID))
}

type UpdateACardParameters struct {
	// The description you choose to give the card.
	Description *string `pjson:"description"`
	// The status to update the Card with.
	Status *UpdateACardParametersStatus `pjson:"status"`
	// The card's updated billing address.
	BillingAddress *UpdateACardParametersBillingAddress `pjson:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *UpdateACardParametersDigitalWallet `pjson:"digital_wallet"`
	jsonFields    map[string]interface{}              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into UpdateACardParameters using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *UpdateACardParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateACardParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The description you choose to give the card.
func (r UpdateACardParameters) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The status to update the Card with.
func (r UpdateACardParameters) GetStatus() (Status UpdateACardParametersStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// The card's updated billing address.
func (r UpdateACardParameters) GetBillingAddress() (BillingAddress UpdateACardParametersBillingAddress) {
	if r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r UpdateACardParameters) GetDigitalWallet() (DigitalWallet UpdateACardParametersDigitalWallet) {
	if r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

func (r UpdateACardParameters) String() (result string) {
	return fmt.Sprintf("&UpdateACardParameters{Description:%s Status:%s BillingAddress:%s DigitalWallet:%s}", core.FmtP(r.Description), core.FmtP(r.Status), r.BillingAddress, r.DigitalWallet)
}

type UpdateACardParametersStatus string

const (
	UpdateACardParametersStatusActive   UpdateACardParametersStatus = "active"
	UpdateACardParametersStatusDisabled UpdateACardParametersStatus = "disabled"
	UpdateACardParametersStatusCanceled UpdateACardParametersStatus = "canceled"
)

type UpdateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 *string `pjson:"line1"`
	// The second line of the billing address.
	Line2 *string `pjson:"line2"`
	// The city of the billing address.
	City *string `pjson:"city"`
	// The US state of the billing address.
	State *string `pjson:"state"`
	// The postal code of the billing address.
	PostalCode *string                `pjson:"postal_code"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// UpdateACardParametersBillingAddress using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *UpdateACardParametersBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateACardParametersBillingAddress into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *UpdateACardParametersBillingAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The first line of the billing address.
func (r UpdateACardParametersBillingAddress) GetLine1() (Line1 string) {
	if r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r UpdateACardParametersBillingAddress) GetLine2() (Line2 string) {
	if r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r UpdateACardParametersBillingAddress) GetCity() (City string) {
	if r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r UpdateACardParametersBillingAddress) GetState() (State string) {
	if r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r UpdateACardParametersBillingAddress) GetPostalCode() (PostalCode string) {
	if r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}

func (r UpdateACardParametersBillingAddress) String() (result string) {
	return fmt.Sprintf("&UpdateACardParametersBillingAddress{Line1:%s Line2:%s City:%s State:%s PostalCode:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.PostalCode))
}

type UpdateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `pjson:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `pjson:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID *string                `pjson:"card_profile_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// UpdateACardParametersDigitalWallet using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *UpdateACardParametersDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateACardParametersDigitalWallet into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *UpdateACardParametersDigitalWallet) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r UpdateACardParametersDigitalWallet) GetEmail() (Email string) {
	if r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r UpdateACardParametersDigitalWallet) GetPhone() (Phone string) {
	if r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r UpdateACardParametersDigitalWallet) GetCardProfileID() (CardProfileID string) {
	if r.CardProfileID != nil {
		CardProfileID = *r.CardProfileID
	}
	return
}

func (r UpdateACardParametersDigitalWallet) String() (result string) {
	return fmt.Sprintf("&UpdateACardParametersDigitalWallet{Email:%s Phone:%s CardProfileID:%s}", core.FmtP(r.Email), core.FmtP(r.Phone), core.FmtP(r.CardProfileID))
}

type CardListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Cards to ones belonging to the specified Account.
	AccountID  *string                   `query:"account_id"`
	CreatedAt  *CardsListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardListParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardListParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardListParams into a url.Values of the query parameters
// associated with this value
func (r *CardListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r CardListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r CardListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Cards to ones belonging to the specified Account.
func (r CardListParams) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r CardListParams) GetCreatedAt() (CreatedAt CardsListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r CardListParams) String() (result string) {
	return fmt.Sprintf("&CardListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), r.CreatedAt)
}

type CardsListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardsListParamsCreatedAt
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardsListParamsCreatedAt into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardsListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *CardsListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r CardsListParamsCreatedAt) GetAfter() (After string) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r CardsListParamsCreatedAt) GetBefore() (Before string) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r CardsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r CardsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r CardsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CardsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type CardList struct {
	// The contents of the list.
	Data *[]Card `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardList using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CardList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardList into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *CardList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardList into a url.Values of the query parameters
// associated with this value
func (r *CardList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r CardList) GetData() (Data []Card) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r CardList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r CardList) String() (result string) {
	return fmt.Sprintf("&CardList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type CardsPage struct {
	*pagination.Page[Card]
}

func (r *CardsPage) Card() *Card {
	return r.Current()
}

func (r *CardsPage) NextPage() (*CardsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CardsPage{page}, nil
	}
}
