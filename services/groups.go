package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
)

type GroupService struct {
	Options []options.RequestOption
}

func NewGroupService(opts ...options.RequestOption) (r *GroupService) {
	r = &GroupService{}
	r.Options = opts
	return
}

// Returns details for the currently authenticated Group.
func (r *GroupService) GetDetails(ctx context.Context, opts ...options.RequestOption) (res *types.Group, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("groups/current"))
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
