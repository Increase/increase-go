package services

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestFilesNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.New(context.TODO(), &requests.CreateAFileParameters{File: fields.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))), Description: fields.F("x"), Purpose: fields.F(requests.CreateAFileParametersPurposeCheckImageFront)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFilesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.Get(
		context.TODO(),
		"file_makxrc67oh9l6sg7w9yc",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFilesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.List(context.TODO(), &requests.FileListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0)), CreatedAt: fields.F(requests.FileListParamsCreatedAt{After: fields.F(time.Now()), Before: fields.F(time.Now()), OnOrAfter: fields.F(time.Now()), OnOrBefore: fields.F(time.Now())}), Purpose: fields.F(requests.FileListParamsPurpose{In: fields.F([]requests.FileListParamsPurposeIn{requests.FileListParamsPurposeInCheckImageFront, requests.FileListParamsPurposeInCheckImageFront, requests.FileListParamsPurposeInCheckImageFront})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
