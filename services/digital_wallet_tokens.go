package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
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
	u, err := url.Parse(fmt.Sprintf("digital_wallet_tokens/%s", digital_wallet_token_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query *types.DigitalWalletTokenListParams, opts ...options.RequestOption) (res *types.DigitalWalletTokensPage, err error) {
	u, err := url.Parse(fmt.Sprintf("digital_wallet_tokens"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	res = &types.DigitalWalletTokensPage{
		Page: &pagination.Page[types.DigitalWalletToken]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
