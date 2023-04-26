package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type RealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

func NewRealTimePaymentsTransferService(opts ...option.RequestOption) (r *RealTimePaymentsTransferService) {
	r = &RealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Create a Real Time Payments Transfer
func (r *RealTimePaymentsTransferService) New(ctx context.Context, body *requests.RealTimePaymentsTransferNewParams, opts ...option.RequestOption) (res *responses.RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "real_time_payments_transfers"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Real Time Payments Transfer
func (r *RealTimePaymentsTransferService) Get(ctx context.Context, real_time_payments_transfer_id string, opts ...option.RequestOption) (res *responses.RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_payments_transfers/%s", real_time_payments_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Real Time Payments Transfers
func (r *RealTimePaymentsTransferService) List(ctx context.Context, query *requests.RealTimePaymentsTransferListParams, opts ...option.RequestOption) (res *responses.Page[responses.RealTimePaymentsTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "real_time_payments_transfers"
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

// List Real Time Payments Transfers
func (r *RealTimePaymentsTransferService) ListAutoPager(ctx context.Context, query *requests.RealTimePaymentsTransferListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.RealTimePaymentsTransfer] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
