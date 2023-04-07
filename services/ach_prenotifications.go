package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type ACHPrenotificationService struct {
	Options []option.RequestOption
}

func NewACHPrenotificationService(opts ...option.RequestOption) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Options = opts
	return
}

// Create an ACH Prenotification
func (r *ACHPrenotificationService) New(ctx context.Context, body *requests.ACHPrenotificationNewParams, opts ...option.RequestOption) (res *responses.ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_prenotifications"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, ach_prenotification_id string, opts ...option.RequestOption) (res *responses.ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_prenotifications/%s", ach_prenotification_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query *requests.ACHPrenotificationListParams, opts ...option.RequestOption) (res *responses.Page[responses.ACHPrenotification], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_prenotifications"
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

// List ACH Prenotifications
func (r *ACHPrenotificationService) ListAutoPager(ctx context.Context, query *requests.ACHPrenotificationListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.ACHPrenotification] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
