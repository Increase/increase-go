package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type CardService struct {
	Options []options.RequestOption
}

func NewCardService(opts ...options.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	return
}

// Create a Card
func (r *CardService) New(ctx context.Context, body *types.CreateACardParameters, opts ...options.RequestOption) (res *types.Card, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("cards"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, body, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, card_id string, opts ...options.RequestOption) (res *types.Card, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("cards/%s", card_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, card_id string, body *types.UpdateACardParameters, opts ...options.RequestOption) (res *types.Card, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("cards/%s", card_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "PATCH", u, body, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Cards
func (r *CardService) List(ctx context.Context, query *types.CardListParams, opts ...options.RequestOption) (res *types.CardsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("cards"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.CardsPage{
		Page: &pagination.Page[types.Card]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

// Retrieve sensitive details for a Card
func (r *CardService) GetSensitiveDetails(ctx context.Context, card_id string, opts ...options.RequestOption) (res *types.CardDetails, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("cards/%s/details", card_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}
