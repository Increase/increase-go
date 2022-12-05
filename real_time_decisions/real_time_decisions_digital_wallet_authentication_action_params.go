package real_time_decisions

type RealTimeDecisionsDigitalWalletAuthenticationActionParams struct {
	// Whether your application was able to deliver the one-time passcode.
	Result *string `json:"result,omitempty"`
}

// Whether your application was able to deliver the one-time passcode.
func (r *RealTimeDecisionsDigitalWalletAuthenticationActionParams) GetResult() (Result string) {
	if r != nil && r.Result != nil {
		Result = *r.Result
	}
	return
}
