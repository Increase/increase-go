package limits

type Limit struct {
	// The Limit identifier.
	Id *string `json:"id,omitempty"`

	// The interval for the metric. This is required if `metric` is `count` or
	// `volume`.
	Interval *string `json:"interval,omitempty"`

	// The metric for the Limit.
	Metric *string `json:"metric,omitempty"`

	// The identifier of the Account Number, Account, or Card the Limit applies to.
	ModelId *string `json:"model_id,omitempty"`

	// The type of the model you wish to associate the Limit with.
	ModelType *string `json:"model_type,omitempty"`

	// The current status of the Limit.
	Status *string `json:"status,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `limit`.
	Type *string `json:"type,omitempty"`

	// The value to evaluate the Limit against.
	Value *int64 `json:"value,omitempty"`
}

// The Limit identifier.
func (r *Limit) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The interval for the metric. This is required if `metric` is `count` or
// `volume`.
func (r *Limit) GetInterval() (Interval string) {
	if r != nil && r.Interval != nil {
		Interval = *r.Interval
	}
	return
}

// The metric for the Limit.
func (r *Limit) GetMetric() (Metric string) {
	if r != nil && r.Metric != nil {
		Metric = *r.Metric
	}
	return
}

// The identifier of the Account Number, Account, or Card the Limit applies to.
func (r *Limit) GetModelId() (ModelId string) {
	if r != nil && r.ModelId != nil {
		ModelId = *r.ModelId
	}
	return
}

// The type of the model you wish to associate the Limit with.
func (r *Limit) GetModelType() (ModelType string) {
	if r != nil && r.ModelType != nil {
		ModelType = *r.ModelType
	}
	return
}

// The current status of the Limit.
func (r *Limit) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `limit`.
func (r *Limit) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// The value to evaluate the Limit against.
func (r *Limit) GetValue() (Value int64) {
	if r != nil && r.Value != nil {
		Value = *r.Value
	}
	return
}
