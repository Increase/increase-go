package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type InboundACHTransferReturn struct {
	// The ID of the Inbound ACH Transfer Return.
	ID fields.Field[string] `json:"id,required"`
	// The ID for the Transaction that is being returned.
	InboundACHTransferTransactionID fields.Field[string] `json:"inbound_ach_transfer_transaction_id,required"`
	// The ID for the transaction refunding the transfer.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
	// The lifecycle status of the transfer.
	Status fields.Field[InboundACHTransferReturnStatus] `json:"status,required"`
	// The reason why this transfer will be returned. This is sent to the initiating
	// bank.
	Reason fields.Field[InboundACHTransferReturnReason] `json:"reason,required"`
	// After the return is submitted to FedACH, this will contain supplemental details.
	Submission fields.Field[InboundACHTransferReturnSubmission] `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_return`.
	Type fields.Field[InboundACHTransferReturnType] `json:"type,required"`
}

// MarshalJSON serializes InboundACHTransferReturn into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *InboundACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturn{ID:%s InboundACHTransferTransactionID:%s TransactionID:%s Status:%s Reason:%s Submission:%s Type:%s}", r.ID, r.InboundACHTransferTransactionID, r.TransactionID, r.Status, r.Reason, r.Submission, r.Type)
}

type InboundACHTransferReturnStatus string

const (
	InboundACHTransferReturnStatusPendingSubmitting InboundACHTransferReturnStatus = "pending_submitting"
	InboundACHTransferReturnStatusSubmitted         InboundACHTransferReturnStatus = "submitted"
)

type InboundACHTransferReturnReason string

const (
	InboundACHTransferReturnReasonAuthorizationRevokedByCustomer                              InboundACHTransferReturnReason = "authorization_revoked_by_customer"
	InboundACHTransferReturnReasonPaymentStopped                                              InboundACHTransferReturnReason = "payment_stopped"
	InboundACHTransferReturnReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   InboundACHTransferReturnReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundACHTransferReturnReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferReturnReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	InboundACHTransferReturnReasonBeneficiaryOrAccountHolderDeceased                          InboundACHTransferReturnReason = "beneficiary_or_account_holder_deceased"
	InboundACHTransferReturnReasonCreditEntryRefusedByReceiver                                InboundACHTransferReturnReason = "credit_entry_refused_by_receiver"
	InboundACHTransferReturnReasonDuplicateEntry                                              InboundACHTransferReturnReason = "duplicate_entry"
	InboundACHTransferReturnReasonCorporateCustomerAdvisedNotAuthorized                       InboundACHTransferReturnReason = "corporate_customer_advised_not_authorized"
)

type InboundACHTransferReturnSubmission struct {
	// The trace number for the submission.
	TraceNumber fields.Field[string] `json:"trace_number,required"`
	// When the ACH transfer return was sent to FedACH.
	SubmittedAt fields.Field[time.Time] `json:"submitted_at,required" format:"date-time"`
}

// MarshalJSON serializes InboundACHTransferReturnSubmission into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *InboundACHTransferReturnSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundACHTransferReturnSubmission) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturnSubmission{TraceNumber:%s SubmittedAt:%s}", r.TraceNumber, r.SubmittedAt)
}

type InboundACHTransferReturnType string

const (
	InboundACHTransferReturnTypeInboundACHTransferReturn InboundACHTransferReturnType = "inbound_ach_transfer_return"
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
