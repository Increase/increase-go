package accounts

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type AccountService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountService(requester core.Requester) (r *AccountService) {
	r = &AccountService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedAccountService struct {
	Accounts *AccountService
}

func (r *PreloadedAccountService) Init(service *AccountService) {
	r.Accounts = service
}

func NewPreloadedAccountService(service *AccountService) (r *PreloadedAccountService) {
	r = &PreloadedAccountService{}
	r.Init(service)
	return
}

//
type Account struct {
	// The current balance of the Account in the minor unit of the currency. For
	// dollars, for example, this is cents.
	Balance *int `json:"balance"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// was created.
	CreatedAt *string `json:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
	// currency.
	Currency *AccountCurrency `json:"currency"`
	// The identifier for the Entity the Account belongs to.
	EntityId *string `json:"entity_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity.
	InformationalEntityId *string `json:"informational_entity_id"`
	// The Account identifier.
	Id *string `json:"id"`
	// The interest accrued but not yet paid, expressed as a string containing a
	// floating-point value.
	InterestAccrued *string `json:"interest_accrued"`
	// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
	// interest was accrued.
	InterestAccruedAt *string `json:"interest_accrued_at"`
	// The name you choose for the Account.
	Name *string `json:"name"`
	// The status of the Account.
	Status *AccountStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `account`.
	Type *AccountType `json:"type"`
}

// The current balance of the Account in the minor unit of the currency. For
// dollars, for example, this is cents.
func (r *Account) GetBalance() (Balance int) {
	if r != nil && r.Balance != nil {
		Balance = *r.Balance
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// was created.
func (r *Account) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
// currency.
func (r *Account) GetCurrency() (Currency AccountCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Entity the Account belongs to.
func (r *Account) GetEntityId() (EntityId string) {
	if r != nil && r.EntityId != nil {
		EntityId = *r.EntityId
	}
	return
}

// The identifier of an Entity that, while not owning the Account, is associated
// with its activity.
func (r *Account) GetInformationalEntityId() (InformationalEntityId string) {
	if r != nil && r.InformationalEntityId != nil {
		InformationalEntityId = *r.InformationalEntityId
	}
	return
}

// The Account identifier.
func (r *Account) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The interest accrued but not yet paid, expressed as a string containing a
// floating-point value.
func (r *Account) GetInterestAccrued() (InterestAccrued string) {
	if r != nil && r.InterestAccrued != nil {
		InterestAccrued = *r.InterestAccrued
	}
	return
}

// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
// interest was accrued.
func (r *Account) GetInterestAccruedAt() (InterestAccruedAt string) {
	if r != nil && r.InterestAccruedAt != nil {
		InterestAccruedAt = *r.InterestAccruedAt
	}
	return
}

// The name you choose for the Account.
func (r *Account) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The status of the Account.
func (r *Account) GetStatus() (Status AccountStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account`.
func (r *Account) GetType() (Type AccountType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type AccountCurrency string

const (
	AccountCurrencyCad AccountCurrency = "CAD"
	AccountCurrencyChf AccountCurrency = "CHF"
	AccountCurrencyEur AccountCurrency = "EUR"
	AccountCurrencyGbp AccountCurrency = "GBP"
	AccountCurrencyJpy AccountCurrency = "JPY"
	AccountCurrencyUsd AccountCurrency = "USD"
)

type AccountStatus string

const (
	AccountStatusOpen   AccountStatus = "open"
	AccountStatusClosed AccountStatus = "closed"
)

type AccountType string

const (
	AccountTypeAccount AccountType = "account"
)

type CreateAnAccountParameters struct {
	// The identifier for the Entity that will own the Account.
	EntityId *string `json:"entity_id,omitempty"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityId *string `json:"informational_entity_id,omitempty"`
	// The name you choose for the Account.
	Name *string `json:"name"`
}

// The identifier for the Entity that will own the Account.
func (r *CreateAnAccountParameters) GetEntityId() (EntityId string) {
	if r != nil && r.EntityId != nil {
		EntityId = *r.EntityId
	}
	return
}

// The identifier of an Entity that, while not owning the Account, is associated
// with its activity. Its relationship to your group must be `informational`.
func (r *CreateAnAccountParameters) GetInformationalEntityId() (InformationalEntityId string) {
	if r != nil && r.InformationalEntityId != nil {
		InformationalEntityId = *r.InformationalEntityId
	}
	return
}

// The name you choose for the Account.
func (r *CreateAnAccountParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

type UpdateAnAccountParameters struct {
	// The new name of the Account.
	Name *string `json:"name,omitempty"`
}

// The new name of the Account.
func (r *UpdateAnAccountParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

type ListAccountsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityId *string `query:"entity_id"`
	// Filter Accounts for those with the specified status.
	Status *ListAccountsQueryStatus `query:"status"`
}

// Return the page of entries after this one.
func (r *ListAccountsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListAccountsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Accounts for those belonging to the specified Entity.
func (r *ListAccountsQuery) GetEntityId() (EntityId string) {
	if r != nil && r.EntityId != nil {
		EntityId = *r.EntityId
	}
	return
}

// Filter Accounts for those with the specified status.
func (r *ListAccountsQuery) GetStatus() (Status ListAccountsQueryStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type ListAccountsQueryStatus string

const (
	ListAccountsQueryStatusOpen   ListAccountsQueryStatus = "open"
	ListAccountsQueryStatusClosed ListAccountsQueryStatus = "closed"
)

//
type AccountList struct {
	// The contents of the list.
	Data *[]Account `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *AccountList) GetData() (Data []Account) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *AccountService) Create(ctx context.Context, body *CreateAnAccountParameters, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.post(
		ctx,
		"/accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedAccountService) Create(ctx context.Context, body *CreateAnAccountParameters, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.Accounts.post(
		ctx,
		"/accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *AccountService) Retrieve(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedAccountService) Retrieve(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.Accounts.get(
		ctx,
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *AccountService) Update(ctx context.Context, account_id string, body *UpdateAnAccountParameters, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedAccountService) Update(ctx context.Context, account_id string, body *UpdateAnAccountParameters, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.Accounts.patch(
		ctx,
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

type AccountsPage struct {
	*pagination.Page[Account]
}

func (r *AccountsPage) Account() *Account {
	return r.Current()
}

func (r *AccountsPage) GetNextPage() (*AccountsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &AccountsPage{page}, nil
	}
}

func (r *AccountService) List(ctx context.Context, query *ListAccountsQuery, opts ...*core.RequestOpts) (res *AccountsPage, err error) {
	page := &AccountsPage{
		Page: &pagination.Page[Account]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/accounts",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedAccountService) List(ctx context.Context, query *ListAccountsQuery, opts ...*core.RequestOpts) (res *AccountsPage, err error) {
	page := &AccountsPage{
		Page: &pagination.Page[Account]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/accounts",
			},
			Requester: r.Accounts.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *AccountService) Close(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/accounts/%s/close", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedAccountService) Close(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *Account, err error) {
	err = r.Accounts.post(
		ctx,
		fmt.Sprintf("/accounts/%s/close", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}
