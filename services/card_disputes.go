package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type CardDisputeService struct {
	Options []options.RequestOption
}

func NewCardDisputeService(opts ...options.RequestOption) (r *CardDisputeService) {
	r = &CardDisputeService{}
	r.Options = opts
	return
}

// Create a Card Dispute
func (r *CardDisputeService) New(ctx context.Context, body *requests.CreateACardDisputeParameters, opts ...options.RequestOption) (res *responses.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_disputes"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Card Dispute
func (r *CardDisputeService) Get(ctx context.Context, card_dispute_id string, opts ...options.RequestOption) (res *responses.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_disputes/%s", card_dispute_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Card Disputes
func (r *CardDisputeService) List(ctx context.Context, query *requests.CardDisputeListParams, opts ...options.RequestOption) (res *responses.CardDisputesPage, err error) {
	opts = append(r.Options, opts...)
	path := "card_disputes"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.CardDisputesPage{
		Page: &pagination.Page[responses.CardDispute]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
