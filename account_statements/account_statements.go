package account_statements

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type AccountStatementService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountStatementService(requester core.Requester) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedAccountStatementService struct {
	AccountStatements *AccountStatementService
}

func (r *PreloadedAccountStatementService) Init(service *AccountStatementService) {
	r.AccountStatements = service
}

func NewPreloadedAccountStatementService(service *AccountStatementService) (r *PreloadedAccountStatementService) {
	r = &PreloadedAccountStatementService{}
	r.Init(service)
	return
}

//
type AccountStatement struct {
	// The Account Statement identifier.
	ID *string `json:"id"`
	// The identifier for the Account this Account Statement belongs to.
	AccountID *string `json:"account_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt *string `json:"created_at"`
	// The identifier of the File containing a PDF of the statement.
	FileID *string `json:"file_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart *string `json:"statement_period_start"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd *string `json:"statement_period_end"`
	// The Account's balance at the start of its statement period.
	StartingBalance *int `json:"starting_balance"`
	// The Account's balance at the start of its statement period.
	EndingBalance *int `json:"ending_balance"`
	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type *AccountStatementType `json:"type"`
}

// The Account Statement identifier.
func (r *AccountStatement) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the Account this Account Statement belongs to.
func (r *AccountStatement) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// Statement was created.
func (r *AccountStatement) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the File containing a PDF of the statement.
func (r *AccountStatement) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
// start of the period the Account Statement covers.
func (r *AccountStatement) GetStatementPeriodStart() (StatementPeriodStart string) {
	if r != nil && r.StatementPeriodStart != nil {
		StatementPeriodStart = *r.StatementPeriodStart
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
// of the period the Account Statement covers.
func (r *AccountStatement) GetStatementPeriodEnd() (StatementPeriodEnd string) {
	if r != nil && r.StatementPeriodEnd != nil {
		StatementPeriodEnd = *r.StatementPeriodEnd
	}
	return
}

// The Account's balance at the start of its statement period.
func (r *AccountStatement) GetStartingBalance() (StartingBalance int) {
	if r != nil && r.StartingBalance != nil {
		StartingBalance = *r.StartingBalance
	}
	return
}

// The Account's balance at the start of its statement period.
func (r *AccountStatement) GetEndingBalance() (EndingBalance int) {
	if r != nil && r.EndingBalance != nil {
		EndingBalance = *r.EndingBalance
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_statement`.
func (r *AccountStatement) GetType() (Type AccountStatementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type AccountStatementType string

const (
	AccountStatementTypeAccountStatement AccountStatementType = "account_statement"
)

type ListAccountStatementsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Account Statements to those belonging to the specified Account.
	AccountID            *string                                         `query:"account_id"`
	StatementPeriodStart *ListAccountStatementsQueryStatementPeriodStart `query:"statement_period_start"`
}

// Return the page of entries after this one.
func (r *ListAccountStatementsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListAccountStatementsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Account Statements to those belonging to the specified Account.
func (r *ListAccountStatementsQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r *ListAccountStatementsQuery) GetStatementPeriodStart() (StatementPeriodStart ListAccountStatementsQueryStatementPeriodStart) {
	if r != nil && r.StatementPeriodStart != nil {
		StatementPeriodStart = *r.StatementPeriodStart
	}
	return
}

type ListAccountStatementsQueryStatementPeriodStart struct {
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
func (r *ListAccountStatementsQueryStatementPeriodStart) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListAccountStatementsQueryStatementPeriodStart) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListAccountStatementsQueryStatementPeriodStart) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListAccountStatementsQueryStatementPeriodStart) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type AccountStatementList struct {
	// The contents of the list.
	Data *[]AccountStatement `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *AccountStatementList) GetData() (Data []AccountStatement) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountStatementList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *AccountStatementService) Retrieve(ctx context.Context, account_statement_id string, opts ...*core.RequestOpts) (res *AccountStatement, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/account_statements/%s", account_statement_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedAccountStatementService) Retrieve(ctx context.Context, account_statement_id string, opts ...*core.RequestOpts) (res *AccountStatement, err error) {
	err = r.AccountStatements.get(
		ctx,
		fmt.Sprintf("/account_statements/%s", account_statement_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type AccountStatementsPage struct {
	*pagination.Page[AccountStatement]
}

func (r *AccountStatementsPage) AccountStatement() *AccountStatement {
	return r.Current()
}

func (r *AccountStatementsPage) GetNextPage() (*AccountStatementsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &AccountStatementsPage{page}, nil
	}
}

func (r *AccountStatementService) List(ctx context.Context, query *ListAccountStatementsQuery, opts ...*core.RequestOpts) (res *AccountStatementsPage, err error) {
	page := &AccountStatementsPage{
		Page: &pagination.Page[AccountStatement]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_statements",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedAccountStatementService) List(ctx context.Context, query *ListAccountStatementsQuery, opts ...*core.RequestOpts) (res *AccountStatementsPage, err error) {
	page := &AccountStatementsPage{
		Page: &pagination.Page[AccountStatement]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_statements",
			},
			Requester: r.AccountStatements.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
