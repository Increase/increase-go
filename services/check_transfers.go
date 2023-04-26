package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type CheckTransferService struct {
	Options []option.RequestOption
}

func NewCheckTransferService(opts ...option.RequestOption) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Options = opts
	return
}

// Create a Check Transfer
func (r *CheckTransferService) New(ctx context.Context, body *requests.CheckTransferNewParams, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "check_transfers"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Check Transfer
func (r *CheckTransferService) Get(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Check Transfers
func (r *CheckTransferService) List(ctx context.Context, query *requests.CheckTransferListParams, opts ...option.RequestOption) (res *responses.Page[responses.CheckTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "check_transfers"
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

// List Check Transfers
func (r *CheckTransferService) ListAutoPager(ctx context.Context, query *requests.CheckTransferListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.CheckTransfer] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Check Transfer
func (r *CheckTransferService) Approve(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/approve", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/cancel", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Request a stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/stop_payment", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
