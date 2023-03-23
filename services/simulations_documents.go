package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *SimulationsDocumentService) New(ctx context.Context, body *requests.SimulateATaxDocumentBeingCreatedParameters, opts ...options.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/documents"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
