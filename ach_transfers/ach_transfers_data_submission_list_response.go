package ach_transfers

type ACHTransfersDataSubmissionListResponse struct {
	// The trace number for the submission.
	TraceNumber *string `json:"trace_number,omitempty"`
}

// The trace number for the submission.
func (r *ACHTransfersDataSubmissionListResponse) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}
