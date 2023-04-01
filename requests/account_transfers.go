package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateAnAccountTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description you choose to give the transfer.
	Description fields.Field[string] `json:"description,required"`
	// The identifier for the account that will receive the transfer.
	DestinationAccountID fields.Field[string] `json:"destination_account_id,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CreateAnAccountTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnAccountTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountTransferParameters{AccountID:%s Amount:%s Description:%s DestinationAccountID:%s RequireApproval:%s}", r.AccountID, r.Amount, r.Description, r.DestinationAccountID, r.RequireApproval)
}

type AccountTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Account Transfers to those that originated from the specified Account.
	AccountID fields.Field[string]                             `query:"account_id"`
	CreatedAt fields.Field[AccountTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes AccountTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountTransferListParams) String() (result string) {
	return fmt.Sprintf("&AccountTransferListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type AccountTransferListParamsCreatedAt struct {
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

// URLQuery serializes AccountTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *AccountTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&AccountTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}
