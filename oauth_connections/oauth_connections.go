package oauth_connections

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type OauthConnectionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewOauthConnectionService(requester core.Requester) (r *OauthConnectionService) {
	r = &OauthConnectionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type OauthConnection struct {
	// The OAuth Connection's identifier.
	ID string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt string `json:"created_at"`
	// The identifier of the Group that has authorized your OAuth application.
	GroupID string `json:"group_id"`
	// Whether the connection is active.
	Status OauthConnectionStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type OauthConnectionType `json:"type"`
}

type OauthConnectionStatus string

const (
	OauthConnectionStatusActive   OauthConnectionStatus = "active"
	OauthConnectionStatusInactive OauthConnectionStatus = "inactive"
)

type OauthConnectionType string

const (
	OauthConnectionTypeOauthConnection OauthConnectionType = "oauth_connection"
)

type ListOauthConnectionsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
}

//
type OauthConnectionList struct {
	// The contents of the list.
	Data []OauthConnection `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *OauthConnectionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *OauthConnectionService) Retrieve(ctx context.Context, oauth_connection_id string, opts ...*core.RequestOpts) (res *OauthConnection, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/oauth_connections/%s", oauth_connection_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type OauthConnectionsPage struct {
	*pagination.Page[OauthConnection]
}

func (r *OauthConnectionsPage) OauthConnection() *OauthConnection {
	return r.Current()
}

func (r *OauthConnectionsPage) GetNextPage() (*OauthConnectionsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &OauthConnectionsPage{page}, nil
	}
}

func (r *OauthConnectionService) List(ctx context.Context, query *ListOauthConnectionsQuery, opts ...*core.RequestOpts) (res *OauthConnectionsPage, err error) {
	page := &OauthConnectionsPage{
		Page: &pagination.Page[OauthConnection]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/oauth_connections",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
