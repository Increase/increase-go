package services

import "context"
import "increase/core"
import "increase/types"
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

func (r *AccountTransferService) Create(ctx context.Context, body *types.CreateAnAccountTransferParameters, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
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

func (r *AccountTransferService) Retrieve(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
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

func (r *AccountTransferService) List(ctx context.Context, query *types.ListAccountTransfersQuery, opts ...*core.RequestOpts) (res *types.AccountTransfersPage, err error) {
	page := &types.AccountTransfersPage{
		Page: &pagination.Page[types.AccountTransfer]{
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
