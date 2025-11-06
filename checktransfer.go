// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// CheckTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckTransferService] method instead.
type CheckTransferService struct {
	Options []option.RequestOption
}

// NewCheckTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckTransferService(opts ...option.RequestOption) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Options = opts
	return
}

// Create a Check Transfer
func (r *CheckTransferService) New(ctx context.Context, body CheckTransferNewParams, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "check_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Check Transfer
func (r *CheckTransferService) Get(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkTransferID == "" {
		err = errors.New("missing required check_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("check_transfers/%s", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Check Transfers
func (r *CheckTransferService) List(ctx context.Context, query CheckTransferListParams, opts ...option.RequestOption) (res *pagination.Page[CheckTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "check_transfers"
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

// List Check Transfers
func (r *CheckTransferService) ListAutoPaging(ctx context.Context, query CheckTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CheckTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Check Transfer
func (r *CheckTransferService) Approve(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkTransferID == "" {
		err = errors.New("missing required check_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("check_transfers/%s/approve", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkTransferID == "" {
		err = errors.New("missing required check_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("check_transfers/%s/cancel", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, checkTransferID string, body CheckTransferStopPaymentParams, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkTransferID == "" {
		err = errors.New("missing required check_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("check_transfers/%s/stop_payment", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check Transfers move funds from your Increase account by mailing a physical
// check.
type CheckTransfer struct {
	// The Check transfer's identifier.
	ID string `json:"id,required"`
	// The identifier of the Account from which funds will be transferred.
	AccountID string `json:"account_id,required"`
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval CheckTransferApproval `json:"approval,required,nullable"`
	// If the Check Transfer was successfully deposited, this will contain the
	// identifier of the Inbound Check Deposit object with details of the deposit.
	ApprovedInboundCheckDepositID string `json:"approved_inbound_check_deposit_id,required,nullable"`
	// How the account's available balance should be checked.
	BalanceCheck CheckTransferBalanceCheck `json:"balance_check,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation CheckTransferCancellation `json:"cancellation,required,nullable"`
	// The check number printed on the check.
	CheckNumber string `json:"check_number,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy CheckTransferCreatedBy `json:"created_by,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CheckTransferCurrency `json:"currency,required"`
	// Whether Increase will print and mail the check or if you will do it yourself.
	FulfillmentMethod CheckTransferFulfillmentMethod `json:"fulfillment_method,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// If the check has been mailed by Increase, this will contain details of the
	// shipment.
	Mailing CheckTransferMailing `json:"mailing,required,nullable"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// Details relating to the physical check that Increase will print and mail. Will
	// be present if and only if `fulfillment_method` is equal to `physical_check`.
	PhysicalCheck CheckTransferPhysicalCheck `json:"physical_check,required,nullable"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The identifier of the Account Number from which to send the transfer and print
	// on the check.
	SourceAccountNumberID string `json:"source_account_number_id,required,nullable"`
	// The lifecycle status of the transfer.
	Status CheckTransferStatus `json:"status,required"`
	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest CheckTransferStopPaymentRequest `json:"stop_payment_request,required,nullable"`
	// After the transfer is submitted, this will contain supplemental details.
	Submission CheckTransferSubmission `json:"submission,required,nullable"`
	// Details relating to the custom fulfillment you will perform. Will be present if
	// and only if `fulfillment_method` is equal to `third_party`.
	ThirdParty CheckTransferThirdParty `json:"third_party,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type CheckTransferType `json:"type,required"`
	// If set, the check will be valid on or before this date. After this date, the
	// check transfer will be stopped and deposits will not be accepted. For checks
	// printed by Increase, this date is included on the check as its expiry.
	ValidUntilDate time.Time              `json:"valid_until_date,required,nullable" format:"date"`
	ExtraFields    map[string]interface{} `json:"-,extras"`
	JSON           checkTransferJSON      `json:"-"`
}

// checkTransferJSON contains the JSON metadata for the struct [CheckTransfer]
type checkTransferJSON struct {
	ID                            apijson.Field
	AccountID                     apijson.Field
	AccountNumber                 apijson.Field
	Amount                        apijson.Field
	Approval                      apijson.Field
	ApprovedInboundCheckDepositID apijson.Field
	BalanceCheck                  apijson.Field
	Cancellation                  apijson.Field
	CheckNumber                   apijson.Field
	CreatedAt                     apijson.Field
	CreatedBy                     apijson.Field
	Currency                      apijson.Field
	FulfillmentMethod             apijson.Field
	IdempotencyKey                apijson.Field
	Mailing                       apijson.Field
	PendingTransactionID          apijson.Field
	PhysicalCheck                 apijson.Field
	RoutingNumber                 apijson.Field
	SourceAccountNumberID         apijson.Field
	Status                        apijson.Field
	StopPaymentRequest            apijson.Field
	Submission                    apijson.Field
	ThirdParty                    apijson.Field
	Type                          apijson.Field
	ValidUntilDate                apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CheckTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type CheckTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                    `json:"approved_by,required,nullable"`
	JSON       checkTransferApprovalJSON `json:"-"`
}

// checkTransferApprovalJSON contains the JSON metadata for the struct
// [CheckTransferApproval]
type checkTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// How the account's available balance should be checked.
type CheckTransferBalanceCheck string

const (
	CheckTransferBalanceCheckFull CheckTransferBalanceCheck = "full"
	CheckTransferBalanceCheckNone CheckTransferBalanceCheck = "none"
)

func (r CheckTransferBalanceCheck) IsKnown() bool {
	switch r {
	case CheckTransferBalanceCheckFull, CheckTransferBalanceCheckNone:
		return true
	}
	return false
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type CheckTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                        `json:"canceled_by,required,nullable"`
	JSON       checkTransferCancellationJSON `json:"-"`
}

// checkTransferCancellationJSON contains the JSON metadata for the struct
// [CheckTransferCancellation]
type checkTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type CheckTransferCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey CheckTransferCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category CheckTransferCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication CheckTransferCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User CheckTransferCreatedByUser `json:"user,required,nullable"`
	JSON checkTransferCreatedByJSON `json:"-"`
}

// checkTransferCreatedByJSON contains the JSON metadata for the struct
// [CheckTransferCreatedBy]
type checkTransferCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CheckTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type CheckTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                           `json:"description,required,nullable"`
	JSON        checkTransferCreatedByAPIKeyJSON `json:"-"`
}

// checkTransferCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [CheckTransferCreatedByAPIKey]
type checkTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type CheckTransferCreatedByCategory string

const (
	CheckTransferCreatedByCategoryAPIKey           CheckTransferCreatedByCategory = "api_key"
	CheckTransferCreatedByCategoryOAuthApplication CheckTransferCreatedByCategory = "oauth_application"
	CheckTransferCreatedByCategoryUser             CheckTransferCreatedByCategory = "user"
)

func (r CheckTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case CheckTransferCreatedByCategoryAPIKey, CheckTransferCreatedByCategoryOAuthApplication, CheckTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type CheckTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                     `json:"name,required"`
	JSON checkTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// checkTransferCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [CheckTransferCreatedByOAuthApplication]
type checkTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type CheckTransferCreatedByUser struct {
	// The email address of the User.
	Email string                         `json:"email,required"`
	JSON  checkTransferCreatedByUserJSON `json:"-"`
}

// checkTransferCreatedByUserJSON contains the JSON metadata for the struct
// [CheckTransferCreatedByUser]
type checkTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type CheckTransferCurrency string

const (
	CheckTransferCurrencyUsd CheckTransferCurrency = "USD"
)

func (r CheckTransferCurrency) IsKnown() bool {
	switch r {
	case CheckTransferCurrencyUsd:
		return true
	}
	return false
}

// Whether Increase will print and mail the check or if you will do it yourself.
type CheckTransferFulfillmentMethod string

const (
	CheckTransferFulfillmentMethodPhysicalCheck CheckTransferFulfillmentMethod = "physical_check"
	CheckTransferFulfillmentMethodThirdParty    CheckTransferFulfillmentMethod = "third_party"
)

func (r CheckTransferFulfillmentMethod) IsKnown() bool {
	switch r {
	case CheckTransferFulfillmentMethodPhysicalCheck, CheckTransferFulfillmentMethodThirdParty:
		return true
	}
	return false
}

// If the check has been mailed by Increase, this will contain details of the
// shipment.
type CheckTransferMailing struct {
	// The ID of the file corresponding to an image of the check that was mailed, if
	// available.
	ImageID string `json:"image_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt time.Time `json:"mailed_at,required" format:"date-time"`
	// The tracking number of the shipment, if available for the shipping method.
	TrackingNumber string                   `json:"tracking_number,required,nullable"`
	JSON           checkTransferMailingJSON `json:"-"`
}

// checkTransferMailingJSON contains the JSON metadata for the struct
// [CheckTransferMailing]
type checkTransferMailingJSON struct {
	ImageID        apijson.Field
	MailedAt       apijson.Field
	TrackingNumber apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CheckTransferMailing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferMailingJSON) RawJSON() string {
	return r.raw
}

// Details relating to the physical check that Increase will print and mail. Will
// be present if and only if `fulfillment_method` is equal to `physical_check`.
type CheckTransferPhysicalCheck struct {
	// The ID of the file for the check attachment.
	AttachmentFileID string `json:"attachment_file_id,required,nullable"`
	// Details for where Increase will mail the check.
	MailingAddress CheckTransferPhysicalCheckMailingAddress `json:"mailing_address,required"`
	// The descriptor that will be printed on the memo field on the check.
	Memo string `json:"memo,required,nullable"`
	// The descriptor that will be printed on the letter included with the check.
	Note string `json:"note,required,nullable"`
	// The payer of the check. This will be printed on the top-left portion of the
	// check and defaults to the return address if unspecified.
	Payer []CheckTransferPhysicalCheckPayer `json:"payer,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The return address to be printed on the check.
	ReturnAddress CheckTransferPhysicalCheckReturnAddress `json:"return_address,required,nullable"`
	// The shipping method for the check.
	ShippingMethod CheckTransferPhysicalCheckShippingMethod `json:"shipping_method,required,nullable"`
	// The text that will appear as the signature on the check in cursive font. If
	// blank, the check will be printed with 'No signature required'.
	SignatureText string `json:"signature_text,required,nullable"`
	// Tracking updates relating to the physical check's delivery.
	TrackingUpdates []CheckTransferPhysicalCheckTrackingUpdate `json:"tracking_updates,required"`
	ExtraFields     map[string]interface{}                     `json:"-,extras"`
	JSON            checkTransferPhysicalCheckJSON             `json:"-"`
}

// checkTransferPhysicalCheckJSON contains the JSON metadata for the struct
// [CheckTransferPhysicalCheck]
type checkTransferPhysicalCheckJSON struct {
	AttachmentFileID apijson.Field
	MailingAddress   apijson.Field
	Memo             apijson.Field
	Note             apijson.Field
	Payer            apijson.Field
	RecipientName    apijson.Field
	ReturnAddress    apijson.Field
	ShippingMethod   apijson.Field
	SignatureText    apijson.Field
	TrackingUpdates  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckJSON) RawJSON() string {
	return r.raw
}

// Details for where Increase will mail the check.
type CheckTransferPhysicalCheckMailingAddress struct {
	// The city of the check's destination.
	City string `json:"city,required,nullable"`
	// The street address of the check's destination.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the address of the check's destination.
	Line2 string `json:"line2,required,nullable"`
	// The name component of the check's mailing address.
	Name string `json:"name,required,nullable"`
	// The postal code of the check's destination.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state of the check's destination.
	State string                                       `json:"state,required,nullable"`
	JSON  checkTransferPhysicalCheckMailingAddressJSON `json:"-"`
}

// checkTransferPhysicalCheckMailingAddressJSON contains the JSON metadata for the
// struct [CheckTransferPhysicalCheckMailingAddress]
type checkTransferPhysicalCheckMailingAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheckMailingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckMailingAddressJSON) RawJSON() string {
	return r.raw
}

type CheckTransferPhysicalCheckPayer struct {
	// The contents of the line.
	Contents string                              `json:"contents,required"`
	JSON     checkTransferPhysicalCheckPayerJSON `json:"-"`
}

// checkTransferPhysicalCheckPayerJSON contains the JSON metadata for the struct
// [CheckTransferPhysicalCheckPayer]
type checkTransferPhysicalCheckPayerJSON struct {
	Contents    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheckPayer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckPayerJSON) RawJSON() string {
	return r.raw
}

// The return address to be printed on the check.
type CheckTransferPhysicalCheckReturnAddress struct {
	// The city of the check's destination.
	City string `json:"city,required,nullable"`
	// The street address of the check's destination.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the address of the check's destination.
	Line2 string `json:"line2,required,nullable"`
	// The name component of the check's return address.
	Name string `json:"name,required,nullable"`
	// The postal code of the check's destination.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state of the check's destination.
	State string                                      `json:"state,required,nullable"`
	JSON  checkTransferPhysicalCheckReturnAddressJSON `json:"-"`
}

// checkTransferPhysicalCheckReturnAddressJSON contains the JSON metadata for the
// struct [CheckTransferPhysicalCheckReturnAddress]
type checkTransferPhysicalCheckReturnAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheckReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckReturnAddressJSON) RawJSON() string {
	return r.raw
}

// The shipping method for the check.
type CheckTransferPhysicalCheckShippingMethod string

const (
	CheckTransferPhysicalCheckShippingMethodUspsFirstClass CheckTransferPhysicalCheckShippingMethod = "usps_first_class"
	CheckTransferPhysicalCheckShippingMethodFedexOvernight CheckTransferPhysicalCheckShippingMethod = "fedex_overnight"
)

func (r CheckTransferPhysicalCheckShippingMethod) IsKnown() bool {
	switch r {
	case CheckTransferPhysicalCheckShippingMethodUspsFirstClass, CheckTransferPhysicalCheckShippingMethodFedexOvernight:
		return true
	}
	return false
}

type CheckTransferPhysicalCheckTrackingUpdate struct {
	// The type of tracking event.
	Category CheckTransferPhysicalCheckTrackingUpdatesCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the tracking event took place.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The postal code where the event took place.
	PostalCode string                                       `json:"postal_code,required"`
	JSON       checkTransferPhysicalCheckTrackingUpdateJSON `json:"-"`
}

// checkTransferPhysicalCheckTrackingUpdateJSON contains the JSON metadata for the
// struct [CheckTransferPhysicalCheckTrackingUpdate]
type checkTransferPhysicalCheckTrackingUpdateJSON struct {
	Category    apijson.Field
	CreatedAt   apijson.Field
	PostalCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheckTrackingUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckTrackingUpdateJSON) RawJSON() string {
	return r.raw
}

// The type of tracking event.
type CheckTransferPhysicalCheckTrackingUpdatesCategory string

const (
	CheckTransferPhysicalCheckTrackingUpdatesCategoryInTransit            CheckTransferPhysicalCheckTrackingUpdatesCategory = "in_transit"
	CheckTransferPhysicalCheckTrackingUpdatesCategoryProcessedForDelivery CheckTransferPhysicalCheckTrackingUpdatesCategory = "processed_for_delivery"
	CheckTransferPhysicalCheckTrackingUpdatesCategoryDelivered            CheckTransferPhysicalCheckTrackingUpdatesCategory = "delivered"
	CheckTransferPhysicalCheckTrackingUpdatesCategoryReturnedToSender     CheckTransferPhysicalCheckTrackingUpdatesCategory = "returned_to_sender"
)

func (r CheckTransferPhysicalCheckTrackingUpdatesCategory) IsKnown() bool {
	switch r {
	case CheckTransferPhysicalCheckTrackingUpdatesCategoryInTransit, CheckTransferPhysicalCheckTrackingUpdatesCategoryProcessedForDelivery, CheckTransferPhysicalCheckTrackingUpdatesCategoryDelivered, CheckTransferPhysicalCheckTrackingUpdatesCategoryReturnedToSender:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type CheckTransferStatus string

const (
	CheckTransferStatusPendingApproval   CheckTransferStatus = "pending_approval"
	CheckTransferStatusCanceled          CheckTransferStatus = "canceled"
	CheckTransferStatusPendingSubmission CheckTransferStatus = "pending_submission"
	CheckTransferStatusRequiresAttention CheckTransferStatus = "requires_attention"
	CheckTransferStatusRejected          CheckTransferStatus = "rejected"
	CheckTransferStatusPendingMailing    CheckTransferStatus = "pending_mailing"
	CheckTransferStatusMailed            CheckTransferStatus = "mailed"
	CheckTransferStatusDeposited         CheckTransferStatus = "deposited"
	CheckTransferStatusStopped           CheckTransferStatus = "stopped"
	CheckTransferStatusReturned          CheckTransferStatus = "returned"
)

func (r CheckTransferStatus) IsKnown() bool {
	switch r {
	case CheckTransferStatusPendingApproval, CheckTransferStatusCanceled, CheckTransferStatusPendingSubmission, CheckTransferStatusRequiresAttention, CheckTransferStatusRejected, CheckTransferStatusPendingMailing, CheckTransferStatusMailed, CheckTransferStatusDeposited, CheckTransferStatusStopped, CheckTransferStatusReturned:
		return true
	}
	return false
}

// After a stop-payment is requested on the check, this will contain supplemental
// details.
type CheckTransferStopPaymentRequest struct {
	// The reason why this transfer was stopped.
	Reason CheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type        CheckTransferStopPaymentRequestType `json:"type,required"`
	ExtraFields map[string]interface{}              `json:"-,extras"`
	JSON        checkTransferStopPaymentRequestJSON `json:"-"`
}

// checkTransferStopPaymentRequestJSON contains the JSON metadata for the struct
// [CheckTransferStopPaymentRequest]
type checkTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferStopPaymentRequestJSON) RawJSON() string {
	return r.raw
}

// The reason why this transfer was stopped.
type CheckTransferStopPaymentRequestReason string

const (
	CheckTransferStopPaymentRequestReasonMailDeliveryFailed CheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	CheckTransferStopPaymentRequestReasonRejectedByIncrease CheckTransferStopPaymentRequestReason = "rejected_by_increase"
	CheckTransferStopPaymentRequestReasonNotAuthorized      CheckTransferStopPaymentRequestReason = "not_authorized"
	CheckTransferStopPaymentRequestReasonUnknown            CheckTransferStopPaymentRequestReason = "unknown"
)

func (r CheckTransferStopPaymentRequestReason) IsKnown() bool {
	switch r {
	case CheckTransferStopPaymentRequestReasonMailDeliveryFailed, CheckTransferStopPaymentRequestReasonRejectedByIncrease, CheckTransferStopPaymentRequestReasonNotAuthorized, CheckTransferStopPaymentRequestReasonUnknown:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

func (r CheckTransferStopPaymentRequestType) IsKnown() bool {
	switch r {
	case CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest:
		return true
	}
	return false
}

// After the transfer is submitted, this will contain supplemental details.
type CheckTransferSubmission struct {
	// Per USPS requirements, Increase will standardize the address to USPS standards
	// and check it against the USPS National Change of Address (NCOA) database before
	// mailing it. This indicates what modifications, if any, were made to the address
	// before printing and mailing the check.
	AddressCorrectionAction CheckTransferSubmissionAddressCorrectionAction `json:"address_correction_action,required"`
	// The address we submitted to the printer. This is what is physically printed on
	// the check.
	SubmittedAddress CheckTransferSubmissionSubmittedAddress `json:"submitted_address,required"`
	// When this check transfer was submitted to our check printer.
	SubmittedAt time.Time                   `json:"submitted_at,required" format:"date-time"`
	ExtraFields map[string]interface{}      `json:"-,extras"`
	JSON        checkTransferSubmissionJSON `json:"-"`
}

// checkTransferSubmissionJSON contains the JSON metadata for the struct
// [CheckTransferSubmission]
type checkTransferSubmissionJSON struct {
	AddressCorrectionAction apijson.Field
	SubmittedAddress        apijson.Field
	SubmittedAt             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CheckTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// Per USPS requirements, Increase will standardize the address to USPS standards
// and check it against the USPS National Change of Address (NCOA) database before
// mailing it. This indicates what modifications, if any, were made to the address
// before printing and mailing the check.
type CheckTransferSubmissionAddressCorrectionAction string

const (
	CheckTransferSubmissionAddressCorrectionActionNone                             CheckTransferSubmissionAddressCorrectionAction = "none"
	CheckTransferSubmissionAddressCorrectionActionStandardization                  CheckTransferSubmissionAddressCorrectionAction = "standardization"
	CheckTransferSubmissionAddressCorrectionActionStandardizationWithAddressChange CheckTransferSubmissionAddressCorrectionAction = "standardization_with_address_change"
	CheckTransferSubmissionAddressCorrectionActionError                            CheckTransferSubmissionAddressCorrectionAction = "error"
)

func (r CheckTransferSubmissionAddressCorrectionAction) IsKnown() bool {
	switch r {
	case CheckTransferSubmissionAddressCorrectionActionNone, CheckTransferSubmissionAddressCorrectionActionStandardization, CheckTransferSubmissionAddressCorrectionActionStandardizationWithAddressChange, CheckTransferSubmissionAddressCorrectionActionError:
		return true
	}
	return false
}

// The address we submitted to the printer. This is what is physically printed on
// the check.
type CheckTransferSubmissionSubmittedAddress struct {
	// The submitted address city.
	City string `json:"city,required"`
	// The submitted address line 1.
	Line1 string `json:"line1,required"`
	// The submitted address line 2.
	Line2 string `json:"line2,required,nullable"`
	// The submitted recipient name.
	RecipientName string `json:"recipient_name,required"`
	// The submitted address state.
	State string `json:"state,required"`
	// The submitted address zip.
	Zip  string                                      `json:"zip,required"`
	JSON checkTransferSubmissionSubmittedAddressJSON `json:"-"`
}

// checkTransferSubmissionSubmittedAddressJSON contains the JSON metadata for the
// struct [CheckTransferSubmissionSubmittedAddress]
type checkTransferSubmissionSubmittedAddressJSON struct {
	City          apijson.Field
	Line1         apijson.Field
	Line2         apijson.Field
	RecipientName apijson.Field
	State         apijson.Field
	Zip           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CheckTransferSubmissionSubmittedAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferSubmissionSubmittedAddressJSON) RawJSON() string {
	return r.raw
}

// Details relating to the custom fulfillment you will perform. Will be present if
// and only if `fulfillment_method` is equal to `third_party`.
type CheckTransferThirdParty struct {
	// The name that you will print on the check.
	RecipientName string                      `json:"recipient_name,required,nullable"`
	ExtraFields   map[string]interface{}      `json:"-,extras"`
	JSON          checkTransferThirdPartyJSON `json:"-"`
}

// checkTransferThirdPartyJSON contains the JSON metadata for the struct
// [CheckTransferThirdParty]
type checkTransferThirdPartyJSON struct {
	RecipientName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CheckTransferThirdParty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferThirdPartyJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer`.
type CheckTransferType string

const (
	CheckTransferTypeCheckTransfer CheckTransferType = "check_transfer"
)

func (r CheckTransferType) IsKnown() bool {
	switch r {
	case CheckTransferTypeCheckTransfer:
		return true
	}
	return false
}

type CheckTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The transfer amount in USD cents.
	Amount param.Field[int64] `json:"amount,required"`
	// Whether Increase will print and mail the check or if you will do it yourself.
	FulfillmentMethod param.Field[CheckTransferNewParamsFulfillmentMethod] `json:"fulfillment_method,required"`
	// The identifier of the Account Number from which to send the transfer and print
	// on the check.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id,required"`
	// How the account's available balance should be checked. If omitted, the default
	// behavior is `balance_check: full`.
	BalanceCheck param.Field[CheckTransferNewParamsBalanceCheck] `json:"balance_check"`
	// The check number Increase should use for the check. This should not contain
	// leading zeroes and must be unique across the `source_account_number`. If this is
	// omitted, Increase will generate a check number for you.
	CheckNumber param.Field[string] `json:"check_number"`
	// Details relating to the physical check that Increase will print and mail. This
	// is required if `fulfillment_method` is equal to `physical_check`. It must not be
	// included if any other `fulfillment_method` is provided.
	PhysicalCheck param.Field[CheckTransferNewParamsPhysicalCheck] `json:"physical_check"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// Details relating to the custom fulfillment you will perform. This is required if
	// `fulfillment_method` is equal to `third_party`. It must not be included if any
	// other `fulfillment_method` is provided.
	ThirdParty param.Field[CheckTransferNewParamsThirdParty] `json:"third_party"`
	// If provided, the check will be valid on or before this date. After this date,
	// the check transfer will be stopped and deposits will not be accepted. For checks
	// printed by Increase, this date is included on the check as its expiry.
	ValidUntilDate param.Field[time.Time] `json:"valid_until_date" format:"date"`
}

func (r CheckTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether Increase will print and mail the check or if you will do it yourself.
type CheckTransferNewParamsFulfillmentMethod string

const (
	CheckTransferNewParamsFulfillmentMethodPhysicalCheck CheckTransferNewParamsFulfillmentMethod = "physical_check"
	CheckTransferNewParamsFulfillmentMethodThirdParty    CheckTransferNewParamsFulfillmentMethod = "third_party"
)

func (r CheckTransferNewParamsFulfillmentMethod) IsKnown() bool {
	switch r {
	case CheckTransferNewParamsFulfillmentMethodPhysicalCheck, CheckTransferNewParamsFulfillmentMethodThirdParty:
		return true
	}
	return false
}

// How the account's available balance should be checked. If omitted, the default
// behavior is `balance_check: full`.
type CheckTransferNewParamsBalanceCheck string

const (
	CheckTransferNewParamsBalanceCheckFull CheckTransferNewParamsBalanceCheck = "full"
	CheckTransferNewParamsBalanceCheckNone CheckTransferNewParamsBalanceCheck = "none"
)

func (r CheckTransferNewParamsBalanceCheck) IsKnown() bool {
	switch r {
	case CheckTransferNewParamsBalanceCheckFull, CheckTransferNewParamsBalanceCheckNone:
		return true
	}
	return false
}

// Details relating to the physical check that Increase will print and mail. This
// is required if `fulfillment_method` is equal to `physical_check`. It must not be
// included if any other `fulfillment_method` is provided.
type CheckTransferNewParamsPhysicalCheck struct {
	// Details for where Increase will mail the check.
	MailingAddress param.Field[CheckTransferNewParamsPhysicalCheckMailingAddress] `json:"mailing_address,required"`
	// The descriptor that will be printed on the memo field on the check.
	Memo param.Field[string] `json:"memo,required"`
	// The name that will be printed on the check in the 'To:' field.
	RecipientName param.Field[string] `json:"recipient_name,required"`
	// The ID of a File to be attached to the check. This must have
	// `purpose: check_attachment`. For details on pricing and restrictions, see
	// https://increase.com/documentation/originating-checks#printing-checks .
	AttachmentFileID param.Field[string] `json:"attachment_file_id"`
	// The descriptor that will be printed on the letter included with the check.
	Note param.Field[string] `json:"note"`
	// The payer of the check. This will be printed on the top-left portion of the
	// check and defaults to the return address if unspecified. This should be an array
	// of up to 4 elements, each of which represents a line of the payer.
	Payer param.Field[[]CheckTransferNewParamsPhysicalCheckPayer] `json:"payer"`
	// The return address to be printed on the check. If omitted this will default to
	// an Increase-owned address that will mark checks as delivery failed and shred
	// them.
	ReturnAddress param.Field[CheckTransferNewParamsPhysicalCheckReturnAddress] `json:"return_address"`
	// How to ship the check. For details on pricing, timing, and restrictions, see
	// https://increase.com/documentation/originating-checks#printing-checks .
	ShippingMethod param.Field[CheckTransferNewParamsPhysicalCheckShippingMethod] `json:"shipping_method"`
	// The text that will appear as the signature on the check in cursive font. If not
	// provided, the check will be printed with 'No signature required'.
	SignatureText param.Field[string]    `json:"signature_text"`
	ExtraFields   map[string]interface{} `json:"-,extras"`
}

func (r CheckTransferNewParamsPhysicalCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details for where Increase will mail the check.
type CheckTransferNewParamsPhysicalCheckMailingAddress struct {
	// The city component of the check's destination address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address component of the check's destination address.
	Line1 param.Field[string] `json:"line1,required"`
	// The postal code component of the check's destination address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state component of the check's destination address.
	State param.Field[string] `json:"state,required"`
	// The second line of the address component of the check's destination address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CheckTransferNewParamsPhysicalCheckMailingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckTransferNewParamsPhysicalCheckPayer struct {
	// The contents of the line.
	Contents param.Field[string] `json:"contents,required"`
}

func (r CheckTransferNewParamsPhysicalCheckPayer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The return address to be printed on the check. If omitted this will default to
// an Increase-owned address that will mark checks as delivery failed and shred
// them.
type CheckTransferNewParamsPhysicalCheckReturnAddress struct {
	// The city of the return address.
	City param.Field[string] `json:"city,required"`
	// The first line of the return address.
	Line1 param.Field[string] `json:"line1,required"`
	// The name of the return address.
	Name param.Field[string] `json:"name,required"`
	// The postal code of the return address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state of the return address.
	State param.Field[string] `json:"state,required"`
	// The second line of the return address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CheckTransferNewParamsPhysicalCheckReturnAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How to ship the check. For details on pricing, timing, and restrictions, see
// https://increase.com/documentation/originating-checks#printing-checks .
type CheckTransferNewParamsPhysicalCheckShippingMethod string

const (
	CheckTransferNewParamsPhysicalCheckShippingMethodUspsFirstClass CheckTransferNewParamsPhysicalCheckShippingMethod = "usps_first_class"
	CheckTransferNewParamsPhysicalCheckShippingMethodFedexOvernight CheckTransferNewParamsPhysicalCheckShippingMethod = "fedex_overnight"
)

func (r CheckTransferNewParamsPhysicalCheckShippingMethod) IsKnown() bool {
	switch r {
	case CheckTransferNewParamsPhysicalCheckShippingMethodUspsFirstClass, CheckTransferNewParamsPhysicalCheckShippingMethodFedexOvernight:
		return true
	}
	return false
}

// Details relating to the custom fulfillment you will perform. This is required if
// `fulfillment_method` is equal to `third_party`. It must not be included if any
// other `fulfillment_method` is provided.
type CheckTransferNewParamsThirdParty struct {
	// The pay-to name you will print on the check. If provided, this is used for
	// [Positive Pay](/documentation/positive-pay). If this is omitted, Increase will
	// be unable to validate the payer name when the check is deposited.
	RecipientName param.Field[string]    `json:"recipient_name"`
	ExtraFields   map[string]interface{} `json:"-,extras"`
}

func (r CheckTransferNewParamsThirdParty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckTransferListParams struct {
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                         `query:"limit"`
	Status param.Field[CheckTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [CheckTransferListParams]'s query parameters as
// `url.Values`.
func (r CheckTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckTransferListParamsCreatedAt struct {
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

// URLQuery serializes [CheckTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CheckTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckTransferListParamsStatus struct {
	// Filter Check Transfers to those that have the specified status. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]CheckTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CheckTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r CheckTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckTransferListParamsStatusIn string

const (
	CheckTransferListParamsStatusInPendingApproval   CheckTransferListParamsStatusIn = "pending_approval"
	CheckTransferListParamsStatusInCanceled          CheckTransferListParamsStatusIn = "canceled"
	CheckTransferListParamsStatusInPendingSubmission CheckTransferListParamsStatusIn = "pending_submission"
	CheckTransferListParamsStatusInRequiresAttention CheckTransferListParamsStatusIn = "requires_attention"
	CheckTransferListParamsStatusInRejected          CheckTransferListParamsStatusIn = "rejected"
	CheckTransferListParamsStatusInPendingMailing    CheckTransferListParamsStatusIn = "pending_mailing"
	CheckTransferListParamsStatusInMailed            CheckTransferListParamsStatusIn = "mailed"
	CheckTransferListParamsStatusInDeposited         CheckTransferListParamsStatusIn = "deposited"
	CheckTransferListParamsStatusInStopped           CheckTransferListParamsStatusIn = "stopped"
	CheckTransferListParamsStatusInReturned          CheckTransferListParamsStatusIn = "returned"
)

func (r CheckTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case CheckTransferListParamsStatusInPendingApproval, CheckTransferListParamsStatusInCanceled, CheckTransferListParamsStatusInPendingSubmission, CheckTransferListParamsStatusInRequiresAttention, CheckTransferListParamsStatusInRejected, CheckTransferListParamsStatusInPendingMailing, CheckTransferListParamsStatusInMailed, CheckTransferListParamsStatusInDeposited, CheckTransferListParamsStatusInStopped, CheckTransferListParamsStatusInReturned:
		return true
	}
	return false
}

type CheckTransferStopPaymentParams struct {
	// The reason why this transfer should be stopped.
	Reason param.Field[CheckTransferStopPaymentParamsReason] `json:"reason"`
}

func (r CheckTransferStopPaymentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why this transfer should be stopped.
type CheckTransferStopPaymentParamsReason string

const (
	CheckTransferStopPaymentParamsReasonMailDeliveryFailed CheckTransferStopPaymentParamsReason = "mail_delivery_failed"
	CheckTransferStopPaymentParamsReasonNotAuthorized      CheckTransferStopPaymentParamsReason = "not_authorized"
	CheckTransferStopPaymentParamsReasonUnknown            CheckTransferStopPaymentParamsReason = "unknown"
)

func (r CheckTransferStopPaymentParamsReason) IsKnown() bool {
	switch r {
	case CheckTransferStopPaymentParamsReasonMailDeliveryFailed, CheckTransferStopPaymentParamsReasonNotAuthorized, CheckTransferStopPaymentParamsReasonUnknown:
		return true
	}
	return false
}
