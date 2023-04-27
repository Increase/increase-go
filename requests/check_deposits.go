package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type CheckDepositNewParams struct {
	// The identifier for the Account to deposit the check in.
	AccountID field.Field[string] `json:"account_id,required"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The currency to use for the deposit.
	Currency field.Field[string] `json:"currency,required"`
	// The File containing the check's front image.
	FrontImageFileID field.Field[string] `json:"front_image_file_id,required"`
	// The File containing the check's back image.
	BackImageFileID field.Field[string] `json:"back_image_file_id,required"`
}

// MarshalJSON serializes CheckDepositNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CheckDepositNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckDepositListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountID field.Field[string]                          `query:"account_id"`
	CreatedAt field.Field[CheckDepositListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckDepositListParams into a url.Values of the query
// parameters associated with this value
func (r CheckDepositListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CheckDepositListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes CheckDepositListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r CheckDepositListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
