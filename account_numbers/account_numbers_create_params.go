package account_numbers

type AccountNumbersCreateParams struct {
	// The Account the Account Number should belong to.
	AccountId *string `json:"account_id,omitempty"`

	// The name you choose for the Account Number.
	Name *string `json:"name,omitempty"`
}

// The Account the Account Number should belong to.
func (r *AccountNumbersCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The name you choose for the Account Number.
func (r *AccountNumbersCreateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}
