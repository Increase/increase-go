package check_deposits

import "context"
import "increase/core"
import "increase/simulations"
import "fmt"
import "increase/pagination"

type CheckDepositService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCheckDepositService(requester core.Requester) (r *CheckDepositService) {
	r = &CheckDepositService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedCheckDepositService struct {
	CheckDeposits *CheckDepositService
}

func (r *PreloadedCheckDepositService) Init(service *CheckDepositService) {
	r.CheckDeposits = service
}

func NewPreloadedCheckDepositService(service *CheckDepositService) (r *PreloadedCheckDepositService) {
	r = &PreloadedCheckDepositService{}
	r.Init(service)
	return
}

type CreateACheckDepositParameters struct {
	// The identifier for the Account to deposit the check in.
	AccountId *string `json:"account_id"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The currency to use for the deposit.
	Currency *string `json:"currency"`
	// The File containing the check's front image.
	FrontImageFileId *string `json:"front_image_file_id"`
	// The File containing the check's back image.
	BackImageFileId *string `json:"back_image_file_id"`
}

// The identifier for the Account to deposit the check in.
func (r *CreateACheckDepositParameters) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The deposit amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r *CreateACheckDepositParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The currency to use for the deposit.
func (r *CreateACheckDepositParameters) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The File containing the check's front image.
func (r *CreateACheckDepositParameters) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The File containing the check's back image.
func (r *CreateACheckDepositParameters) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
}

type ListCheckDepositsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountId *string                          `query:"account_id"`
	CreatedAt *ListCheckDepositsQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListCheckDepositsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListCheckDepositsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Deposits to those belonging to the specified Account.
func (r *ListCheckDepositsQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *ListCheckDepositsQuery) GetCreatedAt() (CreatedAt ListCheckDepositsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListCheckDepositsQueryCreatedAt struct {
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
func (r *ListCheckDepositsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type CheckDepositList struct {
	// The contents of the list.
	Data *[]simulations.CheckDeposit `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *CheckDepositList) GetData() (Data []simulations.CheckDeposit) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CheckDepositList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *CheckDepositService) Create(ctx context.Context, body *CreateACheckDepositParameters, opts ...*core.RequestOpts) (res *simulations.CheckDeposit, err error) {
	err = r.post(
		ctx,
		"/check_deposits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedCheckDepositService) Create(ctx context.Context, body *CreateACheckDepositParameters, opts ...*core.RequestOpts) (res *simulations.CheckDeposit, err error) {
	err = r.CheckDeposits.post(
		ctx,
		"/check_deposits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CheckDepositService) Retrieve(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *simulations.CheckDeposit, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/check_deposits/%s", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedCheckDepositService) Retrieve(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *simulations.CheckDeposit, err error) {
	err = r.CheckDeposits.get(
		ctx,
		fmt.Sprintf("/check_deposits/%s", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *CheckDepositService) List(ctx context.Context, query *ListCheckDepositsQuery, opts ...*core.RequestOpts) (res *simulations.CheckDepositsPage, err error) {
	page := &simulations.CheckDepositsPage{
		Page: &pagination.Page[simulations.CheckDeposit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_deposits",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedCheckDepositService) List(ctx context.Context, query *ListCheckDepositsQuery, opts ...*core.RequestOpts) (res *simulations.CheckDepositsPage, err error) {
	page := &simulations.CheckDepositsPage{
		Page: &pagination.Page[simulations.CheckDeposit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_deposits",
			},
			Requester: r.CheckDeposits.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
