// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
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

// Entry Sets are accounting entries that are transactionally applied.
type BookkeepingEntrySet struct {
	// The entry set identifier.
	ID string `json:"id,required"`
	// The timestamp of the entry set.
	Date time.Time `json:"date,required" format:"date-time"`
	// The entries
	Entries []BookkeepingEntrySetEntry `json:"entries,required"`
	// The transaction identifier, if any.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_entry_set`.
	Type BookkeepingEntrySetType `json:"type,required"`
	JSON bookkeepingEntrySetJSON
}

// bookkeepingEntrySetJSON contains the JSON metadata for the struct
// [BookkeepingEntrySet]
type bookkeepingEntrySetJSON struct {
	ID            apijson.Field
	Date          apijson.Field
	Entries       apijson.Field
	TransactionID apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	Amount int64 `json:"amount,required"`
	JSON   bookkeepingEntrySetEntryJSON
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
	// The date of the transaction. If `transaction_id` is provided, this must match
	// the `created_at` field on that resource.
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
