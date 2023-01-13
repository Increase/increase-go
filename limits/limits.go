package limits

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type LimitService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewLimitService(requester core.Requester) (r *LimitService) {
	r = &LimitService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedLimitService struct {
	Limits *LimitService
}

func (r *PreloadedLimitService) Init(service *LimitService) {
	r.Limits = service
}

func NewPreloadedLimitService(service *LimitService) (r *PreloadedLimitService) {
	r = &PreloadedLimitService{}
	r.Init(service)
	return
}

//
type Limit struct {
	// The Limit identifier.
	Id *string `json:"id"`
	// The interval for the metric. This is required if `metric` is `count` or
	// `volume`.
	Interval *LimitInterval `json:"interval"`
	// The metric for the Limit.
	Metric *LimitMetric `json:"metric"`
	// The identifier of the Account Number, Account, or Card the Limit applies to.
	ModelId *string `json:"model_id"`
	// The type of the model you wish to associate the Limit with.
	ModelType *LimitModelType `json:"model_type"`
	// The current status of the Limit.
	Status *LimitStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `limit`.
	Type *LimitType `json:"type"`
	// The value to evaluate the Limit against.
	Value *int `json:"value"`
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
func (r *Limit) GetModelId() (ModelId string) {
	if r != nil && r.ModelId != nil {
		ModelId = *r.ModelId
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
func (r *Limit) GetValue() (Value int) {
	if r != nil && r.Value != nil {
		Value = *r.Value
	}
	return
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
	Metric *CreateALimitParametersMetric `json:"metric"`
	// The interval for the metric. Required if `metric` is `count` or `volume`.
	Interval *CreateALimitParametersInterval `json:"interval,omitempty"`
	// The identifier of the Account or Account Number you wish to associate the limit
	// with.
	ModelId *string `json:"model_id"`
	// The value to test the limit against.
	Value *int `json:"value"`
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
func (r *CreateALimitParameters) GetModelId() (ModelId string) {
	if r != nil && r.ModelId != nil {
		ModelId = *r.ModelId
	}
	return
}

// The value to test the limit against.
func (r *CreateALimitParameters) GetValue() (Value int) {
	if r != nil && r.Value != nil {
		Value = *r.Value
	}
	return
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
	Status *UpdateALimitParametersStatus `json:"status"`
}

// The status to update the limit with.
func (r *UpdateALimitParameters) GetStatus() (Status UpdateALimitParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type UpdateALimitParametersStatus string

const (
	UpdateALimitParametersStatusInactive UpdateALimitParametersStatus = "inactive"
	UpdateALimitParametersStatusActive   UpdateALimitParametersStatus = "active"
)

type ListLimitsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// The model to retrieve limits for.
	ModelId *string `query:"model_id"`
	// The status to retrieve limits for.
	Status *string `query:"status"`
}

// Return the page of entries after this one.
func (r *ListLimitsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListLimitsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// The model to retrieve limits for.
func (r *ListLimitsQuery) GetModelId() (ModelId string) {
	if r != nil && r.ModelId != nil {
		ModelId = *r.ModelId
	}
	return
}

// The status to retrieve limits for.
func (r *ListLimitsQuery) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

//
type LimitList struct {
	// The contents of the list.
	Data *[]Limit `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
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

func (r *LimitService) Create(ctx context.Context, body *CreateALimitParameters, opts ...*core.RequestOpts) (res *Limit, err error) {
	err = r.post(
		ctx,
		"/limits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedLimitService) Create(ctx context.Context, body *CreateALimitParameters, opts ...*core.RequestOpts) (res *Limit, err error) {
	err = r.Limits.post(
		ctx,
		"/limits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *LimitService) Retrieve(ctx context.Context, limit_id string, opts ...*core.RequestOpts) (res *Limit, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedLimitService) Retrieve(ctx context.Context, limit_id string, opts ...*core.RequestOpts) (res *Limit, err error) {
	err = r.Limits.get(
		ctx,
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *LimitService) Update(ctx context.Context, limit_id string, body *UpdateALimitParameters, opts ...*core.RequestOpts) (res *Limit, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedLimitService) Update(ctx context.Context, limit_id string, body *UpdateALimitParameters, opts ...*core.RequestOpts) (res *Limit, err error) {
	err = r.Limits.patch(
		ctx,
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

type LimitsPage struct {
	*pagination.Page[Limit]
}

func (r *LimitsPage) Limit() *Limit {
	return r.Current()
}

func (r *LimitsPage) GetNextPage() (*LimitsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &LimitsPage{page}, nil
	}
}

func (r *LimitService) List(ctx context.Context, query *ListLimitsQuery, opts ...*core.RequestOpts) (res *LimitsPage, err error) {
	page := &LimitsPage{
		Page: &pagination.Page[Limit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/limits",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedLimitService) List(ctx context.Context, query *ListLimitsQuery, opts ...*core.RequestOpts) (res *LimitsPage, err error) {
	page := &LimitsPage{
		Page: &pagination.Page[Limit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/limits",
			},
			Requester: r.Limits.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
