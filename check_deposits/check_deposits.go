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

type PreloadedCheckDepositService struct {
	CheckDeposits *CheckDepositService
}

func (r *PreloadedCheckDepositService) Init(service *CheckDepositService) {
	r.CheckDeposits = service
}

func NewPreloadedCheckDepositService(service *CheckDepositService) (r *PreloadedCheckDepositService) {
	r = &PreloadedCheckDepositService{}
	r.Init(service)
	return
}

//
type CheckDeposit struct {
	// The deposit's identifier.
	Id *string `json:"id"`
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency *CheckDepositCurrency `json:"currency"`
	// The status of the Check Deposit.
	Status *CheckDepositStatus `json:"status"`
	// The Account the check was deposited into.
	AccountId *string `json:"account_id"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileId *string `json:"front_image_file_id"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileId *string `json:"back_image_file_id"`
	// The ID for the Transaction created by the deposit.
	TransactionId *string `json:"transaction_id"`
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
	Type *CheckDepositType `json:"type"`
}

// The deposit's identifier.
func (r *CheckDeposit) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The deposited amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CheckDeposit) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *CheckDeposit) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
func (r *CheckDeposit) GetCurrency() (Currency CheckDepositCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The status of the Check Deposit.
func (r *CheckDeposit) GetStatus() (Status CheckDepositStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The Account the check was deposited into.
func (r *CheckDeposit) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The ID for the File containing the image of the front of the check.
func (r *CheckDeposit) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The ID for the File containing the image of the back of the check.
func (r *CheckDeposit) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
}

// The ID for the Transaction created by the deposit.
func (r *CheckDeposit) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
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

// A constant representing the object's type. For this resource it will always be
// `check_deposit`.
func (r *CheckDeposit) GetType() (Type CheckDepositType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
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
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CheckDepositDepositAcceptanceCurrency `json:"currency"`
	// The account number printed on the check.
	AccountNumber *string `json:"account_number"`
	// The routing number printed on the check.
	RoutingNumber *string `json:"routing_number"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number.
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
}

// The amount to be deposited in the minor unit of the transaction's currency. For
// dollars, for example, this is cents.
func (r *CheckDepositDepositAcceptance) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CheckDepositDepositAcceptance) GetCurrency() (Currency CheckDepositDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The account number printed on the check.
func (r *CheckDepositDepositAcceptance) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The routing number printed on the check.
func (r *CheckDepositDepositAcceptance) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
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
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *CheckDepositDepositRejectionCurrency `json:"currency"`
	// Why the check deposit was rejected.
	Reason *CheckDepositDepositRejectionReason `json:"reason"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt *string `json:"rejected_at"`
}

// The rejected amount in the minor unit of check's currency. For dollars, for
// example, this is cents.
func (r *CheckDepositDepositRejection) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *CheckDepositDepositRejection) GetCurrency() (Currency CheckDepositDepositRejectionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the check deposit was rejected.
func (r *CheckDepositDepositRejection) GetReason() (Reason CheckDepositDepositRejectionReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was rejected.
func (r *CheckDepositDepositRejection) GetRejectedAt() (RejectedAt string) {
	if r != nil && r.RejectedAt != nil {
		RejectedAt = *r.RejectedAt
	}
	return
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
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `json:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CheckDepositDepositReturnCurrency `json:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositId *string `json:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionId *string `json:"transaction_id"`
	//
	ReturnReason *CheckDepositDepositReturnReturnReason `json:"return_reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *CheckDepositDepositReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *CheckDepositDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CheckDepositDepositReturn) GetCurrency() (Currency CheckDepositDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *CheckDepositDepositReturn) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *CheckDepositDepositReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

func (r *CheckDepositDepositReturn) GetReturnReason() (ReturnReason CheckDepositDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
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
	AccountId *string `json:"account_id"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The currency to use for the deposit.
	Currency *string `json:"currency"`
	// The File containing the check's front image.
	FrontImageFileId *string `json:"front_image_file_id"`
	// The File containing the check's back image.
	BackImageFileId *string `json:"back_image_file_id"`
}

// The identifier for the Account to deposit the check in.
func (r *CreateACheckDepositParameters) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The deposit amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r *CreateACheckDepositParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The currency to use for the deposit.
func (r *CreateACheckDepositParameters) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The File containing the check's front image.
func (r *CreateACheckDepositParameters) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The File containing the check's back image.
func (r *CreateACheckDepositParameters) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
}

type ListCheckDepositsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountId *string                          `query:"account_id"`
	CreatedAt *ListCheckDepositsQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListCheckDepositsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListCheckDepositsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Deposits to those belonging to the specified Account.
func (r *ListCheckDepositsQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *ListCheckDepositsQuery) GetCreatedAt() (CreatedAt ListCheckDepositsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListCheckDepositsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCheckDepositsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type CheckDepositList struct {
	// The contents of the list.
	Data *[]CheckDeposit `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *CheckDepositList) GetData() (Data []CheckDeposit) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
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

func (r *PreloadedCheckDepositService) Create(ctx context.Context, body *CreateACheckDepositParameters, opts ...*core.RequestOpts) (res *CheckDeposit, err error) {
	err = r.CheckDeposits.post(
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

func (r *PreloadedCheckDepositService) Retrieve(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *CheckDeposit, err error) {
	err = r.CheckDeposits.get(
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

func (r *PreloadedCheckDepositService) List(ctx context.Context, query *ListCheckDepositsQuery, opts ...*core.RequestOpts) (res *CheckDepositsPage, err error) {
	page := &CheckDepositsPage{
		Page: &pagination.Page[CheckDeposit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_deposits",
			},
			Requester: r.CheckDeposits.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
