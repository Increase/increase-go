package check_deposits

import "increase/simulations"

type CheckDepositsListResponse struct {
	// The contents of the list.
	Data *[]simulations.CheckDeposit `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *CheckDepositsListResponse) GetData() (Data []simulations.CheckDeposit) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CheckDepositsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
