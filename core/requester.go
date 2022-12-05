package core

import "net/http"

type Agent interface {
	Do(req *http.Request) (*http.Response, error)
}

type Requester interface {
	Get(path string, req *CoreRequest, res interface{}) error
	Post(path string, req *CoreRequest, res interface{}) error
	Patch(path string, req *CoreRequest, res interface{}) error
	Put(path string, req *CoreRequest, res interface{}) error
	Delete(path string, req *CoreRequest, res interface{}) error
	Request(path string, req *CoreRequest, res interface{}) error
}
