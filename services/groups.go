package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/responses"
)

type GroupService struct {
	Options []option.RequestOption
}

func NewGroupService(opts ...option.RequestOption) (r *GroupService) {
	r = &GroupService{}
	r.Options = opts
	return
}

// Returns details for the currently authenticated Group.
func (r *GroupService) GetDetails(ctx context.Context, opts ...option.RequestOption) (res *responses.Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "groups/current"
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
