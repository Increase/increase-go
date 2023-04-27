package requests

import (
	"net/url"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type WireDrawdownRequestNewParams struct {
	// The Account Number to which the recipient should send funds.
	AccountNumberID field.Field[string] `json:"account_number_id,required"`
	// The amount requested from the recipient, in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// A message the recipient will see as part of the request.
	MessageToRecipient field.Field[string] `json:"message_to_recipient,required"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber field.Field[string] `json:"recipient_account_number,required"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber field.Field[string] `json:"recipient_routing_number,required"`
	// The drawdown request's recipient's name.
	RecipientName field.Field[string] `json:"recipient_name,required"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 field.Field[string] `json:"recipient_address_line1"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 field.Field[string] `json:"recipient_address_line2"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 field.Field[string] `json:"recipient_address_line3"`
}

// MarshalJSON serializes WireDrawdownRequestNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r WireDrawdownRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes WireDrawdownRequestListParams into a url.Values of the query
// parameters associated with this value
func (r WireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
