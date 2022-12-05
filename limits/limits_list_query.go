package limits

type LimitsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// The model to retrieve limits for.
	ModelId *string `json:"model_id,omitempty"`

	// The status to retrieve limits for.
	Status *string `json:"status,omitempty"`
}

// Return the page of entries after this one.
func (r *LimitsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *LimitsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// The model to retrieve limits for.
func (r *LimitsListQuery) GetModelId() (ModelId string) {
	if r != nil && r.ModelId != nil {
		ModelId = *r.ModelId
	}
	return
}

// The status to retrieve limits for.
func (r *LimitsListQuery) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
