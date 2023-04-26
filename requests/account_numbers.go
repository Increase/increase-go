package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type AccountNumberNewParams struct {
	// The Account the Account Number should belong to.
	AccountID field.Field[string] `json:"account_id,required"`
	// The name you choose for the Account Number.
	Name field.Field[string] `json:"name,required"`
}

// MarshalJSON serializes AccountNumberNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AccountNumberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNumberUpdateParams struct {
	// The name you choose for the Account Number.
	Name field.Field[string] `json:"name"`
	// This indicates if transfers can be made to the Account Number.
	Status field.Field[AccountNumberUpdateParamsStatus] `json:"status"`
}

// MarshalJSON serializes AccountNumberUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AccountNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNumberUpdateParamsStatus string

const (
	AccountNumberUpdateParamsStatusActive   AccountNumberUpdateParamsStatus = "active"
	AccountNumberUpdateParamsStatusDisabled AccountNumberUpdateParamsStatus = "disabled"
	AccountNumberUpdateParamsStatusCanceled AccountNumberUpdateParamsStatus = "canceled"
)

type AccountNumberListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status field.Field[AccountNumberListParamsStatus] `query:"status"`
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID field.Field[string]                           `query:"account_id"`
	CreatedAt field.Field[AccountNumberListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes AccountNumberListParams into a url.Values of the query
// parameters associated with this value
func (r AccountNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type AccountNumberListParamsStatus string

const (
	AccountNumberListParamsStatusActive   AccountNumberListParamsStatus = "active"
	AccountNumberListParamsStatusDisabled AccountNumberListParamsStatus = "disabled"
	AccountNumberListParamsStatusCanceled AccountNumberListParamsStatus = "canceled"
)

type AccountNumberListParamsCreatedAt struct {
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

// URLQuery serializes AccountNumberListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r AccountNumberListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
