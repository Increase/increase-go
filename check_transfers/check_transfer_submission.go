package check_transfers

type CheckTransferSubmission struct {
	// The identitying number of the check.
	CheckNumber *string `json:"check_number,omitempty"`
}

// The identitying number of the check.
func (r *CheckTransferSubmission) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}
