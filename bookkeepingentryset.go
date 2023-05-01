package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type BookkeepingEntrySetService struct {
	Options []option.RequestOption
}

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

type BookkeepingEntrySet struct {
	// The entry set identifier.
	ID string `json:"id,required"`
	// The transaction identifier, if any.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The timestamp of the entry set.
	Date time.Time `json:"date,required" format:"date-time"`
	// The entries
	Entries []BookkeepingEntrySetEntries `json:"entries,required"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_entry_set`.
	Type BookkeepingEntrySetType `json:"type,required"`
	JSON BookkeepingEntrySetJSON
}

type BookkeepingEntrySetJSON struct {
	ID            apijson.Metadata
	TransactionID apijson.Metadata
	Date          apijson.Metadata
	Entries       apijson.Metadata
	Type          apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntrySet using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BookkeepingEntrySet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookkeepingEntrySetEntries struct {
	// The bookkeeping account impacted by the entry.
	AccountID string `json:"account_id,required"`
	// The amount of the entry in minor units.
	Amount int64 `json:"amount,required"`
	// The entry identifier.
	ID   string `json:"id,required"`
	JSON BookkeepingEntrySetEntriesJSON
}

type BookkeepingEntrySetEntriesJSON struct {
	AccountID apijson.Metadata
	Amount    apijson.Metadata
	ID        apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntrySetEntries
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BookkeepingEntrySetEntries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookkeepingEntrySetType string

const (
	BookkeepingEntrySetTypeBookkeepingEntrySet BookkeepingEntrySetType = "bookkeeping_entry_set"
)

type BookkeepingEntrySetNewParams struct {
	// The date of the transaction. If `transaction_id` is provided, this must match
	// the `created_at` field on that resource.
	Date field.Field[time.Time] `json:"date" format:"date-time"`
	// The identifier of the Transaction related to this entry set, if any.
	TransactionID field.Field[string] `json:"transaction_id"`
	// The bookkeeping entries.
	Entries field.Field[[]BookkeepingEntrySetNewParamsEntries] `json:"entries,required"`
}

// MarshalJSON serializes BookkeepingEntrySetNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r BookkeepingEntrySetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookkeepingEntrySetNewParamsEntries struct {
	// The identifier for the Bookkeeping Account impacted by this entry.
	AccountID field.Field[string] `json:"account_id,required"`
	// The entry amount in the minor unit of the account currency. For dollars, for
	// example, this is cents. Debit entries have positive amounts; credit entries have
	// negative amounts.
	Amount field.Field[int64] `json:"amount,required"`
}
