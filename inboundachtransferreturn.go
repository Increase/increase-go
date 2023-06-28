// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// InboundACHTransferReturnService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewInboundACHTransferReturnService] method instead.
type InboundACHTransferReturnService struct {
	Options []option.RequestOption
}

// NewInboundACHTransferReturnService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
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
func (r *InboundACHTransferReturnService) Get(ctx context.Context, inboundACHTransferReturnID string, opts ...option.RequestOption) (res *InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfer_returns/%s", inboundACHTransferReturnID)
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

// If unauthorized activity occurs via ACH, you can create an Inbound ACH Transfer
// Return and we'll reverse the transaction. You can create an Inbound ACH Transfer
// return the first two days after receiving an Inbound ACH Transfer.
type InboundACHTransferReturn struct {
	// The ID of the Inbound ACH Transfer Return.
	ID string `json:"id,required"`
	// The ID for the Transaction that is being returned.
	InboundACHTransferTransactionID string `json:"inbound_ach_transfer_transaction_id,required"`
	// The reason why this transfer will be returned. This is sent to the initiating
	// bank.
	Reason InboundACHTransferReturnReason `json:"reason,required"`
	// The lifecycle status of the transfer.
	Status InboundACHTransferReturnStatus `json:"status,required"`
	// After the return is submitted to FedACH, this will contain supplemental details.
	Submission InboundACHTransferReturnSubmission `json:"submission,required,nullable"`
	// The ID for the transaction refunding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_return`.
	Type InboundACHTransferReturnType `json:"type,required"`
	JSON inboundACHTransferReturnJSON
}

// inboundACHTransferReturnJSON contains the JSON metadata for the struct
// [InboundACHTransferReturn]
type inboundACHTransferReturnJSON struct {
	ID                              apijson.Field
	InboundACHTransferTransactionID apijson.Field
	Reason                          apijson.Field
	Status                          apijson.Field
	Submission                      apijson.Field
	TransactionID                   apijson.Field
	Type                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *InboundACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why this transfer will be returned. This is sent to the initiating
// bank.
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

// The lifecycle status of the transfer.
type InboundACHTransferReturnStatus string

const (
	InboundACHTransferReturnStatusPendingSubmitting InboundACHTransferReturnStatus = "pending_submitting"
	InboundACHTransferReturnStatusSubmitted         InboundACHTransferReturnStatus = "submitted"
)

// After the return is submitted to FedACH, this will contain supplemental details.
type InboundACHTransferReturnSubmission struct {
	// When the ACH transfer return was sent to FedACH.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	// The trace number for the submission.
	TraceNumber string `json:"trace_number,required"`
	JSON        inboundACHTransferReturnSubmissionJSON
}

// inboundACHTransferReturnSubmissionJSON contains the JSON metadata for the struct
// [InboundACHTransferReturnSubmission]
type inboundACHTransferReturnSubmissionJSON struct {
	SubmittedAt apijson.Field
	TraceNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundACHTransferReturnSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer_return`.
type InboundACHTransferReturnType string

const (
	InboundACHTransferReturnTypeInboundACHTransferReturn InboundACHTransferReturnType = "inbound_ach_transfer_return"
)

type InboundACHTransferReturnNewParams struct {
	// The reason why this transfer will be returned. The most usual return codes are
	// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
	Reason param.Field[InboundACHTransferReturnNewParamsReason] `json:"reason,required"`
	// The transaction identifier of the Inbound ACH Transfer to return to the
	// originating financial institution.
	TransactionID param.Field[string] `json:"transaction_id,required"`
}

func (r InboundACHTransferReturnNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why this transfer will be returned. The most usual return codes are
// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
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
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundACHTransferReturnListParams]'s query parameters as
// `url.Values`.
func (r InboundACHTransferReturnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
