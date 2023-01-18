package external_accounts

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type ExternalAccountService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewExternalAccountService(requester core.Requester) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type ExternalAccount struct {
	// The External Account's identifier.
	ID string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the External Account was created.
	CreatedAt string `json:"created_at"`
	// The External Account's description for display purposes.
	Description string `json:"description"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number"`
	// The destination account number.
	AccountNumber string `json:"account_number"`
	// The type of the account to which the transfer will be sent.
	Funding ExternalAccountFunding `json:"funding"`
	// A constant representing the object's type. For this resource it will always be
	// `external_account`.
	Type ExternalAccountType `json:"type"`
}

type ExternalAccountFunding string

const (
	ExternalAccountFundingChecking ExternalAccountFunding = "checking"
	ExternalAccountFundingSavings  ExternalAccountFunding = "savings"
	ExternalAccountFundingOther    ExternalAccountFunding = "other"
)

type ExternalAccountType string

const (
	ExternalAccountTypeExternalAccount ExternalAccountType = "external_account"
)

type CreateAnExternalAccountParameters struct {
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number"`
	// The account number for the destination account.
	AccountNumber string `json:"account_number"`
	// The type of the destination account. Defaults to `checking`.
	Funding CreateAnExternalAccountParametersFunding `json:"funding,omitempty"`
	// The name you choose for the Account.
	Description string `json:"description"`
}

type CreateAnExternalAccountParametersFunding string

const (
	CreateAnExternalAccountParametersFundingChecking CreateAnExternalAccountParametersFunding = "checking"
	CreateAnExternalAccountParametersFundingSavings  CreateAnExternalAccountParametersFunding = "savings"
	CreateAnExternalAccountParametersFundingOther    CreateAnExternalAccountParametersFunding = "other"
)

type UpdateAnExternalAccountParameters struct {
	// The description you choose to give the external account.
	Description string `json:"description,omitempty"`
}

type ListExternalAccountsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
}

//
type ExternalAccountList struct {
	// The contents of the list.
	Data []ExternalAccount `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *ExternalAccountList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *ExternalAccountService) Create(ctx context.Context, body *CreateAnExternalAccountParameters, opts ...*core.RequestOpts) (res *ExternalAccount, err error) {
	err = r.post(
		ctx,
		"/external_accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *ExternalAccountService) Retrieve(ctx context.Context, external_account_id string, opts ...*core.RequestOpts) (res *ExternalAccount, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/external_accounts/%s", external_account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *ExternalAccountService) Update(ctx context.Context, external_account_id string, body *UpdateAnExternalAccountParameters, opts ...*core.RequestOpts) (res *ExternalAccount, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/external_accounts/%s", external_account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

type ExternalAccountsPage struct {
	*pagination.Page[ExternalAccount]
}

func (r *ExternalAccountsPage) ExternalAccount() *ExternalAccount {
	return r.Current()
}

func (r *ExternalAccountsPage) GetNextPage() (*ExternalAccountsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &ExternalAccountsPage{page}, nil
	}
}

func (r *ExternalAccountService) List(ctx context.Context, query *ListExternalAccountsQuery, opts ...*core.RequestOpts) (res *ExternalAccountsPage, err error) {
	page := &ExternalAccountsPage{
		Page: &pagination.Page[ExternalAccount]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/external_accounts",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
