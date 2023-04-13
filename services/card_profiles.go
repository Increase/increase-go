package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type CardProfileService struct {
	Options []option.RequestOption
}

func NewCardProfileService(opts ...option.RequestOption) (r *CardProfileService) {
	r = &CardProfileService{}
	r.Options = opts
	return
}

// Create a Card Profile
func (r *CardProfileService) New(ctx context.Context, body *requests.CardProfileNewParams, opts ...option.RequestOption) (res *responses.CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_profiles"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Card Profile
func (r *CardProfileService) Get(ctx context.Context, card_profile_id string, opts ...option.RequestOption) (res *responses.CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_profiles/%s", card_profile_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Card Profiles
func (r *CardProfileService) List(ctx context.Context, query *requests.CardProfileListParams, opts ...option.RequestOption) (res *responses.Page[responses.CardProfile], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_profiles"
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

// List Card Profiles
func (r *CardProfileService) ListAutoPager(ctx context.Context, query *requests.CardProfileListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.CardProfile] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
