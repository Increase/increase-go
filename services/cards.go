package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type CardService struct {
	Options []option.RequestOption
}

func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	return
}

// Create a Card
func (r *CardService) New(ctx context.Context, body *requests.CardNewParams, opts ...option.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := "cards"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, card_id string, opts ...option.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, card_id string, body *requests.CardUpdateParams, opts ...option.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Cards
func (r *CardService) List(ctx context.Context, query *requests.CardListParams, opts ...option.RequestOption) (res *responses.Page[responses.Card], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cards"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return
	}
	err = cfg.Execute()
	if err != nil {
		return
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Cards
func (r *CardService) ListAutoPager(ctx context.Context, query *requests.CardListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Card] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve sensitive details for a Card
func (r *CardService) GetSensitiveDetails(ctx context.Context, card_id string, opts ...option.RequestOption) (res *responses.CardDetails, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/details", card_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
