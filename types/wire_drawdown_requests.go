package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type WireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type *WireDrawdownRequestType `pjson:"type"`
	// The Wire drawdown request identifier.
	ID *string `pjson:"id"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberID *string `pjson:"account_number_id"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber *string `pjson:"recipient_account_number"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber *string `pjson:"recipient_routing_number"`
	// The amount being requested in cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency *string `pjson:"currency"`
	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient *string `pjson:"message_to_recipient"`
	// The drawdown request's recipient's name.
	RecipientName *string `pjson:"recipient_name"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 *string `pjson:"recipient_address_line1"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 *string `pjson:"recipient_address_line2"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 *string `pjson:"recipient_address_line3"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission *WireDrawdownRequestSubmission `pjson:"submission"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionID *string `pjson:"fulfillment_transaction_id"`
	// The lifecycle status of the drawdown request.
	Status     *WireDrawdownRequestStatus `pjson:"status"`
	jsonFields map[string]interface{}     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireDrawdownRequest using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireDrawdownRequest into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireDrawdownRequest) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// A constant representing the object's type. For this resource it will always be
// `wire_drawdown_request`.
func (r *WireDrawdownRequest) GetType() (Type WireDrawdownRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// The Wire drawdown request identifier.
func (r *WireDrawdownRequest) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The Account Number to which the recipient of this request is being requested to
// send funds.
func (r *WireDrawdownRequest) GetAccountNumberID() (AccountNumberID string) {
	if r != nil && r.AccountNumberID != nil {
		AccountNumberID = *r.AccountNumberID
	}
	return
}

// The drawdown request's recipient's account number.
func (r *WireDrawdownRequest) GetRecipientAccountNumber() (RecipientAccountNumber string) {
	if r != nil && r.RecipientAccountNumber != nil {
		RecipientAccountNumber = *r.RecipientAccountNumber
	}
	return
}

// The drawdown request's recipient's routing number.
func (r *WireDrawdownRequest) GetRecipientRoutingNumber() (RecipientRoutingNumber string) {
	if r != nil && r.RecipientRoutingNumber != nil {
		RecipientRoutingNumber = *r.RecipientRoutingNumber
	}
	return
}

// The amount being requested in cents.
func (r *WireDrawdownRequest) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
// requested. Will always be "USD".
func (r *WireDrawdownRequest) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The message the recipient will see as part of the drawdown request.
func (r *WireDrawdownRequest) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The drawdown request's recipient's name.
func (r *WireDrawdownRequest) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// Line 1 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine1() (RecipientAddressLine1 string) {
	if r != nil && r.RecipientAddressLine1 != nil {
		RecipientAddressLine1 = *r.RecipientAddressLine1
	}
	return
}

// Line 2 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine2() (RecipientAddressLine2 string) {
	if r != nil && r.RecipientAddressLine2 != nil {
		RecipientAddressLine2 = *r.RecipientAddressLine2
	}
	return
}

// Line 3 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine3() (RecipientAddressLine3 string) {
	if r != nil && r.RecipientAddressLine3 != nil {
		RecipientAddressLine3 = *r.RecipientAddressLine3
	}
	return
}

// After the drawdown request is submitted to Fedwire, this will contain
// supplemental details.
func (r *WireDrawdownRequest) GetSubmission() (Submission WireDrawdownRequestSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the recipient fulfills the drawdown request by sending funds, then this will
// be the identifier of the corresponding Transaction.
func (r *WireDrawdownRequest) GetFulfillmentTransactionID() (FulfillmentTransactionID string) {
	if r != nil && r.FulfillmentTransactionID != nil {
		FulfillmentTransactionID = *r.FulfillmentTransactionID
	}
	return
}

// The lifecycle status of the drawdown request.
func (r *WireDrawdownRequest) GetStatus() (Status WireDrawdownRequestStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r WireDrawdownRequest) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequest{Type:%s ID:%s AccountNumberID:%s RecipientAccountNumber:%s RecipientRoutingNumber:%s Amount:%s Currency:%s MessageToRecipient:%s RecipientName:%s RecipientAddressLine1:%s RecipientAddressLine2:%s RecipientAddressLine3:%s Submission:%s FulfillmentTransactionID:%s Status:%s}", core.FmtP(r.Type), core.FmtP(r.ID), core.FmtP(r.AccountNumberID), core.FmtP(r.RecipientAccountNumber), core.FmtP(r.RecipientRoutingNumber), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MessageToRecipient), core.FmtP(r.RecipientName), core.FmtP(r.RecipientAddressLine1), core.FmtP(r.RecipientAddressLine2), core.FmtP(r.RecipientAddressLine3), r.Submission, core.FmtP(r.FulfillmentTransactionID), core.FmtP(r.Status))
}

type WireDrawdownRequestType string

const (
	WireDrawdownRequestTypeWireDrawdownRequest WireDrawdownRequestType = "wire_drawdown_request"
)

type WireDrawdownRequestSubmission struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData *string                `pjson:"input_message_accountability_data"`
	jsonFields                     map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireDrawdownRequestSubmission
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireDrawdownRequestSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireDrawdownRequestSubmission into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *WireDrawdownRequestSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The input message accountability data (IMAD) uniquely identifying the submission
// with Fedwire.
func (r *WireDrawdownRequestSubmission) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r WireDrawdownRequestSubmission) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequestSubmission{InputMessageAccountabilityData:%s}", core.FmtP(r.InputMessageAccountabilityData))
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
	AccountNumberID *string `pjson:"account_number_id"`
	// The amount requested from the recipient, in cents.
	Amount *int64 `pjson:"amount"`
	// A message the recipient will see as part of the request.
	MessageToRecipient *string `pjson:"message_to_recipient"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber *string `pjson:"recipient_account_number"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber *string `pjson:"recipient_routing_number"`
	// The drawdown request's recipient's name.
	RecipientName *string `pjson:"recipient_name"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 *string `pjson:"recipient_address_line1"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 *string `pjson:"recipient_address_line2"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 *string                `pjson:"recipient_address_line3"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAWireDrawdownRequestParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateAWireDrawdownRequestParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAWireDrawdownRequestParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAWireDrawdownRequestParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account Number to which the recipient should send funds.
func (r *CreateAWireDrawdownRequestParameters) GetAccountNumberID() (AccountNumberID string) {
	if r != nil && r.AccountNumberID != nil {
		AccountNumberID = *r.AccountNumberID
	}
	return
}

// The amount requested from the recipient, in cents.
func (r *CreateAWireDrawdownRequestParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// A message the recipient will see as part of the request.
func (r *CreateAWireDrawdownRequestParameters) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The drawdown request's recipient's account number.
func (r *CreateAWireDrawdownRequestParameters) GetRecipientAccountNumber() (RecipientAccountNumber string) {
	if r != nil && r.RecipientAccountNumber != nil {
		RecipientAccountNumber = *r.RecipientAccountNumber
	}
	return
}

// The drawdown request's recipient's routing number.
func (r *CreateAWireDrawdownRequestParameters) GetRecipientRoutingNumber() (RecipientRoutingNumber string) {
	if r != nil && r.RecipientRoutingNumber != nil {
		RecipientRoutingNumber = *r.RecipientRoutingNumber
	}
	return
}

// The drawdown request's recipient's name.
func (r *CreateAWireDrawdownRequestParameters) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// Line 1 of the drawdown request's recipient's address.
func (r *CreateAWireDrawdownRequestParameters) GetRecipientAddressLine1() (RecipientAddressLine1 string) {
	if r != nil && r.RecipientAddressLine1 != nil {
		RecipientAddressLine1 = *r.RecipientAddressLine1
	}
	return
}

// Line 2 of the drawdown request's recipient's address.
func (r *CreateAWireDrawdownRequestParameters) GetRecipientAddressLine2() (RecipientAddressLine2 string) {
	if r != nil && r.RecipientAddressLine2 != nil {
		RecipientAddressLine2 = *r.RecipientAddressLine2
	}
	return
}

// Line 3 of the drawdown request's recipient's address.
func (r *CreateAWireDrawdownRequestParameters) GetRecipientAddressLine3() (RecipientAddressLine3 string) {
	if r != nil && r.RecipientAddressLine3 != nil {
		RecipientAddressLine3 = *r.RecipientAddressLine3
	}
	return
}

func (r CreateAWireDrawdownRequestParameters) String() (result string) {
	return fmt.Sprintf("&CreateAWireDrawdownRequestParameters{AccountNumberID:%s Amount:%s MessageToRecipient:%s RecipientAccountNumber:%s RecipientRoutingNumber:%s RecipientName:%s RecipientAddressLine1:%s RecipientAddressLine2:%s RecipientAddressLine3:%s}", core.FmtP(r.AccountNumberID), core.FmtP(r.Amount), core.FmtP(r.MessageToRecipient), core.FmtP(r.RecipientAccountNumber), core.FmtP(r.RecipientRoutingNumber), core.FmtP(r.RecipientName), core.FmtP(r.RecipientAddressLine1), core.FmtP(r.RecipientAddressLine2), core.FmtP(r.RecipientAddressLine3))
}

type WireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                 `query:"limit"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireDrawdownRequestListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireDrawdownRequestListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireDrawdownRequestListParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *WireDrawdownRequestListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes WireDrawdownRequestListParams into a url.Values of the query
// parameters associated with this value
func (r *WireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *WireDrawdownRequestListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *WireDrawdownRequestListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r WireDrawdownRequestListParams) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequestListParams{Cursor:%s Limit:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit))
}

type WireDrawdownRequestList struct {
	// The contents of the list.
	Data *[]WireDrawdownRequest `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireDrawdownRequestList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireDrawdownRequestList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireDrawdownRequestList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireDrawdownRequestList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes WireDrawdownRequestList into a url.Values of the query
// parameters associated with this value
func (r *WireDrawdownRequestList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *WireDrawdownRequestList) GetData() (Data []WireDrawdownRequest) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *WireDrawdownRequestList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r WireDrawdownRequestList) String() (result string) {
	return fmt.Sprintf("&WireDrawdownRequestList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type WireDrawdownRequestsPage struct {
	*pagination.Page[WireDrawdownRequest]
}

func (r *WireDrawdownRequestsPage) WireDrawdownRequest() *WireDrawdownRequest {
	return r.Current()
}

func (r *WireDrawdownRequestsPage) NextPage() (*WireDrawdownRequestsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &WireDrawdownRequestsPage{page}, nil
	}
}
