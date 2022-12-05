package accounts

type AccountsUpdateParams struct {
	// The new name of the Account.
	Name *string `json:"name,omitempty"`
}

// The new name of the Account.
func (r *AccountsUpdateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}
