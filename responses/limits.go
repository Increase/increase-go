package responses

import (
	apijson "github.com/increase/increase-go/internal/json"
)

type Limit struct {
	// The Limit identifier.
	ID string `json:"id,required"`
	// The interval for the metric. This is required if `metric` is `count` or
	// `volume`.
	Interval LimitInterval `json:"interval,required,nullable"`
	// The metric for the Limit.
	Metric LimitMetric `json:"metric,required"`
	// The identifier of the Account Number, Account, or Card the Limit applies to.
	ModelID string `json:"model_id,required"`
	// The type of the model you wish to associate the Limit with.
	ModelType LimitModelType `json:"model_type,required"`
	// The current status of the Limit.
	Status LimitStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `limit`.
	Type LimitType `json:"type,required"`
	// The value to evaluate the Limit against.
	Value int64 `json:"value,required"`
	JSON  LimitJSON
}

type LimitJSON struct {
	ID        apijson.Metadata
	Interval  apijson.Metadata
	Metric    apijson.Metadata
	ModelID   apijson.Metadata
	ModelType apijson.Metadata
	Status    apijson.Metadata
	Type      apijson.Metadata
	Value     apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Limit using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Limit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LimitInterval string

const (
	LimitIntervalTransaction LimitInterval = "transaction"
	LimitIntervalDay         LimitInterval = "day"
	LimitIntervalWeek        LimitInterval = "week"
	LimitIntervalMonth       LimitInterval = "month"
	LimitIntervalYear        LimitInterval = "year"
	LimitIntervalAllTime     LimitInterval = "all_time"
)

type LimitMetric string

const (
	LimitMetricCount  LimitMetric = "count"
	LimitMetricVolume LimitMetric = "volume"
)

type LimitModelType string

const (
	LimitModelTypeAccount       LimitModelType = "account"
	LimitModelTypeAccountNumber LimitModelType = "account_number"
	LimitModelTypeCard          LimitModelType = "card"
)

type LimitStatus string

const (
	LimitStatusActive   LimitStatus = "active"
	LimitStatusInactive LimitStatus = "inactive"
)

type LimitType string

const (
	LimitTypeLimit LimitType = "limit"
)

type LimitListResponse struct {
	// The contents of the list.
	Data []Limit `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       LimitListResponseJSON
}

type LimitListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LimitListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
