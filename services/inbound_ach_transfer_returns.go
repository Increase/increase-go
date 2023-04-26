package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type InboundACHTransferReturnService struct {
	Options []option.RequestOption
}

func NewInboundACHTransferReturnService(opts ...option.RequestOption) (r *InboundACHTransferReturnService) {
	r = &InboundACHTransferReturnService{}
	r.Options = opts
	return
}

// Create an ACH Return
func (r *InboundACHTransferReturnService) New(ctx context.Context, body *requests.InboundACHTransferReturnNewParams, opts ...option.RequestOption) (res *responses.InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := "inbound_ach_transfer_returns"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Inbound ACH Transfer Return
func (r *InboundACHTransferReturnService) Get(ctx context.Context, inbound_ach_transfer_return_id string, opts ...option.RequestOption) (res *responses.InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfer_returns/%s", inbound_ach_transfer_return_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound ACH Transfer Returns
func (r *InboundACHTransferReturnService) List(ctx context.Context, query *requests.InboundACHTransferReturnListParams, opts ...option.RequestOption) (res *responses.Page[responses.InboundACHTransferReturn], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_ach_transfer_returns"
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

// List Inbound ACH Transfer Returns
func (r *InboundACHTransferReturnService) ListAutoPager(ctx context.Context, query *requests.InboundACHTransferReturnListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.InboundACHTransferReturn] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
