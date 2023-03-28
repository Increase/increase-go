package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
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
	ID                              pjson.Metadata
	InboundACHTransferTransactionID pjson.Metadata
	TransactionID                   pjson.Metadata
	Status                          pjson.Metadata
	Reason                          pjson.Metadata
	Submission                      pjson.Metadata
	Type                            pjson.Metadata
	Raw                             []byte
	Extras                          map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InboundACHTransferReturn
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TraceNumber pjson.Metadata
	SubmittedAt pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundACHTransferReturnSubmission using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundACHTransferReturnSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InboundACHTransferReturnType string

const (
	InboundACHTransferReturnTypeInboundACHTransferReturn InboundACHTransferReturnType = "inbound_ach_transfer_return"
)

type InboundACHTransferReturnList struct {
	// The contents of the list.
	Data []InboundACHTransferReturn `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       InboundACHTransferReturnListJSON
}

type InboundACHTransferReturnListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InboundACHTransferReturnList
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundACHTransferReturnList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InboundACHTransferReturnsPage struct {
	*pagination.Page[InboundACHTransferReturn]
}

func (r *InboundACHTransferReturnsPage) InboundACHTransferReturn() *InboundACHTransferReturn {
	return r.Current()
}

func (r *InboundACHTransferReturnsPage) NextPage() (*InboundACHTransferReturnsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &InboundACHTransferReturnsPage{page}, nil
	}
}
