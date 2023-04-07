package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type InboundWireDrawdownRequestService struct {
	Options []option.RequestOption
}

func NewInboundWireDrawdownRequestService(opts ...option.RequestOption) (r *InboundWireDrawdownRequestService) {
	r = &InboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Wire Drawdown Request
func (r *InboundWireDrawdownRequestService) Get(ctx context.Context, inbound_wire_drawdown_request_id string, opts ...option.RequestOption) (res *responses.InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_wire_drawdown_requests/%s", inbound_wire_drawdown_request_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) List(ctx context.Context, query *requests.InboundWireDrawdownRequestListParams, opts ...option.RequestOption) (res *responses.Page[responses.InboundWireDrawdownRequest], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_wire_drawdown_requests"
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

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) ListAutoPager(ctx context.Context, query *requests.InboundWireDrawdownRequestListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.InboundWireDrawdownRequest] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
