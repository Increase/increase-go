package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type RoutingNumberService struct {
	Options []options.RequestOption
}

func NewRoutingNumberService(opts ...options.RequestOption) (r *RoutingNumberService) {
	r = &RoutingNumberService{}
	r.Options = opts
	return
}

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) List(ctx context.Context, query *types.RoutingNumberListParams, opts ...options.RequestOption) (res *types.RoutingNumbersPage, err error) {
	u, err := url.Parse(fmt.Sprintf("routing_numbers"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.RoutingNumbersPage{
		Page: &pagination.Page[types.RoutingNumber]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
