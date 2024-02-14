// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// BookkeepingEntrySetService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBookkeepingEntrySetService]
// method instead.
type BookkeepingEntrySetService struct {
	Options []option.RequestOption
}

// NewBookkeepingEntrySetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBookkeepingEntrySetService(opts ...option.RequestOption) (r *BookkeepingEntrySetService) {
	r = &BookkeepingEntrySetService{}
	r.Options = opts
	return
}

// Create a Bookkeeping Entry Set
func (r *BookkeepingEntrySetService) New(ctx context.Context, body BookkeepingEntrySetNewParams, opts ...option.RequestOption) (res *BookkeepingEntrySet, err error) {
	opts = append(r.Options[:], opts...)
	path := "bookkeeping_entry_sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Bookkeeping Entry Set
func (r *BookkeepingEntrySetService) Get(ctx context.Context, bookkeepingEntrySetID string, opts ...option.RequestOption) (res *BookkeepingEntrySet, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("bookkeeping_entry_sets/%s", bookkeepingEntrySetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Bookkeeping Entry Sets
func (r *BookkeepingEntrySetService) List(ctx context.Context, query BookkeepingEntrySetListParams, opts ...option.RequestOption) (res *shared.Page[BookkeepingEntrySet], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bookkeeping_entry_sets"
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

// List Bookkeeping Entry Sets
func (r *BookkeepingEntrySetService) ListAutoPaging(ctx context.Context, query BookkeepingEntrySetListParams, opts ...option.RequestOption) *shared.PageAutoPager[BookkeepingEntrySet] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Entry Sets are accounting entries that are transactionally applied. Your
// compliance setup might require annotating money movements using this API. Learn
// more in our
// [guide to Bookkeeping](https://increase.com/documentation/bookkeeping#bookkeeping).
type BookkeepingEntrySet struct {
	// The entry set identifier.
	ID string `json:"id,required"`
	// When the entry set was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The timestamp of the entry set.
	Date time.Time `json:"date,required" format:"date-time"`
	// The entries.
	Entries []BookkeepingEntrySetEntry `json:"entries,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The transaction identifier, if any.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_entry_set`.
	Type BookkeepingEntrySetType `json:"type,required"`
	JSON bookkeepingEntrySetJSON `json:"-"`
}

// bookkeepingEntrySetJSON contains the JSON metadata for the struct
// [BookkeepingEntrySet]
type bookkeepingEntrySetJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Date           apijson.Field
	Entries        apijson.Field
	IdempotencyKey apijson.Field
	TransactionID  apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BookkeepingEntrySet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookkeepingEntrySetEntry struct {
	// The entry identifier.
	ID string `json:"id,required"`
	// The bookkeeping account impacted by the entry.
	AccountID string `json:"account_id,required"`
	// The amount of the entry in minor units.
	Amount int64                        `json:"amount,required"`
	JSON   bookkeepingEntrySetEntryJSON `json:"-"`
}

// bookkeepingEntrySetEntryJSON contains the JSON metadata for the struct
// [BookkeepingEntrySetEntry]
type bookkeepingEntrySetEntryJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookkeepingEntrySetEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `bookkeeping_entry_set`.
type BookkeepingEntrySetType string

const (
	BookkeepingEntrySetTypeBookkeepingEntrySet BookkeepingEntrySetType = "bookkeeping_entry_set"
)

type BookkeepingEntrySetNewParams struct {
	// The bookkeeping entries.
	Entries param.Field[[]BookkeepingEntrySetNewParamsEntry] `json:"entries,required"`
	// The date of the transaction. Optional if `transaction_id` is provided, in which
	// case we use the `date` of that transaction. Required otherwise.
	Date param.Field[time.Time] `json:"date" format:"date-time"`
	// The identifier of the Transaction related to this entry set, if any.
	TransactionID param.Field[string] `json:"transaction_id"`
}

func (r BookkeepingEntrySetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookkeepingEntrySetNewParamsEntry struct {
	// The identifier for the Bookkeeping Account impacted by this entry.
	AccountID param.Field[string] `json:"account_id,required"`
	// The entry amount in the minor unit of the account currency. For dollars, for
	// example, this is cents. Debit entries have positive amounts; credit entries have
	// negative amounts.
	Amount param.Field[int64] `json:"amount,required"`
}

func (r BookkeepingEntrySetNewParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookkeepingEntrySetListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter to the Bookkeeping Entry Set that maps to this Transaction.
	TransactionID param.Field[string] `query:"transaction_id"`
}

// URLQuery serializes [BookkeepingEntrySetListParams]'s query parameters as
// `url.Values`.
func (r BookkeepingEntrySetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
