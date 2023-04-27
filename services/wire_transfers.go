package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type WireTransferService struct {
	Options []option.RequestOption
}

func NewWireTransferService(opts ...option.RequestOption) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Options = opts
	return
}

// Create a Wire Transfer
func (r *WireTransferService) New(ctx context.Context, body requests.WireTransferNewParams, opts ...option.RequestOption) (res *responses.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "wire_transfers"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Wire Transfer
func (r *WireTransferService) Get(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *responses.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s", wire_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Transfers
func (r *WireTransferService) List(ctx context.Context, query requests.WireTransferListParams, opts ...option.RequestOption) (res *responses.Page[responses.WireTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "wire_transfers"
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

// List Wire Transfers
func (r *WireTransferService) ListAutoPaging(ctx context.Context, query requests.WireTransferListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.WireTransfer] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Wire Transfer
func (r *WireTransferService) Approve(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *responses.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s/approve", wire_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Wire Transfer
func (r *WireTransferService) Cancel(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *responses.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s/cancel", wire_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the reversal of a [Wire Transfer](#wire-transfers) by the Federal
// Reserve due to error conditions. This will also create a
// [Transaction](#transaction) to account for the returned funds. This Wire
// Transfer must first have a `status` of `complete`.'
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *responses.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/wire_transfers/%s/reverse", wire_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Wire Transfer](#wire-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *responses.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/wire_transfers/%s/submit", wire_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
