package services

import (
	"context"
	"increase/options"
	"increase/types"
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
	opts = append(r.Options[:], opts...)
	path := "simulations/documents"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
