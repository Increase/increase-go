package check_deposits

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type CheckDepositService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCheckDepositService(requester core.Requester) (r *CheckDepositService) {
	r = &CheckDepositService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type CheckDeposit struct {
	// The deposit's identifier.
	ID string `json:"id"`
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt string `json:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency CheckDepositCurrency `json:"currency"`
	// The status of the Check Deposit.
	Status CheckDepositStatus `json:"status"`
	// The Account the check was deposited into.
	AccountID string `json:"account_id"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID *string `json:"back_image_file_id"`
	// The ID for the Transaction created by the deposit.
	TransactionID *string `json:"transaction_id"`
	// If your deposit is successfully parsed and accepted by Increase, this will
	// contain details of the parsed check.
	DepositAcceptance *CheckDepositDepositAcceptance `json:"deposit_acceptance"`
	// If your deposit is rejected by Increase, this will contain details as to why it
	// was rejected.
	DepositRejection *CheckDepositDepositRejection `json:"deposit_rejection"`
	// If your deposit is returned, this will contain details as to why it was
	// returned.
	DepositReturn *CheckDepositDepositReturn `json:"deposit_return"`
	// A constant representing the object's type. For this resource it will always be
	// `check_deposit`.
	Type CheckDepositType `json:"type"`
}

// The ID for the File containing the image of the back of the check.
func (r *CheckDeposit) GetBackImageFileID() (BackImageFileID string) {
	if r != nil && r.BackImageFileID != nil {
		BackImageFileID = *r.BackImageFileID
	}
	return
}

// The ID for the Transaction created by the deposit.
func (r *CheckDeposit) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// If your deposit is successfully parsed and accepted by Increase, this will
// contain details of the parsed check.
func (r *CheckDeposit) GetDepositAcceptance() (DepositAcceptance CheckDepositDepositAcceptance) {
	if r != nil && r.DepositAcceptance != nil {
		DepositAcceptance = *r.DepositAcceptance
	}
	return
}

// If your deposit is rejected by Increase, this will contain details as to why it
// was rejected.
func (r *CheckDeposit) GetDepositRejection() (DepositRejection CheckDepositDepositRejection) {
	if r != nil && r.DepositRejection != nil {
		DepositRejection = *r.DepositRejection
	}
	return
}

// If your deposit is returned, this will contain details as to why it was
// returned.
func (r *CheckDeposit) GetDepositReturn() (DepositReturn CheckDepositDepositReturn) {
	if r != nil && r.DepositReturn != nil {
		DepositReturn = *r.DepositReturn
	}
	return
}

type CheckDepositCurrency string

const (
	CheckDepositCurrencyCad CheckDepositCurrency = "CAD"
	CheckDepositCurrencyChf CheckDepositCurrency = "CHF"
	CheckDepositCurrencyEur CheckDepositCurrency = "EUR"
	CheckDepositCurrencyGbp CheckDepositCurrency = "GBP"
	CheckDepositCurrencyJpy CheckDepositCurrency = "JPY"
	CheckDepositCurrencyUsd CheckDepositCurrency = "USD"
)

type CheckDepositStatus string

const (
	CheckDepositStatusPending   CheckDepositStatus = "pending"
	CheckDepositStatusSubmitted CheckDepositStatus = "submitted"
	CheckDepositStatusRejected  CheckDepositStatus = "rejected"
	CheckDepositStatusReturned  CheckDepositStatus = "returned"
)

//
type CheckDepositDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CheckDepositDepositAcceptanceCurrency `json:"currency"`
	// The account number printed on the check.
	AccountNumber string `json:"account_number"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number.
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
}

// An additional line of metadata printed on the check. This typically includes the
// check number.
func (r *CheckDepositDepositAcceptance) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

type CheckDepositDepositAcceptanceCurrency string

const (
	CheckDepositDepositAcceptanceCurrencyCad CheckDepositDepositAcceptanceCurrency = "CAD"
	CheckDepositDepositAcceptanceCurrencyChf CheckDepositDepositAcceptanceCurrency = "CHF"
	CheckDepositDepositAcceptanceCurrencyEur CheckDepositDepositAcceptanceCurrency = "EUR"
	CheckDepositDepositAcceptanceCurrencyGbp CheckDepositDepositAcceptanceCurrency = "GBP"
	CheckDepositDepositAcceptanceCurrencyJpy CheckDepositDepositAcceptanceCurrency = "JPY"
	CheckDepositDepositAcceptanceCurrencyUsd CheckDepositDepositAcceptanceCurrency = "USD"
)

//
type CheckDepositDepositRejection struct {
	// The rejected amount in the minor unit of check's currency. For dollars, for
	// example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CheckDepositDepositRejectionCurrency `json:"currency"`
	// Why the check deposit was rejected.
	Reason CheckDepositDepositRejectionReason `json:"reason"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt string `json:"rejected_at"`
}

type CheckDepositDepositRejectionCurrency string

const (
	CheckDepositDepositRejectionCurrencyCad CheckDepositDepositRejectionCurrency = "CAD"
	CheckDepositDepositRejectionCurrencyChf CheckDepositDepositRejectionCurrency = "CHF"
	CheckDepositDepositRejectionCurrencyEur CheckDepositDepositRejectionCurrency = "EUR"
	CheckDepositDepositRejectionCurrencyGbp CheckDepositDepositRejectionCurrency = "GBP"
	CheckDepositDepositRejectionCurrencyJpy CheckDepositDepositRejectionCurrency = "JPY"
	CheckDepositDepositRejectionCurrencyUsd CheckDepositDepositRejectionCurrency = "USD"
)

type CheckDepositDepositRejectionReason string

const (
	CheckDepositDepositRejectionReasonIncompleteImage             CheckDepositDepositRejectionReason = "incomplete_image"
	CheckDepositDepositRejectionReasonDuplicate                   CheckDepositDepositRejectionReason = "duplicate"
	CheckDepositDepositRejectionReasonPoorImageQuality            CheckDepositDepositRejectionReason = "poor_image_quality"
	CheckDepositDepositRejectionReasonIncorrectAmount             CheckDepositDepositRejectionReason = "incorrect_amount"
	CheckDepositDepositRejectionReasonIncorrectRecipient          CheckDepositDepositRejectionReason = "incorrect_recipient"
	CheckDepositDepositRejectionReasonNotEligibleForMobileDeposit CheckDepositDepositRejectionReason = "not_eligible_for_mobile_deposit"
	CheckDepositDepositRejectionReasonUnknown                     CheckDepositDepositRejectionReason = "unknown"
)

//
type CheckDepositDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt string `json:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CheckDepositDepositReturnCurrency `json:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string `json:"transaction_id"`
	//
	ReturnReason CheckDepositDepositReturnReturnReason `json:"return_reason"`
}

type CheckDepositDepositReturnCurrency string

const (
	CheckDepositDepositReturnCurrencyCad CheckDepositDepositReturnCurrency = "CAD"
	CheckDepositDepositReturnCurrencyChf CheckDepositDepositReturnCurrency = "CHF"
	CheckDepositDepositReturnCurrencyEur CheckDepositDepositReturnCurrency = "EUR"
	CheckDepositDepositReturnCurrencyGbp CheckDepositDepositReturnCurrency = "GBP"
	CheckDepositDepositReturnCurrencyJpy CheckDepositDepositReturnCurrency = "JPY"
	CheckDepositDepositReturnCurrencyUsd CheckDepositDepositReturnCurrency = "USD"
)

type CheckDepositDepositReturnReturnReason string

const (
	CheckDepositDepositReturnReturnReasonACHConversionNotSupported CheckDepositDepositReturnReturnReason = "ach_conversion_not_supported"
	CheckDepositDepositReturnReturnReasonDuplicateSubmission       CheckDepositDepositReturnReturnReason = "duplicate_submission"
	CheckDepositDepositReturnReturnReasonInsufficientFunds         CheckDepositDepositReturnReturnReason = "insufficient_funds"
	CheckDepositDepositReturnReturnReasonNoAccount                 CheckDepositDepositReturnReturnReason = "no_account"
	CheckDepositDepositReturnReturnReasonNotAuthorized             CheckDepositDepositReturnReturnReason = "not_authorized"
	CheckDepositDepositReturnReturnReasonStaleDated                CheckDepositDepositReturnReturnReason = "stale_dated"
	CheckDepositDepositReturnReturnReasonStopPayment               CheckDepositDepositReturnReturnReason = "stop_payment"
	CheckDepositDepositReturnReturnReasonUnknownReason             CheckDepositDepositReturnReturnReason = "unknown_reason"
	CheckDepositDepositReturnReturnReasonUnmatchedDetails          CheckDepositDepositReturnReturnReason = "unmatched_details"
	CheckDepositDepositReturnReturnReasonUnreadableImage           CheckDepositDepositReturnReturnReason = "unreadable_image"
)

type CheckDepositType string

const (
	CheckDepositTypeCheckDeposit CheckDepositType = "check_deposit"
)

type CreateACheckDepositParameters struct {
	// The identifier for the Account to deposit the check in.
	AccountID string `json:"account_id"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount int `json:"amount"`
	// The currency to use for the deposit.
	Currency string `json:"currency"`
	// The File containing the check's front image.
	FrontImageFileID string `json:"front_image_file_id"`
	// The File containing the check's back image.
	BackImageFileID string `json:"back_image_file_id"`
}

type ListCheckDepositsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountID string                          `query:"account_id"`
	CreatedAt ListCheckDepositsQueryCreatedAt `query:"created_at"`
}

type ListCheckDepositsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore string `json:"on_or_before,omitempty"`
}

//
type CheckDepositList struct {
	// The contents of the list.
	Data []CheckDeposit `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *CheckDepositList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *CheckDepositService) Create(ctx context.Context, body *CreateACheckDepositParameters, opts ...*core.RequestOpts) (res *CheckDeposit, err error) {
	err = r.post(
		ctx,
		"/check_deposits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *CheckDepositService) Retrieve(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *CheckDeposit, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/check_deposits/%s", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type CheckDepositsPage struct {
	*pagination.Page[CheckDeposit]
}

func (r *CheckDepositsPage) CheckDeposit() *CheckDeposit {
	return r.Current()
}

func (r *CheckDepositsPage) GetNextPage() (*CheckDepositsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &CheckDepositsPage{page}, nil
	}
}

func (r *CheckDepositService) List(ctx context.Context, query *ListCheckDepositsQuery, opts ...*core.RequestOpts) (res *CheckDepositsPage, err error) {
	page := &CheckDepositsPage{
		Page: &pagination.Page[CheckDeposit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_deposits",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
