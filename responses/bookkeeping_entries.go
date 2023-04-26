package responses

import (
	apijson "github.com/increase/increase-go/core/json"
)

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
	Raw        []byte
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
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntryListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BookkeepingEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
