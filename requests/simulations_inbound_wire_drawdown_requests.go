package requests

import (
	"fmt"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type SimulateAnInboundWireDrawdownRequestBeingCreatedParameters struct {
	// The Account Number to which the recipient of this request is being requested to
	// send funds from.
	RecipientAccountNumberID fields.Field[string] `json:"recipient_account_number_id,required"`
	// The drawdown request's originator's account number.
	OriginatorAccountNumber fields.Field[string] `json:"originator_account_number,required"`
	// The drawdown request's originator's routing number.
	OriginatorRoutingNumber fields.Field[string] `json:"originator_routing_number,required"`
	// The drawdown request's beneficiary's account number.
	BeneficiaryAccountNumber fields.Field[string] `json:"beneficiary_account_number,required"`
	// The drawdown request's beneficiary's routing number.
	BeneficiaryRoutingNumber fields.Field[string] `json:"beneficiary_routing_number,required"`
	// The amount being requested in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency fields.Field[string] `json:"currency,required"`
	// A message from the drawdown request's originator.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	// Line 1 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine1 fields.Field[string] `json:"originator_to_beneficiary_information_line1"`
	// Line 2 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine2 fields.Field[string] `json:"originator_to_beneficiary_information_line2"`
	// Line 3 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine3 fields.Field[string] `json:"originator_to_beneficiary_information_line3"`
	// Line 4 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine4 fields.Field[string] `json:"originator_to_beneficiary_information_line4"`
	// The drawdown request's originator's name.
	OriginatorName fields.Field[string] `json:"originator_name"`
	// Line 1 of the drawdown request's originator's address.
	OriginatorAddressLine1 fields.Field[string] `json:"originator_address_line1"`
	// Line 2 of the drawdown request's originator's address.
	OriginatorAddressLine2 fields.Field[string] `json:"originator_address_line2"`
	// Line 3 of the drawdown request's originator's address.
	OriginatorAddressLine3 fields.Field[string] `json:"originator_address_line3"`
	// The drawdown request's beneficiary's name.
	BeneficiaryName fields.Field[string] `json:"beneficiary_name"`
	// Line 1 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine1 fields.Field[string] `json:"beneficiary_address_line1"`
	// Line 2 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine2 fields.Field[string] `json:"beneficiary_address_line2"`
	// Line 3 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine3 fields.Field[string] `json:"beneficiary_address_line3"`
}

// MarshalJSON serializes
// SimulateAnInboundWireDrawdownRequestBeingCreatedParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnInboundWireDrawdownRequestBeingCreatedParameters{RecipientAccountNumberID:%s OriginatorAccountNumber:%s OriginatorRoutingNumber:%s BeneficiaryAccountNumber:%s BeneficiaryRoutingNumber:%s Amount:%s Currency:%s MessageToRecipient:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorName:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s}", r.RecipientAccountNumberID, r.OriginatorAccountNumber, r.OriginatorRoutingNumber, r.BeneficiaryAccountNumber, r.BeneficiaryRoutingNumber, r.Amount, r.Currency, r.MessageToRecipient, r.OriginatorToBeneficiaryInformationLine1, r.OriginatorToBeneficiaryInformationLine2, r.OriginatorToBeneficiaryInformationLine3, r.OriginatorToBeneficiaryInformationLine4, r.OriginatorName, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.BeneficiaryName, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3)
}
