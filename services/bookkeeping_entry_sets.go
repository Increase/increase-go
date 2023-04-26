package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type BookkeepingEntrySetService struct {
	Options []option.RequestOption
}

func NewBookkeepingEntrySetService(opts ...option.RequestOption) (r *BookkeepingEntrySetService) {
	r = &BookkeepingEntrySetService{}
	r.Options = opts
	return
}

// Create a Bookkeeping Entry Set
func (r *BookkeepingEntrySetService) New(ctx context.Context, body *requests.BookkeepingEntrySetNewParams, opts ...option.RequestOption) (res *responses.BookkeepingEntrySet, err error) {
	opts = append(r.Options[:], opts...)
	path := "bookkeeping_entry_sets"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
