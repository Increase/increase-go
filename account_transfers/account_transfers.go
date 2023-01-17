package account_transfers

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type AccountTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountTransferService(requester core.Requester) (r *AccountTransferService) {
	r = &AccountTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedAccountTransferService struct {
	AccountTransfers *AccountTransferService
}

func (r *PreloadedAccountTransferService) Init(service *AccountTransferService) {
	r.AccountTransfers = service
}

func NewPreloadedAccountTransferService(service *AccountTransferService) (r *PreloadedAccountTransferService) {
	r = &PreloadedAccountTransferService{}
	r.Init(service)
	return
}

//
type AccountTransfer struct {
	// The account transfer's identifier.
	ID *string `json:"id"`
	// The transfer amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The Account to which the transfer belongs.
	AccountID *string `json:"account_id"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *AccountTransferCurrency `json:"currency"`
	// The destination account's identifier.
	DestinationAccountID *string `json:"destination_account_id"`
	// The ID for the transaction receiving the transfer.
	DestinationTransactionID *string `json:"destination_transaction_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// The description that will show on the transactions.
	Description *string `json:"description"`
	// The transfer's network.
	Network *AccountTransferNetwork `json:"network"`
	// The lifecycle status of the transfer.
	Status *AccountTransferStatus `json:"status"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `json:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionID *string `json:"transaction_id"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *AccountTransferApproval `json:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *AccountTransferCancellation `json:"cancellation"`
	// A constant representing the object's type. For this resource it will always be
	// `account_transfer`.
	Type *AccountTransferType `json:"type"`
}

// The account transfer's identifier.
func (r *AccountTransfer) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The transfer amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *AccountTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The Account to which the transfer belongs.
func (r *AccountTransfer) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *AccountTransfer) GetCurrency() (Currency AccountTransferCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The destination account's identifier.
func (r *AccountTransfer) GetDestinationAccountID() (DestinationAccountID string) {
	if r != nil && r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// The ID for the transaction receiving the transfer.
func (r *AccountTransfer) GetDestinationTransactionID() (DestinationTransactionID string) {
	if r != nil && r.DestinationTransactionID != nil {
		DestinationTransactionID = *r.DestinationTransactionID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *AccountTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The description that will show on the transactions.
func (r *AccountTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The transfer's network.
func (r *AccountTransfer) GetNetwork() (Network AccountTransferNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// The lifecycle status of the transfer.
func (r *AccountTransfer) GetStatus() (Status AccountTransferStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *AccountTransfer) GetTemplateID() (TemplateID string) {
	if r != nil && r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *AccountTransfer) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *AccountTransfer) GetApproval() (Approval AccountTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *AccountTransfer) GetCancellation() (Cancellation AccountTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_transfer`.
func (r *AccountTransfer) GetType() (Type AccountTransferType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type AccountTransferCurrency string

const (
	AccountTransferCurrencyCad AccountTransferCurrency = "CAD"
	AccountTransferCurrencyChf AccountTransferCurrency = "CHF"
	AccountTransferCurrencyEur AccountTransferCurrency = "EUR"
	AccountTransferCurrencyGbp AccountTransferCurrency = "GBP"
	AccountTransferCurrencyJpy AccountTransferCurrency = "JPY"
	AccountTransferCurrencyUsd AccountTransferCurrency = "USD"
)

type AccountTransferNetwork string

const (
	AccountTransferNetworkAccount AccountTransferNetwork = "account"
)

type AccountTransferStatus string

const (
	AccountTransferStatusPendingSubmission AccountTransferStatus = "pending_submission"
	AccountTransferStatusPendingApproval   AccountTransferStatus = "pending_approval"
	AccountTransferStatusCanceled          AccountTransferStatus = "canceled"
	AccountTransferStatusRequiresAttention AccountTransferStatus = "requires_attention"
	AccountTransferStatusFlaggedByOperator AccountTransferStatus = "flagged_by_operator"
	AccountTransferStatusComplete          AccountTransferStatus = "complete"
)

//
type AccountTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt *string `json:"approved_at"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r *AccountTransferApproval) GetApprovedAt() (ApprovedAt string) {
	if r != nil && r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}

//
type AccountTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string `json:"canceled_at"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r *AccountTransferCancellation) GetCanceledAt() (CanceledAt string) {
	if r != nil && r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}

type AccountTransferType string

const (
	AccountTransferTypeAccountTransfer AccountTransferType = "account_transfer"
)

type CreateAnAccountTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID *string `json:"account_id"`
	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The description you choose to give the transfer.
	Description *string `json:"description"`
	// The identifier for the account that will receive the transfer.
	DestinationAccountID *string `json:"destination_account_id"`
}

// The identifier for the account that will send the transfer.
func (r *CreateAnAccountTransferParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The transfer amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r *CreateAnAccountTransferParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description you choose to give the transfer.
func (r *CreateAnAccountTransferParameters) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier for the account that will receive the transfer.
func (r *CreateAnAccountTransferParameters) GetDestinationAccountID() (DestinationAccountID string) {
	if r != nil && r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

type ListAccountTransfersQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Account Transfers to those that originated from the specified Account.
	AccountID *string                             `query:"account_id"`
	CreatedAt *ListAccountTransfersQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListAccountTransfersQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListAccountTransfersQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Account Transfers to those that originated from the specified Account.
func (r *ListAccountTransfersQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r *ListAccountTransfersQuery) GetCreatedAt() (CreatedAt ListAccountTransfersQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListAccountTransfersQueryCreatedAt struct {
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
func (r *ListAccountTransfersQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListAccountTransfersQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListAccountTransfersQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListAccountTransfersQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type AccountTransferList struct {
	// The contents of the list.
	Data *[]AccountTransfer `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *AccountTransferList) GetData() (Data []AccountTransfer) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountTransferList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *AccountTransferService) Create(ctx context.Context, body *CreateAnAccountTransferParameters, opts ...*core.RequestOpts) (res *AccountTransfer, err error) {
	err = r.post(
		ctx,
		"/account_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedAccountTransferService) Create(ctx context.Context, body *CreateAnAccountTransferParameters, opts ...*core.RequestOpts) (res *AccountTransfer, err error) {
	err = r.AccountTransfers.post(
		ctx,
		"/account_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *AccountTransferService) Retrieve(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *AccountTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/account_transfers/%s", account_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedAccountTransferService) Retrieve(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *AccountTransfer, err error) {
	err = r.AccountTransfers.get(
		ctx,
		fmt.Sprintf("/account_transfers/%s", account_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type AccountTransfersPage struct {
	*pagination.Page[AccountTransfer]
}

func (r *AccountTransfersPage) AccountTransfer() *AccountTransfer {
	return r.Current()
}

func (r *AccountTransfersPage) GetNextPage() (*AccountTransfersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &AccountTransfersPage{page}, nil
	}
}

func (r *AccountTransferService) List(ctx context.Context, query *ListAccountTransfersQuery, opts ...*core.RequestOpts) (res *AccountTransfersPage, err error) {
	page := &AccountTransfersPage{
		Page: &pagination.Page[AccountTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedAccountTransferService) List(ctx context.Context, query *ListAccountTransfersQuery, opts ...*core.RequestOpts) (res *AccountTransfersPage, err error) {
	page := &AccountTransfersPage{
		Page: &pagination.Page[AccountTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_transfers",
			},
			Requester: r.AccountTransfers.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
