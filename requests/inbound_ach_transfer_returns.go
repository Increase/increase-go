package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type InboundACHTransferReturnNewParams struct {
	// The transaction identifier of the Inbound ACH Transfer to return to the
	// originating financial institution.
	TransactionID field.Field[string] `json:"transaction_id,required"`
	// The reason why this transfer will be returned. The most usual return codes are
	// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
	Reason field.Field[InboundACHTransferReturnNewParamsReason] `json:"reason,required"`
}

// MarshalJSON serializes InboundACHTransferReturnNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r InboundACHTransferReturnNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboundACHTransferReturnNewParamsReason string

const (
	InboundACHTransferReturnNewParamsReasonAuthorizationRevokedByCustomer                              InboundACHTransferReturnNewParamsReason = "authorization_revoked_by_customer"
	InboundACHTransferReturnNewParamsReasonPaymentStopped                                              InboundACHTransferReturnNewParamsReason = "payment_stopped"
	InboundACHTransferReturnNewParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   InboundACHTransferReturnNewParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundACHTransferReturnNewParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferReturnNewParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	InboundACHTransferReturnNewParamsReasonBeneficiaryOrAccountHolderDeceased                          InboundACHTransferReturnNewParamsReason = "beneficiary_or_account_holder_deceased"
	InboundACHTransferReturnNewParamsReasonCreditEntryRefusedByReceiver                                InboundACHTransferReturnNewParamsReason = "credit_entry_refused_by_receiver"
	InboundACHTransferReturnNewParamsReasonDuplicateEntry                                              InboundACHTransferReturnNewParamsReason = "duplicate_entry"
	InboundACHTransferReturnNewParamsReasonCorporateCustomerAdvisedNotAuthorized                       InboundACHTransferReturnNewParamsReason = "corporate_customer_advised_not_authorized"
)

type InboundACHTransferReturnListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes InboundACHTransferReturnListParams into a url.Values of the
// query parameters associated with this value
func (r InboundACHTransferReturnListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
