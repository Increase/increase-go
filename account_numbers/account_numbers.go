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

//
type AccountNumber struct {
	// The identifier for the account this Account Number belongs to.
	AccountID string `json:"account_id"`
	// The account number.
	AccountNumber string `json:"account_number"`
	// The Account Number identifier.
	ID string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt string `json:"created_at"`
	// The name you choose for the Account Number.
	Name string `json:"name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number"`
	// This indicates if payments can be made to the Account Number.
	Status AccountNumberStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type AccountNumberType `json:"type"`
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
	AccountID string `json:"account_id"`
	// The name you choose for the Account Number.
	Name string `json:"name"`
}

type UpdateAnAccountNumberParameters struct {
	// The name you choose for the Account Number.
	Name string `json:"name,omitempty"`
	// This indicates if transfers can be made to the Account Number.
	Status UpdateAnAccountNumberParametersStatus `json:"status,omitempty"`
}

type UpdateAnAccountNumberParametersStatus string

const (
	UpdateAnAccountNumberParametersStatusActive   UpdateAnAccountNumberParametersStatus = "active"
	UpdateAnAccountNumberParametersStatusDisabled UpdateAnAccountNumberParametersStatus = "disabled"
	UpdateAnAccountNumberParametersStatusCanceled UpdateAnAccountNumberParametersStatus = "canceled"
)

type ListAccountNumbersQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status ListAccountNumbersQueryStatus `query:"status"`
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID string `query:"account_id"`
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
	Data []AccountNumber `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
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
