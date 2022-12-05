package limits

type LimitsCreateParams struct {
	// The metric for the limit.
	Metric *string `json:"metric,omitempty"`

	// The interval for the metric. Required if `metric` is `count` or `volume`.
	Interval *string `json:"interval,omitempty"`

	// The identifier of the Account or Account Number you wish to associate the limit
	// with.
	ModelId *string `json:"model_id,omitempty"`

	// The value to test the limit against.
	Value *int64 `json:"value,omitempty"`
}

// The metric for the limit.
func (r *LimitsCreateParams) GetMetric() (Metric string) {
	if r != nil && r.Metric != nil {
		Metric = *r.Metric
	}
	return
}

// The interval for the metric. Required if `metric` is `count` or `volume`.
func (r *LimitsCreateParams) GetInterval() (Interval string) {
	if r != nil && r.Interval != nil {
		Interval = *r.Interval
	}
	return
}

// The identifier of the Account or Account Number you wish to associate the limit
// with.
func (r *LimitsCreateParams) GetModelId() (ModelId string) {
	if r != nil && r.ModelId != nil {
		ModelId = *r.ModelId
	}
	return
}

// The value to test the limit against.
func (r *LimitsCreateParams) GetValue() (Value int64) {
	if r != nil && r.Value != nil {
		Value = *r.Value
	}
	return
}
