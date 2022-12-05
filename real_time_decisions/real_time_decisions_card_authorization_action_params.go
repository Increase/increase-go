package real_time_decisions

type RealTimeDecisionsCardAuthorizationActionParams struct {
	// Whether the card authorization should be approved or declined.
	Decision *string `json:"decision,omitempty"`
}

// Whether the card authorization should be approved or declined.
func (r *RealTimeDecisionsCardAuthorizationActionParams) GetDecision() (Decision string) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
}
