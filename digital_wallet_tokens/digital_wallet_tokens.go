package digital_wallet_tokens

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type DigitalWalletTokenService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewDigitalWalletTokenService(requester core.Requester) (r *DigitalWalletTokenService) {
	r = &DigitalWalletTokenService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	ID string `json:"id"`
	// The identifier for the Card this Digital Wallet Token belongs to.
	CardID string `json:"card_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt string `json:"created_at"`
	// This indicates if payments can be made with the Digital Wallet Token.
	Status DigitalWalletTokenStatus `json:"status"`
	// The digital wallet app being used.
	TokenRequestor DigitalWalletTokenTokenRequestor `json:"token_requestor"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type DigitalWalletTokenType `json:"type"`
}

type DigitalWalletTokenStatus string

const (
	DigitalWalletTokenStatusActive      DigitalWalletTokenStatus = "active"
	DigitalWalletTokenStatusInactive    DigitalWalletTokenStatus = "inactive"
	DigitalWalletTokenStatusSuspended   DigitalWalletTokenStatus = "suspended"
	DigitalWalletTokenStatusDeactivated DigitalWalletTokenStatus = "deactivated"
)

type DigitalWalletTokenTokenRequestor string

const (
	DigitalWalletTokenTokenRequestorApplePay  DigitalWalletTokenTokenRequestor = "apple_pay"
	DigitalWalletTokenTokenRequestorGooglePay DigitalWalletTokenTokenRequestor = "google_pay"
)

type DigitalWalletTokenType string

const (
	DigitalWalletTokenTypeDigitalWalletToken DigitalWalletTokenType = "digital_wallet_token"
)

type ListDigitalWalletTokensQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardID    string                                `query:"card_id"`
	CreatedAt ListDigitalWalletTokensQueryCreatedAt `query:"created_at"`
}

type ListDigitalWalletTokensQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore string `json:"on_or_before,omitempty"`
}

//
type DigitalWalletTokenList struct {
	// The contents of the list.
	Data []DigitalWalletToken `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *DigitalWalletTokenList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *DigitalWalletTokenService) Retrieve(ctx context.Context, digital_wallet_token_id string, opts ...*core.RequestOpts) (res *DigitalWalletToken, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/digital_wallet_tokens/%s", digital_wallet_token_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type DigitalWalletTokensPage struct {
	*pagination.Page[DigitalWalletToken]
}

func (r *DigitalWalletTokensPage) DigitalWalletToken() *DigitalWalletToken {
	return r.Current()
}

func (r *DigitalWalletTokensPage) GetNextPage() (*DigitalWalletTokensPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &DigitalWalletTokensPage{page}, nil
	}
}

func (r *DigitalWalletTokenService) List(ctx context.Context, query *ListDigitalWalletTokensQuery, opts ...*core.RequestOpts) (res *DigitalWalletTokensPage, err error) {
	page := &DigitalWalletTokensPage{
		Page: &pagination.Page[DigitalWalletToken]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/digital_wallet_tokens",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
