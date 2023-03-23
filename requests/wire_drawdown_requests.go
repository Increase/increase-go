package requests

import (
	"fmt"
	"net/url"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type WireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type fields.Field[WireDrawdownRequestType] `json:"type,required"`
	// The Wire drawdown request identifier.
	ID fields.Field[string] `json:"id,required"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber fields.Field[string] `json:"recipient_account_number,required"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber fields.Field[string] `json:"recipient_routing_number,required"`
	// The amount being requested in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency fields.Field[string] `json:"currency,required"`
	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	// The drawdown request's recipient's name.
	RecipientName fields.Field[string] `json:"recipient_name,required,nullable"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 fields.Field[string] `json:"recipient_address_line1,required,nullable"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 fields.Field[string] `json:"recipient_address_line2,required,nullable"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 fields.Field[string] `json:"recipient_address_line3,required,nullable"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission fields.Field[WireDrawdownRequestSubmission] `json:"submission,required,nullable"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionID fields.Field[string] `json:"fulfillment_transaction_id,required,nullable"`
	// The lifecycle status of the drawdown request.
	Status fields.Field[WireDrawdownRequestStatus] `json:"status,required"`
}

// MarshalJSON serializes WireDrawdownRequest into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireDrawdownRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireDrawdownRequest) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequest{Type:%s ID:%s AccountNumberID:%s RecipientAccountNumber:%s RecipientRoutingNumber:%s Amount:%s Currency:%s MessageToRecipient:%s RecipientName:%s RecipientAddressLine1:%s RecipientAddressLine2:%s RecipientAddressLine3:%s Submission:%s FulfillmentTransactionID:%s Status:%s}", r.Type, r.ID, r.AccountNumberID, r.RecipientAccountNumber, r.RecipientRoutingNumber, r.Amount, r.Currency, r.MessageToRecipient, r.RecipientName, r.RecipientAddressLine1, r.RecipientAddressLine2, r.RecipientAddressLine3, r.Submission, r.FulfillmentTransactionID, r.Status)
}

type WireDrawdownRequestType string

const (
	WireDrawdownRequestTypeWireDrawdownRequest WireDrawdownRequestType = "wire_drawdown_request"
)

type WireDrawdownRequestSubmission struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData fields.Field[string] `json:"input_message_accountability_data,required"`
}

// MarshalJSON serializes WireDrawdownRequestSubmission into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *WireDrawdownRequestSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireDrawdownRequestSubmission) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequestSubmission{InputMessageAccountabilityData:%s}", r.InputMessageAccountabilityData)
}

type WireDrawdownRequestStatus string

const (
	WireDrawdownRequestStatusPendingSubmission WireDrawdownRequestStatus = "pending_submission"
	WireDrawdownRequestStatusPendingResponse   WireDrawdownRequestStatus = "pending_response"
	WireDrawdownRequestStatusFulfilled         WireDrawdownRequestStatus = "fulfilled"
	WireDrawdownRequestStatusRefused           WireDrawdownRequestStatus = "refused"
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
