package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type AccountTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID field.Field[string] `json:"account_id,required"`
	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The description you choose to give the transfer.
	Description field.Field[string] `json:"description,required"`
	// The identifier for the account that will receive the transfer.
	DestinationAccountID field.Field[string] `json:"destination_account_id,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval field.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes AccountTransferNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountTransferNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type AccountTransferListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Account Transfers to those that originated from the specified Account.
	AccountID field.Field[string]                             `query:"account_id"`
	CreatedAt field.Field[AccountTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes AccountTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type AccountTransferListParamsCreatedAt struct {
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

// URLQuery serializes AccountTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *AccountTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
