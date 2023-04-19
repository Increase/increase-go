package responses

import (
	pjson "github.com/increase/increase-go/core/json"
)

type BalanceLookupLookupResponse struct {
	// The identifier for the account for which the balance was queried.
	AccountID string `json:"account_id,required"`
	// The Account's current balance, representing the sum of all posted Transactions
	// on the Account.
	CurrentBalance int64 `json:"current_balance,required"`
	// The Account's available balance, representing the current balance less any open
	// Pending Transactions on the Account.
	AvailableBalance int64 `json:"available_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `balance_lookup`.
	Type BalanceLookupLookupResponseType `json:"type,required"`
	JSON BalanceLookupLookupResponseJSON
}

type BalanceLookupLookupResponseJSON struct {
	AccountID        pjson.Metadata
	CurrentBalance   pjson.Metadata
	AvailableBalance pjson.Metadata
	Type             pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceLookupLookupResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BalanceLookupLookupResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BalanceLookupLookupResponseType string

const (
	BalanceLookupLookupResponseTypeBalanceLookup BalanceLookupLookupResponseType = "balance_lookup"
)
