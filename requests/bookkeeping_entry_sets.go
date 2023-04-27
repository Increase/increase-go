package requests

import (
	"time"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
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
