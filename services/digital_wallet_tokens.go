package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *DigitalWalletTokenService) Get(ctx context.Context, digital_wallet_token_id string, opts ...options.RequestOption) (res *responses.DigitalWalletToken, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_wallet_tokens/%s", digital_wallet_token_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query *requests.DigitalWalletTokenListParams, opts ...options.RequestOption) (res *responses.DigitalWalletTokensPage, err error) {
	opts = append(r.Options, opts...)
	path := "digital_wallet_tokens"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.DigitalWalletTokensPage{
		Page: &pagination.Page[responses.DigitalWalletToken]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
