package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type InboundWireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_drawdown_request`.
	Type *InboundWireDrawdownRequestType `pjson:"type"`
	// The Wire drawdown request identifier.
	ID *string `pjson:"id"`
	// The Account Number from which the recipient of this request is being requested
	// to send funds.
	RecipientAccountNumberID *string `pjson:"recipient_account_number_id"`
	// The drawdown request's originator's account number.
	OriginatorAccountNumber *string `pjson:"originator_account_number"`
	// The drawdown request's originator's routing number.
	OriginatorRoutingNumber *string `pjson:"originator_routing_number"`
	// The drawdown request's beneficiary's account number.
	BeneficiaryAccountNumber *string `pjson:"beneficiary_account_number"`
	// The drawdown request's beneficiary's routing number.
	BeneficiaryRoutingNumber *string `pjson:"beneficiary_routing_number"`
	// The amount being requested in cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency *string `pjson:"currency"`
	// A message from the drawdown request's originator.
	MessageToRecipient *string `pjson:"message_to_recipient"`
	// Line 1 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine1 *string `pjson:"originator_to_beneficiary_information_line1"`
	// Line 2 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine2 *string `pjson:"originator_to_beneficiary_information_line2"`
	// Line 3 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine3 *string `pjson:"originator_to_beneficiary_information_line3"`
	// Line 4 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine4 *string `pjson:"originator_to_beneficiary_information_line4"`
	// The drawdown request's originator's name.
	OriginatorName *string `pjson:"originator_name"`
	// Line 1 of the drawdown request's originator's address.
	OriginatorAddressLine1 *string `pjson:"originator_address_line1"`
	// Line 2 of the drawdown request's originator's address.
	OriginatorAddressLine2 *string `pjson:"originator_address_line2"`
	// Line 3 of the drawdown request's originator's address.
	OriginatorAddressLine3 *string `pjson:"originator_address_line3"`
	// The drawdown request's beneficiary's name.
	BeneficiaryName *string `pjson:"beneficiary_name"`
	// Line 1 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine1 *string `pjson:"beneficiary_address_line1"`
	// Line 2 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine2 *string `pjson:"beneficiary_address_line2"`
	// Line 3 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine3 *string                `pjson:"beneficiary_address_line3"`
	jsonFields              map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into InboundWireDrawdownRequest
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundWireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundWireDrawdownRequest into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *InboundWireDrawdownRequest) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_drawdown_request`.
func (r InboundWireDrawdownRequest) GetType() (Type InboundWireDrawdownRequestType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

// The Wire drawdown request identifier.
func (r InboundWireDrawdownRequest) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The Account Number from which the recipient of this request is being requested
// to send funds.
func (r InboundWireDrawdownRequest) GetRecipientAccountNumberID() (RecipientAccountNumberID string) {
	if r.RecipientAccountNumberID != nil {
		RecipientAccountNumberID = *r.RecipientAccountNumberID
	}
	return
}

// The drawdown request's originator's account number.
func (r InboundWireDrawdownRequest) GetOriginatorAccountNumber() (OriginatorAccountNumber string) {
	if r.OriginatorAccountNumber != nil {
		OriginatorAccountNumber = *r.OriginatorAccountNumber
	}
	return
}

// The drawdown request's originator's routing number.
func (r InboundWireDrawdownRequest) GetOriginatorRoutingNumber() (OriginatorRoutingNumber string) {
	if r.OriginatorRoutingNumber != nil {
		OriginatorRoutingNumber = *r.OriginatorRoutingNumber
	}
	return
}

// The drawdown request's beneficiary's account number.
func (r InboundWireDrawdownRequest) GetBeneficiaryAccountNumber() (BeneficiaryAccountNumber string) {
	if r.BeneficiaryAccountNumber != nil {
		BeneficiaryAccountNumber = *r.BeneficiaryAccountNumber
	}
	return
}

// The drawdown request's beneficiary's routing number.
func (r InboundWireDrawdownRequest) GetBeneficiaryRoutingNumber() (BeneficiaryRoutingNumber string) {
	if r.BeneficiaryRoutingNumber != nil {
		BeneficiaryRoutingNumber = *r.BeneficiaryRoutingNumber
	}
	return
}

// The amount being requested in cents.
func (r InboundWireDrawdownRequest) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
// requested. Will always be "USD".
func (r InboundWireDrawdownRequest) GetCurrency() (Currency string) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// A message from the drawdown request's originator.
func (r InboundWireDrawdownRequest) GetMessageToRecipient() (MessageToRecipient string) {
	if r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// Line 1 of the information conveyed from the originator of the message to the
// beneficiary.
func (r InboundWireDrawdownRequest) GetOriginatorToBeneficiaryInformationLine1() (OriginatorToBeneficiaryInformationLine1 string) {
	if r.OriginatorToBeneficiaryInformationLine1 != nil {
		OriginatorToBeneficiaryInformationLine1 = *r.OriginatorToBeneficiaryInformationLine1
	}
	return
}

// Line 2 of the information conveyed from the originator of the message to the
// beneficiary.
func (r InboundWireDrawdownRequest) GetOriginatorToBeneficiaryInformationLine2() (OriginatorToBeneficiaryInformationLine2 string) {
	if r.OriginatorToBeneficiaryInformationLine2 != nil {
		OriginatorToBeneficiaryInformationLine2 = *r.OriginatorToBeneficiaryInformationLine2
	}
	return
}

// Line 3 of the information conveyed from the originator of the message to the
// beneficiary.
func (r InboundWireDrawdownRequest) GetOriginatorToBeneficiaryInformationLine3() (OriginatorToBeneficiaryInformationLine3 string) {
	if r.OriginatorToBeneficiaryInformationLine3 != nil {
		OriginatorToBeneficiaryInformationLine3 = *r.OriginatorToBeneficiaryInformationLine3
	}
	return
}

// Line 4 of the information conveyed from the originator of the message to the
// beneficiary.
func (r InboundWireDrawdownRequest) GetOriginatorToBeneficiaryInformationLine4() (OriginatorToBeneficiaryInformationLine4 string) {
	if r.OriginatorToBeneficiaryInformationLine4 != nil {
		OriginatorToBeneficiaryInformationLine4 = *r.OriginatorToBeneficiaryInformationLine4
	}
	return
}

// The drawdown request's originator's name.
func (r InboundWireDrawdownRequest) GetOriginatorName() (OriginatorName string) {
	if r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

// Line 1 of the drawdown request's originator's address.
func (r InboundWireDrawdownRequest) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

// Line 2 of the drawdown request's originator's address.
func (r InboundWireDrawdownRequest) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

// Line 3 of the drawdown request's originator's address.
func (r InboundWireDrawdownRequest) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

// The drawdown request's beneficiary's name.
func (r InboundWireDrawdownRequest) GetBeneficiaryName() (BeneficiaryName string) {
	if r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// Line 1 of the drawdown request's beneficiary's address.
func (r InboundWireDrawdownRequest) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// Line 2 of the drawdown request's beneficiary's address.
func (r InboundWireDrawdownRequest) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// Line 3 of the drawdown request's beneficiary's address.
func (r InboundWireDrawdownRequest) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r InboundWireDrawdownRequest) String() (result string) {
	return fmt.Sprintf("&InboundWireDrawdownRequest{Type:%s ID:%s RecipientAccountNumberID:%s OriginatorAccountNumber:%s OriginatorRoutingNumber:%s BeneficiaryAccountNumber:%s BeneficiaryRoutingNumber:%s Amount:%s Currency:%s MessageToRecipient:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorName:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s}", core.FmtP(r.Type), core.FmtP(r.ID), core.FmtP(r.RecipientAccountNumberID), core.FmtP(r.OriginatorAccountNumber), core.FmtP(r.OriginatorRoutingNumber), core.FmtP(r.BeneficiaryAccountNumber), core.FmtP(r.BeneficiaryRoutingNumber), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MessageToRecipient), core.FmtP(r.OriginatorToBeneficiaryInformationLine1), core.FmtP(r.OriginatorToBeneficiaryInformationLine2), core.FmtP(r.OriginatorToBeneficiaryInformationLine3), core.FmtP(r.OriginatorToBeneficiaryInformationLine4), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3))
}

type InboundWireDrawdownRequestType string

const (
	InboundWireDrawdownRequestTypeInboundWireDrawdownRequest InboundWireDrawdownRequestType = "inbound_wire_drawdown_request"
)

type InboundWireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                 `query:"limit"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// InboundWireDrawdownRequestListParams using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundWireDrawdownRequestListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundWireDrawdownRequestListParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *InboundWireDrawdownRequestListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes InboundWireDrawdownRequestListParams into a url.Values of
// the query parameters associated with this value
func (r *InboundWireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r InboundWireDrawdownRequestListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r InboundWireDrawdownRequestListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r InboundWireDrawdownRequestListParams) String() (result string) {
	return fmt.Sprintf("&InboundWireDrawdownRequestListParams{Cursor:%s Limit:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit))
}

type InboundWireDrawdownRequestList struct {
	// The contents of the list.
	Data *[]InboundWireDrawdownRequest `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// InboundWireDrawdownRequestList using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InboundWireDrawdownRequestList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundWireDrawdownRequestList into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *InboundWireDrawdownRequestList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes InboundWireDrawdownRequestList into a url.Values of the
// query parameters associated with this value
func (r *InboundWireDrawdownRequestList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r InboundWireDrawdownRequestList) GetData() (Data []InboundWireDrawdownRequest) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r InboundWireDrawdownRequestList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r InboundWireDrawdownRequestList) String() (result string) {
	return fmt.Sprintf("&InboundWireDrawdownRequestList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type InboundWireDrawdownRequestsPage struct {
	*pagination.Page[InboundWireDrawdownRequest]
}

func (r *InboundWireDrawdownRequestsPage) InboundWireDrawdownRequest() *InboundWireDrawdownRequest {
	return r.Current()
}

func (r *InboundWireDrawdownRequestsPage) NextPage() (*InboundWireDrawdownRequestsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &InboundWireDrawdownRequestsPage{page}, nil
	}
}
