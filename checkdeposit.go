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

// CheckDepositService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckDepositService] method instead.
type CheckDepositService struct {
	Options []option.RequestOption
}

// NewCheckDepositService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckDepositService(opts ...option.RequestOption) (r *CheckDepositService) {
	r = &CheckDepositService{}
	r.Options = opts
	return
}

// Create a Check Deposit
func (r *CheckDepositService) New(ctx context.Context, body CheckDepositNewParams, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "check_deposits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Check Deposit
func (r *CheckDepositService) Get(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("check_deposits/%s", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Check Deposits
func (r *CheckDepositService) List(ctx context.Context, query CheckDepositListParams, opts ...option.RequestOption) (res *pagination.Page[CheckDeposit], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "check_deposits"
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

// List Check Deposits
func (r *CheckDepositService) ListAutoPaging(ctx context.Context, query CheckDepositListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CheckDeposit] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Check Deposits allow you to deposit images of paper checks into your account.
type CheckDeposit struct {
	// The deposit's identifier.
	ID string `json:"id,required"`
	// The Account the check was deposited into.
	AccountID string `json:"account_id,required"`
	// The deposited amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Once your deposit is successfully parsed and accepted by Increase, this will
	// contain details of the parsed check.
	DepositAcceptance CheckDepositDepositAcceptance `json:"deposit_acceptance,required,nullable"`
	// If your deposit is rejected by Increase, this will contain details as to why it
	// was rejected.
	DepositRejection CheckDepositDepositRejection `json:"deposit_rejection,required,nullable"`
	// If your deposit is returned, this will contain details as to why it was
	// returned.
	DepositReturn CheckDepositDepositReturn `json:"deposit_return,required,nullable"`
	// After the check is parsed, it is submitted to the Check21 network for
	// processing. This will contain details of the submission.
	DepositSubmission CheckDepositDepositSubmission `json:"deposit_submission,required,nullable"`
	// The description of the Check Deposit, for display purposes only.
	Description string `json:"description,required,nullable"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// Increase will sometimes hold the funds for Check Deposits. If funds are held,
	// this sub-object will contain details of the hold.
	InboundFundsHold CheckDepositInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// If the Check Deposit was the result of an Inbound Mail Item, this will contain
	// the identifier of the Inbound Mail Item.
	InboundMailItemID string `json:"inbound_mail_item_id,required,nullable"`
	// If the Check Deposit was the result of an Inbound Mail Item, this will contain
	// the identifier of the Lockbox that received it.
	LockboxID string `json:"lockbox_id,required,nullable"`
	// The status of the Check Deposit.
	Status CheckDepositStatus `json:"status,required"`
	// The ID for the Transaction created by the deposit.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_deposit`.
	Type CheckDepositType `json:"type,required"`
	JSON checkDepositJSON `json:"-"`
}

// checkDepositJSON contains the JSON metadata for the struct [CheckDeposit]
type checkDepositJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	Amount            apijson.Field
	BackImageFileID   apijson.Field
	CreatedAt         apijson.Field
	DepositAcceptance apijson.Field
	DepositRejection  apijson.Field
	DepositReturn     apijson.Field
	DepositSubmission apijson.Field
	Description       apijson.Field
	FrontImageFileID  apijson.Field
	IdempotencyKey    apijson.Field
	InboundFundsHold  apijson.Field
	InboundMailItemID apijson.Field
	LockboxID         apijson.Field
	Status            apijson.Field
	TransactionID     apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CheckDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkDepositJSON) RawJSON() string {
	return r.raw
}

// Once your deposit is successfully parsed and accepted by Increase, this will
// contain details of the parsed check.
type CheckDepositDepositAcceptance struct {
	// The account number printed on the check. This is an account at the bank that
	// issued the check.
	AccountNumber string `json:"account_number,required"`
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CheckDepositDepositAcceptanceCurrency `json:"currency,required"`
	// The routing number printed on the check. This is a routing number for the bank
	// that issued the check.
	RoutingNumber string `json:"routing_number,required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string                            `json:"serial_number,required,nullable"`
	JSON         checkDepositDepositAcceptanceJSON `json:"-"`
}

// checkDepositDepositAcceptanceJSON contains the JSON metadata for the struct
// [CheckDepositDepositAcceptance]
type checkDepositDepositAcceptanceJSON struct {
	AccountNumber  apijson.Field
	Amount         apijson.Field
	AuxiliaryOnUs  apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	RoutingNumber  apijson.Field
	SerialNumber   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CheckDepositDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkDepositDepositAcceptanceJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CheckDepositDepositAcceptanceCurrency string

const (
	CheckDepositDepositAcceptanceCurrencyCad CheckDepositDepositAcceptanceCurrency = "CAD"
	CheckDepositDepositAcceptanceCurrencyChf CheckDepositDepositAcceptanceCurrency = "CHF"
	CheckDepositDepositAcceptanceCurrencyEur CheckDepositDepositAcceptanceCurrency = "EUR"
	CheckDepositDepositAcceptanceCurrencyGbp CheckDepositDepositAcceptanceCurrency = "GBP"
	CheckDepositDepositAcceptanceCurrencyJpy CheckDepositDepositAcceptanceCurrency = "JPY"
	CheckDepositDepositAcceptanceCurrencyUsd CheckDepositDepositAcceptanceCurrency = "USD"
)

func (r CheckDepositDepositAcceptanceCurrency) IsKnown() bool {
	switch r {
	case CheckDepositDepositAcceptanceCurrencyCad, CheckDepositDepositAcceptanceCurrencyChf, CheckDepositDepositAcceptanceCurrencyEur, CheckDepositDepositAcceptanceCurrencyGbp, CheckDepositDepositAcceptanceCurrencyJpy, CheckDepositDepositAcceptanceCurrencyUsd:
		return true
	}
	return false
}

// If your deposit is rejected by Increase, this will contain details as to why it
// was rejected.
type CheckDepositDepositRejection struct {
	// The rejected amount in the minor unit of check's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was rejected.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CheckDepositDepositRejectionCurrency `json:"currency,required"`
	// The identifier of the associated declined transaction.
	DeclinedTransactionID string `json:"declined_transaction_id,required"`
	// Why the check deposit was rejected.
	Reason CheckDepositDepositRejectionReason `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt time.Time                        `json:"rejected_at,required" format:"date-time"`
	JSON       checkDepositDepositRejectionJSON `json:"-"`
}

// checkDepositDepositRejectionJSON contains the JSON metadata for the struct
// [CheckDepositDepositRejection]
type checkDepositDepositRejectionJSON struct {
	Amount                apijson.Field
	CheckDepositID        apijson.Field
	Currency              apijson.Field
	DeclinedTransactionID apijson.Field
	Reason                apijson.Field
	RejectedAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CheckDepositDepositRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkDepositDepositRejectionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type CheckDepositDepositRejectionCurrency string

const (
	CheckDepositDepositRejectionCurrencyCad CheckDepositDepositRejectionCurrency = "CAD"
	CheckDepositDepositRejectionCurrencyChf CheckDepositDepositRejectionCurrency = "CHF"
	CheckDepositDepositRejectionCurrencyEur CheckDepositDepositRejectionCurrency = "EUR"
	CheckDepositDepositRejectionCurrencyGbp CheckDepositDepositRejectionCurrency = "GBP"
	CheckDepositDepositRejectionCurrencyJpy CheckDepositDepositRejectionCurrency = "JPY"
	CheckDepositDepositRejectionCurrencyUsd CheckDepositDepositRejectionCurrency = "USD"
)

func (r CheckDepositDepositRejectionCurrency) IsKnown() bool {
	switch r {
	case CheckDepositDepositRejectionCurrencyCad, CheckDepositDepositRejectionCurrencyChf, CheckDepositDepositRejectionCurrencyEur, CheckDepositDepositRejectionCurrencyGbp, CheckDepositDepositRejectionCurrencyJpy, CheckDepositDepositRejectionCurrencyUsd:
		return true
	}
	return false
}

// Why the check deposit was rejected.
type CheckDepositDepositRejectionReason string

const (
	CheckDepositDepositRejectionReasonIncompleteImage             CheckDepositDepositRejectionReason = "incomplete_image"
	CheckDepositDepositRejectionReasonDuplicate                   CheckDepositDepositRejectionReason = "duplicate"
	CheckDepositDepositRejectionReasonPoorImageQuality            CheckDepositDepositRejectionReason = "poor_image_quality"
	CheckDepositDepositRejectionReasonIncorrectAmount             CheckDepositDepositRejectionReason = "incorrect_amount"
	CheckDepositDepositRejectionReasonIncorrectRecipient          CheckDepositDepositRejectionReason = "incorrect_recipient"
	CheckDepositDepositRejectionReasonNotEligibleForMobileDeposit CheckDepositDepositRejectionReason = "not_eligible_for_mobile_deposit"
	CheckDepositDepositRejectionReasonMissingRequiredDataElements CheckDepositDepositRejectionReason = "missing_required_data_elements"
	CheckDepositDepositRejectionReasonSuspectedFraud              CheckDepositDepositRejectionReason = "suspected_fraud"
	CheckDepositDepositRejectionReasonDepositWindowExpired        CheckDepositDepositRejectionReason = "deposit_window_expired"
	CheckDepositDepositRejectionReasonRequestedByUser             CheckDepositDepositRejectionReason = "requested_by_user"
	CheckDepositDepositRejectionReasonUnknown                     CheckDepositDepositRejectionReason = "unknown"
)

func (r CheckDepositDepositRejectionReason) IsKnown() bool {
	switch r {
	case CheckDepositDepositRejectionReasonIncompleteImage, CheckDepositDepositRejectionReasonDuplicate, CheckDepositDepositRejectionReasonPoorImageQuality, CheckDepositDepositRejectionReasonIncorrectAmount, CheckDepositDepositRejectionReasonIncorrectRecipient, CheckDepositDepositRejectionReasonNotEligibleForMobileDeposit, CheckDepositDepositRejectionReasonMissingRequiredDataElements, CheckDepositDepositRejectionReasonSuspectedFraud, CheckDepositDepositRejectionReasonDepositWindowExpired, CheckDepositDepositRejectionReasonRequestedByUser, CheckDepositDepositRejectionReasonUnknown:
		return true
	}
	return false
}

// If your deposit is returned, this will contain details as to why it was
// returned.
type CheckDepositDepositReturn struct {
	// The returned amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CheckDepositDepositReturnCurrency `json:"currency,required"`
	// Why this check was returned by the bank holding the account it was drawn
	// against.
	ReturnReason CheckDepositDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                        `json:"transaction_id,required"`
	JSON          checkDepositDepositReturnJSON `json:"-"`
}

// checkDepositDepositReturnJSON contains the JSON metadata for the struct
// [CheckDepositDepositReturn]
type checkDepositDepositReturnJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	ReturnReason   apijson.Field
	ReturnedAt     apijson.Field
	TransactionID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CheckDepositDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkDepositDepositReturnJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type CheckDepositDepositReturnCurrency string

const (
	CheckDepositDepositReturnCurrencyCad CheckDepositDepositReturnCurrency = "CAD"
	CheckDepositDepositReturnCurrencyChf CheckDepositDepositReturnCurrency = "CHF"
	CheckDepositDepositReturnCurrencyEur CheckDepositDepositReturnCurrency = "EUR"
	CheckDepositDepositReturnCurrencyGbp CheckDepositDepositReturnCurrency = "GBP"
	CheckDepositDepositReturnCurrencyJpy CheckDepositDepositReturnCurrency = "JPY"
	CheckDepositDepositReturnCurrencyUsd CheckDepositDepositReturnCurrency = "USD"
)

func (r CheckDepositDepositReturnCurrency) IsKnown() bool {
	switch r {
	case CheckDepositDepositReturnCurrencyCad, CheckDepositDepositReturnCurrencyChf, CheckDepositDepositReturnCurrencyEur, CheckDepositDepositReturnCurrencyGbp, CheckDepositDepositReturnCurrencyJpy, CheckDepositDepositReturnCurrencyUsd:
		return true
	}
	return false
}

// Why this check was returned by the bank holding the account it was drawn
// against.
type CheckDepositDepositReturnReturnReason string

const (
	CheckDepositDepositReturnReturnReasonACHConversionNotSupported CheckDepositDepositReturnReturnReason = "ach_conversion_not_supported"
	CheckDepositDepositReturnReturnReasonClosedAccount             CheckDepositDepositReturnReturnReason = "closed_account"
	CheckDepositDepositReturnReturnReasonDuplicateSubmission       CheckDepositDepositReturnReturnReason = "duplicate_submission"
	CheckDepositDepositReturnReturnReasonInsufficientFunds         CheckDepositDepositReturnReturnReason = "insufficient_funds"
	CheckDepositDepositReturnReturnReasonNoAccount                 CheckDepositDepositReturnReturnReason = "no_account"
	CheckDepositDepositReturnReturnReasonNotAuthorized             CheckDepositDepositReturnReturnReason = "not_authorized"
	CheckDepositDepositReturnReturnReasonStaleDated                CheckDepositDepositReturnReturnReason = "stale_dated"
	CheckDepositDepositReturnReturnReasonStopPayment               CheckDepositDepositReturnReturnReason = "stop_payment"
	CheckDepositDepositReturnReturnReasonUnknownReason             CheckDepositDepositReturnReturnReason = "unknown_reason"
	CheckDepositDepositReturnReturnReasonUnmatchedDetails          CheckDepositDepositReturnReturnReason = "unmatched_details"
	CheckDepositDepositReturnReturnReasonUnreadableImage           CheckDepositDepositReturnReturnReason = "unreadable_image"
	CheckDepositDepositReturnReturnReasonEndorsementIrregular      CheckDepositDepositReturnReturnReason = "endorsement_irregular"
	CheckDepositDepositReturnReturnReasonAlteredOrFictitiousItem   CheckDepositDepositReturnReturnReason = "altered_or_fictitious_item"
	CheckDepositDepositReturnReturnReasonFrozenOrBlockedAccount    CheckDepositDepositReturnReturnReason = "frozen_or_blocked_account"
	CheckDepositDepositReturnReturnReasonPostDated                 CheckDepositDepositReturnReturnReason = "post_dated"
	CheckDepositDepositReturnReturnReasonEndorsementMissing        CheckDepositDepositReturnReturnReason = "endorsement_missing"
	CheckDepositDepositReturnReturnReasonSignatureMissing          CheckDepositDepositReturnReturnReason = "signature_missing"
	CheckDepositDepositReturnReturnReasonStopPaymentSuspect        CheckDepositDepositReturnReturnReason = "stop_payment_suspect"
	CheckDepositDepositReturnReturnReasonUnusableImage             CheckDepositDepositReturnReturnReason = "unusable_image"
	CheckDepositDepositReturnReturnReasonImageFailsSecurityCheck   CheckDepositDepositReturnReturnReason = "image_fails_security_check"
	CheckDepositDepositReturnReturnReasonCannotDetermineAmount     CheckDepositDepositReturnReturnReason = "cannot_determine_amount"
	CheckDepositDepositReturnReturnReasonSignatureIrregular        CheckDepositDepositReturnReturnReason = "signature_irregular"
	CheckDepositDepositReturnReturnReasonNonCashItem               CheckDepositDepositReturnReturnReason = "non_cash_item"
	CheckDepositDepositReturnReturnReasonUnableToProcess           CheckDepositDepositReturnReturnReason = "unable_to_process"
	CheckDepositDepositReturnReturnReasonItemExceedsDollarLimit    CheckDepositDepositReturnReturnReason = "item_exceeds_dollar_limit"
	CheckDepositDepositReturnReturnReasonBranchOrAccountSold       CheckDepositDepositReturnReturnReason = "branch_or_account_sold"
)

func (r CheckDepositDepositReturnReturnReason) IsKnown() bool {
	switch r {
	case CheckDepositDepositReturnReturnReasonACHConversionNotSupported, CheckDepositDepositReturnReturnReasonClosedAccount, CheckDepositDepositReturnReturnReasonDuplicateSubmission, CheckDepositDepositReturnReturnReasonInsufficientFunds, CheckDepositDepositReturnReturnReasonNoAccount, CheckDepositDepositReturnReturnReasonNotAuthorized, CheckDepositDepositReturnReturnReasonStaleDated, CheckDepositDepositReturnReturnReasonStopPayment, CheckDepositDepositReturnReturnReasonUnknownReason, CheckDepositDepositReturnReturnReasonUnmatchedDetails, CheckDepositDepositReturnReturnReasonUnreadableImage, CheckDepositDepositReturnReturnReasonEndorsementIrregular, CheckDepositDepositReturnReturnReasonAlteredOrFictitiousItem, CheckDepositDepositReturnReturnReasonFrozenOrBlockedAccount, CheckDepositDepositReturnReturnReasonPostDated, CheckDepositDepositReturnReturnReasonEndorsementMissing, CheckDepositDepositReturnReturnReasonSignatureMissing, CheckDepositDepositReturnReturnReasonStopPaymentSuspect, CheckDepositDepositReturnReturnReasonUnusableImage, CheckDepositDepositReturnReturnReasonImageFailsSecurityCheck, CheckDepositDepositReturnReturnReasonCannotDetermineAmount, CheckDepositDepositReturnReturnReasonSignatureIrregular, CheckDepositDepositReturnReturnReasonNonCashItem, CheckDepositDepositReturnReturnReasonUnableToProcess, CheckDepositDepositReturnReturnReasonItemExceedsDollarLimit, CheckDepositDepositReturnReturnReasonBranchOrAccountSold:
		return true
	}
	return false
}

// After the check is parsed, it is submitted to the Check21 network for
// processing. This will contain details of the submission.
type CheckDepositDepositSubmission struct {
	// The ID for the File containing the check back image that was submitted to the
	// Check21 network.
	BackFileID string `json:"back_file_id,required"`
	// The ID for the File containing the check front image that was submitted to the
	// Check21 network.
	FrontFileID string `json:"front_file_id,required"`
	// When the check deposit was submitted to the Check21 network for processing.
	// During business days, this happens within a few hours of the check being
	// accepted by Increase.
	SubmittedAt time.Time                         `json:"submitted_at,required" format:"date-time"`
	JSON        checkDepositDepositSubmissionJSON `json:"-"`
}

// checkDepositDepositSubmissionJSON contains the JSON metadata for the struct
// [CheckDepositDepositSubmission]
type checkDepositDepositSubmissionJSON struct {
	BackFileID  apijson.Field
	FrontFileID apijson.Field
	SubmittedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckDepositDepositSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkDepositDepositSubmissionJSON) RawJSON() string {
	return r.raw
}

// Increase will sometimes hold the funds for Check Deposits. If funds are held,
// this sub-object will contain details of the hold.
type CheckDepositInboundFundsHold struct {
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency CheckDepositInboundFundsHoldCurrency `json:"currency,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status CheckDepositInboundFundsHoldStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_funds_hold`.
	Type CheckDepositInboundFundsHoldType `json:"type,required"`
	JSON checkDepositInboundFundsHoldJSON `json:"-"`
}

// checkDepositInboundFundsHoldJSON contains the JSON metadata for the struct
// [CheckDepositInboundFundsHold]
type checkDepositInboundFundsHoldJSON struct {
	Amount                  apijson.Field
	AutomaticallyReleasesAt apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CheckDepositInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkDepositInboundFundsHoldJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
type CheckDepositInboundFundsHoldCurrency string

const (
	CheckDepositInboundFundsHoldCurrencyCad CheckDepositInboundFundsHoldCurrency = "CAD"
	CheckDepositInboundFundsHoldCurrencyChf CheckDepositInboundFundsHoldCurrency = "CHF"
	CheckDepositInboundFundsHoldCurrencyEur CheckDepositInboundFundsHoldCurrency = "EUR"
	CheckDepositInboundFundsHoldCurrencyGbp CheckDepositInboundFundsHoldCurrency = "GBP"
	CheckDepositInboundFundsHoldCurrencyJpy CheckDepositInboundFundsHoldCurrency = "JPY"
	CheckDepositInboundFundsHoldCurrencyUsd CheckDepositInboundFundsHoldCurrency = "USD"
)

func (r CheckDepositInboundFundsHoldCurrency) IsKnown() bool {
	switch r {
	case CheckDepositInboundFundsHoldCurrencyCad, CheckDepositInboundFundsHoldCurrencyChf, CheckDepositInboundFundsHoldCurrencyEur, CheckDepositInboundFundsHoldCurrencyGbp, CheckDepositInboundFundsHoldCurrencyJpy, CheckDepositInboundFundsHoldCurrencyUsd:
		return true
	}
	return false
}

// The status of the hold.
type CheckDepositInboundFundsHoldStatus string

const (
	CheckDepositInboundFundsHoldStatusHeld     CheckDepositInboundFundsHoldStatus = "held"
	CheckDepositInboundFundsHoldStatusComplete CheckDepositInboundFundsHoldStatus = "complete"
)

func (r CheckDepositInboundFundsHoldStatus) IsKnown() bool {
	switch r {
	case CheckDepositInboundFundsHoldStatusHeld, CheckDepositInboundFundsHoldStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_funds_hold`.
type CheckDepositInboundFundsHoldType string

const (
	CheckDepositInboundFundsHoldTypeInboundFundsHold CheckDepositInboundFundsHoldType = "inbound_funds_hold"
)

func (r CheckDepositInboundFundsHoldType) IsKnown() bool {
	switch r {
	case CheckDepositInboundFundsHoldTypeInboundFundsHold:
		return true
	}
	return false
}

// The status of the Check Deposit.
type CheckDepositStatus string

const (
	CheckDepositStatusPending   CheckDepositStatus = "pending"
	CheckDepositStatusSubmitted CheckDepositStatus = "submitted"
	CheckDepositStatusRejected  CheckDepositStatus = "rejected"
	CheckDepositStatusReturned  CheckDepositStatus = "returned"
)

func (r CheckDepositStatus) IsKnown() bool {
	switch r {
	case CheckDepositStatusPending, CheckDepositStatusSubmitted, CheckDepositStatusRejected, CheckDepositStatusReturned:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `check_deposit`.
type CheckDepositType string

const (
	CheckDepositTypeCheckDeposit CheckDepositType = "check_deposit"
)

func (r CheckDepositType) IsKnown() bool {
	switch r {
	case CheckDepositTypeCheckDeposit:
		return true
	}
	return false
}

type CheckDepositNewParams struct {
	// The identifier for the Account to deposit the check in.
	AccountID param.Field[string] `json:"account_id,required"`
	// The deposit amount in USD cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The File containing the check's back image.
	BackImageFileID param.Field[string] `json:"back_image_file_id,required"`
	// The File containing the check's front image.
	FrontImageFileID param.Field[string] `json:"front_image_file_id,required"`
	// The description you choose to give the Check Deposit, for display purposes only.
	Description param.Field[string] `json:"description"`
}

func (r CheckDepositNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckDepositListParams struct {
	// Filter Check Deposits to those belonging to the specified Account.
	AccountID param.Field[string]                          `query:"account_id"`
	CreatedAt param.Field[CheckDepositListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CheckDepositListParams]'s query parameters as `url.Values`.
func (r CheckDepositListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckDepositListParamsCreatedAt struct {
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

// URLQuery serializes [CheckDepositListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CheckDepositListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
