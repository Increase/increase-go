package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulationCardAuthorizeParams struct {
	// The authorization amount in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID field.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID field.Field[string] `json:"digital_wallet_token_id"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID field.Field[string] `json:"event_subscription_id"`
}

// MarshalJSON serializes SimulationCardAuthorizeParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationCardAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type SimulationCardSettlementParams struct {
	// The identifier of the Card to create a settlement on.
	CardID field.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID field.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount field.Field[int64] `json:"amount"`
}

// MarshalJSON serializes SimulationCardSettlementParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationCardSettlementParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}
