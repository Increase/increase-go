package check_transfers

type CheckTransfersSubmissionStopPaymentResponse struct {
	// The identitying number of the check.
	CheckNumber *string `json:"check_number,omitempty"`
}

// The identitying number of the check.
func (r *CheckTransfersSubmissionStopPaymentResponse) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}
