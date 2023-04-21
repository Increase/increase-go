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
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestFileNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.New(context.TODO(), &requests.FileNewParams{File: increase.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))), Description: increase.F("x"), Purpose: increase.F(requests.FileNewParamsPurposeCheckImageFront)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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

func TestFileListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.List(context.TODO(), &requests.FileListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), CreatedAt: increase.F(requests.FileListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}), Purpose: increase.F(requests.FileListParamsPurpose{In: increase.F([]requests.FileListParamsPurposeIn{requests.FileListParamsPurposeInCheckImageFront, requests.FileListParamsPurposeInCheckImageFront, requests.FileListParamsPurposeInCheckImageFront})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
