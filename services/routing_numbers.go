package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type RoutingNumberService struct {
	Options []option.RequestOption
}

func NewRoutingNumberService(opts ...option.RequestOption) (r *RoutingNumberService) {
	r = &RoutingNumberService{}
	r.Options = opts
	return
}

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) List(ctx context.Context, query *requests.RoutingNumberListParams, opts ...option.RequestOption) (res *responses.Page[responses.RoutingNumber], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "routing_numbers"
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

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) ListAutoPager(ctx context.Context, query *requests.RoutingNumberListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.RoutingNumber] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
