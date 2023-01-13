package core

import (
	"context"
	"net/http"
)

type Agent interface {
	Do(req *http.Request) (*http.Response, error)
}

type Requester interface {
	Get(ctx context.Context, path string, req *CoreRequest, res interface{}) error
	Post(ctx context.Context, path string, req *CoreRequest, res interface{}) error
	Patch(ctx context.Context, path string, req *CoreRequest, res interface{}) error
	Put(ctx context.Context, path string, req *CoreRequest, res interface{}) error
	Delete(ctx context.Context, path string, req *CoreRequest, res interface{}) error
	Request(ctx context.Context, path string, req *CoreRequest, res interface{}) error
}
