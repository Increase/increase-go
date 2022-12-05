package check_transfers

type CheckTransfersSubmissionCreateResponse struct {
	// The identitying number of the check.
	CheckNumber *string `json:"check_number,omitempty"`
}

// The identitying number of the check.
func (r *CheckTransfersSubmissionCreateResponse) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}
