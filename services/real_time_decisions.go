package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"io"
	"net/url"
)

type RealTimeDecisionService struct {
	Options []options.RequestOption
}

func NewRealTimeDecisionService(opts ...options.RequestOption) (r *RealTimeDecisionService) {
	r = &RealTimeDecisionService{}
	r.Options = opts
	return
}

// Retrieve a Real-Time Decision
func (r *RealTimeDecisionService) Get(ctx context.Context, real_time_decision_id string, opts ...options.RequestOption) (res *types.RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("real_time_decisions/%s", real_time_decision_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, real_time_decision_id string, body *types.ActionARealTimeDecisionParameters, opts ...options.RequestOption) (res *types.RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("real_time_decisions/%s/action", real_time_decision_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}
