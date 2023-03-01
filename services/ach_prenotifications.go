package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	path := "ach_prenotifications"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, ach_prenotification_id string, opts ...options.RequestOption) (res *types.ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_prenotifications/%s", ach_prenotification_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query *types.ACHPrenotificationListParams, opts ...options.RequestOption) (res *types.ACHPrenotificationsPage, err error) {
	opts = append(r.Options, opts...)
	path := "ach_prenotifications"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.ACHPrenotificationsPage{
		Page: &pagination.Page[types.ACHPrenotification]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
