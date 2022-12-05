package simulations

type ACHTransfersTransactionSourceCardSettlementCreateInboundResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	MerchantCity *string `json:"merchant_city,omitempty"`

	MerchantCountry *string `json:"merchant_country,omitempty"`

	MerchantName *string `json:"merchant_name,omitempty"`

	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`

	MerchantState *string `json:"merchant_state,omitempty"`

	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionId *string `json:"pending_transaction_id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type *string `json:"type,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetPendingTransactionId() (PendingTransactionId string) {
	if r != nil && r.PendingTransactionId != nil {
		PendingTransactionId = *r.PendingTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
