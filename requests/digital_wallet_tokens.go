package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/fields"
	"github.com/increase/increase-go/core/query"
)

type DigitalWalletTokenListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardID    fields.Field[string]                                `query:"card_id"`
	CreatedAt fields.Field[DigitalWalletTokenListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DigitalWalletTokenListParams into a url.Values of the query
// parameters associated with this value
func (r *DigitalWalletTokenListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DigitalWalletTokenListParams) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokenListParams{Cursor:%s Limit:%s CardID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.CardID, r.CreatedAt)
}

type DigitalWalletTokenListParamsCreatedAt struct {
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

// URLQuery serializes DigitalWalletTokenListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *DigitalWalletTokenListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DigitalWalletTokenListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokenListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}
