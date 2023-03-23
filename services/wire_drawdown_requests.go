package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type WireDrawdownRequestService struct {
	Options []options.RequestOption
}

func NewWireDrawdownRequestService(opts ...options.RequestOption) (r *WireDrawdownRequestService) {
	r = &WireDrawdownRequestService{}
	r.Options = opts
	return
}

// Create a Wire Drawdown Request
func (r *WireDrawdownRequestService) New(ctx context.Context, body *requests.CreateAWireDrawdownRequestParameters, opts ...options.RequestOption) (res *responses.WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "wire_drawdown_requests"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Wire Drawdown Request
func (r *WireDrawdownRequestService) Get(ctx context.Context, wire_drawdown_request_id string, opts ...options.RequestOption) (res *responses.WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_drawdown_requests/%s", wire_drawdown_request_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query *requests.WireDrawdownRequestListParams, opts ...options.RequestOption) (res *responses.WireDrawdownRequestsPage, err error) {
	opts = append(r.Options, opts...)
	path := "wire_drawdown_requests"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.WireDrawdownRequestsPage{
		Page: &pagination.Page[responses.WireDrawdownRequest]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}
