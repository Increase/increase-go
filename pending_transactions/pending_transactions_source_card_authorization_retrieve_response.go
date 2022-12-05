package pending_transactions

type PendingTransactionsSourceCardAuthorizationRetrieveResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	MerchantAcceptorId *string `json:"merchant_acceptor_id,omitempty"`

	MerchantCity *string `json:"merchant_city,omitempty"`

	MerchantCountry *string `json:"merchant_country,omitempty"`

	MerchantDescriptor *string `json:"merchant_descriptor,omitempty"`

	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`

	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionId *string `json:"real_time_decision_id,omitempty"`

	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetRealTimeDecisionId() (RealTimeDecisionId string) {
	if r != nil && r.RealTimeDecisionId != nil {
		RealTimeDecisionId = *r.RealTimeDecisionId
	}
	return
}

// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
// purchase), the identifier of the token that was used.
func (r *PendingTransactionsSourceCardAuthorizationRetrieveResponse) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}
