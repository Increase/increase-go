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

type PreloadedCardService struct {
	Cards *CardService
}

func (r *PreloadedCardService) Init(service *CardService) {
	r.Cards = service
}

func NewPreloadedCardService(service *CardService) (r *PreloadedCardService) {
	r = &PreloadedCardService{}
	r.Init(service)
	return
}

//
type Card struct {
	// The card identifier.
	Id *string `json:"id"`
	// The identifier for the account this card belongs to.
	AccountId *string `json:"account_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt *string `json:"created_at"`
	// The card's description for display purposes.
	Description *string `json:"description"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 *string `json:"last4"`
	// The month the card expires in MM format (e.g., August is 08).
	ExpirationMonth *string `json:"expiration_month"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear *string `json:"expiration_year"`
	// This indicates if payments can be made with the card.
	Status *CardStatus `json:"status"`
	// The Card's billing address.
	BillingAddress *CardBillingAddress `json:"billing_address"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CardDigitalWallet `json:"digital_wallet"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type *CardType `json:"type"`
}

// The card identifier.
func (r *Card) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the account this card belongs to.
func (r *Card) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card was created.
func (r *Card) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The card's description for display purposes.
func (r *Card) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The last 4 digits of the Card's Primary Account Number.
func (r *Card) GetLast4() (Last4 string) {
	if r != nil && r.Last4 != nil {
		Last4 = *r.Last4
	}
	return
}

// The month the card expires in MM format (e.g., August is 08).
func (r *Card) GetExpirationMonth() (ExpirationMonth string) {
	if r != nil && r.ExpirationMonth != nil {
		ExpirationMonth = *r.ExpirationMonth
	}
	return
}

// The year the card expires in YYYY format (e.g., 2025).
func (r *Card) GetExpirationYear() (ExpirationYear string) {
	if r != nil && r.ExpirationYear != nil {
		ExpirationYear = *r.ExpirationYear
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

// The Card's billing address.
func (r *Card) GetBillingAddress() (BillingAddress CardBillingAddress) {
	if r != nil && r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
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

// A constant representing the object's type. For this resource it will always be
// `card`.
func (r *Card) GetType() (Type CardType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
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
	CardProfileId *string `json:"card_profile_id"`
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
func (r *CardDigitalWallet) GetCardProfileId() (CardProfileId string) {
	if r != nil && r.CardProfileId != nil {
		CardProfileId = *r.CardProfileId
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
	CardId *string `json:"card_id"`
	// The card number.
	PrimaryAccountNumber *string `json:"primary_account_number"`
	// The month the card expires in MM format (e.g., August is 08).
	ExpirationMonth *string `json:"expiration_month"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear *string `json:"expiration_year"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode *string `json:"verification_code"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type *CardDetailsType `json:"type"`
}

// The identifier for the Card for which sensitive details have been returned.
func (r *CardDetails) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The card number.
func (r *CardDetails) GetPrimaryAccountNumber() (PrimaryAccountNumber string) {
	if r != nil && r.PrimaryAccountNumber != nil {
		PrimaryAccountNumber = *r.PrimaryAccountNumber
	}
	return
}

// The month the card expires in MM format (e.g., August is 08).
func (r *CardDetails) GetExpirationMonth() (ExpirationMonth string) {
	if r != nil && r.ExpirationMonth != nil {
		ExpirationMonth = *r.ExpirationMonth
	}
	return
}

// The year the card expires in YYYY format (e.g., 2025).
func (r *CardDetails) GetExpirationYear() (ExpirationYear string) {
	if r != nil && r.ExpirationYear != nil {
		ExpirationYear = *r.ExpirationYear
	}
	return
}

// The three-digit verification code for the card. It's also known as the Card
// Verification Code (CVC), the Card Verification Value (CVV), or the Card
// Identification (CID).
func (r *CardDetails) GetVerificationCode() (VerificationCode string) {
	if r != nil && r.VerificationCode != nil {
		VerificationCode = *r.VerificationCode
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_details`.
func (r *CardDetails) GetType() (Type CardDetailsType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

type CreateACardParameters struct {
	// The Account the card should belong to.
	AccountId *string `json:"account_id"`
	// The description you choose to give the card.
	Description *string `json:"description,omitempty"`
	// The card's billing address.
	BillingAddress *CreateACardParametersBillingAddress `json:"billing_address,omitempty"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CreateACardParametersDigitalWallet `json:"digital_wallet,omitempty"`
}

// The Account the card should belong to.
func (r *CreateACardParameters) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The description you choose to give the card.
func (r *CreateACardParameters) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The card's billing address.
func (r *CreateACardParameters) GetBillingAddress() (BillingAddress CreateACardParametersBillingAddress) {
	if r != nil && r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r *CreateACardParameters) GetDigitalWallet() (DigitalWallet CreateACardParametersDigitalWallet) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

//
type CreateACardParametersBillingAddress struct {
	// The first line of the billing address.
	Line1 *string `json:"line1"`
	// The second line of the billing address.
	Line2 *string `json:"line2,omitempty"`
	// The city of the billing address.
	City *string `json:"city"`
	// The US state of the billing address.
	State *string `json:"state"`
	// The postal code of the billing address.
	PostalCode *string `json:"postal_code"`
}

// The first line of the billing address.
func (r *CreateACardParametersBillingAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r *CreateACardParametersBillingAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r *CreateACardParametersBillingAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r *CreateACardParametersBillingAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r *CreateACardParametersBillingAddress) GetPostalCode() (PostalCode string) {
	if r != nil && r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}

//
type CreateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `json:"email,omitempty"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `json:"phone,omitempty"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileId *string `json:"card_profile_id,omitempty"`
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r *CreateACardParametersDigitalWallet) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r *CreateACardParametersDigitalWallet) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r *CreateACardParametersDigitalWallet) GetCardProfileId() (CardProfileId string) {
	if r != nil && r.CardProfileId != nil {
		CardProfileId = *r.CardProfileId
	}
	return
}

type UpdateACardParameters struct {
	// The description you choose to give the card.
	Description *string `json:"description,omitempty"`
	// The status to update the Card with.
	Status *UpdateACardParametersStatus `json:"status,omitempty"`
	// The card's updated billing address.
	BillingAddress *UpdateACardParametersBillingAddress `json:"billing_address,omitempty"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *UpdateACardParametersDigitalWallet `json:"digital_wallet,omitempty"`
}

// The description you choose to give the card.
func (r *UpdateACardParameters) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The status to update the Card with.
func (r *UpdateACardParameters) GetStatus() (Status UpdateACardParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The card's updated billing address.
func (r *UpdateACardParameters) GetBillingAddress() (BillingAddress UpdateACardParametersBillingAddress) {
	if r != nil && r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r *UpdateACardParameters) GetDigitalWallet() (DigitalWallet UpdateACardParametersDigitalWallet) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
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
	Line1 *string `json:"line1"`
	// The second line of the billing address.
	Line2 *string `json:"line2,omitempty"`
	// The city of the billing address.
	City *string `json:"city"`
	// The US state of the billing address.
	State *string `json:"state"`
	// The postal code of the billing address.
	PostalCode *string `json:"postal_code"`
}

// The first line of the billing address.
func (r *UpdateACardParametersBillingAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r *UpdateACardParametersBillingAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r *UpdateACardParametersBillingAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r *UpdateACardParametersBillingAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r *UpdateACardParametersBillingAddress) GetPostalCode() (PostalCode string) {
	if r != nil && r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}

//
type UpdateACardParametersDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `json:"email,omitempty"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `json:"phone,omitempty"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileId *string `json:"card_profile_id,omitempty"`
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r *UpdateACardParametersDigitalWallet) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r *UpdateACardParametersDigitalWallet) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r *UpdateACardParametersDigitalWallet) GetCardProfileId() (CardProfileId string) {
	if r != nil && r.CardProfileId != nil {
		CardProfileId = *r.CardProfileId
	}
	return
}

type ListCardsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Cards to ones belonging to the specified Account.
	AccountId *string                  `query:"account_id"`
	CreatedAt *ListCardsQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListCardsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListCardsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Cards to ones belonging to the specified Account.
func (r *ListCardsQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *ListCardsQuery) GetCreatedAt() (CreatedAt ListCardsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListCardsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCardsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCardsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCardsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCardsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type CardList struct {
	// The contents of the list.
	Data *[]Card `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *CardList) GetData() (Data []Card) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
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

func (r *PreloadedCardService) Create(ctx context.Context, body *CreateACardParameters, opts ...*core.RequestOpts) (res *Card, err error) {
	err = r.Cards.post(
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

func (r *PreloadedCardService) Retrieve(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *Card, err error) {
	err = r.Cards.get(
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

func (r *PreloadedCardService) Update(ctx context.Context, card_id string, body *UpdateACardParameters, opts ...*core.RequestOpts) (res *Card, err error) {
	err = r.Cards.patch(
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

func (r *PreloadedCardService) List(ctx context.Context, query *ListCardsQuery, opts ...*core.RequestOpts) (res *CardsPage, err error) {
	page := &CardsPage{
		Page: &pagination.Page[Card]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/cards",
			},
			Requester: r.Cards.Requester,
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

func (r *PreloadedCardService) RetrieveSensitiveDetails(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *CardDetails, err error) {
	err = r.Cards.get(
		ctx,
		fmt.Sprintf("/cards/%s/details", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}
