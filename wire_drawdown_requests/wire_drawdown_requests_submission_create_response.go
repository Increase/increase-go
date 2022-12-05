package wire_drawdown_requests

type WireDrawdownRequestsSubmissionCreateResponse struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data,omitempty"`
}

// The input message accountability data (IMAD) uniquely identifying the submission
// with Fedwire.
func (r *WireDrawdownRequestsSubmissionCreateResponse) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}
