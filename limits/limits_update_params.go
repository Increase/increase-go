package limits

type LimitsUpdateParams struct {
	// The status to update the limit with.
	Status *string `json:"status,omitempty"`
}

// The status to update the limit with.
func (r *LimitsUpdateParams) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
