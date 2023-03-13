package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
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
	path := "cards"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, card_id string, opts ...options.RequestOption) (res *types.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, card_id string, body *types.UpdateACardParameters, opts ...options.RequestOption) (res *types.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Cards
func (r *CardService) List(ctx context.Context, query *types.CardListParams, opts ...options.RequestOption) (res *types.CardsPage, err error) {
	opts = append(r.Options, opts...)
	path := "cards"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CardsPage{
		Page: &pagination.Page[types.Card]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Retrieve sensitive details for a Card
func (r *CardService) GetSensitiveDetails(ctx context.Context, card_id string, opts ...options.RequestOption) (res *types.CardDetails, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/details", card_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
