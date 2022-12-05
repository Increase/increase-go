package simulations

type AccountStatementsCreateParams struct {
	// The identifier of the Account the statement is for.
	AccountId *string `json:"account_id,omitempty"`
}

// The identifier of the Account the statement is for.
func (r *AccountStatementsCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}
