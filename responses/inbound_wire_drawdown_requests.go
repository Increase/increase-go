package responses

import (
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
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
	Type                                    pjson.Metadata
	ID                                      pjson.Metadata
	RecipientAccountNumberID                pjson.Metadata
	OriginatorAccountNumber                 pjson.Metadata
	OriginatorRoutingNumber                 pjson.Metadata
	BeneficiaryAccountNumber                pjson.Metadata
	BeneficiaryRoutingNumber                pjson.Metadata
	Amount                                  pjson.Metadata
	Currency                                pjson.Metadata
	MessageToRecipient                      pjson.Metadata
	OriginatorToBeneficiaryInformationLine1 pjson.Metadata
	OriginatorToBeneficiaryInformationLine2 pjson.Metadata
	OriginatorToBeneficiaryInformationLine3 pjson.Metadata
	OriginatorToBeneficiaryInformationLine4 pjson.Metadata
	OriginatorName                          pjson.Metadata
	OriginatorAddressLine1                  pjson.Metadata
	OriginatorAddressLine2                  pjson.Metadata
	OriginatorAddressLine3                  pjson.Metadata
	BeneficiaryName                         pjson.Metadata
	BeneficiaryAddressLine1                 pjson.Metadata
	BeneficiaryAddressLine2                 pjson.Metadata
	BeneficiaryAddressLine3                 pjson.Metadata
	Raw                                     []byte
	Extras                                  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InboundWireDrawdownRequest
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundWireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InboundWireDrawdownRequestType string

const (
	InboundWireDrawdownRequestTypeInboundWireDrawdownRequest InboundWireDrawdownRequestType = "inbound_wire_drawdown_request"
)

type InboundWireDrawdownRequestList struct {
	// The contents of the list.
	Data []InboundWireDrawdownRequest `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       InboundWireDrawdownRequestListJSON
}

type InboundWireDrawdownRequestListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundWireDrawdownRequestList using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InboundWireDrawdownRequestList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
