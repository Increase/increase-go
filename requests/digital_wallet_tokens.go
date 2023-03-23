package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the Card this Digital Wallet Token belongs to.
	CardID fields.Field[string] `json:"card_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// This indicates if payments can be made with the Digital Wallet Token.
	Status fields.Field[DigitalWalletTokenStatus] `json:"status,required"`
	// The digital wallet app being used.
	TokenRequestor fields.Field[DigitalWalletTokenTokenRequestor] `json:"token_requestor,required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type fields.Field[DigitalWalletTokenType] `json:"type,required"`
}

// MarshalJSON serializes DigitalWalletToken into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DigitalWalletToken) String() (result string) {
	return fmt.Sprintf("&DigitalWalletToken{ID:%s CardID:%s CreatedAt:%s Status:%s TokenRequestor:%s Type:%s}", r.ID, r.CardID, r.CreatedAt, r.Status, r.TokenRequestor, r.Type)
}

type DigitalWalletTokenStatus string

const (
	DigitalWalletTokenStatusActive      DigitalWalletTokenStatus = "active"
	DigitalWalletTokenStatusInactive    DigitalWalletTokenStatus = "inactive"
	DigitalWalletTokenStatusSuspended   DigitalWalletTokenStatus = "suspended"
	DigitalWalletTokenStatusDeactivated DigitalWalletTokenStatus = "deactivated"
)

type DigitalWalletTokenTokenRequestor string

const (
	DigitalWalletTokenTokenRequestorApplePay  DigitalWalletTokenTokenRequestor = "apple_pay"
	DigitalWalletTokenTokenRequestorGooglePay DigitalWalletTokenTokenRequestor = "google_pay"
)

type DigitalWalletTokenType string

const (
	DigitalWalletTokenTypeDigitalWalletToken DigitalWalletTokenType = "digital_wallet_token"
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
