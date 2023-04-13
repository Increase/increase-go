package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type DigitalWalletTokenService struct {
	Options []option.RequestOption
}

func NewDigitalWalletTokenService(opts ...option.RequestOption) (r *DigitalWalletTokenService) {
	r = &DigitalWalletTokenService{}
	r.Options = opts
	return
}

// Retrieve a Digital Wallet Token
func (r *DigitalWalletTokenService) Get(ctx context.Context, digital_wallet_token_id string, opts ...option.RequestOption) (res *responses.DigitalWalletToken, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_wallet_tokens/%s", digital_wallet_token_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query *requests.DigitalWalletTokenListParams, opts ...option.RequestOption) (res *responses.Page[responses.DigitalWalletToken], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "digital_wallet_tokens"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) ListAutoPager(ctx context.Context, query *requests.DigitalWalletTokenListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.DigitalWalletToken] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
