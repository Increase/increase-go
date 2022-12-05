package external_accounts

type ExternalAccountsUpdateParams struct {
	// The description you choose to give the external account.
	Description *string `json:"description,omitempty"`
}

// The description you choose to give the external account.
func (r *ExternalAccountsUpdateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}
