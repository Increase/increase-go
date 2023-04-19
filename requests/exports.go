package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type ExportNewParams struct {
	// The type of Export to create.
	Category field.Field[ExportNewParamsCategory] `json:"category,required"`
	// Options for the created export. Required if `category` is equal to
	// `transaction_csv`.
	TransactionCsv field.Field[ExportNewParamsTransactionCsv] `json:"transaction_csv"`
	// Options for the created export. Required if `category` is equal to
	// `balance_csv`.
	BalanceCsv field.Field[ExportNewParamsBalanceCsv] `json:"balance_csv"`
}

// MarshalJSON serializes ExportNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r ExportNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ExportNewParamsCategory string

const (
	ExportNewParamsCategoryTransactionCsv ExportNewParamsCategory = "transaction_csv"
	ExportNewParamsCategoryBalanceCsv     ExportNewParamsCategory = "balance_csv"
)

type ExportNewParamsTransactionCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID field.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt field.Field[ExportNewParamsTransactionCsvCreatedAt] `json:"created_at"`
}

type ExportNewParamsTransactionCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

type ExportNewParamsBalanceCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID field.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt field.Field[ExportNewParamsBalanceCsvCreatedAt] `json:"created_at"`
}

type ExportNewParamsBalanceCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

type ExportListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes ExportListParams into a url.Values of the query parameters
// associated with this value
func (r ExportListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
