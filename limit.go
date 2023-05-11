package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// LimitService contains methods and other services that help with interacting with
// the increase API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewLimitService] method instead.
type LimitService struct {
	Options []option.RequestOption
}

// NewLimitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
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

// You can set limits at the Account, Account Number, or Card level. Limits applied
// to Accounts will apply to all Account Numbers and Cards in the Account. You can
// specify any number of Limits and they will all be applied to inbound debits and
// card authorizations. Volume and count Limits are designed to prevent
// unauthorized debits.
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
	JSON  limitJSON
}

// limitJSON contains the JSON metadata for the struct [Limit]
type limitJSON struct {
	ID          apijson.Field
	Interval    apijson.Field
	Metric      apijson.Field
	ModelID     apijson.Field
	ModelType   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

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
	Metric param.Field[LimitNewParamsMetric] `json:"metric,required"`
	// The interval for the metric. Required if `metric` is `count` or `volume`.
	Interval param.Field[LimitNewParamsInterval] `json:"interval"`
	// The identifier of the Account or Account Number you wish to associate the limit
	// with.
	ModelID param.Field[string] `json:"model_id,required"`
	// The value to test the limit against.
	Value param.Field[int64] `json:"value,required"`
}

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
	Status param.Field[LimitUpdateParamsStatus] `json:"status,required"`
}

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
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// The model to retrieve limits for.
	ModelID param.Field[string] `query:"model_id"`
	// The status to retrieve limits for.
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [LimitListParams]'s query parameters as `url.Values`.
func (r LimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of Limit objects
type LimitListResponse struct {
	// The contents of the list.
	Data []Limit `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       limitListResponseJSON
}

// limitListResponseJSON contains the JSON metadata for the struct
// [LimitListResponse]
type limitListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
