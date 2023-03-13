package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
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
	path := fmt.Sprintf("real_time_decisions/%s", real_time_decision_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, real_time_decision_id string, body *types.ActionARealTimeDecisionParameters, opts ...options.RequestOption) (res *types.RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s/action", real_time_decision_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
