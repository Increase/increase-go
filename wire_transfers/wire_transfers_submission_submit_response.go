package wire_transfers

type WireTransfersSubmissionSubmitResponse struct {
	// The accountability data for the submission.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data,omitempty"`
}

// The accountability data for the submission.
func (r *WireTransfersSubmissionSubmitResponse) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}
