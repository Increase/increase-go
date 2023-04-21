package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulationInboundWireDrawdownRequestNewParams struct {
	// The Account Number to which the recipient of this request is being requested to
	// send funds from.
	RecipientAccountNumberID field.Field[string] `json:"recipient_account_number_id,required"`
	// The drawdown request's originator's account number.
	OriginatorAccountNumber field.Field[string] `json:"originator_account_number,required"`
	// The drawdown request's originator's routing number.
	OriginatorRoutingNumber field.Field[string] `json:"originator_routing_number,required"`
	// The drawdown request's beneficiary's account number.
	BeneficiaryAccountNumber field.Field[string] `json:"beneficiary_account_number,required"`
	// The drawdown request's beneficiary's routing number.
	BeneficiaryRoutingNumber field.Field[string] `json:"beneficiary_routing_number,required"`
	// The amount being requested in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency field.Field[string] `json:"currency,required"`
	// A message from the drawdown request's originator.
	MessageToRecipient field.Field[string] `json:"message_to_recipient,required"`
	// Line 1 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine1 field.Field[string] `json:"originator_to_beneficiary_information_line1"`
	// Line 2 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine2 field.Field[string] `json:"originator_to_beneficiary_information_line2"`
	// Line 3 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine3 field.Field[string] `json:"originator_to_beneficiary_information_line3"`
	// Line 4 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine4 field.Field[string] `json:"originator_to_beneficiary_information_line4"`
	// The drawdown request's originator's name.
	OriginatorName field.Field[string] `json:"originator_name"`
	// Line 1 of the drawdown request's originator's address.
	OriginatorAddressLine1 field.Field[string] `json:"originator_address_line1"`
	// Line 2 of the drawdown request's originator's address.
	OriginatorAddressLine2 field.Field[string] `json:"originator_address_line2"`
	// Line 3 of the drawdown request's originator's address.
	OriginatorAddressLine3 field.Field[string] `json:"originator_address_line3"`
	// The drawdown request's beneficiary's name.
	BeneficiaryName field.Field[string] `json:"beneficiary_name"`
	// Line 1 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine1 field.Field[string] `json:"beneficiary_address_line1"`
	// Line 2 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine2 field.Field[string] `json:"beneficiary_address_line2"`
	// Line 3 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine3 field.Field[string] `json:"beneficiary_address_line3"`
}

// MarshalJSON serializes SimulationInboundWireDrawdownRequestNewParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r SimulationInboundWireDrawdownRequestNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}
