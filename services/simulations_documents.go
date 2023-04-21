package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationDocumentService struct {
	Options []option.RequestOption
}

func NewSimulationDocumentService(opts ...option.RequestOption) (r *SimulationDocumentService) {
	r = &SimulationDocumentService{}
	r.Options = opts
	return
}

// Simulates an tax document being created for an account.
func (r *SimulationDocumentService) New(ctx context.Context, body *requests.SimulationDocumentNewParams, opts ...option.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/documents"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
