package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateAWireDrawdownRequestParameters struct {
	// The Account Number to which the recipient should send funds.
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The amount requested from the recipient, in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// A message the recipient will see as part of the request.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber fields.Field[string] `json:"recipient_account_number,required"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber fields.Field[string] `json:"recipient_routing_number,required"`
	// The drawdown request's recipient's name.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 fields.Field[string] `json:"recipient_address_line1"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 fields.Field[string] `json:"recipient_address_line2"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 fields.Field[string] `json:"recipient_address_line3"`
}

// MarshalJSON serializes CreateAWireDrawdownRequestParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAWireDrawdownRequestParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAWireDrawdownRequestParameters) String() (result string) {
	return fmt.Sprintf("&CreateAWireDrawdownRequestParameters{AccountNumberID:%s Amount:%s MessageToRecipient:%s RecipientAccountNumber:%s RecipientRoutingNumber:%s RecipientName:%s RecipientAddressLine1:%s RecipientAddressLine2:%s RecipientAddressLine3:%s}", r.AccountNumberID, r.Amount, r.MessageToRecipient, r.RecipientAccountNumber, r.RecipientRoutingNumber, r.RecipientName, r.RecipientAddressLine1, r.RecipientAddressLine2, r.RecipientAddressLine3)
}

type WireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
}

// URLQuery serializes WireDrawdownRequestListParams into a url.Values of the query
// parameters associated with this value
func (r *WireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r WireDrawdownRequestListParams) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequestListParams{Cursor:%s Limit:%s}", r.Cursor, r.Limit)
}
