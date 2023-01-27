package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type DigitalWalletTokenService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewDigitalWalletTokenService(requester core.Requester) (r *DigitalWalletTokenService) {
	r = &DigitalWalletTokenService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Retrieve a Digital Wallet Token
func (r *DigitalWalletTokenService) Get(ctx context.Context, digital_wallet_token_id string, opts ...*core.RequestOpts) (res *types.DigitalWalletToken, err error) {
	path := fmt.Sprintf("/digital_wallet_tokens/%s", digital_wallet_token_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query *types.ListDigitalWalletTokensQuery, opts ...*core.RequestOpts) (res *types.DigitalWalletTokensPage, err error) {
	page := &types.DigitalWalletTokensPage{
		Page: &pagination.Page[types.DigitalWalletToken]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/digital_wallet_tokens",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
