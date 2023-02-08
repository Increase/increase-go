package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
)

type SimulationsDocumentService struct {
	Options []options.RequestOption
}

func NewSimulationsDocumentService(opts ...options.RequestOption) (r *SimulationsDocumentService) {
	r = &SimulationsDocumentService{}
	r.Options = opts
	return
}

// Simulates an tax document being created for an account.
func (r *SimulationsDocumentService) New(ctx context.Context, body *types.SimulateATaxDocumentBeingCreatedParameters, opts ...options.RequestOption) (res *types.Document, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/documents"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}
