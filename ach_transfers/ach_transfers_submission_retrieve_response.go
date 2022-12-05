package ach_transfers

type ACHTransfersSubmissionRetrieveResponse struct {
	// The trace number for the submission.
	TraceNumber *string `json:"trace_number,omitempty"`
}

// The trace number for the submission.
func (r *ACHTransfersSubmissionRetrieveResponse) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}
