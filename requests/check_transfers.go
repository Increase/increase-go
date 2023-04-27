package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type CheckTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID field.Field[string] `json:"account_id,required"`
	// The street address of the check's destination.
	AddressLine1 field.Field[string] `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 field.Field[string] `json:"address_line2"`
	// The city of the check's destination.
	AddressCity field.Field[string] `json:"address_city,required"`
	// The state of the check's destination.
	AddressState field.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip field.Field[string] `json:"address_zip,required"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress field.Field[CheckTransferNewParamsReturnAddress] `json:"return_address"`
	// The transfer amount in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The descriptor that will be printed on the memo field on the check.
	Message field.Field[string] `json:"message,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note field.Field[string] `json:"note"`
	// The name that will be printed on the check.
	RecipientName field.Field[string] `json:"recipient_name,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval field.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CheckTransferNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CheckTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckTransferNewParamsReturnAddress struct {
	// The name of the return address.
	Name field.Field[string] `json:"name,required"`
	// The first line of the return address.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the return address.
	Line2 field.Field[string] `json:"line2"`
	// The city of the return address.
	City field.Field[string] `json:"city,required"`
	// The US state of the return address.
	State field.Field[string] `json:"state,required"`
	// The postal code of the return address.
	Zip field.Field[string] `json:"zip,required"`
}

type CheckTransferListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID field.Field[string]                           `query:"account_id"`
	CreatedAt field.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckTransferListParams into a url.Values of the query
// parameters associated with this value
func (r CheckTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CheckTransferListParamsCreatedAt struct {
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

// URLQuery serializes CheckTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r CheckTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
