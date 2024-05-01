// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
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
func (r *InboundACHTransferService) List(ctx context.Context, query InboundACHTransferListParams, opts ...option.RequestOption) (res *pagination.Page[InboundACHTransfer], err error) {
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
func (r *InboundACHTransferService) ListAutoPaging(ctx context.Context, query InboundACHTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundACHTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
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
	// The inbound ACH transfer's identifier.
	ID string `json:"id,required"`
	// If your transfer is accepted, this will contain details of the acceptance.
	Acceptance InboundACHTransferAcceptance `json:"acceptance,required,nullable"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id,required"`
	// Additional information sent from the originator.
	Addenda InboundACHTransferAddenda `json:"addenda,required,nullable"`
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
	JSON inboundACHTransferJSON `json:"-"`
}

// inboundACHTransferJSON contains the JSON metadata for the struct
// [InboundACHTransfer]
type inboundACHTransferJSON struct {
	ID                                 apijson.Field
	Acceptance                         apijson.Field
	AccountID                          apijson.Field
	AccountNumberID                    apijson.Field
	Addenda                            apijson.Field
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

func (r inboundACHTransferJSON) RawJSON() string {
	return r.raw
}

// If your transfer is accepted, this will contain details of the acceptance.
type InboundACHTransferAcceptance struct {
	// The time at which the transfer was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The id of the transaction for the accepted transfer.
	TransactionID string                           `json:"transaction_id,required"`
	JSON          inboundACHTransferAcceptanceJSON `json:"-"`
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

func (r inboundACHTransferAcceptanceJSON) RawJSON() string {
	return r.raw
}

// Additional information sent from the originator.
type InboundACHTransferAddenda struct {
	// The type of addendum.
	Category InboundACHTransferAddendaCategory `json:"category,required"`
	// Unstructured `payment_related_information` passed through by the originator.
	Freeform InboundACHTransferAddendaFreeform `json:"freeform,required,nullable"`
	JSON     inboundACHTransferAddendaJSON     `json:"-"`
}

// inboundACHTransferAddendaJSON contains the JSON metadata for the struct
// [InboundACHTransferAddenda]
type inboundACHTransferAddendaJSON struct {
	Category    apijson.Field
	Freeform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundACHTransferAddenda) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAddendaJSON) RawJSON() string {
	return r.raw
}

// The type of addendum.
type InboundACHTransferAddendaCategory string

const (
	// Unstructured addendum.
	InboundACHTransferAddendaCategoryFreeform InboundACHTransferAddendaCategory = "freeform"
)

func (r InboundACHTransferAddendaCategory) IsKnown() bool {
	switch r {
	case InboundACHTransferAddendaCategoryFreeform:
		return true
	}
	return false
}

// Unstructured `payment_related_information` passed through by the originator.
type InboundACHTransferAddendaFreeform struct {
	// Each entry represents an addendum received from the originator.
	Entries []InboundACHTransferAddendaFreeformEntry `json:"entries,required"`
	JSON    inboundACHTransferAddendaFreeformJSON    `json:"-"`
}

// inboundACHTransferAddendaFreeformJSON contains the JSON metadata for the struct
// [InboundACHTransferAddendaFreeform]
type inboundACHTransferAddendaFreeformJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundACHTransferAddendaFreeform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAddendaFreeformJSON) RawJSON() string {
	return r.raw
}

type InboundACHTransferAddendaFreeformEntry struct {
	// The payment related information passed in the addendum.
	PaymentRelatedInformation string                                     `json:"payment_related_information,required"`
	JSON                      inboundACHTransferAddendaFreeformEntryJSON `json:"-"`
}

// inboundACHTransferAddendaFreeformEntryJSON contains the JSON metadata for the
// struct [InboundACHTransferAddendaFreeformEntry]
type inboundACHTransferAddendaFreeformEntryJSON struct {
	PaymentRelatedInformation apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InboundACHTransferAddendaFreeformEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAddendaFreeformEntryJSON) RawJSON() string {
	return r.raw
}

// If your transfer is declined, this will contain details of the decline.
type InboundACHTransferDecline struct {
	// The time at which the transfer was declined.
	DeclinedAt time.Time `json:"declined_at,required" format:"date-time"`
	// The id of the transaction for the declined transfer.
	DeclinedTransactionID string `json:"declined_transaction_id,required"`
	// The reason for the transfer decline.
	Reason InboundACHTransferDeclineReason `json:"reason,required"`
	JSON   inboundACHTransferDeclineJSON   `json:"-"`
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

func (r inboundACHTransferDeclineJSON) RawJSON() string {
	return r.raw
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
	// Your integration declined this transfer via the API.
	InboundACHTransferDeclineReasonUserInitiated InboundACHTransferDeclineReason = "user_initiated"
)

func (r InboundACHTransferDeclineReason) IsKnown() bool {
	switch r {
	case InboundACHTransferDeclineReasonACHRouteCanceled, InboundACHTransferDeclineReasonACHRouteDisabled, InboundACHTransferDeclineReasonBreachesLimit, InboundACHTransferDeclineReasonCreditEntryRefusedByReceiver, InboundACHTransferDeclineReasonDuplicateReturn, InboundACHTransferDeclineReasonEntityNotActive, InboundACHTransferDeclineReasonGroupLocked, InboundACHTransferDeclineReasonInsufficientFunds, InboundACHTransferDeclineReasonMisroutedReturn, InboundACHTransferDeclineReasonReturnOfErroneousOrReversingDebit, InboundACHTransferDeclineReasonNoACHRoute, InboundACHTransferDeclineReasonOriginatorRequest, InboundACHTransferDeclineReasonTransactionNotAllowed, InboundACHTransferDeclineReasonUserInitiated:
		return true
	}
	return false
}

// The direction of the transfer.
type InboundACHTransferDirection string

const (
	// Credit
	InboundACHTransferDirectionCredit InboundACHTransferDirection = "credit"
	// Debit
	InboundACHTransferDirectionDebit InboundACHTransferDirection = "debit"
)

func (r InboundACHTransferDirection) IsKnown() bool {
	switch r {
	case InboundACHTransferDirectionCredit, InboundACHTransferDirectionDebit:
		return true
	}
	return false
}

// If you initiate a notification of change in response to the transfer, this will
// contain its details.
type InboundACHTransferNotificationOfChange struct {
	// The new account number provided in the notification of change.
	UpdatedAccountNumber string `json:"updated_account_number,required,nullable"`
	// The new account number provided in the notification of change.
	UpdatedRoutingNumber string                                     `json:"updated_routing_number,required,nullable"`
	JSON                 inboundACHTransferNotificationOfChangeJSON `json:"-"`
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

func (r inboundACHTransferNotificationOfChangeJSON) RawJSON() string {
	return r.raw
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

func (r InboundACHTransferStatus) IsKnown() bool {
	switch r {
	case InboundACHTransferStatusPending, InboundACHTransferStatusDeclined, InboundACHTransferStatusAccepted, InboundACHTransferStatusReturned:
		return true
	}
	return false
}

// If your transfer is returned, this will contain details of the return.
type InboundACHTransferTransferReturn struct {
	// The reason for the transfer return.
	Reason InboundACHTransferTransferReturnReason `json:"reason,required"`
	// The time at which the transfer was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The id of the transaction for the returned transfer.
	TransactionID string                               `json:"transaction_id,required"`
	JSON          inboundACHTransferTransferReturnJSON `json:"-"`
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

func (r inboundACHTransferTransferReturnJSON) RawJSON() string {
	return r.raw
}

// The reason for the transfer return.
type InboundACHTransferTransferReturnReason string

const (
	// The customer's account has insufficient funds. This reason is only allowed for
	// debits. The Nacha return code is R01.
	InboundACHTransferTransferReturnReasonInsufficientFunds InboundACHTransferTransferReturnReason = "insufficient_funds"
	// The originating financial institution asked for this transfer to be returned.
	// The receiving bank is complying with the request. The Nacha return code is R06.
	InboundACHTransferTransferReturnReasonReturnedPerOdfiRequest InboundACHTransferTransferReturnReason = "returned_per_odfi_request"
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

func (r InboundACHTransferTransferReturnReason) IsKnown() bool {
	switch r {
	case InboundACHTransferTransferReturnReasonInsufficientFunds, InboundACHTransferTransferReturnReasonReturnedPerOdfiRequest, InboundACHTransferTransferReturnReasonAuthorizationRevokedByCustomer, InboundACHTransferTransferReturnReasonPaymentStopped, InboundACHTransferTransferReturnReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, InboundACHTransferTransferReturnReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, InboundACHTransferTransferReturnReasonBeneficiaryOrAccountHolderDeceased, InboundACHTransferTransferReturnReasonCreditEntryRefusedByReceiver, InboundACHTransferTransferReturnReasonDuplicateEntry, InboundACHTransferTransferReturnReasonCorporateCustomerAdvisedNotAuthorized:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer`.
type InboundACHTransferType string

const (
	InboundACHTransferTypeInboundACHTransfer InboundACHTransferType = "inbound_ach_transfer"
)

func (r InboundACHTransferType) IsKnown() bool {
	switch r {
	case InboundACHTransferTypeInboundACHTransfer:
		return true
	}
	return false
}

type InboundACHTransferListParams struct {
	// Filter Inbound ACH Tranfers to ones belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound ACH Tranfers to ones belonging to the specified Account Number.
	AccountNumberID param.Field[string]                                `query:"account_number_id"`
	CreatedAt       param.Field[InboundACHTransferListParamsCreatedAt] `query:"created_at"`
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

func (r InboundACHTransferListParamsStatus) IsKnown() bool {
	switch r {
	case InboundACHTransferListParamsStatusPending, InboundACHTransferListParamsStatusDeclined, InboundACHTransferListParamsStatusAccepted, InboundACHTransferListParamsStatusReturned:
		return true
	}
	return false
}

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
	// The customer's account has insufficient funds. This reason is only allowed for
	// debits. The Nacha return code is R01.
	InboundACHTransferTransferReturnParamsReasonInsufficientFunds InboundACHTransferTransferReturnParamsReason = "insufficient_funds"
	// The originating financial institution asked for this transfer to be returned.
	// The receiving bank is complying with the request. The Nacha return code is R06.
	InboundACHTransferTransferReturnParamsReasonReturnedPerOdfiRequest InboundACHTransferTransferReturnParamsReason = "returned_per_odfi_request"
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

func (r InboundACHTransferTransferReturnParamsReason) IsKnown() bool {
	switch r {
	case InboundACHTransferTransferReturnParamsReasonInsufficientFunds, InboundACHTransferTransferReturnParamsReasonReturnedPerOdfiRequest, InboundACHTransferTransferReturnParamsReasonAuthorizationRevokedByCustomer, InboundACHTransferTransferReturnParamsReasonPaymentStopped, InboundACHTransferTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, InboundACHTransferTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, InboundACHTransferTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased, InboundACHTransferTransferReturnParamsReasonCreditEntryRefusedByReceiver, InboundACHTransferTransferReturnParamsReasonDuplicateEntry, InboundACHTransferTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized:
		return true
	}
	return false
}
