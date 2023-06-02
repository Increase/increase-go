package increase_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestFileNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.New(context.TODO(), increase.FileNewParams{
		File:        increase.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Purpose:     increase.F(increase.FileNewParamsPurposeCheckImageFront),
		Description: increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.Get(
		context.TODO(),
		"file_makxrc67oh9l6sg7w9yc",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.List(context.TODO(), increase.FileListParams{
		CreatedAt: increase.F(increase.FileListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}),
		Cursor:    increase.F("string"),
		Limit:     increase.F(int64(0)),
		Purpose:   increase.F(increase.FileListParamsPurpose{In: increase.F([]increase.FileListParamsPurposeIn{increase.FileListParamsPurposeInCheckImageFront, increase.FileListParamsPurposeInCheckImageFront, increase.FileListParamsPurposeInCheckImageFront})}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
