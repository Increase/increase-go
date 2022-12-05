package ach_transfers

type ACHTransfersListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter ACH Transfers to those that originated from the specified Account.
	AccountId *string `json:"account_id,omitempty"`

	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountId *string                         `json:"external_account_id,omitempty"`
	CreatedAt         *ACHTransfersCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *ACHTransfersListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ACHTransfersListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter ACH Transfers to those that originated from the specified Account.
func (r *ACHTransfersListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// Filter ACH Transfers to those made to the specified External Account.
func (r *ACHTransfersListQuery) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

func (r *ACHTransfersListQuery) GetCreatedAt() (CreatedAt ACHTransfersCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
