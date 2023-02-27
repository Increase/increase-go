package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestFilesNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.New(context.TODO(), &types.CreateAFileParameters{File: nil, Description: increase.P("x"), Purpose: increase.P(types.CreateAFileParametersPurposeCheckImageFront)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestFilesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.Get(context.TODO(), "file_makxrc67oh9l6sg7w9yc")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestFilesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Files.List(context.TODO(), &types.FileListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), CreatedAt: &types.FilesListParamsCreatedAt{After: increase.P("2019-12-27T18:11:19.117Z"), Before: increase.P("2019-12-27T18:11:19.117Z"), OnOrAfter: increase.P("2019-12-27T18:11:19.117Z"), OnOrBefore: increase.P("2019-12-27T18:11:19.117Z")}, Purpose: &types.FilesListParamsPurpose{In: &[]types.FilesListParamsPurposeIn{types.FilesListParamsPurposeInCheckImageFront, types.FilesListParamsPurposeInCheckImageFront, types.FilesListParamsPurposeInCheckImageFront}}})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
