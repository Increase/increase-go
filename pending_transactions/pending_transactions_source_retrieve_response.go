package pending_transactions

type PendingTransactionsSourceRetrieveResponse struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *string `json:"category,omitempty"`

	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *PendingTransactionsSourceAccountTransferInstructionRetrieveResponse `json:"account_transfer_instruction,omitempty"`

	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *PendingTransactionsSourceACHTransferInstructionRetrieveResponse `json:"ach_transfer_instruction,omitempty"`

	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *PendingTransactionsSourceCardAuthorizationRetrieveResponse `json:"card_authorization,omitempty"`

	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *PendingTransactionsSourceCheckDepositInstructionRetrieveResponse `json:"check_deposit_instruction,omitempty"`

	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *PendingTransactionsSourceCheckTransferInstructionRetrieveResponse `json:"check_transfer_instruction,omitempty"`

	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *PendingTransactionsSourceCardRouteAuthorizationRetrieveResponse `json:"card_route_authorization,omitempty"`

	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse `json:"wire_drawdown_payment_instruction,omitempty"`

	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *PendingTransactionsSourceWireTransferInstructionRetrieveResponse `json:"wire_transfer_instruction,omitempty"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *PendingTransactionsSourceRetrieveResponse) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
func (r *PendingTransactionsSourceRetrieveResponse) GetAccountTransferInstruction() (AccountTransferInstruction PendingTransactionsSourceAccountTransferInstructionRetrieveResponse) {
	if r != nil && r.AccountTransferInstruction != nil {
		AccountTransferInstruction = *r.AccountTransferInstruction
	}
	return
}

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
func (r *PendingTransactionsSourceRetrieveResponse) GetACHTransferInstruction() (ACHTransferInstruction PendingTransactionsSourceACHTransferInstructionRetrieveResponse) {
	if r != nil && r.ACHTransferInstruction != nil {
		ACHTransferInstruction = *r.ACHTransferInstruction
	}
	return
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
func (r *PendingTransactionsSourceRetrieveResponse) GetCardAuthorization() (CardAuthorization PendingTransactionsSourceCardAuthorizationRetrieveResponse) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
func (r *PendingTransactionsSourceRetrieveResponse) GetCheckDepositInstruction() (CheckDepositInstruction PendingTransactionsSourceCheckDepositInstructionRetrieveResponse) {
	if r != nil && r.CheckDepositInstruction != nil {
		CheckDepositInstruction = *r.CheckDepositInstruction
	}
	return
}

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
func (r *PendingTransactionsSourceRetrieveResponse) GetCheckTransferInstruction() (CheckTransferInstruction PendingTransactionsSourceCheckTransferInstructionRetrieveResponse) {
	if r != nil && r.CheckTransferInstruction != nil {
		CheckTransferInstruction = *r.CheckTransferInstruction
	}
	return
}

// A Deprecated Card Authorization object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_authorization`.
func (r *PendingTransactionsSourceRetrieveResponse) GetCardRouteAuthorization() (CardRouteAuthorization PendingTransactionsSourceCardRouteAuthorizationRetrieveResponse) {
	if r != nil && r.CardRouteAuthorization != nil {
		CardRouteAuthorization = *r.CardRouteAuthorization
	}
	return
}

// A Wire Drawdown Payment Instruction object. This field will be present in the
// JSON response if and only if `category` is equal to
// `wire_drawdown_payment_instruction`.
func (r *PendingTransactionsSourceRetrieveResponse) GetWireDrawdownPaymentInstruction() (WireDrawdownPaymentInstruction PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse) {
	if r != nil && r.WireDrawdownPaymentInstruction != nil {
		WireDrawdownPaymentInstruction = *r.WireDrawdownPaymentInstruction
	}
	return
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
func (r *PendingTransactionsSourceRetrieveResponse) GetWireTransferInstruction() (WireTransferInstruction PendingTransactionsSourceWireTransferInstructionRetrieveResponse) {
	if r != nil && r.WireTransferInstruction != nil {
		WireTransferInstruction = *r.WireTransferInstruction
	}
	return
}
