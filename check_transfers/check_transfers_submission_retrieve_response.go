package check_transfers

type CheckTransfersSubmissionRetrieveResponse struct {
	// The identitying number of the check.
	CheckNumber *string `json:"check_number,omitempty"`
}

// The identitying number of the check.
func (r *CheckTransfersSubmissionRetrieveResponse) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}
