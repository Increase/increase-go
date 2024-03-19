// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type Page[T any] struct {
	Data []T `json:"data"`
	// A pointer to a place in the list.
	NextCursor string   `json:"next_cursor,nullable"`
	JSON       pageJSON `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// pageJSON contains the JSON metadata for the struct [Page[T]]
type pageJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageJSON) RawJSON() string {
	return r.raw
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("cursor", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Page[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}
