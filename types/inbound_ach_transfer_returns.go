package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type InboundACHTransferReturn struct {
	// The ID of the Inbound ACH Transfer Return.
	ID *string `pjson:"id"`
	// The ID for the Transaction that is being returned.
	InboundACHTransferTransactionID *string `pjson:"inbound_ach_transfer_transaction_id"`
	// The ID for the transaction refunding the transfer.
	TransactionID *string `pjson:"transaction_id"`
	// The lifecycle status of the transfer.
	Status *InboundACHTransferReturnStatus `pjson:"status"`
	// The reason why this transfer will be returned. This is sent to the initiating
	// bank.
	Reason *InboundACHTransferReturnReason `pjson:"reason"`
	// After the return is submitted to FedACH, this will contain supplemental details.
	Submission *InboundACHTransferReturnSubmission `pjson:"submission"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_return`.
	Type       *InboundACHTransferReturnType `pjson:"type"`
	jsonFields map[string]interface{}        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into InboundACHTransferReturn
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundACHTransferReturn into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *InboundACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ID of the Inbound ACH Transfer Return.
func (r *InboundACHTransferReturn) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The ID for the Transaction that is being returned.
func (r *InboundACHTransferReturn) GetInboundACHTransferTransactionID() (InboundACHTransferTransactionID string) {
	if r != nil && r.InboundACHTransferTransactionID != nil {
		InboundACHTransferTransactionID = *r.InboundACHTransferTransactionID
	}
	return
}

// The ID for the transaction refunding the transfer.
func (r *InboundACHTransferReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The lifecycle status of the transfer.
func (r *InboundACHTransferReturn) GetStatus() (Status InboundACHTransferReturnStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The reason why this transfer will be returned. This is sent to the initiating
// bank.
func (r *InboundACHTransferReturn) GetReason() (Reason InboundACHTransferReturnReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// After the return is submitted to FedACH, this will contain supplemental details.
func (r *InboundACHTransferReturn) GetSubmission() (Submission InboundACHTransferReturnSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer_return`.
func (r *InboundACHTransferReturn) GetType() (Type InboundACHTransferReturnType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r InboundACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturn{ID:%s InboundACHTransferTransactionID:%s TransactionID:%s Status:%s Reason:%s Submission:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.InboundACHTransferTransactionID), core.FmtP(r.TransactionID), core.FmtP(r.Status), core.FmtP(r.Reason), r.Submission, core.FmtP(r.Type))
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
	TraceNumber *string `pjson:"trace_number"`
	// When the ACH transfer return was sent to FedACH.
	SubmittedAt *string                `pjson:"submitted_at"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// InboundACHTransferReturnSubmission using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundACHTransferReturnSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundACHTransferReturnSubmission into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *InboundACHTransferReturnSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The trace number for the submission.
func (r *InboundACHTransferReturnSubmission) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

// When the ACH transfer return was sent to FedACH.
func (r *InboundACHTransferReturnSubmission) GetSubmittedAt() (SubmittedAt string) {
	if r != nil && r.SubmittedAt != nil {
		SubmittedAt = *r.SubmittedAt
	}
	return
}

func (r InboundACHTransferReturnSubmission) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturnSubmission{TraceNumber:%s SubmittedAt:%s}", core.FmtP(r.TraceNumber), core.FmtP(r.SubmittedAt))
}

type InboundACHTransferReturnType string

const (
	InboundACHTransferReturnTypeInboundACHTransferReturn InboundACHTransferReturnType = "inbound_ach_transfer_return"
)

type CreateAnACHReturnParameters struct {
	// The transaction identifier of the Inbound ACH Transfer to return to the
	// originating financial institution.
	TransactionID *string `pjson:"transaction_id"`
	// The reason why this transfer will be returned. The most usual return codes are
	// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
	Reason     *CreateAnACHReturnParametersReason `pjson:"reason"`
	jsonFields map[string]interface{}             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAnACHReturnParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CreateAnACHReturnParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnACHReturnParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnACHReturnParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transaction identifier of the Inbound ACH Transfer to return to the
// originating financial institution.
func (r *CreateAnACHReturnParameters) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The reason why this transfer will be returned. The most usual return codes are
// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
func (r *CreateAnACHReturnParameters) GetReason() (Reason CreateAnACHReturnParametersReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r CreateAnACHReturnParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHReturnParameters{TransactionID:%s Reason:%s}", core.FmtP(r.TransactionID), core.FmtP(r.Reason))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                 `query:"limit"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// InboundACHTransferReturnListParams using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundACHTransferReturnListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundACHTransferReturnListParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *InboundACHTransferReturnListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes InboundACHTransferReturnListParams into a url.Values of the
// query parameters associated with this value
func (r *InboundACHTransferReturnListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *InboundACHTransferReturnListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *InboundACHTransferReturnListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r InboundACHTransferReturnListParams) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturnListParams{Cursor:%s Limit:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit))
}

type InboundACHTransferReturnList struct {
	// The contents of the list.
	Data *[]InboundACHTransferReturn `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into InboundACHTransferReturnList
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundACHTransferReturnList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes InboundACHTransferReturnList into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *InboundACHTransferReturnList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes InboundACHTransferReturnList into a url.Values of the query
// parameters associated with this value
func (r *InboundACHTransferReturnList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *InboundACHTransferReturnList) GetData() (Data []InboundACHTransferReturn) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *InboundACHTransferReturnList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r InboundACHTransferReturnList) String() (result string) {
	return fmt.Sprintf("&InboundACHTransferReturnList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
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
