package real_time_decisions

type RealTimeDecisionsCardAuthorizationRetrieveResponse struct {
	// Whether or not the authorization was approved.
	Decision *string `json:"decision,omitempty"`

	// The identifier of the Card that is being authorized.
	CardId *string `json:"card_id,omitempty"`

	// The identifier of the Account the authorization will debit.
	AccountId *string `json:"account_id,omitempty"`

	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount *int64 `json:"presentment_amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency *string `json:"presentment_currency,omitempty"`

	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount *int64 `json:"settlement_amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency *string `json:"settlement_currency,omitempty"`

	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`

	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor *string `json:"merchant_descriptor,omitempty"`

	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorId *string `json:"merchant_acceptor_id,omitempty"`

	// The city the merchant resides in.
	MerchantCity *string `json:"merchant_city,omitempty"`

	// The country the merchant resides in.
	MerchantCountry *string `json:"merchant_country,omitempty"`
}

// Whether or not the authorization was approved.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetDecision() (Decision string) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
}

// The identifier of the Card that is being authorized.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The identifier of the Account the authorization will debit.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The amount of the attempted authorization in the currency the card user sees at
// the time of purchase, in the minor unit of that currency. For dollars, for
// example, this is cents.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetPresentmentAmount() (PresentmentAmount int64) {
	if r != nil && r.PresentmentAmount != nil {
		PresentmentAmount = *r.PresentmentAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
// user sees at the time of purchase.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetPresentmentCurrency() (PresentmentCurrency string) {
	if r != nil && r.PresentmentCurrency != nil {
		PresentmentCurrency = *r.PresentmentCurrency
	}
	return
}

// The amount of the attempted authorization in the currency it will be settled in.
// This currency is the same as that of the Account the card belongs to.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetSettlementAmount() (SettlementAmount int64) {
	if r != nil && r.SettlementAmount != nil {
		SettlementAmount = *r.SettlementAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
// transaction will be settled in.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetSettlementCurrency() (SettlementCurrency string) {
	if r != nil && r.SettlementCurrency != nil {
		SettlementCurrency = *r.SettlementCurrency
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

// The city the merchant resides in.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *RealTimeDecisionsCardAuthorizationRetrieveResponse) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}
