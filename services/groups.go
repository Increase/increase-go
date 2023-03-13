package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
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
	path := "groups/current"
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
