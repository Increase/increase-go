package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"io"
	"net/url"
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
func (r *CardDisputeService) New(ctx context.Context, body *types.CreateACardDisputeParameters, opts ...options.RequestOption) (res *types.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("card_disputes"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, content, opts...)
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

// Retrieve a Card Dispute
func (r *CardDisputeService) Get(ctx context.Context, card_dispute_id string, opts ...options.RequestOption) (res *types.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("card_disputes/%s", card_dispute_id))
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

// List Card Disputes
func (r *CardDisputeService) List(ctx context.Context, query *types.CardDisputeListParams, opts ...options.RequestOption) (res *types.CardDisputesPage, err error) {
	u, err := url.Parse(fmt.Sprintf("card_disputes"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CardDisputesPage{
		Page: &pagination.Page[types.CardDispute]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
