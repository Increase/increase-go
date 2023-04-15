package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
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
	return pjson.MarshalRoot(r)
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
	return pjson.MarshalRoot(r)
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
	AccountID field.Field[string] `query:"account_id"`
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
