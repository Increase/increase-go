package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type InboundACHTransferReturn struct {
	// The ID of the Inbound ACH Transfer Return.
	ID string `json:"id,required"`
	// The ID for the Transaction that is being returned.
	InboundACHTransferTransactionID string `json:"inbound_ach_transfer_transaction_id,required"`
	// The ID for the transaction refunding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The lifecycle status of the transfer.
	Status InboundACHTransferReturnStatus `json:"status,required"`
	// The reason why this transfer will be returned. This is sent to the initiating
	// bank.
	Reason InboundACHTransferReturnReason `json:"reason,required"`
	// After the return is submitted to FedACH, this will contain supplemental details.
	Submission InboundACHTransferReturnSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_return`.
	Type InboundACHTransferReturnType `json:"type,required"`
	JSON InboundACHTransferReturnJSON
}

type InboundACHTransferReturnJSON struct {
	ID                              apijson.Metadata
	InboundACHTransferTransactionID apijson.Metadata
	TransactionID                   apijson.Metadata
	Status                          apijson.Metadata
	Reason                          apijson.Metadata
	Submission                      apijson.Metadata
	Type                            apijson.Metadata
	raw                             string
	Extras                          map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InboundACHTransferReturn
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	TraceNumber string `json:"trace_number,required"`
	// When the ACH transfer return was sent to FedACH.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        InboundACHTransferReturnSubmissionJSON
}

type InboundACHTransferReturnSubmissionJSON struct {
	TraceNumber apijson.Metadata
	SubmittedAt apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundACHTransferReturnSubmission using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InboundACHTransferReturnSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundACHTransferReturnType string

const (
	InboundACHTransferReturnTypeInboundACHTransferReturn InboundACHTransferReturnType = "inbound_ach_transfer_return"
)

type InboundACHTransferReturnListResponse struct {
	// The contents of the list.
	Data []InboundACHTransferReturn `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       InboundACHTransferReturnListResponseJSON
}

type InboundACHTransferReturnListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundACHTransferReturnListResponse using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundACHTransferReturnListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
