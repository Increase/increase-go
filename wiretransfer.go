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

// WireTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWireTransferService] method
// instead.
type WireTransferService struct {
	Options []option.RequestOption
}

// NewWireTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWireTransferService(opts ...option.RequestOption) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Options = opts
	return
}

// Create a Wire Transfer
func (r *WireTransferService) New(ctx context.Context, body WireTransferNewParams, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "wire_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Wire Transfer
func (r *WireTransferService) Get(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s", wire_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Transfers
func (r *WireTransferService) List(ctx context.Context, query WireTransferListParams, opts ...option.RequestOption) (res *shared.Page[WireTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "wire_transfers"
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

// List Wire Transfers
func (r *WireTransferService) ListAutoPaging(ctx context.Context, query WireTransferListParams, opts ...option.RequestOption) *shared.PageAutoPager[WireTransfer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Wire Transfer
func (r *WireTransferService) Approve(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s/approve", wire_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Wire Transfer
func (r *WireTransferService) Cancel(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s/cancel", wire_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the reversal of a [Wire Transfer](#wire-transfers) by the Federal
// Reserve due to error conditions. This will also create a
// [Transaction](#transaction) to account for the returned funds. This Wire
// Transfer must first have a `status` of `complete`.'
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/wire_transfers/%s/reverse", wire_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Wire Transfer](#wire-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/wire_transfers/%s/submit", wire_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Wire transfers move funds between your Increase account and any other account
// accessible by Fedwire.
type WireTransfer struct {
	// The wire transfer's identifier.
	ID string `json:"id,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency WireTransferCurrency `json:"currency,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The beneficiary's name.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval WireTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation WireTransferCancellation `json:"cancellation,required,nullable"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal WireTransferReversal `json:"reversal,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The transfer's network.
	Network WireTransferNetwork `json:"network,required"`
	// The lifecycle status of the transfer.
	Status WireTransferStatus `json:"status,required"`
	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission WireTransferSubmission `json:"submission,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type WireTransferType `json:"type,required"`
	JSON wireTransferJSON
}

// wireTransferJSON contains the JSON metadata for the struct [WireTransfer]
type wireTransferJSON struct {
	ID                      apijson.Field
	MessageToRecipient      apijson.Field
	Amount                  apijson.Field
	Currency                apijson.Field
	AccountNumber           apijson.Field
	BeneficiaryName         apijson.Field
	BeneficiaryAddressLine1 apijson.Field
	BeneficiaryAddressLine2 apijson.Field
	BeneficiaryAddressLine3 apijson.Field
	AccountID               apijson.Field
	ExternalAccountID       apijson.Field
	RoutingNumber           apijson.Field
	Approval                apijson.Field
	Cancellation            apijson.Field
	Reversal                apijson.Field
	CreatedAt               apijson.Field
	Network                 apijson.Field
	Status                  apijson.Field
	Submission              apijson.Field
	TransactionID           apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *WireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferCurrency string

const (
	WireTransferCurrencyCad WireTransferCurrency = "CAD"
	WireTransferCurrencyChf WireTransferCurrency = "CHF"
	WireTransferCurrencyEur WireTransferCurrency = "EUR"
	WireTransferCurrencyGbp WireTransferCurrency = "GBP"
	WireTransferCurrencyJpy WireTransferCurrency = "JPY"
	WireTransferCurrencyUsd WireTransferCurrency = "USD"
)

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type WireTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       wireTransferApprovalJSON
}

// wireTransferApprovalJSON contains the JSON metadata for the struct
// [WireTransferApproval]
type wireTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       wireTransferCancellationJSON
}

// wireTransferCancellationJSON contains the JSON metadata for the struct
// [WireTransferCancellation]
type wireTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your transfer is reversed, this will contain details of the reversal.
type WireTransferReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string `json:"wire_transfer_id,required"`
	JSON           wireTransferReversalJSON
}

// wireTransferReversalJSON contains the JSON metadata for the struct
// [WireTransferReversal]
type wireTransferReversalJSON struct {
	Amount                                                apijson.Field
	CreatedAt                                             apijson.Field
	Description                                           apijson.Field
	InputCycleDate                                        apijson.Field
	InputSequenceNumber                                   apijson.Field
	InputSource                                           apijson.Field
	InputMessageAccountabilityData                        apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	FinancialInstitutionToFinancialInstitutionInformation apijson.Field
	TransactionID                                         apijson.Field
	WireTransferID                                        apijson.Field
	raw                                                   string
	ExtraFields                                           map[string]apijson.Field
}

func (r *WireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferNetwork string

const (
	WireTransferNetworkWire WireTransferNetwork = "wire"
)

type WireTransferStatus string

const (
	WireTransferStatusCanceled          WireTransferStatus = "canceled"
	WireTransferStatusRequiresAttention WireTransferStatus = "requires_attention"
	WireTransferStatusPendingApproval   WireTransferStatus = "pending_approval"
	WireTransferStatusRejected          WireTransferStatus = "rejected"
	WireTransferStatusReversed          WireTransferStatus = "reversed"
	WireTransferStatusComplete          WireTransferStatus = "complete"
	WireTransferStatusPendingCreating   WireTransferStatus = "pending_creating"
)

// After the transfer is submitted to Fedwire, this will contain supplemental
// details.
type WireTransferSubmission struct {
	// The accountability data for the submission.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// When this wire transfer was submitted to Fedwire.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        wireTransferSubmissionJSON
}

// wireTransferSubmissionJSON contains the JSON metadata for the struct
// [WireTransferSubmission]
type wireTransferSubmissionJSON struct {
	InputMessageAccountabilityData apijson.Field
	SubmittedAt                    apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *WireTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

type WireTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// The transfer amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient param.Field[string] `json:"message_to_recipient,required"`
	// The beneficiary's name.
	BeneficiaryName param.Field[string] `json:"beneficiary_name,required"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 param.Field[string] `json:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 param.Field[string] `json:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 param.Field[string] `json:"beneficiary_address_line3"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
}

func (r WireTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WireTransferListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID param.Field[string]                          `query:"external_account_id"`
	CreatedAt         param.Field[WireTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes [WireTransferListParams]'s query parameters as `url.Values`.
func (r WireTransferListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type WireTransferListParamsCreatedAt struct {
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

// URLQuery serializes [WireTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r WireTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of Wire Transfer objects
type WireTransferListResponse struct {
	// The contents of the list.
	Data []WireTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       wireTransferListResponseJSON
}

// wireTransferListResponseJSON contains the JSON metadata for the struct
// [WireTransferListResponse]
type wireTransferListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
