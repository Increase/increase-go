package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type CardAuthorizeParams struct {
	// The authorization amount in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID field.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID field.Field[string] `json:"digital_wallet_token_id"`
}

// MarshalJSON serializes CardAuthorizeParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CardAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardSettlementParams struct {
	// The identifier of the Card to create a settlement on.
	CardID field.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID field.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount field.Field[int64] `json:"amount"`
}

// MarshalJSON serializes CardSettlementParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CardSettlementParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}
