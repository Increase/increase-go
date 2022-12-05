package main

import (
	"context"
	"fmt"
	"increase/core"
)

type RequestOpts = core.RequestOpts

func main() {
	// client := core.CoreClient{
	// 	BaseURL: "https://google.com",
	// }

	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(&RequestOpts{
			ExtraBody: map[string]interface{}{
				"goodbye": "world",
			},
			Headers: map[string]string{
				"x-foo": "bar",
			},
			Context: context.TODO(),
		}),
		Body: map[string]interface{}{
			"hello": "world",
		},
		Method: "post",
		Path:   "/",
	}

	req.InitializeHTTPRequest()

	req.Request.Header.Add("X-Bar", "Foo")

	bytes, err := req.MarshalBody()
	fmt.Printf("%v %s %v\n", req.Params.Headers, string(bytes), err)

	req.FillRequestHeaders()
	fmt.Printf("%v\n", req.Request.Header)
}
