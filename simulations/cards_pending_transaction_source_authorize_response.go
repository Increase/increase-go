package simulations

type CardsPendingTransactionSourceAuthorizeResponse struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *string `json:"category,omitempty"`

	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *CardsPendingTransactionSourceAccountTransferInstructionAuthorizeResponse `json:"account_transfer_instruction,omitempty"`

	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *CardsPendingTransactionSourceACHTransferInstructionAuthorizeResponse `json:"ach_transfer_instruction,omitempty"`

	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *CardsPendingTransactionSourceCardAuthorizationAuthorizeResponse `json:"card_authorization,omitempty"`

	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *CardsPendingTransactionSourceCheckDepositInstructionAuthorizeResponse `json:"check_deposit_instruction,omitempty"`

	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *CardsPendingTransactionSourceCheckTransferInstructionAuthorizeResponse `json:"check_transfer_instruction,omitempty"`

	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *CardsPendingTransactionSourceCardRouteAuthorizationAuthorizeResponse `json:"card_route_authorization,omitempty"`

	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse `json:"wire_drawdown_payment_instruction,omitempty"`

	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *CardsPendingTransactionSourceWireTransferInstructionAuthorizeResponse `json:"wire_transfer_instruction,omitempty"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetAccountTransferInstruction() (AccountTransferInstruction CardsPendingTransactionSourceAccountTransferInstructionAuthorizeResponse) {
	if r != nil && r.AccountTransferInstruction != nil {
		AccountTransferInstruction = *r.AccountTransferInstruction
	}
	return
}

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetACHTransferInstruction() (ACHTransferInstruction CardsPendingTransactionSourceACHTransferInstructionAuthorizeResponse) {
	if r != nil && r.ACHTransferInstruction != nil {
		ACHTransferInstruction = *r.ACHTransferInstruction
	}
	return
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetCardAuthorization() (CardAuthorization CardsPendingTransactionSourceCardAuthorizationAuthorizeResponse) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetCheckDepositInstruction() (CheckDepositInstruction CardsPendingTransactionSourceCheckDepositInstructionAuthorizeResponse) {
	if r != nil && r.CheckDepositInstruction != nil {
		CheckDepositInstruction = *r.CheckDepositInstruction
	}
	return
}

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetCheckTransferInstruction() (CheckTransferInstruction CardsPendingTransactionSourceCheckTransferInstructionAuthorizeResponse) {
	if r != nil && r.CheckTransferInstruction != nil {
		CheckTransferInstruction = *r.CheckTransferInstruction
	}
	return
}

// A Deprecated Card Authorization object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_authorization`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetCardRouteAuthorization() (CardRouteAuthorization CardsPendingTransactionSourceCardRouteAuthorizationAuthorizeResponse) {
	if r != nil && r.CardRouteAuthorization != nil {
		CardRouteAuthorization = *r.CardRouteAuthorization
	}
	return
}

// A Wire Drawdown Payment Instruction object. This field will be present in the
// JSON response if and only if `category` is equal to
// `wire_drawdown_payment_instruction`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetWireDrawdownPaymentInstruction() (WireDrawdownPaymentInstruction CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse) {
	if r != nil && r.WireDrawdownPaymentInstruction != nil {
		WireDrawdownPaymentInstruction = *r.WireDrawdownPaymentInstruction
	}
	return
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
func (r *CardsPendingTransactionSourceAuthorizeResponse) GetWireTransferInstruction() (WireTransferInstruction CardsPendingTransactionSourceWireTransferInstructionAuthorizeResponse) {
	if r != nil && r.WireTransferInstruction != nil {
		WireTransferInstruction = *r.WireTransferInstruction
	}
	return
}
