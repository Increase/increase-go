package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type CardDisputeService struct {
	Options []option.RequestOption
}

func NewCardDisputeService(opts ...option.RequestOption) (r *CardDisputeService) {
	r = &CardDisputeService{}
	r.Options = opts
	return
}

// Create a Card Dispute
func (r *CardDisputeService) New(ctx context.Context, body *requests.CardDisputeNewParams, opts ...option.RequestOption) (res *responses.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_disputes"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Card Dispute
func (r *CardDisputeService) Get(ctx context.Context, card_dispute_id string, opts ...option.RequestOption) (res *responses.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_disputes/%s", card_dispute_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Card Disputes
func (r *CardDisputeService) List(ctx context.Context, query *requests.CardDisputeListParams, opts ...option.RequestOption) (res *responses.Page[responses.CardDispute], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_disputes"
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

// List Card Disputes
func (r *CardDisputeService) ListAutoPager(ctx context.Context, query *requests.CardDisputeListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.CardDispute] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
