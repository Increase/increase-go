package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
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
	opts = append(r.Options, opts...)
	path := "routing_numbers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.RoutingNumbersPage{
		Page: &pagination.Page[types.RoutingNumber]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
