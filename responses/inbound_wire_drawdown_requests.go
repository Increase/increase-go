package responses

import (
	apijson "github.com/increase/increase-go/internal/json"
)

type InboundWireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_drawdown_request`.
	Type InboundWireDrawdownRequestType `json:"type,required"`
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The Account Number from which the recipient of this request is being requested
	// to send funds.
	RecipientAccountNumberID string `json:"recipient_account_number_id,required"`
	// The drawdown request's originator's account number.
	OriginatorAccountNumber string `json:"originator_account_number,required"`
	// The drawdown request's originator's routing number.
	OriginatorRoutingNumber string `json:"originator_routing_number,required"`
	// The drawdown request's beneficiary's account number.
	BeneficiaryAccountNumber string `json:"beneficiary_account_number,required"`
	// The drawdown request's beneficiary's routing number.
	BeneficiaryRoutingNumber string `json:"beneficiary_routing_number,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// A message from the drawdown request's originator.
	MessageToRecipient string `json:"message_to_recipient,required,nullable"`
	// Line 1 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// Line 2 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// Line 3 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// Line 4 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	// The drawdown request's originator's name.
	OriginatorName string `json:"originator_name,required,nullable"`
	// Line 1 of the drawdown request's originator's address.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// Line 2 of the drawdown request's originator's address.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// Line 3 of the drawdown request's originator's address.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The drawdown request's beneficiary's name.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// Line 1 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// Line 2 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// Line 3 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	JSON                    InboundWireDrawdownRequestJSON
}

type InboundWireDrawdownRequestJSON struct {
	Type                                    apijson.Metadata
	ID                                      apijson.Metadata
	RecipientAccountNumberID                apijson.Metadata
	OriginatorAccountNumber                 apijson.Metadata
	OriginatorRoutingNumber                 apijson.Metadata
	BeneficiaryAccountNumber                apijson.Metadata
	BeneficiaryRoutingNumber                apijson.Metadata
	Amount                                  apijson.Metadata
	Currency                                apijson.Metadata
	MessageToRecipient                      apijson.Metadata
	OriginatorToBeneficiaryInformationLine1 apijson.Metadata
	OriginatorToBeneficiaryInformationLine2 apijson.Metadata
	OriginatorToBeneficiaryInformationLine3 apijson.Metadata
	OriginatorToBeneficiaryInformationLine4 apijson.Metadata
	OriginatorName                          apijson.Metadata
	OriginatorAddressLine1                  apijson.Metadata
	OriginatorAddressLine2                  apijson.Metadata
	OriginatorAddressLine3                  apijson.Metadata
	BeneficiaryName                         apijson.Metadata
	BeneficiaryAddressLine1                 apijson.Metadata
	BeneficiaryAddressLine2                 apijson.Metadata
	BeneficiaryAddressLine3                 apijson.Metadata
	raw                                     string
	Extras                                  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InboundWireDrawdownRequest
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundWireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundWireDrawdownRequestType string

const (
	InboundWireDrawdownRequestTypeInboundWireDrawdownRequest InboundWireDrawdownRequestType = "inbound_wire_drawdown_request"
)

type InboundWireDrawdownRequestListResponse struct {
	// The contents of the list.
	Data []InboundWireDrawdownRequest `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       InboundWireDrawdownRequestListResponseJSON
}

type InboundWireDrawdownRequestListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundWireDrawdownRequestListResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundWireDrawdownRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
