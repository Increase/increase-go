package pending_transactions

type PendingTransactionsDataSourceListResponse struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *string `json:"category,omitempty"`

	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *PendingTransactionsDataSourceAccountTransferInstructionListResponse `json:"account_transfer_instruction,omitempty"`

	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *PendingTransactionsDataSourceACHTransferInstructionListResponse `json:"ach_transfer_instruction,omitempty"`

	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *PendingTransactionsDataSourceCardAuthorizationListResponse `json:"card_authorization,omitempty"`

	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *PendingTransactionsDataSourceCheckDepositInstructionListResponse `json:"check_deposit_instruction,omitempty"`

	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *PendingTransactionsDataSourceCheckTransferInstructionListResponse `json:"check_transfer_instruction,omitempty"`

	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *PendingTransactionsDataSourceCardRouteAuthorizationListResponse `json:"card_route_authorization,omitempty"`

	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *PendingTransactionsDataSourceWireDrawdownPaymentInstructionListResponse `json:"wire_drawdown_payment_instruction,omitempty"`

	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *PendingTransactionsDataSourceWireTransferInstructionListResponse `json:"wire_transfer_instruction,omitempty"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *PendingTransactionsDataSourceListResponse) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
func (r *PendingTransactionsDataSourceListResponse) GetAccountTransferInstruction() (AccountTransferInstruction PendingTransactionsDataSourceAccountTransferInstructionListResponse) {
	if r != nil && r.AccountTransferInstruction != nil {
		AccountTransferInstruction = *r.AccountTransferInstruction
	}
	return
}

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
func (r *PendingTransactionsDataSourceListResponse) GetACHTransferInstruction() (ACHTransferInstruction PendingTransactionsDataSourceACHTransferInstructionListResponse) {
	if r != nil && r.ACHTransferInstruction != nil {
		ACHTransferInstruction = *r.ACHTransferInstruction
	}
	return
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
func (r *PendingTransactionsDataSourceListResponse) GetCardAuthorization() (CardAuthorization PendingTransactionsDataSourceCardAuthorizationListResponse) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
func (r *PendingTransactionsDataSourceListResponse) GetCheckDepositInstruction() (CheckDepositInstruction PendingTransactionsDataSourceCheckDepositInstructionListResponse) {
	if r != nil && r.CheckDepositInstruction != nil {
		CheckDepositInstruction = *r.CheckDepositInstruction
	}
	return
}

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
func (r *PendingTransactionsDataSourceListResponse) GetCheckTransferInstruction() (CheckTransferInstruction PendingTransactionsDataSourceCheckTransferInstructionListResponse) {
	if r != nil && r.CheckTransferInstruction != nil {
		CheckTransferInstruction = *r.CheckTransferInstruction
	}
	return
}

// A Deprecated Card Authorization object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_authorization`.
func (r *PendingTransactionsDataSourceListResponse) GetCardRouteAuthorization() (CardRouteAuthorization PendingTransactionsDataSourceCardRouteAuthorizationListResponse) {
	if r != nil && r.CardRouteAuthorization != nil {
		CardRouteAuthorization = *r.CardRouteAuthorization
	}
	return
}

// A Wire Drawdown Payment Instruction object. This field will be present in the
// JSON response if and only if `category` is equal to
// `wire_drawdown_payment_instruction`.
func (r *PendingTransactionsDataSourceListResponse) GetWireDrawdownPaymentInstruction() (WireDrawdownPaymentInstruction PendingTransactionsDataSourceWireDrawdownPaymentInstructionListResponse) {
	if r != nil && r.WireDrawdownPaymentInstruction != nil {
		WireDrawdownPaymentInstruction = *r.WireDrawdownPaymentInstruction
	}
	return
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
func (r *PendingTransactionsDataSourceListResponse) GetWireTransferInstruction() (WireTransferInstruction PendingTransactionsDataSourceWireTransferInstructionListResponse) {
	if r != nil && r.WireTransferInstruction != nil {
		WireTransferInstruction = *r.WireTransferInstruction
	}
	return
}
