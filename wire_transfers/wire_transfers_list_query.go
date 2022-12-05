package wire_transfers

type WireTransfersListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Wire Transfers to those belonging to the specified Account.
	AccountId *string `json:"account_id,omitempty"`

	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountId *string                          `json:"external_account_id,omitempty"`
	CreatedAt         *WireTransfersCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *WireTransfersListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *WireTransfersListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Wire Transfers to those belonging to the specified Account.
func (r *WireTransfersListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// Filter Wire Transfers to those made to the specified External Account.
func (r *WireTransfersListQuery) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

func (r *WireTransfersListQuery) GetCreatedAt() (CreatedAt WireTransfersCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
