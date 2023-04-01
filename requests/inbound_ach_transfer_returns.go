package requests

import (
	"fmt"
	"net/url"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateAnACHReturnParameters struct {
	// The transaction identifier of the Inbound ACH Transfer to return to the
	// originating financial institution.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
	// The reason why this transfer will be returned. The most usual return codes are
	// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
	Reason fields.Field[CreateAnACHReturnParametersReason] `json:"reason,required"`
}

// MarshalJSON serializes CreateAnACHReturnParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnACHReturnParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnACHReturnParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHReturnParameters{TransactionID:%s Reason:%s}", r.TransactionID, r.Reason)
}

type CreateAnACHReturnParametersReason string

const (
	CreateAnACHReturnParametersReasonAuthorizationRevokedByCustomer                              CreateAnACHReturnParametersReason = "authorization_revoked_by_customer"
	CreateAnACHReturnParametersReasonPaymentStopped                                              CreateAnACHReturnParametersReason = "payment_stopped"
	CreateAnACHReturnParametersReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   CreateAnACHReturnParametersReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	CreateAnACHReturnParametersReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity CreateAnACHReturnParametersReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	CreateAnACHReturnParametersReasonBeneficiaryOrAccountHolderDeceased                          CreateAnACHReturnParametersReason = "beneficiary_or_account_holder_deceased"
	CreateAnACHReturnParametersReasonCreditEntryRefusedByReceiver                                CreateAnACHReturnParametersReason = "credit_entry_refused_by_receiver"
	CreateAnACHReturnParametersReasonDuplicateEntry                                              CreateAnACHReturnParametersReason = "duplicate_entry"
	CreateAnACHReturnParametersReasonCorporateCustomerAdvisedNotAuthorized                       CreateAnACHReturnParametersReason = "corporate_customer_advised_not_authorized"
)

type InboundACHTransferReturnListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
}

// URLQuery serializes InboundACHTransferReturnListParams into a url.Values of the
// query parameters associated with this value
func (r *InboundACHTransferReturnListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r InboundACHTransferReturnListParams) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturnListParams{Cursor:%s Limit:%s}", r.Cursor, r.Limit)
}
