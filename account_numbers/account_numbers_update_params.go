package account_numbers

type AccountNumbersUpdateParams struct {
	// The name you choose for the Account Number.
	Name *string `json:"name,omitempty"`

	// This indicates if transfers can be made to the Account Number.
	Status *string `json:"status,omitempty"`
}

// The name you choose for the Account Number.
func (r *AccountNumbersUpdateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// This indicates if transfers can be made to the Account Number.
func (r *AccountNumbersUpdateParams) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
