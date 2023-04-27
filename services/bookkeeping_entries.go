package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type BookkeepingEntryService struct {
	Options []option.RequestOption
}

func NewBookkeepingEntryService(opts ...option.RequestOption) (r *BookkeepingEntryService) {
	r = &BookkeepingEntryService{}
	r.Options = opts
	return
}

// List Bookkeeping Entries
func (r *BookkeepingEntryService) List(ctx context.Context, query requests.BookkeepingEntryListParams, opts ...option.RequestOption) (res *responses.Page[responses.BookkeepingEntry], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bookkeeping_entries"
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

// List Bookkeeping Entries
func (r *BookkeepingEntryService) ListAutoPaging(ctx context.Context, query requests.BookkeepingEntryListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.BookkeepingEntry] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
