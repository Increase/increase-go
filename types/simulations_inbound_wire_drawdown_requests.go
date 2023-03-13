package types

import (
	"fmt"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
)

type SimulateAnInboundWireDrawdownRequestBeingCreatedParameters struct {
	// The Account Number to which the recipient of this request is being requested to
	// send funds from.
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

// UnmarshalJSON deserializes the provided bytes into
// SimulateAnInboundWireDrawdownRequestBeingCreatedParameters using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// SimulateAnInboundWireDrawdownRequestBeingCreatedParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account Number to which the recipient of this request is being requested to
// send funds from.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetRecipientAccountNumberID() (RecipientAccountNumberID string) {
	if r != nil && r.RecipientAccountNumberID != nil {
		RecipientAccountNumberID = *r.RecipientAccountNumberID
	}
	return
}

// The drawdown request's originator's account number.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorAccountNumber() (OriginatorAccountNumber string) {
	if r != nil && r.OriginatorAccountNumber != nil {
		OriginatorAccountNumber = *r.OriginatorAccountNumber
	}
	return
}

// The drawdown request's originator's routing number.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorRoutingNumber() (OriginatorRoutingNumber string) {
	if r != nil && r.OriginatorRoutingNumber != nil {
		OriginatorRoutingNumber = *r.OriginatorRoutingNumber
	}
	return
}

// The drawdown request's beneficiary's account number.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetBeneficiaryAccountNumber() (BeneficiaryAccountNumber string) {
	if r != nil && r.BeneficiaryAccountNumber != nil {
		BeneficiaryAccountNumber = *r.BeneficiaryAccountNumber
	}
	return
}

// The drawdown request's beneficiary's routing number.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetBeneficiaryRoutingNumber() (BeneficiaryRoutingNumber string) {
	if r != nil && r.BeneficiaryRoutingNumber != nil {
		BeneficiaryRoutingNumber = *r.BeneficiaryRoutingNumber
	}
	return
}

// The amount being requested in cents.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
// requested. Will always be "USD".
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// A message from the drawdown request's originator.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// Line 1 of the information conveyed from the originator of the message to the
// beneficiary.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorToBeneficiaryInformationLine1() (OriginatorToBeneficiaryInformationLine1 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine1 != nil {
		OriginatorToBeneficiaryInformationLine1 = *r.OriginatorToBeneficiaryInformationLine1
	}
	return
}

// Line 2 of the information conveyed from the originator of the message to the
// beneficiary.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorToBeneficiaryInformationLine2() (OriginatorToBeneficiaryInformationLine2 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine2 != nil {
		OriginatorToBeneficiaryInformationLine2 = *r.OriginatorToBeneficiaryInformationLine2
	}
	return
}

// Line 3 of the information conveyed from the originator of the message to the
// beneficiary.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorToBeneficiaryInformationLine3() (OriginatorToBeneficiaryInformationLine3 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine3 != nil {
		OriginatorToBeneficiaryInformationLine3 = *r.OriginatorToBeneficiaryInformationLine3
	}
	return
}

// Line 4 of the information conveyed from the originator of the message to the
// beneficiary.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorToBeneficiaryInformationLine4() (OriginatorToBeneficiaryInformationLine4 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine4 != nil {
		OriginatorToBeneficiaryInformationLine4 = *r.OriginatorToBeneficiaryInformationLine4
	}
	return
}

// The drawdown request's originator's name.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

// Line 1 of the drawdown request's originator's address.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

// Line 2 of the drawdown request's originator's address.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

// Line 3 of the drawdown request's originator's address.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

// The drawdown request's beneficiary's name.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// Line 1 of the drawdown request's beneficiary's address.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// Line 2 of the drawdown request's beneficiary's address.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// Line 3 of the drawdown request's beneficiary's address.
func (r *SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r SimulateAnInboundWireDrawdownRequestBeingCreatedParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnInboundWireDrawdownRequestBeingCreatedParameters{RecipientAccountNumberID:%s OriginatorAccountNumber:%s OriginatorRoutingNumber:%s BeneficiaryAccountNumber:%s BeneficiaryRoutingNumber:%s Amount:%s Currency:%s MessageToRecipient:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorName:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s}", core.FmtP(r.RecipientAccountNumberID), core.FmtP(r.OriginatorAccountNumber), core.FmtP(r.OriginatorRoutingNumber), core.FmtP(r.BeneficiaryAccountNumber), core.FmtP(r.BeneficiaryRoutingNumber), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MessageToRecipient), core.FmtP(r.OriginatorToBeneficiaryInformationLine1), core.FmtP(r.OriginatorToBeneficiaryInformationLine2), core.FmtP(r.OriginatorToBeneficiaryInformationLine3), core.FmtP(r.OriginatorToBeneficiaryInformationLine4), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3))
}
