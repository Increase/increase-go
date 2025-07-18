// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// WireTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWireTransferService] method instead.
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
func (r *WireTransferService) Get(ctx context.Context, wireTransferID string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if wireTransferID == "" {
		err = errors.New("missing required wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("wire_transfers/%s", wireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Wire Transfers
func (r *WireTransferService) List(ctx context.Context, query WireTransferListParams, opts ...option.RequestOption) (res *pagination.Page[WireTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *WireTransferService) ListAutoPaging(ctx context.Context, query WireTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[WireTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Wire Transfer
func (r *WireTransferService) Approve(ctx context.Context, wireTransferID string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if wireTransferID == "" {
		err = errors.New("missing required wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("wire_transfers/%s/approve", wireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Wire Transfer
func (r *WireTransferService) Cancel(ctx context.Context, wireTransferID string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if wireTransferID == "" {
		err = errors.New("missing required wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("wire_transfers/%s/cancel", wireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Wire transfers move funds between your Increase account and any other account
// accessible by Fedwire.
type WireTransfer struct {
	// The wire transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval WireTransferApproval `json:"approval,required,nullable"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// The beneficiary's name.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation WireTransferCancellation `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy WireTransferCreatedBy `json:"created_by,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency WireTransferCurrency `json:"currency,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The ID of an Inbound Wire Drawdown Request in response to which this transfer
	// was sent.
	InboundWireDrawdownRequestID string `json:"inbound_wire_drawdown_request_id,required,nullable"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required,nullable"`
	// The transfer's network.
	Network WireTransferNetwork `json:"network,required"`
	// The originator's address line 1.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The originator's address line 2.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The originator's address line 3.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator's name.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal WireTransferReversal `json:"reversal,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The Account Number that was passed to the wire's recipient.
	SourceAccountNumberID string `json:"source_account_number_id,required,nullable"`
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
	JSON wireTransferJSON `json:"-"`
}

// wireTransferJSON contains the JSON metadata for the struct [WireTransfer]
type wireTransferJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountNumber                apijson.Field
	Amount                       apijson.Field
	Approval                     apijson.Field
	BeneficiaryAddressLine1      apijson.Field
	BeneficiaryAddressLine2      apijson.Field
	BeneficiaryAddressLine3      apijson.Field
	BeneficiaryName              apijson.Field
	Cancellation                 apijson.Field
	CreatedAt                    apijson.Field
	CreatedBy                    apijson.Field
	Currency                     apijson.Field
	ExternalAccountID            apijson.Field
	IdempotencyKey               apijson.Field
	InboundWireDrawdownRequestID apijson.Field
	MessageToRecipient           apijson.Field
	Network                      apijson.Field
	OriginatorAddressLine1       apijson.Field
	OriginatorAddressLine2       apijson.Field
	OriginatorAddressLine3       apijson.Field
	OriginatorName               apijson.Field
	PendingTransactionID         apijson.Field
	Reversal                     apijson.Field
	RoutingNumber                apijson.Field
	SourceAccountNumberID        apijson.Field
	Status                       apijson.Field
	Submission                   apijson.Field
	TransactionID                apijson.Field
	Type                         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *WireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireTransferJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type WireTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                   `json:"approved_by,required,nullable"`
	JSON       wireTransferApprovalJSON `json:"-"`
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

func (r wireTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                       `json:"canceled_by,required,nullable"`
	JSON       wireTransferCancellationJSON `json:"-"`
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

func (r wireTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type WireTransferCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey WireTransferCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category WireTransferCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication WireTransferCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User WireTransferCreatedByUser `json:"user,required,nullable"`
	JSON wireTransferCreatedByJSON `json:"-"`
}

// wireTransferCreatedByJSON contains the JSON metadata for the struct
// [WireTransferCreatedBy]
type wireTransferCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WireTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type WireTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                          `json:"description,required,nullable"`
	JSON        wireTransferCreatedByAPIKeyJSON `json:"-"`
}

// wireTransferCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [WireTransferCreatedByAPIKey]
type wireTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type WireTransferCreatedByCategory string

const (
	WireTransferCreatedByCategoryAPIKey           WireTransferCreatedByCategory = "api_key"
	WireTransferCreatedByCategoryOAuthApplication WireTransferCreatedByCategory = "oauth_application"
	WireTransferCreatedByCategoryUser             WireTransferCreatedByCategory = "user"
)

func (r WireTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case WireTransferCreatedByCategoryAPIKey, WireTransferCreatedByCategoryOAuthApplication, WireTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type WireTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                    `json:"name,required"`
	JSON wireTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// wireTransferCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [WireTransferCreatedByOAuthApplication]
type wireTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type WireTransferCreatedByUser struct {
	// The email address of the User.
	Email string                        `json:"email,required"`
	JSON  wireTransferCreatedByUserJSON `json:"-"`
}

// wireTransferCreatedByUserJSON contains the JSON metadata for the struct
// [WireTransferCreatedByUser]
type wireTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For wire transfers this is always equal to `usd`.
type WireTransferCurrency string

const (
	WireTransferCurrencyCad WireTransferCurrency = "CAD"
	WireTransferCurrencyChf WireTransferCurrency = "CHF"
	WireTransferCurrencyEur WireTransferCurrency = "EUR"
	WireTransferCurrencyGbp WireTransferCurrency = "GBP"
	WireTransferCurrencyJpy WireTransferCurrency = "JPY"
	WireTransferCurrencyUsd WireTransferCurrency = "USD"
)

func (r WireTransferCurrency) IsKnown() bool {
	switch r {
	case WireTransferCurrencyCad, WireTransferCurrencyChf, WireTransferCurrencyEur, WireTransferCurrencyGbp, WireTransferCurrencyJpy, WireTransferCurrencyUsd:
		return true
	}
	return false
}

// The transfer's network.
type WireTransferNetwork string

const (
	WireTransferNetworkWire WireTransferNetwork = "wire"
)

func (r WireTransferNetwork) IsKnown() bool {
	switch r {
	case WireTransferNetworkWire:
		return true
	}
	return false
}

// If your transfer is reversed, this will contain details of the reversal.
type WireTransferReversal struct {
	// The amount that was reversed in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire, set by the reversing bank.
	Description string `json:"description,required"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The Fedwire cycle date for the wire reversal. The "Fedwire day" begins at 9:00
	// PM Eastern Time on the evening before the `cycle date`.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// The Fedwire cycle date for the wire transfer that is being reversed by this
	// message.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// The sending bank's reference number for the wire reversal.
	SenderReference string `json:"sender_reference,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string                   `json:"wire_transfer_id,required"`
	JSON           wireTransferReversalJSON `json:"-"`
}

// wireTransferReversalJSON contains the JSON metadata for the struct
// [WireTransferReversal]
type wireTransferReversalJSON struct {
	Amount                                                apijson.Field
	CreatedAt                                             apijson.Field
	Description                                           apijson.Field
	FinancialInstitutionToFinancialInstitutionInformation apijson.Field
	InputCycleDate                                        apijson.Field
	InputMessageAccountabilityData                        apijson.Field
	InputSequenceNumber                                   apijson.Field
	InputSource                                           apijson.Field
	OriginatorRoutingNumber                               apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	SenderReference                                       apijson.Field
	TransactionID                                         apijson.Field
	WireTransferID                                        apijson.Field
	raw                                                   string
	ExtraFields                                           map[string]apijson.Field
}

func (r *WireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wireTransferReversalJSON) RawJSON() string {
	return r.raw
}

// The lifecycle status of the transfer.
type WireTransferStatus string

const (
	WireTransferStatusPendingApproval   WireTransferStatus = "pending_approval"
	WireTransferStatusCanceled          WireTransferStatus = "canceled"
	WireTransferStatusPendingReviewing  WireTransferStatus = "pending_reviewing"
	WireTransferStatusRejected          WireTransferStatus = "rejected"
	WireTransferStatusRequiresAttention WireTransferStatus = "requires_attention"
	WireTransferStatusPendingCreating   WireTransferStatus = "pending_creating"
	WireTransferStatusReversed          WireTransferStatus = "reversed"
	WireTransferStatusSubmitted         WireTransferStatus = "submitted"
	WireTransferStatusComplete          WireTransferStatus = "complete"
)

func (r WireTransferStatus) IsKnown() bool {
	switch r {
	case WireTransferStatusPendingApproval, WireTransferStatusCanceled, WireTransferStatusPendingReviewing, WireTransferStatusRejected, WireTransferStatusRequiresAttention, WireTransferStatusPendingCreating, WireTransferStatusReversed, WireTransferStatusSubmitted, WireTransferStatusComplete:
		return true
	}
	return false
}

// After the transfer is submitted to Fedwire, this will contain supplemental
// details.
type WireTransferSubmission struct {
	// The accountability data for the submission.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// When this wire transfer was submitted to Fedwire.
	SubmittedAt time.Time                  `json:"submitted_at,required" format:"date-time"`
	JSON        wireTransferSubmissionJSON `json:"-"`
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

func (r wireTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `wire_transfer`.
type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

func (r WireTransferType) IsKnown() bool {
	switch r {
	case WireTransferTypeWireTransfer:
		return true
	}
	return false
}

type WireTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The transfer amount in USD cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The beneficiary's name.
	BeneficiaryName param.Field[string] `json:"beneficiary_name,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient param.Field[string] `json:"message_to_recipient,required"`
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 param.Field[string] `json:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 param.Field[string] `json:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 param.Field[string] `json:"beneficiary_address_line3"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// The ID of an Inbound Wire Drawdown Request in response to which this transfer is
	// being sent.
	InboundWireDrawdownRequestID param.Field[string] `json:"inbound_wire_drawdown_request_id"`
	// The originator's address line 1. This is only necessary if you're transferring
	// from a commingled account. Otherwise, we'll use the associated entity's details.
	OriginatorAddressLine1 param.Field[string] `json:"originator_address_line1"`
	// The originator's address line 2. This is only necessary if you're transferring
	// from a commingled account. Otherwise, we'll use the associated entity's details.
	OriginatorAddressLine2 param.Field[string] `json:"originator_address_line2"`
	// The originator's address line 3. This is only necessary if you're transferring
	// from a commingled account. Otherwise, we'll use the associated entity's details.
	OriginatorAddressLine3 param.Field[string] `json:"originator_address_line3"`
	// The originator's name. This is only necessary if you're transferring from a
	// commingled account. Otherwise, we'll use the associated entity's details.
	OriginatorName param.Field[string] `json:"originator_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number"`
	// The ID of an Account Number that will be passed to the wire's recipient
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id"`
}

func (r WireTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WireTransferListParams struct {
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID param.Field[string]                          `query:"account_id"`
	CreatedAt param.Field[WireTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WireTransferListParams]'s query parameters as `url.Values`.
func (r WireTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
