package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
)

type DigitalWalletTokenService struct {
	Options []options.RequestOption
}

func NewDigitalWalletTokenService(opts ...options.RequestOption) (r *DigitalWalletTokenService) {
	r = &DigitalWalletTokenService{}
	r.Options = opts
	return
}

// Retrieve a Digital Wallet Token
func (r *DigitalWalletTokenService) Get(ctx context.Context, digital_wallet_token_id string, opts ...options.RequestOption) (res *types.DigitalWalletToken, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_wallet_tokens/%s", digital_wallet_token_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query *types.DigitalWalletTokenListParams, opts ...options.RequestOption) (res *types.DigitalWalletTokensPage, err error) {
	opts = append(r.Options, opts...)
	path := "digital_wallet_tokens"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.DigitalWalletTokensPage{
		Page: &pagination.Page[types.DigitalWalletToken]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
