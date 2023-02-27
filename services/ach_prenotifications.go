package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type ACHPrenotificationService struct {
	Options []options.RequestOption
}

func NewACHPrenotificationService(opts ...options.RequestOption) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Options = opts
	return
}

// Create an ACH Prenotification
func (r *ACHPrenotificationService) New(ctx context.Context, body *types.CreateAnACHPrenotificationParameters, opts ...options.RequestOption) (res *types.ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("ach_prenotifications"))
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

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, ach_prenotification_id string, opts ...options.RequestOption) (res *types.ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("ach_prenotifications/%s", ach_prenotification_id))
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

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query *types.ACHPrenotificationListParams, opts ...options.RequestOption) (res *types.ACHPrenotificationsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("ach_prenotifications"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.ACHPrenotificationsPage{
		Page: &pagination.Page[types.ACHPrenotification]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
