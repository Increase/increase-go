package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type LimitService struct {
	Options []option.RequestOption
}

func NewLimitService(opts ...option.RequestOption) (r *LimitService) {
	r = &LimitService{}
	r.Options = opts
	return
}

// Create a Limit
func (r *LimitService) New(ctx context.Context, body LimitNewParams, opts ...option.RequestOption) (res *Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := "limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Limit
func (r *LimitService) Get(ctx context.Context, limit_id string, opts ...option.RequestOption) (res *Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("limits/%s", limit_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Limit
func (r *LimitService) Update(ctx context.Context, limit_id string, body LimitUpdateParams, opts ...option.RequestOption) (res *Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("limits/%s", limit_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Limits
func (r *LimitService) List(ctx context.Context, query LimitListParams, opts ...option.RequestOption) (res *shared.Page[Limit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "limits"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Limits
func (r *LimitService) ListAutoPaging(ctx context.Context, query LimitListParams, opts ...option.RequestOption) *shared.PageAutoPager[Limit] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

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
	return apiquery.Marshal(r)
}

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
