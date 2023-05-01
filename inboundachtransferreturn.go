package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type InboundACHTransferReturnService struct {
	Options []option.RequestOption
}

func NewInboundACHTransferReturnService(opts ...option.RequestOption) (r *InboundACHTransferReturnService) {
	r = &InboundACHTransferReturnService{}
	r.Options = opts
	return
}

// Create an ACH Return
func (r *InboundACHTransferReturnService) New(ctx context.Context, body InboundACHTransferReturnNewParams, opts ...option.RequestOption) (res *InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := "inbound_ach_transfer_returns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Inbound ACH Transfer Return
func (r *InboundACHTransferReturnService) Get(ctx context.Context, inbound_ach_transfer_return_id string, opts ...option.RequestOption) (res *InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfer_returns/%s", inbound_ach_transfer_return_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound ACH Transfer Returns
func (r *InboundACHTransferReturnService) List(ctx context.Context, query InboundACHTransferReturnListParams, opts ...option.RequestOption) (res *shared.Page[InboundACHTransferReturn], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_ach_transfer_returns"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Inbound ACH Transfer Returns
func (r *InboundACHTransferReturnService) ListAutoPaging(ctx context.Context, query InboundACHTransferReturnListParams, opts ...option.RequestOption) *shared.PageAutoPager[InboundACHTransferReturn] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

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
	return apiquery.Marshal(r)
}

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
