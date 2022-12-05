package real_time_decisions

type RealTimeDecisionsDigitalWalletTokenActionParams struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval *RealTimeDecisionsDigitalWalletTokenApprovalActionParams `json:"approval,omitempty"`

	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline *RealTimeDecisionsDigitalWalletTokenDeclineActionParams `json:"decline,omitempty"`
}

// If your application approves the provisioning attempt, this contains metadata
// about the digital wallet token that will be generated.
func (r *RealTimeDecisionsDigitalWalletTokenActionParams) GetApproval() (Approval RealTimeDecisionsDigitalWalletTokenApprovalActionParams) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your application declines the provisioning attempt, this contains details
// about the decline.
func (r *RealTimeDecisionsDigitalWalletTokenActionParams) GetDecline() (Decline RealTimeDecisionsDigitalWalletTokenDeclineActionParams) {
	if r != nil && r.Decline != nil {
		Decline = *r.Decline
	}
	return
}
