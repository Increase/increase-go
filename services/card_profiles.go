package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
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
	path := "card_profiles"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Card Profile
func (r *CardProfileService) Get(ctx context.Context, card_profile_id string, opts ...options.RequestOption) (res *types.CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_profiles/%s", card_profile_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Card Profiles
func (r *CardProfileService) List(ctx context.Context, query *types.CardProfileListParams, opts ...options.RequestOption) (res *types.CardProfilesPage, err error) {
	opts = append(r.Options, opts...)
	path := "card_profiles"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CardProfilesPage{
		Page: &pagination.Page[types.CardProfile]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
