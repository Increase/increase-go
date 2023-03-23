package requests

import (
	"fmt"
	"net/url"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type Limit struct {
	// The Limit identifier.
	ID fields.Field[string] `json:"id,required"`
	// The interval for the metric. This is required if `metric` is `count` or
	// `volume`.
	Interval fields.Field[LimitInterval] `json:"interval,required,nullable"`
	// The metric for the Limit.
	Metric fields.Field[LimitMetric] `json:"metric,required"`
	// The identifier of the Account Number, Account, or Card the Limit applies to.
	ModelID fields.Field[string] `json:"model_id,required"`
	// The type of the model you wish to associate the Limit with.
	ModelType fields.Field[LimitModelType] `json:"model_type,required"`
	// The current status of the Limit.
	Status fields.Field[LimitStatus] `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `limit`.
	Type fields.Field[LimitType] `json:"type,required"`
	// The value to evaluate the Limit against.
	Value fields.Field[int64] `json:"value,required"`
}

// MarshalJSON serializes Limit into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Limit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Limit) String() (result string) {
	return fmt.Sprintf("&Limit{ID:%s Interval:%s Metric:%s ModelID:%s ModelType:%s Status:%s Type:%s Value:%s}", r.ID, r.Interval, r.Metric, r.ModelID, r.ModelType, r.Status, r.Type, r.Value)
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
