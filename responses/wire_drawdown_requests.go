package responses

import (
	pjson "github.com/increase/increase-go/core/json"
)

type WireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type WireDrawdownRequestType `json:"type,required"`
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberID string `json:"account_number_id,required"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber string `json:"recipient_account_number,required"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber string `json:"recipient_routing_number,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The drawdown request's recipient's name.
	RecipientName string `json:"recipient_name,required,nullable"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 string `json:"recipient_address_line1,required,nullable"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 string `json:"recipient_address_line2,required,nullable"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 string `json:"recipient_address_line3,required,nullable"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission WireDrawdownRequestSubmission `json:"submission,required,nullable"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionID string `json:"fulfillment_transaction_id,required,nullable"`
	// The lifecycle status of the drawdown request.
	Status WireDrawdownRequestStatus `json:"status,required"`
	JSON   WireDrawdownRequestJSON
}

type WireDrawdownRequestJSON struct {
	Type                     pjson.Metadata
	ID                       pjson.Metadata
	AccountNumberID          pjson.Metadata
	RecipientAccountNumber   pjson.Metadata
	RecipientRoutingNumber   pjson.Metadata
	Amount                   pjson.Metadata
	Currency                 pjson.Metadata
	MessageToRecipient       pjson.Metadata
	RecipientName            pjson.Metadata
	RecipientAddressLine1    pjson.Metadata
	RecipientAddressLine2    pjson.Metadata
	RecipientAddressLine3    pjson.Metadata
	Submission               pjson.Metadata
	FulfillmentTransactionID pjson.Metadata
	Status                   pjson.Metadata
	Raw                      []byte
	Extras                   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireDrawdownRequest using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireDrawdownRequestType string

const (
	WireDrawdownRequestTypeWireDrawdownRequest WireDrawdownRequestType = "wire_drawdown_request"
)

type WireDrawdownRequestSubmission struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	JSON                           WireDrawdownRequestSubmissionJSON
}

type WireDrawdownRequestSubmissionJSON struct {
	InputMessageAccountabilityData pjson.Metadata
	Raw                            []byte
	Extras                         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireDrawdownRequestSubmission
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireDrawdownRequestSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireDrawdownRequestStatus string

const (
	WireDrawdownRequestStatusPendingSubmission WireDrawdownRequestStatus = "pending_submission"
	WireDrawdownRequestStatusPendingResponse   WireDrawdownRequestStatus = "pending_response"
	WireDrawdownRequestStatusFulfilled         WireDrawdownRequestStatus = "fulfilled"
	WireDrawdownRequestStatusRefused           WireDrawdownRequestStatus = "refused"
)

type WireDrawdownRequestListResponse struct {
	// The contents of the list.
	Data []WireDrawdownRequest `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       WireDrawdownRequestListResponseJSON
}

type WireDrawdownRequestListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireDrawdownRequestListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *WireDrawdownRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
