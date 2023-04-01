package requests

import (
	"fmt"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type SimulateAnAuthorizationOnACardParameters struct {
	// The authorization amount in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID fields.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id"`
}

// MarshalJSON serializes SimulateAnAuthorizationOnACardParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SimulateAnAuthorizationOnACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAnAuthorizationOnACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnAuthorizationOnACardParameters{Amount:%s CardID:%s DigitalWalletTokenID:%s}", r.Amount, r.CardID, r.DigitalWalletTokenID)
}

type SimulateSettlingACardAuthorizationParameters struct {
	// The identifier of the Card to create a settlement on.
	CardID fields.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID fields.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount fields.Field[int64] `json:"amount"`
}

// MarshalJSON serializes SimulateSettlingACardAuthorizationParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateSettlingACardAuthorizationParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateSettlingACardAuthorizationParameters) String() (result string) {
	return fmt.Sprintf("&SimulateSettlingACardAuthorizationParameters{CardID:%s PendingTransactionID:%s Amount:%s}", r.CardID, r.PendingTransactionID, r.Amount)
}
