package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsDocumentService struct {
	Options []option.RequestOption
}

func NewSimulationsDocumentService(opts ...option.RequestOption) (r *SimulationsDocumentService) {
	r = &SimulationsDocumentService{}
	r.Options = opts
	return
}

// Simulates an tax document being created for an account.
func (r *SimulationsDocumentService) New(ctx context.Context, body *requests.DocumentNewParams, opts ...option.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/documents"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
