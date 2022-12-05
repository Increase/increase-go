package real_time_decisions

type RealTimeDecisionsDigitalWalletTokenDeclineActionParams struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason *string `json:"reason,omitempty"`
}

// Why the tokenization attempt was declined. This is for logging purposes only and
// is not displayed to the end-user.
func (r *RealTimeDecisionsDigitalWalletTokenDeclineActionParams) GetReason() (Reason string) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}
