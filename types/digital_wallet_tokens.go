package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/pagination"
)

type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	ID *string `pjson:"id"`
	// The identifier for the Card this Digital Wallet Token belongs to.
	CardID *string `pjson:"card_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt *string `pjson:"created_at"`
	// This indicates if payments can be made with the Digital Wallet Token.
	Status *DigitalWalletTokenStatus `pjson:"status"`
	// The digital wallet app being used.
	TokenRequestor *DigitalWalletTokenTokenRequestor `pjson:"token_requestor"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type       *DigitalWalletTokenType `pjson:"type"`
	jsonFields map[string]interface{}  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DigitalWalletToken using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *DigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DigitalWalletToken into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *DigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Digital Wallet Token identifier.
func (r *DigitalWalletToken) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the Card this Digital Wallet Token belongs to.
func (r *DigitalWalletToken) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card was created.
func (r *DigitalWalletToken) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This indicates if payments can be made with the Digital Wallet Token.
func (r *DigitalWalletToken) GetStatus() (Status DigitalWalletTokenStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The digital wallet app being used.
func (r *DigitalWalletToken) GetTokenRequestor() (TokenRequestor DigitalWalletTokenTokenRequestor) {
	if r != nil && r.TokenRequestor != nil {
		TokenRequestor = *r.TokenRequestor
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `digital_wallet_token`.
func (r *DigitalWalletToken) GetType() (Type DigitalWalletTokenType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r DigitalWalletToken) String() (result string) {
	return fmt.Sprintf("&DigitalWalletToken{ID:%s CardID:%s CreatedAt:%s Status:%s TokenRequestor:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.CardID), core.FmtP(r.CreatedAt), core.FmtP(r.Status), core.FmtP(r.TokenRequestor), core.FmtP(r.Type))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardID     *string                                 `query:"card_id"`
	CreatedAt  *DigitalWalletTokensListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DigitalWalletTokenListParams
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *DigitalWalletTokenListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DigitalWalletTokenListParams into an array of bytes using
// the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DigitalWalletTokenListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return the page of entries after this one.
func (r *DigitalWalletTokenListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *DigitalWalletTokenListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Digital Wallet Tokens to ones belonging to the specified Card.
func (r *DigitalWalletTokenListParams) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

func (r *DigitalWalletTokenListParams) GetCreatedAt() (CreatedAt DigitalWalletTokensListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r DigitalWalletTokenListParams) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokenListParams{Cursor:%s Limit:%s CardID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.CardID), r.CreatedAt)
}

type DigitalWalletTokensListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DigitalWalletTokensListParamsCreatedAt using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *DigitalWalletTokensListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DigitalWalletTokensListParamsCreatedAt into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *DigitalWalletTokensListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *DigitalWalletTokensListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *DigitalWalletTokensListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *DigitalWalletTokensListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *DigitalWalletTokensListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r DigitalWalletTokensListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokensListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type DigitalWalletTokenList struct {
	// The contents of the list.
	Data *[]DigitalWalletToken `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DigitalWalletTokenList using
// the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *DigitalWalletTokenList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DigitalWalletTokenList into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DigitalWalletTokenList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The contents of the list.
func (r *DigitalWalletTokenList) GetData() (Data []DigitalWalletToken) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *DigitalWalletTokenList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r DigitalWalletTokenList) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokenList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type DigitalWalletTokensPage struct {
	*pagination.Page[DigitalWalletToken]
}

func (r *DigitalWalletTokensPage) DigitalWalletToken() *DigitalWalletToken {
	return r.Current()
}

func (r *DigitalWalletTokensPage) NextPage() (*DigitalWalletTokensPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &DigitalWalletTokensPage{page}, nil
	}
}
