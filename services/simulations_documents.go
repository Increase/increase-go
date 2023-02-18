package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"io"
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
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("simulations/documents"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, content, opts...)
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
