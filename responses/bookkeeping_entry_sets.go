package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

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
	Raw           []byte
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
	Raw       []byte
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
