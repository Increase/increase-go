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

type CardProfileService struct {
	Options []options.RequestOption
}

func NewCardProfileService(opts ...options.RequestOption) (r *CardProfileService) {
	r = &CardProfileService{}
	r.Options = opts
	return
}

// Create a Card Profile
func (r *CardProfileService) New(ctx context.Context, body *types.CreateACardProfileParameters, opts ...options.RequestOption) (res *types.CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("card_profiles"))
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

// Retrieve a Card Profile
func (r *CardProfileService) Get(ctx context.Context, card_profile_id string, opts ...options.RequestOption) (res *types.CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("card_profiles/%s", card_profile_id))
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

// List Card Profiles
func (r *CardProfileService) List(ctx context.Context, query *types.CardProfileListParams, opts ...options.RequestOption) (res *types.CardProfilesPage, err error) {
	u, err := url.Parse(fmt.Sprintf("card_profiles"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CardProfilesPage{
		Page: &pagination.Page[types.CardProfile]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
