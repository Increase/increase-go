package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateACheckDepositParameters struct {
	// The identifier for the Account to deposit the check in.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The currency to use for the deposit.
	Currency fields.Field[string] `json:"currency,required"`
	// The File containing the check's front image.
	FrontImageFileID fields.Field[string] `json:"front_image_file_id,required"`
	// The File containing the check's back image.
	BackImageFileID fields.Field[string] `json:"back_image_file_id,required"`
}

// MarshalJSON serializes CreateACheckDepositParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateACheckDepositParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACheckDepositParameters) String() (result string) {
	return fmt.Sprintf("&CreateACheckDepositParameters{AccountID:%s Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s}", r.AccountID, r.Amount, r.Currency, r.FrontImageFileID, r.BackImageFileID)
}

type CheckDepositListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountID fields.Field[string]                          `query:"account_id"`
	CreatedAt fields.Field[CheckDepositListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckDepositListParams into a url.Values of the query
// parameters associated with this value
func (r *CheckDepositListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckDepositListParams) String() (result string) {
	return fmt.Sprintf("&CheckDepositListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type CheckDepositListParamsCreatedAt struct {
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

// URLQuery serializes CheckDepositListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CheckDepositListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckDepositListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CheckDepositListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}
