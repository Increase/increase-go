package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/core/query"
	"increase/pagination"
	"net/url"
)

type Limit struct {
	// The Limit identifier.
	ID *string `pjson:"id"`
	// The interval for the metric. This is required if `metric` is `count` or
	// `volume`.
	Interval *LimitInterval `pjson:"interval"`
	// The metric for the Limit.
	Metric *LimitMetric `pjson:"metric"`
	// The identifier of the Account Number, Account, or Card the Limit applies to.
	ModelID *string `pjson:"model_id"`
	// The type of the model you wish to associate the Limit with.
	ModelType *LimitModelType `pjson:"model_type"`
	// The current status of the Limit.
	Status *LimitStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `limit`.
	Type *LimitType `pjson:"type"`
	// The value to evaluate the Limit against.
	Value      *int64                 `pjson:"value"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Limit using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Limit) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Limit into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Limit) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Limit identifier.
func (r *Limit) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The interval for the metric. This is required if `metric` is `count` or
// `volume`.
func (r *Limit) GetInterval() (Interval LimitInterval) {
	if r != nil && r.Interval != nil {
		Interval = *r.Interval
	}
	return
}

// The metric for the Limit.
func (r *Limit) GetMetric() (Metric LimitMetric) {
	if r != nil && r.Metric != nil {
		Metric = *r.Metric
	}
	return
}

// The identifier of the Account Number, Account, or Card the Limit applies to.
func (r *Limit) GetModelID() (ModelID string) {
	if r != nil && r.ModelID != nil {
		ModelID = *r.ModelID
	}
	return
}

// The type of the model you wish to associate the Limit with.
func (r *Limit) GetModelType() (ModelType LimitModelType) {
	if r != nil && r.ModelType != nil {
		ModelType = *r.ModelType
	}
	return
}

// The current status of the Limit.
func (r *Limit) GetStatus() (Status LimitStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `limit`.
func (r *Limit) GetType() (Type LimitType) {
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

func (r Limit) String() (result string) {
	return fmt.Sprintf("&Limit{ID:%s Interval:%s Metric:%s ModelID:%s ModelType:%s Status:%s Type:%s Value:%s}", core.FmtP(r.ID), core.FmtP(r.Interval), core.FmtP(r.Metric), core.FmtP(r.ModelID), core.FmtP(r.ModelType), core.FmtP(r.Status), core.FmtP(r.Type), core.FmtP(r.Value))
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
	Metric *CreateALimitParametersMetric `pjson:"metric"`
	// The interval for the metric. Required if `metric` is `count` or `volume`.
	Interval *CreateALimitParametersInterval `pjson:"interval"`
	// The identifier of the Account or Account Number you wish to associate the limit
	// with.
	ModelID *string `pjson:"model_id"`
	// The value to test the limit against.
	Value      *int64                 `pjson:"value"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateALimitParameters using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CreateALimitParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateALimitParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateALimitParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The metric for the limit.
func (r *CreateALimitParameters) GetMetric() (Metric CreateALimitParametersMetric) {
	if r != nil && r.Metric != nil {
		Metric = *r.Metric
	}
	return
}

// The interval for the metric. Required if `metric` is `count` or `volume`.
func (r *CreateALimitParameters) GetInterval() (Interval CreateALimitParametersInterval) {
	if r != nil && r.Interval != nil {
		Interval = *r.Interval
	}
	return
}

// The identifier of the Account or Account Number you wish to associate the limit
// with.
func (r *CreateALimitParameters) GetModelID() (ModelID string) {
	if r != nil && r.ModelID != nil {
		ModelID = *r.ModelID
	}
	return
}

// The value to test the limit against.
func (r *CreateALimitParameters) GetValue() (Value int64) {
	if r != nil && r.Value != nil {
		Value = *r.Value
	}
	return
}

func (r CreateALimitParameters) String() (result string) {
	return fmt.Sprintf("&CreateALimitParameters{Metric:%s Interval:%s ModelID:%s Value:%s}", core.FmtP(r.Metric), core.FmtP(r.Interval), core.FmtP(r.ModelID), core.FmtP(r.Value))
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
	Status     *UpdateALimitParametersStatus `pjson:"status"`
	jsonFields map[string]interface{}        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into UpdateALimitParameters using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *UpdateALimitParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateALimitParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateALimitParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The status to update the limit with.
func (r *UpdateALimitParameters) GetStatus() (Status UpdateALimitParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r UpdateALimitParameters) String() (result string) {
	return fmt.Sprintf("&UpdateALimitParameters{Status:%s}", core.FmtP(r.Status))
}

type UpdateALimitParametersStatus string

const (
	UpdateALimitParametersStatusInactive UpdateALimitParametersStatus = "inactive"
	UpdateALimitParametersStatusActive   UpdateALimitParametersStatus = "active"
)

type LimitListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// The model to retrieve limits for.
	ModelID *string `query:"model_id"`
	// The status to retrieve limits for.
	Status     *string                `query:"status"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into LimitListParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LimitListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes LimitListParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *LimitListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes LimitListParams into a url.Values of the query parameters
// associated with this value
func (r *LimitListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *LimitListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *LimitListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// The model to retrieve limits for.
func (r *LimitListParams) GetModelID() (ModelID string) {
	if r != nil && r.ModelID != nil {
		ModelID = *r.ModelID
	}
	return
}

// The status to retrieve limits for.
func (r *LimitListParams) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r LimitListParams) String() (result string) {
	return fmt.Sprintf("&LimitListParams{Cursor:%s Limit:%s ModelID:%s Status:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.ModelID), core.FmtP(r.Status))
}

type LimitList struct {
	// The contents of the list.
	Data *[]Limit `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into LimitList using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LimitList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes LimitList into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *LimitList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes LimitList into a url.Values of the query parameters
// associated with this value
func (r *LimitList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *LimitList) GetData() (Data []Limit) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *LimitList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r LimitList) String() (result string) {
	return fmt.Sprintf("&LimitList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
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
