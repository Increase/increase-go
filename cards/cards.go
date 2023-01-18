package cards

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type CardService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardService(requester core.Requester) (r *CardService) {
	r = &CardService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type Card struct {
	// The card identifier.
	ID string `json:"id"`
	// The identifier for the account this card belongs to.
	AccountID string `json:"account_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt string `json:"created_at"`
	// The card's description for display purposes.
	Description *string `json:"description"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 string `json:"last4"`
	// The month the card expires in MM format (e.g., August is 08).
	ExpirationMonth string `json:"expiration_month"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear string `json:"expiration_year"`
	// This indicates if payments can be made with the card.
	Status *CardStatus `json:"status"`
	// The Card's billing address.
	BillingAddress CardBillingAddress `json:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CardDigitalWallet `json:"digital_wallet"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type CardType `json:"type"`
}

// The card's description for display purposes.
func (r *Card) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// This indicates if payments can be made with the card.
func (r *Card) GetStatus() (Status CardStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r *Card) GetDigitalWallet() (DigitalWallet CardDigitalWallet) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

//
type CardBillingAddress struct {
	// The first line of the billing address.
	Line1 *string `json:"line1"`
	// The second line of the billing address.
	Line2 *string `json:"line2"`
	// The city of the billing address.
	City *string `json:"city"`
	// The US state of the billing address.
	State *string `json:"state"`
	// The postal code of the billing address.
	PostalCode *string `json:"postal_code"`
}

// The first line of the billing address.
func (r *CardBillingAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r *CardBillingAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r *CardBillingAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r *CardBillingAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r *CardBillingAddress) GetPostalCode() (PostalCode string) {
	if r != nil && r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}

//
type CardDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `json:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `json:"phone"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID *string `json:"card_profile_id"`
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r *CardDigitalWallet) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r *CardDigitalWallet) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r *CardDigitalWallet) GetCardProfileID() (CardProfileID string) {
	if r != nil && r.CardProfileID != nil {
		CardProfileID = *r.CardProfileID
	}
	return
}

type CardType string

const (
	CardTypeCard CardType = "card"
)

//
type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID string `json:"card_id"`
	// The card number.
	PrimaryAccountNumber string `json:"primary_account_number"`
	// The month the card expires in MM format (e.g., August is 08).
	ExpirationMonth string `json:"expiration_month"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear string `json:"expiration_year"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode string `json:"verification_code"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type CardDetailsType `json:"type"`
}

type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

type CreateACardParameters struct {
	// The Account the card should belong to.
	AccountID string `json:"account_id"`
	// The description you choose to give the card.
	Description string `json:"description,omitempty"`
	// The card's billing address.
	BillingAddress CreateACardParametersBillingAddress `json:"billing_address,omitempty"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet CreateACardParametersDigitalWallet `json:"digital_wallet,omitempty"`
}

//
type CreateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 string `json:"line1"`
	// The second line of the billing address.
	Line2 string `json:"line2,omitempty"`
	// The city of the billing address.
	City string `json:"city"`
	// The US state of the billing address.
	State string `json:"state"`
	// The postal code of the billing address.
	PostalCode string `json:"postal_code"`
}

//
type CreateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email,omitempty"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string `json:"phone,omitempty"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID string `json:"card_profile_id,omitempty"`
}

type UpdateACardParameters struct {
	// The description you choose to give the card.
	Description string `json:"description,omitempty"`
	// The status to update the Card with.
	Status UpdateACardParametersStatus `json:"status,omitempty"`
	// The card's updated billing address.
	BillingAddress UpdateACardParametersBillingAddress `json:"billing_address,omitempty"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet UpdateACardParametersDigitalWallet `json:"digital_wallet,omitempty"`
}

type UpdateACardParametersStatus string

const (
	UpdateACardParametersStatusActive   UpdateACardParametersStatus = "active"
	UpdateACardParametersStatusDisabled UpdateACardParametersStatus = "disabled"
	UpdateACardParametersStatusCanceled UpdateACardParametersStatus = "canceled"
)

//
type UpdateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 string `json:"line1"`
	// The second line of the billing address.
	Line2 string `json:"line2,omitempty"`
	// The city of the billing address.
	City string `json:"city"`
	// The US state of the billing address.
	State string `json:"state"`
	// The postal code of the billing address.
	PostalCode string `json:"postal_code"`
}

//
type UpdateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email,omitempty"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string `json:"phone,omitempty"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID string `json:"card_profile_id,omitempty"`
}

type ListCardsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// Filter Cards to ones belonging to the specified Account.
	AccountID string                  `query:"account_id"`
	CreatedAt ListCardsQueryCreatedAt `query:"created_at"`
}

type ListCardsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore string `json:"on_or_before,omitempty"`
}

//
type CardList struct {
	// The contents of the list.
	Data []Card `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *CardList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *CardService) Create(ctx context.Context, body *CreateACardParameters, opts ...*core.RequestOpts) (res *Card, err error) {
	err = r.post(
		ctx,
		"/cards",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *CardService) Retrieve(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *Card, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/cards/%s", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *CardService) Update(ctx context.Context, card_id string, body *UpdateACardParameters, opts ...*core.RequestOpts) (res *Card, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/cards/%s", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

type CardsPage struct {
	*pagination.Page[Card]
}

func (r *CardsPage) Card() *Card {
	return r.Current()
}

func (r *CardsPage) GetNextPage() (*CardsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &CardsPage{page}, nil
	}
}

func (r *CardService) List(ctx context.Context, query *ListCardsQuery, opts ...*core.RequestOpts) (res *CardsPage, err error) {
	page := &CardsPage{
		Page: &pagination.Page[Card]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/cards",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *CardService) RetrieveSensitiveDetails(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *CardDetails, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/cards/%s/details", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}
