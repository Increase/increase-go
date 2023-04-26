package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type ProgramService struct {
	Options []option.RequestOption
}

func NewProgramService(opts ...option.RequestOption) (r *ProgramService) {
	r = &ProgramService{}
	r.Options = opts
	return
}

// Retrieve a Program
func (r *ProgramService) Get(ctx context.Context, program_id string, opts ...option.RequestOption) (res *responses.Program, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("programs/%s", program_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Programs
func (r *ProgramService) List(ctx context.Context, query *requests.ProgramListParams, opts ...option.RequestOption) (res *responses.Page[responses.Program], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "programs"
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

// List Programs
func (r *ProgramService) ListAutoPager(ctx context.Context, query *requests.ProgramListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Program] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
