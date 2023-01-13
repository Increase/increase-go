package account_numbers

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type AccountNumberService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountNumberService(requester core.Requester) (r *AccountNumberService) {
	r = &AccountNumberService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedAccountNumberService struct {
	AccountNumbers *AccountNumberService
}

func (r *PreloadedAccountNumberService) Init(service *AccountNumberService) {
	r.AccountNumbers = service
}

func NewPreloadedAccountNumberService(service *AccountNumberService) (r *PreloadedAccountNumberService) {
	r = &PreloadedAccountNumberService{}
	r.Init(service)
	return
}

//
type AccountNumber struct {
	// The identifier for the account this Account Number belongs to.
	AccountId *string `json:"account_id"`
	// The account number.
	AccountNumber *string `json:"account_number"`
	// The Account Number identifier.
	Id *string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt *string `json:"created_at"`
	// The name you choose for the Account Number.
	Name *string `json:"name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// This indicates if payments can be made to the Account Number.
	Status *AccountNumberStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type *AccountNumberType `json:"type"`
}

// The identifier for the account this Account Number belongs to.
func (r *AccountNumber) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The account number.
func (r *AccountNumber) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The Account Number identifier.
func (r *AccountNumber) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// Number was created.
func (r *AccountNumber) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The name you choose for the Account Number.
func (r *AccountNumber) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *AccountNumber) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// This indicates if payments can be made to the Account Number.
func (r *AccountNumber) GetStatus() (Status AccountNumberStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_number`.
func (r *AccountNumber) GetType() (Type AccountNumberType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type AccountNumberStatus string

const (
	AccountNumberStatusActive   AccountNumberStatus = "active"
	AccountNumberStatusDisabled AccountNumberStatus = "disabled"
	AccountNumberStatusCanceled AccountNumberStatus = "canceled"
)

type AccountNumberType string

const (
	AccountNumberTypeAccountNumber AccountNumberType = "account_number"
)

type CreateAnAccountNumberParameters struct {
	// The Account the Account Number should belong to.
	AccountId *string `json:"account_id"`
	// The name you choose for the Account Number.
	Name *string `json:"name"`
}

// The Account the Account Number should belong to.
func (r *CreateAnAccountNumberParameters) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The name you choose for the Account Number.
func (r *CreateAnAccountNumberParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

type UpdateAnAccountNumberParameters struct {
	// The name you choose for the Account Number.
	Name *string `json:"name,omitempty"`
	// This indicates if transfers can be made to the Account Number.
	Status *UpdateAnAccountNumberParametersStatus `json:"status,omitempty"`
}

// The name you choose for the Account Number.
func (r *UpdateAnAccountNumberParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// This indicates if transfers can be made to the Account Number.
func (r *UpdateAnAccountNumberParameters) GetStatus() (Status UpdateAnAccountNumberParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type UpdateAnAccountNumberParametersStatus string

const (
	UpdateAnAccountNumberParametersStatusActive   UpdateAnAccountNumberParametersStatus = "active"
	UpdateAnAccountNumberParametersStatusDisabled UpdateAnAccountNumberParametersStatus = "disabled"
	UpdateAnAccountNumberParametersStatusCanceled UpdateAnAccountNumberParametersStatus = "canceled"
)

type ListAccountNumbersQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status *ListAccountNumbersQueryStatus `query:"status"`
	// Filter Account Numbers to those belonging to the specified Account.
	AccountId *string `query:"account_id"`
}

// Return the page of entries after this one.
func (r *ListAccountNumbersQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListAccountNumbersQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// The status to retrieve Account Numbers for.
func (r *ListAccountNumbersQuery) GetStatus() (Status ListAccountNumbersQueryStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// Filter Account Numbers to those belonging to the specified Account.
func (r *ListAccountNumbersQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

type ListAccountNumbersQueryStatus string

const (
	ListAccountNumbersQueryStatusActive   ListAccountNumbersQueryStatus = "active"
	ListAccountNumbersQueryStatusDisabled ListAccountNumbersQueryStatus = "disabled"
	ListAccountNumbersQueryStatusCanceled ListAccountNumbersQueryStatus = "canceled"
)

//
type AccountNumberList struct {
	// The contents of the list.
	Data *[]AccountNumber `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *AccountNumberList) GetData() (Data []AccountNumber) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountNumberList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *AccountNumberService) Create(ctx context.Context, body *CreateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *AccountNumber, err error) {
	err = r.post(
		ctx,
		"/account_numbers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedAccountNumberService) Create(ctx context.Context, body *CreateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *AccountNumber, err error) {
	err = r.AccountNumbers.post(
		ctx,
		"/account_numbers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *AccountNumberService) Retrieve(ctx context.Context, account_number_id string, opts ...*core.RequestOpts) (res *AccountNumber, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedAccountNumberService) Retrieve(ctx context.Context, account_number_id string, opts ...*core.RequestOpts) (res *AccountNumber, err error) {
	err = r.AccountNumbers.get(
		ctx,
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *UpdateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *AccountNumber, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedAccountNumberService) Update(ctx context.Context, account_number_id string, body *UpdateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *AccountNumber, err error) {
	err = r.AccountNumbers.patch(
		ctx,
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

type AccountNumbersPage struct {
	*pagination.Page[AccountNumber]
}

func (r *AccountNumbersPage) AccountNumber() *AccountNumber {
	return r.Current()
}

func (r *AccountNumbersPage) GetNextPage() (*AccountNumbersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &AccountNumbersPage{page}, nil
	}
}

func (r *AccountNumberService) List(ctx context.Context, query *ListAccountNumbersQuery, opts ...*core.RequestOpts) (res *AccountNumbersPage, err error) {
	page := &AccountNumbersPage{
		Page: &pagination.Page[AccountNumber]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_numbers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedAccountNumberService) List(ctx context.Context, query *ListAccountNumbersQuery, opts ...*core.RequestOpts) (res *AccountNumbersPage, err error) {
	page := &AccountNumbersPage{
		Page: &pagination.Page[AccountNumber]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_numbers",
			},
			Requester: r.AccountNumbers.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
