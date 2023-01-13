package wire_transfers

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type WireTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewWireTransferService(requester core.Requester) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedWireTransferService struct {
	WireTransfers *WireTransferService
}

func (r *PreloadedWireTransferService) Init(service *WireTransferService) {
	r.WireTransfers = service
}

func NewPreloadedWireTransferService(service *WireTransferService) (r *PreloadedWireTransferService) {
	r = &PreloadedWireTransferService{}
	r.Init(service)
	return
}

//
type WireTransfer struct {
	// The wire transfer's identifier.
	Id *string `json:"id"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient"`
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency *WireTransferCurrency `json:"currency"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
	// The Account to which the transfer belongs.
	AccountId *string `json:"account_id"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountId *string `json:"external_account_id"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *WireTransferApproval `json:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *WireTransferCancellation `json:"cancellation"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal *WireTransferReversal `json:"reversal"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// The transfer's network.
	Network *WireTransferNetwork `json:"network"`
	// The lifecycle status of the transfer.
	Status *WireTransferStatus `json:"status"`
	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission *WireTransferSubmission `json:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateId *string `json:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionId *string `json:"transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type *WireTransferType `json:"type"`
}

// The wire transfer's identifier.
func (r *WireTransfer) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *WireTransfer) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The transfer amount in USD cents.
func (r *WireTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For wire transfers this is always equal to `usd`.
func (r *WireTransfer) GetCurrency() (Currency WireTransferCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The destination account number.
func (r *WireTransfer) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The Account to which the transfer belongs.
func (r *WireTransfer) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The identifier of the External Account the transfer was made to, if any.
func (r *WireTransfer) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *WireTransfer) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *WireTransfer) GetApproval() (Approval WireTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *WireTransfer) GetCancellation() (Cancellation WireTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// If your transfer is reversed, this will contain details of the reversal.
func (r *WireTransfer) GetReversal() (Reversal WireTransferReversal) {
	if r != nil && r.Reversal != nil {
		Reversal = *r.Reversal
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *WireTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The transfer's network.
func (r *WireTransfer) GetNetwork() (Network WireTransferNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// The lifecycle status of the transfer.
func (r *WireTransfer) GetStatus() (Status WireTransferStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// After the transfer is submitted to Fedwire, this will contain supplemental
// details.
func (r *WireTransfer) GetSubmission() (Submission WireTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *WireTransfer) GetTemplateId() (TemplateId string) {
	if r != nil && r.TemplateId != nil {
		TemplateId = *r.TemplateId
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *WireTransfer) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `wire_transfer`.
func (r *WireTransfer) GetType() (Type WireTransferType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
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

//
type WireTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt *string `json:"approved_at"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r *WireTransferApproval) GetApprovedAt() (ApprovedAt string) {
	if r != nil && r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}

//
type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string `json:"canceled_at"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r *WireTransferCancellation) GetCanceledAt() (CanceledAt string) {
	if r != nil && r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}

//
type WireTransferReversal struct {
	// The amount that was reversed.
	Amount *int `json:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `json:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `json:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `json:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `json:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `json:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `json:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `json:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string `json:"previous_message_input_source"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation *string `json:"receiver_financial_institution_information"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation *string `json:"financial_institution_to_financial_institution_information"`
}

// The amount that was reversed.
func (r *WireTransferReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *WireTransferReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *WireTransferReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *WireTransferReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *WireTransferReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *WireTransferReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *WireTransferReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *WireTransferReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *WireTransferReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *WireTransferReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *WireTransferReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *WireTransferReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
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

//
type WireTransferSubmission struct {
	// The accountability data for the submission.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data"`
}

// The accountability data for the submission.
func (r *WireTransferSubmission) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

type CreateAWireTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountId *string `json:"account_id"`
	// The account number for the destination account.
	AccountNumber *string `json:"account_number,omitempty"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number,omitempty"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountId *string `json:"external_account_id,omitempty"`
	// The transfer amount in cents.
	Amount *int `json:"amount"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient"`
	// The beneficiary's name.
	BeneficiaryName *string `json:"beneficiary_name,omitempty"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 *string `json:"beneficiary_address_line1,omitempty"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 *string `json:"beneficiary_address_line2,omitempty"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 *string `json:"beneficiary_address_line3,omitempty"`
}

// The identifier for the account that will send the transfer.
func (r *CreateAWireTransferParameters) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The account number for the destination account.
func (r *CreateAWireTransferParameters) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *CreateAWireTransferParameters) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The ID of an External Account to initiate a transfer to. If this parameter is
// provided, `account_number` and `routing_number` must be absent.
func (r *CreateAWireTransferParameters) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

// The transfer amount in cents.
func (r *CreateAWireTransferParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *CreateAWireTransferParameters) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The beneficiary's name.
func (r *CreateAWireTransferParameters) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// The beneficiary's address line 1.
func (r *CreateAWireTransferParameters) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// The beneficiary's address line 2.
func (r *CreateAWireTransferParameters) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// The beneficiary's address line 3.
func (r *CreateAWireTransferParameters) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

type ListWireTransfersQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountId *string `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountId *string                          `query:"external_account_id"`
	CreatedAt         *ListWireTransfersQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListWireTransfersQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListWireTransfersQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Wire Transfers to those belonging to the specified Account.
func (r *ListWireTransfersQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// Filter Wire Transfers to those made to the specified External Account.
func (r *ListWireTransfersQuery) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

func (r *ListWireTransfersQuery) GetCreatedAt() (CreatedAt ListWireTransfersQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListWireTransfersQueryCreatedAt struct {
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
func (r *ListWireTransfersQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListWireTransfersQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListWireTransfersQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListWireTransfersQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type WireTransferList struct {
	// The contents of the list.
	Data *[]WireTransfer `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *WireTransferList) GetData() (Data []WireTransfer) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *WireTransferList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *WireTransferService) Create(ctx context.Context, body *CreateAWireTransferParameters, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.post(
		ctx,
		"/wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedWireTransferService) Create(ctx context.Context, body *CreateAWireTransferParameters, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.WireTransfers.post(
		ctx,
		"/wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *WireTransferService) Retrieve(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/wire_transfers/%s", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedWireTransferService) Retrieve(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.WireTransfers.get(
		ctx,
		fmt.Sprintf("/wire_transfers/%s", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type WireTransfersPage struct {
	*pagination.Page[WireTransfer]
}

func (r *WireTransfersPage) WireTransfer() *WireTransfer {
	return r.Current()
}

func (r *WireTransfersPage) GetNextPage() (*WireTransfersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &WireTransfersPage{page}, nil
	}
}

func (r *WireTransferService) List(ctx context.Context, query *ListWireTransfersQuery, opts ...*core.RequestOpts) (res *WireTransfersPage, err error) {
	page := &WireTransfersPage{
		Page: &pagination.Page[WireTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/wire_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedWireTransferService) List(ctx context.Context, query *ListWireTransfersQuery, opts ...*core.RequestOpts) (res *WireTransfersPage, err error) {
	page := &WireTransfersPage{
		Page: &pagination.Page[WireTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/wire_transfers",
			},
			Requester: r.WireTransfers.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Simulates the reversal of an Wire Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `complete`.
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/reverse", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the reversal of an Wire Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `complete`.
func (r *PreloadedWireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.WireTransfers.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/reverse", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the submission of a Wire Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/submit", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the submission of a Wire Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or `pending_creating`.
func (r *PreloadedWireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.WireTransfers.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/submit", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}
