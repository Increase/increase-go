package responses

import (
	pjson "github.com/increase/increase-go/core/json"
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
	AccountID  pjson.Metadata
	Amount     pjson.Metadata
	EntrySetID pjson.Metadata
	ID         pjson.Metadata
	Type       pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntry using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BookkeepingEntry) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingEntryListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BookkeepingEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
