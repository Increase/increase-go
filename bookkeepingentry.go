package increase

import (
	"context"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
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
func (r *BookkeepingEntryService) List(ctx context.Context, query BookkeepingEntryListParams, opts ...option.RequestOption) (res *shared.Page[BookkeepingEntry], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bookkeeping_entries"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *BookkeepingEntryService) ListAutoPaging(ctx context.Context, query BookkeepingEntryListParams, opts ...option.RequestOption) *shared.PageAutoPager[BookkeepingEntry] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type BookkeepingEntry struct {
	// The identifier for the Account the Entry belongs to.
	AccountID string `json:"account_id,required"`
	// The Entry amount in the minor unit of its currency. For dollars, for example,
	// this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier for the Account the Entry belongs to.
	EntrySetID string `json:"entry_set_id,required"`
	// The entry identifier.
	ID string `json:"id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_entry`.
	Type BookkeepingEntryType `json:"type,required"`
	JSON BookkeepingEntryJSON
}

type BookkeepingEntryJSON struct {
	AccountID  apijson.Metadata
	Amount     apijson.Metadata
	EntrySetID apijson.Metadata
	ID         apijson.Metadata
	Type       apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntry using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BookkeepingEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookkeepingEntryType string

const (
	BookkeepingEntryTypeBookkeepingEntry BookkeepingEntryType = "bookkeeping_entry"
)

type BookkeepingEntryListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes BookkeepingEntryListParams into a url.Values of the query
// parameters associated with this value
func (r BookkeepingEntryListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type BookkeepingEntryListResponse struct {
	// The contents of the list.
	Data []BookkeepingEntry `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       BookkeepingEntryListResponseJSON
}

type BookkeepingEntryListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntryListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BookkeepingEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
