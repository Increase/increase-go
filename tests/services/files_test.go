package services

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestFilesNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.New(context.TODO(), &types.CreateAFileParameters{File: increase.P("todo: files not implemented"), Description: increase.P("x"), Purpose: increase.P(types.CreateAFileParametersPurposeCheckImageFront)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestFilesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.Get(
		context.TODO(),
		"file_makxrc67oh9l6sg7w9yc",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestFilesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.List(context.TODO(), &types.FileListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), CreatedAt: increase.P(types.FileListParamsCreatedAt{After: increase.P(time.Now()), Before: increase.P(time.Now()), OnOrAfter: increase.P(time.Now()), OnOrBefore: increase.P(time.Now())}), Purpose: increase.P(types.FileListParamsPurpose{In: increase.P([]types.FileListParamsPurposeIn{types.FileListParamsPurposeInCheckImageFront, types.FileListParamsPurposeInCheckImageFront, types.FileListParamsPurposeInCheckImageFront})})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
