package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type ACHTransferService struct {
	Options []option.RequestOption
}

func NewACHTransferService(opts ...option.RequestOption) (r *ACHTransferService) {
	r = &ACHTransferService{}
	r.Options = opts
	return
}

// Create an ACH Transfer
func (r *ACHTransferService) New(ctx context.Context, body *requests.ACHTransferNewParams, opts ...option.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_transfers"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s", ach_transfer_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List ACH Transfers
func (r *ACHTransferService) List(ctx context.Context, query *requests.ACHTransferListParams, opts ...option.RequestOption) (res *responses.Page[responses.ACHTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_transfers"
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

// List ACH Transfers
func (r *ACHTransferService) ListAutoPager(ctx context.Context, query *requests.ACHTransferListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.ACHTransfer] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Approve(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/approve", ach_transfer_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Cancels an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Cancel(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/cancel", ach_transfer_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}
