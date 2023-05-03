package increase

import (
	"context"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// BookkeepingEntryService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBookkeepingEntryService] method
// instead.
type BookkeepingEntryService struct {
	Options []option.RequestOption
}

// NewBookkeepingEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
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

// Entries are T-account entries recording debits and credits.
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
	JSON bookkeepingEntryJSON
}

// bookkeepingEntryJSON contains the JSON metadata for the struct
// [BookkeepingEntry]
type bookkeepingEntryJSON struct {
	AccountID  apijson.Field
	Amount     apijson.Field
	EntrySetID apijson.Field
	ID         apijson.Field
	Type       apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *BookkeepingEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookkeepingEntryType string

const (
	BookkeepingEntryTypeBookkeepingEntry BookkeepingEntryType = "bookkeeping_entry"
)

type BookkeepingEntryListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [BookkeepingEntryListParams]'s query parameters as
// `url.Values`.
func (r BookkeepingEntryListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of Bookkeeping Entry objects
type BookkeepingEntryListResponse struct {
	// The contents of the list.
	Data []BookkeepingEntry `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       bookkeepingEntryListResponseJSON
}

// bookkeepingEntryListResponseJSON contains the JSON metadata for the struct
// [BookkeepingEntryListResponse]
type bookkeepingEntryListResponseJSON struct {
	Data       apijson.Field
	NextCursor apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *BookkeepingEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
