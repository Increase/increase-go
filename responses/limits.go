package responses

import (
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
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
	ID        pjson.Metadata
	Interval  pjson.Metadata
	Metric    pjson.Metadata
	ModelID   pjson.Metadata
	ModelType pjson.Metadata
	Status    pjson.Metadata
	Type      pjson.Metadata
	Value     pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Limit using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Limit) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type LimitList struct {
	// The contents of the list.
	Data []Limit `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       LimitListJSON
}

type LimitListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LimitList using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LimitList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LimitsPage struct {
	*pagination.Page[Limit]
}

func (r *LimitsPage) Limit() *Limit {
	return r.Current()
}

func (r *LimitsPage) NextPage() (*LimitsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &LimitsPage{page}, nil
	}
}
