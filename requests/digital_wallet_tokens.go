package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type DigitalWalletTokenListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardID    field.Field[string]                                `query:"card_id"`
	CreatedAt field.Field[DigitalWalletTokenListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DigitalWalletTokenListParams into a url.Values of the query
// parameters associated with this value
func (r *DigitalWalletTokenListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DigitalWalletTokenListParamsCreatedAt struct {
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

// URLQuery serializes DigitalWalletTokenListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *DigitalWalletTokenListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
