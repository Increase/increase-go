// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
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
	opts = append(r.Options[:], opts...)
	path := "check_deposits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Check Deposit
func (r *CheckDepositService) Get(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options, opts...)
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
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency CheckDepositCurrency `json:"currency,required"`
	// If your deposit is successfully parsed and accepted by Increase, this will
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
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
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
	Currency          apijson.Field
	DepositAcceptance apijson.Field
	DepositRejection  apijson.Field
	DepositReturn     apijson.Field
	DepositSubmission apijson.Field
	FrontImageFileID  apijson.Field
	IdempotencyKey    apijson.Field
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

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
type CheckDepositCurrency string

const (
	// Canadian Dollar (CAD)
	CheckDepositCurrencyCad CheckDepositCurrency = "CAD"
	// Swiss Franc (CHF)
	CheckDepositCurrencyChf CheckDepositCurrency = "CHF"
	// Euro (EUR)
	CheckDepositCurrencyEur CheckDepositCurrency = "EUR"
	// British Pound (GBP)
	CheckDepositCurrencyGbp CheckDepositCurrency = "GBP"
	// Japanese Yen (JPY)
	CheckDepositCurrencyJpy CheckDepositCurrency = "JPY"
	// US Dollar (USD)
	CheckDepositCurrencyUsd CheckDepositCurrency = "USD"
)

func (r CheckDepositCurrency) IsKnown() bool {
	switch r {
	case CheckDepositCurrencyCad, CheckDepositCurrencyChf, CheckDepositCurrencyEur, CheckDepositCurrencyGbp, CheckDepositCurrencyJpy, CheckDepositCurrencyUsd:
		return true
	}
	return false
}

// If your deposit is successfully parsed and accepted by Increase, this will
// contain details of the parsed check.
type CheckDepositDepositAcceptance struct {
	// The account number printed on the check.
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
	// The routing number printed on the check.
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
	// Canadian Dollar (CAD)
	CheckDepositDepositAcceptanceCurrencyCad CheckDepositDepositAcceptanceCurrency = "CAD"
	// Swiss Franc (CHF)
	CheckDepositDepositAcceptanceCurrencyChf CheckDepositDepositAcceptanceCurrency = "CHF"
	// Euro (EUR)
	CheckDepositDepositAcceptanceCurrencyEur CheckDepositDepositAcceptanceCurrency = "EUR"
	// British Pound (GBP)
	CheckDepositDepositAcceptanceCurrencyGbp CheckDepositDepositAcceptanceCurrency = "GBP"
	// Japanese Yen (JPY)
	CheckDepositDepositAcceptanceCurrencyJpy CheckDepositDepositAcceptanceCurrency = "JPY"
	// US Dollar (USD)
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
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	Reason         apijson.Field
	RejectedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
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
	// Canadian Dollar (CAD)
	CheckDepositDepositRejectionCurrencyCad CheckDepositDepositRejectionCurrency = "CAD"
	// Swiss Franc (CHF)
	CheckDepositDepositRejectionCurrencyChf CheckDepositDepositRejectionCurrency = "CHF"
	// Euro (EUR)
	CheckDepositDepositRejectionCurrencyEur CheckDepositDepositRejectionCurrency = "EUR"
	// British Pound (GBP)
	CheckDepositDepositRejectionCurrencyGbp CheckDepositDepositRejectionCurrency = "GBP"
	// Japanese Yen (JPY)
	CheckDepositDepositRejectionCurrencyJpy CheckDepositDepositRejectionCurrency = "JPY"
	// US Dollar (USD)
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
	// The check's image is incomplete.
	CheckDepositDepositRejectionReasonIncompleteImage CheckDepositDepositRejectionReason = "incomplete_image"
	// This is a duplicate check submission.
	CheckDepositDepositRejectionReasonDuplicate CheckDepositDepositRejectionReason = "duplicate"
	// This check has poor image quality.
	CheckDepositDepositRejectionReasonPoorImageQuality CheckDepositDepositRejectionReason = "poor_image_quality"
	// The check was deposited with the incorrect amount.
	CheckDepositDepositRejectionReasonIncorrectAmount CheckDepositDepositRejectionReason = "incorrect_amount"
	// The check is made out to someone other than the account holder.
	CheckDepositDepositRejectionReasonIncorrectRecipient CheckDepositDepositRejectionReason = "incorrect_recipient"
	// This check was not eligible for mobile deposit.
	CheckDepositDepositRejectionReasonNotEligibleForMobileDeposit CheckDepositDepositRejectionReason = "not_eligible_for_mobile_deposit"
	// This check is missing at least one required field.
	CheckDepositDepositRejectionReasonMissingRequiredDataElements CheckDepositDepositRejectionReason = "missing_required_data_elements"
	// This check is suspected to be fraudulent.
	CheckDepositDepositRejectionReasonSuspectedFraud CheckDepositDepositRejectionReason = "suspected_fraud"
	// This check's deposit window has expired.
	CheckDepositDepositRejectionReasonDepositWindowExpired CheckDepositDepositRejectionReason = "deposit_window_expired"
	// The check was rejected for an unknown reason.
	CheckDepositDepositRejectionReasonUnknown CheckDepositDepositRejectionReason = "unknown"
)

func (r CheckDepositDepositRejectionReason) IsKnown() bool {
	switch r {
	case CheckDepositDepositRejectionReasonIncompleteImage, CheckDepositDepositRejectionReasonDuplicate, CheckDepositDepositRejectionReasonPoorImageQuality, CheckDepositDepositRejectionReasonIncorrectAmount, CheckDepositDepositRejectionReasonIncorrectRecipient, CheckDepositDepositRejectionReasonNotEligibleForMobileDeposit, CheckDepositDepositRejectionReasonMissingRequiredDataElements, CheckDepositDepositRejectionReasonSuspectedFraud, CheckDepositDepositRejectionReasonDepositWindowExpired, CheckDepositDepositRejectionReasonUnknown:
		return true
	}
	return false
}

// If your deposit is returned, this will contain details as to why it was
// returned.
type CheckDepositDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
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
	// Canadian Dollar (CAD)
	CheckDepositDepositReturnCurrencyCad CheckDepositDepositReturnCurrency = "CAD"
	// Swiss Franc (CHF)
	CheckDepositDepositReturnCurrencyChf CheckDepositDepositReturnCurrency = "CHF"
	// Euro (EUR)
	CheckDepositDepositReturnCurrencyEur CheckDepositDepositReturnCurrency = "EUR"
	// British Pound (GBP)
	CheckDepositDepositReturnCurrencyGbp CheckDepositDepositReturnCurrency = "GBP"
	// Japanese Yen (JPY)
	CheckDepositDepositReturnCurrencyJpy CheckDepositDepositReturnCurrency = "JPY"
	// US Dollar (USD)
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
	// The check doesn't allow ACH conversion.
	CheckDepositDepositReturnReturnReasonACHConversionNotSupported CheckDepositDepositReturnReturnReason = "ach_conversion_not_supported"
	// The account is closed.
	CheckDepositDepositReturnReturnReasonClosedAccount CheckDepositDepositReturnReturnReason = "closed_account"
	// The check has already been deposited.
	CheckDepositDepositReturnReturnReasonDuplicateSubmission CheckDepositDepositReturnReturnReason = "duplicate_submission"
	// Insufficient funds
	CheckDepositDepositReturnReturnReasonInsufficientFunds CheckDepositDepositReturnReturnReason = "insufficient_funds"
	// No account was found matching the check details.
	CheckDepositDepositReturnReturnReasonNoAccount CheckDepositDepositReturnReturnReason = "no_account"
	// The check was not authorized.
	CheckDepositDepositReturnReturnReasonNotAuthorized CheckDepositDepositReturnReturnReason = "not_authorized"
	// The check is too old.
	CheckDepositDepositReturnReturnReasonStaleDated CheckDepositDepositReturnReturnReason = "stale_dated"
	// The payment has been stopped by the account holder.
	CheckDepositDepositReturnReturnReasonStopPayment CheckDepositDepositReturnReturnReason = "stop_payment"
	// The reason for the return is unknown.
	CheckDepositDepositReturnReturnReasonUnknownReason CheckDepositDepositReturnReturnReason = "unknown_reason"
	// The image doesn't match the details submitted.
	CheckDepositDepositReturnReturnReasonUnmatchedDetails CheckDepositDepositReturnReturnReason = "unmatched_details"
	// The image could not be read.
	CheckDepositDepositReturnReturnReasonUnreadableImage CheckDepositDepositReturnReturnReason = "unreadable_image"
	// The check endorsement was irregular.
	CheckDepositDepositReturnReturnReasonEndorsementIrregular CheckDepositDepositReturnReturnReason = "endorsement_irregular"
	// The check present was either altered or fake.
	CheckDepositDepositReturnReturnReasonAlteredOrFictitiousItem CheckDepositDepositReturnReturnReason = "altered_or_fictitious_item"
	// The account this check is drawn on is frozen.
	CheckDepositDepositReturnReturnReasonFrozenOrBlockedAccount CheckDepositDepositReturnReturnReason = "frozen_or_blocked_account"
	// The check is post dated.
	CheckDepositDepositReturnReturnReasonPostDated CheckDepositDepositReturnReturnReason = "post_dated"
	// The endorsement was missing.
	CheckDepositDepositReturnReturnReasonEndorsementMissing CheckDepositDepositReturnReturnReason = "endorsement_missing"
	// The check signature was missing.
	CheckDepositDepositReturnReturnReasonSignatureMissing CheckDepositDepositReturnReturnReason = "signature_missing"
	// The bank suspects a stop payment will be placed.
	CheckDepositDepositReturnReturnReasonStopPaymentSuspect CheckDepositDepositReturnReturnReason = "stop_payment_suspect"
	// The bank cannot read the image.
	CheckDepositDepositReturnReturnReasonUnusableImage CheckDepositDepositReturnReturnReason = "unusable_image"
	// The check image fails the bank's security check.
	CheckDepositDepositReturnReturnReasonImageFailsSecurityCheck CheckDepositDepositReturnReturnReason = "image_fails_security_check"
	// The bank cannot determine the amount.
	CheckDepositDepositReturnReturnReasonCannotDetermineAmount CheckDepositDepositReturnReturnReason = "cannot_determine_amount"
	// The signature is inconsistent with prior signatures.
	CheckDepositDepositReturnReturnReasonSignatureIrregular CheckDepositDepositReturnReturnReason = "signature_irregular"
	// The check is a non-cash item and cannot be drawn against the account.
	CheckDepositDepositReturnReturnReasonNonCashItem CheckDepositDepositReturnReturnReason = "non_cash_item"
	// The bank is unable to process this check.
	CheckDepositDepositReturnReturnReasonUnableToProcess CheckDepositDepositReturnReturnReason = "unable_to_process"
	// The check exceeds the bank or customer's limit.
	CheckDepositDepositReturnReturnReasonItemExceedsDollarLimit CheckDepositDepositReturnReturnReason = "item_exceeds_dollar_limit"
	// The bank sold this account and no longer services this customer.
	CheckDepositDepositReturnReturnReasonBranchOrAccountSold CheckDepositDepositReturnReturnReason = "branch_or_account_sold"
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
	// When the check deposit was submitted to the Check21 network for processing.
	// During business days, this happens within a few hours of the check being
	// accepted by Increase.
	SubmittedAt time.Time                         `json:"submitted_at,required" format:"date-time"`
	JSON        checkDepositDepositSubmissionJSON `json:"-"`
}

// checkDepositDepositSubmissionJSON contains the JSON metadata for the struct
// [CheckDepositDepositSubmission]
type checkDepositDepositSubmissionJSON struct {
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

// The status of the Check Deposit.
type CheckDepositStatus string

const (
	// The Check Deposit is pending review.
	CheckDepositStatusPending CheckDepositStatus = "pending"
	// The Check Deposit has been deposited.
	CheckDepositStatusSubmitted CheckDepositStatus = "submitted"
	// The Check Deposit has been rejected.
	CheckDepositStatusRejected CheckDepositStatus = "rejected"
	// The Check Deposit has been returned.
	CheckDepositStatusReturned CheckDepositStatus = "returned"
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
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The File containing the check's back image.
	BackImageFileID param.Field[string] `json:"back_image_file_id,required"`
	// The currency to use for the deposit.
	Currency param.Field[string] `json:"currency,required"`
	// The File containing the check's front image.
	FrontImageFileID param.Field[string] `json:"front_image_file_id,required"`
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
