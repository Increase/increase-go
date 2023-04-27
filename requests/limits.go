package requests

import (
	"net/url"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type LimitNewParams struct {
	// The metric for the limit.
	Metric field.Field[LimitNewParamsMetric] `json:"metric,required"`
	// The interval for the metric. Required if `metric` is `count` or `volume`.
	Interval field.Field[LimitNewParamsInterval] `json:"interval"`
	// The identifier of the Account or Account Number you wish to associate the limit
	// with.
	ModelID field.Field[string] `json:"model_id,required"`
	// The value to test the limit against.
	Value field.Field[int64] `json:"value,required"`
}

// MarshalJSON serializes LimitNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r LimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LimitNewParamsMetric string

const (
	LimitNewParamsMetricCount  LimitNewParamsMetric = "count"
	LimitNewParamsMetricVolume LimitNewParamsMetric = "volume"
)

type LimitNewParamsInterval string

const (
	LimitNewParamsIntervalTransaction LimitNewParamsInterval = "transaction"
	LimitNewParamsIntervalDay         LimitNewParamsInterval = "day"
	LimitNewParamsIntervalWeek        LimitNewParamsInterval = "week"
	LimitNewParamsIntervalMonth       LimitNewParamsInterval = "month"
	LimitNewParamsIntervalYear        LimitNewParamsInterval = "year"
	LimitNewParamsIntervalAllTime     LimitNewParamsInterval = "all_time"
)

type LimitUpdateParams struct {
	// The status to update the limit with.
	Status field.Field[LimitUpdateParamsStatus] `json:"status,required"`
}

// MarshalJSON serializes LimitUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r LimitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LimitUpdateParamsStatus string

const (
	LimitUpdateParamsStatusInactive LimitUpdateParamsStatus = "inactive"
	LimitUpdateParamsStatusActive   LimitUpdateParamsStatus = "active"
)

type LimitListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// The model to retrieve limits for.
	ModelID field.Field[string] `query:"model_id"`
	// The status to retrieve limits for.
	Status field.Field[string] `query:"status"`
}

// URLQuery serializes LimitListParams into a url.Values of the query parameters
// associated with this value
func (r LimitListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
