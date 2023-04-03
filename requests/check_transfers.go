package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateACheckTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The street address of the check's destination.
	AddressLine1 fields.Field[string] `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 fields.Field[string] `json:"address_line2"`
	// The city of the check's destination.
	AddressCity fields.Field[string] `json:"address_city,required"`
	// The state of the check's destination.
	AddressState fields.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip fields.Field[string] `json:"address_zip,required"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress fields.Field[CreateACheckTransferParametersReturnAddress] `json:"return_address"`
	// The transfer amount in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The descriptor that will be printed on the memo field on the check.
	Message fields.Field[string] `json:"message,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note fields.Field[string] `json:"note"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CreateACheckTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateACheckTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACheckTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateACheckTransferParameters{AccountID:%s AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s ReturnAddress:%s Amount:%s Message:%s Note:%s RecipientName:%s RequireApproval:%s}", r.AccountID, r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.ReturnAddress, r.Amount, r.Message, r.Note, r.RecipientName, r.RequireApproval)
}

type CreateACheckTransferParametersReturnAddress struct {
	// The name of the return address.
	Name fields.Field[string] `json:"name,required"`
	// The first line of the return address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the return address.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the return address.
	City fields.Field[string] `json:"city,required"`
	// The US state of the return address.
	State fields.Field[string] `json:"state,required"`
	// The postal code of the return address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateACheckTransferParametersReturnAddress) String() (result string) {
	return fmt.Sprintf("&CreateACheckTransferParametersReturnAddress{Name:%s Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Name, r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CheckTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID fields.Field[string]                           `query:"account_id"`
	CreatedAt fields.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *CheckTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckTransferListParams) String() (result string) {
	return fmt.Sprintf("&CheckTransferListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type CheckTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes CheckTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CheckTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CheckTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}
