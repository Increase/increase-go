package events

type EventsCategoryListQuery struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In *[]string `json:"in,omitempty"`
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *EventsCategoryListQuery) GetIn() (In []string) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}
