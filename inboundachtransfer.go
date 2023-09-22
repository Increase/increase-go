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

// InboundACHTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewInboundACHTransferService] method
// instead.
type InboundACHTransferService struct {
	Options []option.RequestOption
}

// NewInboundACHTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboundACHTransferService(opts ...option.RequestOption) (r *InboundACHTransferService) {
	r = &InboundACHTransferService{}
	r.Options = opts
	return
}

// Retrieve an Inbound ACH Transfer
func (r *InboundACHTransferService) Get(ctx context.Context, inboundACHTransferID string, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfers/%s", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound ACH Transfers
func (r *InboundACHTransferService) List(ctx context.Context, query InboundACHTransferListParams, opts ...option.RequestOption) (res *shared.Page[InboundACHTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_ach_transfers"
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

// List Inbound ACH Transfers
func (r *InboundACHTransferService) ListAutoPaging(ctx context.Context, query InboundACHTransferListParams, opts ...option.RequestOption) *shared.PageAutoPager[InboundACHTransfer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Decline an Inbound ACH Transfer
func (r *InboundACHTransferService) Decline(ctx context.Context, inboundACHTransferID string, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfers/%s/decline", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create a notification of change for an Inbound ACH Transfer
func (r *InboundACHTransferService) NotificationOfChange(ctx context.Context, inboundACHTransferID string, body InboundACHTransferNotificationOfChangeParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfers/%s/notification_of_change", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return an Inbound ACH Transfer
func (r *InboundACHTransferService) TransferReturn(ctx context.Context, inboundACHTransferID string, body InboundACHTransferTransferReturnParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfers/%s/transfer_return", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// An Inbound ACH Transfer is an ACH transfer initiated outside of Increase to your
// account.
type InboundACHTransfer struct {
	// The inbound ach transfer's identifier.
	ID string `json:"id,required"`
	// If your transfer is accepted, this will contain details of the acceptance.
	Acceptance InboundACHTransferAcceptance `json:"acceptance,required,nullable"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The time at which the transfer will be automatically resolved.
	AutomaticallyResolvesAt time.Time `json:"automatically_resolves_at,required" format:"date-time"`
	// If your transfer is declined, this will contain details of the decline.
	Decline InboundACHTransferDecline `json:"decline,required,nullable"`
	// The direction of the transfer.
	Direction InboundACHTransferDirection `json:"direction,required"`
	// If you initiate a notification of change in response to the transfer, this will
	// contain its details.
	NotificationOfChange InboundACHTransferNotificationOfChange `json:"notification_of_change,required,nullable"`
	// The descriptive date of the transfer.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date,required,nullable"`
	// The additional information included with the transfer.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	// The description of the transfer.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// The id of the company that initiated the transfer.
	OriginatorCompanyID string `json:"originator_company_id,required"`
	// The name of the company that initiated the transfer.
	OriginatorCompanyName string `json:"originator_company_name,required"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required"`
	// The id of the receiver of the transfer.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the receiver of the transfer.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// The status of the transfer.
	Status InboundACHTransferStatus `json:"status,required"`
	// The trace number of the transfer.
	TraceNumber string `json:"trace_number,required"`
	// If your transfer is returned, this will contain details of the return.
	TransferReturn InboundACHTransferTransferReturn `json:"transfer_return,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer`.
	Type InboundACHTransferType `json:"type,required"`
	JSON inboundACHTransferJSON
}

// inboundACHTransferJSON contains the JSON metadata for the struct
// [InboundACHTransfer]
type inboundACHTransferJSON struct {
	ID                                 apijson.Field
	Acceptance                         apijson.Field
	AccountNumberID                    apijson.Field
	Amount                             apijson.Field
	AutomaticallyResolvesAt            apijson.Field
	Decline                            apijson.Field
	Direction                          apijson.Field
	NotificationOfChange               apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	OriginatorRoutingNumber            apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	Status                             apijson.Field
	TraceNumber                        apijson.Field
	TransferReturn                     apijson.Field
	Type                               apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your transfer is accepted, this will contain details of the acceptance.
type InboundACHTransferAcceptance struct {
	// The time at which the transfer was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The id of the transaction for the accepted transfer.
	TransactionID string `json:"transaction_id,required"`
	JSON          inboundACHTransferAcceptanceJSON
}

// inboundACHTransferAcceptanceJSON contains the JSON metadata for the struct
// [InboundACHTransferAcceptance]
type inboundACHTransferAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundACHTransferAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your transfer is declined, this will contain details of the decline.
type InboundACHTransferDecline struct {
	// The time at which the transfer was declined.
	DeclinedAt time.Time `json:"declined_at,required" format:"date-time"`
	// The id of the transaction for the declined transfer.
	DeclinedTransactionID string `json:"declined_transaction_id,required"`
	// The reason for the transfer decline.
	Reason InboundACHTransferDeclineReason `json:"reason,required"`
	JSON   inboundACHTransferDeclineJSON
}

// inboundACHTransferDeclineJSON contains the JSON metadata for the struct
// [InboundACHTransferDecline]
type inboundACHTransferDeclineJSON struct {
	DeclinedAt            apijson.Field
	DeclinedTransactionID apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InboundACHTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason for the transfer decline.
type InboundACHTransferDeclineReason string

const (
	// The account number is canceled.
	InboundACHTransferDeclineReasonACHRouteCanceled InboundACHTransferDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	InboundACHTransferDeclineReasonACHRouteDisabled InboundACHTransferDeclineReason = "ach_route_disabled"
	// The transaction would cause an Increase limit to be exceeded.
	InboundACHTransferDeclineReasonBreachesLimit InboundACHTransferDeclineReason = "breaches_limit"
	// A credit was refused. This is a reasonable default reason for decline of
	// credits.
	InboundACHTransferDeclineReasonCreditEntryRefusedByReceiver InboundACHTransferDeclineReason = "credit_entry_refused_by_receiver"
	// A rare return reason. The return this message refers to was a duplicate.
	InboundACHTransferDeclineReasonDuplicateReturn InboundACHTransferDeclineReason = "duplicate_return"
	// The account's entity is not active.
	InboundACHTransferDeclineReasonEntityNotActive InboundACHTransferDeclineReason = "entity_not_active"
	// Your account is inactive.
	InboundACHTransferDeclineReasonGroupLocked InboundACHTransferDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	InboundACHTransferDeclineReasonInsufficientFunds InboundACHTransferDeclineReason = "insufficient_funds"
	// A rare return reason. The return this message refers to was misrouted.
	InboundACHTransferDeclineReasonMisroutedReturn InboundACHTransferDeclineReason = "misrouted_return"
	// The originating financial institution made a mistake and this return corrects
	// it.
	InboundACHTransferDeclineReasonReturnOfErroneousOrReversingDebit InboundACHTransferDeclineReason = "return_of_erroneous_or_reversing_debit"
	// The account number that was debited does not exist.
	InboundACHTransferDeclineReasonNoACHRoute InboundACHTransferDeclineReason = "no_ach_route"
	// The originating financial institution asked for this transfer to be returned.
	InboundACHTransferDeclineReasonOriginatorRequest InboundACHTransferDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	InboundACHTransferDeclineReasonTransactionNotAllowed InboundACHTransferDeclineReason = "transaction_not_allowed"
	// The user initiated the decline.
	InboundACHTransferDeclineReasonUserInitiated InboundACHTransferDeclineReason = "user_initiated"
)

// The direction of the transfer.
type InboundACHTransferDirection string

const (
	// Credit
	InboundACHTransferDirectionCredit InboundACHTransferDirection = "credit"
	// Debit
	InboundACHTransferDirectionDebit InboundACHTransferDirection = "debit"
)

// If you initiate a notification of change in response to the transfer, this will
// contain its details.
type InboundACHTransferNotificationOfChange struct {
	// The new account number provided in the notification of change.
	UpdatedAccountNumber string `json:"updated_account_number,required,nullable"`
	// The new account number provided in the notification of change.
	UpdatedRoutingNumber string `json:"updated_routing_number,required,nullable"`
	JSON                 inboundACHTransferNotificationOfChangeJSON
}

// inboundACHTransferNotificationOfChangeJSON contains the JSON metadata for the
// struct [InboundACHTransferNotificationOfChange]
type inboundACHTransferNotificationOfChangeJSON struct {
	UpdatedAccountNumber apijson.Field
	UpdatedRoutingNumber apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundACHTransferNotificationOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the transfer.
type InboundACHTransferStatus string

const (
	// The Inbound ACH Transfer is awaiting action, will transition automatically if no
	// action is taken.
	InboundACHTransferStatusPending InboundACHTransferStatus = "pending"
	// The Inbound ACH Transfer has been declined.
	InboundACHTransferStatusDeclined InboundACHTransferStatus = "declined"
	// The Inbound ACH Transfer is accepted.
	InboundACHTransferStatusAccepted InboundACHTransferStatus = "accepted"
	// The Inbound ACH Transfer has been returned.
	InboundACHTransferStatusReturned InboundACHTransferStatus = "returned"
)

// If your transfer is returned, this will contain details of the return.
type InboundACHTransferTransferReturn struct {
	// The reason for the transfer return.
	Reason InboundACHTransferTransferReturnReason `json:"reason,required"`
	// The time at which the transfer was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The id of the transaction for the returned transfer.
	TransactionID string `json:"transaction_id,required"`
	JSON          inboundACHTransferTransferReturnJSON
}

// inboundACHTransferTransferReturnJSON contains the JSON metadata for the struct
// [InboundACHTransferTransferReturn]
type inboundACHTransferTransferReturnJSON struct {
	Reason        apijson.Field
	ReturnedAt    apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundACHTransferTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason for the transfer return.
type InboundACHTransferTransferReturnReason string

const (
	// The customer no longer authorizes this transaction. The Nacha return code is
	// R07.
	InboundACHTransferTransferReturnReasonAuthorizationRevokedByCustomer InboundACHTransferTransferReturnReason = "authorization_revoked_by_customer"
	// The customer asked for the payment to be stopped. This reason is only allowed
	// for debits. The Nacha return code is R08.
	InboundACHTransferTransferReturnReasonPaymentStopped InboundACHTransferTransferReturnReason = "payment_stopped"
	// The customer advises that the debit was unauthorized. The Nacha return code is
	// R10.
	InboundACHTransferTransferReturnReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InboundACHTransferTransferReturnReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// The payee is deceased. The Nacha return code is R14.
	InboundACHTransferTransferReturnReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferTransferReturnReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// The account holder is deceased. The Nacha return code is R15.
	InboundACHTransferTransferReturnReasonBeneficiaryOrAccountHolderDeceased InboundACHTransferTransferReturnReason = "beneficiary_or_account_holder_deceased"
	// The customer refused a credit entry. This reason is only allowed for credits.
	// The Nacha return code is R23.
	InboundACHTransferTransferReturnReasonCreditEntryRefusedByReceiver InboundACHTransferTransferReturnReason = "credit_entry_refused_by_receiver"
	// The account holder identified this transaction as a duplicate. The Nacha return
	// code is R24.
	InboundACHTransferTransferReturnReasonDuplicateEntry InboundACHTransferTransferReturnReason = "duplicate_entry"
	// The corporate customer no longer authorizes this transaction. The Nacha return
	// code is R29.
	InboundACHTransferTransferReturnReasonCorporateCustomerAdvisedNotAuthorized InboundACHTransferTransferReturnReason = "corporate_customer_advised_not_authorized"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer`.
type InboundACHTransferType string

const (
	InboundACHTransferTypeInboundACHTransfer InboundACHTransferType = "inbound_ach_transfer"
)

type InboundACHTransferListParams struct {
	// Filter Inbound ACH Tranfers to ones belonging to the specified Account.
	AccountID param.Field[string]                                `query:"account_id"`
	CreatedAt param.Field[InboundACHTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Inbound ACH Transfers to those with the specified status.
	Status param.Field[InboundACHTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [InboundACHTransferListParams]'s query parameters as
// `url.Values`.
func (r InboundACHTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundACHTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [InboundACHTransferListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r InboundACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter Inbound ACH Transfers to those with the specified status.
type InboundACHTransferListParamsStatus string

const (
	// The Inbound ACH Transfer is awaiting action, will transition automatically if no
	// action is taken.
	InboundACHTransferListParamsStatusPending InboundACHTransferListParamsStatus = "pending"
	// The Inbound ACH Transfer has been declined.
	InboundACHTransferListParamsStatusDeclined InboundACHTransferListParamsStatus = "declined"
	// The Inbound ACH Transfer is accepted.
	InboundACHTransferListParamsStatusAccepted InboundACHTransferListParamsStatus = "accepted"
	// The Inbound ACH Transfer has been returned.
	InboundACHTransferListParamsStatusReturned InboundACHTransferListParamsStatus = "returned"
)

type InboundACHTransferNotificationOfChangeParams struct {
	// The updated account number to send in the notification of change.
	UpdatedAccountNumber param.Field[string] `json:"updated_account_number"`
	// The updated routing number to send in the notification of change.
	UpdatedRoutingNumber param.Field[string] `json:"updated_routing_number"`
}

func (r InboundACHTransferNotificationOfChangeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboundACHTransferTransferReturnParams struct {
	// The reason why this transfer will be returned. The most usual return codes are
	// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
	Reason param.Field[InboundACHTransferTransferReturnParamsReason] `json:"reason,required"`
}

func (r InboundACHTransferTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why this transfer will be returned. The most usual return codes are
// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
type InboundACHTransferTransferReturnParamsReason string

const (
	// The customer no longer authorizes this transaction. The Nacha return code is
	// R07.
	InboundACHTransferTransferReturnParamsReasonAuthorizationRevokedByCustomer InboundACHTransferTransferReturnParamsReason = "authorization_revoked_by_customer"
	// The customer asked for the payment to be stopped. This reason is only allowed
	// for debits. The Nacha return code is R08.
	InboundACHTransferTransferReturnParamsReasonPaymentStopped InboundACHTransferTransferReturnParamsReason = "payment_stopped"
	// The customer advises that the debit was unauthorized. The Nacha return code is
	// R10.
	InboundACHTransferTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InboundACHTransferTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// The payee is deceased. The Nacha return code is R14.
	InboundACHTransferTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferTransferReturnParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// The account holder is deceased. The Nacha return code is R15.
	InboundACHTransferTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased InboundACHTransferTransferReturnParamsReason = "beneficiary_or_account_holder_deceased"
	// The customer refused a credit entry. This reason is only allowed for credits.
	// The Nacha return code is R23.
	InboundACHTransferTransferReturnParamsReasonCreditEntryRefusedByReceiver InboundACHTransferTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	// The account holder identified this transaction as a duplicate. The Nacha return
	// code is R24.
	InboundACHTransferTransferReturnParamsReasonDuplicateEntry InboundACHTransferTransferReturnParamsReason = "duplicate_entry"
	// The corporate customer no longer authorizes this transaction. The Nacha return
	// code is R29.
	InboundACHTransferTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized InboundACHTransferTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
)
