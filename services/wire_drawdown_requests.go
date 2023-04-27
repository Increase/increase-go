package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type WireDrawdownRequestService struct {
	Options []option.RequestOption
}

func NewWireDrawdownRequestService(opts ...option.RequestOption) (r *WireDrawdownRequestService) {
	r = &WireDrawdownRequestService{}
	r.Options = opts
	return
}

// Create a Wire Drawdown Request
func (r *WireDrawdownRequestService) New(ctx context.Context, body requests.WireDrawdownRequestNewParams, opts ...option.RequestOption) (res *responses.WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "wire_drawdown_requests"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Wire Drawdown Request
func (r *WireDrawdownRequestService) Get(ctx context.Context, wire_drawdown_request_id string, opts ...option.RequestOption) (res *responses.WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_drawdown_requests/%s", wire_drawdown_request_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query requests.WireDrawdownRequestListParams, opts ...option.RequestOption) (res *responses.Page[responses.WireDrawdownRequest], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "wire_drawdown_requests"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) ListAutoPaging(ctx context.Context, query requests.WireDrawdownRequestListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.WireDrawdownRequest] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
