package requests

import (
	"fmt"
	"net/url"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateALimitParameters struct {
	// The metric for the limit.
	Metric fields.Field[CreateALimitParametersMetric] `json:"metric,required"`
	// The interval for the metric. Required if `metric` is `count` or `volume`.
	Interval fields.Field[CreateALimitParametersInterval] `json:"interval"`
	// The identifier of the Account or Account Number you wish to associate the limit
	// with.
	ModelID fields.Field[string] `json:"model_id,required"`
	// The value to test the limit against.
	Value fields.Field[int64] `json:"value,required"`
}

// MarshalJSON serializes CreateALimitParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateALimitParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateALimitParameters) String() (result string) {
	return fmt.Sprintf("&CreateALimitParameters{Metric:%s Interval:%s ModelID:%s Value:%s}", r.Metric, r.Interval, r.ModelID, r.Value)
}

type CreateALimitParametersMetric string

const (
	CreateALimitParametersMetricCount  CreateALimitParametersMetric = "count"
	CreateALimitParametersMetricVolume CreateALimitParametersMetric = "volume"
)

type CreateALimitParametersInterval string

const (
	CreateALimitParametersIntervalTransaction CreateALimitParametersInterval = "transaction"
	CreateALimitParametersIntervalDay         CreateALimitParametersInterval = "day"
	CreateALimitParametersIntervalWeek        CreateALimitParametersInterval = "week"
	CreateALimitParametersIntervalMonth       CreateALimitParametersInterval = "month"
	CreateALimitParametersIntervalYear        CreateALimitParametersInterval = "year"
	CreateALimitParametersIntervalAllTime     CreateALimitParametersInterval = "all_time"
)

type UpdateALimitParameters struct {
	// The status to update the limit with.
	Status fields.Field[UpdateALimitParametersStatus] `json:"status,required"`
}

// MarshalJSON serializes UpdateALimitParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateALimitParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateALimitParameters) String() (result string) {
	return fmt.Sprintf("&UpdateALimitParameters{Status:%s}", r.Status)
}

type UpdateALimitParametersStatus string

const (
	UpdateALimitParametersStatusInactive UpdateALimitParametersStatus = "inactive"
	UpdateALimitParametersStatusActive   UpdateALimitParametersStatus = "active"
)

type LimitListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// The model to retrieve limits for.
	ModelID fields.Field[string] `query:"model_id"`
	// The status to retrieve limits for.
	Status fields.Field[string] `query:"status"`
}

// URLQuery serializes LimitListParams into a url.Values of the query parameters
// associated with this value
func (r *LimitListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r LimitListParams) String() (result string) {
	return fmt.Sprintf("&LimitListParams{Cursor:%s Limit:%s ModelID:%s Status:%s}", r.Cursor, r.Limit, r.ModelID, r.Status)
}
