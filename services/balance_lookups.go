package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type BalanceLookupService struct {
	Options []option.RequestOption
}

func NewBalanceLookupService(opts ...option.RequestOption) (r *BalanceLookupService) {
	r = &BalanceLookupService{}
	r.Options = opts
	return
}

// Look up the balance for an Account
func (r *BalanceLookupService) Lookup(ctx context.Context, body *requests.BalanceLookupLookupParams, opts ...option.RequestOption) (res *responses.BalanceLookupLookupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "balance_lookups"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
